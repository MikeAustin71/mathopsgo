package mathops

import (
	"errors"
	"fmt"
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
	Base 				BigIntNum
	XNumber 		BigIntNum
  YExponent 	BigIntNum
}

// LogBaseOfX - Computes the log[base](xNum). 'base' and 'xNum' are passed as
// BigIntNum types.
//
// 'maxPrecision' is an uint specifying the precision or integerNum of digits to the
// right of the decimal place in the result.
//
//
// Calculation Formula:
// ====================
//			log[base](x) = y
//
//			x = base^y
//			'y' is the exponent of base required to set base^y equal to x.
// 			The calculation result 'y' of log[base](xNum) is returned as a BigIntNum.
//
// Input Parameters:
// =================
//
// base		BigIntNum	- The base of the logarithm. 'base' must be an integer integerNum
// 										greater than one (base > +1).
//
// xNum		BigIntNum - The 'x' value such that base^y = 'x'.  xNum must be greater
//										than zero.
//
// Return Values:
// ==============
// y			BigIntNum - If the function completes successfully, the result 'y' will be
//										returned as a BigIntNum.
//
// 				error			- If the function encounters an error, an error message will be returned
//										as type 'error'. If the function completes successfully, this value
//										will be set to 'nil'.
//
//          log Value:  10.258265535016932
// Expected log Value:  10.258265535022667
// base:  4
// xNum:  1500000
func (bLog BigIntMathLogarithms) LogBaseOfX(
	base, xNum BigIntNum, maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntMathLogarithms.LogBaseOfX() "
	maxInternalPrecision := xNum.GetPrecisionUint() + uint(200)
	bigZero := BigIntNum{}.NewZero(0)

	digits, residualXNum, err := bLog.GetIntDigits(base, xNum, maxInternalPrecision)

	if err != nil {
		return bigZero,
		fmt.Errorf(ePrefix + "- Error='%v'", err.Error() )
	}

	if residualXNum.IsZero() {
		return digits, nil
	}

	result := digits.CopyOut()

	digits.SetBigInt(big.NewInt(0), 0)

	bigTen := BigIntNum{}.NewInt(10, 0)

	residualXNum, err = BigIntMathPower{}.Pwr(residualXNum, bigTen, maxInternalPrecision)

	if err != nil {
		return bigZero,
			fmt.Errorf(ePrefix +
				"Error returned by BigIntMathPower{}.Pwr(residualXNum, bigTen, maxInternalPrecision). " +
				"residualXNum='%v', maxPrecision='%v' Error='%v'",
				residualXNum.GetNumStr(), maxPrecision, err.Error())
	}

	maxPrecision++

	for x := uint(0); x < maxPrecision; x++ {


		digits, residualXNum, err = bLog.GetNextDecimalDigit(base, residualXNum, maxInternalPrecision)

		if err != nil {
			return bigZero,
				fmt.Errorf(ePrefix + "- Error='%v'", err.Error() )
		}

		result.MultiplyByTenToPowerAdd(1, digits)

		if residualXNum.IsZero() {
			result.SetPrecision(x+1)
			return result, nil
		}


	}


	result.ShiftPrecisionLeft(maxPrecision)

	maxPrecision--

	result.RoundToDecPlace(maxPrecision)


	return result, nil
}


// GetIntDigits - This function is designed to return all the logarithm
// digits associated with the integer portion of XNum.
func (bLog BigIntMathLogarithms) GetIntDigits(
		base, xNum BigIntNum,
		maxPrecision uint) (digits, newResidualXNum BigIntNum, err error) {

	ePrefix := "BigIntMathLogarithms.GetIntDigits() "
	var errX error
	digits = BigIntNum{}.NewZero(0)
	newResidualXNum = xNum.CopyOut()
	err = nil

	for newResidualXNum.Cmp(base) > -1 {

		newResidualXNum, errX =
			BigIntMathDivide{}.BigIntNumFracQuotient(newResidualXNum, base, maxPrecision)

		if errX != nil {
			err =
				fmt.Errorf(ePrefix +
					"Error returned by BigIntMathDivide{}.BigIntNumFracQuotient(newResidualXNum, " +
					"base, maxPrecision ) newResidualXNum='%v', base='%v' maxPrecision='%v' Error='%v' ",
					newResidualXNum.GetNumStr(), base.GetNumStr(), maxPrecision, errX.Error())

				return digits, newResidualXNum, err
		}

		digits.Increment()

	}

	return digits, newResidualXNum, nil
}

