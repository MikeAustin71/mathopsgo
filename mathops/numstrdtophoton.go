package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math"
  "math/big"
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

// parseSignedBigIntPrecision
//
//		Receives a signed *Big Int number and a *Big Int precision
//		parameter. This method then generates and returns a new
//		instance of NumStrDto.
//
//		'precision'
//		===========
//
//		'precision' determines the number of digits to the right of the
//		decimal place.
//
//		   signedBigInt    precision      result
//		    946254            3            946.254
//		    946254            0            946254
//		   -946254            3           -946.254
//		   -946254            0           -946254
//
//
//		Maximum Precision Value
//		=======================
//
//		Input parameter 'precision' is a signed *Big Int value.
//		The maximum limit for a 'precision' value is positive number,
//		2,147,483,647 or	2^31 - 1. This is also the maximum
//		allowable limit for a signed 32-bit integer.
//
//		If the 'precision' value exceeds the maximum allowable limit,
//		an error will be returned.
//
//	 Likewise, if 'precison' is less than zero, an error will be
//	 returned.
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
//		configured with the Numeric Separators copied from input
//		parameter 'numSeps'. If these Numeric Separators prove
//		to be invalid, an error will be returned.
func (nStrDtoPhoton *numStrDtoPhoton) parseSignedBigIntPrecision(
  numSeps NumericSeparatorDto,
  signedBigInt *big.Int,
  precision *big.Int,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  nStrDtoPhoton.lock.Lock()

  defer nStrDtoPhoton.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoPhoton.parseSignedBigIntPrecision",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if signedBigInt == nil {

    return NumStrDto{}, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'signedBigInt'",
    }
  }

  bigIntIsNilPtr, bigIntExceedMax32BitInt, bigIntLessThanZero :=
    new(MathProcessUtility).DoesBigIntExceedMax32BitInt(
      precision)

  if bigIntIsNilPtr {

    return NumStrDto{}, &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'precision'",
    }
  }

  if bigIntExceedMax32BitInt {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input Parmater 'precision' is INVALID!\n" +
          "The 'precision' value exceeds maximum allowable limit of 2,147,483,647\n" +
          fmt.Sprintf("precision= %v", precision.Text(10)),
      }
  }

  if bigIntLessThanZero {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input Parmater 'precision' is INVALID!\n" +
          "The 'precision' value is less than zero.\n" +
          fmt.Sprintf("precision= %v", precision.Text(10)),
      }
  }

  precisionUint := uint(precision.Uint64())

  return new(numStrDtoQuark).parseSignedBigInt(
    numSeps, signedBigInt, precisionUint, ePrefix)
}

