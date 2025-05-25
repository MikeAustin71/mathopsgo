package mathops

import (
  "math/big"
  "testing"
)

func TestBigIntMathMultiply_MultiplyNumStrSeries_01(t *testing.T) {

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
    t.Errorf("Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
      "multiplierStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      multiplierStr, expectedNumSeps.String(), err.Error())
    return
  }

  var ia IntAry
  var expectedBigINum, result BigIntNum

  for i := 0; i < lenArray; i++ {

    ia, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
        "multiplicandStrs[%d], expectedNumSeps)\n"+
        "multiplicandStrs[%v]='%v'\n"+
        "expectedNumSeps='%v'\n"+
        "Error='%v'\n\n",
        i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }

  } // end of For loop

  expectedBigINum, err = new(BigIntNum).NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:"+
      "expectedBigINum, err = new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumStr,\n"+
      "    expectedNumSeps, expectedNumSeps)\n"+
      "expectedNumStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      expectedNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  result, err = new(BigIntMathMultiply).MultiplyNumStrSeries(
    expectedNumSeps,
    expectedNumSeps,
    multiplierStr,
    multiplicandStrs[0],
    multiplicandStrs[1],
    multiplicandStrs[2],
    multiplicandStrs[3],
    multiplicandStrs[4],
    multiplicandStrs[5])

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, err = new(BigIntMathMultiply).MultiplyNumStrSeries(\n"+
      "     9-parameters)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  var isEqualToResult bool

  isEqualToResult, err = expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "isEqualToResult, err = expectedBigINum.Equal(result)\n"+
      "Error='%v'", err.Error())
    return
  }

  if !isEqualToResult {
    t.Errorf("Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  var resultBigInt *big.Int

  resultBigInt, err = result.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultBigInt, err = result.GetBigInt()\n"+
      "Error='%v'", err.Error())
    return
  }

  if expectedBigINum.bigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))

    return
  }

  var resultSignValue int

  resultSignValue, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultSignValue, err = result.GetSign()\n"+
      "Error='%v'", err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, resultSignValue)

    return
  }

  var actualNumStr string

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
      "Error='%v'", err.Error())

    return
  }

  expectedNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'", err.Error())

    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      expectedNumStr, actualNumStr)
    return
  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'", err.Error())

    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_02(t *testing.T) {

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
    t.Errorf("Error returned by:"+
      "iaResult, err := new(IntAry).\n"+
      "  NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
      "multiplierStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      multiplierStr, expectedNumSeps.String(), err.Error())

    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    ia, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
        "  multiplicandStrs[%d], expectedNumSeps)\n"+
        "multiplicandStrs[%v]='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
        i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())

      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "i='%v'; multiplicandStrs[i]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of for loop

  expectedBigINum, err := new(BigIntNum).
    NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "expectedNumStr, expectedNumSeps, expectedNumSeps)\n"+
      "expectedNumStr='%v' expectedNumSeps='%v'\nError='%v'\n\n",
      expectedNumStr, expectedNumSeps.String(), err.Error())

    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
    expectedNumSeps,
    expectedNumSeps,
    multiplierStr,
    multiplicandStrs[0],
    multiplicandStrs[1],
    multiplicandStrs[2],
    multiplicandStrs[3],
    multiplicandStrs[4],
    multiplicandStrs[5])

  if err != nil {
    t.Errorf("Error returned by new(IntAry).MultiplyNumStrSeries(...)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      expectedBigINumSign, result.sign)
    return
  }

  var actualNumStr string

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v'\n"+
      "Instead, actualNumStr='%v'\n\n",
      expectedNumStr, actualNumStr)
    return
  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'. Instead, NumSeps='%v'",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_03(t *testing.T) {

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

  expectedNumSeps := NumericSeparatorDto{}

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStrWithNumSeps(\n"+
      "  multiplierStr, expectedNumSeps)\n"+
      "multiplierStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      multiplierStr, expectedNumSeps.String(), err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    ia, err = new(IntAry).
      NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
        "  multiplicandStrs[%d], expectedNumSeps)\n"+
        "multiplicandStrs[%v]='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
        i, i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of for loop

  var expectedBigINum BigIntNum

  expectedBigINum, err = new(BigIntNum).
    NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err = new(BigIntNum).\n"+
      "  NewNumStrWithNumSeps(expectedNumStr,\n"+
      "    expectedNumSeps, expectedNumSeps)\n"+
      "expectedNumStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      expectedNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
    expectedNumSeps,
    expectedNumSeps,
    multiplierStr,
    multiplicandStrs[0],
    multiplicandStrs[1],
    multiplicandStrs[2],
    multiplicandStrs[3],
    multiplicandStrs[4],
    multiplicandStrs[5])

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(\n"+
      "         9-parameters)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumEqualResult, err := expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  var resultBigInt *big.Int

  resultBigInt, err = result.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultBigInt, err = result.GetBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINum.bigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'\n"+
      "Instead, Decimal= '%s'\n\n",
      expectedBigINum.bigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
    return
  }

  var actualNumStr string

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  expectedNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      expectedNumStr, actualNumStr)
    return
  }

  expectedNumSeps.SetDefaultsIfEmpty()

  var actualNumSeps NumericSeparatorDto

  actualNumSeps, err = result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumSeps, err = result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'\n\n",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_04(t *testing.T) {

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

  expectedBigINumSign := -1

  expectedNumSeps := NumericSeparatorDto{}
  expectedNumSeps.DecimalSeparator = '.'
  expectedNumSeps.ThousandsSeparator = ','
  expectedNumSeps.CurrencySymbol = '$'

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).
    NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "    NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
      "multiplierStr='%v'; expectedNumSeps='%v'\nError='%v'\n\n",
      multiplierStr, expectedNumSeps.String(), err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    ia, err = new(IntAry).
      NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "new(IntAry).NewNumStrWithNumSeps("+
        "multiplicandStrs[i], expectedNumSeps)\n"+
        "multiplicandStrs[%v]='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "i='%v'; multiplicandStrs[i]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }
  } // End of for loop

  expectedBigINum, err := new(BigIntNum).
    NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "NewNumStrWithNumSeps(expectedNumStr, expectedNumSeps,\n"+
      "    expectedNumSeps)\n"+
      "expectedNumStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      expectedNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
    expectedNumSeps,
    expectedNumSeps,
    multiplierStr,
    multiplicandStrs[0],
    multiplicandStrs[1],
    multiplicandStrs[2],
    multiplicandStrs[3],
    multiplicandStrs[4],
    multiplicandStrs[5])

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(...)\n"+
      "multiplierStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      multiplierStr, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumEqualResult, err :=\n"+
      "   expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("Error: Expected Decimal='%s'.\nInstead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  var resultSignValue int

  resultSignValue, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultSignValue, err = result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'.\n\n",
      expectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStr, err := iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'\n\n",
      expectedNumStr, actualNumStr)
    return
  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumSeps, err := result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'.\n\n",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrSeries_05(t *testing.T) {

  var err error

  // multiplier = -5,123456
  multiplierStr := "-5,123456"
  // multiplicandStrs
  multiplicandStrs := []string{
    "1,879",
    "3,824",
    "21,756",
    "2,1234567",
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

  lenArray := len(multiplicandStrs)

  iaResult, err := new(IntAry).NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResult, err := new(IntAry)."+
      "  NewNumStrWithNumSeps(multiplierStr, expectedNumSeps)\n"+
      "multiplierStr='%v' expectedNumSeps='%v'\nError='%v'\n\n",
      multiplierStr, expectedNumSeps.String(), err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    ia, err = new(IntAry).NewNumStrWithNumSeps(multiplicandStrs[i], expectedNumSeps)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "ia, err = new(IntAry).NewNumStrWithNumSeps(\n"+
        "    multiplicandStrs[i], expectedNumSeps)\n"+
        "multiplicandStrs[%v]='%v'\nexpectedNumSeps='%v'\n"+
        "Error='%v'\n\n",
        i, multiplicandStrs[i], expectedNumSeps.String(), err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }
  } // End of for loop

  expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(
    expectedNumStr, expectedNumSeps, expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).NewNumStrWithNumSeps(\n"+
      "  expectedNumStr, expectedNumSeps, expectedNumSeps)\n"+
      "expectedNumStr='%v'; expectedNumSeps='%v'\nError='%v'\n\n",
      expectedNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(
    expectedNumSeps,
    expectedNumSeps,
    multiplierStr,
    multiplicandStrs[0],
    multiplicandStrs[1],
    multiplicandStrs[2],
    multiplicandStrs[3],
    multiplicandStrs[4],
    multiplicandStrs[5])

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrSeries(...)\n"+
      "expectedNumStr='%v'\nexpectedNumSeps='%v'\nError='%v'\n\n",
      expectedNumStr, expectedNumSeps.String(), err.Error())
    return
  }

  expectedBigINumEqualResult, err := expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumEqualResult, err := \n"+
      "    expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedBigINumEqualResult {
    t.Errorf("Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected Decimal='%s'.\n"+
      "Instead, Decimal= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      expectedBigINumSign, result.sign)
    return
  }

  actualNumStr, err := result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v'\n"+
      "Instead, actualNumStr='%v'\n\n",
      expectedNumStr, actualNumStr)
    return
  }

  actualNumSeps, err := result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumSeps, err := \n"+
      "  result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'.\n\n",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_01(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875.94572
  expectedNumStr := "2875.94572"

  expectedSignValue := 1

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
    return
  }

  multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'  Error='%v'. ", multiplicandStr, err.Error())
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
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
    return
  }

  expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr) "+
      "expectedNumStr='%v'  Error='%v'. ", expectedNumStr, err.Error())
    return
  }

  var multiplierNumStr, multiplicandNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'. \n\n", err.Error())
    return
  }

  multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplicandNumStrDto.GetNumStr()\n"+
      "Error='%v'. \n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto"+
      "(multiplierNumStrDto, multiplicandNumStrDto) "+
      "multiplierNumStrDto='%v' multiplicandNumStrDto='%v' Error='%v'. ",
      multiplierNumStr, multiplicandNumStr, err.Error())
    return
  }

  var resultNumStr string

  expectedNumStr, err = expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by resultNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != resultNumStr {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, resultNumStr)
    return
  }

  expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, resultNumStr)
    return
  }

  var signVal int

  signVal, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by result.GetSign()\n"+
      "Error='%v'\n\n ", err.Error())
    return
  }

  if expectedSignValue != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, signVal)
    return
  }

  var actualNumStr string

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n\n ", err.Error())
    return
  }

  var iaResultNumStr string

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n\n ", err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_02(t *testing.T) {
  // multiplier = 57638422123.327890123
  multiplierStr := "57638422123.327890123"

  // multiplicand = 537621943.12345
  multiplicandStr := "537621943.12345"

  // product = 30987680500513189125.14259702468435
  expectedNumStr := "30987680500513189125.14259702468435"

  expectedSignValue := 1

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'\nError='%v'\n", multiplierStr, err.Error())
    return
  }

  multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'\nError='%v'\n\n", multiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr) "+
      "multiplicandStr='%v'\nError='%v'\n\n", multiplicandStr, err.Error())
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
    t.Errorf("Error returned by iaMultiplier.Multiply() "+
      "Error='%v'. ", err.Error())
    return
  }

  expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\nError='%v'\n\n", expectedNumStr, err.Error())
    return
  }

  var multiplierNumStr, multiplicandNumStr, resultNumStr,
    actualNumStr, iaResultNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplicandNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto(\n"+
      "  multiplierNumStrDto, multiplicandNumStrDto)\n"+
      "multiplierNumStrDto='%v'; multiplicandNumStrDto='%v'\nError='%v'\n\n",
      multiplierNumStr, multiplicandNumStr, err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedNumStr, err = expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != resultNumStr {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, resultNumStr)
    return
  }

  expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetBigInt()\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, resultNumStr)
    return
  }

  var resultSignVal int

  resultSignVal, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by result.GetSign()\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  if expectedSignValue != resultSignVal {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, resultSignVal)
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by iaResult.GetNumStr()\n"+
      "Error='%v'\n",
      err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_03(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "57638422123.327890123"

  // multiplicand = -537621943.12345
  multiplicandStr := "-537621943.12345"

  // product = -30987680500513189125.14259702468435
  expectedNumStr := "-30987680500513189125.14259702468435"

  expectedSignValue := -1

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n", multiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n", multiplicandStr, err.Error())
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
    t.Errorf("Error returned by iaMultiplier.Multiply()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\nError='%v'\n", expectedNumStr, err.Error())
    return
  }

  var result BigIntNum

  var multiplilerNumStr, multiplicandNumStr, actualNumStr,
    resultNumStr, iaResultNumStr string

  multiplilerNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplicandNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err = new(BigIntMathMultiply).
    MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).MultiplyNumStrDto(\n"+
      "multiplierNumStrDto, multiplicandNumStrDto)\n"+
      "multiplierNumStrDto='%v'; multiplicandNumStrDto='%v'\nError='%v'\n",
      multiplilerNumStr, multiplicandNumStr, err.Error())
    return
  }

  expectedNumStr, err = expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedNumStr, err = expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != resultNumStr {
    t.Errorf("Error: Expected BigIntNum='%s'.\nInstead, BigIntNum= '%s'\n\n",
      expectedNumStr, resultNumStr)
    return
  }

  expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetBigInt()\n"+
      "Error='%v'\n\n",
      err.Error())
    return
  }

  if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, resultNumStr)
    return
  }

  var resultSignVal int

  resultSignVal, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by result.GetSign()\n"+
      "Error='%v'\n\n",
      err.Error())
    return
  }

  if expectedSignValue != resultSignVal {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, result.sign)
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n\n",
      err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n\n",
      err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_04(t *testing.T) {
  // multiplier = 89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = -247632
  multiplicandStr := "-247632"

  // product = 22197234145.3632
  expectedNumStr := "22197234145.3632"

  expectedSignValue := 1

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n", multiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
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
    t.Errorf("Error returned by iaMultiplier.Multiply()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).\n"+
      "NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\nError='%v'\n\n",
      expectedNumStr, err.Error())
    return
  }

  var multiplierNumStr, multiplicandNumStr, actualNumStr,
    resultNumStr, iaResultNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by multiplicandNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

  if err != nil {
    t.Errorf("Error returned by new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)\n"+
      "multiplierNumStrDto='%v'\nmultiplicandNumStrDto='%v'\nError='%v'\n\n",
      multiplierNumStr, multiplicandNumStr, err.Error())
    return
  }

  expectedNumStr, err = expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != resultNumStr {
    t.Errorf("Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      expectedNumStr, resultNumStr)
    return
  }

  expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by expectedNumStrDto.GetBigInt() "+
      "Error='%v'. ",
      err.Error())
    return
  }

  if expectedNumStrDtoBigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      expectedNumStr, resultNumStr)
    return
  }

  var resultSignVal int

  resultSignVal, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedSignValue != resultSignVal {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedSignValue, resultSignVal)
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by\n"+
      "actualNumStr, err= result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by\n"+
      "iaResultNumStr, err = result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_05(t *testing.T) {
  // multiplier = -89637.9876
  multiplierStr := "-89637.9876"

  // multiplicand = 0.00
  multiplicandStr := "0.00"

  // product = 0
  expectedNumStr := "0"

  expectedSignValue := 1

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).\n"+
      "NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by new(NumStrDto).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
    return
  }

  iaMultiplier, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplier, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  iaMultiplicand, err := new(IntAry).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaMultiplicand, err := new(IntAry).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
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
    t.Errorf("Error returned by:\n"+
      "  iaMultiplier.Multiply()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  expectedNumStrDto, err := new(NumStrDto).NewNumStr(expectedNumStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(expectedNumStr)\n"+
      "expectedNumStr='%v'\nError='%v'\n\n",
      expectedNumStr, err.Error())
    return
  }

  var multiplierNumStr, multiplicandNumStr,
    resultNumStr, iaResultNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStr, err = multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrDto(\n"+
      "  multiplierNumStrDto, multiplicandNumStrDto)\n"+
      "multiplierNumStrDto='%v'\nmultiplicandNumStrDto='%v'\n"+
      "Error='%v'\n\n",
      multiplierNumStr, multiplicandNumStr, err.Error())
  }

  expectedNumStr, err = expectedNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStr, err = expectedNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStrDto.GetNumStr() != result.GetNumStr() {
    t.Errorf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedNumStr, resultNumStr)
    return
  }

  expectedNumStrDtoBigInt, err := expectedNumStrDto.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumStrDtoBigInt, err := \n"+
      "  expectedNumStrDto.GetBigInt()\n"+
      "Error='%v'\n\n",
      err.Error())
    return
  }

  var resultBigInt *big.Int

  resultBigInt, err = result.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultBigInt, err = result.GetBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStrDtoBigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("Comparison Error: Expected BigIntNum='%s'.\n"+
      "Instead, BigIntNum= '%s'.\n\n",
      expectedNumStr, resultNumStr)
    return
  }

  var resultSignVal int

  resultSignVal, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultSignVal, err = result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedSignValue != resultSignVal {
    t.Errorf("Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      expectedSignValue, resultSignVal)
    return
  }

  actualBigInt, err := result.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by:"+
      "actualBigInt, err := result.GetBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  iaBigInt, err := iaResult.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaBigInt, err := iaResult.GetBigInt()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResultNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if actualBigInt.Cmp(iaBigInt) != 0 {
    t.Errorf("Error: Expected actualBigInt='%v'\n"+
      "Instead, actualBigInt='%v'\n\n",
      iaResultNumStr, actualBigInt)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_07(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875,94572
  expectedNumStr := "2875,94572"

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "err = multiplierNumStrDto.SetNumericSeparatorsDto(\n"+
      "  expectedNumSeps)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  var multiplierNumStr, multiplicandNumStr, actualNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStrDto, err = multiplierNumStrDto.\n"+
      "  GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandNumStrDto, err = multiplicandNumStrDto.\n"+
      "  GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDto(multiplierNumStrDto, multiplicandNumStrDto)

  if err != nil {
    t.Errorf("Error returned by:"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrDto(\n"+
      "  multiplierNumStrDto, multiplicandNumStrDto)\n"+
      "multiplierNumStrDto='%v'\nmultiplicandNumStrDto='%v'\nError='%v'.\n\n",
      multiplierNumStr, multiplicandNumStr, err.Error())
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'.\n"+
      "Instead, NumStr='%v'\n\n",
      expectedNumStr, actualNumStr)
    return
  }

  var actualNumSeps NumericSeparatorDto

  actualNumSeps, err = result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumSeps, err = result.\n"+
      "  GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'.",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDto_06(t *testing.T) {
  // multiplier = 123.32
  multiplierStr := "123.32"

  // multiplicand = 23.321
  multiplicandStr := "23.321"

  // product = 2875,94572
  expectedNumStr := "2875,94572"

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  multiplicandNumStrDto, err := new(NumStrDto).NewNumStr(multiplicandStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplicandStr)\n"+
      "multiplicandStr='%v'\nError='%v'\n\n",
      multiplicandStr, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  expectedNumSepsCopy := NumericSeparatorDto{}

  expectedNumSepsCopy, err = expectedNumSeps.CopyOut(false)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedNumSepsCopy, err = expectedNumSeps.\n"+
      "  CopyOut(false)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  usaNumSeps := NumericSeparatorDto{}

  usaNumSeps.DecimalSeparator = '.'
  usaNumSeps.ThousandsSeparator = ','
  usaNumSeps.CurrencySymbol = '$'

  err = multiplierNumStrDto.SetNumericSeparatorsDto(usaNumSeps)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "err = multiplierNumStrDto.SetNumericSeparatorsDto(\n"+
      "  expectedNumSeps)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  var multiplierNumStr, multiplicandNumStr, actualNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStrDto, err = multiplierNumStrDto.\n"+
      "  GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  multiplicandNumStr, err = multiplicandNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplicandNumStrDto, err = multiplicandNumStrDto.\n"+
      "  GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDto(
    multiplierNumStrDto, multiplicandNumStrDto, expectedNumSepsCopy)

  if err != nil {
    t.Errorf("Error returned by:"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrDto(\n"+
      "  multiplierNumStrDto, multiplicandNumStrDto, expectedNumSepsCopy)\n"+
      "multiplierNumStrDto='%v'\nmultiplicandNumStrDto='%v'\n"+
      "expectedNumSepsCopy='%v'\nError='%v'.\n\n",
      multiplierNumStr, multiplicandNumStr, expectedNumSepsCopy.String(), err.Error())
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v'.\n"+
      "Instead, NumStr='%v'\n\n",
      expectedNumStr, actualNumStr)
    return
  }

  var actualNumSeps NumericSeparatorDto

  actualNumSeps, err = result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumSeps, err = result.\n"+
      "  GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'.\n"+
      "Instead, NumSeps='%v'.",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_01(t *testing.T) {

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

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'.\n\n",
      multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  nDtoArray := make([]NumStrDto, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "iaResult, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  for i := 0; i < lenArray; i++ {

    nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by:"+
        "nDtoArray[%d], err = new(NumStrDto).NewNumStr(multiplicandStrs[%d]\n)"+
        "i='%v';multiplicandStrs[i]='%v'\nError='%v'.\n\n",
        i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err := new(IntAry).NewNumStrDto(nDtoArray[i])

    if err != nil {
      t.Errorf("Error returned by:"+
        "ia, err := new(IntAry).NewNumStrDto(nDtoArray[i])\n"+
        "i='%v'; multiplicandStrs[i]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }
  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      expectedBigINumStr, err.Error())
    return
  }

  var multiplierNumStr, actualNumStr, iaResultNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStr, err = multiplierNumStrDto.\n"+
      "  GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

  if err != nil {
    t.Errorf("Error returned by:"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(\n"+
      "  multiplierNumStrDto, nDtoArray)\n"+
      "multiplierNumStrDto='%v'\nError='%v'\n\n",
      multiplierNumStr, err.Error())
  }

  var expectedBigINumIsEqualToResult bool

  expectedBigINumIsEqualToResult, err = expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumIsEqualToResult, err = \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedBigINumIsEqualToResult {
    t.Errorf("Error: Expected NumStrDto='%s'.\n"+
      "Instead, NumStrDto= '%s'\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  var resultBigInt *big.Int

  resultBigInt, err = result.GetBigInt()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumIsEqualToResult, err = \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINum.bigInt.Cmp(resultBigInt) != 0 {
    t.Errorf("Comparison Error: Expected NumStrDto='%s'.\n"+
      "Instead, NumStrDto= '%s'\n\n",
      expectedBigINum.bigInt.Text(10), resultBigInt.Text(10))
  }

  var resultSignValue int

  resultSignValue, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultSignValue, err = result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'",
      expectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResultNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_02(t *testing.T) {

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

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n", multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)

  nDtoArray := make([]NumStrDto, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "nDtoArray[%d], err = new(NumStrDto).NewNumStr(multiplicandStrs[%d])"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = nDtoArray[i].GetIntAry()

    if err != nil {
      t.Errorf("Error returned by:"+
        "ia, err = nDtoArray[%d].GetIntAry()\n "+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n",
        i, multiplicandStrs[i], err.Error())
      return
    }

  }

  expectedBigINum, err := new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "expectedBigINum, err := new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      expectedBigINumStr, err.Error())
    return
  }

  var multiplierNumStr, actualNumStr, iaResultNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStr, err = \n"+
      "  multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

  if err != nil {
    t.Errorf("Error returned by:"+
      "result, err := new(BigIntMathMultiply).MultiplyNumStrDtoArray(\n"+
      "    multiplierNumStrDto, nDtoArray)\n"+
      "multiplierNumStrDto='%v'\nError='%v'\n\n",
      multiplierNumStr, err.Error())
    return
  }

  var expectedBigINumEqualsResult bool

  expectedBigINumEqualsResult, err = expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:"+
      "expectedBigINumEqualsResult, err = \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  if expectedBigINum.bigInt.Cmp(result.bigInt) != 0 {
    t.Errorf("Comparison Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  var resultSignValue int

  resultSignValue, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by:"+
      "resultSignValue, err = result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:"+
      "resultSignValue, err = result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:"+
      "iaResultNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_03(t *testing.T) {

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

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "multiplierNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  nDtoArray := make([]NumStrDto, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResult, err := new(IntAry).NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n",
      multiplierStr, err.Error())
    return
  }

  var ia IntAry
  var expectedBigINum, result BigIntNum
  var multiplierNumStr, actualNumStr, iaResultNumStr string
  var expectedBigINumEqualsResult bool
  var compareResult int

  for i := 0; i < lenArray; i++ {

    nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "nDtoArray[%d], err = new(NumStrDto).NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = nDtoArray[i].GetIntAry()

    if err != nil {
      t.Errorf("Error returned by:"+
        "ia, err = nDtoArray[%d].GetIntAry()\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }
  } // End of for loop

  expectedBigINum, err = new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "expectedBigINum, err = new(BigIntNum).NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n",
      expectedBigINumStr, err.Error())
    return
  }

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:"+
      "multiplierNumStr, err = multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n", err.Error())
    return
  }

  result, err = new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

  if err != nil {
    t.Errorf("Error returned by:"+
      "result, err = new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrDtoArray(multiplierNumStrDto,\n"+
      "    nDtoArray)\n"+
      "multiplierNumStrDto='%v'\nError='%v'\n\n",
      multiplierNumStr, err.Error())
    return
  }

  expectedBigINumEqualsResult, err = expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:"+
      "expectedBigINumEqualsResult, err = expectedBigINum.Equal(result)\n"+
      "Error='%v'\n", err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("Error: Expected NumStrDto='%s'. Instead, NumStrDto= '%s'. ",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  compareResult, err = expectedBigINum.CmpBigInt(result)

  if err != nil {
    t.Errorf("Error returned by:"+
      "compareResult, err = expectedBigINum.CmpBigInt(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if compareResult != 0 {
    t.Errorf("Comparison Error: Expected NumStrDto='%s'.\n"+
      "Instead, NumStrDto= '%s'\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  if expectedBigINumSign != result.sign {
    t.Errorf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.sign)
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:"+
      "actualNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:"+
      "iaResultNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_04(t *testing.T) {

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

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  nDtoArray := make([]NumStrDto, lenArray)

  iaResult, err := new(IntAry).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResult, err := new(IntAry).\n"+
      "  NewNumStr(multiplierStr)\n"+
      "multiplierStr='%v'\nError='%v'\n\n", multiplierStr, err.Error())
    return
  }

  var ia IntAry

  for i := 0; i < lenArray; i++ {

    nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "nDtoArray[%d], err = new(NumStrDto).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, i, i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err = nDtoArray[i].GetIntAry()

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "ia, err = nDtoArray[%d].GetIntAry()\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, i, multiplicandStrs[i], err.Error())
      return
    }

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "err = iaResult.MultiplyThisBy(&ia, -1, -1)\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, multiplicandStrs[i], err.Error())
      return
    }

  } // End of for loop

  var expectedBigINum, result BigIntNum

  expectedBigINum, err = new(BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINum, err = new(BigIntNum).\n"+
      "  NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr='%v'\nError='%v'\n\n",
      expectedBigINumStr, err.Error())
    return
  }

  var multiplierNumStr, actualNumStr, iaResultNumStr string

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "multiplierNumStr, err = multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err = new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "result, err = new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrDtoArray(multiplierNumStrDto,\n"+
      "    nDtoArray)\n"+
      "multiplierNumStrDto='%v'\nError='%v'\n\n",
      multiplierNumStr, err.Error())
    return
  }

  var expectedBigINumEqualsResult bool

  expectedBigINumEqualsResult, err = expectedBigINum.Equal(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "expectedBigINumEqualsResult, err = \n"+
      "  expectedBigINum.Equal(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedBigINumEqualsResult {
    t.Errorf("Error: Expected NumStrDto='%s'.\n"+
      "Instead, NumStrDto= '%s'\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  var compareResult int

  compareResult, err = expectedBigINum.CmpBigInt(result)

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "compareResult, err =  expectedBigINum.\n"+
      "  CmpBigInt(result)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if compareResult != 0 {
    t.Errorf("Comparison Error: Expected NumStrDto='%s'.\n"+
      "Instead, NumStrDto= '%s'.\n\n",
      expectedBigINum.bigInt.Text(10), result.bigInt.Text(10))
    return
  }

  var resultSignValue int

  resultSignValue, err = result.GetSign()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "resultSignValue, err = result.GetSign()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedBigINumSign != resultSignValue {
    t.Errorf("Error: Expected number sign='%v'.\n"+
      "Instead, number sign='%v'\n\n",
      expectedBigINumSign, resultSignValue)
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "actualNumStr, err = result.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  iaResultNumStr, err = iaResult.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:\n"+
      "iaResultNumStr, err = iaResult.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if iaResultNumStr != actualNumStr {
    t.Errorf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResultNumStr, actualNumStr)
  }

  return
}

