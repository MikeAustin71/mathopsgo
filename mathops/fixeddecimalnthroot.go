package mathops

import (
	"fmt"
	"math/big"
)

type FixedDecimalNthRoot struct {
	NthRoot   BigIntFixedDecimal
	Radicand  BigIntFixedDecimal
	Root      BigIntFixedDecimal
	fracMask1 *big.Int
	fracMask2 *big.Int
	fracMask3 *big.Int
	fracMask4 *big.Int
	fracMask5 *big.Int
}

// FormatFractionalIntegerFromRadicand - Performs two tasks:
// 	(1) Formats the fractional part of the radicand for bundle allocation.
//  (2) Creates calculation constants which will be stored in parent
//      structure for used in method GetNextFractionalBundleFromRadicand()
//
// The fractionalRadicand is received as a type *big.Int integer with
// an associated precision specification. Together these two parameters
// define a decimal fraction such as 0.0123456. This represents the fraction
// part of the radicand. The method then proceeds to reformat this decimal
// fraction as an integer value with a leading '11' value. For example the
// decimal fraction 0.0123456 would be sent to this method as an integer value
// of '123456' with a precision value of '7'. This value would then be reformatted
// as the integer value 11012345600. In addition to the prefix '11', notice that
// two trailing zeros were added. This is due to the requirement that all fractional
// values must be evenly divisible by the nthRoot. In this example, the nthRoot is
// '3'. Upon completion, this method will return a formattedFracInteger of '11012345600'
// and a fracPrecision of '9'.
//
// Example:
// 	decimal fraction:  0.0123456
//
// 	Input Values
//  ------------
// 	fracRadicand = 123456
//  fracRadicandPrecision = 7
//  nthRoot = 3
//
//  Return Values
//  -------------
//  formattedFracInteger = 11012345600 - Notice the leading integer '1' placeholder
//                                       in the return value. Aso notice that the
//                                       fractional digits have trailing zeros added
//                                       such that the fractional digits are evenly
//                                       divisible by the nthRoot.
//
// fracPrecision = 9									 	 Only the number of fractional digits in the
//                                       'formattedFracInteger' are counted. The
//                                       leading '11' integers are NOT counted.
//
// Input Parameters
// ================
// fracRadicand				*big.Int	- The fractional numeric digits of the radicand expressed
//                            		as a type *big.Int integer value. fracRadicand must be
//                                greater than or equal to zero
//
// fracRadicandPrecision 	uint	- The precision specification for input parameter 'fracRadicand'.
//																Taken together, 'fracRadicand' and 'fracRadicandPrecision' defined
//                                a fixed length decimal value.
//                                Example: fracRadicand = 123456; fracRadicandPrecision=7 defines
//                                         a value of 0.0123456
//
// nthRoot						*big.Int	- The nthRoot must be defined as positive value greater than or equal
//                                to '2'.
//
// Return Values
// =============
//
// formattedFracInteger *big.Int	- If this function completes successfully, this return value
//                                  will be populated with the formatted fractional integer value.
//                                  In the example above, this would be '11012345600'.
//
// fracPrecision        *big.Int  - The precision associated with the returned 'formattedFracInteger'.
//
// err									error			- If the function completes successfully, this return value will be
//                                  set to 'nil'. If an error occurs, the returned error instance
//                                  will include an appropriate error message.
//
func (fdNthRoot *FixedDecimalNthRoot) FormatFractionalIntegerFromRadicand(
	fracRadicand *big.Int,
	fracRadicandPrecision uint,
	nthRoot *big.Int) (formattedFracInteger, fracPrecision *big.Int, err error) {

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundleFromRadicand() "

	bigZero := big.NewInt(0)

	formattedFracInteger = big.NewInt(0)
	fracPrecision = big.NewInt(0)
	err = nil

	if nthRoot.Cmp(big.NewInt(2)) == -1 {

		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'nthRoot' is less than Two!. "+
			"nthRoot='%v' ", nthRoot.Text(10))

		return formattedFracInteger, fracPrecision, err
	}

	cmpFracRadicandZero := fracRadicand.Cmp(bigZero)

	if cmpFracRadicandZero == -1 {

		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'fracRadicand' is less than ZERO!. "+
			"fracRadicand='%v' ", fracRadicand.Text(10))

		return formattedFracInteger, fracPrecision, err
	}

	if cmpFracRadicandZero == 0 {
		fdNthRoot.fracMask1 = big.NewInt(0)
		fdNthRoot.fracMask2 = big.NewInt(0)
		fdNthRoot.fracMask3 = big.NewInt(0)
		fdNthRoot.fracMask4 = big.NewInt(0)
		fdNthRoot.fracMask5 = big.NewInt(0)

		// formattedFracInteger == 0; fracPrecision==0; err==nil
		return formattedFracInteger, fracPrecision, err
	}


	bigTen := big.NewInt(10)

	scale := big.NewInt(0).Exp(bigTen, fracPrecision, nil)
	bigEleven := big.NewInt(11)
	leadingOneOne := big.NewInt(0).Mul(bigEleven, scale)
	formattedFracInteger = big.NewInt(0).Add(fracRadicand, leadingOneOne)
	fracPrecision = big.NewInt(int64(fracRadicandPrecision))

	precisionRemainder := big.NewInt(0).Rem(fracPrecision, nthRoot)
	if precisionRemainder.Cmp(bigZero) == 1 {
		nthRootMinusRmdr := big.NewInt(0).Sub(nthRoot, precisionRemainder)
		scale = big.NewInt(0).Exp(bigTen, nthRootMinusRmdr, nil)
		formattedFracInteger.Mul(formattedFracInteger, scale)
		fracPrecision.Add(fracPrecision, nthRootMinusRmdr)
	}

	bigOne := big.NewInt(1)
	fdNthRoot.fracMask1 =
		big.NewInt(0).Mul(
			bigEleven,
			big.NewInt(0).Exp(
				bigTen,
				nthRoot,
				nil))

	fdNthRoot.fracMask2 = big.NewInt(0).Exp(
		bigTen,
		big.NewInt(0).Add(
			nthRoot,
			bigOne), nil)

	fdNthRoot.fracMask3 =
		big.NewInt(0).Exp(
			bigTen,
			big.NewInt(0).Sub(
				fracPrecision,
				nthRoot),
			nil)

	fdNthRoot.fracMask4 =
		big.NewInt(0).Mul(fdNthRoot.fracMask3, bigEleven)

	fdNthRoot.fracMask5 =
		big.NewInt(0).Exp(
			bigTen,
			big.NewInt(0).Add(
				fracPrecision,
				bigOne),
			nil)

	return formattedFracInteger, fracPrecision, err
}

