package mathops

import (
	"math/big"
	"fmt"
	"errors"
)

// BigIntMathNthRoot - Used to extract square roots and nth roots of positive and negative
// real numbers. Currently nth roots must be positive and the value must NOT exceed the
// maximum int-32 value of +2,147,483,647. As a practical matter, your computer may run out
// of memory before you hit this maximum.
//
// The technique employed to calculate nth roots is known as the
// "shifting nth-root algorithm".
//
// This source file is located in source code repository:
//
// https://github.com/MikeAustin71/mathhlpr.git
//
//
// See: https://en.wikipedia.org/wiki/Shifting_nth_root_algorithm
//
// Dependencies: intAry - intary.go
//
type BigIntMathNthRoot struct {
	NthRoot            		BigIntNum
	OriginalRadicand   		BigIntNum
	BundleRadicand     		*big.Int
	BundleAddOnPrecision 	*big.Int
	BundleLength       		*big.Int
	ResultBInt         		*big.Int
	AcutalResultPrecision *big.Int
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

func (nthrt *BigIntMathNthRoot) Empty() {
	nthrt.NthRoot = BigIntNum{}.NewZero(0)
	nthrt.OriginalRadicand = BigIntNum{}.NewZero(0)
	nthrt.BundleRadicand = big.NewInt(0)
	nthrt.BundleAddOnPrecision = big.NewInt(0)
	nthrt.BundleLength = big.NewInt(0)
	nthrt.ResultBInt = big.NewInt(0)
	nthrt.ResultBINum = BigIntNum{}.NewZero(0)
	nthrt.RequestedPrecision = 0
	nthrt.BigOne = big.NewInt(0)
	nthrt.Big10 = big.NewInt(0)
	nthrt.Big10ToNthPower = big.NewInt(0)
	nthrt.BigZero = big.NewInt(0)
	nthrt.YPrime = big.NewInt(0)
	nthrt.Minuend = big.NewInt(0)
	nthrt.Subtrahend = big.NewInt(0)
	nthrt.R = big.NewInt(0)
	nthrt.RPrime = big.NewInt(0)
	nthrt.BaseNum = big.NewInt(0)
	nthrt.Alpha = big.NewInt(0)
	nthrt.Beta = big.NewInt(0)
}

// GetNthRoot  - Calculates the Nth Root of a real number ('num')
// passed to the method as Type BigIntNum.  In addition, the caller must supply
// input parameters for 'nthRoot' and 'maxPrecision'.
//
// 'nthRoot' specifies the root which will be calculated for parameter, 'num'. Example,
// square root, cube root, 4th root, 9th root etc. 'nthRoot' is a BigIntNum Type which
// must be a positive integer number with a value greater than one ('1') and less than
// or equal to +2,147,483,647. As a practical matter, your computer may run out of memory
// before you reach this maximum limit. If 'nthRoot' fails to meet this criteria an error
// will be returned.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an instance to Type BigIntNum .
//
func (nthrt *BigIntMathNthRoot) GetNthRoot(
					radicand, nthRoot BigIntNum,
						maxPrecision uint) (BigIntNum, error) {


	ePrefix := "BigIntMathNthRoot.GetNthRoot() "

	result, err := nthrt.calcNthRootGateway(radicand, nthRoot, maxPrecision)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error returned by nthrt.calcNthRootGateway(radicand, nthRoot, maxPrecision). " +
				"radicand='%v' nthRoot='%v' maxPrecision='%v' ",
					radicand.GetNumStr(), nthRoot.GetNumStr(), maxPrecision)
	}

	return result, nil
}


