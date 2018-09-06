package mathops

import (
	"errors"
	"fmt"
	"math/big"
)

/*

This file contains Probability source code for 'Permutations'.  For 'Combinations' see
source code file MikeAustin71\mathopsgo\mathops\probability_02.go

														Permutations
  													============

Reference
#########

	https://www.mathsisfun.com/combinatorics/combinations-permutations.html
	https://www.youtube.com/watch?v=XqQTXW7XfYA&list=PL06A16C388F14E6FE&index=21
	https://www.probabilitycourse.com/courses.php

$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
  *** This permutation calculation assumes NO REPETITIONS! ***
$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

 In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'

              						n!
 							nPr	 =		------
 												(n-r)!
             ----------------------

 				Where n is the number of things to choose from,
				and we choose r of them, repetition is NOT allowed,
 				and order matters.


$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$
 *** This Permutation calculation assumes REPETITIONS ARE ALLOWED! ***
$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$

In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'


 									nPr	 =		n^r
                  -------------

 				Where n is the number of things to choose from,
				and we choose r of them, repetition is allowed,
 				and order matters.


*/

type Probability struct {
	NoOfTrials                   BigIntNum
	PossibleOutcomesPerTrial     BigIntNum
	NoSuccessfulOutcomesPerTrial BigIntNum
	TotPossibleOutcomes          BigIntNum
	TotSuccessfulOutcomes        BigIntNum
	PercentCertainty             BigIntNum
}

// PermutationsNoRepsBigInt - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. This calculation
// assumes that repetitions ARE NOT allowed.
//
// Input parameters 'numOfItems' and 'numOfItemsPicked' are both passed as type *big.Int. Both
// input parameters must be non-zero, positive integer numbers. 'numOfItems' must be equal to or
// greater than 'numOfItemsPicked'.
//
// Input Parameters
// ================
//
// numOfItems 				*big.Int	- Must be a positive integer number greater than zero.
//                                'numOfItems' must be greater than or equal to
// 																'numOfItemsPicked'.
//
// numOfItemsPicked 	*big.Int	- Must be a positive integer number greater than zero.
// 																'numOfItemsChosen' must be less than or equal to
// 																'numOfItems'.
//
//
// Returns
// =======
//
// BigIntNum			- If the calculation is successful, the result is returned as a
//                  BigIntNum type. If the calculation fails, the error return is
//                  populated.
//
// error					- If the calculation is successful, this return value is 'nil'. If
//                  the calculation fails, 'error' is populated with an appropriate
//                  error message.
//
// Calculation
// ===========
//
// In the following permutation formula, n= 'numOfItems' and r = 'numOfItemsPicked'
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 				Where n is the number of things to select from,
//				and we pick r of them, repetition is NOT allowed,
// 				and order matters.
//
// *** This calculation assumes NO REPETITIONS! ***
//
// Example
// =======
// How many ways can first and second place be awarded to 10 people?
//
//                         10!
// Answer = 10P2 =       ------ =   10 x 9 = 90
//                         8!
//
// Note: 0! = 1
//
func (prob Probability) PermutationsNoRepsBigInt(
	numOfItems, numOfItemsPicked *big.Int) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsNoRepsBigInt() "
	bigZero := big.NewInt(0)

	if numOfItems.Cmp(bigZero) == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItemsPicked.Cmp(bigZero) == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItems.Cmp(bigZero) < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItemsPicked.Cmp(bigZero) < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	if numOfItems.Cmp(numOfItemsPicked) < 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'! ")

	}

	// Numerator

	nUpperLimit := big.NewInt(0).Set(numOfItems)
	nLowerLimit := big.NewInt(1)

	if numOfItems.Cmp(numOfItemsPicked) == 0 {
		//                             n!
		// This is equivalent to   ------------
		//                             1

		result, err := NFactorial{}.CalcFactorialValueBigInt(nUpperLimit, nLowerLimit)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix+"Error returned by NFactorial{}.CalcFactorialValueBigInt("+
					"nUpperLimit, nLowerLimit). upperLimit='%v' lowerLimit='%v' Error='%v'. \n",
					numOfItems, 1, err.Error())
		}

		return result, nil

	}

	// numOfItems MUST BE GREATER THAN numOfItemsPicked

	nLowerLimit = big.NewInt(0).Sub(nUpperLimit, numOfItemsPicked)

	result, err := NFactorial{}.CalcFactorialValueBigInt(nUpperLimit, nLowerLimit)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by NFactorial{}.CalcFactorialValueBigInt("+
				"nUpperLimit, nLowerLimit). upperLimit='%v' lowerLimit='%v' Error='%v'. \n",
				nUpperLimit.Text(10), nLowerLimit.Text(10), err.Error())
	}

	return result, nil
}

