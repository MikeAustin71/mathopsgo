package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntMathAdd_AddBigInts_01(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigInts_01"
  // n1Str := 123456.789
  b1Str := "123456789"
  b1Precision := uint(3)

  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  // n2Str := 987.123456
  b2Str := "987123456"
  b2Precision := uint(6)

  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  // Result := 124443.912456
  expectedResultStr := "124443912456"
  expectedPrecision := uint(6)
  expectedSign := 1

  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
      "expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
    return
  }

  result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigInts_02(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigInts_02"
  // n1Str := 123456.789
  b1Str := "123456789"
  b1Precision := uint(3)

  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  // n2Str := -987.123456
  b2Str := "-987123456"
  b2Precision := uint(6)
  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  // Result := 122469.665544
  expectedResultStr := "122469665544"
  expectedPrecision := uint(6)
  expectedSign := 1
  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
      "expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
    return
  }

  result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigInts_03(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigInts_03"
  // n1Str := -123456.789
  b1Str := "-123456789"
  b1Precision := uint(3)

  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  // n2Str := 987.123456
  b2Str := "987123456"
  b2Precision := uint(6)

  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  // Result := -122469.665544
  expectedResultStr := "-122469665544"
  expectedPrecision := uint(6)
  expectedSign := -1

  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
      "expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
    return
  }

  result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigInts_04(t *testing.T) {
  ePrefix := "TestBigIntMathAdd_AddBigInts_04"
  // n1Str := -123456.789
  b1Str := "-123456789"
  b1Precision := uint(3)

  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  // n2Str := -987.123456
  b2Str := "-987123456"
  b2Precision := uint(6)

  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  // Result := -124443.912456
  expectedResultStr := "-124443912456"
  expectedPrecision := uint(6)
  expectedSign := -1

  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
      "expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
    return
  }

  result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigInts(b1Big, b1Precision, b2Big, b2Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNums_01(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNums_01"
  // n1Str := 123456.789
  b1Str := "123456789"
  b1Precision := uint(3)

  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      err.Error())
    return
  }

  b1NumStr, err := b1Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1NumStr, err := b1Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // n2Str := 987.123456
  b2Str := "987123456"
  b2Precision := uint(6)

  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  b2NumStr, err := b2Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2NumStr, err := b2Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // Result := 124443.912456
  expectedResultStr := "124443912456"
  expectedPrecision := uint(6)
  expectedSign := 1

  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
      "expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
    return
  }

  result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)\n"+
      "b1Num= '%v'\n"+
      "b2Num= '%v'\n"+
      "Error='%v'\n\n",
      b1NumStr,
      b2NumStr,
      ePrefix,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNums_02(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNums_02"

  // n1Str := 123456.789
  b1Str := "123456789"
  b1Precision := uint(3)

  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      err.Error())
    return
  }

  b1NumStr, err := b1Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1NumStr, err := b1Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // n2Str := -987.123456
  b2Str := "-987123456"
  b2Precision := uint(6)

  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  b2NumStr, err := b2Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2NumStr, err := b2Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // Result := 122469.665544
  expectedResultStr := "122469665544"
  expectedPrecision := uint(6)
  expectedSign := 1

  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
      "expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
    return
  }

  result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)\n"+
      "b1Num= '%v'\n"+
      "b2Num= '%v'\n"+
      "Error='%v'\n\n",
      b1NumStr,
      b2NumStr,
      ePrefix,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNums_03(t *testing.T) {
  ePrefix := "TestBigIntMathAdd_AddBigIntNums_03"
  // n1Str := -123456.789
  b1Str := "-123456789"
  b1Precision := uint(3)
  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      err.Error())
    return
  }

  b1NumStr, err := b1Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1NumStr, err := b1Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // n2Str := 987.123456
  b2Str := "987123456"
  b2Precision := uint(6)
  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  b2NumStr, err := b2Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2NumStr, err := b2Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // Result := -122469.665544
  expectedResultStr := "-122469665544"
  expectedPrecision := uint(6)
  expectedSign := -1
  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)\n"+
      "b1Num= '%v'\n"+
      "b2Num= '%v'\n"+
      "Error='%v'\n\n",
      b1NumStr,
      b2NumStr,
      ePrefix,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNums_04(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNums_04"
  // n1Str := -123456.789
  b1Str := "-123456789"
  b1Precision := uint(3)
  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Big, oK := big.NewInt(0).SetString(b1Str, 10)\n"+
      "b1Str='%v'\n\n", ePrefix, b1Str)
    return
  }

  b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Num, err := new(BigIntNum).NewBigInt(b1Big, b1Precision)\n"+
      "b1Big= '%v'\n"+
      "b1Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Big.Text(10),
      b1Precision,
      err.Error())
    return
  }

  b1NumStr, err := b1Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1NumStr, err := b1Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // n2Str := -987.123456
  b2Str := "-987123456"
  b2Precision := uint(6)

  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Big, oK := big.NewInt(0).SetString(b2Str, 10)\n"+
      "b2Str='%v'\n\n", ePrefix, b2Str)
    return
  }

  b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Num, err := new(BigIntNum).NewBigInt(b2Big, b2Precision)\n"+
      "b2Big= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b2Big.Text(10),
      b2Precision,
      err.Error())
    return
  }

  b2NumStr, err := b2Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2NumStr, err := b2Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // Result := -124443.912456
  expectedResultStr := "-124443912456"
  expectedPrecision := uint(6)
  expectedSign := -1

  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)\n"+
      "expectedResultStr='%v'\n\n", ePrefix, expectedResultStr)
    return
  }

  result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)\n"+
      "b1Num= '%v'\n"+
      "b2Num= '%v'\n"+
      "Error='%v'\n\n",
      b1NumStr,
      b2NumStr,
      ePrefix,
      err.Error())
    return
  }

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.bigInt = '%v'\n"+
      "Instead, result.bigInt = '%v'\n\n",
      ePrefix,
      biExpectedResult.Text(10),
      result.bigInt.Text(10))
    return
  }

  if expectedPrecision != result.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.precision = '%v'\n"+
      "Instead, result.precision = '%v'\n\n",
      ePrefix,
      expectedPrecision,
      result.precision)
    return
  }

  if expectedSign != result.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected result.sign = '%v'\n"+
      "Instead, result.sign = '%v'\n\n",
      ePrefix,
      expectedSign,
      result.sign)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNums_05(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNums_05"

  // n1Str := 123456.789
  b1Str := "123456.789"

  b1Num, err := new(BigIntNum).NewNumStr(b1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1Num, err := new(BigIntNum).NewNumStr(b1Str)\n"+
      "b1Str= '%v'\n"+
      "Error='%v'\n\n", ePrefix,
      b1Str,
      err.Error())
    return
  }

  b1NumStr, err := b1Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1NumStr, err := b1Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = b1Num.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = b1Num.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\nError='%v'\n\n", ePrefix,
      expectedNumSeps.String(),
      err.Error())
    return
  }

  // n2Str := 987.123456
  b2Str := "987.123456"

  b2Num, err := new(BigIntNum).NewNumStr(b2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2Num, err := new(BigIntNum).NewNumStr(b2Str)\n"+
      "b2Str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      b2Str,
      err.Error())
    return
  }

  b2NumStr, err := b2Num.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2NumStr, err := b2Num.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  // Result := 124443.912456
  expectedResultStr := "124443,912456"

  result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)\n"+
      "b1Num= '%v'\n"+
      "b2Num= '%v'\n"+
      "Error='%v'\n\n",
      b1NumStr,
      b2NumStr,
      ePrefix,
      err.Error())
    return
  }

  actualResultNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultStr != actualResultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualResultNumStr = '%v'\n"+
      "Instead, actualResultNumStr = '%v'\n\n",
      ePrefix,
      expectedResultStr,
      actualResultNumStr)
    return
  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "result= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      actualResultNumStr,
      err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNumSeps = '%v'\n"+
      "Instead, actualNumSeps = '%v'\n\n",
      ePrefix,
      expectedNumSeps.String(),
      actualNumSeps.String())
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNumArray_01"

  numStrs := []string{"45.8",
    "1.45962",
    "58.71",
    "-37.62174",
    "89.8",
  }

  expectedTotalStr := "158.14788"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
      "expectedTotalStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedTotalStr,
      err.Error())
    return
  }

  expectedBNumStr, err := expectedBNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBNumStr, err := expectedBNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedResultNumStr, err := expectedBNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumStr, err := expectedBNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  const lenBigNums = 5

  bNums := make([]BigIntNum, lenBigNums)

  for i := 0; i < lenBigNums; i++ {

    bNums[i], err = new(BigIntNum).NewNumStr(numStrs[i])

    if err != nil {

      t.Errorf("%v\n"+
        "Error returned by new(BigIntNum).NewNumStr(numStrs[%d])\n"+
        "numStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        i,
        numStrs[i],
        err.Error())
    }
  }

  results, err := new(BigIntMathAdd).AddBigIntNumArray(bNums)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "results, err := new(BigIntMathAdd).AddBigIntNumArray(bNums)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualResultNumStr, err := results.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResultNumStr, err := results.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != actualResultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualResultNumStr = '%v'\n"+
      "Instead, actualResultNumStr = '%v'\n\n",
      ePrefix,
      expectedResultNumStr,
      actualResultNumStr)
    return
  }

  resultsEqualToBNum, err := expectedBNum.Equal(results)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultsEqualToBNum, err = expectedBNum.Equal(results)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !resultsEqualToBNum {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected results = '%v'\n"+
      "Instead, results = '%v'\n\n",
      ePrefix,
      expectedBNum.bigInt.Text(10),
      results.bigInt.Text(10))
    return
  }

  if expectedBNumStr != actualResultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected resultsNumStr = '%v'\n"+
      "Instead, resultsNumStr = '%v'\n\n",
      ePrefix,
      expectedBNumStr,
      actualResultNumStr)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNumArray_02"

  numStrs := []string{"-978425.648941",
    "33.12",
    "-804.1",
    "32567",
    "-41.859",
  }

  expectedTotalStr := "-946671.487941"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
      "expectedTotalStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedTotalStr,
      err.Error())
    return
  }

  expectedResultNumStr, err := expectedBNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumStr, err := expectedBNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  const lenBigNums = 5

  bNums := make([]BigIntNum, lenBigNums)

  for i := 0; i < lenBigNums; i++ {

    bNums[i], err = new(BigIntNum).NewNumStr(numStrs[i])

    if err != nil {

      t.Errorf("%v\n"+
        "Error returned by new(BigIntNum).NewNumStr(numStrs[%d])\n"+
        "numStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        i,
        numStrs[i],
        err.Error())
    }
  }

  results, err := new(BigIntMathAdd).AddBigIntNumArray(bNums)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "results, err := new(BigIntMathAdd).AddBigIntNumArray(bNums)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualResultNumStr, err := results.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResultNumStr, err := results.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResultNumStr != actualResultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualResultNumStr = '%v'\n"+
      "Instead, actualResultNumStr = '%v'\n\n",
      ePrefix,
      expectedResultNumStr,
      actualResultNumStr)
    return
  }

  resultsEqualToBNum, err := expectedBNum.Equal(results)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultsEqualToBNum, err = expectedBNum.Equal(results)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !resultsEqualToBNum {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected results = '%v'\n"+
      "Instead, results = '%v'\n\n",
      ePrefix,
      expectedBNum.bigInt.Text(10),
      results.bigInt.Text(10))
    return
  }

  if expectedResultNumStr != actualResultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected resultsNumStr = '%v'\n"+
      "Instead, resultsNumStr = '%v'\n\n",
      ePrefix,
      expectedResultNumStr,
      actualResultNumStr)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumArray_03(t *testing.T) {
  ePrefix := "TestBigIntMathAdd_AddBigIntNumArray_03"
  numStrs := []string{"-978425.648941",
    "33.12",
    "-804.1",
    "32567",
    "-41.859",
  }

  expectedTotalStr := "-946671,487941"

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  const lenBigNums = 5
  var err error
  bNums := make([]BigIntNum, lenBigNums)

  for i := 0; i < lenBigNums; i++ {

    bNums[i], err = new(BigIntNum).NewNumStr(numStrs[i])

    if err != nil {

      t.Errorf("%v\n"+
        "Error returned by\n"+
        "new(BigIntNum).NewNumStr(numStrs[%d])\n"+
        "numStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        i,
        numStrs[i],
        err.Error())
    }
  }

  err = bNums[0].SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bNums[0].SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedNumSeps.String(),
      err.Error())
    return
  }

  results, err := new(BigIntMathAdd).AddBigIntNumArray(bNums)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "results, err := new(BigIntMathAdd).AddBigIntNumArray(bNums)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  actualResultNumStr, err := results.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualResultNumStr, err := results.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedTotalStr != actualResultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualResultNumStr = '%v'\n"+
      "Instead, actualResultNumStr = '%v'\n\n",
      ePrefix,
      expectedTotalStr,
      actualResultNumStr)
    return
  }

  actualNumSeps, err := results.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumSeps, err := results.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNumSeps = '%v'\n"+
      "Instead, actualNumSeps = '%v'\n\n",
      ePrefix,
      expectedNumSeps.String(),
      actualNumSeps.String())
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNumOutputToArray_01"

  var err error

  // multiplier = 0
  addendStr := "5"
  // bNumStrs
  bNumStrs := []string{
    "5",
    "10.123",
    "15",
    "253.692",
    "35",
    "55",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "10",
    "15.123",
    "20",
    "258.692",
    "40",
    "60",
  }

  addendBiNum, err := new(BigIntNum).NewNumStr(addendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "addendBiNum, err := new(BigIntNum).NewNumStr(addendStr)\n"+
      "addendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendStr,
      err.Error())
    return
  }

  addendBiNumStr, err := addendBiNum.NewNumStr(addendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "addendBiNumStr, err := addendBiNum.NewNumStr(addendStr)\n"+
      "addendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendStr,
      err.Error())
    return
  }

  lenArray := len(bNumStrs)
  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(bNumStrs[i])

    if err != nil {

      t.Errorf("%v\n"+
        "Error returned by\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(bNumStrs[%d])\n"+
        "bNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        i,
        i,
        bNumStrs[i],
        err.Error())
    }
  }

  result, err := new(BigIntMathAdd).AddBigIntNumOutputToArray(
    addendBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNumOutputToArray(addendBiNum, bINumArray)\n"+
      "addendBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendBiNumStr,
      err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected resultNumStr = '%v'\n"+
        "Instead, resultNumStr = '%v'\n"+
        "Cycle Index = '%v'\n\n",
        ePrefix,
        expectedNumStrs[j],
        resultNumStr,
        j)
      return
    }
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNumOutputToArray_02"

  var err error

  // multiplier = 0
  addendStr := "3.1"
  // bNumStrs
  bNumStrs := []string{
    "5",
    "10.123",
    "0",
    "253.692",
    "35",
    "55",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "8.1",
    "13.223",
    "3.1",
    "256.792",
    "38.1",
    "58.1",
  }

  addendBiNum, err := new(BigIntNum).NewNumStr(addendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "addendBiNum, err := new(BigIntNum).NewNumStr(addendStr)\n"+
      "addendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendStr,
      err.Error())
    return
  }

  addendBiNumStr, err := addendBiNum.NewNumStr(addendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "addendBiNumStr, err := addendBiNum.NewNumStr(addendStr)\n"+
      "addendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendStr,
      err.Error())
    return
  }

  lenArray := len(bNumStrs)
  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(bNumStrs[i])

    if err != nil {

      t.Errorf("%v\n"+
        "Error returned by\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(bNumStrs[%d])\n"+
        "bNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        i,
        i,
        bNumStrs[i],
        err.Error())
    }
  }

  result, err := new(BigIntMathAdd).AddBigIntNumOutputToArray(
    addendBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNumOutputToArray(addendBiNum, bINumArray)\n"+
      "addendBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendBiNumStr,
      err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected resultNumStr = '%v'\n"+
        "Instead, resultNumStr = '%v'\n"+
        "Cycle Index = '%v'\n\n",
        ePrefix,
        expectedNumStrs[j],
        resultNumStr,
        j)
      return
    }
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathAdd_AddBigIntNumOutputToArray_03"

  var err error

  // multiplier = 0
  addendStr := "5"
  // bNumStrs
  bNumStrs := []string{
    "5",
    "10.123",
    "15",
    "253.692",
    "35",
    "55",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "10",
    "15,123",
    "20",
    "258,692",
    "40",
    "60",
  }

  addendBiNum, err := new(BigIntNum).NewNumStr(addendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "addendBiNum, err := new(BigIntNum).NewNumStr(addendStr)\n"+
      "addendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendStr,
      err.Error())
    return
  }

  addendBiNumStr, err := addendBiNum.NewNumStr(addendStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "addendBiNumStr, err := addendBiNum.NewNumStr(addendStr)\n"+
      "addendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendStr,
      err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = addendBiNum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = addendBiNum.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(bNumStrs)
  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(bNumStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(bNumStrs[%d])\n"+
        "bNumStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, i, bNumStrs[i], err.Error())
      return
    }
  }

  result, err := new(BigIntMathAdd).AddBigIntNumOutputToArray(
    addendBiNum, bINumArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathAdd).AddBigIntNumOutputToArray(addendBiNum, bINumArray)\n"+
      "addendBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      addendBiNumStr,
      err.Error())
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
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    if expectedNumStrs[j] != resultNumStr {

      t.Errorf("%v\n"+
        "Error: Unexpected Result!\n"+
        "Expected resultNumStr = '%v'\n"+
        "Instead, resultNumStr = '%v'\n"+
        "Cycle Index = '%v'\n\n",
        ePrefix,
        expectedNumStrs[j],
        resultNumStr,
        j)
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
        "Error: Unexpected Result!\n"+
        "Numeric Serparators don't match.\n"+
        "Expected actualNumSeps = '%v'\n"+
        "Instead, actualNumSeps = '%v'\n"+
        "Cycle Index = '%v'\n\n",
        ePrefix,
        expectedNumSeps.String(),
        actualNumSeps.String(),
        j)
      return
    }
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumSeries_01(t *testing.T) {
  ePrefix := "TestBigIntMathAdd_AddBigIntNumSeries_01"
  n1Str := "45.8"
  n2Str := "1.45962"
  n3Str := "58.71"
  n4Str := "-37.62174"
  n5Str := "89.8"
  expectedTotalStr := "158.14788"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)\n"+
      "expectedTotalStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedTotalStr, err.Error())
    return
  }

  expectedResultNumStr, err := expectedBNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumStr, err := expectedBNum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  b1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, err := new(BigIntNum).NewNumStr(n1Str)\n"+
      "n1Str= '%v'\n"+
      "Error='%v'\n\n", ePrefix, n1Str, err.Error())
    return
  }

  b1NumStr, err := b1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1NumStr, err := b1.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  b2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, err := new(BigIntNum).NewNumStr(n2Str)\n"+
      "n2Str= '%v'\n"+
      "Error='%v'\n\n", ePrefix, n2Str, err.Error())
    return
  }

  b2NumStr, err := b2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2NumStr, err := b2.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  b3, err := new(BigIntNum).NewNumStr(n3Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b3, err := new(BigIntNum).NewNumStr(n3Str)\n"+
      "n3Str= '%v'\n"+
      "Error='%v'\n\n", ePrefix, n3Str, err.Error())
    return
  }

  b3NumStr, err := b3.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b3NumStr, err := b3.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  b4, err := new(BigIntNum).NewNumStr(n4Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b4, err := new(BigIntNum).NewNumStr(n4Str)\n"+
      "n4Str= '%v'\n"+
      "Error='%v'\n\n", ePrefix, n4Str, err.Error())
    return
  }

  b4NumStr, err := b4.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b4NumStr, err := b4.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  b5, err := new(BigIntNum).NewNumStr(n5Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b5, err := new(BigIntNum).NewNumStr(n5Str)\n"+
      "n5Str= '%v'\n"+
      "Error='%v'\n\n", ePrefix, n5Str, err.Error())
    return
  }

  b5NumStr, err := b5.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b5NumStr, err := b5.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  total, err := new(BigIntMathAdd).AddBigIntNumSeries(b1, b2, b3, b4, b5)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "total, err := new(BigIntMathAdd).AddBigIntNumSeries(b1, b2, b3, b4, b5)\n"+
      "b1= '%v'\n"+
      "b2= '%v'\n"+
      "b3= '%v'\n"+
      "b4= '%v'\n"+
      "b5= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      b1NumStr,
      b2NumStr,
      b3NumStr,
      b4NumStr,
      b5NumStr,
      err.Error())
    return
  }

  totalNumStr, err := total.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "totalNumStr, err := total.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBNumEqualsTotal, err := expectedBNum.Equal(total)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBNumEqualsTotal, err := expectedBNum.Equal(total)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBNumEqualsTotal {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected total = '%v'\n"+
      "Instead, total = '%v'\n\n",
      ePrefix, expectedResultNumStr, totalNumStr)
  }

  return
}

