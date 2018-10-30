package mathops

import (
	"errors"
	"fmt"
	"math/big"
)

type BigIntMath struct {
	Input  *big.Int
	Output *big.Int
}

// GetMagnitude - Returns the magnitude of a *big.Int number passed as an input
// parameter ('initialValue'. Magnitude is defined here as the power of 10 which generates a value
// less than or equal to the 'target' *big.Int number.
//
// 													10^magnitude  <= initialValue
//
// The value of magnitude is returned as a *big.Int number to the calling function.
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
// Input Parameters
// ================
//
// initialValue	*big.Int 	- An integer of type *big.Int. This method will analyze this
//                         	integer and return it's magnitude.
//
// Return Values
// =============
//
// magnitude 		*big.Int	- 10 raised to the power of magnitude will yield a value which
//                          is less than or equal to the input parameter 'initialValue'.
//
// err					error			- If during the completion of this calculation an error is encountered,
//           								return value 'magnitude' will be set to zero and this error object
//                          will be populated with an appropriate error message. If the method
//                          completes successfully, this return value, 'err' will be set to
//                          'nil'.
//
func (bIntMath BigIntMath) GetMagnitude(initialValue *big.Int) (magnitude *big.Int, err error) {

	ePrefix := "BigIntMath) GetMagnitude() "

	magnitude = big.NewInt(0)
	err = nil

	if initialValue == nil {
		err = errors.New(ePrefix + "Error: 'initialValue' is NOT Initialized! initialValue=nil")
		return magnitude, err
	}

	target := big.NewInt(0).Set(initialValue)
	compareResult := target.Cmp(big.NewInt(0))

	// Convert to absolute value
	if compareResult == -1 {
		target.Neg(target)
	}

	if compareResult == 0 {
		// magnitude = 0; err=nil
		return magnitude, err
	}

	bigTen := big.NewInt(10)

	if target.Cmp(bigTen) == -1 {
		// magnitude = 0; err=nil
		return magnitude, err
	}

	bitLen := target.BitLen()

	if bitLen <= 0 {
		err = fmt.Errorf(ePrefix+"Error: target.BitLen() = %v - Negative value!",
			bitLen)
		return magnitude, err
	}

	tenToPowerPrecision := big.NewInt(0)
	bigZero := big.NewInt(0)

	// Note: logTwo is a global constant

	var errx error

	// ************************************
	// target MUST BE <= 2^(bit length)
	// ************************************
	magnitude, tenToPowerPrecision, errx =
		BigIntMathMultiply{}.BigIntMultiply(
			big.NewInt(int64(bitLen)),
			big.NewInt(0),
			logTwo.GetInteger(),
			logTwo.GetPrecisionBigInt())

	if errx != nil {
		err = fmt.Errorf(ePrefix +
			"%v", errx.Error())
		magnitude = big.NewInt(0)
		return magnitude, err
	}

	if tenToPowerPrecision.Cmp(bigZero) == 1 {
		scale :=
			big.NewInt(0).Exp(
				bigTen,
				tenToPowerPrecision,
				nil)

		magnitude.Quo(magnitude, scale)
	}

	testNum := big.NewInt(0).Exp(bigTen, magnitude, nil)

	if testNum.Cmp(target) == 1 {
		magnitude.Sub(magnitude, big.NewInt(1))
	}

	return magnitude, err
}


// RoundToMaxPrecision - Applies maximum precision to a *big.Int number
// and associated numeric precision, 'bigIntNum' and 'bigIntNumPrecision'.
//
// Precision as used here defines the number of digits to the right of the
// decimal place. If 'bigIntNumPrecision' exceeds 'maxPrecision', the
// 'bigIntNum' and 'bigIntNumPrecision' pair are rounded to 'maxPrecision'
// and returned as 'result' and 'resultPrecision'.
//
// Examples:
// =========
//
//  bigIntNum	bigIntNumPrecision	maxPrecision	result		resultPrecision
//	5255						3                  	2					526					2
//  52671						4										6					52671				4
//
func (bIntMath BigIntMath) RoundToMaxPrecision(
	bigIntNum,
	bigIntNumPrecision,
	maxPrecision *big.Int) (result, resultPrecision *big.Int, err error) {

	ePrefix := "BigIntMath.RoundToMaxPrecision() "
	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil


	if bigIntNum == nil {
		bigIntNum = big.NewInt(0)
		bigIntNumPrecision = big.NewInt(0)
	}

	bigZero := big.NewInt(0)

	if bigIntNum.Cmp(bigZero) == 0 {
		return result, resultPrecision, err
	}

	if bigIntNumPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Input parameter 'bigIntNumPrecision' is LESS THAN ZERO! " +
			"bigIntNumPrecision='%v' ", bigIntNumPrecision.Text(10))
		return result, resultPrecision, err
	}

	result = big.NewInt(0).Set(bigIntNum)
	resultPrecision = big.NewInt(0).Set(bigIntNumPrecision)

	if resultPrecision.Cmp(maxPrecision) == 1 {
		delta := big.NewInt(0).Sub(resultPrecision, maxPrecision)
		delta.Sub(delta, big.NewInt(1))
		bigTen := big.NewInt(10)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		result.Quo(result, scale)
		bigFive := big.NewInt(5)
		if result.Cmp(bigZero) == -1 {
			bigFive.Neg(bigFive)
		}
		result.Add(result, bigFive)
		result.Quo(result, bigTen)
		resultPrecision = big.NewInt(0).Set(maxPrecision)
	}

	err = nil

	return result, resultPrecision, err
}