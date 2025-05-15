package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type bigIntFixedDecAtom struct {
  lock *sync.Mutex
}

// isValid
//
// This method performs validity testing on an instance of
// BigIntFixedDecimal passed as input parameter, 'bigIFxDec'.
//
// If this BigIntFixedDecimal instance fails the validity
// tests, an error will be returned.
//
// If this BigIntFixedDecimal instance passes all validity
// tests, an error value of 'nil' will be returned.
func (bigIFdAtom *bigIntFixedDecAtom) isBigIntFxDecValid(
  bigIFxDec *BigIntFixedDecimal,
  errPrefDto *ePref.ErrPrefixDto) error {

  if bigIFdAtom.lock == nil {
    bigIFdAtom.lock = new(sync.Mutex)
  }

  bigIFdAtom.lock.Lock()

  defer bigIFdAtom.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntFixedDecAtom.isBigIntFxDecValid",
    "")

  if err != nil {
    return err
  }

  if bigIFxDec == nil {

    err = &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec'",
    }

    return err
  }

  if bigIFxDec.integerNum == nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "Error: This instance of BigIntFixedDecimal is INVALID!\n" +
        "BigIntFixedDecimal.integerNum is a 'nil' pointer.\n" +
        "This BigIntFixedDecimal instance FAILED Validation Testing!",
    }
  }

  if bigIFxDec.decimalSeparator == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "This BigIntFixedDecimal Instance is Invalid!\n" +
        "'bigIFxDec.decimalSeparator' is empty with a Zero value.\n" +
        "FATAL ERROR!",
    }
  }

  if bigIFxDec.thousandsSeparator == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "This BigIntFixedDecimal Instance is Invalid!\n" +
        "'bigIFxDec.thousandsSeparator' is empty with a Zero value.\n" +
        "FATAL ERROR!",
    }
  }

  if bigIFxDec.currencySymbol == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "This BigIntFixedDecimal Instance is Invalid!\n" +
        "'bigIFxDec.currencySymbol' is empty with a Zero value.\n" +
        "FATAL ERROR!",
    }
  }

  return nil
}
