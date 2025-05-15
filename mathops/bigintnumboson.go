package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumBoson struct {
	lock *sync.Mutex
}

// roundToDecimalPlace
//
// Rounds the current BigIntNum instance to a specified number of
// decimal places.
//
// 'precision' equals the number of digits to the right of the
// decimal place.
//
// Example:
//
//	integer= 123456; precision = 3; Numeric Value= 123.456
//
// If the value of BigIntNum.bigInt is zero ('0'), that zero value
// will remain unaltered. However, the BigIntNum.precision value
// will be set equal to input parameter, 'precision'.
//
// If the number of decimal places specified for rounding
// ('precision') is equal to the current BigIntNum.precision, no
// action is taken.
//
// If the number of decimal places specified for rounding
// ('precision') is greater than the current BigIntNum.precision
// value, trailing zeros are added to the current BigIntNum.bigInt
// value and BigIntNum.precision is set equal to input parameter,
// 'precision'.
//
// Finally, if the number of decimal places specified for rounding
// ('precision') is less than the current BigIntNum.precision
// value, the fractional digits will be rounded in accordance with
// the input parameter, 'precision'.
//
// Examples:
//
//		  Original        'precision'        Resulting
//		   Value         input parameter       Value
//
//		 654.123456            9              654.123456000
//	   654.123456            4              654.1235
//
//		-654.123456            9             -654.123456000
//		-654.123456            4             -654.1235
//
//		   0                   3                0.000
//		   0.000000            0                0
//
// Existing numeric separators (decimal separator, thousands
// separator and currency symbol) remain unchanged and are not
// altered by this method.
//
//	NOTE
//	====
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumBoson *bigIntNumBoson) roundToDecimalPlace(
	bNum *BigIntNum,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.roundToDecimalPlace",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if bNum.precision == precision {
		// Nothing to do. Specified 'precision' is
		// already implemented.

		return nil
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		bNum,
		ePrefix.XCpy("Setting 'bNum'"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numSeps, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		bNum, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err := new(bigIntNumAtom).getNumericSeparatorsDto(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	// if bigInt == zero, set precision and return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		bNum2, err := new(bigIntNumMechanics).newBigInt(
			big.NewInt(0),
			precision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bNum2, err := new(bigIntNumMechanics).newBigInt(\n" +
					"    big.NewInt(0), precision, ePrefix)",
				ErrContext: "if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {",
				ErrMessage: err.Error(),
			}
		}

		err = new(bigIntNumUtility).bigIntNumCopyIn(
			bNum,
			&bNum2,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumUtility).bigIntNumCopyIn(\n" +
					"    bNum,  &bNum2, ePrefix)",
				ErrContext: "if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {",
				ErrMessage: err.Error(),
			}
		}

		err = new(bigIntNumAtom).setNumericSeparatorsDto(
			bNum,
			numSeps,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
					"    bNum, numSeps, ePrefix)",
				ErrContext: "if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bNum.precision < precision {

		deltaPrecision := precision - bNum.precision

		err = new(bigIntNumProton).bigIntNumExtendPrecision(
			bNum,
			deltaPrecision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumProton).bigIntNumExtendPrecision(\n" +
					"    bNum, deltaPrecision, ePrefix)",
				ErrContext: "if bNum.precision < precision {",
				ErrMessage: err.Error(),
			}
		}

		err = new(bigIntNumAtom).setNumericSeparatorsDto(
			bNum,
			numSeps,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
					"    bNum, numSeps, ePrefix)",
				ErrContext: "if bNum.precision < precision {",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	// Must be: bNum.precision >  precision

	bigNumRound5, err := new(bigIntNumMechanics).newBigInt(
		big.NewInt(5),
		precision+1,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bigNumRound5, err := new(bigIntNumMechanics).newBigInt(\n" +
				"    big.NewInt(5), precision+1, ePrefix)",
			ErrContext: "Must be: bNum.precision > precision",
			ErrMessage: err.Error(),
		}
	}

	bigNumBase, err := new(bigIntNumMechanics).newBigInt(
		bNum.absBigInt,
		bNum.precision,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bigNumBase, err := new(bigIntNumMechanics).newBigInt(\n" +
				"    bNum.absBigInt, bNum.precision, ePrefix)",
			ErrContext: "Must be: bNum.precision > precision",
			ErrMessage: err.Error(),
		}
	}

	result, err := new(BigIntMathAdd).AddBigIntNums(
		bigNumBase, bigNumRound5)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "result, err := new(BigIntMathAdd).AddBigIntNums(\n" +
				"    bigNumBase, bigNumRound5)",
			ErrContext: "Must be: bNum.precision > precision",
			ErrMessage: err.Error(),
		}
	}

	// 10^deltaPrecision
	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision-precision)), nil)

	result.bigInt = big.NewInt(0).Quo(result.bigInt, scaleVal)

	if bNum.sign < 0 {
		result.bigInt = big.NewInt(0).Neg(result.bigInt)
	}

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
				"    bNum, numSeps, numSeps)",
			ErrContext: "10^deltaPrecision",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		result.bigInt,
		precision,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: " err = new(bigIntNumNanobot).setBigInt(\n" +
				"    bNum, result.bigInt, precision, ePrefix)",
			ErrContext: "10^deltaPrecision",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setCurrencySymbol
