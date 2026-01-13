package mathops

import (
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
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
//  bigIntMathNthRootBoson.setupBundles
//  bigIntMathNthRootBoson.calcBundleLength
//  bigIntMathNthRootBoson.findNextRoot
//
//  ------------------------------------
//
//  bigIntMathNthRootQuark.getNextBundleBigIntValue
//  bigIntMathNthRootQuark.calcPrecision
//
//  -------------------------------------
//
//  bigIntMathNthRootProton.initializeBigIntMathNthRoot
//  bigIntMathNthRootProton.doRootExtraction
//  bigIntMathNthRootProton.empty
//
//  -------------------------------------
//
//  bigIntMathNthRootNeutron.calcNthRootGateway
//
//  -------------------------------------
//
//  bigIntMathNthRootAtom.calcPositiveFractionalNthRoot
//
//  -------------------------------------
//
//  bigIntMathNthRootNanobot.calcPositiveIntegerNthRoot
//
//  -------------------------------------
//
//  bigIntMathNthRootMacrobot.calcPositiveNthRoot
//
//  bigIntMathNthRootMacrobot.calcNegativeNthRoot

// Empty
//
//	Resets all the internal member variables to their initial or zero
//	values for the current instance of BigIntMathNthRoot.
func (nthrt *BigIntMathNthRoot) Empty() {

	new(bigIntMathNthRootProton).empty(nthrt)
}

// GetNthRoot
//
//	Calculates the Nth Root of a real number ('radicand') passed to
//	the method as Type BigIntNum.  The calling function must supply
//	input parameters for 'radicand', 'nthRoot' and 'maxPrecision'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will copy the Numeric Separators configured
//	for input parameter 'radicand' to the returned instance of
//	BigIntNum.
//
//	Input Parameters
//	================
//
//	radicand                 BigIntNum
//	  The radicand value from which the nth Root will be taken.
//	            nthRootResult^nthRoot = radicand
//
//	nthRoot                  BigIntNum
//	  Specifies the root which will be calculated for parameter,
//	  'radicand'. Examples: square root, cube root, 4th root, 9th
//	  root, etc.
//
//	  'nthRoot' is a BigIntNum Type, which may be a positive or
//	  negative number. In addition, the nthRoot may be either an
//	  integer number or a fractional number.
//
//	  The nthRoot must be a numeric value greater than one ('1') or
//	  less than minus one (-1). nthRoots with a value of zero will
//	  always return an nthRoot result of zero. Nth Root values of +1
//	  or -1 will generate an error.
//
//	  If the radicand is negative and the nthRoot value is an even
//	  number (evenly divisible by 2 with no remainder), an error will
//	  be returned since the result of such a calculation is an
//	  imaginary number.
//
//	maxPrecision             uint
//	  Specifies the maximum number of decimals to the right of the
//	  decimal point to which the Nth root result will be calculated.
//
//	Return Values
//	=============
//
//		BigIntNum
//	  If the calculation is successful, the nth root result will be
//	  returned as a BigIntNum type. This returned BigIntNum nth root
//	  will contain numeric separators (decimal separator, thousands
//	  separator and currency symbol) copied from input parameter,
//	  'radicand'.
//
//	error
//	  If the calculation completes successfully, the 'error' type
//	  returned will be set equal to 'nil'. If an error is encountered,
//	  the returned 'error' type will contain an appropriate error
//	  message.
func (nthrt *BigIntMathNthRoot) GetNthRoot(
	radicand BigIntNum,
	nthRoot BigIntNum,
	maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathNthRoot.GetNthRoot",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntMathNthRootMechanics).getNthRoot(
		nthrt, &radicand, true, &nthRoot, true, maxPrecision, ePrefix)
}
