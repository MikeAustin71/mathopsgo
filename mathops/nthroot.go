package mathops

import (
  "fmt"
  "math/big"

  ePref "github.com/MikeAustin71/errpref"
)

/*

  NthRootOp
	=========

	The source code repository for nthroot.go is located at:
			https://github.com/MikeAustin71/mathopsgo.git

	The source file decimal.go is located in directory:
		MikeAustin71/mathopsgo/mathops/nthroot.go


*/

// NthRootOp - Used to extract square roots and nth roots of positive and negative
// real numbers.
//
// Currently nth roots may be passed as positive or negative values. The maximum
// Nth Root integer value is +2,147,483,647. The minimum Nth Root integer value
// is two (-2,147,483,648).
//
// The technique employed to calculate nth roots is known as the
// "shifting nth-root algorithm".
//
// See: https://en.wikipedia.org/wiki/Shifting_nth_root_algorithm
//
// Dependencies: intAry - intary.go
type NthRootOp struct {
  NthRootIntAry      *IntAry
  NthRootBigInt      *big.Int
  NthRootInt         int
  Radicand           *IntAry
  BaseNumBundles     [][]int
  LenBaseNumBundles  int
  BaseNumBundlesIdx  int
  ResultAry          IntAry
  ResultIdx          int
  ResultPrecision    int
  RequestedPrecision int
  BigOne             *big.Int
  Big10              *big.Int
  Big10ToNthPower    *big.Int
  BigZero            *big.Int
  Big3               *big.Int
  Y                  *big.Int // Root Extracted thusfar
  YPrime             *big.Int // Next Value of Y
  Minuend            *big.Int
  Subtrahend         *big.Int
  R                  *big.Int // Let R be the remainder
  RPrime             *big.Int // Let RPrime be the new value of r for next iteration
  BaseNum            *big.Int // Base Number System - always 10
  Alpha              *big.Int // Next n-digits of the radicand
  Beta               *big.Int // Next Digit of the root
}

func (nthrt *NthRootOp) Empty() {

  nthrt.NthRootInt = 0

  nthrt.NthRootBigInt = big.NewInt(0)

  nRt := new(IntAry).New()

  nthrt.NthRootIntAry = &nRt

  nRt = new(IntAry).New()

  nthrt.Radicand = &nRt
  nthrt.BaseNumBundles = make([][]int, 0, 500)
  nthrt.LenBaseNumBundles = 0
  nthrt.BaseNumBundlesIdx = 0
  nthrt.ResultAry = new(IntAry).New()
  nthrt.ResultIdx = 0
  nthrt.ResultPrecision = 0
  nthrt.RequestedPrecision = 0
  nthrt.Big10 = big.NewInt(0)
  nthrt.Big10ToNthPower = big.NewInt(0)
}

// GetNthRootFloat32 - Calculates the Nth Root of a positive real number ('radicand')
// passed to the method as a type float32. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float32 parameter, 'radicand', which
// will be input and positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'radicand'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootFloat32(radicand float32, precision, nthRoot, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetNthRootFloat32",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithFloat32(radicand, precision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)",
        ErrContext: fmt.Sprintf("nthRoot= '%v'",
          nthRoot),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)

}

// GetNthRootFloat64 - Calculates the Nth Root of a positive real number ('radicand')
// passed to the method as a type float64. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float64 parameter, 'radicand', which
// will be input and positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'radicand'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootFloat64(radicand float64, precision, nthRoot, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetNthRootFloat64",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithFloat64(radicand, precision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithFloat64(radicand, precision)",
        ErrContext: fmt.Sprintf("nthRoot= %v precision=%v ", nthRoot, precision),
        ErrMessage: err.Error(),
      }
  }

  iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)",
        ErrContext: fmt.Sprintf("nthRoot= %v ", nthRoot),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)
}

// GetNthRootBigFloat - Calculates the Nth Root of a positive real number ('radicand')
// passed to the method as type Big Float (*big.Float). In addition, the caller
// must supply input parameters for the 'nthRoot' and 'maxPrecision'.
//
// Nth root specifies the root which will be calculated for parameter, 'radicand'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootBigFloat(radicand *big.Float, nthRoot, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetNthRootBigFloat",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithFloatBig(radicand, -1)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithFloatBig(radicand, -1)",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, -1),
        ErrMessage: err.Error(),
      }
  }

  iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)",
        ErrContext: fmt.Sprintf("nthRoot= '%v' precision= '%v'",
          nthRoot, 0),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)
}

// GetNthRootInt - Calculates the Nth Root of a positive real number ('radicand')
// passed to the method as a type int. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the int64 parameter, 'radicand', which
// will be positioned to the right of the decimal place. If 'precision' is a negative
// value, an error will be returned.
//
// Nth root specifies the root which will be calculated for parameter, 'radicand'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootInt(radicand, precision, nthRoot, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetNthRootIn",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  if precision < 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is INVALID!\n"+
          "'precision' is a negative value.\n"+
          "precision= '%v'",
          precision),
      }
  }

  err = ai.SetIntAryWithInt(radicand, uint(precision))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithInt(radicand, uint(precision))",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, uint(precision)),
        ErrMessage: err.Error(),
      }
  }

  iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)",
        ErrContext: fmt.Sprintf("nthRoot= '%v' precision= '%v'",
          nthRoot, 0),
        ErrMessage: err.Error(),
      }
  }

  iaResult, err := nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaResult, err := nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iaResult, nil
}

