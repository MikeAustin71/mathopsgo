package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryNanobot struct {
	lock *sync.Mutex
}

// setInternalFlags - Sets Array Lengths and
// test for zero values
//
//	IMPORTANT
//	=========
//
//	The calling function is responsible for verifying the validit
//	of 'ia', the Intary object.
func (iaNanobot *intAryNanobot) setInternalFlags(
	ia *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaNanobot.lock == nil {
		iaNanobot.lock = new(sync.Mutex)
	}

	iaNanobot.lock.Lock()

	defer iaNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNanobot.setInternalFlags()",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia'",
		}
	}

	err = new(intAryElectron).setSignificantDigitIdxs(
		ia,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryElectron).\n" +
				"  setSignificantDigitIdxs( ia, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