func TestBigIntMathAdd_AddBigIntNumSeries_02(t *testing.T) {
  n1Str := "-978425.648941"
  n2Str := "33.12"
  n3Str := "-804.1"
  n4Str := "32567"
  n5Str := "-41.859"
  expectedTotalStr := "-946671.487941"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  b1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'.", n1Str, err.Error())
  }

  b2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'.", n2Str, err.Error())
  }

  b3, err := new(BigIntNum).NewNumStr(n3Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n3Str). "+
      "n3Str='%v' Error='%v'.", n3Str, err.Error())
  }

  b4, err := new(BigIntNum).NewNumStr(n4Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n4Str). "+
      "n4Str='%v' Error='%v'.", n4Str, err.Error())
  }

  b5, err := new(BigIntNum).NewNumStr(n5Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n5Str). "+
      "n5Str='%v' Error='%v'.", n5Str, err.Error())
  }

  total := new(BigIntMathAdd).AddBigIntNumSeries(b1, b2, b3, b4, b5)

  if !expectedBNum.Equal(total) {
    t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedResultNumStr, total.bigInt.Text(10))
  }

}

func TestBigIntMathAdd_AddBigIntNumSeries_03(t *testing.T) {
  n1Str := "45.8"
  n2Str := "1.45962"
  n3Str := "58.71"
  n4Str := "-37.62174"
  n5Str := "89.8"
  expectedTotalStr := "158.14788"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  b1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'.", n1Str, err.Error())
  }

  b2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'.", n2Str, err.Error())
  }

  b3, err := new(BigIntNum).NewNumStr(n3Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n3Str). "+
      "n3Str='%v' Error='%v'.", n3Str, err.Error())
  }

  b4, err := new(BigIntNum).NewNumStr(n4Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n4Str). "+
      "n4Str='%v' Error='%v'.", n4Str, err.Error())
  }

  b5, err := new(BigIntNum).NewNumStr(n5Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n5Str). "+
      "n5Str='%v' Error='%v'.", n5Str, err.Error())
  }

  total := new(BigIntMathAdd).AddBigIntNumSeries(b1, b2, b3, b4, b5)

  if !expectedBNum.Equal(total) {
    t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedResultNumStr, total.bigInt.Text(10))
  }

}

