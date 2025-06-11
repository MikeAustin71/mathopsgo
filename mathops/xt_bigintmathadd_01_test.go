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

  // n1Str := -123456.789
  b1Str := "-123456789"
  b1Precision := uint(3)
  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
  }

  b1Num := new(BigIntNum).NewBigInt(b1Big, b1Precision)

  // n2Str := 987.123456
  b2Str := "987123456"
  b2Precision := uint(6)
  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
  }

  b2Num := new(BigIntNum).NewBigInt(b2Big, b2Precision)

  // Result := -122469.665544
  expectedResultStr := "-122469665544"
  expectedPrecision := uint(6)
  expectedSign := -1
  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  result := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
      biExpectedResult.Text(10), result.bigInt.Text(10))
  }

  if expectedPrecision != result.precision {
    t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
      expectedPrecision, result.precision)
  }

  if expectedSign != result.sign {
    t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
      expectedSign, result.sign)
  }

}

func TestBigIntMathAdd_AddBigIntNums_04(t *testing.T) {

  // n1Str := -123456.789
  b1Str := "-123456789"
  b1Precision := uint(3)
  b1Big, oK := big.NewInt(0).SetString(b1Str, 10)

  if !oK {
    t.Error("Error returned by big.NewInt(0).SetString(b1Str, 10)")
  }

  b1Num := new(BigIntNum).NewBigInt(b1Big, b1Precision)

  // n2Str := -987.123456
  b2Str := "-987123456"
  b2Precision := uint(6)
  b2Big, oK := big.NewInt(0).SetString(b2Str, 10)

  if !oK {
    t.Error("Error returned by big.NewInt(0).SetString(b2Str, 10)")
  }

  b2Num := new(BigIntNum).NewBigInt(b2Big, b2Precision)

  // Result := -124443.912456
  expectedResultStr := "-124443912456"
  expectedPrecision := uint(6)
  expectedSign := -1
  biExpectedResult, oK := big.NewInt(0).SetString(expectedResultStr, 10)

  result := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  if biExpectedResult.Cmp(result.bigInt) != 0 {
    t.Errorf("Error: Expected Result='%v'.  Instead, Result='%v'. ",
      biExpectedResult.Text(10), result.bigInt.Text(10))
  }

  if expectedPrecision != result.precision {
    t.Errorf("Error: Expected Result precision='%v'. Instead, Result precision='%v'. ",
      expectedPrecision, result.precision)
  }

  if expectedSign != result.sign {
    t.Errorf("Error: Expected Recult sign='%v'. Instead, Result sign='%v' ",
      expectedSign, result.sign)
  }

}

func TestBigIntMathAdd_AddBigIntNums_05(t *testing.T) {

  // n1Str := 123456.789
  b1Str := "123456.789"

  b1Num, err := new(BigIntNum).NewNumStr(b1Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(b1Str). "+
      "b1Str='%v' Error='%v'", b1Str, err.Error())
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
    t.Errorf("Error returned by b1Num.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v'", err.Error())
  }

  // n2Str := 987.123456
  b2Str := "987.123456"

  b2Num, err := new(BigIntNum).NewNumStr(b2Str)

  // Result := 124443.912456
  expectedResultStr := "124443,912456"

  result := new(BigIntMathAdd).AddBigIntNums(b1Num, b2Num)

  actualResultStr := result.GetNumStr()

  if expectedResultStr != actualResultStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
      expectedResultStr, actualResultStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathAdd_AddBigIntNumArray_01(t *testing.T) {

  numStrs := []string{"45.8",
    "1.45962",
    "58.71",
    "-37.62174",
    "89.8",
  }

  expectedTotalStr := "158.14788"
  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  const lenBigNums = 5

  bNums := make([]BigIntNum, lenBigNums)

  for i := 0; i < lenBigNums; i++ {

    bNums[i], err = new(BigIntNum).NewNumStr(numStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(numStrs[i]). "+
        "i='%v' numStrs[i]='%v' Error='%v'",
        i, numStrs[i], err.Error())
    }

  }

  results := new(BigIntMathAdd).AddBigIntNumArray(bNums)

  actualResultNumStr := results.GetNumStr()

  if expectedResultNumStr != actualResultNumStr {
    t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
      expectedResultNumStr, actualResultNumStr)
  }

  if !expectedBNum.Equal(results) {
    t.Errorf("Error: Expected BigIntNum to equal actual BigIntNum. It Did NOT! "+
      "BigIntTotal='%s'. Instead, BigIntTotal='%s'. ",
      expectedBNum.bigInt.Text(10), results.bigInt.Text(10))
  }

}

