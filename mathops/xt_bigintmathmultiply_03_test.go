package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumSeries_01"

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
  originalExpectedBigINumStr := "128"

  originalExpectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        ePrefix,
        i, i, i,
        multiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v'\nError='%v'\n\n",
        ePrefix, iaNumStr, err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResultBigINum, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResultBigINum {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: *big.Int values do NOT match!\n"+
      "       Expected BigIntNum= '%s'.\n"+
      "         Result BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  if originalExpectedBigINumStr != expectedBigINumStr ||
    originalExpectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedBigINumStr)
    return

  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumSeries_02"

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
  originalExpectedBigINumStr := "11995826664.26376575446779648"

  originalExpectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).
      NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        ePrefix,
        i, i, i,
        multiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v'\nError='%v'\n\n",
        ePrefix, iaNumStr, err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResultBigINum, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResultBigINum {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: *big.Int values do NOT match!\n"+
      "       Expected BigIntNum= '%s'.\n"+
      "         Result BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  expectedBigINumCmpBigIntResultBigInt, err := expectedBigINum.CmpBigInt(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumCmpBigIntResultBigInt != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedBigINumStr != expectedBigINumStr ||
    originalExpectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedBigINumStr)
    return

  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumSeries_03"

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
  originalExpectedBigINumStr := "2212352.1767579232"

  originalExpectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        ePrefix,
        i, i, i,
        multiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v'\nError='%v'\n\n",
        ePrefix, iaNumStr, err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResultBigINum, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResultBigINum {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: *big.Int values do NOT match!\n"+
      "       Expected BigIntNum= '%s'.\n"+
      "         Result BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  expectedBigINumCmpBigIntResultBigInt, err := expectedBigINum.CmpBigInt(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumCmpBigIntResultBigInt != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedBigINumStr != expectedBigINumStr ||
    originalExpectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedBigINumStr)
    return

  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
    return
  }

  err = iaResult.OptimizeIntArrayLen(true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.OptimizeIntArrayLen(true)\n"+
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumSeries_04"

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
  originalExpectedBigINumStr := "-20408.5138429311978576052224"

  originalExpectedBigINumSign := -1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        ePrefix,
        i, i, i,
        multiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v'\nError='%v'\n\n",
        ePrefix, iaNumStr, err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResultBigINum, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResultBigINum {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: *big.Int values do NOT match!\n"+
      "       Expected BigIntNum= '%s'.\n"+
      "         Result BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  expectedBigINumCmpBigIntResultBigInt, err := expectedBigINum.CmpBigInt(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumCmpBigIntResultBigInt != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedBigINumStr != expectedBigINumStr ||
    originalExpectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedBigINumStr)
    return

  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
    return
  }

  err = iaResult.OptimizeIntArrayLen(true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.OptimizeIntArrayLen(true)\n"+
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyBigIntNumSeries_05"

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
  originalExpectedBigINumStr := "2212352,1767579232"

  originalExpectedBigINumSign := 1

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  lenArray := len(multiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'multiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "bINumArray[%d], err = new(BigIntNum).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        ePrefix,
        i, i, i,
        multiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = bINumArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix, i, err.Error())
      return
    }

    iaNumStr, err = ia.GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "iaNumStr, err = ia.GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia='%v'\nError='%v'\n\n",
        ePrefix, iaNumStr, err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedBigINumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedBigINumStr, err.Error())
    return
  }

  expectedBigINumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumStr, err := expectedBigINum.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumSign, err := expectedBigINum.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumSign, err := expectedBigINum.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyBigIntNumSeries(...)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedBigINumEqualResultBigINum, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualResultBigINum, err := \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedBigINumEqualResultBigINum {
    t.Errorf("%v\n"+
      "Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: *big.Int values do NOT match!\n"+
      "       Expected BigIntNum= '%s'.\n"+
      "         Result BigIntNum= '%s'.\n",
      ePrefix,
      expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  expectedBigINumCmpBigIntResultBigInt, err := expectedBigINum.CmpBigInt(result)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedBigINumCmpBigIntResultBigInt != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedBigINumBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedBigINumStr != expectedBigINumStr ||
    originalExpectedBigINumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedBigINumStr)
    return

  }

  if originalExpectedBigINumSign != resultSignValue ||
    originalExpectedBigINumSign != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedBigINumSign,
      resultSignValue,
      expectedBigINumSign)
    return
  }

  err = iaResult.OptimizeIntArrayLen(true)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.OptimizeIntArrayLen(true)\n"+
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)
    return
  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected and Actual Numeric SeparatorsExpected do NOT match!\n"+
      "Expected NumSeps='%v'.\n"+
      "Actual NumSeps='%v'.\n\n",
      ePrefix, expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyDecimal_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimal_01"

  // multiplier = 123.32
  originalMultiplierStr := "123.32"

  // multiplicand = 23.321
  originalMultiplicandStr := "23.321"

  // product = 2875.94572
  originalExpectedDecimalNumStr := "2875.94572"

  originalExpectedDecimalSign := 1

  multiplierDecimal, err := new(Decimal).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  multiplicandDecimal, err := new(Decimal).
    NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplierNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplier and Original Multiplier String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplierStr='%v' \n"+
      "Converted IntAry Multiplier= '%v'\n",
      ePrefix, originalMultiplierStr, iaMultiplierNumStr)
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplicandNumStr != originalMultiplicandStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplicand and Original Multiplicand String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplicandStr='%v' \n"+
      "Converted IntAry Multiplicand= '%v'\n",
      ePrefix, originalMultiplicandStr, iaMultiplicandNumStr)
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(&iaMultiplier,\n"+
      "  &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier= '%v'\n"+
      "iaMultiplicand= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMultiplierNumStr,
      iaMultiplicandNumStr,
      err.Error())
    return
  }

  expectedDecimal, err := new(Decimal).NewNumStr(originalExpectedDecimalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedDecimalNumStr, err.Error())
    return
  }

  expectedDecimalNumStr, err := expectedDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalNumStr, err := expectedDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalBigInt, err := expectedDecimal.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalSign, err := expectedDecimal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalSign, err :=expectedDecimal.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandDecimalNumStr, err := multiplicandDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimalNumStr, err = multiplicandDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimal(\n"+
      "    multiplierDecimal, multiplicandDecimal)\n"+
      "multiplierDecimal= '%v'\n"+
      "multiplicandDecimal= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierDecimalNumStr,
      multiplicandDecimalNumStr,
      err.Error())

    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultDecimal, err := result.GetDecimal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultDecimal, err := result.GetDecimal()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalCmpResultDecimal, err := expectedDecimal.Cmp(resultDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedDecimalCmpResultDecimal != 0 {

    t.Errorf("%v\n"+
      "Error: Expected Decimal and Actual Result Decimal are unequal!\n"+
      "Expected Decimal= '%s'.\n"+
      "Result Decimal  = '%s'.\n\n",
      ePrefix,
      expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalNumStr != expectedDecimalNumStr ||
    originalExpectedDecimalNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedDecimalNumStr,
      resultNumStr,
      originalExpectedDecimalNumStr)
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected Decimal *big.Int='%s'.\n"+
      "'result' Decimal *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalSign != resultSignValue ||
    originalExpectedDecimalSign != expectedDecimalSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedDecimalSign='%v'\n",
      ePrefix,
      originalExpectedDecimalSign,
      resultSignValue,
      expectedDecimalSign)
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

  }

  return
}