// calcNthRootGateway - Calculate the nth root of a radicand. The result is returned
// as a BigIntNum.
//
// Input Parameters
// ================
//
// radicand - BigIntNum 	A 'radicand' is the number under a radical symbol (√)
//
// nthRoot 	- BigIntNum 	n√ The index or root. This must be a positive integer value.
//
//maxPrecision	- uint		The maximum precision specified for the result.
//
//
func (nthrt *BigIntMathNthRoot) calcNthRootGateway(
							radicand, nthRoot BigIntNum,
									maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntMathNthRoot.calcNthRootGateway() "

	if nthRoot.GetSign() < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Cannot process a negative Nth Root Value. " +
				"Input parameter, nthRoot='%v' ",
					nthRoot.GetNumStr())
	}

	// If the radicand is zero, the result will always be zero
	if radicand.IsZero() {
		return radicand, nil
	}

	// If nthRoot is zero, the result will always be '1'
	if nthRoot.IsZero() {
		return BigIntNum{}.NewOne(maxPrecision), nil
	}

	if nthRoot.precision > 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "-Error: input parameter 'nthRoot' is a floating point number. " +
			"Precision is greater than Zero. NthRoot calculations can only be performed when 'nthRoot' is an " +
			"integer number.")
	}

	bigINumOne := BigIntNum{}.NewOne(0)

	if nthRoot.CmpBigInt(bigINumOne) == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix +
				"Error - Input Parameter 'nthRoot' INVALID! 'nthRoot' cannot equal 1. ")
	}

	if radicand.GetSign() == -1 {

		bigINumTwo := BigIntNum{}.NewTwo(0)

		_, mod, err := BigIntMathDivide{}.BigIntNumQuotientMod(nthRoot, bigINumTwo, 0)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntMathDivide{}.BigIntNumQuotientMod(nthRoot, bigINumTwo, 0) " +
					"Error='%v' ", err.Error())
		}

		if mod.IsZero() {
			return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error - Cannot compute nthRoot of a negative number " +
				"when nthRoot is even. nthRoot can only be extracted from negative numbers when nthRoot" +
				" is odd. Original Number= %v  nthRoot= %v", radicand.GetNumStr(), nthRoot.GetNumStr())
		}

	}

	err := nthrt.initialize(radicand, nthRoot, maxPrecision)

	if err != nil {
		return BigIntNum{}.NewZero(0),
		fmt.Errorf(ePrefix + "-Error returned from initialization. Error= %v", err.Error())
	}

	err = nthrt.doRootExtraction()

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "-Error returned from nthrt.doRootExtraction() - %v", err.Error())
	}

	return nthrt.ResultBINum, nil
}


// initialize - Initializes the data fields of the BigIntMathNthRoot structure and validates the
// original number passed to the Nth Root Calculation. This method assumes that radicand and
// nthRoot have already been validated.
//
func (nthrt *BigIntMathNthRoot) initialize(
				radicand, nthRoot BigIntNum,
						maxPrecision uint) error {

  ePrefix := "BigIntMathNthRoot.initialize() "
	var err error

	nthrt.OriginalRadicand = radicand.CopyOut()
	nthrt.NthRoot = nthRoot.CopyOut()
	nthrt.RequestedPrecision = maxPrecision

	nthrt.BundleRadicand, err =
		nthrt.setupBundles(nthrt.OriginalRadicand.GetAbsoluteBigIntValue(),
													nthrt.NthRoot.GetAbsoluteBigIntValue())

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from nthrt.setupBundles(radicand.bigInt, " +
			"nthRoot.bigInt). Error= %v", err)
	}

	nthrt.BundleAddOnPrecision =
		big.NewInt(int64(maxPrecision + nthrt.OriginalRadicand.GetPrecisionUint() + uint(5)))

	nthrt.BundleLength, err = nthrt.caclBundleLength(
							nthrt.BundleRadicand,
								nthrt.NthRoot.GetAbsoluteBigIntValue(),
									nthrt.BundleAddOnPrecision)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned from nthrt.caclBundleLength(...). Error= %v",
				err.Error())
	}

	radicandPrecision := big.NewInt(int64(nthrt.OriginalRadicand.GetPrecisionUint()))

	nthrt.AcutalResultPrecision, err =
		nthrt.calcPrecision(radicandPrecision,
													nthrt.BundleAddOnPrecision,
														nthrt.NthRoot.GetAbsoluteBigIntValue())

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from nthrt.calcPrecision(). Error= %v", err)
	}

	// Set constants for calculations

	nthrt.BigOne = big.NewInt(1)
	nthrt.Big10 = big.NewInt(10)
	nthrt.Big10ToNthPower = big.NewInt(0).Exp(
														nthrt.Big10,
															nthrt.NthRoot.GetAbsoluteBigIntValue(),
																nil)
	nthrt.BigZero = big.NewInt(0)
	nthrt.ResultBInt = big.NewInt(0)
	nthrt.ResultBINum = BigIntNum{}.NewZero(0)
	nthrt.Y = big.NewInt(0)
	nthrt.YPrime = big.NewInt(0)
	nthrt.R = big.NewInt(0)
	nthrt.RPrime = big.NewInt(0)
	nthrt.BaseNum = big.NewInt(10)
	nthrt.Alpha = big.NewInt(0)
	nthrt.Beta = big.NewInt(0)
	return nil

}

