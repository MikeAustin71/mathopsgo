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
	NthRoot               BigIntNum
	OriginalRadicand      BigIntNum
	SetupRadicand					BigIntNum
	IntBundleRadicand     BigIntNum
	FracBundleRadicand		BigIntNum
	BundleAddOnPrecision  *big.Int
	BundleLength          *big.Int
	ResultBInt            *big.Int
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
	nthrt.IntBundleRadicand = BigIntNum{}.NewZero(0)
	nthrt.FracBundleRadicand = BigIntNum{}.NewZero(0)
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
func (nthrt BigIntMathNthRoot) GetNthRoot(
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

	_, err := nthrt.initialize(radicand, nthRoot, maxPrecision)

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
						maxPrecision uint) (setupRadicand BigIntNum, err error) {

  ePrefix := "BigIntMathNthRoot.initialize() "
	var errx error
	precisionAdjustment := big.NewInt(0)
	nthrt.OriginalRadicand = radicand.CopyOut()
	nthrt.NthRoot = nthRoot.CopyOut()
	nthrt.RequestedPrecision = maxPrecision

	nthrt.SetupRadicand, nthrt.IntBundleRadicand, nthrt.FracBundleRadicand, precisionAdjustment, errx =
															nthrt.setupBundles(nthrt.OriginalRadicand,
																										nthrt.NthRoot)
	if errx != nil {
		setupRadicand = BigIntNum{}.NewZero(0)
		err = fmt.Errorf(ePrefix +
			"Error returned from nthrt.setupBundles(OriginalRadicand, NthRoot). " +
			"OriginalRadicand='%v' NthRoot='%v' Error= %v",
			nthrt.OriginalRadicand.GetNumStr(), nthrt.NthRoot.GetNumStr(), errx.Error())
		return setupRadicand, err
	}

	nthrt.BundleAddOnPrecision =
		big.NewInt(int64(maxPrecision +  uint(2)))


	nthrt.BundleLength, errx = nthrt.calcBundleLength(
							nthrt.SetupRadicand,
									nthrt.NthRoot,
										nthrt.BundleAddOnPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned from nthrt.calcBundleLength(SetupRadicand, NthRoot, AddOnPrecision). " +
			"SetupRadicand='%v' NthRoot='%v' AddOnPrecision='%v' Error= %v",
			nthrt.SetupRadicand.GetNumStr(), nthrt.NthRoot.GetNumStr(),
			nthrt.BundleAddOnPrecision.Text(10), errx.Error())
		setupRadicand = BigIntNum{}.NewZero(0)
		return setupRadicand, err
	}


	radicandPrecision := big.NewInt(int64(nthrt.SetupRadicand.GetPrecisionUint()))

	nthrt.AcutalResultPrecision, errx =
		nthrt.calcPrecision(radicandPrecision,
													precisionAdjustment,
														nthrt.BundleAddOnPrecision,
															nthrt.NthRoot.GetAbsoluteBigIntValue())

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned from nthrt.calcPrecision(). Error= %v", errx)
		setupRadicand = BigIntNum{}.NewZero(0)
		return setupRadicand, err
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
	return nthrt.SetupRadicand, nil

}

