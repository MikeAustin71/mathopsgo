package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyIntArySeries_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntArySeries_01"

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

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult.IsValid(ePrefix)\n"+
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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(
    multiplierIntAry,
    multiplicandIntArray[0],
    multiplicandIntArray[1],
    multiplicandIntArray[2],
    multiplicandIntArray[3],
    multiplicandIntArray[4],
    multiplicandIntArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(\n"+
      "multiplierIntAry= '%v'\n"+
      "  multiplicandIntArray[0-5]\n"+
      "Error='%v'\n\n", ePrefix, multiplierIntAryNumStr, err.Error())
    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if iaResultNumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultBigINumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntArySeries_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntArySeries_02"

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
  expectedBigINumStr := "11995826664.26376575446779648"

  expectedBigINumSign := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult.IsValid(ePrefix)\n"+
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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(
    multiplierIntAry,
    multiplicandIntArray[0],
    multiplicandIntArray[1],
    multiplicandIntArray[2],
    multiplicandIntArray[3],
    multiplicandIntArray[4],
    multiplicandIntArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(\n"+
      "multiplierIntAry= '%v'\n"+
      "  multiplicandIntArray[0-5]\n"+
      "Error='%v'\n\n", ePrefix, multiplierIntAryNumStr, err.Error())
    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if iaResultNumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultBigINumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntArySeries_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntArySeries_03"

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
  expectedBigINumStr := "2212352.1767579232"

  expectedBigINumSign := 1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult.IsValid(ePrefix)\n"+
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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(
    multiplierIntAry,
    multiplicandIntArray[0],
    multiplicandIntArray[1],
    multiplicandIntArray[2],
    multiplicandIntArray[3],
    multiplicandIntArray[4],
    multiplicandIntArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(\n"+
      "multiplierIntAry= '%v'\n"+
      "  multiplicandIntArray[0-5]\n"+
      "Error='%v'\n\n", ePrefix, multiplierIntAryNumStr, err.Error())
    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if iaResultNumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultBigINumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntArySeries_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntArySeries_04"

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
  expectedBigINumStr := "-20408.5138429311978576052224"

  expectedBigINumSign := -1

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult.IsValid(ePrefix)\n"+
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

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(
    multiplierIntAry,
    multiplicandIntArray[0],
    multiplicandIntArray[1],
    multiplicandIntArray[2],
    multiplicandIntArray[3],
    multiplicandIntArray[4],
    multiplicandIntArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(\n"+
      "multiplierIntAry= '%v'\n"+
      "  multiplicandIntArray[0-5]\n"+
      "Error='%v'\n\n", ePrefix, multiplierIntAryNumStr, err.Error())
    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if iaResultNumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultBigINumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyIntArySeries_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyIntArySeries_05"

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

  // product = -20408,5138429311978576052224
  expectedBigINumStr := "-20408,5138429311978576052224"

  expectedBigINumSign := -1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAry, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  err = multiplierIntAry.IsValid(ePrefix + "\nValidating multiplierIntAry")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierIntAry.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierIntAryNumStr, err := multiplierIntAry.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenMultiplicandStrsAry := len(multiplicandStrs)

  multiplicandIntArray := make([]IntAry, lenMultiplicandStrsAry)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenMultiplicandStrsAry; i++ {

    multiplicandIntArray[i], err = new(IntAry).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "multiplicandIntArray[%d], err = new(IntAry).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = multiplicandIntArray[i].CopyOut()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = multiplicandIntArray[%d].CopyOut()\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n", ePrefix, i, multiplicandStrs[i], err.Error())
      return
    }

  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
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

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
      "expectedNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(
    multiplierIntAry,
    multiplicandIntArray[0],
    multiplicandIntArray[1],
    multiplicandIntArray[2],
    multiplicandIntArray[3],
    multiplicandIntArray[4],
    multiplicandIntArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyIntArySeries(\n"+
      "multiplierIntAry= '%v'\n"+
      "  multiplicandIntArray[0-5]\n"+
      "Error='%v'\n\n", ePrefix, multiplierIntAryNumStr, err.Error())
    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating result")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  err = resultBigINum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if iaResultNumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultBigINumStr)

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrs_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrs_01"

  var err error

  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedBigINumStr := "2875.94572"

  expectedBigINumSign := 1

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      usaNumSeps.String(),
      err.Error())

    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(\n"+
      "  multiplierStr, multiplicandStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, usaNumSeps.String(), err.Error())

    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if expectedBigINumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
      "resultBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBigINumStr, err.Error())

    return
  }

  expectedResultNumSeps, err := usaNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumSeps, err := usaNumSeps.CopyOut(false)\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, usaNumSeps.String(), err.Error())

    return
  }

  if !expectedResultNumSeps.Equal(resultBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected resultBigINumSeps Numeric Separators = '%v'\n"+
      "  Actual resultBigINumSeps Separators = '%v'\n\n",
      ePrefix, expectedResultNumSeps.String(), resultBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrs_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrs_02"

  var err error
  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedBigINumStr := "30987680500513189125.14259702468435"

  expectedBigINumSign := 1

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      usaNumSeps.String(),
      err.Error())

    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(\n"+
      "  multiplierStr, multiplicandStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, usaNumSeps.String(), err.Error())

    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if expectedBigINumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
      "resultBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBigINumStr, err.Error())

    return
  }

  expectedResultNumSeps, err := usaNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumSeps, err := usaNumSeps.CopyOut(false)\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, usaNumSeps.String(), err.Error())

    return
  }

  if !expectedResultNumSeps.Equal(resultBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected resultBigINumSeps Numeric Separators = '%v'\n"+
      "  Actual resultBigINumSeps Separators = '%v'\n\n",
      ePrefix, expectedResultNumSeps.String(), resultBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrs_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrs_02"

  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedBigINumStr := "-30987680500513189125.14259702468435"

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINumSign := -1

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      usaNumSeps.String(),
      err.Error())

    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(\n"+
      "  multiplierStr, multiplicandStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, usaNumSeps.String(), err.Error())

    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if expectedBigINumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
      "resultBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBigINumStr, err.Error())

    return
  }

  expectedResultNumSeps, err := usaNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumSeps, err := usaNumSeps.CopyOut(false)\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, usaNumSeps.String(), err.Error())

    return
  }

  if !expectedResultNumSeps.Equal(resultBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected resultBigINumSeps Numeric Separators = '%v'\n"+
      "  Actual resultBigINumSeps Separators = '%v'\n\n",
      ePrefix, expectedResultNumSeps.String(), resultBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrs_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrs_02"

  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedBigINumStr := "22197234145.3632"

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      usaNumSeps.String(),
      err.Error())

    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(\n"+
      "  multiplierStr, multiplicandStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, usaNumSeps.String(), err.Error())

    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if expectedBigINumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
      "resultBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBigINumStr, err.Error())

    return
  }

  expectedResultNumSeps, err := usaNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumSeps, err := usaNumSeps.CopyOut(false)\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, usaNumSeps.String(), err.Error())

    return
  }

  if !expectedResultNumSeps.Equal(resultBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected resultBigINumSeps Numeric Separators = '%v'\n"+
      "  Actual resultBigINumSeps Separators = '%v'\n\n",
      ePrefix, expectedResultNumSeps.String(), resultBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrs_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrs_05"

  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0
  expectedBigINumStr := "0"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(\n"+
      "  multiplierStr, multiplicandStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if expectedBigINumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
      "resultBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBigINumStr, err.Error())

    return
  }

  expectedResultNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())

    return
  }

  if !expectedResultNumSeps.Equal(resultBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected resultBigINumSeps Numeric Separators = '%v'\n"+
      "  Actual resultBigINumSeps Separators = '%v'\n\n",
      ePrefix, expectedResultNumSeps.String(), resultBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrs_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrs_06"

  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  expectedBigINumSign := 1

  // product = 2875.94572
  expectedBigINumStr := "2875.94572"

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(\n"+
      "  multiplierStr, multiplicandStr, expectedNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if expectedBigINumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
      "resultBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBigINumStr, err.Error())

    return
  }

  expectedResultNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())

    return
  }

  if !expectedResultNumSeps.Equal(resultBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected resultBigINumSeps Numeric Separators = '%v'\n"+
      "  Actual resultBigINumSeps Separators = '%v'\n\n",
      ePrefix, expectedResultNumSeps.String(), resultBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrs_07(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrs_07"

  // multiplier = 123,32
  multiplierStr := "123,32"

  // multiplicand = 23,321
  multiplicandStr := "23,321"

  // product = 2875,94572
  expectedBigINumStr := "2875,94572"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err := expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err := expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      expectedNumSeps.String(),
      err.Error())

    return
  }

  resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINum, err := new(BigIntMathMultiply).MultiplyNumStr(\n"+
      "  multiplierStr, multiplicandStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "multiplicandStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())

    return
  }

  err = resultBigINum.IsValid(ePrefix + "\nValidating resultBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = resultBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigINumStr, err := resultBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumStr, err := resultBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(resultBigINum)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.Equal(resultBigINum)\n"+
      "expectedBigINum= '%v'"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr,
      resultBigINumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  resultBigINumBigInt, err := resultBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumBigInt, err := resultBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: BigInt Numbers Not Equal\n"+
      "Because expectedBigINumBigInt.Cmp(resultBigINumBigInt) != 0\n"+
      "Expected resultBigINumBigInt = '%v'\n"+
      "  Actual resultBigINumBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigINumBigInt.Text(10))

    return
  }

  resultBigINumSignValue, err := resultBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSignValue, err := resultBigINum.GetSign()\n"+
      "resultBigINum= '%v'"+
      "Error='%v'\n\n", ePrefix, resultBigINumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultBigINumSignValue {
    t.Errorf("%v\n"+
      "Error: Sign Values Not Equal\n"+
      "Because expectedBigINumSign != resultBigINumSignValue \n"+
      "Expected resultBigINumSignValue = '%v'\n"+
      "  Actual resultBigINumSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultBigINumSignValue)

    return
  }

  if expectedBigINumStr != resultBigINumStr {
    t.Errorf("%v\n"+
      "Error: Number String Values Not Equal\n"+
      "Because iaResultNumStr != resultBigINumStr \n"+
      "Expected resultBigINumStr = '%v'\n"+
      "  Actual resultBigINumStr = '%v'\n\n",
      ePrefix, expectedBigINumStr, resultBigINumStr)

    return
  }

  resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigINumSeps, err := resultBigINum.GetNumericSeparatorsDto()\n"+
      "resultBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultBigINumStr, err.Error())

    return
  }

  expectedResultNumSeps, err := expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumSeps, err := expectedNumSeps.CopyOut(false)\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())

    return
  }

  if !expectedResultNumSeps.Equal(resultBigINumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedQuoNumSeps.Equal(actualQuoNumSeps) == 'false'\n"+
      "Expected resultBigINumSeps Numeric Separators = '%v'\n"+
      "  Actual resultBigINumSeps Separators = '%v'\n\n",
      ePrefix, expectedResultNumSeps.String(), resultBigINumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrArray_01"

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

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.IsValid(ePrefix)\n"+
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

  var iaMultiplicand IntAry

  var iaMultiplicandNumStr string

  for i := 0; i < lenArray; i++ {

    iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[%d], usaNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaMultiplicand.IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "\n"+
        "iaResult= '%v'\n"+
        "iaMultiplicand= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
      return
    }

    iaResultNumStr, err = iaResult.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaResultNumStr, err = iaResult.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, usaNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierStr,
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

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps := result.GetNumericSeparatorsDto()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !usaNumSeps.Equal(resultNumSeps)\n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrArray_02"

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
  expectedBigINumStr := "11995826664.26376575446779648"

  expectedBigINumSign := 1

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.IsValid(ePrefix)\n"+
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

  var iaMultiplicand IntAry

  var iaMultiplicandNumStr string

  for i := 0; i < lenArray; i++ {

    iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[%d], usaNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaMultiplicand.IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "\n"+
        "iaResult= '%v'\n"+
        "iaMultiplicand= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
      return
    }

    iaResultNumStr, err = iaResult.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaResultNumStr, err = iaResult.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, usaNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierStr,
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

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps := result.GetNumericSeparatorsDto()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !usaNumSeps.Equal(resultNumSeps)\n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrArray_03"

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
  expectedBigINumStr := "2212352.1767579232"

  expectedBigINumSign := 1

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.IsValid(ePrefix)\n"+
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

  var iaMultiplicand IntAry

  var iaMultiplicandNumStr string

  for i := 0; i < lenArray; i++ {

    iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[%d], usaNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaMultiplicand.IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "\n"+
        "iaResult= '%v'\n"+
        "iaMultiplicand= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
      return
    }

    iaResultNumStr, err = iaResult.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaResultNumStr, err = iaResult.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, usaNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierStr,
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

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps := result.GetNumericSeparatorsDto()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !usaNumSeps.Equal(resultNumSeps)\n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrArray_04"

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
  expectedBigINumStr := "-20408.5138429311978576052224"

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  expectedBigINumSign := -1

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.IsValid(ePrefix)\n"+
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

  var iaMultiplicand IntAry

  var iaMultiplicandNumStr string

  for i := 0; i < lenArray; i++ {

    iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[%d], usaNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaMultiplicand.IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "\n"+
        "iaResult= '%v'\n"+
        "iaMultiplicand= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
      return
    }

    iaResultNumStr, err = iaResult.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaResultNumStr, err = iaResult.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &usaNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, usaNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierStr,
      usaNumSeps.String(),
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

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps := result.GetNumericSeparatorsDto()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStr, err.Error())
    return
  }

  if !usaNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !usaNumSeps.Equal(resultNumSeps)\n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, usaNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrArray_05"

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
  expectedBigINumStr := "11995826664.26376575446779648"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  usaNumSeps := new(NumericSeparatorDto).NewUSADefaults()

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierStr, usaNumSeps.String(), err.Error())
    return
  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.IsValid(ePrefix)\n"+
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

  var iaMultiplicand IntAry

  var iaMultiplicandNumStr string

  for i := 0; i < lenArray; i++ {

    iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], usaNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[%d], usaNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "usaNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, multiplicandStrs[i], usaNumSeps.String(), err.Error())
      return
    }

    err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaMultiplicand.IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "\n"+
        "iaResult= '%v'\n"+
        "iaMultiplicand= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
      return
    }

    iaResultNumStr, err = iaResult.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaResultNumStr, err = iaResult.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrArray(multiplierStr, multiplicandStrs, usaNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "usaNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierStr,
      usaNumSeps.String(),
      err.Error())

    return
  }

  err = result.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = result.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, expectedNumSeps.String(), err.Error())
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

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps := result.GetNumericSeparatorsDto()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !expectedNumSeps.Equal(resultNumSeps)\n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyNumStrArray_06"

  var err error

  // multiplier = 37,9876
  multiplierStr := "37,9876"

  // multiplicandStrs
  multiplicandStrs := []string{
    "-27,9",
    "48,123456",
    "59,48721",
    "-3",
    "19,1",
    "69",
  }

  // product = 11995826664,26376575446779648
  expectedBigINumStr := "11995826664,26376575446779648"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = expectedNumSeps.IsValid("Validating expectedNumSeps")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedNumSeps.IsValid('Validating expectedNumSeps')\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedNumSeps.String(), err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, multiplierStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = iaResult.IsValid(ePrefix + "\nValidating iaResult")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.IsValid(ePrefix)\n"+
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

  var iaMultiplicand IntAry

  var iaMultiplicandNumStr string

  for i := 0; i < lenArray; i++ {

    iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicand, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[%d], expectedNumSeps)\n"+
        "multiplicandStrs[%d]= '%v'\n"+
        "expectedNumSeps= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix, i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
      return
    }

    err = iaMultiplicand.IsValid(ePrefix + "\nValidating iaMultiplicand")

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaMultiplicand.IsValid(ePrefix)\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    iaMultiplicandNumStr, err = iaMultiplicand.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&iaMultiplicand, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "\n"+
        "iaResult= '%v'\n"+
        "iaMultiplicand= '%v'\n"+
        "Error= '%v'\n\n",
        ePrefix, iaResultNumStr, iaMultiplicandNumStr, err.Error())
      return
    }

    iaResultNumStr, err = iaResult.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaResultNumStr, err = iaResult.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedBigINumStr, &expectedNumSeps)\n"+
      "expectedBigINumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedBigINumStr, expectedNumSeps.String(), err.Error())
    return
  }

  err = expectedBigINum.IsValid(ePrefix + "\nValidating expectedBigINum")

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = expectedBigINum.IsValid(ePrefix)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigIntNumberStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)\n"+
      "multiplierStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierStr,
      expectedNumSeps.String(),
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

  expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err := expectedBigINum.Equal(result)\n"+
      "expectedBigINum= '%v'"+
      "result= '%v'"+
      "Error='%v'\n\n", ePrefix, expectedBigIntNumberStr, resultNumStr, err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumEqualsResult = 'false'\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigIntNumberStr, resultNumStr)

    return
  }

  resultActualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultActualBigInt, err := result.GetBigInt()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "expectedBigINum= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, expectedBigIntNumberStr, err.Error())
    return
  }

  if expectedBigINumBigInt.Cmp(resultActualBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumBigInt.Cmp(resultActualBigInt) != 0\n"+
      "Expected resultActualBigInt = '%v'\n"+
      "  Actual resultActualBigInt = '%v'\n\n",
      ePrefix, expectedBigINumBigInt.Text(10), resultActualBigInt.Text(10))

    return
  }

  resultSign, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSign, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n", ePrefix, resultNumStr, err.Error())
    return
  }

  if expectedBigINumSign != resultSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultSign\n"+
      "Expected resultSign = '%v'\n"+
      "  Actual resultSign = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSign)

    return
  }

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because iaResultNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

    return
  }

  resultNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultNumSeps := result.GetNumericSeparatorsDto()\n"+
      "result= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, resultNumStr, err.Error())
    return
  }

  if !expectedNumSeps.Equal(resultNumSeps) {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because !expectedNumSeps.Equal(resultNumSeps)\n"+
      "Expected resultNumSeps = '%v'\n"+
      "  Actual resultNumSeps = '%v'\n\n",
      ePrefix, expectedNumSeps.String(), resultNumSeps.String())

    return
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_01(t *testing.T) {

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

  expectedNumSeps := new(NumericSeparatorDto).New()

  result, err :=
    new(BigIntMathMultiply).MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrOutputToArray"+
      "(multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  lenArray := len(multiplicandStrs)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j] {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j])
    }
  }
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_02(t *testing.T) {

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

  expectedNumSeps := new(NumericSeparatorDto).New()

  result, err :=
    new(BigIntMathMultiply).MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrOutputToArray"+
      "(multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  lenArray := len(multiplicandStrs)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j] {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j])
    }

  }
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_03(t *testing.T) {

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

  expectedNumSeps := new(NumericSeparatorDto).New()

  result, err :=
    new(BigIntMathMultiply).MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrOutputToArray"+
      "(multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  lenArray := len(multiplicandStrs)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j] {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j])
    }

  }
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_04(t *testing.T) {

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

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  result, err :=
    new(BigIntMathMultiply).MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrOutputToArray"+
      "(multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  lenArray := len(multiplicandStrs)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j] {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j])
    }

  }
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_05(t *testing.T) {

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

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  result, err :=
    new(BigIntMathMultiply).MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrOutputToArray"+
      "(multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  lenArray := len(multiplicandStrs)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j] {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j])
    }

  }
}

func TestBigIntMathMultiply_MultiplyNumStrOutputToArray_06(t *testing.T) {

  var err error

  // multiplier = -31,2
  multiplierStr := "-31,2"
  // multiplicandStrs
  multiplicandStrs := []string{
    "100,1",
    "-26",
    "3,924",
    "8",
    "5297,123",
    "-4,896",
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

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'
  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  result, err :=
    new(BigIntMathMultiply).MultiplyNumStrOutputToArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrOutputToArray"+
      "(multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  lenArray := len(multiplicandStrs)

  for j := 0; j < lenArray; j++ {

    if expectedNumStrs[j] != result[j] {
      t.Errorf("Error: Expected NumStr[%v]='%v'. Instead NumStr[%v]='%v'. ",
        j, expectedNumStrs[j], j, result[j])
    }

  }
}
