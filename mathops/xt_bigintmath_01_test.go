package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntMath_ArithmeticGeometricMean_01(t *testing.T) {
  ePrefix := "TestBigIntMath_ArithmeticGeometricMean_01"
  aNum := big.NewInt(24)
  aNumPrecision := big.NewInt(0)
  gNum := big.NewInt(6)
  gNumPrecision := big.NewInt(0)
  maxInternalPrecision := big.NewInt(150)
  targetPrecision := big.NewInt(46)
  expectedValue := "13.4581714817256154207668131569743992430538388544"

  agMean, agMeanPrecision, gValue, gValuePrecision, _, err :=
    new(BigIntMath).ArithmeticGeometricMean(
      aNum,
      aNumPrecision,
      gNum,
      gNumPrecision,
      maxInternalPrecision,
      targetPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "agMean, agMeanPrecision, gValue, gValuePrecision, _, err :=\n"+
      "new(BigIntMath).ArithmeticGeometricMean(aNum, aNumPrecision, gNum,\n"+
      "gNumPrecision, maxInternalPrecision, targetPrecision)\n"+
      "aNumPrecision= '%v'\n"+
      "gNumPrecision= '%v'\n"+
      "maxInternalPrecision= '%v'\n"+
      "targetPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, aNumPrecision, gNumPrecision, maxInternalPrecision, targetPrecision, err.Error())
    return
  }

  binAGMean, err := new(BigIntNum).NewBigIntBigPrecision(agMean, agMeanPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "binAGMean, err := new(BigIntNum).NewBigIntBigPrecision(agMean, agMeanPrecision)\n"+
      "agMean='%v'\n"+
      "agMeanPrecision='%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, agMean.Text(10), agMeanPrecision.Text(10), err.Error())
    return
  }

  binAGMeanNumStr, err := binAGMean.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "binAGMeanNumStr, err := binAGMean.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  binGValue, err := new(BigIntNum).NewBigIntBigPrecision(gValue, gValuePrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "binGValue, err := new(BigIntNum).NewBigIntBigPrecision(gValue, gValuePrecision)\n"+
      "gValue='%v'\n"+
      "gValuePrecision='%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, gValue.Text(10), gValuePrecision.Text(10), err.Error())
    return
  }

  binGValueNumStr, err := binGValue.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "binGValueNumStr, err := binGValue.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedValue != binAGMeanNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected agMean = '%v'\n"+
      "Instead, agMean = '%v'\n\n",
      ePrefix,
      expectedValue,
      binAGMeanNumStr)
    return
  }

  if expectedValue != binGValueNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected gMean = '%v'\n"+
      "Instead, gMean = '%v'\n\n",
      ePrefix,
      expectedValue,
      binGValueNumStr)
  }

  return
}

