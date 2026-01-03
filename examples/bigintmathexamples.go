package examples

import (
  "fmt"
  "github.com/mikeaustin71/mathops"
  "math/big"
  "time"
)

// ExampleNumStrDtoBigIntParse02
// Example # 2
func ExampleNumStrDtoBigIntParse02() {

  ePrefix := "bigintmathexamples.ExampleNumStrDtoBigIntParse02()"

  num1Str := "-123456789"

  precision := uint(15)

  bNum1, isOk := big.NewInt(0).SetString(num1Str, 10)

  if !isOk {
    fmt.Printf("%v\n"+
      "Error: isOk == false\n"+
      "Error returned by:\n"+
      "bNum1, isOk := big.NewInt(0).SetString(num1Str, 10)\n"+
      "num1Str= '%v'\n\n",
      ePrefix,
      num1Str)
    return
  }

  nDto2, err := new(mathops.NumStrDto).ParseSignedBigInt(bNum1, precision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nDto2, err := new(mathops.NumStrDto).\n"+
      "  ParseSignedBigInt(bNum1, precision)\n"+
      "bNum1= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bNum1.Text(10),
      precision,
      err.Error())
    return
  }

  nDto2NumStr, err := nDto2.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nDto2NumStr, err := nDto2.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Original NumStr: ", num1Str)
  fmt.Println("   nDto2 NumStr: ", nDto2NumStr)
  fmt.Println(" spec precision: ", precision)

}

func ExampleNumStrDtoBigIntNumParse01() {

  ePrefix := "bigintmathexamples.ExampleNumStrDtoBigIntNumParse01()"

  num1Str := "0.000"

  bNum1, err := new(mathops.BigIntNum).NewNumStr(num1Str)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1, err := new(mathops.BigIntNum).NewNumStr(num1Str)\n"+
      "num1Str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Str,
      err.Error())
    return
  }

  bNum1NumStr, err := bNum1.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1NumStr, err := bNum1.GetNumStr()"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nDto2, err := new(mathops.NumStrDto).ParseBigIntNum(bNum1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nDto2, err := new(mathops.NumStrDto).\n"+
      " ParseBigIntNum(bNum1)\n"+
      "bNum1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bNum1NumStr,
      err.Error())
    return
  }

  nDto2NumStr, err := nDto2.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nDto2NumStr, err := nDto2.GetNumStr()"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Original NumStr: ", num1Str)
  fmt.Println("   nDto2 NumStr: ", nDto2NumStr)

}

// ExampleRoundPrecision01
func ExampleRoundPrecision01() {

  ePrefix := "bigintmathexamples.ExampleRoundPrecision01()"

  num1Str := "654.123456"
  expectedNumStr := "654.123"
  newPrecisionUint := uint(3)

  bNum1, err := new(mathops.BigIntNum).NewNumStr(num1Str)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1, err := new(mathops.BigIntNum).NewNumStr(num1Str)\n"+
      "num1Str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Str,
      err.Error())
    return
  }

  bNum1NumStr, err := bNum1.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1NumStr, err := bNum1.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  expectedNum, err := new(mathops.BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedNum, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedNumStr,
      err.Error())
    return
  }

  expectedNumNumStr, err := expectedNum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedNumNumStr, err := expectedNum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Old BNum1: ", bNum1NumStr)

  err = bNum1.RoundToDecPlace(newPrecisionUint)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = bNum1.RoundToDecPlace(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      newPrecisionUint,
      err.Error())
    return
  }

  bNum1NumStr, err = bNum1.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1NumStr, err = bNum1.GetNumStr()\n"+
      "Initialize #2\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("New BNum1: ", bNum1NumStr)

  bNum1PrecisionUint, err := bNum1.GetPrecisionUint()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1PrecisionUint, err := bNum1.GetPrecisionUint()\n"+
      "bNum1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bNum1NumStr,
      err.Error())
    return
  }

  expectedNumIsEqualTobNum1, err := expectedNum.Equal(bNum1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedNumIsEqualTobNum1, err := \n"+
      "  expectedNum.Equal(bNum1)\n"+
      "expectedNum= '%v'\n"+
      "bNum1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedNumNumStr,
      bNum1NumStr,
      err.Error())
    return
  }

  if !expectedNumIsEqualTobNum1 {
    fmt.Printf("%v\n"+
      "Error: expectedNum IS NOT EQUAL TO bNum1!\n"+
      "Because expectedNumIsEqualTobNum1 == false\n"+
      "Expected bNum1 = '%v'\n"+
      "  Actual bNum1 = '%v'\n\n",
      ePrefix, expectedNumNumStr, bNum1NumStr)

    return
  }

  if newPrecisionUint != bNum1PrecisionUint {
    fmt.Printf("%v\n"+
      "Error: Precision Uint Values DON'T MATCH!\n"+
      "Because newPrecisionUint != bNum1PrecisionUint\n"+
      "Expected bNum1PrecisionUint = '%v'\n"+
      "  Actual bNum1PrecisionUint = '%v'\n\n",
      ePrefix, newPrecisionUint, bNum1PrecisionUint)

    return
  }

  return
}

