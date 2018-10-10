package mathops

import "math/big"

type FixedDecimalNthRoot struct {
	NthRoot BigIntFixedDecimal
	Radicand BigIntFixedDecimal
	Root     BigIntFixedDecimal
}



// GetNextIntegerBundle - Returns the next bundle for an integer value.
func (fdNRt FixedDecimalNthRoot) GetNextIntegerBundle(
	integer, nRoot *big.Int) ( nextBundle, integerResidual *big.Int, err error) {

	nextBundle = big.NewInt(0)
	integerResidual = big.NewInt(0)
	err = nil




	return nextBundle, integerResidual, err
}