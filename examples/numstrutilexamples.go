package examples

import (
  "fmt"

  "github.com/mikeaustin71/mathops"
)

func TestNumberDelimiter() {
  ns := mathops.NumStrUtility{}
  n := "1234567890"

  result := ns.DlimDecCurrStr(n, ',', '.', '$')
  fmt.Println(" Original Number: ", n)
  fmt.Println("Delimited Number: ", result)
  //expected result == "1,234,567,890"
}

func TestCurrencyDelimiter() {
  ns := mathops.NumStrUtility{}
  n := "$1234567890.25"

  result := ns.DlimDecCurrStr(n, ',', '.', '$')
  fmt.Println(" Original Number: ", n)
  fmt.Println("Delimited Number: ", result)
  //expected result == "$1,234,567,890.25"

}

func TestDNumStr() {
  ns := mathops.NumStrUtility{}
  n := int64(1234567890)
  expected := "1,234,567,890"

  result := ns.DLimI64(n, ',')

  fmt.Println("Original int64:", n)
  fmt.Println("Expected Result from ns.DLimI64:", expected)
  fmt.Println("Result from ns.DLimI64:", result)
}

func TestDNumStrEvenThousands() {
  ns := mathops.NumStrUtility{}
  n := int64(123456)
  expected := "123,456"

  result := ns.DLimI64(n, ',')

  fmt.Println("Original int64:", n)
  fmt.Println("Expected Result from ns.DLimI64:", expected)
  fmt.Println("Result from ns.DLimI64:", result)
}

func TestNumStrUtilityParseNumString() {
  strs := []string{"123456.654321",
    "123456",
    "0.123456",
    ".123456",
    "1 2   3 4",
    "1 2   3 4 . 1 2 3 4 5 6",
    "-32.495",
    "-.4219",
    "-0.713",
    "Exit Status  1",
    "0",
    "5",
    "24.95",
    "+24.95",
    "Nothing"}

  for _, s := range strs {

    TestParseAndPrintOutNumStrs(s)

  }

}

/* Output
******************************************************
    Original Input Str:  123456.654321
         nStr.NumStrIn:  123456.654321
        nStr.NumStrOut:  123456.654321
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [49 50 51 52 53 54 54 53 52 51 50 49]
      nStr.AbsIntRunes:  [49 50 51 52 53 54]
     nStr.AbsFracRunes:  [54 53 52 51 50 49]
            nStr.precision:  6
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  123456
         nStr.NumStrIn:  123456
        nStr.NumStrOut:  123456
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [49 50 51 52 53 54]
      nStr.AbsIntRunes:  [49 50 51 52 53 54]
     nStr.AbsFracRunes:  []
            nStr.precision:  0
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  false
******************************************************


******************************************************
    Original Input Str:  0.123456
         nStr.NumStrIn:  0.123456
        nStr.NumStrOut:  0.123456
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [48 49 50 51 52 53 54]
      nStr.AbsIntRunes:  [48]
     nStr.AbsFracRunes:  [49 50 51 52 53 54]
            nStr.precision:  6
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  .123456
         nStr.NumStrIn:  .123456
        nStr.NumStrOut:  0.123456
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [49 50 51 52 53 54]
      nStr.AbsIntRunes:  [48]
     nStr.AbsFracRunes:  [49 50 51 52 53 54]
            nStr.precision:  6
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  1 2   3 4
         nStr.NumStrIn:  1234
        nStr.NumStrOut:  1234
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [49 50 51 52]
      nStr.AbsIntRunes:  [49 50 51 52]
     nStr.AbsFracRunes:  []
            nStr.precision:  0
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  false
******************************************************


******************************************************
    Original Input Str:  1 2   3 4 . 1 2 3 4 5 6
         nStr.NumStrIn:  1234.123456
        nStr.NumStrOut:  1234.123456
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [49 50 51 52 49 50 51 52 53 54]
      nStr.AbsIntRunes:  [49 50 51 52]
     nStr.AbsFracRunes:  [49 50 51 52 53 54]
            nStr.precision:  6
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  -32.495
         nStr.NumStrIn:  -32.495
        nStr.NumStrOut:  -32.495
          nStr.signVal:  -1
   nStr.AbsAllNumRunes:  [51 50 52 57 53]
      nStr.AbsIntRunes:  [51 50]
     nStr.AbsFracRunes:  [52 57 53]
            nStr.precision:  3
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  -.4219
         nStr.NumStrIn:  -.4219
        nStr.NumStrOut:  -0.4219
          nStr.signVal:  -1
   nStr.AbsAllNumRunes:  [52 50 49 57]
      nStr.AbsIntRunes:  [48]
     nStr.AbsFracRunes:  [52 50 49 57]
            nStr.precision:  4
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  -0.713
         nStr.NumStrIn:  -0.713
        nStr.NumStrOut:  -0.713
          nStr.signVal:  -1
   nStr.AbsAllNumRunes:  [48 55 49 51]
      nStr.AbsIntRunes:  [48]
     nStr.AbsFracRunes:  [55 49 51]
            nStr.precision:  3
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  Exit Status  1
         nStr.NumStrIn:  ExitStatus1
        nStr.NumStrOut:  1
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [49]
      nStr.AbsIntRunes:  [49]
     nStr.AbsFracRunes:  []
            nStr.precision:  0
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  false
******************************************************


******************************************************
    Original Input Str:  0
         nStr.NumStrIn:  0
        nStr.NumStrOut:  0
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [48]
      nStr.AbsIntRunes:  [48]
     nStr.AbsFracRunes:  []
            nStr.precision:  0
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  false
******************************************************


******************************************************
    Original Input Str:  5
         nStr.NumStrIn:  5
        nStr.NumStrOut:  5
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [53]
      nStr.AbsIntRunes:  [53]
     nStr.AbsFracRunes:  []
            nStr.precision:  0
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  false
******************************************************


******************************************************
    Original Input Str:  24.95
         nStr.NumStrIn:  24.95
        nStr.NumStrOut:  24.95
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [50 52 57 53]
      nStr.AbsIntRunes:  [50 52]
     nStr.AbsFracRunes:  [57 53]
            nStr.precision:  2
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  +24.95
         nStr.NumStrIn:  +24.95
        nStr.NumStrOut:  24.95
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  [50 52 57 53]
      nStr.AbsIntRunes:  [50 52]
     nStr.AbsFracRunes:  [57 53]
            nStr.precision:  2
nStr.HasNumericDigits :  true
nStr.IsFractionalValue:  true
******************************************************


******************************************************
    Original Input Str:  Nothing
         nStr.NumStrIn:  Nothing
        nStr.NumStrOut:  0
          nStr.signVal:  1
   nStr.AbsAllNumRunes:  []
      nStr.AbsIntRunes:  [48]
     nStr.AbsFracRunes:  []
            nStr.precision:  0
nStr.HasNumericDigits :  false
nStr.IsFractionalValue:  false
******************************************************

*/

