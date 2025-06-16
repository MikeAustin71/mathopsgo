package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type numStrDtoPhoton struct {
  lock sync.Mutex
}

// formatForMathOps
//
//	Receives two NumStrDto objects and converts their number
//	strings such that both have the same number of integer and
//	fractional digits while maintaining their original numeric
//	values. This transform will facilitate the performance of
//	string based math operations such as addition and subtraction.
//
//	The return values represent the formatted NumStrDto objects.
//	The first NumStrDto returned always contains the larger
//	absolute value. The second NumStrDto always contains the
//	absolute numeric value which is less than or equal to the first
//	NumStrDto object returned.
//
//	The third parameter returned by this method is an integer
//	('compare') value which will always be set to '1' or '0'.
//
//	'1' indicates that the absolute value of the first NumStrDto
//	object ('n1DtoOut') returned by this method is greater than the
//	second NumStrDto object ('n2DtoOut') returned by this method.
//
//	If the returned integer ('compare') value returned is zero, it
//	signals that the absolute values (not the signed values) of
//	both returned NumStrDto objects are equal.
//
//	If the absolute value of 'n1Dto' is less than 'n2Dto', return
//	value 'n1DtoOut' will be populated with 'n2Dto' values, return
//	value 'n2DtoOut' will be populated with 'n1Dto' values and return
//	parameter 'isOrderReversed' will be set to 'true'.
func (nStrDtoPhoton *numStrDtoPhoton) formatForMathOps(
  numSeps NumericSeparatorDto,
  n1Dto *NumStrDto,
  validateN1Dto bool,
  n2Dto *NumStrDto,
  validateN2Dto bool,
  errPrefDto *ePref.ErrPrefixDto) (
  n1DtoOut NumStrDto,
  n2DtoOut NumStrDto,
  compare int,
  isOrderReversed bool,
  err error) {

  nStrDtoPhoton.lock.Lock()

  defer nStrDtoPhoton.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoPhoton.formatForMathOps",
    "")

  if err != nil {
    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      err
  }

  if n1Dto == nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n1Dto'",
      }
  }

  if n2Dto == nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'n2Dto'",
      }
  }

  nStrElectron := new(numStrDtoElectron)

  if validateN1Dto {

    err = nStrElectron.isValidNumStrDto(
      n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  n1Dto, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateN2Dto {

    err = nStrElectron.isValidNumStrDto(
      n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
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

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Numeric Separators input paramter 'numSeps' is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  lenN1AllRunes := 0
  lenN1IntRunes := 0
  lenN1FracRunes := 0
  lenN2AllRunes := 0
  lenN2IntRunes := 0
  lenN2FracRunes := 0

  compare, err = new(numStrDtoAtom).compareAbsoluteValues(
    n1Dto, false, n2Dto, false, ePrefix.XCpy("n1Dto vs. n2Dto"))

  if err != nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "compare, err = new(numStrDtoAtom).compareAbsoluteValues(\n" +
          "n1Dto, true, n2Dto, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }

  }

  // Original code
  //if compare == 1 {
  //	n1DtoOut = n1Dto.CopyOut()
  //	n2DtoOut = n2Dto.CopyOut()
  //} else if compare == -1 {
  //	n1DtoOut = n2Dto.CopyOut()
  //	n2DtoOut = n1Dto.CopyOut()
  //	isOrderReversed = true
  //	compare = 1
  //} else {
  //	// compare must be zero
  //	n1DtoOut = n1Dto.CopyOut()
  //	n2DtoOut = n2Dto.CopyOut()
  //}

  nStrDtoMolecule := new(numStrDtoMolecule)

  if compare == -1 {

    // n1DtoOut = n2Dto.CopyOut()
    err = nStrDtoMolecule.copy(&n1DtoOut, n2Dto, false, ePrefix.XCpy("n2Dto->n1DtoOut"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = nStrDtoMolecule.copy(&n1DtoOut, n2Dto, false, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    //n2DtoOut = n1Dto.CopyOut()
    err = nStrDtoMolecule.copy(&n2DtoOut, n1Dto, false, ePrefix.XCpy("n1Dto->n2DtoOut"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = nStrDtoMolecule.copy(&n2DtoOut, n1Dto, false, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    isOrderReversed = true
    compare = 1

  } else {

    // n1DtoOut = n1Dto.CopyOut()
    err = nStrDtoMolecule.copy(&n1DtoOut, n1Dto, false, ePrefix.XCpy("n2Dto->n1DtoOut"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = nStrDtoMolecule.copy(&n1DtoOut, n1Dto, false, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    //n2DtoOut = n2Dto.CopyOut()
    err = nStrDtoMolecule.copy(&n2DtoOut, n2Dto, false, ePrefix.XCpy("n1Dto->n2DtoOut"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = nStrDtoMolecule.copy(&n2DtoOut, n2Dto, false, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  nStrGluon := new(numStrDtoGluon)

  // n1DtoOutAbsIntRunes := n1DtoOut.GetAbsIntRunes()
  n1DtoOutAbsIntRunes, err := nStrGluon.getAbsIntRunes(
    &n1DtoOut, true, ePrefix)

  if err != nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1DtoOutAbsIntRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsIntRunes(&n1DtoOut, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  //n1DtoOutAbsFracRunes := n1DtoOut.GetAbsFracRunes()
  n1DtoOutAbsFracRunes, err := nStrGluon.getAbsFracRunes(
    &n1DtoOut, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1DtoOutAbsFracRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsFracRunes(&n1DtoOut, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  //n2DtoOutAbsIntRunes := n2DtoOut.GetAbsIntRunes()
  n2DtoOutAbsIntRunes, err := nStrGluon.getAbsIntRunes(
    &n2DtoOut, true, ePrefix)

  if err != nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n2DtoOutAbsIntRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsIntRunes(&n2DtoOut, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // n2DtoOutAbsFracRunes := n2DtoOut.GetAbsFracRunes()
  n2DtoOutAbsFracRunes, err := nStrGluon.getAbsFracRunes(
    &n2DtoOut, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n2DtoOutAbsFracRunes, err := new(numStrDtoGluon).\n" +
          "  getAbsFracRunes(&n2DtoOut, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if n1DtoOut.precision > n2DtoOut.precision {

    deltaPrecision := n1DtoOut.precision - n2DtoOut.precision

    for i := uint(0); i < deltaPrecision; i++ {

      n2DtoOut.absAllNumRunes = append(n2DtoOut.absAllNumRunes, '0')

      n2DtoOutAbsFracRunes = append(n2DtoOutAbsFracRunes, '0')
    }

    lenN2AllRunes = len(n2DtoOut.absAllNumRunes)

    lenN2IntRunes = len(n2DtoOutAbsIntRunes)

    lenN2FracRunes = len(n2DtoOutAbsFracRunes)

    n2DtoOut.precision = n1DtoOut.precision

    //err = n2DtoOut.IsValid(ePrefix)
    err = nStrElectron.isValidNumStrDto(
      &n2DtoOut, ePrefix.XCpy("Validating 'n2DtoOut'"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).\n" +
            "  isValidNumStrDto(&n2DtoOut, ePrefix)",
          ErrContext: "Error: Intermediate calculation result.\n" +
            "'n2DtoOut' is not valid. Failed Validation Tests.",
          ErrMessage: err.Error(),
        }
    }

    lenN1AllRunes = len(n1DtoOut.absAllNumRunes)

    lenN1IntRunes = len(n1DtoOutAbsIntRunes)

    lenN1FracRunes = len(n1DtoOutAbsFracRunes)

  } else if n1DtoOut.precision < n2DtoOut.precision {

    deltaPrecision := n2DtoOut.precision - n1DtoOut.precision

    for i := uint(0); i < deltaPrecision; i++ {

      n1DtoOut.absAllNumRunes = append(n1DtoOut.absAllNumRunes, '0')

      n1DtoOutAbsFracRunes = append(n1DtoOutAbsFracRunes, '0')
    }

    lenN1AllRunes = len(n1DtoOut.absAllNumRunes)

    lenN1IntRunes = len(n1DtoOutAbsIntRunes)

    lenN1FracRunes = len(n1DtoOutAbsFracRunes)

    n1DtoOut.precision = n2DtoOut.precision

    //err = n1DtoOut.IsValid(ePrefix)
    //
    err = nStrElectron.isValidNumStrDto(
      &n1DtoOut, ePrefix.XCpy("Validating 'n1DtoOut'"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).\n" +
            "  isValidNumStrDto(&n1DtoOut, ePrefix)",
          ErrContext: "Error: Intermediate calculation result.\n" +
            "'n1DtoOut' is not valid. Failed Validation Tests.",
          ErrMessage: err.Error(),
        }
    }

    lenN2AllRunes = len(n2DtoOut.absAllNumRunes)

    lenN2IntRunes = len(n2DtoOutAbsIntRunes)

    lenN2FracRunes = len(n2DtoOutAbsFracRunes)

  } else {
    // n1DtoOut.precision == n2DtoOut.precision

    lenN1AllRunes = len(n1DtoOut.absAllNumRunes)

    lenN1IntRunes = len(n1DtoOutAbsIntRunes)

    lenN1FracRunes = len(n1DtoOutAbsFracRunes)

    lenN2AllRunes = len(n2DtoOut.absAllNumRunes)

    lenN2IntRunes = len(n2DtoOutAbsIntRunes)

    lenN2FracRunes = len(n2DtoOutAbsFracRunes)

  }

  if lenN2IntRunes > lenN1IntRunes {

    var absAllRunes []rune

    var absIntRunes []rune

    deltaRunes := lenN2IntRunes - lenN1IntRunes

    for i := 0; i < deltaRunes; i++ {

      absAllRunes = append(absAllRunes, '0')

      absIntRunes = append(absIntRunes, '0')
    }

    for j := 0; j < lenN1AllRunes; j++ {

      absAllRunes = append(absAllRunes, n1DtoOut.absAllNumRunes[j])

      if j < lenN1IntRunes {

        absIntRunes = append(absIntRunes, n1DtoOutAbsIntRunes[j])
      }

    }

    n1DtoOut.absAllNumRunes = absAllRunes

    n1DtoOutAbsIntRunes = absIntRunes

    lenN1AllRunes = len(n1DtoOut.absAllNumRunes)

    lenN1IntRunes = len(n1DtoOutAbsIntRunes)

    // err = n1DtoOut.IsValid(ePrefix)
    err = nStrElectron.isValidNumStrDto(
      &n1DtoOut, ePrefix.XCpy("Validating 'n1DtoOut'"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).\n" +
            "  isValidNumStrDto(&n1DtoOut, ePrefix)",
          ErrContext: "Error-2: Intermediate calculation result\n" +
            "'n1DtoOut' is not valid. Failed Validation Tests.",
          ErrMessage: err.Error(),
        }
    }

    // End of if lenN2IntRunes > lenN1IntRunes
  } else if lenN1IntRunes > lenN2IntRunes {

    var absAllRunes []rune

    var absIntRunes []rune

    deltaRunes := lenN1IntRunes - lenN2IntRunes

    for i := 0; i < deltaRunes; i++ {

      absAllRunes = append(absAllRunes, '0')

      absIntRunes = append(absIntRunes, '0')
    }

    for j := 0; j < lenN2AllRunes; j++ {

      absAllRunes = append(absAllRunes, n2DtoOut.absAllNumRunes[j])

      if j < lenN2IntRunes {

        absIntRunes = append(absIntRunes, n2DtoOutAbsIntRunes[j])
      }
    }

    n2DtoOut.absAllNumRunes = absAllRunes

    n2DtoOutAbsIntRunes = absIntRunes

    lenN2AllRunes = len(n2DtoOut.absAllNumRunes)

    lenN2IntRunes = len(n2DtoOutAbsIntRunes)

    //err := n2DtoOut.IsValid(ePrefix)
    err = nStrElectron.isValidNumStrDto(
      &n2DtoOut, ePrefix.XCpy("Validating 'n2DtoOut'"))

    if err != nil {

      return NumStrDto{},
        NumStrDto{},
        0,
        false,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).\n" +
            "  isValidNumStrDto(&n2DtoOut, ePrefix)",
          ErrContext: "Error-2: Intermediate calculation result\n" +
            "'n2DtoOut' is not valid. Failed Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  } // End of else if lenN1IntRunes > lenN2IntRunes

  if lenN1AllRunes != lenN2AllRunes {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "lenN1AllRunes != lenN2AllRunes\n",
        ErrMessage: "Error: n1 and n2 AllNumRune arrays are NOT equal in length.\n" +
          fmt.Sprintf("n1 length= '%v' n2 length= '%v'",
            lenN1AllRunes, lenN2AllRunes),
      }
  }

  if lenN1IntRunes != lenN2IntRunes {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "lenN1IntRunes != lenN2IntRunes\n",
        ErrMessage: "Error: n1 and n2 Integer Rune arrays are NOT equal in length.\n" +
          fmt.Sprintf("n1 IntRunes length= '%v'\nn2 IntRunes length= '%v'",
            lenN1IntRunes, lenN2IntRunes),
      }
  }

  if lenN1FracRunes != lenN2FracRunes {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "lenN1FracRunes != lenN2FracRunes\n",
        ErrMessage: "Error: n1 and n2 Fractional Rune arrays are NOT equal in length.\n" +
          fmt.Sprintf("n1 FracRunes length= '%v'\nn2 FracRunes length= '%v'",
            lenN1FracRunes, lenN2FracRunes),
      }
  }

  if n1DtoOut.precision != n2DtoOut.precision {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: " n1DtoOut.precision != n2DtoOut.precision\n",
        ErrMessage: "Error: n1 and n2 'precision' values are NOT equal.\n" +
          fmt.Sprintf("n1DtoOut precision= '%v'\nn2DtoOut precision= '%v'",
            n1DtoOut.precision, n2DtoOut.precision),
      }

  }

  //err = n1DtoOut.IsValid(ePrefix + "n1DtoOut - ")
  err = nStrElectron.isValidNumStrDto(
    &n1DtoOut, ePrefix.XCpy("Validating 'n1DtoOut'"))

  if err != nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).\n" +
          "  isValidNumStrDto(&n1DtoOut, ePrefix)",
        ErrContext: "Error: Final calculation result.\n" +
          "'n1DtoOut' is not valid. Failed Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  // err = n2DtoOut.IsValid(ePrefix + "n2DtoOut - ")
  err = nStrElectron.isValidNumStrDto(
    &n2DtoOut, ePrefix.XCpy("Validating 'n2DtoOut'"))

  if err != nil {

    return NumStrDto{},
      NumStrDto{},
      0,
      false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).\n" +
          "  isValidNumStrDto(&n2DtoOut, ePrefix)",
        ErrContext: "Error: Final calculation result.\n" +
          "'n2DtoOut' is not valid. Failed Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return n1DtoOut, n2DtoOut, compare, isOrderReversed, nil
}
