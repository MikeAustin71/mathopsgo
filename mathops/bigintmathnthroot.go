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
//
//  x bigIntMathNthRootBoson.setupBundles
//  x bigIntMathNthRootBoson.calcBundleLength
//  x bigIntMathNthRootBoson.findNextRoot
//
//  ------------------------------------
//
//  x bigIntMathNthRootQuark.getNextBundleBigIntValue
//  x bigIntMathNthRootQuark.calcPrecision
//
//  -------------------------------------
//
//  x bigIntMathNthRootProton.initializeBigIntMathNthRoo
//  x bigIntMathNthRootProton.doRootExtraction
//
//  -------------------------------------
//
//  x bigIntMathNthRootNeutron.calcNthRootGateway
//
//  -------------------------------------
//
//  x bigIntMathNthRootAtom.calcPositiveFractionalNthRoot
//

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
