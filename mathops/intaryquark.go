package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryQuark struct {
	lock *sync.Mutex
}

// compareAbsoluteValues
//
//	Compares the absolute values of two IntAry instances.
//
//	Returns:
//	 0  = Current IntAry value is equal to the passed IntAry value.
//	 1  = Current IntAry value is greater than the passed IntAry value.
//	-1  = Current IntAry value is less than the passed IntAry value.
func (iaQuark *intAryQuark) compareAbsoluteValues(
	intAry1 *IntAry,
	validateIntAry1 bool,
	intAry2 *IntAry,
	validateIntAry2 bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if iaQuark.lock == nil {
		iaQuark.lock = new(sync.Mutex)
	}

	iaQuark.lock.Lock()

	defer iaQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryQuark.compareAbsoluteValues()",
		"")

	if err != nil {
		return -1, err
	}

	if intAry1 == nil {

		return -1,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry1'",
			}
	}

	if intAry2 == nil {

		return -1,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry2'",
			}
	}

	iAryElectron := new(intAryElectron)
	iAryNanobot := new(intAryNanobot)

	if validateIntAry1 {

		err = iAryElectron.isValidIntAry(intAry1, ePrefix.XCpy("Validating 'intAry1'").String())

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryElectron.isValidIntAry(\n" +
						"  intAry1, ePrefix.XCpy(Validating 'intAry1').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	} else {

		err = iAryNanobot.setInternalFlags(
			intAry1, ePrefix.XCpy("Setting 'intAry1' Flags"))

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryNanobot.setInternalFlags(\n" +
						"  intAry1, ePrefix.XCpy(Setting 'intAry1' Flags))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	if validateIntAry2 {

		err = iAryElectron.isValidIntAry(intAry2, ePrefix.XCpy("Validating 'intAry2'").String())

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryElectron.isValidIntAry(\n" +
						"  intAry2, ePrefix.XCpy(Validating 'intAry2').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	} else {

		err = iAryNanobot.setInternalFlags(
			intAry2, ePrefix.XCpy("Setting 'intAry2' Flags"))

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryNanobot.setInternalFlags(\n" +
						"  intAry2, ePrefix.XCpy(Setting 'intAry2' Flags))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	err = iAryElectron.setIntAryLength(intAry1, ePrefix.XCpy("Setting 'ia' IntAry Length"))

	if err != nil {

		return -1,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = iAryElectron.setIntAryLength(\n" +
					"ia, ePrefix.XCpy(Setting 'ia' IntAry Length))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = iAryElectron.setIntAryLength(intAry2, ePrefix.XCpy("Setting 'iAry2' IntAry Length"))

	if err != nil {

		return -1,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = iAryElectron.setIntAryLength(\n" +
					"iAry2, ePrefix.XCpy(Setting 'iAry2' IntAry Length))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// Not needed. Validation sets internal flags
	//iAry2.SetIsZeroValue()
	//ia.SetIsZeroValue()

	if intAry1.isZeroValue && intAry2.isZeroValue {
		return 0, nil
	}

	iaIntLen := intAry1.intAryLen - intAry1.precision

	iAry2IntLen := intAry2.intAryLen - intAry2.precision

	// Integer Lengths are Equal
	if iaIntLen == iAry2IntLen {
		for i := 0; i < iaIntLen; i++ {
			if intAry1.intAry[i] > intAry2.intAry[i] {
				return 1, nil
			}

			if intAry2.intAry[i] > intAry1.intAry[i] {
				return -1, nil
			}
		}
	}

	deltaStartIdx := 0

	// ia Integer Length is Greater than IAry2 Integer Length
	if iaIntLen > iAry2IntLen {
		deltaStartIdx = iaIntLen - iAry2IntLen

		for j := 0; j < iaIntLen; j++ {

			if j < deltaStartIdx {

				if intAry1.intAry[j] > 0 {
					return 1, nil
				}

			} else {
				// 'i' must be >= deltaStartIdx

				if intAry1.intAry[j] > intAry2.intAry[j-deltaStartIdx] {
					return 1, nil
				}

				if intAry2.intAry[j-deltaStartIdx] > intAry1.intAry[j] {
					return -1, nil
				}
			}
		}
	}

	// iAry2 Integer Length is Greater Than ia Integer Length
	if iAry2IntLen > iaIntLen {
		deltaStartIdx = iAry2IntLen - iaIntLen

		for k := 0; k < iAry2IntLen; k++ {

			if k < deltaStartIdx {
				if intAry2.intAry[k] > 0 {
					return -1, nil
				}

			} else {
				// 'i' must be >= deltaStartIdx

				if intAry2.intAry[k] > intAry1.intAry[k-deltaStartIdx] {
					return -1, nil
				}

				if intAry1.intAry[k-deltaStartIdx] > intAry2.intAry[k] {
					return 1, nil
				}
			}
		}
	}

	// If precision is zero, the intAry's are equivalent
	if intAry1.precision == 0 && intAry2.precision == 0 {
		return 0, nil
	}

	// Integer Values are Equivalent. Now test
	// digits to the right of the decimal point.

	// Test fractional digits to right of decimal point
	iaFracIdx := iaIntLen
	iAry2FracIdx := iAry2IntLen
	// Test for case of Equal precision
	if intAry1.precision == intAry2.precision {
		for m := 0; m < intAry1.precision; m++ {

			if intAry1.intAry[iaFracIdx] > intAry2.intAry[iAry2FracIdx] {
				return 1, nil
			}

			if intAry2.intAry[iAry2FracIdx] > intAry1.intAry[iaFracIdx] {
				return -1, nil
			}

			iaFracIdx++
			iAry2FracIdx++
		}
	}

	iaFracIdx = iaIntLen
	iAry2FracIdx = iAry2IntLen
	// Test for case where ia precision Greater than iAry2 precision
	if intAry1.precision > intAry2.precision {

		for i := 0; i < intAry1.precision; i++ {

			if i < intAry2.precision {

				if intAry1.intAry[iaFracIdx] > intAry2.intAry[iAry2FracIdx] {
					return 1, nil
				}

				if intAry2.intAry[iAry2FracIdx] > intAry1.intAry[iaFracIdx] {
					return -1, nil
				}

				iaFracIdx++
				iAry2FracIdx++

			} else {
				if intAry1.intAry[iaFracIdx] > 0 {
					return 1, nil
				}

				iaFracIdx++
			}
		}
	}

	iaFracIdx = iaIntLen
	iAry2FracIdx = iAry2IntLen
	// Test for case where iAry2 precision Greater than ia precision
	if intAry2.precision > intAry1.precision {

		for i := 0; i < intAry2.precision; i++ {

			if i < intAry1.precision {

				if intAry1.intAry[iaFracIdx] > intAry2.intAry[iAry2FracIdx] {
					return 1, nil
				}

				if intAry2.intAry[iAry2FracIdx] > intAry1.intAry[iaFracIdx] {
					return -1, nil
				}

				iaFracIdx++
				iAry2FracIdx++

			} else {
				if intAry2.intAry[iAry2FracIdx] > 0 {
					return -1, nil
				}

				iAry2FracIdx++
			}
		}

	}

	// The two absolute numeric values must be equal
	return 0, nil
}

// setIntAryToZero
//
//		Sets the value of the intAry object to zero ('0').
//
//	 Note: Input paramter 'numSeps' will be subjected to validation
//	 testing.
func (iaQuark *intAryQuark) setIntAryToZero(
	intAry *IntAry,
	precision uint,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaQuark.lock == nil {
		iaQuark.lock = new(sync.Mutex)
	}

	iaQuark.lock.Lock()

	defer iaQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryQuark.setIntAryToZero()",
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

	err = numSeps.IsValid(ePrefix.XCpy("Validating input param 'numSeps'").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(Validating input param 'numSeps').String())",
			ErrContext: "Input parameter 'numSeps' is INVALID!\n" +
				"'numSeps' FAILED validation testing.",
			ErrMessage: err.Error(),
		}
	}

	intAry.intAryLen = 1 + int(precision)

	intAry.precision = int(precision)

	intAry.intAry = make([]uint8, intAry.intAryLen)

	intAry.signVal = 1

	err = new(intAryPhoton).setNumericSeparatorsDto(
		intAry, numSeps, false, ePrefix.XCpy("Setting 'intAry'"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryMechanics).\n" +
				"  setNumericSeparatorsDto(intAry, ePrefix.String())",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryNanobot).setInternalFlags(
		intAry, ePrefix.XCpy("Setting 'intAry' Flags"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryNanobot).\n" +
				"  setInternalFlags(intAry, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
