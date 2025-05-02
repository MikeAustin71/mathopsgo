package mathops

import (
	"fmt"
	"math/big"
	"sync"
)

type bigIntMathMultiplyElectron struct {
	lock *sync.Mutex
}

// New - Creates a BigIntMathMultiply instance with data
// variables initialized to zero.
func (bIMathMultiplyElec *bigIntMathMultiplyElectron) newBigIntMathMultiply() (BigIntMathMultiply, error) {

	ePrefix := "bigIntMathMultiplyElectron.newBigIntMathMultiply()"

	if bIMathMultiplyElec.lock == nil {
		bIMathMultiplyElec.lock = new(sync.Mutex)
	}

	bIMathMultiplyElec.lock.Lock()

	defer bIMathMultiplyElec.lock.Unlock()

	b2Math := new(BigIntMathMultiply)

	b2Math.Input = new(BigIntPair).New()

	baseZero := big.NewInt(0)

	var err error

	b2Math.Result, err = new(BigIntNum).NewBigInt(baseZero, 0)

	if err != nil {

		return BigIntMathMultiply{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" b2Math.Result, err = new(BigIntNum).NewBigInt(baseZero, 0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

	}

	return *b2Math, nil
}