// setupBundles - The purpose of this method is to modify and format the radicand for separation
// into bundles of integers for computation of the Nth Root.
//
func (nthrt *BigIntMathNthRoot) setupBundles(
			radicand, nthRoot *big.Int) (formattedRadicand *big.Int, err error) {

	// FormatTargetInt
	ePrefix := "BigIntMathNthRoot.setupBundles() "
	formattedRadicand = big.NewInt(0)
	err = nil
	magnitude, errx := BigIntMath{}.GetMagnitude(radicand)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMath{}.GetMagnitude(target). " +
			"target='%v' Error='%v' ",
			radicand.Text(10), err.Error())
		return formattedRadicand, err
	}

	baseTen := big.NewInt(10)
	formattedRadicand = big.NewInt(0).Set(radicand)
	numOfDigits := big.NewInt(0).Add(magnitude, big.NewInt(1))

	// If necessary, add trailing zeros to the radicand to match Nth Root
	// requirements.
	for numOfDigits.Cmp(nthRoot) == -1 {

		formattedRadicand = big.NewInt(0).Mul(formattedRadicand, baseTen )
		numOfDigits = big.NewInt(0).Add(numOfDigits, big.NewInt(1))

	}

	return formattedRadicand, nil

}

// caclBundleLength - Calculates the final bundle length. These are bundles
// of integers that are packaged for submission to the Nth Root calculation
// routine.
//
// The final bundle length is equal to the actual number of bundles associated
// with the radicand plus bundles added on for additional precision in the
// final nthRoot result.j
//
func (nthrt  *BigIntMathNthRoot) caclBundleLength(
				bundleRadicand,
					nthRoot,
						bundleAddonPrecision *big.Int) (bundleLength *big.Int, err error) {

	bundleLength = big.NewInt(0)
	err = nil

	ePrefix := "ExampleBundleCount_03() "
	magnitude, errx := BigIntMath{}.GetMagnitude(bundleRadicand)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned by BigIntMath{}.GetMagnitude(target). " +
			"bundleRadicand='%v' Error='%v' ",
			bundleRadicand.Text(10), err.Error())
		return bundleLength, err
	}

	bigOne := big.NewInt(1)

	numOfDigits := big.NewInt(0).Add(magnitude, bigOne)

	quotient, mod := big.NewInt(0).QuoRem(
		numOfDigits,
		nthRoot,
		big.NewInt(0))


	bundleLength = big.NewInt(0).Set(quotient)

	if mod.Cmp(big.NewInt(0)) == 1 {
		bundleLength = big.NewInt(0).Add(bundleLength, bigOne)
	}

	finalBundleLength := big.NewInt(0).Add(bundleLength, bundleAddonPrecision)

	return finalBundleLength, nil

}

// calcPrecision - This method will calculate the actual precision of the
// result output by the Nth Root calculation.  Actual result precision is
// equal to the Bundle Add On Precision + the quotient of radicand precision
// divided by nthRoot.
//			bundleAddOnPrecion + Quotient(radicandPrecision/nthRoot)
//
func (nthrt *BigIntMathNthRoot) calcPrecision(
			radicandPrecision,
				bundleAddOnPrecision,
					nthRoot *big.Int) (actualPrecision *big.Int, err error)  {

	actualPrecision = big.NewInt(0)
	err = nil

	if nthRoot.Cmp(big.NewInt(0)) == 0 {
		ePrefix := "BigIntMathNthRoot.calcPrecision() "
		err = errors.New(ePrefix +
			"Error: nthRoot=0. Division by nthRoot is Division by Zero and will FAIL!")
		return actualPrecision, err

	}

	quotient := big.NewInt(0).Quo(radicandPrecision, nthRoot)

	actualPrecision = big.NewInt(0).Add(radicandPrecision, quotient)
	err = nil

	return actualPrecision, err
}

func (nthrt *BigIntMathNthRoot) doRootExtraction() error {

	ePrefix := "BigIntMathNthRoot) doRootExtraction() "
	nthrt.Y = big.NewInt(0)
	nthrt.Minuend = big.NewInt(0)
	nthrt.Subtrahend = big.NewInt(0)
	nthrt.R = big.NewInt(0)
	bigOne := big.NewInt(1)

	for i := big.NewInt(0); i.Cmp(nthrt.BundleLength)==-1; i = big.NewInt(0).Add(i,bigOne) {

		err:= nthrt.findNextRoot()

		if err != nil {
			return fmt.Errorf(ePrefix +
				"Error returned by nthrt.findNextRoot(i). " +
				"Error='%v' ", err.Error())
		}

	}

	if nthrt.OriginalRadicand.GetSign() < 1 {
		nthrt.ResultBInt = big.NewInt(0).Neg(nthrt.ResultBInt)
	}

	actualPrecision := uint(nthrt.AcutalResultPrecision.Int64())

	nthrt.ResultBINum = BigIntNum{}.NewBigInt(nthrt.ResultBInt, actualPrecision)

	nthrt.ResultBINum.RoundToDecPlace(nthrt.RequestedPrecision)

	return nil
}

