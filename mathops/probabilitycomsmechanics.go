package mathops

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type probabilityComsMechanics struct {
	lock *sync.Mutex
}

func (probComsMech *probabilityComsMechanics) combinationsNoRepsBigInt(
	numOfItems *big.Int,
	numOfItemsChosen *big.Int,
	numSeps NumericSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if probComsMech.lock == nil {
		probComsMech.lock = new(sync.Mutex)
	}

	probComsMech.lock.Lock()

	defer probComsMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"probabilityComsMechanics.combinationsNoRepsBigInt()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

}
