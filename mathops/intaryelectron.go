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
	ia *IntAry) {

	if iAryElectron.lock == nil {
		iAryElectron.lock = new(sync.Mutex)
	}

	iAryElectron.lock.Lock()

	defer iAryElectron.lock.Unlock()

	if ia == nil {
		return
	}

	ia.intAry = []uint8{}
	ia.intAryLen = 0
	ia.integerLen = 0
	ia.significantIntegerLen = 0
	ia.significantFractionLen = 0
	ia.firstDigitIdx = -1
	ia.lastDigitIdx = -1
	ia.isZeroValue = true
	ia.isIntegerZeroValue = true
	ia.precision = 0
	ia.signVal = 1

	ia.SetDecimalSeparator('.')
	ia.SetThousandsSeparator(',')
	ia.SetCurrencySymbol('$')

	return
}

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
