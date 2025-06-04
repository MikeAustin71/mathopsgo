package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryMolecule struct {
  lock sync.Mutex
}

// roundToPrecision
//
//	Rounds the value of the current IntAry instance to a precision
//	specified by the 'roundToPrecision' parameter.
func (iaMolecule *intAryMolecule) roundToPrecision(
  intAry *IntAry,
  validateIntAry bool,
  roundToPrecision int,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaMolecule.lock.Lock()

  defer iaMolecule.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMolecule.roundToPrecision()",
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

  if roundToPrecision < 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Input parameter 'roundToPrecision' is less than ZERO!\n"+
        "roundToPrecision= %v\n", roundToPrecision),
    }
  }

  if intAry.precision == 0 {
    return nil
  }

  err = new(intAryUtility).selectIntAryValidation(
    intAry,
    "intAry",
    validateIntAry,
    ePrefix)

  if err != nil {
    return err
  }

  numSeps, err := new(intAryPhoton).getNumericSeparatorsDto(intAry, true, ePrefix.XCpy("Numeric Separators intAry -> numSeps"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err := new(intAryPhoton).\n" +
        "  getNumericSeparatorsDto(intAry, true,\n" +
        "  ePrefix.XCpy(Numeric Separators intAry -> numSeps))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nsProfile := NumSepsProfileSelection{
    SourceObjectName:         "intAry",
    OutputNumSepsName:        "numSeps",
    UseDefaultNumSeps:        false,
    SetDefaultNumSepsIfEmpty: true,
    ValidateNumSeps:          false,
    OverrideNumSeps:          numSeps,
  }

  if intAry.isZeroValue {

    err = new(intAryQuark).setIntAryToZero(
      intAry,
      nil,
      nsProfile,
      uint(roundToPrecision),
      ePrefix.XCpy("Setting 'ia' to Zero"))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryQuark).setIntAryToZero(\n" +
          "  ia, nil, nsProfile-numSeps, 0,\n" +
          "  ePrefix.XCpy(Setting 'ia' to Zero))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    return nil
  }

  if roundToPrecision == intAry.precision {
    return nil
  }

  if roundToPrecision > intAry.precision {

    deltaPrecision := roundToPrecision - intAry.precision

    for i := 0; i < deltaPrecision; i++ {

      intAry.intAry = append(intAry.intAry, 0)
    }

    intAry.intAryLen = len(intAry.intAry)

    intAry.precision = roundToPrecision

    err = new(intAryNanobot).setInternalFlags(
      intAry, ePrefix.XCpy("Setting 'ia' Flags"))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
          "  ia, ePrefix.XCpy(Setting 'ia' Flags))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    return nil
  }

  // roundToPrecision must be < ia.precision

  intLen := intAry.intAryLen - intAry.precision

  newIntAryLen := intLen + roundToPrecision

  fracIdx := intLen

  fracRoundIdx := fracIdx + roundToPrecision

  t := make([]uint8, intAry.intAryLen+1)

  n1 := uint8(0)

  n2 := uint8(0)

  carry := uint8(0)

  for i := fracRoundIdx; i >= 0; i-- {

    n1 = intAry.intAry[i]

    if i == fracRoundIdx {

      n2 = n1 + 5

    } else {

      n2 = n1 + carry

    }

    if n2 > 9 {

      carry = 1

      n2 = n2 - 10

    } else {

      carry = 0
    }

    t[i+1] = n2
  }

  intAry.intAry = []uint8{}

  if carry > 0 {

    t[0] = carry

    intAry.intAry = t[0 : newIntAryLen+1]

  } else {

    intAry.intAry = t[1 : newIntAryLen+1]
  }

  intAry.precision = roundToPrecision

  intAry.intAryLen = len(intAry.intAry)

  err = new(intAryNanobot).setInternalFlags(
    intAry, ePrefix.XCpy("Setting 'ia' Flags"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(\n" +
        "  ia, ePrefix.XCpy(Setting 'ia' Flags))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