// GetNthRootInt64 - Calculates the Nth Root of a positive real number ('radicand')
// passed to the method as a type int64. In addition, the caller must supply
// input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the int64 parameter, 'radicand', which
// will be positioned to the right of the decimal place. 'precision' MUST BE a positive
// value. Negative values will trigger an error.
//
// Nth root specifies the root which will be calculated for parameter, 'radicand'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootInt64(radicand int64, precision, nthRoot, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetNthRootInt64",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if precision < 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error: Input parameter 'precision' is INVALID!\n" +
          "'precision' is a negative value.\n" +
          fmt.Sprintf("precision='%v'", precision),
        ErrMessage: "",
      }
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithInt64(radicand, uint(precision))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithInt64(radicand, uint(precision))",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, uint(precision)),
        ErrMessage: err.Error(),
      }
  }

  iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)",
        ErrContext: fmt.Sprintf("nthRoot= '%v' precision= '%v'",
          nthRoot, 0),
        ErrMessage: err.Error(),
      }
  }

  iaResult, err := nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaResult, err := nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return iaResult, nil
}

// GetNthRootBigInt - Calculates the Nth Root of a positive real number ('radicand')
// passed to the method as type Big Int (*big.Int). In addition, the caller
// must supply input parameters for 'precision', 'nthRoot' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the *big.Int parameter, 'radicand', which
// will be positioned to the right of the decimal place.
//
// Nth root specifies the root which will be calculated for parameter, 'radicand'. Example,
// square root, cube root, 4th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootBigInt(radicand *big.Int, precision, nthRoot, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetNthRootBigInt",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithBigInt(radicand, precision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithBigInt(radicand, precision)",
        ErrContext: fmt.Sprintf("radicand= %v, precision=%v",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewInt(nthRoot, 0)",
        ErrContext: fmt.Sprintf("nthRoot= %v, precision=%v",
          nthRoot, 0),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetNthRootIntAry(&ai, &iaNthRoot, maxPrecision)
}

// GetNthRootIntAry  - Calculates the Nth Root of a real number ('radicand')
// passed to the method as a pointer to type intAry.  In addition, the
// caller must supply input parameters for 'nthRoot' and 'maxPrecision'.
//
// 'nthRoot' specifies the root which will be calculated for parameter,
// 'radicand'. Example, square root, cube root, 4th root, 9th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the
// decimal place to which the Nth root will be calculated.  If 'maxPrecision'
// is less than zero, 'maxPrecision' will be automatically set to a value
// of '4,096'.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetNthRootIntAry(
  radicand, nthRoot *IntAry, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetNthRootIntAry",
    "")

  if err != nil {
    return IntAry{}, err
  }

  err = nthrt.calcNthRootGateway(radicand, nthRoot, maxPrecision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.calcNthRootGateway(radicand, nthRoot, maxPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  resultIntAry, err := nthrt.ResultAry.CopyOut()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "resultIntAry, err := nthrt.ResultAry.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return resultIntAry, nil
}

// GetSquareRootFloat32 - Calculates the Square Root of a positive real number ('radicand')
// passed to the method as a type float32. In addition, the caller must supply
// input parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float32 parameter, 'radicand', which
// will be input and positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootFloat32(
  radicand float32,
  precision, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootFloat32",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithFloat32(radicand, precision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithFloat32(radicand, precision)",
        ErrContext: fmt.Sprintf("adicand= %v, precision=%v",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetSquareRootIntAry(&ai, maxPrecision)

}

// GetSquareRootFloat64 - Calculates the Square Root of a positive real number ('radicand')
// passed to the method as a type float64. In addition, the caller must supply
// input parameter 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the float64 parameter, 'radicand', which
// will be input and positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootFloat64(
  radicand float64,
  precision, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootFloat64",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithFloat64(radicand, precision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithFloat64(radicand, precision)",
        ErrContext: fmt.Sprintf("radicand= %v, precision=%v",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootBigFloat - Calculates the Square Root of a positive real number ('radicand')
// passed to the method as type Big Float (*big.Float). In addition, the caller
// must supply input parameter for 'maxPrecision'.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootBigFloat(radicand *big.Float, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootBigFloat",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithFloatBig(radicand, -1)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithFloatBig(radicand, -1)",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, -1),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootInt - Calculates the Square Root of a positive real number ('radicand')
// passed to the method as type Int. In addition, the caller must supply input
// parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the 'int' parameter, 'radicand', which
// will be positioned to the right of the decimal place. If 'precision' is a negative
// value an error will be returned.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootInt(radicand int, precision, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootInt",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  if precision < 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error: Input parameter 'precision' is INVALID!\n" +
          "'precision' is a negative value.\n" +
          fmt.Sprintf("precision= '%v'", precision),
        ErrMessage: "",
      }
  }

  err = ai.SetIntAryWithInt(radicand, uint(precision))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: " err = ai.SetIntAryWithInt(radicand, uint(precision))",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootInt32 - Calculates the Square Root of a positive real number ('radicand')