func TestParseAndPrintOutNumStrs(str string) {

  ePrefix := "TestParseAndPrintOutNumStrs"

  nu := mathops.NumStrUtility{}

  nStr, err := nu.ParseNumString(str)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nStr, err := nu.ParseNumString(str)\n"+
      "str= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      str,
      err.Error())
    return
  }

  PrintNumStrDtoContents(str, nStr)

  return
}

func TestConvertNumStrToDecimal(str string) {

  nsu := mathops.NumStrUtility{}

  nsu.CurrencySymbol = '$'
  nsu.DecimalSeparator = '.'
  nsu.ThousandsSeparator = ','

  dec, err := nsu.ConvertNumStrToDecimal(str)

  if err != nil {
    panic(fmt.Errorf("TestConvertNumStrToDecimal() Error from nsu.ConvertNumStrToDecimal(str). str='%v' Error: %v ", str, err))
  }

  PrintDecimalContents(dec)

}

func TestScaleNumStr(numStr string, precision uint) {

  ePrefix := "TestScaleNumStr"

  nu := mathops.NumStrUtility{}

  nsDto, err := nu.ScaleNumStr(numStr, precision, true)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nsDto, err := nu.ScaleNumStr(\n"+
      " numStr, precision, true)\n"+
      "numStr= '%v'\n"+
      "precision= '%v'\n"+
      "roundResult= 'true'\n"+
      "Error='%v'\n\n",
      ePrefix,
      numStr,
      precision,
      err.Error())
    return
  }

  fmt.Printf("numStr= '%v' precision='%v' \n", numStr, precision)

  PrintNumStrDtoContents(numStr, nsDto)

}

func AddTwoDecimals(numStr1, numStr2 string) {

  ePrefix := "AddTwoDecimals"

  d1 := mathops.Decimal{}

  err := d1.SetNumStr(numStr1)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err := d1.SetNumStr(numStr1)\n"+
      "numStr1= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      numStr1,
      err.Error())
    return
  }

  d2 := mathops.Decimal{}

  err = d2.SetNumStr(numStr2)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = d2.SetNumStr(numStr2)\n"+
      "numStr2= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      numStr2,
      err.Error())
    return
  }

  d3, err := d1.Add(d2)
  fmt.Println("numStr1: ", numStr1)

  PrintDecimalContents(d1)

  fmt.Println("numStr2: ", numStr2)

  PrintDecimalContents(d2)

  fmt.Println()
  fmt.Println()

  PrintDecimalContents(d3)

  return
}