func TestBigIntMathMultiply_MultiplyDecimal_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimal_02"

  // multiplier = 57638422123.327890123
  originalMultiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  originalMultiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  originalExpectedDecimalNumStr := "30987680500513189125.14259702468435"

  originalExpectedDecimalSign := 1

  multiplierDecimal, err := new(Decimal).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  multiplicandDecimal, err := new(Decimal).
    NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimal, err := \n"+
      "  new(Decimal).NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplierNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplier and Original Multiplier String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplierStr='%v' \n"+
      "Converted IntAry Multiplier= '%v'\n",
      ePrefix, originalMultiplierStr, iaMultiplierNumStr)
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplicandNumStr != originalMultiplicandStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplicand and Original Multiplicand String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplicandStr='%v' \n"+
      "Converted IntAry Multiplicand= '%v'\n",
      ePrefix, originalMultiplicandStr, iaMultiplicandNumStr)
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(&iaMultiplier,\n"+
      "  &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier= '%v'\n"+
      "iaMultiplicand= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMultiplierNumStr,
      iaMultiplicandNumStr,
      err.Error())
    return
  }

  expectedDecimal, err := new(Decimal).NewNumStr(originalExpectedDecimalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedBigINumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedDecimalNumStr, err.Error())
    return
  }

  expectedDecimalNumStr, err := expectedDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalNumStr, err := expectedDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalBigInt, err := expectedDecimal.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalSign, err := expectedDecimal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalSign, err :=expectedDecimal.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandDecimalNumStr, err := multiplicandDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimalNumStr, err = multiplicandDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimal(\n"+
      "    multiplierDecimal, multiplicandDecimal)\n"+
      "multiplierDecimal= '%v'\n"+
      "multiplicandDecimal= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierDecimalNumStr,
      multiplicandDecimalNumStr,
      err.Error())

    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultDecimal, err := result.GetDecimal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultDecimal, err := result.GetDecimal()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalCmpResultDecimal, err := expectedDecimal.Cmp(resultDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedDecimalCmpResultDecimal != 0 {

    t.Errorf("%v\n"+
      "Error: Expected Decimal and Actual Result Decimal are unequal!\n"+
      "Expected Decimal= '%s'.\n"+
      "Result Decimal  = '%s'.\n\n",
      ePrefix,
      expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalNumStr != expectedDecimalNumStr ||
    originalExpectedDecimalNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedDecimalNumStr,
      resultNumStr,
      originalExpectedDecimalNumStr)
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected Decimal *big.Int='%s'.\n"+
      "'result' Decimal *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalSign != resultSignValue ||
    originalExpectedDecimalSign != expectedDecimalSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedDecimalSign='%v'\n",
      ePrefix,
      originalExpectedDecimalSign,
      resultSignValue,
      expectedDecimalSign)
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

  }

  return
}

