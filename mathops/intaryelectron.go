package mathops

import (
	"fmt"
	"sync"
)

type intAryElectron struct {
	lock *sync.Mutex
}

func (iAryElectron *intAryElectron) newIntAry() IntAry {

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

	ePrefix := "intAryElectron.isValidIntAry()"

	if iAry == nil {

		return fmt.Errorf("%v\n"+
			"Called Method: %v\n"+
			"IntAry Validation Test.\n"+
			"Error: Input parameter 'iAry' is a nil pointer!\n"+
			"IntAry object FAILED Validation Test!\n",
			errName,
			ePrefix)
	}

	if len(errName) == 0 {

		errName = "intAryElectron.IsValid()"

	}

	iAry.SetInternalFlags()

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