func TestBigIntMathAdd_AddBigIntNumSeries_04(t *testing.T) {
  n1Str := "45.8"
  n2Str := "1.45962"
  n3Str := "58.71"
  n4Str := "-37.62174"
  n5Str := "89.8"
  expectedResultNumStr := "158,14788"

  b1, err := new(BigIntNum).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'.", n1Str, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = b1.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by b1.SetNumericSeparatorsDto(expectedNumSeps) "+
      "Error='%v' ", err.Error())
  }

  b2, err := new(BigIntNum).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'.", n2Str, err.Error())
  }

  b3, err := new(BigIntNum).NewNumStr(n3Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n3Str). "+
      "n3Str='%v' Error='%v'.", n3Str, err.Error())
  }

  b4, err := new(BigIntNum).NewNumStr(n4Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n4Str). "+
      "n4Str='%v' Error='%v'.", n4Str, err.Error())
  }

  b5, err := new(BigIntNum).NewNumStr(n5Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(n5Str). "+
      "n5Str='%v' Error='%v'.", n5Str, err.Error())
  }

  total := new(BigIntMathAdd).AddBigIntNumSeries(b1, b2, b3, b4, b5)

  actualNumStr := total.GetNumStr()

  if expectedResultNumStr != actualNumStr {
    t.Errorf("Error: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedResultNumStr, actualNumStr)
  }

  actualNumSeps := total.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathAdd_AddDecimal_01(t *testing.T) {
  n1Str := "123456.789"

  n2Str := "987.123456"

  expectedNumStr := "124443.912456"
  expectedPrecision := uint(6)
  expectedSign := 1

  dec1, err := Decimal{}.NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v' ", n1Str, err.Error())
  }

  dec2, err := Decimal{}.NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v' ", n2Str, err.Error())
  }

  result, err := new(BigIntMathAdd).AddDecimal(dec1, dec2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimal(dec1, dec2). "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != result.precision {
    t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
      expectedPrecision, result.precision)
  }

  if expectedSign != result.sign {
    t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
      expectedSign, result.sign)
  }

}

