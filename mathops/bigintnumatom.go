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
	callingMethodName string) error {

	if bIntNumAtom.lock == nil {
		bIntNumAtom.lock = new(sync.Mutex)
	}

	bIntNumAtom.lock.Lock()

	defer bIntNumAtom.lock.Unlock()

	ePrefix := "bigIntNumAtom.IsValid()"

	var err error

	if len(callingMethodName) > 0 {
		ePrefix = "Active Method: " + ePrefix + "\nCalling Method Chain:\n " + callingMethodName
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