// ExampleSetPrecision01
func ExampleSetPrecision01() {

  ePrefix := "ExampleSetPrecision01()"

  num1Str := "654.123456"
  expectedNumStr := "654.123"
  newPrecisionUint := uint(3)

  bNum1, err := new(mathops.BigIntNum).NewNumStr(num1Str)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(num1Str)\n"+
      "num1Str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Str,
      err.Error())
    return
  }

  bNum1NumStr, err := bNum1.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1NumStr, err = bNum1.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  expectedNum, err := new(mathops.BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(expectedNumStr). Error='%v' ",
      err.Error())
    return
  }

  expectedNumNumStr, err := expectedNum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedNumNumStr, err := expectedNum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Old BNum1: ", bNum1NumStr)

  err = bNum1.SetPrecision(newPrecisionUint)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = bNum1.SetPrecision(newPrecision)\n"+
      "newPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      newPrecisionUint,
      err.Error())
    return
  }

  bNum1NumStr, err = bNum1.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1NumStr, err = bNum1.GetNumStr()\n"+
      "bNum1NumStr Initialization #2\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  bNum1PrecisionUint, err := bNum1.GetPrecisionUint()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bNum1PrecisionUint, err := bNum1.GetPrecisionUint()\n"+
      "bNum1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bNum1NumStr,
      err.Error())
    return
  }

  fmt.Println("New BNum1: ", bNum1NumStr)

  expectedNumIsEqualTobNum1, err := expectedNum.Equal(bNum1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedNumIsEqualTobNum1, err := expectedNum.Equal(bNum1)\n"+
      "expectedNum= '%v'\n"+
      "bNum1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedNumNumStr,
      bNum1NumStr,
      err.Error())
    return
  }

  if !expectedNumIsEqualTobNum1 {
    fmt.Printf("%v\n"+
      "Error: expectedNum and bNum1 ARE NOT EQUAL!\n"+
      "Eqivalency Test FAILED!!\n"+
      "Because expectedNumIsEqualTobNum1 == false\n"+
      "Expected bNum1 = '%v'\n"+
      "  Actual bNum1 = '%v'\n\n",
      ePrefix, expectedNumNumStr, bNum1NumStr)

    return
  }

  if newPrecisionUint != bNum1PrecisionUint {
    fmt.Printf("%v\n"+
      "Error: Expected and Actual Uint Precision Values DON'T MATCH!\n"+
      "Because newPrecisionUint != bNum1PrecisionUint\n"+
      "Expected bNum1PrecisionUint = '%v'\n"+
      "  Actual bNum1PrecisionUint = '%v'\n\n",
      ePrefix, newPrecisionUint, bNum1PrecisionUint)

    return
  }

  return
}

// ExampleBigIntCurrencyStr01
// Example Method
func ExampleBigIntCurrencyStr01(num1Str, expectedNumStr string, mode mathops.NegativeValueFmtMode) {

  ePrefix := "ExampleBigIntCurrencyStr01()"

  bINum, err := new(mathops.BigIntNum).NewNumStr(num1Str)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(mathops.BigIntNum).NewNumStr(num1Str)\n"+
      "num1Str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Str,
      err.Error())
    return
  }

  outStr, err := bINum.FormatCurrencyStr(mode)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "outStr, err := bINum.FormatCurrencyStr(mode)\n"+
      "mode= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      mode.String(),
      err.Error())
    return
  }

  fmt.Println("       Original NumStr: ", "'", num1Str, "'")
  fmt.Println("       Expected NumStr: ", "'", expectedNumStr, "'")
  fmt.Println("         Actual NumStr: ", "'", outStr, "'")
  fmt.Println("  Actual String Length: ", len(outStr))
  fmt.Println("Expected String Length: ", len(expectedNumStr))
  return

}

// ExampleBigIntThouStr01
func ExampleBigIntThouStr01(num1Str, expectedNumStr string, mode mathops.NegativeValueFmtMode) {

  ePrefix := "ExampleBigIntThouStr01()"

  bINum, err := new(mathops.BigIntNum).NewNumStr(num1Str)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(mathops.BigIntNum).NewNumStr(num1Str)\n"+
      "num1Str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Str,
      err.Error())
    return
  }

  bINumNumStr, err := bINum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bINumNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  outStr, err := bINum.FormatThousandsStr(mode)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "outStr, err := bINum.FormatThousandsStr(mode)\n"+
      "bINum= '%v'\n"+
      "mode= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bINumNumStr,
      mode.String(),
      err.Error())
    return
  }

  fmt.Println("       Original NumStr: ", "'", num1Str, "'")
  fmt.Println("       Expected NumStr: ", "'", expectedNumStr, "'")
  fmt.Println("         Actual NumStr: ", "'", outStr, "'")
  fmt.Println("  Actual String Length: ", len(outStr))
  fmt.Println("Expected String Length: ", len(expectedNumStr))
  return

}

func ExampleBigIntNumString03(
  bInt *big.Int,
  precision uint,
  expectedNumStr string,
  mode mathops.NegativeValueFmtMode) {

  ePrefix := "ExampleBigIntNumString03()"

  bINum, err := new(mathops.BigIntNum).NewBigInt(bInt, precision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(mathops.BigIntNum).\n"+
      "  NewBigInt(bInt, precision)\n"+
      "bInt= '%v'\n"+
      "precision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bInt.Text(10),
      precision,
      err.Error())
    return
  }

  outStr, err := bINum.FormatNumStr(mode)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "outStr, err := bINum.FormatNumStr(mode)\n"+
      "mode= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      mode.String(),
      err.Error())
    return
  }

  fmt.Println("       Expected NumStr: ", "'", expectedNumStr, "'")
  fmt.Println("         Actual NumStr: ", "'", outStr, "'")
  fmt.Println("  Actual String Length: ", len(outStr))
  fmt.Println("Expected String Length: ", len(expectedNumStr))
  return

}

// ExampleBigIntNumString02
func ExampleBigIntNumString02(num1Str, expectedNumStr string, mode mathops.NegativeValueFmtMode) {

  ePrefix := "ExampleBigIntNumString02()"

  bINum, err := new(mathops.BigIntNum).NewNumStr(num1Str)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(mathops.BigIntNum).NewNumStr(num1Str)\n"+
      "num1Str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Str,
      err.Error())
    return
  }

  outStr, err := bINum.FormatNumStr(mode)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "outStr, err := bINum.FormatNumStr(mode)\n"+
      "mode= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      mode.String(),
      err.Error())
    return
  }

  fmt.Println("       Original NumStr: ", "'", num1Str, "'")
  fmt.Println("       Expected NumStr: ", "'", expectedNumStr, "'")
  fmt.Println("         Actual NumStr: ", "'", outStr, "'")
  fmt.Println("  Actual String Length: ", len(outStr))
  fmt.Println("Expected String Length: ", len(expectedNumStr))
  return
}

// ExampleBigIntNumString01
func ExampleBigIntNumString01(num1Str string) {

  ePrefix := "ExampleBigIntNumString01()"

  bINum := new(mathops.BigIntNum)

  fmt.Println("original numStr: ", num1Str)

  expectedNumSeps := new(mathops.NumericSeparatorDto).NewUSADefaults()

  err := bINum.SetNumStr(num1Str, expectedNumSeps)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err := bINum.SetNumStr(num1Str, expectedNumSeps)\n"+
      "num1Str= '%v'\n"+
      "expectedNumSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      num1Str,
      expectedNumSeps.String(),
      err.Error())
    return
  }

  bINumNumStr, err := bINum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bINumNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("  actual numStr: ", bINumNumStr)

}

