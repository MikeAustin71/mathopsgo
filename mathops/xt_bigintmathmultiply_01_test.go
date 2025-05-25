package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntMathMultiply_BigIntMultiply_01(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedBigINumStr := "2875.94572"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'\n Error='%v'\n\n", multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'\nError='%v'\n\n", multiplicandStr, err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'\nError='%v'\n\n", expectedBigINumStr, err.Error())
    return
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

  multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

  multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, resultPrecision, err := new(BigIntMathMultiply).\n"+
      "BigIntMultiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetIntegerValue()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(result) != 0 {
    t.Errorf("Error: Expected result='%v'.\n"+
      "Instead, result= '%s'.\n\n",
      expectedBigINumBigInt.Text(10), result.Text(10))
    return
  }

  expectedBigINumPrecision, err := expectedBigINum.GetPrecisionUint()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumPrecision, err := expectedBigINum.GetPrecisionUint()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  uintResultPrecision := uint(resultPrecision.Uint64())

  if expectedBigINumPrecision != uintResultPrecision {
    t.Errorf("Error: Expected result precision='%v'.\n"+
      "Instead, result precision='%v'\n\n",
      expectedBigINumPrecision, uintResultPrecision)
    return
  }
}

func TestBigIntMathMultiply_BigIntMultiply_02(t *testing.T) {
  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedBigINumStr := "30987680500513189125.14259702468435"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "multiplicandBiNum, err := new(BigIntNum).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
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
    t.Errorf("Error returned by:"+
      "err = iaMultiplier.Multiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      expectedBigINumStr, err.Error())
    return
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

  multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierPrecision, err := multiplierBiNum.\n"+
      "  GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

  multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandPrecision, err := multiplicandBiNum.\n"+
      "  GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, resultPrecision, err := new(BigIntMathMultiply).\n"+
      "  BigIntMultiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumInt, err := expectedBigINum.GetIntegerValue()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumInt.Cmp(result) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINumInt.Text(10), result.Text(10))
    return
  }

  expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumPrecisionBigInt, err :=\n"+
      "  expectedBigINum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumPrecisionBigInt.Cmp(resultPrecision) != 0 {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINumPrecisionBigInt.Text(10), resultPrecision.Text(10))
    return
  }
}

func TestBigIntMathMultiply_BigIntMultiply_03(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedBigINumStr := "-30987680500513189125.14259702468435"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n",
      multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).\n"+
      "  NewNumStr(multiplicandStr)\n\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
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
    t.Errorf("Error returned by:"+
      "err = iaMultiplier.Multiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      expectedBigINumStr, err.Error())
    return
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

  multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

  multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, resultPrecision, err := new(BigIntMathMultiply).\n"+
      "BigIntMultiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumInt, err := expectedBigINum.GetIntegerValue()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumInt.Cmp(result) != 0 {
    t.Errorf("Error: Expected result='%v'.\n"+
      "Instead, result= '%s'.\n\n",
      expectedBigINumInt.Text(10), result.Text(10))
    return
  }

  expectedBigINumPrecisionInt, err := expectedBigINum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumPrecisionInt, err := expectedBigINum.\n"+
      "  GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumPrecisionInt.Cmp(resultPrecision) != 0 {
    t.Errorf("Expected result precision='%v'.\n"+
      "Instead, result precision='%v'\n\n",
      expectedBigINumPrecisionInt.Text(10), resultPrecision.Text(10))

  }

  return
}

func TestBigIntMathMultiply_BigIntMultiply_04(t *testing.T) {
  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedBigINumStr := "22197234145.3632"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'\n"+
      "\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'\n"+
      "\nError='%v'\n\n", multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n", multiplicandStr, err.Error())
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
    t.Errorf("Error returned by:\n"+
      "iaMultiplier.Multiply(...)"+
      "\nError='%v'\n\n", err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\n"+
      "\nError='%v'\n\n", expectedBigINumStr, err.Error())
    return
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

  multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

  multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandPrecision, err := multiplicandBiNum.\n"+
      "  GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, resultPrecision, err := new(BigIntMathMultiply).\n"+
      "  BigIntMultiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumInt, err := expectedBigINum.\n"+
      "  GetIntegerValue()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumInt.Cmp(result) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINumInt.Text(10), result.Text(10))
    return
  }

  expectedBigINumPrecisionBigInt, err := expectedBigINum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumPrecisionBigInt, err := expectedBigINum.\n"+
      "  GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumPrecisionBigInt.Cmp(resultPrecision) != 0 {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINumPrecisionBigInt.Text(10), resultPrecision.Text(10))
  }

  return
}

