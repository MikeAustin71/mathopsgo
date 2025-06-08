package mathops

import (
  "bytes"
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryAtom struct {
  lock *sync.Mutex
}

// equal
//
//	Receives two instances of IntAry and compares the values of all
//	member data fields to determine if they are equivalent in all
//	respects.
//
//	Returns 'true' if all member field values of 'iAry1' are equal
//	to the corresponding field values of 'iAry2'.
//
//	Note that the BackUp fields for both compared IntAry objects
//	are NOT included in the 'Equals' comparison.
//
//	If any errors are encountered, a boolean value of 'false' is
//	returned
func (iaAtom *intAryAtom) equal(
  iAry1 *IntAry,
  validateiAry1 bool,
  iAry2 *IntAry,
  validateiAry2 bool,
  errPrefDto *ePref.ErrPrefixDto) (bool, error) {

  if iaAtom.lock == nil {
    iaAtom.lock = new(sync.Mutex)
  }

  iaAtom.lock.Lock()

  defer iaAtom.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryNeutron.equal()",
    "")

  if err != nil {
    return false, err
  }

  if iAry1 == nil {

    return false,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'iAry1'",
      }
  }

  if iAry2 == nil {

    return false,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'iAry2'",
      }
  }

  iaElectron := new(intAryElectron)

  if validateiAry1 {

    err = iaElectron.isValidIntAry(iAry1, ePrefix.XCpy("Validating 'iAry1'").String())

    if err != nil {

      return false,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = iaElectron.isValidIntAry(iAry1, ePrefix.XCpy(Validating 'iAry1').String())",
          ErrContext: "Input parameter 'iAry1' is INVALID!\n" +
            "'iAry1' FAILED Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateiAry2 {

    err = iaElectron.isValidIntAry(iAry2, ePrefix.XCpy("Validating 'iAry1'").String())

    if err != nil {

      return false,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = iaElectron.isValidIntAry(iAry2, ePrefix.XCpy(Validating 'iAry1').String())",
          ErrContext: "Input parameter 'iAry2' is INVALID!\n" +
            "'iAry2' FAILED Validation Tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  iaNanobot := new(intAryNanobot)

  err = iaNanobot.setInternalFlags(iAry1, ePrefix.XCpy("Setting flags 'iAry1'"))

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = iaNanobot.setInternalFlags(iAry1, ePrefix.XCpy(Setting flags 'iAry1'))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = iaNanobot.setInternalFlags(iAry2, ePrefix.XCpy("Setting flags 'iAry2'"))

  if err != nil {

    return false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = iaNanobot.setInternalFlags(iAry2," +
          "ePrefix.XCpy(Setting flags 'iAry2'))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return new(intAryBoson).dataFieldEqualityTest(iAry1, iAry2), nil
}

// getRawNumStr
//
//	Returns the current value of the input parameter IntAry object
//	('intAry') as a number string.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. These digits
//	are formatted in a way that facilitates conversion to a
//	corresponding numeric value.
//
//	The number string returned by this method will contain a
//	decimal separator to separate integer and fractional
//	components of the numeric value. The returned number string
//	will not contain 'thousands' separators or 'currency' symbols.
//
//	Validation Testing
//	==================
//
//	If input parameter 'validateIntAry' is set to true, this
//	method will subject 'intAry' to validation tests.
func (iaAtom *intAryAtom) getRawNumStr(
  intAry *IntAry,
  validateIntAry bool,
  errPrefDto *ePref.ErrPrefixDto) (string, error) {

  if iaAtom.lock == nil {
    iaAtom.lock = new(sync.Mutex)
  }

  iaAtom.lock.Lock()

  defer iaAtom.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryAtom.getRawNumStr()",
    "")

  if err != nil {
    return "", err
  }

  if intAry == nil {

    return "",
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'intAry'",
      }
  }

  if validateIntAry {

    err = new(intAryElectron).isValidIntAry(intAry, ePrefix.XCpy("Validating 'intAry'").String())

    if err != nil {

      return "",
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
            "  intAry, ePrefix.XCpy(Validating 'intAry').String())",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  // We now know that 'intAry' has valid NumSeps

  err = new(intAryNanobot).setInternalFlags(intAry, ePrefix.XCpy("Set intAry Flags"))

  if err != nil {
    return "", err
  }

  var buffer bytes.Buffer

  if intAry.signVal < 0 {
    buffer.WriteRune('-')
  }

  intLen := intAry.intAryLen - intAry.precision

  for i := 0; i < intLen; i++ {
    buffer.WriteRune(rune(intAry.intAry[i] + 48))
  }

  if intAry.precision > 0 {
    buffer.WriteRune(intAry.decimalSeparator)

    for j := 0; j < intAry.precision; j++ {
      buffer.WriteRune(rune(intAry.intAry[intLen] + 48))
      intLen++
    }

  }

  return buffer.String(), nil
}

// OptimizeIntArrayLen
//
//	 Eliminates Leading zeros from the front or integer portion
//	 of the integer string.
//
//	 If parameter 'optimizeFracDigits' is set equal to 'true',
//	 trailing zeros to the right of the decimal place will also be
//	 eliminated.
//
//		Validation Testing
//		==================
//
//		If input parameter 'validateIntAry' is set to true, this
//		method will subject 'intAry' to validation tests.
func (iaAtom *intAryAtom) optimizeIntArrayLen(
  intAry *IntAry,
  validateIntAry bool,
  optimizeFracDigits bool,
  validateResult bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaAtom.lock == nil {
    iaAtom.lock = new(sync.Mutex)
  }

  iaAtom.lock.Lock()

  defer iaAtom.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryAtom.optimizeIntArrayLen()",
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

  err = new(intAryUtility).selectIntAryValidation(
    intAry,
    "intAry",
    validateIntAry,
    ePrefix.XCpy("Validating 'intAry' on Startup"))

  if err != nil {
    return err
  }

  if intAry.isZeroValue {
    return nil
  }

  integerLen :=
    intAry.intAryLen - intAry.precision - intAry.firstDigitIdx

  if optimizeFracDigits {

    intAry.intAry = intAry.intAry[intAry.firstDigitIdx : intAry.lastDigitIdx+1]
    intAry.intAryLen = intAry.lastDigitIdx - intAry.firstDigitIdx + 1

  } else {

    intAry.intAry = intAry.intAry[intAry.firstDigitIdx:]
    intAry.intAryLen = intAry.intAryLen - intAry.firstDigitIdx
  }

  intAry.precision = intAry.intAryLen - integerLen

  err = new(intAryUtility).selectIntAryValidation(
    intAry,
    "intAry",
    validateResult,
    ePrefix.XCpy("Final Result Validation"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
        "intAry, \"intAry\", validateResult='%v' ePrefix", validateResult),
      ErrContext: "Error: The Final Result is INVALID!\n" +
        "Final Result 'ia' FAILED Validation Tests",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// setEqualArrayLengths
//
//	Compares an intAry object to the current intAry and ensures
//	that the lengths of both IntArrays are equal.
func (iaAtom *intAryAtom) setEqualArrayLengths(
  ia *IntAry,
  validateIa bool,
  iAry2 *IntAry,
  validateIAry2 bool,
  validateResult bool,
  errPrefDto *ePref.ErrPrefixDto) error {

  if iaAtom.lock == nil {
    iaAtom.lock = new(sync.Mutex)
  }

  iaAtom.lock.Lock()

  defer iaAtom.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryAtom.setEqualArrayLengths",
    "")

  if err != nil {
    return err
  }

  if ia == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'ia'",
    }
  }

  if iAry2 == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'iAry2'",
    }
  }

  // This will always set interal flags
  err = new(intAryUtility).selectIntAryValidation(
    ia,
    "ia",
    validateIa,
    ePrefix.XCpy("Validating 'ia' on Startup"))

  if err != nil {
    return err
  }

  // This will always set interal flags
  err = new(intAryUtility).selectIntAryValidation(
    iAry2,
    "iAry2",
    validateIAry2,
    ePrefix.XCpy("Validating 'iAry2' on Startup"))

  if err != nil {
    return err
  }

  iaIntLen := ia.intAryLen - ia.precision

  iAry2IntLen := iAry2.intAryLen - iAry2.precision

  if iaIntLen > iAry2IntLen {

    //iAry2.AddArrayLengthLeft(iaIntLen - iAry2IntLen)

    err = new(intAryNeutron).addArrayLengthLeft(iAry2, false, iaIntLen-iAry2IntLen, false, ePrefix)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNeutron).addArrayLengthLeft(\n" +
          "iAry2, validateIa=false, iaIntLen - iAry2IntLen, validateResult=false, ePrefix )",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  if iAry2IntLen > iaIntLen {
    //ia.AddArrayLengthLeft(iAry2IntLen - iaIntLen)

    err = new(intAryNeutron).addArrayLengthLeft(ia, false, iAry2IntLen-iaIntLen, false, ePrefix)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNeutron).addArrayLengthLeft(\n" +
          "ia, validateIa=false, iAry2IntLen - iaIntLen, validateResult=false, ePrefix )",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }
  }

  if ia.precision > iAry2.precision {

    //iAry2.AddArrayLengthRight(ia.precision - iAry2.precision)

    err = new(intAryNeutron).addArrayLengthRight(ia, false, ia.precision-iAry2.precision, false, ePrefix)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNeutron).addArrayLengthRight(\n" +
          "ia, validateIa=false, ia.precision - iAry2.precision, validateResult=false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    iAry2.precision = ia.precision
  }

  if iAry2.precision > ia.precision {

    //ia.AddArrayLengthRight(iAry2.precision - ia.precision)

    err = new(intAryNeutron).addArrayLengthRight(ia, false, iAry2.precision-ia.precision, false, ePrefix)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = new(intAryNeutron).addArrayLengthRight(\n" +
          "ia, validateIa=false, iAry2.precision - ia.precision, validateResult=false, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    ia.precision = iAry2.precision
  }

  // This will always set Internal flags
  err = new(intAryUtility).selectIntAryValidation(
    ia,
    "ia",
    validateResult,
    ePrefix.XCpy("Validating Final Calculation Result on 'ia'"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
        "ia, \"ia\", validateResult='%v' ePrefix", validateResult),
      ErrContext: "Error: The 'ia' Final Result is INVALID!\n" +
        "Final Result 'ia' FAILED Validation Tests",
      ErrMessage: err.Error(),
    }
  }

  // This will always set Internal flags
  err = new(intAryUtility).selectIntAryValidation(
    iAry2,
    "iAry2",
    validateResult,
    ePrefix.XCpy("Validating Final Calculation Result on 'iAry2'"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: fmt.Sprintf("err = new(intAryUtility).selectIntAryValidation(\n"+
        "iAry2, \"iAry2\", validateResult='%v' ePrefix", validateResult),
      ErrContext: "Error: The 'iAry2' Final Result is INVALID!\n" +
        "Final Result 'iAry2' FAILED Validation Tests",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
