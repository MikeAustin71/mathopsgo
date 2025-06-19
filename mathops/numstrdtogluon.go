package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type numStrDtoGluon struct {
  lock sync.Mutex
}

// getAbsAllNumRunes
//
//	Returns an array of runes representing all the integer and
//	fractional digits included in the current NumStrDto instance.
//	The rune array returned will consist of numeric digits with no
//	sign value prefixed. This effectively returns the absolute
//	value of all integer and fractional digits combined in one rune
//	array (there is no decimal point).
func (nStrDtoGluon *numStrDtoGluon) getAbsAllNumRunes(
  numStrDto *NumStrDto,
  validateNumStrDto bool,
  errPrefDto *ePref.ErrPrefixDto) ([]rune, error) {

  nStrDtoGluon.lock.Lock()

  defer nStrDtoGluon.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoGluon.getAbsAllNumRunes()",
    "")

  if err != nil {
    return []rune{}, err
  }

  if numStrDto == nil {

    return []rune{},
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

      return []rune{},
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

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "len(numStrDto.absAllNumRunes) == 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The internal rune array of numeric characters is empty.",
        }
    }

    precision := int(numStrDto.precision)

    if precision < 0 {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "precision < 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is less than zero.",
        }
    }

    if precision > lenAbsAllNumRunes {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is greater than the internal numeric digits array.",
          ErrMessage: "",
        }
    }
  }

  outRunes := make([]rune, lenAbsAllNumRunes)

  for i := 0; i < lenAbsAllNumRunes; i++ {
    outRunes[i] = numStrDto.absAllNumRunes[i]
  }

  return outRunes, nil
}

// getAbsFracRunes
//
//	Returns all the fractional digits to the right of the decimal
//	place in the current NumStrDto instance as an array of runes.
//	The rune array is not signed; that is, the rune array does not
//	contain a '+' or '-' character in the first array position. The
//	rune array is therefore said to represent the absolute value of
//	the fractional digits in the current NumStrDto numeric value.
func (nStrDtoGluon *numStrDtoGluon) getAbsFracRunes(
  numStrDto *NumStrDto,
  validateNumStrDto bool,
  errPrefDto *ePref.ErrPrefixDto) ([]rune, error) {

  nStrDtoGluon.lock.Lock()

  defer nStrDtoGluon.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoGluon.getAbsFracRunes()",
    "")

  if err != nil {
    return []rune{}, err
  }

  if numStrDto == nil {

    return []rune{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'numStrDto'",
      }
  }

  precision := int(numStrDto.precision)

  lenAllNums := len(numStrDto.absAllNumRunes)

  if validateNumStrDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

    if err != nil {

      return []rune{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  numStrDto, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  } else {

    if lenAllNums == 0 {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "len(numStrDto.absAllNumRunes) == 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The internal rune array of numeric characters is empty.",
        }
    }

    if precision < 0 {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "precision < 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is less than zero.",
        }
    }

    if precision > lenAllNums {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is greater than the internal numeric digits array.",
          ErrMessage: "",
        }
    }

  }

  absFracRunes := make([]rune, precision)

  lenIntNums := lenAllNums - precision

  for i := lenIntNums; i < lenAllNums; i++ {
    absFracRunes[i-lenIntNums] = numStrDto.absAllNumRunes[i]
  }

  return absFracRunes, nil
}

// getAbsFracRunesLength
//
//	Returns the length of the fractional digits in the number
//	string.
func (nStrDtoGluon *numStrDtoGluon) getAbsFracRunesLength(
  numStrDto *NumStrDto,
  validateNumStrDto bool,
  errPrefDto *ePref.ErrPrefixDto) (int, error) {

  nStrDtoGluon.lock.Lock()

  defer nStrDtoGluon.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoGluon.getAbsIntRunes()",
    "")

  if err != nil {
    return 0, err
  }

  if numStrDto == nil {

    return 0,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'numStrDto'",
      }
  }

  lenAllNums := len(numStrDto.absAllNumRunes)

  precision := int(numStrDto.precision)

  if validateNumStrDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

    if err != nil {

      return 0,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  numStrDto, ePrefix)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  } else {

    if lenAllNums == 0 {

      return 0,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "len(numStrDto.absAllNumRunes) == 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The internal rune array of numeric characters is empty.",
        }
    }

    if precision < 0 {

      return 0,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "precision < 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is less than zero.",
        }
    }

    if precision > lenAllNums {

      return 0,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is greater than the internal numeric digits array.",
          ErrMessage: "",
        }
    }

  }

  return precision, nil
}

