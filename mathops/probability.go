package mathops

import (
	"fmt"
	"errors"
	"math/big"
)

/*
	Permutations
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
		PercentCertainty 	BigIntNum
		NoOfConstraintOutcomes BigIntNum
		NoOfPossibleOutcomes   BigIntNum
}


// Permutations - Calculates the number of permutations associated with a specified number of
// possible outcomes and a given sub-set of those outcomes which meet particular constraints.
//
//	Permutations
//	============
//
//              n!
// nPr	 =		------
// 						(n-r)!
//
//
func (prob Probability) Permutations(numPossibleOutcomes, numOfConstraintOutcomes uint64) (BigIntNum, error) {

	ePrefix := "Probability.Permutations() "

	if numOfConstraintOutcomes > numPossibleOutcomes {
		return BigIntNum{}.NewZero(0),
			errors.New(ePrefix + "Error: 'numOfPossibleOutcomes' exceeds 'numOfConstraintOutcomes'! ")

	}

	n := FactorialDto{}
	n.UpperLimit = numPossibleOutcomes
	n.LowerLimit = 1

	nMinusR := FactorialDto{}

	if numPossibleOutcomes == numOfConstraintOutcomes {
		nMinusR.UpperLimit = 1
	} else {
		nMinusR.UpperLimit = numPossibleOutcomes - numOfConstraintOutcomes
	}

	nMinusR.LowerLimit = 1

	n.LowerLimit = nMinusR.UpperLimit

	upperLimit := big.NewInt(0).SetUint64(n.UpperLimit)
	lowerLimit := big.NewInt(0).SetUint64(n.LowerLimit)

	result, err := NFactorial{}.CalcFactorialValueBigInt(upperLimit, lowerLimit)

	if err != nil {
		return BigIntNum{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error returned by NFactorial{}.CalcFactorialValueBigInt(...). " +
				"upperLimit='%v' lowerLimit='%v' Error='%v'. \n",n.UpperLimit, n.LowerLimit, err.Error())
	}

	return result, nil
}
