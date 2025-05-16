package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type bigIntFixedDecBoson struct {
	lock *sync.Mutex
}

// setNumericSeparatorsToDefaultIfEmpty
//
// If numeric separators were previously set to zero or nil,
// this method will set those numeric separators to the USA
// defaults. This means that the Decimal separator is set
// to ('.'), the Thousands separator is set to (',') and the
// currency symbol is set to '$'.
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that all numeric separators
// are set to valid values.
func (bigIFdBoson *bigIntFixedDecBoson) setNumericSeparatorsToDefaultIfEmpty(
	bigIFxDec *BigIntFixedDecimal,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdBoson.lock == nil {
		bigIFdBoson.lock = new(sync.Mutex)
	}

	bigIFdBoson.lock.Lock()

	defer bigIFdBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecBoson.setNumericSeparatorsToDefaultIfEmpty",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}

		return err
	}

	if bigIFxDec.decimalSeparator == 0 {
		bigIFxDec.decimalSeparator = '.'
	}

	if bigIFxDec.thousandsSeparator == 0 {
		bigIFxDec.thousandsSeparator = ','
	}

	if bigIFxDec.currencySymbol == 0 {
		bigIFxDec.currencySymbol = '$'
	}

	return nil
}

// setNumericSeparatorsDto
//
// Sets the values of numeric separators:
//
//	decimal place separator
//	thousands separator
//	currency symbol
//
// These numeric separators are configured based on values
// transmitted through input parameter 'customSeparators'.
//
// If any of the values contained in input parameter
// 'customSeparators' is set to zero, an error will be returned.
//
//		Input Parameters
//		================
//
//	 bigIFxDec 						*BigIntFixedDecimal
//
//		This instance of BigIntFixedDecimal will be configured
//		with the numeric separator values passed by input parameter
//	 'numSeparators'
//
//		numSeparators					NumericSeparatorDto
//
//		This instance of NumericSeparatorDto holds the decimal
//		separator, thousands seprator and currency symbol which
//		will be used to configure the current instance of
//		BigIntFixedDecimal.
//
//		  type NumericSeparatorDto struct {
//		    DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//		    ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
//		    CurrencySymbol     rune // Currency Symbol
//		  }
func (bigIFdBoson *bigIntFixedDecBoson) setNumericSeparatorsDto(
	bigIFxDec *BigIntFixedDecimal,
	numSeparators NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdBoson.lock == nil {
		bigIFdBoson.lock = new(sync.Mutex)
	}

	bigIFdBoson.lock.Lock()

	defer bigIFdBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecBoson.setNumericSeparatorsDto",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}

		return err
	}

	if numSeparators.DecimalSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: " if customSeparators.DecimalSeparator == 0 {",
			ErrMessage: "Error: Input parameter 'customSeparators.DecimalSeparator' is set to '0' - Invalid rune!",
		}
	}

	if numSeparators.ThousandsSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: " if customSeparators.ThousandsSeparator == 0 {",
			ErrMessage: "Error: Input parameter 'customSeparators.ThousandsSeparator' is set to '0' - Invalid rune!",
		}
	}

	if numSeparators.CurrencySymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: " if customSeparators.CurrencySymbol == 0 {",
			ErrMessage: "Error: Input parameter 'customSeparators.CurrencySymbol' is set to '0' - Invalid rune!",
		}

	}

	bigIFxDec.decimalSeparator = numSeparators.DecimalSeparator

	bigIFxDec.thousandsSeparator = numSeparators.ThousandsSeparator

	bigIFxDec.currencySymbol = numSeparators.CurrencySymbol

	return nil

}

