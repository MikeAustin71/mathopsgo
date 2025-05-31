package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryMechanics struct {
  lock *sync.Mutex
}

// setInternalFlags - Sets Array Lengths and
// test for zero values
func (iaMech *intAryMechanics) setInternalFlags(
  ia *IntAry,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaMech.lock == nil {
    iaMech.lock = new(sync.Mutex)
  }

  iaMech.lock.Lock()

  defer iaMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMechanics.setInternalFlags()",
    "")

  if err != nil {
    return err
  }

  err = new(intAryElectron).setSignificantDigitIdxs(
    ia,
    ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryElectron).\n" +
        "  setSignificantDigitIdxs( ia, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