//
// Configures the input parameter rune as the currency symbol for
// an instance of BigIntNum. It is used when generating number
// strings for display. The BigIntNum type is passed as input
// parameter 'bNum'.
//
// FYI, in the USA, the currency symbol is the dollar sign ('$').
//
// Note: If a zero value is submitted as input, an error will be
// returned.
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
func (bIntNumBoson *bigIntNumBoson) setCurrencySymbol(
	bNum *BigIntNum,
	currencySymbol rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.setCurrencySymbol",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if currencySymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'currencySymbol' is INVALID!\n" +
				"'currencySymbol', of type 'rune', is equal to zero.",
		}
	}

	bNum.currencySymbol = currencySymbol

	return nil
}

// setDecimalSeparator
//
// Configures the input parameter rune as the Decimal Separator for
// an instance of BigIntNum. The BigIntNum instance is passed as
// input parameter, 'bNum'.
//
// The Decimal Separator is used to separate the integer and
// fractional elements of a number string.
//
// In the USA, the Decimal Separator is a period character ('.').
//
//	Example: 123.45
//
// Note: If a zero value is submitted for input parameter
// 'decimalSeparator', an error will be returned.
func (bIntNumBoson *bigIntNumBoson) setDecimalSeparator(
	bNum *BigIntNum,
	decimalSeparator rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.setDecimalSeparator",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if decimalSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'decimalSeparator' is INVALID!\n" +
				"'decimalSeparator', of type 'rune', is equal to zero.",
		}
	}

	bNum.decimalSeparator = decimalSeparator

	return nil
}

// setDecimalSymbol
//
// Configures the input parameter rune as the Thousands Separator
// for an instance of BigIntNum. The BigIntNum instance is passed
// as input parameter, 'bNum'.
//
// The Decimal Separator is used to separate the integer and
// fractional elements of a number string.
//
// In the USA, the Thousands Separator is the comma character (',').
// Example: 1,000,000
//
// Note: If a zero value is submitted for input parameter
// 'thousandsSeparator', an error will be returned.
func (bIntNumBoson *bigIntNumBoson) setThousandsSeparator(
	bNum *BigIntNum,
	thousandsSeparator rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.setThousandsSeparator",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if thousandsSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'thousandsSeparator' is INVALID!\n" +
				"'thousandsSeparator', of type 'rune', is equal to zero.",
		}
	}

	bNum.thousandsSeparator = thousandsSeparator

	return nil
}

// shiftPrecisionLeft
//
// Shifts precision of the current BigIntNum numeric value to the
// left by 'shiftLeftPlaces' decimal places. This is a 'relative'
// shift-left operation. The shift left operation is therefore
// performed with the current decimal place position as the starting
// point.
//
// This operation is equivalent to:
//
//	result = Decimal value / 10^shiftLeftPlaces
//	                 or
//	signed number divided by 10 raised to the
//	power of shiftLeftPlaces.
//
// This method performs a relative shift left of the decimal place
// position. Be careful, this is NOT Shift Number Left operation.
// Instead, this is a Shift Precision Left operation which means
// that the decimal place will be shifted left.
//
// See Examples below.
//
//	 Input Parameters
//	 ================
//
//			shiftLeftPlaces int	- The number of positions the decimal place will be
//														shifted left from its current position.
//
//	 Examples
//	 ========
//
//			shift-left
//
//		   signed Number   places      Result
//
//		    "123456.789"      3      "123.456789"
//		    "123456.789"      2      "1234.56789"
//		    "123456.789"      6      "0.123456789"
//		    "123456789"       6      "123.456789"
//		    "123"             5      "0.00123"
//		    "0"               3      "0"
//		    "123456.789"      0      "123456.789" - zero has no effect on original number string
//
//		   "-123456.789"      0     "-123456.789"
//		   "-123456.789"      3     "-123.456789"
//		   "-123456789"       6     "-123.456789"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
//	NOTE
//	====
//
//	This method does NOT test the validity of 'bNum', an
//	instance of type BigIntNum. The calling method must
//	do this!
func (bIntNumBoson *bigIntNumBoson) shiftPrecisionLeft(
	bNum *BigIntNum,
	shiftLeftPlaces uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.shiftPrecisionLeft",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	isbNumZero := false

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		isbNumZero = true
	}

	if shiftLeftPlaces == 0 || isbNumZero {

		return nil
	}

	newPrecision := bNum.precision + shiftLeftPlaces

	return new(bigIntNumNanobot).setBigInt(
		bNum,
		bNum.bigInt,
		newPrecision,
		ePrefix.XCpy(fmt.Sprintf("Setting 'bNum'; bNum.bigInt= '%v' newPrecision= '%v'",
			bNum.bigInt.Text(10), newPrecision)))
}

