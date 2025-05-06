package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.isBigIntNumValid",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if bNum.bigInt == nil {

		bNum.bigInt = big.NewInt(0)

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.bigInt' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum.bigInt was reset to zero.\n",
			ePrefix.String())
	}

	if bNum.sign != -1 && bNum.sign != 1 {

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.sign' is NOT equal to +1 or -1 !\n"+
			"FATAL ERROR!\n",
			ePrefix.String())

	}

	if bNum.absBigInt == nil {

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.absBigInt' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix.String())
	}

	if bNum.scaleFactor == nil {

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.scaleFactor' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix.String())

	}

	return nil
}

// getIntegerPart
//
// Returns a BigIntNum equal to the integer value of
// the input parameter 'bNum' (type BigIntNum).
//
// Examples:
//
//					 Current
//					BigIntNum				 		Return
//		 			  Value		  		    Value
//					----------				---------
//
//	         123.456						 123
//					 -123.456						-123
//					  123								 123
//					 -123								-123
//
//		NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumAtom *bigIntNumAtom) getIntegerPart(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.getIntegerPart",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&ReturnBasicError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "FATAL ERROR: Input parameter 'bNum' is a nil pointer.",
			}
	}

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		return new(bigIntNumMechanics).newBigInt(
			big.NewInt(0), 0, ePrefix)
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision)), nil)

	quotient := big.NewInt(0).Quo(bNum.bigInt, scaleVal)

	bNum2, err := new(bigIntNumMechanics).
		newBigInt(quotient, 0, ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	return bNum2, nil
}

// getBigIntNumStr - Converts a BigIntNum value to string of
// numbers which includes the decimal place and decimal digits
// if they exist.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must do this!
func (bIntNumAtom *bigIntNumAtom) getBigIntNumStr(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.getBigIntNumStr()",
		"")

	if err != nil {
		return "", err
	}

	if bNum == nil {

		return "",
			fmt.Errorf("%v\n"+
				"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
				ePrefix.String())
	}

	numStr, err := new(bigIntNumMolecule).formatBigIntNumStr(
		bNum,
		LEADMINUSNEGVALFMTMODE,
		ePrefix)

	if err != nil {

		return "", err
	}

	return numStr, nil
}

// GetNumericSeparatorsDto - Returns a structure containing the
// character or rune values for decimal place separator, thousands
// separator and currency symbol.
//
//	NOTE:
//
// ================
// This method does NOT test the validity of 'bNumDestination'
// BigIntNum instance. The calling method must do this!
func (bIntNumAtom *bigIntNumAtom) getNumericSeparatorsDto(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (NumericSeparatorDto, error) {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.getBigIntNumStr()",
		"")

	if err != nil {
		return NumericSeparatorDto{}, err
	}

	if bNum == nil {

		return NumericSeparatorDto{},
			fmt.Errorf("%v\n"+
				"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
				ePrefix.String())
	}

	numSeps := NumericSeparatorDto{}
	numSeps.DecimalSeparator = bNum.decimalSeparator
	numSeps.ThousandsSeparator = bNum.thousandsSeparator
	numSeps.CurrencySymbol = bNum.currencySymbol

	return numSeps, nil
}

// getSciNotationStr
//
// Returns a string expressing the current BigIntNum numerical
// value as scientific notation.
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//
//											scientific notation string. If the value of 'mantissaLen'
//											is less than two ('2'), this method will automatically set
//											the 'mantissaLen' to a default value of two ('2').
//
//											Example Scientific Notation:
//											----------------------------
//
//	 										scientific notation string: '2.652e+8'
//
//	 										significand = '2.652'
//	 										significand integer digit = '2'
//												mantissa		= significand factional digits = '.652'
//	 										exponent    = '8'  (10^8)
func (bIntNumAtom *bigIntNumAtom) getSciNotationStr(
	bNum *BigIntNum,
	mantissaLen uint,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.getSciNotationStr",
		"")

	if err != nil {
		return "", err
	}

	if bNum == nil {

		return "",
			&ReturnBasicError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "FATAL ERROR: Input parameter 'bNum' is a nil pointer.",
			}
	}

	//sciNotation, err := bNum.GetSciNotationNumber(mantissaLen)

	sciNotation, err := new(bigIntNumProton).
		bigIntNumGetSciNotationNumber(bNum, mantissaLen, ePrefix)

	if err != nil {
		return "", err
	}

	result, err := sciNotation.GetSciNotationStr(mantissaLen)

	if err != nil {
		return "",
			&ReturnBasicError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "  result, err := sciNotation.\n" +
					"    GetSciNotationStr(mantissaLen)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}

// setNumericSeparators - Used to assign values for the Decimal and Thousands separators as well
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
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.setNumericSeparators",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())
	}

	if decimalSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is set to '0' - Invalid rune!\n",
			ePrefix.String())
	}

	if thousandsSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'thousandsSeparator' is set to '0' - Invalid rune!\n",
			ePrefix.String())

	}

	if currencySymbol == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.CurrencySymbol' is set to '0' - Invalid rune!\n",
			ePrefix.String())

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
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.setNumericSeparatorsDto",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())
	}

	if customSeparators.DecimalSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.DecimalSeparator' is set to '0' - Invalid rune!\n",
			ePrefix.String())
	}

	if customSeparators.ThousandsSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.ThousandsSeparator' is set to '0' - Invalid rune!\n",
			ePrefix.String())

	}

	if customSeparators.CurrencySymbol == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.CurrencySymbol' is set to '0' - Invalid rune!\n",
			ePrefix.String())

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
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.setNumericSeparatorsToDefaultIfEmpty",
		"")

	if err != nil {
		return err
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
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.setNumericSeparatorsToUSADefault",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.",
			ePrefix)
	}

	bNum.decimalSeparator = '.'

	bNum.thousandsSeparator = ','

	bNum.currencySymbol = '$'

	return nil
}
