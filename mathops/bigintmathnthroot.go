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
	NthRoot            	BigIntNum
	OriginalNum        	BigIntNum
	NthRootIntVal			 	int
	BaseNumBundles     	[][]int
	LenBaseNumBundles  	int
	BaseNumBundlesIdx  	int
	ResultBInt					*big.Int
	ResultPrecision    	int
	ResultBINum					BigIntNum
	RequestedPrecision 	int
	BigOne             	*big.Int
	Big10              	*big.Int
	Big10ToNthPower    	*big.Int
	BigZero            	*big.Int
	Big3               	*big.Int
	Y                  	*big.Int // Root Extracted thusfar
	YPrime             	*big.Int // Next Value of Y
	Minuend            	*big.Int
	Subtrahend         	*big.Int
	R                  	*big.Int // Let R be the remainder
	RPrime             	*big.Int // Let RPrime be the new value of r for next iteration
	BaseNum            	*big.Int // Base Number System - always 10
	Alpha              	*big.Int // Next n-digits of the radicand
	Beta               	*big.Int // Next Digit of the root
}

func (nthrt *BigIntMathNthRoot) Empty() {
	nthrt.NthRoot = BigIntNum{}.NewZero(0)
	nthrt.OriginalNum = BigIntNum{}.NewZero(0)
	nthrt.BaseNumBundles = make([][]int, 0, 500)
	nthrt.LenBaseNumBundles = 0
	nthrt.BaseNumBundlesIdx = 0
	nthrt.ResultBInt = big.NewInt(0)
	nthrt.ResultBINum = BigIntNum{}.NewZero(0)
	nthrt.ResultPrecision = 0
	nthrt.RequestedPrecision = 0
	nthrt.Big10 = big.NewInt(0)
	nthrt.Big10ToNthPower = big.NewInt(0)
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
// nthRoot 	- BigIntNum 	n√ The index or root.
//
//maxPrecision	- uint		The maximum precision specified for the result.
//
//
func (nthrt *BigIntMathNthRoot) calcNthRootGateway(
							radicand, nthRoot BigIntNum,
									maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntMathNthRoot.calcNthRootGateway() "

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

	err := nthrt.initializeAndExtract(radicand, nthRoot, maxPrecision)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned from nthrt.initializeAndExtract(num, " +
				"nthRoot, maxPrecision). num='%v' nthRoot='%v'  Error= %v",
				radicand.GetNumStr(), nthRoot.GetNumStr(), err.Error() )
	}

	return nthrt.ResultBINum, nil
}


// initializeAndExtract - Validates input parameters and sets up structures used
// to complete the nthRoot calculation.
func (nthrt *BigIntMathNthRoot) initializeAndExtract(
					num , nthRoot BigIntNum, maxPrecision uint) error {

	ePrefix := "BigIntMathNthRoot.initializeAndExtract() "

	err := nthrt.initialize(num, nthRoot, maxPrecision)

	if err != nil {
		return fmt.Errorf(ePrefix + "-Error returned from initialization. Error= %v", err.Error())
	}

	err = nthrt.doRootExtraction()

	if err != nil {
		return fmt.Errorf(ePrefix + "-Error returned from nthrt.doRootExtraction() - %v", err.Error())
	}

	return nil

}

