package mathops

import (
	"math/big"
	"fmt"
)

type BigIntMath struct {
	Input 		*big.Int
	Output		*big.Int
}

// GetMagnitude - Returns the magnitude of a *big.Int number passed as an input
// parameter. Magnitude is defined here as the power of 10 which generates a value
// less than or equal to the 'target' *big.Int number.
//
// 													10^magnitude  <= target
//
// The value of magnitude is returned as a *big.Int number to the calling function.
//
// If target is a negative number, an error will be returned.
//
// WARNING: This method fails on 'target' numbers with values greater than 10^10,000,000
//
//
// Examples:
// =========
//
//			   			target				magnitude
//              ------        ---------
//
//			  		 963,256						5
//									 2						0
//									32						1
// 			 8,456,123,921					  9
//
//
//
func (bIntMath BigIntMath) GetMagnitude(target *big.Int) (*big.Int, error) {

	compareResult := target.Cmp(big.NewInt(0))

	if compareResult == -1 {
		ePrefix := "BigIntMath.GetMagnitude() "
		return big.NewInt(0),
			fmt.Errorf(ePrefix + "Error: Input parameter 'target' is a negative number. " +
				"target='%v' ", target.Text(10))
	}

	if compareResult == 0 {
		return big.NewInt(0), nil
	}

	// Conversion constant 3.3219281
	conversionConst := BigIntNum{}.NewInt64Exponent(33219281, -7 )
	biNumBitLen := BigIntNum{}.NewInt64Exponent(int64(target.BitLen()), 0)

	bINumNextMag, _ :=
		BigIntMathDivide{}.BigIntNumIntQuotient(biNumBitLen, conversionConst)

	base10 := big.NewInt(10)
	bigOne := big.NewInt(1)
	lastMag := big.NewInt(0).Sub(bINumNextMag.bigInt, big.NewInt(1))

	for true {

		val := big.NewInt(0).Exp(base10, bINumNextMag.bigInt, nil)

		if val.Cmp(target) == 1 {
			break
		}

		lastMag = big.NewInt(0).Set(bINumNextMag.bigInt)
		bINumNextMag.bigInt = big.NewInt(0).Add(bINumNextMag.bigInt, bigOne)

	}

	return lastMag, nil
}
