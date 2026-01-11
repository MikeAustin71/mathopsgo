package mathops

import (
  "fmt"
  "sync"

  ePref "github.com/MikeAustin71/errpref"
)

type intAryUtility struct {
  theLock sync.Mutex
}

// selectIntAryValidation
//
// Designed to perform conditional validation on IntAry objects.
//
// Specific characteristics of the validation are engineered
// through user input.
func (iaUtility *intAryUtility) selectIntAryValidation(
  intAry *IntAry,
  intAryName string,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaUtility.theLock.Lock()

  defer iaUtility.theLock.Unlock()
  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryUtility.selectIntAryValidation()",
    "")

  if err != nil {
    return err
  }

  if intAry == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'intAry'",
    }
  }

  if intAryName == "" {
    intAryName = "intAry"
  }

  if !validateIntAry {

    err = new(intAryElectron).setSignificantDigitIdxs(
      intAry,
      ePrefix)

    if err != nil {
      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: fmt.Sprintf("err = new(intAryElectron).\n"+
          "  setSignificantDigitIdxs(\n"+
          "  intAryName=%s, ePrefix)", intAryName),
        ErrContext: fmt.Sprintf(
          "Error Setting Internal Flags on '%s'", intAryName),
        ErrMessage: err.Error(),
      }
    }

    return nil
  } // End Of if !validateIntAry

  // We need to validate the IntAry Object.

  // This sets internal flags
  err = new(intAryElectron).isValidIntAry(
    intAry,
    "Validating "+intAryName)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryElectron).isValidIntAry(\n"+
        "Validating '%s')", intAryName),
      ErrContext: fmt.Sprintf("Input parameter '%s' is INVALID!\n"+
        "'%s' FAILED Validation Tests.", intAryName, intAryName),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// selectNumericSeparators
//
// Designed to perform conditional generation of Numeric Separators
// for use with IntAry objects.
//
// Specific characteristics of the returned Numeric Separator Dto
// objects are engineered through user input.
func (iaUtility *intAryUtility) selectNumericSeparators(
  intAry *IntAry,
  nsProfile NumSepsProfileSelection,
  errPrefDto *ePref.ErrPrefixDto) (NumericSeparatorDto, error) {

  iaUtility.theLock.Lock()

  defer iaUtility.theLock.Unlock()
  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryUtility.selectNumericSeparators()",
    "")

  if err != nil {
    return NumericSeparatorDto{}, err
  }

  if nsProfile.SourceObjectName == "" {
    nsProfile.SourceObjectName = "intAry"
  }

  if nsProfile.OutputNumSepsName == "" {
    nsProfile.OutputNumSepsName = "numSeps"
  }

  var numSeps NumericSeparatorDto

  err = nsProfile.OverrideNumSeps.IsValid(ePrefix.String())

  if err == nil {
    // This a valid NumericSeparatorDto instance. Use This
    err = numSeps.CopyIn(&nsProfile.OverrideNumSeps, true)

    if err != nil {

      return NumericSeparatorDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: " err = numSeps.CopyIn(&nsProfile.OverrideNumSeps, true)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    return numSeps, nil
  }

  if nsProfile.UseDefaultNumSeps {

    numSeps = new(NumericSeparatorDto).NewUSADefaults()

    return numSeps, nil
  }

  // If not using defaults, must use source object
  if intAry == nil {

    return NumericSeparatorDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  // Take NumSeps from 'intAry'

  tagLine := fmt.Sprintf("%s Numeric Separators -> %s", nsProfile.SourceObjectName, nsProfile.OutputNumSepsName)

  retFuncName := fmt.Sprintf("%s, err = new(intAryPhoton).getNumericSeparatorsDto(\n%s, ePrefix.XCpy(%s))",
    nsProfile.OutputNumSepsName, nsProfile.SourceObjectName, tagLine)

  numSeps, err = new(intAryPhoton).getNumericSeparatorsDto(intAry, nsProfile.SetDefaultNumSepsIfEmpty, ePrefix.XCpy(tagLine))

  if err != nil {

    return NumericSeparatorDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: retFuncName,
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if nsProfile.ValidateNumSeps {

    tagLine = fmt.Sprintf("%s\nValidating Numeric Separators %s",
      ePrefix.String(), nsProfile.OutputNumSepsName)

    retFuncName = fmt.Sprintf(
      "err = %s.IsValid(ePrefix)", nsProfile.OutputNumSepsName)

    err = numSeps.IsValid(tagLine)

    if err != nil {

      return NumericSeparatorDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: retFuncName,
          ErrContext: tagLine,
          ErrMessage: err.Error(),
        }
    }
  }

  return numSeps, nil
}
