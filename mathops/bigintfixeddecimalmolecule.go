package mathops

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
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

// areBigIntFixDecsEqualValue
//
// Receives pointers to two instances of BigIntFixedDecimal and proceeds
// to compare the numeric values. If the two values are equal, this
// method returns 'true'.
//
// If the two values are not equal, this method returns 'false'.
//
// If the two 'precision' values are not equal, this method returns 'false'.
//
// The Numeric Separators are NOT included in this comparison. If the two
// values are equal but the Numeric Separators are unequal, this method
// will return true.
func (bigIFdMolecule *bigIntFixedDecMolecule) areBigIntFixDecsEqualValue(
	bigIFxDecBase *BigIntFixedDecimal,
	validateBigIFxDecBase bool,
	bigIFxDecComparison *BigIntFixedDecimal,
	validateBigIFxDecComparison bool,
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

	if bigIFxDecBase == nil {

		return false, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}
	}

	if bigIFxDecComparison == nil {

		return false, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}
	}

	bIFxDecAtom := new(bigIntFixedDecAtom)

	if validateBigIFxDecBase {

		err = bIFxDecAtom.isBigIntFxDecValid(
			bigIFxDecBase,
			ePrefix.XCpy("Validating bigIFxDecBase"))

		if err != nil {

			return false,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = bIFxDecAtom.isBigIntFxDecValid(\n" +
						"bigIFxDecBase, ePrefix.XCpy('Validating bigIFxDecBase')\n ",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateBigIFxDecComparison {

		err = bIFxDecAtom.isBigIntFxDecValid(
			bigIFxDecComparison,
			ePrefix.XCpy("Validating bigIFxDecComparison"))

		if err != nil {

			return false,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = bIFxDecAtom.isBigIntFxDecValid(\n" +
						"bigIFxDecComparison, ePrefix.XCpy('Validating bigIFxDecComparison')\n ",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	if bigIFxDecBase.precision != bigIFxDecComparison.precision {

		return false, nil
	}

	if bigIFxDecBase.integerNum.Cmp(bigIFxDecComparison.integerNum) != 0 {

		return false, nil
	}

	return true, nil
}
