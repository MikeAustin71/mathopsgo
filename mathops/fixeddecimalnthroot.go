package mathops

import (
	"errors"
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
func (fdNRt FixedDecimalNthRoot) FormatFractionalIntegerFromRadicand(
	fracInteger,
	nthRoot *big.Int) (formattedFracInteger, fracTotalDigits *big.Int, err error) {

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundleFromRadicand() "

	formattedFracInteger = big.NewInt(0).Set(fracInteger)
	fracTotalDigits = big.NewInt(0)
	err = nil
	var errX error

	fracTotalDigits, errX = BigIntMath{}.GetMagnitude(fracInteger)

	if errX != nil {
		err = fmt.Errorf(ePrefix +
			"Error returned from %v", errX.Error())
		formattedFracInteger.Set(big.NewInt(0))
		return formattedFracInteger, fracTotalDigits, err
	}

	fracTotalDigits.Add(fracTotalDigits, big.NewInt(1))
	scratch := big.NewInt(0)

	_, remainingDigits := big.NewInt(0).QuoRem(fracTotalDigits, nthRoot, scratch)

	if remainingDigits.Cmp(big.NewInt(0)) == 1 {
		remainingDigits.Sub(nthRoot, remainingDigits)
		scale := big.NewInt(0).Exp(big.NewInt(10), remainingDigits, nil)

		formattedFracInteger.Mul(formattedFracInteger,scale)
		fracTotalDigits.Add(fracTotalDigits, remainingDigits)

	}


	return formattedFracInteger, fracTotalDigits, err
}

// GetNextIntegerBundleFromRadicand - Returns the next bundle for an integer value.
func (fdNRt FixedDecimalNthRoot) GetNextIntegerBundleFromRadicand(
	integer,
	intTotalDigits,
	nthRoot *big.Int) (nextBundle,
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

	if nthRoot.Cmp(bigZero) == 0 {
		err = errors.New(ePrefix + "Error: Input Parameter nthRoot = ZERO!")
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
