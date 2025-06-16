package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type numStrDtoMechanics struct {
  lock sync.Mutex
}

// addNumStrDto
//
//	Adds the numeric values of input parameters 'n1Dto' and 'n2Dto'. The
//	final result is stored in parameter 'n1Dto'.
func (numStrDtoMech *numStrDtoMechanics) addNumStrDto(
  numSeps NumericSeparatorDto,
  n1Dto *NumStrDto,
  validateN1Dto bool,
  n2Dto *NumStrDto,
  validateN2Dto bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  numStrDtoMech.lock.Lock()

  defer numStrDtoMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoMechanics.addNumStrDto()",
    "")

  if err != nil {
    return err
  }

  if n1Dto == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'n1Dto'",
    }
  }

  if n2Dto == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'n2Dto'",
    }
  }

  if validateN1Dto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "  n1Dto, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  if validateN2Dto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "  n2Dto, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
      ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
        "'numSeps' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  nResult, err := new(numStrDtoBoson).addNumStrs(
    numSeps, n1Dto, false, n2Dto, false, ePrefix.XCpy("n1Dto+n2Dto"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "nResult, err := new(numStrDtoBoson).addNumStrs(\n" +
        "  numSeps, n1Dto, false, n2Dto, false, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(numStrDtoMolecule).copy(n1Dto, &nResult, false, ePrefix.XCpy("nResult->n1Dto"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
        "  n1Dto, &nResult, false, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// addNumStrs
//
//		Adds the values represented by two NumStrDto objects and
//		returns the result as a new instance of NumStrDto.
//
//	   n1Dto + n2Dto = Returned NumStrDto
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and
//		convert them into numeric values.
//
//		The final NumStrDto result returned by this method will be
//		configured with the Numeric Separators passed by input
//	 	parameter 'numSeps'. If these Numeric Separators are
//		determined to be invalid, an error will be returned.
func (numStrDtoMech *numStrDtoMechanics) addNumStrs(
  numSeps NumericSeparatorDto,
  n1Dto *NumStrDto,
  validateN1Dto bool,
  n2Dto *NumStrDto,
  validateN2Dto bool,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  numStrDtoMech.lock.Lock()

  defer numStrDtoMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoMechanics.addNumStrs()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if n1Dto == nil {

    return NumStrDto{}, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'n1Dto'",
    }
  }

  if n2Dto == nil {

    return NumStrDto{}, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'n2Dto'",
    }
  }

  if validateN1Dto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

    if err != nil {

      return NumStrDto{}, &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "  n1Dto, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  if validateN2Dto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

    if err != nil {

      return NumStrDto{}, &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "  n2Dto, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{}, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
      ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
        "'numSeps' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  var totalNStrDto NumStrDto

  totalNStrDto, err = new(numStrDtoBoson).addNumStrs(
    numSeps, n1Dto, false, n2Dto, false, ePrefix.XCpy("n1Dto + n2Dto"))

  if err != nil {
    return NumStrDto{}, &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "totalNStrDto, err = new(numStrDtoBoson).addNumStrs(\n" +
        "  numSeps, n1Dto, false, n2Dto, false, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return totalNStrDto, nil
}