// PermutationsWithRepsBigInt - Calculates the number of permutations associated with
// a collection of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS
// significant. This calculation assumes that REPETITIONS ARE ALLOWED.
//
// Input Parameters
// ================
//
// numOfItems 				*big.Int	- Must be a positive integer number greater than zero.
//                                'numOfItems' must be greater than or equal to
// 																'numOfItemsPicked'.
//
// numOfItemsPicked 	*big.Int	- Must be a positive integer number greater than zero.
// 																'numOfItemsChosen' must be less than or equal to
// 																'numOfItems'.
//
// Returns
// =======
//
// BigIntNum			- If the calculation is successful, the result is returned as a
//                  BigIntNum type. If the calculation fails, the error return is
//                  populated.
//
// error					- If the calculation is successful, this return value is 'nil'. If
//                  the calculation fails, 'error' is populated with an appropriate
//                  error message.
//
// Calculation
// ===========
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is allowed,
// 				and order matters.
//
//
// *** This calculation assumes REPETITIONS ARE ALLOWED! ***
//
//  Example
//  =======
// 				Combination lock with 3 numbers: there are 10 numbers to choose from
// 				(0,1,2,3,4,5,6,7,8,9) and we choose 3 of them (repetitions allowed):
// 							10 × 10 × ... (3 times) = 10^3 = 1,000 permutations
//
func (prob Probability) PermutationsWithRepsBigInt(
	numOfItems, numOfItemsPicked *big.Int) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsNoRepsBigInt() "
	bigZero := big.NewInt(0)

	if numOfItems.Cmp(bigZero) == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItemsPicked.Cmp(bigZero) == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItems.Cmp(bigZero) < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItemsPicked.Cmp(bigZero) < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	if numOfItems.Cmp(numOfItemsPicked) < 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'! ")

	}

	resultBigInt := big.NewInt(0).Exp(numOfItems, numOfItemsPicked, nil)

	result := BigIntNum{}.NewBigInt(resultBigInt, 0)

	return result, nil
}

// PermutationsBigIntNum - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input parameters
// 'numOfItems' and 'numOfItemsPicked' are passed as type BigIntNum. Both input parameters must
// be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than
// 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition IS allowed,
// 				and order matters.
//
func (prob Probability) PermutationsBigIntNum(
	numOfItems, numOfItemsPicked BigIntNum,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsBigIntNum() "

	if numOfItems.IsZero() {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItemsPicked.IsZero() {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItems.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItemsPicked.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	if numOfItems.Cmp(numOfItemsPicked) < 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'! ")

	}

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(numOfItems.bigInt, numOfItemsPicked.bigInt)
	}

	return Probability{}.PermutationsWithRepsBigInt(numOfItems.bigInt, numOfItemsPicked.bigInt)
}

