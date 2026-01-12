package mathops

import "testing"

func TestBigIntMathMultiply_MultiplyDecimalArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_01"

  var err error

  // multiplier = 2
  originalMultiplierStr := "2"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "2",
    "2",
    "2",
    "2",
    "2",
    "2",
  }

  // product = 128
  originalExpectedNumStr := "128"

  originalExpectedSignValue := 1

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

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := \n"+
      "    multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\nError='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d]))\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
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

  expectedBigINum, err := new(BigIntNum).
    NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).MultiplyDecimalArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
      "    multiplierDecimal, decimalArray)\n"+
      "multiplierDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedSignValue != resultSignValue ||
    originalExpectedSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_02"

  var err error

  // multiplier = 37.9876
  originalMultiplierStr := "37.9876"

  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "-27.9",
    "48.123456",
    "59.48721",
    "-3",
    "19.1",
    "69",
  }

  // product = 11995826664.26376575446779648
  originalExpectedNumStr := "11995826664.26376575446779648"

  originalExpectedSignValue := 1

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

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := \n"+
      "    multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).
      NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d]))\n"+
        "originalMultiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
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

  expectedBigINum, err := new(BigIntNum).
    NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "    NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).
    MultiplyDecimalArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
      "    multiplierDecimal, decimalArray)\n"+
      "multiplierDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedSignValue != resultSignValue ||
    originalExpectedSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_03"

  var err error

  // multiplier = 10.1
  originalMultiplierStr := "10.1"

  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "2",
    "5.8",
    "68.7",
    "3.1234567",
    "8.0",
    "11",
  }

  // product = 2212352.1767579232
  originalExpectedNumStr := "2212352.1767579232"

  originalExpectedSignValue := 1

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

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := \n"+
      "    multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).
      NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d]))\n"+
        "originalMultiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
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

  expectedBigINum, err := new(BigIntNum).
    NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "    NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).
    MultiplyDecimalArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
      "    multiplierDecimal, decimalArray)\n"+
      "multiplierDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedSignValue != resultSignValue ||
    originalExpectedSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_04"
  var err error

  // multiplier = -5.123456
  originalMultiplierStr := "-5.123456"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "1.879",
    "3.824",
    "21.756",
    "2.1234567",
    "6",
    "2",
  }

  // product = -20408.5138429311978576052224
  originalExpectedNumStr := "-20408.5138429311978576052224"

  originalExpectedSignValue := -1

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

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

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := \n"+
      "    multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).
    NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d]))\n"+
        "originalMultiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
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

  expectedBigINum, err := new(BigIntNum).
    NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "    NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).MultiplyDecimalArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
      "    multiplierDecimal, decimalArray)\n"+
      "multiplierDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedSignValue != resultSignValue ||
    originalExpectedSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalArray_05"

  var err error

  // multiplier = -5.123456
  originalMultiplierStr := "-5.123456"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "1.879",
    "3.824",
    "21.756",
    "2.1234567",
    "6",
    "2",
  }

  // product = -20408.5138429311978576052224
  originalExpectedNumStr := "-20408,5138429311978576052224"

  originalExpectedSignValue := -1

  frenchNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  frenchNumSeps.DecimalSeparator = frenchDecSeparator
  frenchNumSeps.ThousandsSeparator = frenchThousandsSeparator
  frenchNumSeps.CurrencySymbol = frenchCurrencySymbol

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

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

  multiplierDecimalNumStr, err := multiplierDecimal.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimalNumStr, err := \n"+
      "    multiplierDecimal.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  err = multiplierDecimal.SetNumericSeparatorsDto(frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierDecimal.SetNumericSeparatorsDto(frenchNumSeps)\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  var ia IntAry
  var iaNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d]))\n"+
        "originalMultiplicandStrs[%d]='%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
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

  err = iaResult.SetNumericSeparatorsDto(frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = iaResult.SetNumericSeparatorsDto(frenchNumSeps)\n"+
      "frenchNumSeps= '%v'\n"+
      "Error= '%v'\n\n",
      ePrefix, frenchNumSeps.String(), err.Error())
    return
  }

  expectedBigINum, err := new(BigIntNum).
    NewNumStrWithNumSeps(originalExpectedNumStr, &frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      " NewNumStrWithNumSeps(originalExpectedNumStr, &frenchNumSeps)\n"+
      "originalExpectedNumStr= '%v'\n"+
      "frenchNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalExpectedNumStr, frenchNumSeps.String(), err.Error())
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

  // multiplierDecimal set with frenchNumSeps
  result, err := new(BigIntMathMultiply).MultiplyDecimalArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalArray(\n"+
      "    multiplierDecimal, decimalArray)\n"+
      "multiplierDecimalNumStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
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

  if originalExpectedNumStr != expectedBigINumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExpectedNumStr != expectedBigINumStr\n"+
      "Expected expectedBigINumStr = '%v'\n"+
      "  Actual expectedBigINumStr = '%v'\n\n",
      ePrefix, originalExpectedNumStr, expectedBigINumStr)

    return
  }

  if originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExpectedNumStr != resultNumStr\n"+
      "Expected resultNumStr = '%v'\n"+
      "  Actual resultNumStr = '%v'\n\n",
      ePrefix, originalExpectedNumStr, resultNumStr)

    return
  }

  if originalExpectedSignValue != resultSignValue {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExpectedSignValue != resultSignValue\n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, originalExpectedSignValue, resultSignValue)

    return
  }

  if originalExpectedSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because originalExpectedSignValue != expectedBigINumSign\n"+
      "Expected expectedBigINumSign = '%v'\n"+
      "  Actual expectedBigINumSign = '%v'\n\n",
      ePrefix, originalExpectedSignValue, expectedBigINumSign)

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
      "\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !frenchNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'.\n\n",
      ePrefix, frenchNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalOutputToArray_01"
  var err error

  // multiplier = 2
  originalMultiplierStr := "2"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
  }

  // Expected Results Array
  originalExpectedNumStrs := []string{
    "2",
    "4",
    "6",
    "8",
    "10",
    "12",
  }

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

  } // End of first for loop

  result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)\n"+
      "multiplierDecimal='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if originalExpectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected  Result NumStr[%v]='%v'.\n"+
        "Instead NumStr[%v]='%v'. ",
        ePrefix, j, originalExpectedNumStrs[j], j, resultNumStr)
    }

  } // End of 2nd for loop

  return
}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalOutputToArray_02"

  var err error

  // multiplier = 8
  originalMultiplierStr := "8"

  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "100.1",
    "-26",
    "3.924",
    "8",
    "5297.123",
    "-4.896",
  }

  // Expected Results Array
  originalExpectedNumStrs := []string{
    "800.8",
    "-208",
    "31.392",
    "64",
    "42376.984",
    "-39.168",
  }

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

  } // End of first for loop

  result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)\n"+
      "multiplierDecimal='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if originalExpectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected  Result NumStr[%v]='%v'.\n"+
        "Instead NumStr[%v]='%v'. ",
        ePrefix, j, originalExpectedNumStrs[j], j, resultNumStr)
    }

  } // End of 2nd for loop

  return
}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalOutputToArray_03"

  var err error

  // multiplier = -31.2
  originalMultiplierStr := "-31.2"

  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "100.1",
    "-26",
    "3.924",
    "8",
    "5297.123",
    "-4.896",
  }

  // Expected Results Array
  originalExpectedNumStrs := []string{
    "-3123.12",
    "811.2",
    "-122.4288",
    "-249.6",
    "-165270.2376",
    "152.7552",
  }

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

  } // End of first for loop

  result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)\n"+
      "multiplierDecimal='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if originalExpectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected  Result NumStr[%v]='%v'.\n"+
        "Instead NumStr[%v]='%v'. ",
        ePrefix, j, originalExpectedNumStrs[j], j, resultNumStr)
    }

  } // End of 2nd for loop

  return
}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalOutputToArray_04"

  var err error

  // multiplier = 283
  originalMultiplierStr := "283"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "0",
    "-26",
    "0",
    "8",
    "5297.123",
    "0",
  }

  // Expected Results Array
  originalExpectedNumStrs := []string{
    "0",
    "-7358",
    "0",
    "2264",
    "1499085.809",
    "0",
  }

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

  } // End of first for loop

  result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)\n"+
      "multiplierDecimal='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if originalExpectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected  Result NumStr[%v]='%v'.\n"+
        "Instead NumStr[%v]='%v'. ",
        ePrefix, j, originalExpectedNumStrs[j], j, resultNumStr)
    }

  } // End of 2nd for loop

}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalOutputToArray_05"

  var err error

  // multiplier = 0
  originalMultiplierStr := "0"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "5",
    "-26",
    "9",
    "8",
    "5297.123",
    "37",
  }

  // Expected Results Array
  originalExpectedNumStrs := []string{
    "0",
    "0",
    "0",
    "0",
    "0",
    "0",
  }

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

  } // End of first for loop

  result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)\n"+
      "multiplierDecimal='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if originalExpectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected  Result NumStr[%v]='%v'.\n"+
        "Instead NumStr[%v]='%v'. ",
        ePrefix, j, originalExpectedNumStrs[j], j, resultNumStr)
    }

  } // End of 2nd for loop

  return
}

