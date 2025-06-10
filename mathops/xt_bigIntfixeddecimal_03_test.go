package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntFixedDecimal_MultiplyByTenToPwr_01(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_MultiplyByTenToPwr_01"
  expectedNumStr := "105.6752"
  num := 1056752
  precision := uint(4)
  exponent := uint(0)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_02(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_MultiplyByTenToPwr_02"
  expectedNumStr := "-105.6752"
  num := -1056752
  precision := uint(4)
  exponent := uint(0)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_03(t *testing.T) {

  ePrefix := "fixedDec"
  expectedNumStr := "1056.752"
  num := 1056752
  precision := uint(4)
  exponent := uint(1)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_04(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_MultiplyByTenToPwr_04"
  expectedNumStr := "-1056.752"
  num := -1056752
  precision := uint(4)
  exponent := uint(1)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_05(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_MultiplyByTenToPwr_05"
  expectedNumStr := "10567.52"
  num := 1056752
  precision := uint(4)
  exponent := uint(2)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_06(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_MultiplyByTenToPwr_06"
  expectedNumStr := "-10567.52"
  num := -1056752
  precision := uint(4)
  exponent := uint(2)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_07(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_MultiplyByTenToPwr_07"
  expectedNumStr := "10567520000"
  num := 1056752
  precision := uint(4)
  exponent := uint(8)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTenToPwr_08(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_MultiplyByTenToPwr_08"
  expectedNumStr := "0"
  num := 0
  precision := uint(4)
  exponent := uint(5)

  fixedDec := new(BigIntFixedDecimal).NewInt(num, precision)

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.MultiplyByTenToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec.MultiplyByTenToPower(exponent)\n"+
      "fixedDecNumStr= '%v\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, fixedDecNumStr, exponent, err.Error())
    return
  }

  fixedDecNumStr2, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr2, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr2 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr2 = '%v'\n"+
      "Instead, fixedDecNumStr2 = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr2)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_01(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_MultiplyByTwoToPower_01"
  // multiplicand = 23.321
  multiplicandBInt := big.NewInt(23321)
  multiplicandPrecision := uint(3)
  exponent := uint(5)
  expectedResult := "746.272"

  fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision ='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandBInt.Text(10), multiplicandPrecision, err.Error())
    return
  }

  err = fixedDec.MultiplyByTwoToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fixedDec.MultiplyByTwoToPower(exponent)\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n", ePrefix, exponent, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedResult, fixedDecNumStr)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_02(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_MultiplyByTwoToPower_02"
  // multiplicand = 8
  multiplicandBInt := big.NewInt(8)
  multiplicandPrecision := uint(0)
  exponent := uint(10)
  expectedResult := "8192"

  fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision ='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandBInt.Text(10), multiplicandPrecision, err.Error())
    return
  }

  err = fixedDec.MultiplyByTwoToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fixedDec.MultiplyByTwoToPower(exponent)\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n", ePrefix, exponent, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedResult, fixedDecNumStr)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_03(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_MultiplyByTwoToPower_03"
  // multiplicand = 9.871234
  multiplicandBInt := big.NewInt(9871234)
  multiplicandPrecision := uint(6)
  exponent := uint(1)
  expectedResult := "19.742468"

  fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision ='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandBInt.Text(10), multiplicandPrecision, err.Error())
    return
  }

  err = fixedDec.MultiplyByTwoToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fixedDec.MultiplyByTwoToPower(exponent)\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n", ePrefix, exponent, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedResult, fixedDecNumStr)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_04(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_MultiplyByTwoToPower_04"
  // multiplicand = -9.871234
  multiplicandBInt := big.NewInt(-9871234)
  multiplicandPrecision := uint(6)
  exponent := uint(3)
  expectedResult := "-78.969872"

  fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision ='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandBInt.Text(10), multiplicandPrecision, err.Error())
    return
  }

  err = fixedDec.MultiplyByTwoToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fixedDec.MultiplyByTwoToPower(exponent)\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n", ePrefix, exponent, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedResult, fixedDecNumStr)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_05(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_MultiplyByTwoToPower_05"
  // multiplicand = 8
  multiplicandBInt := big.NewInt(8)
  multiplicandPrecision := uint(0)
  exponent := uint(0)
  expectedResult := "8"

  fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision ='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandBInt.Text(10), multiplicandPrecision, err.Error())
    return
  }

  err = fixedDec.MultiplyByTwoToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fixedDec.MultiplyByTwoToPower(exponent)\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n", ePrefix, exponent, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedResult, fixedDecNumStr)
  }

  return
}