func TestBigIntMathMultiply_MultiplyNumStrDtoArray_05(t *testing.T) {

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

  // product = 11995826664,26376575446779648
  expectedNumStr := "11995826664,26376575446779648"

  multiplierNumStrDto, err := new(NumStrDto).NewNumStr(multiplierStr)

  if err != nil {
    t.Errorf("Error returned by:"+
      "multiplierNumStrDto, err := new(NumStrDto).\n"+
      "  NewNumStr(multiplierStr)\n "+
      "multiplierStr='%v'\nError='%v'\n\n",
      multiplierStr, err.Error())
    return
  }

  expectedNumSeps := NumericSeparatorDto{}
  frenchDecSeparator := ','
  frenchThousandsSeparator := ' '
  frenchCurrencySymbol := '€'

  expectedNumSeps.DecimalSeparator = frenchDecSeparator
  expectedNumSeps.ThousandsSeparator = frenchThousandsSeparator
  expectedNumSeps.CurrencySymbol = frenchCurrencySymbol

  err = multiplierNumStrDto.SetNumericSeparatorsDto(expectedNumSeps)

  if err != nil {
    t.Errorf("Error returned by:"+
      "err = multiplierNumStrDto.SetNumericSeparatorsDto(\n"+
      "  expectedNumSeps)\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  nDtoArray := make([]NumStrDto, lenArray)

  var result BigIntNum
  var multiplierNumStr, actualNumStr string
  var actualNumSeps NumericSeparatorDto

  for i := 0; i < lenArray; i++ {

    nDtoArray[i], err = new(NumStrDto).NewNumStr(multiplicandStrs[i])

    if err != nil {
      t.Errorf("Error returned by:\n"+
        "nDtoArray[%d], err = new(NumStrDto).\n"+
        "  NewNumStr(multiplicandStrs[%d])\n"+
        "multiplicandStrs[%d]='%v'\nError='%v'\n\n",
        i, i, i, multiplicandStrs[i], err.Error())
      return
    }
  } // End of for loop

  multiplierNumStr, err = multiplierNumStrDto.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:"+
      "multiplierNumStr, err = multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  result, err = new(BigIntMathMultiply).MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)

  if err != nil {
    t.Errorf("Error returned by:"+
      "result, err = new(BigIntMathMultiply).\n"+
      "  MultiplyNumStrDtoArray(multiplierNumStrDto, nDtoArray)\n"+
      "multiplierNumStrDto='%v'\nError='%v'\n\n",
      multiplierNumStr, err.Error())
    return
  }

  actualNumStr, err = result.GetNumStr()

  if err != nil {
    t.Errorf("Error returned by:"+
      "multiplierNumStr, err = multiplierNumStrDto.GetNumStr()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if expectedNumStr != actualNumStr {
    t.Errorf("Error: Expected NumStr='%v' "+
      "Instead, NumStr='%v'",
      expectedNumStr, actualNumStr)
    return
  }

  actualNumSeps, err = result.GetNumericSeparatorsDto()

  if err != nil {
    t.Errorf("Error returned by:"+
      "actualNumSeps, err = result.GetNumericSeparatorsDto()\n"+
      "Error='%v'\n\n", err.Error())
    return
  }

  if !expectedNumSeps.Equal(actualNumSeps) {
    t.Errorf("Error: Expected NumSeps='%v'\n"+
      "Instead, NumSeps='%v'\n\n",
      expectedNumSeps.String(), actualNumSeps.String())
  }

  return
}