// initialize - Initializes the data fields of the BigIntMathNthRoot structure and validates the
// original number passed to the Nth Root Calculation. This method assumes that radicand and
// nthRoot have already been validated.
//
func (nthrt *BigIntMathNthRoot) initialize(
				radicand, nthRoot BigIntNum,
						maxPrecision uint) error {

  ePrefix := "BigIntMathNthRoot.initialize() "


	intNthRootVal, err := nthRoot.GetInt()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error Returned from nthrt.NthRoot.GetInt() - Error= %v", err)
	}

	if intNthRootVal == 1 {
		return fmt.Errorf(ePrefix +
			"Error - Input Parameter 'nthRoot' INVALID! 'nthRoot' cannot equal 1. " +
			"nthRoot= %v", intNthRootVal)

	}

	if radicand.GetSign() == -1 {

 	  mod := intNthRootVal / 2

		isNthRootEven := false

		if intNthRootVal == mod * 2 {
			isNthRootEven = true
		}

		if isNthRootEven {
			return fmt.Errorf(ePrefix + "Error - Cannot compute nthRoot of a negative number " +
				"when nthRoot is even. nthRoot can only be extracted from negative numbers when nthRoot" +
				" is odd. Original Number= %v  nthRoot= %v", radicand.GetNumStr(), nthRoot.GetNumStr())
		}

	}

	nthrt.NthRoot = nthRoot.CopyOut()
	nthrt.OriginalNum = radicand.CopyOut()
	nthrt.NthRootIntVal = intNthRootVal
	nthrt.RequestedPrecision = int(maxPrecision)

	err = nthrt.bundleInts()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from nthrt.bundleInts(). Error= %v", err)
	}

	err = nthrt.bundleFracs()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned from nthrt.bundleFracs(). Error= %v", err)
	}

	err = nthrt.calcPrecision()

	if err != nil {
		return fmt.Errorf("BigIntMathNthRoot.initialize() - Error returned from nthrt.calcPrecision(). Error= %v", err)
	}

	// Set constants for calculations

	nthrt.BigOne = big.NewInt(1)
	nthrt.Big3 = big.NewInt(3)
	nthrt.Big10 = big.NewInt(10)
	nthrt.Big10ToNthPower = big.NewInt(0).Exp(nthrt.Big10, nthrt.NthRoot.bigInt, nil)
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

// bundleInts - computes the bundle size and creates the
// bundle array elements necessary to process the integer
// digits in the original number.
func (nthrt *BigIntMathNthRoot) bundleInts() error {

	ePrefix := "BigIntMathNthRoot.bundleInts() "

	intNums := nthrt.OriginalNum.GetIntegerPart()
	//intNums, err := nthrt.OriginalNum.GetIntegerDigits()

	intNumOfDigits := intNums.GetNumberOfDigits()

	if intNumOfDigits < 1 {
		return errors.New(ePrefix + " intNums Number of Digits is less than 1")
	}

	bundleSize := 0
	bIntNthRtValue, err := nthrt.NthRoot.GetInt()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by nthrt.NthRoot.GetInt(). Error='%v' ", err.Error())
	}

	if intNumOfDigits <= bIntNthRtValue {

		bundleSize = 1

	} else {

		bundleSize = intNumOfDigits / bIntNthRtValue

		if intNumOfDigits > ((intNumOfDigits / bIntNthRtValue) * bIntNthRtValue) {

			bundleSize++

		}

	}

	nthrt.BaseNumBundles = make([][]int, bundleSize)

	bundleIdx := bundleSize - 1
	intAryIdx := 0
	intAry, _ := intNums.GetIntAry()
	u8Ary, _ := intAry.GetIntAry()
	for j := intNumOfDigits - 1; j >= 0; j -= bIntNthRtValue {
		bundle := make([]int, bIntNthRtValue)
		intAryIdx = j
		for k := bIntNthRtValue - 1; k >= 0; k-- {
			if intAryIdx < 0 {
				break
			}

			bundle[k] = int(u8Ary[intAryIdx])
			intAryIdx--
		}

		nthrt.BaseNumBundles[bundleIdx] = append(nthrt.BaseNumBundles[bundleIdx], bundle...)

		bundleIdx--
	}

	return nil
}

// bundleFracs - computes the numeric bundle size and creates
// the bundle array elements required to process the fractional
// digits in the original number.
func (nthrt *BigIntMathNthRoot) bundleFracs() error {

	ePrefix := "BigIntMathNthRoot.bundleFracs() "

	if nthrt.OriginalNum.GetPrecision() < 1 {
		return nil
	}

	fracNums := nthrt.OriginalNum.GetFractionalPart()

	intNthRootVal, err := nthrt.NthRoot.GetInt()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error Returned from nthrt.NthRoot.GetInt() - Error= %v", err)
	}

	iFAry, _ := fracNums.GetIntAry()

	iFAryStats := iFAry.GetIntAryStats()

	u8Ary, _ := iFAry.GetIntAry()

	for i := 1; i < iFAryStats.IntAryLen; i += intNthRootVal {

		bundle := make([]int, intNthRootVal)

		for j := 0; j < intNthRootVal; j++ {

			if i+j < iFAryStats.IntAryLen {
				bundle[j] = int(u8Ary[i+j])
			}
		}

		nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)

	}

	return nil
}

