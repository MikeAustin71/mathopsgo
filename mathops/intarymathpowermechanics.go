package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "sync"
)

type intAryMathPwrMechanics struct {
  lock sync.Mutex
}

// pwrTwoNegativeFractionalExponent
//
//  Raises 'base' to the power of 'exponent'.
//
//  Input parameter 'exponent' is expected to represent a negative
//  fractional value, i.e., a value less than zero with digits to
//  the right of the decimal point.
//
//  If 'exponent' is a positive value, i.e., a greater than -1, an
//  error will be returned.
//
//  Also, if 'exponent' has a precision value less than '1', an
//  error will be returned.
//
//  Input parameter 'maxResultPrecision' will round the result to
//  this number of decimal places after the decimal point if the
//  result is greater than 'maxResultPrecision'.  If the value of
//  'maxResultPrecision' is less than zero, it will be
//  automatically set to a value of '4096'.
//
//  Input parameter 'minResultPrecision' signals that if the result
//  precision is less than 'minResultPrecision', zeros will be
//  added to the right of the decimal place in order to implement
//  the 'minResultPrecision' specification. If the value of
//  'minResultPrecision' is less than zero, 'minResultPrecision'
//  will be automatically set to a value of zero.
//
//  The result of the power operation is returned in the input
//  parameter 'base'. During this procedure the original value of
//  'base' is destroyed.
//
//  The returned IntAry object 'base' will contain same numeric
//  separators (i.e. decimal separator, thousands separator and
//  currency symbol) as those in the original 'base' instance. As
//  such, the 'base' numeric separators will remain unchanged.
func (iaMathMech *intAryMathPwrMechanics) pwrTwoNegativeFractionalExponent(
  base *IntAry,
  exponent *IntAry,
  minResultPrecision,
  maxResultPrecision int,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaMathMech.lock.Lock()

  defer iaMathMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMathPwrMechanics.pwrTwoNegativeFractionalExponent",
    "")

  if err != nil {
    return err
  }

  if base == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'base'",
    }
  }

  if exponent == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'exponent'",
    }
  }

  exponentGetSignVal, err := exponent.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "exponentGetSignVal, err := exponent.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  exponentNumStr, err := exponent.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
      ErrMessage: err.Error(),
    }
  }

  if exponentGetSignVal != -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "if exponent Sign Value != 1",
      ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a positive fractional value. "+
        "Instead, 'exponent' is negative! exponent='%v'", exponentNumStr),
    }
  }

  exponentPrecisionVal := exponent.GetPrecision()

  if exponentPrecisionVal < 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "if exponent precision Value < 1",
      ErrMessage: fmt.Sprintf("Error: 'exponent' precision value is expected to be greater than zero.\n"+
        "precision= '%v'", exponentPrecisionVal),
    }
  }

  numSeps, err := base.GetNumericSeparatorsDto()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "numSeps, err := base.GetNumericSeparatorsDto()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if maxResultPrecision < 0 {
    maxResultPrecision = 4096
  }

  if minResultPrecision < 0 {
    minResultPrecision = 0
  }

  if minResultPrecision > maxResultPrecision {
    minResultPrecision = maxResultPrecision
  }

  internalMaxPrecision := maxResultPrecision + 100

  // Set exponent to a positive value
  newExponent, err := exponent.CopyOut()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newExponent, err := exponent.CopyOut()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = newExponent.ChangeSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = newExponent.ChangeSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  newInverseBase, err := base.Inverse(internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newInverseBase, err := base.Inverse(internalMaxPrecision)",
      ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'",
        internalMaxPrecision),
      ErrMessage: err.Error(),
    }
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(&newExponent)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "fracIntAry, err := new(FracIntAry).NewFracIntAry(&newExponent)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  internalMaxPrecision += 5

  err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)",
      ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'",
        internalMaxPrecision),
      ErrMessage: err.Error(),
    }
  }

  internalMaxPrecision += 5

  newBase, err := NthRootOp{}.NewNthRoot(&newInverseBase, &fracIntAry.Denominator, internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "newBase, err := NthRootOp{}.NewNthRoot(\n" +
        "&newInverseBase, &fracIntAry.Denominator, internalMaxPrecision)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  expBigInt, err := fracIntAry.Numerator.GetBigInt()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "expBigInt, err := fracIntAry.Numerator.GetBigInt()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  internalMaxPrecision += 5

  err = new(intAryMathPwrNanobot).pwrByTwos(&newBase, expBigInt, maxResultPrecision, internalMaxPrecision, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryMathPwrNanobot).pwrByTwos(\n" +
        "  &newBase, expBigInt, maxResultPrecision, internalMaxPrecision,\n" +
        "  ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = base.CopyIn(&newBase, false)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = base.CopyIn(&newBase, false)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if base.GetPrecision() > maxResultPrecision {

    err = base.RoundToPrecision(maxResultPrecision)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = base.RoundToPrecision(maxResultPrecision)",
        ErrContext: fmt.Sprintf("maxResultPrecision= '%v'",
          maxResultPrecision),
        ErrMessage: err.Error(),
      }
    }

  }

  if base.GetPrecision() < minResultPrecision {

    err = base.SetPrecision(minResultPrecision, false)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = base.SetPrecision(minResultPrecision, false)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

  }

  err = base.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = base.SetNumericSeparatorsDto(numSeps)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// pwrTwoPositiveFractionalExponent