func ExampleSubtraction01() {

  ePrefix := "ExampleSubtraction01()"

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"
  minuendPrecision := uint(7)

  subtrahend0 := "737.21"
  subtrahend1 := "9637591.879546"
  subtrahend2 := "28"
  subtrahend3 := "5284.9765"
  subtrahend4 := "-189291837.12"
  subtrahend5 := "7638932.12398765"

  // result = 15,3035,620.80650965
  expectedBigINumStr := "153035620.80650965"
  expectedBigINumPrecision := uint(8)
  expectedBigINumSign := 1

  minuendNDto, err := new(mathops.NumStrDto).NewNumStr(minuendStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "minuendNDto, err := new(mathops.NumStrDto).NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      minuendStr,
      err.Error())
    return
  }

  minuendNDtoNumStr, err := minuendNDto.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "minuendNDtoNumStr, err := minuendNDto.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Minuend NumStrDto: ", minuendNDtoNumStr)

  bMinuend, err := minuendNDto.GetBigInt()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bMinuend, err := minuendNDto.GetBigInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println(" Minuend BigIntNum: ", bMinuend.Text(10))

  lenSubtrahends := 6
  subtrahendAry := make([]mathops.BigIntNum, lenSubtrahends)

  subtrahendAry[0], err = new(mathops.BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[0], err = new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend0,
      err.Error())
    return
  }

  subtrahendAry[1], err = new(mathops.BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[1], err = new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend1,
      err.Error())
    return
  }

  subtrahendAry[2], err = new(mathops.BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[2], err = new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend2,
      err.Error())
    return
  }

  subtrahendAry[3], err = new(mathops.BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[3], err = new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend3,
      err.Error())
    return
  }

  subtrahendAry[4], err = new(mathops.BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[4], err = new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend4,
      err.Error())
    return
  }

  subtrahendAry[5], err = new(mathops.BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "subtrahendAry[5], err = new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend5,
      err.Error())
    return
  }

  expectedNDto, err := new(mathops.NumStrDto).NewNumStr(expectedBigINumStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedNDto, err := new(mathops.NumStrDto).\n"+
      " NewNumStr(expectedBigINumStr)\n"+
      "expectedBigINumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumStr,
      err.Error())
    return
  }

  expectedNDtoNumStr, err := expectedNDto.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedNDtoNumStr, err := expectedNDto.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  expectedBigI, err := expectedNDto.GetBigInt()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedBigI, err := expectedNDto.GetBigInt()\n"+
      "expectedNDto= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix, expectedNDtoNumStr, err.Error())
    return
  }

  expectedBigINum, err := new(mathops.BigIntNum).NewBigInt(expectedBigI, expectedBigINumPrecision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINum, err := new(mathops.BigIntNum).\n"+
      "  NewBigInt(expectedBigI, expectedBigINumPrecision)\n"+
      "expectedBigI= '%v'\n"+
      "expectedBigINumPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigI.Text(10),
      expectedBigINumPrecision,
      err.Error())
    return
  }

  expectedBigINumNumStr, err := expectedBigINum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  minuendBiNum, err := new(mathops.BigIntNum).NewBigInt(bMinuend, minuendPrecision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNum, err := new(mathops.BigIntNum).\n"+
      "  NewBigInt(bMinuend, minuendPrecision)\n"+
      "bMinuend= '%v'\n"+
      "minuendPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bMinuend.Text(10),
      minuendPrecision,
      err.Error())
    return
  }

  minuendBiNumNumStr, err := minuendBiNum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumNumStr, err := minuendBiNum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  result, err := new(mathops.BigIntMathSubtract).SubtractBigIntNumArray(minuendBiNum, subtrahendAry)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(mathops.BigIntMathSubtract).\n"+
      " SubtractBigIntNumArray(minuendBiNum, subtrahendAry)\n"+
      "minuendBiNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      minuendBiNumNumStr,
      err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultSignValue, err := result.GetSign()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultSignValue, err := result.GetSign()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  minuendBiNumNumStr, err = minuendBiNum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "minuendBiNumNumStr, err = minuendBiNum.GetNumStr()\n"+
      "Initialization #2\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Minuend: ", minuendBiNumNumStr)

  var subtrahendAryElementStr string

  for i := 0; i < lenSubtrahends; i++ {

    subtrahendAryElementStr, err = subtrahendAry[i].GetNumStr()

    if err != nil {
      fmt.Printf("%v\n"+
        "Error returned by:\n"+
        "subtrahendAryElementStr, err = subtrahendAry[i].GetNumStr()\n"+
        "i= '%v'\n"+
        "Error='%v'\n\n",
        ePrefix,
        i,
        err.Error())
      return
    }

    fmt.Printf("Subtrahend[%v]='%v' \n", i, subtrahendAryElementStr)
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err = result.GetNumStr()\n"+
      "Initialization #2\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Result: ", resultNumStr)

  expectedBigINumEqualsResult, err := expectedBigINum.CmpBigInt(result)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedBigINumEqualsResult, err :=\n"+
      "  expectedBigINum.CmpBigInt(result)\n"+
      "expectedBigINum= '%v'\n"+
      "result= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedBigINumNumStr,
      resultNumStr,
      err.Error())
    return
  }

  if expectedBigINumEqualsResult != 0 {
    fmt.Printf("%v\n"+
      "Error: Expected and Actual Big Int Numbers ARE NOT EQUAL!\n"+
      "Because expectedBigINumEqualsResult != 0\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedBigINumNumStr, resultNumStr)

    return
  }

  if expectedBigINumSign != resultSignValue {
    fmt.Printf("%v\n"+
      "Error: Expected and Actual Sign Values ARE NOT EQUAL!\n"+
      "Because expectedBigINumSign != resultSignValue\n"+
      "Expected resultSignValue = '%v'\n"+
      "  Actual resultSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultSignValue)

    return
  }

  return
}

func ExampleSubtraction02() {

  ePrefix := "ExampleSubtraction02()"

  minuendStr := "-1718973642.1234567"

  iaMinuend, err := new(mathops.IntAry).NewNumStr(minuendStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaMinuend, err := new(mathops.IntAry).\n"+
      " NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      minuendStr,
      err.Error())
    return
  }

  iaMinuendNumStr, err := iaMinuend.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err := iaMinuend.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  bigIMinuend, err := new(mathops.BigIntNum).NewNumStr(minuendStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigIMinuend, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(minuendStr)\n"+
      "minuendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      minuendStr,
      err.Error())
    return
  }

  bigIMinuendNumStr, err := bigIMinuend.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigIMinuendNumStr, err := bigIMinuend.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  subtrahend0 := "-28934682.721"
  subtrahend1 := "424.987654321"
  subtrahend2 := "-987"
  subtrahend3 := "62.94"
  subtrahend4 := "-999999999.99999"
  subtrahend5 := "-9638932.371"

  subtrahendAry := make([]mathops.BigIntNum, 6)
  // Confirmed Subtraction Result
  // ia Result5:    2,757,547,756.287792379
  // Array Result:  2,757,547,756.287792379

  iaSub0, err := new(mathops.IntAry).NewNumStr(subtrahend0)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub0, err := new(mathops.IntAry).\n"+
      " NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend0,
      err.Error())
    return
  }

  iaSub1, err := new(mathops.IntAry).NewNumStr(subtrahend1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub1, err := new(mathops.IntAry).\n"+
      " NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend1,
      err.Error())
    return
  }

  iaSub1NumStr, err := iaSub1.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub1NumStr, err := iaSub1.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  iaSub2, err := new(mathops.IntAry).NewNumStr(subtrahend2)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub2, err := new(mathops.IntAry).\n"+
      " NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend2,
      err.Error())
    return
  }

  iaSub3, err := new(mathops.IntAry).NewNumStr(subtrahend3)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub3, err := new(mathops.IntAry).\n"+
      " NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend3,
      err.Error())
    return
  }

  iaSub3NumStr, err := iaSub3.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub3NumStr, err := iaSub3.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix,
      err.Error())
    return
  }

  iaSub4, err := new(mathops.IntAry).NewNumStr(subtrahend4)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub4, err := new(mathops.IntAry).\n"+
      " NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend4,
      err.Error())
    return
  }

  iaSub4NumStr, err := iaSub4.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub4NumStr, err := iaSub4.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix,
      err.Error())
    return
  }

  iaSub5, err := new(mathops.IntAry).NewNumStr(subtrahend5)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub5, err := new(mathops.IntAry).\n"+
      " NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend5,
      err.Error())
    return
  }

  iaSub5NumStr, err := iaSub5.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaSub5NumStr, err := iaSub5.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix,
      err.Error())
    return
  }

  bigISub0, err := new(mathops.BigIntNum).NewNumStr(subtrahend0)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub0, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend0)\n"+
      "subtrahend0= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend0,
      err.Error())
    return
  }

  bigISub0NumStr, err := bigISub0.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub0NumStr, err := bigISub0.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  subtrahendAry[0] = bigISub0

  bigISub1, err := new(mathops.BigIntNum).NewNumStr(subtrahend1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub1, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend1)\n"+
      "subtrahend1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend1,
      err.Error())
    return
  }

  bigISub1NumStr, err := bigISub1.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub1NumStr, err := bigISub1.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  subtrahendAry[1] = bigISub1

  bigISub2, err := new(mathops.BigIntNum).NewNumStr(subtrahend2)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub2, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend2)\n"+
      "subtrahend2= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend2,
      err.Error())
    return
  }

  bigISub2NumStr, err := bigISub2.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub2NumStr, err := bigISub2.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  subtrahendAry[2] = bigISub2

  bigISub3, err := new(mathops.BigIntNum).NewNumStr(subtrahend3)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub3, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend3)\n"+
      "subtrahend3= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend3,
      err.Error())
    return
  }

  bigISub3NumStr, err := bigISub3.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  subtrahendAry[3] = bigISub3

  bigISub4, err := new(mathops.BigIntNum).NewNumStr(subtrahend4)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub4, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend4)\n"+
      "subtrahend4= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend4,
      err.Error())
    return
  }

  bigISub4NumStr, err := bigISub4.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub4NumStr, err := bigISub4.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  subtrahendAry[4] = bigISub4

  bigISub5, err := new(mathops.BigIntNum).NewNumStr(subtrahend5)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub5, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(subtrahend5)\n"+
      "subtrahend5= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      subtrahend5,
      err.Error())
    return
  }

  bigISub5NumStr, err := bigISub5.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigISub5NumStr, err := bigISub5.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix,
      err.Error())
    return
  }

  subtrahendAry[5] = bigISub5

  bPair, err := new(mathops.BigIntPair).NewBigIntNum(bigIMinuend, bigISub0)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bPair, err := new(mathops.BigIntPair).\n"+
      "  NewBigIntNum(bigIMinuend, bigISub0)\n"+
      "bigIMinuend= '%v'\n"+
      "bigISub0= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bigIMinuendNumStr,
      bigISub0NumStr,
      err.Error())
    return
  }

  result, err := new(mathops.BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(mathops.BigIntMathSubtract).SubtractPair(bPair)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("Original Minuend: ", iaMinuendNumStr)
  fmt.Println("        bigISub0: ", bigISub0NumStr)
  fmt.Println("         Result0: ", resultNumStr)
  _ = iaMinuend.SubtractFromThis(&iaSub0)
  fmt.Println("      ia Result0: ", iaMinuendNumStr)

  err = iaMinuend.SubtractFromThis(&iaSub1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.SubtractFromThis(&iaSub1)\n"+
      "iaMinuend= '%v'\n"+
      "iaSub1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      iaSub1NumStr,
      err.Error())
    return
  }

  bPair, err = new(mathops.BigIntPair).NewBigIntNum(result, bigISub1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bPair, err = new(mathops.BigIntPair).\n"+
      "  NewBigIntNum(result, bigISub1)\n"+
      "result= '%v'\n"+
      "bigISub1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      resultNumStr,
      bigISub1NumStr,
      err.Error())
    return
  }

  result, err = new(mathops.BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err = new(mathops.BigIntMathSubtract).\n"+
      "  SubtractPair(bPair)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Initialization #2\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("        bigISub1: ", bigISub1NumStr)
  fmt.Println("         Result1: ", resultNumStr)
  fmt.Println("      ia Result1: ", iaMinuendNumStr)

  err = iaMinuend.SubtractFromThis(&iaSub2)

  bPair, err = new(mathops.BigIntPair).NewBigIntNum(result, bigISub2)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bPair, err = new(mathops.BigIntPair).\n"+
      " NewBigIntNum(result, bigISub2)\n"+
      "target= '%v'\n"+
      "result= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      resultNumStr,
      bigISub2NumStr,
      err.Error())
    return
  }

  result, err = new(mathops.BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err = new(mathops.BigIntMathSubtract).SubtractPair(bPair)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Initialization #3\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("       bigISub2: ", bigISub2NumStr)
  fmt.Println("        Result2: ", resultNumStr)
  fmt.Println("     ia Result2: ", iaMinuendNumStr)

  err = iaMinuend.SubtractFromThis(&iaSub3)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.SubtractFromThis(&iaSub3)\n"+
      "iaMinuend= '%v'\n"+
      "xrayStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaMinuendNumStr,
      iaSub3NumStr,
      err.Error())
    return
  }

  iaMinuendNumStr, err = iaMinuend.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaMinuendNumStr, err = iaMinuend.GetNumStr()\n"+
      "Initialization # 2-b\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  bPair, err = new(mathops.BigIntPair).NewBigIntNum(result, bigISub3)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bPair, err = new(mathops.BigIntPair).\n"+
      " NewBigIntNum(result, bigISub3)\n"+
      "result= '%v'\n"+
      "bigISub3= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      resultNumStr,
      bigISub3NumStr,
      err.Error())
    return
  }

  result, err = new(mathops.BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err = new(mathops.BigIntMathSubtract).SubtractPair(bPair)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err = result.GetNumStr()\n"+
      "Initialization # 3-b\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("       bigISub3: ", bigISub3NumStr)
  fmt.Println("        Result3: ", resultNumStr)
  fmt.Println("     ia Result3: ", iaMinuendNumStr)

  err = iaMinuend.SubtractFromThis(&iaSub4)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.SubtractFromThis(&iaSub4)\n"+
      "iaSub4= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaSub4NumStr,
      err.Error())
    return
  }

  bPair, err = new(mathops.BigIntPair).NewBigIntNum(result, bigISub4)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bPair, err = new(mathops.BigIntPair).\n"+
      " NewBigIntNum(result, bigISub4)\n"+
      "result= '%v'\n"+
      "bigISub4= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      resultNumStr,
      bigISub4NumStr,
      err.Error())
    return
  }

  result, err = new(mathops.BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err = new(mathops.BigIntMathSubtract).\n"+
      " SubtractPair(bPair)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err = result.GetNumStr()\n"+
      "Initialization # 5-a\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("       bigISub4: ", bigISub4NumStr)
  fmt.Println("        Result4: ", resultNumStr)
  fmt.Println("     ia Result4: ", iaMinuendNumStr)

  err = iaMinuend.SubtractFromThis(&iaSub5)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = iaMinuend.SubtractFromThis(&iaSub5)\n"+
      "iaSub5= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaSub5NumStr,
      err.Error())
    return
  }

  bPair, err = new(mathops.BigIntPair).NewBigIntNum(result, bigISub5)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bPair, err = new(mathops.BigIntPair).\n"+
      " NewBigIntNum(result, bigISub5)\n"+
      "result= '%v'\n"+
      "bigISub5= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      resultNumStr,
      bigISub5NumStr,
      err.Error())
    return
  }

  result, err = new(mathops.BigIntMathSubtract).SubtractPair(bPair)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err = new(mathops.BigIntMathSubtract).\n"+
      "  SubtractPair(bPair)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Initialization # 6-a\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("       bigISub5: ", bigISub5NumStr)
  fmt.Println("        Result5: ", resultNumStr)
  fmt.Println("     ia Result5: ", iaMinuendNumStr)

  result, err = new(mathops.BigIntMathSubtract).SubtractBigIntNumArray(bigIMinuend, subtrahendAry)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err = new(mathops.BigIntMathSubtract).\n"+
      " SubtractBigIntNumArray(bigIMinuend, subtrahendAry)\n"+
      "bigIMinuend= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bigIMinuendNumStr,
      err.Error())
    return
  }

  resultNumStr, err = result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err = result.GetNumStr()\n"+
      "Initialization # 7 \n"+
      "Error='%v'\n\n",
      ePrefix,
      err.Error())
    return
  }

  fmt.Println("   Array Result: ", resultNumStr)

}

