package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type bigIntMathNthRootBoson struct {
  lock sync.Mutex
}

// Experimental 2
func (bIMathNthrtBoson *bigIntMathNthRootBoson) setupBundles(
  radicand *BigIntNum,
  validateRadicand bool,
  nthRoot *BigIntNum,
  validateNthRoot bool,
  errPrefDto *ePref.ErrPrefixDto) (setupRadicand BigIntNum,
  intBundleRadicand BigIntNum,
  fracBundleRadicand BigIntNum,
  precisionAdjustment *big.Int,
  err error) {

  bIMathNthrtBoson.lock.Lock()

  defer bIMathNthrtBoson.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathNthRootBoson.setupBundles",
    "")

  if err != nil {
    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0), err
  }

  if radicand == nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'radicand'",
      }
  }

  if nthRoot == nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'nthRoot'",
      }
  }

  if validateRadicand {

    err = radicand.IsValid(ePrefix.XCpy("Validating 'radicand'").String())

    if err != nil {

      return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = radicand.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'radicand'\").String()",
          ErrContext: "Error: Input parameter 'radicand' is INVALID!\n" +
            "'radicand' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNthRoot {

    err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

    if err != nil {

      return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = nthRoot.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'nthRoot'\").String()",
          ErrContext: "Error: Input parameter 'nthRoot' is INVALID!\n" +
            "'nthRoot' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  radicandNumStr, err := radicand.GetNumStr()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandNumStr, err := radicand.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  intBundleRadicand = new(BigIntNum).New()

  fracBundleRadicand = new(BigIntNum).New()

  precisionAdjustment = big.NewInt(0)

  modX := big.NewInt(0)

  bigTen := big.NewInt(10)

  scaleValue := big.NewInt(0)

  radicandAbsoluteBInt, err := radicand.GetAbsoluteBigIntValue()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandAbsoluteBInt, err := radicand.GetAbsoluteBigIntValue()",
        ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  radicandPrecisionUint, err := radicand.GetPrecisionUint()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandPrecisionUint, err := radicand.GetPrecisionUint()",
        ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  setupRadicand, err = new(BigIntNum).NewBigInt(
    radicandAbsoluteBInt, radicandPrecisionUint)

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "setupRadicand, err = new(BigIntNum).NewBigInt(\n" +
          "  radicandAbsoluteBInt, radicandPrecisionUint)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  setupRadicandNumStr, err := setupRadicand.GetNumStr()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "setupRadicandNumStr, err := setupRadicand.GetNumStr()",
        ErrContext: "#1 GetNumStr()",
        ErrMessage: err.Error(),
      }
  }

  err = setupRadicand.TrimTrailingFracZeros()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = setupRadicand.TrimTrailingFracZeros()",
        ErrContext: fmt.Sprintf("setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  setupRadicandNumStr, err = setupRadicand.GetNumStr()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "setupRadicandNumStr, err = setupRadicand.GetNumStr()",
        ErrContext: fmt.Sprintf("#2 setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  err = setupRadicand.SetExpectedToActualNumberOfDigits()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = setupRadicand.SetExpectedToActualNumberOfDigits()",
        ErrContext: fmt.Sprintf("setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  setupRadicandNumStr, err = setupRadicand.GetNumStr()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "setupRadicandNumStr, err = setupRadicand.GetNumStr()",
        ErrContext: fmt.Sprintf("#3 setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  setupRadicandPrecisionUint, err := setupRadicand.GetPrecisionUint()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "setupRadicandPrecisionUint, err := setupRadicand.GetPrecisionUint()",
        ErrContext: fmt.Sprintf("setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  if setupRadicandPrecisionUint == 0 {

    intBundleRadicand, err = setupRadicand.CopyOut()

    if err != nil {

      return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "intBundleRadicand, err = setupRadicand.CopyOut()",
          ErrContext: fmt.Sprintf("setupRadicand= '%v'", setupRadicandNumStr),
          ErrMessage: err.Error(),
        }
    }

    /*
       fmt.Println("setupRadicand: ", setupRadicand.GetNumStr())
       fmt.Println("intBundleRadicand:", intBundleRadicand.GetNumStr())
       fmt.Println("fracBundleRadicand:", fracBundleRadicand.GetNumStr())
    */

    return setupRadicand, intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
  }

  setupRadicandTotalDigits, _, err := setupRadicand.GetActualNumberOfDigits()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "setupRadicandTotalDigits, _, err := \n" +
          "  setupRadicand.GetActualNumberOfDigits()",
        ErrContext: fmt.Sprintf("setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  // Precision must be greater than zero
  radicandPrecision, err := setupRadicand.GetPrecisionBigInt()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandPrecision, err := setupRadicand.GetPrecisionBigInt()",
        ErrContext: fmt.Sprintf("setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  setupRadicandIntegerDigits :=
    big.NewInt(0).Sub(setupRadicandTotalDigits, radicandPrecision)

  expectedFractionalDigits := big.NewInt(0).Set(radicandPrecision)

  scaleValue = big.NewInt(0).Exp(bigTen, radicandPrecision, nil)

  setupRadicandAbsoluteBInt, err := setupRadicand.GetAbsoluteBigIntValue()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "setupRadicandAbsoluteBInt, err :=\n" +
          "  setupRadicand.GetAbsoluteBigIntValue()",
        ErrContext: fmt.Sprintf("setupRadicand= '%v'", setupRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  intBundleRadicandBigInt, fracBundleRadicandBigInt :=
    big.NewInt(0).QuoRem(setupRadicandAbsoluteBInt, scaleValue, modX)

  intBundleRadicand, err = new(BigIntNum).NewBigInt(intBundleRadicandBigInt, 0)

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "intBundleRadicand, err = new(BigIntNum).\n" +
          "  NewBigInt(intBundleRadicandBigInt, 0)",
        ErrContext: fmt.Sprintf("intBundleRadicandBigInt= '%v'",
          intBundleRadicandBigInt.Text(10)),
        ErrMessage: err.Error(),
      }
  }

  err = intBundleRadicand.SetExpectedNumberOfDigits(setupRadicandIntegerDigits)

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = intBundleRadicand.SetExpectedNumberOfDigits(\n" +
          "  setupRadicandIntegerDigits)",
        ErrContext: fmt.Sprintf("setupRadicandIntegerDigits= '%v'",
          setupRadicandIntegerDigits.Text(10)),
        ErrMessage: err.Error(),
      }
  }

  intBundleRadicandNumStr, err := intBundleRadicand.GetNumStr()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "intBundleRadicandNumStr, err := \n" +
          "  intBundleRadicand.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nthRootAbsoluteBigInt, err := nthRoot.GetAbsoluteBigIntValue()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootAbsoluteBigInt, err := nthRoot.GetAbsoluteBigIntValue()",
        ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
        ErrMessage: err.Error(),
      }
  }

  mod := big.NewInt(0).Rem(radicandPrecision, nthRootAbsoluteBigInt)

  if mod.Cmp(big.NewInt(0)) == 1 {

    delta := big.NewInt(0).Sub(nthRootAbsoluteBigInt, mod)

    scaleVal := big.NewInt(0).Exp(big.NewInt(10), delta, nil)

    fracBundleRadicandBigInt = big.NewInt(0).Mul(fracBundleRadicandBigInt, scaleVal)

    expectedFractionalDigits = big.NewInt(0).Add(expectedFractionalDigits, delta)
  }

  fracBundleRadicand, err = new(BigIntNum).NewBigInt(fracBundleRadicandBigInt, 0)

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "fracBundleRadicand, err = new(BigIntNum).NewBigInt(\n" +
          "  fracBundleRadicandBigInt, 0)",
        ErrContext: fmt.Sprintf("fracBundleRadicandBigInt= '%v'",
          fracBundleRadicandBigInt.Text(10)),
        ErrMessage: err.Error(),
      }
  }

  fracBundleRadicandNumStr, err := fracBundleRadicand.GetNumStr()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "fracBundleRadicandNumStr, err := fracBundleRadicand.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  intBundleRadicandIsZero, err := intBundleRadicand.IsZero()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("intBundleRadicand= '%v'", intBundleRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  fracBundleRadicandIsZero, err := fracBundleRadicand.IsZero()

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("fracBundleRadicand= '%v'",
          fracBundleRadicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  if intBundleRadicandIsZero &&
    fracBundleRadicandIsZero {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Both intBundleRadicand and fracBundleRadicand are ZERO!",
      }
  }

  err = fracBundleRadicand.SetExpectedNumberOfDigits(expectedFractionalDigits)

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = fracBundleRadicand.SetExpectedNumberOfDigits(\n" +
          "  expectedFractionalDigits)",
        ErrContext: fmt.Sprintf("expectedFractionalDigits = '%v'",
          expectedFractionalDigits.Text(10)),
        ErrMessage: err.Error(),
      }
  }

  err = setupRadicand.IsValid(ePrefix.XCpy("Validating 'setupRadicand'").String())

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = setupRadicand.IsValid(ePrefix.XCpy(\n" +
          "  \"Validating 'setupRadicand'\").String())",
        ErrContext: "Error: Final calculated result 'setupRadicand' is INVALID!\n" +
          "'setupRadicand' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  err = intBundleRadicand.IsValid(ePrefix.XCpy("Validating 'intBundleRadicand'").String())

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = intBundleRadicand.IsValid(ePrefix.XCpy(\n" +
          "  \"Validating 'intBundleRadicand'\").String())",
        ErrContext: "Error: Final calculated result 'intBundleRadicand' is INVALID!\n" +
          "'intBundleRadicand' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  err = fracBundleRadicand.IsValid(ePrefix.XCpy("Validating 'fracBundleRadicand'").String())

  if err != nil {

    return BigIntNum{}, BigIntNum{}, BigIntNum{}, big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = fracBundleRadicand.IsValid(ePrefix.XCpy(\n" +
          "  \"Validating 'fracBundleRadicand'\").String())",
        ErrContext: "Error: Final calculated result 'fracBundleRadicand' is INVALID!\n" +
          "'fracBundleRadicand' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return setupRadicand, intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
}

// calcBundleLength
//
//	Calculates the final bundle length. These are bundles of integers
//	that are packaged for submission to the Nth Root calculation
//	routine.
//
//	The final bundle length is equal to the actual number of bundles
//	associated with the radicand plus bundles added on for additional
//	precision in the final nthRoot result.
func (bIMathNthrtBoson *bigIntMathNthRootBoson) calcBundleLength(
  radicand *BigIntNum,
  validateRadicand bool,
  nthRoot *BigIntNum,
  validateNthRoot bool,
  bundleAddonPrecision *big.Int,
  errPrefDto *ePref.ErrPrefixDto) (totalBundleLength *big.Int,
  fracBundleLength *big.Int, err error) {

  bIMathNthrtBoson.lock.Lock()

  defer bIMathNthrtBoson.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathNthRootBoson.calcBundleLength",
    "")

  if err != nil {
    return big.NewInt(0), big.NewInt(0), err
  }

  totalBundleLength = big.NewInt(0)

  fracBundleLength = big.NewInt(0)

  if radicand == nil {

    return totalBundleLength, fracBundleLength,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'radicand'",
      }
  }

  if nthRoot == nil {

    return totalBundleLength, fracBundleLength,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'nthRoot'",
      }
  }

  if validateRadicand {

    err = radicand.IsValid(ePrefix.XCpy("Validating 'radicand'").String())

    if err != nil {

      return totalBundleLength, fracBundleLength,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = radicand.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'radicand'\").String()",
          ErrContext: "Error: Input parameter 'radicand' is INVALID!\n" +
            "'radicand' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateNthRoot {

    err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

    if err != nil {

      return totalBundleLength, fracBundleLength,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = nthRoot.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'nthRoot'\").String()",
          ErrContext: "Error: Input parameter 'nthRoot' is INVALID!\n" +
            "'nthRoot' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  radicandNumStr, err := radicand.GetNumStr()

  if err != nil {

    return totalBundleLength, fracBundleLength,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandNumStr, err := radicand.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return totalBundleLength, fracBundleLength,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  modX := big.NewInt(0)

  bigZero := big.NewInt(0)

  bigOne := big.NewInt(1)

  numOfDigits := big.NewInt(1)

  intBundleLength := big.NewInt(0)

  intBundleLengthMod := big.NewInt(0)

  bigIntAbsoluteNthRoot, err := nthRoot.GetAbsoluteBigIntValue()

  if err != nil {

    return totalBundleLength, fracBundleLength,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bigIntAbsoluteNthRoot, err := nthRoot.GetAbsoluteBigIntValue()",
        ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
        ErrMessage: err.Error(),
      }
  }

  radicandPrecisionUint, err := radicand.GetPrecisionUint()

  if err != nil {

    return totalBundleLength, fracBundleLength,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandPrecisionUint, err := radicand.GetPrecisionUint()",
        ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  scaleValue := big.NewInt(0).Exp(
    big.NewInt(10),
    big.NewInt(int64(radicandPrecisionUint)),
    nil)

  radicandBigIntAbsolute, err := radicand.GetAbsoluteBigIntValue()

  if err != nil {

    return totalBundleLength, fracBundleLength,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "radicandBigIntAbsolute, err := radicand.GetAbsoluteBigIntValue()",
        ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
        ErrMessage: err.Error(),
      }
  }

  // break radicand into integer digits
  intRadicand :=
    big.NewInt(0).Quo(
      radicandBigIntAbsolute,
      scaleValue)

  if intRadicand.Cmp(bigZero) == 1 {

    numOfDigits, err = new(BigIntMath).GetMagnitude(intRadicand)

    if err != nil {

      return totalBundleLength, fracBundleLength,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "numOfDigits, err = new(BigIntMath).GetMagnitude(intRadicand)",
          ErrContext: fmt.Sprintf("intRadicand = '%v'", intRadicand.Text(10)),
          ErrMessage: err.Error(),
        }
    }

    // Convert integer radicand to number of digits
    numOfDigits = big.NewInt(0).Add(numOfDigits, bigOne)

    intBundleLength, intBundleLengthMod = big.NewInt(0).QuoRem(
      numOfDigits,
      bigIntAbsoluteNthRoot,
      modX)

    // For the integer bundle calculation add one to bundle length if
    // there are any remaining digits.
    // Example  5 / 3 = Quotient of 1 Mod of 2. Add 1 to bundle size
    if intBundleLengthMod.Cmp(bigZero) == 1 {

      intBundleLength = big.NewInt(0).Add(intBundleLength, bigOne)
    }
  }

  radicandPrecision := big.NewInt(int64(radicandPrecisionUint))

  if radicandPrecision.Cmp(bigZero) == 1 {

    fracBundleLength = big.NewInt(0).Quo(radicandPrecision, bigIntAbsoluteNthRoot)
  }

  /*
  	fmt.Println("        intBundleLength: ", intBundleLength.Text(10))
  	fmt.Println("       fracBundleLength: ", fracBundleLength.Text(10))
  	fmt.Println( "  bundleAddonPrecision: ", bundleAddonPrecision.Text(10))
  */

  totalBundleLength = big.NewInt(0).Add(intBundleLength, fracBundleLength)

  totalBundleLength = big.NewInt(0).Add(totalBundleLength, bundleAddonPrecision)

  return totalBundleLength, fracBundleLength, err
}

// findNextRoot - Called by doRootExtraction() to find the next
// root value.
func (bIMathNthrtBoson *bigIntMathNthRootBoson) findNextRoot(
  nthrt *BigIntMathNthRoot,
  errPrefDto *ePref.ErrPrefixDto) error {

  bIMathNthrtBoson.lock.Lock()

  defer bIMathNthrtBoson.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathNthRootBoson.findNextRoot",
    "")

  if err != nil {
    return err
  }

  if nthrt == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ErrContext:    "",
      ParameterName: "'nthrt'",
    }
  }

  var bundle *big.Int

  bundle, nthrt.IntBundleRadicand, nthrt.FracBundleRadicand, err =
    new(bigIntMathNthRootQuark).
      getNextBundleBigIntValue(
        nthrt,
        &nthrt.IntBundleRadicand,
        true,
        &nthrt.FracBundleRadicand,
        true,
        &nthrt.NthRoot,
        true,
        ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "bundle, nthrt.IntBundleRadicand, nthrt.FracBundleRadicand, err =\n" +
        "  getNextBundleBigIntValue(nthrt, &nthrt.IntBundleRadicand, true,\n" +
        "  &nthrt.FracBundleRadicand, true, &nthrt.NthRoot, true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthrtNthRootAbsoluteBigInt, err := nthrt.NthRoot.GetAbsoluteBigIntValue()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "nthrtNthRootAbsoluteBigInt, err :=\n" +
        "  nthrt.NthRoot.GetAbsoluteBigIntValue()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  // alpha = next n-digits of radicand
  nthrt.Alpha = big.NewInt(0).Set(bundle)

  nthrt.RPrime = big.NewInt(-1)

  itatr := big.NewInt(9)
  term1a := big.NewInt(0)
  term1b := big.NewInt(0)

  term2a := big.NewInt(0)
  term2a1 := big.NewInt(0)
  term2a2 := big.NewInt(0)

  term2b := big.NewInt(0)
  term2b1 := big.NewInt(0)
  term2b2 := big.NewInt(0)

  term1a = big.NewInt(0).Mul(nthrt.Big10ToNthPower, nthrt.R)
  term1b = big.NewInt(0).Set(nthrt.Alpha)

  nthrt.Minuend = big.NewInt(0).Add(term1a, term1b)

  for itatr.Cmp(nthrt.BigZero) > -1 &&
    nthrt.RPrime.Cmp(nthrt.BigZero) == -1 {

    nthrt.Beta = big.NewInt(0).Set(itatr)
    nthrt.YPrime = big.NewInt(0).Mul(nthrt.Y, nthrt.Big10)
    nthrt.YPrime = big.NewInt(0).Add(nthrt.YPrime, nthrt.Beta)

    term2a1 = big.NewInt(0).Mul(nthrt.BaseNum, nthrt.Y)
    term2a2 = big.NewInt(0).Add(term2a1, nthrt.Beta)

    term2a = big.NewInt(0).Exp(term2a2,
      big.NewInt(0).Set(nthrtNthRootAbsoluteBigInt),
      nil)

    term2b1 = big.NewInt(0).Set(nthrt.Big10ToNthPower)

    term2b2 = big.NewInt(0).Exp(nthrt.Y,
      big.NewInt(0).Set(nthrtNthRootAbsoluteBigInt),
      nil)

    term2b = big.NewInt(0).Mul(term2b1, term2b2)

    nthrt.Subtrahend = big.NewInt(0).Sub(term2a, term2b)

    nthrt.RPrime = big.NewInt(0).Sub(nthrt.Minuend, nthrt.Subtrahend)

    itatr = big.NewInt(0).Sub(itatr, nthrt.BigOne)
  }

  nthrt.R = big.NewInt(0).Set(nthrt.RPrime)

  nthrt.Y = big.NewInt(0).Set(nthrt.YPrime)

  nthrt.ResultBInt = big.NewInt(0).Mul(nthrt.ResultBInt, nthrt.Big10)

  nthrt.ResultBInt = big.NewInt(0).Add(nthrt.ResultBInt, nthrt.Beta)

  return nil
}