// passed to the method as type Int32. In addition, the caller must supply input
// parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the 'int32' parameter, 'radicand', which
// will be positioned to the right of the decimal place. 'precision' MUST BE a positive
// value. Negative values will trigger an error.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootInt32(radicand int32, precision, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootInt32",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if precision < 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is INVALID!\n"+
          "'precision' is less than zero."+
          "precision='%v' ", precision),
      }
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithInt32(radicand, uint(precision))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithInt32(radicand, uint(precision))",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  iaResult, err := nthrt.GetSquareRootIntAry(&ai, maxPrecision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaResult, err := nthrt.GetSquareRootIntAry(&ai, maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  return iaResult, nil
}

// GetSquareRootInt64 - Calculates the Square Root of a positive real number ('radicand')
// passed to the method as type Int64. In addition, the caller must supply input
// parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the 'int64' parameter, 'radicand', which
// will be positioned to the right of the decimal place. 'precision' MUST BE a positive
// value. Negative values will trigger an error.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootInt64(radicand int64, precision, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootInt64",
    "")

  if err != nil {
    return IntAry{}, err
  }

  if precision < 0 {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' is INVALID!\n"+
          "'precision' is less than zero."+
          "precision='%v'", precision),
      }
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithInt64(radicand, uint(precision))

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithInt64(radicand, uint(precision))",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  iaResult, err := nthrt.GetSquareRootIntAry(&ai, maxPrecision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaResult, err := nthrt.GetSquareRootIntAry(&ai, maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  return iaResult, nil
}

// GetSquareRootBigInt - Calculates the Square Root of a positive real number ('radicand')
// passed to the method as a pointer to type Big Int (*big.Int). In addition, the caller
// must supply input parameters for 'precision' and 'maxPrecision'.
//
// 'precision' specifies the number of digits in the *big.Int parameter, 'radicand', which
// will be positioned to the right of the decimal place.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootBigInt(radicand *big.Int, precision, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootBigInt",
    "")

  if err != nil {
    return IntAry{}, err
  }

  ai := new(IntAry).New()

  err = ai.SetIntAryWithBigInt(radicand, precision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = ai.SetIntAryWithBigInt(radicand, precision)",
        ErrContext: fmt.Sprintf("radicand= '%v' precision= '%v'",
          radicand, precision),
        ErrMessage: err.Error(),
      }
  }

  return nthrt.GetSquareRootIntAry(&ai, maxPrecision)
}

