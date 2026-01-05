package mathops

import (
	"fmt"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type intAryMathSubtractMechanics struct {
	lock sync.Mutex
}

// SubtractTotal - This method performs a subtraction operation subtracting
// input parameter 'ia2' from input parameter 'ia1'. The result, or difference,
// is returned through use of a pointer in 'ia1'. This means that the original
// value of 'ia1' will be overwritten and destroyed by the subtraction operation.
//
// The returned 'ia1' IntAry will contain numeric separators (decimal separator,
// thousands separator and currency symbol) from the original 'ia1' IntAry
// instance. This means that the numeric separators contained in the original
// ia1 IntAry will remain unchanged.
func (iaSubtractMech *intAryMathSubtractMechanics) subtractTotal(
	ia1 *IntAry,
	validateIa1 bool,
	ia2 *IntAry,
	validateIa2 bool,
	validateResult bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaSubtractMech.lock.Lock()

	defer iaSubtractMech.lock.Unlock()

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

	//numSeps := ia1.GetNumericSeparatorsDto()
	numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(ia1, true, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(\n" +
				"ia1, setDefaultsIfEmpty=true, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryAtom).setEqualArrayLengths(
		ia1, false, ia2, false, true, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryAtom).setEqualArrayLengths(\n" +
				"  ia1, validateIa=false, ia2, validateIa2=false,\n" +
				"  validateResult=true, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	ia1PrecisionUint := uint(ia1.precision)

	if ia1.isZeroValue && ia2.isZeroValue {

		//ia1.SetIntAryToZero(ia1.GetPrecisionUint())

		nsProfile := NumSepsProfileSelection{
			SourceObjectName:         "ia1",
			OutputNumSepsName:        "numSeps",
			UseDefaultNumSeps:        false,
			SetDefaultNumSepsIfEmpty: true,
			ValidateNumSeps:          false,
			OverrideNumSeps:          numSeps,
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

	compare, err := new(intAryQuark).compareAbsoluteValues(
		ia1, false, ia2, false, ePrefix.XCpy("Comparing ia1 vs ia2"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "compare, err := new(intAryQuark).compareAbsoluteValues(",
			ErrContext: "ia1, validateIntAry1=false, ia2, validateIntAry2=false, ePrefix)",
			ErrMessage: err.Error(),
		}
	}

	isZeroResult := false

	// Largest Value in now in N1 slot
	newSignVal := ia1.signVal

	doAdd := false

	doReverseNums := false

	if compare == 1 {
		// compare == + 1
		// Absolute Value: N1 > N2

		if ia1.signVal == 1 && ia2.signVal == 1 {

			doAdd = false

			newSignVal = 1

		} else if ia1.signVal == -1 && ia2.signVal == 1 {

			doAdd = true

			newSignVal = -1

		} else if ia1.signVal == -1 && ia2.signVal == -1 {

			doAdd = false

			newSignVal = -1

		} else {
			// Must Be ia1.signVal == 1 && ia2.signVal == -1

			doAdd = true

			newSignVal = 1
		}

	} else if compare == -1 {

		// Absolute Values: N2 > N1
		if ia1.signVal == 1 && ia2.signVal == 1 {

			doAdd = false

			doReverseNums = true

			newSignVal = -1

		} else if ia1.signVal == -1 && ia2.signVal == 1 {

			doAdd = true

			newSignVal = -1

		} else if ia1.signVal == -1 && ia2.signVal == -1 {

			doAdd = false

			doReverseNums = true

			newSignVal = 1

		} else {
			// Must Be ia1.signVal == 1 && ia2.signVal == -1

			doAdd = true

			newSignVal = 1
		}

	} else {
		// Must be compare == 0
		// Absolute Values: N1==N2
		if ia1.signVal == 1 && ia2.signVal == 1 {

			doAdd = false

			newSignVal = 1

			isZeroResult = true

		} else if ia1.signVal == -1 && ia2.signVal == 1 {

			doAdd = true

			newSignVal = -1

		} else if ia1.signVal == -1 && ia2.signVal == -1 {

			doAdd = false

			newSignVal = 1

			isZeroResult = true

		} else {
			// Must Be ia1.signVal == 1 && ia2.signVal == -1

			doAdd = true

			newSignVal = 1

		}
	}

	err = new(intAryMathSubtractNanobot).addToSubtract(ia1, false, ia2, false, newSignVal, doAdd, isZeroResult, doReverseNums, true, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryMathSubtractNanobot).addToSubtract(\n" +
				"  ia1, validateIa1=false, ia2, validateIa2=false, newSignVal,\n" +
				"  doAdd, isZeroResult, doReverseNums, validateReslt=true, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(intAryPhoton).setNumericSeparatorsDto(
		ia1, numSeps, true, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "new(intAryPhoton).setNumericSeparatorsDto(\n" +
				"  ia1, numSeps, false, ePrefix))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
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