// shiftPrecisionRight
//
// Shifts precision of the BigIntNum input parameter 'bNum'
// numeric value to the right by 'shiftRightPlaces' decimal
// places. This is a 'relative' shift-right operation. The
// shift right operation is therefore performed with the
// current decimal place position as the starting point.
//
// This is equivalent to:
//
//	   result = Decimal value X 10^shiftRightPrecision
//
//		                      or
//
//	   Decimal numeric value multiplied by 10 raised to
//	   the power of shiftRightPrecision.
//
// This method performs a relative shift right of the decimal
// place position. Be careful, this is NOT a Shift Number Right
// operation. This is Shift Precision Right which means that the
// decimal place will be shifted right.
//
// See Examples below.
//
//		 Input Parameters
//		 ================
//
//			shiftRightPlaces      int
//
//	   The number of positions the decimal place will be
//	   shifted right from its current position.
//
//		 Examples
//		 ========
//
//			shift-right
//
//		  signed Number     places    Result
//
//		  "123456.789"        3       "123456789"
//		  "123456.789"        2       "12345678.9"
//		  "123456.7896"       7       "1234567896000"
//		  "123456789"         6       "123456789000000"
//		  "123"               5       "12300000"
//		  "0"                 3       "0"
//		  "123456.789"        0       "123456.789" - zero has no effect on original number string
//
//		 "-123456.789         0      "-123456.789"
//		 "-123456.789         3      "-123456789"
//		 "-123456789"         6      "-123456789000000"
//
//		 Numeric Separators
//		 ==================
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) for the current instance of BigIntNum will
// remain unchanged and are not altered by this method.
//
//	NOTE
//	====
//
//	This method does NOT test the validity of 'bNum', an
//	instance of type BigIntNum. The calling method must
//	do this!
func (bIntNumBoson *bigIntNumBoson) shiftPrecisionRight(
	bNum *BigIntNum,
	shiftRightPlaces uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.shiftPrecisionRight",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	isbNumZero := false

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		isbNumZero = true
	}

	if shiftRightPlaces == 0 || isbNumZero {

		return nil

	}

	bigINanobot := new(bigIntNumNanobot)

	if shiftRightPlaces <= bNum.precision {

		newPrecision := bNum.precision - shiftRightPlaces

		err = bigINanobot.setBigInt(
			bNum,
			bNum.bigInt,
			newPrecision,
			ePrefix.XCpy(fmt.Sprintf("Setting 'bNum'; bNum.bigInt= '%v' newPrecision= '%v'",
				bNum.bigInt.Text(10), newPrecision)))

		return err
	}

	// shiftRightPlaces > bNum.precision

	newPrecision := shiftRightPlaces - bNum.precision

	bigITen := big.NewInt(10)

	exponent := big.NewInt(int64(newPrecision))

	scaleFactor := big.NewInt(0).Exp(bigITen, exponent, nil)

	newValue := big.NewInt(0).Mul(bNum.bigInt, scaleFactor)

	err = bigINanobot.setBigInt(
		bNum,
		newValue,
		0,
		ePrefix.XCpy(fmt.Sprintf("Setting 'bNum'; newValue= '%v' newPrecision= '%v'",
			newValue.Text(10), newPrecision)))

	return err
}

