package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntPair_Compare_01(t *testing.T) {

  ePrefix := "TestBigIntPair_Compare_01"
  n1Str := "4"
  n2Str := "-24"
  expectedBig1Compare := 1
  expectedPrecision1Compare := 0
  expectedBig1AbsCompare := -1

  b1, oK := big.NewInt(0).SetString(n1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(n1Str, 10)\n"+
      "n1Str='%v'\n\n", ePrefix, n1Str)
    return
  }

  b2, oK := big.NewInt(0).SetString(n2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(n2Str, 10)\n"+
      "n2Str='%v'\n\n", ePrefix, n2Str)
    return
  }

  bPair, err := new(BigIntPair).NewBase(b1, 0, b2, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBase(b1, 0, b2, 0)\n"+
      "b1= '%v'\nb2= '%v'\nError='%v'\n\n",
      ePrefix, b1.Text(10), b2.Text(10), err.Error())
    return
  }

  if expectedBig1Compare != bPair.Big1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1Compare = '%v'\n"+
      "Instead, bPair.Big1Compare = '%v'\n\n",
      ePrefix, expectedBig1Compare, bPair.Big1Compare)
    return
  }

  if expectedPrecision1Compare != bPair.Precision1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Comparisons don't match.\n"+
      "Expected bPair.Precision1Compare = '%v'\n"+
      "Instead, bPair.Precision1Compare = '%v'\n\n",
      ePrefix, expectedPrecision1Compare, bPair.Precision1Compare)
    return
  }

  if expectedBig1AbsCompare != bPair.Big1AbsCompare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1AbsCompare = '%v'\n"+
      "Instead, bPair.Big1AbsCompare = '%v'\n\n",
      ePrefix, expectedBig1AbsCompare, bPair.Big1AbsCompare)
  }

  return
}

func TestBigIntPair_Compare_02(t *testing.T) {
  ePrefix := "TestBigIntPair_Compare_02"
  n1Str := "-24"
  n2Str := "4"
  expectedBig1Compare := -1
  expectedPrecision1Compare := 0
  expectedBig1AbsCompare := 1

  b1, oK := big.NewInt(0).SetString(n1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(n1Str, 10)\n"+
      "n1Str='%v'\n\n", ePrefix, n1Str)
    return
  }

  b2, oK := big.NewInt(0).SetString(n2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(n2Str, 10)\n"+
      "n2Str='%v'\n\n", ePrefix, n2Str)
    return
  }

  bPair, err := new(BigIntPair).NewBase(b1, 0, b2, 0)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBase(b1, 0, b2, 0)\n"+
      "b1= '%v'\nb2= '%v'\nError='%v'\n\n",
      ePrefix, b1.Text(10), b2.Text(10), err.Error())
    return
  }

  if expectedBig1Compare != bPair.Big1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1Compare = '%v'\n"+
      "Instead, bPair.Big1Compare = '%v'\n\n",
      ePrefix, expectedBig1Compare, bPair.Big1Compare)
    return
  }

  if expectedPrecision1Compare != bPair.Precision1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result! Precision Comparisons don't match.\n"+
      "Expected bPair.Precision1Compare = '%v'\n"+
      "Instead, bPair.Precision1Compare = '%v'\n\n",
      ePrefix, expectedPrecision1Compare, bPair.Precision1Compare)
    return
  }

  if expectedBig1AbsCompare != bPair.Big1AbsCompare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1AbsCompare = '%v'\n"+
      "Instead, bPair.Big1AbsCompare = '%v'\n\n",
      ePrefix, expectedBig1AbsCompare, bPair.Big1AbsCompare)
  }

  return
}

