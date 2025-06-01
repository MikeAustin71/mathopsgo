package mathops

import (
	"bytes"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryAtom struct {
	lock *sync.Mutex
}

// getRawNumStr
//
//	Returns the current value of the input parameter IntAry object
//	('intAry') as a number string.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. These digits
//	are formatted in a way that facilitates conversion to a
//	corresponding numeric value.
//
//	The number string returned by this method will contain a
//	decimal separator to separate integer and fractional
//	components of the numeric value. The returned number string
//	will not contain 'thousands' separators or 'currency' symbols.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaAtom *intAryAtom) getRawNumStr(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	if iaAtom.lock == nil {
		iaAtom.lock = new(sync.Mutex)
	}

	iaAtom.lock.Lock()

	defer iaAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryAtom.getRawNumStr()",
		"")

	if err != nil {
		return "", err
	}

	if intAry == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	if validateIntAry {

		err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
						"  intAry, ePrefix.XCpy(Validating 'intAry').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	// We now know that 'intAry' has valid NumSeps

	err = new(intAryNanobot).setInternalFlags(intAry, ePrefix.XCpy("Set intAry Flags"))

	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer

	if intAry.signVal < 0 {
		buffer.WriteRune('-')
	}

	intLen := intAry.intAryLen - intAry.precision

	for i := 0; i < intLen; i++ {
		buffer.WriteRune(rune(intAry.intAry[i] + 48))
	}

	if intAry.precision > 0 {
		buffer.WriteRune(intAry.decimalSeparator)

		for j := 0; j < intAry.precision; j++ {
			buffer.WriteRune(rune(intAry.intAry[intLen] + 48))
			intLen++
		}

	}

	return buffer.String(), nil
}

// OptimizeIntArrayLen
//
//	 Eliminates Leading zeros from the front or integer portion
//	 of the integer string.
//
//	 If parameter 'optimizeFracDigits' is set equal to 'true',
//	 trailing zeros to the right of the decimal place will also be
//	 eliminated.
//
//		Validation Testing
//		==================
//
//		If input parameter 'validateIntAry' is set to true, this
//		method will subject 'intAry' to validation tests.
func (iaAtom *intAryAtom) optimizeIntArrayLen(
	intAry *IntAry,
	validateIntAry bool,
	optimizeFracDigits bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaAtom.lock == nil {
		iaAtom.lock = new(sync.Mutex)
	}

	iaAtom.lock.Lock()

	defer iaAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryAtom.optimizeIntArrayLen()",
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

	if validateIntAry {

		err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
					"  intAry, ePrefix.XCpy(Validating 'intAry').String())",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	// We now know that 'intAry' is valid

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix.XCpy("Set intAry Flags"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
				"intAry, ePrefix.XCpy(Set intAry Flags))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if intAry.isZeroValue {
		return nil
	}

	integerLen :=
		intAry.intAryLen - intAry.precision - intAry.firstDigitIdx

	if optimizeFracDigits {

		intAry.intAry = intAry.intAry[intAry.firstDigitIdx : intAry.lastDigitIdx+1]
		intAry.intAryLen = intAry.lastDigitIdx - intAry.firstDigitIdx + 1

	} else {

		intAry.intAry = intAry.intAry[intAry.firstDigitIdx:]
		intAry.intAryLen = intAry.intAryLen - intAry.firstDigitIdx
	}

	intAry.precision = intAry.intAryLen - integerLen

	return nil
}
