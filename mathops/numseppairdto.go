package mathops

import ePref "github.com/MikeAustin71/errpref"

// NumericSeparatorPairDto
//
//	Consists of a pair of type NumericSeparatorDto
//	instances.
//
//	InputSeparators are primarily used to parse number
//	strings into numeric values.
//
//	OutputSeparators are used to format number types
//	when they are converted to number strings for
//	display purposes.
type NumericSeparatorPairDto struct {
	InputSeparators  NumericSeparatorDto
	OutputSeparators NumericSeparatorDto
}

// GetInputSeparators
//
//	Implements the IGetNumSeparators interface. This method is
//	called when the encapsulated Numeric Separators for the
//	current instance of NumericSeparatorPairDto are used as
//	'input' Numeric Separators. Input Numeric Separators are
//	primarily used to parse number strings into numeric values.
//
//	Input Parameters
//	================
//
//	None
//
//	Output Parameters
//	=================
//
//	*NumericSeparatorDto
//	  This returned instance of NumericSeparatorDto is designed to
//	  used as input Numeric Separators when parsing number strings
//	  and converting those strings to a numeric value.
//
//	error
//	  If no errors are encountered, this returned error parameter
//	  will be set to 'nil'.
func (numSepPair *NumericSeparatorPairDto) GetInputSeparators() (*NumericSeparatorDto, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"NumericSeparatorPairDto.GetInputSeparators",
		"")

	if err != nil {
		return &NumericSeparatorDto{}, err
	}

	err = new(numSepsDtoElectron).isValidNumStrDto(
		&numSepPair.InputSeparators,
		ePrefix.XCpy("Validating Input Separators"))

	if err != nil {

		return &NumericSeparatorDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "The Input Numeric Separators are invalid!",
				ErrMessage: err.Error(),
			}
	}

	inputNumSeps := &NumericSeparatorDto{}

	inputNumSeps.CurrencySymbol = numSepPair.InputSeparators.CurrencySymbol
	inputNumSeps.DecimalSeparator = numSepPair.InputSeparators.DecimalSeparator
	inputNumSeps.ThousandsSeparator = numSepPair.InputSeparators.ThousandsSeparator

	return inputNumSeps, nil
}

// GetOutputSeparators
//
//	Implements the IGetNumSeparators interface. This method is
//	called when the encapsulated Numeric Separators for the
//	current instance of NumericSeparatorPairDto are used as output
//	Numeric Separators. Output Numeric Separators are primarily
//	used format number types returned from functions. When these
//	number types are later converted to number strings for display
//	purposes, the output formatting for decimal separators,
//	thousands separators and currency symbols will be controlled
//	by these output Numeric Separators.
//
//	Input Parameters
//	================
//
//	None
//
//	Output Parameters
//	=================
//
//	*NumericSeparatorDto
//	  This returned instance of NumericSeparatorDto is designed to
//	  used in formatting numeric types returned by other methods.
//
//	error
//	  If no errors are encountered, this returned error parameter
//	  will be set to 'nil'.
func (numSepPair *NumericSeparatorPairDto) GetOutputSeparators() (*NumericSeparatorDto, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"NumericSeparatorDto.GetOutputSeparators",
		"")

	if err != nil {
		return &NumericSeparatorDto{}, err
	}

	err = new(numSepsDtoElectron).isValidNumStrDto(
		&numSepPair.OutputSeparators,
		ePrefix.XCpy("Validating Output Separators"))

	if err != nil {

		return &NumericSeparatorDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "The Output Numeric Separators are invalid!",
				ErrMessage: err.Error(),
			}
	}

	outputNumSeps := &NumericSeparatorDto{}

	outputNumSeps.CurrencySymbol = numSepPair.OutputSeparators.CurrencySymbol
	outputNumSeps.DecimalSeparator = numSepPair.OutputSeparators.DecimalSeparator
	outputNumSeps.ThousandsSeparator = numSepPair.OutputSeparators.ThousandsSeparator

	return outputNumSeps, nil
}
