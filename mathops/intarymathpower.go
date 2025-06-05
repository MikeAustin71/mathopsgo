package mathops

import (
  "errors"
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
)

type IntAryMathPower struct {
  Input  IntAryPair
  Result IntAry
}

// MinimumRequiredPrecision - designed to be used with the power function
// below. This method will compute the minimum number of decimal places
// required to support the result of raising a 'base' value to a specified
// 'exponent'. Both the 'base' and the 'exponent' are passed to this function
// as type *IntAry.
//
// For example, raising the value 3.12 to the power of 4 means that the
// result will require at least 8-decimal places to the right of the
// decimal in order to display a correct result. In the following example
// with base ='3.12' and exponent = '4', this method will return '8'.
//
//	Example: 3.12^4 = 94.75854336 (8-digits to the right of the decimal)
//
// The calculated minimum required precision is returned as a positive
// value of type 'int'.
//
// If the minimum required precision exceeds the maximum positive value for
// IntAry precision ( +2,147,483,646, which equals 2^31 − 2), an error
// message is returned in addition to the maximum positive value for IntAry
// precision (+2,147,483,646).
func (iaPwr *IntAryMathPower) MinimumRequiredPrecision(
  base, exponent *IntAry) (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAryMathPower.MinimumRequiredPrecision()",
    "")

  if err != nil {
    return 0, err
  }

  //maxValue := 2147483646

  signVal := 1

  basePrecisionUint, err := base.GetPrecisionUint()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "basePrecisionUint, err := base.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  basePrecision, err := new(IntAry).NewUint(basePrecisionUint, signVal, 0)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "basePrecision, err := new(IntAry).NewUint(basePrecisionUint, signVal, 0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  tExponent, err := exponent.CopyOut()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "tExponent, err := exponent.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  exponentSignVal, err := tExponent.GetSign()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "exponentSignVal, err := tExponent.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if exponentSignVal == -1 {

    err = tExponent.ChangeSign()

    if err != nil {

      return 0,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = tExponent.ChangeSign()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  iaResult := IntAry{}

  maxPrecision := tExponent.GetPrecision() + 5

  err = new(IntAryMathMultiply).Multiply(&basePrecision, &tExponent, &iaResult, 0, maxPrecision)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = new(IntAryMathMultiply).Multiply(&basePrecision, &tExponent, &iaResult, 0, maxPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  iaResultPrecisionUint, err := iaResult.GetPrecisionUint()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaResultPrecisionUint, err := iaResult.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if iaResultPrecisionUint > 0 {

    err = iaResult.RoundToPrecision(0)

    if err != nil {

      return 0,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = iaResult.RoundToPrecision(0)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

  }

  intVal, err := iaResult.GetInt()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "intVal, err := iaResult.GetInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return intVal, nil
}

// Pwr
//
//	Raises input parameter 'base' to the power of 'exponent'. This method
//	uses the 'Power By Twos' technique.
//	See:
//	https://en.wikipedia.org/wiki/Exponentiation_by_squaring
//	https://en.wikipedia.org/wiki/Exponentiation_by_squaring#Computation_by_powers_of_2
//
//	This method is based on revised code taken in part from Ye Lin
//	Aung.
//	https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
//
//	This algorithm was modified by Mike Rapp to achieve improved
//	performance.
//
//	The result of raising 'base' to the power of 'exponent' will
//	return the result in 'base'. As such the original value of
//	'base' will be overwritten.
//
//	The returned value 'base' will contain the same numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) as that of the original 'base' instance.'base' numeric
//	separators will therefore remain unchanged.
//
//	Example
//	=======
//
//	The 'power' calculation is computed as follows:
//
//	        base = base^exponent
//
//	maxResultPrecision
//	==================
//
//	Input parameter 'maxResultPrecision' will round the result to
//	this number of decimal places after the decimal point if the
//	result is greater than 'maxResultPrecision'.
//
//	If the value of 'maxResultPrecision' is less than zero, it will
//	be automatically set to a value of '4096'.
//
//	minResultPrecision
//	==================
//
//	Input parameter 'minResultPrecision' signals that if the result
//	precision is less than 'minResultPrecision', zeros will be added
//	to the right of the decimal place in order to implement the
//	'minResultPrecision' specification.
//
//	If the value of 'minResultPrecision' is less than zero,
//
// 'minResultPrecision' will be automatically set to a value of zero.
func (iaPwr *IntAryMathPower) Pwr(
  base *IntAry,
  exponent *IntAry,
  minResultPrecision,
  maxResultPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "IntAryMathPower.Pwr",
    "")

  if err != nil {
    return err
  }

  err = base.IsValid(ePrefix.XCpy(" base IntAry Error").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = base.IsValid(ePrefix.XCpy( base IntAry Error).String())",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = exponent.IsValid(ePrefix.XCpy(" exponent IntAry Error").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = exponent.IsValid(ePrefix.XCpy(exponent IntAry Error).String())",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  baseIsZero, err := base.IsZero()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "baseIsZero, err :=  base.IsZero()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if baseIsZero {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "base == Zero",
      ErrContext: "",
      ErrMessage: "Error: Input parameter 'base' has a zero value!",
    }
  }

  exponentIsZero, err := exponent.IsZero()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "exponentIsZero, err := exponent.IsZero()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if exponentIsZero {

    err = base.SetIntAryToOne(minResultPrecision)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = base.SetIntAryToOne(minResultPrecision)",
        ErrContext: fmt.Sprintf("minResultPrecision = '%v'", minResultPrecision),
        ErrMessage: err.Error(),
      }
    }

    return nil
  }

  iaOne, err := new(IntAry).NewOne(exponent.GetPrecision())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "iaOne, err := new(IntAry).NewOne(exponent.GetPrecision())",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  exponentIsEqualIaOne, err := exponent.Equal(&iaOne)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "exponentIsEqualIaOne, err := exponent.Equal(&iaOne)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if exponentIsEqualIaOne {
    return nil
  }

  exponentPrecision := exponent.GetPrecision()

  exponentSign, err := exponent.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "exponentSign, err := exponent.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if exponentPrecision == 0 && exponentSign == 1 {
    return iaPwr.pwrTwoPositiveIntegerExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  if exponentPrecision == 0 && exponentSign == -1 {
    return iaPwr.pwrTwoNegativeIntegerExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  if exponentPrecision > 0 && exponentSign == 1 {
    return iaPwr.pwrTwoPositiveFractionalExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  if exponentPrecision > 0 && exponentSign == -1 {
    return iaPwr.pwrTwoNegativeFractionalExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  return &FuncReturnError{
    ErrPrefix:  ePrefix.String(),
    ReturnFunc: ")",
    ErrContext: "",
    ErrMessage: "Error: Input parameters failed to match valid calculation types!",
  }
}

// PwrByMultiplication - raises base to the power of exponent using
// repetitive multiplication. This method may be slower than the method
// IntAryMathPower.Pwr(); however, this method is capable of handling
// very large exponents. Effectively, input parameter 'base' is raised
// to the power of exponent.
//
//	result = base^exponent
//
// The result of this operation is returned as pointer to an IntAry
// instance. This returned IntAry instance will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from input parameter 'base'.
//
// Input parameter 'maxResultPrecision' will round the result to this
// number of decimal places after the decimal point if the result is
// greater than 'maxResultPrecision'.  If the value of 'maxResultPrecision'
// is less than zero, it will be automatically set to a value of '4096'.
//
// Input parameter 'minResultPrecision' signals that if the result precision
// is less than 'minResultPrecision', zeros will be added to the right of
// the decimal place in order to implement the 'minResultPrecision'
// specification. If the value of 'minResultPrecision' is less than zero,
// 'minResultPrecision' will be automatically set to a value of zero.
func (iaPwr *IntAryMathPower) PwrByMultiplication(
  base, exponent *IntAry,
  minResultPrecision, maxResultPrecision int) (*IntAry, error) {

  ePrefix := "IntAryMathPower.PwrByMultiplication() "
  iaReturn := IntAry{}.NewZero(0)

  iaReturn.SetNumericSeparatorsDto(base.GetNumericSeparatorsDto())

  err := base.IsValid(ePrefix + "Invalid 'base' - ")

  if err != nil {
    return &iaReturn, err
  }

  err = exponent.IsValid(ePrefix + "Invalid 'exponent' - ")

  if err != nil {
    return &iaReturn, err
  }

  if base.IsZero() {
    return &iaReturn,
      errors.New(ePrefix + "'base' is zero value. INVALID INPUT!")
  }

  if exponent.IsZero() {
    iaReturn.SetIntAryToOne(minResultPrecision)
    return &iaReturn, nil
  }

  iaOne := IntAry{}.NewOne(exponent.GetPrecision())

  if exponent.Equals(&iaOne) {
    iaReturn = base.CopyOut()
    return &iaReturn, nil
  }

  exponentPrecision := exponent.GetPrecision()
  exponentSign := exponent.GetSign()

  if exponentPrecision == 0 && exponentSign == 1 {
    return iaPwr.pwrMultiplyPositiveIntegerExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  if exponentPrecision == 0 && exponentSign == -1 {
    return iaPwr.pwrMultiplyNegativeIntegerExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  if exponentPrecision > 0 && exponentSign == 1 {
    return iaPwr.pwrMultiplyPositiveFractionalExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  if exponentPrecision > 0 && exponentSign == -1 {
    return iaPwr.pwrMultiplyNegativeFractionalExponent(
      base,
      exponent,
      minResultPrecision,
      maxResultPrecision)
  }

  return &iaReturn,
    errors.New(ePrefix + "Error: input parameters failed to match valid calculation types!")
}

// pwrMultiplyPositiveIntegerExponent - raises 'base' to the power of 'exponent'.
// Input parameter 'exponent' is expected to represent a positive integer.
// If 'exponent' is NOT a positive integer, an error will be thrown.
//
// If 'exponent' is a fractional value, i.e., it has digits to the right of the
// decimal place, it is by definition NOT an integer value and an error will
// be thrown.
//
// This method uses simple multiplication to generate the result.
//
// Input parameter 'maxResultPrecision' will round the result to this
// number of decimal places after the decimal point if the result is
// greater than 'maxResultPrecision'.  If the value of 'maxResultPrecision'
// is less than zero, it will be automatically set to a value of '4096'.
//
// Input parameter 'minResultPrecision' signals that if the result precision
// is less than 'minResultPrecision', zeros will be added to the right of
// the decimal place in order to implement the 'minResultPrecision'
// specification. If the value of 'minResultPrecision' is less than zero,
// 'minResultPrecision' will be automatically set to a value of zero.
//
// The result of the power operation is returned as a pointer to a new
// 'result' IntAry. None of the input parameters are altered by this
// operation. The returned 'result' IntAry will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from
// input parameter 'base'.
//
// Note: This method does not perform tests for base==0, exponent==0 or exponent==1.
// It is assumed that these tests were performed before calling this method.
func (iaPwr *IntAryMathPower) pwrMultiplyPositiveIntegerExponent(
  base, exponent *IntAry,
  minResultPrecision, maxResultPrecision int) (*IntAry, error) {

  ePrefix := "IntAryMathPower.pwrMultiplyPositiveIntegerExponent() "
  iaErrReturn := IntAry{}.NewZero(0)

  if exponent.GetSign() != 1 {
    return &iaErrReturn,
      fmt.Errorf(ePrefix+
        "Error: 'exponent' is expected to be a positive integer. "+
        "Instead, 'exponent' is negative! exponent='%v'",
        exponent.GetNumStr())
  }

  if exponent.GetPrecision() != 0 {
    return &iaErrReturn,
      fmt.Errorf(ePrefix+
        "Error: 'exponent' is expected to be an integer value. "+
        "Instead, 'exponent' is a fractional value! exponent='%v'",
        exponent.GetNumStr())
  }

  numSeps := base.GetNumericSeparatorsDto()

  if maxResultPrecision < 0 {
    maxResultPrecision = 4096
  }

  if minResultPrecision < 0 {
    minResultPrecision = 0
  }

  if minResultPrecision > maxResultPrecision {
    minResultPrecision = maxResultPrecision
  }

  result := IntAry{}.NewOne(0)

  internalMaxPrecision := maxResultPrecision + 100

  opExponent := exponent.CopyOut()

  for !opExponent.IsZero() {

    IntAryMathMultiply{}.MultiplyInPlace(&result, base, minResultPrecision, internalMaxPrecision)

    err := opExponent.DecrementIntegerOne()

    if err != nil {
      return &iaErrReturn,
        fmt.Errorf(ePrefix+
          "Error returned by opExponent.DecrementIntegerOne() "+
          "Error='%v' ", err.Error())
    }

  }

  if result.GetPrecision() > maxResultPrecision {
    result.RoundToPrecision(maxResultPrecision)
  }

  if result.GetPrecision() < minResultPrecision {
    result.SetPrecision(minResultPrecision, false)
  }

  err := result.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return &iaErrReturn,
      fmt.Errorf(ePrefix+
        "Error returned by result.SetNumericSeparatorsDto(numSeps) "+
        "Error='%v' ", err.Error())
  }

  return &result, nil
}
