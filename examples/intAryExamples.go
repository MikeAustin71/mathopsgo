package examples

import (
  "fmt"
  "time"

  "github.com/mikeaustin71/mathops"
)

func ExampleIntAryPwrByTwos01(
  baseStr, exponentStr, expectedResult string,
  minResultPrecision, maxResultPrecision int) {

  ePrefix := "ExampleIntAryPwrByTwos01"

  iaBase, err := new(mathops.IntAry).NewNumStr(baseStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaBase, err := new(mathops.IntAry).NewNumStr(baseStr)\n"+
      "baseStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      baseStr,
      err.Error())
    return
  }

  iaBaseNumStr, err := iaBase.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaBaseNumStr, err := iaBase.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  iaExponent, err := new(mathops.IntAry).NewNumStr(exponentStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaExponent, err := new(mathops.IntAry).NewNumStr(exponentStr)\n"+
      "exponentStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      exponentStr,
      err.Error())
    return
  }

  iaExponentNumStr, err := iaExponent.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaExponentNumStr, err := iaExponent.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  var t1 time.Time
  var t0 time.Time

  t0 = time.Now()

  err = new(mathops.IntAryMathPower).Pwr(
    &iaBase,
    &iaExponent,
    minResultPrecision,
    maxResultPrecision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = new(mathops.IntAryMathPower).Pwr(\n"+
      " &iaBase, &iaExponent, minResultPrecision, maxResultPrecision)\n"+
      "iaBase= '%v'\n"+
      "iaExponent= '%v'\n"+
      "minResultPrecision= '%v'\n"+
      "maxResultPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaBaseNumStr,
      iaExponentNumStr,
      minResultPrecision,
      maxResultPrecision,
      err.Error())
    return
  }

  t1 = time.Now()

  actualResult, err := iaBase.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  timeStr := CodeDurationToStr(t1.Sub(t0))

  fmt.Println()
  fmt.Println("IntAryMathPower{}.Pwr()")
  fmt.Println("=======================")

  if expectedResult != actualResult {

    fmt.Println("XXX FAILURE XXX")

  } else {

    fmt.Println("*** SUCCESS ***")

  }

  fmt.Println("=======================")
  fmt.Println("           base: ", baseStr)
  fmt.Println("       exponent: ", exponentStr)
  fmt.Println("  Actual Result: ", actualResult)
  fmt.Println("Expected Result: ", expectedResult)
  fmt.Println("   Elapsed Time: ", timeStr)

}

func ExampleIntAryGetMagnitude01(numStr string, expectedMagnitude int) {

  ePrefix := "ExampleIntAryGetMagnitude01"

  iaNum, err := new(mathops.IntAry).NewNumStr(numStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaNum, err := new(mathops.IntAry).\n"+
      " NewNumStr(numStr)\n"+
      "numStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      numStr,
      err.Error())
    return
  }

  iaNumNumStr, err := iaNum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaNumNumStr, err := iaNum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  bInt, err := iaNum.GetBigInt()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bInt, err := iaNum.GetBigInt()\n"+
      "iaNum= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaNumNumStr,
      err.Error())
    return
  }

  bIMagnitude, err := new(mathops.BigIntMath).GetMagnitude(bInt)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bIMagnitude, err := new(mathops.BigIntMath).\n"+
      " GetMagnitude(bInt)\n"+
      "bInt= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bInt.Text(10),
      err.Error())
    return
  }

  magnitude, err := iaNum.GetMagnitudeDigits()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := iaNum.GetMagnitudeDigits()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  iaMagnitude, err := iaNum.GetMagnitude()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaMagnitude, err := iaNum.GetMagnitude()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("       Target Number: ", numStr)
  fmt.Println("    Actual Magnitude: ", magnitude)
  fmt.Println("  Expected Magnitude: ", expectedMagnitude)
  fmt.Println(" IntAry.GetMagnitude: ", iaMagnitude)
  fmt.Println("-------------------------------------")
  fmt.Println("BigIntMath Magnitude: ", bIMagnitude.Text(10))

}