// GetNextDecimalDigit - Private function designed to return the next
// logarithm digit to the right of the decimal place in XNum.
func (bLog BigIntMathLogarithms) GetNextDecimalDigit(
	base, residualXNum BigIntNum,
	maxPrecision uint) (digit, newResidualXNum BigIntNum, err error) {

	ePrefix := "BigIntMathLogarithms.getDecimalDigits() "
	digit = BigIntNum{}.NewZero(0)

	newResidualXNum = residualXNum.CopyOut()
	err = nil
	var errX error
	bigTen := BigIntNum{}.NewInt(10, 0)

	if newResidualXNum.Cmp(base) == -1 {

		newResidualXNum, errX = BigIntMathPower{}.Pwr(newResidualXNum, bigTen, maxPrecision)

		if errX != nil {
			err =
				fmt.Errorf(ePrefix +
					"Error returned by BigIntMathPower{}.Pwr(newResidualXNum, bigTen, maxPrecision) #1. " +
					"newResidualXNum='%v', maxPrecision='%v' Error='%v'",
					newResidualXNum.GetNumStr(), maxPrecision, errX.Error())
			return digit, newResidualXNum, err
		}

		return digit, newResidualXNum, err
	}




	for newResidualXNum.Cmp(base) > -1 {

		//fmt.Println("Before newResidualXNum= ", newResidualXNum.GetNumStr())

		newResidualXNum, errX =
			BigIntMathDivide{}.BigIntNumFracQuotient(newResidualXNum, base, maxPrecision)

		if errX != nil {

			err =
				fmt.Errorf(ePrefix +
					"Error returne by BigIntMathDivide{}.BigIntNumFracQuotient(newResidualXNum, " +
					"base, maxPrecision ) newResidualXNum='%v', base='%v' maxPrecision='%v' Error='%v' ",
					newResidualXNum.GetNumStr(), base.GetNumStr(), maxPrecision, errX.Error())

			return digit, newResidualXNum, err
		}

		//fmt.Println("After newResidualXNum= ", newResidualXNum.GetNumStr())

		digit.Increment()

	}

	//fmt.Println("   Last residualXNum: ", newResidualXNum.GetNumStr())


	if !newResidualXNum.IsZero() && newResidualXNum.Cmp(base) == -1	{

		newResidualXNum, errX = BigIntMathPower{}.Pwr(newResidualXNum, bigTen, maxPrecision)

		if errX != nil {
			err =
				fmt.Errorf(ePrefix +
					"Error returned by BigIntMathPower{}.Pwr(residualXNum, bigTen, maxPrecision) #2. " +
					"residualXNum='%v', maxPrecision='%v' Error='%v'",
					residualXNum.GetNumStr(), maxPrecision, errX.Error())
			return digit, newResidualXNum, err
		}

	}



	err = nil

	return digit, newResidualXNum, err
}

// EPwrXFromMaclaurinSeries = Computes value of e^exponent where e is the mathematical
// constant, "Euler's integerNum".
//
// "Euler's integerNum" or 'e' truncated to 50 decimal places is:
//     2.71828182845904523536028747135266249775724709369995
// Rounded to 50-decimal places is:
//     2.71828182845904523536028747135266249775724709369996
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
//                      be raised in order to compute the value of e^exponent.
//
// n				int64			- The integerNum of cycles in the Taylor series which will
//                      be used to compute the value of e^exponent
//
// Return Values
// =============
//
// BigIntNum					- If successful, this function returns the value of e^exponent
//											as a type BigIntNum.
//
// error							- If an error is encountered an error message will be formatted
//											and returned as type 'error'. If the function successfully
//											completes the calculation, this value will be set to 'nil'.
//
func (bLog BigIntMathLogarithms) EPwrXFromMaclaurinSeries(
	exponent BigIntNum, nCycles int64) (BigIntNum, error) {

	ePrefix := "BigIntMathLogarithms.EPwrXFromMaclaurinSeries() "
	// sum = 1 + x
	sum := BigIntMathAdd{}.AddBigIntNums(BigIntNum{}.NewInt(1, 0), exponent)
	nFactorial := BigIntNum{}.NewInt(1,0)

	x := exponent.CopyOut()

	for n:= int64(2); n <= nCycles ; n++ {

		nFactorial = BigIntMathMultiply{}.MultiplyBigIntNums(nFactorial, BigIntNum{}.NewInt64(n, 0))

		x = BigIntMathMultiply{}.MultiplyBigIntNums(x, exponent)

		factor, err := BigIntMathDivide{}.BigIntNumFracQuotient(x, nFactorial, 1200)

		if err != nil {
			return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntMathDivide{}.BigIntNumFracQuotient(binOne, nFactorial, 200) "+
				"nFactorial='%v' Error='%v' ",
				nFactorial.GetNumStr(), err.Error())
		}

		sum = BigIntMathAdd{}.AddBigIntNums(sum, factor)

		if sum.GetPrecisionUint() > 500 {
			sum.RoundToDecPlace(500)
		}

	}

	return sum, nil
}

