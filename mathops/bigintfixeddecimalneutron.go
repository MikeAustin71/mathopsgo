package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "sync"
)

type bigIntFixedDecNeutron struct {
  lock *sync.Mutex
}

// DivideByTenToPower
//
// Divides the numeric value of from input parameter 'bigIFxDec'
// (type BigIntFixedDecimal) by 10 to the power of 'exponent'.
//
//		result = BigIntFixedDecimal / 10^exponent
//
//	 Input Parameters
//	 ================
//
//	 bigIFxDec         *BigIntFixedDecimal
//
//	 An instance of 'BigIntFixedDecimal'. The numeric value of
//	 'bigIFxDec' will be divided by 10 to the power of 'exponent'.
//	 The resulting numeric value will be then be stored in
//	 'bigIFxDec'. The original value of 'bigIFxDec' will be
//	 overwritten.
//
//
//	 exponent          uint
//
//	 The value of BigIntFixedDecimal instance 'bigIFxDec' will be
//	 divided by ten raised to the power of 'exponent'.
//
// This method will destroy and overwrite the previous value of
// the current BigIntFixedDecimal instance with the results of
// this calculation.
//
//	NOTE
//	====
//
// This method does NOT test the validity of 'bigIFxDec', an
// instance of type BigIntNum. The calling method must
// do this!
func (bigIFdNeutron *bigIntFixedDecNeutron) divideByTenToPower(
  bigIFxDec *BigIntFixedDecimal,
  exponent uint,
  errPrefDto *ePref.ErrPrefixDto) error {

  if bigIFdNeutron.lock == nil {
    bigIFdNeutron.lock = new(sync.Mutex)
  }

  bigIFdNeutron.lock.Lock()

  defer bigIFdNeutron.lock.Unlock()

  var ePrefix *ePref.ErrPrefixDto

  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
    errPrefDto,
    "bigIntFixedDecNeutron.divideByTenToPower",
    "")

  if err != nil {
    return err
  }

  if bigIFxDec == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec'",
    }
  }

  if bigIFxDec.integerNum == nil {

    return &InputPtrNilError{
      ErrPrefix:     ePrefix.String(),
      ParameterName: "'bigIFxDec.integerNum'",
    }
  }

  if bigIFxDec.integerNum.Cmp(big.NewInt(0)) == 0 {
    return nil
  }

  scale :=
    big.NewInt(0).Exp(
      big.NewInt(10),
      big.NewInt(int64(exponent)), nil)

  factor := new(BigIntFixedDecimal).New(scale, 0)

  bigIFxDec2, err := new(bigIntFixedDecUtility).copyOut(
    bigIFxDec,
    ePrefix.XCpy("Copying 'bigIFxDec' -> bigIFxDec2"))

  if err != nil {
    return err
  }

  newPrecision := bigIFxDec2.precision + exponent

  result, err :=
    BigIntMathDivide{}.FixedDecimalFracQuotient(
      bigIFxDec2, factor, newPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "result, err := BigIntMathDivide{}.FixedDecimalFracQuotient(\n" +
        "bigIFxDec2, factor, newPrecision)",
      ErrContext: fmt.Sprintf("factor= '%v' newPrecision= '%v'", factor, newPrecision),
      ErrMessage: err.Error(),
    }
  }

  return new(bigIntFixedDecUtility).copyIn(
    bigIFxDec,
    &result,
    ePrefix.XCpy("Copying 'result' -> 'bigIFxDec'"))
}