// GetSquareRootIntAry - Calculates the square Root of a positive real number ('radicand')
// passed to the method as a pointer to type intAry. In addition, the caller
// must supply input parameters for 'maxPrecision'.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the square root will be calculated.
//
// The calculation result is returned as an intAry object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) GetSquareRootIntAry(radicand *IntAry, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.GetSquareRootIntAry",
    "")

  if err != nil {
    return IntAry{}, err
  }

  iaNthRoot, err := new(IntAry).NewTwo(0)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iaNthRoot, err := new(IntAry).NewTwo(0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = nthrt.calcNthRootGateway(radicand, &iaNthRoot, maxPrecision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.calcNthRootGateway(radicand, &iaNthRoot, maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  resultIntAry, err := nthrt.ResultAry.CopyOut()

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "resultIntAry, err = nthrt.ResultAry.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return resultIntAry, nil
}

// NewNthRoot
//
//	Returns the result of a Nth Root calculation. The result
//	is returned as a type 'IntAry'.
//
//	This method calculates the Nth Root of a real number ('radicand')
//	passed to the method as a pointer to type intAry.  In addition,
//	the caller must supply input parameters for 'nthRoot' and
//	'maxPrecision'.
//
//	'nthRoot' specifies the root which will be calculated for
//	parameter, 'radicand'. Example, square root, cube root, 4th
//	root, 9th root etc.
//
//	'maxPrecision' specifies the number of decimals to the right of
//	the decimal place to which the Nth root will be calculated. If
//	'maxPrecision' is less than zero, 'maxPrecision' will be
//	automatically set to a value of '4,096'.
//
//	The calculation result is returned as an intAry object.
//
//	Note: A negative 'radicand' value with an even nthRoot will
//	generate an error.
func (nthrt *NthRootOp) NewNthRoot(radicand, nthRoot *IntAry, maxPrecision int) (IntAry, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.NewNthRoot",
    "")

  if err != nil {
    return IntAry{}, err
  }

  resultIntAry, err := new(NthRootOp).GetNthRootIntAry(radicand, nthRoot, maxPrecision)

  if err != nil {

    return IntAry{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "resultIntAry, err := nthRtOp.GetNthRootIntAry(radicand, nthRoot, maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  return resultIntAry, nil
}

// SetNthRootIntAry  - Calculates the Nth Root of a number ('radicand') passed to the
// method as a pointer to type intAry.  In addition, the caller must supply input
// parameters for 'nthRoot' and 'maxPrecision'.
//
// The difference between this method, 'SetNthRootIntAry' and 'OriginalNthRoot' is in
// the return value.  This method, 'SetNthRootIntAry' does not return the result. Instead,
// the calculation result is stored in the NthRootOp intAry Object, NthRootOp.ResultAry.
// This method is primarily for use by other low level routines seeking to improve performance
// by avoiding the return of a new intAry object.
//
// Nth root specifies the root which will be calculated for parameter, 'radicand'. Example,
// square root, cube root, 4th root, 9th root etc.
//
// 'maxPrecision' specifies the number of decimals to the right of the decimal place to
// which the Nth root will be calculated.
//
// The calculation result is stored in the NthRootOp field, 'NthRootOp.ResultAry'.
// 'NthRootOp.ResultAry' is an intAry Object.
//
// Note: A negative 'radicand' value with an even nthRoot will generate an error.
func (nthrt *NthRootOp) SetNthRootIntAry(radicand, nthRoot *IntAry, maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.SetNthRootIntAry",
    "")

  if err != nil {
    return err
  }

  err = nthrt.calcNthRootGateway(radicand, nthRoot, maxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.calcNthRootGateway(radicand, nthRoot, maxPrecision)",
      ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// calcNthRootGateway - This method is the primary means by which the nth root
// calculation is accessed. All screening and validation of input parameters 'radicand'
// and 'nthRoot' are performed here. If 'radicand' and 'nthRoot' pass all tests for
// validity, this method proceeds to perform the nth root calculation and store the result
// in the NthRootOp data structure. The final result of the nth root calculation is therefore
// stored in data structure element, 'NthRootOp.ResultAry'.
//
// Note: If maxPrecision is less than zero it will automatically be set to a value of
// 4,096 decimal places.
func (nthrt *NthRootOp) calcNthRootGateway(radicand, nthRoot *IntAry, maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.calcNthRootGateway",
    "")

  if err != nil {
    return err
  }

  if maxPrecision < 0 {
    maxPrecision = 4096
  }

  nthrt.ResultAry, err = new(IntAry).NewInt32(0, 0)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthrt.ResultAry, err = new(IntAry).NewInt32(0, 0)",
      ErrContext: fmt.Sprintf("int32Num= '%v' precision= '%v'",
        0, 0),
      ErrMessage: err.Error(),
    }
  }

  radicandIsZero, err := radicand.IsZero()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "radicandIsZero, err := radicand.IsZero()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  // If radicand is zero, the result will always be zero.
  if radicandIsZero {

    err = nthrt.ResultAry.SetIntAryToZero(uint(maxPrecision))

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.ResultAry.SetIntAryToZero(uint(maxPrecision))",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
    }

    return nil
  }

  nthRootIsZero, err := nthRoot.IsZero()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootIsZero, err = nthRoot.IsZero()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRootIsZero {

    err = nthrt.ResultAry.SetIntAryToOne(maxPrecision)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.ResultAry.SetIntAryToOne(maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
    }

    return nil
  }

  nthRootPrecision := nthRoot.GetPrecision()

  nthRootSign, err := nthRoot.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSign, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRootPrecision == 0 && nthRootSign == -1 {

    return nthrt.calcNegativeIntegerNthRoot(radicand, nthRoot, maxPrecision)

  } else if nthRoot.precision > 0 && nthRootSign == -1 {

    return nthrt.calcNegativeFractionalNthRoot(radicand, nthRoot, maxPrecision)

  } else if nthRoot.precision == 0 && nthRootSign == 1 {

    return nthrt.calcPositiveIntegerNthRoot(radicand, nthRoot, maxPrecision)

  } else if nthRoot.precision > 0 && nthRootSign == 1 {

    return nthrt.calcPositiveFractionalNthRoot(radicand, nthRoot, maxPrecision)

  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return &FuncReturnError{
    ErrPrefix:  ePrefix.String(),
    ReturnFunc: "",
    ErrContext: "",
    ErrMessage: fmt.Sprintf("Error - 'nthRoot' configuration failed to match acceptable calculation patterns! "+
      "nthRoot='%v' ", nthRootNumStr),
  }

}

// calcPositiveIntegerNthRoot - Calculates the Nth Root of a radicand where
// nth root is both positive and an integer value.
func (nthrt *NthRootOp) calcPositiveIntegerNthRoot(
  radicand, nthRoot *IntAry,
  maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.calcPositiveIntegerNthRoot",
    "")

  if err != nil {
    return err
  }

  nthRootSignVal, err := nthRoot.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSignVal, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRootSignVal != 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSignVal, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Expected postive 'nthRoot'.\n"+
        "Instead, 'nthRoot' is negative!\n"+
        "nthRoot= %v", nthRootNumStr),
    }
  }

  radicandSignVal, err := radicand.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "radicandSignVal, err := radicand.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if radicandSignVal == -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "radicandSignVal, err := radicand.GetSign()",
      ErrContext: "",
      ErrMessage: "Error: Cannot calculate nthRoot of a negative number when nthRoot is even.\n" +
        "Sign Value of radicand is -1",
    }
  }

  if maxPrecision < 0 {
    maxPrecision = 4096
  }

  err = nthrt.initialize(radicand, nthRoot, maxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.initialize(radicand, nthRoot, maxPrecision)",
      ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.doRootExtraction()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.doRootExtraction()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// calcNegativeIntegerNthRoot - Calculates the Nth Root of a radicand where
// nth root is both negative and an integer value.
func (nthrt *NthRootOp) calcNegativeIntegerNthRoot(
  radicand, nthRoot *IntAry,
  maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.calcNegativeIntegerNthRoot",
    "")

  if err != nil {
    return err
  }

  nthRootSignVal, err := nthRoot.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSignVal, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRootSignVal != -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error expected negative 'nthRoot'.\n"+
        "Instead, 'nthRoot' is positive!"+
        "nthRoot= %v", nthRootNumStr),
    }
  }

  radicandSignVal, err := radicand.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "radicandSignVal, err := radicand.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if radicandSignVal == -1 {

    nthRootIsEvenNum, err := nthRoot.IsEvenNumber()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthRootIsEvenNum, err := nthRoot.IsEvenNumber()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    if nthRootIsEvenNum {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Invalid Entry!\n" +
          "Cannot calculate nthRoot of a negative number when nthRoot is even.",
      }
    }
  }

  if maxPrecision < 0 {
    maxPrecision = 4096
  }

  // Change sign from negative (-) to positive (+)
  err = nthRoot.ChangeSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthRoot.ChangeSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.initialize(radicand, nthRoot, maxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.initialize(radicand, nthRoot, maxPrecision)",
      ErrContext: "Error returned from initialization.",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.doRootExtraction()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.doRootExtraction()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  result, err := nthrt.ResultAry.Inverse(maxPrecision + 100)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "result, err := nthrt.ResultAry.Inverse(maxPrecision + 100)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if result.GetPrecision() > maxPrecision {

    err = result.RoundToPrecision(maxPrecision)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = result.RoundToPrecision(maxPrecision)",
        ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
        ErrMessage: err.Error(),
      }
    }

  }

  // Change sign from positive (+), back to negative (-)
  err = nthRoot.ChangeSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthRoot.ChangeSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthrt.Empty()

  err = nthrt.NthRootIntAry.CopyIn(nthRoot, false)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.NthRootIntAry.CopyIn(nthRoot, false)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.ResultAry.CopyIn(&result, false)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.ResultAry.CopyIn(&result, false)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthrt.Radicand, err = radicand.CopyOutPtr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthrt.Radicand, err = radicand.CopyOutPtr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthrt.ResultPrecision = nthrt.ResultAry.GetPrecision()

  nthrt.RequestedPrecision = maxPrecision

  nthrt.NthRootBigInt, err = nthrt.NthRootIntAry.GetBigInt()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthrt.NthRootBigInt, err = nthrt.NthRootIntAry.GetBigInt()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// calcPositiveFractionalNthRoot - Calculates the Nth Root of a radicand where
