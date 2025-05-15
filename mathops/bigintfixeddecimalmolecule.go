package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntFixedDecMolecule struct {
	lock *sync.Mutex
}

// isBigIntFxDecZero
//
// Returns true only if the current BigIntFixedDecimal numeric
// value is equal to zero. If the current BigIntFixedDecimal
// is invalid, an error will be returned.
//
// This method does NOT test the validity of 'bigIFxDec', an
// instance of type BigIntNum. The calling method must
// do this!
func (bigIFdMolecule *bigIntFixedDecMolecule) isBigIntFxDecZero(
	bigIFxDec *BigIntFixedDecimal,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

	if bigIFdMolecule.lock == nil {
		bigIFdMolecule.lock = new(sync.Mutex)
	}

	bigIFdMolecule.lock.Lock()

	defer bigIFdMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecMolecule.isBigIntFxDecZero",
		"")

	if err != nil {
		return false, err
	}

	if bigIFxDec == nil {

		return false, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}
	}

	if bigIFxDec.integerNum == nil {

		return false, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec.integerNum'",
		}
	}

	if bigIFxDec.integerNum.Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}

	return false, nil
}
