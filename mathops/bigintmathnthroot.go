package mathops

import (
  "errors"
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
)

// BigIntMathNthRoot - Used to extract square roots and nth roots of positive and negative
// real numbers. Nth Roots may be either integer or fractional numbers. In addition, Nth
// Roots may be either positive or integer numbers.
//
// The technique employed to calculate nth roots is known as the
// "shifting nth-root algorithm".
//
// This source file is located in source code repository:
//
// https://github.com/MikeAustin71/mathopsgo.git
//
// See: https://en.wikipedia.org/wiki/Shifting_nth_root_algorithm
type BigIntMathNthRoot struct {
  NthRoot               BigIntNum
  OriginalRadicand      BigIntNum
  SetupRadicand         BigIntNum
  IntBundleRadicand     BigIntNum
  FracBundleRadicand    BigIntNum
  BundleAddOnPrecision  *big.Int
  FracBundleLength      *big.Int
  TotalBundleLength     *big.Int
  ResultBInt            *big.Int
  ActualResultPrecision *big.Int
  FracPrecision         *big.Int
  //ResultPrecision    int
  ResultBINum        BigIntNum
  RequestedPrecision uint
  BigOne             *big.Int
  Big10              *big.Int
  Big10ToNthPower    *big.Int
  BigZero            *big.Int
  Y                  *big.Int // Root Extracted thusfar
  YPrime             *big.Int // Next Value of Y
  Minuend            *big.Int
  Subtrahend         *big.Int
  R                  *big.Int // Let R be the remainder
  RPrime             *big.Int // Let RPrime be the new value of r for next iteration
  BaseNum            *big.Int // Base Number System - always 10
  Alpha              *big.Int // Next n-digits of the radicand
  Beta               *big.Int // Next Digit of the root
}

// Low-Level Routines
//  x bigIntMathNthRootBoson.setupBundles
//  bigIntMathNthRootBoson.calcBundleLength
//  bigIntMathNthRootBoson.findNextRoot

// Empty
//
//	Resets all the internal member variables to their initial or zero
//	values for the current instance of BigIntMathNthRoot.
func (nthrt *BigIntMathNthRoot) Empty() {

  new(bigIntMathNthRootProton).empty(nthrt)
}

// GetNthRoot
//
//	 Calculates the Nth Root of a real number ('radicand') passed to
//	 the method as Type BigIntNum.  The calling function must supply
//	 input parameters for 'radicand', 'nthRoot' and 'maxPrecision'.
//
//	 Input Parameters
//	 ================
//
//	 radicand                 BigIntNum
//	   The radicand value from which the nth Root will be taken.
//	             nthRootResult^nthRoot = radicand
//
//	 nthRoot                  BigIntNum
//	   Specifies the root which will be calculated for parameter,
//	   'radicand'. Examples: square root, cube root, 4th root, 9th root
//	   etc.
//
//	   'nthRoot' is a BigIntNum Type which may be a positive or
//	   negative number. In addition, the nthRoot may be either an
//	   integer number or a fractional number.
//
//	   The nthRoot must be a numeric value greater than one ('1') or
//	   less than minus one (-1). nthRoots with a value of zero will
//	   always return an nthRoot result of zero. Nth Root values of +1
//	   or -1 will generate an error.
//
//	   If the radicand is negative and the nthRoot value is an even
//	   number (evenly divisible by 2 with no remainder), an error will
//	   be returned since the result of such a calculation is an
//	   imaginary number.
//
//	 maxPrecision             uint
//	   Specifies the maximum number of decimals to the right of the
//	   decimal point to which the Nth root result will be calculated.
//
//	 Return Values
//	 =============
//
//		BigIntNum
//	   If the calculation is successful, the nth root result will be
//	   returned as a BigIntNum type. This returned BigIntNum nth root
//	   will contain numeric separators (decimal separator, thousands
//	   separator and currency symbol) copied from input parameter,
//	   'radicand'.
//
//	 error
//	   If the calculation completes successfully, the 'error' type
//	   returned will be set equal to 'nil'. If an error is encountered,
//	   the returned 'error' type will contain an appropriate error
//	   message.
func (nthrt *BigIntMathNthRoot) GetNthRoot(
  radicand BigIntNum,
  nthRoot BigIntNum,
  maxPrecision uint) (BigIntNum, error) {

  ePrefix := "BigIntMathNthRoot.OriginalNthRoot() "

  if radicand.GetSign() == -1 {

    isEvenNum, err := nthRoot.IsEvenNumber()

    if err != nil {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+
          "Error returned by nthRoot.IsEvenNumber() "+
          "nthRoot='%v' Error='%v'\n", nthRoot.GetNumStr(), err.Error())
    }

    if isEvenNum {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+
          "INVALID ENTRY - Cannot calculate nthRoot of a negative radicand when nthRoot is even. "+
          "Original Number= %v  nthRoot= %v\n", radicand.GetNumStr(), nthRoot.GetNumStr())
    }

  }

  // If the radicand is zero, the result will always be zero
  if radicand.IsZero() {
    return radicand, nil
  }

  numSeps := radicand.GetNumericSeparatorsDto()

  bigINumOne := new(BigIntNum).NewOne(0)

  var err error

  err = bigINumOne.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix + "Error returned by bigINumOne.SetNumericSeparatorsDto(numSeps).")
  }

  // If nthRoot is zero, the result will always be '1'
  if nthRoot.IsZero() {
    return bigINumOne, nil
  }

  // Error if nthRoot == 1
  if nthRoot.Cmp(bigINumOne) == 0 {
    return new(BigIntNum).NewZero(0),
      errors.New(ePrefix +
        "Error - Input Parameter 'nthRoot' INVALID! 'nthRoot' cannot equal 1.\n")
  }

  var nthRootResult BigIntNum

  if nthRoot.GetSign() == -1 {

    nthRootResult, err = nthrt.calcNegativeNthRoot(radicand, nthRoot, maxPrecision)

    if err != nil {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+"Error returned by nthrt.calcNegativeNthRoot(...). "+
          "Error='%v' \n", err.Error())
    }

  } else {

    nthRootResult, err = nthrt.calcPositiveNthRoot(radicand, nthRoot, maxPrecision)

    if err != nil {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+"Error returned by nthrt.calcPositiveNthRoot(...). "+
          "Error='%v' \n", err.Error())
    }

  }

  err = nthRootResult.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix + "Error returned by nthRootResult.SetNumericSeparatorsDto(numSeps).")
  }

  return nthRootResult, nil
}

