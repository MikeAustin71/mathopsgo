package mathops

import (
	"fmt"
	"math/big"
	"errors"
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

	if base.IsZero() {
		return result,
		 errors.New(ePrefix + "Error: 'base' input parameter is Zero. INVALID INPUT!")
	}

	if exponent.IsZero() {
		return BigIntNum{}.NewOne(0), nil
	}

	bigOne := BigIntNum{}.NewOne(exponent.GetPrecisionUint())

	if exponent.Equal(bigOne) {
		result = base.CopyOut()
		return base, nil
	}

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
		 			" Error='%v' \n", err.Error())
		 }

		} else {
			// fractional exponent must be a negative number

			result, err = bIPwr.raiseToNegativeFractionalPower(base, exponent, maxPrecision)

			if err != nil {
				return BigIntNum{},
					fmt.Errorf(ePrefix	+
						"Error='%v' \n", err.Error())
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
		 	"Error='%v' ",err.Error())
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
		return BigIntNum{}.NewZero(0),
		fmt.Errorf(ePrefix + "Error returned by bIPwr.raiseToPositive" +
			"IntegerPower(base, bINumNumerator)" +
			"base='%v' bINumNumerator='%v' Error='%v' \n",
				base.GetNumStr(), bINumNumerator.GetNumStr(), err.Error())
	}

	nthRoot := BigIntNum{}.NewBigInt(ratExponent.Denom(), 0)

	// If nthRoot is zero, the result will always be '1'
	if nthRoot.IsZero() {
		return BigIntNum{}.NewOne(maxPrecision), nil
	}

	bigINumOne := BigIntNum{}.NewOne(0)

	// Error if nthRoot == 1
	if nthRoot.Cmp(bigINumOne) == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix +
				"Error - Input Parameter 'nthRoot' INVALID! 'nthRoot' cannot equal 1.\n")
	}

	if pwr1Result.GetSign() == -1 {

		isEvenNum, err := nthRoot.IsEvenNumber()

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix +
					"Error returned by nthRoot.IsEvenNumber() " +
					"nthRoot='%v' Error='%v'\n",nthRoot.GetNumStr(), err.Error())
		}

		if isEvenNum {
			return BigIntNum{}.NewZero(0),
				errors.New(ePrefix +
					"INVALID ENTRY - Cannot calculate nthRoot of a negative number when nthRoot is even. \n")
		}

	}

	biNumResult, err := BigIntMathNthRoot{}.GetNthRoot(pwr1Result, nthRoot, maxPrecision)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix +
					"Error returned by dBigIntMathNthRoot{}.GetNthRoot(...) " +
				"Error='%v'",	err.Error())
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

	bigINewPrecision := big.NewInt(0).Mul(bigIBasePrecision, exponent.absBigInt)

	newPrecision := uint(bigINewPrecision.Int64())

	result := big.NewInt(0).Exp(base.bigInt, exponent.absBigInt, nil)

	bINumResult1 := BigIntNum{}.NewBigInt(result, newPrecision)

	inverseResult, err := bINumResult1.GetInverse(maxPrecision)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by bINumResult1.GetInverse(maxPrecision)" +
				"bINumResult1='%v' maxPrecision='%v' Error='%v'",
				 bINumResult1.GetNumStr(), maxPrecision, err.Error())
	}

	return inverseResult, nil
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