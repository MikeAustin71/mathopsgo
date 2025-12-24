package mathops

import ePref "github.com/MikeAustin71/errpref"

// SetNumericSeparators
//
//	This method is used to set and configure the Numeric Separator
//	for this instance of type Probability. The values contained in
//	input parameter 'numSeps' will be used to configure the type
//	Probability intenal member variable 'NumSeps'.
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
//	If the input parameter 'numSeps' is judged to be invalid, an error
//	will be returned.
func (prob *Probability) SetNumericSeparators(numSeps NumericSeparatorDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"Probability.CombinationsNoRepsBigInt",
		"")

	if err != nil {
		return err
	}

	return new(probabilityUtilsElectron).setNumericSeparators(
		prob,
		numSeps,
		false,
		ePrefix)
}
