package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumNanobot struct {
	lock *sync.Mutex
}

// NewWithNumSeps
//
// Returns a new BigIntNum instance initialized to zero.
//
// Input parameter 'numSeps' will be used to seed the new
// BigIntNum instance with numeric separators (decimal
// separator, thousands separator and currency symbol).
//
// If input parameter 'numSeps' is determined to be empty,
// the new returned instance of BigIntNum, will be
// automatically configured with USA default numeric
// separators (decimal separator, thousands separator
// and currency symbol).
func (bIntNumNano *bigIntNumNanobot) newWithNumSeps(
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumNano.lock == nil {
		bIntNumNano.lock = new(sync.Mutex)
	}

	bIntNumNano.lock.Lock()

	defer bIntNumNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumMechanics.newWithNumSeps()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSeps.SetDefaultsIfEmpty()

	b := new(bigIntNumMechanics).new()

	new(bigIntNumElectron).empty(&b)

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		&b, numSeps, ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	return b, nil
}

// setBigFloat
//
// Sets the value of a BigIntNum (input parameter 'bNum' using a
// *big.Float floating point input parameter.  The maximum
// precision of the generated output number is specified by the
// input parameter, 'maxPrecision'.
//
// Input Parameters
// ================
//
//	bigFloat *big.Float
//
//		This *big.Float value will be converted into an instance of
//		BigIntNum.
//
//	maxPrecision uint
//
//		The maximum precision for the resulting BigIntNum after
//		conversion of input parameter 'bigFloat'. Final precision
//		will never be greater than 'maxPrecision'; however, actual
//		precision may be less than 'maxPrecision'.
//
// Background
// ==========
//
// As part of converting a BigFloat to a BigIntNum number,
// the Accuracy flag is analyzed to determine if rounding errors
// associated with the conversion. The internal Accuracy Flag is
// set as:
//
//			Below Accuracy == -1	(Returns an error)
//	   Exact Accuracy == 0		(No Error Returned)
//	   Above Accuracy == +1		(Returns an error)
//
// If Accuracy == 0, no error is issued by this method. However, if
// Accuracy == -1 or Accuracy == +1, an error will be returned. All
// conversions must be exact.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) contained in the original BigIntNum ('bNum')
// will remain unchanged and will not be altered by this method.
func (bIntNumNano *bigIntNumNanobot) setBigFloat(
	bNum *BigIntNum,
	bigFloat *big.Float,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumNano.lock == nil {
		bIntNumNano.lock = new(sync.Mutex)
	}

	bIntNumNano.lock.Lock()

	defer bIntNumNano.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNanobot.setBigInt",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {
		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if bigFloat == nil {
		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigFloat'",
		}
	}

	rat, accuracyFlag := bigFloat.Rat(nil)

	if accuracyFlag == -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "rat, accuracyFlag := bigFloat.Rat(nil)",
			ErrMessage: "Error: Conversion of input parameter 'bigFloat' resulted in Accuracy Flag == -1\n" +
				"or 'Below Accuracy'. Conversion is NOT Exact!",
		}
	}

	if accuracyFlag == 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "rat, accuracyFlag := bigFloat.Rat(nil)",
			ErrMessage: "Error: Conversion of input parameter 'bigFloat' resulted in Accuracy Flag == +1\n" +
				"or 'Above Accuracy'. Conversion is NOT Exact!",
		}
	}

	err = bNum.SetBigRat(rat, maxPrecision)

	err = new(bigIntNumMolecule).setBigRat(
		bNum,
		rat,
		maxPrecision,
		ePrefix)

	return err
}

