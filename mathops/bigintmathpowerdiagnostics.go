package mathops

import (
  "fmt"
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

    var isBaseAbsOne bool

    isBaseAbsOne, err = new(bigIntNumMolecule).isBIntNumAbsOne(
      base,
      ePrefix)

    if err != nil {

      return BigIntMathPowerProfile{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "isBaseAbsOne, err = new(bigIntNumMolecule).isBIntNumAbsOne(\n" +
            "  base, ePrefix)",
          ErrContext: "Error returned while testing 'base' for a value of one (1)!",
          ErrMessage: err.Error(),
        }

    }

    if isBaseAbsOne && base.sign == 1 {
      powerCalcProfile.BaseIsPlusOne = true

    } else if isBaseAbsOne && base.sign == -1 {
      powerCalcProfile.BaseIsMinusOne = true

    } else {

      err = fmt.Errorf("%v\n"+
        "base is INVALID. base.sign is NOT set to 1 or -1!\n"+
        "base.sign = %v\n",
        ePrefix,
        base.sign)

      return BigIntMathPowerProfile{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "Error: base.sign is NOT set to 1 or -1!",
          ErrMessage: err.Error(),
        }

    }

  }

  if powerCalcProfile.BaseIsPlusOne || powerCalcProfile.BaseIsMinusOne {
    powerCalcProfile.BaseIsAbsOne = true
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

    var isExponentAbsOne bool

    isExponentAbsOne, err = new(bigIntNumMolecule).isBIntNumAbsOne(
      exponent,
      ePrefix)

    if err != nil {

      return BigIntMathPowerProfile{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "isExponentAbsOne, err = new(bigIntNumMolecule).isBIntNumAbsOne(\n" +
            "  exponent, ePrefix)",
          ErrContext: "Error returned while testing 'exponent' for a value of one (1)!",
          ErrMessage: err.Error(),
        }

    }

    if isExponentAbsOne && exponent.sign == 1 {
      powerCalcProfile.ExponentIsPlusOne = true

    } else if isExponentAbsOne && exponent.sign == -1 {
      powerCalcProfile.ExponentIsMinusOne = true

    } else {

      err = fmt.Errorf("%v\n"+
        "exponent is INVALID. exponent.sign is NOT set to 1 or -1!\n"+
        "exponent.sign = %v\n",
        ePrefix,
        exponent.sign)

      return BigIntMathPowerProfile{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "Error: exponent.sign is NOT set to 1 or -1!",
          ErrMessage: err.Error(),
        }

    }

  }

  if powerCalcProfile.ExponentIsPlusOne || powerCalcProfile.ExponentIsMinusOne {
    powerCalcProfile.ExponentIsAbsOne = true
  }

  if exponent.precision == 0 {
    powerCalcProfile.ExponentIsInteger = true
  }

  if exponent.sign == -1 {
    powerCalcProfile.ExponentIsNegative = true
  }

  // Set Exponent Calculation Type Code

  pwrCalcClassify := new(bigIntMathPowerCalcClassify)

  err = pwrCalcClassify.classifyExponentCalcType(&powerCalcProfile, ePrefix)

  if err != nil {

    return BigIntMathPowerProfile{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = pwrCalcClassify.classifyExponentCalcType(&powerCalcProfile, ePrefix)",
        ErrContext: "Error returned while classifying the type of exponentiation calculation!",
        ErrMessage: err.Error(),
      }
  }

  return powerCalcProfile, nil
}
