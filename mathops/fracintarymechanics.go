package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type fracIntAryMechanics struct {
	lock sync.Mutex
}

// getRationalValue
//
//	Converts the FracIntAry instance numeric value and returns the
//	value as a big rational	number (*big.Rat).
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
func (fracIntMech *fracIntAryMechanics) getRationalValue(
	fIa *FracIntAry,
	maxPrecision int,
	errPrefDto *ePref.ErrPrefixDto) (*big.Rat, error) {

	fracIntMech.lock.Lock()

	defer fracIntMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"fracIntAryMechanics.getRationalValue()",
		"")

	if err != nil {
		return big.NewRat(1, 1), err
	}

	if fIa == nil {

		return big.NewRat(1, 1), &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'fIa'",
		}
	}

	if maxPrecision < -1 {

		return big.NewRat(1, 1),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "maxPrecision < -1",
				ErrMessage: fmt.Sprintf("Error: Input parameter 'maxPrecision' is INVALID!\n"+
					"'maxPrecision' is less than -1\n"+
					"maxPrecision= %v", maxPrecision),
			}
	}

	if maxPrecision == -1 {
		maxPrecision = 4096
	}

	if fIa.Numerator.GetPrecision() == 0 &&
		fIa.Denominator.GetPrecision() == 0 {

		numeratorNumStr, err := fIa.Numerator.GetNumStr()

		if err != nil {

			return big.NewRat(1, 1),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "numeratorNumStr, err := fIa.Numerator.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		denominatorNumStr, err := fIa.Denominator.GetNumStr()

		if err != nil {

			return big.NewRat(1, 1),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "denominatorNumStr, err := fIa.Denominator.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		fRat, ok := big.NewRat(1, 1).SetString(numeratorNumStr + "/" + denominatorNumStr)

		if !ok {

			return big.NewRat(1, 1),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "fRat, ok := big.NewRat(1, 1).SetString(numeratorNumStr + \"/\" + denominatorNumStr)",
					ErrContext: "",
					ErrMessage: fmt.Sprintf("Function SetString() Failed!\n"+
						"numeratorNumStr = %v\ndenominatorNumStr = %v\n",
						numeratorNumStr, denominatorNumStr),
				}
		}

		return fRat, nil
	}

	newFloat, err := fIa.Numerator.DivideThisBy(&fIa.Denominator, 0, maxPrecision)

	if err != nil {

		return big.NewRat(1, 1),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newFloat, err := fIa.Numerator.DivideThisBy(&fIa.Denominator, 0, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newFloatNumStr, err := newFloat.GetNumStr()

	if err != nil {

		return big.NewRat(1, 1),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newFloatNumStr, err := newFloat.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fRat, ok := big.NewRat(1, 1).SetString(newFloatNumStr)

	if !ok {

		return big.NewRat(1, 1),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fRat, ok := big.NewRat(1, 1).SetString(newFloatNumStr)",
				ErrContext: "",
				ErrMessage: fmt.Sprintf("Function SetString() Failed!\n"+
					"newFloatNumStr= '%v'", newFloatNumStr),
			}
	}

	return fRat, nil
}

// isValidFracInt
func (fracIntMech *fracIntAryMechanics) isValidFracInt(
	fIa *FracIntAry,
	callingFunctions string) error {

	fracIntMech.lock.Lock()

	defer fracIntMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		callingFunctions+"\n",
		"fracIntAryMechanics.isValidFracInt",
		"")

	if err != nil {
		return err
	}

	if fIa == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'fIa'",
		}
	}

	var firstError, secondError error

	err = fIa.Numerator.IsValid(ePrefix.XCpy("Validating Numerator").String())

	if err != nil {

		firstError = &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = fIa.Numerator.IsValid(\n" +
				"ePrefix.XCpy(Validating Numerator).String())",
			ErrContext: "Error fIa.Numerator is INVALID!",
			ErrMessage: err.Error(),
		}
	}

	var err2 error

	err2 = fIa.Denominator.IsValid(ePrefix.XCpy("Validating Denominator").String())

	if err2 != nil {
		secondError = &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = fIa.Denominator.IsValid(\n" +
				"ePrefix.XCpy(Validating Denominator).String())",
			ErrContext: "Error fIa.Numerator is INVALID!",
			ErrMessage: err2.Error(),
		}
	}

	if err != nil && err2 != nil {
		return fmt.Errorf("%w\n%w", firstError, secondError)
	}

	if err != nil {
		return firstError
	}

	if err2 != nil {
		return secondError
	}

	return nil
}