func TestBigIntMathAdd_AddDecimal_02(t *testing.T) {
  n1Str := "123456.789"

  n2Str := "-987.123456"

  expectedNumStr := "122469.665544"
  expectedPrecision := uint(6)
  expectedSign := 1

  dec1, err := Decimal{}.NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v' ", n1Str, err.Error())
  }

  dec2, err := Decimal{}.NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v' ", n2Str, err.Error())
  }

  result, err := new(BigIntMathAdd).AddDecimal(dec1, dec2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimal(dec1, dec2). "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != result.precision {
    t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
      expectedPrecision, result.precision)
  }

  if expectedSign != result.sign {
    t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
      expectedSign, result.sign)
  }

}

func TestBigIntMathAdd_AddDecimal_03(t *testing.T) {
  n1Str := "-123456.789"

  n2Str := "987.123456"

  expectedNumStr := "-122469.665544"
  expectedPrecision := uint(6)
  expectedSign := -1

  dec1, err := Decimal{}.NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v' ", n1Str, err.Error())
  }

  dec2, err := Decimal{}.NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v' ", n2Str, err.Error())
  }

  result, err := new(BigIntMathAdd).AddDecimal(dec1, dec2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimal(dec1, dec2). "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != result.precision {
    t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
      expectedPrecision, result.precision)
  }

  if expectedSign != result.sign {
    t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
      expectedSign, result.sign)
  }

}

