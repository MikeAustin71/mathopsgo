package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type numStrDtoElectron struct {
  lock sync.Mutex
}

// emptyNumStrDto
//
//	Receives a pointer to an instance of NumStrDto and proceeds to
//	reset all internal member data values to their zero or initial
//	states.
func (nStrDtoElectron *numStrDtoElectron) emptyNumStrDto(
  numStrDto *NumStrDto) {

  nStrDtoElectron.lock.Lock()

  defer nStrDtoElectron.lock.Unlock()

  if numStrDto == nil {
    return
  }

  numStrDto.signVal = 0
  numStrDto.absAllNumRunes = []rune{}
  numStrDto.precision = 0
  numStrDto.decimalSeparator = 0
  numStrDto.thousandsSeparator = 0
  numStrDto.currencySymbol = 0

  return
}

// isValidNumStrDto
//
//	Performs a diagnostic review of the current NumStrDto instance
//	and returns 'nil' if the NumStrDto object is valid in all
//	respects.
//
//	If the NumStrDto instance is judged invalid, an error message
//	is returned.
func (nStrDtoElectron *numStrDtoElectron) isValidNumStrDto(
  numStrDto *NumStrDto,
  errPrefDto *ePref.ErrPrefixDto) error {

  nStrDtoElectron.lock.Lock()

  defer nStrDtoElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoElectron.isValidNumStrDto()",
    "")

  if err != nil {
    return err
  }

  if numStrDto == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'numStrDto'",
    }
  }

  if numStrDto.thousandsSeparator == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "numStrDto.thousandsSeparator == 0",
      ErrMessage: "Error: Thousands Separator cannot be '0'\n" +
        "NumStrDto Numeric Separators are Invalid!",
    }
  }

  if numStrDto.decimalSeparator == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "numStrDto.decimalSeparator == 0",
      ErrMessage: "Error: Decimal Separator cannot be '0'\n" +
        "NumStrDto Numeric Separators are Invalid!",
    }
  }

  if numStrDto.currencySymbol == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "numStrDto.currencySymbol == 0",
      ErrMessage: "Error: Currency Separator cannot be '0'\n" +
        "NumStrDto Numeric Separators are Invalid!",
    }
  }

  lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

  if lenAbsAllNumRunes == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: " len(numStrDto.absAllNumRunes) == 0",
      ErrMessage: "Error: Number string is a ZERO length array!\n" +
        "NumStrDto object is invalid.",
    }
  }

  if int(numStrDto.precision) >= lenAbsAllNumRunes {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "precision > number string length",
      ErrMessage: "Error: 'precision' does not correlate with number string.\n" +
        "NumStrDto object is invalid!",
    }
  }

  if numStrDto.signVal != 1 && numStrDto.signVal != -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "numStrDto.signVal != 1 && numStrDto.signVal != -1",
      ErrMessage: "Error: sign Value is INVALID. Should be +1 or -1.\n" +
        "NumStrDto object is invalid!",
    }
  }

  for i := 0; i < lenAbsAllNumRunes; i++ {

    if numStrDto.absAllNumRunes[i] < '0' || numStrDto.absAllNumRunes[i] > '9' {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Number string character < 0 Or > 9",
        ErrMessage: "Error: Non-Numeric character found in number string!\n" +
          "NumStrDto object is invalid!",
      }
    }
  }

  return nil
}

// isNumStrZeroValue
//
//	Returns 'true' if all the digits in the number string for the
//	current NumStrDto instance are zero.
func (nStrDtoElectron *numStrDtoElectron) isNumStrZeroValue(
  numStrDto *NumStrDto,
  validateNumStrDto bool,
  errPrefDto *ePref.ErrPrefixDto) (bool, error) {

  nStrDtoElectron.lock.Lock()

  defer nStrDtoElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoElectron.isNumStrZeroValue()",
    "")

  if err != nil {
    return true, err
  }

  if numStrDto == nil {

    return true,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'numStrDto'",
      }
  }

  lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

  if validateNumStrDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

    if err != nil {

      return true,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  numStrDto, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  } else {

    if lenAbsAllNumRunes == 0 {

      return true,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "len(numStrDto.absAllNumRunes) == 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The internal rune array of numeric characters is empty.\n" +
            "This instance of NumStrDto is either corrupted or uninitialized!",
        }
    }

    precision := int(numStrDto.precision)

    if precision < 0 {

      return true,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "precision < 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is less than zero.",
        }
    }

    if precision > lenAbsAllNumRunes {

      return true,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is greater than the internal numeric digits array.",
          ErrMessage: "",
        }
    }
  }

  for i := 0; i < lenAbsAllNumRunes; i++ {

    if numStrDto.absAllNumRunes[i] != '0' {

      return false, nil

    }
  }

  return true, nil
}

