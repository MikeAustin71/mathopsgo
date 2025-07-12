package mathops

import (
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
)

/*

Calculate Any Logarithm manually:
	https://www.youtube.com/watch?v=TUAFzuMVem0

Properties of Logarithms:
https://www.youtube.com/watch?v=AAW7WRFBKdw&t=6s

https://en.wikipedia.org/wiki/Logarithm
log[b](x) = y exactly if b^y = x

log[10](450) = 2.6532125137753436793763169117857
precision = 31

LogN(1500000;4) = log[4](1500000) = 10.2582655350227

*/

type BigIntMathLogarithms struct {
  Base      BigIntNum
  XNumber   BigIntNum
  YExponent BigIntNum
}

// BigIntNumLogBaseOfX - Computes the log[base](xNum). 'base' and 'xNum' are passed as
// BigIntNum types.
//
// 'maxPrecision' is an uint specifying the precision or integerNum of digits to the
// right of the decimal place in the result.
//
// Calculation Formula:
// ====================
//
//	log[base](x) = y
//
//	x = base^y
//	'y' is the exponent of base required to set base^y equal to x.
//	The calculation result 'y' of log[base](xNum) is returned as a BigIntNum.
//
// Input Parameters:
// =================
//
// base		BigIntNum	- The base of the logarithm. 'base' must be an integer integerNum
//
//	greater than one (base > +1).
//
// xNum		BigIntNum - The 'x' value such that base^y = 'x'.  xNum must be greater
//
//	than zero.
//
// Return Values:
// ==============
// y			BigIntNum - If the function completes successfully, the result 'y' will be
//
//											returned as a BigIntNum.
//
//					error			- If the function encounters an error, an error message will be returned
//											as type 'error'. If the function completes successfully, this value
//											will be set to 'nil'.
//
//	        log Value:  10.25826553502266513433782776044
//
// Expected log Value:  10.2582655350227
//
// base:  4
// xNum:  1500000
func (bLog BigIntMathLogarithms) BigIntNumLogBaseOfX(
  base, xNum BigIntNum, maxPrecision uint) (result BigIntNum, err error) {

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathLogarithms.BigIntNumLogBaseOfX",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  biMaxPrecision := big.NewInt(0).SetUint64(uint64(maxPrecision))
  factor := big.NewInt(4)

  cycles := big.NewInt(40)
  if biMaxPrecision.Cmp(big.NewInt(10)) == 1 {
    cycles = big.NewInt(0).Mul(biMaxPrecision, factor)
  }

  maxInternalPrecision := big.NewInt(0).Mul(cycles, factor)

  var baseBigInt, basePrecisionBigInt,
    xNumBigInt, xNumPrecisionBigInt *big.Int

  baseBigInt, err = base.GetIntegerValue()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "baseBigInt, err = base.GetIntegerValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  basePrecisionBigInt, err = base.GetPrecisionBigInt()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "basePrecisionBigInt, err = base.GetPrecisionBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xNumBigInt, err = xNum.GetIntegerValue()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "xNumBigInt, err = xNum.GetIntegerValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xNumPrecisionBigInt, err = xNum.GetPrecisionBigInt()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "xNumPrecisionBigInt, err = xNum.GetPrecisionBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  biResult, biResultPrecision, errX :=
    BigIntMathLogarithms{}.LogBaseOfXByDivide(
      baseBigInt,
      basePrecisionBigInt,
      xNumBigInt,
      xNumPrecisionBigInt,
      maxInternalPrecision,
      biMaxPrecision,
      cycles)

  if errX != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "biResult, biResultPrecision, errX := \n" +
          "  BigIntMathLogarithms{}.LogBaseOfXByDivide()",
        ErrContext: "",
        ErrMessage: errX.Error(),
      }
  }

  result, errX =
    new(BigIntNum).NewBigIntBigPrecision(
      biResult,
      biResultPrecision)

  if errX != nil {

    err = fmt.Errorf("%v\n"+
      "Error returned by: \n"+
      "result, errX = new(BigIntNum).NewBigIntBigPrecision(...)\n"+
      "Error= %v\n",
      ePrefix,
      errX.Error())
  }

  return result, nil
}

//	log Value:  10.2582655350226651343378277604421
//
// Expected log Value:  10.2582655350227
// base:  4
// xNum:  1500000
func (bLog BigIntMathLogarithms) BigIntLogBaseOfX(
  base *big.Int,
  basePrecision *big.Int,
  xNum *big.Int,
  xNumPrecision *big.Int,
  maxPrecision *big.Int) (logResult,
  logResultPrecision *big.Int,
  err error) {

  logResult = big.NewInt(0)
  logResultPrecision = big.NewInt(0)

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathLogarithms.BigIntLogBaseOfX",
    "")

  if err != nil {
    return logResult, logResultPrecision, err
  }

  bigZero := big.NewInt(0)

  if base == nil {

    return logResult, logResultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'base'",
      }
  }

  if basePrecision == nil {

    return logResult, logResultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'basePrecision'",
      }
  }

  if xNum == nil {

    return logResult, logResultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'xNum'",
      }
  }

  if xNumPrecision == nil {

    return logResult, logResultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'xNumPrecision'",
      }
  }

  if maxPrecision == nil {

    return logResult, logResultPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'maxPrecision'",
      }
  }

  if base.Cmp(bigZero) != 1 {

    err = fmt.Errorf("%v\n"+
      "Error: Input Parameter 'base' is INVALID!\n"+
      "base='%v'\n",
      ePrefix,
      base.Text(10))

    return logResult, logResultPrecision, err
  }

  if basePrecision.Cmp(bigZero) == -1 {

    err = fmt.Errorf("%v\n"+
      "Error: Input Parameter 'basePrecision' is negative and INVALID!\n"+
      "basePrecision='%v'\n",
      ePrefix,
      basePrecision.Text(10))

    return logResult, logResultPrecision, err

  }

  if xNum.Cmp(bigZero) == 0 {

    err = fmt.Errorf("%v\n"+
      "Error: Input Parameter 'xNum' is INVALID!\n"+
      "xNum='%v'\n",
      ePrefix,
      xNum.Text(10))

    return logResult, logResultPrecision, err
  }

  if xNumPrecision.Cmp(bigZero) == -1 {

    err = fmt.Errorf("%v\n"+
      "Error: Input Parameter 'xNumPrecision' is LESS THAN ZERO!\n"+
      "xNumPrecision='%v'\n",
      ePrefix,
      xNumPrecision.Text(10))

    return logResult, logResultPrecision, err
  }

  if maxPrecision.Cmp(bigZero) == -1 {

    err = fmt.Errorf("%v\n"+
      "Error: Input Parameter 'maxPrecision' is LESS THAN ZERO!\n"+
      "maxPrecision='%v'\n",
      ePrefix,
      maxPrecision.Text(10))

    return logResult, logResultPrecision, err
  }

  factor := big.NewInt(4)

  cycles := big.NewInt(40)

  if maxPrecision.Cmp(big.NewInt(10)) == 1 {

    cycles = big.NewInt(0).Mul(maxPrecision, factor)
  }

  maxInternalPrecision := big.NewInt(0).Mul(cycles, factor)

  var errX error

  logResult, logResultPrecision, errX =
    bLog.LogBaseOfXByDivide(
      base,
      basePrecision,
      xNum,
      xNumPrecision,
      maxInternalPrecision,
      maxPrecision,
      cycles)

  if errX != nil {

    err = fmt.Errorf("%v\n"+
      "Error returned from logResult, logResultPrecision, errX = bLog.LogBaseOfXByDivide(...)\n"+
      "Error= %v\n",
      ePrefix,
      errX.Error())

    return logResult, logResultPrecision, err
  }

  return logResult, logResultPrecision, nil
}

