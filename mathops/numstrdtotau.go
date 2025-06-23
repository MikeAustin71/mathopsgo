package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math"
  "sync"
)

type numStrDtoTau struct {
  lock sync.Mutex
}

// lowLevelAddition
//
//	 Performs addition on two NumStrDto values and returns the
//	 total value as a new instance of NumStrDto.
//
//	 Assumptions
//	 ===========
//
//	 Since this is a low-level routine, the following assumptions
//	 apply to input parameters:
//
//	 1. 'n1NumDto' and 'n2NumDto' have the same sign value
//
//	 2. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    have equal length. Equal length can be achieved by calling
//	    new(numStrDtoPhoton).formatForMathOps() before calling
//	    this method.
//
//	 3. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    contain valid 'runes' with value >= '0' && value <= '9'.
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
//		parameter 'numSeps'. If these Numeric Separators are determined
//	 to be invalid, an error will be returned.
func (nStrDtoTau *numStrDtoTau) lowLevelAddition(
  numSeps NumericSeparatorDto,
  n1NumDto *NumStrDto,
  validateN1NumDto bool,
  n2NumDto *NumStrDto,
  validateN2NumDto bool,
  precision uint,
  newSignVal int,
  validateFinalResult bool,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  nStrDtoTau.lock.Lock()

  defer nStrDtoTau.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoTau.lowLevelAddition()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if n1NumDto == nil {

    return NumStrDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n1NumDto'",
      }
  }

  if n2NumDto == nil {

    return NumStrDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n2NumDto'",
      }
  }

  if validateN1NumDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n1NumDto, ePrefix.XCpy("Validating 'n1NumDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  n1NumDto, ePrefix)",
          ErrContext: "Error: Input parameter 'n1NumDto' is invalid!\n" +
            "'n1NumDto' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateN2NumDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n2NumDto, ePrefix.XCpy("Validating 'n2NumDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  n2NumDto, ePrefix)",
          ErrContext: "Error: Input parameter 'n2NumDto' is invalid!\n" +
            "'n2NumDto' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  lenN1AllRunes := len(n1NumDto.absAllNumRunes)

  if lenN1AllRunes == 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "n1NumDto.absAllNumRunes == 0",
        ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
          "'n1NumDto' internal rune array is empty with zero length. ",
      }

  }

  lenN2AllRunes := len(n2NumDto.absAllNumRunes)

  if lenN1AllRunes != lenN2AllRunes {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "n1NumDto.absAllNumRunes != n2NumDto.absAllNumRunes",
        ErrMessage: "Error: Input parameters 'n1NumDto' and 'n2NumDto'\n" +
          "have runes arrays of unequal length.\n" +
          "'n1NumDto' and 'n2NumDto' internal rune arrays must have equal length. ",
      }
  }

  n3IntAry := make([]int, lenN1AllRunes+1)

  carry := 0

  n1 := 0

  n2 := 0

  n3 := 0

  for j := lenN1AllRunes - 1; j >= 0; j-- {

    n1 = int(n1NumDto.absAllNumRunes[j]) - 48

    if n1 < 0 || n1 > 9 {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "",
          ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
            "The 'n1NumDto' internal rune array contains invalid characters!\n" +
            fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
              j, n1NumDto.absAllNumRunes[j]),
        }

    }

    n2 = int(n2NumDto.absAllNumRunes[j]) - 48

    if n2 < 0 || n2 > 9 {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "",
          ErrMessage: "Error: Input parameter 'n2NumDto' is invalid!\n" +
            "The 'n2NumDto' internal rune array contains invalid characters!\n" +
            fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
              j, n2NumDto.absAllNumRunes[j]),
        }
    }

    n3 = n1 + n2 + carry

    carry = 0

    if n3 > 9 {

      n3 = n3 - 10

      carry = 1
    }

    n3IntAry[j+1] = n3
  }

  if carry > 0 {

    n3IntAry[0] = carry
  }

  // return nDto.FindIntArraySignificantDigitLimits(n3IntAry, precision, newSignVal)

  nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(
    numSeps, n3IntAry, precision, newSignVal, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(",
        ErrContext: "  numSeps, n3IntAry, precision, newSignVal, ePrefix)",
        ErrMessage: err.Error(),
      }
  }

  isZeroValue, err := new(numStrDtoElectron).isNumStrZeroValue(
    &nOutDto, false, ePrefix.XCpy("nDtoOut=0 ?"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: " isZeroValue, err :=  new(numStrDtoElectron).isNumStrZeroValue(\n" +
          "  &nDtoOut, true, ePrefix.XCpy(\"nDtoOut=0 ?\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if isZeroValue {
    newSignVal = 1
  }

  err = new(numStrDtoAtom).setSignValue(
    &nOutDto, false, newSignVal, ePrefix.XCpy("n1DtoSetup"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: " err = new(numStrDtoAtom).setSignValue(&nDtoOut, true, 1,\n" +
          "  ePrefix.XCpy(\"nDtoOut\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if validateFinalResult {

    err = new(numStrDtoElectron).isValidNumStrDto(
      &nOutDto, ePrefix.XCpy("Validating Final Result 'nOutDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  &nOutDto, ePrefix)",
          ErrContext: "Error: Final Calculated Result 'nOutDto' is invalid!\n" +
            "'nOutDto' FAILED final validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  return nOutDto, nil

}

// lowLevelAddition
//
//	 Performs addition on two NumStrDto values and returns the
//	 total value as a new instance of NumStrDto.
//
//	 Assumptions
//	 ===========
//
//	 Since this is a low-level routine, the following assumptions
//	 apply to input parameters:
//
//	 1. 'n1NumDto' and 'n2NumDto' have the same sign value
//
//	 2. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    have equal length. Equal length can be achieved by calling
//	    new(numStrDtoPhoton).formatForMathOps() before calling
//	    this method.
//
//	 3. With respect to numeric values 'n1NumDto' must be greater
//	    than or equal to 'n2NumDto'. Again, this relationship can
//	    be achieved by calling
//	    new(numStrDtoPhoton).formatForMathOps() before calling this
//	    method.
//
//	 4. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    contain valid 'runes' with value >= '0' && value <= '9'.
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
//		parameter 'numSeps'. If these Numeric Separators are determined
//	 to be invalid, an error will be returned.
func (nStrDtoTau *numStrDtoTau) lowLevelSubtraction(
  numSeps NumericSeparatorDto,
  n1NumDto *NumStrDto,
  validateN1NumDto bool,
  n2NumDto *NumStrDto,
  validateN2NumDto bool,
  precision uint,
  newSignVal int,
  validateFinalResult bool,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  nStrDtoTau.lock.Lock()

  defer nStrDtoTau.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoTau.lowLevelSubtraction()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if n1NumDto == nil {

    return NumStrDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n1NumDto'",
      }
  }

  if n2NumDto == nil {

    return NumStrDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n2NumDto'",
      }
  }

  if validateN1NumDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n1NumDto, ePrefix.XCpy("Validating 'n1NumDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  n1NumDto, ePrefix)",
          ErrContext: "Error: Input parameter 'n1NumDto' is invalid!\n" +
            "'n1NumDto' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateN2NumDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n2NumDto, ePrefix.XCpy("Validating 'n2NumDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  n2NumDto, ePrefix)",
          ErrContext: "Error: Input parameter 'n2NumDto' is invalid!\n" +
            "'n2NumDto' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  lenN1AllRunes := len(n1NumDto.absAllNumRunes)

  if lenN1AllRunes == 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "n1NumDto.absAllNumRunes == 0",
        ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
          "'n1NumDto' internal rune array is empty with zero length. ",
      }

  }

  lenN2AllRunes := len(n2NumDto.absAllNumRunes)

  if lenN1AllRunes != lenN2AllRunes {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "n1NumDto.absAllNumRunes != n2NumDto.absAllNumRunes",
        ErrMessage: "Error: Input parameters 'n1NumDto' and 'n2NumDto'\n" +
          "have runes arrays of unequal length.\n" +
          "'n1NumDto' and 'n2NumDto' internal rune arrays must have equal length. ",
      }
  }

  n1IntAry := make([]int, lenN1AllRunes)

  n2IntAry := make([]int, lenN1AllRunes)

  n3IntAry := make([]int, lenN1AllRunes)

  for i := 0; i < lenN1AllRunes; i++ {

    n1IntAry[i] = int(n1NumDto.absAllNumRunes[i]) - 48

    n2IntAry[i] = int(n2NumDto.absAllNumRunes[i]) - 48

  }

  carry := 0
  n1 := 0
  n2 := 0
  n3 := 0

  // Main Subtraction Routine
  for j := lenN1AllRunes - 1; j >= 0; j-- {

    n1 = n1IntAry[j]

    if n1 < 0 || n1 > 9 {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "",
          ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
            "The 'n1NumDto' internal rune array contains invalid characters!\n" +
            fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
              j, n1NumDto.absAllNumRunes[j]),
        }

    }

    n2 = n2IntAry[j]

    if n2 < 0 || n2 > 9 {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "",
          ErrMessage: "Error: Input parameter 'n2NumDto' is invalid!\n" +
            "The 'n2NumDto' internal rune array contains invalid characters!\n" +
            fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
              j, n2NumDto.absAllNumRunes[j]),
        }
    }

    n3 = 0

    if n1-carry-n2 < 0 {

      n1 += 10

      n3 = n1 - n2 - carry

      carry = 1

    } else {

      n3 = n1 - n2 - carry

      carry = 0
    }

    n3IntAry[j] = n3

  }

  //nOutDto, err := nDto.FindIntArraySignificantDigitLimits(n3IntAry, precision, newSignVal)

  nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(
    numSeps, n3IntAry, precision, newSignVal, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(",
        ErrContext: "  numSeps, n3IntAry, precision, newSignVal, ePrefix)",
        ErrMessage: err.Error(),
      }
  }

  isZeroValue, err := new(numStrDtoElectron).isNumStrZeroValue(
    &nOutDto, false, ePrefix.XCpy("nDtoOut=0 ?"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: " isZeroValue, err :=  new(numStrDtoElectron).isNumStrZeroValue(\n" +
          "  &nDtoOut, true, ePrefix.XCpy(\"nDtoOut=0 ?\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if isZeroValue {
    newSignVal = 1
  }

  //nDtoOut.SetSignValue(newSignVal)

  err = new(numStrDtoAtom).setSignValue(
    &nOutDto, false, newSignVal, ePrefix.XCpy("n1DtoSetup"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: " err = new(numStrDtoAtom).setSignValue(&nDtoOut, true, 1,\n" +
          "  ePrefix.XCpy(\"nDtoOut\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if validateFinalResult {

    err = new(numStrDtoElectron).isValidNumStrDto(
      &nOutDto, ePrefix.XCpy("Validating Final Result 'nOutDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  &nOutDto, ePrefix)",
          ErrContext: "Error: Final Calculated Result 'nOutDto' is invalid!\n" +
            "'nOutDto' FAILED final validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  return nOutDto, nil
}

// multiplyNumStrs
//
//	Multiplies two NumStrDto instances and returns the result as a
//	separate NumStrDto instance.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators provided by the
//	current instance of 'nDto'. If these Numeric Separators are
//	determined to be invalid, an error will be returned.
func (nStrDtoTau *numStrDtoTau) multiplyNumStrs(
  numSeps NumericSeparatorDto,
  n1NumDto *NumStrDto,
  validateN1NumDto bool,
  n2NumDto *NumStrDto,
  validateN2NumDto bool,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  nStrDtoTau.lock.Lock()

  defer nStrDtoTau.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoTau.multiplyNumStrs()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if n1NumDto == nil {

    return NumStrDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n1NumDto'",
      }
  }

  if n2NumDto == nil {

    return NumStrDto{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n2NumDto'",
      }
  }

  if validateN1NumDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n1NumDto, ePrefix.XCpy("Validating 'n1NumDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  n1NumDto, ePrefix)",
          ErrContext: "Error: Input parameter 'n1NumDto' is invalid!\n" +
            "'n1NumDto' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateN2NumDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      n2NumDto, ePrefix.XCpy("Validating 'n2NumDto'"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  n2NumDto, ePrefix)",
          ErrContext: "Error: Input parameter 'n2NumDto' is invalid!\n" +
            "'n2NumDto' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  lenN1AbsAllRunes := len(n1NumDto.absAllNumRunes)

  lenN2AbsAllRunes := len(n2NumDto.absAllNumRunes)

  var n1Setup NumStrDto
  var n2Setup NumStrDto

  if lenN2AbsAllRunes > lenN1AbsAllRunes {

    //n1Setup = n2Dto.CopyOut()

    err = new(numStrDtoMolecule).copy(&n1Setup, n2NumDto, false, ePrefix.XCpy("n2NumDto->n1Setup"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
            "  &n1Setup, n2NumDto, false, ePrefix.XCpy(\"n2NumDto->n1Setup\"))",
          ErrContext: "Error Copying n2NumDto->n1Setup",
          ErrMessage: err.Error(),
        }
    }

    //n2Setup = n1Dto.CopyOut()

    err = new(numStrDtoMolecule).copy(&n2Setup, n1NumDto, false, ePrefix.XCpy("n1NumDto->n2Setup"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
            "  &n2Setup, n1NumDto, false, ePrefix.XCpy(\"n1NumDto->n2Setup\"))",
          ErrContext: "Error Copying n1NumDto->n2Setup",
          ErrMessage: err.Error(),
        }
    }

  } else {
    // Must be lenN1AbsAllRunes >= lenN2AbsAllRunes

    // n1Setup = n1Dto.CopyOut()

    err = new(numStrDtoMolecule).copy(&n1Setup, n1NumDto, false, ePrefix.XCpy("n1NumDto->n1Setup"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
            "  &n1Setup, n1NumDto, false, ePrefix.XCpy(\"n1NumDto->n1Setup\"))",
          ErrContext: "Error Copying n1NumDto->n1Setup",
          ErrMessage: err.Error(),
        }
    }

    //n2Setup = n2Dto.CopyOut()

    err = new(numStrDtoMolecule).copy(&n2Setup, n2NumDto, false, ePrefix.XCpy("n2NumDto->n2Setup"))

    if err != nil {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
            "  &n2Setup, n2NumDto, false, ePrefix.XCpy(\"n2NumDto->n2Setup\"))",
          ErrContext: "Error Copying n2NumDto->n2Setup",
          ErrMessage: err.Error(),
        }
    }

  }

  //else {
  //  // Must be lenN1AbsAllRunes == lenN2AbsAllRunes
  //
  //  n1Setup = n1Dto.CopyOut()
  //
  //  n2Setup = n2Dto.CopyOut()
  //
  //}

  newPrecision := n1Setup.precision + n2Setup.precision

  newSignVal := 1

  if n1Setup.signVal == n2Setup.signVal {

    newSignVal = 1

  } else {

    // Must be n1Setup.signVal != n2Setup.signVal
    newSignVal = -1

  }

  lenN1AbsAllRunes = len(n1Setup.absAllNumRunes)

  lenN2AbsAllRunes = len(n2Setup.absAllNumRunes)

  lenLevels := lenN2AbsAllRunes

  lenNumPlaces := (lenN1AbsAllRunes + lenN2AbsAllRunes) + 1

  intMAry := make([][]int, lenLevels)

  for i := 0; i < lenLevels; i++ {
    intMAry[i] = make([]int, lenNumPlaces)
  }

  intFinalAry := make([]int, lenNumPlaces+1)

  carry := 0

  levels := 0

  place := 0

  n1 := 0

  n2 := 0

  n3 := 0

  n4 := 0

  for i := lenN2AbsAllRunes - 1; i >= 0; i-- {

    place = (lenNumPlaces - 1) - levels

    for j := lenN1AbsAllRunes - 1; j >= 0; j-- {

      n1 = int(n1Setup.absAllNumRunes[j]) - 48

      n2 = int(n2Setup.absAllNumRunes[i]) - 48

      n3 = (n1 * n2) + carry

      //n4 = int(math.Mod(float64(n3), float64(10.00)))
      n4 = int(math.Mod(float64(n3), 10.00))

      intMAry[levels][place] = n4

      //carry = int(n3 / 10)
      carry = n3 / 10

      place--
    }

    intMAry[levels][place] = carry
    carry = 0
    levels++
  }

  carry = 0
  n1 = 0
  n2 = 0
  n3 = 0
  n4 = 0

  for i := 0; i < lenLevels; i++ {

    for j := lenNumPlaces - 1; j >= 0; j-- {

      n1 = intFinalAry[j+1]

      n2 = intMAry[i][j]

      n3 = n1 + n2 + carry

      n4 = 0

      if n3 > 9 {

        //n4 = int(math.Mod(float64(n3), float64(10.0)))
        n4 = int(math.Mod(float64(n3), 10.0))

        carry = n3 / 10

      } else {

        n4 = n3

        carry = 0
      }

      intFinalAry[j+1] = n4
    }

    if carry > 0 {

      intFinalAry[0] = carry
    }

  }

  //numStrOut, err := nDto.FindIntArraySignificantDigitLimits(intFinalAry, newPrecision, newSignVal)
  //
  //if err != nil {
  //  return NumStrDto{},
  //    fmt.Errorf(ePrefix+
  //      "- Error returned from nDto.FindIntArraySignificantDigitLimits(intFinalAry,newPrecision, "+
  //      "newSignVal). Error= %v", err)
  //}

  numStrOut, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(
    numSeps, intFinalAry, newPrecision, newSignVal, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numStrOut, err := new(numStrDtoAtom).\n" +
          "  findIntArraySignificantDigitLimits(\n" +
          "  numSeps, intFinalAry, newPrecision, newSignVal, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    &numStrOut, ePrefix.XCpy("Validating Final Result 'numStrOut'"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "  &numStrOut, ePrefix)",
        ErrContext: "Error: Final Result 'numStrOut' is invalid!\n" +
          "'numStrOut' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return numStrOut, nil
}

// shiftPrecisionLeft
//
//	Shifts the relative position of a decimal point within a number
//	string. The position of the decimal point is shifted
//	'shiftPrecision' positions to the left of the current decimal
//	point position.
//
//	This is equivalent to:
//	         result = signedNumStr / 10^precision
//	                          or
//	 signedNumStr divided by 10 raised to the power of precision.
//
//	Examples
//	========
//
//	                  Shift-Left
//	signedNumStr      precision         Result
//
//	"123456.789"          3           "123.456789"
//	"123456.789"          2           "1234.56789"
//	"123456.789"          6           "0.123456789"
//	"123456789"           6           "123.456789"
//	"123"                 5           "0.00123"
//	"0"                   3           "0.000"
//	"0.000"               2           "0.00000"
//	"123456.789"          0           "123456.789"
//	      zero 'shiftPrecision' has no effect on
//	          the original number string
//
//	"-123456.789"         3           "-123.456789"
//	"-123456789"          6           "-123.456789"
//
//	Numeric Separators
//	==================
//
//	The Numeric Separators originally configured for the current
//	instance of NumStrDto will be copied to the returned instance
//	of NumStrDto.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	numSeps                  NumericSeparatorDto
//	 This parameter contains the Numeric Separators which will be
//	 used to populate the returned instance of 'NumStrDto'.
//
//	signedNumStr             string
//	  A valid number string. The leading digit may optionally be a
//	  '+' or '-' indicating numeric sign value. If '+' or '-'
//	  characters are not present in the first character position,
//	  the number is assumed to represent a positive	numeric value
//	  ('+').
//
//	shiftPrecision           uint
//	  The number of digits by which the current decimal point
//	  position in the number string, 'signedNumStr' will be shifted
//	  to the left.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//	  This object encapsulates an error prefix string
//	  which is included in all returned error
//	  messages. Usually, it contains the name of the
//	  calling method or methods listed as a function
//	  chain.
//
//	  If no error prefix information is needed, set
//	  this parameter to 'nil'.
//
//	  Type ErrPrefixDto is included in the 'errpref'
//	  software package:
//	    "github.com/MikeAustin71/errpref".
//
//	Return Values
//	=============
//
//	NumStrDto
//	  This method returns the result of the Shift Left precision
//	  operation in the form of a new 'NumStrDto' instance.
//
//	error
//	  If a processing error is encountered, this returned error
//	  object will be configured with an appropriate error message.
func (nStrDtoTau *numStrDtoTau) shiftPrecisionLeft(
  numSeps NumericSeparatorDto,
  signedNumStr string,
  shiftLeftPrecision uint,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  nStrDtoTau.lock.Lock()

  defer nStrDtoTau.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoTau.shiftPrecisionLeft()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if len(signedNumStr) == 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(signedNumStr) == 0",
        ErrMessage: "Error: Input parameter 'signedNumStr' is INVALID!\n" +
          "'signedNumStr' is a zero length string.",
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  if new(MathProcessUtility).DoesUintExceedMax32BitInt(shiftLeftPrecision) {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'shiftLeftPrecision' is INVALID!\n" +
          "'shiftLeftPrecision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
          fmt.Sprintf("shiftLeftPrecision= '%v'", shiftLeftPrecision),
      }
  }

  n1, err := new(numStrDtoQuark).parseNumStr(numSeps, signedNumStr, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1, err := new(numStrDtoQuark).parseNumStr(\n" +
          "  numSeps, signedNumStr, ePrefix)",
        ErrContext: fmt.Sprintf("signedNumStr= '%s'\n"+
          "numSeps= '%s'", signedNumStr, numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  u64TotalSpecPrecision := uint64(shiftLeftPrecision) + uint64(n1.precision)

  if new(MathProcessUtility).DoesUint64ExceedMax32BitInt(u64TotalSpecPrecision) {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Computed parameter 'u64TotalSpecPrecision' is INVALID!\n" +
          "'u64TotalSpecPrecision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
          "'u64TotalSpecPrecision' is the sum of 'shiftLeftPrecision' plus 'n1.precision'\n" +
          fmt.Sprintf("u64TotalSpecPrecision= '%v'", u64TotalSpecPrecision),
      }
  }

  n2 := new(numStrDtoMolecule).newZeroNumStrDto(numSeps, 0)

  n2.signVal = n1.signVal

  n2.precision = shiftLeftPrecision + n1.precision

  iTotalSpecPrecision := int(n2.precision)

  lenAbsAllNumRunes := len(n1.absAllNumRunes)

  lenAbsIntRunes, err := new(numStrDtoGluon).getAbsIntRunesLength(
    &n1, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "lenAbsIntRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsIntRunesLength(&n1, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  lenAbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunesLength(
    &n1, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "lenAbsFracRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsFracRunesLength(&n1, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  n1NumStrIsZeroValue, err := new(numStrDtoElectron).isNumStrZeroValue(
    &n1, true, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1NumStrIsZeroValue, err := new(numStrDtoElectron).\n" +
          "  isNumStrZeroValue(&n1, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if n1NumStrIsZeroValue {

    n3 := new(numStrDtoMolecule).newZeroNumStrDto(numSeps, n2.precision)

    return n3, nil
  }

  if iTotalSpecPrecision == lenAbsAllNumRunes {

    n2.absAllNumRunes = append(n2.absAllNumRunes, '0')

  } else if iTotalSpecPrecision > lenAbsAllNumRunes {

    deltaPrecision := iTotalSpecPrecision - lenAbsAllNumRunes + 1

    for i := 0; i < deltaPrecision; i++ {

      n2.absAllNumRunes = append(n2.absAllNumRunes, '0')
    }

  }

  for j := 0; j < lenAbsAllNumRunes; j++ {

    n2.absAllNumRunes = append(n2.absAllNumRunes, n1.absAllNumRunes[j])

  }

  lenAbsAllNumRunes = len(n2.absAllNumRunes)

  lenAbsFracRunes = iTotalSpecPrecision

  lenAbsIntRunes = lenAbsAllNumRunes - lenAbsFracRunes

  if lenAbsIntRunes <= 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "lenAbsIntRunes <= 0",
        ErrMessage: "Error: Calculated number of integer digits is less than or equal to ZERO.\n" +
          fmt.Sprintf("lenAbsIntRunes= '%v' ", lenAbsIntRunes),
      }
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    &n2, ePrefix.XCpy("Validating 'n2'"))

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "  &n2, ePrefix)",
        ErrContext: "Error: The Final Result NumStrDto instance ('n2') is INVALID!\n" +
          "'n2' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  actualLenAbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunesLength(
    &n2, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "actualLenAbsFracRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsFracRunesLength(&n2, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if actualLenAbsFracRunes != iTotalSpecPrecision {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "actualLenAbsFracRunes != iTotalSpecPrecision",
        ErrMessage: "Error: Calculated number of fractional digits is INVALID!\n" +
          "fractional digits not equal to requested fractional digits.\n" +
          fmt.Sprintf("Calculated Fractional Digits= '%v'\n"+
            "Requested Fractional Digits= '%v'",
            actualLenAbsFracRunes, iTotalSpecPrecision),
      }
  }

  if uint(lenAbsFracRunes) != n2.precision {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "uint(lenAbsFracRunes) != n2.precision",
        ErrMessage: "Error: Calculated number of fractional digits is INVALID!\n" +
          "fractional digits not equal to requested fractional digits.\n" +
          fmt.Sprintf("Calculated Fractional Digits= '%v'\n"+
            "Requested Fractional Digits= '%v'",
            lenAbsFracRunes, n2.precision),
      }
  }

  return n2, nil
}

