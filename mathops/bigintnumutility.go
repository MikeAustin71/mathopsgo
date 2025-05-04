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

// CopyIn - Receives an incoming BigIntNum type and
// copies the value into the current BigIntNum instance.
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
		"bigIntNumElectron.resetBigIntNum()",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNumSource,
		ePrefix.XCpy(" - bNumSource"))

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
		"bigIntNumElectron.resetBigIntNum()",
		"")

	if err != nil {
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
