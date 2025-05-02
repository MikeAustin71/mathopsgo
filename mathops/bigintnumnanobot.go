package mathops

import (
	"fmt"
	"math/big"
	"sync"
)

type bigIntNumNanobot struct {
	lock *sync.Mutex
}

// SetBigInt - Sets the value of the current BigIntNum instance using
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
	callingMethodChain string) error {

	if bIntNumNano.lock == nil {
		bIntNumNano.lock = new(sync.Mutex)
	}

	bIntNumNano.lock.Lock()

	defer bIntNumNano.lock.Unlock()

	ePrefix := "bigIntNumNanobot.setBigInt"

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if bigI == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'bigI' is a nil pointer!\n",
			ePrefix)

	}

	numSeps := bNum.GetNumericSeparatorsDto()

	bNum.Empty()

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

	err := bNum.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned from:\n"+
			"bNum.SetNumericSeparatorsDto(numSeps)\n"+
			"Error= %v\n",
			ePrefix,
			err.Error())
	}

	return err
}