func TestBigIntMathMultiply_BigIntMultiply_05(t *testing.T) {
  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0.00
  expectedBigINumStr := "0"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandBiNum, err := new(BigIntNum).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "\nError='%v'\n\n",
      multiplicandStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      multiplicandStr, err.Error())
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
    t.Errorf("Error returned by:\n"+
      "\terr = iaMultiplier.Multiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return

  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\n"+
      "Error='%v'\n\n", expectedBigINumStr, err.Error())
    return
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)

  multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierPrecision, err := multiplierBiNum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)

  multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandPrecision, err := multiplicandBiNum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, resultPrecision, err := new(BigIntMathMultiply).\n"+
      "BigIntMultiply(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINumInt, err := expectedBigINum.GetIntegerValue()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumInt, err := expectedBigINum.GetIntegerValue()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumInt.Cmp(result) != 0 {
    t.Errorf("Error: Expected result='%v'.\n"+
      "Instead, result= '%s'.\n",
      expectedBigINumInt.Text(10), result.Text(10))
    return
  }

  expectedBigINumBigIPrecision, err := expectedBigINum.GetPrecisionBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumBigIPrecision, err :=\n"+
      "  expectedBigINum.GetPrecisionBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumBigIPrecision.Cmp(resultPrecision) != 0 {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINumBigIPrecision, resultPrecision.Text(10))
  }

  return
}

