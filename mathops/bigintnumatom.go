package mathops

import (
	"fmt"
	"math/big"
	"sync"
)

type bigIntNumAtom struct {
	lock *sync.Mutex
}

// isBigIntNumValid - returns a boolean value signaling
// whether the current BigIntNum object is valid.
func (bIntNumAtom *bigIntNumAtom) isBigIntNumValid(
	bNum *BigIntNum,
	callingMethodChain string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "Active Method: bigIntNumAtom.isBigIntNumValid()"

	var err error

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if bNum.bigInt == nil {

		bNum.bigInt = big.NewInt(0)

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.bigInt' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum.bigInt was reset to zero.\n",
			ePrefix)

	}

	if bNum.sign != -1 && bNum.sign != 1 {

		err = bNum.Reset()

		if err != nil {
			return fmt.Errorf("%v\n"+
				"This BigIntNum Instance is Invalid!\n"+
				"'bNum.sign' is NOT equal to +1 or -1 !\n"+
				"FATAL ERROR!\n"+
				"Attmpted reset of bNum to default values FAILED!.\n",
				ePrefix)

		}

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.sign' is NOT equal to +1 or -1 !\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix)
	}

	if bNum.absBigInt == nil {

		err = bNum.Reset()

		if err != nil {

			return fmt.Errorf("%v\n"+
				"This BigIntNum Instance is Invalid!\n"+
				"'bNum.absBigInt' is 'nil'!\n"+
				"FATAL ERROR!\n"+
				"Attmpted reset of bNum to default values FAILED!.\n",
				ePrefix)

		}

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.absBigInt' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix)
	}

	if bNum.scaleFactor == nil {

		err = bNum.Reset()

		if err != nil {
			return fmt.Errorf("%v\n"+
				"This BigIntNum Instance is Invalid!\n"+
				"'bNum.scaleFactor' is 'nil'!\n"+
				"FATAL ERROR!\n"+
				"Attmpted reset of bNum to default values FAILED!.\n",
				ePrefix)

		}

		return fmt.Errorf("%v\n"+
			"This BigIntNum Instance is Invalid!\n"+
			"'bNum.scaleFactor' is 'nil'!\n"+
			"FATAL ERROR!\n"+
			"bNum was successfully reset to default values (Zero).\n",
			ePrefix)

	}

	return nil
}

// setNumericSeparatorsDto - Sets the values of numeric separators:
//
//	decimal place separator
//	thousands separator
//	currency symbol
//
// based on values transmitted through input parameter 'customSeparators'.
//
// If any of the values contained in input parameter 'customSeparators' is set
// to zero, an error will be returned.
//
// NOTE:
// This is a low-level operation. It is assumed that 'bNum'
// has already been validated.
func (bIntNumAtom *bigIntNumAtom) setNumericSeparatorsDto(
	bNum *BigIntNum,
	customSeparators NumericSeparatorDto,
	callingMethodChain string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "Active Method: BigIntNum.SetNumericSeparatorsDto() "

	if len(callingMethodChain) > 0 {
		ePrefix = ePrefix + "\nCalling Method Chain:\n " + callingMethodChain
	}

	if bNum == nil {

		return fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix)
	}

	if customSeparators.DecimalSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.DecimalSeparator' is set to '0' - Invalid rune!\n",
			ePrefix)
	}

	if customSeparators.ThousandsSeparator == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.ThousandsSeparator' is set to '0' - Invalid rune!\n",
			ePrefix)

	}

	if customSeparators.CurrencySymbol == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'customSeparators.CurrencySymbol' is set to '0' - Invalid rune!\n",
			ePrefix)

	}

	bNum.decimalSeparator = customSeparators.DecimalSeparator

	bNum.thousandsSeparator = customSeparators.ThousandsSeparator

	bNum.currencySymbol = customSeparators.CurrencySymbol

	return nil
}
