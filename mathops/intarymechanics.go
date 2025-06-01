package mathops

import (
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryMechanics struct {
  lock *sync.Mutex
}

// decrementIntegerOne
//
//	Decrements the numeric value of the IntAry instance ('intAry')
//	by subtracting '1'.
func (iaMech *intAryMechanics) decrementIntegerOne(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaMech.lock == nil {
    iaMech.lock = new(sync.Mutex)
  }

  iaMech.lock.Lock()

  defer iaMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMechanics.decrementIntegerOne()",
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

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(
      intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy(Validating 'intAry').String())",
        ErrContext: "IntAry instanace 'intAry' is INVALID!\n" +
          "'intAry' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
    }

  }

  err = new(intAryNanobot).setInternalFlags(
    intAry, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(intAry, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if intAry.isZeroValue || intAry.isIntegerZeroValue {
    intAry.signVal = -1
  }

  intLen := intAry.intAryLen - intAry.precision
  intIdx := intLen - 1
  lastIdx := intAry.intAryLen - 1

  n1 := 0
  n2 := 0
  carry := 0

  intAry.isZeroValue = true
  intAry.isIntegerZeroValue = true

  for i := lastIdx; i >= 0; i-- {
    n1 = int(intAry.intAry[i])

    if i > intIdx {
      //  i > intIdx
      // This must be a fractional digit
      // Retain fractional digits

      if n1 != 0 {
        intAry.isZeroValue = false
      }

      continue

    } else if i == intIdx {

      n2 = n1 + (-1 * intAry.signVal)

      if n2 < 0 {
        n2 = n1 + 10 - 1
        carry = 1

      } else if n2 > 9 {
        n2 = n1 + 1 - 10
        carry = 1

      } else {
        carry = 0
      }

    } else {
      // Must be i < intIdx

      n2 = n1 + ((intAry.signVal * carry) * -1)

      if n2 < 0 {
        n2 = n1 + 10 - carry
        carry = 1
      } else if n2 > 9 {
        n2 = n1 - 10 + carry
        carry = 1
      } else {
        carry = 0
      }

    }

    if n2 != 0 {
      intAry.isZeroValue = false
      intAry.isIntegerZeroValue = false
    }

    intAry.intAry[i] = uint8(n2)

  }

  if intAry.isZeroValue && carry == 0 {
    intAry.signVal = 1
  }

  if carry > 0 {

    intAry.intAry = append([]uint8{1}, intAry.intAry...)
    intAry.intAryLen++

  } else if intAry.intAry[0] == 0 && intLen > 1 {
    intAry.intAry = intAry.intAry[1:]
    intAry.intAryLen--
  }

  return nil
}

// getMagnitudeDigits
//
//	Returns the number of digits in the integer portion of the
//	numeric value in the IntAry instance passed as input parameter
//	'ia'.
//
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaMech *intAryMechanics) getMagnitudeDigits(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (int, error) {

  if iaMech.lock == nil {
    iaMech.lock = new(sync.Mutex)
  }

  iaMech.lock.Lock()

  defer iaMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMechanics.getMagnitudeDigits()",
    "")

  if err != nil {
    return 0, err
  }

  if intAry == nil {

    return 0,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'ia'",
      }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return 0,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "  intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  err = new(intAryNanobot).setInternalFlags(
    intAry, ePrefix)

  if err != nil {

    return 0, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = new(intAryNanobot).setInternalFlags(ia, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  intAryMagnitude :=
    intAry.intAryLen - intAry.precision - intAry.firstDigitIdx

  return intAryMagnitude, nil
}

// setNumericSeparators
//
//	Used to assign values for the Decimal and Thousands separators
//	as well as the Currency Symbol to be used in displaying the
//	current number string.
//
//	Different nations and cultures use different symbols to delimit
//	numerical values. In the USA and many other countries, a period
//	character ('.') is used to delimit integer and fractional digits
//	within a numeric value (123.45). Likewise, thousands may be
//	delimited by a comma (','). Currency signs very by nationality.
//	For instance, the USA, Canada and several other countries use
//	the dollar sign ($) as a currency symbol.
//
//	For a list of major world currency symbols see:
//
//	  MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//
//	  http://www.xe.com/symbols.php
//
//	Note: If zero values are submitted as input for separator values,
//	those values will default to USA standards.
//
//	USA Examples
//	============
//
//	Decimal Separator period ('.')    = 123.456
//	Thousands Separator comma (',')   = 1,000,000,000
//	Currency Symbol dollar sign ('$') = $123
func (iaMech *intAryMechanics) setNumericSeparators(
  intAry *IntAry,
  decimalSeparator,
  thousandsSeparator,
  currencySymbol rune,
  callingFunction string) error {

  if iaMech.lock == nil {
    iaMech.lock = new(sync.Mutex)
  }

  iaMech.lock.Lock()

  defer iaMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    callingFunction,
    "intAryMechanics.setNumericSeparators()",
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

  if decimalSeparator == 0 {
    decimalSeparator = '.'
  }

  if thousandsSeparator == 0 {
    thousandsSeparator = ','
  }

  if currencySymbol == 0 {
    currencySymbol = '$'
  }

  intAry.SetDecimalSeparator(decimalSeparator)

  intAry.SetThousandsSeparator(thousandsSeparator)

  intAry.SetCurrencySymbol(currencySymbol)

  return nil
}

// setNumericSeparatorsDto
//
// Sets the values of numeric separators:
//
//	  decimal point separator
//	  thousands separator
//	  currency symbol
//
//	These vales are based on data transmitted through input
//	parameter 'customSeparators' of type NumericSeparatorDto.
//
//	If any of the values contained in input parameter
//	'customSeparators' are set to zero, an error will be returned.
func (iaMech *intAryMechanics) setNumericSeparatorsDto(
  intAry *IntAry,
  customSeparators NumericSeparatorDto,
  callingFunction string) error {

  if iaMech.lock == nil {
    iaMech.lock = new(sync.Mutex)
  }

  iaMech.lock.Lock()

  defer iaMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    callingFunction,
    "intAryMechanics.setNumericSeparatorsDto()",
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

  err = customSeparators.IsValid(ePrefix.XCpy("Validating 'customSeparators'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = customSeparators.IsValid(ePrefix.XCpy(\n" +
        "  Validating 'customSeparators').String())",
      ErrContext: "Input parameter 'customSeparators' is invalid.\n" +
        "'customSeparators' FAILED validation tests.",
      ErrMessage: err.Error(),
    }
  }

  intAry.decimalSeparator = customSeparators.DecimalSeparator

  intAry.thousandsSeparator = customSeparators.ThousandsSeparator

  intAry.currencySymbol = customSeparators.CurrencySymbol

  return nil
}

// setAbsoluteValue
//
//	Converts the numeric value of the IntAry object passed as
//	input parameter 'IntAry' to its absolute value.
func (iaMech *intAryMechanics) setAbsoluteValue(
  intAry *IntAry,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaMech.lock == nil {
    iaMech.lock = new(sync.Mutex)
  }

  iaMech.lock.Lock()

  defer iaMech.lock.Unlock()
  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMechanics.setAbsoluteValue()",
    "")

  if err != nil {
    return err
  }

  err = new(intAryNeutron).setSign(intAry, 1, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryNeutron).\n" +
        "   setSign(intAry, 1, ePrefix)\n",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
