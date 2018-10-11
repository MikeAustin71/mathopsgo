package mathops

import (
	"errors"
	"math/big"
)

type FixedDecimalNthRoot struct {
	NthRoot BigIntFixedDecimal
	Radicand BigIntFixedDecimal
	Root     BigIntFixedDecimal
}



// GetNextIntegerBundle - Returns the next bundle for an integer value.
func (fdNRt FixedDecimalNthRoot) GetNextIntegerBundle(
	integer,
	intTotalDigits,
	nthRoot *big.Int) ( nextBundle,
											integerResidual,
											intResidualTotalDigits *big.Int,
											err error) {

	nextBundle = big.NewInt(0)
	integerResidual = big.NewInt(0)
	intResidualTotalDigits = big.NewInt(0)
	err = nil

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundle() "

	bigZero := big.NewInt(0)

	if nthRoot.Cmp(bigZero) == 0 {
		err = errors.New(ePrefix + "Error: Input Parameter nthRoot = ZERO!")
		return nextBundle, integerResidual, intResidualTotalDigits, err
	}

	numOfBundlesInThisInteger := big.NewInt(0)
	remainingDigits := big.NewInt(0)
	scratch := big.NewInt(0)
	factor := big.NewInt(0)

	numOfBundlesInThisInteger, remainingDigits = big.NewInt(0).QuoRem(intTotalDigits, nthRoot, scratch)

	cmpNumOfBundles := numOfBundlesInThisInteger.Cmp(bigZero)
	cmpRemainDigits := remainingDigits.Cmp(bigZero)

	if cmpNumOfBundles == 0 && cmpRemainDigits == 0 {
		return nextBundle, integerResidual, intResidualTotalDigits, err
	}


	if cmpRemainDigits == 1 {
		// We have remaining digits.
		// or an incomplete bundle.
		// The number of digits here
		// is less than nthRoot.

		factor.Set(remainingDigits)

	} else {
		// cmpNumOfBundles MUST BE == 1
		// The number of digits in this
		// bundle will be = nthRoot
		factor.Set(nthRoot)
	}

	delta := intTotalDigits.Sub(intTotalDigits, factor)
	scale := big.NewInt(0).Exp(big.NewInt(10), delta, nil)
	nextBundle,integerResidual = big.NewInt(0).QuoRem(integer,scale, scratch)
	intResidualTotalDigits = big.NewInt(0).Sub(delta, big.NewInt(1))
	err = nil
	return nextBundle, integerResidual, intResidualTotalDigits, err
}