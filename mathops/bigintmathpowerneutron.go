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