// LogBaseOfXByDivide - Generates logs for specified
// bases and X-numbers. Calculation methodology derrived
// from following:
//
// Calculate Any Logarithm manually -
// https://www.youtube.com/watch?v=TUAFzuMVem0
func (bLog BigIntMathLogarithms) LogBaseOfXByDivide(
  base *big.Int,
  basePrecision *big.Int,
  xNum *big.Int,
  xNumPrecision *big.Int,
  maxInternalPrecision *big.Int,
  maxPrecision *big.Int,
  cycles *big.Int) (logResult,
  logResultPrecision *big.Int,
  err error) {

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathLogarithms.LogBaseOfXByDivide",
    "")

  if err != nil {
    return big.NewInt(0), big.NewInt(0), err
  }

  logResult = big.NewInt(0)
  logResultPrecision = big.NewInt(0)
  bigZero := big.NewInt(0)

  if base == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'base'",
      }
  }

  if basePrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'basePrecision'",
      }
  }

  if xNum == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'xNum'",
      }
  }

  if xNumPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'xNumPrecision'",
      }
  }

  if maxInternalPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'maxInternalPrecision'",
      }
  }

  if maxPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'maxPrecision'",
      }
  }

  if cycles == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'cycles'",
      }
  }

  if basePrecision.Cmp(bigZero) == -1 {

    return big.NewInt(0), big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("basePrecision='%v'", basePrecision.Text(10)),
        ErrMessage: "Error: Input parameter 'basePrecision' is LESS THAN ZERO!",
      }
  }

  if xNumPrecision.Cmp(bigZero) == -1 {

    err = fmt.Errorf("%v\n"+
      "Error: Input parameter 'xNumPrecision' is LESS THAN ZERO!\n"+
      "xNumPrecision='%v'\n",
      ePrefix,
      xNumPrecision.Text(10))

    return logResult, logResultPrecision, err
  }

  if maxInternalPrecision.Cmp(bigZero) == -1 {

    err = fmt.Errorf("%v\n"+
      "Error: Input parameter 'maxInternalPrecision' is LESS THAN ZERO!\n"+
      "maxInternalPrecision='%v'\n",
      ePrefix,
      maxInternalPrecision.Text(10))

    return logResult, logResultPrecision, err
  }

  if maxPrecision.Cmp(bigZero) == -1 {

    err = fmt.Errorf("%v\n"+
      "Error: Input parameter 'maxPrecision' is LESS THAN ZERO!\n"+
      "maxPrecision='%v'\n",
      ePrefix,
      maxPrecision.Text(10))

    return logResult, logResultPrecision, err
  }

  if cycles.Cmp(big.NewInt(5)) == -1 {
    cycles.Mul(maxPrecision, big.NewInt(5))
  }

  tXNum := big.NewInt(0).Set(xNum)
  tXNumPrecision := big.NewInt(0).Set(xNumPrecision)
  bigOne := big.NewInt(1)

  iBase, iBasePrecision, err :=
    new(BigIntMathDivide).BigIntFracQuotient(
      bigOne,
      big.NewInt(0),
      base,
      basePrecision,
      maxInternalPrecision)

  if err != nil {

    return big.NewInt(0), big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "iBase, iBasePrecision, err := new(BigIntMathDivide).BigIntFracQuotient(...)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  ri := big.NewInt(0)
  riPrecision := big.NewInt(0)

  cmpNums, err :=
    new(BigIntMath).BigIntPrecisionCmp(
      tXNum,
      tXNumPrecision,
      base,
      basePrecision)

  if err != nil {

    return tXNum, tXNumPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "cmpNums, err := new(BigIntMath).BigIntPrecisionCmp(...)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  for cmpNums == 1 {

    tXNum, tXNumPrecision, err =
      new(BigIntMathMultiply).BigIntMultiply(
        tXNum,
        tXNumPrecision,
        iBase,
        iBasePrecision)

    if err != nil {

      return tXNum, tXNumPrecision,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "tXNum, tXNumPrecision, errX = new(BigIntMathMultiply).BigIntMultiply(...)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    tXNum, tXNumPrecision, err =
      new(BigIntMath).RoundToMaxPrecision(
        tXNum,
        tXNumPrecision,
        maxInternalPrecision,
        true)

    if err != nil {

      return tXNum, tXNumPrecision,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "tXNum, tXNumPrecision, errX = new(BigIntMath).RoundToMaxPrecision(...)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    ri.Add(ri, bigOne)

    cmpNums, err =
      new(BigIntMath).BigIntPrecisionCmp(
        tXNum,
        tXNumPrecision,
        base,
        basePrecision)
  }

  if err != nil {

    return tXNum, tXNumPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "cmpNums, err = new(BigIntMath).BigIntPrecisionCmp(...)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  p := big.NewInt(1)
  pPrecision := big.NewInt(0)
  oneHalf := big.NewInt(5)
  oneHalfPrecision := big.NewInt(1)
  uint64Cycles := cycles.Uint64()

  for i := uint64(0); i < uint64Cycles; i++ {

    tXNum.Mul(tXNum, tXNum)
    tXNumPrecision.Add(tXNumPrecision, tXNumPrecision)

    tXNum, tXNumPrecision, err =
      new(BigIntMath).RoundToMaxPrecision(
        tXNum,
        tXNumPrecision,
        maxInternalPrecision,
        true)

    if err != nil {

      return tXNum, tXNumPrecision,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "tXNum, tXNumPrecision, errX = new(BigIntMath).RoundToMaxPrecision(...)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    p.Mul(p, oneHalf)
    pPrecision.Add(pPrecision, oneHalfPrecision)

    p, pPrecision, err =
      new(BigIntMath).RoundToMaxPrecision(
        p,
        pPrecision,
        maxInternalPrecision,
        true)

    if err != nil {

      return tXNum, tXNumPrecision,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "p, pPrecision, errX = new(BigIntMath).RoundToMaxPrecision(...)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    cmpNums, err =
      new(BigIntMath).BigIntPrecisionCmp(
        tXNum,
        tXNumPrecision,
        base,
        basePrecision)

    if err != nil {

      return tXNum, tXNumPrecision,
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "cmpNums, err = new(BigIntMath).BigIntPrecisionCmp(...)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    if cmpNums == 1 {

      ri, riPrecision, err =
        new(BigIntMathAdd).BigIntAdd(
          ri,
          riPrecision,
          p,
          pPrecision)

      if err != nil {

        return logResult, logResultPrecision,
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "ri, riPrecision, err = new(BigIntMathAdd).BigIntAdd(...)",
            ErrContext: "",
            ErrMessage: err.Error(),
          }
      }

      tXNum.Mul(tXNum, iBase)
      tXNumPrecision.Add(tXNumPrecision, iBasePrecision)

      tXNum, tXNumPrecision, err =
        new(BigIntMath).RoundToMaxPrecision(
          tXNum,
          tXNumPrecision,
          maxInternalPrecision,
          true)

      if err != nil {

        return logResult, logResultPrecision,
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "tXNum, tXNumPrecision, err =new(BigIntMath).RoundToMaxPrecision(...)",
            ErrContext: "",
            ErrMessage: err.Error(),
          }
      }
    }
  }

  logResult, logResultPrecision, err =
    new(BigIntMath).RoundToMaxPrecision(
      ri,
      riPrecision,
      maxPrecision,
      true)

  if err != nil {

    return logResult, logResultPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "logResult, logResultPrecision, err = new(BigIntMath).RoundToMaxPrecision(...)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return logResult, logResultPrecision, nil
}

/*
func (bLog BigIntMathLogarithms) LogBaseOfXByDivide(
	base,
	basePrecision,
	xNum,
	xNumPrecision,
	maxInternalPrecision,
	maxPrecision ,
	cycles *big.Int) (logResult,
													logResultPrecision *big.Int,
													err error) {

	ePrefix := "BigIntMathLogarithms.LogBaseOfXByDivide() "
	logResult = big.NewInt(0)
	logResultPrecision = big.NewInt(0)
	err = nil
	tXNum := big.NewInt(0).Set(xNum)
	tXNumPrecision := big.NewInt(0).Set(xNumPrecision)
	bigOne := big.NewInt(1)

	iBase, iBasePrecision, errX :=
		BigIntMathDivide{}.BigIntFracQuotient(
			bigOne,
			big.NewInt(0),
			base,
			basePrecision,
			maxInternalPrecision)

	if errX != nil {
		err = fmt.Errorf(ePrefix + "%v", errX)
		return logResult, logResultPrecision, err
	}

	ri := big.NewInt(0)
	riPrecision := big.NewInt(0)

	cmpNums :=
		BigIntMath{}.BigIntPrecisionCmp(
			tXNum,
			tXNumPrecision,
			base,
			basePrecision)

	for cmpNums == 1 {

		tXNum, tXNumPrecision, errX =
			BigIntMathMultiply{}.BigIntMultiply(
				tXNum,
				tXNumPrecision,
				iBase,
				iBasePrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v", errX)
			return logResult, logResultPrecision, err
		}

		tXNum, tXNumPrecision, errX =
			BigIntMath{}.RoundToMaxPrecision(
				tXNum,
				tXNumPrecision,
				maxInternalPrecision,
				true)

		if errX != nil {
			err = fmt.Errorf(ePrefix +
				"%v", errX.Error())
			return logResult, logResultPrecision, err
		}

		ri.Add(ri, bigOne	)

		cmpNums =
			BigIntMath{}.BigIntPrecisionCmp(
				tXNum,
				tXNumPrecision,
				base,
				basePrecision)
	}

	p := big.NewInt(1)
	pPrecision := big.NewInt(0)
	oneHalf := big.NewInt(5)
	oneHalfPrecision := big.NewInt(1)
	uint64Cycles := cycles.Uint64()

	for i := uint64(0); i < uint64Cycles; i++ {

		tXNum, tXNumPrecision, errX =
			BigIntMathMultiply{}.BigIntMultiply(
				tXNum,
				tXNumPrecision,
				tXNum,
				tXNumPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v", errX)
			return logResult, logResultPrecision, err
		}

		p, pPrecision, errX =
			BigIntMathMultiply{}.BigIntMultiply(
				p,
				pPrecision,
				oneHalf,
				oneHalfPrecision)

		if errX != nil {
			err = fmt.Errorf(ePrefix + "%v", errX)
			return logResult, logResultPrecision, err
		}

		cmpNums =
			BigIntMath{}.BigIntPrecisionCmp(
				tXNum,
				tXNumPrecision,
				base,
				basePrecision)

		if cmpNums==1 {

			ri, riPrecision, errX =
				BigIntMathAdd{}.BigIntAdd(
					ri,
					riPrecision,
					p,
					pPrecision)

			tXNum, tXNumPrecision, errX =
				BigIntMathMultiply{}.BigIntMultiply(
					tXNum,
					tXNumPrecision,
					iBase,
					iBasePrecision)

			if errX != nil {
				err = fmt.Errorf(ePrefix + "%v", errX)
				return logResult, logResultPrecision, err
			}

		}

		tXNum, tXNumPrecision, errX =
			BigIntMath{}.RoundToMaxPrecision(
				tXNum,
				tXNumPrecision,
				maxInternalPrecision,
				true)

		if errX != nil {
			err = fmt.Errorf(ePrefix +
				"%v", errX.Error())
			return logResult, logResultPrecision, err
		}

	}

	logResult, logResultPrecision, errX =
		BigIntMath{}.RoundToMaxPrecision(
			ri,
			riPrecision,
			maxPrecision,
			true)

	if errX != nil {
		logResult.Set(big.NewInt(0))
		logResultPrecision.Set(big.NewInt(0))
		err = fmt.Errorf(ePrefix +
			"%v", errX.Error())
		return logResult, logResultPrecision, err
	}

	err = nil

	return logResult, logResultPrecision, err
}
*/

// EPwrXFromMaclaurinSeries = Computes value of e^exponent where e is the mathematical
// constant, "Euler's integerNum".
//
// "Euler's integerNum" or 'e' truncated to 50 decimal places is:
//
//	2.71828182845904523536028747135266249775724709369995
//
// Rounded to 50-decimal places is:
//
//	2.71828182845904523536028747135266249775724709369996
//
// Reference: https://en.wikipedia.org/wiki/E_(mathematical_constant)
//
// 2.7182818284590452353602874713526624977572470936999595749669676277240766303
// 5354759457138217852516642742746
// Reference: https://oeis.org/A001113/list
//
// This function works best where 'exponent' is near zero.
//
// Input Parameters
// ================
// exponent	BigIntNum - The exponent to which the mathematical constant 'e' will
//
//	be raised in order to compute the value of e^exponent.
//
// n				int64			- The integerNum of cycles in the Taylor series which will
//
//	be used to compute the value of e^exponent
//
// Return Values
// =============
//
// BigIntNum					- If successful, this function returns the value of e^exponent
//
//	as a type BigIntNum.
//
// error							- If an error is encountered an error message will be formatted
//
//	and returned as type 'error'. If the function successfully
//	completes the calculation, this value will be set to 'nil'.
func (bLog BigIntMathLogarithms) EPwrXFromMaclaurinSeries(
  exponent BigIntNum, nCycles int64) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathLogarithms.EPwrXFromMaclaurinSeries",
    "")

  if err != nil {
    return BigIntNum{}, err
  }
  // sum = 1 + x

  bigIntNumOne, err := new(BigIntNum).NewInt(1, 0)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "bigIntNumOne, err := new(BigIntNum).NewInt(1, 0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := exponent.GetNumericSeparatorsDto()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := exponent.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numSeps.SetDefaultsIfEmpty()

  sum, err := new(BigIntMathAdd).AddBigIntNums(bigIntNumOne, exponent)

  if err != nil {
    return BigIntNum{},
      fmt.Errorf("%v\n"+
        "Error returned by sum, err := BigIntMathAdd{}.AddBigIntNums()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())
  }

  nFactorial, err := new(BigIntNum).NewInt(1, 0)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nFactorial, err := new(BigIntNum).NewInt(1, 0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  x, err := exponent.CopyOut()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "x, err := exponent.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  for n := int64(2); n <= nCycles; n++ {

    bigIntNum64, err := new(BigIntNum).NewInt64(n, 0)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "bigIntNum64, err := new(BigIntNum).NewInt64(n, 0)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    nFactorial, err = new(BigIntMathMultiply).MultiplyBigIntNums(nFactorial, bigIntNum64)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "nFactorial, err = new(BigIntMathMultiply).MultiplyBigIntNums(nFactorial, bigIntNum64)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    x, err = new(BigIntMathMultiply).MultiplyBigIntNums(x, exponent)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "x, err = new(BigIntMathMultiply).MultiplyBigIntNums(x, exponent)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    nFactorialNumStr, err := nFactorial.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "nFactorialNumStr, err := nFactorial.GetNumStr()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    factor, err := new(BigIntMathDivide).BigIntNumFracQuotient(x, nFactorial, numSeps, 1200)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "factor, err := new(BigIntMathDivide).BigIntNumFracQuotient(x, nFactorial, numSeps, 1200)",
          ErrContext: fmt.Sprintf("nFactorial='%v'", nFactorialNumStr),
          ErrMessage: err.Error(),
        }
    }

    sum, err = new(BigIntMathAdd).AddBigIntNums(sum, factor)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "sum, err = new(BigIntMathAdd).AddBigIntNums(sum, factor)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    sumPrecisionUint, err := sum.GetPrecisionUint()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "sumPrecisionUint, err := sum.GetPrecisionUint()",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }

    if sumPrecisionUint > 500 {

      err = sum.RoundToDecPlace(500)

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "err = sum.RoundToDecPlace(500)",
            ErrContext: "",
            ErrMessage: err.Error(),
          }
      }
    }
  }

  return sum, nil
}