//                             Stage 1                                             //
// ******************************************************************************* //

// calcPositiveNthRoot - Calculates the nth root of a radicand where the nth root
// is a positive value.
func (nthrt *BigIntMathNthRoot) calcPositiveNthRoot(radicand, nthRoot BigIntNum,
  maxPrecision uint) (BigIntNum, error) {

  ePrefix := "BigIntMathNthRoot.calcPositiveNthRoot() "

  if nthRoot.GetSign() == -1 {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+"Error: This method only calculates nthRoot results for positive "+
        "nthRoot values. The entry for nthRoot is negative. nthRoot='%v'\n", nthRoot.GetNumStr())
  }

  if nthRoot.GetPrecisionUint() > 0 {
    return nthrt.calcPositiveFractionalNthRoot(radicand, nthRoot, maxPrecision)
  }

  return nthrt.calcPositiveIntegerNthRoot(radicand, nthRoot, maxPrecision)

}

// calcNegativeNthRoot - calculates the nth root result of a radicand where the
// nth root is a negative value.
func (nthrt *BigIntMathNthRoot) calcNegativeNthRoot(radicand, nthRoot BigIntNum,
  maxPrecision uint) (BigIntNum, error) {

  ePrefix := "BigIntMathNthRoot.calcNegativeNthRoot() "

  if nthRoot.GetSign() != -1 {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+"Error: This method only calculates nthRoot results for positive "+
        "nthRoot values. The entry for nthRoot is negative. nthRoot='%v'\n", nthRoot.GetNumStr())
  }

  newNthRoot := nthRoot.CopyOut()

  newNthRoot.ChangeSign()

  var nthRootResult BigIntNum
  var err error

  if newNthRoot.GetPrecisionUint() == 0 {
    nthRootResult, err = nthrt.calcPositiveIntegerNthRoot(radicand, newNthRoot, maxPrecision)
    if err != nil {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+"Error returned by calcPositiveIntegerNthRoot(...) "+
          "Error='%v' \n", err.Error())
    }
  } else {
    nthRootResult, err = nthrt.calcPositiveFractionalNthRoot(radicand, newNthRoot, maxPrecision)
    if err != nil {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+"Error returned by calcPositiveFractionalNthRoot(...) "+
          "Error='%v' \n", err.Error())
    }
  }

  inverse, err := nthRootResult.Inverse(maxPrecision)

  if err != nil {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+"Error returned by nthRootResult.Inverse(maxPrecision) "+
        "maxPrecision='%v' Error='%v' \n", maxPrecision, err.Error())

  }

  return inverse, nil
}

