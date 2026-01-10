package mathops

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathAddNanobot struct {
	lock *sync.Mutex
}

// addPairNoNumSeps
//
//	Receives a BigIntPair and proceeds to add the values
//	b1.BigIntNum to b2.BigIntNum.
//
//	The result of this addition is returned as a type BigIntNum.
//
//	The BigIntNum 'result' returned by this addition operation
//	will contain default numeric separators (decimal separator,
//	thousands separator and currency symbol).
func (bIMathAddNano *bigIntMathAddNanobot) addPairNoNumSeps(
	bPair BigIntPair,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIMathAddNano.lock == nil {
		bIMathAddNano.lock = new(sync.Mutex)
	}

	bIMathAddNano.lock.Lock()

	defer bIMathAddNano.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathAddNanobot.addPairNoNumSeps",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = bPair.IsValid("Validating 'bPair'")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bPair.IsValid(\"Validating 'bPair'\")",
				ErrContext: "Input parameter 'bPair' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	err = bPair.MakePrecisionsEqual()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err := bPair.MakePrecisionsEqual()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigI1, err := bPair.GetBig1BigInt()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigI1, err := bPair.GetBig1BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigI2, err := bPair.GetBig2BigInt()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigI2, err := bPair.GetBig2BigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	b3 := big.NewInt(0).Add(bigI1, bigI2)

	big2Precision, err := bPair.Big2.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "big2Precision, err := bPair.Big2.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bResult, err := new(BigIntNum).NewBigInt(b3, big2Precision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bResult, err := new(BigIntNum).NewBigInt(b3, big2Precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = bResult.IsValid("Validating 'bResult'")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = bResult.IsValid(\"Validating 'bResult'\")",
				ErrContext: "Final result 'bResult' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	return bResult, nil
}