func ExampleSubtractIntAryArray01() {

  ePrefix := "ExampleSubtractIntAryArray01()"

  var err error

  // minuend = -18,973,642.1234567
  minuendStr := "-18973642.1234567"

  subtrahend0 := "737.21"
  subtrahend1 := "9637591.879546"
  subtrahend2 := "28"
  subtrahend3 := "5284.9765"
  subtrahend4 := "-189291837.12"
  subtrahend5 := "7638932.12398765"

  // result = 153,035,620.80650965
  expectedBigINumStr := "153035620.80650965"

  expectedBigINumSign := 1

  iaMinuend, err := new(mathops.IntAry).NewNumStr(minuendStr)

  if err != nil {
    fmt.Printf("Error returned by IntAry{}.NewNumStr(minuendStr) "+
      "minuendStr='%v' Error='%v' ", minuendStr, err.Error())
    return
  }

  expectedBigINum, err := new(mathops.BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    fmt.Printf("Error returned by BigIntNum{}.NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v' Error='%v' ", expectedBigINumStr, err.Error())
    return
  }

  lenSubtrahends := 6
  subtrahendAry := make([]mathops.IntAry, lenSubtrahends)

  subtrahendAry[0], err = new(mathops.IntAry).NewNumStr(subtrahend0)

  if err != nil {
    fmt.Printf("Error returned from IntAry{}.NewNumStr(subtrahend0). "+
      "subtrahend0='%v' Error='%v'. ",
      subtrahend0, err.Error())
    return
  }

  subtrahendAry[1], err = new(mathops.IntAry).NewNumStr(subtrahend1)

  if err != nil {
    fmt.Printf("Error returned from IntAry{}.NewNumStr(subtrahend1). "+
      "subtrahend1='%v' Error='%v'. ",
      subtrahend1, err.Error())
    return
  }

  subtrahendAry[2], err = new(mathops.IntAry).NewNumStr(subtrahend2)

  if err != nil {
    fmt.Printf("Error returned from IntAry{}.NewNumStr(subtrahend2). "+
      "subtrahend2='%v' Error='%v'. ",
      subtrahend2, err.Error())
    return
  }

  subtrahendAry[3], err = new(mathops.IntAry).NewNumStr(subtrahend3)

  if err != nil {
    fmt.Printf("Error returned from IntAry{}.NewNumStr(subtrahend3). "+
      "subtrahend3='%v' Error='%v'. ",
      subtrahend3, err.Error())
    return
  }

  subtrahendAry[4], err = new(mathops.IntAry).NewNumStr(subtrahend4)

  if err != nil {
    fmt.Printf("Error returned from IntAry{}.NewNumStr(subtrahend4). "+
      "subtrahend4='%v' Error='%v'. ",
      subtrahend4, err.Error())
    return
  }

  subtrahendAry[5], err = new(mathops.IntAry).NewNumStr(subtrahend5)

  if err != nil {
    fmt.Printf("Error returned from IntAry{}.NewNumStr(subtrahend5). "+
      "subtrahend5='%v' Error='%v'. ",
      subtrahend5, err.Error())
    return
  }

  result, err := new(mathops.BigIntMathSubtract).SubtractIntAryArray(iaMinuend, subtrahendAry)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(mathops.BigIntMathSubtract).\n"+
      " SubtractIntAryArray(iaMinuend, subtrahendAry)\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultNumStr, err := result.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultNumStr, err := result.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  expectedBigInt, err := expectedBigINum.GetBigInt()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedBigInt, err := expectedBigINum.GetBigInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  resultBigInt, err := result.GetBigInt()

  if err != nil {
    fmt.Printf("Error returned by result.GetBigInt(). Error='%v' ", err.Error())
    return
  }

  resultNumberSignValue, err := result.GetSign()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "resultNumberSignValue, err := result.GetSign()\n"+
      "result= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      resultNumStr,
      err.Error())
    return
  }

  if expectedBigInt.Cmp(resultBigInt) != 0 {
    fmt.Printf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigInt.Text(10), resultBigInt.Text(10))
    return
  }

  if expectedBigINumSign != resultNumberSignValue {
    fmt.Printf("%v\n"+
      "Error: Unexpected Result!\n"+
      "Because expectedBigINumSign != resultNumberSignValue\n"+
      "Expected resultNumberSignValue = '%v'\n"+
      "  Actual resultNumberSignValue = '%v'\n\n",
      ePrefix, expectedBigINumSign, resultNumberSignValue)

    return
  }

  fmt.Println("BigIntMathSubtract{}.SubtractIntAryArray")
  fmt.Println("========================================")
  fmt.Println("  Actual Result: ", resultNumStr)
  fmt.Println("Expected Result: ", expectedBigINumStr)
}

