package mathops

import (
	"errors"
	"fmt"
	"math/big"
)

/*
	PermutationsUint64
  ============

	nPr	 =     n!
            ----
            (n-r)!

	n things taken r at a time

  Combinations
  ============

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


// PermutationsBigInt - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int64. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than
// 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsBigInt(numOfItems, numOfItemsPicked *big.Int) (BigIntNum, error) {

	ePrefix := "Probability.PermutationsBigInt() "
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

// PermutationsBigIntNum - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type BigIntNum. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than
// 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsBigIntNum(numOfItems, numOfItemsPicked BigIntNum) (BigIntNum, error) {

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

	return Probability{}.PermutationsBigInt(numOfItems.bigInt, numOfItemsPicked.bigInt)
}


// PermutationsInt - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than 
// 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsInt(numOfItems, numOfItemsPicked int) (BigIntNum, error) {
	
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
	
	return Probability{}.PermutationsBigInt(n, r)
	
}

// PermutationsInt32 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int32. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than 
// 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsInt32(numOfItems, numOfItemsPicked int32) (BigIntNum, error) {
	
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

	return Probability{}.PermutationsBigInt(n, r)
}

// PermutationsInt64 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type int64. Both input parameters
// must be non-zero, positive integer numbers. 'numOfItems' must be equal to or greater than 
// 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//	PermutationsUint64
//	============
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsInt64(numOfItems, numOfItemsPicked int64) (BigIntNum, error) {
	
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

	return Probability{}.PermutationsBigInt(n, r)
}


// PermutationsUint - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type uint. 'numOfItems'
// must be equal to or greater than 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsUint(numOfItems, numOfItemsPicked uint) (BigIntNum, error) {

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

	numOfItemsBI := big.NewInt(int64(numOfItems))
	numOfItemsPickedBI := big.NewInt( int64(numOfItemsPicked))

	return Probability{}.PermutationsBigInt(numOfItemsBI, numOfItemsPickedBI)
}

// PermutationsUint32 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type uint32. 'numOfItems'
// must be equal to or greater than 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsUint32(numOfItems, numOfItemsPicked uint32) (BigIntNum, error) {

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

	numOfItemsBI := big.NewInt(int64(numOfItems))
	numOfItemsPickedBI := big.NewInt( int64(numOfItemsPicked))

	return Probability{}.PermutationsBigInt(numOfItemsBI, numOfItemsPickedBI)
}


// PermutationsUint64 - Calculates the number of permutations associated with a collection of
// 'numOfItems' from which one picks 'numOfItemsPicked'. Order IS significant. Input
// parameters 'numOfItems' and 'numOfItemsPicked' are passed as type uint64. 'numOfItems'
// must be equal to or greater than 'numOfItemsPicked'.
//
// In the following permutation formula, n= 'numOfItems'  and r = 'numOfItemsPicked'
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
// Note: 0! = 1
//
func (prob Probability) PermutationsUint64(numOfItems, numOfItemsPicked uint64) (BigIntNum, error) {

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

	numOfItemsBI := big.NewInt(0).SetUint64(numOfItems)
	numOfItemsPickedBI := big.NewInt(0).SetUint64(numOfItemsPicked)

	return Probability{}.PermutationsBigInt(numOfItemsBI, numOfItemsPickedBI)
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