func TestBigIntMathAdd_AddBigIntNumArray_02(t *testing.T) {

  numStrs := []string{"-978425.648941",
    "33.12",
    "-804.1",
    "32567",
    "-41.859",
  }

  expectedTotalStr := "-946671.487941"
  expectedBNum, err := new(BigIntNum).NewNumStr(expectedTotalStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedTotalStr). "+
      "expectedTotalStr='%v' Error='%v'.", expectedTotalStr, err.Error())

  }

  expectedResultNumStr := expectedBNum.GetNumStr()

  const lenBigNums = 5

  bNums := make([]BigIntNum, lenBigNums)

  for i := 0; i < lenBigNums; i++ {

    bNums[i], err = new(BigIntNum).NewNumStr(numStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(numStrs[i]). "+
        "i='%v' numStrs[i]='%v' Error='%v'",
        i, numStrs[i], err.Error())
    }

  }

  results := new(BigIntMathAdd).AddBigIntNumArray(bNums)

  actualResultNumStr := results.GetNumStr()

  if expectedResultNumStr != actualResultNumStr {
    t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
      expectedResultNumStr, actualResultNumStr)
  }

  if !expectedBNum.Equal(results) {
    t.Errorf("Error: Expected BigIntNum != actual BigIntNum. "+
      "BigIntTotal='%v'. Instead, BigIntTotal='%v'. ",
      expectedBNum.bigInt.Text(10), results.bigInt.Text(10))
  }

}

func TestBigIntMathAdd_AddBigIntNumArray_03(t *testing.T) {

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
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(numStrs[i]). "+
        "i='%v' numStrs[i]='%v' Error='%v'",
        i, numStrs[i], err.Error())
    }

  }

  err = bNums[0].SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by bNums[0].SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v'", err.Error())
  }

  results := new(BigIntMathAdd).AddBigIntNumArray(bNums)

  actualResultNumStr := results.GetNumStr()

  if expectedTotalStr != actualResultNumStr {
    t.Errorf("Error: Expected Total='%v'. Instead, Total='%v'. ",
      expectedTotalStr, actualResultNumStr)
  }

  actualNumSeps := results.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected numSeps='%v'. Instead, numSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_01(t *testing.T) {

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
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(addendStr) "+
      "addendStr='%v'  Error='%v'. ", addendStr, err.Error())
  }

  lenArray := len(bNumStrs)
  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(bNumStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(bNumStrs[i]) "+
        "i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
    }

  }

  result := new(BigIntMathAdd).AddBigIntNumOutputToArray(addendBiNum, bINumArray)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j].GetNumStr() {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j].GetNumStr())
    }

  }

}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_02(t *testing.T) {

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
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(addendStr) "+
      "addendStr='%v'  Error='%v'. ", addendStr, err.Error())
  }

  lenArray := len(bNumStrs)
  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(bNumStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(bNumStrs[i]) "+
        "i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
    }

  }

  result := new(BigIntMathAdd).AddBigIntNumOutputToArray(addendBiNum, bINumArray)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j].GetNumStr() {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j].GetNumStr())
    }

  }

}

func TestBigIntMathAdd_AddBigIntNumOutputToArray_03(t *testing.T) {

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
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(addendStr) "+
      "addendStr='%v'  Error='%v'. ", addendStr, err.Error())
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
    t.Errorf("Error returned by addendBiNum."+
      "SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenArray := len(bNumStrs)
  bINumArray := make([]BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(bNumStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(bNumStrs[i]) "+
        "i='%v'  bNumStrs[i]='%v'  Error='%v'. ", i, bNumStrs[i], err.Error())
    }

  }

  result := new(BigIntMathAdd).AddBigIntNumOutputToArray(addendBiNum, bINumArray)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j].GetNumStr() {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j].GetNumStr())
    }

    actualNumSeps := result[j].GetNumericSeparatorsDto()

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("Error: Expected numSeps='%v' Instead, numSeps='%v'",
        expectedNumSeps.String(), actualNumSeps.String())
    }

  }

}

func TestBigIntMathAdd_AddBigIntNumSeries_01(t *testing.T) {
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