func ExampleBigIntRounding_01() {
  nStr := "0.000"
  expectedNumStr := "0.00"
  roundToDec := uint(2)

  bINum1, _ := new(mathops.BigIntNum).NewNumStr(nStr)

  bINum1.RoundToDecPlace(roundToDec)

  actualNumStr := bINum1.GetNumStr()

  fmt.Println("Expected NumStr: ", expectedNumStr)
  fmt.Println("  Actual NumStr: ", actualNumStr)

}

func ExampleBigIntAdd_01() {
  //n1Str := "0.000123"

  //n2Str := ".001"

  b1 := big.NewInt(123)
  b2 := big.NewInt(1)

  expectedResultStr := "1123"
  expectedPrecision := uint(6)

  result := mathops.BigIntMathAdd{}.AddBigInts(b1, 6, b2, 3)

  ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)
}

func ExampleBigIntAddNumStr_01() {
  n1Str := "0.000123"

  n2Str := "1"

  // Result = 	1.000123
  expectedResultStr := "1000123"
  expectedPrecision := uint(6)

  dto := mathops.NumericSeparatorDto{}

  numSeps := dto.New()

  result, err := mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str, numSeps)

  if err != nil {
    fmt.Printf("Error returned by mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
      "Error='%v' ", err.Error())
    return
  }

  ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

  fmt.Println("Successful Completion")
}