func TestBigIntMathAdd_AddDecimal_04(t *testing.T) {
  n1Str := "123456.789"

  n2Str := "987.123456"

  expectedNumStr := "124443,912456"
  expectedPrecision := uint(6)
  expectedSign := 1

  dec1, err := Decimal{}.NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v' ", n1Str, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = dec1.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by dec1.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v'", err.Error())
  }

  dec2, err := Decimal{}.NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v' ", n2Str, err.Error())
  }

  result, err := new(BigIntMathAdd).AddDecimal(dec1, dec2)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimal(dec1, dec2). "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v' ",
      expectedNumStr, actualNumStr)
  }

  if expectedPrecision != result.precision {
    t.Errorf("Error: Expected Result precision='%v'.  Instead, precision='%v'. ",
      expectedPrecision, result.precision)
  }

  if expectedSign != result.sign {
    t.Errorf("Error: Expected Result sign='%v'. Instead, sign='%v'. ",
      expectedSign, result.sign)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathAdd_AddDecimalArray_01(t *testing.T) {

  numStrAry := []string{
    "45.8",
    "1.45962",
    "58.71",
    "-37.62174",
    "89.8",
  }

  lenStrAry := len(numStrAry)

  expectedTotalStr := "158.14788"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  decAry := make([]Decimal, lenStrAry)

  for i := 0; i < lenStrAry; i++ {

    dec, err := Decimal{}.NewNumStr(numStrAry[i])

    if err != nil {

      if err != nil {
        t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) "+
          "i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
      }

    }

    decAry[i] = dec

  }

  total, err := new(BigIntMathAdd).AddDecimalArray(decAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalArray(decAry). "+
      "Error='%v' ", err.Error())
  }

  if !expectedBNum.Equal(total) {
    t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
  }

  actualTotalNumstr := total.GetNumStr()

  if expectedResultNumStr != actualTotalNumstr {
    t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedResultNumStr, actualTotalNumstr)
  }

}

