package mathops

import (
  "sync"

  ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathAddMacrobot struct {
  lock *sync.Mutex
}

// addBigIntNums
//
//	Adds two BigIntNums and returns the result in a new BigIntNum
//	instance.
//
//	The BigIntNum 'result' returned by this addition operation will
//	contain numeric separators (decimal separator, thousands
//	separator and currency symbol) which were copied from input
//	parameter 'b1'.
func (bIMathAddMacro *bigIntMathAddMacrobot) addBigIntNums(
  b1 BigIntNum,
  b2 BigIntNum,
  errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

  if bIMathAddMacro.lock == nil {
    bIMathAddMacro.lock = new(sync.Mutex)
  }

  bIMathAddMacro.lock.Lock()

  defer bIMathAddMacro.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathAddMacrobot.addBigIntNums",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  // NewBigIntNum validates b1 and b2
  bPair, err := new(BigIntPair).NewBigIntNum(b1, b2)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(b1, b2)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  addResult, err := new(bigIntMathAddMicrobot).addPair(bPair, ePrefix)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(b1, b2)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return addResult, nil
}