func (bLog BigIntMathLogarithms) EPwrXFromTaylorSeries(
  exponent BigIntNum, a BigIntNum, nCycles int64) (BigIntNum, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathLogarithms.EPwrXFromTaylorSeries",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  err = exponent.IsValid(ePrefix.XCpy("Validating exponent").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = exponent.IsValid(ePrefix)",
        ErrContext: "Error: Input parameter 'exponent' is invalid!\n" +
          "'exponent' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := exponent.GetNumericSeparatorsDto()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := exponent.GetNumericSeparatorsDto()",
        ErrContext: "Error extracting numeric separators from 'exponent'.",
        ErrMessage: err.Error(),
      }
  }

  numSeps.SetDefaultsIfEmpty()

  err = a.IsValid(ePrefix.XCpy("Validating 'a'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = a.IsValid(ePrefix)",
        ErrContext: "Error: Input parameter 'a' is invalid!\n" +
          "'a' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  e := eulersNumber1k.GetBigIntNum()

  eIsZero, err := e.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "eIsZero, err := e.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if eIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "eulersNumber1k == 0",
        ErrMessage: "Error: eulersNumber1k is ZERO!",
      }
  }

  aValue, err := a.GetUInt()

  if err != nil {

    return BigIntNum{},
      fmt.Errorf("%v\n"+
        "Error returned by aValue, err := a.GetUInt()\n"+
        "Error= %v\n",
        ePrefix,
        err.Error())

  }

  internalMaxPrecision := uint(20000)

  outputMaxPrecision := uint(1100)

  eIntegerValue, err := e.GetIntegerValue()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "eIntegerValue, err := e.GetIntegerValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  ePrecisionUint, err := e.GetPrecisionUint()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "ePrecisionUint, err := e.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  ePwrBigInt, ePwrBigIntPrecision, err :=
    new(BigIntMathPower).BigIntPwrIteration(
      eIntegerValue,
      ePrecisionUint,
      aValue,
      internalMaxPrecision,
      outputMaxPrecision)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "ePwrBigInt, ePwrBigIntPrecision, err :=\n" +
          "  new(BigIntMathPower).BigIntPwrIteration(\n" +
          "   eIntegerValue, ePrecisionUint, aValue, internalMaxPrecision, outputMaxPrecision)",
        ErrContext: fmt.Sprintf("eIntegerValue= '%v'\n"+
          "  ePrecisionUint= '%v'\n"+
          "  aValue= '%v'\n"+
          "  internalMaxPrecision= '%v'\n"+
          "  outputMaxPrecision= '%v'\n",
          eIntegerValue, ePrecisionUint, aValue, internalMaxPrecision, outputMaxPrecision),
        ErrMessage: err.Error(),
      }
  }

  eToPwr, err := new(BigIntNum).NewBigInt(ePwrBigInt, ePwrBigIntPrecision)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "eToPwr, err := new(BigIntNum).NewBigInt(ePwrBigInt, ePwrBigIntPrecision)",
        ErrContext: fmt.Sprintf("ePwrBigInt = '%v'\n"+
          "ePwrBigIntPrecision = '%v'\n",
          ePwrBigInt.Text(10), ePwrBigIntPrecision),
        ErrMessage: err.Error(),
      }
  }

  // eToPwr, err := BigIntMathPower{}.BigIntNumPwr(e, a, 500)

  //eToPwr := BigIntMathMultiply{}.MultiplyBigIntNums(e.CopyOut(), e.CopyOut())

  // sum = 0
  sum, err := new(BigIntNum).NewInt(0, 0)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "sum, err := new(BigIntNum).NewInt(0, 0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  x, err := exponent.CopyOut()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "x, err := exponent.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xNumStr, err := x.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "xNumStr, err := x.GetNumStr()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  aNumStr, err := a.GetNumStr()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "aNumStr, err := a.GetNumStr()\n",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xMinusA, err := new(BigIntMathSubtract).SubtractBigIntNums(x, a)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "xMinusA, err := new(BigIntMathSubtract).SubtractBigIntNums(x, a)",
        ErrContext: fmt.Sprintf("x= '%v'\n"+
          "a= '%v'",
          xNumStr, aNumStr),
        ErrMessage: err.Error(),
      }
  }

  xMinusANth, err := new(BigIntNum).NewInt(0, 0)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "xMinusANth, err := new(BigIntNum).NewInt(0, 0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  nFact, err := new(BigIntNum).NewInt(0, 0)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "nFact, err := new(BigIntNum).NewInt(0, 0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  for n := int64(0); n <= nCycles; n++ {

    if n == 0 {

      xMinusANth, err = new(BigIntNum).NewInt(1, 0)

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "xMinusANth, err = new(BigIntNum).NewInt(1, 0)",
            ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
              "n == 0", n),
            ErrMessage: err.Error(),
          }
      }

      nFact, err = new(BigIntNum).NewInt(1, 0)

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "nFact, err = new(BigIntNum).NewInt(1, 0)",
            ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
              "n == 0", n),
            ErrMessage: err.Error(),
          }
      }

    } else if n == 1 {

      xMinusANth, err = xMinusA.CopyOut()

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "xMinusANth, err = xMinusA.CopyOut()",
            ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
              "n == '%d'", n, n),
            ErrMessage: err.Error(),
          }
      }

      nFact, err = new(BigIntNum).NewInt64(n, 0)

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "nFact, err = new(BigIntNum).NewInt64(n, 0)",
            ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
              "n == '%d'", n, n),
            ErrMessage: err.Error(),
          }
      }

    } else {

      xMinusANth, err = new(BigIntMathMultiply).MultiplyBigIntNums(xMinusANth, xMinusA)

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix: ePrefix.String(),
            ReturnFunc: "xMinusANth, err = new(BigIntMathMultiply).\n" +
              "  MultiplyBigIntNums(xMinusANth, xMinusA)",
            ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
              "n == '%d'", n, n),
            ErrMessage: err.Error(),
          }
      }

      bINumN, err := new(BigIntNum).NewInt64(n, 0)

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "bINumN, err := new(BigIntNum).NewInt64(n, 0)",
            ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
              "n == '%d'", n, n),
            ErrMessage: err.Error(),
          }
      }

      nFact, err = new(BigIntMathMultiply).MultiplyBigIntNums(nFact, bINumN)

      if err != nil {

        return BigIntNum{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "nFact, err = new(BigIntMathMultiply).MultiplyBigIntNums(nFact, bINumN)",
            ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
              "n == '%d'", n, n),
            ErrMessage: err.Error(),
          }
      }

    }

    eToPwrNumStr, err := eToPwr.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "eToPwrNumStr, err := eToPwr.GetNumStr()",
          ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
            "n == '%d'", n, n),
          ErrMessage: err.Error(),
        }
    }

    nFactNumStr, err := nFact.GetNumStr()

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
            "n == '%d'", n, n),
          ErrMessage: err.Error(),
        }
    }

    factor1, err := new(BigIntMathDivide).BigIntNumFracQuotient(eToPwr, nFact, numSeps, 501)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "factor1, err := new(BigIntMathDivide).\n" +
            "  BigIntNumFracQuotient(eToPwr, nFact, numSeps,501)",
          ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
            "eToPwr = '%v'\n"+
            "nFact= '%v'\n"+
            "maxPrecision= 501", n, eToPwrNumStr, nFactNumStr),
          ErrMessage: err.Error(),
        }
    }

    factor2, err := new(BigIntMathMultiply).MultiplyBigIntNums(factor1, xMinusANth)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "factor2, err := new(BigIntMathMultiply).\n" +
            "  MultiplyBigIntNums(factor1, xMinusANth)",
          ErrContext: fmt.Sprintf("Cycle No= '%d'", n),
          ErrMessage: err.Error(),
        }
    }

    sum, err = new(BigIntMathAdd).AddBigIntNums(sum, factor2)

    if err != nil {

      return BigIntNum{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "sum, err = new(BigIntMathAdd).\n" +
            "  AddBigIntNums(sum, factor2)",
          ErrContext: fmt.Sprintf("Cycle No= '%d'", n),
          ErrMessage: err.Error(),
        }
    }

  }

  err = sum.RoundToDecPlace(500)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = sum.RoundToDecPlace(500)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = sum.IsValid(ePrefix.XCpy("Validating final result 'sum'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = sum.IsValid(ePrefix)",
        ErrContext: "Error: Final result 'sum' is invalid!\n" +
          "'sum' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return sum, nil
}

// EPwrXFromTaylorSeriesFixedDecimal
//
//	Returns 'exponentX'
func (bLog BigIntMathLogarithms) EPwrXFromTaylorSeriesFixedDecimal(
  exponentX BigIntFixedDecimal,
  a uint,
  nCycles uint) (BigIntFixedDecimal, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathLogarithms.EPwrXFromTaylorSeriesFixedDecimal",
    "")

  if err != nil {
    return BigIntFixedDecimal{}, err
  }

  err = exponentX.IsValid(ePrefix.XCpy("Validating exponentX").String())

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = exponentX.IsValid(ePrefix)",
        ErrContext: "Error: Input parameter 'exponentX' is invalid!\n" +
          "'exponentX' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := exponentX.GetNumericSeparatorsDto()

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := exponentX.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  numSeps.SetDefaultsIfEmpty()

  e := eulersNumber1k.GetFixedDecimal()

  eIsZero, err := e.IsZero()

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "eIsZero, err := e.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if eIsZero {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "A call to the following method produced and invalid result:\n" +
          "e := eulersNumber1k.GetFixedDecimal()\n" +
          "'e' was returned as a ZERO value. Therefore,\n" +
          "EulersNumber1050 Constant is ZERO!",
      }
  }

  internalMaxPrecision := uint(20000)

  outputMaxPrecision := uint(1500)

  eIntegerValue, err := e.GetIntegerValue()

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "eIntegerValue, err := e.GetIntegerValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  ePrecisionUint, err := e.GetPrecisionUint()

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "ePrecisionUint, err := e.GetPrecisionUint()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  ePwrBigInt, ePwrBigIntPrecision, err :=
    new(BigIntMathPower).BigIntPwrIteration(
      eIntegerValue,
      ePrecisionUint,
      a,
      internalMaxPrecision,
      outputMaxPrecision)

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "ePwrBigInt, ePwrBigIntPrecision, err :=\n" +
          "  new(BigIntMathPower).BigIntPwrIteration(\n" +
          "   eIntegerValue, ePrecisionUint, a, internalMaxPrecision, outputMaxPrecision)",
        ErrContext: fmt.Sprintf("eIntegerValue= '%v'\n"+
          "  ePrecisionUint= '%v'\n"+
          "  a= '%v'\n"+
          "  internalMaxPrecision= '%v'\n"+
          "  outputMaxPrecision= '%v'\n",
          eIntegerValue, ePrecisionUint, a, internalMaxPrecision, outputMaxPrecision),
        ErrMessage: err.Error(),
      }
  }

  eToPwr, err := new(BigIntFixedDecimal).New(ePwrBigInt, ePwrBigIntPrecision)

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "eToPwr, err := new(BigIntFixedDecimal).New(ePwrBigInt, ePwrBigIntPrecision)",
        ErrContext: fmt.Sprintf("ePwrBigInt= '%v'\n"+
          "ePwrBigIntPrecision= '%v'", ePwrBigInt, ePwrBigIntPrecision),
        ErrMessage: err.Error(),
      }
  }

  fixedDecA, err := new(BigIntFixedDecimal).New(big.NewInt(int64(a)), 0)

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "fixedDecA, err := new(BigIntFixedDecimal).\n" +
          "New(big.NewInt(int64(a)), 0)",
        ErrContext: fmt.Sprintf("a= '%v'\n", a),
        ErrMessage: err.Error(),
      }
  }

  // sum = 0
  sum := new(BigIntFixedDecimal).NewZero(0)

  xNum, err := exponentX.CopyOut()

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "xNum, err := exponentX.CopyOut()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xMinusA, err := new(BigIntMathSubtract).FixedDecimalSubtract(xNum, fixedDecA)

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "xMinusA, err := new(BigIntMathSubtract).\n" +
          "FixedDecimalSubtract(xNum, fixedDecA)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xMinusANth := new(BigIntFixedDecimal).NewZero(0)

  nFact := new(BigIntFixedDecimal).NewZero(0)

  for n := uint(0); n < nCycles; n++ {

    if n == 0 {

      xMinusANth = new(BigIntFixedDecimal).NewInt(1, 0)

      nFact = new(BigIntFixedDecimal).NewInt(1, 0)

    } else if n == 1 {

      xMinusANth, err = xMinusA.CopyOut()

      if err != nil {

        return BigIntFixedDecimal{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "xMinusANth, err = xMinusA.CopyOut()",
            ErrContext: fmt.Sprintf("n= '%v'", n),
            ErrMessage: err.Error(),
          }
      }

      nFact = new(BigIntFixedDecimal).NewUInt(n, 0)

    } else {

      xMinusANth, err = new(BigIntMathMultiply).FixedDecimalMultiply(xMinusANth, xMinusA)

      if err != nil {

        return BigIntFixedDecimal{},
          &FuncReturnError{
            ErrPrefix:  ePrefix.String(),
            ReturnFunc: "xMinusANth, err = new(BigIntMathMultiply).FixedDecimalMultiply(xMinusANth, xMinusA)",
            ErrContext: fmt.Sprintf("n= '%v'", n),
            ErrMessage: err.Error(),
          }
      }

      bIFixdec, err := new(BigIntFixedDecimal).New(big.NewInt(int64(n)), 0, numSeps)

      if err != nil {

        return BigIntFixedDecimal{},
          &FuncReturnError{
            ErrPrefix: ePrefix.String(),
            ReturnFunc: "bIFixdec, err := new(BigIntFixedDecimal).\n" +
              "  New(big.NewInt(int64(n)), 0, numSeps)",
            ErrContext: fmt.Sprintf("n= '%v'", n),
            ErrMessage: err.Error(),
          }
      }

      nFact, err = new(BigIntMathMultiply).FixedDecimalMultiply(
        nFact, bIFixdec)

      if err != nil {

        return BigIntFixedDecimal{},
          &FuncReturnError{
            ErrPrefix: ePrefix.String(),
            ReturnFunc: "nFact, err = new(BigIntMathMultiply).\n" +
              "FixedDecimalMultiply(nFact, bIFixdec)",
            ErrContext: fmt.Sprintf("n= '%v'", n),
            ErrMessage: err.Error(),
          }
      }

    }

    eToPwrNumStr, err := eToPwr.GetNumStr()

    if err != nil {

      return BigIntFixedDecimal{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "eToPwrNumStr, err := eToPwr.GetNumStr()",
          ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
            "n == '%d'", n, n),
          ErrMessage: err.Error(),
        }
    }

    nFactNumStr, err := nFact.GetNumStr()

    if err != nil {

      return BigIntFixedDecimal{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "",
          ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
            "n == '%d'", n, n),
          ErrMessage: err.Error(),
        }
    }

    factor1, err := new(BigIntMathDivide).FixedDecimalFracQuotient(eToPwr, nFact, numSeps, 501)

    if err != nil {

      return BigIntFixedDecimal{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "factor1, err := new(BigIntMathDivide).\n" +
            "FixedDecimalFracQuotient(eToPwr, nFact, numSeps,501)",
          ErrContext: fmt.Sprintf("Cycle No= '%d'\n"+
            "eToPwr = '%v'\n"+
            "nFact= '%v'\n"+
            "maxPrecision= 501", n, eToPwrNumStr, nFactNumStr),
          ErrMessage: err.Error(),
        }
    }

    factor2, err := new(BigIntMathMultiply).FixedDecimalMultiply(factor1, xMinusANth)

    if err != nil {

      return BigIntFixedDecimal{},
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "factor2, err := new(BigIntMathMultiply).\n" +
            "  FixedDecimalMultiply(factor1, xMinusANth)",
          ErrContext: fmt.Sprintf("Cycle No= '%d'", n),
          ErrMessage: err.Error(),
        }
    }

    sum, err = new(BigIntMathAdd).FixedDecimalAdd(sum, factor2)

    if err != nil {

      return BigIntFixedDecimal{},
        &FuncReturnError{
          ErrPrefix:  ePrefix.String(),
          ReturnFunc: "sum, err = new(BigIntMathAdd).FixedDecimalAdd(sum, factor2)",
          ErrContext: fmt.Sprintf("Cycle No= '%d'", n),
          ErrMessage: err.Error(),
        }
    }

  }

  err = sum.RoundToDecPlace(500)

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = sum.RoundToDecPlace(500)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = sum.IsValid(ePrefix.XCpy("Validating final result 'sum'").String())

  if err != nil {

    return BigIntFixedDecimal{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = sum.IsValid(ePrefix)",
        ErrContext: "Error: Final result 'sum' is invalid!\n" +
          "'sum' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  return sum, err
}

func (bLog BigIntMathLogarithms) EPwrXFromTaylorSeriesBigInt(
  exponentX *big.Int,
  exponentXPrecision *big.Int,
  a *big.Int,
  nCycles *big.Int,
  internalMaxPrecision *big.Int,
  maxPrecision *big.Int) (xNum *big.Int, xNumPrecision *big.Int, err error) {

  xNum = big.NewInt(0)
  xNumPrecision = big.NewInt(0)

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "BigIntMathLogarithms.EPwrXFromTaylorSeriesBigInt",
    "")

  if err != nil {
    return xNum, xNumPrecision, err
  }

  if exponentX == nil {

    return xNum, xNumPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'exponentX'",
      }
  }

  if exponentXPrecision == nil {

    return xNum, xNumPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'exponentXPrecision'",
      }
  }

  if a == nil {

    return xNum, xNumPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'a'",
      }
  }

  if nCycles == nil {

    return xNum, xNumPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'nCycles'",
      }
  }

  if internalMaxPrecision == nil {

    return xNum, xNumPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'internalMaxPrecision'",
      }
  }

  if maxPrecision == nil {

    return xNum, xNumPrecision,
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'maxPrecision'",
      }
  }

  e, ePrecision := eulersNumber1k.GetBigIntPrecision()

  bigZero := big.NewInt(0)

  if e.Cmp(bigZero) == 0 {

    return xNum, xNumPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: 'e' returned from eulersNumber1k.GetBigIntPrecision()\n" +
          "has ZERO value. Therefore, EulersNumber1050 Constant is ZERO!",
      }
  }

  eToPwr, eToPwrPrecision, err := new(BigIntMathPower).BigIntegerPwrIteration(
    e,
    ePrecision,
    a,
    internalMaxPrecision,
    maxPrecision)

  if err != nil {

    return xNum, xNumPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "eToPwr, eToPwrPrecision, err := new(BigIntMathPower).\n" +
          "  BigIntegerPwrIteration(e, ePrecision, a, internalMaxPrecision, maxPrecision)",
        ErrContext: fmt.Sprintf("e= '%v'\n"+
          "ePrecision= '%v'\n"+
          "a= '%v'\n"+
          "internalMaxPrecision= '%v'\n"+
          "maxPrecision= '%v'\n",
          e.Text(10), ePrecision, a.Text(10), internalMaxPrecision, maxPrecision),
        ErrMessage: err.Error(),
      }
  }

  xNum = big.NewInt(0)

  xNumPrecision = big.NewInt(0)

  xMinusA, xMinusAPrecision, err := new(BigIntMathSubtract).
    BigIntSubtract(exponentX, exponentXPrecision, a, big.NewInt(0))

  if err != nil {

    return xNum, xNumPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "xMinusA, xMinusAPrecision, err := new(BigIntMathSubtract).\n" +
          "  BigIntSubtract(exponentX, exponentXPrecision, a, big.NewInt(0))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xMinusANth := big.NewInt(0)

  xMinusANthPrecision := big.NewInt(0)

  nFact := big.NewInt(0)

  nFactPrecision := big.NewInt(0)

  bigOne := big.NewInt(1)

  factor1 := big.NewInt(0)

  factor1Precision := big.NewInt(0)

  factor2 := big.NewInt(0)

  factor2Precision := big.NewInt(0)

  for n := big.NewInt(0); n.Cmp(nCycles) == -1; n.Add(n, bigOne) {

    if n.Cmp(bigZero) == 0 {

      xMinusANth = big.NewInt(1)

      xMinusANthPrecision = big.NewInt(0)

    } else if n.Cmp(bigOne) == 0 {

      xMinusANth = big.NewInt(0).Set(xMinusA)

      xMinusANth = big.NewInt(0).Set(xMinusAPrecision)

    } else {

      xMinusANth, xMinusANthPrecision, err =
        new(BigIntMathMultiply).BigIntMultiply(
          xMinusANth,
          xMinusANthPrecision,
          xMinusA,
          xMinusAPrecision)

      if err != nil {

        return big.NewInt(0), big.NewInt(0),
          &FuncReturnError{
            ErrPrefix: ePrefix.String(),
            ReturnFunc: "xMinusANth, xMinusANthPrecision, err =\n" +
              "  new(BigIntMathMultiply).BigIntMultiply(\n" +
              "   xMinusANth, xMinusANthPrecision, xMinusA, xMinusAPrecision)",
            ErrContext: "ID-2\n",
            ErrMessage: err.Error(),
          }
      }

      nFact, nFactPrecision, err =
        new(BigIntMathMultiply).BigIntMultiply(
          nFact,
          nFactPrecision,
          n,
          big.NewInt(0))

      if err != nil {

        return big.NewInt(0), big.NewInt(0),
          &FuncReturnError{
            ErrPrefix: ePrefix.String(),
            ReturnFunc: "nFact, nFactPrecision, err =\n" +
              "  new(BigIntMathMultiply).BigIntMultiply(\n" +
              "   nFact, nFactPrecision, nFact, nFactPrecision, nFactPrecision)",
            ErrContext: "ID-3",
            ErrMessage: err.Error(),
          }
      }
    }

    factor1, factor1Precision, err =
      new(BigIntMathDivide).BigIntFracQuotient(
        eToPwr,
        eToPwrPrecision,
        nFact,
        nFactPrecision,
        internalMaxPrecision)

    if err != nil {

      return big.NewInt(0), big.NewInt(0),
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "factor1, factor1Precision, err =\n" +
            "  new(BigIntMathDivide).BigIntFracQuotient(\n" +
            "   eToPwr, eToPwrPrecision, nFact, nFactPrecision, internalMaxPrecision)",
          ErrContext: "ID-4",
          ErrMessage: err.Error(),
        }
    }

    factor2, factor2Precision, err =
      new(BigIntMathMultiply).BigIntMultiply(
        factor1,
        factor1Precision,
        xMinusANth,
        xMinusANthPrecision)

    if err != nil {

      return big.NewInt(0), big.NewInt(0),
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "factor2, factor2Precision, err =\n" +
            "  new(BigIntMathMultiply).BigIntMultiply(\n" +
            "   factor1, factor1Precision, xMinusANth, xMinusANthPrecision)",
          ErrContext: "ID-5",
          ErrMessage: err.Error(),
        }
    }

    xNum, xNumPrecision, err = new(BigIntMathAdd).BigIntAdd(
      xNum, xNumPrecision, factor2, factor2Precision)

    if err != nil {

      return big.NewInt(0), big.NewInt(0),
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "xNum, xNumPrecision, err = \n" +
            "  new(BigIntMathAdd).BigIntAdd(\n" +
            "   xNum, xNumPrecision, factor2, factor2Precision)",
          ErrContext: "ID-6",
          ErrMessage: err.Error(),
        }
    }
  }

  xNum, xNumPrecision, err =
    new(BigIntMath).RoundToMaxPrecision(
      xNum, xNumPrecision, maxPrecision, true)

  if err != nil {

    return big.NewInt(0), big.NewInt(0),
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "xNum, xNumPrecision, err = \n" +
          "  new(BigIntMath).RoundToMaxPrecision(\n" +
          "   xNum, xNumPrecision, maxPrecision, true)",
        ErrContext: "ID-7",
        ErrMessage: err.Error(),
      }
  }

  return xNum, xNumPrecision, err
}

