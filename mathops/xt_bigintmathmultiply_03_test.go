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

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_03(t *testing.T) {

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStrs[i]) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

    ia, err := bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.CmpBigInt(result) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
  }

  actualNumStr := result.GetNumStr()

  iaResult.OptimizeIntArrayLen(true)

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_04(t *testing.T) {

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

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStrs[i]) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

    ia, err := bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
  }

  result := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  if !expectedBigINum.Equal(result) {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
  }

  if expectedBigINum.CmpBigInt(result) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
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

}

func TestBigIntMathMultiply_MultiplyBigIntNumSeries_05(t *testing.T) {

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
  expectedNumStr := "2212352,1767579232"

  multiplierBiNum, err := new(BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
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
  bINumArray := make([]BigIntNum, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by new(BigIntNum).NewNumStr(multiplicandStrs[i]) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

    ia, err := bINumArray[i].GetIntAry()

    if err != nil {
      t.Errorf("Error returned by bINumArray[i].GetIntAryElements() "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  result := new(BigIntMathMultiply).MultiplyBigIntNumSeries(
    multiplierBiNum,
    bINumArray[0],
    bINumArray[1],
    bINumArray[2],
    bINumArray[3],
    bINumArray[4],
    bINumArray[5])

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}

func TestBigIntMathMultiply_MultiplyDecimal_01(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedNumStr := "2875.94572"

  expectedSignValue := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandDecimal, err := new(Decimal).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedDecimal, err := new(Decimal).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(expectedNumStr) "+
      "expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimal"+
      "(multiplierDecimal, multiplicandDecimal) "+
      "multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
      multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
  }

  if expectedDecimal.GetNumStr() != result.GetNumStr() {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedDecimal.GetBigInt() "+
      "Error='%v'. ",
      err.Error())
  }

  if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  if expectedSignValue != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyDecimal_02(t *testing.T) {
  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedNumStr := "30987680500513189125.14259702468435"

  expectedSignValue := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandDecimal, err := new(Decimal).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedDecimal, err := new(Decimal).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(expectedNumStr) "+
      "expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimal"+
      "(multiplierDecimal, multiplicandDecimal) "+
      "multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
      multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
  }

  if expectedDecimal.GetNumStr() != result.GetNumStr() {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedDecimal.GetBigInt() "+
      "Error='%v'. ",
      err.Error())
  }

  if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  if expectedSignValue != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyDecimal_03(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedNumStr := "-30987680500513189125.14259702468435"

  expectedSignValue := -1

  multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandDecimal, err := new(Decimal).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedDecimal, err := new(Decimal).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(expectedNumStr) "+
      "expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimal"+
      "(multiplierDecimal, multiplicandDecimal) "+
      "multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
      multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
  }

  if expectedDecimal.GetNumStr() != result.GetNumStr() {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedDecimal.GetBigInt() "+
      "Error='%v'. ",
      err.Error())
  }

  if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  if expectedSignValue != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyDecimal_04(t *testing.T) {
  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedNumStr := "22197234145.3632"

  expectedSignValue := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandDecimal, err := new(Decimal).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedDecimal, err := new(Decimal).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(expectedNumStr) "+
      "expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimal"+
      "(multiplierDecimal, multiplicandDecimal) "+
      "multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
      multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
  }

  if expectedDecimal.GetNumStr() != result.GetNumStr() {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedDecimal.GetBigInt() "+
      "Error='%v'. ",
      err.Error())
  }

  if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  if expectedSignValue != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, result.sign)
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
  }

}

func TestBigIntMathMultiply_MultiplyDecimal_05(t *testing.T) {
  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0.00
  expectedNumStr := "0"

  expectedSignValue := 1

  multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  multiplicandDecimal, err := new(Decimal).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  iaResult := new(IntAry).New()

  err = iaMultiplier.Multiply(
    &iaMultiplier,
    &iaMultiplicand,
    &iaResult,
    -1,
    -1)

  if err != nil {
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
  }

  expectedDecimal, err := new(Decimal).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(expectedNumStr) "+
      "expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimal"+
      "(multiplierDecimal, multiplicandDecimal) "+
      "multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
      multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
  }

  if expectedDecimal.GetNumStr() != result.GetNumStr() {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  expectedDecimalBigInt, err := expectedDecimal.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedDecimal.GetBigInt() "+
      "Error='%v'. ",
      err.Error())
  }

  if expectedDecimalBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedDecimal.GetNumStr(), result.GetNumStr())
  }

  if expectedSignValue != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, result.sign)
  }

  actualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by result.GetBigInt() "+
      "Error='%v'. ", err.Error())
  }

  iaBigInt, err := iaResult.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by iaResult.GetBigInt() "+
      "Error='%v'. ", err.Error())
  }

  if actualBigInt.Cmp(iaBigInt) != 0 {
    t.Errorf("Error: Expected actualBigInt='%v' "+
      "Instead, actualBigInt='%v'",
      iaResult.GetNumStr(), actualBigInt)
  }

}

func TestBigIntMathMultiply_MultiplyDecimal_06(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedNumStr := "2875,94572"

  multiplierDecimal, err := new(Decimal).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
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
    t.Errorf("Error returned by multiplierBiNum.SetNumericSeparatorsDto(expectedNumSeps). "+
      "Error='%v' ", err.Error())
  }

  multiplicandDecimal, err := new(Decimal).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(Decimal).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
  }

  result, err := new(BigIntMathMultiply).MultiplyDecimal(multiplierDecimal, multiplicandDecimal)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyDecimal"+
      "(multiplierDecimal, multiplicandDecimal) "+
      "multiplierDecimal='%v' multiplicandDecimal='%v' Error='%v'. ",
      multiplierDecimal.GetNumStr(), multiplicandDecimal.GetNumStr(), err.Error())
  }

  actualNumStr := result.GetNumStr()

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'. Instead, NumStr='%v'. ",
      expectedNumStr, actualNumStr)
  }

  actualNumSeps := result.GetNumericSeparatorsDto()

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'. ",
      expectedNumSeps.String(), actualNumSeps.String())
  }

}