func (bLog BigIntMathLogarithms) EPwrXFromTaylorSeries(
	exponent, a BigIntNum, nCycles int64) (BigIntNum, error) {

	ePrefix := "BigIntMathLogarithms.EPwrXFromTaylorSeries() "

	e := eulersNumber1050.GetBigIntNum()

	if e.IsZero() {
		return BigIntNum{}.NewZero(0),
		errors.New(ePrefix + "eulersNumber1050 is ZERO!")
	}

	aValue, err := a.GetUInt()

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by a.GetUInt. Error='%v'",
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
			outputMaxPrecision )

	eToPwr := BigIntNum{}.NewBigInt(ePwrBigInt, ePwrBigIntPrecision)

	// eToPwr, err := BigIntMathPower{}.Pwr(e, a, 500)

	//eToPwr := BigIntMathMultiply{}.MultiplyBigIntNums(e.CopyOut(), e.CopyOut())

	// sum = 0
	sum := BigIntNum{}.NewInt(0,0)

	x := exponent.CopyOut()

	xMinusA := BigIntMathSubtract{}.SubtractBigIntNums(x, a)

	xMinusANth := BigIntNum{}.NewInt(0, 0)

	nFact := BigIntNum{}.NewInt(0, 0)

	for n:= int64(0); n <= nCycles ; n++ {

		if n==0 {
			xMinusANth = BigIntNum{}.NewInt(1,0)
			nFact = BigIntNum{}.NewInt(1,0)
		} else if n==1 {
			xMinusANth = xMinusA.CopyOut()
			nFact = BigIntNum{}.NewInt64(n, 0)
		} else {
			xMinusANth = BigIntMathMultiply{}.MultiplyBigIntNums(xMinusANth, xMinusA)
			nFact = BigIntMathMultiply{}.MultiplyBigIntNums(nFact, BigIntNum{}.NewInt64(n, 0))
		}

		factor1, err := BigIntMathDivide{}.BigIntNumFracQuotient(eToPwr, nFact, 500)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntMathDivide{}.BigIntNumFracQuotient(eToPwr, nFact, 500) "+
					"eToPwr='%v' nFact='%v' Error='%v' ",
					eToPwr.GetNumStr(), nFact.GetNumStr(), err.Error())
		}

		factor2:= BigIntMathMultiply{}.MultiplyBigIntNums(factor1, xMinusANth)

		sum = BigIntMathAdd{}.AddBigIntNums(sum, factor2)

	}

	sum.RoundToDecPlace(500)

	return sum, nil

}



// EPwrXFromTaylorSeriesBigInt
func (bLog BigIntMathLogarithms) EPwrXFromTaylorSeriesBigInt(
	exponentX BigIntFixedDecimal,
	a, nCycles uint) (BigIntFixedDecimal, error) {

	ePrefix := "BigIntMathLogarithms.EPwrXFromTaylorSeries() "

	e := eulersNumber1050.GetFixedDecimal()

	if e.IsZero() {
		return BigIntFixedDecimal{}.NewZero(0),
		errors.New(ePrefix + "EulersNumber1050 Constant is ZERO!")
	}

	internalMaxPrecision := uint(20000)

	outputMaxPrecision := uint(1500)

	ePwrBigInt, ePwrBigIntPrecision :=
		BigIntMathPower{}.BigIntPwrIteration(
			e.GetInteger(),
			e.GetPrecision(),
			a,
			internalMaxPrecision,
			outputMaxPrecision )

	eToPwr :=  BigIntFixedDecimal{}.New(ePwrBigInt, ePwrBigIntPrecision)

	fixedDecA := BigIntFixedDecimal{}.New(big.NewInt(int64(a)), 0)

	// sum = 0
	sum := BigIntFixedDecimal{}.NewZero(0)

	xNum := exponentX.CopyOut()

	xMinusA := BigIntMathSubtract{}.FixedDecimalSubtract(xNum, fixedDecA )

	xMinusANth := BigIntFixedDecimal{}.NewZero(0)

	nFact := BigIntFixedDecimal{}.NewZero(0)

	for n:=uint(0); n < nCycles; n++ {

		if n==0 {

			xMinusANth = BigIntFixedDecimal{}.NewInt(1, 0)
			nFact = BigIntFixedDecimal{}.NewInt(1, 0)

		} else if n==1 {

			xMinusANth = xMinusA.CopyOut()
			nFact =  BigIntFixedDecimal{}.NewUInt(n, 0)

		} else {

			xMinusANth = BigIntMathMultiply{}.FixedDecimalMultiply(xMinusANth, xMinusA)

			nFact =
				BigIntMathMultiply{}.FixedDecimalMultiply(
					nFact,
					BigIntFixedDecimal{}.New(big.NewInt(int64(n)), 0))

		}

		factor1, err := BigIntMathDivide{}.FixedDecimalFracQuotient(eToPwr, nFact,500)

		if err != nil {
			return BigIntFixedDecimal{}.NewZero(0),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntMathDivide{}.BigIntNumFracQuotient(eToPwr, nFact, 500) "+
					"eToPwr='%v' nFact='%v' Error='%v' ",
					eToPwr.GetNumStr(), nFact.GetNumStr(), err.Error())
		}

		factor2 := BigIntMathMultiply{}.FixedDecimalMultiply(factor1, xMinusANth)

		sum = BigIntMathAdd{}.FixedDecimalAdd(sum, factor2)

	}

	sum.RoundToDecPlace(500)

	return sum, nil
}