// nth root is both negative and an integer value.
func (nthrt *NthRootOp) calcPositiveFractionalNthRoot(
  radicand, nthRoot *IntAry,
  maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.calcPositiveFractionalNthRoot",
    "")

  if err != nil {
    return err
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRoot.GetPrecision() < 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Expected fractional 'nthRoot'.\n"+
        "Instead, 'nthRoot' is an integer value.\n"+
        "nthRoot='%v' ", nthRootNumStr),
    }
  }

  if maxPrecision < 0 {
    maxPrecision = 4096
  }

  internalMaxPrecision := maxPrecision + 100

  fracIntAry, err := new(FracIntAry).NewFracIntAry(nthRoot)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "fracIntAry, err := new(FracIntAry).NewFracIntAry(nthRoot)",
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

  newRadicand, err := radicand.CopyOut()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newRadicand, err := radicand.CopyOut()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathPower).Pwr(
    &newRadicand,
    &fracIntAry.Denominator,
    0,
    internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathPower).Pwr(\n" +
        "  &newRadicand, &fracIntAry.Denominator, 0, internalMaxPrecision)",
      ErrContext: fmt.Sprintf("minResultPrecision= '0' internalMaxPrecision= '%v' ",
        internalMaxPrecision),
      ErrMessage: err.Error(),
    }
  }

  newRadicandSignVal, err := newRadicand.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newRadicandSignVal, err :=  newRadicand.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if newRadicandSignVal == -1 {

    fracIntAryNumeratorIsEvenNum, err := fracIntAry.Numerator.IsEvenNumber()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "fracIntAryNumeratorIsEvenNum, err := fracIntAry.Numerator.IsEvenNumber()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    if fracIntAryNumeratorIsEvenNum {
      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: INVALID ENTRY\n" +
          "Cannot calculate nthRoot of a negative number when nthRoot is even.",
      }
    }
  }

  err = nthrt.initialize(&newRadicand, &fracIntAry.Numerator, maxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.initialize(&newRadicand, &fracIntAry.Numerator, maxPrecision)",
      ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.doRootExtraction()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.doRootExtraction()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// calcNegativeFractionalNthRoot - Calculates the Nth Root of a radicand where
// nth root is both negative and an integer value.
func (nthrt *NthRootOp) calcNegativeFractionalNthRoot(
  radicand, nthRoot *IntAry,
  maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.calcNegativeFractionalNthRoot",
    "")

  if err != nil {
    return err
  }

  if maxPrecision < 0 {
    maxPrecision = 4096
  }

  internalMaxPrecision := maxPrecision + 100

  nthRootSignVal, err := nthRoot.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSignVal, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRootSignVal != -1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Expected negative 'nthRoot'. Instead, 'nthRoot' is positive! "+
        "nthRoot='%v' ", nthRootNumStr),
    }
  }

  // Change sign from negative (-) to positive (+)
  err = nthRoot.ChangeSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthRoot.ChangeSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRoot.GetPrecision() < 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthRoot.ChangeSign()",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Expected fractional 'nthRoot'.\n"+
        "Instead, 'nthRoot' is an integer value."+
        "nthRoot='%v' ", nthRootNumStr),
    }
  }

  fracIntAry, err := new(FracIntAry).NewFracIntAry(nthRoot)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "fracIntAry, err := new(FracIntAry).NewFracIntAry(nthRoot)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)",
      ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'", internalMaxPrecision),
      ErrMessage: err.Error(),
    }
  }

  newRadicand, err := radicand.CopyOut()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newRadicand, err := radicand.CopyOut()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(IntAryMathPower).Pwr(
    &newRadicand,
    &fracIntAry.Denominator,
    0,
    internalMaxPrecision)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(IntAryMathPower).Pwr(\n" +
        "  &newRadicand, &fracIntAry.Denominator, 0, internalMaxPrecision)",
      ErrContext: fmt.Sprintf("minResultPrecision= '%v' internalMaxPrecision= '%v'",
        0, internalMaxPrecision),
      ErrMessage: err.Error(),
    }
  }

  newRadicandSignVal, err := newRadicand.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newRadicandSignVal, err :=  newRadicand.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if newRadicandSignVal == -1 {

    fracIntAryNumeratorIsEvenNum, err := fracIntAry.Numerator.IsEvenNumber()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "fracIntAryNumeratorIsEvenNum, err := fracIntAry.Numerator.IsEvenNumber()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    if fracIntAryNumeratorIsEvenNum {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: INVALID ENTRY!\n" +
          "Cannot calculate nthRoot of a negative number when nthRoot is even.",
      }
    }

    err = nthrt.initialize(&newRadicand, &fracIntAry.Numerator, internalMaxPrecision)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.initialize(&newRadicand, &fracIntAry.Numerator, internalMaxPrecision)",
        ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'", internalMaxPrecision),
        ErrMessage: err.Error(),
      }
    }

    err = nthrt.doRootExtraction()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.doRootExtraction()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    result, err := nthrt.ResultAry.Inverse(internalMaxPrecision + 10)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "result, err := nthrt.ResultAry.Inverse(internalMaxPrecision + 10)",
        ErrContext: fmt.Sprintf("internalMaxPrecision= '%v' + 10", internalMaxPrecision),
        ErrMessage: err.Error(),
      }
    }

    if result.GetPrecision() > maxPrecision {

      err = result.RoundToPrecision(maxPrecision)

      if err != nil {

        return &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "err = result.RoundToPrecision(maxPrecision)",
          ErrContext: fmt.Sprintf("maxPrecision= '%v'", maxPrecision),
          ErrMessage: err.Error(),
        }
      }

    }

    // Change sign from positive (+), back to negative (-)
    err = fracIntAry.Numerator.ChangeSign()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = fracIntAry.Numerator.ChangeSign()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    nthrt.Empty()

    err = nthrt.NthRootIntAry.CopyIn(&fracIntAry.Numerator, false)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.NthRootIntAry.CopyIn(&fracIntAry.Numerator, false)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    err = nthrt.ResultAry.CopyIn(&result, false)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = nthrt.ResultAry.CopyIn(&result, false)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    nthrt.Radicand, err = newRadicand.CopyOutPtr()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthrt.Radicand, err = newRadicand.CopyOutPtr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    nthrt.ResultPrecision = nthrt.ResultAry.GetPrecision()

    nthrt.RequestedPrecision = maxPrecision

    nthrt.NthRootBigInt, err = nthrt.NthRootIntAry.GetBigInt()

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nthrt.NthRootBigInt, err = nthrt.NthRootIntAry.GetBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
    }

    return nil
  }

  return nil
}