func TestBigIntMathMultiply_BigIntMultiply_06(t *testing.T) {

  multiplierStr := "0.00000"

  multiplicandStr := "0.00"

  // product = 0.00
  expectedBigINumStr := "0"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
  multiplierPrecision := multiplierBiNum.GetPrecisionBigInt()
  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
  multiplicandPrecision := multiplicandBiNum.GetPrecisionBigInt()

  result, resultPrecision, err := new(BigIntMathMultiply).BigIntMultiply(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiply(...) "+
      "Error='%v'. ", err.Error())
  }

  if expectedBigINum.GetIntegerValue().Cmp(result) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINum.GetIntegerValue().Text(10), result.Text(10))
  }

  if expectedBigINum.GetPrecisionBigInt().Cmp(resultPrecision) != 0 {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINum.GetPrecisionUint(), resultPrecision.Text(10))
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_01(t *testing.T) {

  expectedNumStr := "105.6752"
  num := 1056752
  precision := uint(4)
  exponent := big.NewInt(0)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyByTenToPwr_02(t *testing.T) {

  expectedNumStr := "-105.6752"
  num := -1056752
  precision := uint(4)
  exponent := big.NewInt(0)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_03(t *testing.T) {

  expectedNumStr := "1056.752"
  num := 1056752
  precision := uint(4)
  exponent := big.NewInt(1)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_04(t *testing.T) {

  expectedNumStr := "-1056.752"
  num := -1056752
  precision := uint(4)
  exponent := big.NewInt(1)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_05(t *testing.T) {

  expectedNumStr := "10567.52"
  num := 1056752
  precision := uint(4)
  exponent := big.NewInt(2)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_06(t *testing.T) {

  expectedNumStr := "-10567.52"
  num := -1056752
  precision := uint(4)
  exponent := big.NewInt(2)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_07(t *testing.T) {

  expectedNumStr := "10567520000"
  num := 1056752
  precision := uint(4)
  exponent := big.NewInt(8)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_08(t *testing.T) {

  expectedNumStr := "0"
  num := 0
  precision := uint(4)
  exponent := big.NewInt(5)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }
}

func TestBigIntMathMultiply_MultiplyByTenToPwr_09(t *testing.T) {

  expectedNumStr := "10.56752"
  num := 1056752
  precision := uint(0)
  exponent := big.NewInt(-5)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyByTenToPwr_10(t *testing.T) {

  expectedNumStr := "0.01056752"
  num := 1056752
  precision := uint(3)
  exponent := big.NewInt(-5)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyByTenToPwr_11(t *testing.T) {

  expectedNumStr := "-0.01056752"
  num := -1056752
  precision := uint(3)
  exponent := big.NewInt(-5)

  fixDec := new(BigIntFixedDecimal).NewInt(num, precision)

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTenToPower(
      fixDec.GetInteger(), fixDec.GetPrecisionBigInt(), exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTenToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntFixedDecimal).NewBigIntPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntFixedDecimal).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_01(t *testing.T) {

  // multiplicand = 23.321
  multiplicandBInt := big.NewInt(23321)
  multiplicandPrecision := big.NewInt(3)
  exponent := uint(5)
  expectedResult := "746.272"

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
      multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != result.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, result.GetNumStr())
  }

}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_02(t *testing.T) {

  // multiplicand = 8
  multiplicandBInt := big.NewInt(8)
  multiplicandPrecision := big.NewInt(0)
  exponent := uint(10)
  expectedResult := "8192"

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
      multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != result.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, result.GetNumStr())
  }

}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_03(t *testing.T) {

  // multiplicand = 9.871234
  multiplicandBInt := big.NewInt(9871234)
  multiplicandPrecision := big.NewInt(6)
  exponent := uint(1)
  expectedResult := "19.742468"

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
      multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != result.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, result.GetNumStr())
  }

}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_04(t *testing.T) {

  // multiplicand = -9.871234
  multiplicandBInt := big.NewInt(-9871234)
  multiplicandPrecision := big.NewInt(6)
  exponent := uint(3)
  expectedResult := "-78.969872"

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
      multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != result.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, result.GetNumStr())
  }

}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_05(t *testing.T) {

  // multiplicand = 8
  multiplicandBInt := big.NewInt(8)
  multiplicandPrecision := big.NewInt(0)
  exponent := uint(0)
  expectedResult := "8"

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
      multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != result.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, result.GetNumStr())
  }

}

func TestBigIntMathMultiply_BigIntMultiplyByTwoToPower_06(t *testing.T) {

  // (0.12345 x 2^15 = 4045.2096)
  multiplicandBInt := big.NewInt(12345)
  multiplicandPrecision := big.NewInt(5)
  exponent := uint(15)
  expectedResult := "4045.2096"

  product, productPrecision, err :=
    new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
      multiplicandBInt, multiplicandPrecision, exponent)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(...) "+
      "Error='%v' ", err.Error())
  }

  result, err := new(BigIntNum).NewBigIntBigPrecision(product, productPrecision)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != result.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, result.GetNumStr())
  }

}

func TestBigIntMathMultiply_FixedDecimalMultiply_01(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedBigINumStr := "2875.94572"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  multiplier :=
    new(BigIntFixedDecimal).New(
      multiplierBiNum.GetIntegerValue(),
      multiplierBiNum.GetPrecisionUint())

  multiplicand :=
    new(BigIntFixedDecimal).New(
      multiplicandBiNum.GetIntegerValue(),
      multiplicandBiNum.GetPrecisionUint())

  result := new(BigIntMathMultiply).FixedDecimalMultiply(
    multiplier,
    multiplicand)

  if expectedBigINum.GetIntegerValue().Cmp(result.GetInteger()) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINum.GetIntegerValue().Text(10), result.GetInteger().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINum.GetPrecisionUint(), result.GetPrecision())
  }
}

