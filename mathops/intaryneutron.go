package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type intAryNeutron struct {
  lock *sync.Mutex
}

func (iaNeutron *intAryNeutron) copyToBackup(
  iaDestination *IntAry,
  iaSource *IntAry,
  validateSourceIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaNeutron.lock == nil {
    iaNeutron.lock = new(sync.Mutex)
  }

  iaNeutron.lock.Lock()

  defer iaNeutron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.getBigInt()",
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
      ParameterName: "'iaDestination'",
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
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(iaSource, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  iaDestination.BackUp.Empty()

  iaDestination.BackUp.intAry = make([]uint8, iaSource.intAryLen)
  for i := 0; i < iaSource.intAryLen; i++ {
    iaDestination.BackUp.intAry[i] = iaSource.intAry[i]
  }

  iaDestination.BackUp.intAryLen = iaSource.intAryLen
  iaDestination.BackUp.integerLen = iaSource.integerLen
  iaDestination.BackUp.significantIntegerLen = iaSource.significantIntegerLen
  iaDestination.BackUp.significantFractionLen = iaSource.significantFractionLen
  iaDestination.BackUp.firstDigitIdx = iaSource.firstDigitIdx
  iaDestination.BackUp.lastDigitIdx = iaSource.lastDigitIdx
  iaDestination.BackUp.isZeroValue = iaSource.isZeroValue
  iaDestination.BackUp.isIntegerZeroValue = iaSource.isIntegerZeroValue
  iaDestination.BackUp.precision = iaSource.precision
  iaDestination.BackUp.signVal = iaSource.signVal
  iaDestination.BackUp.decimalSeparator = iaSource.decimalSeparator
  iaDestination.BackUp.thousandsSeparator = iaSource.thousandsSeparator
  iaDestination.BackUp.currencySymbol = iaSource.currencySymbol

  return nil
}

// getBigInt
//
//		Returns the current value of this intAry object expressed
//		as a signed integer number of type *big.Int.
//
//	 IMPORTANT
//	 =========
//
//	 The calling function is responsible for validating
//	 'ia'.
func (iaNeutron *intAryNeutron) getBigInt(
  ia *IntAry,
  errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

  if iaNeutron.lock == nil {
    iaNeutron.lock = new(sync.Mutex)
  }

  iaNeutron.lock.Lock()

  defer iaNeutron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.getBigInt()",
    "")

  if err != nil {
    return big.NewInt(0), err
  }

  if ia == nil {

    return big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'ia'",
      }
  }

  lenIntAry := len(ia.intAry)

  if lenIntAry != ia.intAryLen {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: " if lenIntAry != ia.intAryLen",
        ErrMessage: "Error: The actual length of 'intAry' does not match ia.intAryLen.\n" +
          "This instance of 'ia' is INVALID!",
      }

  }

  if lenIntAry == 0 {

    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: " if lenIntAry != ia.intAryLen",
        ErrMessage: "Error: The actual length of 'intAry' is ZERO.\n" +
          "This instance of 'ia' is INVALID!",
      }
  }

  result := big.NewInt(0).SetInt64(0)

  big10 := big.NewInt(0).SetInt64(10)

  for i := 0; i < ia.intAryLen; i++ {
    result = big.NewInt(0).Mul(result, big10)
    result = big.NewInt(0).Add(result, big.NewInt(0).SetInt64(int64(ia.intAry[i])))

  }

  if ia.signVal == -1 {

    result = big.NewInt(0).Neg(result)
  }

  return result, nil
}