// initialize - Initializes the data fields of the NthRootOp structure and validates the
// radicand and nthRoot numerical values passed to the Nth Root Calculation.
func (nthrt *NthRootOp) initialize(radicand, nthRoot *IntAry, maxPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.initialize",
    "")

  if err != nil {
    return err
  }

  err = radicand.IsValid(ePrefix.XCpy("Validating 'radicand'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = radicand.IsValid(ePrefix.XCpy( \"Validating 'radicand'\").String())",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthRoot.IsValid(ePrefix.XCpy(\"Validating 'nthRoot'\").String())",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthRootSignVal, err := nthRoot.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSignVal, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthRootNumStr, err := nthRoot.GetNumStr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRootSignVal < 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: nthRoot is INVALID!\n"+
        "nthRoot is a negative number!\n"+
        "nthRoot='%v'",
        nthRootNumStr),
    }
  }

  if maxPrecision < 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: 'maxPrecision' is INVALID!\n"+
        "'maxPrecision' is less than zero!\n"+
        "maxPrecision='%v'",
        maxPrecision),
    }
  }

  nthRootIsOne, err := nthRoot.IsOne()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootIsOne, err := nthRoot.IsOne()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  if nthRootIsOne {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: fmt.Sprintf("Error: Input Parameter 'nthRoot' INVALID!\n"+
        "'nthRoot' cannot equal 1. "+
        "nthRoot= %v", nthRootNumStr),
    }
  }

  nthrt.NthRootInt, err = nthRoot.GetInt()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthrt.NthRootInt, err = nthRoot.GetInt()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthrt.NthRootIntAry, err = nthRoot.CopyOutPtr()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthrt.NthRootIntAry, err = nthRoot.CopyOutPtr()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  bINum, err := nthRoot.GetBigInt()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "bINum, err := nthRoot.GetBigInt()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthrt.NthRootBigInt = big.NewInt(0).Set(bINum)

  nthrt.Radicand = radicand

  nthrt.RequestedPrecision = maxPrecision

  err = nthrt.bundleInts()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.bundleInts()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.bundleFracs()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.bundleFracs()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.calcPrecision()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.calcPrecision()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  // Set constants for calculations

  nthrt.BigOne = big.NewInt(1)

  nthrt.Big3 = big.NewInt(3)

  nthrt.Big10 = big.NewInt(10)

  nthrt.Big10ToNthPower = big.NewInt(0).Exp(nthrt.Big10, big.NewInt(int64(nthrt.NthRootInt)), nil)

  nthrt.BigZero = big.NewInt(0)

  nthrt.ResultAry = new(IntAry).New()

  err = nthrt.ResultAry.SetPrecision(nthrt.RequestedPrecision, false)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = nthrt.ResultAry.SetPrecision(\n" +
        "  nthrt.RequestedPrecision, false)",
      ErrContext: fmt.Sprintf("nthrt.RequestedPrecision= '%v'",
        nthrt.RequestedPrecision),
      ErrMessage: err.Error(),
    }
  }

  radicandSignVal, err := radicand.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSignVal, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.ResultAry.SetSign(radicandSignVal)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthRootSignVal, err := nthRoot.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  nthrt.Y = big.NewInt(0)

  nthrt.YPrime = big.NewInt(0)

  nthrt.R = big.NewInt(0)

  nthrt.RPrime = big.NewInt(0)

  nthrt.BaseNum = big.NewInt(10)

  nthrt.Alpha = big.NewInt(0)

  nthrt.Beta = big.NewInt(0)

  return nil
}

