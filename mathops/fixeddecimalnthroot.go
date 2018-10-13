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
}

// FormatFractionalDigitsFromRadicand - Performs two tasks:
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
// ** IMPORTANT **
// This function must be called prior to calling GetNextFractionalBundleFromRadicand()
//
func (fdNthRoot *FixedDecimalNthRoot) FormatFractionalDigitsFromRadicand(
	fracRadicand *big.Int,
	fracRadicandPrecision uint,
	nthRoot *big.Int) (formattedFracInteger, fracPrecision *big.Int, err error) {

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundleFromRadicand() "

	formattedFracInteger = big.NewInt(0)
	fracPrecision = big.NewInt(0)
	err = nil
	fdNthRoot.fracMask1 = big.NewInt(0)
	fdNthRoot.fracMask2 = big.NewInt(0)

	if nthRoot.Cmp(big.NewInt(2)) == -1 {

		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'nthRoot' is less than Two!. "+
			"nthRoot='%v' ", nthRoot.Text(10))

		return formattedFracInteger, fracPrecision, err
	}

	bigZero := big.NewInt(0)

	cmpFracRadicandZero := fracRadicand.Cmp(bigZero)

	if cmpFracRadicandZero == -1 {

		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'fracRadicand' is less than ZERO!. "+
			"fracRadicand='%v' ", fracRadicand.Text(10))

		return formattedFracInteger, fracPrecision, err
	}

	if cmpFracRadicandZero == 0 {

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

	// Only fracMask1 and fracMask2 are
	// calculation constants
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
//                residualFracNum = 1.456
//                residualFracPrecision = 3
//
// Input Parameters
// ================
//
// fmtFracNum					*big.Int	- The Radicand fractional digits formatted with a
//                                  leading integer, '11'. Example: radicand fractional
//                                  digits, '123456' with a precision of '9', must be
//                                  formatted as '11.000123456'. If 'fmtFracNum'
//                                  is less than '11', an error will be returned.
//                                  A value of '11' signals that all fractional digits
//                                  have been processed and the next bundle will be
//                                  set to zero.
//
// fracPrecision				*big.Int	- The number of digits to the right of the decimal
//                                  place in fmtFracNum. If fracPrecision is
//                                  NOT evenly divisible by nthRoot, an error will
//                                  be returned.
//
// nthRoot							*big.Int	- The nthRoot for which this root will be calculated.
//																	'nthRoot' must be greater than or equal to two.
//
// Return Values
// =============
//
// nextBundle						*big.Int	- The next bundle of digits to be processed for the
//                                  nthRoot calculation.
//
// residualFracNum			*big.Int	- The remaining fractional digits to be processed
//
// residualFracPrecision *big.Int	- The precision specification for 'residualFracNum.
//
func (fdNthRoot FixedDecimalNthRoot) GetNextFractionalBundleFromRadicand(
	fmtFracNum,
	fracPrecision,
	nthRoot *big.Int) (
	nextBundle,
	residualFracNum,
	residualFracPrecision *big.Int,
	err error) {

	nextBundle = big.NewInt(0)
	residualFracNum = big.NewInt(0)
	residualFracPrecision = big.NewInt(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.GetNextFractionalBundleFromRadicand() "

	if nthRoot.Cmp(big.NewInt(2)) == -1 {
		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'nthRoot' is less than 2!! "+
			"nthRoot='%v' ", nthRoot.Text(10))

		return nextBundle, residualFracNum, residualFracPrecision, err
	}

	bigZero := big.NewInt(0)
	bigEleven := big.NewInt(0)

	if fracPrecision.Cmp(bigZero) == 0 ||
		fmtFracNum.Cmp(bigEleven) == 0 {
		// If fracPrecision is zero, it signals there are no more digits
		// left to process. fmtFracNum == 11 also signals
		// that there are no fractional digits to process. In either
		// of these cases, nextBundle == Zero
		err = nil
		return nextBundle, residualFracNum, residualFracPrecision, err
	}

	fracPrecisionRemainder := big.NewInt(0).Rem(fracPrecision, nthRoot)

	if fracPrecisionRemainder.Cmp(bigZero) == 1 {
		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'fracPrecision' is NOT evely divisible by 'nthRoot'! "+
			"fracPrecision='%v' nthRoot='%v' ", fracPrecision.Text(10), nthRoot.Text(10))

		return nextBundle, residualFracNum, residualFracPrecision, err
	}

	bigTen := big.NewInt(10)
	bigOne := big.NewInt(1)

	// fdNthRoot.fracMask1 and fdNthRoot.fracMask2 are static and
	// were previously calculated in the call to FormatFractionalDigitsFromRadicand()

	fracMask3 := big.NewInt(0).Exp(
		bigTen,
		big.NewInt(0).Sub(
			fracPrecision,
			nthRoot), nil)

	fracMask4 := big.NewInt(0).Mul(bigEleven, fracMask3)

	fracMask5 := big.NewInt(0).Exp(
		bigTen,
		big.NewInt(0).Sub(
			fracPrecision,
			bigOne),
		nil)

	adjustedBundle := big.NewInt(0).Quo(fmtFracNum, fracMask3)
	nextBundle = big.NewInt(0).Sub(adjustedBundle, fdNthRoot.fracMask1)
	adjustedBundle2 := big.NewInt(0).Sub(adjustedBundle, fdNthRoot.fracMask2)
	adjustedBundle3 := big.NewInt(0).Mul(adjustedBundle2, fracMask3)
	adjustedFracNum1 := big.NewInt(0).Sub(fmtFracNum, adjustedBundle3)
	adjustedFracNum2 := big.NewInt(0).Add(fracMask4, adjustedFracNum1)
	residualFracNum = big.NewInt(0).Sub(adjustedFracNum2, fracMask5)
	residualFracPrecision = big.NewInt(0).Sub(fracPrecision, nthRoot)
	err = nil

	return nextBundle, residualFracNum, residualFracPrecision, err
}

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
