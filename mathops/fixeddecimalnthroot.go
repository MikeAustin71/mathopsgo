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
	integerMagnitude,
	nthRoot *big.Int,
	intBundleNo int) ( nextBundle,
											integerResidual,
											magnitudeResidual *big.Int,
											thisBundleNo int,
											err error) {

	nextBundle = big.NewInt(0)
	integerResidual = big.NewInt(0)
	magnitudeResidual = big.NewInt(0)
	thisBundleNo = 0
	err = nil

	ePrefix := "FixedDecimalNthRoot.GetNextIntegerBundle() "

	bigZero := big.NewInt(0)

	if nthRoot.Cmp(bigZero) == 0 {
		err = errors.New(ePrefix + "Error: Input Parameter nthRoot = ZERO!")
		return nextBundle, integerResidual, magnitudeResidual, thisBundleNo, err
	}

	totalDigits := big.NewInt(0).Add(integerMagnitude, big.NewInt(1))
	numOfBundlesInThisInteger := big.NewInt(0)
	remainingDigits := big.NewInt(0)
	scratch := big.NewInt(0)
	factor := big.NewInt(0)

	numOfBundlesInThisInteger, remainingDigits = big.NewInt(0).QuoRem(integerMagnitude, nthRoot, scratch)

	cmpNumOfBundles := numOfBundlesInThisInteger.Cmp(bigZero)
	cmpRemainDigits := remainingDigits.Cmp(bigZero)

	if cmpNumOfBundles == 0 && cmpRemainDigits == 0 {
		thisBundleNo = -1
		return nextBundle, integerResidual, magnitudeResidual, thisBundleNo, err
	}

	if cmpRemainDigits == 1 {
		factor.Set(remainingDigits)
	} else {
		// cmpNumOfBundles MUST BE == 1
		factor.Set(nthRoot)
	}

	delta := totalDigits.Sub(totalDigits, factor)
	scale := big.NewInt(0).Exp(big.NewInt(10), delta, nil)
	nextBundle,integerResidual = big.NewInt(0).QuoRem(integer,scale, scratch)
	magnitudeResidual = big.NewInt(0).Sub(delta, big.NewInt(1))
	thisBundleNo = intBundleNo + 1
	err = nil
	return nextBundle, integerResidual, magnitudeResidual, thisBundleNo, err
}