// setNumericSeparators
//
// Used to assign values for the Decimal separators, Thousands
// separators and Currency symbol characters to the instance of
// BigIntFixedDecimal passed as input parameter 'bigIFxDec'.
//
// These numeric separator characters are used to display numeric
// values in number strings.
//
// Different nations and cultures use different symbols to delimit numerical
// values. In the USA and many other countries, a period character ('.') is
// used to delimit integer and fractional digits within a numeric value
// (123.45). Likewise, thousands may be delimited by a comma (','). Currency
// signs very by nationality. For instance, the USA, Canada and several other
// countries use the dollar sign ($) as a currency symbol.
//
// For a list of major world currency symbols see:
//
//	MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//
//	http://www.xe.com/symbols.php
//
// Note: If zero values are submitted as input for separator values, an
// error will be returned.
//
//	USA Examples
//	============
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
func (bigIFdBoson *bigIntFixedDecBoson) setNumericSeparators(
	bigIFxDec *BigIntFixedDecimal,
	decimalSeparator rune,
	thousandsSeparator rune,
	currencySymbol rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdBoson.lock == nil {
		bigIFdBoson.lock = new(sync.Mutex)
	}

	bigIFdBoson.lock.Lock()

	defer bigIFdBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecBoson.setNumericSeparators",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}

		return err
	}

	if decimalSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'decimalSeparator' is set to '0' - Invalid rune!",
		}
	}

	if thousandsSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'thousandsSeparator' is set to '0' - Invalid rune!",
		}
	}

	if currencySymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'currencySymbol' is set to '0' - Invalid rune!",
		}

	}

	bigIFxDec.decimalSeparator = decimalSeparator
	bigIFxDec.thousandsSeparator = thousandsSeparator
	bigIFxDec.currencySymbol = currencySymbol

	return nil
}

// setNumSepSymbol
//
// Sets the number separator symbol on an instance of
// BigIntFixedDecimal passed as input parameter 'bigIFxDec'.
//
//	Input Parameters
//	================
//
//	bigIFxDec           *BigIntFixedDecimal
//
//	The instance of BigIntFixedDecimal which will be configured
//	with the numeric separator character passed as input parameter
//	'numSepSymbol'.
//
//	numSepSymbolType    NumSepSymbolCode
//
//	An integer value enumeration used to desigate the specific
//	numeric separator in 'bigIFxDec' which will set to the value
//	of input parameter 'numSepSymbol'.
//
//	  const (
//	    // DECIMALSYMBOL
//	    // Symbol for the separator character used to
//	    // separate integer and fractional segments of
//	    // a floating point number or curreny value.
//	    DECIMALSYMBOL NumSepSymbolCode = iota
//
//	    // THOUSANDSYMBOL
//	    // Symbol for the separator character used to
//	    // separate thousands in a numeric presentation
//	    // where the numeric value is greater than 999
//	    THOUSANDSYMBOL
//
//	    // CURRENCYSYMBOL
//	    // Symbol for the character used to designate a
//	    // numeric value as currency.
//	    CURRENCYSYMBOL
//	  )
//
//
//	numSepSymbol        rune
//
//	The specific numer separator character which will be
//	transferred to 'bigIFxDec'.
func (bigIFdBoson *bigIntFixedDecBoson) setNumSepSymbol(
	bigIFxDec *BigIntFixedDecimal,
	numSepSymbolType NumSepSymbolCode,
	numSepSymbol rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdBoson.lock == nil {
		bigIFdBoson.lock = new(sync.Mutex)
	}

	bigIFdBoson.lock.Lock()

	defer bigIFdBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecBoson.setNumericSeparators",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}
	}

	if numSepSymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("numSepSymbolType is eqaul to %v separator.",
				numSepSymbolType.String()),
			ErrMessage: "Error: Input parameter 'numSepSymbol' is INVALID!\n" +
				"'numSepSymbol' is empty and has a zero value!",
		}
	}

	switch numSepSymbolType {

	case DECIMALSYMBOL:
		bigIFxDec.decimalSeparator = numSepSymbol
	case THOUSANDSYMBOL:
		bigIFxDec.thousandsSeparator = numSepSymbol
	case CURRENCYSYMBOL:
		bigIFxDec.currencySymbol = numSepSymbol
	default:
		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf(
				"Input parameter 'numSepSymbolType' has an unknown value designator.\n"+
					"Value of 'numSepSymbolType' = %v ", numSepSymbolType),
			ErrMessage: "Error: Input parameter 'numSepSymbolType' is INVALID!",
		}

	}

	return nil
}