func ExampleBigIntAddNumStr_02() {
  n1Str := "0.000123"

  n2Str := ".001"

  // Result = 	"0.001123"
  expectedResultStr := "001123"
  expectedPrecision := uint(6)

  dto := &mathops.NumericSeparatorDto{}
  numSeps := dto.New()

  result, err := mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str, numSeps)

  if err != nil {
    fmt.Printf("Error returned by mathops.BigIntMathAdd{}.AddNumStr(n1Str, n2Str). "+
      "Error='%v' ", err.Error())
    return
  }

  ExamplePrintBasicMathResult(expectedResultStr, expectedPrecision, result)

}

func ExamplePrintBasicMathResult(expectedResultStr string,
  expectedPrecision uint,
  result mathops.BigIntNum) {

  getNumStr := result.GetNumStr()

  fmt.Println("            expected result: ", expectedResultStr)
  fmt.Println("              result.bigInt: ", result.GetNumStr())
  fmt.Println("         result.GetNumStr(): ", getNumStr)
  fmt.Println("")
  fmt.Println("         Expected precision: ", expectedPrecision)
  fmt.Println("           result.precision: ", result.GetPrecisionUint())
  fmt.Println("")

}

func ExampleBigIntDivideModulo_03(
  numStrDividend,
  numStrDivisor string,
  maxPrecision uint,
  expectedResult string) {

  bINDividend, err := new(mathops.BigIntNum).NewNumStr(numStrDividend)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(numStrDividend) "+
      " numStrDividend='%v' Error='%v. ",
      numStrDividend, err.Error())
    return
  }

  bINDivisor, err := new(mathops.BigIntNum).NewNumStr(numStrDivisor)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(numStrDivisor) "+
      " numStrDivisor='%v' Error='%v. ",
      numStrDividend, err.Error())
    return
  }

  modulo, err := mathops.BigIntMathDivide{}.BigIntNumModulo(
    bINDividend,
    bINDivisor,
    maxPrecision)

  if err != nil {
    fmt.Printf("Error returned bymathops.BigIntMathDivide{}.BigIntNumModulo(...) "+
      "Error='%v. ",
      err.Error())
    return
  }

  PrintBigIntNumModulo(
    "Raw Results",
    bINDividend,
    bINDivisor,
    modulo,
    maxPrecision,
    expectedResult)

  modulo.TrimTrailingFracZeros()

  PrintBigIntNumModulo(
    "Optimized Results",
    bINDividend,
    bINDivisor,
    modulo,
    maxPrecision,
    expectedResult)

}

