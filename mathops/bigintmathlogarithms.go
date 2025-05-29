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
	base,
	basePrecision,
	xNum,
	xNumPrecision,
	maxPrecision *big.Int) (logResult,
	logResultPrecision *big.Int,
	err error) {

	logResult = big.NewInt(0)
	logResultPrecision = big.NewInt(0)
	err = nil
	ePrefix := "BigIntMathLogarithms.BigIntLogBaseOfX() "

	bigZero := big.NewInt(0)

	if base == nil {
		base = big.NewInt(0)
	}

	if base.Cmp(bigZero) != 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'base' is INVALID!\n"+
			"base='%v'\n",
			ePrefix,
			base.Text(10))

		return logResult, logResultPrecision, err
	}

	if basePrecision == nil {
		basePrecision = big.NewInt(0)
	}

	if basePrecision.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'basePrecision' is negative and INVALID!\n"+
			"basePrecision='%v'\n",
			ePrefix,
			basePrecision.Text(10))

		return logResult, logResultPrecision, err

	}

	if xNum == nil {
		xNum = big.NewInt(0)
	}

	if xNum.Cmp(bigZero) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'xNum' is INVALID!\n"+
			"xNum='%v'\n",
			ePrefix,
			xNum.Text(10))

		return logResult, logResultPrecision, err
	}

	if xNumPrecision == nil {
		xNumPrecision = big.NewInt(0)
	}

	if xNumPrecision.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'xNumPrecision' is LESS THAN ZERO!\n"+
			"xNumPrecision='%v'\n",
			ePrefix,
			xNumPrecision.Text(10))

		return logResult, logResultPrecision, err
	}

	if maxPrecision == nil {
		maxPrecision = big.NewInt(0)
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

	err = nil

	return logResult, logResultPrecision, err
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
	exponent, a BigIntNum, nCycles int64) (BigIntNum, error) {

	ePrefix := "BigIntMathLogarithms.EPwrXFromTaylorSeries() "

	e := eulersNumber1k.GetBigIntNum()

	if e.IsZero() {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error: eulersNumber1k is ZERO!\n",
				ePrefix)
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

	ePwrBigInt, ePwrBigIntPrecision :=
		BigIntMathPower{}.BigIntPwrIteration(
			e.GetIntegerValue(),
			e.GetPrecisionUint(),
			aValue,
			internalMaxPrecision,
			outputMaxPrecision)

	eToPwr, err := new(BigIntNum).NewBigInt(ePwrBigInt, ePwrBigIntPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "eToPwr, err := new(BigIntNum).NewBigInt(ePwrBigInt, ePwrBigIntPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// eToPwr, err := BigIntMathPower{}.Pwr(e, a, 500)

	//eToPwr := BigIntMathMultiply{}.MultiplyBigIntNums(e.CopyOut(), e.CopyOut())

	// sum = 0
	sum := BigIntNum{}.NewInt(0, 0)

	x := exponent.CopyOut()

	xMinusA := BigIntMathSubtract{}.SubtractBigIntNums(x, a)

	xMinusANth := BigIntNum{}.NewInt(0, 0)

	nFact := BigIntNum{}.NewInt(0, 0)

	for n := int64(0); n <= nCycles; n++ {

		if n == 0 {

			xMinusANth = BigIntNum{}.NewInt(1, 0)
			nFact = BigIntNum{}.NewInt(1, 0)

		} else if n == 1 {

			xMinusANth = xMinusA.CopyOut()
			nFact = BigIntNum{}.NewInt64(n, 0)

		} else {

			xMinusANth = BigIntMathMultiply{}.MultiplyBigIntNums(xMinusANth, xMinusA)

			nFact = BigIntMathMultiply{}.MultiplyBigIntNums(nFact, BigIntNum{}.NewInt64(n, 0))

		}

		factor1, err := BigIntMathDivide{}.BigIntNumFracQuotient(eToPwr, nFact, 500)

		if err != nil {

			return BigIntNum{}.NewZero(0),
				fmt.Errorf("%v\n"+
					"Error returned by factor1, err := BigIntMathDivide{}.BigIntNumFracQuotient(eToPwr, nFact, 500)\n"+
					"eToPwr='%v'\nnFact='%v'\nError= %v\n",
					ePrefix,
					eToPwr.GetNumStr(),
					nFact.GetNumStr(),
					err.Error())
		}

		factor2 := BigIntMathMultiply{}.MultiplyBigIntNums(factor1, xMinusANth)

		sum, err = BigIntMathAdd{}.AddBigIntNums(sum, factor2)

		if err != nil {

			return BigIntNum{}.NewZero(0),
				fmt.Errorf("%v\n"+
					"Error returned by sum, err = BigIntMathAdd{}.AddBigIntNums(sum, factor2)\n"+
					"Error= %v\n",
					ePrefix,
					err.Error())
		}

	}

	sum.RoundToDecPlace(500)

	return sum, nil

}

// EPwrXFromTaylorSeriesFixedDecimal
// Returns 'exponentX'
func (bLog BigIntMathLogarithms) EPwrXFromTaylorSeriesFixedDecimal(
	exponentX BigIntFixedDecimal,
	a, nCycles uint) (BigIntFixedDecimal, error) {

	ePrefix := "BigIntMathLogarithms.EPwrXFromTaylorSeries()"

	e := eulersNumber1k.GetFixedDecimal()

	if e.IsZero() {
		return BigIntFixedDecimal{},
			fmt.Errorf("%v\n"+
				"A call to the following method produced and invalid result:\n"+
				"e := eulersNumber1k.GetFixedDecimal()\n"+
				"'e' was returned as a ZERO value. Therefore,\n"+
				"EulersNumber1050 Constant is ZERO!\n",
				ePrefix)
	}

	internalMaxPrecision := uint(20000)

	outputMaxPrecision := uint(1500)

	ePwrBigInt, ePwrBigIntPrecision :=
		BigIntMathPower{}.BigIntPwrIteration(
			e.GetIntegerValue(),
			e.GetPrecisionUint(),
			a,
			internalMaxPrecision,
			outputMaxPrecision)

	eToPwr := new(BigIntFixedDecimal).New(ePwrBigInt, ePwrBigIntPrecision)

	fixedDecA := new(BigIntFixedDecimal).New(big.NewInt(int64(a)), 0)

	// sum = 0
	sum := new(BigIntFixedDecimal).NewZero(0)

	xNum := exponentX.CopyOut()

	xMinusA := BigIntMathSubtract{}.FixedDecimalSubtract(xNum, fixedDecA)

	xMinusANth := new(BigIntFixedDecimal).NewZero(0)

	nFact := new(BigIntFixedDecimal).NewZero(0)

	for n := uint(0); n < nCycles; n++ {

		if n == 0 {

			xMinusANth = new(BigIntFixedDecimal).NewInt(1, 0)
			nFact = new(BigIntFixedDecimal).NewInt(1, 0)

		} else if n == 1 {

			xMinusANth = xMinusA.CopyOut()
			nFact = new(BigIntFixedDecimal).NewUInt(n, 0)

		} else {

			xMinusANth = BigIntMathMultiply{}.FixedDecimalMultiply(xMinusANth, xMinusA)

			nFact =
				BigIntMathMultiply{}.FixedDecimalMultiply(
					nFact,
					new(BigIntFixedDecimal).New(big.NewInt(int64(n)), 0))

		}

		factor1, err := BigIntMathDivide{}.FixedDecimalFracQuotient(eToPwr, nFact, 500)

		if err != nil {
			return new(BigIntFixedDecimal).NewZero(0),
				fmt.Errorf("%v\n"+
					"Error returned by factor1, err := BigIntMathDivide{}.BigIntNumFracQuotient(eToPwr, nFact, 500)\n"+
					"\neToPwr='%v'\nnFact='%v'\nError= %v\n",
					ePrefix,
					eToPwr.GetNumStr(),
					nFact.GetNumStr(),
					err.Error())
		}

		factor2 := BigIntMathMultiply{}.FixedDecimalMultiply(factor1, xMinusANth)

		sum, err = BigIntMathAdd{}.FixedDecimalAdd(sum, factor2)

		if err != nil {
			err = fmt.Errorf("%v\n"+
				"Error returned by sum, err = BigIntMathAdd{}.FixedDecimalAdd(...)\n"+
				"Error='%v'\n",
				ePrefix,
				err.Error())
		}

	}

	err := sum.RoundToDecPlace(500)

	if err != nil {

		return sum,
			fmt.Errorf("%v\n"+
				"Error returned by err := sum.RoundToDecPlace(500)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return sum, err
}

func (bLog BigIntMathLogarithms) EPwrXFromTaylorSeriesBigInt(
	exponentX,
	exponentXPrecision,
	a,
	nCycles,
	internalMaxPrecision,
	maxPrecision *big.Int) (xNum, xNumPrecision *big.Int, err error) {

	xNum = big.NewInt(0)
	xNumPrecision = big.NewInt(0)
	err = nil

	ePrefix := "BigIntMathLogarithms.EPwrXFromTaylorSeries()"

	e, ePrecision := eulersNumber1k.GetBigIntPrecision()

	bigZero := big.NewInt(0)

	if e.Cmp(bigZero) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: 'e' returned from eulersNumber1k.GetBigIntPrecision()\n"+
			"has ZERO value. Therefore, EulersNumber1050 Constant is ZERO!\n",
			ePrefix)

		return xNum, xNumPrecision, err
	}

	eToPwr, eToPwrPrecision, errX :=
		BigIntMathPower{}.BigIntegerPwrIteration(
			e,
			ePrecision,
			a,
			internalMaxPrecision,
			maxPrecision)

	if errX != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by: \n"+
			" eToPwr, eToPwrPrecision, errX :=BigIntMathPower{}.BigIntegerPwrIteration()\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return xNum, xNumPrecision, err
	}

	xNum = big.NewInt(0)
	xNumPrecision = big.NewInt(0)

	xMinusA, xMinusAPrecision, errX :=
		BigIntMathSubtract{}.BigIntSubtract(exponentX, exponentXPrecision, a, big.NewInt(0))

	if errX != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by: \n"+
			"  xMinusA, xMinusAPrecision, errX := BigIntMathSubtract{}.BigIntSubtract(...)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return xNum, xNumPrecision, err
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

			xMinusANth, xMinusANthPrecision, errX =
				BigIntMathMultiply{}.BigIntMultiply(
					xMinusANth,
					xMinusANthPrecision,
					xMinusA,
					xMinusAPrecision)

			if errX != nil {

				xNum = big.NewInt(0)
				xNumPrecision = big.NewInt(0)

				err = fmt.Errorf("%v\n"+
					"Error returned by: \n"+
					"xMinusANth, xMinusANthPrecision, errX = BigIntMathMultiply{}.BigIntMultiply(...)\n"+
					"ID-1\n"+
					"Error= %v\n",
					ePrefix,
					errX.Error())

				return xNum, xNumPrecision, err
			}

			nFact, nFactPrecision, errX =
				BigIntMathMultiply{}.BigIntMultiply(
					nFact,
					nFactPrecision,
					n,
					big.NewInt(0))

			if errX != nil {

				err = fmt.Errorf("%v\n"+
					"Error returned by: \n"+
					"nFact, nFactPrecision, errX = BigIntMathMultiply{}.BigIntMultiply(...)\n"+
					"ID-2\n"+
					"Error= %v\n",
					ePrefix,
					errX.Error())

				return xNum, xNumPrecision, err
			}

		}

		factor1, factor1Precision, errX =
			BigIntMathDivide{}.BigIntFracQuotient(
				eToPwr,
				eToPwrPrecision,
				nFact,
				nFactPrecision,
				internalMaxPrecision)

		if errX != nil {
			xNum = big.NewInt(0)
			xNumPrecision = big.NewInt(0)

			err = fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" factor1, factor1Precision, errX = BigIntMathDivide{}.BigIntFracQuotient(...)\n"+
				"Error= %v\n", errX.Error())

			return xNum, xNumPrecision, err
		}

		factor2, factor2Precision, errX =
			BigIntMathMultiply{}.BigIntMultiply(
				factor1,
				factor1Precision,
				xMinusANth,
				xMinusANthPrecision)

		if errX != nil {
			xNum = big.NewInt(0)
			xNumPrecision = big.NewInt(0)
			err = fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" factor2, factor2Precision, errX = BigIntMathMultiply{}.BigIntMultiply(...)\n"+
				"Error%v\n",
				ePrefix,
				errX.Error())

			return xNum, xNumPrecision, err
		}

		xNum, xNumPrecision, errX = BigIntMathAdd{}.BigIntAdd(xNum, xNumPrecision, factor2, factor2Precision)

		if errX != nil {

			xNum = big.NewInt(0)
			xNumPrecision = big.NewInt(0)

			err = fmt.Errorf("%v\n"+
				"Error returned by \n"+
				" xNum, xNumPrecision, errX = BigIntMathAdd{}.BigIntAdd(...)\n"+
				"Error= %v\n",
				ePrefix,
				errX.Error())

			return xNum, xNumPrecision, err
		}

	}

	xNum, xNumPrecision, errX = BigIntMath{}.RoundToMaxPrecision(xNum, xNumPrecision, maxPrecision, true)

	if errX != nil {
		xNum = big.NewInt(0)
		xNumPrecision = big.NewInt(0)

		err = fmt.Errorf("%v\n"+
			"Error returned by:\n"+
			" xNum, xNumPrecision, errX = BigIntMath{}.RoundToMaxPrecision(...)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return xNum, xNumPrecision, err
	}

	err = nil
	return xNum, xNumPrecision, err
}