// GetNextFractionalBundleFromRadicand - Receives the fractional digits of a
// radicand and extracts the next value. If all the digits have been allocated
// to bundles and the value of input parameter 'fracInteger' is zero,
// this method returns 'fracIntResidue' as zero.
//
// This method assumes that input parameter 'fracInteger' containing the fractional
// digits of a radicand has been properly formatted. This means that number of
// digits in the 'fracInteger' is evenly divisible by 'nthRoot'. It also assumes
// a placeholder value of '1' as the only digit to the left of the decimal
//
// Example:
//    For actual fractional digits '123456'
//		InputValues =
//								formattedFracDigits = 1.000123456
//						    fracPrecision = 9
//								nthRoot = 3
//
//    Return Values:
//								nextBundle = 123
//                fracIntResidue = 1.456
//                fracIntResiduePrecision = 3
//
// Input Parameters
// ================
//
// formattedFracDigits	*big.Int	- The Radicand fractional digits formatted with a
//                                  leading integer, '1'. Example: radicand fractional
//                                  digits, '123456' with a precision of '9', must be
//                                  formatted as '1.000123456'. If 'formattedFracDigits'
//                                  is less than '1', an error will be returned.
//
// fracPrecision				*big.Int	- The number of digits to the right of the decimal
//                                  place in formattedFracDigits. If fracPrecision is
//                                  NOT evenly divisible by nthRoot, an error will
//                                  be returned.
//
/*
func (fdNthRoot FixedDecimalNthRoot) GetNextFractionalBundleFromRadicand(
	formattedFracDigits,
	fracPrecision,
	nthRoot *big.Int,) (
		nextBundle,
		fracIntResidual,
		fracIntResidualPrecision *big.Int,
		err error){

	nextBundle = big.NewInt(0)
	fracIntResidual = big.NewInt(0)
	fracIntResidualPrecision = big.NewInt(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.GetNextFractionalBundleFromRadicand() "


	if nthRoot.Cmp(big.NewInt(2)) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'nthRoot' is less than 2!! " +
			"nthRoot='%v' ", nthRoot.Text(10))

		return nextBundle, fracIntResidual, fracIntResidualPrecision, err
	}

	bigTen := big.NewInt(10)
	oneToPrecision := big.NewInt(0).Exp(bigTen, fracPrecision, nil)


	if formattedFracDigits.Cmp(oneToPrecision) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'formattedFracDigits' is less than 1!! " +
			"formattedFracDigits='%v' ", formattedFracDigits.Text(10))

		return nextBundle, fracIntResidual, fracIntResidualPrecision, err
	}

	bigZero := big.NewInt(0)

	if fracPrecision.Cmp(bigZero) ==  1 {
		// If fracPrecision is zero, it signals there are no more digits
		// left to process. nextBundle == Zero
		err = nil
		return nextBundle, fracIntResidual, fracIntResidualPrecision, err
	}

	scratchNum := big.NewInt(0)
	quotient,remainder := big.NewInt(0).QuoRem(fracPrecision, nthRoot,scratchNum)

	if remainder.Cmp(bigZero) == 1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'fracPrecision' is NOT evenly divisible by 'nthRoot'!!! " +
			"fracPrecision='%v' nthRoot='%v' ", fracPrecision.Text(10), nthRoot.Text(10))

		return nextBundle, fracIntResidual, fracIntResidualPrecision, err
	}

	nthRootMinusPrecision := big.NewInt(0).Sub(fracPrecision, nthRoot)

	scale := big.NewInt(0).Exp(bigTen, nthRootMinusPrecision, nil)

	nextBundle = big.NewInt(0).Quo(formattedFracDigits, nthRootMinusPrecision)

	return nextBundle, fracIntResidual, fracIntResidualPrecision, err
}
*/