func TestBigIntMathMultiply_MultiplyDecimal_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimal_03"

  // multiplier = 123.32
  originalMultiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  originalMultiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  originalExpectedDecimalNumStr := "-30987680500513189125.14259702468435"

  originalExpectedDecimalSign := -1

  multiplierDecimal, err := new(Decimal).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  multiplicandDecimal, err := new(Decimal).
    NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplierNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplier and Original Multiplier String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplierStr='%v' \n"+
      "Converted IntAry Multiplier= '%v'\n",
      ePrefix, originalMultiplierStr, iaMultiplierNumStr)
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplicandNumStr != originalMultiplicandStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplicand and Original Multiplicand String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplicandStr='%v' \n"+
      "Converted IntAry Multiplicand= '%v'\n",
      ePrefix, originalMultiplicandStr, iaMultiplicandNumStr)
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(&iaMultiplier,\n"+
      "  &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier= '%v'\n"+
      "iaMultiplicand= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMultiplierNumStr,
      iaMultiplicandNumStr,
      err.Error())
    return
  }

  expectedDecimal, err := new(Decimal).NewNumStr(originalExpectedDecimalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedDecimalNumStr, err.Error())
    return
  }

  expectedDecimalNumStr, err := expectedDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalNumStr, err := expectedDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalBigInt, err := expectedDecimal.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalSign, err := expectedDecimal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalSign, err :=expectedDecimal.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandDecimalNumStr, err := multiplicandDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimalNumStr, err = multiplicandDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimal(\n"+
      "    multiplierDecimal, multiplicandDecimal)\n"+
      "multiplierDecimal= '%v'\n"+
      "multiplicandDecimal= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierDecimalNumStr,
      multiplicandDecimalNumStr,
      err.Error())

    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultDecimal, err := result.GetDecimal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultDecimal, err := result.GetDecimal()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalCmpResultDecimal, err := expectedDecimal.Cmp(resultDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedDecimalCmpResultDecimal != 0 {

    t.Errorf("%v\n"+
      "Error: Expected Decimal and Actual Result Decimal are unequal!\n"+
      "Expected Decimal= '%s'.\n"+
      "Result Decimal  = '%s'.\n\n",
      ePrefix,
      expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalNumStr != expectedDecimalNumStr ||
    originalExpectedDecimalNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedDecimalNumStr,
      resultNumStr,
      originalExpectedDecimalNumStr)
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected Decimal *big.Int='%s'.\n"+
      "'result' Decimal *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalSign != resultSignValue ||
    originalExpectedDecimalSign != expectedDecimalSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedDecimalSign='%v'\n",
      ePrefix,
      originalExpectedDecimalSign,
      resultSignValue,
      expectedDecimalSign)
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

  }

  return
}

