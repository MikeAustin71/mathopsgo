package mathops

import (
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
