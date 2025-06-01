package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryProton struct {
	lock *sync.Mutex
}

// copy
//
//	Makes a deep copy of the current IntAry instance.
//
//	The copy operation copies a source IntAry instance to a
//	destination IntAry instance.
//
//	If input parameter 'validateSourceIntAry' is set to 'true',
//	validation tests will be applied to input parameter 'iaSource'.
//
//	If input parameter 'copyToBackup' is set to 'true', a copy
//	of 'iaSource' will be copied to 'iaDestination.BackUp'.
func (iaProton *intAryProton) copy(
	iaDestination *IntAry,
	iaSource *IntAry,
	validateSourceIntAry bool,
	copyToBackup bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaProton.lock.Lock()

	defer iaProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryProton.copy()",
		"")

	if err != nil {
		return err
	}

	if iaDestination == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'iaDestination'",
		}
	}

	if iaSource == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'iaSource'",
		}
	}

	if validateSourceIntAry {

		err = new(intAryElectron).isValidIntAry(
			iaSource, ePrefix.XCpy("Validating 'iaSource'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia, ePrefix.XCpy(Validating 'ia').String())",
				ErrContext: "IntAry instanace 'iaSource' is INVALID!\n" +
					"'iaSource' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	err = new(intAryNanobot).setInternalFlags(
		iaSource, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(iaSource, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	new(intAryElectron).empty(iaDestination)

	iaDestination.intAry = make([]uint8, iaSource.intAryLen)

	for i := 0; i < iaSource.intAryLen; i++ {
		iaDestination.intAry[i] = iaSource.intAry[i]
	}

	iaDestination.intAryLen = iaSource.intAryLen
	iaDestination.integerLen = iaSource.integerLen
	iaDestination.significantIntegerLen = iaSource.significantIntegerLen
	iaDestination.significantFractionLen = iaSource.significantFractionLen
	iaDestination.firstDigitIdx = iaSource.firstDigitIdx
	iaDestination.lastDigitIdx = iaSource.lastDigitIdx
	iaDestination.isZeroValue = iaSource.isZeroValue
	iaDestination.isIntegerZeroValue = iaSource.isIntegerZeroValue
	iaDestination.precision = iaSource.precision
	iaDestination.signVal = iaSource.signVal
	iaDestination.decimalSeparator = iaSource.decimalSeparator
	iaDestination.thousandsSeparator = iaSource.thousandsSeparator
	iaDestination.currencySymbol = iaSource.currencySymbol

	if copyToBackup {

		err = new(intAryNeutron).copyToBackup(
			iaDestination,
			iaSource,
			false,
			ePrefix)
	}

	return err
}

// CopyOutDigits
//
//	Makes a deep copy of the IntAry instance passed as input
//	parameter 'iaSource'.
//
//	This method uses input parameter 'digitsToCopy' to copy a
//	specified number of digits to the new copy.
//
//	If input parameter 'validateSourceIntAry' is set to true, input
//	parameter 'iaSource' will be subjected to validation tests.
//
// If input parameter 'copyToBackup' is set to true, a duplicate
// copy of the transformed IntAry oject ('iaDestination') will be
// stored in the 'BackUp' Field of 'iaDestination'.
func (iaProton *intAryProton) copyOutDigits(
	iaSource *IntAry,
	validateSourceIntAry bool,
	digitsToCopy int,
	copyToBackup bool,
	errPrefDto *ePref.ErrPrefixDto) (iaDestination IntAry, err error) {

	iaProton.lock.Lock()

	defer iaProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryProton.copyOutDigits()",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if iaSource == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'iaSource'",
			}
	}

	iaElectron := new(intAryElectron)

	if validateSourceIntAry {

		err = iaElectron.isValidIntAry(
			iaSource, ePrefix.XCpy("Validating 'iaSource'").String())

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia, ePrefix.XCpy(Validating 'ia').String())",
					ErrContext: "IntAry instanace 'iaSource' is INVALID!\n" +
						"'iaSource' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	iaNanobot := new(intAryNanobot)

	err = iaNanobot.setInternalFlags(
		iaSource, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(intAryNanobot).setInternalFlags(iaSource, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iaSrcAryLen := len(iaSource.intAry)

	if digitsToCopy > iaSrcAryLen {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "digitsToCopy > iaSrcAryLen",
				ErrMessage: "Error: The number of digits to copy is greater\n" +
					"than the length of the iaSource IntAry.",
			}
	}

	iAry2 := iaElectron.newIntAry()

	iAry2.intAry = make([]uint8, digitsToCopy)

	for i := 0; i < digitsToCopy; i++ {

		iAry2.intAry[i] = iaSource.intAry[i]

	}

	iAry2.intAryLen = digitsToCopy

	if iaSource.integerLen < digitsToCopy {

		iAry2.precision = digitsToCopy - iaSource.integerLen

	} else {
		iAry2.precision = 0
	}

	iAry2.signVal = iaSource.signVal
	iAry2.decimalSeparator = iaSource.decimalSeparator
	iAry2.thousandsSeparator = iaSource.thousandsSeparator
	iAry2.currencySymbol = iaSource.currencySymbol
	iAry2.BackUp = new(BackUpIntAry).New()

	err = iaNanobot.setInternalFlags(
		&iAry2, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(intAryNanobot).setInternalFlags(&iAry2, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if copyToBackup {
		err = new(intAryNeutron).copyToBackup(
			&iAry2, &iAry2, false, ePrefix)

		if err != nil {

			return IntAry{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryNeutron).copyToBackup(\n" +
						"  &iAry2, &iAry2, false, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	return iAry2, nil
}