//                             Stage 2                                             //
// ******************************************************************************* //

func (nthrt *BigIntMathNthRoot) calcPositiveIntegerNthRoot(radicand, nthRoot BigIntNum,
  maxPrecision uint) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathNthRoot.calcPositiveIntegerNthRoot",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if nthRoot.GetSign() != 1 {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+
        "Error: Expected 'nthRoot' to be positive. nthRoot is negative! "+
        "nthRoot='%v' ", nthRoot.GetNumStr())
  }

  if nthRoot.GetPrecisionInt() != 0 {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+
        "Error: Expected 'nthRoot' to be integer value. nthRoot is fractional value! "+
        "nthRoot='%v' ", nthRoot.GetNumStr())
  }

  if radicand.GetSign() == -1 {

    isEvenNum, err := nthRoot.IsEvenNumber()

    if err != nil {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+
          "Error returned by nthRoot.IsEvenNumber() "+
          "nthRoot='%v' Error='%v'\n", nthRoot.GetNumStr(), err.Error())
    }

    if isEvenNum {
      return new(BigIntNum).NewZero(0),
        fmt.Errorf(ePrefix+
          "INVALID ENTRY - Cannot calculate nthRoot of a negative number when nthRoot is even. "+
          "Original Number= %v  nthRoot= %v\n", radicand.GetNumStr(), nthRoot.GetNumStr())
    }

  }

  bINumResult, err := new(bigIntMathNthRootNeutron).calcNthRootGateway(nthrt, &radicand, true, &nthRoot, true, maxPrecision, ePrefix)

  if err != nil {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+
        "Error returned by nthrt.calcNthRootGateway(radicand, nthRoot, maxPrecision) "+
        "Error='%v' ", err.Error())
  }

  return bINumResult, nil
}

// calcPrecision - This method will calculate the actual precision of the
// result output by the Nth Root calculation.  Actual result precision is
// equal to the Bundle Add On Precision + the quotient of radicand precision
// divided by nthRoot.
//
//	bundleAddOnPrecision + Quotient(radicandPrecision/nthRoot)
func (nthrt *BigIntMathNthRoot) calcPrecision(
  radicandPrecision,
  fracBundleLength,
  precisionAdjustment,
  bundleAddOnPrecision,
  nthRoot *big.Int) (actualPrecision *big.Int, err error) {

  actualPrecision = big.NewInt(0)
  err = nil
  bigZero := big.NewInt(0)

  if nthRoot.Cmp(bigZero) == 0 {
    ePrefix := "BigIntMathNthRoot.calcPrecision() "
    err = errors.New(ePrefix +
      "Error: nthRoot=0. Division by nthRoot is Division by zero and will FAIL!")
    return actualPrecision, err

  }

  actualPrecision = big.NewInt(0).Add(bundleAddOnPrecision, fracBundleLength)
  actualPrecision = big.NewInt(0).Add(actualPrecision, precisionAdjustment)

  if actualPrecision.Cmp(bigZero) < 0 {
    actualPrecision = big.NewInt(0)
  }

  /*
  	fmt.Println("    fracBundleLength: ", fracBundleLength.Text(10))
  	fmt.Println("bundleAddOnPrecision: ", bundleAddOnPrecision.Text(10))
  	fmt.Println("precision Adjustment: ", precisionAdjustment.Text(10))
  	fmt.Println("    actual precision: ", actualPrecision.Text(10))
  */
  err = nil

  return actualPrecision, err
}

// getNextBundleBigIntValue - Calculate and return the next bundle of
// numeric digits for nth Root extraction calculations.
//