func TestBigIntMathMultiply_FixedDecimalMultiply_02(t *testing.T) {
  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedBigINumStr := "30987680500513189125.14259702468435"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  multiplier :=
    new(BigIntFixedDecimal).New(
      multiplierBiNum.GetIntegerValue(),
      multiplierBiNum.GetPrecisionUint())

  multiplicand :=
    new(BigIntFixedDecimal).New(
      multiplicandBiNum.GetIntegerValue(),
      multiplicandBiNum.GetPrecisionUint())

  result := new(BigIntMathMultiply).FixedDecimalMultiply(
    multiplier,
    multiplicand)

  if expectedBigINum.GetIntegerValue().Cmp(result.GetInteger()) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINum.GetIntegerValue().Text(10), result.GetInteger().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINum.GetPrecisionUint(), result.GetPrecision())
  }
}

func TestBigIntMathMultiply_FixedDecimalMultiply_03(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedBigINumStr := "-30987680500513189125.14259702468435"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  multiplier :=
    new(BigIntFixedDecimal).New(
      multiplierBiNum.GetIntegerValue(),
      multiplierBiNum.GetPrecisionUint())

  multiplicand :=
    new(BigIntFixedDecimal).New(
      multiplicandBiNum.GetIntegerValue(),
      multiplicandBiNum.GetPrecisionUint())

  result := new(BigIntMathMultiply).FixedDecimalMultiply(
    multiplier,
    multiplicand)

  if expectedBigINum.GetIntegerValue().Cmp(result.GetInteger()) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINum.GetIntegerValue().Text(10), result.GetInteger().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINum.GetPrecisionUint(), result.GetPrecision())
  }
}

func TestBigIntMathMultiply_FixedDecimalMultiply_04(t *testing.T) {
  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedBigINumStr := "22197234145.3632"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  multiplier :=
    new(BigIntFixedDecimal).New(
      multiplierBiNum.GetIntegerValue(),
      multiplierBiNum.GetPrecisionUint())

  multiplicand :=
    new(BigIntFixedDecimal).New(
      multiplicandBiNum.GetIntegerValue(),
      multiplicandBiNum.GetPrecisionUint())

  result := new(BigIntMathMultiply).FixedDecimalMultiply(
    multiplier,
    multiplicand)

  if expectedBigINum.GetIntegerValue().Cmp(result.GetInteger()) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINum.GetIntegerValue().Text(10), result.GetInteger().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINum.GetPrecisionUint(), result.GetPrecision())
  }
}

func TestBigIntMathMultiply_FixedDecimalMultiply_05(t *testing.T) {
  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0.00
  expectedBigINumStr := "0"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  multiplier :=
    new(BigIntFixedDecimal).New(
      multiplierBiNum.GetIntegerValue(),
      multiplierBiNum.GetPrecisionUint())

  multiplicand :=
    new(BigIntFixedDecimal).New(
      multiplicandBiNum.GetIntegerValue(),
      multiplicandBiNum.GetPrecisionUint())

  result := new(BigIntMathMultiply).FixedDecimalMultiply(
    multiplier,
    multiplicand)

  if expectedBigINum.GetIntegerValue().Cmp(result.GetInteger()) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINum.GetIntegerValue().Text(10), result.GetInteger().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINum.GetPrecisionUint(), result.GetPrecision())
  }
}

func TestBigIntMathMultiply_FixedDecimalMultiply_06(t *testing.T) {

  multiplierStr := "0.00000"

  multiplicandStr := "0.00"

  // product = 0.00
  expectedBigINumStr := "0"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  multiplier :=
    new(BigIntFixedDecimal).New(
      multiplierBiNum.GetIntegerValue(),
      multiplierBiNum.GetPrecisionUint())

  multiplicand :=
    new(BigIntFixedDecimal).New(
      multiplicandBiNum.GetIntegerValue(),
      multiplicandBiNum.GetPrecisionUint())

  result := new(BigIntMathMultiply).FixedDecimalMultiply(
    multiplier,
    multiplicand)

  if expectedBigINum.GetIntegerValue().Cmp(result.GetInteger()) != 0 {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%s'. ",
      expectedBigINum.GetIntegerValue().Text(10), result.GetInteger().Text(10))
  }

  if expectedBigINum.GetPrecisionUint() != result.GetPrecision() {
    t.Errorf("Expected result precision='%v'. Instead, result precision='%v'",
      expectedBigINum.GetPrecisionUint(), result.GetPrecision())
  }
}

