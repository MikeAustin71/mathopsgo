package mathops

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
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

	bINum, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINum, err := new(bigIntNumMechanics).newZero(" +
					"    0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		&bINum, numSeps, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(" +
					"    &bINum, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bINum, nil
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

// setBigInt
//
// Sets the value of the BigIntNum instance passed as input
// parameter 'bNum'. The new value for 'bNum' is extracted from
// input parameters bigI (*big.Int) and precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
//	 Input Parameters
//	 ===============
//
//	 bigI                     *big.Int
//	  'bigI' is a type *big.Int and represents the integer value
//	  of the number; that is, the numeric value without decimal
//	  digits.
//
//	 precision                uint
//	   This unsigned integer (always a positive value) identifies
//	   the location of the decimal place in the integer value
//	   'bigI'. The decimal place location is calculated by
//	    starting with the right most digit in the integer number
//	    and counting	left, 'precision' places. Example:
//
//			   Integer Value    precision    Numeric Value
//			     123456             3           123.456
//		                 123456 x 10^-3 =  123.456
//
//		Numeric Separators
//		==================
//
//	 Existing numeric separators (decimal separator, thousands
//	 separator and currency symbol), contained in 'bNum', remain
//	 unchanged and are not altered by this method.
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

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigI'",
		}
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		bNum, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numSeps := NumericSeparatorDto{}

	numSeps.DecimalSeparator = bNum.decimalSeparator

	numSeps.ThousandsSeparator = bNum.thousandsSeparator

	numSeps.CurrencySymbol = bNum.currencySymbol

	new(bigIntNumElectron).empty(bNum)

	bNum.bigInt.Set(bigI)

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

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
				"    bNum, numSeps, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).isBigIntNumValid(\n" +
				"    bNum, ePrefix)",
			ErrContext: "Final bNum Validation",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setBigIntNumSeps
//
// Sets the value of the BigIntNum instance passed as input
// parameter 'bNum'. The new value for 'bNum' is extracted from
// input parameters bigI (*big.Int) and precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
//	 Input Parameters
//	 ===============
//
//	 bNum                     *BigIntNum
//	   This instance of BigIntNum will be reconfigured using
//	   the following input parameters.
//
//	 bigI                     *big.Int
//	  'bigI' is a type *big.Int and represents the integer value
//	  of the number; that is, the numeric value without decimal
//	  digits.
//
//	 precision                uint
//	   This unsigned integer (always a positive value) identifies
//	   the location of the decimal place in the integer value
//	   'bigI'. The decimal place location is calculated by
//	    starting with the right most digit in the integer number
//	    and counting	left, 'precision' places. Example:
//
//			   Integer Value    precision    Numeric Value
//			     123456             3           123.456
//		                 123456 x 10^-3 =  123.456
//
//	 numSepsDto               NumericSeparatorDto
//	   Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	   instance. A NumericSeparatorDto contains symbols or characters
//	   for the decimal separator, thousands separator and currency
//	   symbol. These separators are used when parsing number strings
//	   into numeric values or displaying numeric values in number
//	   strings.
//
//	   If any of the 'numSeps' Numeric Separator Components are
//	   invalid, those components will be automatically reset to USA
//	   default values.
//
//	   The input parameter ('bNum') will be reconfigured with
//	   Numeric Separators provided by 'numSepsDto'.
func (bIntNumNano *bigIntNumNanobot) setBigIntNumSeps(
	bNum *BigIntNum,
	bigInt *big.Int,
	precision uint,
	numSepsDto NumericSeparatorDto,
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
		"bigIntNumNanobot.setBigIntNumSeps",
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

	if bigInt == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigInt'",
		}
	}

	new(bigIntNumElectron).empty(bNum)

	bNum.bigInt = big.NewInt(0).Set(bigInt)

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

	numSepsDto.SetDefaultsIfEmpty()

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSepsDto,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
				"    bNum, numSepsDto, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
