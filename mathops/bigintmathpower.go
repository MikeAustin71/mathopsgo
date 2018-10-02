package mathops

import (
	"errors"
	"fmt"
	"math/big"
)

type BigIntMathPower struct {
	Base     BigIntNum
	Exponent BigIntNum
	Result   BigIntNum
}

// BigIntPwr - Raises input parameter 'base' to the power of input parameter 'exponent'.
// This method of raising a base to an exponent uses iterative multiplication and manages
// the internal precision of each iterative multiplication. If, during the process of
// multiplying the base time itself, the internal precision exceeds the 'internalMaxPrecision'
// limit, that intermediate number is rounded down to 'internalMaxPrecision'.
//
// If the precision of the final result exceeds the limit imposed by input parameter,
// 'outputMaxPrecision', that final result will be rounded to 'outputMaxPrecision'
// digits to the right of the decimal place.
//
// Input Parameter
// ===============
//
// base							*big.Int	-	The base which will be raised to the power of 'exponent'.
//
//                        							baseToPwr = base^exponent
//
// basePrecision				uint	- The number of digits to the right of the decimal place
//                        			in the numeric sequence represented by 'base'.
//
// exponent							uint	- This function will raise 'base' to the power of 'exponent'.
//
//                        							baseToPwr = base^exponent
//
//
// internalMaxPrecision uint	- This value is imposed as a limit on the precision of
//                              internal calculations necessary to compute the result
//                              of this power operation. If during the calculation an
//                              interim or intermediate result is generated which exceeds
//                              this limit, that intermediate result will be rounded to
//             									'internalMaxPrecision'. The term precision defines the
//                              number of digits to the right of the decimal place.
//
//                              If 'internalMaxPrecision' is less than 'outputMaxPrecision',
//                              'internalMaxPrecision' will be automatically set to a value
//                              of 'outputMaxPrecision' + 100.
//
// outputMaxPrecision		uint	- This value is imposed as a limit on the precision of
//                              the final calculated result of the power operation.
// 															If the number of digits to the right of the decimal
//                              point in the final calculated result exceeds this limit,
// 															that final result will be rounded to 'outputMaxPrecision'
//                              digits to the right of the decimal place. The term precision
// 															defines the number of digits to the right of the decimal
// 															place.
//
// Return Values
// =============
//
// baseToPwr						*big.Int	-	This function returns the result of 'base' raised
//                               		to the power of 'exponent'. This result, 'baseToPwr'
//                                  is returned as a type *big.Int.
//
//                                  				baseToPwr = base^exponent
//
// baseToPwrPrecision		uint			- Specifies the number of digits to the the right of the
//                                  decimal place in the numeric sequence represented
// 																	by the calculation result, 'baseToPwr'.
//
func (bIPwr BigIntMathPower) BigIntPwr(
	base *big.Int,
	basePrecision,
	exponent,
	internalMaxPrecision,
	outputMaxPrecision uint) (baseToPwr *big.Int, baseToPwrPrecision uint) {

	baseToPwrPrecision = 0

	exponent++

	bigIZero := big.NewInt(0)


	if base.Cmp(bigIZero) == 0 {
		baseToPwr = big.NewInt(0)
		baseToPwrPrecision = 0
		return baseToPwr, baseToPwrPrecision
	}

	if base.Cmp(big.NewInt(1)) == 0 {
		baseToPwr = big.NewInt(0).Set(base)
		baseToPwrPrecision = basePrecision
		return baseToPwr, baseToPwrPrecision
	}

	if internalMaxPrecision < outputMaxPrecision {
		internalMaxPrecision = outputMaxPrecision + 100
	}

	bigIPlusFive := big.NewInt(5)
	bigIMinusFive := big.NewInt(-5)
	bigITen := big.NewInt(10)
	roundFactor := big.NewInt(0)
	cmpResult := 0

	for i := uint(0); i < exponent; i++ {

		if i == 0 {
			baseToPwr = big.NewInt(1)
			baseToPwrPrecision = 0
		} else if i == 1 {
			baseToPwr = big.NewInt(0).Set(base)
			baseToPwrPrecision = basePrecision
		} else {
			baseToPwr = big.NewInt(0).Mul(baseToPwr, base)
			baseToPwrPrecision = baseToPwrPrecision + basePrecision
		}

		if baseToPwrPrecision > internalMaxPrecision {
			// beforeRounding := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)
			delta := baseToPwrPrecision - internalMaxPrecision - 1
			scale := big.NewInt(0).Exp(bigITen, big.NewInt(int64(delta)), nil)

			cmpResult = baseToPwr.Cmp(bigIZero)
			if cmpResult == 1 {
				// baseToPwr is GREATER Than Zero
				roundFactor = big.NewInt(0).Mul(bigIPlusFive, scale)
			} else if cmpResult == -1 {
				// baseToPwr is LESS Than Zero
				roundFactor = big.NewInt(0).Mul(bigIMinusFive, scale)
			}	else {
				baseToPwrPrecision = 0
				continue
			}

			baseToPwr = big.NewInt(0).Add(baseToPwr, roundFactor)
			scale = big.NewInt(0).Mul(scale, bigITen)
			baseToPwr = big.NewInt(0).Quo(baseToPwr, scale)
			baseToPwrPrecision = internalMaxPrecision
			/*
			afterRounding := BigIntNum{}.NewBigInt(baseToPwr, baseToPwrPrecision)
			fmt.Println()
			fmt.Println("BeforeRounding: ", beforeRounding.GetNumStr())
			fmt.Println(" AfterRounding: ", afterRounding.GetNumStr())
			fmt.Println("-----------------------------------------------")
			*/
		}

	}

	if baseToPwrPrecision > outputMaxPrecision {
		delta := baseToPwrPrecision - outputMaxPrecision - 1
		scale := big.NewInt(0).Exp(bigITen, big.NewInt(int64(delta)), nil)
		round := big.NewInt(0).Mul(bigIPlusFive, scale)
		if baseToPwr.Cmp(bigIZero) == -1 {
			round = big.NewInt(0).Mul(round, big.NewInt(-1))
		}
		baseToPwr = big.NewInt(0).Add(baseToPwr, round)
		scale = big.NewInt(0).Mul(scale, bigITen)
		baseToPwr = big.NewInt(0).Quo(baseToPwr, scale)
		baseToPwrPrecision = outputMaxPrecision

	}

	return baseToPwr, baseToPwrPrecision
}