func TestBigIntMathMultiply_MultiplyDecimalOutputToArray_06(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalOutputToArray_06"

  var err error

  // multiplier = 283
  originalMultiplierStr := "283"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "0",
    "-26",
    "0",
    "8",
    "5297.123",
    "0",
  }

  // Expected Results Array
  originalExpectedNumStrs := []string{
    "0",
    "-7358",
    "0",
    "2264",
    "1499085,809",
    "0",
  }

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
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

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

  } // End of first for loop

  result, err := new(BigIntMathMultiply).MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).\n"+
      "  MultiplyDecimalOutputToArray(multiplierDecimal, decimalArray)\n"+
      "multiplierDecimal='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, multiplierDecimalNumStr, err.Error())
    return
  }

  var resultNumStr string

  for j := 0; j < lenArray; j++ {

    resultNumStr, err = result[j].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "resultNumStr, err = result[%d].GetNumStr()\n"+
        "Error='%v'\n\n",
        ePrefix, j, err.Error())
      return
    }

    if originalExpectedNumStrs[j] != resultNumStr {
      t.Errorf("%v\n"+
        "Error: Expected  Result NumStr[%v]='%v'.\n"+
        "Instead NumStr[%v]='%v'. ",
        ePrefix, j, originalExpectedNumStrs[j], j, resultNumStr)
    }

    actualNumSeps, err := result[j].GetNumericSeparatorsDto()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "actualNumSeps, err := result[%d].GetNumericSeparatorsDto()\n"+
        "Error='%v'\n\n", ePrefix, j, err.Error())
      return
    }

    if !expectedNumSeps.Equal(actualNumSeps) {
      t.Errorf("%v\n"+
        "Error: Expected NumSeps='%v'.\n"+
        "Instead, NumSeps='%v'. Loop Index='%v'",
        ePrefix, expectedNumSeps.String(), actualNumSeps.String(), j)
      return
    }
  }
}

