package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type fracIntAryNanobot struct {
	lock sync.Mutex
}

// getLowestCommonDenom
//
//	Returns a FracIntAry which represents the lowest common
//	denominator for the current FracIntAry.
//
//	Note: if 'maxPrecision' is less than 0, it is automatically
//	converted to '4,096' decimal places.
func (fracIntNanobot *fracIntAryNanobot) getLowestCommonDenom(
	fIa *FracIntAry,
	maxPrecision int,
	errPrefDto *ePref.ErrPrefixDto) (FracIntAry, error) {

	fracIntNanobot.lock.Lock()

	defer fracIntNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fracIntAryNanobot.getLowestCommonDenom()",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	if fIa == nil {

		return FracIntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'fIa'",
			}
	}

	if maxPrecision < 0 {
		maxPrecision = 4096
	}

	ratFrac, err := new(fracIntAryMechanics).getRationalValue(
		fIa, maxPrecision, ePrefix)

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "ratFrac, err := new(fracIntAryMechanics).getRationalValue(\n" +
					"  fIa, maxPrecision, ePrefix)",
				ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	newFAry, err := fIa.NewBigInts(ratFrac.Num(), ratFrac.Denom())

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "newFAry, err := \n" +
					"  fIa.NewBigInts(ratFrac.Num(), ratFrac.Denom())",
				ErrContext: fmt.Sprintf("ratFrac.Num()= '%v'\n"+
					"ratFrac.Denom()= '%v'", ratFrac.Num().String(), ratFrac.Denom().String()),
				ErrMessage: err.Error(),
			}
	}

	return newFAry, nil
}

// copyOut
//
//	Creates and returns a copy of the FracIntAry passed as an input
//	parameter.
func (fracIntNanobot *fracIntAryNanobot) copyOut(
	fIa *FracIntAry,
	errPrefDto *ePref.ErrPrefixDto) (FracIntAry, error) {

	fracIntNanobot.lock.Lock()

	defer fracIntNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fracIntAryNanobot.copyOut()",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	if fIa == nil {

		return FracIntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'fIa'",
			}
	}

	newFrac := FracIntAry{}

	newFrac.Numerator, err = fIa.Numerator.CopyOut()

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newFrac.Numerator, err = fIa.Numerator.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newFrac.Denominator, err = fIa.Denominator.CopyOut()

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newFrac.Denominator, err = fIa.Denominator.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return newFrac, nil
}

// CopyIn
//
//	Receives a pointer to an incoming FracIntAry ('fIa2') and
//	copies the	values into the FracIntAry 'fIa'.
func (fracIntNanobot *fracIntAryNanobot) copyIn(
	fIaDestination *FracIntAry,
	fIa2Source *FracIntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	fracIntNanobot.lock.Lock()

	defer fracIntNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fracIntAryNanobot.copyOut()",
		"")

	if err != nil {
		return err
	}

	if fIaDestination == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'fIaDestination'",
		}
	}

	if fIa2Source == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'fIa2Source'",
		}
	}

	fIaDestination.Numerator, err = fIa2Source.Numerator.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "fIaDestination.Numerator, err =\n" +
				"fIa2Source.Numerator.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	fIaDestination.Denominator, err = fIa2Source.Denominator.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "fIaDestination.Denominator, err =\n" +
				" fIa2Source.Denominator.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