// FixedDecimalPwr - Raises input parameter 'base' to the power of input parameter 'exponent'.
// This method of raising a base to an exponent uses iterative multiplication and manages
// the internal precision of each iterative multiplication. If, during the process of
// multiplying the base times itself, the internal precision exceeds the 'internalMaxPrecision'
// limit, that intermediate number is rounded down to 'internalMaxPrecision' digits to the
// right of the decimal place.
//
// If the precision of the final result exceeds the limit imposed by input parameter,
// 'outputMaxPrecision', that final result will be rounded to 'outputMaxPrecision'
// digits to the right of the decimal place.
//
// Input Parameter
// ===============
//
// base	BigIntFixedDecimal	-	The base which will be raised to the power of 'exponent'.
//                            The BigIntFixedDecimal type describes a numeric value with
//                            a fixed number of digits to the right of the decimal place.
//                            The type includes a *big.Int integer value and a precision
//                            specification.
//
//                        							baseToPwr = base^exponent
//
//
// exponent							uint	- This function will raise 'base' to the power of 'exponent'.
//
//                        							baseToPwr = base^exponent
//
//
// internalMaxPrecision uint	- This value is imposed as a limit on the precision of
//                              internal calculations necessary to compute the result
//                              of this power operation. If during the calculation an
//                              interim or intermediate result is generated which exceeds
//                              this limit, that intermediate result will be rounded to
//             									'internalMaxPrecision' digits to the right of the decimal
//                              place. The term precision defines the number of digits to
// 															the right of the decimal place.
//
//                              If 'internalMaxPrecision' is less than 'outputMaxPrecision',
//                              'internalMaxPrecision' will be automatically set to a value
//                              of 'outputMaxPrecision' + 100.
//
// outputMaxPrecision		uint	- This value is imposed as a limit on the precision of
//                              the final calculated result of the power operation.
// 															If the number of digits to the right of the decimal
//                              point in the final calculated result exceeds this limit,
// 															that final result will be rounded to 'outputMaxPrecision'
//                              digits to the right of the decimal place. The term precision
// 															defines the number of digits to the right of the decimal
// 															place.
//
// Return Values
// =============
//
// baseToPwr	BigIntFixedDecimal	-	This function returns the result of 'base' raised
//                               		to the power of 'exponent'. This result, 'baseToPwr'
//                                  is returned as a type BigIntFixedDecimal. The BigIntFixedDecimal
// 																	type describes a numeric value with a fixed number of digits
// 																	to the right of the decimal place.
//
//                                  				baseToPwr = base^exponent
//
func (bIPwr BigIntMathPower) FixedDecimalPwr(
	base BigIntFixedDecimal,
	exponent,
	internalMaxPrecision,
	outputMaxPrecision uint) (baseToPwr BigIntFixedDecimal) {

	baseToPwr = BigIntFixedDecimal{}.NewZero(0)

	bigIPwr, bigIPwrPrecision := BigIntMathPower{}.BigIntPwr(
		base.GetInteger(), base.GetPrecision(), exponent, internalMaxPrecision, outputMaxPrecision)

	baseToPwr.SetNumericValue(bigIPwr, bigIPwrPrecision)

	return baseToPwr
}

