package mathops

import (
	"fmt"
	"math/big"
	"sync"
)

type bigIntNumAtom struct {
	lock *sync.Mutex
}

// isBigIntNumValid - returns a boolean value signaling
// whether the current BigIntNum object is valid.
func (bIntNumAtom *bigIntNumAtom) isBigIntNumValid(
	bNum *BigIntNum,
	callingMethodChain string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "Active Method: bigIntNumAtom.isBigIntNumValid()"

	var err error

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if bNum.bigInt == nil {

		bNum.bigInt = big.NewInt(0)

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.bigInt' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum.bigInt was reset to zero.\n",
			ePrefix)

	}

	if bNum.sign != -1 && bNum.sign != 1 {

		err = bNum.Reset()

		if err != nil {
			return fmt.Errorf("%v\n"+
				"This BigIntNum Instance is Invalid!\n"+
				"'bNum.sign' is NOT equal to +1 or -1 !\n"+
				"FATAL ERROR!\n"+
				"Attmpted reset of bNum to default values FAILED!.\n",
				ePrefix)

		}

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.sign' is NOT equal to +1 or -1 !\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix)
	}

	if bNum.absBigInt == nil {

		err = bNum.Reset()

		if err != nil {

			return fmt.Errorf("%v\n"+
				"This BigIntNum Instance is Invalid!\n"+
				"'bNum.absBigInt' is 'nil'!\n"+
				"FATAL ERROR!\n"+
				"Attmpted reset of bNum to default values FAILED!.\n",
				ePrefix)

		}

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.absBigInt' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix)
	}

	if bNum.scaleFactor == nil {

		err = bNum.Reset()

		if err != nil {
			return fmt.Errorf("%v\n"+
				"This BigIntNum Instance is Invalid!\n"+
				"'bNum.scaleFactor' is 'nil'!\n"+
				"FATAL ERROR!\n"+
				"Attmpted reset of bNum to default values FAILED!.\n",
				ePrefix)

		}

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.scaleFactor' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix)

	}

	return nil
}

// SetNumericSeparators - Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current number string.
//
// Different nations and cultures use different symbols to delimit numerical values. In the
// USA and many other countries, a period character ('.') is used to delimit integer and
// fractional digits within a numeric value (123.45). Likewise, thousands may be delimited
// by a comma (','). Currency signs very by nationality. For instance, the USA, Canada and
// several other countries use the dollar sign ($) as a currency symbol.
//
// For a list of major world currency symbols see:
//
//		MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//	 http://www.xe.com/symbols.php
//
// Note: If zero values are submitted as input for separator values, those values will default
// to USA standards.
//
// USA Examples:
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
func (bIntNumAtom *bigIntNumAtom) setNumericSeparators(
	bNum *BigIntNum,
	decimalSeparator rune,
	thousandsSeparator rune,
	currencySymbol rune,
	callingMethodChain string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "Active Method: bigIntNumAtom.isBigIntNumValid()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if decimalSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is set to '0' - Invalid rune!\n",
			ePrefix)
	}

	if thousandsSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'thousandsSeparator' is set to '0' - Invalid rune!\n",
			ePrefix)

	}

	if currencySymbol == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.CurrencySymbol' is set to '0' - Invalid rune!\n",
			ePrefix)

	}

	bNum.decimalSeparator = decimalSeparator

	bNum.thousandsSeparator = thousandsSeparator

	bNum.currencySymbol = currencySymbol

	return nil
}

// setNumericSeparatorsDto - Sets the values of numeric separators:
//
//	decimal place separator
//	thousands separator
//	currency symbol
//
// based on values transmitted through input parameter 'customSeparators'.
//
// If any of the values contained in input parameter 'customSeparators' is set
// to zero, an error will be returned.
//
// NOTE:
// This is a low-level operation. It is assumed that 'bNum'
// has already been validated.
func (bIntNumAtom *bigIntNumAtom) setNumericSeparatorsDto(
	bNum *BigIntNum,
	customSeparators NumericSeparatorDto,
	callingMethodChain string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "Active Method: bigIntNumAtom.setNumericSeparatorsDto() "

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if customSeparators.DecimalSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.DecimalSeparator' is set to '0' - Invalid rune!\n",
			ePrefix)
	}

	if customSeparators.ThousandsSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.ThousandsSeparator' is set to '0' - Invalid rune!\n",
			ePrefix)

	}

	if customSeparators.CurrencySymbol == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.CurrencySymbol' is set to '0' - Invalid rune!\n",
			ePrefix)

	}

	bNum.decimalSeparator = customSeparators.DecimalSeparator

	bNum.thousandsSeparator = customSeparators.ThousandsSeparator

	bNum.currencySymbol = customSeparators.CurrencySymbol

	return nil
}

// setNumericSeparatorsToDefaultIfEmpty - If numeric separators are
// set to zero or nil, this method will set those numeric
// separators to the USA defaults. This means that the
// Decimal separator is set to ('.'), the Thousands separator
// is set to (',') and the currency symbol is set to '$'.
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that numeric separators
// are set to valid values.
func (bIntNumAtom *bigIntNumAtom) setNumericSeparatorsToDefaultIfEmpty(
	bNum *BigIntNum,
	callingMethodChain string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "Active Method: bigIntNumAtom.setNumericSeparatorsToDefaultIfEmpty()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	if bNum.currencySymbol == 0 {
		bNum.currencySymbol = '$'
	}

	return nil
}

// setNumericSeparatorsToUSADefault - Sets Numeric separators:
//
//	Decimal Point Separator
//	Thousands Separator
//	Currency Symbol
//
// to the United States of America (USA) defaults.
//
// Call specific methods to set numeric separators for other countries or
// cultures:
//
//	bNum.SetDecimalSeparator()
//	bNum.SetThousandsSeparator()
//	bNum.SetCurrencySymbol()
func (bIntNumAtom *bigIntNumAtom) setNumericSeparatorsToUSADefault(
	bNum *BigIntNum,
	callingMethodChain string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "Active Method: bigIntNumAtom.setNumericSeparatorsToUSADefault()"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	bNum.decimalSeparator = '.'

	bNum.thousandsSeparator = ','

	bNum.currencySymbol = '$'

	return nil
}