func (nthrt *BigIntMathNthRoot) findNextRoot() error {

	ePrefix := "BigIntMathNthRoot.findNextRoot() "
	bundle, err := nthrt.getNextBundleBigIntValue()

	return fmt.Errorf(ePrefix + "Error returned by nthrt.getNextBundleBigIntValue(bundleIdx). " +
		"Error='%v' ", err.Error())

	// alpha = next n-digits of radicand
	nthrt.Alpha = big.NewInt(0).Set(bundle)

	nthrt.RPrime = big.NewInt(-1)

	itatr := big.NewInt(9)
	term_1a := big.NewInt(0)
	term_1b := big.NewInt(0)

	term_2a1 := big.NewInt(0)
	term_2a2 := big.NewInt(0)
	term_2a := big.NewInt(0)

	term_2b := big.NewInt(0)
	term_2b1 := big.NewInt(0)
	term_2b2 := big.NewInt(0)

	term_1a = big.NewInt(0).Mul(nthrt.Big10ToNthPower, nthrt.R)
	term_1b = big.NewInt(0).Set(nthrt.Alpha)
	nthrt.Minuend = big.NewInt(0).Add(term_1a, term_1b)

	for itatr.Cmp(nthrt.BigZero) > -1 &&
		nthrt.RPrime.Cmp(nthrt.BigZero) == -1 {

		nthrt.Beta = big.NewInt(0).Set(itatr)
		nthrt.YPrime = big.NewInt(0).Mul(nthrt.Y, nthrt.Big10)
		nthrt.YPrime = big.NewInt(0).Add(nthrt.YPrime, nthrt.Beta)

		term_2a1 = big.NewInt(0).Mul(nthrt.BaseNum, nthrt.Y)
		term_2a2 = big.NewInt(0).Add(term_2a1, nthrt.Beta)

		term_2a = big.NewInt(0).Exp(term_2a2,
																big.NewInt(0).Set(nthrt.NthRoot.GetAbsoluteBigIntValue()),
																nil)

		term_2b1 = big.NewInt(0).Set(nthrt.Big10ToNthPower)

		term_2b2 = big.NewInt(0).Exp(nthrt.Y,
								big.NewInt(0).Set(nthrt.NthRoot.GetAbsoluteBigIntValue()),
								nil)

		term_2b = big.NewInt(0).Mul(term_2b1, term_2b2)

		nthrt.Subtrahend = big.NewInt(0).Sub(term_2a, term_2b)

		nthrt.RPrime = big.NewInt(0).Sub(nthrt.Minuend, nthrt.Subtrahend)

		itatr = big.NewInt(0).Sub(itatr, nthrt.BigOne)
	}

	nthrt.R = big.NewInt(0).Set(nthrt.RPrime)
	nthrt.Y = big.NewInt(0).Set(nthrt.YPrime)
	nthrt.ResultBInt = big.NewInt(0).Mul(nthrt.ResultBInt, nthrt.Big10)
	nthrt.ResultBInt = big.NewInt(0).Add(nthrt.ResultBInt, nthrt.Beta )

	return nil
}


func (nthrt *BigIntMathNthRoot) getNextBundleBigIntValue() (*big.Int, error) {

	ePrefix := "BigIntMathNthRoot) getNextBundleBigIntValue() "

	if nthrt.BundleRadicand.Cmp(nthrt.BigZero) == 0 {
		return big.NewInt(0), nil
	}

	radicand := big.NewInt(0).Set(nthrt.BundleRadicand)

	magnitude, err := BigIntMath{}.GetMagnitude(radicand)

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntMath{}.GetMagnitude(radicand). " +
				"radicand='%v' Error='%v' ",
						radicand.Text(10), err.Error())

	}

	numOfDigits := big.NewInt(0).Add(magnitude, big.NewInt(1))

	digitsQuotient := big.NewInt(0).Quo(numOfDigits, nthrt.NthRoot.GetAbsoluteBigIntValue())

	divisor := big.NewInt(0).Exp(nthrt.Big10, digitsQuotient, nil)

	modX := big.NewInt(0)

	nextBundleVal, mod := big.NewInt(0).QuoRem(nthrt.BundleRadicand, divisor, modX)

	nthrt.BundleRadicand  = big.NewInt(0).Set(mod)

	return nextBundleVal, nil
}