func (bLog BigIntMathLogarithms) BigIntNumNatLogOfX(
  xNum BigIntNum,
  maxPrecision uint) (lnOfX BigIntNum, err error) {

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CopyIn",
    "")

  if err != nil {
    return BigIntNum{}, err
  }

  err = xNum.IsValid(ePrefix.XCpy("Validating 'xNum'").String())

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = xNum.IsValid(ePrefix)",
        ErrContext: "Error: Input parameter 'xNum' is invalid!\n" +
          "'xNum' FAILED validation tests.",
        ErrMessage: err.Error(),
      }
  }

  lnOfX, err = new(BigIntNum).NewZero(0)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "lnOfX, err = new(BigIntNum).NewZero(0)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  xNumIsZero, err := xNum.IsZero()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "xNumIsZero, err := xNum.IsZero()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if xNumIsZero {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "",
        ErrMessage: "Error: Input parameter 'xNum' is zero!\n",
      }
  }

  biXNumIntValue, err := xNum.GetIntegerValue()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "biXNumIntValue, err := xNum.GetIntegerValue()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  biXNumPrecision, err := xNum.GetPrecisionBigInt()

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "biXNumPrecision, err := xNum.GetPrecisionBigInt()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  maxFinalPrecision := big.NewInt(0).SetUint64(uint64(maxPrecision))

  /*
  		Assume maxPrecision = 50
  	  m = 90
  		agMeanMaxOutputPrecision 		= 190
  		agMeanMaxInternalPrecision 	= 950
  	  s4DivPrecisionMaxInternalPrecision = 2950
  		piDivideMaxPrecision = 6950

  */

  m, _, err := new(BigIntMath).RoundToMaxPrecision(
    big.NewInt(0).Mul(maxFinalPrecision, big.NewInt(18)),
    big.NewInt(1),
    big.NewInt(0),
    true)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "m, _, err := new(BigIntMath).RoundToMaxPrecision(\n" +
          "  big.NewInt(0).Mul(maxFinalPrecision, big.NewInt(18)),\n" +
          "  big.NewInt(1), big.NewInt(0), true)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  agMeanMaxOutputPrecision :=
    big.NewInt(0).Add(maxFinalPrecision, big.NewInt(100))

  agMeanMaxInternalPrecision :=
    big.NewInt(0).Mul(agMeanMaxOutputPrecision, big.NewInt(5))

  s4DivPrecisionMaxInternalPrecision :=
    big.NewInt(0).Add(agMeanMaxInternalPrecision, big.NewInt(2000))

  piDivideMaxPrecision :=
    big.NewInt(0).Add(s4DivPrecisionMaxInternalPrecision, big.NewInt(4000))

  biResult, biResultPrecision, err :=
    BigIntMathLogarithms{}.SaskiKanadaNatLogOfX(
      biXNumIntValue,
      biXNumPrecision,
      m,
      s4DivPrecisionMaxInternalPrecision,
      agMeanMaxInternalPrecision,
      agMeanMaxOutputPrecision,
      piDivideMaxPrecision,
      maxFinalPrecision)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "biResult, biResultPrecision, err :=\n" +
          "  BigIntMathLogarithms{}.SaskiKanadaNatLogOfX(\n" +
          "  biXNumIntValue, biXNumPrecision, m, s4DivPrecisionMaxInternalPrecision,\n" +
          "  agMeanMaxInternalPrecision, agMeanMaxOutputPrecision, piDivideMaxPrecision,\n" +
          "  maxFinalPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  lnOfX, err = new(BigIntNum).NewBigIntBigPrecision(biResult, biResultPrecision)

  if err != nil {

    return BigIntNum{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "lnOfX, err = new(BigIntNum).NewBigIntBigPrecision(\n" +
          "  biResult, biResultPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return lnOfX, nil
}

// SaskiKanadaNatLogOfX
// https://en.wikipedia.org/wiki/Natural_logarithm
// Sasaki, T.; Kanada, Y. (1982).
// "Practically fast multiple-precision evaluation of log(x)".
// Journal of Information Processing. 5 (4): 247–250. Retrieved 2011-03-30.
func (bLog BigIntMathLogarithms) SaskiKanadaNatLogOfX(
  xNum *big.Int,
  xNumPrecision *big.Int,
  m *big.Int,
  s4DivPrecisionMaxInternalPrecision *big.Int,
  agMeanMaxInternalPrecision *big.Int,
  agMeanMaxOutputPrecision *big.Int,
  piDivideMaxPrecision *big.Int,
  maxFinalResultPrecision *big.Int) (lnOfX *big.Int, lnOfXPrecision *big.Int, err error) {

  lnOfX = big.NewInt(0)
  lnOfXPrecision = big.NewInt(0)

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CopyIn",
    "")

  if err != nil {
    return big.NewInt(0), big.NewInt(0), err
  }

  if xNum == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'xNum'",
      }
  }

  if xNumPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'xNumPrecision'",
      }
  }

  if m == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'m'",
      }
  }

  if s4DivPrecisionMaxInternalPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'s4DivPrecisionMaxInternalPrecision'",
      }
  }

  if agMeanMaxInternalPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'agMeanMaxInternalPrecision'",
      }
  }

  if agMeanMaxOutputPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'agMeanMaxOutputPrecision'",
      }
  }

  if piDivideMaxPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'piDivideMaxPrecision'",
      }
  }

  if maxFinalResultPrecision == nil {

    return big.NewInt(0), big.NewInt(0),
      &InputPtrNilError{
        ErrPrefix:     ePrefix.String(),
        ParameterName: "'maxFinalResultPrecision'",
      }
  }

  bigZero := big.NewInt(0)

  if m.Cmp(bigZero) == -1 {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("m= '%v'", m.Text(10)),
        ErrMessage: "Error: Input parameter 'm' is INVALID!\n" +
          "'m' has a negative value.",
      }
  }

  if xNumPrecision.Cmp(bigZero) == -1 {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("xNumPrecision= '%v'", xNumPrecision.Text(10)),
        ErrMessage: "Error: Input parameter 'xNumPrecision' is INVALID!\n" +
          "'xNumPrecision' has a negative value.",
      }
  }

  if piDivideMaxPrecision.Cmp(bigZero) == -1 {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("piDivideMaxPrecision= '%v'", piDivideMaxPrecision.Text(10)),
        ErrMessage: "Error: Input parameter 'piDivideMaxPrecision' is INVALID!\n" +
          "'piDivideMaxPrecision' has a negative value.",
      }
  }

  if maxFinalResultPrecision.Cmp(bigZero) == -1 {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: fmt.Sprintf("maxFinalResultPrecision= '%v'", maxFinalResultPrecision.Text(10)),
        ErrMessage: "Error: Input parameter 'maxFinalResultPrecision' is INVALID!\n" +
          "'maxFinalResultPrecision' has a negative value.",
      }
  }

  sFactor2Tom := big.NewInt(0).Exp(big.NewInt(2), m, nil)

  s := big.NewInt(0).Mul(xNum, sFactor2Tom)

  sPrecision := big.NewInt(0).Set(xNumPrecision)

  s, sPrecision, err = new(BigIntMath).RoundToMaxPrecision(
    s,
    sPrecision,
    big.NewInt(0).Add(s4DivPrecisionMaxInternalPrecision, big.NewInt(20)),
    true)

  if err != nil {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "s, sPrecision, err = new(BigIntMath).\n" +
          "RoundToMaxPrecision(s, sPrecision,\n" +
          "  big.NewInt(0).Add(s4DivPrecisionMaxInternalPrecision, big.NewInt(20)),\n" +
          "  true)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  fourDivS, fourDivSPrecision, err :=
    new(BigIntMathDivide).BigIntFracQuotient(
      big.NewInt(4), big.NewInt(0), s, sPrecision, s4DivPrecisionMaxInternalPrecision)

  if err != nil {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "fourDivS, fourDivSPrecision, err :=\n" +
          "new(BigIntMathDivide).BigIntFracQuotient(\n" +
          "  big.NewInt(4), big.NewInt(0), s, sPrecision,\n" +
          "  s4DivPrecisionMaxInternalPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  // Uses maxInternal Precision
  agMean, agMeanPrecision, _, _, _, err :=
    new(BigIntMath).ArithmeticGeometricMean(
      big.NewInt(1),
      big.NewInt(0),
      fourDivS,
      fourDivSPrecision,
      agMeanMaxInternalPrecision,
      agMeanMaxOutputPrecision)

  if err != nil {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "agMean, agMeanPrecision, _, _, _, err :=\n" +
          " new(BigIntMath).ArithmeticGeometricMean(\n" +
          "  big.NewInt(1), big.NewInt(0), fourDivS, fourDivSPrecision,\n" +
          "  agMeanMaxInternalPrecision, agMeanMaxOutputPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  denomFactorM := big.NewInt(0).Mul(big.NewInt(2), agMean)

  denomFactorMPrecision := big.NewInt(0).Set(agMeanPrecision)

  // Uses MaxInternal Precision
  factor1, factor1Precision, err :=
    new(BigIntMathDivide).BigIntFracQuotient(
      piNumber20k.GetInteger(),
      piNumber20k.GetPrecisionBigInt(),
      denomFactorM,
      denomFactorMPrecision,
      piDivideMaxPrecision)

  if err != nil {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "factor1, factor1Precision, err :=\n" +
          " new(BigIntMathDivide).BigIntFracQuotient(\n" +
          "   piNumber20k.GetInteger(), piNumber20k.GetPrecisionBigInt(),\n" +
          "   denomFactorM, denomFactorMPrecision, piDivideMaxPrecision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  factor2 := big.NewInt(0).Mul(m, natLogTwo20k.GetInteger())

  factor2Precision := big.NewInt(0).Set(natLogTwo20k.GetPrecisionBigInt())

  lnOfX, lnOfXPrecision, err =
    new(BigIntMathSubtract).BigIntSubtract(
      factor1,
      factor1Precision,
      factor2,
      factor2Precision)

  if err != nil {

    return lnOfX, lnOfXPrecision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "lnOfX, lnOfXPrecision, err = \n" +
          " new(BigIntMathSubtract).BigIntSubtract(\n" +
          "  factor1, factor1Precision, factor2, factor2Precision)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  if lnOfXPrecision.Cmp(maxFinalResultPrecision) == 1 {

    lnOfX, lnOfXPrecision, err =
      new(BigIntMath).RoundToMaxPrecision(
        lnOfX,
        lnOfXPrecision,
        maxFinalResultPrecision,
        true)

    if err != nil {

      return lnOfX, lnOfXPrecision,
        &FuncReturnError{
          ErrPrefix: ePrefix.String(),
          ReturnFunc: "lnOfX, lnOfXPrecision, err =\n" +
            "lnOfX, lnOfXPrecision, maxFinalResultPrecision, true)",
          ErrContext: "",
          ErrMessage: err.Error(),
        }
    }
  }

  return lnOfX, lnOfXPrecision, nil
}