func (nthrt *BigIntMathNthRoot) getNextBundleBigIntValue(

  intBundleRadicand, fracBundleRadicand, nthRoot BigIntNum) (
  nextBundleValue *big.Int, newIntBundleRadicand, newFracBundleRadicand BigIntNum, err error) {

  ePrefix := "BigIntMathNthRoot.getNextBundleBigIntValue() "
  nextBundleValue = big.NewInt(0)
  newIntBundleRadicand = new(BigIntNum).NewZero(0)
  newIntBundleRadicand.SetExpectedToActualNumberOfDigits()
  newFracBundleRadicand = new(BigIntNum).NewZero(0)
  newFracBundleRadicand.SetExpectedToActualNumberOfDigits()
  err = nil
  modX := big.NewInt(0)

  var errx error
  var magnitude *big.Int
  var numberOfDigits *big.Int
  var tempRadicand *big.Int
  var digitsQuotient *big.Int
  var exponent *big.Int
  var divisor *big.Int

  intExpectedNumOfDigits := intBundleRadicand.GetExpectedNumberOfDigits()
  intActualNumOfDigits, _, errx := intBundleRadicand.GetActualNumberOfDigits()
  if errx != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by intBundleRadicand.GetActualNumberOfDigits(). "+
      "Error='%v' ", errx.Error())

    return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
  }

  intExpectedActualDelta := big.NewInt(0).Sub(intExpectedNumOfDigits, intActualNumOfDigits)
  intModExpectedNthRoot := big.NewInt(0).Rem(intExpectedNumOfDigits, nthRoot.GetAbsoluteBigIntValue())
  intIsZeroValue := intBundleRadicand.IsZero()

  bigZero := big.NewInt(0)

  fracExpectedNumOfDigits := fracBundleRadicand.GetExpectedNumberOfDigits()
  fracActualNumOfDigits, _, errx := fracBundleRadicand.GetActualNumberOfDigits()

  if errx != nil {
    err = fmt.Errorf(ePrefix+
      "Error returned by fracBundleRadicand.GetActualNumberOfDigits(). "+
      "Error='%v' ", errx.Error())

    return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
  }

  fracExpectedActualDelta := big.NewInt(0).Sub(fracExpectedNumOfDigits, fracActualNumOfDigits)
  fracIsZeroValue := fracBundleRadicand.IsZero()

  if (intExpectedActualDelta.Cmp(bigZero) > 0 &&
    intExpectedActualDelta.Cmp(nthRoot.GetAbsoluteBigIntValue()) >= 0) ||
    (intExpectedActualDelta.Cmp(bigZero) > 0 &&
      intIsZeroValue) {

    // fmt.Println("Calc nthRoot IntegerDeltaDigits >=0")

    nextBundleValue = big.NewInt(0)
    newIntBundleRadicand = intBundleRadicand.CopyOut()
    newFracBundleRadicand = fracBundleRadicand.CopyOut()
    intExpectedNumOfDigits = big.NewInt(0).Sub(intExpectedNumOfDigits, nthRoot.GetAbsoluteBigIntValue())

    if intExpectedNumOfDigits.Cmp(intActualNumOfDigits) == -1 {
      intExpectedNumOfDigits = big.NewInt(0).Set(intActualNumOfDigits)
    }

    newIntBundleRadicand.SetExpectedNumberOfDigits(intExpectedNumOfDigits)

    err = nil

  } else if !intBundleRadicand.IsZero() {

    // fmt.Println("Calc inBundleRadicand is NOT zero")

    magnitude, errx = BigIntMath{}.GetMagnitude(intBundleRadicand.GetAbsoluteBigIntValue())

    if errx != nil {
      err = fmt.Errorf(ePrefix+
        "Error returned by BigIntMath{}.GetMagnitudeDigits(intBundleRadicand.GetAbsoluteBigIntValue()). "+
        "intBundleRadicand='%v' Error='%v'",
        intBundleRadicand.GetNumStr(), errx.Error())

      nextBundleValue = big.NewInt(0)
      newIntBundleRadicand = new(BigIntNum).NewZero(0)
      newFracBundleRadicand = new(BigIntNum).NewZero(0)

      return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
    }

    intActualNumOfDigits = big.NewInt(0).Add(magnitude, big.NewInt(1))
    digitsQuotient = big.NewInt(0).Quo(magnitude, nthrt.NthRoot.GetAbsoluteBigIntValue())
    exponent = big.NewInt(0).Mul(digitsQuotient, nthrt.NthRoot.GetAbsoluteBigIntValue())
    divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

    nextBundleValue, tempRadicand =
      big.NewInt(0).QuoRem(intBundleRadicand.GetAbsoluteBigIntValue(), divisor, modX)

    numberOfDigits, errx = BigIntMath{}.GetMagnitude(nextBundleValue)

    if errx != nil {
      err = fmt.Errorf(ePrefix+"Error returned by BigIntMath{}.GetMagnitudeDigits(nextBundleValue). "+
        "nextBundleValue='%v' Error='%v'",
        nextBundleValue.Text(10), errx.Error())

      nextBundleValue = big.NewInt(0)
      newIntBundleRadicand = new(BigIntNum).NewZero(0)
      newFracBundleRadicand = new(BigIntNum).NewZero(0)

      return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
    }

    numberOfDigits = big.NewInt(0).Add(numberOfDigits, big.NewInt(1))

    if intModExpectedNthRoot.Cmp(bigZero) > 0 {
      intExpectedNumOfDigits = big.NewInt(0).Sub(intExpectedNumOfDigits, numberOfDigits)
    } else {
      intExpectedNumOfDigits = big.NewInt(0).Sub(intExpectedNumOfDigits, nthrt.NthRoot.GetAbsoluteBigIntValue())
    }

    newIntBundleRadicand = new(BigIntNum).NewBigInt(tempRadicand, 0)
    newIntBundleRadicand.SetExpectedNumberOfDigits(intExpectedNumOfDigits)
    newFracBundleRadicand = fracBundleRadicand.CopyOut()
    err = nil

  } else if (fracExpectedActualDelta.Cmp(bigZero) > 0 &&
    fracExpectedActualDelta.Cmp(nthRoot.GetAbsoluteBigIntValue()) >= 0) ||
    (fracExpectedActualDelta.Cmp(bigZero) > 0 && fracIsZeroValue) {

    // fmt.Println("Calc nthRoot FracDeltaDigits >= 0")

    nextBundleValue = big.NewInt(0)
    newIntBundleRadicand = intBundleRadicand.CopyOut()
    newFracBundleRadicand = fracBundleRadicand.CopyOut()
    fracExpectedNumOfDigits = big.NewInt(0).Sub(fracExpectedNumOfDigits, nthRoot.GetAbsoluteBigIntValue())

    if fracExpectedNumOfDigits.Cmp(fracActualNumOfDigits) == -1 {
      fracExpectedNumOfDigits = big.NewInt(0).Set(fracActualNumOfDigits)
    }

    newFracBundleRadicand.SetExpectedNumberOfDigits(fracExpectedNumOfDigits)
    nthrt.FracPrecision = big.NewInt(0).Add(nthrt.FracPrecision, big.NewInt(1))
    err = nil

  } else if !fracBundleRadicand.IsZero() {

    // fmt.Println("Calc fracBundleRadicand Is NOT zero")

    magnitude, errx = BigIntMath{}.GetMagnitude(fracBundleRadicand.GetAbsoluteBigIntValue())

    if errx != nil {

      err = fmt.Errorf(ePrefix+
        "Error returned by BigIntMath{}.GetMagnitudeDigits(fracBundleRadicand.GetAbsoluteBigIntValue()). "+
        "fracBundleRadicand='%v' Error='%v' ",
        fracBundleRadicand.GetNumStr(), errx.Error())

      nextBundleValue = big.NewInt(0)
      newIntBundleRadicand = new(BigIntNum).NewZero(0)
      newFracBundleRadicand = new(BigIntNum).NewZero(0)
      return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
    }

    digitsQuotient = big.NewInt(0).Quo(magnitude, nthrt.NthRoot.GetAbsoluteBigIntValue())
    exponent = big.NewInt(0).Mul(digitsQuotient, nthrt.NthRoot.GetAbsoluteBigIntValue())
    divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

    nextBundleValue, tempRadicand =
      big.NewInt(0).QuoRem(fracBundleRadicand.GetAbsoluteBigIntValue(), divisor, modX)

    newIntBundleRadicand = new(BigIntNum).NewZero(0)
    newIntBundleRadicand.SetExpectedToActualNumberOfDigits()
    fracExpectedNumOfDigits = big.NewInt(0).Sub(fracExpectedNumOfDigits, nthrt.NthRoot.GetAbsoluteBigIntValue())
    newFracBundleRadicand = new(BigIntNum).NewBigInt(tempRadicand, 0)
    newFracBundleRadicand.SetExpectedNumberOfDigits(fracExpectedNumOfDigits)
    nthrt.FracPrecision = big.NewInt(0).Add(nthrt.FracPrecision, big.NewInt(1))
    err = nil

  } else {

    // fmt.Println("Calc else zero")

    nextBundleValue = big.NewInt(0)
    newIntBundleRadicand = new(BigIntNum).NewZero(0)
    newIntBundleRadicand.SetExpectedToActualNumberOfDigits()
    newFracBundleRadicand = new(BigIntNum).NewZero(0)
    newFracBundleRadicand.SetExpectedToActualNumberOfDigits()
    nthrt.FracPrecision = big.NewInt(0).Add(nthrt.FracPrecision, big.NewInt(1))
    err = nil
  }

  return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
}
