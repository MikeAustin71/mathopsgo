package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryPhoton struct {
	lock *sync.Mutex
}

// getNumericSeparatorsDto
//
//	 Receives a pointer to IntAry and extracts the Numeric
//	 Separators. These separators are consolidated and returned as
//	 a NumericSeparatorDto structure containing the character or
//	 'rune' values for decimal point separator, thousands
//	 separator and currency symbol.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//		them into numeric values.
func (iaPhoton *intAryPhoton) getNumericSeparatorsDto(
	intAry *IntAry,
	errPrefDto *ePref.ErrPrefixDto) (NumericSeparatorDto, error) {

	if iaPhoton.lock == nil {
		iaPhoton.lock = new(sync.Mutex)
	}

	iaPhoton.lock.Lock()

	defer iaPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryPhoton.getNumericSeparatorsDto()",
		"")

	if err != nil {
		return NumericSeparatorDto{}, err
	}

	if intAry == nil {

		return NumericSeparatorDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	numSeps := NumericSeparatorDto{}
	numSeps.DecimalSeparator = intAry.GetDecimalSeparator()
	numSeps.ThousandsSeparator = intAry.GetThousandsSeparator()
	numSeps.CurrencySymbol = intAry.GetCurrencySymbol()

	return numSeps, nil
}

// CompareSignedValues
//
//	Compares two IntAry signed numeric values.
//
//	Returns:
//	 0  = Current IntAry value is equal to the passed IntAry value.
//	 1  = Current IntAry value is greater than the passed IntAry value.
//	-1  = Current IntAry value is less than the passed IntAry value.
func (iaPhoton *intAryPhoton) compareSignedValues(
	intAry1 *IntAry,
	validateIntAry1 bool,
	intAry2 *IntAry,
	validateIntAry2 bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if iaPhoton.lock == nil {
		iaPhoton.lock = new(sync.Mutex)
	}

	iaPhoton.lock.Lock()

	defer iaPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryPhoton.getNumericSeparatorsDto()",
		"")

	if err != nil {
		return -1, err
	}

	if intAry1 == nil {

		return -1,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry1'",
			}
	}

	if intAry2 == nil {

		return -1,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry2'",
			}
	}

	var doTier2ValidationIntAry1, doTier2ValidationIntAry2 bool

	if validateIntAry1 {
		doTier2ValidationIntAry1 = false
	} else {
		doTier2ValidationIntAry1 = true
	}

	if validateIntAry2 {
		doTier2ValidationIntAry2 = false
	} else {
		doTier2ValidationIntAry2 = true
	}

	iAryElectron := new(intAryElectron)

	iAryNanobot := new(intAryNanobot)

	if validateIntAry1 {

		err = iAryElectron.isValidIntAry(intAry1, ePrefix.XCpy("Validating 'intAry1'").String())

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryElectron.isValidIntAry(\n" +
						"  intAry1, ePrefix.XCpy(Validating 'intAry1').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	} else {

		err = iAryNanobot.setInternalFlags(
			intAry1, ePrefix.XCpy("Setting 'intAry1' Flags"))

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryNanobot.setInternalFlags(\n" +
						"  intAry1, ePrefix.XCpy(Setting 'intAry1' Flags))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateIntAry2 {

		err = iAryElectron.isValidIntAry(intAry2, ePrefix.XCpy("Validating 'intAry2'").String())

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryElectron.isValidIntAry(\n" +
						"  intAry2, ePrefix.XCpy(Validating 'intAry2').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	} else {

		err = iAryNanobot.setInternalFlags(
			intAry2, ePrefix.XCpy("Setting 'intAry2' Flags"))

		if err != nil {

			return -1,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = iAryNanobot.setInternalFlags(\n" +
						"  intAry2, ePrefix.XCpy(Setting 'intAry2' Flags))",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	iCompare, err := new(intAryQuark).compareAbsoluteValues(
		intAry1, doTier2ValidationIntAry1, intAry2, doTier2ValidationIntAry2, ePrefix)

	if intAry1.isZeroValue && intAry2.isZeroValue {
		return 0, nil
	}

	if intAry1.signVal != intAry2.signVal {

		if intAry1.signVal == 1 {

			return 1, nil

		} else {

			return -1, nil

		}
	}

	// Must be ia.signVal == iAry2.signVal

	if intAry1.signVal == 1 {
		return iCompare, nil
	}

	// Must be ia.signVal && iAry2.signVal == -1

	return iCompare * -1, nil
}

// setNumericSeparators
//
//	Used to assign values for the Decimal and Thousands separators
//	as well as the Currency Symbol to be used in displaying the
//	current number string.
//
//	Different nations and cultures use different symbols to delimit
//	numerical values. In the USA and many other countries, a period
//	character ('.') is used to delimit integer and fractional digits
//	within a numeric value (123.45). Likewise, thousands may be
//	delimited by a comma (','). Currency signs very by nationality.
//	For instance, the USA, Canada and several other countries use
//	the dollar sign ($) as a currency symbol.
//
//	For a list of major world currency symbols see:
//
//	  MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//
//	  http://www.xe.com/symbols.php
//
//	Note: If zero values are submitted as input for separator values,
//	those values will default to USA standards.
//
//	USA Examples
//	============
//
//	Decimal Separator period ('.')    = 123.456
//	Thousands Separator comma (',')   = 1,000,000,000
//	Currency Symbol dollar sign ('$') = $123
func (iaPhoton *intAryPhoton) setNumericSeparators(
	intAry *IntAry,
	decimalSeparator,
	thousandsSeparator,
	currencySymbol rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaPhoton.lock == nil {
		iaPhoton.lock = new(sync.Mutex)
	}

	iaPhoton.lock.Lock()

	defer iaPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryPhoton.setNumericSeparators()",
		"")

	if err != nil {
		return err
	}

	if intAry == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intAry'",
		}
	}

	var errMsg string

	if decimalSeparator == 0 {
		errMsg += "Input parameter 'decimalSeparator' is empty.\n"
	}

	if thousandsSeparator == 0 {
		errMsg += "Input parameter 'thousandsSeparator' is empty.\n"
	}

	if currencySymbol == 0 {
		errMsg += "Input parameter 'currencySymbol' is empty.\n"
	}

	if len(errMsg) > 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error:\n%v", errMsg),
		}
	}

	intAry.decimalSeparator = decimalSeparator

	intAry.thousandsSeparator = thousandsSeparator

	intAry.currencySymbol = currencySymbol

	return nil
}