// calcPrecision - calculates the bundle size and creates
// the bundle array elements necessary to process the nth
// root to the requested number of decimal places to the
// right of the decimal point.
func (nthrt *BigIntMathNthRoot) calcPrecision() error {

	ePrefix := "BigIntMathNthRoot.calcPrecision() "

	originalNumPrecision := nthrt.OriginalNum.GetPrecision()

	if originalNumPrecision < 0 {
		return fmt.Errorf(ePrefix +
			"- Existing precision is less than zero! OriginalNum.precision= %v",
			originalNumPrecision)
	}

	intNthRootVal, err := nthrt.NthRoot.GetInt()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error Returned from nthrt.NthRoot.GetInt() - Error= %v", err)
	}


	existingPrecision := originalNumPrecision / intNthRootVal

	if originalNumPrecision != existingPrecision*intNthRootVal {
		existingPrecision++
	}

	bundle := make([]int, intNthRootVal)

	if nthrt.RequestedPrecision <= existingPrecision {

		nthrt.RequestedPrecision = existingPrecision

	}

	nthrt.RequestedPrecision++

	deltaPrecision := nthrt.RequestedPrecision - existingPrecision

	for i := 0; i < deltaPrecision; i++ {
		nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)
	}

	nthrt.LenBaseNumBundles = len(nthrt.BaseNumBundles)

	nthrt.ResultPrecision = nthrt.RequestedPrecision

	return nil
}

func (nthrt *BigIntMathNthRoot) doRootExtraction() error {

	// ePrefix := "BigIntMathNthRoot) doRootExtraction() "
	nthrt.Y = big.NewInt(0)
	nthrt.Minuend = big.NewInt(0)
	nthrt.Subtrahend = big.NewInt(0)
	nthrt.R = big.NewInt(0)

	for i := 0; i < nthrt.LenBaseNumBundles; i++ {

		nthrt.findNextRoot(i)

	}

	if nthrt.OriginalNum.GetSign() < 1 {
		nthrt.ResultBInt = big.NewInt(0).Neg(nthrt.ResultBInt)
	}

	nthrt.ResultBINum = BigIntNum{}.NewBigInt(nthrt.ResultBInt, uint(nthrt.RequestedPrecision))

	nthrt.ResultBINum.RoundToDecPlace( uint(nthrt.RequestedPrecision)-1)

	return nil
}

func (nthrt *BigIntMathNthRoot) findNextRoot(bundleIdx int) {

	bundle := nthrt.getBundleBigInt(bundleIdx)

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

		term_2a = big.NewInt(0).Exp(term_2a2, big.NewInt(int64(nthrt.NthRootIntVal)), nil)

		term_2b1 = big.NewInt(0).Set(nthrt.Big10ToNthPower)

		term_2b2 = big.NewInt(0).Exp(nthrt.Y,big.NewInt(int64(nthrt.NthRootIntVal)), nil)

		term_2b = big.NewInt(0).Mul(term_2b1, term_2b2)

		nthrt.Subtrahend = big.NewInt(0).Sub(term_2a, term_2b)

		nthrt.RPrime = big.NewInt(0).Sub(nthrt.Minuend, nthrt.Subtrahend)

		itatr = big.NewInt(0).Sub(itatr, nthrt.BigOne)
	}

	nthrt.R = big.NewInt(0).Set(nthrt.RPrime)
	nthrt.Y = big.NewInt(0).Set(nthrt.YPrime)
	nthrt.ResultBInt = big.NewInt(0).Mul(nthrt.ResultBInt, nthrt.Big10)
	nthrt.ResultBInt = big.NewInt(0).Add(nthrt.ResultBInt, nthrt.Beta )

}


func (nthrt *BigIntMathNthRoot) getBundleBigInt(idx int) *big.Int {

	bigBundleVal := big.NewInt(0)

	for i := 0; i < nthrt.NthRootIntVal; i++ {

		bigBundleVal = big.NewInt(0).Mul(bigBundleVal, nthrt.Big10)
		bigBundleVal = big.NewInt(0).Add(bigBundleVal, big.NewInt(0).SetInt64(int64(nthrt.BaseNumBundles[idx][i])))

	}

	return bigBundleVal
}

