package mathops

import (
	"fmt"
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

// FracIntAry - A fraction represented by a numerator and a denominator.
// Both numerator and denominator are of type intAry
type FracIntAry struct {
	Numerator   IntAry
	Denominator IntAry
}

// IsValid
//
//	This method performs validity testing on the
//	current instance of FracIntAry.
//
//	If the current FracIntAry instance is valid,
//	this method returns 'nil'.
//
//	If the current FracIntAry instance is invalid,
//	this method includes an appropriate error message
//	in the returned 'error' object.
func (fIa *FracIntAry) IsValid(callingFunctions string) error {

	callingFunctions += "\nFracIntAry.IsValid()"

	return new(fracIntAryMechanics).isValidFracInt(fIa, callingFunctions)
}

// NewBigInts
//
//	Creates a new FracIntAry type from two *big.Int types passed as
//	input parameters.
func (fIa *FracIntAry) NewBigInts(numerator, denominator *big.Int) (FracIntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.NewBigInts",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	iaNumerator, err := new(IntAry).NewBigInt(numerator, 0)

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaNumerator, err := new(IntAry).NewBigInt(numerator, 0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iaDenominator, err := new(IntAry).NewBigInt(denominator, 0)

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaDenominator, err := new(IntAry).NewBigInt(denominator, 0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newFracIntAry, err := new(FracIntAry).NewIntArys(&iaNumerator, &iaDenominator)

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newFracIntAry, err := new(FracIntAry).NewIntArys(&iaNumerator, &iaDenominator)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return newFracIntAry, nil
}

// NewNumStrs
//
//	Creates a new FracIntAry type by passing input parameters
//	numerator and denominator as number strings.
func (fIa *FracIntAry) NewNumStrs(numerator string, denominator string) (FracIntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.NewNumStrs",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	fIa2 := FracIntAry{}

	fIa2.Numerator, err = new(IntAry).NewNumStr(numerator)

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fIa2.Numerator, err = new(IntAry).NewNumStr(numerator)",
				ErrContext: fmt.Sprintf("numerator= '%v'", numerator),
				ErrMessage: err.Error(),
			}
	}

	fIa2.Denominator, err = new(IntAry).NewNumStr(denominator)

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fIa2.Denominator, err = new(IntAry).NewNumStr(denominator)",
				ErrContext: fmt.Sprintf("denominator= '%v'", denominator),
				ErrMessage: err.Error(),
			}
	}

	return fIa2, nil
}

// NewIntArys
//
//	Creates a type FracIntAry by passing numerator and denominator
//	input parameters of type *IntAry
func (fIa *FracIntAry) NewIntArys(numerator *IntAry, denominator *IntAry) (FracIntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.NewIntArys",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	fIa2 := FracIntAry{}

	fIa2.Numerator, err = numerator.CopyOut()

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fIa2.Numerator, err = numerator.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fIa2.Denominator, err = denominator.CopyOut()

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fIa2.Denominator, err = denominator.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fIa2, nil
}

// NewFracIntAry
//
//	Creates a FracIntAry instance from a single IntAry object. The
//	IntAry input parameter is converted into an equivalent
//	fraction.
//
//	  ia
//	 ----
//	  1
func (fIa *FracIntAry) NewFracIntAry(ia *IntAry) (FracIntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.NewFracIntAry",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	fIa2 := FracIntAry{}

	if ia.GetPrecision() == 0 {

		fIa2.Numerator, err = ia.CopyOut()

		if err != nil {

			return FracIntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "fIa2.Numerator, err = ia.CopyOut()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		fIa2.Denominator, err = new(IntAry).NewOne(0)

		if err != nil {

			return FracIntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "fIa2.Denominator, err = new(IntAry).NewOne(0)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return fIa2, nil
	}

	precision := ia.GetPrecision()

	fIa2.Numerator, err = ia.CopyOut()

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fIa2.Numerator, err = ia.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if precision > 0 {

		err = fIa2.Numerator.ShiftPrecisionRight(uint(precision))

		if err != nil {

			return FracIntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = fIa2.Numerator.ShiftPrecisionRight(uint(precision))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	fIa2.Denominator, err = new(IntAry).NewOne(0)

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fIa2.Denominator, err = new(IntAry).NewOne(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(IntAryMathMultiply).MultiplyByTenToPower(&fIa2.Denominator, uint(precision))

	if err != nil {

		return FracIntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(IntAryMathMultiply).\n" +
					"  MultiplyByTenToPower(&fIa2.Denominator, uint(precision))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fIa2, nil
}

// CopyOut
//
//	Creates and returns a copy of the current FracIntAry.
func (fIa *FracIntAry) CopyOut() (FracIntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.CopyOut",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	return new(fracIntAryNanobot).copyOut(
		fIa, ePrefix)
}

// CopyIn
//
//	Receives a pointer to an incoming FracIntAry and copies the
//	values into the current FracIntAry.
func (fIa *FracIntAry) CopyIn(fIa2 *FracIntAry) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.CopyIn",
		"")

	if err != nil {
		return err
	}

	return new(fracIntAryNanobot).copyIn(fIa, fIa2, ePrefix)
}

// GetRationalValue
//
//	Converts the fraction and returns the value as a big rational
//	number (*big.Rat).
//
//	maxPrecision
//	============
//
//	Input parameter 'maxPrecision' determines the maximum number of
//	decimal places to the right of the decimal point contained in
//	the result.
//
//	If the value of 'maxPrecision' is -1, maximum precision will
//	default to 4096 decimal places.
//
//	'maxPrecision' values less than -1 will trigger an error.
func (fIa *FracIntAry) GetRationalValue(maxPrecision int) (*big.Rat, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.GetRationalValue",
		"")

	if err != nil {
		return big.NewRat(1, 1), err
	}

	return new(fracIntAryMechanics).getRationalValue(
		fIa, maxPrecision, ePrefix)
}

// GetLowestCommonDenom
//
//	Returns a FracIntAry which represents the lowest common
//	denominator for the current FracIntAry.
//
//	Note: if 'maxPrecision' is less than 0, it is automatically
//	converted to '4,096' decimal places.
func (fIa *FracIntAry) GetLowestCommonDenom(maxPrecision int) (FracIntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.GetLowestCommonDenom",
		"")

	if err != nil {
		return FracIntAry{}, err
	}

	return new(fracIntAryNanobot).getLowestCommonDenom(
		fIa, maxPrecision, ePrefix)
}

// ReduceToLowestCommonDenom
//
//	Converts the value of the current FracIntAry to its lowest
//	common denominator.
//
//	Note: if 'maxPrecision' is less than 0, it is automatically
//	converted to '4,096' decimal places.
func (fIa *FracIntAry) ReduceToLowestCommonDenom(maxPrecision int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FracIntAry.ReduceToLowestCommonDenom",
		"")

	if err != nil {
		return err
	}

	if maxPrecision < 0 {
		maxPrecision = 4096
	}

	fracIntNanobot := new(fracIntAryNanobot)

	fIaLCD, err := fracIntNanobot.getLowestCommonDenom(
		fIa, maxPrecision, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "fIaLCD, err := new(fracIntAryNanobot).\n" +
				"getLowestCommonDenom(fIa, maxPrecision, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = fracIntNanobot.copyIn(fIa, &fIaLCD, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = fracIntNanobot.copyIn(fIa, &fIaLCD, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