// trimTrailingFracZeros
//
// This method will delete non-significant trailing zeros from
// the fractional digits of the current BigIntNum numerical
// value.
//
//	Examples
//	========
//
//	  Initial Value   Trimmed Value
//
//	  456.123000         456.123
//	    0.000              0
//	    7.0                7
//	 -456.123000        -456.123
//
//	NOTE
//	====
//
//	This method does NOT test the validity of 'bNum', an
//	instance of type BigIntNum. The calling method must
//	do this!
func (bIntNumBoson *bigIntNumBoson) trimTrailingFracZeros(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.trimTrailingFracZeros",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if bNum.precision == 0 {
		return nil
	}

	biBaseZero := big.NewInt(0)

	if bNum.bigInt.Cmp(biBaseZero) == 0 {

		bNum.precision = 0

		bNum.scaleFactor = big.NewInt(1)

		return nil
	}

	// bNum.precision must be greater than zero
	biBase10 := big.NewInt(10)

	scrap := big.NewInt(0)

	newBigIntNum, mod10 :=
		big.NewInt(0).QuoRem(bNum.bigInt, biBase10, scrap)

	doReset := false

	for mod10.Cmp(biBaseZero) == 0 && bNum.precision > 0 {

		bNum.bigInt.Set(newBigIntNum)

		bNum.precision--

		newBigIntNum, mod10 = big.NewInt(0).QuoRem(bNum.bigInt, biBase10, scrap)

		doReset = true
	}

	if doReset {

		if bNum.sign < 0 {

			bNum.absBigInt = big.NewInt(0).Neg(bNum.bigInt)

		} else {

			bNum.absBigInt = big.NewInt(0).Set(bNum.bigInt)

		}

		bigPrecision := big.NewInt(0).SetInt64(int64(bNum.precision))

		bNum.scaleFactor = big.NewInt(0).Exp(biBase10, bigPrecision, nil)

	}

	return nil
}

// truncToDecPlace
//
// Truncates BigIntNum input parameter 'bNum' to the number of
// decimal places specified by input parameter 'precision'. No
// rounding occurs. The trailing digits are simply truncated or
// deleted in order to achieve the specified number of decimal
// places.
//
// 'precision' equals the number of digits to the right of the
// decimal place.
//
// If the value of BigIntNum.bigInt is zero ('0'), that zero value
// will remain unaltered. However, BigIntNum.precision will be set
// equal to input parameter, 'precision'.
//
// If the number of decimal places specified for truncation
// ('precision') is equal to the current BigIntNum.precision, no
// action is taken and the original BigIntNum numeric value
// remains unchanged.
//
// If the number of decimal places specified for truncation
// ('precision') is greater than the current BigIntNum.precision,
// trailing zeros are added to the current BigIntNum.bigInt value
// and BigIntNum.precision is set equal to input parameter,
// 'precision'.
//
// If 'precision' is less than the current BigIntNum.precision
// value, the current BigIntNum numeric value is truncated to
// the specified 'precision' value and NO rounding occurs.
//
//	Examples
//	========
//
//	   Original           'newPrecision'        Resulting
//	    Value             input parameter         Value
//
//	  654.123456                9              654.123456000
//	  654.123456                4              654.1234 (no rounding)
//
//	 -654.123456                9             -654.123456000
//	 -654.123456                4             -654.1234 (no rounding)
//
//	    0                       3                0.000
//	    0.000000                0                0
//
//	Numeric Separators
//	==================
//
// Existing numeric separators (decimal separator, thousands
// separator and currency symbol) in the current instance of
// BigIntNum will remain unchanged and are not altered by this
// method.
func (bIntNumBoson *bigIntNumBoson) truncToDecPlace(
	bNum *BigIntNum,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.truncToDecPlace",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if bNum.precision == precision {

		// Nothing to do. Specified 'precision' is already implemented.

		return nil
	}

	// bigInt == zero, set precision and return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		bNum.precision = precision

		return nil
	}

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bNum.precision < precision {

		deltaPrecision := precision - bNum.precision

		return new(bigIntNumProton).bigIntNumExtendPrecision(
			bNum,
			deltaPrecision,
			ePrefix.XCpy(fmt.Sprintf("Extending 'bNum'. deltaPrecision= '%v'",
				deltaPrecision)))
	}

	// Must be bNum.precision > precision
	base10 := big.NewInt(10)

	deltaPrecision := big.NewInt(int64(bNum.precision - precision))

	newBigInt := big.NewInt(0).Set(bNum.absBigInt)

	newScaleVal := big.NewInt(0).Exp(base10, deltaPrecision, nil)

	newBigInt = big.NewInt(0).Quo(newBigInt, newScaleVal)

	if bNum.sign < 1 {
		newBigInt = big.NewInt(0).Neg(newBigInt)
	}

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		newBigInt,
		precision,
		ePrefix.XCpy(fmt.Sprintf("Setting 'bNum'. newBigInt= '%v' precision= '%v'",
			newBigInt.Text(10), precision)))

	return err
}
