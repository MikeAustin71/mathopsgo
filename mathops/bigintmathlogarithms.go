package mathops

import (
	"fmt"
	"math/big"
)

/*
https://en.wikipedia.org/wiki/Logarithm
log[b](x) = y exactly if b^y = x

log[10](450) = 2.6532125137753436793763169117857
precision = 31
 */

type BigIntMathLogarithms struct {
	Base 				BigIntNum
	XNumber 		BigIntNum
  YExponent 	BigIntNum
}

// LogBaseOfX - Computes the log[base](xNum). 'base' and 'xNum' are passed as
// BigIntNum types.
//
// 'maxPrecision' is an uint specifying the precision or number of digits to the
// right of the decimal place in the result.
//
// The calculation result of log[base](xNum) is returned as a BigIntNum.
//
func (bLog BigIntMathLogarithms) LogBaseOfX(
	base, xNum BigIntNum, maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntMathLogarithms.LogBaseOfX() "
	maxInternalPrecision := xNum.GetPrecisionUint() + uint(100)
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

	if newResidualXNum.Cmp(base) == -1 {

		bigTen := BigIntNum{}.NewInt(10, 0)

		residualXNum, errX = BigIntMathPower{}.Pwr(residualXNum, bigTen, maxPrecision)

		if errX != nil {
			err =
				fmt.Errorf(ePrefix +
					"Error returned by BigIntMathPower{}.Pwr(residualXNum, bigTen, maxPrecision) #1. " +
					"residualXNum='%v', maxPrecision='%v' Error='%v'",
					residualXNum.GetNumStr(), maxPrecision, errX.Error())
			return digit, newResidualXNum, err
		}

		return digit, residualXNum, err
	}

	for newResidualXNum.Cmp(base) > -1 {

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

		digit.Increment()

	}

	err = nil

	return digit, residualXNum, err
}