func TestBigIntPair_Compare_03(t *testing.T) {
  ePrefix := "TestBigIntPair_Compare_03"
  n1Str := "40000"
  n1Precision := uint(4)
  n2Str := "240"
  n2Precision := uint(1)

  expectedMaxPrecision := uint(4)
  expectedBig1Compare := 1
  expectedPrecision1Compare := 1
  expectedBig1AbsCompare := 1

  b1, oK := big.NewInt(0).SetString(n1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(n1Str, 10)\n"+
      "n1Str='%v'\n\n", ePrefix, n1Str)
    return
  }

  b2, oK := big.NewInt(0).SetString(n2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(n2Str, 10)\n"+
      "n2Str='%v'\n\n", ePrefix, n2Str)
    return
  }

  bPair, err := new(BigIntPair).NewBase(b1, n1Precision, b2, n2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBase(b1, 0, b2, 0)\n"+
      "b1= '%v'\nb2= '%v'\nError='%v'\n\n",
      ePrefix, b1.Text(10), b2.Text(10), err.Error())
    return
  }

  if expectedBig1Compare != bPair.Big1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1Compare = '%v'\n"+
      "Instead, bPair.Big1Compare = '%v'\n\n",
      ePrefix, expectedBig1Compare, bPair.Big1Compare)
    return
  }

  if expectedPrecision1Compare != bPair.Precision1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result! Precision Comparisons don't match.\n"+
      "Expected bPair.Precision1Compare = '%v'\n"+
      "Instead, bPair.Precision1Compare = '%v'\n\n",
      ePrefix, expectedPrecision1Compare, bPair.Precision1Compare)
    return
  }

  if expectedBig1AbsCompare != bPair.Big1AbsCompare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1AbsCompare = '%v'\n"+
      "Instead, bPair.Big1AbsCompare = '%v'\n\n",
      ePrefix, expectedBig1AbsCompare, bPair.Big1AbsCompare)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n",
      ePrefix, err.Error())
    return
  }

  expectedBig1Compare = -1
  expectedPrecision1Compare = 0
  expectedBig1AbsCompare = -1

  if expectedMaxPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Comparisons don't match.\n"+
      "After Equalizing precision\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, expectedMaxPrecision, bPair.Big2.precision)
    return
  }

  if expectedBig1Compare != bPair.Big1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "After Equalizing precision,\n"+
      "Expected bPair.Big1Compare = '%v'\n"+
      "Instead, bPair.Big1Compare = '%v'\n"+
      "bPair.Big1='%s'\n"+
      "bPair.Big2='%s'\n\n",
      ePrefix, expectedBig1Compare, bPair.Big1Compare,
      bPair.Big1.bigInt.Text(10), bPair.Big2.bigInt.Text(10))
    return
  }

  if expectedBig1AbsCompare != bPair.Big1AbsCompare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "After Equalizing precision,\n"+
      "Expected bPair.Big1AbsCompare = '%v'\n"+
      "Instead, bPair.Big1AbsCompare = '%v'\n"+
      "bPair.Big1='%s'\n"+
      "bPair.Big2='%s'\n\n",
      ePrefix, expectedBig1Compare, bPair.Big1Compare,
      bPair.Big1.bigInt.Text(10), bPair.Big2.bigInt.Text(10))
    return
  }

  if expectedPrecision1Compare != bPair.Precision1Compare {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "After Equalizing precision,\n"+
      "Expected bPair.Precision1Compare = '%v'\n"+
      "Instead, bPair.Precision1Compare = '%v'\n"+
      "bPair.Big1='%s'\n"+
      "bPair.Big2='%s'\n\n",
      ePrefix, expectedPrecision1Compare, bPair.Precision1Compare,
      bPair.Big1.bigInt.Text(10), bPair.Big2.bigInt.Text(10))
  }

  return
}