// Experimental 2
func (nthrt *BigIntMathNthRoot) setupBundles(
	radicand, nthRoot BigIntNum) (setupRadicand ,
																	intBundleRadicand,
																		fracBundleRadicand BigIntNum,
																			precisionAdjustment *big.Int,
	err error) {

	ePrefix := "BigIntMathNthRoot.setupBundles() "

	 
	intBundleRadicand = BigIntNum{}.NewZero(0)
	fracBundleRadicand = BigIntNum{}.NewZero(0)
	precisionAdjustment = big.NewInt(0)
	err = nil

	var errx error
	numOfDigits := big.NewInt(0)
	numOfAllDigits := big.NewInt(0)
	numOfFracDigits := big.NewInt(0)
	modX := big.NewInt(0)
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	bigTen := big.NewInt(10)
	scaleValue := big.NewInt(0)
	nthRootBigInt := nthRoot.GetAbsoluteBigIntValue()
	setupRadicand =
		BigIntNum{}.NewBigInt(
			radicand.GetAbsoluteBigIntValue(),
			radicand.GetPrecisionUint())

	setupRadicand.TrimTrailingFracZeros()
	setupRadicand.SetExpectedToActualNumberOfDigits()

	if setupRadicand.GetPrecisionUint() == 0 {
		intBundleRadicand = setupRadicand.CopyOut()
		// fmt.Println("setupRadicand: ", setupRadicand.GetNumStr())
		// fmt.Println("intBundleRadicand:", intBundleRadicand.GetNumStr())
		// fmt.Println("fracBundleRadicand:", fracBundleRadicand.GetNumStr())
		return setupRadicand, intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	// Precision must be greater than zero
	radicandPrecision := setupRadicand.GetPrecisionBigInt()

	numOfAllDigits, _, errx = setupRadicand.GetActualNumberOfDigits()

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by  setupRadicand.GetActualNumberOfDigits() " +
			"Error='%v'", errx.Error())

		return BigIntNum{}.NewZero(0),  intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	scaleValue = big.NewInt(0).Exp(bigTen, radicandPrecision, nil)

	fracDigits := big.NewInt(0).Rem(setupRadicand.GetAbsoluteBigIntValue(), scaleValue)

	numOfFracDigits, errx = BigIntMath{}.GetMagnitude(fracDigits)

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by  BigIntMath{}.GetMagnitude(" +
			"fracDigits) Error='%v'", errx.Error())
		return BigIntNum{}.NewZero(0), intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	numOfFracDigits = big.NewInt(0).Add(numOfFracDigits, bigOne)

	if radicandPrecision.Cmp(bigZero) == 1 &&
		(numOfAllDigits.Cmp(radicandPrecision) < 1  ||
			numOfFracDigits.Cmp(radicandPrecision) == -1) {

		nthRootMultiple := big.NewInt(0).Quo(radicandPrecision, nthRootBigInt)

		/*
		if numOfFracDigits.Cmp(radicandPrecision) == -1 {
			precisionAdjustment = big.NewInt(0).Add(nthRootMultiple,bigOne)
			precisionAdjustment = big.NewInt(0).Neg(precisionAdjustment)

		}
		*/

		exponent := big.NewInt(0).Mul(nthRootBigInt,nthRootMultiple)
		scaleValue = big.NewInt(0).Exp(bigTen, exponent, nil)
		bINumScaleValue := BigIntNum{}.NewBigInt(scaleValue, 0)
		setupRadicand = BigIntMathMultiply{}.MultiplyBigIntNums(setupRadicand, bINumScaleValue)
		setupRadicand.SetExpectedToActualNumberOfDigits()
		radicandPrecision = setupRadicand.GetPrecisionBigInt()
		//fmt.Println("bINumSetupRadicand: ", setupRadicand.GetNumStr())
		//precisionAdjustment = big.NewInt(0).Neg(nthRootMultiple)
	}

	scaleValue = big.NewInt(0).Exp(bigTen,	radicandPrecision, nil)
	intBundleRadicandBigInt, fracBundleRadicandBigInt :=
		big.NewInt(0).QuoRem(setupRadicand.GetAbsoluteBigIntValue(), scaleValue, modX)

	intBundleRadicand = BigIntNum{}.NewBigInt(intBundleRadicandBigInt, 0)

	fracBundleRadicand = BigIntNum{}.NewBigInt(fracBundleRadicandBigInt, 0)

	if intBundleRadicand.IsZero() &&
		fracBundleRadicand.IsZero() {
		err = errors.New(ePrefix + "Error: Both intBundleRadicand and fracBundleRadicand are ZERO!")
		return BigIntNum{}.NewZero(0), intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	intBundleRadicand.SetExpectedToActualNumberOfDigits()
	fracBundleRadicand.SetExpectedToActualNumberOfDigits()

	if fracBundleRadicand.IsZero() {
		err = nil
		return setupRadicand, intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	numOfDigits, _, errx = fracBundleRadicand.GetActualNumberOfDigits()

	if errx != nil {
		err =	fmt.Errorf(ePrefix +
			"Error returned by fracBundleRadicand.GetActualNumberOfDigits(). " +
			"fracBundleRadicand='%v' Error='%v' ",
			fracBundleRadicand.GetNumStr(), errx.Error())

		intBundleRadicand = BigIntNum{}.NewZero(0)
		fracBundleRadicand = BigIntNum{}.NewZero(0)
		precisionAdjustment = big.NewInt(0)

		return BigIntNum{}.NewZero(0), intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	mod := big.NewInt(0).Rem(numOfDigits,nthRoot.GetAbsoluteBigIntValue() )

	if mod.Cmp(big.NewInt(0))==1 {
		delta := big.NewInt(0).Sub(nthRoot.GetAbsoluteBigIntValue(), mod)
		scaleVal := big.NewInt(0).Exp(big.NewInt(10), delta, nil)
		fracBundleRadicandBigInt = big.NewInt(0).Mul(fracBundleRadicand.GetAbsoluteBigIntValue(),scaleVal)
		fracBundleRadicand = BigIntNum{}.NewBigInt(fracBundleRadicandBigInt, 0)
		fracBundleRadicand.SetExpectedToActualNumberOfDigits()
	}

	err = nil
	return setupRadicand, intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
}

// setupBundles - Experimental
/*
func (nthrt *BigIntMathNthRoot) setupBundles(
							radicand, nthRoot BigIntNum) (intBundleRadicand,
																							fracBundleRadicand,
																								precisionAdjustment *big.Int,
																									err error) {

	ePrefix := "BigIntMathNthRoot.setupBundles() "

	intBundleRadicand = big.NewInt(0)
	fracBundleRadicand = big.NewInt(0)
	precisionAdjustment = big.NewInt(0)
	err = nil

	var errx error
	numOfDigits := big.NewInt(0)
	modX := big.NewInt(0)
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	bigTen := big.NewInt(10)
	scaleValue := big.NewInt(0)
	nthRootBigInt := nthRoot.GetAbsoluteBigIntValue()
	bINumSetUpRadicand :=
			BigIntNum{}.NewBigInt(
										radicand.GetAbsoluteBigIntValue(),
											radicand.GetPrecisionUint())

	radicandPrecision := big.NewInt(int64(bINumSetUpRadicand.GetPrecisionUint()))

	setUpRadicand := bINumSetUpRadicand.GetAbsoluteBigIntValue()

	numOfDigits, errx = BigIntMath{}.GetMagnitude(bINumSetUpRadicand.GetAbsoluteBigIntValue())

	if errx != nil {
		err = fmt.Errorf(ePrefix + "Error returned by  BigIntMath{}.GetMagnitude(" +
			"radicand.GetAbsoluteBigIntValue()) Error='%v'", errx.Error())
	}

	numOfDigits = big.NewInt(0).Add(numOfDigits, bigOne)

	if numOfDigits.Cmp(radicandPrecision) < 1  {

		nthRootMultiple := big.NewInt(0).Quo(radicandPrecision, nthRootBigInt)
		exponent := big.NewInt(0).Mul(nthRootBigInt,nthRootMultiple)
		scaleValue = big.NewInt(0).Exp(bigTen, exponent, nil)
		bINumScaleValue := BigIntNum{}.NewBigInt(scaleValue, 0)
		bINumSetUpRadicand = BigIntMathMultiply{}.MultiplyBigIntNums(bINumSetUpRadicand, bINumScaleValue)
		setUpRadicand = bINumSetUpRadicand.GetAbsoluteBigIntValue()
		radicandPrecision = big.NewInt(int64(bINumSetUpRadicand.GetPrecisionUint()))
		//precisionAdjustment = big.NewInt(0).Neg(nthRootMultiple)
	}

	scaleValue = big.NewInt(0).Exp(bigTen,	radicandPrecision, nil)

	intBundleRadicand, fracBundleRadicand =
			big.NewInt(0).QuoRem(setUpRadicand, scaleValue, modX)

	if intBundleRadicand.Cmp(bigZero) == 0 &&
				fracBundleRadicand.Cmp(bigZero) == 0 {
		err = errors.New(ePrefix + "Error: Both intBundleRadicand and fracBundleRadicand are ZERO!")
		return intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	numOfDigits, errx = BigIntMath{}.GetMagnitude(fracBundleRadicand)

	if errx != nil {
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntMath{}.GetMagnitude(fracBundleRadicand). " +
			"fracBundleRadicand='%v' Error='%v' ",
			fracBundleRadicand.Text(10), errx.Error())

		intBundleRadicand = big.NewInt(0)
		fracBundleRadicand = big.NewInt(0)
		precisionAdjustment = big.NewInt(0)

		return intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
	}

	numOfDigits = big.NewInt(0).Add(numOfDigits, bigOne)

	mod := big.NewInt(0).Rem(numOfDigits,nthRoot.GetAbsoluteBigIntValue() )

	if mod.Cmp(big.NewInt(0))==1 {
		delta := big.NewInt(0).Sub(nthRoot.GetAbsoluteBigIntValue(), mod)
		scaleVal := big.NewInt(0).Exp(big.NewInt(10), delta, nil)
		fracBundleRadicand = big.NewInt(0).Mul(fracBundleRadicand,scaleVal)
	}

	err = nil

	return intBundleRadicand, fracBundleRadicand, precisionAdjustment, err
}
*/

// setupBundles - The purpose of this method is to modify and format the radicand for separation
// into bundles of integers for computation of the Nth Root.
//
/*
func (nthrt *BigIntMathNthRoot) setupBundles(
			radicand, nthRoot BigIntNum) (intBundleRadicand, fracBundleRadicand *big.Int, err error) {

	intBundleRadicand = big.NewInt(0)
	fracBundleRadicand = big.NewInt(0)
	err = nil

	modX := big.NewInt(0)
	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(radicand.GetPrecisionUint())), nil)

	intBundleRadicand, fracBundleRadicand = big.NewInt(0).QuoRem(radicand.GetAbsoluteBigIntValue(), scaleVal, modX)


	magnitude, errx := BigIntMath{}.GetMagnitude(fracBundleRadicand)

	if errx != nil {
		ePrefix := "BigIntMathNthRoot.setupBundles() "
		err =	fmt.Errorf(ePrefix +
			"Error returned by BigIntMath{}.GetMagnitude(fracBundleRadicand). " +
			"fracBundleRadicand='%v' Error='%v' ",
			fracBundleRadicand.Text(10), err.Error())

		intBundleRadicand = big.NewInt(0)
		fracBundleRadicand = big.NewInt(0)

		return intBundleRadicand, fracBundleRadicand, err
	}


	numOfDigits := big.NewInt(0).Add(magnitude, big.NewInt(1))

	mod := big.NewInt(0).Rem(numOfDigits,nthRoot.GetAbsoluteBigIntValue() )


	if mod.Cmp(big.NewInt(0))==1 {
		delta := big.NewInt(0).Sub(nthRoot.GetAbsoluteBigIntValue(), mod)
		scaleVal := big.NewInt(0).Exp(big.NewInt(10), delta, nil)
		fracBundleRadicand = big.NewInt(0).Mul(fracBundleRadicand,scaleVal)
	}

	err = nil

	return intBundleRadicand, fracBundleRadicand, err
}
*/

// calcBundleLength - Calculates the final bundle length. These are bundles
// of integers that are packaged for submission to the Nth Root calculation
// routine.
//
// The final bundle length is equal to the actual number of bundles associated
// with the radicand plus bundles added on for additional precision in the
// final nthRoot result.
//
func (nthrt  *BigIntMathNthRoot) calcBundleLength(
				radicand,
						nthRoot BigIntNum,
							bundleAddonPrecision *big.Int) (bundleLength *big.Int, err error) {

	ePrefix := "BigIntMathNthRoot.calcBundleLength() "
	bundleLength = big.NewInt(0)
	err = nil
	var errx error

	modX := big.NewInt(0)
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)
	numOfDigits := big.NewInt(1)
	intBundleLength := big.NewInt(0)
	intBundleLengthMod := big.NewInt(0)
	fracBundleLength := big.NewInt(0)
	bigIntNthRoot := nthRoot.GetAbsoluteBigIntValue()

	scaleValue := big.NewInt(0).Exp(
									big.NewInt(10),
										big.NewInt(int64(radicand.GetPrecisionUint())),
											nil)

	radicandBigInt := radicand.GetAbsoluteBigIntValue()

	// break radicand into integer digits
	intRadicand :=
		big.NewInt(0).Quo(
			radicandBigInt,
					scaleValue)

	if intRadicand.Cmp(bigZero) == 1 {

		numOfDigits, errx = BigIntMath{}.GetMagnitude(intRadicand)

		if errx != nil {
			err = fmt.Errorf(ePrefix	+
				"Error returned by BigIntMath{}.GetMagnitude(intRadicand). " +
				"intRadicand='%v' Error='%v'",
					intRadicand.Text(10), err.Error())
		}

		// Convert integer radicand to number of digits
		numOfDigits = big.NewInt(0).Add(numOfDigits, bigOne)


		intBundleLength, intBundleLengthMod = big.NewInt(0).QuoRem(
																						numOfDigits,
																							bigIntNthRoot,
																								modX)

		// For the integer bundle calculation add one to bundle length if
		// there are any remaining digits.
		// Example  5 / 3 = Quotient of 1 Mod of 2. Add 1 to bundle size
		if intBundleLengthMod.Cmp(bigZero) == 1 {
			intBundleLength = big.NewInt(0).Add(intBundleLength, bigOne)
		}

	}

	// break radicand into fractional digits
	fracRadicand :=
		big.NewInt(0).Rem(
			radicandBigInt,
			 scaleValue)

	if fracRadicand.Cmp(bigZero) == 1 {

	numOfDigits, errx = BigIntMath{}.GetMagnitude(fracRadicand)

	if errx != nil {
		err = fmt.Errorf(ePrefix	+
			"Error returned by BigIntMath{}.GetMagnitude(fracRadicand). " +
			"fracRadicand='%v' Error='%v'",
			fracRadicand.Text(10), err.Error())
	}


	// Convert the fractional radicand component to number of digits
	numOfDigits = big.NewInt(0).Add(numOfDigits, bigOne)

	// Frac Bundle Length is equal to 'Quotient' of Number of Digits divided
	// by Nth Root. The Modulo of this division is ignored for fractional
	// digits.
	fracBundleLength = big.NewInt(0).Quo(numOfDigits, bigIntNthRoot)

	}

	fmt.Println("        intBundleLength: ", intBundleLength.Text(10))
	fmt.Println("       fracBundleLength: ", fracBundleLength.Text(10))

	bundleLength = big.NewInt(0).Add(intBundleLength, fracBundleLength)

	fmt.Println("      finalBundleLength: ", bundleLength.Text(10))

	bundleLength = big.NewInt(0).Add(bundleLength, bundleAddonPrecision)
	err = nil

	return bundleLength, err
}

// calcPrecision - This method will calculate the actual precision of the
// result output by the Nth Root calculation.  Actual result precision is
// equal to the Bundle Add On Precision + the quotient of radicand precision
// divided by nthRoot.
//			bundleAddOnPrecion + Quotient(radicandPrecision/nthRoot)
//
func (nthrt *BigIntMathNthRoot) calcPrecision(
			radicandPrecision,
			 precisionAdjustment,
					bundleAddOnPrecision,
						nthRoot *big.Int) (actualPrecision *big.Int, err error)  {

	actualPrecision = big.NewInt(0)
	err = nil
	bigZero := big.NewInt(0)

	if nthRoot.Cmp(bigZero) == 0 {
		ePrefix := "BigIntMathNthRoot.calcPrecision() "
		err = errors.New(ePrefix +
			"Error: nthRoot=0. Division by nthRoot is Division by Zero and will FAIL!")
		return actualPrecision, err

	}

	quotient := big.NewInt(0).Quo(radicandPrecision, nthRoot)

	actualPrecision = big.NewInt(0).Add(bundleAddOnPrecision, quotient)
	actualPrecision = big.NewInt(0).Add(actualPrecision, precisionAdjustment)
	if actualPrecision.Cmp(bigZero) < 0 {
		actualPrecision = big.NewInt(0)
	}

	err = nil

	return actualPrecision, err
}

func (nthrt *BigIntMathNthRoot) doRootExtraction() error {

	ePrefix := "BigIntMathNthRoot.doRootExtraction() "
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
	var err error
	var bundle *big.Int

	bundle, nthrt.IntBundleRadicand, nthrt.FracBundleRadicand, err =
						nthrt.getNextBundleBigIntValue(nthrt.IntBundleRadicand,
																						nthrt.FracBundleRadicand,
																							nthrt.NthRoot)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by nthrt.getNextBundleBigIntValue(nthrt.IntBundleRadicand). " +
			"Error='%v' ", err.Error())
	}


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

func (nthrt *BigIntMathNthRoot) getNextBundleBigIntValue(
	intBundleRadicand, fracBundleRadicand, nthRoot BigIntNum) (
	nextBundleValue *big.Int, newIntBundleRadicand, newFracBundleRadicand BigIntNum, err error) {

	ePrefix := "BigIntMathNthRoot.getNextBundleBigIntValue() "
	nextBundleValue = big.NewInt(0)
	newIntBundleRadicand = BigIntNum{}.NewZero(0)
	newIntBundleRadicand.SetExpectedToActualNumberOfDigits()
	newFracBundleRadicand = BigIntNum{}.NewZero(0)
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

	expectedNumberOfDigits := intBundleRadicand.GetExpectedNumberOfDigits()
	actualNumberOfDigits, _, errx  := intBundleRadicand.GetActualNumberOfDigits()

	if expectedNumberOfDigits.Cmp(actualNumberOfDigits) == 1 {
		nextBundleValue = big.NewInt(0)
		newIntBundleRadicand = intBundleRadicand.CopyOut()
		newFracBundleRadicand = fracBundleRadicand.CopyOut()

		expectedNumberOfDigits = big.NewInt(0).Sub(expectedNumberOfDigits, nthRoot.GetAbsoluteBigIntValue() )
		if expectedNumberOfDigits.Cmp(actualNumberOfDigits) == -1 {
			expectedNumberOfDigits = big.NewInt(0).Set(actualNumberOfDigits)
		}

		newIntBundleRadicand.SetExpectedNumberOfDigits(expectedNumberOfDigits)

		err = nil

		//return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err

	} else if !intBundleRadicand.IsZero() {

		magnitude, errx = BigIntMath{}.GetMagnitude(intBundleRadicand.GetAbsoluteBigIntValue())

		if errx != nil {
			err = fmt.Errorf(ePrefix +
				"Error returned by BigIntMath{}.GetMagnitude(intBundleRadicand.GetAbsoluteBigIntValue()). " +
				"intBundleRadicand='%v' Error='%v'",
				intBundleRadicand.GetNumStr(), errx.Error())

			nextBundleValue = big.NewInt(0)
			newIntBundleRadicand = BigIntNum{}.NewZero(0)
			newFracBundleRadicand = BigIntNum{}.NewZero(0)

			return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
		}
		actualNumberOfDigits = big.NewInt(0).Add(magnitude, big.NewInt(1))
		digitsQuotient = big.NewInt(0).Quo(magnitude, nthrt.NthRoot.GetAbsoluteBigIntValue())
		exponent = big.NewInt(0).Mul(digitsQuotient, nthrt.NthRoot.GetAbsoluteBigIntValue())
		divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

		nextBundleValue, tempRadicand =
				big.NewInt(0).QuoRem(intBundleRadicand.GetAbsoluteBigIntValue(), divisor, modX)

		numberOfDigits, errx = BigIntMath{}.GetMagnitude(nextBundleValue)

		if errx != nil {
			err = fmt.Errorf(ePrefix + "Error returned by BigIntMath{}.GetMagnitude(nextBundleValue). " +
				"nextBundleValue='%v' Error='%v'",
					nextBundleValue.Text(10), errx.Error())

			nextBundleValue = big.NewInt(0)
			newIntBundleRadicand = BigIntNum{}.NewZero(0)
			newFracBundleRadicand = BigIntNum{}.NewZero(0)

			return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
		}

		numberOfDigits = big.NewInt(0).Add(numberOfDigits, big.NewInt(1))

		expectedNumberOfDigits = big.NewInt(0).Sub(actualNumberOfDigits, numberOfDigits)

		newIntBundleRadicand = BigIntNum{}.NewBigInt(tempRadicand, 0)
		newIntBundleRadicand.SetExpectedNumberOfDigits(expectedNumberOfDigits)
		newFracBundleRadicand = fracBundleRadicand.CopyOut()
		err = nil

	} else if !fracBundleRadicand.IsZero() {

		magnitude, errx = BigIntMath{}.GetMagnitude(fracBundleRadicand.GetAbsoluteBigIntValue())

		if errx != nil {

			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntMath{}.GetMagnitude(fracBundleRadicand.GetAbsoluteBigIntValue()). " +
				"fracBundleRadicand='%v' Error='%v' ",
				fracBundleRadicand.GetNumStr(), errx.Error())

			nextBundleValue = big.NewInt(0)
			newIntBundleRadicand = BigIntNum{}.NewZero(0)
			newFracBundleRadicand = BigIntNum{}.NewZero(0)
			return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
		}

		digitsQuotient = big.NewInt(0).Quo(magnitude, nthrt.NthRoot.GetAbsoluteBigIntValue())
		exponent = big.NewInt(0).Mul(digitsQuotient, nthrt.NthRoot.GetAbsoluteBigIntValue())
		divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

		nextBundleValue, tempRadicand =
			big.NewInt(0).QuoRem(fracBundleRadicand.GetAbsoluteBigIntValue(), divisor, modX)

		newIntBundleRadicand = BigIntNum{}.NewZero(0)
		newIntBundleRadicand.SetExpectedToActualNumberOfDigits()

		newFracBundleRadicand = BigIntNum{}.NewBigInt(tempRadicand, 0)
		newFracBundleRadicand.SetExpectedToActualNumberOfDigits()

		err = nil

	} else {
		nextBundleValue = big.NewInt(0)
		newIntBundleRadicand = BigIntNum{}.NewZero(0)
		newIntBundleRadicand.SetExpectedToActualNumberOfDigits()
		newFracBundleRadicand = BigIntNum{}.NewZero(0)
		newFracBundleRadicand.SetExpectedToActualNumberOfDigits()
		err = nil
	}


	fmt.Println("--------------------")
	fmt.Println("Original IntBundleRadicand: ",nthrt.IntBundleRadicand.GetNumStr())
	fmt.Println("           nextBundleValue: ", nextBundleValue.Text(10))
	fmt.Println("      newIntBundleRadicand: ", newIntBundleRadicand.GetNumStr())
	fmt.Println("     newFracBundleRadicand: ", newFracBundleRadicand.GetNumStr())
	fmt.Println("")

	return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
}


