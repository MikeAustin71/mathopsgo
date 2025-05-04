package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumElectron struct {
	lock *sync.Mutex
}

// Empty - Resets the BigIntNum data fields to their
// uninitialized or zero state.
func (bINumElectron *bigIntNumElectron) empty(
	bNum *BigIntNum) {

	if bINumElectron.lock == nil {
		bINumElectron.lock = new(sync.Mutex)
	}

	bINumElectron.lock.Lock()

	defer bINumElectron.lock.Unlock()

	if bNum == nil {
		return
	}

	bNum.bigInt = big.NewInt(0)

	bNum.absBigInt = big.NewInt(0)

	bNum.scaleFactor = big.NewInt(1)

	bNum.numberOfExpectedDigits = big.NewInt(0)

	bNum.sign = 1

	bNum.precision = 0

	bNum.decimalSeparator = '.'

	bNum.thousandsSeparator = ','

	bNum.currencySymbol = '$'

}

// resetBigIntNum - Resets the current BigIntNum to a new
// valid BigIntNum using the BigIntNum components
// BigIntNum.bigInt and BigIntNum.precision. This
// method is usually called after method bNum.IsValid()
// returns false.
func (bINumElectron *bigIntNumElectron) resetBigIntNum(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bINumElectron.lock == nil {
		bINumElectron.lock = new(sync.Mutex)
	}

	bINumElectron.lock.Lock()

	defer bINumElectron.lock.Unlock()

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

	bINumNanobot := new(bigIntNumNanobot)

	if bNum.bigInt == nil {

		err = bINumNanobot.setBigInt(
			bNum,
			big.NewInt(0),
			uint(0),
			ePrefix)

		if err != nil {

			return err

		}

		return nil
	}

	if bNum.sign != 1 && bNum.sign != -1 {

		newNum := big.NewInt(0).Set(bNum.bigInt)

		err = bINumNanobot.setBigInt(
			bNum,
			newNum,
			bNum.precision,
			ePrefix)

		if err != nil {

			return err

		}

		return nil
	}

	if bNum.absBigInt == nil || bNum.scaleFactor == nil {

		newNum := big.NewInt(0).Set(bNum.bigInt)

		err = bINumNanobot.setBigInt(
			bNum,
			newNum,
			bNum.precision,
			ePrefix)

		if err != nil {

			return err

		}
	}

	return nil
}