// shiftPrecisionRight
//
//	Shifts the existing precision of a number string. The position
//	of the decimal point is shifted 'shiftRightPrecision' positions
//	to the right.
//
//	This is equivalent to:
//
//	       result = signedNumStr X 10^shiftRightPrecision
//	                             or
//	signedNumStr Multiplied by 10 raised to the power of 'shiftRightPrecision'.
//
//	Examples
//	========
//
//	signedNumStr    shiftRightPrecision     Result
//
//	"123456.789"             3            "123456789"
//	"123456.789"             2            "12345678.9"
//	"123456.789"             6            "123456789000"
//	"123456789"              6            "123456789000000"
//	"123"                    5            "12300000"
//	"0"                      3            "0"
//	"-123456.789"            3            "-123456789"
//	"-123456789"             6            "-123456789000000"
//
//	       zero ('0') 'shiftRightPrecision' has
//	      no effect on the original number string
//
//	"123456.789"             0            "123456.789"
//	"-123456.789"            0            "-123456.789"
//
//	Numeric Separators
//	==================
//
//	The Numeric Separators originally configured for the current
//	instance of NumStrDto will be copied to the returned instance
//	of NumStrDto.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	numSeps                  NumericSeparatorDto
//	 This parameter contains the Numeric Separators which will be
//	 used to populate the returned instance of 'NumStrDto'.
//
//	signedNumStr             string
//	  A valid number string. The leading digit may optionally be a
//	  '+' or '-' indicating numeric sign value. If '+' or '-'
//	  characters are not present in the first character position,
//	  the number is assumed to represent a positive	numeric value
//	  ('+').
//
//	shiftRightPrecision      uint
//	  The number of digits by which the current decimal point
//	  position in the number string, 'signedNumStr' will be shifted
//	  to the right.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//	  This object encapsulates an error prefix string
//	  which is included in all returned error
//	  messages. Usually, it contains the name of the
//	  calling method or methods listed as a function
//	  chain.
//
//	  If no error prefix information is needed, set
//	  this parameter to 'nil'.
//
//	  Type ErrPrefixDto is included in the 'errpref'
//	  software package:
//	    "github.com/MikeAustin71/errpref".
//
//	Return Values
//	=============
//
//	NumStrDto
//	  This method returns the result of the Shift Left precision
//	  operation in the form of a new 'NumStrDto' instance.
//
//	error
//	  If a processing error is encountered, this returned error
//	  object will be configured with an appropriate error message.
func (nStrDtoTau *numStrDtoTau) shiftPrecisionRight(
  numSeps NumericSeparatorDto,
  signedNumStr string,
  shiftRightPrecision uint,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  nStrDtoTau.lock.Lock()

  defer nStrDtoTau.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoTau.shiftPrecisionRight()",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if len(signedNumStr) == 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "len(signedNumStr) == 0",
        ErrMessage: "Error: Input parameter 'signedNumStr' is INVALID!\n" +
          "'signedNumStr' is a zero length string.",
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  if new(MathProcessUtility).DoesUintExceedMax32BitInt(shiftRightPrecision) {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'shiftRightPrecision' is INVALID!\n" +
          "'shiftRightPrecision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
          fmt.Sprintf("shiftRightPrecision= '%v'", shiftRightPrecision),
      }
  }

  n1, err := new(numStrDtoQuark).parseNumStr(numSeps, signedNumStr, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1, err := new(numStrDtoQuark).parseNumStr(\n" +
          "  numSeps, signedNumStr, ePrefix)",
        ErrContext: fmt.Sprintf("signedNumStr= '%s'\n"+
          "numSeps= '%s'", signedNumStr, numSeps.String()),
        ErrMessage: err.Error(),
      }
  }

  n2 := new(numStrDtoMolecule).newZeroNumStrDto(numSeps, 0)

  iTotalSpecPrecision := 0

  iPrecision := int(shiftRightPrecision)

  iN1Precision := int(n1.precision)

  if iN1Precision > 0 && iPrecision < iN1Precision {

    iTotalSpecPrecision = iN1Precision - iPrecision

  } else {

    iTotalSpecPrecision = 0

  }

  n2.signVal = n1.signVal

  n2.precision = uint(iTotalSpecPrecision)

  lenAbsAllNumRunes := len(n1.absAllNumRunes)

  n1NumStrIsZeroValue, err := new(numStrDtoElectron).isNumStrZeroValue(
    &n1, true, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1NumStrIsZeroValue, err := new(numStrDtoElectron).\n" +
          "  isNumStrZeroValue(&n1, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if n1NumStrIsZeroValue {

    // return nDto.GetZeroNumStrDto(0), nil

    n3 := new(numStrDtoMolecule).newZeroNumStrDto(numSeps, n2.precision)

    return n3, nil
  }

  if int(shiftRightPrecision) > int(n1.precision) {

    for i := 0; i < lenAbsAllNumRunes; i++ {

      n2.absAllNumRunes = append(n2.absAllNumRunes, n1.absAllNumRunes[i])

    }

    deltaPrecision := int(shiftRightPrecision) - int(n1.precision)

    for i := 0; i < deltaPrecision; i++ {

      n2.absAllNumRunes = append(n2.absAllNumRunes, '0')

    }

  } else {

    for i := 0; i < lenAbsAllNumRunes; i++ {

      n2.absAllNumRunes = append(n2.absAllNumRunes, n1.absAllNumRunes[i])

    }
  }

  lenAbsAllNumRunes = len(n2.absAllNumRunes)

  lenAbsFracRunes := iTotalSpecPrecision

  lenAbsIntRunes := lenAbsAllNumRunes - lenAbsFracRunes

  if lenAbsIntRunes <= 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "lenAbsIntRunes <= 0",
        ErrMessage: "Error: Calculated number of integer digits is less than or equal to ZERO.\n" +
          fmt.Sprintf("lenAbsIntRunes= '%v' ", lenAbsIntRunes),
      }
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    &n2, ePrefix.XCpy("Validating 'n2'"))

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "  &n2, ePrefix)",
        ErrContext: "Error: The Final Result NumStrDto instance ('n2') is INVALID!\n" +
          "'n2' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  actualLenAbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunesLength(
    &n2, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "actualLenAbsFracRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsFracRunesLength(&n2, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if actualLenAbsFracRunes != iTotalSpecPrecision {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "actualLenAbsFracRunes != iTotalSpecPrecision",
        ErrMessage: "Error: Calculated number of fractional digits is INVALID!\n" +
          "fractional digits not equal to requested fractional digits.\n" +
          fmt.Sprintf("Calculated Fractional Digits= '%v'\n"+
            "Requested Fractional Digits= '%v'",
            actualLenAbsFracRunes, iTotalSpecPrecision),
      }
  }

  if uint(lenAbsFracRunes) != n2.precision {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "uint(lenAbsFracRunes) != n2.precision",
        ErrMessage: "Error: Calculated number of fractional digits is INVALID!\n" +
          "fractional digits not equal to requested fractional digits.\n" +
          fmt.Sprintf("Calculated Fractional Digits= '%v'\n"+
            "Requested Fractional Digits= '%v'",
            lenAbsFracRunes, n2.precision),
      }
  }

  return n2, nil
}