/*
func (nthrt *BigIntMathNthRoot) getNextBundleBigIntValue(
				intBundleRadicand, fracBundleRadicand *big.Int) (
					nextBundleValue, newIntBundleRadicand, newFracBundleRadicand *big.Int, err error) {

	ePrefix := "BigIntMathNthRoot) getNextBundleBigIntValue() "
	nextBundleValue = big.NewInt(0)
	newIntBundleRadicand = big.NewInt(0)
	newFracBundleRadicand = big.NewInt(0)
	err = nil
	bigZero := big.NewInt(0)
	modX := big.NewInt(0)

	var errx error
	var magnitude *big.Int
	var digitsQuotient *big.Int
	var exponent *big.Int
	var divisor *big.Int

	if intBundleRadicand.Cmp(bigZero) == 1 {

		magnitude, errx = BigIntMath{}.GetMagnitude(intBundleRadicand)

		if errx != nil {

			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntMath{}.GetMagnitude(intBundleRadicand). " +
				"intBundleRadicand='%v' Error='%v' ",
				intBundleRadicand.Text(10), err.Error())

			return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
		}


		digitsQuotient = big.NewInt(0).Quo(magnitude, nthrt.NthRoot.GetAbsoluteBigIntValue())
		exponent = big.NewInt(0).Mul(digitsQuotient, nthrt.NthRoot.GetAbsoluteBigIntValue())
		divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

		nextBundleValue, newIntBundleRadicand = big.NewInt(0).QuoRem(intBundleRadicand, divisor, modX)

		newFracBundleRadicand = big.NewInt(0).Set(fracBundleRadicand)

		err = nil

	} else if fracBundleRadicand.Cmp(bigZero) == 1 {

		magnitude, errx = BigIntMath{}.GetMagnitude(fracBundleRadicand)

		if errx != nil {

			err =	fmt.Errorf(ePrefix +
				"Error returned by BigIntMath{}.GetMagnitude(fracBundleRadicand). " +
				"fracBundleRadicand='%v' Error='%v' ",
				fracBundleRadicand.Text(10), err.Error())

			return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
		}

		digitsQuotient = big.NewInt(0).Quo(magnitude, nthrt.NthRoot.GetAbsoluteBigIntValue())
		exponent = big.NewInt(0).Mul(digitsQuotient, nthrt.NthRoot.GetAbsoluteBigIntValue())
		divisor = big.NewInt(0).Exp(nthrt.Big10, exponent, nil)

		nextBundleValue, newFracBundleRadicand = big.NewInt(0).QuoRem(fracBundleRadicand, divisor, modX)

		newIntBundleRadicand = big.NewInt(0)

		err = nil

	} else {
		nextBundleValue = big.NewInt(0)
		newIntBundleRadicand = big.NewInt(0)
		newFracBundleRadicand = big.NewInt(0)
		err = nil
	}


	fmt.Println("--------------------")
	fmt.Println("nthrt.IntBundleRadicand: ",nthrt.IntBundleRadicand.Text(10))
	fmt.Println("        nextBundleValue: ", nextBundleValue.Text(10))
	fmt.Println("   newIntBundleRadicand: ", newIntBundleRadicand.Text(10))
	fmt.Println("  newFracBundleRadicand: ", newFracBundleRadicand.Text(10))
	fmt.Println("")

	return nextBundleValue, newIntBundleRadicand, newFracBundleRadicand, err
}
*/