func TestBigIntMathAdd_AddDecimalArray_02(t *testing.T) {

  numStrAry := []string{
    "-978425.648941",
    "33.12",
    "-804.1",
    "32567",
    "-41.859",
  }

  lenStrAry := len(numStrAry)

  expectedTotalStr := "-946671.487941"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  decAry := make([]Decimal, lenStrAry)

  for i := 0; i < lenStrAry; i++ {

    dec, err := Decimal{}.NewNumStr(numStrAry[i])

    if err != nil {

      if err != nil {
        t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) "+
          "i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
      }

    }

    decAry[i] = dec

  }

  total, err := new(BigIntMathAdd).AddDecimalArray(decAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalArray(decAry). "+
      "Error='%v' ", err.Error())
  }

  if !expectedBNum.Equal(total) {
    t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
  }

  actualTotalNumstr := total.GetNumStr()

  if expectedResultNumStr != actualTotalNumstr {
    t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedResultNumStr, actualTotalNumstr)
  }

}

func TestBigIntMathAdd_AddDecimalArray_03(t *testing.T) {

  numStrAry := []string{
    "45.8",
    "1.45962",
    "58.71",
    "-37.62174",
    "89.8",
  }

  lenStrAry := len(numStrAry)

  expectedTotalStr := "158,14788"

  decAry := make([]Decimal, lenStrAry)

  for i := 0; i < lenStrAry; i++ {

    dec, err := Decimal{}.NewNumStr(numStrAry[i])

    if err != nil {

      if err != nil {
        t.Errorf("Error returned by Decimal{}.NewNumStr(numStrAry[i]) "+
          "i='%v' numStrAry[i]='%v' Error='%v' ", i, numStrAry[i], err.Error())
      }

    }

    decAry[i] = dec

  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err := decAry[0].SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by decAry[0].SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  total, err := new(BigIntMathAdd).AddDecimalArray(decAry)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalArray(decAry). "+
      "Error='%v' ", err.Error())
  }

  actualTotalNumStr := total.GetNumStr()

  if expectedTotalStr != actualTotalNumStr {
    t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedTotalStr, actualTotalNumStr)
  }

  actualNumSeps := total.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathAdd_AddDecimalOutputToArray_01(t *testing.T) {

  var err error

  // addendStr = 5
  addendStr := "5"

  // decNumStrs
  decNumStrs := []string{
    "5",
    "10.123",
    "15",
    "253.692",
    "35",
    "55",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "10",
    "15.123",
    "20",
    "258.692",
    "40",
    "60",
  }

  decAddend, err := Decimal{}.NewNumStr(addendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) "+
      "addendStr='%v'  Error='%v'. ", addendStr, err.Error())
  }

  lenArray := len(decNumStrs)
  decArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decArray[i], err = Decimal{}.NewNumStr(decNumStrs[i])

    if err != nil {
      t.Errorf("Error returned by Decimal{}.NewNumStr(decNumStrs[i]) "+
        "i='%v'  decNumStrs[i]='%v'  Error='%v'. ", i, decNumStrs[i], err.Error())
    }

  }

  result, err := new(BigIntMathAdd).AddDecimalOutputToArray(decAddend, decArray)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalOutputToArray("+
      "decAddend, decArray) addendStr='%v'  Error='%v'. ",
      addendStr, err.Error())
  }

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j].GetNumStr() {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j].GetNumStr())
    }
  }
}