// GetNextIntegerBundleFromRadicand - Returns the next bundle for an integer value.
func (fdNthRoot *FixedDecimalNthRoot) GetNextIntegerBundleFromRadicand(
	integer,
	intTotalDigits,
	nthRoot *big.Int) (
	nextBundle,
	nextBundleTotDigits,
	integerResidual,
	intResidualTotalDigits *big.Int,
	err error) {

	nextBundle = big.NewInt(0)
	nextBundleTotDigits = big.NewInt(0)
	integerResidual = big.NewInt(0)
	intResidualTotalDigits = big.NewInt(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundleFromRadicand() "

	bigZero := big.NewInt(0)

	if integer.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix+"Error: Input Parameter 'integer' is less than zero! "+
			"integer='%v' ", integer.Text(10))

		return nextBundle, nextBundleTotDigits, integerResidual, intResidualTotalDigits, err
	}

	if intTotalDigits.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix+"Error: Input Parameter 'intTotalDigits' is less than zero! "+
			"intTotalDigits='%v' ", intTotalDigits.Text(10))

		return nextBundle, nextBundleTotDigits, integerResidual, intResidualTotalDigits, err
	}

	if nthRoot.Cmp(big.NewInt(2)) == -1 {
		err = fmt.Errorf(ePrefix+"Error: Input Parameter nthRoot is less than two! "+
			"nthRoot='%v' ", nthRoot.Text(10))

		return nextBundle, nextBundleTotDigits, integerResidual, intResidualTotalDigits, err
	}

	numOfBundlesInThisInteger := big.NewInt(0)
	remainingDigits := big.NewInt(0)
	scratch := big.NewInt(0)

	numOfBundlesInThisInteger, remainingDigits = big.NewInt(0).QuoRem(intTotalDigits, nthRoot, scratch)

	cmpNumOfBundles := numOfBundlesInThisInteger.Cmp(bigZero)
	cmpRemainDigits := remainingDigits.Cmp(bigZero)

	if cmpNumOfBundles == 0 && cmpRemainDigits == 0 {
		return nextBundle, nextBundleTotDigits, integerResidual, intResidualTotalDigits, err
	}

	if cmpRemainDigits == 1 {
		// We have remaining digits.
		// or an incomplete bundle.
		// The number of digits here
		// is less than nthRoot.

		nextBundleTotDigits.Set(remainingDigits)

	} else {
		// cmpNumOfBundles MUST BE == 1
		// The number of digits in this
		// bundle will be = nthRoot
		nextBundleTotDigits.Set(nthRoot)
	}

	intResidualTotalDigits.Sub(intTotalDigits, nextBundleTotDigits)
	scale := big.NewInt(0).Exp(big.NewInt(10), intResidualTotalDigits, nil)
	nextBundle, integerResidual = big.NewInt(0).QuoRem(integer, scale, scratch)

	err = nil
	return nextBundle, nextBundleTotDigits, integerResidual, intResidualTotalDigits, err
}
