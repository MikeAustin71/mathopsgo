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
  expectedNumStr := "-20408,5138429311978576052224"

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

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumStr, &expectedNumSeps)\n"+
      "expectedNumStr= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n", ePrefix, expectedNumStr, expectedNumSeps.String(), err.Error())
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
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedNumStr := "2875.94572"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by .NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStr"+
      "(multiplierStr, multiplicandStr, expectedNumSeps) "+
      "multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrs_02(t *testing.T) {
  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedNumStr := "30987680500513189125.14259702468435"

  expectedNumSeps := NumericSeparatorDto{}

  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStr"+
      "(multiplierStr, multiplicandStr, expectedNumSeps) "+
      "multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
  }

  expectedNumSeps.SetDefaultsIfEmpty()

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

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrs_03(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedNumStr := "-30987680500513189125.14259702468435"

  expectedNumSeps := new(NumericSeparatorDto).New()

  expectedBigINumSign := -1

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStr"+
      "(multiplierStr, multiplicandStr, expectedNumSeps) "+
      "multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrs_04(t *testing.T) {
  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedNumStr := "22197234145.3632"

  expectedNumSeps := new(NumericSeparatorDto).New()

  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStr"+
      "(multiplierStr, multiplicandStr, expectedNumSeps) "+
      "multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrs_05(t *testing.T) {
  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0
  expectedNumStr := "0"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStr"+
      "(multiplierStr, multiplicandStr, expectedNumSeps) "+
      "multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrs_06(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedNumStr := "2875.94572"

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStr"+
      "(multiplierStr, multiplicandStr, expectedNumSeps) "+
      "multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathMultiply_MultiplyNumStrs_07(t *testing.T) {
  // multiplier = 123,32
  multiplierStr := "123,32"

  // multiplicand = 23,321
  multiplicandStr := "23,321"

  // product = 2875,94572
  expectedNumStr := "2875,94572"

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  expectedBigINumSign := 1

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStr(multiplierStr, multiplicandStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStr"+
      "(multiplierStr, multiplicandStr, expectedNumSeps) "+
      "multiplierStr='%v' multiplicandStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, multiplicandStr, expectedNumSeps.String(), err.Error())
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

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathMultiply_MultiplyNumStrArray_01(t *testing.T) {

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
  expectedNumStr := "128"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  for i := 0; i < lenArray; i++ {

    ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps) "+
        "multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrArray("+
      "multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
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

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrArray_02(t *testing.T) {

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
  expectedNumStr := "11995826664.26376575446779648"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  for i := 0; i < lenArray; i++ {

    ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
        "multiplicandStrs[i], expectedNumSeps) "+
        "multiplicandStrs[%v]='%v' expectedNumSeps='%v'  Error='%v'. ",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v' Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrArray("+
      "multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
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

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrArray_03(t *testing.T) {

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
  expectedNumStr := "2212352.1767579232"

  expectedBigINumSign := 1

  expectedNumSeps := new(NumericSeparatorDto).New()

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  for i := 0; i < lenArray; i++ {

    ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
        "multiplicandStrs[i], expectedNumSeps) "+
        "multiplicandStrs[%v]='%v' expectedNumSeps='%v'  Error='%v'. ",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrArray("+
      "multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
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

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrArray_04(t *testing.T) {

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
  expectedNumStr := "-20408.5138429311978576052224"

  expectedNumSeps := new(NumericSeparatorDto).New()

  expectedBigINumSign := -1

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  for i := 0; i < lenArray; i++ {

    ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
        "multiplicandStrs[i], expectedNumSeps) "+
        "multiplicandStrs[%v]='%v' expectedNumSeps='%v'  Error='%v'. ",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrArray("+
      "multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v' Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
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

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
}

func TestBigIntMathMultiply_MultiplyNumStrArray_05(t *testing.T) {

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
  expectedNumStr := "11995826664.26376575446779648"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}

  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  for i := 0; i < lenArray; i++ {

    ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
        "multiplicandStrs[i], expectedNumSeps) "+
        "multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrArray("+
      "multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
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

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathMultiply_MultiplyNumStrArray_06(t *testing.T) {

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
  expectedNumStr := "11995826664,26376575446779648"

  expectedBigINumSign := 1

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())
  }

  for i := 0; i < lenArray; i++ {

    ia, err := new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by new(IntAry).NewNumStrWithNumSeps("+
        "multiplicandStrs[i], expectedNumSeps) "+
        "multiplicandStrs[%v]='%v' expectedNumSeps='%v' Error='%v'. ",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by Decimal{}.NewNumStrWithNumSeps("+
      "expectedNumStr, expectedNumSeps) "+
      "expectedNumStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      expectedNumStr, expectedNumSeps.String(), err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrArray(multiplierStr, multiplicandStrs, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrArray("+
      "multiplierStr, multiplicandStrs, expectedNumSeps) "+
      "multiplierStr='%v' expectedNumSeps='%v'  Error='%v'. ",
      multiplierStr, expectedNumSeps.String(), err.Error())

    return
  }

  booResult, err := expectedBigINum.Equal(result)

  if err != nil {

    t.Errorf("Error returned by expectedBigINum.Equal(result)\n"+
      "Error: %v", err.Error())
    return
  }

  if !booResult {
    t.Errorf("Error: Expected expectedBigINum bigInt ='%s'. Instead, result bigInt = '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'. Instead, Decimal= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error:  Expected iaResultNumStr='%v' "+
      "Instead, actual NumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }
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
