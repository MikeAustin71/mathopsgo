package mathops

import (
	"fmt"
	"math/big"
)

type FixedDecimalNthRoot struct {
	NthRoot  BigIntFixedDecimal
	Radicand BigIntFixedDecimal
	Root     BigIntFixedDecimal
}

// FormatFractionalIntegerFromRadicand - Formats the fractional part
// of the radicand for bundle allocation.
//
// The fractionalRadicand is received as a type *big.Int integer with
// an associated precision specification.
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
//  formattedFracInteger = 1012345600  - Notice the leading integer '1' placeholder
//                                       in the return value. Aso notice that the
//                                       fractional digits have trailing zeros added
//                                       such that the fractional digits are evenly
//                                       divisible by the nthRoot
// fracTotalDigits = 9									 Only the number of fractional digits in the
//                                       'formattedFracInteger' are counted. The
//                                       leading '1' is NOT counted.
//
// Input Parameters
// ================
// fracRadicand				*big.Int	= The fractional numeric digits of the radicand expressed
//                            		as a type *big.Int integer value. fracRadicand must be
//                                greater than or equal to zero
//
// fracRadicandPrecision 	uint	= The precision specification for input parameter 'fracRadicand'.
//																Taken together, 'fracRadicand' and 'fracRadicandPrecision' defined
//                                a fixed length decimal value.
//                                Example: fracRadicand = 123456; fracRadicandPrecision=7 defines
//                                         a value of 0.0123456
//
// nthRoot						*big.Int	= The nthRoot must be defined as positive value greater than or equal
//                                to '2'.
//
func (fdNRt FixedDecimalNthRoot) FormatFractionalIntegerFromRadicand(
	fracRadicand *big.Int,
	fracRadicandPrecision uint,
	nthRoot *big.Int) (formattedFracInteger, fracPrecision *big.Int, err error) {

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundleFromRadicand() "

	bigZero := big.NewInt(0)

	if nthRoot.Cmp(big.NewInt(2)) == -1 {
		formattedFracInteger = big.NewInt(0)
		fracPrecision = big.NewInt(0)

		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'nthRoot' is less than Two!. " +
			"nthRoot='%v' ", nthRoot.Text(10))

		return formattedFracInteger, fracPrecision, err
	}


	if fracRadicand.Cmp(bigZero) == -1 {
		formattedFracInteger = big.NewInt(0)
		fracPrecision = big.NewInt(0)

		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'fracRadicand' is less than ZERO!. " +
			"fracRadicand='%v' ", fracRadicand.Text(10))

		return formattedFracInteger, fracPrecision, err
	}

	formattedFracInteger = big.NewInt(0).Set(fracRadicand)
	fracPrecision = big.NewInt(int64(fracRadicandPrecision))
	err = nil
	bigTen := big.NewInt(10)

	leadingOne := big.NewInt(1)
	scale := big.NewInt(0).Exp(bigTen, fracPrecision, nil )
	leadingOne.Mul(leadingOne, scale)
  formattedFracInteger.Add(formattedFracInteger,leadingOne)
	remainderFracPrecisionModNthRoot := big.NewInt(0).Rem(fracPrecision, nthRoot)
	if remainderFracPrecisionModNthRoot.Cmp(bigZero)== 1 {
		nthRootMinusRmdr := big.NewInt(0).Sub(nthRoot, remainderFracPrecisionModNthRoot)
		scale = big.NewInt(0).Exp(bigTen, nthRootMinusRmdr, nil)
		formattedFracInteger.Mul(formattedFracInteger, scale)
		fracPrecision.Add(fracPrecision, nthRootMinusRmdr)
	}


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
//								fracDigits = 1.123456
//						    fracPrecision = 6
//								nthRoot = 3
//
//    Return Values:
//								nextBundle = 123
//                fracIntResidue = 1.456
//                fracIntResiduePrecision = 3
//
func (fdNRt FixedDecimalNthRoot) GetNextFractionalBundleFromRadicand(
	fracDigits,
	fracPrecision,
	nthRoot *big.Int,) (
		nextBundle,
		fracIntResiduePrecision,
		fracIntResidueTotalDigits *big.Int,
		err error){

	nextBundle = big.NewInt(0)
	fracIntResiduePrecision = big.NewInt(0)
	fracIntResidueTotalDigits = big.NewInt(0)
	err = nil



	return nextBundle, fracIntResiduePrecision, fracIntResidueTotalDigits, err
}

// GetNextIntegerBundleFromRadicand - Returns the next bundle for an integer value.
func (fdNRt FixedDecimalNthRoot) GetNextIntegerBundleFromRadicand(
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
		err = fmt.Errorf(ePrefix + "Error: Input Parameter 'integer' is less than zero! " +
			"integer='%v' ", integer.Text(10))

		return nextBundle, nextBundleTotDigits, integerResidual, intResidualTotalDigits, err
	}

	if intTotalDigits.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix + "Error: Input Parameter 'intTotalDigits' is less than zero! " +
			"intTotalDigits='%v' ", intTotalDigits.Text(10))

		return nextBundle, nextBundleTotDigits, integerResidual, intResidualTotalDigits, err
	}

	if nthRoot.Cmp(big.NewInt(2)) == -1 {
		err = fmt.Errorf(ePrefix + "Error: Input Parameter nthRoot is less than two! " +
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