func TestBigIntMathMultiply_FixedDecimalMultiply_07(t *testing.T) {
  // multiplier = 0.12345
  multiplierStr := "0.12345"

  // multiplicand = 32768
  multiplicandStr := "32768"

  // product = 4045.2096
  expectedBigINumStr := "4045.2096"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  multiplier :=
    new(BigIntFixedDecimal).New(
      multiplierBiNum.GetIntegerValue(),
      multiplierBiNum.GetPrecisionUint())

  multiplicand :=
    new(BigIntFixedDecimal).New(
      multiplicandBiNum.GetIntegerValue(),
      multiplicandBiNum.GetPrecisionUint())

  result := new(BigIntMathMultiply).FixedDecimalMultiply(
    multiplier,
    multiplicand)

  if expectedBigINum.GetNumStr() != result.GetNumStr() {
    t.Errorf("Error: Expected result='%v'. Instead, result= '%v'. ",
      expectedBigINum.GetNumStr(), result.GetNumStr())
  }
}

func TestBigIntMathMultiply_MultiplyBigInts_01(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedBigINumStr := "2875.94572"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
  multiplierPrecision := multiplierBiNum.GetPrecisionUint()
  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
  multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

  numSepsDto := NumericSeparatorDto{}

  numSepsDto.SetDefaultsIfEmpty()

  result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision,
    numSepsDto)

  if err != nil {
    t.Errorf("Error returned by BigIntMathMultiply.MultiplyBigIntsBigIntNum()\n"+
      "Error: %v",
      err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

}

func TestBigIntMathMultiply_MultiplyBigInts_02(t *testing.T) {
  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedBigINumStr := "30987680500513189125.14259702468435"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
  multiplierPrecision := multiplierBiNum.GetPrecisionUint()
  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
  multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

  numSepsDto := NumericSeparatorDto{}

  numSepsDto.SetDefaultsIfEmpty()

  result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision,
    numSepsDto)

  if err != nil {
    t.Errorf("Error returned by BigIntMathMultiply.MultiplyBigIntsBigIntNum()\n"+
      "Error: %v",
      err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyBigInts_03(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedBigINumStr := "-30987680500513189125.14259702468435"

  expectedBigINumSign := -1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
  multiplierPrecision := multiplierBiNum.GetPrecisionUint()
  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
  multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

  numSepsDto := NumericSeparatorDto{}
  numSepsDto.SetDefaultsIfEmpty()

  result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision,
    numSepsDto)

  if err != nil {
    t.Errorf("Error returned by BigIntMathMultiply.MultiplyBigIntsBigIntNum()\n"+
      "Error: %v",
      err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyBigInts_04(t *testing.T) {
  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedBigINumStr := "22197234145.3632"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
  multiplierPrecision := multiplierBiNum.GetPrecisionUint()
  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
  multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

  numSepsDto := NumericSeparatorDto{}

  numSepsDto.SetDefaultsIfEmpty()

  result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision,
    numSepsDto)

  if err != nil {
    t.Errorf("Error returned by BigIntMathMultiply.MultiplyBigIntsBigIntNum()\n"+
      "Error: %v",
      err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyBigInts_05(t *testing.T) {
  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0.00
  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandBiNum, err := new(BigIntNum).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  bIMultiplier := big.NewInt(0).Set(multiplierBiNum.bigInt)
  multiplierPrecision := multiplierBiNum.GetPrecisionUint()
  biMultiplicand := big.NewInt(0).Set(multiplicandBiNum.bigInt)
  multiplicandPrecision := multiplicandBiNum.GetPrecisionUint()

  numSepsDto := NumericSeparatorDto{}

  numSepsDto.SetDefaultsIfEmpty()

  result, err := new(BigIntMathMultiply).MultiplyBigIntsBigIntNum(
    bIMultiplier,
    multiplierPrecision,
    biMultiplicand,
    multiplicandPrecision,
    numSepsDto)

  if err != nil {
    t.Errorf("Error returned by BigIntMathMultiply.MultiplyBigIntsBigIntNum()\n"+
      "Error: %v",
      err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}
