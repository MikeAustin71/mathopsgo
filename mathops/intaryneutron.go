package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type intAryNeutron struct {
	lock *sync.Mutex
}

// copyToBackup
//
//	This method receives two IntAry objects, 'iaDestination' and
//	'iaSource'. It then proceeds to copy the primary data fields
//	from 'iaSource' to the 'BackUp' fields of 'iaDestination'.
func (iaNeutron *intAryNeutron) copyToBackup(
	iaDestination *IntAry,
	iaSource *IntAry,
	validateSourceIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaNeutron.lock == nil {
		iaNeutron.lock = new(sync.Mutex)
	}

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getBigInt()",
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
			ParameterName: "'iaDestination'",
		}
	}

	if validateSourceIntAry {

		err = new(intAryElectron).isValidIntAry(
			iaSource, ePrefix.XCpy("Validating 'iaSource'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(intAryElectron).isValidIntAry(iaSource, ePrefix.XCpy(Validating 'iaSource').String())",
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

	iaDestination.BackUp.Empty()

	iaDestination.BackUp.intAry = make([]uint8, iaSource.intAryLen)

	for i := 0; i < iaSource.intAryLen; i++ {

		iaDestination.BackUp.intAry[i] = iaSource.intAry[i]
	}

	iaDestination.BackUp.intAryLen = iaSource.intAryLen
	iaDestination.BackUp.integerLen = iaSource.integerLen
	iaDestination.BackUp.significantIntegerLen = iaSource.significantIntegerLen
	iaDestination.BackUp.significantFractionLen = iaSource.significantFractionLen
	iaDestination.BackUp.firstDigitIdx = iaSource.firstDigitIdx
	iaDestination.BackUp.lastDigitIdx = iaSource.lastDigitIdx
	iaDestination.BackUp.isZeroValue = iaSource.isZeroValue
	iaDestination.BackUp.isIntegerZeroValue = iaSource.isIntegerZeroValue
	iaDestination.BackUp.precision = iaSource.precision
	iaDestination.BackUp.signVal = iaSource.signVal
	iaDestination.BackUp.decimalSeparator = iaSource.decimalSeparator
	iaDestination.BackUp.thousandsSeparator = iaSource.thousandsSeparator
	iaDestination.BackUp.currencySymbol = iaSource.currencySymbol

	return nil
}

// equal
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
func (iaNeutron *intAryNeutron) equal(
	iAry1 *IntAry,
	validateiAry1 bool,
	iAry2 *IntAry,
	validateiAry2 bool,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

	if iaNeutron.lock == nil {
		iaNeutron.lock = new(sync.Mutex)
	}

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.equal()",
		"")

	if err != nil {
		return false, err
	}

	if iAry1 == nil {

		return false,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'iAry1'",
			}
	}

	if iAry2 == nil {

		return false,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'iAry2'",
			}
	}

	iaElectron := new(intAryElectron)

	if validateiAry1 {

		err = iaElectron.isValidIntAry(iAry1, ePrefix.XCpy("Validating 'iAry1'").String())

		if err != nil {

			return false,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = iaElectron.isValidIntAry(iAry1, ePrefix.XCpy(Validating 'iAry1').String())",
					ErrContext: "Input parameter 'iAry1' is INVALID!\n" +
						"'iAry1' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateiAry2 {

		err = iaElectron.isValidIntAry(iAry2, ePrefix.XCpy("Validating 'iAry1'").String())

		if err != nil {

			return false,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = iaElectron.isValidIntAry(iAry2, ePrefix.XCpy(Validating 'iAry1').String())",
					ErrContext: "Input parameter 'iAry2' is INVALID!\n" +
						"'iAry2' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	iaNanobot := new(intAryNanobot)

	err = iaNanobot.setInternalFlags(iAry1, ePrefix.XCpy("Setting flags 'iAry1'"))

	if err != nil {

		return false,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = iaNanobot.setInternalFlags(iAry1, ePrefix.XCpy(Setting flags 'iAry1'))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = iaNanobot.setInternalFlags(iAry2, ePrefix.XCpy("Setting flags 'iAry2'"))

	if err != nil {

		return false,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = iaNanobot.setInternalFlags(iAry2," +
					"ePrefix.XCpy(Setting flags 'iAry2'))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return new(intAryBoson).dataFieldEqualityTest(iAry1, iAry2), nil
}

// getBigInt
//
//	Returns the current value of this intAry object expressed
//	as a signed integer number of type *big.Int.
func (iaNeutron *intAryNeutron) getBigInt(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	if iaNeutron.lock == nil {
		iaNeutron.lock = new(sync.Mutex)
	}

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.getBigInt()",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if intAry == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'ia'",
			}
	}

	if validateIntAry {

		err = new(intAryElectron).isValidIntAry(
			intAry, ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return big.NewInt(0),
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
						"   intAry, ePrefix.XCpy(Validating 'intAry').String())",
					ErrContext: "Input parameter 'intAry' is INVALID!\n" +
						"'intAry' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	lenIntAry := len(intAry.intAry)

	if lenIntAry != intAry.intAryLen {

		return big.NewInt(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: " if lenIntAry != ia.intAryLen",
				ErrMessage: "Error: The actual length of 'intAry' does not match ia.intAryLen.\n" +
					"This instance of 'ia' is INVALID!",
			}

	}

	if lenIntAry == 0 {

		return big.NewInt(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: " if lenIntAry != ia.intAryLen",
				ErrMessage: "Error: The actual length of 'intAry' is ZERO.\n" +
					"This instance of 'ia' is INVALID!",
			}
	}

	result := big.NewInt(0).SetInt64(0)

	big10 := big.NewInt(0).SetInt64(10)

	for i := 0; i < intAry.intAryLen; i++ {
		result = big.NewInt(0).Mul(result, big10)
		result = big.NewInt(0).Add(result, big.NewInt(0).SetInt64(int64(intAry.intAry[i])))

	}

	if intAry.signVal == -1 {

		result = big.NewInt(0).Neg(result)
	}

	return result, nil
}

// setSign
//
//	Used to change the sign value of the intAry  object passed as
//	input parameter 'intAry'. The new sign value will be set
//	according the input parameter, 'signVal'.
//
//	'signVal' has only two valid values, -1 or +1. If
//	any value other than -1 or +1 is detected, an error
//	will be returned.
func (iaNeutron *intAryNeutron) setSign(
	intAry *IntAry,
	signVal int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaNeutron.lock == nil {
		iaNeutron.lock = new(sync.Mutex)
	}

	iaNeutron.lock.Lock()

	defer iaNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNeutron.equal()",
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

	if signVal != -1 && signVal != 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("'signVal' == %d", signVal),
			ErrMessage: "Error: Input parameter 'signVal' is INVALID.\n" +
				"'signVal' must be either -1 or +1.",
		}

	}

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(intAry, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if intAry.isZeroValue {
		return nil
	}

	intAry.signVal = signVal

	return nil
}