func PrintDecimalContents(dec mathops.Decimal) {

  ePrefix := "PrintDecimalContents"

  fmt.Println()
  fmt.Println("******************************************************")

  fmt.Println("               dec.IsValid: ", dec.GetIsValid())

  decNumStr, err := dec.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "str, err := dec.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("             dec.GetNumStr: ", decNumStr)

  sgn, err := dec.GetSign()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "sgn, err := dec.GetSign()\n"+
      "dec= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      decNumStr,
      err.Error())
    return
  }

  fmt.Println("               dec.GetSign: ", sgn)

  decAllDigitsNumStr, err := dec.GetSignedAllDigitsStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "decAllDigitsNumStr, err := dec.GetSignedAllDigitsStr()\n"+
      "decNumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      decNumStr,
      err.Error())
    return
  }

  fmt.Println(" dec.GetSignedAllDigitsStr: ", decAllDigitsNumStr)

  bScaleVal, err := dec.GetScaleVal()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bScaleVal, err := dec.GetScaleVal()\n"+
      "decNumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      decNumStr,
      err.Error())
    return
  }

  fmt.Println("         dec.GetScaleVal(): ", bScaleVal.Text(10))

  precision, err := dec.GetPrecision()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "precision, err := dec.GetPrecision()\n"+
      "precision is an 'int'\n"+
      "decNumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      decNumStr,
      err.Error())
    return
  }

  fmt.Println("          dec.GetPrecisionInt: ", precision)

  bf, err := dec.GetBigFloat()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bf, err := dec.GetBigFloat()\n"+
      "decNumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      decNumStr,
      err.Error())
    return
  }

  fmt.Println("           dec.GetBigFloat: ", bf.Text('e', 16))

  decBigFloatNumStr, err := dec.GetBigFloatString(uint(precision))

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "decBigFloatNumStr, err := \n"+
      " dec.GetBigFloatString(uint(precision))\n"+
      "decNumStr= '%v'\n"+
      "uint(precision)= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      decNumStr,
      uint(precision),
      err.Error())
    return
  }

  fmt.Println("        dec.GetFloatString: ", decBigFloatNumStr)
  fmt.Println("******************************************************")
  fmt.Println()

  return
}

func PrintNumStrDtoContents(originalInputStr string, nStr mathops.NumStrDto) {

  ePrefix := "PrintNumStrDtoContents"

  err := nStr.IsValid("Validating PrintNumStrDtoContents() 'nStr'")

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err := nStr.IsValid(\"Validating PrintNumStrDtoContents() 'nStr'\")\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nStrNumStr, err := nStr.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nStrSignValue, err := nStr.GetSign()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nStrSignValue, err := nStr.GetSign()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nStrAbsAllNumRunes, err := nStr.GetAbsAllNumRunes()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nStrAbsAllNumRunes, err := nStr.GetAbsAllNumRunes()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nStrAbsIntRunes, err := nStr.GetAbsIntRunes()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nStrAbsIntRunes, err := nStr.GetAbsIntRunes()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nStrAbsFracRunes, err := nStr.GetAbsFracRunes()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nStrAbsFracRunes, err := nStr.GetAbsFracRunes()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nStrPrecisionInt := nStr.GetPrecision()
	
  fmt.Println()
  fmt.Println("******************************************************")
  fmt.Println("          Original Input Str: ", originalInputStr)
  fmt.Println("            nStr.GetNumStr(): ", nStrNumStr)
  fmt.Println("              nStr.GetSign(): ", nStrSignValue)
  fmt.Println("    nStr.GetAbsAllNumRunes(): ", nStrAbsAllNumRunes)
  fmt.Println("       nStr.GetAbsIntRunes(): ", nStrAbsIntRunes)
  fmt.Println("      nStr.GetAbsFracRunes(): ", nStrAbsFracRunes)
  fmt.Println("         nStr.GetPrecisionInt(): ", nStrPrecisionInt)
  fmt.Println("     nStr.HasNumericDigits(): ", nStr.HasNumericDigits())
  fmt.Println("      nStr.IsFractionalValue: ", nStr.IsFractionalValue())
  fmt.Println("nStr.GetThousandsSeparator(): ", nStr.GetThousandsSeparator())
  fmt.Println("  nStr.GetDecimalSeparator(): ", nStr.GetDecimalSeparator())
  fmt.Println("    nStr.GetCurrencySymbol(): ", nStr.GetCurrencySymbol())
  fmt.Println("******************************************************")
  fmt.Println("            nStr.IsValid(): ", "nStr IS VALID!")
  fmt.Println("******************************************************")
  fmt.Println()

  return
}
