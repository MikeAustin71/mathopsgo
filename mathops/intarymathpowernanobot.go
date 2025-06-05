package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type intAryMathPwrNanobot struct {
  lock sync.Mutex
}

// pwrByTwos
//
//  Raises a *big.Int 'base', to the specified 'power' using the
//  Exponentiation by squaring algorithm. The result of this
//  operation is stored in input parameter 'ia'.
//
//  See:
//  https://en.wikipedia.org/wiki/Exponentiation_by_squaring
//  https://en.wikipedia.org/wiki/Exponentiation_by_squaring#Computation_by_powers_of_2
//  This method is based on revised code taken in part from Ye Lin Aung.
//  https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
//
// This algorithm modified by Mike Rapp to achieve improved performance.
//
//  Input Parameters:
//  =================
//
//  ia                       *IntAry
//    The base which will be raised to an exponent specified by
//    input parameter 'power'. The final calculation result will
//    be stored in this IntAry instance.
//
//  power                    *big.Int
//    The input parameter 'power' may be either a positive or
//    negative integer.
//
//  maxResultPrecision       int
//    'maxResultPrecision' will determine the maximum number of
//    digits to the right of the decimal place in the result.
//
//    Valid values are -1 and values >= zero ('0')
//
//    Values less than -1 will trigger an error.
//
//    A value of -1 signals that no limit will be placed on
//    the number of decimals places to right of the decimal
//    point in the result.
//
//  internalPrecision        int
//    'internalPrecision' will control the number of digits of
//    accuracy to the right of the decimal point maintained by
//    internal multiplication operations used in raising the intAry
//    value to the designated power.
//
//    Valid values are -1 and values >= zero ('0')
//
//    Values less than -1 will trigger an error.
//
//    A value of -1 signals that no limit will be placed on
//    the number of decimals places to right of the decimal
//    point during internal multiplication operations.
//
//  Return Values
//  =============
//
//  error
//    If, during the execution of this method, an error is
//    identified, this method will set the returned error to a non
//    'nil' value and return an appropriate error message.
func (iaMathPwrNano *intAryMathPwrNanobot) pwrByTwos(
  ia *IntAry,
  power *big.Int,
  maxResultPrecision,
  internalPrecision int,
  errPrefDto *ePref.ErrPrefixDto) error {

  iaMathPwrNano.lock.Lock()

  defer iaMathPwrNano.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "intAryMathPwrNanobot.pwrByTwos()",
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

  if power == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'power'",
    }
  }

  if maxResultPrecision < -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: " maxResultPrecision < -1",
      ErrMessage: fmt.Sprintf("Error: Input Parameter 'maxResultPrecision' is INVALID!\n"+
        "'maxResultPrecision' is less than -1.\n"+
        "maxResultPrecision= %v", maxResultPrecision),
    }
  }

  if internalPrecision < -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "internalPrecision < -1",
      ErrMessage: fmt.Sprintf("Error: Parameter 'internalPrecision' is INVALID!\n"+
        "'internalPrecision' is less than -1.\n"+
        "internalPrecision= %v", internalPrecision),
    }
  }

  numSeps, err := ia.GetNumericSeparatorsDto()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "numSeps, err := ia.GetNumericSeparatorsDto()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if maxResultPrecision == internalPrecision {
    internalPrecision += 20
  }

  err = ia.SetInternalFlags()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = ia.SetInternalFlags()",
      ErrContext: "#1 SetInternalFlags()",
      ErrMessage: err.Error(),
    }
  }

  tPower := big.NewInt(0).Set(power)
  one := big.NewInt(1)
  zero := big.NewInt(0)
  two := big.NewInt(2)

  if ia.isZeroValue {

    err = ia.SetIntAryToZero(0)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ia.SetIntAryToZero(0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    err = ia.SetNumericSeparatorsDto(numSeps)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ia.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    return nil
  }

  if tPower.Cmp(two) == -1 {

    if tPower.Cmp(zero) == -1 {

      ia2, err := ia.Inverse(internalPrecision)

      if err != nil {

        return &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "ia2, err := ia.Inverse(internalPrecision)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
      }

      err = ia.CopyIn(&ia2, false)

      if err != nil {

        return &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = ia.CopyIn(&ia2, false)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
      }

      tPower = big.NewInt(0).Mul(tPower, big.NewInt(-1))

    } else if tPower.Cmp(one) == 0 {
      // no change in value. x^1 == x

      err = ia.SetNumericSeparatorsDto(numSeps)

      if err != nil {

        return &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = ia.SetNumericSeparatorsDto(numSeps)",
          ErrContext: "#2 SetNumSeps",
          ErrMessage: err.Error(),
        }
      }

      return nil

    } else if tPower.Cmp(zero) == 0 {

      err = ia.SetIntAryToOne(0)

      if err != nil {

        return &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = ia.SetIntAryToOne(0)",
          ErrContext: "#2 SetIntAryToOne",
          ErrMessage: err.Error(),
        }
      }

      err = ia.SetNumericSeparatorsDto(numSeps)

      if err != nil {

        return &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "ia.SetNumericSeparatorsDto(numSeps)",
          ErrContext: "#3 SetNumericSeparatorsDto",
          ErrMessage: err.Error(),
        }
      }

      return nil
    }

  }

  tBase, err := ia.CopyOut()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "tBase, err := ia.CopyOut()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = ia.SetIntAryToOne(0)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = ia.SetIntAryToOne(0)",
      ErrContext: "#2 SetIntAryToOne",
      ErrMessage: err.Error(),
    }
  }

  for tPower.Cmp(zero) == 1 {
    //temp, _:= intAry{}.NewNumStr("0")

    if big.NewInt(0).Mod(tPower, two).Cmp(one) == 0 {
      //temp = big.NewInt(0).Mul(result, tBase)
      //result = big.NewInt(0).Set(temp)

      err = ia.MultiplyThisBy(&tBase, -1, internalPrecision)

      if err != nil {

        return &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = ia.SetIntAryToOne(0)",
          ErrContext: "#2 SetIntAryToOne",
          ErrMessage: err.Error(),
        }
      }

      //fmt.Println("ia precision = ", ia.GetPrecisionInt())

      if tPower.Cmp(one) == 0 {
        if maxResultPrecision > -1 && maxResultPrecision < ia.GetPrecision() {

          err = ia.SetPrecision(maxResultPrecision, true)

          if err != nil {

            return &FuncReturnError{
              ErrPrefix:  ePrefix.String(),
              ReturnFunc: "err = ia.SetPrecision(maxResultPrecision, true)",
              ErrContext: "",
              ErrMessage: err.Error(),
            }
          }
        }

        err = ia.SetNumericSeparatorsDto(numSeps)

        if err != nil {

          return &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = ia.SetNumericSeparatorsDto(numSeps)",
            ErrContext: "#4 SetNumericSeparatorsDto",
            ErrMessage: err.Error(),
          }
        }

        return nil
      }
    }

    err = tBase.MultiplyThisBy(&tBase, -1, internalPrecision)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = tBase.MultiplyThisBy(&tBase, -1, internalPrecision)",
        ErrContext: "#2 MultiplyThisBy",
        ErrMessage: err.Error(),
      }
    }

    tPower = big.NewInt(0).Div(tPower, two)
  }

  err = ia.SetNumericSeparatorsDto(numSeps)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: " err = ia.SetNumericSeparatorsDto(numSeps)",
      ErrContext: "#5 SetNumericSeparatorsDto",
      ErrMessage: err.Error(),
    }
  }

  if maxResultPrecision == -1 {
    return nil
  }

  if maxResultPrecision < ia.GetPrecision() {

    err = ia.SetPrecision(maxResultPrecision, true)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: " err = ia.SetNumericSeparatorsDto(numSeps)",
        ErrContext: "#5 SetNumericSeparatorsDto",
        ErrMessage: err.Error(),
      }
    }
  }

  return nil
}