func (bLog BigIntMathLogarithms) BigIntNumNatLogOfX(
	xNum BigIntNum,
	maxPrecision uint) (lnOfX BigIntNum, err error) {

	ePrefix := "BigIntMathLogarithms.BigIntNumNatLogOfX() "
	lnOfX = BigIntNum{}.NewZero(0)
	err = nil

	errX := xNum.IsValid(ePrefix)

	if errX != nil {

		err = fmt.Errorf("%v\n"+
			"'xNum' is INVALID!\n"+
			"Error returned from errX := xNum.IsValid(ePrefix)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return lnOfX, err
	}

	if xNum.IsZero() {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'xNum' is ZERO!\n",
			ePrefix)

		return lnOfX, err
	}

	biXNum := xNum.GetIntegerValue()
	biXNumPrecision := xNum.GetPrecisionBigInt()
	maxFinalPrecision := big.NewInt(0).SetUint64(uint64(maxPrecision))

	/*
			Assume maxPrecision = 50
		  m = 90
			agMeanMaxOutputPrecision 		= 190
			agMeanMaxInternalPrecision 	= 950
		  s4DivPrecisionMaxInternalPrecision = 2950
			piDivideMaxPrecision = 6950

	*/

	m, _, err :=
		BigIntMath{}.RoundToMaxPrecision(
			big.NewInt(0).Mul(maxFinalPrecision, big.NewInt(18)),
			big.NewInt(1),
			big.NewInt(0),
			true)

	if err != nil {

		return lnOfX,
			fmt.Errorf("%v\n"+
				"Error returned by m, _, err := BigIntMath{}.RoundToMaxPrecision(...)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

	}

	agMeanMaxOutputPrecision :=
		big.NewInt(0).Add(maxFinalPrecision, big.NewInt(100))

	agMeanMaxInternalPrecision :=
		big.NewInt(0).Mul(agMeanMaxOutputPrecision, big.NewInt(5))

	s4DivPrecisionMaxInternalPrecision :=
		big.NewInt(0).Add(agMeanMaxInternalPrecision, big.NewInt(2000))

	piDivideMaxPrecision :=
		big.NewInt(0).Add(s4DivPrecisionMaxInternalPrecision, big.NewInt(4000))

	biResult, biResultPrecision, errX :=
		BigIntMathLogarithms{}.SaskiKanadaNatLogOfX(
			biXNum,
			biXNumPrecision,
			m,
			s4DivPrecisionMaxInternalPrecision,
			agMeanMaxInternalPrecision,
			agMeanMaxOutputPrecision,
			piDivideMaxPrecision,
			maxFinalPrecision)

	if errX != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by:\n"+
			" biResult, biResultPrecision, errX := BigIntMathLogarithms{}.SaskiKanadaNatLogOfX(...)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return lnOfX, err
	}

	lnOfX, errX = BigIntNum{}.NewBigIntBigPrecision(biResult, biResultPrecision)

	if errX != nil {

		lnOfX = BigIntNum{}.NewZero(0)

		err = fmt.Errorf("%v\n"+
			"Error returned by:\n"+
			" lnOfX, errX = BigIntNum{}.NewBigIntBigPrecision(biResult, biResultPrecision)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return lnOfX, err
	}

	err = nil

	return lnOfX, err
}

// SaskiKanadaNatLogOfX
// https://en.wikipedia.org/wiki/Natural_logarithm
// Sasaki, T.; Kanada, Y. (1982).
// "Practically fast multiple-precision evaluation of log(x)".
// Journal of Information Processing. 5 (4): 247–250. Retrieved 2011-03-30.
func (bLog BigIntMathLogarithms) SaskiKanadaNatLogOfX(
	xNum,
	xNumPrecision,
	m,
	s4DivPrecisionMaxInternalPrecision,
	agMeanMaxInternalPrecision,
	agMeanMaxOutputPrecision,
	piDivideMaxPrecision,
	maxFinalResultPrecision *big.Int) (lnOfX, lnOfXPrecision *big.Int, err error) {

	ePrefix := "BigIntMathLogarithms.SaskiKanadaNatLogOfX()"

	lnOfX = big.NewInt(0)
	lnOfXPrecision = big.NewInt(0)
	err = nil

	bigZero := big.NewInt(0)

	if xNumPrecision == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'xNumPrecision' is nil and INVALID!\n",
			ePrefix)

		return lnOfX, lnOfXPrecision, err
	}

	if xNum == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'xNum' is nil and INVALID!\n",
			ePrefix)

		return lnOfX, lnOfXPrecision, err
	}

	if m == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'm' is nil and INVALID!\n",
			ePrefix)

		return lnOfX, lnOfXPrecision, err
	}

	if piDivideMaxPrecision == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'piDivideMaxPrecision' is nil and INVALID!\n",
			ePrefix)

		return lnOfX, lnOfXPrecision, err
	}

	if maxFinalResultPrecision == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxFinalResultPrecision' is nil and INVALID!\n",
			ePrefix)

		return lnOfX, lnOfXPrecision, err
	}

	if m.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'm' is negative and INVALID!\n"+
			"m='%v'\n",
			ePrefix,
			m.Text(10))

		return lnOfX, lnOfXPrecision, err
	}

	if xNumPrecision.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'xNumPrecision' is negative and INVALID!\n"+
			"xNumPrecision='%v'\n",
			ePrefix,
			xNumPrecision.Text(10))

		return lnOfX, lnOfXPrecision, err
	}

	if piDivideMaxPrecision.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'piDivideMaxPrecision' is negative and INVALID!\n"+
			"piDivideMaxPrecision='%v'\n",
			ePrefix,
			piDivideMaxPrecision.Text(10))

		return lnOfX, lnOfXPrecision, err
	}

	if maxFinalResultPrecision.Cmp(bigZero) == -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxFinalResultPrecision' is negative and INVALID!\n"+
			"maxFinalResultPrecision='%v'\n",
			ePrefix,
			maxFinalResultPrecision.Text(10))

		return lnOfX, lnOfXPrecision, err
	}

	sFactor2Tom := big.NewInt(0).Exp(big.NewInt(2), m, nil)

	s := big.NewInt(0).Mul(xNum, sFactor2Tom)
	sPrecision := big.NewInt(0).Set(xNumPrecision)

	s, sPrecision, errX :=
		BigIntMath{}.RoundToMaxPrecision(
			s,
			sPrecision,
			big.NewInt(0).Add(s4DivPrecisionMaxInternalPrecision, big.NewInt(20)),
			true)

	fourDivS, fourDivSPrecision, errX :=
		BigIntMathDivide{}.BigIntFracQuotient(big.NewInt(4), big.NewInt(0), s, sPrecision, s4DivPrecisionMaxInternalPrecision)

	if errX != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by: \n"+
			" fourDivS, fourDivSPrecision, errX :=BigIntMathDivide{}.BigIntFracQuotient(...)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return lnOfX, lnOfXPrecision, err
	}

	// Uses maxInternal Precision
	agMean, agMeanPrecision, _, _, _, errX :=
		BigIntMath{}.ArithmeticGeometricMean(
			big.NewInt(1),
			big.NewInt(0),
			fourDivS,
			fourDivSPrecision,
			agMeanMaxInternalPrecision,
			agMeanMaxOutputPrecision)

	if errX != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by:\n"+
			" agMean, agMeanPrecision, _, _, _, errX :=BigIntMath{}.ArithmeticGeometricMean(...)\n"+
			"%v", errX.Error())

		return lnOfX, lnOfXPrecision, err
	}

	denomFactorM := big.NewInt(0).Mul(big.NewInt(2), agMean)
	denomFactorMPrecision := big.NewInt(0).Set(agMeanPrecision)

	// Uses MaxInternal Precision
	factor1, factor1Precision, errX :=
		BigIntMathDivide{}.BigIntFracQuotient(
			piNumber20k.GetInteger(),
			piNumber20k.GetPrecisionBigInt(),
			denomFactorM,
			denomFactorMPrecision,
			piDivideMaxPrecision)

	if errX != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by:\n"+
			" factor1, factor1Precision, errX := BigIntMathDivide{}.BigIntFracQuotient(...)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return lnOfX, lnOfXPrecision, err
	}

	factor2 := big.NewInt(0).Mul(m, natLogTwo20k.GetInteger())
	factor2Precision := big.NewInt(0).Set(natLogTwo20k.GetPrecisionBigInt())

	lnOfX, lnOfXPrecision, errX =
		BigIntMathSubtract{}.BigIntSubtract(
			factor1,
			factor1Precision,
			factor2,
			factor2Precision)

	if errX != nil {
		lnOfX = big.NewInt(0)
		lnOfXPrecision = big.NewInt(0)

		err = fmt.Errorf("%v\n"+
			"Error returned by: \n"+
			" lnOfX, lnOfXPrecision, errX = BigIntMathSubtract{}.BigIntSubtract(...)\n"+
			"Error= %v\n",
			ePrefix,
			errX.Error())

		return lnOfX, lnOfXPrecision, err
	}

	if lnOfXPrecision.Cmp(maxFinalResultPrecision) == 1 {

		lnOfX, lnOfXPrecision, errX =
			BigIntMath{}.RoundToMaxPrecision(
				lnOfX,
				lnOfXPrecision,
				maxFinalResultPrecision,
				true)

		if errX != nil {

			lnOfX = big.NewInt(0)
			lnOfXPrecision = big.NewInt(0)

			err = fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" lnOfX, lnOfXPrecision, errX = BigIntMath{}.RoundToMaxPrecision(...)\n"+
				"Error= %v\n",
				ePrefix,
				errX.Error())

			return lnOfX, lnOfXPrecision, err
		}
	}

	err = nil

	return lnOfX, lnOfXPrecision, err
}
