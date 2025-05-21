package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

type bigIntNumUtility struct {
	lock *sync.Mutex
}

// getBigIntNumOfDigits
//
// Returns the number of digits in the numeric value of the
// current BigIntNum instance. The count only includes numeric
// digits and as such, EXCLUDES number signs ('-' or '+' ),
// thousands separators (',') and decimal separators.
//
// Examples:
// =========
//
//	Result=
//
// Numeric String          Number of
//
//	Value                Numeric Digits
//
// =============           ==============
//
//	        123.45								5
//	  1,234,567                  7
//	 -1,234,567.8				 			  8
//	          0									1
//		         0.00               1
//		       012.34               4
//		         0.1234						  4
//		         0.123400						6
//		         0.0123400					6
//		 1,234,567.800						 10
//		         5                  1
//
//	 The returned integer number will always be a positive number.
//	 Also, GetActualNumberOfDigits() will be faster for larger
//	 numbers.
//
//		NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIngNumUtil *bigIntNumUtility) getBigIntNumOfDigits(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIngNumUtil.lock == nil {
		bIngNumUtil.lock = new(sync.Mutex)
	}

	bIngNumUtil.lock.Lock()

	defer bIngNumUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumElectron.getBigIntNumOfDigits()",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())

		return 0, err
	}

	scratchNum := big.NewInt(0).Set(bNum.absBigInt)
	baseZero := big.NewInt(0)
	digitCnt := 0

	if scratchNum.Cmp(baseZero) == 0 {

		digitCnt = 1

		return digitCnt, nil
	}

	baseTen := big.NewInt(10)

	for scratchNum.Cmp(baseZero) == 1 {
		scratchNum = big.NewInt(0).Quo(scratchNum, baseTen)
		digitCnt++
	}

	return digitCnt, nil
}

// bigIntNumChangeSign - Changes the sign of the current BigIntNum value.
//
// If the value of BigIntNum is zero, the sign will remain unchanged
// and this method will return with no action taken.
//
// If the sign of the current BigIntNum value is positive (+), the sign
// will be changed to negative (-). Likewise, if the current sign is
// negative (-), the sign will be changed to positive (+).
//
//	NOTE:
//
// ================
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIngNumUtil *bigIntNumUtility) bigIntNumChangeSign(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIngNumUtil.lock == nil {
		bIngNumUtil.lock = new(sync.Mutex)
	}

	bIngNumUtil.lock.Lock()

	defer bIngNumUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumElectron.bigIntNumChangeSign()",
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

	var isZero bool

	isZero, err = new(bigIntNumMolecule).isBIntNumZero(
		bNum,
		ePrefix)

	if err != nil {

		return err
	}

	if isZero {

		bNum.sign = 1

		return nil
	}

	bNum.bigInt = big.NewInt(0).Neg(bNum.bigInt)

	if bNum.bigInt.Cmp(big.NewInt(0)) == -1 {

		bNum.sign = -1

	} else {

		bNum.sign = 1

	}

	return nil
}

// CopyIn - Receives an incoming BigIntNum type and
// copies the value into the current BigIntNum instance.
//
//	NOTE:
//
// ================
//
// This method does NOT test the validity of 'bNumDestination'
// BigIntNum instance. The calling method must do this!
func (bIngNumUtil *bigIntNumUtility) bigIntNumCopyIn(
	bNumDestination *BigIntNum,
	bNumSource *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIngNumUtil.lock == nil {
		bIngNumUtil.lock = new(sync.Mutex)
	}

	bIngNumUtil.lock.Lock()

	defer bIngNumUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumUtility.bigIntNumCopyIn()",
		"")

	if err != nil {
		return err
	}

	if bNumDestination == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNumDestination'",
		}
	}

	if bNumSource == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNumSource'",
		}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNumSource,
		ePrefix.XCpy(" Validating bNumSource"))

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error: The source BigIntNum for the CopyIn operation is invalid.\n"+
			"Error: %v",
			ePrefix,
			err)
	}

	bNumDestination.bigInt = big.NewInt(0).Set(bNumSource.bigInt)

	bNumDestination.absBigInt = big.NewInt(0).Set(bNumSource.absBigInt)

	bNumDestination.precision = bNumSource.precision

	bNumDestination.scaleFactor = big.NewInt(0).Set(bNumSource.scaleFactor)

	bNumDestination.numberOfExpectedDigits = big.NewInt(0).Set(bNumSource.numberOfExpectedDigits)

	bNumDestination.sign = bNumSource.sign

	bNumDestination.decimalSeparator = bNumSource.decimalSeparator

	bNumDestination.thousandsSeparator = bNumSource.thousandsSeparator

	bNumDestination.currencySymbol = bNumSource.currencySymbol

	return nil
}