// MinimumRequiredPrecision - designed to be used with the power function
// below. This method will compute the minimum number of decimal places
// required to support the result of raising a 'base' value to a specified
// exponent. Both the 'base' and the 'exponent' are passed to this function
// as type BigIntNum.
//
// For example, raising the value 3.12 to the power of 4 means that the
// result will require at least 8-decimal places to the right of the
// decimal in order to display a correct result. In the following example
// with base ='3.12' and exponent = '4', this method will return '8'.
//
// 		Example: 3.12^4 = 94.75854336 (8-digits to the right of the decimal)
//
// The calculated minimum required precision is returned as type 'uint'.
//
// If the minimum required precision exceeds the maximum value for type
// 'uint' (+4,294,967,295, which equals 2^32 − 1), an error message is returned
// in addition to the maximum uint value (+4,294,967,295).
//
func (bIPwr BigIntMathPower) MinimumRequiredPrecision(
	base, exponent BigIntNum) (uint, error) {

	basePrecision := BigIntNum{}.NewUint(base.GetPrecisionUint(), 0)

	tExponent := exponent.CopyOut()

	if tExponent.GetSign() == -1 {
		tExponent.ChangeSign()
	}

	minRequiredPrecision :=
		BigIntMathMultiply{}.MultiplyBigIntNums(basePrecision, tExponent)

	minRequiredPrecision.RoundToDecPlace(0)

	return minRequiredPrecision.GetUInt()
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
//
// The return value, a type BigIntNum, represents the result of the base^exponent operation described above.
// This returned BigIntNum 'result' will contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter,'base'.
//
func (bIPwr BigIntMathPower) Pwr(base, exponent BigIntNum, maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntMathPower.Pwr() "

	result := BigIntNum{}.NewWithNumSeps(base.GetNumericSeparatorsDto())
	var err error

	if base.IsZero() {
		// If 'base' is zero, return zero value.
		return result, nil
	}

	if exponent.IsZero() {
		return BigIntNum{}.NewOne(0), nil
	}

	numSeps := base.GetNumericSeparatorsDto()

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
					fmt.Errorf(ePrefix+
						"Error returned by bIPwr.raiseToPositiveIntegerPower("+
						"base, exponent, maxPrecision) "+
						"base='%v' exponent='%v' Error='%v' ",
						base.GetNumStr(), exponent.GetNumStr(), err.Error())
			}

		} else {
			// exponent must be a negative number
			result, err = bIPwr.raiseToNegativeIntegerPower(base, exponent, maxPrecision)

			if err != nil {
				return BigIntNum{}.NewZero(0),
					fmt.Errorf(ePrefix+"Error returned by bIPwr.raiseToNegativeIntegerPower("+
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
					fmt.Errorf(ePrefix+
						" Error='%v' \n", err.Error())
			}

		} else {
			// fractional exponent must be a negative number

			result, err = bIPwr.raiseToNegativeFractionalPower(base, exponent, maxPrecision)

			if err != nil {
				return BigIntNum{},
					fmt.Errorf(ePrefix+
						"Error='%v' \n", err.Error())
			}

		}
	}

	if result.precision > maxPrecision {
		result.SetPrecision(maxPrecision)
	}

	err = result.SetNumericSeparatorsDto(numSeps)

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
			fmt.Errorf(ePrefix+"Error - The Exponent is an Integer Number. "+
				"Only Fractional Exponents can be processed by this method! exponent='%v' ",
				exponent.GetNumStr())
	}

	if exponent.GetSign() > 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix+"Error = The Exponent is a Positive Number. "+
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
			fmt.Errorf(ePrefix+
				"Error='%v' ", err.Error())
	}

	inverseResult, err := bINumResult.GetInverse(maxPrecision)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix+"Error returned by bINumResult.GetInverse(maxPrecision). "+
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
			fmt.Errorf(ePrefix+"Error - The Exponent is an Integer Number. "+
				"Only Fractional Exponents can be processed by this method! exponent='%v' ",
				exponent.GetNumStr())
	}

	if exponent.GetSign() < 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix+"Error = The Exponent is a Negative Number. "+
				"Only Positive Exponents can be processed by this method! exponent='%v'",
				exponent.GetNumStr())
	}

	modXZero := big.NewInt(0)

	scaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(exponent.precision)),
		modXZero)

	ratExponent := big.NewRat(1, 1).SetFrac(exponent.bigInt, scaleFactor)

	bINumNumerator := BigIntNum{}.NewBigInt(ratExponent.Num(), 0)

	pwr1Result, err := bIPwr.raiseToPositiveIntegerPower(base, bINumNumerator)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by bIPwr.raiseToPositive"+
				"IntegerPower(base, bINumNumerator)"+
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
				fmt.Errorf(ePrefix+
					"Error returned by nthRoot.IsEvenNumber() "+
					"nthRoot='%v' Error='%v'\n", nthRoot.GetNumStr(), err.Error())
		}

		if isEvenNum {
			return BigIntNum{}.NewZero(0),
				errors.New(ePrefix +
					"INVALID ENTRY - Cannot calculate nthRoot of a negative number when nthRoot is even. \n")
		}

	}

	// fmt.Println("pwr1Result: ", pwr1Result.GetNumStr())
	// fmt.Println("   nthRoot: ", nthRoot.GetNumStr())
	biNumResult, err := BigIntMathNthRoot{}.GetNthRoot(pwr1Result, nthRoot, maxPrecision)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by dBigIntMathNthRoot{}.NthRoot(...) "+
				"Error='%v'", err.Error())
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
			fmt.Errorf(ePrefix+"Error - The Exponent is a Fractional Number. "+
				"Only Integer Exponents can be processed by this method! exponent='%v' ",
				exponent.GetNumStr())
	}

	if exponent.GetSign() > 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix+"Error = The Exponent is a Positive Number. "+
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
			fmt.Errorf(ePrefix+"Error returned by bINumResult1.GetInverse(maxPrecision)"+
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
			fmt.Errorf(ePrefix+"Error - The Exponent is a Fractional Number. "+
				"Only Integer Exponents can be processed by this method! exponent='%v' ",
				exponent.GetNumStr())
	}

	if exponent.GetSign() < 0 {
		return BigIntNum{},
			fmt.Errorf(ePrefix+"Error = The Exponent is a Negative Number. "+
				"Only Positive Exponents can be processed by this method! exponent='%v'",
				exponent.GetNumStr())
	}

	bigIBasePrecision := big.NewInt(int64(base.GetPrecisionUint()))

	bigINewPrecision := big.NewInt(0).Mul(bigIBasePrecision, exponent.bigInt)

	newPrecision := uint(bigINewPrecision.Int64())
	//"           base: ", base.GetNumStr())
	// fmt.Println("    base bigInt: ", base.bigInt.Text(10))
	// fmt.Println("       exponent: ", exponent.GetNumStr())
	// fmt.Println("exponent.bigInt: ", exponent.bigInt.Text(10))

	result := big.NewInt(0).Exp(base.bigInt, exponent.bigInt, nil)

	return BigIntNum{}.NewBigInt(result, newPrecision), nil
}