// bundleInts - computes the bundle size and creates the
// bundle array elements necessary to process the integer
// digits in the original number.
func (nthrt *NthRootOp) bundleInts() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.bundleInts",
    "")

  if err != nil {
    return err
  }

  intNums, err := nthrt.Radicand.GetIntegerDigits()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "intNums, err := nthrt.Radicand.GetIntegerDigits()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  intNumsStats := intNums.GetIntAryStats()

  if intNumsStats.IntAryLen < 1 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "",
      ErrMessage: "Error: intNums array length less than 1",
    }
  }

  bundleSize := 0

  if intNumsStats.IntAryLen <= nthrt.NthRootInt {

    bundleSize = 1

  } else {

    bundleSize = intNumsStats.IntAryLen / nthrt.NthRootInt

    if intNumsStats.IntAryLen > ((intNumsStats.IntAryLen / nthrt.NthRootInt) * nthrt.NthRootInt) {

      bundleSize++

    }

  }

  nthrt.BaseNumBundles = make([][]int, bundleSize)

  bundleIdx := bundleSize - 1
  intAryIdx := 0

  intAry, _, err := intNums.GetIntAryElements()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "intAry, _, err := intNums.GetIntAryElements()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  for j := intNumsStats.IntAryLen - 1; j >= 0; j -= nthrt.NthRootInt {

    bundle := make([]int, nthrt.NthRootInt)

    intAryIdx = j

    for k := nthrt.NthRootInt - 1; k >= 0; k-- {

      if intAryIdx < 0 {
        break
      }

      bundle[k] = int(intAry[intAryIdx])
      intAryIdx--
    }

    nthrt.BaseNumBundles[bundleIdx] = append(nthrt.BaseNumBundles[bundleIdx], bundle...)

    bundleIdx--
  }

  return nil
}

// bundleFracs - computes the numeric bundle size and creates
// the bundle array elements required to process the fractional
// digits in the original number.
func (nthrt *NthRootOp) bundleFracs() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.bundleFracs",
    "")

  if err != nil {
    return err
  }

  if nthrt.Radicand.GetPrecision() < 1 {
    return nil
  }

  fracNums, err := nthrt.Radicand.GetFractionalDigits()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "fracNums, err := nthrt.Radicand.GetFractionalDigits()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  iFracNumStats := fracNums.GetIntAryStats()

  iFAry, _, err := fracNums.GetIntAryElements()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "iFAry, _, err := fracNums.GetIntAryElements()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  for i := 1; i < iFracNumStats.IntAryLen; i += nthrt.NthRootInt {

    bundle := make([]int, nthrt.NthRootInt)

    for j := 0; j < nthrt.NthRootInt; j++ {

      if i+j < iFracNumStats.IntAryLen {

        bundle[j] = int(iFAry[i+j])

      }
    }

    nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)

  }

  return nil
}

// CalcPrecision - calculates the bundle size and creates
// the bundle array elements necessary to process the nth
// root to the requested number of decimal places to the
// right of the decimal point.
func (nthrt *NthRootOp) calcPrecision() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.calcPrecision",
    "")

  if err != nil {
    return err
  }

  if nthrt.Radicand.GetPrecision() < 0 {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "nthrt.Radicand.GetPrecision() < 0",
      ErrMessage: fmt.Sprintf("Error: Existing precision is less than zero!\nnthrt.Radicand.precision= %v", nthrt.Radicand.GetPrecision()),
    }
  }

  existingPrecision := nthrt.Radicand.GetPrecision() / nthrt.NthRootInt

  if nthrt.Radicand.GetPrecision() != existingPrecision*nthrt.NthRootInt {
    existingPrecision++
  }

  bundle := make([]int, nthrt.NthRootInt)

  if nthrt.RequestedPrecision <= existingPrecision {

    // nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)

    nthrt.RequestedPrecision = existingPrecision

    // return nil
  }

  nthrt.RequestedPrecision++

  deltaPrecision := nthrt.RequestedPrecision - existingPrecision

  for i := 0; i < deltaPrecision; i++ {

    nthrt.BaseNumBundles = append(nthrt.BaseNumBundles, bundle)

  }

  nthrt.LenBaseNumBundles = len(nthrt.BaseNumBundles)

  nthrt.ResultPrecision = nthrt.RequestedPrecision

  return nil
}