func TestBigIntPair_NewBase_01(t *testing.T) {

  ePrefix := "TestBigIntPair_NewBase_01"
  //n1Str:="1234567.8901"
  num1Str := "12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := 1
  b1AbsBigInt := big.NewInt(0).Set(b1)

  //n2Str:="7654321.95"
  num2Str := "765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := 1
  b2AbsBigInt := big.NewInt(0).Set(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "76543219500"

  bPair, err := new(BigIntPair).NewBase(b1, b1Precision, b2, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBase(b1, b1Precision, b2, b2Precision)\n"+
      "b1= '%v'\n"+
      "b1Precision= '%v'\n"+
      "b2= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      b1.Text(10),
      b1Precision,
      b2.Text(10),
      b2Precision, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Comparisons don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}
func TestBigIntPair_NewBase_02(t *testing.T) {

  ePrefix := "TestBigIntPair_NewBase_02"

  //n1Str:="-1234567.8901"
  num1Str := "-12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := -1
  b1AbsBigInt := big.NewInt(0).Neg(b1)

  //n2Str:="-7654321.95"
  num2Str := "-765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := -1
  b2AbsBigInt := big.NewInt(0).Neg(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "-76543219500"

  bPair, err := new(BigIntPair).NewBase(b1, b1Precision, b2, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBase(b1, b1Precision, b2, b2Precision)\n"+
      "b1= '%v'\n"+
      "b1Precision= '%v'\n"+
      "b2= '%v'\n"+
      "b2Precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      b1.Text(10),
      b1Precision,
      b2.Text(10),
      b2Precision, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result! Precision Comparisons don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewBigIntNum_01(t *testing.T) {

  ePrefix := "TestBigIntPair_NewBigIntNum_01"

  //n1Str:="1234567.8901"
  num1Str := "12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := 1
  b1AbsBigInt := big.NewInt(0).Set(b1)

  //n2Str:="7654321.95"
  num2Str := "765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := 1
  b2AbsBigInt := big.NewInt(0).Set(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "76543219500"

  b1BigIntNum, err := new(BigIntNum).NewBigInt(b1, b1Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1BigIntNum, err := new(BigIntNum).NewBigInt(b1, b1Precision)\n"+
      "b1='%v'\n"+
      "b1Precision= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, b1.Text(10), b1Precision, err.Error())
    return
  }

  b1BigIntNumStr, err := b1BigIntNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1BigIntNumStr, err := b1BigIntNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  b2BigIntNum, err := new(BigIntNum).NewBigInt(b2, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2BigIntNum, err := new(BigIntNum).NewBigInt(b2, b2Precision)\n"+
      "b2='%v'\n"+
      "b2Precision= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, b2.Text(10), b2Precision, err.Error())
    return
  }

  b2BigIntNumStr, err := b2BigIntNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2BigIntNumStr, err := b2BigIntNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(b1BigIntNum, b2BigIntNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(b1BigIntNum, b2BigIntNum)\n"+
      "b1BigIntNum='%v'\n"+
      "b2BigIntNum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, b1BigIntNumStr, b2BigIntNumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewBigIntNum_02(t *testing.T) {

  ePrefix := "TestBigIntPair_NewBigIntNum_02"
  //n1Str:="-1234567.8901"
  num1Str := "-12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := -1
  b1AbsBigInt := big.NewInt(0).Neg(b1)

  //n2Str:="-7654321.95"
  num2Str := "-765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := -1
  b2AbsBigInt := big.NewInt(0).Neg(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "-76543219500"

  b1BigIntNum, err := new(BigIntNum).NewBigInt(b1, b1Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1BigIntNum, err := new(BigIntNum).NewBigInt(b1, b1Precision)\n"+
      "b1='%v'\n"+
      "b1Precision= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, b1.Text(10), b1Precision, err.Error())
    return
  }

  b1BigIntNumStr, err := b1BigIntNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1BigIntNumStr, err := b1BigIntNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  b2BigIntNum, err := new(BigIntNum).NewBigInt(b2, b2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2BigIntNum, err := new(BigIntNum).NewBigInt(b2, b2Precision)\n"+
      "b2='%v'\n"+
      "b2Precision= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, b2.Text(10), b2Precision, err.Error())
    return
  }

  b2BigIntNumStr, err := b2BigIntNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2BigIntNumStr, err := b2BigIntNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewBigIntNum(b1BigIntNum, b2BigIntNum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewBigIntNum(b1BigIntNum, b2BigIntNum)\n"+
      "b1BigIntNum='%v'\n"+
      "b2BigIntNum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, b1BigIntNumStr, b2BigIntNumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewNumStr_01(t *testing.T) {

  ePrefix := "TestBigIntPair_NewNumStr_01"
  n1Str := "1234567.8901"
  num1Str := "12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := 1
  b1AbsBigInt := big.NewInt(0).Set(b1)

  n2Str := "7654321.95"
  num2Str := "765432195"
  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := 1
  b2AbsBigInt := big.NewInt(0).Set(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "76543219500"

  bPair, err := new(BigIntPair).NewNumStr(n1Str, n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewNumStr(n1Str, n2Str)\n"+
      "n1Str='%v'\n"+
      "n2Str= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, n2Str, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewNumStr_02(t *testing.T) {

  ePrefix := "TestBigIntPair_NewNumStr_02"

  n1Str := "-1234567.8901"
  num1Str := "-12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := -1
  b1AbsBigInt := big.NewInt(0).Neg(b1)

  n2Str := "-7654321.95"
  num2Str := "-765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := -1
  b2AbsBigInt := big.NewInt(0).Neg(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "-76543219500"

  bPair, err := new(BigIntPair).NewNumStr(n1Str, n2Str)

  if err != nil {
    t.Errorf("Error returned by new(BigIntPair).NewNumStr(n1Str, n2Str). "+
      "Error='%v'. ", err.Error())
  }

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewNumStr(n1Str, n2Str)\n"+
      "n1Str='%v'\n"+
      "n2Str= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, n2Str, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewNumStrDto_01(t *testing.T) {

  ePrefix := "TestBigIntPair_NewNumStrDto_01"
  n1Str := "1234567.8901"
  num1Str := "12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := 1
  b1AbsBigInt := big.NewInt(0).Set(b1)

  n2Str := "7654321.95"
  num2Str := "765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := 1
  b2AbsBigInt := big.NewInt(0).Set(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "76543219500"

  nx1Dto, err := new(NumStrDto).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx1Dto, err := new(NumStrDto).NewNumStr(n1Str)\n"+
      "n1Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, err.Error())
    return
  }

  nx1DtoNumStr, err := nx1Dto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx1DtoNumStr, err := nx1Dto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nx2Dto, err := new(NumStrDto).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx1Dto, err := new(NumStrDto).NewNumStr(n2Str)\n"+
      "n2Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n2Str, err.Error())
    return
  }

  nx2DtoNumStr, err := nx2Dto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx2DtoNumStr, err := nx2Dto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewNumStrDto(nx1Dto, nx2Dto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewNumStrDto(nx1Dto, nx2Dto)\n"+
      "nx1Dto='%v'\n"+
      "nx2Dto= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, nx1DtoNumStr, nx2DtoNumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewNumStrDto_02(t *testing.T) {

  ePrefix := "TestBigIntPair_NewNumStrDto_02"
  n1Str := "-1234567.8901"
  num1Str := "-12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := -1
  b1AbsBigInt := big.NewInt(0).Neg(b1)

  n2Str := "-7654321.95"
  num2Str := "-765432195"
  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := -1
  b2AbsBigInt := big.NewInt(0).Neg(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "-76543219500"

  nx1Dto, err := new(NumStrDto).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx1Dto, err := new(NumStrDto).NewNumStr(n1Str)\n"+
      "n1Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, err.Error())
    return
  }

  nx1DtoNumStr, err := nx1Dto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx1DtoNumStr, err := nx1Dto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  nx2Dto, err := new(NumStrDto).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx1Dto, err := new(NumStrDto).NewNumStr(n2Str)\n"+
      "n2Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n2Str, err.Error())
    return
  }

  nx2DtoNumStr, err := nx2Dto.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "nx2DtoNumStr, err := nx2Dto.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewNumStrDto(nx1Dto, nx2Dto)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewNumStrDto(nx1Dto, nx2Dto)\n"+
      "nx1Dto='%v'\n"+
      "nx2Dto= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, nx1DtoNumStr, nx2DtoNumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewIntAry_01(t *testing.T) {
  ePrefix := "TestBigIntPair_NewIntAry_01"
  n1Str := "1234567.8901"
  num1Str := "12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := 1
  b1AbsBigInt := big.NewInt(0).Set(b1)

  n2Str := "7654321.95"
  num2Str := "765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := 1
  b2AbsBigInt := big.NewInt(0).Set(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "76543219500"

  ia1, err := new(IntAry).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia1, err := new(IntAry).NewNumStr(n1Str)\n"+
      "n1Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, err.Error())
    return
  }

  ia1NumStr, err := ia1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia1NumStr, err := ia1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  ia2, err := new(IntAry).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia2, err := new(IntAry).NewNumStr(n2Str)\n"+
      "n2Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n2Str, err.Error())
    return
  }

  ia2NumStr, err := ia2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia2NumStr, err := ia2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewIntAry(ia1, ia2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewIntAry(ia1, ia2)\n"+
      "ia1='%v'\n"+
      "ia2= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, ia1NumStr, ia2NumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_NewIntAry_02(t *testing.T) {

  ePrefix := "TestBigIntPair_NewIntAry_02"
  n1Str := "-1234567.8901"
  num1Str := "-12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := -1
  b1AbsBigInt := big.NewInt(0).Neg(b1)

  n2Str := "-7654321.95"
  num2Str := "-765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := -1
  b2AbsBigInt := big.NewInt(0).Neg(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "-76543219500"

  ia1, err := new(IntAry).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia1, err := new(IntAry).NewNumStr(n1Str)\n"+
      "n1Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, err.Error())
    return
  }

  ia1NumStr, err := ia1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia1NumStr, err := ia1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  ia2, err := new(IntAry).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia2, err := new(IntAry).NewNumStr(n2Str)\n"+
      "n2Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n2Str, err.Error())
    return
  }

  ia2NumStr, err := ia2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "ia2NumStr, err := ia2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewIntAry(ia1, ia2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewIntAry(ia1, ia2)\n"+
      "ia1='%v'\n"+
      "ia2= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, ia1NumStr, ia2NumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_Decimal_01(t *testing.T) {

  ePrefix := "TestBigIntPair_Decimal_01"

  n1Str := "1234567.8901"
  num1Str := "12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := 1
  b1AbsBigInt := big.NewInt(0).Set(b1)

  n2Str := "7654321.95"
  num2Str := "765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := 1
  b2AbsBigInt := big.NewInt(0).Set(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "76543219500"

  dec1, err := new(Decimal).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1, err := new(Decimal).NewNumStr(n1Str)\n"+
      "n1Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, err.Error())
    return
  }

  dec1NumStr, err := dec1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1NumStr, err := dec1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2, err := new(Decimal).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2, err := new(Decimal).NewNumStr(n2Str)\n"+
      "n2Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n2Str, err.Error())
    return
  }

  dec2NumStr, err := dec2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2NumStr, err := dec2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewDecimal(dec1, dec2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewDecimal(dec1, dec2)\n"+
      "dec1='%v'\n"+
      "dec2= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, dec1NumStr, dec2NumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}

func TestBigIntPair_Decimal_02(t *testing.T) {

  ePrefix := "TestBigIntPair_Decimal_02"
  n1Str := "-1234567.8901"
  num1Str := "-12345678901"

  b1, oK := big.NewInt(0).SetString(num1Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b1, oK := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str='%v'\n\n", ePrefix, num1Str)
    return
  }

  b1Precision := uint(4)
  b1Sign := -1
  b1AbsBigInt := big.NewInt(0).Neg(b1)

  n2Str := "-7654321.95"
  num2Str := "-765432195"

  b2, oK := big.NewInt(0).SetString(num2Str, 10)

  if !oK {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "b2, oK := big.NewInt(0).SetString(num2Str, 10)\n"+
      "num2Str='%v'\n\n", ePrefix, num2Str)
    return
  }

  b2Precision := uint(2)
  b2Sign := -1
  b2AbsBigInt := big.NewInt(0).Neg(b2)

  reconciledPrecision := uint(4)

  n2StrReconciled := "-76543219500"

  dec1, err := new(Decimal).NewNumStr(n1Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1, err := new(Decimal).NewNumStr(n1Str)\n"+
      "n1Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n1Str, err.Error())
    return
  }

  dec1NumStr, err := dec1.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec1NumStr, err := dec1.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  dec2, err := new(Decimal).NewNumStr(n2Str)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2, err := new(Decimal).NewNumStr(n2Str)\n"+
      "n2Str='%v'\n"+
      "Error= '%v'\n\n", ePrefix, n2Str, err.Error())
    return
  }

  dec2NumStr, err := dec2.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "dec2NumStr, err := dec2.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  bPair, err := new(BigIntPair).NewDecimal(dec1, dec2)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(BigIntPair).NewDecimal(dec1, dec2)\n"+
      "dec1='%v'\n"+
      "dec2= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, dec1NumStr, dec2NumStr, err.Error())
    return
  }

  if bPair.Big1.bigInt.Cmp(b1) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big1.bigInt = '%v'\n"+
      "Instead, bPair.Big1.bigInt = '%v'\n\n",
      ePrefix,
      b1.Text(10),
      bPair.Big1.bigInt.Text(10))
    return
  }

  if b1AbsBigInt.Cmp(bPair.Big1.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Comparisons don't match.\n"+
      "Expected bPair.Big1.absBigInt = '%v'\n"+
      "Instead, bPair.Big1.absBigInt = '%v'\n\n",
      ePrefix,
      b1AbsBigInt.Text(10),
      bPair.Big1.absBigInt.Text(10))
    return
  }

  if b1Precision != bPair.Big1.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big1.precision = '%v'\n"+
      "Instead, bPair.Big1.precision = '%v'\n\n",
      ePrefix, b1Precision, bPair.Big1.precision)
    return
  }

  if b1Sign != bPair.Big1.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big1.sign = '%v'\n"+
      "Instead, bPair.Big1.sign = '%v'\n\n",
      ePrefix, b1Sign, bPair.Big1.sign)
    return
  }

  if bPair.Big2.bigInt.Cmp(b2) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected bPair.Big2.bigInt = '%v'\n"+
      "Instead, bPair.Big2.bigInt = '%v'\n\n",
      ePrefix,
      b2.Text(10),
      bPair.Big2.bigInt.Text(10))
    return
  }

  if b2AbsBigInt.Cmp(bPair.Big2.absBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Absolute Values don't match.\n"+
      "Expected bPair.Big2.absBigInt = '%v'\n"+
      "Instead, bPair.Big2.absBigInt = '%v'\n\n",
      ePrefix,
      b2AbsBigInt.Text(10),
      bPair.Big2.absBigInt.Text(10))
    return
  }

  if b2Precision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, b2Precision, bPair.Big2.precision)
    return
  }

  if b2Sign != bPair.Big2.sign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Sign Values don't match.\n"+
      "Expected bPair.Big2.sign = '%v'\n"+
      "Instead, bPair.Big2.sign = '%v'\n\n",
      ePrefix, b2Sign, bPair.Big2.sign)
    return
  }

  err = bPair.MakePrecisionsEqual()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = bPair.MakePrecisionsEqual()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  if reconciledPrecision != bPair.Big2.precision {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Reconciled Precision Values don't match.\n"+
      "Expected bPair.Big2.precision = '%v'\n"+
      "Instead, bPair.Big2.precision = '%v'\n\n",
      ePrefix, reconciledPrecision, bPair.Big2.precision)
    return
  }

  actualBig2Numstr := bPair.Big2.bigInt.Text(10)

  if n2StrReconciled != actualBig2Numstr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Number Strings don't match.\n"+
      "Expected actualBig2Numstr = '%v'\n"+
      "Instead, actualBig2Numstr = '%v'\n\n",
      ePrefix, n2StrReconciled, actualBig2Numstr)
  }

  return
}
