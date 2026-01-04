package examples

import (
  "fmt"
  "math/big"

  "github.com/mikeaustin71/mathops"
)

func ExampleBINumSetFromIntFracStrings(
  intStr, fracStr, expectedResult string,
  signVal int,
  numSeps mathops.NumericSeparatorDto) {

  ePrefix := "ExampleBINumSetFromIntFracStrings"

  bINum := mathops.BigIntNum{}

  err := bINum.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err := bINum.SetNumericSeparatorsDto(numSeps)\n"+
      "numSeps= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      numSeps.String(),
      err.Error())
    return
  }

  err = bINum.SetIntFracStrings(intStr, fracStr, signVal)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetIntFracStrings(intStr, fracStr, signVal)\n"+
      "intStr= '%v'\n"+
      "fracStr= '%v'\n"+
      "signVal= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      intStr,
      fracStr,
      signVal,
      err.Error())
    return
  }

  bINumNumSeps, err := bINum.GetNumericSeparatorsDto()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  bINumNumStr, err := bINum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println()
  fmt.Println("BigIntNum{}.ShiftPrecisionRight()")
  fmt.Println("---------------------------------")

  if expectedResult == bINumNumStr {
    fmt.Println("*** SUCCESS ***")
  } else {
    fmt.Println("@@@ FAILURE @@@")
  }

  fmt.Println("-------------------------")
  fmt.Println("      Int String: ", intStr)
  fmt.Println("     Frac String: ", fracStr)
  fmt.Println("   Actual Result: ", bINumNumStr)
  fmt.Println(" Expected Result: ", expectedResult)
  fmt.Println("-------------------------")

  if !numSeps.Equal(bINumNumSeps) {
    fmt.Println("ERROR - Numeric Separators NOT Equal")
  } else {
    fmt.Println("SUCCESS - Numeric Separators ARE Equal!")
  }

  fmt.Println("-------------------------")
  fmt.Println("  Actual Numeric Separators: ", bINumNumSeps.String())
  fmt.Println("Expected Numeric Separators: ", numSeps.String())

  return
}

func ExampleShiftPrecisionRight(baseNumStr string, shiftPlacesLeft uint, expectedResult string) {

  ePrefix := "ExampleShiftPrecisionRight()"

  bigIntNum, err := new(mathops.BigIntNum).NewNumStr(baseNumStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  err = bigIntNum.ShiftPrecisionRight(shiftPlacesLeft)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = bigIntNum.ShiftPrecisionRight(shiftPlacesLeft)\n"+
      "shiftPlacesLeft= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      shiftPlacesLeft,
      err.Error())
    return
  }

  actualResult, err := bigIntNum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "binRootPrecisionInt, err := binRoot.GetPrecisionInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println()
  fmt.Println("BigIntNum{}.ShiftPrecisionRight()")
  fmt.Println("---------------------------------")

  if expectedResult == actualResult {
    fmt.Println("*** SUCCESS ***")
  } else {
    fmt.Println("@@@ FAILURE @@@")
  }
  fmt.Println("-------------------------")
  fmt.Println("      Base Number: ", baseNumStr)
  fmt.Println("Shift Places Left: ", shiftPlacesLeft)
  fmt.Println("    Actual Result: ", actualResult)
  fmt.Println("  Expected Result: ", expectedResult)

}

func ExampleShiftPrecisionLeft(baseNumStr string, shiftPlacesLeft uint, expectedResult string) {

  ePrefix := "ExampleShiftPrecisionLeft()"

  bigIntNum, err := new(mathops.BigIntNum).NewNumStr(baseNumStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bigIntNum, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(baseNumStr)\n"+
      "baseNumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      baseNumStr,
      err.Error())
    return
  }

  err = bigIntNum.ShiftPrecisionLeft(shiftPlacesLeft)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = bigIntNum.ShiftPrecisionLeft(shiftPlacesLeft)\n"+
      "shiftPlacesLeft= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      shiftPlacesLeft,
      err.Error())
    return
  }

  actualResult, err := bigIntNum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "actualResult, err := bigIntNum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println()
  fmt.Println("BigIntNum{}.ShiftPrecisionLeft()")
  fmt.Println("-------------------------")

  if expectedResult == actualResult {
    fmt.Println("*** SUCCESS ***")
  } else {
    fmt.Println("@@@ FAILURE @@@")
  }
  fmt.Println("-------------------------")
  fmt.Println("      Base Number: ", baseNumStr)
  fmt.Println("Shift Places Left: ", shiftPlacesLeft)
  fmt.Println("    Actual Result: ", actualResult)
  fmt.Println("  Expected Result: ", expectedResult)
}