// CopyOut - Makes a deep copy of the current BigIntNum instance
// and returns it as a new BigIntNum instance.
//
//	NOTE:
//
// ================
// This method tests the validity of the 'bNum'
// BigIntNum instance.
func (bIngNumUtil *bigIntNumUtility) bigIntNumCopyOut(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIngNumUtil.lock == nil {
		bIngNumUtil.lock = new(sync.Mutex)
	}

	bIngNumUtil.lock.Lock()

	defer bIngNumUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumUtility.bigIntNumCopyOut()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy(" - bNum"))

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error: The source BigIntNum for the CopyOut operation is invalid.\n"+
				"Error: %v",
				ePrefix.String(),
				err)
	}

	b2, err := new(bigIntNumMechanics).newBigInt(
		big.NewInt(0).Set(bNum.bigInt),
		bNum.precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	b2.decimalSeparator = bNum.decimalSeparator

	b2.thousandsSeparator = bNum.thousandsSeparator

	b2.currencySymbol = bNum.currencySymbol

	b2.numberOfExpectedDigits =
		big.NewInt(0).Set(bNum.numberOfExpectedDigits)

	return b2, nil
}

// setBigIntBigPrecision
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated *big.Int precision.
//
// The 'precision' parameter specifies the number of digits to the
// right of the decimal place. The Numeric value is equal to:
//
//	bigI x 10^(precision x -1)
//
// This effectively locates the decimal place by counting from the
// extreme right of the integer number, 'precision' places to the
// left. See the example below.
//
//	Input Parameters
//	================
//
//	bNum          *BigIntNum
//	An instance of BigIntNum which will be reconfigured
//	according to values prvovided by input parameters 'bigI'
//	and 'precision'.
//
//	bigInt        *big.Int
//
//	'bigI' is a type *big.Int and represents the integer value of
//	the number; that is, the numeric value without decimal digits.
//
//
//	precision     *big.In
//
//	This integer value (always a positive value) identifies the
//	location of the decimal place in the integer value 'bigI'. The
//	decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places.
//
//	  Integer Value    precision    Numeric Value
//
//		    123456           3            123.456
//	             123456 x 10^-3  =    123.456
//
//	If precision is greater than the maximum value of an unsigned
//	integer (+4,294,967,295,	which equals 2^32 − 1), an error will
//	be triggered. Also, if the 'precision' value is less than zero,
//	an error will be triggered.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators specify the symbols or characters (runes)
//	used for the decimal separator, thousands separator and
//	currency symbol. These separators are used when displaying
//	numeric values in number strings.
//
//	The original Numeric Separators configured in 'bNum' and will
//	not be altered by this method. However, if any of the Numeric
//	Separator values are invalid (set to rune value of zero),
//	those Numeric Separators will be automatically reset to USA
//	defaults.
func (bIngNumUtil *bigIntNumUtility) setBigIntBigPrecision(
	bNum *BigIntNum,
	bigInt *big.Int,
	precision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIngNumUtil.lock == nil {
		bIngNumUtil.lock = new(sync.Mutex)
	}

	bIngNumUtil.lock.Lock()

	defer bIngNumUtil.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumUtility.setBigIntBigPrecision",
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

	if precision == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'precision'",
		}
	}

	if precision.Cmp(big.NewInt(0)) == -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("precision='%v'",
				precision.Text(10)),
			ErrMessage: "Error: Input parameter 'precision' IS LESS THAN ZERO!",
		}
	}

	maxUint32 := big.NewInt(0).SetUint64(uint64(math.MaxUint32))

	if precision.Cmp(maxUint32) == 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("precision='%v'",
				precision.Text(10)),
			ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' exceeds maximum limit of '%v'!",
				maxUint32.Text(10)),
		}
	}

	uintPrecision := uint(precision.Uint64())

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		bigInt,
		uintPrecision,
		ePrefix.XCpy(
			fmt.Sprintf("Setting 'bNum' 'bigI'='%v' precision= '%v'",
				bigInt.Text(10), uintPrecision)))

	return err
}