// setNumericSeparatorsDto
//
// Sets the values of numeric separators:
//
//		  decimal point separator
//		  thousands separator
//		  currency symbol
//
//		These vales are based on data transmitted through input
//		parameter 'customSeparators' of type NumericSeparatorDto.
//
//		If input parameter 'validateNumericSeparators' is set to true,
//	 'customSeparators' will be subjected to validatin testing.
func (iaPhoton *intAryPhoton) setNumericSeparatorsDto(
	intAry *IntAry,
	customSeparators NumericSeparatorDto,
	validateNumericSeparators bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaPhoton.lock == nil {
		iaPhoton.lock = new(sync.Mutex)
	}

	iaPhoton.lock.Lock()

	defer iaPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryPhoton.setNumericSeparatorsDto()",
		"")

	if err != nil {
		return err
	}

	if intAry == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intAry'",
		}
	}

	if validateNumericSeparators {

		err = customSeparators.IsValid(ePrefix.XCpy("Validating 'customSeparators'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = customSeparators.IsValid(ePrefix.XCpy(\n" +
					"  Validating 'customSeparators').String())",
				ErrContext: "Input parameter 'customSeparators' is invalid.\n" +
					"'customSeparators' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	intAry.decimalSeparator = customSeparators.DecimalSeparator

	intAry.thousandsSeparator = customSeparators.ThousandsSeparator

	intAry.currencySymbol = customSeparators.CurrencySymbol

	return nil
}

// setNumericSeparatorsToDefaultIfEmpty
//
//	 Receives a pointer to an IntAry object. By defintion, the
//	 IntAry contains fields for Numeric Separators.
//
//		If numeric separators are set to zero or nil, this method will
//		set those numeric separators to the USA defaults. This means
//		that the Decimal separator is set to a period ('.'), the
//		Thousands separator is set to a comma (',') and the currency
//		symbol is set to the dollar sign ('$').
//
//		If the numeric separators were previously set to a value other
//		than zero or nil, that value is not altered by this method.
//
//		Effectively, this method ensures that numeric separators are
//		set to valid values.
func (iaPhoton *intAryPhoton) setNumericSeparatorsToDefaultIfEmpty(
	intAry *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaPhoton.lock == nil {
		iaPhoton.lock = new(sync.Mutex)
	}

	iaPhoton.lock.Lock()

	defer iaPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryPhoton.setNumericSeparatorsToDefaultIfEmpty()",
		"")

	if err != nil {
		return err
	}

	if intAry == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intAry'",
		}
	}

	if intAry.decimalSeparator == 0 {
		intAry.decimalSeparator = '.'
	}

	if intAry.thousandsSeparator == 0 {
		intAry.thousandsSeparator = ','
	}

	if intAry.currencySymbol == 0 {
		intAry.currencySymbol = '$'
	}

	return nil
}

