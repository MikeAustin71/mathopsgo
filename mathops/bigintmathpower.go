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

// Pwr - Raises 'base' to the power of 'exponent'.  Both 'base' and 'exponent' are Type BigIntNum.
// Upon computing the result of 'base' raised to the power of 'exponent' (base^exponent), the result
// is returned as Type BigIntNum.
//
// Examples:
// =========
//
//	base				exponent				maxPrecision	  		result
// -------			---------  			-------------	 			-------
//
//	2							 4								17								16
//  2							-4								17								 0.0625
//  3.7						 2.8							30								38.991040735983142451443031376258
//  3.7						-2.8							32 								 0.02564691737189623971146450249457
// -2							-3.8							32								 0.07179364718731468792491418417362
// -2              3.8              30								13.928809012737986226180320279676
func (bIPwr BigIntMathPower) Pwr(base, exponent BigIntNum, maxPrecision uint) (BigIntNum, error) {
	ePrefix := "BigIntMathPower.Pwr() "

	result := BigIntNum{}
	var err error

	if exponent.GetPrecisionUint() == uint(0) {

		if exponent.GetSign() > 0 {
			// exponent is a positive number
			result, err = bIPwr.raiseToPositiveIntegerPower(base, exponent)

			if err != nil {
				return BigIntNum{}.NewZero(0),
					fmt.Errorf(ePrefix +
						"Error returned by bIPwr.raiseToPositiveIntegerPower(" +
						"base, exponent, maxPrecision) " +
						"base='%v' exponent='%v' Error='%v' ",
						base.GetNumStr(), exponent.GetNumStr(), err.Error())
			}

		} else {
			// exponent must be a negative number
			result, err = bIPwr.raiseToNegativeIntegerPower(base, exponent, maxPrecision)

			if err != nil {
				return BigIntNum{}.NewZero(0),
					fmt.Errorf(ePrefix + "Error returned by bIPwr.raiseToNegativeIntegerPower("+
						"base, exponent, maxPrecision) "+
						"base='%v' exponent='%v' maxPrecision='%v' Error='%v' ",
						base.GetNumStr(), exponent.GetNumStr(), maxPrecision, err.Error())
			}

		}

	} else {
		// precision must be greater than zero. exponent is a fractional number

		if exponent.GetSign() > 0 {
			// fractional exponent is a positive number

		 result, err = bIPwr.raiseToPositiveFractionalPower(base, exponent, maxPrecision)

		 if err != nil {
		 	return BigIntNum{},
		 		fmt.Errorf(ePrefix	+
		 			"Error returned by bIPwr.raiseToPositiveFractionalPower(base, " +
		 				"exponent, maxPrecision). base='%v' exponent='%v' " +
		 				"maxPrecision='%v' Error='%v' ",
		 					base.GetNumStr(), exponent.GetNumStr(), maxPrecision, err.Error())
		 }

		} else {
			// fractional exponent must be a negative number

			result, err = bIPwr.raiseToNegativeFractionalPower(base, exponent, maxPrecision)

			if err != nil {
				return BigIntNum{},
					fmt.Errorf(ePrefix	+
						"Error returned by bIPwr.raiseToNegativeFractionalPower(base, " +
						"exponent, maxPrecision). base='%v' exponent='%v' " +
						"maxPrecision='%v' Error='%v' ",
						base.GetNumStr(), exponent.GetNumStr(), maxPrecision, err.Error())
			}

		}
	}

	if result.precision > maxPrecision {
		result.SetPrecision(maxPrecision)
	}

	return result, nil
}

// raiseToNegativeFractionalPower - Assumes that input parameter 'exponent' is negative and a
// fractional number (precision > 0 - has fractional digits). If 'exponent' is positive or if
// 'exponent' is an integer number, an error is returned.
//
// If 'exponent' is both a negative number and a fractional number, this method proceeds to
// raise input parameter 'base' to the power of 'exponent' and return the result as a BigIntNum
// type.
//
func (bIPwr BigIntMathPower) raiseToNegativeFractionalPower(
		base,
			exponent BigIntNum,
				maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntMathPower.raiseToNegativeFractionalPower() "
	
	if exponent.GetPrecisionUint() == 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error - The Exponent is an Integer Number. " +
				"Only Fractional Exponents can be processed by this method! exponent='%v' ",
				exponent.GetNumStr())
	}

	if exponent.GetSign() > 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error = The Exponent is a Positive Number. " +
				"Only Negative Exponents can be processed by this method! exponent='%v'",
				exponent.GetNumStr())
	}

	exponentAbsVal := exponent.GetAbsoluteBigIntNumValue()
	bINumResult, err :=
							bIPwr.raiseToPositiveFractionalPower(
											base,
												exponentAbsVal,
													maxPrecision)


	if err != nil {
		return BigIntNum{},
		 fmt.Errorf(ePrefix +
		 	"Error returned by bIPwr.raiseToPositiveFractionalPower(base," +
			"exponent.GetAbsoluteBigIntNumValue(), maxPrecision)" +
			"base='%v' exponentAbsValue='%v' maxPrecision='%v' Error='%v' ",
			base.GetNumStr(), exponentAbsVal.GetNumStr(),
			err.Error())
	}

	inverseResult, err := bINumResult.GetInverse(maxPrecision)

	if err != nil {
		return BigIntNum{},
		fmt.Errorf(ePrefix + "Error returned by bINumResult.GetInverse(maxPrecision). " +
			"bINumResult='%v' maxPrecision='%v' Error='%v' ",
				bINumResult.GetNumStr(), maxPrecision, err.Error())
	}

	return inverseResult, nil
}


