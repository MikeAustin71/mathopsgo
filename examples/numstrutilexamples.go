package common

import (
	
	"fmt"
	"../mathops"
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

func TestNumStrUtility_ParseNumString() {
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
	nu := mathops.NumStrUtility{}

	nStr, err := nu.ParseNumString(str)

	if err != nil {
		fmt.Printf("Error from nu.ParseNumString(str). str='%v' Error: %v\n", str, err.Error())
		return
	}

	PrintNumStrDtoContents(str, nStr)

}

func TestConvertNumStrToDecimal(str string) {

	nsu := mathops.NumStrUtility{}
	dec, err := nsu.ConvertNumStrToDecimal(str)

	if err != nil {
		panic(fmt.Errorf("TestConvertNumStrToDecimal() Error from nsu.ConvertNumStrToDecimal(str). str='%v' Error: %v ", str, err))
	}

	PrintDecimalContents(dec)

}

func TestScaleNumStr(numStr string, precision uint) {

	nu := mathops.NumStrUtility{}

	nsDto, err := nu.ScaleNumStr(numStr, precision, true)

	if err != nil {
		fmt.Printf("TestScaleNumStr(): Error returned from nu.ScaleNumStr(). numStr='%v' precision='%v' Error: %v", numStr, precision, err)
		return
	}

	fmt.Printf("numStr= '%v' precision='%v' \n", numStr, precision)

	PrintNumStrDtoContents(numStr, nsDto)

}

func AddTwoDecimals(numStr1, numStr2 string) {

	d1 := mathops.Decimal{}
	err := d1.SetNumStr(numStr1)

	if err != nil {
		fmt.Printf("Error from SetNumStr(numStr1). Error = %v", err)
		return
	}

	d2 := mathops.Decimal{}

	err = d2.SetNumStr(numStr2)

	if err != nil {
		fmt.Printf("Error from 2nd SetNumStr(numStr1). Error = %v", err)
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

}

func PrintDecimalContents(dec mathops.Decimal) {

	fmt.Println()
	fmt.Println("******************************************************")

	fmt.Println("               dec.IsValid: ", dec.GetIsValid())
	str := dec.GetNumStr()
	fmt.Println("             dec.GetNumStr: ", str)
	sgn := dec.GetSign()
	fmt.Println("               dec.GetSign: ", sgn)
	str, _ = dec.GetSignedAllDigitsStr()
	fmt.Println(" dec.GetSignedAllDigitsStr: ", str)
	bScaleVal, _ := dec.GetScaleVal()
	fmt.Println("         dec.GetScaleVal(): ", bScaleVal.String())
	precision := dec.GetPrecision()
	fmt.Println("          dec.GetPrecision: ", precision)
	bf, _ := dec.GetBigFloat()
	fmt.Println("           dec.GetBigFloat: ", bf.Text('e', 16))
	str, _ = dec.GetBigFloatString(uint(precision))
	fmt.Println("        dec.GetFloatString: ", str)
	fmt.Println("******************************************************")
	fmt.Println()

}

func PrintNumStrDtoContents(originalInputStr string, nStr mathops.NumStrDto) {
	fmt.Println()
	fmt.Println("******************************************************")
	fmt.Println("    Original Input Str: ", originalInputStr)
	fmt.Println("         nStr.NumStrIn: ", nStr.NumStrIn)
	fmt.Println("        nStr.NumStrOut: ", nStr.NumStrOut)
	fmt.Println("          nStr.signVal: ", nStr.SignVal)
	fmt.Println("   nStr.AbsAllNumRunes: ", nStr.AbsAllNumRunes)
	fmt.Println("      nStr.AbsIntRunes: ", nStr.AbsIntRunes)
	fmt.Println("     nStr.AbsFracRunes: ", nStr.AbsFracRunes)
	fmt.Println("        nStr.precision: ", nStr.GetPrecision())
	fmt.Println("nStr.HasNumericDigits : ", nStr.HasNumericDigits)
	fmt.Println("nStr.IsFractionalValue: ", nStr.IsFractionalValue)
	fmt.Println("******************************************************")
	fmt.Println()

}