func ExampleSetBigFloat01(bigFloat *big.Float, maxPrecision uint, expectedResult string) {

  ePrefix := "ExampleSetBigFloat01()"

  bINum, err := new(mathops.BigIntNum).NewZero(0)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bINum, err := new(mathops.BigIntNum).NewZero(0)\n"+
      "precision= '0'\n"+
      "Error='%v'\n\n",
      ePrefix,
      err.Error())
    return
  }

  err = bINum.SetBigFloat(bigFloat, maxPrecision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = bINum.SetBigFloat(bigFloat, maxPrecision)\n"+
      "target= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      bigFloat.Text('f', -1),
      maxPrecision,
      err.Error())
    return
  }

  actualNumStr, err := bINum.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "actualNumStr, err := bINum.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println()
  fmt.Println("BigIntNum{}.SetBigFloat()")
  fmt.Println("-------------------------")

  if expectedResult == actualNumStr {
    fmt.Println("*** SUCCESS ***")
  } else {
    fmt.Println("@@@ FAILURE @@@")
  }
  fmt.Println("-------------------------")
  fmt.Println("       bigFloat: ", bigFloat.Text('f', -1))
  fmt.Println("   maxPrecision: ", maxPrecision)
  fmt.Println("  Actual Result: ", actualNumStr)
  fmt.Println("Expected Result: ", expectedResult)

  return
}

func ExampleBundleCount01(bINumTarget, bINumNthRoot mathops.BigIntNum, expectedBundleCnt string) {

  ePrefix := "ExampleBundleCount01()"

  target, err := bINumTarget.GetBigInt()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "target, err := bINumTarget.GetBigInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nthRoot, err := bINumNthRoot.GetBigInt()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nthRoot, err := bINumNthRoot.GetBigInt()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  fmt.Println("      Original Target: ", target.Text(10))

  newTarget, err := ExampleBundleCount02(target, nthRoot)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "newTarget, err := ExampleBundleCount02(target, nthRoot)\n"+
      "target= '%v'\n"+
      "nthRoot= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      target.Text(10),
      nthRoot.Text(10),
      err.Error())
    return
  }

  fmt.Println("           New Target: ", newTarget.Text(10))

  bundleCnt, err := ExampleBundleCount03(newTarget, nthRoot)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "bundleCnt, err := ExampleBundleCount03(newTarget, nthRoot)\n"+
      "newTarget= '%v'\n"+
      "nthRoot= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      newTarget.Text(10),
      nthRoot.Text(10),
      err.Error())
    return
  }

  fmt.Println("         Bundle Count: ", bundleCnt.Text(10))
  fmt.Println("Expected Bundle Count: ", expectedBundleCnt)

  return
}

func ExampleBundlePrecisionCount03(
  intBundleCnt *big.Int,
  maxPrecision uint) (totalBundleCnt, precisionBundleCnt *big.Int, err error) {

  err = nil

  totalBundleCnt = big.NewInt(0).Add(intBundleCnt, big.NewInt(int64(maxPrecision)))

  precisionBundleCnt = big.NewInt(0)

  return
}

func ExampleBundleCount03(target, nthRoot *big.Int) (bundleCnt *big.Int, err error) {
  // Guaranteed: Magnitude of target is Greater Than or Equal to nthRoot
  bundleCnt = big.NewInt(0)
  err = nil

  ePrefix := "ExampleBundleCount03() "

  magnitude, err := new(mathops.BigIntMath).GetMagnitude(target)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(mathops.BigIntMath).\n"+
      " GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      target.Text(10),
      err.Error())
    return
  }

  bigOne := big.NewInt(1)

  numOfDigits := big.NewInt(0).Add(magnitude, bigOne)

  quotient, mod := big.NewInt(0).QuoRem(
    numOfDigits,
    nthRoot,
    big.NewInt(0))

  bundleCnt = big.NewInt(0).Set(quotient)

  if mod.Cmp(big.NewInt(0)) == 1 {
    bundleCnt = big.NewInt(0).Add(bundleCnt, bigOne)
  }

  return bundleCnt, nil
}

