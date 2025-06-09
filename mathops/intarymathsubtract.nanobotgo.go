package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryMathSubtractNanobot struct {
	lock sync.Mutex
}

// addToSubtract
//
//	Adds or subtracts two IntAry instances and returns the result
//	in the first IntAry parameter,'ia1'.
//
//	If input parameter 'doAdd' is 'true', this method performs
//	addition. If 'doAdd' is 'false', this method performs
//	subtraction.
func (iaSubtractNano *intAryMathSubtractNanobot) addToSubtract(
	ia1 *IntAry,
	validateIa1 bool,
	ia2 *IntAry,
	validateIa2 bool,
	newSignVal int,
	doAdd bool,
	isZeroResult bool,
	doReverseNums bool,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaSubtractNano.lock.Lock()

	defer iaSubtractNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathSubtractMechanics.addToSubtract()",
		"")

	if err != nil {
		return err
	}

	if ia1 == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia1'",
		}
	}

	if ia2 == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia2'",
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia1,
		"ia1",
		validateIa1,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia1, 'ia1', validateIa1= '%v', ePrefix)", validateIa1),
			ErrContext: "Valiate 'ia1' on Startup",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryUtility).selectIntAryValidation(
		ia2,
		"ia2",
		validateIa2,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"ia2, 'ia2', validateIa2= '%v', ePrefix)", validateIa2),
			ErrContext: "Valiate 'ia2' on Startup",
			ErrMessage: err.Error(),
		}
	}

	if isZeroResult {

		ia1PrecisionUint := uint(ia1.precision)

		nsProfile := NumSepsProfileSelection{
			SourceObjectName:         "ia1",
			OutputNumSepsName:        "numSeps",
			UseDefaultNumSeps:        false,
			SetDefaultNumSepsIfEmpty: true,
			ValidateNumSeps:          false,
			OverrideNumSeps:          NumericSeparatorDto{},
		}

		err = new(intAryQuark).setIntAryToZero(
			ia1,
			nil,
			nsProfile,
			ia1PrecisionUint,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: fmt.Sprintf("err = new(intAryQuark).setIntAryToZero(\n"+
					"ia1, nil, nsProfile, precision= '%v', ePrefix", ia1PrecisionUint),
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	ia1.signVal = newSignVal

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	for j := ia1.intAryLen - 1; j >= 0; j-- {

		if doReverseNums {

			n2 = int(ia1.intAry[j])
			n1 = int(ia2.intAry[j])

		} else {
			n1 = int(ia1.intAry[j])
			n2 = int(ia2.intAry[j])

		}

		if doAdd {
			// doAdd == true
			// Do Addition

			n3 = n1 + n2 + carry

			if n3 > 9 {
				n3 = n1 + n2 + carry - 10
				carry = 1

			} else {
				carry = 0
			}

		} else {
			// doAdd == false
			// Do Subtraction
			n3 = n1 - n2 - carry

			if n3 < 0 {
				n3 = n1 + 10 - n2 - carry
				carry = 1
			} else {
				carry = 0
			}
		}

		ia1.intAry[j] = uint8(n3)

	}

	if carry > 0 {
		ia1.intAry = append([]uint8{1}, ia1.intAry...)
		ia1.intAryLen++
	}

	if ia1.intAry[0] == 0 {

		err = new(intAryElectron).setSignificantDigitIdxs(
			ia1,
			ePrefix.XCpy("Set Digit Indexes 'ia1'"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryElectron).\n" +
					"  setSignificantDigitIdxs( ia1, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		ia1.intAry = ia1.intAry[ia1.firstDigitIdx:]

	}

	// Even if no validation specified, this will set
	// internal flags
	err = new(intAryUtility).selectIntAryValidation(
		ia1,
		"ia1",
		validateResult,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
				"  ia1,\"ia1\", validateResult='%v', ePrefix)", validateResult),
			ErrContext: "Validating Final Result 'ia1' on Exit",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