// setBigIntBigPrecisionNumSeps
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated precision (also of type *big.Int).
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigInt x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
//		 Precision Example:
//		 ==================
//
//		     Integer Value    precision     Numeric Value
//		 			  123456					 3					  123.456
//		 	             123456 x 10^-3  =    123.456
//
//		 Numeric Seprators
//		 =================
//
//		 The returned BigIntNum instance will be configured with the
//		 Numeric Separators provided by input parameter 'numSeps'.
//		 Numeric Separatpors consist of decimal separators, thousands
//		 separators and a currency symbol.
//
//		 Input Parameters
//		 ================
//
//		 bNum                     *BigIntNum
//		   An instance of BigIntNum which will be reconfigured
//			  according to values prvovided by input parameters 'bigInt'
//			  'precision' and 'numSeps'.
//
//
//		 bigInt                   *big.Int
//		   'bigInt' is a type *big.Int and represents the integer
//				value of the number; that is, the numeric value without
//				decimal digits.
//
//
//		 precision                *big.Int
//		   This integer value (always a positive value) identifies
//		   the location of the decimal place in the integer value 'bigInt'.
//		   The decimal place location is calculated by starting with the
//		   right most digit in the integer number and counting	left,
//		   'precision' places. If precision is greater than the maximum
//		   value of an unsigned integer (+4,294,967,295,	which equals
//		   2^32 − 1), an error will be triggered. Also, if the 'precision'
//		   value is less than zero, an error will be triggered.
//
//		     Integer Value		precision			Numeric Value
//		       123456					 3					  123.456
//		                123456 x 10^-3  =    123.456
//
//
//		 numSeps			NumericSeparatorDto
//		   Numeric Separators specify the symbols or characters (runes)
//		   used for the decimal separator, thousands separator and
//		   currency symbol. These separators are used when displaying
//		   numeric values in number strings.
//
//		   BigIntNum instance, 'bNum', will be reconfigured with the
//		   Numeric Separators provided by 'numSeps'.
//
//	    If any of the 'numSep' values are invalid, they will be
//	    reset to USA default values.
//
//		 Return Parameters
//		 =================
//
//		 error
//		   If no errors are encountered during execution, this method
//		   return an error value of 'nil'.
func (bIngNumUtil *bigIntNumUtility) setBigIntBigPrecisionNumSeps(
	bNum *BigIntNum,
	bigInt *big.Int,
	precision *big.Int,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIngNumUtil.lock == nil {
		bIngNumUtil.lock = new(sync.Mutex)
	}

	bIngNumUtil.lock.Lock()

	defer bIngNumUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumUtility.setBigIntBigPrecisionNumSeps",
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

	if precision == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'precision'",
		}
	}

	numSeps.SetDefaultsIfEmpty()

	if precision.Cmp(big.NewInt(0)) == -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("precision= '%v'", precision.Text(10)),
			ErrMessage: "Error: Input parameter 'precision' IS LESS THAN ZERO!",
		}
	}

	maxUint32 := big.NewInt(0).SetUint64(uint64(math.MaxUint32))

	if precision.Cmp(maxUint32) == 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("precision= '%v'",
				precision.Text(10)),
			ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' exceeds maximum limit of '%v' !",
				maxUint32.Text(10)),
		}
	}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		bNum,
		bigInt,
		uint(precision.Uint64()),
		numSeps,
		ePrefix.XCpy(fmt.Sprintf("Setting bNum bigInt= '%v'  precision= '%v'",
			bigInt.Text(10), precision.Text(10))))

	return err
}
