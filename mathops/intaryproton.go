package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryProton struct {
  lock *sync.Mutex
}

// copy
//
// Makes a deep copy of the current IntAry instance with backup and
// returns it as a new IntAry object.
func (iaProton *intAryProton) copy(
  iaDestination *IntAry,
  iaSource *IntAry,
  validateSourceIntAry bool,
  copyBackup bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaProton.lock.Lock()

  defer iaProton.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryProton.copy",
    "")

  if err != nil {
    return err
  }

  if iaDestination == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'iaDestination'",
    }
  }

  if iaSource == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'iaSource'",
    }
  }

  if validateSourceIntAry {

    err = new(intAryElectron).isValidIntAry(
      iaSource, ePrefix.XCpy("Validating 'iaSource'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(ia, ePrefix.XCpy(Validating 'ia').String())",
        ErrContext: "IntAry instanace 'iaSource' is INVALID!\n" +
          "'iaSource' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
    }
  }

  err = new(intAryNanobot).setInternalFlags(
    iaSource, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(ia, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  new(intAryElectron).empty(iaDestination)

  iaDestination.intAry = make([]uint8, iaSource.intAryLen)

  for i := 0; i < iaSource.intAryLen; i++ {
    iaDestination.intAry[i] = iaSource.intAry[i]
  }

  iaDestination.intAryLen = iaSource.intAryLen
  iaDestination.integerLen = iaSource.integerLen
  iaDestination.significantIntegerLen = iaSource.significantIntegerLen
  iaDestination.significantFractionLen = iaSource.significantFractionLen
  iaDestination.firstDigitIdx = iaSource.firstDigitIdx
  iaDestination.lastDigitIdx = iaSource.lastDigitIdx
  iaDestination.isZeroValue = iaSource.isZeroValue
  iaDestination.isIntegerZeroValue = iaSource.isIntegerZeroValue
  iaDestination.precision = iaSource.precision
  iaDestination.signVal = iaSource.signVal
  iaDestination.decimalSeparator = iaSource.decimalSeparator
  iaDestination.thousandsSeparator = iaSource.thousandsSeparator
  iaDestination.currencySymbol = iaSource.currencySymbol

  if copyBackup {

    err = new(intAryNeutron).copyToBackup(
      iaDestination,
      iaSource,
      false,
      ePrefix)

  }

  err = new(intAryNeutron).copyToBackup(
    iaDestination,
    iaSource,
    false,
    ePrefix)

  return err
}