func TestBigIntMathMultiply_MultiplyDecimal_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimal_04"

  // multiplier = 89637.9876
  originalMultiplierStr := "-89637.9876"

  // multiplicand = -247632
  originalMultiplicandStr := "-247632"

  // product = 22197234145.3632
  originalExpectedDecimalNumStr := "22197234145.3632"

  originalExpectedDecimalSign := 1

  multiplierDecimal, err := new(Decimal).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  multiplicandDecimal, err := new(Decimal).
    NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplierNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplier and Original Multiplier String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplierStr='%v' \n"+
      "Converted IntAry Multiplier= '%v'\n",
      ePrefix, originalMultiplierStr, iaMultiplierNumStr)
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplicandNumStr != originalMultiplicandStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplicand and Original Multiplicand String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplicandStr='%v' \n"+
      "Converted IntAry Multiplicand= '%v'\n",
      ePrefix, originalMultiplicandStr, iaMultiplicandNumStr)
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(&iaMultiplier,\n"+
      "  &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier= '%v'\n"+
      "iaMultiplicand= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMultiplierNumStr,
      iaMultiplicandNumStr,
      err.Error())
    return
  }

  expectedDecimal, err := new(Decimal).NewNumStr(originalExpectedDecimalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedDecimalNumStr, err.Error())
    return
  }

  expectedDecimalNumStr, err := expectedDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalNumStr, err := expectedDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalBigInt, err := expectedDecimal.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalSign, err := expectedDecimal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalSign, err :=expectedDecimal.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandDecimalNumStr, err := multiplicandDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimalNumStr, err = multiplicandDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimal(\n"+
      "    multiplierDecimal, multiplicandDecimal)\n"+
      "multiplierDecimal= '%v'\n"+
      "multiplicandDecimal= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierDecimalNumStr,
      multiplicandDecimalNumStr,
      err.Error())

    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultDecimal, err := result.GetDecimal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultDecimal, err := result.GetDecimal()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalCmpResultDecimal, err := expectedDecimal.Cmp(resultDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedDecimalCmpResultDecimal != 0 {

    t.Errorf("%v\n"+
      "Error: Expected Decimal and Actual Result Decimal are unequal!\n"+
      "Expected Decimal= '%s'.\n"+
      "Result Decimal  = '%s'.\n\n",
      ePrefix,
      expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalNumStr != expectedDecimalNumStr ||
    originalExpectedDecimalNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedDecimalNumStr,
      resultNumStr,
      originalExpectedDecimalNumStr)
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected Decimal *big.Int='%s'.\n"+
      "'result' Decimal *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalSign != resultSignValue ||
    originalExpectedDecimalSign != expectedDecimalSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedDecimalSign='%v'\n",
      ePrefix,
      originalExpectedDecimalSign,
      resultSignValue,
      expectedDecimalSign)
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

  if iaResultNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: iaResult and result number strings do NOT match!\n"+
      "iaResultNumStr='%v'\n"+
      "  resultNumStr='%v'\n\n",
      ePrefix, iaResultNumStr, resultNumStr)

  }

  return
}