// setPrecision
//
//		Parses the incoming number string and applies the designated
//		'precision'.
//
//		'precision' determines the number of digits to the right of the
//		decimal place. The boolean parameter 'roundResult' is used to
//		apply rounding in those cases where 'precision' dictates a
//		reduction in the number of digits to the right of the decimal
//		place. See 'Examples' below.
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
//		configured with the Numeric Separators provided by the
//		current instance of 'nDto'. If these Numeric Separators are
//		determined to be invalid, an error will be returned.
//
//		Examples
//		========
//
//		         ------------ Input Parameters ------------
//
//		Example
//		 No       signedNumStr    precision   roundResult     Final Result
//
//		  1       "123456789"         7          false          "123456789.0000000"
//		  2       "123456789"         7          true           "123456789.0000000"
//		  3      "-123456789"         7          false         "-123456789.0000000"
//		  4      "-123456789"         7          true          "-123456789.0000000"
//		  5       "123456.789"        2          true           "123456.79"
//		  6       "123456.789"        2          false          "123456.78"
//		  7       "123456.789"        5          false          "123456.78900"
//		  8       "123.456789"        1          false          "123.4"
//		  9       "123.456789"        1          true           "123.5"
//		 10      "-123.456789"        1          false         "-123.4"
//		 11      "-123.456789"        1          true          "-123.5"
//		 12       "123456.789"        0          true           "123457"
//		 13      "-123456.789"        0          true          "-123457"
//		 14       "123456.789"        0          false          "123456"
//		 15      "-123456.789"        0          false         "-123456"
//		 16       "123457"            1          false          "123457.0"
//		 17       "123457"            1          true           "123457.0"
//		 18      "-123457"            1          false         "-123457.0"
//		 19      "-123457"            1          true          "-123457.0"
//
//		Input Parameters
//		================
//
//		signedNumStr             string
//		  A valid number string
//
//		precision                uint
//		  The 'precision' values designates the number of places to the
//		  right of the decimal point which will be realized upon
//		  completion of this operation.
//
//	   If 'precision' exceeds the maximum limit of 2,147,483,647
//	   an error will be returned
//
//		roundResult              bool
//		  If the 'precision' value is less than the current number of
//		  places to the right of the decimal point, this method will
//		  truncate the existing fractional digits. If 'roundResult' is
//		  set to true, this truncation operation will include rounding
//		  the last digit.
//
//		Return Values
//		=============
//
//		NumStrDto
//		  This returned instance of NumStrDto will contain the result
//		  of the 'set precision' operation described above.
//
//		error
//		  If an error is encountered during processing, this returned
//		  error object will be configured with an appropriate error
//		  message
func (nStrDtoPhoton *numStrDtoPhoton) setPrecision(
  numSeps NumericSeparatorDto,
  signedNumStr string,
  precision uint,
  roundResult bool,
  errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

  nStrDtoPhoton.lock.Lock()

  defer nStrDtoPhoton.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoPhoton.formatForMathOps",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  if len(signedNumStr) == 0 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'signedNumStr' is INVALID!\n" +
          "'signedNumStr' is an empty, zero length string.",
      }
  }

  bigIntMax := big.NewInt(int64(math.MaxInt32))

  bigIntPrecision := big.NewInt(0).SetUint64(uint64(precision))

  compareResult := bigIntPrecision.Cmp(bigIntMax)

  if compareResult == 1 {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
          fmt.Sprintf("'precision' exceeds the maximum limit of %v\n"+
            "precision= '%v", bigIntMax.Text(10), bigIntPrecision.Text(10)),
      }

  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{}, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
      ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
        "Numeric Separators object 'numSeps' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  // n1, err := n0.ParseNumStr(signedNumStr)
  n1, err := new(numStrDtoQuark).parseNumStr(
    numSeps, signedNumStr, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1, err := new(numStrDtoQuark).parseNumStr(\n" +
          "  numSeps, signedNumStr, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  //n2 := new(NumStrDto).New()
  n2 := new(numStrDtoMolecule).newZeroNumStrDto(numSeps, precision)

  n2.signVal = n1.signVal

  //n2.precision = precision
  //n2.thousandsSeparator = nDto.thousandsSeparator
  //n2.decimalSeparator = nDto.decimalSeparator
  //n2.currencySymbol = nDto.currencySymbol

  //n2AbsIntRunes := n2.GetAbsIntRunes()
  n2AbsIntRunes, err := new(numStrDtoGluon).getAbsIntRunes(
    &n2, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n2AbsIntRunes, err := new(numStrDtoGluon).getAbsIntRunes(\n" +
          "  &n2, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // n2AbsFracRunes := n2.GetAbsFracRunes()
  n2AbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunes(
    &n2, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n2AbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunes(\n" +
          "  &n2, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  iSpecPrecision := int(precision)

  lenN1AbsAllNumRunes := len(n1.absAllNumRunes)

  //n1AbsIntRunes := n1.GetAbsIntRunes()

  n1AbsIntRunes, err := new(numStrDtoGluon).getAbsIntRunes(
    &n1, true, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1AbsIntRunes, err := new(numStrDtoGluon).getAbsIntRunes(\n" +
          "  &n1, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // n1AbsFracRunes := n1.GetAbsFracRunes()
  n1AbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunes(
    &n1, false, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "n1AbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunes(\n" +
          "  &n1, false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  lenN1AbsIntRunes := len(n1AbsIntRunes)

  lenN1AbsFracRunes := len(n1AbsFracRunes)

  totalRunes := 0

  if roundResult && lenN1AbsFracRunes > 0 &&
    iSpecPrecision < lenN1AbsFracRunes {

    absAllNumsToRound, isOk := big.NewInt(0).SetString(string(n1.absAllNumRunes), 10)

    if !isOk {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "absAllNumsToRound, isOk := big.NewInt(0).SetString(string(n1.absAllNumRunes), 10)",
          ErrContext: "",
          ErrMessage: "Error: Failed to convert string to big.Int().\n" +
            fmt.Sprintf("string(n1.absAllNumRunes)= '%v'",
              string(n1.absAllNumRunes)),
        }
    }

    bigDeltaPrecision := big.NewInt(int64(lenN1AbsFracRunes - iSpecPrecision - 1))

    base10 := big.NewInt(int64(10))

    roundUp5 := big.NewInt(int64(5))

    roundScaleFactor := big.NewInt(0).Exp(base10, bigDeltaPrecision, nil)

    roundUpNum := big.NewInt(0).Mul(roundUp5, roundScaleFactor)

    roundedAbsAllNums := big.NewInt(0).Add(absAllNumsToRound, roundUpNum)

    actualDeltaPrecision := big.NewInt(int64(lenN1AbsFracRunes - iSpecPrecision))

    actualDeltaScaleFactor := big.NewInt(0).Exp(base10, actualDeltaPrecision, nil)

    actualAbsAllNums := big.NewInt(0).Div(roundedAbsAllNums, actualDeltaScaleFactor)

    n1.absAllNumRunes = []rune{}

    n1AbsIntRunes = []rune{}

    n1AbsFracRunes = []rune{}

    n1.absAllNumRunes = []rune(actualAbsAllNums.String())

    lenN1AbsAllNumRunes = len(n1.absAllNumRunes)

    for i := 0; i < lenN1AbsAllNumRunes; i++ {

      if i < lenN1AbsIntRunes {

        n1AbsIntRunes = append(n1AbsIntRunes, n1.absAllNumRunes[i])

      } else {

        n1AbsFracRunes = append(n1AbsFracRunes, n1.absAllNumRunes[i])
      }
    }

    lenN1AbsIntRunes = len(n1AbsIntRunes)

    lenN1AbsFracRunes = len(n1AbsFracRunes)

    if lenN1AbsAllNumRunes != (lenN1AbsIntRunes + lenN1AbsFracRunes) {

      return NumStrDto{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "lenN1AbsAllNumRunes != (lenN1AbsIntRunes + lenN1AbsFracRunes)",
          ErrMessage: "Error on rounding. Length of IntRunes + FracRunes not equal to Total Runes\n" +
            "lenN1AbsAllNumRunes != (lenN1AbsIntRunes + lenN1AbsFracRunes)" +
            fmt.Sprintf("lenN1AbsAllNumRunes = '%v'\nlenN1AbsIntRunes+lenN1AbsFracRunes = '%v'",
              lenN1AbsAllNumRunes, lenN1AbsIntRunes+lenN1AbsFracRunes),
        }
    }
  } // End Of
  //if roundResult && lenN1AbsFracRunes > 0 &&
  //	iSpecPrecision < lenN1AbsFracRunes {

  if lenN1AbsIntRunes == 0 {

    n2.absAllNumRunes = append(n2.absAllNumRunes, '0')

    n2AbsIntRunes = append(n2AbsIntRunes, '0')
  }

  totalRunes = lenN1AbsIntRunes + iSpecPrecision

  for i := 0; i < totalRunes; i++ {

    if i < lenN1AbsAllNumRunes {

      n2.absAllNumRunes = append(n2.absAllNumRunes, n1.absAllNumRunes[i])

    } else {

      n2.absAllNumRunes = append(n2.absAllNumRunes, '0')

    }

    if i < lenN1AbsIntRunes {

      n2AbsIntRunes = append(n2AbsIntRunes, n1.absAllNumRunes[i])

    } else {

      if i < lenN1AbsAllNumRunes {

        n2AbsFracRunes = append(n2AbsFracRunes, n1.absAllNumRunes[i])

      } else {

        n2AbsFracRunes = append(n2AbsFracRunes, '0')

      }
    }
  } // End Of
  // for i := 0; i < totalRunes; i++

  // err = n2.IsValid(ePrefix)

  err = new(numStrDtoElectron).isValidNumStrDto(
    &n2, ePrefix.XCpy("Validating Final Result n2"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
          "&n2, ePrefix)",
        ErrContext: "Error: Final result 'n2' is INVALID!\n" +
          "'n2' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return n2, nil
}
