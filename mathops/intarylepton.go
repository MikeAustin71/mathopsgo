package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryLepton struct {
	lock *sync.Mutex
}

// isValidNumericValues
//
//	Examines an IntAry object and returns an error if	that IntAry
//	object is found to contain invalid numeric values.
//
//	IMPORTANT
//	=========
//
//	This method ONLY tests validity of IntAry internal numeric
//	value fields. It does NOT test the validity of the internal
//	Numeric Separator fields. To perform vaidity testing on Numeric
//	Sepators, see method 'intAryLepton.isValidNumSeps'.
//
//	To call both isValidIntAry and isValidNumSeps, call
//	'intAryLepton.isValidIntary'
func (iAryLepton *intAryLepton) isValidNumericValues(
	intAry *IntAry,
	errName string) error {

	if iAryLepton.lock == nil {
		iAryLepton.lock = new(sync.Mutex)
	}

	iAryLepton.lock.Lock()

	defer iAryLepton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errName,
		"intAryLepton.isValidNumericValues()",
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

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(ia, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if intAry.signVal != -1 && intAry.signVal != 1 {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: sign Value is INVALID!\n"+
			"sign Value= '%v'\n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix,
			intAry.signVal)
	}

	if intAry.precision < 0 {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: precision Value is INVALID!\n"+
			"'precision' value is Less Than Zero!\n"+
			"'precision' Value= '%v'\n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix,
			intAry.precision)
	}

	if intAry.precision >= intAry.intAryLen {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: 'precision' value is greater than or equal to IntArray length.\n"+
			"iAry.precision= '%v'\n"+
			"iAry.intAryLen= '%v' \n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix,
			intAry.precision,
			intAry.intAryLen)

	}

	if intAry.integerLen == 0 {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: IntAry integer length is zero.\n"+
			"Missing leading integer zero!\n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix)
	}

	return nil
}

// isValidNumSeps
//
//	This validation method only tests the validity of Numeric
//	Separators contained in the IntAry object passed as input
//	parameter 'intAry'.
func (iAryLepton *intAryLepton) isValidNumSeps(
	intAry *IntAry,
	errName string) error {

	if iAryLepton.lock == nil {
		iAryLepton.lock = new(sync.Mutex)
	}

	iAryLepton.lock.Lock()

	defer iAryLepton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errName,
		"intAryLepton.isValidNumSeps()",
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

	var errorString string

	if intAry.decimalSeparator == 0 {

		errorString += "  Decimal Separator is invalid.\n"

	}

	if intAry.thousandsSeparator == 0 {
		errorString += "  Thousands Separator is invalid.\n"
	}

	if intAry.currencySymbol == 0 {
		errorString += "  Currency Symbol is invalid.\n"
	}

	if errorString != "" {

		return fmt.Errorf("%v\n"+
			"Error: Numeric Separators invalid.\n"+
			"%v\n", ePrefix.String(), errorString)
	}

	return nil
}
