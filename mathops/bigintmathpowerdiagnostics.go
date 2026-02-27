package mathops

import (
  "sync"

  ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathPowerDiagnostics struct {
  lock *sync.Mutex
}

// profilePowerCalc
//
// Analyzes the components of exponentiation (a power calculation)
// and returns a BigIntMathPowerProfile structure. This structure
// contains information about the components of the power
// calculation.
func (bIntPwrDiag *bigIntMathPowerDiagnostics) profilePowerCalc(
  base *BigIntNum,
  exponent *BigIntNum,
  errPrefDto *ePref.ErrPrefixDto) (BigIntMathPowerProfile, error) {

  if bIntPwrDiag.lock == nil {
    bIntPwrDiag.lock = new(sync.Mutex)
  }

  bIntPwrDiag.lock.Lock()

  defer bIntPwrDiag.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathPowerDiagnostics.profilePowerCalc",
    "")

  if err != nil {
    return BigIntMathPowerProfile{}, err
  }

  if base == nil {

    return BigIntMathPowerProfile{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'base'",
      }
  }

  if exponent == nil {

    return BigIntMathPowerProfile{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'exponent'",
      }
  }

  powerCalcProfile := BigIntMathPowerProfile{}

  // Analyze 'base'

  err = new(bigIntNumAtom).isBigIntNumValid(
    base,
    ePrefix.XCpy("Validating 'base'"))

  if err != nil {

    return BigIntMathPowerProfile{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(bigIntNumAtom).isBigIntNumValid(base, ePrefix)",
        ErrContext: "Input Parameter 'base' is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  // Analyze 'exponent'

  err = new(bigIntNumAtom).isBigIntNumValid(
    exponent,
    ePrefix.XCpy("Validating 'exponent'"))

  if err != nil {

    return BigIntMathPowerProfile{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(bigIntNumAtom).isBigIntNumValid(exponent, ePrefix)",
        ErrContext: "Input Parameter 'exponent' is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  powerCalcProfile.BaseIsZero, err = new(bigIntNumMolecule).isBIntNumZero(
    base,
    ePrefix)

  if err != nil {

    return BigIntMathPowerProfile{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "powerCalcProfile.BaseIsZero, err = new(bigIntNumMolecule).isBIntNumZero(\n" +
          "  base, ePrefix)",
        ErrContext: "Error returned while testing 'base' for zero value!",
        ErrMessage: err.Error(),
      }

  }

  if !powerCalcProfile.BaseIsZero {

    powerCalcProfile.BaseIsOne, err = new(bigIntNumMolecule).isBIntNumAbsOne(
      base,
      ePrefix)

    if err != nil {

      return BigIntMathPowerProfile{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "powerCalcProfile.BaseIsOne, err = new(bigIntNumMolecule).isBIntNumOne(\n" +
            "  base, ePrefix)",
          ErrContext: "Error returned while testing 'base' for a value of one (1)!",
          ErrMessage: err.Error(),
        }

    }

  }

  if base.precision == 0 {
    powerCalcProfile.BaseIsInteger = true
  }

  if base.sign == -1 {
    powerCalcProfile.BaseIsNegative = true
  }

  powerCalcProfile.ExponentIsZero, err = new(bigIntNumMolecule).isBIntNumZero(
    exponent,
    ePrefix)

  if err != nil {

    return BigIntMathPowerProfile{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "powerCalcProfile.ExponentIsZero, err = new(bigIntNumMolecule).isBIntNumZero(\n" +
          "  exponent, ePrefix)",
        ErrContext: "Error returned while testing 'exponent' for zero value!",
        ErrMessage: err.Error(),
      }

  }

  if !powerCalcProfile.ExponentIsZero {

    powerCalcProfile.ExponentIsOne, err = new(bigIntNumMolecule).isBIntNumAbsOne(
      exponent,
      ePrefix)

    if err != nil {

      return BigIntMathPowerProfile{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "powerCalcProfile.ExponentIsOne, err = new(bigIntNumMolecule).isBIntNumOne(\n" +
            "  exponent, ePrefix)",
          ErrContext: "Error returned while testing 'exponent' for a value of one (1)!",
          ErrMessage: err.Error(),
        }

    }

  }

  if exponent.precision == 0 {
    powerCalcProfile.ExponentIsInteger = true
  }

  if exponent.sign == -1 {
    powerCalcProfile.ExponentIsNegative = true
  }

  // Calc Type # 1
  if !powerCalcProfile.BaseIsNegative &&
    powerCalcProfile.BaseIsInteger &&
    !powerCalcProfile.ExponentIsNegative &&
    powerCalcProfile.ExponentIsInteger {

    powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoPlusInt

  }

  // Calc Type # 2
  if !powerCalcProfile.BaseIsNegative &&
    powerCalcProfile.BaseIsInteger &&
    !powerCalcProfile.ExponentIsNegative &&
    !powerCalcProfile.ExponentIsInteger {

    powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoPlusFrac

  }

  // Calc Type # 3
  if !powerCalcProfile.BaseIsNegative &&
    powerCalcProfile.BaseIsInteger &&
    powerCalcProfile.ExponentIsNegative &&
    powerCalcProfile.ExponentIsInteger {

    powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoMinusInt

  }

  // Calc Type # 4
  if !powerCalcProfile.BaseIsNegative &&
    powerCalcProfile.BaseIsInteger &&
    powerCalcProfile.ExponentIsNegative &&
    !powerCalcProfile.ExponentIsInteger {

    powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoMinusFrac

  }
 
  // Calc Type # 5
  if !powerCalcProfile.BaseIsNegative &&
    !powerCalcProfile.BaseIsInteger &&
    !powerCalcProfile.ExponentIsNegative &&
    powerCalcProfile.ExponentIsInteger {

    powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusFracExpoPlusInt

  }

  return powerCalcProfile, err
}
