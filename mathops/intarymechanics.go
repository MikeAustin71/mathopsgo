package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryMechanics struct {
	lock *sync.Mutex
}

// getMagnitudeDigits
//
//	Returns the number of digits in the integer portion of the
//	numeric value in the IntAry instance passed as input parameter
//	'ia'.
//
//	IMPORTANT
//	=========
//
//	The calling function is responsible for verifying the validit
//	of 'ia', the Intary object.
func (iaMech *intAryMechanics) getMagnitudeDigits(
	ia *IntAry,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if iaMech.lock == nil {
		iaMech.lock = new(sync.Mutex)
	}

	iaMech.lock.Lock()

	defer iaMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMechanics.getMagnitudeDigits()",
		"")

	if err != nil {
		return 0, err
	}

	err = new(intAryNanobot).setInternalFlags(
		ia, ePrefix)

	if err != nil {

		return 0, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(ia, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	iaMagnitude := ia.intAryLen - ia.precision - ia.firstDigitIdx

	return iaMagnitude, nil
}
