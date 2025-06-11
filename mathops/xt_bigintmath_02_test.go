package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntMath_TruncateTrailingFractionalZeros_01(t *testing.T) {

  ePrefix := "TestBigIntMath_TruncateTrailingFractionalZeros_01"
  num := big.NewInt(56123456000)
  numPrecision := big.NewInt(9)

  expectedNum := big.NewInt(56123456)
  expectedNumPrecision := big.NewInt(6)

  actualNum, actualNumPrecision, err :=
    new(BigIntMath).TruncateTrailingFractionalZeros(num, numPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNum, actualNumPrecision, err := new(BigIntMath).\n"+
      "  TruncateTrailingFractionalZeros( num, numPrecision)\n"+
      "num= '%v'\n"+
      "numPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num,
      numPrecision,
      err.Error())
    return
  }

  if expectedNum.Cmp(actualNum) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNum = '%v'\n"+
      "Instead, actualNum = '%v'\n\n",
      ePrefix,
      expectedNum.Text(10),
      actualNum.Text(10))
    return
  }

  if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNumPrecision = '%v'\n"+
      "Instead, actualNumPrecision = '%v'\n\n",
      ePrefix,
      expectedNumPrecision.Text(10),
      actualNumPrecision.Text(10))
  }

  return
}

func TestBigIntMath_TruncateTrailingFractionalZeros_02(t *testing.T) {

  ePrefix := "TestBigIntMath_TruncateTrailingFractionalZeros_02"
  num := big.NewInt(-56123456000)
  numPrecision := big.NewInt(9)

  expectedNum := big.NewInt(-56123456)
  expectedNumPrecision := big.NewInt(6)

  actualNum, actualNumPrecision, err :=
    new(BigIntMath).TruncateTrailingFractionalZeros(num, numPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNum, actualNumPrecision, err := new(BigIntMath).\n"+
      "  TruncateTrailingFractionalZeros( num, numPrecision)\n"+
      "num= '%v'\n"+
      "numPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num,
      numPrecision,
      err.Error())
    return
  }

  if expectedNum.Cmp(actualNum) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNum = '%v'\n"+
      "Instead, actualNum = '%v'\n\n",
      ePrefix,
      expectedNum.Text(10),
      actualNum.Text(10))
    return
  }

  if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNumPrecision = '%v'\n"+
      "Instead, actualNumPrecision = '%v'\n\n",
      ePrefix,
      expectedNumPrecision.Text(10),
      actualNumPrecision.Text(10))
  }

  return
}

func TestBigIntMath_TruncateTrailingFractionalZeros_03(t *testing.T) {

  ePrefix := "TestBigIntMath_TruncateTrailingFractionalZeros_03"
  num := big.NewInt(56123456)
  numPrecision := big.NewInt(6)

  expectedNum := big.NewInt(56123456)
  expectedNumPrecision := big.NewInt(6)

  actualNum, actualNumPrecision, err :=
    new(BigIntMath).TruncateTrailingFractionalZeros(num, numPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNum, actualNumPrecision, err := new(BigIntMath).\n"+
      "  TruncateTrailingFractionalZeros( num, numPrecision)\n"+
      "num= '%v'\n"+
      "numPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num,
      numPrecision,
      err.Error())
    return
  }

  if expectedNum.Cmp(actualNum) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNum = '%v'\n"+
      "Instead, actualNum = '%v'\n\n",
      ePrefix,
      expectedNum.Text(10),
      actualNum.Text(10))
    return
  }

  if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNumPrecision = '%v'\n"+
      "Instead, actualNumPrecision = '%v'\n\n",
      ePrefix,
      expectedNumPrecision.Text(10),
      actualNumPrecision.Text(10))
  }

  return
}

func TestBigIntMath_TruncateTrailingFractionalZeros_04(t *testing.T) {

  ePrefix := "TestBigIntMath_TruncateTrailingFractionalZeros_04"
  num := big.NewInt(-56123456)
  numPrecision := big.NewInt(6)

  expectedNum := big.NewInt(-56123456)
  expectedNumPrecision := big.NewInt(6)

  actualNum, actualNumPrecision, err :=
    new(BigIntMath).TruncateTrailingFractionalZeros(num, numPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNum, actualNumPrecision, err := new(BigIntMath).\n"+
      "  TruncateTrailingFractionalZeros( num, numPrecision)\n"+
      "num= '%v'\n"+
      "numPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num,
      numPrecision,
      err.Error())
    return
  }

  if expectedNum.Cmp(actualNum) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNum = '%v'\n"+
      "Instead, actualNum = '%v'\n\n",
      ePrefix,
      expectedNum.Text(10),
      actualNum.Text(10))
    return
  }

  if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNumPrecision = '%v'\n"+
      "Instead, actualNumPrecision = '%v'\n\n",
      ePrefix,
      expectedNumPrecision.Text(10),
      actualNumPrecision.Text(10))
  }

  return
}

func TestBigIntMath_TruncateTrailingFractionalZeros_05(t *testing.T) {

  ePrefix := "TestBigIntMath_TruncateTrailingFractionalZeros_05"
  num := big.NewInt(0)
  numPrecision := big.NewInt(6)

  expectedNum := big.NewInt(0)
  expectedNumPrecision := big.NewInt(0)

  actualNum, actualNumPrecision, err :=
    new(BigIntMath).TruncateTrailingFractionalZeros(num, numPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNum, actualNumPrecision, err := new(BigIntMath).\n"+
      "  TruncateTrailingFractionalZeros( num, numPrecision)\n"+
      "num= '%v'\n"+
      "numPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num,
      numPrecision,
      err.Error())
    return
  }

  if expectedNum.Cmp(actualNum) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNum = '%v'\n"+
      "Instead, actualNum = '%v'\n\n",
      ePrefix,
      expectedNum.Text(10),
      actualNum.Text(10))
    return
  }

  if expectedNumPrecision.Cmp(actualNumPrecision) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected actualNumPrecision = '%v'\n"+
      "Instead, actualNumPrecision = '%v'\n\n",
      ePrefix,
      expectedNumPrecision.Text(10),
      actualNumPrecision.Text(10))
  }

  return
}
