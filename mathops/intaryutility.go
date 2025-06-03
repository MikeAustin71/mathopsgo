package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
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

  var tagLine1, tagLine2, retFuncName string

  if !validateIntAry {

    tagLine1 = fmt.Sprintf("Setting '%s' Flags", intAryName)

    tagLine2 = fmt.Sprintf("Error Setting Flags on '%s'", intAryName)

    retFuncName = fmt.Sprintf("err = new(intAryNanobot).setInternalFlags(\n"+
      "  %s, ePrefix.XCpy(Setting '%s' Flags))", intAryName, intAryName)

    err = new(intAryNanobot).setInternalFlags(
      intAry, ePrefix.XCpy(tagLine1))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: retFuncName,
        ErrContext: tagLine2,
        ErrMessage: err.Error(),
      }
    }

    return nil
  } // End Of if !validateIntAry

  // We need to validate the IntAry Object.
  tagLine1 = fmt.Sprintf("Validating '%s'", intAryName)

  tagLine2 = fmt.Sprintf("Input parameter '%s' is INVALID!\n"+
    "'%s' FAILED Validation Tests.", intAryName, intAryName)

  retFuncName = fmt.Sprintf("err = new(intAryElectron).isValidIntAry(\n"+
    "%s, ePrefix.XCpy(Validating '%s').String())", intAryName, intAryName)

  err = new(intAryElectron).isValidIntAry(
    intAry,
    ePrefix.XCpy(tagLine1).String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: retFuncName,
      ErrContext: tagLine2,
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
  nsProfile NumSepProfileSelection,
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