// PermutationsDecimal - Calculates the number of permutations associated with a collection
// of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type Decimal. Both input
// parameters must be non-zero, positive integer numbers. 'numOfItems' must be equal to or
// greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type Decimal.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition IS allowed,
// 				and order matters.
//
func (prob Probability) PermutationsDecimal(
	numOfItems, numOfItemsPicked Decimal,
	allowRepetitions bool) (Decimal, error) {

	ePrefix := "Probability.PermutationsDecimal() "

	if numOfItems.GetSign() == -1 {
		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItems.IsZero() {
		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItems.GetPrecisionUint() > 0 {
		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is NOT an Integer!")
	}

	if numOfItemsPicked.GetSign() == -1 {
		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	if numOfItemsPicked.IsZero() {
		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItemsPicked.GetPrecisionUint() > 0 {
		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!")
	}

	n, err := numOfItems.GetBigInt()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItems.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	r, err := numOfItemsPicked.GetBigInt()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItemsPicked.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	if r.Cmp(n) == 1 {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.  "+
				"numOfItems='%v' numOfItemsPicked='%v' \n",
				numOfItems.GetNumStr(), numOfItemsPicked.GetNumStr())
	}

	var result BigIntNum

	if !allowRepetitions {

		result, err = Probability{}.PermutationsNoRepsBigInt(n, r)

		if err != nil {
			return Decimal{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsNoRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	} else {

		result, err = Probability{}.PermutationsWithRepsBigInt(n, r)

		if err != nil {
			return Decimal{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsWithRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	}

	iaDecimal, err := result.GetDecimal()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by result.GetDecimal(). "+
				"Error='%v' \n", err.Error())
	}

	return iaDecimal, nil
}

// PermutationsIntAry - Calculates the number of permutations associated with a collection
// of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type IntAry. Both input
// parameters must be non-zero, positive integer numbers. 'numOfItems' must be equal to or
// greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type IntAry.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition IS allowed,
// 				and order matters.
//
//
//
func (prob Probability) PermutationsIntAry(
	numOfItems, numOfItemsPicked IntAry,
	allowRepetitions bool) (IntAry, error) {

	ePrefix := "Probability.PermutationsIntAry() "

	if numOfItems.GetSign() == -1 {
		return IntAry{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItems.IsZero() {
		return IntAry{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItems.GetPrecisionUint() > 0 {
		return IntAry{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is NOT an Integer!")
	}

	if numOfItemsPicked.GetSign() == -1 {
		return IntAry{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	if numOfItemsPicked.IsZero() {
		return IntAry{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItemsPicked.GetPrecisionUint() > 0 {
		return IntAry{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!")
	}

	n, err := numOfItems.GetBigInt()

	if err != nil {
		return IntAry{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItems.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	r, err := numOfItemsPicked.GetBigInt()

	if err != nil {
		return IntAry{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItemsPicked.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	if r.Cmp(n) == 1 {
		return IntAry{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.  "+
				"numOfItems='%v' numOfItemsPicked='%v' \n",
				numOfItems.GetNumStr(), numOfItemsPicked.GetNumStr())
	}

	var result BigIntNum

	if !allowRepetitions {

		result, err = Probability{}.PermutationsNoRepsBigInt(n, r)

		if err != nil {
			return IntAry{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsNoRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	} else {

		result, err = Probability{}.PermutationsWithRepsBigInt(n, r)

		if err != nil {
			return IntAry{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsWithRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	}

	iaResult, err := result.GetIntAry()

	if err != nil {
		return IntAry{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by result.GetIntAry(). "+
				"Error='%v' \n", err.Error())
	}

	return iaResult, nil
}

// PermutationsINumMgr - Calculates the number of permutations associated with a collection
// of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type INumMgr. Both input
// parameters must be non-zero, positive integer numbers. 'numOfItems' must be equal to or
// greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition IS allowed,
// 				and order matters.
//
//
//
func (prob Probability) PermutationsINumMgr(
	numOfItems, numOfItemsPicked INumMgr,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsINumMgr() "

	if numOfItems.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItems.IsZero() {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItems.GetPrecisionUint() > 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is NOT an Integer!")
	}

	if numOfItemsPicked.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	if numOfItemsPicked.IsZero() {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItemsPicked.GetPrecisionUint() > 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!")
	}

	n, err := numOfItems.GetBigInt()

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItems.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	r, err := numOfItemsPicked.GetBigInt()

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItemsPicked.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	if r.Cmp(n) == 1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.  "+
				"numOfItems='%v' numOfItemsPicked='%v' \n",
				numOfItems.GetNumStr(), numOfItemsPicked.GetNumStr())
	}

	var result BigIntNum

	if !allowRepetitions {

		result, err = Probability{}.PermutationsNoRepsBigInt(n, r)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsNoRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	} else {

		result, err = Probability{}.PermutationsWithRepsBigInt(n, r)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsWithRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	}

	return result, nil
}

// PermutationsNumStrDto - Calculates the number of permutations associated with a collection
// of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type NumStrDto. Both input
// parameters must be non-zero, positive integer numbers. 'numOfItems' must be equal to or
// greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type NumStrDto.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition IS allowed,
// 				and order matters.
//
//
//
func (prob Probability) PermutationsNumStrDto(
	numOfItems, numOfItemsPicked NumStrDto,
	allowRepetitions bool) (NumStrDto, error) {

	ePrefix := "Probability.PermutationsNumStrDto() "

	if numOfItems.GetSign() == -1 {
		return NumStrDto{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItems.IsZero() {
		return NumStrDto{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItems.GetPrecisionUint() > 0 {
		return NumStrDto{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is NOT an Integer!")
	}

	if numOfItemsPicked.GetSign() == -1 {
		return NumStrDto{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	if numOfItemsPicked.IsZero() {
		return NumStrDto{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItemsPicked.GetPrecisionUint() > 0 {
		return NumStrDto{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!")
	}

	n, err := numOfItems.GetBigInt()

	if err != nil {
		return NumStrDto{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItems.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	r, err := numOfItemsPicked.GetBigInt()

	if err != nil {
		return NumStrDto{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by numOfItemsPicked.GetBigInt(). "+
				"Error='%v' \n", err.Error())
	}

	if r.Cmp(n) == 1 {
		return NumStrDto{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.  "+
				"numOfItems='%v' numOfItemsPicked='%v' \n",
				numOfItems.GetNumStr(), numOfItemsPicked.GetNumStr())
	}

	var result BigIntNum

	if !allowRepetitions {

		result, err = Probability{}.PermutationsNoRepsBigInt(n, r)

		if err != nil {
			return NumStrDto{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsNoRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	} else {

		result, err = Probability{}.PermutationsWithRepsBigInt(n, r)

		if err != nil {
			return NumStrDto{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsWithRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	}

	iaResult, err := result.GetNumStrDto()

	if err != nil {
		return NumStrDto{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by result.GetNumStrDto(). "+
				"Error='%v' \n", err.Error())
	}

	return iaResult, nil
}

// PermutationsNumberStr - Calculates the number of permutations associated with a collection
// of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant.
//
// Input parameters 'numOfItems' and 'numOfItemsPicked' are passed as strings. These strings
// must be formatted as valid number strings. Number strings may be prefixed by a plus (+) or
// minus (-) and must consist of a string of numeric digits which may be delimited by a thousands
// separator. If the numeric value is a fractional value, the fractional digits must be
// preceded by a period ('.') or decimal separator.
//
// Both input parameters must be non-zero, positive integer numbers. 'numOfItems' must be equal
// to or greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition IS allowed,
// 				and order matters.
//
func (prob Probability) PermutationsNumberStr(
	numOfItems, numOfItemsPicked string,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsNumberStr() "

	nBigINum, err := BigIntNum{}.NewNumStr(numOfItems)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by BigIntNum{}.NewNumStr(numOfItems) "+
				"Error='%v' \n", err.Error())
	}

	rBigINum, err := BigIntNum{}.NewNumStr(numOfItemsPicked)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by BigIntNum{}.NewNumStr(numOfItemsPicked) "+
				"Error='%v' \n", err.Error())
	}

	if nBigINum.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error: Input parameter 'numOfItems' is LESS THAN ZERO! "+
				"numOfItems='%v'", numOfItems)
	}

	if nBigINum.IsZero() {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if nBigINum.GetPrecisionUint() > 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItems' is NOT an Integer!")
	}

	if rBigINum.GetSign() == -1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO! "+
				"numOfItemsPicked='%v' ", numOfItemsPicked)
	}

	if rBigINum.IsZero() {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if rBigINum.GetPrecisionUint() > 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: Input parameter 'numOfItemsPicked' is NOT an Integer!")
	}

	if rBigINum.Cmp(nBigINum) == 1 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: 'numOfItemsPicked' is GREATER THAN 'numOfItems'.  "+
				"numOfItems='%v' numOfItemsPicked='%v' \n",
				numOfItems, numOfItemsPicked)
	}

	var result BigIntNum

	if !allowRepetitions {

		result, err = Probability{}.PermutationsNoRepsBigInt(nBigINum.bigInt, rBigINum.bigInt)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsNoRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	} else {

		result, err = Probability{}.PermutationsWithRepsBigInt(nBigINum.bigInt, rBigINum.bigInt)

		if err != nil {
			return BigIntNum{}.NewZero(0),
				fmt.Errorf(ePrefix+
					"Error returned by Probability{}.PermutationsWithRepsBigInt(numOfItems, numOfItemsPicked). "+
					"Error='%v' \n", err.Error())
		}

	}

	return result, nil
}

// PermutationsInt - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than
// 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition IS allowed,
// 				and order matters.
//
//
//
func (prob Probability) PermutationsInt(
	numOfItems, numOfItemsPicked int,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsInt() "

	if numOfItems == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItemsPicked == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItems < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItemsPicked < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	n := big.NewInt(int64(numOfItems))
	r := big.NewInt(int64(numOfItemsPicked))

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(n, r)
	}

	return Probability{}.PermutationsWithRepsBigInt(n, r)
}

// PermutationsInt32 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int32. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than
// 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//           			-----------------
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is allowed,
// 				and order matters.
//
func (prob Probability) PermutationsInt32(
	numOfItems, numOfItemsPicked int32,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsInt32() "

	if numOfItems == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItemsPicked == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItems < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItemsPicked < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	n := big.NewInt(int64(numOfItems))
	r := big.NewInt(int64(numOfItemsPicked))

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(n, r)
	}

	return Probability{}.PermutationsWithRepsBigInt(n, r)
}

// PermutationsInt64 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int64. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than
// 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is allowed,
// 				and order matters.
//
func (prob Probability) PermutationsInt64(
	numOfItems, numOfItemsPicked int64,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsInt64() "

	if numOfItems == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is ZERO!")
	}

	if numOfItemsPicked == 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is ZERO!")
	}

	if numOfItems < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItems' is LESS THAN ZERO!")
	}

	if numOfItemsPicked < 0 {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'numOfItemsPicked' is LESS THAN ZERO!")
	}

	n := big.NewInt(numOfItems)
	r := big.NewInt(numOfItemsPicked)

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(n, r)
	}

	return Probability{}.PermutationsWithRepsBigInt(n, r)
}

// PermutationsUint - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type uint. 'numOfItems'
// must be equal to or greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is allowed,
// 				and order matters.
//
func (prob Probability) PermutationsUint(
	numOfItems, numOfItemsPicked uint,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsUint() "

	if numOfItems == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is ZERO! ")
	}

	if numOfItemsPicked == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItemsPicked' is ZERO! ")
	}

	if numOfItems < numOfItemsPicked {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'! ")

	}

	n := big.NewInt(int64(numOfItems))
	r := big.NewInt(int64(numOfItemsPicked))

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(n, r)
	}

	return Probability{}.PermutationsWithRepsBigInt(n, r)
}

// PermutationsUint32 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type uint32. 'numOfItems'
// must be equal to or greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is allowed,
// 				and order matters.
//
func (prob Probability) PermutationsUint32(
	numOfItems, numOfItemsPicked uint32,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsUint32() "

	if numOfItems == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is ZERO! ")
	}

	if numOfItemsPicked == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItemsPicked' is ZERO! ")
	}

	if numOfItems < numOfItemsPicked {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'! ")

	}

	n := big.NewInt(int64(numOfItems))
	r := big.NewInt(int64(numOfItemsPicked))

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(n, r)
	}

	return Probability{}.PermutationsWithRepsBigInt(n, r)
}

// PermutationsUint64 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type uint64. 'numOfItems'
// must be equal to or greater than 'numOfItemsPicked'.
//
// The input parameter 'allowRepetitions' is a boolean value which will determine whether the
// calculation results will allow repetitions or not. The formula for the permutation will
// therefore vary depending on whether repetitions are allowed.
//
// The result of this permutation calculation is returned as a type BigIntNum.
//
// In the following permutation formulas, n= 'numOfItems'  and r = 'numOfItemsPicked'. The actual
// formula applied depends on whether input parameter 'allowRepetitions' is 'true' or 'false'.
//
//      ====================================================================
// 			'allowRepetitions' = false
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
//
// 							Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
//      ====================================================================
// 			'allowRepetitions' = true
//
// 									nPr	 =		n^r
//                  -------------
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is allowed,
// 				and order matters.
//
func (prob Probability) PermutationsUint64(
	numOfItems, numOfItemsPicked uint64,
	allowRepetitions bool) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsUint64() "

	if numOfItems == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is ZERO! ")
	}

	if numOfItemsPicked == 0 {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItemsPicked' is ZERO! ")
	}

	if numOfItems < numOfItemsPicked {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is LESS THAN 'numOfItemsPicked'! ")

	}

	n := big.NewInt(int64(numOfItems))
	r := big.NewInt(int64(numOfItemsPicked))

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(n, r)
	}

	return Probability{}.PermutationsWithRepsBigInt(n, r)
}
