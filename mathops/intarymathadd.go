package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

type IntAryMathAdd struct {
	Input  IntAryPair
	Result IntAry
}

// AddManyArray
//
//		Adds the contents of the []IntAry, 'iaMany' to the value of
//		'total'.
//
//		Validation Testing
//		==================
//
//	 No validation testing is performed on 'total' or 'iaMany'.
//
//		Numeric Separators
//		==================
//
//		The original 'total' numeric separators will remain unchanged.
//		Numeric separators consist of the decimal separator, thousands
//		separator and currency symbol.
func (iaAdd *IntAryMathAdd) AddManyArray(total *IntAry, iaMany []IntAry) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathAdd.AddManyArray()",
		"")

	if err != nil {
		return err
	}

	lAry := len(iaMany)

	iaAddMech := new(intAryMathAddMechanics)

	for i := 0; i < lAry; i++ {

		err = iaAddMech.runTotal(total, false, &iaMany[i], false, ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: fmt.Sprintf("err = iaAddMech.runTotal(total, false, &iaMany[%d], false, ePrefix)", i),
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	return nil
}

// AddMany
//
//	 Adds multiple IntAry instances and returns the result in input
//	 parameter 'total'.
//
//		Validation Testing
//		==================
//
//	 No validation testing is performed on 'total' or 'iaMany'.
//
//		Numeric Separators
//		==================
//
//	 The original 'total' numeric separators will remain unchanged.
//	 Numeric separators consist of the decimal separator, thousands
//	 separator and currency symbol.
func (iaAdd *IntAryMathAdd) AddMany(total *IntAry, iaMany ...*IntAry) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathAdd.AddMany()",
		"")

	if err != nil {
		return err
	}

	iaAddMech := new(intAryMathAddMechanics)

	for idx, iAry := range iaMany {

		err = iaAddMech.runTotal(total, false, iAry, false, ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = iaAddMech.runTotal(total, false, iAry, false, ePrefix)",
				ErrContext: fmt.Sprintf("'iaMany' series Cycle Index= '%d'", idx),
				ErrMessage: err.Error(),
			}
		}
	}

	return nil
}

// Add
//
//	Adds two IntAry instances and returns the result as a new
//	IntAry Instance.
//
//	The return value 'IntAry' will contain the numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from input parameter 'ia1'.
func (iaAdd *IntAryMathAdd) Add(ia1 *IntAry, ia2 *IntAry) (IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathAdd.AddMany()",
		"")

	if err != nil {
		return IntAry{}, err
	}

	ia3, err := ia1.CopyOut()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "ia3, err := ia1.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(intAryMathAddMechanics).
		runTotal(&ia3, false, ia2, false, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryMathAddMechanics).\n" +
					"  runTotal(&ia3, false, ia2, false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return ia3, nil
}

// RunTotal
//
//	Adds to IntAry input parameters and returns the results in the
//	first parameter, 'ia'.
//
//	Validation Testing
//	==================
//
//	This method will NOT perform validation tests on input
//	parameter 'ia' and 'ia2'.
//
//	Numeric Separators
//	==================
//
//	The returned addition 'result' in 'ia' will contain numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) copied from the original 'ia' structure. In other
//	words, the 'ia' numeric separators will remain unchanged.
func (iaAdd *IntAryMathAdd) RunTotal(ia, ia2 *IntAry) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathAdd.RunTotal()",
		"")

	if err != nil {
		return err
	}

	return new(intAryMathAddMechanics).runTotal(
		ia, false, ia2, false, ePrefix)
}
