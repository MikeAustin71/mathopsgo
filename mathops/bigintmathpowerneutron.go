package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type bigIntMathPowerNeutron struct {
  lock sync.Mutex
}

// bigIntNumMinRequiredPrecision
//
//  Designed to be used with the power function (Pwr). This method
//  will compute the minimum number of decimal places required to
//  support the result of raising a 'base' value to a specified
//  exponent. Both the 'base' and the 'exponent' are passed to this
//  function as type BigIntNum.
//
//  For example, raising the value 3.12 to the power of 4 means that
//  the result will require at least 8-decimal places to the right of
//  the decimal in order to display a correct result. In the following
//  example with base ='3.12' and exponent = '4', this method will
//  return '8'.
//
//  Example
//  =======
//
//  3.12^4 = 94.75854336 (2x4 = 8-digits to the right of the decimal)
//         Minimum Required Precision = precision x exponent
//
//  The calculated minimum required precision is returned as type
//  'uint'.
//
//  If the minimum required precision exceeds the maximum value for
//  type 'uint' (+4,294,967,295, which equals 2^32 − 1), an error
//  message is returned in addition to the maximum uint value
//  (+4,294,967,295).
func (bIMathPwrNeutron *bigIntMathPowerNeutron) bigIntNumMinRequiredPrecision(
  base *BigIntNum,
  validateBase bool,
  exponent *BigIntNum,
  validateExponent bool,
  errPrefDto *ePref.ErrPrefixDto) (uint, error) {

  bIMathPwrNeutron.lock.Lock()

  defer bIMathPwrNeutron.lock.Unlock()

  var err error

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathPowerNeutron.bigIntNumMinRequiredPrecision",
    "")

  if err != nil {
    return 0, err
  }

  if base == nil {

    return 0,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'base'",
      }
  }

  if exponent == nil {

    return 0,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'exponent'",
      }
  }

  if validateBase {

    err = base.IsValid(ePrefix.XCpy("Validating 'base'").String())

    if err != nil {

      return 0,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = base.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'base'\").String()",
          ErrContext: "Error: Input parameter 'base' is INVALID!\n" +
            "'base' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateExponent {

    err = exponent.IsValid(ePrefix.XCpy("Validating 'exponent'").String())

    if err != nil {

      return 0,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = exponent.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'exponent'\").String()",
          ErrContext: "Error: Input parameter 'exponent' is INVALID!\n" +
            "'exponent' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

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

  basePrecision, err := new(BigIntNum).NewUint(basePrecisionUint, 0)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "basePrecision, err := new(BigIntNum).\n" +
          "  NewUint(basePrecisionUint, 0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  basePrecisionNumStr, err := basePrecision.GetNumStr()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "basePrecisionNumStr, err := basePrecision.GetNumStr()",
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

  tExponentNumStr, err := tExponent.GetNumStr()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "tExponentNumStr, err := tExponent.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  tExponentSignValue, err := tExponent.GetSign()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "tExponentSignValue, err := tExponent.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if tExponentSignValue == -1 {

    err = tExponent.ChangeSign()

    if err != nil {

      return 0,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = tExponent.ChangeSign()",
          ErrContext: "tExponentSignValue == -1",
          ErrMessage: err.Error(),
        }
    }
  }

  minRequiredPrecision, err := new(BigIntMathMultiply).
    MultiplyBigIntNums(basePrecision, tExponent)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "minRequiredPrecision, err := new(BigIntMathMultiply).\n" +
          "  MultiplyBigIntNums(basePrecision, tExponent)",
        ErrContext: fmt.Sprintf("basePrecision= '%v'\n"+
          "tExponent= '%v'", basePrecisionNumStr, tExponentNumStr),
        ErrMessage: err.Error(),
      }
  }

  err = minRequiredPrecision.RoundToDecPlace(0)

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = minRequiredPrecision.RoundToDecPlace(0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  minRequiredPrecisionUint, err := minRequiredPrecision.GetUInt()

  if err != nil {

    return 0,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "minRequiredPrecisionUint, err := minRequiredPrecision.GetUInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return minRequiredPrecisionUint, nil
}

// bigIntPwrIteration
//
//	Raises input parameter 'base' to the power of input parameter
//	'exponent'.
//
//	This method of raising a base to an exponent uses iterative
//	multiplication and manages the internal precision of each
//	iterative multiplication. If, during the process of multiplying
//	the base time itself, the internal precision exceeds the
//	'internalMaxPrecision' limit, that intermediate number is rounded
//	down to 'internalMaxPrecision'.
//
//	If the precision of the final result exceeds the limit imposed by
//	input parameter, 'outputMaxPrecision', that final result will be
//	rounded to 'outputMaxPrecision' digits to the right of the decimal
//	place.
//
//	Input Parameter
//	===============
//
//	base                     *big.Int
//	  The base which will be raised to the power of 'exponent'.
//	             baseToPwr = base^exponent
//
//	basePrecision            uint
//	  The number of digits to the right of the decimal place in the
//	  numeric sequence represented by 'base'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and
//	  'basePrecision' define a numeric value with a fixed number of
//	  decimal digits to the right of the decimal place.
//
//	exponent                 uint
//	  This function will raise 'base' to the power of 'exponent'.
//	              baseToPwr = base^exponent
//
//	internalMaxPrecision     uint
//	  This value is imposed as a limit on the precision of internal
//	  calculations necessary to compute the result of this power
//	  operation. If during the calculation an interim or intermediate
//	  result is generated which exceeds this limit, that intermediate
//	  result will be rounded to 'internalMaxPrecision'. The term
//	  precision defines the number of digits to the right of the
//	  decimal place.
//
//	  If 'internalMaxPrecision' is less than 'outputMaxPrecision',
//	  'internalMaxPrecision' will be automatically set to a value of
//	  'outputMaxPrecision' + 100.
//
//	outputMaxPrecision       uint
//	  This value is imposed as a limit on the precision of the final
//	  calculated result of the power operation. If the number of
//	  digits to the right of the decimal point in the final calculated
//	  result exceeds this limit, that final result will be rounded to
//	  'outputMaxPrecision' digits to the right of the decimal place.
//	  The term precision defines the number of digits to the right of
//	  the decimal place.
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
//	baseToPwr                *big.Int
//	  This function returns the result of 'base' raised to the power
//	  of 'exponent'. This result, 'baseToPwr' is returned as a type
//	  *big.Int.
//	                   baseToPwr = base^exponent
//
//	baseToPwrPrecision       uint
//	  Specifies the number of digits to the right of the decimal place
//	  in the numeric sequence represented by the calculation result,
//	  'baseToPwr'.
//
//	err                      error
//	  If the calculation encounters an error, an appropriate error
//	  message will be formatted and returned. If the calculation
//	  completes successfully, this return value will be set to 'nil'.
func (bIMathPwrNeutron *bigIntMathPowerNeutron) bigIntPwrIteration(
  base *big.Int,
  basePrecision uint,
  exponent uint,
  internalMaxPrecision uint,
  outputMaxPrecision uint,
  errPrefDto *ePref.ErrPrefixDto) (
  baseToPwr *big.Int, baseToPwrPrecision uint, err error) {

  bIMathPwrNeutron.lock.Lock()

  defer bIMathPwrNeutron.lock.Unlock()

  baseToPwr = big.NewInt(0)

  baseToPwrPrecision = 0

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathPowerNeutron.bigIntPwrIteration",
    "")

  if err != nil {
    return baseToPwr, baseToPwrPrecision, err
  }

  if base == nil {

    return baseToPwr, baseToPwrPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'base'",
      }
  }

  exponent++

  bigIZero := big.NewInt(0)

  if base.Cmp(bigIZero) == 0 {

    baseToPwr = big.NewInt(0)

    baseToPwrPrecision = 0

    return baseToPwr, baseToPwrPrecision, nil
  }

  if base.Cmp(big.NewInt(1)) == 0 {

    baseToPwr = big.NewInt(0).Set(base)

    baseToPwrPrecision = basePrecision

    return baseToPwr, baseToPwrPrecision, nil
  }

  if internalMaxPrecision < outputMaxPrecision {

    internalMaxPrecision = outputMaxPrecision + 100
  }

  bigIPlusFive := big.NewInt(5)

  bigIMinusFive := big.NewInt(-5)

  bigITen := big.NewInt(10)

  roundFactor := big.NewInt(0)

  cmpResult := 0

  for i := uint(0); i < exponent; i++ {

    if i == 0 {

      baseToPwr = big.NewInt(1)

      baseToPwrPrecision = 0

    } else if i == 1 {

      baseToPwr = big.NewInt(0).Set(base)

      baseToPwrPrecision = basePrecision

    } else {

      baseToPwr = big.NewInt(0).Mul(baseToPwr, base)

      baseToPwrPrecision = baseToPwrPrecision + basePrecision
    }

    if baseToPwrPrecision > internalMaxPrecision {

      // beforeRounding := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)
      delta := baseToPwrPrecision - internalMaxPrecision - 1

      scale := big.NewInt(0).Exp(bigITen, big.NewInt(int64(delta)), nil)

      cmpResult = baseToPwr.Cmp(bigIZero)

      if cmpResult == 1 {

        // baseToPwr is GREATER Than zero
        roundFactor = big.NewInt(0).Mul(bigIPlusFive, scale)

      } else if cmpResult == -1 {

        // baseToPwr is LESS Than zero
        roundFactor = big.NewInt(0).Mul(bigIMinusFive, scale)

      } else {

        baseToPwrPrecision = 0
        continue
      }

      baseToPwr = big.NewInt(0).Add(baseToPwr, roundFactor)

      scale = big.NewInt(0).Mul(scale, bigITen)

      baseToPwr = big.NewInt(0).Quo(baseToPwr, scale)

      baseToPwrPrecision = internalMaxPrecision
    }
  }

  if baseToPwrPrecision > outputMaxPrecision {

    delta := baseToPwrPrecision - outputMaxPrecision - 1

    scale := big.NewInt(0).Exp(bigITen, big.NewInt(int64(delta)), nil)

    round := big.NewInt(0).Mul(bigIPlusFive, scale)

    if baseToPwr.Cmp(bigIZero) == -1 {

      round = big.NewInt(0).Mul(round, big.NewInt(-1))

    }

    baseToPwr = big.NewInt(0).Add(baseToPwr, round)

    scale = big.NewInt(0).Mul(scale, bigITen)

    baseToPwr = big.NewInt(0).Quo(baseToPwr, scale)

    baseToPwrPrecision = outputMaxPrecision
  }

  return baseToPwr, baseToPwrPrecision, nil
}

