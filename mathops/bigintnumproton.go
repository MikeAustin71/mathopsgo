package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type bigIntNumProton struct {
	lock *sync.Mutex
}

// Cmp - Performs a comparison of two BigIntNum numeric values
// and returns an integer value indicating the relationship
// between the two numeric values (i.e. greater than, less than,
// or equal).
//
// Note: Unlike method CmpBigInt() below, this method does more than
// just compare the root *big.Int. In making the comparision, this
// method takes into account, numeric sign values and precision. Therefore,
// this method effectively compares numeric values. As such, this method
// provides a true and comprehensive picture of the relationship between
// two BigIntNum values.
//
// Return Values:
// bNum == bigIntNum 				Return  0
// bNum > bigIntNum					Return +1
// bNum < bigIntNum					Return -1
func (bIntNumProton *bigIntNumProton) bigIntNumCmp(
	bNum *BigIntNum,
	bigIntNum2 *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumCmp",
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

	if bigIntNum2 == nil {

		err = fmt.Errorf("%v\n"+
			"FATAL ERROR: Input parameter 'bigIntNum' is a nil pointer.\n",
			ePrefix.String())

		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bigIntNum2,
		ePrefix.XCpy(" Validating Input Parameter 'bigIntNum'"))

	if err != nil {
		return 0, err
	}

	b1Sign := bNum.sign
	b2Sign := bigIntNum2.sign

	if b1Sign != b2Sign {

		if b1Sign > b2Sign {

			return b1Sign, nil
		}

		return b2Sign, nil
	}

	// The signs must be equal

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return 0,
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				"  bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(bNum, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	difference := BigIntMathSubtract{}.SubtractBigIntNums(
		bNum2,
		*bigIntNum2)

	var differenceIsZero bool

	differenceIsZero, err = new(bigIntNumMolecule).isBIntNumZero(
		&difference,
		ePrefix)

	if err != nil {

		return 0, err
	}

	if differenceIsZero {
		return 0, nil
	}

	return difference.GetSign(), nil
}
