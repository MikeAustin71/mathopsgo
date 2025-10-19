package mathops

import (
  "math/big"
  "testing"
)

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_01(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_01"

  radicand := big.NewInt(4567)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(666)

  nthRootPrecision := big.NewInt(3)

  maxPrecision := big.NewInt(26)

  //                                       1         2         3
  //                              1234567890123456789012345678901234567
  expectedResultNumStr := "312565.80127327376619628998765762"

  expectedPrecisionUint := uint(26)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_02(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_02"

  radicand := big.NewInt(45123)

  radicandPrecision := big.NewInt(3)

  nthRoot := big.NewInt(25)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "4.5894346095427589142057717684748"

  expectedPrecisionUint := uint(31)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_03(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_03"

  radicand := big.NewInt(-45123)

  radicandPrecision := big.NewInt(3)

  nthRoot := big.NewInt(25)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "4.5894346095427589142057717684748"

  expectedPrecisionUint := uint(31)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_04(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_04"

  radicand := big.NewInt(659)

  radicandPrecision := big.NewInt(1)

  nthRoot := big.NewInt(251)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "1.1815865924660052902023955150749"

  expectedPrecisionUint := uint(31)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_05(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_05"

  radicand := big.NewInt(679)

  radicandPrecision := big.NewInt(1)

  nthRoot := big.NewInt(353)

  nthRootPrecision := big.NewInt(2)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "3.3032639583686665394923710881736"

  expectedPrecisionUint := uint(31)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_06(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_06"

  radicand := big.NewInt(-679)

  radicandPrecision := big.NewInt(1)

  nthRoot := big.NewInt(353)

  nthRootPrecision := big.NewInt(2)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "3.3032639583686665394923710881736"

  expectedPrecisionUint := uint(31)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_07(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_07"

  radicand := big.NewInt(9645321)

  radicandPrecision := big.NewInt(1)

  nthRoot := big.NewInt(912)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(30)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "1.163101209960992659162495181682"

  expectedPrecisionUint := uint(30)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_08(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_08"

  radicand := big.NewInt(-9645321)

  radicandPrecision := big.NewInt(1)

  nthRoot := big.NewInt(9124)

  nthRootPrecision := big.NewInt(2)

  maxPrecision := big.NewInt(31)

  //                                   1         2         3
  //                        0.1234567890123456789012345678901234567
  expectedResultNumStr := "-1.1630241704961631471493809157491"

  expectedPrecisionUint := uint(31)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_09(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_09"

  radicand := big.NewInt(0)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(52)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResultNumStr := "0"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_10(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_10"

  radicand := big.NewInt(1)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(52)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "1"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_11(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_11"

  radicand := big.NewInt(-1)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(52)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  //                                  1         2         3
  //                       0.1234567890123456789012345678901234567
  expectedResultNumStr := "1"

  expectedPrecisionUint := uint(0)

  expectedSignValue := 1

  fixDecNthRoot := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fixDecNthRoot.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, resultPrecision, err :=\n"+
      "  fixDecNthRoot.CalculatePositiveFractionalNthRoot(\n"+
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

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_12(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_12"

  radicand := big.NewInt(15)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(0)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "radicand= 15; nthRoot= 0.0\n"+
      "nthroot= 0 should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_13(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_13"

  radicand := big.NewInt(15)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(0)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(31)

  _, _, err :=
    new(FixedDecimalNthRoot).CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "radicand= 15; nthRoot= 0\n"+
      "nthroot= 0 should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_14(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_14"

  radicand := big.NewInt(15)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(-67)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  _, _, err :=
    new(FixedDecimalNthRoot).CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "radicand= 15; nthRoot= -6.7\n"+
      "Negative nthroot should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_15(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_15"

  radicand := big.NewInt(15)

  radicandPrecision := big.NewInt(0)

  nthRoot := big.NewInt(12)

  nthRootPrecision := big.NewInt(0)

  maxPrecision := big.NewInt(31)

  _, _, err :=
    new(FixedDecimalNthRoot).CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Error: Expected error return from integer nthroot. Instead, NO ERROR RETURNED! ")
  }

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "radicand= 15; nthRoot= 12\n"+
      "Integer nthroot should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_16(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_16"

  radicand := big.NewInt(15)

  radicandPrecision := big.NewInt(-1)

  nthRoot := big.NewInt(12)

  nthRootPrecision := big.NewInt(1)

  maxPrecision := big.NewInt(31)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Error: Expected error return from negative radicand precision. Instead, NO ERROR RETURNED! ")
  }

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "radicandPrecision= -1; nthRoot= 1.2\n"+
      "negative radicand precision should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_17(t *testing.T) {

  ePrefix := "TestFixedDecimalNthRoot_CalculatePositiveFractionalNthRoot_17"

  radicand := big.NewInt(15)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(12)
  nthRootPrecision := big.NewInt(-2)
  maxPrecision := big.NewInt(31)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculatePositiveFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Errorf("%v\n"+
      "Expected an Error, BUT NO ERROR WAS RETURNED!\n"+
      "nthRootPrecision= -2\n"+
      "negative nthRoot precision should produce an error.\n\n", ePrefix)
    return
  }

  return
}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_01(t *testing.T) {

  radicand := big.NewInt(38432)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(-372)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.05855387604340335560518923124485"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_02(t *testing.T) {

  radicand := big.NewInt(-38432)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(-372)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                             1         2         3
  //                    1234567890123456789012345678901234567
  expectedResult := "-0.05855387604340335560518923124485"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_03(t *testing.T) {

  radicand := big.NewInt(945384)
  radicandPrecision := big.NewInt(5)
  nthRoot := big.NewInt(-5723)
  nthRootPrecision := big.NewInt(3)
  maxPrecision := big.NewInt(31)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.6753494112877003591875450895632"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_04(t *testing.T) {

  radicand := big.NewInt(-945384)
  radicandPrecision := big.NewInt(5)
  nthRoot := big.NewInt(-5723)
  nthRootPrecision := big.NewInt(3)
  maxPrecision := big.NewInt(31)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.6753494112877003591875450895632"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_05(t *testing.T) {

  radicand := big.NewInt(64)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(-22)
  nthRootPrecision := big.NewInt(1)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.15101118055055591416115527630953"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_06(t *testing.T) {

  radicand := big.NewInt(-6475)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(-382)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.10052943765859122324489541725654"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_07(t *testing.T) {

  radicand := big.NewInt(-5291)
  radicandPrecision := big.NewInt(2)
  nthRoot := big.NewInt(-389)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.36052149780869144809980868847077"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_08(t *testing.T) {

  radicand := big.NewInt(5291)
  radicandPrecision := big.NewInt(2)
  nthRoot := big.NewInt(-389)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0.36052149780869144809980868847077"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_09(t *testing.T) {

  radicand := big.NewInt(-5291)
  radicandPrecision := big.NewInt(2)
  nthRoot := big.NewInt(-2)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_10(t *testing.T) {

  radicand := big.NewInt(0)
  radicandPrecision := big.NewInt(3)
  nthRoot := big.NewInt(-389)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "0"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_11(t *testing.T) {

  radicand := big.NewInt(1)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(-389)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "1"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_12(t *testing.T) {

  radicand := big.NewInt(-1)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(-389)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  //                            1         2         3
  //                   1234567890123456789012345678901234567
  expectedResult := "1"

  fdNr := FixedDecimalNthRoot{}

  result, resultPrecision, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err != nil {
    t.Errorf("Error returned by fdNr.CalculateNegativeFractionalNthRoot(...) "+
      "Error='%v' ", err.Error())
  }

  resultBiNum, err := BigIntNum{}.NewBigIntBigPrecision(result, resultPrecision)

  if err != nil {
    t.Errorf("Error returned by .BigIntNum{}.NewBigIntBigPrecision(...) "+
      "Error='%v' ", err.Error())
  }

  if expectedResult != resultBiNum.GetNumStr() {
    t.Errorf("Error: expected result='%v'. Instead, result='%v'",
      expectedResult, resultBiNum.GetNumStr())
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_13(t *testing.T) {

  radicand := big.NewInt(58971)
  radicandPrecision := big.NewInt(-1)
  nthRoot := big.NewInt(-389)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Error: Expected error return due to negative radicand precision. NO ERROR RETURNED!")
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_14(t *testing.T) {

  radicand := big.NewInt(58971)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(-389)
  nthRootPrecision := big.NewInt(-2)
  maxPrecision := big.NewInt(32)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Error: Expected error return due to negative nthRoot precision. NO ERROR RETURNED!")
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_15(t *testing.T) {

  radicand := big.NewInt(58971)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(0)
  nthRootPrecision := big.NewInt(2)
  maxPrecision := big.NewInt(32)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Error: Expected error return due to zero nthRoot. NO ERROR RETURNED!")
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_16(t *testing.T) {

  radicand := big.NewInt(58971)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(52)
  nthRootPrecision := big.NewInt(1)
  maxPrecision := big.NewInt(32)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Error: Expected error return due to positive nthRoot. NO ERROR RETURNED!")
  }

}

func TestFixedDecimalNthRoot_CalculateNegativeFractionalNthRoot_17(t *testing.T) {

  radicand := big.NewInt(58971)
  radicandPrecision := big.NewInt(0)
  nthRoot := big.NewInt(52)
  nthRootPrecision := big.NewInt(0)
  maxPrecision := big.NewInt(32)

  fdNr := FixedDecimalNthRoot{}

  _, _, err :=
    fdNr.CalculateNegativeFractionalNthRoot(
      radicand,
      radicandPrecision,
      nthRoot,
      nthRootPrecision,
      maxPrecision)

  if err == nil {
    t.Error("Error: Expected error return due to integer nthRoot. NO ERROR RETURNED!")
  }

}