// setNumericSeparatorsToUSADefault
//
//		Sets Numeric Separators to the United States of America (USA)
//		defaults.
//
//		Sets Numeric separators:
//		  Decimal Point Separator
//		  Thousands Separator
//		  Currency Symbol
//
//	 Call specific methods to set numeric separators for other countries or
//	 cultures:
//
//		ia.SetDecimalSeparator()
//		ia.SetThousandsSeparator()
//		ia.SetCurrencySymbol()
func (iaPhoton *intAryPhoton) setNumericSeparatorsToUSADefault(
	intAry *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaPhoton.lock == nil {
		iaPhoton.lock = new(sync.Mutex)
	}

	iaPhoton.lock.Lock()

	defer iaPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryPhoton.setNumericSeparatorsToUSADefault()",
		"")

	if err != nil {
		return err
	}

	if intAry == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'intAry'",
		}
	}

	intAry.decimalSeparator = '.'

	intAry.thousandsSeparator = ','

	intAry.currencySymbol = '$'

	return nil
}

// setNumSepSymbol
//
//	Sets the number separator symbol on an instance of
//	BigIntFixedDecimal passed as input parameter 'bigIFxDec'.
//
//	Input Parameters
//	================
//
//	intAry               *IntAry
//	  The instance of BigIntNum which will be configured with
//	  the numeric separator character passed as input parameter
//	  'numSepSymbol'.
//
//	numSepSymbolType    NumSepSymbolCode
//	  An integer value enumeration used to desigate the specific
//	  numeric separator in 'bNum' which will set to the value
//	  of input parameter 'numSepSymbol'.
//
//	    const (
//	      // DECIMALSYMBOL
//	      // Symbol for the separator character used to
//	      // separate integer and fractional segments of
//	      // a floating point number or curreny value.
//	      DECIMALSYMBOL NumSepSymbolCode = iota
//
//	      // THOUSANDSYMBOL
//	      // Symbol for the separator character used to
//	      // separate thousands in a numeric presentation
//	      // where the numeric value is greater than 999
//	      THOUSANDSYMBOL
//
//	      // CURRENCYSYMBOL
//	      // Symbol for the character used to designate a
//	      // numeric value as currency.
//	      CURRENCYSYMBOL
//	    )
//
//	numSepSymbol        rune
//	  The specific numer separator character which will be
//	  transferred to 'intAry'.
func (iaPhoton *intAryPhoton) setNumSepSymbol(
	intAry *IntAry,
	numSepSymbolType NumSepSymbolCode,
	numSepSymbol rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaPhoton.lock == nil {
		iaPhoton.lock = new(sync.Mutex)
	}

	iaPhoton.lock.Lock()

	defer iaPhoton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryPhoton.setNumSepSymbol",
		"")

	if err != nil {
		return err
	}

	if intAry == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'intAry'",
		}
	}

	if numSepSymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("numSepSymbolType is eqaul to %v separator.",
				numSepSymbolType.String()),
			ErrMessage: "Error: Input parameter 'numSepSymbol' is INVALID!\n" +
				"'numSepSymbol' is empty and has a zero value.",
		}
	}

	switch numSepSymbolType {

	case DECIMALSYMBOL:
		intAry.decimalSeparator = numSepSymbol
	case THOUSANDSYMBOL:
		intAry.thousandsSeparator = numSepSymbol
	case CURRENCYSYMBOL:
		intAry.currencySymbol = numSepSymbol
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
