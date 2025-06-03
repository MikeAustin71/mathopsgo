package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryMathAddUtility struct {
	lock *sync.Mutex
}

// addToSubtract
//
//	Adds or subtracts two IntAry instances and returns the result
//	in the first IntAry parameter.
func (iaAddUtil *intAryMathAddUtility) addToSubtract(
	ia1 *IntAry,
	validateIa1 bool,
	ia2 *IntAry,
	validateIa2 bool,
	newSignVal int,
	doAdd bool,
	isZeroResult bool,
	doReverseNums bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaAddUtil.lock == nil {
		iaAddUtil.lock = new(sync.Mutex)
	}

	iaAddUtil.lock.Lock()

	defer iaAddUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathAddUtility.addToSubtract",
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

	if validateIa1 {

		err = ia1.IsValid(ePrefix.XCpy("Validating 'ia1'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia1.IsValid(ePrefix.XCpy(Validating 'ia1').String())",
				ErrContext: "Input parameter 'ia1' is INVALID!\n" +
					"'ia1' FAILED Validatin Tests.",
				ErrMessage: err.Error(),
			}
		}

	} else {

		err = ia1.SetInternalFlags()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia1.SetInternalFlags()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if validateIa2 {

		err = ia2.IsValid(ePrefix.XCpy("Validating 'ia2'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia2.IsValid(ePrefix.XCpy(Validating 'ia2').String())",
				ErrContext: "Input parameter 'ia2' is INVALID!\n" +
					"'ia2' FAILED Validatin Tests.",
				ErrMessage: err.Error(),
			}
		}

	} else {

		err = ia2.SetInternalFlags()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia2.SetInternalFlags()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	var ia1PrecisionUint uint

	if isZeroResult {

		ia1PrecisionUint, err = ia1.GetPrecisionUint()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "ia1PrecisionUint, err = ia1.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		err = ia1.SetIntAryToZero(ia1PrecisionUint)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia1.SetIntAryToZero(ia1PrecisionUint)",
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

		err = ia1.SetSignificantDigitIdxs()

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = ia1.SetSignificantDigitIdxs()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		ia1.intAry = ia1.intAry[ia1.firstDigitIdx:]
	}

	err = ia1.SetInternalFlags()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = ia1.SetInternalFlags()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