// raiseToPositiveFractionalPower - Assumes that input parameter 'exponent' is positive and a
// fractional number (precision > 0 - has fractional digits). If 'exponent' is negative or if
// 'exponent' is an integer number, an error is returned.
//
// If 'exponent' is both a positive number and a fractional number, this method proceeds to
// raise input parameter 'base' to the power of 'exponent' and return the result as a BigIntNum
// type.
//
func (bIPwr BigIntMathPower) raiseToPositiveFractionalPower(
		base,
			exponent BigIntNum,
				maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntMathPower.raiseToPositiveFractionalPower() "

	if exponent.GetPrecisionUint() == 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error - The Exponent is an Integer Number. " +
				"Only Fractional Exponents can be processed by this method! exponent='%v' ",
				exponent.GetNumStr())
	}

	if exponent.GetSign() < 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error = The Exponent is a Negative Number. " +
				"Only Positive Exponents can be processed by this method! exponent='%v'",
				exponent.GetNumStr())
	}

	modXZero := big.NewInt(0)

	scaleFactor := big.NewInt(0).Exp(
				big.NewInt(10),
						big.NewInt(int64(exponent.precision)),
							modXZero )

	ratExponent := big.NewRat(1, 1 ).SetFrac(exponent.bigInt, scaleFactor)



	bINumNumerator := BigIntNum{}.NewBigInt(ratExponent.Num(), 0)

	pwr1Result, err := bIPwr.raiseToPositiveIntegerPower(base, bINumNumerator)

	if err != nil {
		return BigIntNum{},
		fmt.Errorf(ePrefix + "Error returned by bIPwr.raiseToPositive" +
			"IntegerPower(base, bINumNumerator)" +
			"base='%v' bINumNumerator='%v' Error='%v' ",
				base.GetNumStr(), bINumNumerator.GetNumStr(), err.Error())
	}

	denominator := BigIntNum{}.NewBigInt(ratExponent.Denom(), 0)

	nthRoot, err := denominator.GetUnsignedInt()

	if err != nil {
		return BigIntNum{},
		fmt.Errorf(ePrefix + "Error returned by denominator.GetUnsignedInt() Error='%v'",
			err.Error())
	}

	iaPwrResult1, err := IntAry{}.NewBigInt(pwr1Result.bigInt, int(pwr1Result.precision))

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by IntAry{}.NewBigInt(pwrResult1, 0). " +
				"pwrResult1='%v' Error='%v'", pwr1Result.GetNumStr(), err.Error())
	}

	nr := NthRootOp{}

  iaResult, err := nr.GetNthRootIntAry(&iaPwrResult1, nthRoot, maxPrecision)

  if err != nil {
  	return BigIntNum{},
  		fmt.Errorf(ePrefix + "Error returned by nr.GetNthRootIntAry(&iaPwrResult1, " +
  			"nthRoot, maxPrecision). iaPwrResult1='%v' nthRoot='%v' maxPrecision='%v' " +
  			"Error='%v' ", iaPwrResult1.GetNumStr(), nthRoot, maxPrecision, err.Error())
	}

	biNumResult, err := iaResult.GetBigIntNum()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by nr.GetNthRootIntAry(&iaPwrResult1, " +
				"nthRoot, maxPrecision). iaPwrResult1='%v' nthRoot='%v' maxPrecision='%v' " +
				"Error='%v' ", iaPwrResult1.GetNumStr(), nthRoot, maxPrecision, err.Error())
	}

	return biNumResult, nil
}

// raiseToNegativeIntegerPower - Assumes that input parameter 'exponent' is negative and an
// integer number (precision==0 - no fractional digits). If 'exponent' is positive or if
// it is NOT an integer number, an error is returned.
//
// If 'exponent' is both a negative number and an integer number, this method proceeds to
// raise input parameter 'base' to the power of 'exponent' and return the result as a BigIntNum
// type.
//
func (bIPwr BigIntMathPower) raiseToNegativeIntegerPower(
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

	result, err := bINumResult1.GetInverse(maxPrecision)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by bINumResult1.GetInverse(maxPrecision)" +
				"bINumResult1='%v' maxPrecision='%v' Error='%v'",
				 bINumResult1.GetNumStr(), maxPrecision, err.Error())
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
//  -2          3.8
func (bIPwr BigIntMathPower) raiseToPositiveIntegerPower(
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