func ExampleBigIntDivideIntQuotient_02(numStrDividend, numStrDivisor, expectedResult string) {

  bINDividend, err := new(mathops.BigIntNum).NewNumStr(numStrDividend)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(numStrDividend) "+
      " numStrDividend='%v' Error='%v. ",
      numStrDividend, err.Error())
    return
  }

  bINDivisor, err := new(mathops.BigIntNum).NewNumStr(numStrDivisor)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(numStrDivisor) "+
      " numStrDivisor='%v' Error='%v. ",
      numStrDividend, err.Error())
    return
  }

  quotient, err := mathops.BigIntMathDivide{}.BigIntNumIntQuotient(
    bINDividend,
    bINDivisor)

  if err != nil {
    fmt.Printf("Error returned bymathops.BigIntMathDivide{}.BigIntNumIntQuotient(...) "+
      "Error='%v. ",
      err.Error())
    return
  }

  PrintBigIntNumIntQuotient(
    "Raw Results",
    bINDividend,
    bINDivisor,
    quotient,
    expectedResult)

  quotient.TrimTrailingFracZeros()

  PrintBigIntNumIntQuotient(
    "Optimized Results",
    bINDividend,
    bINDivisor,
    quotient,
    expectedResult)

}

func ExampleBigIntDivideQuotientModulo_01(numStrDividend, numStrDivisor string, maxPrecision uint) {

  bINDividend, err := new(mathops.BigIntNum).NewNumStr(numStrDividend)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(numStrDividend) "+
      " numStrDividend='%v' Error='%v. ",
      numStrDividend, err.Error())
    return
  }

  bINDivisor, err := new(mathops.BigIntNum).NewNumStr(numStrDivisor)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(numStrDivisor) "+
      " numStrDivisor='%v' Error='%v. ",
      numStrDividend, err.Error())
    return
  }

  quotient, modulo, err := mathops.BigIntMathDivide{}.BigIntNumQuotientMod(
    bINDividend,
    bINDivisor,
    maxPrecision)

  if err != nil {
    fmt.Printf("Error returned bymathops.BigIntMathDivide{}.BigIntNumQuotientMod(...) "+
      "Error='%v. ",
      err.Error())
    return
  }

  PrintBigIntNumQuotientMod(
    "Raw Results",
    bINDividend,
    bINDivisor,
    quotient,
    modulo,
    maxPrecision)

  quotient.TrimTrailingFracZeros()
  modulo.TrimTrailingFracZeros()

  PrintBigIntNumQuotientMod(
    "Optimized Results",
    bINDividend,
    bINDivisor,
    quotient,
    modulo,
    maxPrecision)

}

func ExampleBigIntMultiply_02() {

  var err error

  // multiplier = 2
  multiplierStr := "2"
  // multiplicandStrs
  multiplicandStrs := []string{
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
  }

  // product = 128
  expectedBigINumStr := "128"

  expectedBigINumSign := 1

  multiplierBiNum, err := new(mathops.BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
    return
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]mathops.BigIntNum, lenArray)

  iaResult, err := new(mathops.IntAry).NewNumStr(multiplierStr)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.IntAry).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
    return
  }

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(mathops.BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(multiplicandStrs[i]) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
      return
    }

    ia, err := bINumArray[i].GetIntAry()

    if err != nil {
      fmt.Printf("Error returned by bINumArray[i].GetIntAryElements() "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
      return
    }

    fmt.Println("After ia bINumArray[i]=", bINumArray[i].GetNumStr())

    err = iaResult.MultiplyThisBy(&ia, -1, -1)

    if err != nil {
      fmt.Printf("Error returned by iaResult.MultiplyThisBy(&ia, -1, -1) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
      return
    }

  }

  expectedBigINum, err := new(mathops.BigIntNum).NewNumStr(expectedBigINumStr)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(expectedBigINumStr) "+
      "expectedBigINumStr='%v'  Error='%v'. ", expectedBigINumStr, err.Error())
    return
  }

  for k := 0; k < len(bINumArray); k++ {
    fmt.Println("PreLoad bINumArray[k]=", bINumArray[k].GetNumStr())
  }

  result := new(mathops.BigIntMathMultiply).MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

  if !expectedBigINum.Equal(result) {
    fmt.Printf("Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.GetNumStr(), result.GetNumStr())
    return
  }

  if expectedBigINum.CmpBigInt(result) != 0 {
    fmt.Printf("Comparison Error: Expected BigIntNum='%s'. Instead, BigIntNum= '%s'. ",
      expectedBigINum.GetNumStr(), result.GetNumStr())
    return
  }

  if expectedBigINumSign != result.GetSign() {
    fmt.Printf("Error: Expected number sign='%v'. Instead, number sign='%v'",
      expectedBigINumSign, result.GetSign())
    return
  }

  actualNumStr := result.GetNumStr()

  if iaResult.GetNumStr() != actualNumStr {
    fmt.Printf("Error: Expected actualNumStr='%v' "+
      "Instead, actualNumStr='%v'",
      iaResult.GetNumStr(), actualNumStr)
    return
  }

}

func ExampleBigIntMultiply_01() {
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

  // expectedBigINumSign := 1

  multiplierBiNum, err := new(mathops.BigIntNum).NewNumStr(multiplierStr)

  if err != nil {
    fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(multiplierStr) "+
      "multiplierStr='%v'  Error='%v'. ", multiplierStr, err.Error())
  }

  lenArray := len(multiplicandStrs)
  bINumArray := make([]mathops.BigIntNum, lenArray)

  for i := 0; i < lenArray; i++ {

    bINumArray[i], err = new(mathops.BigIntNum).NewNumStr(multiplicandStrs[i])

    if err != nil {
      fmt.Printf("Error returned by new(mathops.BigIntNum).NewNumStr(multiplicandStrs[i]) "+
        "i='%v'  multiplicandStrs[i]='%v'  Error='%v'. ", i, multiplicandStrs[i], err.Error())
    }

  }

  result := new(mathops.BigIntMathMultiply).MultiplyBigIntNumArray(multiplierBiNum, bINumArray)

  fmt.Println("Expected Result: ", expectedBigINumStr)
  fmt.Println("  Actual Result: ", result.GetNumStr())

}