func TestBigIntMathMultiply_MultiplyDecimal_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimal_05"

  // multiplier = -89637.9876
  originalMultiplierStr := "-89637.9876"

  // multiplicand = 0.00
  originalMultiplicandStr := "0.00"

  // product = 0.00
  originalExpectedDecimalNumStr := "0"

  originalExpectedDecimalSign := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  multiplicandDecimal, err := new(Decimal).NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplierNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplier and Original Multiplier String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplierStr='%v' \n"+
      "Converted IntAry Multiplier= '%v'\n",
      ePrefix, originalMultiplierStr, iaMultiplierNumStr)
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplicandNumStr != originalMultiplicandStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplicand and Original Multiplicand String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplicandStr='%v' \n"+
      "Converted IntAry Multiplicand= '%v'\n",
      ePrefix, originalMultiplicandStr, iaMultiplicandNumStr)
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(&iaMultiplier,\n"+
      "  &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier= '%v'\n"+
      "iaMultiplicand= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMultiplierNumStr,
      iaMultiplicandNumStr,
      err.Error())
    return
  }

  expectedDecimal, err := new(Decimal).NewNumStr(originalExpectedDecimalNumStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(expectedNumStr) "+
      "expectedNumStr='%v'  Error='%v'. ", originalExpectedDecimalNumStr, err.Error())
  }

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedDecimalNumStr, err.Error())
    return
  }

  expectedDecimalNumStr, err := expectedDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalNumStr, err := expectedDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalBigInt, err := expectedDecimal.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalSign, err := expectedDecimal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalSign, err :=expectedDecimal.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandDecimalNumStr, err := multiplicandDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimalNumStr, err = multiplicandDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimal(\n"+
      "    multiplierDecimal, multiplicandDecimal)\n"+
      "multiplierDecimal= '%v'\n"+
      "multiplicandDecimal= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierDecimalNumStr,
      multiplicandDecimalNumStr,
      err.Error())

    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultDecimal, err := result.GetDecimal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultDecimal, err := result.GetDecimal()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalCmpResultDecimal, err := expectedDecimal.Cmp(resultDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedDecimalCmpResultDecimal != 0 {

    t.Errorf("%v\n"+
      "Error: Expected Decimal and Actual Result Decimal are unequal!\n"+
      "Expected Decimal= '%s'.\n"+
      "Result Decimal  = '%s'.\n\n",
      ePrefix,
      expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalNumStr != expectedDecimalNumStr ||
    originalExpectedDecimalNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedDecimalNumStr,
      resultNumStr,
      originalExpectedDecimalNumStr)
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected Decimal *big.Int='%s'.\n"+
      "'result' Decimal *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalSign != resultSignValue ||
    originalExpectedDecimalSign != expectedDecimalSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedDecimalSign='%v'\n",
      ePrefix,
      originalExpectedDecimalSign,
      resultSignValue,
      expectedDecimalSign)
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

  if resultNumStr != iaResultNumStr {
    t.Errorf("%v\n"+
      "Error: 'ia' and 'result' number strings are Unequal!\n"+
      "'result' number string = %s\n"+
      "    'ia' number string = %s\n\n",
      ePrefix,
      resultNumStr,
      iaResultNumStr)
    return
  }

  iaBigInt, err := iaResult.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaBigInt, err := iaResult.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if resultBigInt.Cmp(iaBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: 'ia' and 'result' *big.Int values do NOT match!\n"+
      "result *big.Int ='%v'\n"+
      "    ia *big.int ='%v'\n\n",
      ePrefix, resultBigInt.Text(10), iaBigInt.Text(10))
  }

  return
}