// bigIntegerPwrIteration
//
//	Raises input parameter 'base' to the power of input parameter
//	'exponent'. This version of the function uses *big.Int type for
//	all input parameters.
//
//	This method of raising a base to an exponent uses iterative
//	multiplication and manages the internal precision of each
//	iterative multiplication. If, during the process of multiplying
//	the base times itself, the internal precision exceeds the
//	'internalMaxPrecision' limit, that intermediate number is rounded
//	to 'internalMaxPrecision'.
//
//	If the precision of the final result exceeds the limit imposed by
//	input parameter, 'outputMaxPrecision', that final result will be
//	rounded to 'outputMaxPrecision' digits to the right of the
//	decimal place.
//
//	!!! WARNING !!!
//	===============
//
//	Currently, this method will only accept positive numeric values
//	for input parameter 'exponent'.
//
//	Input Parameter
//	===============
//
//	base                     *big.Int
//	  The base which will be raised to the power of 'exponent'.
//	                  baseToPwr = base^exponent
//
//	basePrecision            *big.Int
//	  The number of digits to the right of the decimal place in the
//	  numeric sequence represented by 'base'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and
//	  'basePrecision' define a numeric value with a fixed number of
//	  decimal digits to the right of the decimal place.
//
//	exponent                 *big.Int
//	  'exponent' is an integer value with zero precision.
//
//	   This function will raise 'base' to the power of 'exponent'.
//									baseToPwr = base^exponent
//
//	   !!! WARNING !!!
//	     Currently, 'exponent' MUST BE a positive numeric value.
//
//	internalMaxPrecision     *big.Int
//	  This value is imposed as a limit on the precision of internal
//	  calculations necessary to compute the result of this power
//	  operation. If during the calculation an interim or intermediate
//	  result is generated which exceeds this limit, that intermediate
//	  result will be rounded to 'internalMaxPrecision'.
//
//	  The term 'precision' defines the number of digits to the right
//	  of the decimal place.
//
//	  If 'internalMaxPrecision' is less than 'outputMaxPrecision',
//	  'internalMaxPrecision' will be automatically set to a value of
//	  'outputMaxPrecision' + 100.
//
//	outputMaxPrecision       *big.Int
//	  This value is imposed as a limit on the precision of the final
//	  calculated result of the power operation. If the number of
//	  digits to the right of the decimal point in the final calculated
//	  result exceeds this limit, that final result will be rounded to
//	  'outputMaxPrecision' digits to the right of the decimal place.
//
//	  The term 'precision' defines the number of digits to the right
//	  of the decimal place. If 'outputMaxPrecision' is less than zero,
//	  an error will be returned.
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
//	baseToPwr                *big.Int
//	  This function returns the result of 'base' raised to the power
//	  of 'exponent'. This result, 'baseToPwr' is returned as a type
//	  *big.Int.
//	                    baseToPwr = base^exponent
//
//	baseToPwrPrecision       *big.Int
//	  Specifies the number of digits to the right of the decimal place
//	  in the numeric sequence represented by the calculation result,
//	  'baseToPwr'.
//
//	err                      error
//	  If the function fails to complete successfully, this value is
//	  configured with an appropriate error message and returned to the
//	  caller. If the function completes successfully, this value is
//	  set to 'nil'.
func (bIMathPwrNeutron *bigIntMathPowerNeutron) bigIntegerPwrIteration(
  base *big.Int,
  basePrecision *big.Int,
  exponent *big.Int,
  internalMaxPrecision *big.Int,
  outputMaxPrecision *big.Int,
  errPrefDto *ePref.ErrPrefixDto) (baseToPwr *big.Int, baseToPwrPrecision *big.Int, err error) {

  bIMathPwrNeutron.lock.Lock()

  defer bIMathPwrNeutron.lock.Unlock()

  baseToPwr = big.NewInt(0)

  baseToPwrPrecision = big.NewInt(0)

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathPowerNeutron.bigIntegerPwrIteration",
    "")

  if err != nil {
    return baseToPwr, baseToPwrPrecision, err
  }

  if base == nil {

    return baseToPwr, baseToPwrPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'base'",
      }
  }

  if basePrecision == nil {

    return baseToPwr, baseToPwrPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'basePrecision'",
      }
  }

  if exponent == nil {

    return baseToPwr, baseToPwrPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'exponent'",
      }
  }

  if internalMaxPrecision == nil {

    return baseToPwr, baseToPwrPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'internalMaxPrecision'",
      }
  }

  if outputMaxPrecision == nil {

    return baseToPwr, baseToPwrPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'outputMaxPrecision'",
      }
  }

  bigZero := big.NewInt(0)

  bigOne := big.NewInt(1)

  if exponent.Cmp(bigZero) == -1 {

    return baseToPwr, baseToPwrPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("exponent = '%v'", exponent.Text(10)),
        ErrMessage: "Error: Input parameter 'exponent' is out of range!\n" +
          "'exponent' is a negative number.",
      }
  }

  if basePrecision.Cmp(bigZero) == -1 {

    return baseToPwr, baseToPwrPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("basePrecision = '%v'", basePrecision.Text(10)),
        ErrMessage: "Error: Input parameter 'basePrecision' is out of range!\n" +
          "'basePrecision' is a negative number.",
      }
  }

  if outputMaxPrecision.Cmp(bigZero) == -1 {

    return baseToPwr, baseToPwrPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("outputMaxPrecision = '%v'", outputMaxPrecision.Text(10)),
        ErrMessage: "Error: Input parameter 'outputMaxPrecision' is out of range!\n" +
          "'outputMaxPrecision' is a negative number.",
      }
  }

  if base.Cmp(bigZero) == 0 {
    return baseToPwr, baseToPwrPrecision, err
  }

  if base.Cmp(bigOne) == 0 &&
    baseToPwrPrecision.Cmp(bigZero) == 0 {

    baseToPwr = big.NewInt(0).Set(bigOne)

    baseToPwrPrecision.Set(basePrecision)

    return baseToPwr, baseToPwrPrecision, err
  }

  if base.Cmp(big.NewInt(-1)) == 0 &&
    baseToPwrPrecision.Cmp(bigZero) == 0 {

    baseToPwr = big.NewInt(-1)

    baseToPwrPrecision.Set(basePrecision)

    return baseToPwr, baseToPwrPrecision, err
  }

  if internalMaxPrecision.Cmp(outputMaxPrecision) != 1 {

    internalMaxPrecision = big.NewInt(0).Add(outputMaxPrecision, big.NewInt(100))
  }

  cycles := big.NewInt(0).Add(exponent, bigOne)

  bigIPlusFive := big.NewInt(5)

  bigIMinusFive := big.NewInt(-5)

  bigITen := big.NewInt(10)

  roundFactor := big.NewInt(0)

  cmpResult := 0

  for i := big.NewInt(0); i.Cmp(cycles) == -1; i.Add(i, bigOne) {

    if i.Cmp(bigZero) == 0 {

      baseToPwr = big.NewInt(1)

      baseToPwrPrecision = big.NewInt(0)

    } else if i.Cmp(bigOne) == 0 {

      baseToPwr = big.NewInt(0).Set(base)

      baseToPwrPrecision = big.NewInt(0).Set(basePrecision)

    } else {

      baseToPwr = big.NewInt(0).Mul(baseToPwr, base)

      baseToPwrPrecision.Add(baseToPwrPrecision, basePrecision)
    }

    if baseToPwrPrecision.Cmp(internalMaxPrecision) == 1 {

      // beforeRounding := new(BigIntNum).NewBigInt(baseToPwr, baseToPwrPrecision)

      delta := big.NewInt(0).Sub(baseToPwrPrecision, internalMaxPrecision)

      delta.Sub(delta, bigOne)

      scale := big.NewInt(0).Exp(bigITen, delta, nil)

      cmpResult = baseToPwr.Cmp(bigZero)

      if cmpResult == 1 {

        // baseToPwr is GREATER Than zero
        roundFactor = big.NewInt(0).Mul(bigIPlusFive, scale)

      } else if cmpResult == -1 {
        // baseToPwr is LESS Than zero

        roundFactor = big.NewInt(0).Mul(bigIMinusFive, scale)

      } else {

        baseToPwrPrecision = big.NewInt(0)

        continue
      }

      baseToPwr = big.NewInt(0).Add(baseToPwr, roundFactor)

      scale = big.NewInt(0).Mul(scale, bigITen)

      baseToPwr = big.NewInt(0).Quo(baseToPwr, scale)

      baseToPwrPrecision = big.NewInt(0).Set(internalMaxPrecision)

    }
  }

  if baseToPwrPrecision.Cmp(outputMaxPrecision) == 1 {

    delta := big.NewInt(0).Sub(baseToPwrPrecision, outputMaxPrecision)

    delta.Sub(delta, bigOne)

    scale := big.NewInt(0).Exp(bigITen, delta, nil)

    roundFactor = big.NewInt(0).Mul(bigIPlusFive, scale)

    if baseToPwr.Cmp(bigZero) == -1 {
      roundFactor = big.NewInt(0).Mul(roundFactor, big.NewInt(-1))
    }

    baseToPwr = big.NewInt(0).Add(baseToPwr, roundFactor)

    scale = big.NewInt(0).Mul(scale, bigITen)

    baseToPwr = big.NewInt(0).Quo(baseToPwr, scale)

    baseToPwrPrecision = big.NewInt(0).Set(outputMaxPrecision)

  }

  return baseToPwr, baseToPwrPrecision, err
}

