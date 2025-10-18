package mathops

import (
  "math/big"
  "testing"
)

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_01(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_01"

  radicand := big.NewInt(842567)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-3)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "0.01058762420531197554077114264317"

  expectedPrecisionUint := uint(32)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_02(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_02"

  radicand := big.NewInt(-842567)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-3)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                   1         2         3
  //                          1234567890123456789012345678901234567
  expectedResultNumStr := "-0.01058762420531197554077114264317"

  expectedPrecisionUint := uint(32)

  expectedSignValue := -1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_03(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_03"

  radicand := big.NewInt(9857324656)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-5)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                  1         2         3
  //                         1234567890123456789012345678901234567
  expectedResultNumStr := "0.01002878192918682057341883690632"

  expectedPrecisionUint := uint(32)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_04(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_04"

  radicand := big.NewInt(-9857324656)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-5)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                   1         2         3
  //                          1234567890123456789012345678901234567
  expectedResultNumStr := "-0.01002878192918682057341883690632"

  expectedPrecisionUint := uint(32)

  expectedSignValue := -1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_05(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_05"

  radicand := big.NewInt(25614)

  radicandPrecision := big.NewInt(2)

  nthRoot := big.NewInt(-3)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                  1         2         3
  //                         1234567890123456789012345678901234567
  expectedResultNumStr := "0.15746143256077581282726778712212"

  expectedPrecisionUint := uint(32)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_06(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_06"

  radicand := big.NewInt(-25614)

  radicandPrecision := big.NewInt(2)

  nthRoot := big.NewInt(-3)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                   1         2         3
  //                          1234567890123456789012345678901234567
  expectedResultNumStr := "-0.15746143256077581282726778712212"

  expectedPrecisionUint := uint(32)

  expectedSignValue := -1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_07(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_07"

  radicand := big.NewInt(648912476123)

  radicandPrecision := big.NewInt(3)

  nthRoot := big.NewInt(-9)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                         1234567890123456789012345678901234567
  expectedResultNumStr := "0.1049223981473805598342831898137"

  expectedPrecisionUint := uint(31)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_08(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_08"

  radicand := big.NewInt(-648912476123)

  radicandPrecision := big.NewInt(3)

  nthRoot := big.NewInt(-9)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(31)

  //                                   1         2         3
  //                          1234567890123456789012345678901234567
  expectedResultNumStr := "-0.1049223981473805598342831898137"
  //Qalculate              −0.1049223981473805598342831898137

  expectedPrecisionUint := uint(31)

  expectedSignValue := -1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_09(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_09"

  radicand := big.NewInt(27)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-3)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                  1         2         3
  //                         1234567890123456789012345678901234567
  expectedResultNumStr := "0.33333333333333333333333333333333"

  expectedPrecisionUint := uint(32)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_10(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_10"

  radicand := big.NewInt(59487637924562)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-7)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(32)

  //                                  1         2         3
  //                         1234567890123456789012345678901234567
  expectedResultNumStr := "0.01077022443956376776023853141449"
  // Qalculate             0.010770224439563767760238531414492521

  expectedPrecisionUint := uint(32)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_11(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_11"

  radicand := big.NewInt(1)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-7)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(2)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "1"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_12(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_12"

  radicand := big.NewInt(-1)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-7)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(2)

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedResultNumStr := "-1"

  expectedPrecisionUint := uint(0)

  expectedSignValue := -1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculateNegativeIntegerNthRoot(\n"+
      "   radicand, radicandPrecision, nthRoot, nthRootPrecision,\n"+
      "   maxPrecision\n"+
      "radicand= '%v'\n"+
      "radicandPrecision= '%v'\n"+
      "nthRoot= '%v'\n"+
      "nthRootPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      radicand.Text(10),
      radicandPrecision.Text(10),
      nthRoot.Text(10),
      nthRootPrecision.Text(10),
      maxPrecision.Text(10),
      err.Error())

    return
  }

  resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNum, err := new(BigIntNum).NewBigIntBigPrecision(\n"+
      "  result, resultPrecision)\n"+
      "result= '%v'\n"+
      "resultPrecision= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix,
      result.Text(10),
      resultPrecision.Text(10),
      err.Error())

    return
  }

  err = resultBiNum.IsValid("Validating resultBiNum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBiNum.IsValid('Validating resultBiNum')\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumNumberStr, err := resultBiNum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumNumberStr, err := resultBiNum.GetNumStr()\n"+
      "Error= '%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumPrecisionUint, err := resultBiNum.GetPrecisionUint()\n"+
      "resultBiNum= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  resultBiNumSignValue, err := resultBiNum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBiNumSignValue, err := resultBiNum.GetSign()\n"+
      "resultBiNum= '%v\n"+
      "Error= '%v'\n\n", ePrefix, resultBiNumNumberStr, err.Error())
    return
  }

  if expectedResultNumStr != resultBiNumNumberStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedResultNumStr != resultBiNumNumberStr\n"+
      "Expected resultBiNumNumberStr = '%v'\n"+
      "  Actual resultBiNumNumberStr = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultBiNumNumberStr)

    return
  }

  if expectedPrecisionUint != resultBiNumPrecisionUint {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedPrecisionUint != resultBiNumPrecisionUint\n"+
      "Expected resultBiNumPrecisionUint = '%v'\n"+
      "  Actual resultBiNumPrecisionUint = '%v'\n\n",
      ePrefix, expectedPrecisionUint, resultBiNumPrecisionUint)

    return
  }

  if expectedSignValue != resultBiNumSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedSignValue != resultBiNumSignValue\n"+
      "Expected resultBiNumSignValue = '%v'\n"+
      "  Actual resultBiNumSignValue = '%v'\n\n",
      ePrefix, expectedSignValue, resultBiNumSignValue)

    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_13(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_13"

  radicand := big.NewInt(0)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-7)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(2)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "radicand= 0; nthRoot= -7\n"+
      "radicand= 0 should be an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_14(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculateNegativeIntegerNthRoot_14"

  radicand := big.NewInt(92)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(0)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(2)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Expected Error to be returned due to zero nthRoot. NO ERROR RECEIVED!")
  }

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "radicand= 92; nthRoot= 0\n"+
      "nthRoot= 0 should be an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_01(t *testing.T) {
  radicand := big.NewInt(842567)
  radicandPrecision := big.NewInt(3)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "9.4449895520307989885751143526805"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_02(t *testing.T) {
  radicand := big.NewInt(-842567)
  radicandPrecision := big.NewInt(3)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "-9.4449895520307989885751143526805"

  fdNr := FixedDecimalNthRoot{}.New()

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_03(t *testing.T) {
  radicand := big.NewInt(357)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "7.0939709447507098368744825374903"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_04(t *testing.T) {
  radicand := big.NewInt(-357)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "-7.0939709447507098368744825374903"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_05(t *testing.T) {
  radicand := big.NewInt(16)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(2)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "4"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_06(t *testing.T) {
  radicand := big.NewInt(531441)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(12)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "3"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_07(t *testing.T) {
  radicand := big.NewInt(78123456)
  radicandPrecision := big.NewInt(6)
  nthRoot := big.NewInt(4)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "2.9730030983116531418548508649862"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_08(t *testing.T) {
  radicand := big.NewInt(-78123456)
  radicandPrecision := big.NewInt(6)
  nthRoot := big.NewInt(4)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(31)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Expected an error return from even nthRoot and negative radicand. NO ERROR RETURNED! ")
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_09(t *testing.T) {
  radicand := big.NewInt(-78123456)
  radicandPrecision := big.NewInt(6)
  nthRoot := big.NewInt(5)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(30)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "-2.390871799107522442017212407745"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_10(t *testing.T) {
  radicand := big.NewInt(1)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(5)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(30)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "1"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_11(t *testing.T) {
  radicand := big.NewInt(1)
  radicandPrecision := big.NewInt(5)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.02154434690031883721759293566519"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_12(t *testing.T) {
  radicand := big.NewInt(-1)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "-1"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_13(t *testing.T) {
  radicand := big.NewInt(-1)
  radicandPrecision := big.NewInt(5)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(32)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "-0.02154434690031883721759293566519"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_14(t *testing.T) {
  radicand := big.NewInt(0)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(3)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(32)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "0"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculatePositiveIntegerNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  actualNumStr := resultBiNum.GetNumStr()

  if expectedResult != actualNumStr {
    t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
      expectedResult, actualNumStr)
  }

}

func TestFixedDecimalNthRoot_CalculatePositiveIntegerNthRoot_15(t *testing.T) {
  radicand := big.NewInt(37)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(0)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(32)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculatePositiveIntegerNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Expected error return from nthRoot==0. NO ERROR RETURNED! ")
  }

}