func ExampleBundleCount02(target, nthRoot *big.Int) (newTarget *big.Int, err error) {
  // FormatTargetInt

  ePrefix := "ExampleBundleCount02() "

  newTarget = big.NewInt(0)
  err = nil
  magnitude, err := new(mathops.BigIntMath).GetMagnitude(target)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "magnitude, err := new(mathops.BigIntMath).\n"+
      " GetMagnitude(target)\n"+
      "target= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      target.Text(10),
      err.Error())
    return
  }

  baseTen := big.NewInt(10)
  newTarget = big.NewInt(0).Set(target)
  numOfDigits := big.NewInt(0).Add(magnitude, big.NewInt(1))

  for numOfDigits.Cmp(nthRoot) == -1 {

    newTarget = big.NewInt(0).Mul(newTarget, baseTen)
    numOfDigits = big.NewInt(0).Add(numOfDigits, big.NewInt(1))

  }

  return newTarget, nil
}

func ExampleBigIntNumNthRoot01(
  radicandStr, nthRootStr string,
  maxPrecision uint,
  expectedNumStr string) {

  ePrefix := "ExampleBigIntNumNthRoot01() "

  radicand, err := new(mathops.BigIntNum).NewNumStr(radicandStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "radicand, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(radicandStr)\n"+
      "radicandStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      radicandStr,
      err.Error())
    return
  }

  radicandNumStr, err := radicand.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "radicandNumStr, err := radicand.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  nthRoot, err := new(mathops.BigIntNum).NewNumStr(nthRootStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nthRoot, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(nthRootStr)\n"+
      "nthRootStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      nthRootStr,
      err.Error())
    return
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "nthRootNumStr, err := nthRoot.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  result, err := new(mathops.BigIntMathNthRoot).GetNthRoot(radicand, nthRoot, maxPrecision)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "result, err := new(mathops.BigIntMathNthRoot).\n"+
      "  GetNthRoot(radicand, nthRoot, maxPrecision)\n"+
      "radicand= '%v'\n"+
      "nthRoot= '%v'\n"+
      "maxPrecision= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      radicandNumStr,
      nthRootNumStr,
      maxPrecision,
      err.Error())
    return
  }

  err = result.IsValid("Validating result")

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "err = result.IsValid(\"Validating result\")\n"+
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

  expectedResult, err := new(mathops.BigIntNum).NewNumStr(expectedNumStr)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedResult, err := new(mathops.BigIntNum).\n"+
      " NewNumStr(expectedNumStr)\n"+
      "expectedNumStr= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedNumStr,
      err.Error())
    return
  }

  expectedResultNumStr, err := expectedResult.GetNumStr()

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedResultNumStr, err := expectedResult.GetNumStr()\n"+
      "Error='%v'\n\n",
      ePrefix, err.Error())
    return
  }

  if expectedNumStr != expectedResultNumStr {
    fmt.Printf("%v\n"+
      "Error: 'expectedResult' Initialization FAILED!\n"+
      "Because expectedNumStr != expectedResultNumStr\n"+
      "Expected expectedResultNumStr = '%v'\n"+
      "  Actual expectedResultNumStr = '%v'\n\n",
      ePrefix, expectedNumStr, expectedResultNumStr)

    return
  }

  fmt.Println("*** BigIntMathNthRoot ***")
  fmt.Println("Expected Result: ", expectedNumStr)
  fmt.Println("  Actual Result: ", resultNumStr)
  fmt.Println("  Max Precision: ", maxPrecision)

  fmt.Println("*** BigIntMathNthRoot Outcome ***")

  expectedResultEqualsActualResult, err := expectedResult.Equal(result)

  if err != nil {
    fmt.Printf("%v\n"+
      "Error returned by:\n"+
      "expectedResultEqualsActualResult, err :=\n"+
      " expectedResult.Equal(result)\n"+
      "expectedResult= '%v'\n"+
      "result= '%v'\n"+
      "Error='%v'\n\n",
      ePrefix,
      expectedResultNumStr,
      resultNumStr,
      err.Error())
    return
  }

  if !expectedResultEqualsActualResult {

    fmt.Printf("%v\n"+
      "Error: Expected vs Actual Results DON'T MATCH!\n"+
      "Because expectedResultEqualsActualResult == false\n"+
      "Expected result = '%v'\n"+
      "  Actual result = '%v'\n\n",
      ePrefix, expectedResultNumStr, resultNumStr)

    fmt.Println("           Base: ", radicandNumStr)
    fmt.Println("        NthRootInt: ", nthRootNumStr)

    return

  }

  fmt.Println("SUCCESS - Expected Result Matched Actual Result!")
  fmt.Println()
  fmt.Println("           Base: ", radicandNumStr)
  fmt.Println("        NthRootInt: ", nthRootNumStr)
}