// setBigInt - Sets the value of the current BigIntNum instance using
// the input parameters *big.Int integer and precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
// Input Parameters
// bigI *big.Int	- 'bigI' is a type *big.Int and represents the integer
//
//	value of the number; that is, the numeric value without decimal
//	digits.
//
// precision uint	- This unsigned integer (always a positive value) identifies
//
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places. Example:
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bIntNumNano *bigIntNumNanobot) setBigInt(
	bNum *BigIntNum,
	bigI *big.Int,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumNano.lock == nil {
		bIntNumNano.lock = new(sync.Mutex)
	}

	bIntNumNano.lock.Lock()

	defer bIntNumNano.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNanobot.setBigInt",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {
		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if bigI == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'bigI' is a nil pointer!\n",
			ePrefix.String())

	}

	numSeps := NumericSeparatorDto{}

	numSeps.DecimalSeparator = bNum.decimalSeparator

	numSeps.ThousandsSeparator = bNum.thousandsSeparator

	numSeps.CurrencySymbol = bNum.currencySymbol

	new(bigIntNumElectron).empty(bNum)

	bNum.bigInt = big.NewInt(0).Set(bigI)

	bNum.precision = precision

	base10 := big.NewInt(0).SetInt64(int64(10))

	bigPrecision := big.NewInt(0).SetInt64(int64(bNum.precision))

	bNum.scaleFactor = big.NewInt(0).Exp(base10, bigPrecision, nil)

	bNum.numberOfExpectedDigits = big.NewInt(0)

	result := bNum.bigInt.Cmp(big.NewInt(0))

	if result == -1 {

		bNum.sign = -1
		minusOne := big.NewInt(0).SetInt64(-1)
		bNum.absBigInt = big.NewInt(0).Mul(bNum.bigInt, minusOne)

	} else {

		bNum.sign = 1
		bNum.absBigInt = big.NewInt(0).Set(bNum.bigInt)

	}

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned from:\n"+
			"err := new(bigIntNumAtom).setNumericSeparatorsDto(\n"+
			"     bNum, numSeps, ePrefix)\n"+
			"Error= %v\n",
			ePrefix.String(),
			err.Error())
	}

	return err
}

// setIntFracStrings
//
// Sets the value of the BigIntNum instance passed as input parameter
// 'bNum' using the numeric value represented by separate integer
// and fractional components passed as strings.
//
// Input parameters 'intStr' and 'fracStr' are strings representing the integer and
// fractional components. They are combined by this method to create a numeric value
// which is assigned to the current BigIntNum instance.
//
// Input parameter 'signVal' must be set to one of two values: +1 or -1. This value is
// used to signal the sign of the resulting numeric value. +1 generates a positive number
// and -1 generates a negative number. If input parameters 'inStr' or 'fracStr' contain
// a leading minus or plus sign character, it will be ignored. The sign of the resulting
// numeric value is controlled strictly by input parameter, 'signVal'.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an instance of type BigIntNum.
// The calling method must do this!
func (bIntNumNano *bigIntNumNanobot) setIntFracStrings(
	bNum *BigIntNum,
	intStr,
	fracStr string,
	signVal int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumNano.lock == nil {
		bIntNumNano.lock = new(sync.Mutex)
	}

	bIntNumNano.lock.Lock()

	defer bIntNumNano.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNanobot.setIntFracStrings",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {
		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	cleanIntRuneAry := make([]rune, 0, 100)

	zeroChar := uint8('0')
	nineChar := uint8('9')

	lStr := len(intStr)

	if lStr == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "lStr := len(intStr)",
			ErrContext: "if lStr == 0 {",
			ErrMessage: "Error: Input Parameter 'intStr' is zero Length!",
		}
	}

	isFirstRune := true

	// Create pure number string from 'intStr'
	for i := 0; i < lStr; i++ {

		if intStr[i] >= zeroChar &&
			intStr[i] <= nineChar {

			if isFirstRune && signVal == -1 {
				cleanIntRuneAry = append(cleanIntRuneAry, '-')
			}

			isFirstRune = false

			cleanIntRuneAry = append(cleanIntRuneAry, rune(intStr[i]))
		}
	}

	if len(cleanIntRuneAry) == 0 {
		cleanIntRuneAry = append(cleanIntRuneAry, '0')
	}

	lStr = len(fracStr)

	if lStr > 0 {

		isFirstRune = true

		for j := 0; j < lStr; j++ {

			if fracStr[j] >= zeroChar &&
				fracStr[j] <= nineChar {

				decNumSep, err := bNum.GetDecimalSeparator()

				if err != nil {

					return &FuncReturnError{
						ErrPrefix:  ePrefix.String(),
						ReturnFunc: "decNumSep, err := bNum.GetDecimalSeparator()",
						ErrContext: "",
						ErrMessage: err.Error(),
					}
				}

				if isFirstRune {

					cleanIntRuneAry = append(cleanIntRuneAry, decNumSep)

					isFirstRune = false
				}

				cleanIntRuneAry = append(cleanIntRuneAry, rune(fracStr[j]))
			}

		}
	}

	err = new(bigIntNumMolecule).setNumStr(
		bNum,
		string(cleanIntRuneAry),
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "\terr = new(bigIntNumMolecule).setNumStr(bNum," +
				"    string(cleanIntRuneAry), ePrefix)",
			ErrContext: fmt.Sprintf("cleanIntRuneAry='%v'", string(cleanIntRuneAry)),
			ErrMessage: err.Error(),
		}
	}

	return nil
}
