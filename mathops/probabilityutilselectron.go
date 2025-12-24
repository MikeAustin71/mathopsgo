package mathops

import (
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type probabilityUtilsElectron struct {
	lock *sync.Mutex
}

// setNumericSeparators
//
//	This method is used to set and configure the Numeric Separators
//	for an instance of type Probability passed as an input . The
//	values contained in input parameter 'numSeps' will be used to
//	configure the type Probability intenal member variable
//	'prob.NumSeps'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators consist of the decimal separator, thousands
//	seprator, and currency symbol used to cofigure custom number
//	strings.
//
//	Types returned by various Probability operational functions will
//	be configued with Numeric Separator values configured through this
//	method.
//
//	If the input parameter 'newNumSeps' is judged to be invalid, and
//	input parameter 'setNumSepDefaultUSAIfEmpty' is set to 'false', an
//	error will be returned.
//
//		If the input parameter 'newNumSeps' is judged to be invalid, and
//	input parameter 'setNumSepDefaultUSAIfEmpty' is set to 'true',
//	prob.NumSeps will be set to default USA Numeric Separators.
//	USA default values are listed as follows:
//	  decimal separator ('.')
//	  thousands separator (',')
//	  currency symbol ('$').
//
//	Input Parameters
//	================
//
//	prob                         *Probability
//	 Pointer to an instance of type Probability
//
//	newNumSeps                   NumericSeparatorDto
//	  The Numeric Separators contained in 'newNumSeps' will be
//	  copied to 'prob.NumSeps'. This will effectively configure
//	  the Probability instance (prob) with new Numeric Separators.
//
//	setNumSepDefaultUSAIfEmpty   bool
//	  If this boolean value is set to 'true', any empty values
//	  contained in 'newNumSeps' will be automatically set to USA
//	  default Numeric Separators.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//	  This object encapsulates an error prefix string
//	  which is included in all returned error
//	  messages. Usually, it contains the name of the
//	  calling method or methods listed as a function
//	  chain.
//
//	  If no error prefix information is needed, set
//	  this parameter to 'nil'.
//
//	  Type ErrPrefixDto is included in the 'errpref'
//	  software package:
//	    "github.com/MikeAustin71/errpref".
//
//	Return Values
//	=============
//
//	error
//	  If an error is encountered during processing, this returend
//	  'error' object will be configured with an appropriate error
//	  message.
//
//	  If the returned error object is set to 'nil', it signals that
//	  the Probability Numeric Seprators (prob.NumSeps) were successfully
//	  configured.
func (probUtilsElectron *probabilityUtilsElectron) setNumericSeparators(
	prob *Probability,
	newNumSeps NumericSeparatorDto,
	setNumSepDefaultUSAIfEmpty bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if probUtilsElectron.lock == nil {
		probUtilsElectron.lock = new(sync.Mutex)
	}

	probUtilsElectron.lock.Lock()

	defer probUtilsElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"probabilityUtilsElectron.setNumericSeparators()",
		"")

	if err != nil {
		return err
	}

	if prob == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'prob'",
		}
	}

	if setNumSepDefaultUSAIfEmpty {

		newNumSeps.SetDefaultsIfEmpty()

	}

	err = newNumSeps.IsValid("Validating 'newNumSeps'")

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = newNumSeps.IsValid(\"Validating 'newNumSeps'\")",
			ErrContext: "Input Parameter 'newNumSeps' is invalid",
			ErrMessage: err.Error(),
		}
	}

	prob.NumSeps.DecimalSeparator = newNumSeps.DecimalSeparator

	prob.NumSeps.ThousandsSeparator = newNumSeps.ThousandsSeparator

	prob.NumSeps.CurrencySymbol = newNumSeps.CurrencySymbol

	return nil
}
