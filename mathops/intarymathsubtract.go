package mathops

import (
	"fmt"

	ePref "github.com/MikeAustin71/errpref"
)

type IntAryMathSubtract struct {
	Input  IntAryPair
	Result IntAry
}

// Subtract - Receives two input parameters of type *IntAry. The second parameter
// 'subtrahend' is subtracted from the first parameter, 'minuend'. The result, or
// difference, is returned as a new IntAry instance.
//
// 					'minuend' - 'subtrahend' = difference or result
//
// The two input parameters, 'miunuend' and 'subtrahend', are assumed to be valid
// IntAry instances properly initialized. No validation is performed on 'minuend'
// or 'subtrahend'.
//
// The result of this subtraction operation is returned as an IntAry instance. This
// IntAry will contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter 'minuend'.

func (iaSubtract *IntAryMathSubtract) Subtract(minuend *IntAry, subtrahend *IntAry) (IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathSubtract.Subtract",
		"")

	if err != nil {
		return IntAry{}, err
	}

	//ia3 := minuend.CopyOut()

	ia3 := new(intAryElectron).newIntAry()

	err = new(intAryProton).copy(
		&ia3, minuend, true, false,
		ePrefix.XCpy("minuend->ia3"))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryProton).copy(\n" +
					"  &ia3, minuend, validateMinuend=true,\n" +
					"  copyToBackup=false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(intAryMathSubtractMechanics).subtractTotal(&ia3, false, subtrahend, true, true, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryMathSubtractMechanics).subtractTotal(\n" +
					"  &ia3, validateIa3=false, subtrahend, validateSubtrahend=true,\n" +
					"  validateResult=true, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return ia3, nil
}

// SubtractTotal
//
//	This method performs a subtraction operation subtracting input
//	parameter 'ia2' from input parameter 'ia1'. The result, or
//	difference, is returned through use of a pointer in 'ia1'. This
//	means that the original value of 'ia1' will be overwritten and
//	destroyed by the subtraction operation.
//
//	The returned 'ia1' IntAry will contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from the original 'ia1' IntAry 'ia1' instance. This
//	means that the numeric separators contained in the original
//	'ia1' IntAry will remain unchanged.
//
//	If input paramter 'validateIa1' is set to true, IntAry object
//	'ia1' will be subjected to validation testing.
//
//	If input paramter 'validateIa2' is set to true, IntAry object
//	'ia2' will be subjected to validation testing.
//
//	If input paramter 'validateIa2' is set to true, the final
//	result ('ia1'), after performing the subtraction operation,
//	will be subjected to validation testing.
func (iaSubtract *IntAryMathSubtract) SubtractTotal(
	ia1 *IntAry,
	validateIa1 bool,
	ia2 *IntAry,
	validateIa2 bool,
	validateResult bool) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathSubtract.SubtractTotal",
		"")

	if err != nil {
		return err
	}

	err = new(intAryMathSubtractMechanics).subtractTotal(
		ia1, validateIa1, ia2, true, true, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: fmt.Sprintf("err = new(intAryMathSubtractMechanics).subtractTotal(\n"+
				"ia1, validateIa1=%v, ia2, validateIa2=%v, validateResult=%v, ePrefix)",
				validateIa1, validateIa2, validateResult),
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