func TestBigIntFixedDecimal_MultiplyByTwoToPower_06(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_MultiplyByTwoToPower_06"

  // (0.12345 x 2^15 = 4045.2096)
  multiplicandBInt := big.NewInt(12345)
  multiplicandPrecision := uint(5)
  exponent := uint(15)
  expectedResult := "4045.2096"

  fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).New(multiplicandBInt, multiplicandPrecision)\n"+
      "multiplicandBInt= '%v'\n"+
      "multiplicandPrecision ='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplicandBInt.Text(10), multiplicandPrecision, err.Error())
    return
  }

  err = fixedDec.MultiplyByTwoToPower(exponent)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fixedDec.MultiplyByTwoToPower(exponent)\n"+
      "exponent='%v'\n"+
      "Error='%v'\n\n", ePrefix, exponent, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedResult != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedResult, fixedDecNumStr)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_01(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_01"
  numStr := "89765.123456789012"
  expectedNumStr := numStr
  expectedPrecision := uint(12)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_02(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_NewNumStr_02"
  numStr := "-89765.123456789012"
  expectedNumStr := numStr
  expectedPrecision := uint(12)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_03(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_NewNumStr_03"
  numStr := "0.123456789012"
  expectedNumStr := numStr
  expectedPrecision := uint(12)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_04(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_04"
  numStr := ".123456789012"
  expectedNumStr := "0.123456789012"
  expectedPrecision := uint(12)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_05(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_NewNumStr_05"
  numStr := "-.123456789012"
  expectedNumStr := "-0.123456789012"
  expectedPrecision := uint(12)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_06(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_06"
  numStr := "10"
  expectedNumStr := "10"
  expectedPrecision := uint(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_07(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_07"
  numStr := "-52"
  expectedNumStr := "-52"
  expectedPrecision := uint(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_08(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_08"
  numStr := "-00052.1234"
  expectedNumStr := "-52.1234"
  expectedPrecision := uint(4)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_09(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_NewNumStr_09"
  numStr := "(00052.1234)"
  expectedNumStr := "-52.1234"
  expectedPrecision := uint(4)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_10(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_10"
  numStr := "(00052.1234"
  expectedNumStr := "52.1234"
  expectedPrecision := uint(4)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_11(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_11"
  numStr := "+00052.1234"
  expectedNumStr := "52.1234"
  expectedPrecision := uint(4)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_12(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_NewNumStr_12"
  numStr := "-00052.1234567890123456"
  expectedNumStr := "-52.1234567890123456"
  expectedPrecision := uint(16)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_13(t *testing.T) {
  ePrefix := "TestBigIntFixedDecimal_NewNumStr_13"
  numStr := "52 . 123 4567 8901 23456"
  expectedNumStr := "52.1234567890123456"
  expectedPrecision := uint(16)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_14(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_14"
  numStr := "5 2"
  expectedNumStr := "52"
  expectedPrecision := uint(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_15(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_15"
  numStr := "    (52)     "
  expectedNumStr := "-52"
  expectedPrecision := uint(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_16(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_16"
  numStr := "    (52)    1234567 "
  expectedNumStr := "-52"
  expectedPrecision := uint(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_17(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_17"
  expectedPrecision := uint(1024)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(EulersNum50kStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = fixedDec.RoundToDecPlace(expectedPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = fixedDec.RoundToDecPlace(expectedPrecision)\n"+
      "expectedPrecision= '%v'\nError='%v'\n\n",
      ePrefix, expectedPrecision, err.Error())
    return
  }

  fixedDecNumStr, err := fixedDec.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecNumStr, err := fixedDec.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  fdEuler1K := GetEulersNum1k()

  fdEuler1KNumStr, err := fdEuler1K.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fdEuler1KNumStr, err := fdEuler1K.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if fdEuler1KNumStr != fixedDecNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecNumStr = '%v'\n"+
      "Instead, fixedDecNumStr = '%v'\n\n",
      ePrefix, fdEuler1KNumStr, fixedDecNumStr)
    return
  }

  fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecPrecisionUint, err := fixedDec.GetPrecisionUint()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedPrecision != fixedDecPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected fixedDecPrecisionUint = '%v'\n"+
      "Instead, fixedDecPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecision, fixedDecPrecisionUint)
  }

  return
}

func TestBigIntFixedDecimal_NewNumStr_18(t *testing.T) {

  ePrefix := "TestBigIntFixedDecimal_NewNumStr_18"
  numStr := "abcdefghijklmnop"

  _, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error from INVALID Number String. NO ERROR RETURNED!",
      ePrefix)
  }
}