// setCurrencySymbol
//
//	Assigns the input parameter rune as the currency symbol to be
//	used by the instance of NumStrDto passed as input parameter,
//	'numStrDto'.
//
//	If a zero value is submitted as input for the 'currencySymbol'
//	rune, an error will be returned.
//
//	For a list of Major Currency Unicode Symbols, see constants
//	located in: /mathops/mathopsconstants.go
//
//	In the USA, the currency symbol is the dollar sign ('$').
//
//	USA Example
//	===========
//
//	  $123.45
func (nStrDtoElectron *numStrDtoElectron) setCurrencySymbol(
  numStrDto *NumStrDto,
  currencySymbol rune,
  errPrefDto *ePref.ErrPrefixDto) error {

  nStrDtoElectron.lock.Lock()

  defer nStrDtoElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "nStrDtoElectron.setCurrencySymbol()",
    "")

  if err != nil {
    return err
  }

  if numStrDto == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'numStrDto'",
    }
  }

  if currencySymbol == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "Error: Input parameter 'currencySymbol' is INVALID!\n" +
        "rune 'currencySymbol' has a value of zero.",
    }

  }

  numStrDto.currencySymbol = currencySymbol

  return nil
}

// setDecimalSeparator
//
//	Assigns the input parameter rune as the decimal separator to be
//	used by the instance of NumStrDto passed as input parameter,
//	'numStrDto'.
//
//	If a zero value is submitted as input for the 'decimalSeparator'
//	rune, an error will be returned.
//
//	Decimal separators separate integer and fractional components
//	of a numeric value.
//
//	In the USA, the decimal separator is the period character ('.').
//
//	USA Example
//	===========
//
//	123.45892
func (nStrDtoElectron *numStrDtoElectron) setDecimalSeparator(
  numStrDto *NumStrDto,
  decimalSeparator rune,
  errPrefDto *ePref.ErrPrefixDto) error {

  nStrDtoElectron.lock.Lock()

  defer nStrDtoElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoElectron.setDecimalSeparator()",
    "")

  if err != nil {
    return err
  }

  if numStrDto == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'numStrDto'",
    }
  }

  if decimalSeparator == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "Error: Input parameter 'decimalSeparator' is INVALID!\n" +
        "rune 'decimalSeparator' has a value of zero.",
    }

  }

  numStrDto.decimalSeparator = decimalSeparator

  return nil
}

// setThousandsSeparator
//
//	Assigns the input parameter rune as the thousands separator to
//	be used by the instance of NumStrDto passed as input parameter,
//	'numStrDto'.
//
//	If a zero value is submitted as input for the 'thousandsSeparator'
//	rune, an error will be returned.
//
//	In the USA, the thousands separator is the comma character
//	(',').
//
//	USA Example
//	===========
//
//	123,456,789
func (nStrDtoElectron *numStrDtoElectron) setThousandsSeparator(
  numStrDto *NumStrDto,
  thousandsSeparator rune,
  errPrefDto *ePref.ErrPrefixDto) error {

  nStrDtoElectron.lock.Lock()

  defer nStrDtoElectron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoElectron.setThousandsSeparator()",
    "")

  if err != nil {
    return err
  }

  if numStrDto == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'numStrDto'",
    }
  }

  if thousandsSeparator == 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "Error: Input parameter 'ghousandsSeparator' is INVALID!\n" +
        "rune 'ghousandsSeparator' has a value of zero.",
    }

  }

  numStrDto.thousandsSeparator = thousandsSeparator

  return nil
}
