package mathops

import (
	"fmt"
	"math/big"
)

type BigIntMathPower struct {
	Base BigIntNum
	Exponent  BigIntNum
	Result BigIntNum
}


func (bIPwr BigIntMathPower) Pwr(base, exponent BigIntNum, maxPrecision uint) (BigIntNum, error) {
	ePrefix := "BigIntMathPower.Pwr() "

	result := BigIntNum{}
	var err error

	if exponent.GetPrecisionUint() == uint(0) &&
			exponent.GetSign() > 0 {
		result, err = bIPwr.raiseToPositiveIntegerPower(base, exponent)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix +
					"Error returned by bIPwr.raiseToPositiveIntegerPower(" +
					"base, exponent, maxPrecision) " +
					"base='%v' exponent='%v' Error='%v' ",
					base.GetNumStr(), exponent.GetNumStr(), err.Error())
		}

	} else if exponent.GetPrecisionUint() == uint(0) &&
		exponent.GetSign() < 0 {
		result, err = bIPwr.raiseToNegativeIntegerPower(base, exponent, maxPrecision)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix + "Error returned by bIPwr.raiseToNegativeIntegerPower(" +
					"base, exponent, maxPrecision) " +
					"base='%v' exponent='%v' maxPrecision='%v' Error='%v' ",
						base.GetNumStr(), exponent.GetNumStr(), maxPrecision, err.Error())
		}

	} else {
		result = BigIntNum{}.NewOne()
	}



	return result, nil
}

// raiseToNegativeIntegerPower - Assumes that input parameter 'exponent' is negative and an
// integer number (precision==0 - no fractional digits). If 'exponent' is positive or if
// it is NOT an integer number, an error is returned.
//
// If 'exponent' is both a negative number and an integer number, this method proceeds to
// raise input parameter 'base' to the power of 'exponent' and return the result as a BigIntNum
// type.
//
func (bIPwr *BigIntMathPower) raiseToNegativeIntegerPower(
															base,
																exponent BigIntNum,
																	maxPrecision uint) (BigIntNum, error) {

  ePrefix := "BigIntMathPower.raiseToNegativeIntegerPower()"

	if exponent.GetPrecisionUint() > 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error - The Exponent is a Fractional Number. " +
				"Only Integer Exponents can be processed by this method! exponent='%v' ",
				exponent.GetNumStr())
	}

	if exponent.GetSign() > 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error = The Exponent is a Positive Number. " +
				"Only Negative Exponents can be processed by this method! exponent='%v'",
				exponent.GetNumStr())
	}

	bigIBasePrecision := big.NewInt(int64(base.GetPrecisionUint()))

	bigINewPrecision := big.NewInt(0).Mul(bigIBasePrecision, exponent.bigInt)

	newPrecision := uint(bigINewPrecision.Int64())

	result1 := big.NewInt(0).Exp(base.bigInt, exponent.absBigInt, nil)

	bINumResult1 := BigIntNum{}.NewBigInt(result1, newPrecision)

	bINumOne := BigIntNum{}.NewOne()

	result, err := BigIntMathDivide{}.BigIntNumFracQuotient(bINumOne, bINumResult1, maxPrecision)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by BigIntMathDivide{}.BigIntNumFracQuotient" +
				"(bINumOne, bINumResult1, maxPrecision) " +
				"bINumOne='%v' bINumResult1='%v' maxPrecision='%v' Error='%v'",
				bINumOne.GetNumStr(), bINumResult1.GetNumStr(), maxPrecision, err.Error())

	}

	return result, nil
}

// raiseToPositiveIntegerPower - Assumes that input parameter 'exponent' is a positive
// integer number. If 'exponent' is negative or if 'exponent' is NOT an integer (precision > 0),
// an error will be triggered.
//
// If the exponent is both positive and an integer number, this method proceeds to raise the 'base'
// parameter to the power of 'exponent' and returns the result as a 'BigIntNum' type.
//
//	Examples:
//  =========
//
//	base^exponent = result
//
//	Base			Exponent		Result
//	 2					2						 4
//   3					4						81
//   4.2				3						74.088
// 	10					3					1000
//  -4.2				3					 -74.088
//	-2.9				4					  70.7281
//
func (bIPwr *BigIntMathPower) raiseToPositiveIntegerPower(
				base,
					exponent BigIntNum) (BigIntNum, error) {

	ePrefix := "BigIntMathPower.raiseToPositiveIntegerPower() "

	if exponent.GetPrecisionUint() > 0 {
		return BigIntNum{},
		 fmt.Errorf(ePrefix + "Error - The Exponent is a Fractional Number. " +
		 	"Only Integer Exponents can be processed by this method! exponent='%v' ",
		 		exponent.GetNumStr())
	}

	if exponent.GetSign() < 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error = The Exponent is a Negative Number. " +
				"Only Positive Exponents can be processed by this method! exponent='%v'",
					exponent.GetNumStr())
	}

	bigIBasePrecision := big.NewInt(int64(base.GetPrecisionUint()))

	bigINewPrecision := big.NewInt(0).Mul(bigIBasePrecision, exponent.bigInt)

	newPrecision := uint(bigINewPrecision.Int64())

	result := big.NewInt(0).Exp(base.bigInt, exponent.bigInt, nil)

	return BigIntNum{}.NewBigInt(result, newPrecision), nil
}