func TestBigIntMathMultiply_MultiplyDecimalSeries_01(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalSeries_01"
  var err error

  // multiplier = 2
  originalMultiplierStr := "2"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "2",
    "2",
    "2",
    "2",
    "2",
    "2",
  }

  // product = 128
  originalExpectedNumStr := "128"

  originalExpectedNumSignValue := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected iaNumStr='%v'.\n"+
      "Instead, iaNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, iaNumStr)
    return
  }

  var ia IntAry
  var decimalNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    decimalNumStr, err = decimalArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " decimalNumStr, err = decimalArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    if decimalNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected decimalNumStr='%v'.\n"+
        "Instead, decimalNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], decimalNumStr)
      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        err.Error())

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

    if iaNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected iaNumStr='%v'.\n"+
        "Instead, iaNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], iaNumStr)
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia= '%v'\n"+
        "Loop Index= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        iaNumStr,
        i,
        err.Error())

      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
    multiplierDecimal,
    decimalArray[0],
    decimalArray[1],
    decimalArray[2],
    decimalArray[3],
    decimalArray[4],
    decimalArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(...)\n"+
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedNumSignValue != resultSignValue ||
    originalExpectedNumSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedNumSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalSeries_02(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalSeries_02"

  var err error

  // multiplier = 37.9876
  originalMultiplierStr := "37.9876"

  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "-27.9",
    "48.123456",
    "59.48721",
    "-3",
    "19.1",
    "69",
  }

  // product = 11995826664.26376575446779648
  originalExpectedNumStr := "11995826664.26376575446779648"

  originalExpectedNumSignValue := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected iaNumStr='%v'.\n"+
      "Instead, iaNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, iaNumStr)
    return
  }

  var ia IntAry
  var decimalNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    decimalNumStr, err = decimalArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " decimalNumStr, err = decimalArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    if decimalNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected decimalNumStr='%v'.\n"+
        "Instead, decimalNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], decimalNumStr)
      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        err.Error())

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

    if iaNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected iaNumStr='%v'.\n"+
        "Instead, iaNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], iaNumStr)
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia= '%v'\n"+
        "Loop Index= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        iaNumStr,
        i,
        err.Error())

      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
    multiplierDecimal,
    decimalArray[0],
    decimalArray[1],
    decimalArray[2],
    decimalArray[3],
    decimalArray[4],
    decimalArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(...)\n"+
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedNumSignValue != resultSignValue ||
    originalExpectedNumSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedNumSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalSeries_03(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalSeries_03"

  var err error

  // multiplier = 10.1
  originalMultiplierStr := "10.1"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "2",
    "5.8",
    "68.7",
    "3.1234567",
    "8.0",
    "11",
  }

  // product = 2212352.1767579232
  originalExpectedNumStr := "2212352.1767579232"

  originalExpectedNumSignValue := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected iaNumStr='%v'.\n"+
      "Instead, iaNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, iaNumStr)
    return
  }

  var ia IntAry
  var decimalNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    decimalNumStr, err = decimalArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " decimalNumStr, err = decimalArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    if decimalNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected decimalNumStr='%v'.\n"+
        "Instead, decimalNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], decimalNumStr)
      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        err.Error())

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

    if iaNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected iaNumStr='%v'.\n"+
        "Instead, iaNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], iaNumStr)
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia= '%v'\n"+
        "Loop Index= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        iaNumStr,
        i,
        err.Error())

      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
    multiplierDecimal,
    decimalArray[0],
    decimalArray[1],
    decimalArray[2],
    decimalArray[3],
    decimalArray[4],
    decimalArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(...)\n"+
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedNumSignValue != resultSignValue ||
    originalExpectedNumSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedNumSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalSeries_04(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalSeries_04"

  var err error

  // multiplier = -5.123456
  originalMultiplierStr := "-5.123456"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "1.879",
    "3.824",
    "21.756",
    "2.1234567",
    "6",
    "2",
  }

  // product = -20408.5138429311978576052224
  originalExpectedNumStr := "-20408.5138429311978576052224"

  originalExpectedNumSignValue := -1

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected iaNumStr='%v'.\n"+
      "Instead, iaNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, iaNumStr)
    return
  }

  var ia IntAry
  var decimalNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    decimalNumStr, err = decimalArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " decimalNumStr, err = decimalArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    if decimalNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected decimalNumStr='%v'.\n"+
        "Instead, decimalNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], decimalNumStr)
      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        err.Error())

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

    if iaNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected iaNumStr='%v'.\n"+
        "Instead, iaNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], iaNumStr)
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia= '%v'\n"+
        "Loop Index= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        iaNumStr,
        i,
        err.Error())

      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
    multiplierDecimal,
    decimalArray[0],
    decimalArray[1],
    decimalArray[2],
    decimalArray[3],
    decimalArray[4],
    decimalArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(...)\n"+
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedNumSignValue != resultSignValue ||
    originalExpectedNumSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedNumSignValue,
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

