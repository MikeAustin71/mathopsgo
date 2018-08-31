package mathops

import (
	"errors"
	"fmt"
	"math/big"
)

/*
https://www.mathsisfun.com/combinatorics/combinations-permutations.html

	Permutations
  ============
  This calculation assumes NO REPETITIONS!

  n things taken r at a time

	nPr	 =     n!
            ----
            (n-r)!





  Combinations
  ============
 This calculation assumes NO REPETITIONS!

	nCr		=            n!
                -----------
								(n-r)! (r!)


possibilities
	5 coin tosses = 2^5 possibilities
  probability of getting k-heads in n flips.


 */

type Probability struct {
	  NoOfTrials                   BigIntNum
		PossibleOutcomesPerTrial     BigIntNum
		NoSuccessfulOutcomesPerTrial BigIntNum
	  TotPossibleOutcomes					 BigIntNum
	  TotSuccessfulOutcomes				 BigIntNum
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
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//              						n!
// 							nPr	 =		------
// 												(n-r)!
//             ----------------------
// 										Note: 0! = 1
//
// 				Where n is the number of things to choose from,
//				and we choose r of them, repetition is NOT allowed,
// 				and order matters.
//
// *** This calculation assumes NO REPETITIONS! ***
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
				fmt.Errorf(ePrefix + "Error returned by NFactorial{}.CalcFactorialValueBigInt(" +
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
			fmt.Errorf(ePrefix + "Error returned by NFactorial{}.CalcFactorialValueBigInt(" +
				"nUpperLimit, nLowerLimit). upperLimit='%v' lowerLimit='%v' Error='%v'. \n",
				nUpperLimit.Text(10), nLowerLimit.Text(10), err.Error())
	}

	return result, nil
}

// PermutationsWithRepsBigInt - Calculates the number of permutations associated with a collection
// of 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. This calculation
// assumes that REPETITIONS ARE ALLOWED.
//
// Input parameters 'numOfItems' and 'numOfItemsPicked' are passed as type *big.Int. Both input
// parameters must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater
// than 'numOfItemsPicked'.
//
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

	n :=  big.NewInt(int64(numOfItems))
	r :=  big.NewInt(int64(numOfItemsPicked))

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

	n :=  big.NewInt(int64(numOfItems))
	r :=  big.NewInt(int64(numOfItemsPicked))

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

	n :=  big.NewInt(numOfItems)
	r :=  big.NewInt(numOfItemsPicked)

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
	r := big.NewInt( int64(numOfItemsPicked))

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
	r := big.NewInt( int64(numOfItemsPicked))

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
	r := big.NewInt( int64(numOfItemsPicked))

	if !allowRepetitions {
		return Probability{}.PermutationsNoRepsBigInt(n, r)
	}

	return Probability{}.PermutationsWithRepsBigInt(n, r)
}

// Combinations - Calculates the number of combinations associated with a collection of 'numOfItems'
// from which one chooses 'numOfItemsChosen'. Order is NOT significant.
//
//	Combinations
//	============
//
//                 n!
// nCr	 =		-----------
// 						 (n-r)! r!
//
//
// Note: 0! = 1
//
// *** This calculation assumes NO REPETITIONS! ***
//
func (prob Probability) Combinations(numOfItems, numOfItemsChosen uint64) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsUint64() "

	if numOfItems < numOfItemsChosen {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfItems' is LESS THAN 'numOfItemsChosen'! ")

	}

	//          5!           5!         5!
	//      ---------  =  -------  =  ----- = 1
	//      (5-5)! 5!      0! 5!        5!

	if numOfItems == numOfItemsChosen {
		return BigIntNum{}.NewOne(0), nil
	}

	// Numerator
	n := FactorialDto{}
	n.UpperLimit = numOfItems
	n.LowerLimit = 1

	// Denominator

	nMinusR := FactorialDto{}

	if numOfItems == numOfItemsChosen {
		nMinusR.UpperLimit = 1
	} else {
		nMinusR.UpperLimit = numOfItems - numOfItemsChosen
	}

	nMinusR.LowerLimit = 1

	if nMinusR.UpperLimit > 1 {

		n.LowerLimit = nMinusR.UpperLimit

	}


	upperLimit := big.NewInt(0).SetUint64(n.UpperLimit)
	lowerLimit := big.NewInt(0).SetUint64(n.LowerLimit)

	numerator, err := NFactorial{}.CalcFactorialValueBigInt(upperLimit, lowerLimit)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error returned by NFactorial{}.CalcFactorialValueBigInt(...). " +
				"upperLimit='%v' lowerLimit='%v' Error='%v'. \n",n.UpperLimit, n.LowerLimit, err.Error())
	}

	r := FactorialDto{}
	r.UpperLimit = numOfItemsChosen
	r.LowerLimit = 1

	if r.UpperLimit < 2 {
		return numerator, nil
	}



	return numerator, nil
}
