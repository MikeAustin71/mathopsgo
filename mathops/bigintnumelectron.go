package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumElectron struct {
	lock *sync.Mutex
}

// bigIntNumCeiling - Returns the ceiling integer value of the current BigIntNum
// instance.
//
// Ceiling is defined as: The least, or lowest value integer, which is greater
// than or equal to the numeric value of the current BigIntNum.
// Reference Wikipedia:
//
//	https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
// Examples
// ========
//
//							Initial 		 Ceiling
//	 					 Value				Value
//							-------      -------
//	 						5.95					6
//	 						5.05					6
//	 						5							5
//						 -5.05			 	 -5
//	 						2.4				  	3
//	 						2.9					 	3
//						 -2.7				 	 -2
//						 -2						 -2
//
//	NOTE:
//
// ================
// This method does NOT test the validity of 'bNumDestination'
// BigIntNum instance. The calling method must do this!
func (bINumElectron *bigIntNumElectron) bigIntNumCeiling(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

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
		"bigIntNumElectron.bigIntNumCeiling()",
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

	isbNumZero, err := new(bigIntNumMolecule).isBIntNumZero(
		bNum,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	if isbNumZero {

		bInt2, err := new(BigIntNum).NewBigInt(big.NewInt(0), 0)

		if err != nil {

			return BigIntNum{}, err
		}

		return bInt2, nil
	}

	// bNum is NOT Zero
	if bNum.precision == 0 {

		bInt3, err := new(bigIntNumUtility).bigIntNumCopyOut(
			bNum,
			ePrefix)

		if err != nil {

			return BigIntNum{}, err
		}

		return bInt3, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision)), nil)

	absQuotient := big.NewInt(0).Quo(bNum.absBigInt, scaleVal)

	// absQuotient IS NOT EQUAL TO bNum.absBigInt

	bigINumMech := new(bigIntNumMechanics)

	if bNum.sign > 0 {
		// bNum is positive

		absQuotient = big.NewInt(0).Add(absQuotient, big.NewInt(1))

		bInt4, err := bigINumMech.newBigInt(
			absQuotient, 0, ePrefix)

		if err != nil {

			return BigIntNum{}, err

		}

		return bInt4, nil
	}

	// bNum is negative
	bInt5, err := bigINumMech.
		newBigInt(big.NewInt(0).Neg(absQuotient),
			0, ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	return bInt5, nil
}

// bigIntNumEqual - Compares two BigIntNum instances and
// returns 'true' if the two instances are equal in all
// respects.
//
// Be careful, two BigIntNum instances could have equal
// values with different precisions. In that case this
// method would return 'false'. To test for equivalent
// values, see method BigIntNum.EqualValue(), below.
//
// If they are not Equal, the method returns 'false'.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bINumElectron *bigIntNumElectron) bigIntNumEqual(
	bNum *BigIntNum,
	b2 *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

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
		return false, err
	}

	if bNum == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())

		return false, err
	}

	if b2 == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'b2' is a nil pointer.\n",
			ePrefix.String())

		return false, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		b2,
		ePrefix.XCpy("Validating input parameter 'b2'"))

	if err != nil {
		return false, err
	}

	if bNum.bigInt.Cmp(b2.bigInt) != 0 {
		return false, nil
	}

	if bNum.absBigInt.Cmp(b2.absBigInt) != 0 {
		return false, nil
	}

	if bNum.scaleFactor.Cmp(b2.scaleFactor) != 0 {
		return false, nil
	}

	if bNum.sign != b2.sign {
		return false, nil
	}

	if bNum.precision != b2.precision {
		return false, nil
	}

	return true, nil
}

// EqualValue - Compares the values of the current BigIntNum instance
// and the input parameter BigIntNum, 'b2'. If the two numeric values
// are equal, this method returns 'true'.
func (bINumElectron *bigIntNumElectron) bigIntNumEqualValue(
	bNum *BigIntNum,
	b2 *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

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
		return false, err
	}

	if bNum == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())

		return false, err
	}

	if b2 == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'b2' is a nil pointer.\n",
			ePrefix.String())

		return false, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		b2,
		ePrefix.XCpy("Validating input parameter 'b2'"))

	if err != nil {
		return false, err
	}

	bNum3, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {
		return false, err
	}

	difference, err := BigIntMathSubtract{}.SubtractBigIntNums(
		bNum3,
		*b2)

	if err != nil {

		return false,
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" difference, err := BigIntMathSubtract{}.SubtractBigIntNums(\n"+
				"   bNum3, *b2\n"+
				"Error= %v\n",
				ePrefix.String(),
				err.Error())
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		&difference,
		ePrefix)

	if err != nil {

		return false,
			fmt.Errorf("%v\n"+
				"Validation Error on current Decimal instance 'difference'.\n"+
				"'difference' is a calculated variable which failed the valiadtion test.\n"+
				"Error returned by: \n"+
				" err := new(bigIntNumAtom).isBigIntNumValid(&difference,ePrefix)\n"+
				"Validation Error= %v\n",
				ePrefix,
				err.Error())

	}

	if difference.bigInt.Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}

	return false, nil
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
//
//	NOTE
//	====
//
//	The newly configured instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol).
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

	if bNum == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bNum' is a nil pointer.\n",
			ePrefix.String())

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
