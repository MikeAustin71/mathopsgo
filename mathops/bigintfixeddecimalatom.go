package mathops

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntFixedDecAtom struct {
	lock *sync.Mutex
}

// isValid
//
// This method performs validity testing on an instance of
// BigIntFixedDecimal passed as input parameter, 'bigIFxDec'.
//
// If this BigIntFixedDecimal instance fails the validity
// tests, an error will be returned.
//
// If this BigIntFixedDecimal instance passes all validity
// tests, an error value of 'nil' will be returned.
func (bigIFdAtom *bigIntFixedDecAtom) isBigIntFxDecValid(
	bigIFxDec *BigIntFixedDecimal,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdAtom.lock == nil {
		bigIFdAtom.lock = new(sync.Mutex)
	}

	bigIFdAtom.lock.Lock()

	defer bigIFdAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecAtom.isBigIntFxDecValid",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}

		return err
	}

	if bigIFxDec.integerNum == nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: This instance of BigIntFixedDecimal is INVALID!\n" +
				"BigIntFixedDecimal.integerNum is a 'nil' pointer.\n" +
				"This BigIntFixedDecimal instance FAILED Validation Testing!",
		}
	}

	if bigIFxDec.decimalSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "This BigIntFixedDecimal Instance is Invalid!\n" +
				"'bigIFxDec.decimalSeparator' is empty with a Zero value.\n" +
				"FATAL ERROR!",
		}
	}

	if bigIFxDec.thousandsSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "This BigIntFixedDecimal Instance is Invalid!\n" +
				"'bigIFxDec.thousandsSeparator' is empty with a Zero value.\n" +
				"FATAL ERROR!",
		}
	}

	if bigIFxDec.currencySymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "This BigIntFixedDecimal Instance is Invalid!\n" +
				"'bigIFxDec.currencySymbol' is empty with a Zero value.\n" +
				"FATAL ERROR!",
		}
	}

	return nil
}

// setNumericValue
//
// Sets the 'integerNum' and 'precision' values for the
// BigIntFixedDecimal instance passed as input parameter,
// 'bigIFxDec'. Taken together, 'integerNum' and 'precision'
// describe a numeric value with a fixed number of fractional
// digits to the right of the decimal place.
//
//	Numeric Separators
//	==================
//
//	The numeric separators used to configure BigIntFixedDecimal
//	instance 'bigIFxDec' will be taken from input parameter,
//	'numSepsDto'. If any member elements of numSepsDto are
//	invalid, an error will be returned.
func (bigIFdAtom *bigIntFixedDecAtom) setNumericValue(
	bigIFxDec *BigIntFixedDecimal,
	integer *big.Int,
	precision uint,
	numSepsDto NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdAtom.lock == nil {
		bigIFdAtom.lock = new(sync.Mutex)
	}

	bigIFdAtom.lock.Lock()

	defer bigIFdAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecAtom.setNumericValue",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}
	}

	if integer == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'integer'",
		}

	}

	err = numSepsDto.IsValid(
		ePrefix.XCpy("Testing numSepsDto").String())

	if err != nil {
		return err
	}

	bigIFxDec.integerNum = big.NewInt(0).Set(integer)

	bigIFxDec.precision = precision

	return new(bigIntFixedDecBoson).setNumericSeparatorsDto(
		bigIFxDec,
		numSepsDto,
		ePrefix.XCpy("Setting 'bigIFxDec'"))
}