func TestBigIntMathMultiply_MultiplyDecimal_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimal_06"

  // multiplier = 123.32
  originalMultiplierStr := "123.32"

  // multiplicand = 23.321
  originalMultiplicandStr := "23.321"

  // product = 2875.94572
  originalExpectedDecimalNumStr := "2875,94572"

  originalExpectedDecimalSign := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = multiplierDecimal.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierDecimal.SetNumericSeparatorsDto(expectedNumSeps)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandDecimal, err := new(Decimal).
    NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", originalMultiplicandStr, err.Error())
  }

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaMultiplierNumStr, err := iaMultiplier.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplierNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplier and Original Multiplier String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplierStr='%v' \n"+
      "Converted IntAry Multiplier= '%v'\n",
      ePrefix, originalMultiplierStr, iaMultiplierNumStr)
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(originalMultiplicandStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(originalMultiplicandStr)\n"+
      "originalMultiplicandStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplicandStr, err.Error())
    return
  }

  iaMultiplicandNumStr, err := iaMultiplicand.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaMultiplierNumStr, err := iaMultiplier.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaMultiplicandNumStr != originalMultiplicandStr {
    t.Errorf("%v\n"+
      "Error: Converted IntAry Multiplicand and Original Multiplicand String,\n"+
      "are NOT equal!\n"+
      "      originalMultiplicandStr='%v' \n"+
      "Converted IntAry Multiplicand= '%v'\n",
      ePrefix, originalMultiplicandStr, iaMultiplicandNumStr)
    return
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaMultiplier.Multiply(&iaMultiplier,\n"+
      "  &iaMultiplicand, &iaResult, -1, -1)\n"+
      "iaMultiplier= '%v'\n"+
      "iaMultiplicand= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMultiplierNumStr,
      iaMultiplicandNumStr,
      err.Error())
    return
  }

  expectedDecimal, err := new(Decimal).NewNumStr(originalExpectedDecimalNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimal, err := new(Decimal).\n"+
      "  NewNumStr(originalExpectedBigINumStr)\n"+
      "originalExpectedDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedDecimalNumStr, err.Error())
    return
  }

  expectedDecimalNumStr, err := expectedDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalNumStr, err := expectedDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalBigInt, err := expectedDecimal.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalSign, err := expectedDecimal.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedDecimalSign, err :=expectedDecimal.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  multiplicandDecimalNumStr, err := multiplicandDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplicandDecimalNumStr, err = multiplicandDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).
    MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimal(\n"+
      "    multiplierDecimal, multiplicandDecimal)\n"+
      "multiplierDecimal= '%v'\n"+
      "multiplicandDecimal= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      multiplierDecimalNumStr,
      multiplicandDecimalNumStr,
      err.Error())

    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
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

  resultSignValue, err := result.GetSign()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  resultDecimal, err := result.GetDecimal()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "resultDecimal, err := result.GetDecimal()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  expectedDecimalCmpResultDecimal, err := expectedDecimal.Cmp(resultDecimal)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if expectedDecimalCmpResultDecimal != 0 {

    t.Errorf("%v\n"+
      "Error: Expected Decimal and Actual Result Decimal are unequal!\n"+
      "Expected Decimal= '%s'.\n"+
      "Result Decimal  = '%s'.\n\n",
      ePrefix,
      expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected BigIntNum *big.Int='%s'.\n"+
      "   Actual 'result' *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalNumStr != expectedDecimalNumStr ||
    originalExpectedDecimalNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Decimal Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedDecimalNumStr,
      resultNumStr,
      originalExpectedDecimalNumStr)
    return

  }

  if expectedDecimalBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("%v\n"+
      "Error: Comparision of expected vs result *big.Int Unequal!\n"+
      "Expected Decimal *big.Int='%s'.\n"+
      "'result' Decimal *big.Int= '%s'. ",
      ePrefix, expectedDecimalBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if originalExpectedDecimalSign != resultSignValue ||
    originalExpectedDecimalSign != expectedDecimalSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedDecimalSign='%v'\n",
      ePrefix,
      originalExpectedDecimalSign,
      resultSignValue,
      expectedDecimalSign)
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

  if iaResultNumStr != resultNumStr ||
    iaResultNumStr != originalExpectedDecimalNumStr {
    t.Errorf("%v\n"+
      "Error: 'ia', 'result and 'original' Number Strings are unequal!\n"+
      "             'ia' Number String = '%s'.\n"+
      "         'result' Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      iaResultNumStr,
      resultNumStr,
      originalExpectedDecimalNumStr)
    return

  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected NumSeps='%v'.\n"+
      "Instead, Actual NumSeps='%v'.\n\n",
      ePrefix, expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}
