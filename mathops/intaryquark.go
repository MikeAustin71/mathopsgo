package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryQuark struct {
	lock *sync.Mutex
}

// setIntAryToZero
//
//	Sets the value of the intAry object to zero ('0').
func (iaQuark *intAryQuark) setIntAryToZero(
	intAry *IntAry,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaQuark.lock == nil {
		iaQuark.lock = new(sync.Mutex)
	}

	iaQuark.lock.Lock()

	defer iaQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryQuark.setIntAryToZero()",
		"")

	if err != nil {
		return err
	}

	intAry.intAryLen = 1 + int(precision)
	intAry.precision = int(precision)
	intAry.intAry = make([]uint8, intAry.intAryLen)
	intAry.signVal = 1

	err = new(intAryBoson).setNumericSeparatorsToDefaultIfEmpty(
		intAry, ePrefix.String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryBoson).\n" +
				"  setNumericSeparatorsToDefaultIfEmpty(intAry, ePrefix.String())",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).\n" +
				"  setInternalFlags(intAry, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

func (iaQuark *intAryQuark) setNumericSeparatorsToUSADefault(
	intAry *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaQuark.lock == nil {
		iaQuark.lock = new(sync.Mutex)
	}

	iaQuark.lock.Lock()

	defer iaQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryQuark.setNumericSeparatorsToUSADefault()",
		"")

	if err != nil {
		return err
	}

	if intAry == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intAry'",
		}
	}

	intAry.SetDecimalSeparator('.')
	intAry.SetThousandsSeparator(',')
	intAry.SetCurrencySymbol('$')

	return nil
}
