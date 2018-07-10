package mathops

import (
	"math/big"
	"fmt"
)

type BigIntMath struct {
	Input 		*big.Int
	Output		*big.Int
}

// GetMagnitudeDigits - Returns the magnitude of a *big.Int number passed as an input
// parameter. Magnitude is defined here as the power of 10 which generates a value
// less than or equal to the 'target' *big.Int number.
//
// 			10^magnitude  <= target
//
// The value of magnitude is returned as a *big.Int number to the calling function.
//
// If target is a negative number, an error will be returned.
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

	base10 := big.NewInt(10)
	bigOne := big.NewInt(1)
	lastMag := big.NewInt(0)
	nextMag := big.NewInt(0)

	for true {

		val := big.NewInt(0).Exp(base10, nextMag, nil)

		if val.Cmp(target) == 1 {
			break
		}

		lastMag = big.NewInt(0).Set(nextMag)
		nextMag = big.NewInt(0).Add(bigOne, nextMag)

	}

	return lastMag, nil
}