func TestBigIntMath_GetMagnitude_01(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_01"
  target := big.NewInt(98327123)
  expectedMagnitude := big.NewInt(7)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target='%v'\n"+
      "Error= '%v'\n\n", ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected magnitude = '%v'\n"+
      "Instead, magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_02(t *testing.T) {
  ePrefix := "TestBigIntMath_GetMagnitude_02"
  target := big.NewInt(0)
  expectedMagnitude := big.NewInt(0)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target='%v'\n"+
      "Error= '%v'\n\n", ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected magnitude = '%v'\n"+
      "Instead, magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_03(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_03"
  target := big.NewInt(82)
  expectedMagnitude := big.NewInt(1)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target='%v'\n"+
      "Error= '%v'\n\n", ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected magnitude = '%v'\n"+
      "Instead, magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_04(t *testing.T) {
  ePrefix := "TestBigIntMath_GetMagnitude_04"
  target := big.NewInt(-5)
  expectedMagnitude := big.NewInt(0)
  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target='%v'\n"+
      "Error= '%v'\n\n", ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected magnitude = '%v'\n"+
      "Instead, magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_05(t *testing.T) {
  ePrefix := "TestBigIntMath_GetMagnitude_05"
  target := big.NewInt(2)
  expectedMagnitude := big.NewInt(0)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target='%v'\n"+
      "Error= '%v'\n\n", ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected magnitude = '%v'\n"+
      "Instead, magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_06(t *testing.T) {
  ePrefix := "TestBigIntMath_GetMagnitude_06"
  target := big.NewInt(85652647928)
  expectedMagnitude := big.NewInt(10)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      " magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target='%v'\n"+
      "Error= '%v'\n\n", ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected magnitude = '%v'\n"+
      "Instead, magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_07(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_07"
  numStr := "8565264792812345678901234567890"

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  expectedMagnitude := big.NewInt(30)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_08(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_08"

  numStr := "-8565264792812345678901234567890"

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  expectedMagnitude := big.NewInt(30)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_09(t *testing.T) {
  ePrefix := "TestBigIntMath_GetMagnitude_09"
  numStr := "131072"
  expectedMagnitude := big.NewInt(5)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_10(t *testing.T) {
  ePrefix := "TestBigIntMath_GetMagnitude_10"
  numStr := "131071"
  expectedMagnitude := big.NewInt(5)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_11(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_11"
  numStr := "999999999999"
  expectedMagnitude := big.NewInt(11)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_12(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_12"
  numStr := "1999999999999"
  expectedMagnitude := big.NewInt(12)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_13(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_13"
  numStr := "9999999999999"
  expectedMagnitude := big.NewInt(12)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_14(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_14"
  numStr := "7"
  expectedMagnitude := big.NewInt(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_15(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_15"
  numStr := "8"
  expectedMagnitude := big.NewInt(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_16(t *testing.T) {

  ePrefix := "TestBigIntMath_GetMagnitude_16"
  numStr := "9"
  expectedMagnitude := big.NewInt(0)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_GetMagnitude_17(t *testing.T) {
  ePrefix := "TestBigIntMath_GetMagnitude_17"
  numStr := "10"
  expectedMagnitude := big.NewInt(1)

  fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDec, err := new(BigIntFixedDecimal).NewNumStr(numStr, '.')\n"+
      "numStr='%v'\n"+
      "Error='%v'\n\n", ePrefix, numStr, err.Error())
    return
  }

  fixedDecIntValue, err := fixedDec.GetIntegerValue()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "fixedDecIntValue, err := fixedDec.GetIntegerValue()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  target := big.NewInt(0).Set(fixedDecIntValue)

  magnitude, err := new(BigIntMath).GetMagnitude(target)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(BigIntMath).GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, target.Text(10), err.Error())
    return
  }

  if expectedMagnitude.Cmp(magnitude) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Magnitude = '%v'\n"+
      "Instead, Magnitude = '%v'\n\n",
      ePrefix,
      expectedMagnitude.Text(10),
      magnitude.Text(10))
  }

  return
}

func TestBigIntMath_BigIntPrecisionCmp_01(t *testing.T) {
  ePrefix := "TestBigIntMath_BigIntPrecisionCmp_01"
  num1 := big.NewInt(5)
  num1Precision := big.NewInt(0)
  num2 := big.NewInt(2)
  num2Precision := big.NewInt(0)
  expectedResult := 1

  cmpResult, err :=
    new(BigIntMath).BigIntPrecisionCmp(
      num1,
      num1Precision,
      num2,
      num2Precision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "cmpResult, err := new(BigIntMath).BigIntPrecisionCmp(\n"+
      " num1, num1Precision,  num2, num2Precision)\n"+
      " num1= '%v'\n"+
      "num1Precision= '%v'\n,"+
      "num2= '%v'\n,"+
      "num2Precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1.Text(10),
      num1Precision.Text(10),
      num2.Text(10),
      num2Precision.Text(10),
      err.Error())
    return
  }

  if expectedResult != cmpResult {
    t.Errorf("Error: Expected compare result='%v'. "+
      "Instead, compare result='%v' ",
      expectedResult, cmpResult)
  }

  if expectedResult != cmpResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Expected Compare Result = '%v'\n"+
      "Instead, Compare Result = '%v'\n\n",
      ePrefix,
      expectedResult,
      cmpResult)
  }

  return

}

func TestBigIntMath_BigIntPrecisionCmp_02(t *testing.T) {
  num1 := big.NewInt(52)
  num1Precision := big.NewInt(1)
  num2 := big.NewInt(5105)
  num2Precision := big.NewInt(3)
  expectedResult := 1

  cmpResult :=
    new(BigIntMath).BigIntPrecisionCmp(
      num1,
      num1Precision,
      num2,
      num2Precision)

  if expectedResult != cmpResult {
    t.Errorf("Error: Expected compare result='%v'. "+
      "Instead, compare result='%v' ",
      expectedResult, cmpResult)
  }

}

func TestBigIntMath_BigIntPrecisionCmp_03(t *testing.T) {
  num1 := big.NewInt(-52)
  num1Precision := big.NewInt(1)
  num2 := big.NewInt(5105)
  num2Precision := big.NewInt(3)
  expectedResult := -1

  cmpResult :=
    new(BigIntMath).BigIntPrecisionCmp(
      num1,
      num1Precision,
      num2,
      num2Precision)

  if expectedResult != cmpResult {
    t.Errorf("Error: Expected compare result='%v'. "+
      "Instead, compare result='%v' ",
      expectedResult, cmpResult)
  }

}

func TestBigIntMath_BigIntPrecisionCmp_04(t *testing.T) {
  num1 := big.NewInt(52346789)
  num1Precision := big.NewInt(1)
  num2 := big.NewInt(8234001)
  num2Precision := big.NewInt(3)
  expectedResult := 1

  cmpResult :=
    new(BigIntMath).BigIntPrecisionCmp(
      num1,
      num1Precision,
      num2,
      num2Precision)

  if expectedResult != cmpResult {
    t.Errorf("Error: Expected compare result='%v'. "+
      "Instead, compare result='%v' ",
      expectedResult, cmpResult)
  }

}

func TestBigIntMath_BigIntPrecisionCmp_05(t *testing.T) {
  num1 := big.NewInt(52346789)
  num1Precision := big.NewInt(1)
  num2 := big.NewInt(52346789)
  num2Precision := big.NewInt(1)
  expectedResult := 0

  cmpResult :=
    new(BigIntMath).BigIntPrecisionCmp(
      num1,
      num1Precision,
      num2,
      num2Precision)

  if expectedResult != cmpResult {
    t.Errorf("Error: Expected compare result='%v'. "+
      "Instead, compare result='%v' ",
      expectedResult, cmpResult)
  }

}

func TestBigIntMath_BigIntPrecisionCmp_06(t *testing.T) {
  num1 := big.NewInt(-51)
  num1Precision := big.NewInt(1)
  num2 := big.NewInt(-52)
  num2Precision := big.NewInt(1)
  expectedResult := 1

  cmpResult :=
    new(BigIntMath).BigIntPrecisionCmp(
      num1,
      num1Precision,
      num2,
      num2Precision)

  if expectedResult != cmpResult {
    t.Errorf("Error: Expected compare result='%v'. "+
      "Instead, compare result='%v' ",
      expectedResult, cmpResult)
  }

}

func TestBigIntMath_BigIntPrecisionCmp_07(t *testing.T) {
  num1 := big.NewInt(-51)
  num1Precision := big.NewInt(1)
  num2 := big.NewInt(-51)
  num2Precision := big.NewInt(1)
  expectedResult := 0

  cmpResult :=
    new(BigIntMath).BigIntPrecisionCmp(
      num1,
      num1Precision,
      num2,
      num2Precision)

  if expectedResult != cmpResult {
    t.Errorf("Error: Expected compare result='%v'. "+
      "Instead, compare result='%v' ",
      expectedResult, cmpResult)
  }

}

func TestBigIntMath_RoundToMaxPrecision_01(t *testing.T) {
  num1 := big.NewInt(762939453125)
  num1Precision := big.NewInt(17)
  maxPrecision := big.NewInt(16)
  expectedResult := "0.0000076293945313"

  result, resultPrecision, err :=
    new(BigIntMath).RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_RoundToMaxPrecision_02(t *testing.T) {
  binNum1, err := new(BigIntNum).NewNumStr("276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(29)
  expectedResult := "0.00276213586400995126719079829"

  result, resultPrecision, err :=
    new(BigIntMath).RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_RoundToMaxPrecision_03(t *testing.T) {

  binNum1, err := new(BigIntNum).NewNumStr("276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(30)
  expectedResult := "0.002762135864009951267190798289"

  result, resultPrecision, err :=
    new(BigIntMath).RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_RoundToMaxPrecision_04(t *testing.T) {

  binNum1, err := new(BigIntNum).NewNumStr("-276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "-0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(30)
  expectedResult := "-0.002762135864009951267190798289"

  result, resultPrecision, err :=
    new(BigIntMath).RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_RoundToMaxPrecision_05(t *testing.T) {
  binNum1, err := new(BigIntNum).NewNumStr("-276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "-0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(29)
  expectedResult := "-0.00276213586400995126719079829"

  result, resultPrecision, err :=
    new(BigIntMath).RoundToMaxPrecision(num1, num1Precision, maxPrecision, false)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_RoundToMaxPrecision_06(t *testing.T) {

  num1 := big.NewInt(1230000)

  num1Precision := big.NewInt(5)
  // "12.30000" -> "
  maxPrecision := big.NewInt(3)
  expectedResult := "12.3"

  result, resultPrecision, err :=
    new(BigIntMath).RoundToMaxPrecision(num1, num1Precision, maxPrecision, true)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_RoundToMaxPrecision_07(t *testing.T) {

  num1 := big.NewInt(-1230000)

  num1Precision := big.NewInt(5)
  // "-12.30000" -> "
  maxPrecision := big.NewInt(3)
  expectedResult := "-12.3"

  result, resultPrecision, err :=
    new(BigIntMath).RoundToMaxPrecision(num1, num1Precision, maxPrecision, true)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }
}

func TestBigIntMath_TruncateToMaxPrecision_01(t *testing.T) {
  num1 := big.NewInt(762939453125)
  num1Precision := big.NewInt(17)
  maxPrecision := big.NewInt(16)
  expectedResult := "0.0000076293945312"

  result, resultPrecision, err :=
    new(BigIntMath).TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from Truncate to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_TruncateToMaxPrecision_02(t *testing.T) {
  binNum1, err := new(BigIntNum).NewNumStr("276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(29)
  expectedResult := "0.00276213586400995126719079828"

  result, resultPrecision, err :=
    new(BigIntMath).TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from Truncate to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_TruncateToMaxPrecision_03(t *testing.T) {

  binNum1, err := new(BigIntNum).NewNumStr("276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(30)
  expectedResult := "0.002762135864009951267190798289"

  result, resultPrecision, err :=
    new(BigIntMath).TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from Truncate to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_TruncateToMaxPrecision_04(t *testing.T) {

  binNum1, err := new(BigIntNum).NewNumStr("-276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "-0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(30)
  expectedResult := "-0.002762135864009951267190798289"

  result, resultPrecision, err :=
    new(BigIntMath).TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}

func TestBigIntMath_TruncateToMaxPrecision_05(t *testing.T) {
  binNum1, err := new(BigIntNum).NewNumStr("-276213586400995126719079828947")

  if err != nil {
    t.Errorf("Error returned from binNum1: %v", err.Error())
  }

  num1 := binNum1.GetIntegerValue()

  num1Precision := big.NewInt(32)
  // "-0.00276213586400995126719079828947"
  maxPrecision := big.NewInt(29)
  expectedResult := "-0.00276213586400995126719079828"

  result, resultPrecision, err :=
    new(BigIntMath).TruncateToMaxPrecision(num1, num1Precision, maxPrecision)

  if err != nil {
    t.Errorf("Error returned from Round to Max Precision: %v", err.Error())
  }

  binResult, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned from new(BigIntNum).NewBigIntBigPrecision: %v",
      err.Error())
  }

  if expectedResult != binResult.GetNumStr() {
    t.Errorf("Error: Expected result ='%v'. Instead, result='%v'. ",
      expectedResult, binResult.GetNumStr())
  }

}
