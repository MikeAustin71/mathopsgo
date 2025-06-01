package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryElectron struct {
	lock *sync.Mutex
}

// Empty - Basically resets all the fields of the intAry
// structure to their 'zero' values.
func (iAryElectron *intAryElectron) empty(
	intAry *IntAry) {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	if intAry == nil {
		return
	}

	intAry.intAry = []uint8{}
	intAry.intAryLen = 0
	intAry.integerLen = 0
	intAry.significantIntegerLen = 0
	intAry.significantFractionLen = 0
	intAry.firstDigitIdx = -1
	intAry.lastDigitIdx = -1
	intAry.isZeroValue = true
	intAry.isIntegerZeroValue = true
	intAry.precision = 0
	intAry.signVal = 1

	intAry.SetDecimalSeparator('.')
	intAry.SetThousandsSeparator(',')
	intAry.SetCurrencySymbol('$')

	intAry.BackUp = new(BackUpIntAry).New()

	return
}

// equals
//
//	Receives two instances of IntAry and compares the values of all
//	member data fields to determine if they are equivalent in all
//	respects.
//
//	Returns 'true' if all member field values of 'iAry1' are equal
//	to the corresponding field values of 'iAry2'.
//
//	Note that the BackUp fields for both compared IntAry objects
//	are NOT included in the 'Equals' comparison.
//
//	If any errors are encountered, a boolean value of 'false' is
//	returned
func (iAryElectron *intAryElectron) equals(
	iAry1 *IntAry,
	iAry2 *IntAry) bool {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	if iAry1 == nil || iAry2 == nil {
		return false
	}

	iaNanobot := new(intAryNanobot)
	var err error

	err = iaNanobot.setInternalFlags(iAry1, nil)

	if err != nil {
		return false
	}

	err = iaNanobot.setInternalFlags(iAry2, nil)

	if err != nil {
		return false
	}

	return new(intAryBoson).dataFieldEqualityTest(iAry1, iAry2)
}

// newIntAry
//
// Generates a new instance of IntAry
func (iAryElectron *intAryElectron) newIntAry() IntAry {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	iAry := IntAry{}
	iAry.intAry = []uint8{}
	iAry.intAryLen = 0
	iAry.integerLen = 0
	iAry.significantIntegerLen = 0
	iAry.significantFractionLen = 0
	iAry.firstDigitIdx = -1
	iAry.lastDigitIdx = -1
	iAry.isZeroValue = true
	iAry.isIntegerZeroValue = true
	iAry.precision = 0
	iAry.signVal = 1
	iAry.decimalSeparator = '.'
	iAry.thousandsSeparator = ','
	iAry.currencySymbol = '$'
	iAry.BackUp = new(BackUpIntAry).New()

	return iAry
}

// isValidIntAry
//
//	Examines an intAry object and returns an error if
//	that intAry object is found to be invalid.
func (iAryElectron *intAryElectron) isValidIntAry(
	iAry *IntAry,
	errName string) error {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errName,
		"intAryElectron.isValidIntAry()",
		"")

	if err != nil {
		return err
	}

	if iAry == nil {

		return fmt.Errorf("%v\n"+
			"IntAry Validation Test.\n"+
			"Error: Input parameter 'iAry' is a nil pointer!\n"+
			"IntAry object FAILED Validation Test!\n",
			ePrefix.String())
	}

	err = new(intAryNanobot).setInternalFlags(
		iAry, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(ia, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if iAry.signVal != -1 && iAry.signVal != 1 {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: sign Value is INVALID!\n"+
			"sign Value= '%v'\n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix,
			iAry.signVal)
	}

	if iAry.precision < 0 {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: precision Value is INVALID!\n"+
			"'precision' value is Less Than Zero!\n"+
			"'precision' Value= '%v'\n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix,
			iAry.precision)
	}

	if iAry.precision >= iAry.intAryLen {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: 'precision' value is greater than or equal to IntArray length.\n"+
			"iAry.precision= '%v'\n"+
			"iAry.intAryLen= '%v' \n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix,
			iAry.precision,
			iAry.intAryLen)

	}

	if iAry.integerLen == 0 {

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

// setIntAryLength
//
//	Calculates the current IntAry string length and sets internal
//	variable 'ia.intAryLen'.
func (iAryElectron *intAryElectron) setIntAryLength(
	intAry *IntAry,
	callingFunction string) error {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		callingFunction,
		"intAryElectron.setIntAryLength()",
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

	intAry.intAryLen = len(intAry.intAry)

	return nil
}

// setSignificantDigitIdxs - Finds the first
// significant digit (the first numeric digit
// greater than zero) and sets index value in
// the local field variable, 'firstDigitIdx'.
//
// In addition, this method also identifies the
// Last Significant Digit (the last non-zero value
// in the intAry) and records that index in the
// local field variable, 'lastDigitIdx'.
func (iAryElectron *intAryElectron) setSignificantDigitIdxs(
	ia *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryElectron.setSignificantDigitIdxs()",
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

	err = ia.SetNumericSeparatorsToDefaultIfEmpty()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia.SetNumericSeparatorsToDefaultIfEmpty()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	ia.intAryLen = len(ia.intAry)

	if ia.intAryLen == ia.precision {

		ia.intAry = append([]uint8{0}, ia.intAry...)

		ia.intAryLen++
	}

	if ia.intAryLen < ia.precision {

		deltaZeros := ia.precision - ia.intAryLen + 1

		zeroAry := make([]uint8, deltaZeros)

		ia.intAry = append(zeroAry, ia.intAry...)

		ia.intAryLen += deltaZeros
	}

	ia.firstDigitIdx = -1
	ia.lastDigitIdx = -1

	ia.integerLen = 0
	ia.significantIntegerLen = 0
	ia.significantFractionLen = 0

	lastIntIdx := ia.intAryLen - ia.precision - 1
	ia.isZeroValue = true
	ia.isIntegerZeroValue = true
	ia.integerLen = ia.intAryLen - ia.precision

	for i := 0; i < ia.intAryLen; i++ {
		if ia.intAry[i] > 0 {

			ia.isZeroValue = false

			if i < ia.integerLen {

				ia.isIntegerZeroValue = false
			}
		}

		// At minimum, there should be a single
		// leading zero before the decimal point.
		// Example 0.000.
		if i == lastIntIdx && ia.intAry[i] == 0 {

			if ia.firstDigitIdx == -1 {

				ia.firstDigitIdx = i
			}

		}

		if ia.intAry[i] > 0 {

			if ia.firstDigitIdx == -1 {

				ia.firstDigitIdx = i
			}

			ia.lastDigitIdx = i
		}

	}

	ia.significantIntegerLen = ia.intAryLen - ia.precision - ia.firstDigitIdx

	if ia.lastDigitIdx >= ia.integerLen {

		ia.significantFractionLen = ia.lastDigitIdx - ia.integerLen + 1
	} else {

		ia.significantFractionLen = 0
	}

	return nil
}