// bigIntToNegativeIntegerPower
//
//	Raises 'base' to the power of a negative integer exponent,
//	'exponent'.
//
//	As stated in the function name, this method expects to process
//	only negative, integer exponents.
//
//	Examples
//	========
//
//	base  basePrecision  exponent  exponentPrecision  result     resultPrecision
//
//	 5         0            -2            0             4              2
//	                          5^-2 = 0.04
//
//	1131       2            -3            0           00069121    8 (to 8-decimal places)
//	            11.31^-3= 0.00069121345785745610965099525880031
//
//		            The actual number of decimal places returned in
//	            the result is controlled by	input parameter,
//	            'maxPrecision'.
//
//	Input Parameters
//	================
//
//	base                     *big.Int
//	  The base which will be raised to the power of a negative integer
//	  exponent.
//
//	basePrecision            uint
//	  The precision specification for 'base'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and 'basePrecision'
//	  define a numeric value with a fixed number of decimal digits to
//	  the right of the decimal place.
//
//	exponent                 *big.Int
//	  The exponent to which 'base' will be raised by this calculation.
//	  By method definition, 'exponent' must be a negative value. If
//	  exponent is greater than -1, an error will be triggered.
//
//	exponentPrecision        uint
//	  The precision specification for 'exponent'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and 'basePrecision'
//	  define a numeric value with a fixed number of decimal digits to
//	  the right of the decimal place. For this method,
//	  'exponentPrecision' MUST BE SET TO ZERO, thereby designating
//	  'exponent' as an integer value. Values greater than zero will
//	  trigger an error.
//
//	maxPrecision             uint
//	  When this method calculates 'base' raised to the power of
//	  'exponent', the maximum number of decimal digits to the right of
//	  the decimal place in the resulting value will be limited by
//	  'maxPrecision'.
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
//	result                   *big.Int
//	  If the calculation completes successfully, this return value
//	  will be populated with the value of 'base' raised to the power
//	  of 'exponent'.
//
//	resultPrecision          uint
//	  The precision specification for 'result'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'result' and
//	  'resultPrecision' define a numeric value with a fixed number of
//	  decimal digits to the right of the decimal place.
//
//	err                      error
//	  If the calculation encounters an error, an appropriate error
//	  message will be formatted and returned. If the calculation
//	  completes successfully, this return value will be set to 'nil'.
func (bIMathPwrNeutron *bigIntMathPowerNeutron) bigIntToNegativeIntegerPower(
  base *big.Int,
  basePrecision *big.Int,
  exponent *big.Int,
  exponentPrecision *big.Int,
  maxPrecision *big.Int,
  errPrefDto *ePref.ErrPrefixDto) (
  result *big.Int, resultPrecision *big.Int, err error) {

  bIMathPwrNeutron.lock.Lock()

  defer bIMathPwrNeutron.lock.Unlock()

  result = big.NewInt(0)
  resultPrecision = big.NewInt(0)

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathPowerNeutron.bigIntToNegativeIntegerPower",
    "")

  if err != nil {
    return result, resultPrecision, err
  }

  if base == nil {

    return result, resultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'base'",
      }
  }

  if basePrecision == nil {

    return result, resultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'basePrecision'",
      }
  }

  if exponent == nil {

    return result, resultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'exponent'",
      }
  }

  if exponentPrecision == nil {

    return result, resultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'exponentPrecision'",
      }
  }

  if maxPrecision == nil {

    return result, resultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'maxPrecision'",
      }
  }

  bigOne := big.NewInt(1)

  bigZero := big.NewInt(0)

  if base.Cmp(bigZero) == 0 {
    // base is zero result is zero
    return result, resultPrecision, err
  }

  cmpExponentZero := exponent.Cmp(bigZero)

  if cmpExponentZero == 1 {

    return result, resultPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("exponent= '%v'", exponent.Text(10)),
        ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
          "'exponent' is a positive value.\n" +
          "Only negative exponents can be used with this method.",
      }
  }

  if exponentPrecision.Cmp(bigZero) == 1 {

    return result, resultPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("exponent= '%v'", exponent.Text(10)),
        ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
          "'exponent' is NOT an integer. It has fractional digits.\n" +
          "Only negative, integer exponents can be used with this method.",
      }
  }

  if maxPrecision.Cmp(bigZero) == -1 {

    return result, resultPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision.Text(10)),
        ErrMessage: "Error: Input parameter 'maxPrecision' is INVALID!\n" +
          "'maxPrecision' is a negative value.\n" +
          "Only positive 'maxPrecision' values can be used with this method.",
      }
  }

  if cmpExponentZero == 0 {
    // Any number raised to a zero power is one
    result = big.NewInt(1)
    return result, resultPrecision, err
  }

  if exponent.Cmp(big.NewInt(-1)) == 0 {

    // if exponent == -1, result is equal to inverseBigIntNum of base
    result, resultPrecision, err =
      new(BigIntMathDivide).BigIntFracQuotient(bigOne,
        big.NewInt(0),
        base,
        basePrecision,
        maxPrecision)

    if err != nil {

      return result, resultPrecision,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "result, resultPrecision, err =\n" +
            "  new(BigIntMathDivide).BigIntFracQuotient(\n" +
            "bigOne, big.NewInt(0), base, basePrecision, maxPrecision)",
          ErrContext: fmt.Sprintf("bigOne= '1'\n"+
            "dividendPrecision= '0'\n"+
            "base= '%v'\n"+
            "basePrecision= '%v'\n"+
            "maxPrecision= '%v'\n",
            base.Text(10), basePrecision.Text(10),
            maxPrecision.Text(10)),
          ErrMessage: err.Error(),
        }
    }

    return result, resultPrecision, err
  }

  tempExponent := big.NewInt(0).Set(exponent)

  // Exponent now positive integer value
  tempExponent.Neg(tempExponent)

  tempResult := big.NewInt(0).Exp(base, tempExponent, nil)

  tempResultPrecision := big.NewInt(0).Mul(
    basePrecision, tempExponent)

  result, resultPrecision, err =
    new(BigIntMathDivide).BigIntFracQuotient(
      bigOne,
      big.NewInt(0),
      tempResult,
      tempResultPrecision,
      maxPrecision)

  if err != nil {

    return result, resultPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "result, resultPrecision, err =\n" +
          "  new(BigIntMathDivide).BigIntFracQuotient(\n" +
          "bigOne, big.NewInt(0), tempResult, tempResultPrecision, maxPrecision)",
        ErrContext: fmt.Sprintf("bigOne= '1'\n"+
          "dividendPrecision= '0'\n"+
          "tempResult= '%v'\n"+
          "tempResultPrecision= '%v'\n"+
          "maxPrecision= '%v'\n",
          tempResult.Text(10), tempResultPrecision.Text(10),
          maxPrecision.Text(10)),
        ErrMessage: err.Error(),
      }
  }

  return result, resultPrecision, err
}

