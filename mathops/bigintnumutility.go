package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())

		return err
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

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNumDestination' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if bNumSource == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNumSource' is a nil pointer.\n",
			ePrefix.String())

		return err
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

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())

		return BigIntNum{}, err
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