func ExampleBigIntNumPower_01(baseStr, exponentStr, expectedStr string, maxPrecision uint) {

  bINumBase, err := new(mathops.BigIntNum).NewNumStr(baseStr)

  if err != nil {
    fmt.Printf("Error returned by BigIntNum{}.NewNumStr(baseStr). "+
      "baseStr='%v' Error='%v' \n", baseStr, err.Error())
    return
  }

  fmt.Println("Base= ", bINumBase.GetNumStr())

  bINumExponent, err := new(mathops.BigIntNum).NewNumStr(exponentStr)

  if err != nil {
    fmt.Printf("Error returned by BigIntNum{}.NewNumStr(exponentStr). "+
      "exponentStr='%v' Error='%v' \n", exponentStr, err.Error())
    return
  }

  fmt.Println("Exponent= ", bINumExponent.GetNumStr())

  var t0 time.Time
  var t1 time.Time

  t0 = time.Now()
  result, err := mathops.BigIntMathPower{}.BigIntNumPwr(bINumBase, bINumExponent, maxPrecision)
  t1 = time.Now()

  if err != nil {
    fmt.Printf("Error returned by BigIntMathPower{}.BigIntNumPwr(bINumBase, bINumExponent, maxPrecision). "+
      "bINumBase='%v' bINumExponent='%v' maxPrecision='%v' Error='%v' \n",
      bINumBase.GetNumStr(), bINumExponent.GetNumStr(), maxPrecision, err.Error())
    return
  }

  str := CodeDurationToStr(t1.Sub(t0))

  fmt.Println()
  fmt.Println("*** BigIntMathPower{}.BigIntNumPwr() ***")
  fmt.Println("===============================")
  if expectedStr != result.GetNumStr() {
    fmt.Println("XXX FAILURE XXX")
  } else {
    fmt.Println("*** SUCCESS ***")
  }
  fmt.Println("===============================")

  fmt.Println("Expected Result: ", expectedStr)
  fmt.Println("  Actual Result: ", result.GetNumStr())
  fmt.Println("   Elapsed Time: ", str)

  bigIntAbs := bINumExponent.GetAbsoluteBigIntValue()

  fmt.Println("              Base: ", bINumBase.GetNumStr())
  fmt.Println("          Exponent: ", bINumExponent.GetNumStr())
  fmt.Println("     Exponent Sign: ", bINumExponent.GetSign())
  fmt.Println("Exponent AbsBigInt: ", bigIntAbs.Text(10))
  fmt.Println("     Max Precision: ", maxPrecision)

}

func ExampleDecimalDivide_01() {
  // str1 / str2
  /*
  	str1 := "575.63"
  	str2 := "2014.123"
  	ePrecision := 20
  	expected := "0.28579684557497233287"
  */

  str1 := "975.69"
  str2 := "589.7654321"
  expected := "1.654369597"
  ePrecision := 9

  d1, err := mathops.Decimal{}.NewNumStr(str1)

  if err != nil {
    fmt.Printf("Error returned from mathops.Decimal{}.NewNumStr(str1). "+
      "str1='%v' Error='%v \n", str1, err.Error())
    return
  }

  d2, err := mathops.Decimal{}.NewNumStr(str2)

  if err != nil {
    fmt.Printf("Error returned from mathops.Decimal{}.NewNumStr(str2). "+
      "str2='%v' Error='%v \n", str2, err.Error())
    return
  }

  d3, err := d1.Divide(d2, uint(ePrecision))

  if err != nil {
    fmt.Printf("Error returned from d1.Divide(d2, ePrecision). "+
      " Error='%v \n", err.Error())
    return
  }

  actualResult := d3.GetNumStr()

  ia1, err := new(mathops.IntAry).NewNumStr(str1)

  if err != nil {
    fmt.Printf("Error returned from new(mathops.IntAry).NewNumStr(str1). "+
      "str1='%v' Error='%v \n", str1, err.Error())
    return
  }

  ia2, err := new(mathops.IntAry).NewNumStr(str2)

  if err != nil {
    fmt.Printf("Error returned from new(mathops.IntAry).NewNumStr(str2). "+
      "str2='%v' Error='%v \n", str2, err.Error())
    return
  }

  ia3, err := ia1.DivideThisBy(&ia2, 0, ePrecision)

  if err != nil {
    fmt.Printf("Error returned from ia1.DivideThisBy(&ia2, 0, 20 ). "+
      "Error='%v \n", err.Error())
    return
  }

  chkResult := ia3.GetNumStr()

  fmt.Println("			    Dividend: ", str1)
  fmt.Println(" 			   Divisor: ", str2)
  fmt.Println("	   Actual Result: ", actualResult)
  fmt.Println("	 Expected Result: ", expected)
  fmt.Println("	    Check Result: ", chkResult)

}

func PrintBigIntNumModulo(
  title string,
  dividend,
  divisor,
  modulo mathops.BigIntNum,
  maxPrecision uint,
  expectedResult string) {

  fmt.Println(title)
  fmt.Println("**************************************************")
  fmt.Println("Results of BigIntMathDivide.BigIntNumModulo() ")
  fmt.Println("**************************************************")
  fmt.Println("         Dividend: ", dividend.GetNumStr())
  fmt.Println("          Divisor: ", divisor.GetNumStr())
  fmt.Println("           Modulo: ", modulo.GetNumStr())
  fmt.Println("  Expected Result: ", expectedResult)
  fmt.Println("    Max Precision: ", maxPrecision)
}

func PrintBigIntNumIntQuotient(
  title string,
  dividend,
  divisor,
  quotient mathops.BigIntNum,
  expectedResult string) {

  fmt.Println(title)
  fmt.Println("**************************************************")
  fmt.Println("Results of BigIntMathDivide.BigIntNumIntQuotient() ")
  fmt.Println("**************************************************")
  fmt.Println("         Dividend: ", dividend.GetNumStr())
  fmt.Println("          Divisor: ", divisor.GetNumStr())
  fmt.Println("         Quotient: ", quotient.GetNumStr())
  fmt.Println("  Expected Result: ", expectedResult)

}

func PrintBigIntNumQuotientMod(
  title string,
  dividend,
  divisor,
  quotient,
  modulo mathops.BigIntNum,
  maxPrecision uint) {

  fmt.Println(title)
  fmt.Println("**************************************************")
  fmt.Println("Results of BigIntMathDivide.BigIntNumQuotientMod() ")
  fmt.Println("**************************************************")
  fmt.Println("         Dividend: ", dividend.GetNumStr())
  fmt.Println("          Divisor: ", divisor.GetNumStr())
  fmt.Println("         Quotient: ", quotient.GetNumStr())
  fmt.Println("           Modulo: ", modulo.GetNumStr())
  fmt.Println("    Max Precision: ", maxPrecision)

}
