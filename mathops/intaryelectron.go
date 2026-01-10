package mathops

import (
	"sync"

	ePref "github.com/MikeAustin71/errpref"
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

	intAry.decimalSeparator = '.'
	intAry.thousandsSeparator = ','
	intAry.currencySymbol = '$'

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
//		Examines an IntAry object and returns an error if	that IntAry
//		object is found to be invalid.
//
//		IMPORTANT
//		=========
//
//		This method tests both the validity of IntAry internal
//	 	numeric value fields 'AND' the validity of the internal
//	 	Numeric Separator fields.
//
//	 	To perform vaidity testing separately on numeric values
//	 	and Numeric Sepators, see the following methods:
//
//	 	    'intAryLepton.isValidNumericValues'
//
//	 	    'intAryLepton.isValidNumSeps'
//
//	 To call both isValidIntAry and isValidNumSeps, call
//	 'intAryLepton.isValidIntary'
func (iAryElectron *intAryElectron) isValidIntAry(
	intAry *IntAry,
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

	var iaLepton = new(intAryLepton)

	err = iaLepton.isValidNumericValues(
		intAry, ePrefix.XCpy("Testing Numeric Values").String())

	if err != nil {
		return err
	}

	err = iaLepton.isValidNumSeps(
		intAry, ePrefix.XCpy("Testing Numeric Separators").String())

	return err
}

// setIntAryLength
//
//	Calculates the current IntAry string length and sets internal
//	variable 'ia.intAryLen'.
func (iAryElectron *intAryElectron) setIntAryLength(
	intAry *IntAry,
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

	err = new(intAryPhoton).setNumericSeparatorsToDefaultIfEmpty(
		ia, ePrefix.XCpy("Set 'ia' Numeric Separators"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryPhoton).setNumericSeparatorsToDefaultIfEmpty(\n" +
				"ia, ePrefix.XCpy(Set 'ia' Numeric Separators))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	ia.intAryLen = len(ia.intAry)

	if ia.intAryLen == 0 {

		ia.firstDigitIdx = -1
		ia.lastDigitIdx = -1

		ia.integerLen = 0
		ia.significantIntegerLen = 0
		ia.significantFractionLen = 0
		ia.isZeroValue = true
		ia.isIntegerZeroValue = true
		ia.precision = 0

		return nil
	}

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

// setSignificantDigitIdxsNoErrors
//
//	This method differs from:
//	  intAryElection.setSignificantDigitIdxs
//	in that this method does NOT return an error.
//
//	This method finds the first significant digit (the first
//	numeric digit greater than zero) and sets index value in the
//	local field variable, 'firstDigitIdx'.
//
//	In addition, this method also identifies the Last Significant
//	Digit (the last non-zero value in the intAry) and records that
//	index in the local field variable, 'lastDigitIdx'. Finally,
//	this method will set all internal flags.
//
//	If any of the Numeric Separators are invalid, theu will be
//	reset to USA Defaults.
func (iAryElectron *intAryElectron) setSignificantDigitIdxsNoErrors(
	ia *IntAry) {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	if ia == nil {

		return
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

	if ia.decimalSeparator == 0 {
		ia.decimalSeparator = '.'
	}

	if ia.thousandsSeparator == 0 {
		ia.thousandsSeparator = '.'
	}

	if ia.currencySymbol == 0 {
		ia.currencySymbol = '$'
	}

	return
}
