package mathops

import (
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathAddMicrobot struct {
	lock *sync.Mutex
}

// addPair
//
//	Receives a BigIntPair instance and proceeds to add b1.BigIntNum
//	to b2.BigIntNum.
//
//	The result is returned as type BigIntNum.
//
//	The BigIntNum 'result' returned by this addition operation will
//	contain numeric separators (decimal separator, thousands separator
//	and currency symbol) copied from b1.BigIntNum.
func (bIMathAddMicro *bigIntMathAddMicrobot) addPair(
	bPair BigIntPair,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIMathAddMicro.lock == nil {
		bIMathAddMicro.lock = new(sync.Mutex)
	}

	bIMathAddMicro.lock.Lock()

	defer bIMathAddMicro.lock.Unlock()

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

	numSeps, err := bPair.Big1.GetNumericSeparatorsDto()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := bPair.Big1.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	err = numSeps.IsValid("Validating 'numSeps'")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(\"Validating 'numSeps'\")",
				ErrContext: "'numSeps' is INVALID!'\n" +
					"'numSeps' was extracted from bPair.Big1",
				ErrMessage: err.Error(),
			}
	}

	finalResult, err := new(bigIntMathAddNanobot).addPairNoNumSeps(bPair, ePrefix)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "finalResult, err := new(bigIntMathAddNanobot).\n" +
					"  addPairNoNumSeps(bPair, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return finalResult,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = finalResult.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return finalResult, nil
}