//
//	Raises 'base' to the power of 'exponent'.
//
//	   Result = base^exponent
//
//	exponent
//	========
//
//	Input parameter 'exponent' is expected to represent a positive
//	fractional value, i.e., a value greater than -1 with digits to
//	the right of the decimal point.
//
//	If 'exponent' is a negative value, i.e., a value less than
//	zero, an error will be returned.
//
//	Also, if 'exponent' has a precision value less than '1', an
//	error will be returned.
//
//	maxResultPrecision
//	==================
//
//	Input parameter 'maxResultPrecision' will round the result to
//	this number of decimal places after the decimal point if the
//	result is greater than 'maxResultPrecision'.  If the value of
//	'maxResultPrecision' is less than zero, it will be
//	automatically set to a value of '4096'.
//
//	minResultPrecision
//	==================
//
//	Input parameter 'minResultPrecision' signals that if the result
//	precision is less than 'minResultPrecision', zeros will be
//	added to the right of the decimal place in order to implement
//	the 'minResultPrecision' specification. If the value of
//	'minResultPrecision' is less than zero, 'minResultPrecision'
//	will be automatically set to a value of zero.
//
//	The result of the power operation is returned in the input
//	parameter 'base'. During this procedure the original value
//	of 'base' is destroyed.
//
//	Numeric Separators
//	==================
//
//	The returned IntAry object 'base' will contain same numeric
//	separators ( i.e. decimal separator, thousands separator and
//	currency symbol) as those in the original 'base' instance. As
//	such, the 'base' numeric separators will remain unchanged.
func (iaMathMech *intAryMathPwrMechanics) pwrTwoPositiveFractionalExponent(
  base *IntAry,
  exponent *IntAry,
  minResultPrecision,
  maxResultPrecision int,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaMathMech.lock.Lock()

  defer iaMathMech.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMathPwrMechanics.pwrTwoPositiveFractionalExponent",
    "")

  if err != nil {
    return err
  }

  if base == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'base'",
    }
  }

  if exponent == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'exponent'",
    }
  }

  exponentSignVal, err := exponent.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "exponentSignVal, err := exponent.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  exponentNumStr, err := exponent.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
      ErrMessage: err.Error(),
    }
  }

  if exponentSignVal != 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "if exponent Sign Value != 1",
      ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a positive fractional value. "+
        "Instead, 'exponent' is negative! exponent='%v'", exponentNumStr),
    }
  }

  exponentPrecisionVal := exponent.GetPrecision()

  if exponentPrecisionVal < 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "if exponent precision Value < 1",
      ErrMessage: fmt.Sprintf("Error: 'exponent' precision value is expected to be greater than zero.\n"+
        "precision= '%v'", exponentPrecisionVal),
    }
  }

  numSeps, err := base.GetNumericSeparatorsDto()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "numSeps, err := ia.GetNumericSeparatorsDto()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if maxResultPrecision < 0 {
    maxResultPrecision = 4096
  }

  if minResultPrecision < 0 {
    minResultPrecision = 0
  }

  if minResultPrecision > maxResultPrecision {
    minResultPrecision = maxResultPrecision
  }

  internalMaxPrecision := maxResultPrecision + 100

  // returns exponent / 1
  fracIntAry, err := new(FracIntAry).NewFracIntAry(exponent)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "fracIntAry, err := new(FracIntAry).NewFracIntAry(exponent)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  internalMaxPrecision += 5

  newBase, err := NthRootOp{}.NewNthRoot(base, &fracIntAry.Denominator, internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newBase, err := NthRootOp{}.NewNthRoot(...)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = base.CopyIn(&newBase, false)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = base.CopyIn(&newBase, false)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  expBigInt, err := fracIntAry.Numerator.GetBigInt()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "expBigInt, err := fracIntAry.Numerator.GetBigInt()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  internalMaxPrecision += 5

  err = new(intAryMathPwrNanobot).
    pwrByTwos(base, expBigInt, maxResultPrecision, internalMaxPrecision, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(intAryMathPwrNanobot).pwrByTwos(\n" +
        "  base, expBigInt, maxResultPrecision, internalMaxPrecision, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if base.GetPrecision() > maxResultPrecision {

    err = base.RoundToPrecision(maxResultPrecision)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = base.RoundToPrecision(maxResultPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

  }

  if base.GetPrecision() < minResultPrecision {

    err = base.SetPrecision(minResultPrecision, false)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = base.SetPrecision(minResultPrecision, false)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

  }

  err = base.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = base.SetNumericSeparatorsDto(numSeps)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}