func TestBigIntMathAdd_AddDecimalOutputToArray_02(t *testing.T) {

  var err error

  // addendStr = 3.1
  addendStr := "3.1"

  // decNumStrs
  decNumStrs := []string{
    "5",
    "10.123",
    "0",
    "253.692",
    "35",
    "55",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "8.1",
    "13.223",
    "3.1",
    "256.792",
    "38.1",
    "58.1",
  }

  decAddend, err := Decimal{}.NewNumStr(addendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) "+
      "addendStr='%v'  Error='%v'. ", addendStr, err.Error())
  }

  lenArray := len(decNumStrs)
  decsArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decsArray[i], err = Decimal{}.NewNumStr(decNumStrs[i])

    if err != nil {
      t.Errorf("Error returned by Decimal{}.NewNumStr(decNumStrs[i]) "+
        "i='%v'  decNumStrs[i]='%v'  Error='%v'. ", i, decNumStrs[i], err.Error())
    }

  }

  result, err := new(BigIntMathAdd).AddDecimalOutputToArray(decAddend, decsArray)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalOutputToArray("+
      "decAddend, decsArray) addendStr='%v'  Error='%v'. ",
      addendStr, err.Error())
  }

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j].GetNumStr() {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j].GetNumStr())
    }
  }
}

func TestBigIntMathAdd_AddDecimalOutputToArray_03(t *testing.T) {

  var err error

  // addendStr = 5
  addendStr := "5"

  // decNumStrs
  decNumStrs := []string{
    "5",
    "10.123",
    "15",
    "253.692",
    "35",
    "55",
  }

  // Expected Results Array
  expectedNumStrs := []string{
    "10",
    "15,123",
    "20",
    "258,692",
    "40",
    "60",
  }

  decAddend, err := Decimal{}.NewNumStr(addendStr)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStr(addendStr) "+
      "addendStr='%v'  Error='%v'. ", addendStr, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = decAddend.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by decAddend.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenArray := len(decNumStrs)
  decArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decArray[i], err = Decimal{}.NewNumStr(decNumStrs[i])

    if err != nil {
      t.Errorf("Error returned by Decimal{}.NewNumStr(decNumStrs[i]) "+
        "i='%v'  decNumStrs[i]='%v'  Error='%v'. ", i, decNumStrs[i], err.Error())
    }

  }

  result, err := new(BigIntMathAdd).AddDecimalOutputToArray(decAddend, decArray)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalOutputToArray("+
      "decAddend, decArray) addendStr='%v'  Error='%v'. ",
      addendStr, err.Error())
  }

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j].GetNumStr() {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j].GetNumStr())
    }

    actualNumSeps := result[j].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'.",
        expectedNumSeps.String(), actualNumSeps.String())
    }

  }
}