func (nthrt *NthRootOp) doRootExtraction() error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.doRootExtraction",
    "")

  if err != nil {
    return err
  }

  nthrt.Y = big.NewInt(0)

  nthrt.Minuend = big.NewInt(0)

  nthrt.Subtrahend = big.NewInt(0)

  nthrt.R = big.NewInt(0)

  for i := 0; i < nthrt.LenBaseNumBundles; i++ {

    err = nthrt.findNextRoot(i)

    if err != nil {

      return &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: fmt.Sprintf("err = nthrt.findNextRoot(%d)", i),
        ErrContext: "for i := 0; i < nthrt.LenBaseNumBundles; i++",
        ErrMessage: err.Error(),
      }
    }

  }

  nthrtRadicandSignVal, err := nthrt.Radicand.GetSign()

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "nthrtRadicandSignVal, err := nthrt.Radicand.GetSign()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.ResultAry.SetSign(nthrtRadicandSignVal)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.ResultAry.SetSign(nthrtRadicandSignVal)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.ResultAry.OptimizeIntArrayLen(false)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.ResultAry.OptimizeIntArrayLen(false)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = nthrt.ResultAry.RoundToPrecision(nthrt.RequestedPrecision - 1)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = nthrt.ResultAry.RoundToPrecision(nthrt.RequestedPrecision - 1)",
      ErrContext: fmt.Sprintf("nthrt.RequestedPrecision= '%v'\n"+
        "nthrt.RequestedPrecision-1= '%v'", nthrt.RequestedPrecision, nthrt.RequestedPrecision-1),
      ErrMessage: err.Error(),
    }
  }

  return nil
}

func (nthrt *NthRootOp) findNextRoot(bundleIdx int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NthRootOp.findNextRoot",
    "")

  if err != nil {
    return err
  }

  bundle := nthrt.getBundleBigInt(bundleIdx)

  // alpha = next n-digits of radicand
  nthrt.Alpha = big.NewInt(0).Set(bundle)

  nthrt.RPrime = big.NewInt(-1)

  // nthrt.R already set
  // n = nthrt.NthRootInt already set
  // nthrt.Y already set
  // nthrt.BaseNum = 10
  // nthrt.Big10ToNthPower = BaseNum^n

  itatr := big.NewInt(9)
  term1a := big.NewInt(0)
  term1b := big.NewInt(0)

  term2a1 := big.NewInt(0)
  term2a2 := big.NewInt(0)
  term2a := big.NewInt(0)

  term2b := big.NewInt(0)
  term2b1 := big.NewInt(0)
  term2b2 := big.NewInt(0)

  term1a = big.NewInt(0).Mul(nthrt.Big10ToNthPower, nthrt.R)
  term1b = big.NewInt(0).Set(nthrt.Alpha)
  nthrt.Minuend = big.NewInt(0).Add(term1a, term1b)

  for itatr.Cmp(nthrt.BigZero) > -1 &&
    nthrt.RPrime.Cmp(nthrt.BigZero) == -1 {

    nthrt.Beta = big.NewInt(0).Set(itatr)
    nthrt.YPrime = big.NewInt(0).Mul(nthrt.Y, nthrt.Big10)
    nthrt.YPrime = big.NewInt(0).Add(nthrt.YPrime, nthrt.Beta)

    term2a1 = big.NewInt(0).Mul(nthrt.BaseNum, nthrt.Y)
    term2a2 = big.NewInt(0).Add(term2a1, nthrt.Beta)
    term2a = big.NewInt(0).Exp(term2a2, big.NewInt(int64(nthrt.NthRootInt)), nil)

    term2b1 = big.NewInt(0).Set(nthrt.Big10ToNthPower)
    term2b2 = big.NewInt(0).Exp(nthrt.Y, big.NewInt(int64(nthrt.NthRootInt)), nil)

    term2b = big.NewInt(0).Mul(term2b1, term2b2)

    nthrt.Subtrahend = big.NewInt(0).Sub(term2a, term2b)

    nthrt.RPrime = big.NewInt(0).Sub(nthrt.Minuend, nthrt.Subtrahend)

    itatr = big.NewInt(0).Sub(itatr, nthrt.BigOne)
  }

  nthrt.R = big.NewInt(0).Set(nthrt.RPrime)

  nthrt.Y = big.NewInt(0).Set(nthrt.YPrime)

  err = nthrt.ResultAry.AppendToIntAry(uint8(nthrt.Beta.Int64()))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "newRadicand, err := radicand.CopyOut()",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

func (nthrt *NthRootOp) getBundleBigInt(idx int) *big.Int {

  bigBundleVal := big.NewInt(0)

  for i := 0; i < nthrt.NthRootInt; i++ {

    bigBundleVal = big.NewInt(0).Mul(bigBundleVal, nthrt.Big10)
    bigBundleVal = big.NewInt(0).Add(bigBundleVal, big.NewInt(0).SetInt64(int64(nthrt.BaseNumBundles[idx][i])))

  }

  return bigBundleVal
}
