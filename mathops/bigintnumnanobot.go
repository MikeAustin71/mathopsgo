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

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())
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