func TestBigIntMathMultiply_MultiplyDecimalSeries_05(t *testing.T) {

  ePrefix := "TestBigIntMathMultiply_MultiplyDecimalSeries_05"

  var err error

  // multiplier = 10.1
  originalMultiplierStr := "10.1"
  // multiplicandStrs
  originalMultiplicandStrs := []string{
    "2",
    "5.8",
    "68.7",
    "3.1234567",
    "8.0",
    "11",
  }

  // product = 2212352,1767579232
  originalExpectedNumStr := "2212352,1767579232"

  originalExpectedNumSignValue := 1

  frenchNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  frenchNumSeps.DecimalSeparator = frenchDecSeparator
  frenchNumSeps.ThousandsSeparator = frenchThousandsSeparator
  frenchNumSeps.CurrencySymbol = frenchCurrencySymbol

  multiplierDecimal, err := new(Decimal).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "multiplierDecimal, err := new(Decimal).\n"+
      "    NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
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

  if multiplierDecimalNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected MultiplierDecimalNumStr='%v'.\n"+
      "Instead, MultiplierDecimalNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, multiplierDecimalNumStr)
    return
  }

  err = multiplierDecimal.SetNumericSeparatorsDto(frenchNumSeps)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "err = multiplierDecimal.SetNumericSeparatorsDto(frenchNumSeps)\n"+
      "frenchNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, frenchNumSeps.String(), err.Error())
    return
  }

  lenArray := len(originalMultiplicandStrs)

  if lenArray == 0 {
    t.Errorf("%v\n"+
      "Fatal Error: 'originalMultiplicandStrs' is an empty array!\n\n",
      ePrefix)
    return
  }

  decimalArray := make([]Decimal, lenArray)

  iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(originalMultiplierStr)\n"+
      "originalMultiplierStr='%v'\n"+
      "Error='%v'\n\n",
      ePrefix, originalMultiplierStr, err.Error())
    return
  }

  iaNumStr, err := iaResult.GetNumStr()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "iaNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if iaNumStr != originalMultiplierStr {
    t.Errorf("%v\n"+
      "Error: Expected iaNumStr='%v'.\n"+
      "Instead, iaNumStr='%v'.\n\n",
      ePrefix, originalMultiplierStr, iaNumStr)
    return
  }

  var ia IntAry
  var decimalNumStr string

  for i := 0; i < lenArray; i++ {

    decimalArray[i], err = new(Decimal).NewNumStr(originalMultiplicandStrs[i])

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "decimalArray[%d], err = new(Decimal).\n"+
        "    NewNumStr(originalMultiplicandStrs[%d])\n"+
        "originalMultiplicandStrs[%d]= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i, i, i,
        originalMultiplicandStrs[i],
        err.Error())

      return
    }

    decimalNumStr, err = decimalArray[i].GetNumStr()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        " decimalNumStr, err = decimalArray[%d].GetNumStr()\n"+
        "Error='%v'\n\n", ePrefix, i, err.Error())
      return
    }

    if decimalNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected decimalNumStr='%v'.\n"+
        "Instead, decimalNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], decimalNumStr)
      return
    }

    ia, err = decimalArray[i].GetIntAry()

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "ia, err = decimalArray[%d].GetIntAry()\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        err.Error())

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

    if iaNumStr != originalMultiplicandStrs[i] {
      t.Errorf("%v\n"+
        "Error: Expected iaNumStr='%v'.\n"+
        "Instead, iaNumStr='%v'.\n\n",
        ePrefix, originalMultiplicandStrs[i], iaNumStr)
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("%v\n"+
        "Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "ia= '%v'\n"+
        "Loop Index= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        iaNumStr,
        i,
        err.Error())

      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStr(originalExpectedNumStr)

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(originalExpectedNumStr)\n"+
      "originalExpectedNumStr='%v'\nError='%v'\n\n",
      ePrefix, originalExpectedNumStr, err.Error())
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

  result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(
    multiplierDecimal,
    decimalArray[0],
    decimalArray[1],
    decimalArray[2],
    decimalArray[3],
    decimalArray[4],
    decimalArray[5])

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyDecimalSeries(...)\n"+
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

  if originalExpectedNumStr != expectedBigINumStr ||
    originalExpectedNumStr != resultNumStr {
    t.Errorf("%v\n"+
      "Error: BigIntNum Number Strings are unequal!\n"+
      "         Expected Number String = '%s'.\n"+
      "           Result Number String = '%s'.\n"+
      "Original Expected Number String = '%s'",
      ePrefix,
      expectedBigINumStr,
      resultNumStr,
      originalExpectedNumStr)
    return

  }

  if originalExpectedNumSignValue != resultSignValue ||
    originalExpectedNumSignValue != expectedBigINumSign {
    t.Errorf("%v\n"+
      "Error: Expected Number Signs do NOT match!\n"+
      "originalExpectedBigINumSign sign='%v'.\n"+
      "            resultSignValue sign='%v'\n"+
      "             expectedBigINumSign='%v'\n",
      ePrefix,
      originalExpectedNumSignValue,
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

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("%v\n"+
      "Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", ePrefix, err.Error())
    return
  }

  if !frenchNumSeps.Equal(actualNumSeps) {
    t.Errorf("%v\n"+
      "Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'.\n\n",
      ePrefix, frenchNumSeps.String(), actualNumSeps.String())
  }

  return
}
