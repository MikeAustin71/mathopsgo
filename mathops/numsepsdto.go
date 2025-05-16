package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

// NumericSeparatorDto
// Numeric Separator Data Transfer Object
// Used to transmit symbols used for decimal point,
// thousands separator and currency symbol.
type NumericSeparatorDto struct {
	DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
	ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
	CurrencySymbol     rune // Currency Symbol
}

// Equal - Compares two NumericSeparatorDto's and returns 'true' if they
// are equivalent.
func (numSep *NumericSeparatorDto) Equal(numSep2 NumericSeparatorDto) bool {

	if numSep.DecimalSeparator != numSep2.DecimalSeparator {
		return false
	}

	if numSep.ThousandsSeparator != numSep2.ThousandsSeparator {
		return false
	}

	if numSep.CurrencySymbol != numSep2.CurrencySymbol {
		return false
	}

	return true
}

// IsValid
//
// This method will test the current instance of
// NumericSeparatorDto to determine if all member variables are
// valid. If any member elements are determined to be invalid,
// an error will be returned.
//
// If all member elements of the current NumericSeparatorDto
// instance are determined to be valid, the method will return
// a 'nil' value.
func (numSep *NumericSeparatorDto) IsValid(callingMethodName string) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	if len(callingMethodName) > 0 {
		callingMethodName = "NumericSeparatorDto.IsValid" + "\n" + callingMethodName
	} else {
		callingMethodName = "NumericSeparatorDto.IsValid"
	}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		callingMethodName,
		"")

	if err != nil {
		return err
	}

	if numSep.DecimalSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numSep.DecimalSeparator == 0 {",
			ErrMessage: "Error: NumericSeparatorDto member element 'numSep.DecimalSeparator' is INVALID!",
		}
	}

	if numSep.ThousandsSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numSep.ThousandsSeparator == 0 {",
			ErrMessage: "Error: NumericSeparatorDto member element 'numSep.ThousandsSeparator' is INVALID!",
		}
	}

	if numSep.CurrencySymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numSep.CurrencySymbol == 0 {",
			ErrMessage: "Error: CurrencySymbol member element 'numSep.ThousandsSeparator' is INVALID!",
		}
	}

	return nil
}

// New - Returns a new instance of NumericSeparatorDto. The
// rune values are automatically set to USA defaults.
func (numSep *NumericSeparatorDto) New() NumericSeparatorDto {

	n2 := NumericSeparatorDto{}

	n2.SetDefaultsIfEmpty()

	return n2
}

// SetDefaultsIfEmpty - If any of the NumericSeparatorDTo rune values
// are zero, this method will set those elements to USA default values.
func (numSep *NumericSeparatorDto) SetDefaultsIfEmpty() {

	if numSep.DecimalSeparator == 0 {
		numSep.DecimalSeparator = '.'
	}

	if numSep.ThousandsSeparator == 0 {
		numSep.ThousandsSeparator = ','
	}

	if numSep.CurrencySymbol == 0 {
		numSep.CurrencySymbol = '$'
	}
}

// String - Provides a formatted listing of the contents from the current
// NumericSeparatorDto instance.
func (numSep *NumericSeparatorDto) String() string {
	return fmt.Sprintf("Decimal Separator: %q  Thousands Separator: %q  Currency Symbol: %q",
		numSep.DecimalSeparator, numSep.ThousandsSeparator, numSep.CurrencySymbol)
}