// getAbsIntRunes
//
//	Returns all the integer digits included in the current
//	NumStrDto numeric value as an array of runes. The returned rune
//	array does not contain a sign value in the first position and
//	therefore represents the absolute or positive value of all the
//	integer digits. The integer digits of a NumStrDto numeric
//	includes all the digits to the left of the decimal point.
//
//	If the current NumStrDto consists of zero integers and
//	fractional digits (Example: '0.123456'), this method will
//	return a rune array consisting one array element with a '0'
//	value.
func (nStrDtoGluon *numStrDtoGluon) getAbsIntRunes(
  numStrDto *NumStrDto,
  validateNumStrDto bool,
  errPrefDto *ePref.ErrPrefixDto) ([]rune, error) {

  nStrDtoGluon.lock.Lock()

  defer nStrDtoGluon.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoGluon.getAbsIntRunes()",
    "")

  if err != nil {
    return []rune{}, err
  }

  if numStrDto == nil {

    return []rune{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'numStrDto'",
      }
  }

  lenAllNums := len(numStrDto.absAllNumRunes)

  precision := int(numStrDto.precision)

  if validateNumStrDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

    if err != nil {

      return []rune{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
            "  numStrDto, ePrefix)",
          ErrContext: "Error: Input parameter 'numStrDto' is INVALID!",
          ErrMessage: err.Error(),
        }
    }
  } else {

    if lenAllNums == 0 {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "len(numStrDto.absAllNumRunes) == 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The internal rune array of numeric characters is empty.",
        }
    }

    if precision < 0 {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "precision < 0",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is less than zero.",
        }
    }

    if precision > lenAllNums {

      return []rune{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "precision > lenAllNums",
          ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
            "The 'precision' value is greater than the internal numeric digits array.",
        }
    }

  }

  lenIntNum := lenAllNums - precision

  absIntRunes := make([]rune, lenIntNum)

  for i := 0; i < lenIntNum; i++ {
    absIntRunes[i] = numStrDto.absAllNumRunes[i]
  }

  return absIntRunes, nil
}

// getScaleFactor
//
//	Returns the scale factor for the number string encapsulated by
//	the current instance of NumStrDto.
//
//	Scale factor is defined by 10 raised to the power of
//	'precision' (nDto.precision).  nDto.precision is the number of
//	digits to the right of the decimal point.
//
//	This method will fail and return an error if the current instance
//	of NumStrDto is invalid.
func (nStrDtoGluon *numStrDtoGluon) getScaleFactor(
  nDto *NumStrDto,
  validateNumStrDto bool,
  errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

  nStrDtoGluon.lock.Lock()

  defer nStrDtoGluon.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "numStrDtoGluon.getScaleFactor()",
    "")

  if err != nil {
    return big.NewInt(0), err
  }

  if nDto == nil {

    return big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'nDto'",
      }
  }

  if validateNumStrDto {

    err = new(numStrDtoElectron).isValidNumStrDto(
      nDto, ePrefix.XCpy("Validating 'nDto'"))

    if err != nil {
      return big.NewInt(0),
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
            "'nDto' FAILED Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if nDto.precision == 0 {
    return big.NewInt(int64(1)), nil
  }

  base10 := big.NewInt(0).SetInt64(int64(10))

  bigPrecision := big.NewInt(0).SetInt64(int64(nDto.precision))

  scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

  return scaleFactor, nil
}