func TestBigIntMathAdd_AddDecimalSeries_01(t *testing.T) {
  n1Str := "45.8"
  n2Str := "1.45962"
  n3Str := "58.71"
  n4Str := "-37.62174"
  n5Str := "89.8"
  expectedTotalStr := "158.14788"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  dec1, err := Decimal{}.NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'. ",
      n1Str, err.Error())
  }

  dec2, err := Decimal{}.NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'. ",
      n2Str, err.Error())
  }

  dec3, err := Decimal{}.NewNumStr(n3Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n3Str). "+
      "n3Str='%v' Error='%v'. ",
      n3Str, err.Error())
  }

  dec4, err := Decimal{}.NewNumStr(n4Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). "+
      "n4Str='%v' Error='%v'. ",
      n4Str, err.Error())
  }

  dec5, err := Decimal{}.NewNumStr(n5Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). "+
      "n5Str='%v' Error='%v'. ",
      n5Str, err.Error())
  }

  total, err := new(BigIntMathAdd).AddDecimalSeries(dec1, dec2, dec3, dec4, dec5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalSeries(dec1, "+
      "dec2, dec3, dec4, dec5). Error='%v' ", err.Error())
  }

  if !expectedBNum.Equal(total) {
    t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
  }

  actualTotalNumstr := total.GetNumStr()

  if expectedResultNumStr != actualTotalNumstr {
    t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedResultNumStr, actualTotalNumstr)
  }

}

func TestBigIntMathAdd_AddDecimalSeries_02(t *testing.T) {
  n1Str := "-978425.648941"
  n2Str := "33.12"
  n3Str := "-804.1"
  n4Str := "32567"
  n5Str := "-41.859"
  expectedTotalStr := "-946671.487941"

  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  dec1, err := Decimal{}.NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'. ",
      n1Str, err.Error())
  }

  dec2, err := Decimal{}.NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'. ",
      n2Str, err.Error())
  }

  dec3, err := Decimal{}.NewNumStr(n3Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n3Str). "+
      "n3Str='%v' Error='%v'. ",
      n3Str, err.Error())
  }

  dec4, err := Decimal{}.NewNumStr(n4Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). "+
      "n4Str='%v' Error='%v'. ",
      n4Str, err.Error())
  }

  dec5, err := Decimal{}.NewNumStr(n5Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). "+
      "n5Str='%v' Error='%v'. ",
      n5Str, err.Error())
  }

  total, err := new(BigIntMathAdd).AddDecimalSeries(dec1, dec2, dec3, dec4, dec5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalSeries(dec1, "+
      "dec2, dec3, dec4, dec5). Error='%v' ", err.Error())
  }

  if !expectedBNum.Equal(total) {
    t.Errorf("Error - Incorrect Total: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedBNum.bigInt.Text(10), total.bigInt.Text(10))
  }

  actualTotalNumstr := total.GetNumStr()

  if expectedResultNumStr != actualTotalNumstr {
    t.Errorf("Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedResultNumStr, actualTotalNumstr)
  }

}

func TestBigIntMathAdd_AddDecimalSeries_03(t *testing.T) {
  n1Str := "45.8"
  n2Str := "1.45962"
  n3Str := "58.71"
  n4Str := "-37.62174"
  n5Str := "89.8"

  expectedTotalStr := "158,14788"

  dec1, err := Decimal{}.NewNumStr(n1Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n1Str). "+
      "n1Str='%v' Error='%v'. ",
      n1Str, err.Error())
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = dec1.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by dec1.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  dec2, err := Decimal{}.NewNumStr(n2Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n2Str). "+
      "n2Str='%v' Error='%v'. ",
      n2Str, err.Error())
  }

  dec3, err := Decimal{}.NewNumStr(n3Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n3Str). "+
      "n3Str='%v' Error='%v'. ",
      n3Str, err.Error())
  }

  dec4, err := Decimal{}.NewNumStr(n4Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n4Str). "+
      "n4Str='%v' Error='%v'. ",
      n4Str, err.Error())
  }

  dec5, err := Decimal{}.NewNumStr(n5Str)

  if err != nil {
    t.Errorf("Error returned by  Decimal{}.NewNumStr(n5Str). "+
      "n5Str='%v' Error='%v'. ",
      n5Str, err.Error())
  }

  total, err := new(BigIntMathAdd).AddDecimalSeries(dec1, dec2, dec3, dec4, dec5)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathAdd).AddDecimalSeries(dec1, "+
      "dec2, dec3, dec4, dec5). Error='%v' ", err.Error())
  }

  actualTotalStr := total.GetNumStr()

  if expectedTotalStr != actualTotalStr {
    t.Errorf("Error: Expected total='%v'. Instead, "+
      "total='%v'. ",
      expectedTotalStr, actualTotalStr)
  }

  actualNumSeps := total.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}