func ExampleIntAryMultiplyPower01(baseStr, exponentStr, expectedResultStr string, minPrecision, maxPrecision int) {

  ePrefix := "ExampleIntAryMultiplyPower01()"

  var t1 time.Time
  var t0 time.Time

  iaBase, err := new(mathops.IntAry).NewNumStr(baseStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaBase, err := new(mathops.IntAry).\n"+
      " NewNumStr(baseStr)\n"+
      "baseStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      baseStr,
      err.Error())
    return
  }

  iaBaseNumStr, err := iaBase.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaBaseNumStr, err := iaBase.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  iaExponent, err := new(mathops.IntAry).NewNumStr(exponentStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaExponent, err := new(mathops.IntAry).\n"+
      " NewNumStr(exponentStr)\n"+
      "exponentStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      exponentStr,
      err.Error())
    return
  }

  iaExponentNumStr, err := iaExponent.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "iaExponentNumStr, err := iaExponent.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  t0 = time.Now()

  actualResult, err := new(mathops.IntAryMathPower).PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "actualResult, err := new(mathops.IntAryMathPower).\n"+
      " PwrByMultiplication(&iaBase, &iaExponent, minPrecision, maxPrecision)\n"+
      "iaBase= '%v'\n"+
      "iaExponent= '%v'\n"+
      "minPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      iaBaseNumStr,
      iaExponentNumStr,
      minPrecision,
      maxPrecision,
      err.Error())
    return
  }

  t1 = time.Now()

  actualResultNumStr, err := actualResult.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "actualResultNumStr, err := actualResult.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println()
  fmt.Println("PwrByMultiplication()")
  fmt.Println("=====================")

  if expectedResultStr == actualResultNumStr {

    fmt.Println("*** SUCCESS ***")

  } else {

    fmt.Println("XXX FAILURE XXX")

  }
  fmt.Println("=====================")

  codeDurationStr := CodeDurationToStr(t1.Sub(t0))

  fmt.Println("     Input Base: ", baseStr)
  fmt.Println(" Input Exponent: ", exponentStr)
  fmt.Println("         Result: ", actualResultNumStr)
  fmt.Println("Expected Result: ", expectedResultStr)
  fmt.Println("   Elapsed Time: ", codeDurationStr)

}

func ExampleIntAryDivide01(dividendStr, divisorStr, eQuotient string, minPrecision, maxPrecision int) {

  ePrefix := "ExampleIntAryDivide01()"

  dividend, err := new(mathops.IntAry).NewNumStr(dividendStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "dividend, err := new(mathops.IntAry).\n"+
      " NewNumStr(dividendStr)\n"+
      "dividendStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendStr,
      err.Error())
    return
  }

  dividendNumStr, err := dividend.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "dividendNumStr, err := dividend.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  divisor, err := new(mathops.IntAry).NewNumStr(divisorStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "divisor, err := new(mathops.IntAry).\n"+
      " NewNumStr(divisorStr)\n"+
      "divisorStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      divisorStr,
      err.Error())
    return
  }

  divisorNumStr, err := divisor.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "divisorNumStr, err := divisor.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  quotient, err := new(mathops.IntAryMathDivide).Divide(
    &dividend,
    &divisor,
    minPrecision,
    maxPrecision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "quotient, err := new(mathops.IntAryMathDivide).Divide(\n"+
      " &dividend, &divisor, minPrecision, maxPrecision)\n"+
      "dividend= '%v'\n"+
      "divisor= '%v'\n"+
      "minPrecision= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      dividendNumStr,
      divisorNumStr,
      minPrecision,
      maxPrecision,
      err.Error())
    return
  }

  actualStr, err := quotient.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "actualStr, err := quotient.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if eQuotient != actualStr {
    fmt.Printf("%v\n"+
      "Error: 'eQuotient' and 'actualStr' ARE NOT EQUAL!\n"+
      "Because eQuotient != actualStr\n"+
      "Expected actualStr = '%v'\n"+
      "  Actual actualStr = '%v'\n\n",
      ePrefix, eQuotient, actualStr)

    return
  }

  fmt.Println("*** SUCCESS ***")
  fmt.Println("Expected Quotient: ", eQuotient)
  fmt.Println("  Actual Quotient: ", actualStr)
  fmt.Println("         Dividend: ", dividendStr)
  fmt.Println("          Divisor: ", divisorStr)

  return
}