// bigIntNumRaiseToPositiveIntegerPower
//
//	Assumes that input parameter 'exponent' is a positive integer
//	number. If 'exponent' is negative or if 'exponent' is NOT an
//	integer (precision > 0), an error will be triggered.
//
//	If the exponent is both positive and an integer number, this
//	method proceeds to raise the 'base' parameter to the power of
//	'exponent' and returns the result as a 'BigIntNum' type.
//
//	Examples
//	=========
//
//	base^exponent = result
//
//	  Base    Exponent    Result
//
//	    2        2            4
//	    3        4           81
//	    4.2      3           74.088
//	   10        3         1000
//	   -4.2      3          -74.088
//	   -2.9      4           70.7281
//	   -2        3.8        ERROR - Exponent is Fraction
func (bIMathPwrNeutron *bigIntMathPowerNeutron) bigIntNumRaiseToPositiveIntegerPower(
  base *BigIntNum,
  validateBase bool,
  exponent *BigIntNum,
  validateExponent bool,
  errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

  bIMathPwrNeutron.lock.Lock()

  defer bIMathPwrNeutron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntMathPowerNeutron.bigIntNumRaiseToPositiveIntegerPower",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  if base == nil {

    return BigIntNum{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'base'",
      }
  }

  if exponent == nil {

    return BigIntNum{},
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ErrContext:    "",
        ParameterName: "'exponent'",
      }
  }

  if validateBase {

    err = base.IsValid(ePrefix.XCpy("Validating 'base'").String())

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = base.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'base'\").String()",
          ErrContext: "Error: Input parameter 'base' is INVALID!\n" +
            "'base' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  if validateExponent {

    err = exponent.IsValid(ePrefix.XCpy("Validating 'exponent'").String())

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "err = exponent.IsValid(\n" +
            "  ePrefix.XCpy(\"Validating 'exponent'\").String()",
          ErrContext: "Error: Input parameter 'exponent' is INVALID!\n" +
            "'exponent' FAILED validation tests.",
          ErrMessage: err.Error(),
        }
    }
  }

  exponentNumStr, err := exponent.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "exponentNumStr, err := exponent.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  exponentPrecisionUint, err := exponent.GetPrecisionUint()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "exponentPredisionUint, err := exponent.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if exponentPrecisionUint > 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
        ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
          "'exponent' is an fractional number.\n" +
          "Only integer exponents can be processed by this method.",
      }
  }

  exponentSignValue, err := exponent.GetSign()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "exponentSignValue, err = exponent.GetSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if exponentSignValue < 0 {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
        ErrMessage: "Error: Input parameter 'exponent' is a negative number.\n" +
          "Only Positive Exponents can be processed by this method!",
      }
  }

  bigIBasePrecision := big.NewInt(int64(exponentPrecisionUint))

  bigINewPrecision := big.NewInt(0).Mul(bigIBasePrecision, exponent.bigInt)

  newPrecision := uint(bigINewPrecision.Int64())

  result := big.NewInt(0).Exp(base.bigInt, exponent.bigInt, nil)

  bINumResult, err := new(BigIntNum).NewBigInt(result, newPrecision)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "bINumResult, err := new(BigIntNum).\n" +
          "  NewBigInt(result, newPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = bINumResult.IsValid(ePrefix.XCpy("Validating 'bINumResult'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "err = bINumResult.IsValid(\n" +
          "  ePrefix.XCpy(\"Validating 'bINumResult'\").String()",
        ErrContext: "Error: Final Calculation Result 'bINumResult' is INVALID!\n" +
          "'bINumResult' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return bINumResult, nil
}
