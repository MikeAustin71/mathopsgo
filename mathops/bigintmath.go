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

	bitLen := target.BitLen()

	if bitLen <= 0 {
		err = fmt.Errorf(ePrefix + "Error: target.BitLen() = %v - Negative value!",
			bitLen)
		return magnitude, err
	}

	if bitLen > -1 && bitLen < 4  {
		// magnitude == 0  err==nil
		return magnitude, err
	}

	bigTen := big.NewInt(10)

	tenToPowerPrecision := uint(0)

	// Note: logTwo is a global constant
	// ************************************
	// target MUST BE <= 2^(bit length + 1)
	// ************************************
	//bitLen++

	magnitude, tenToPowerPrecision =
		BigIntMathMultiply{}.BigIntMultiply(
			big.NewInt(int64(bitLen)),
			0,
			logTwo.GetInteger(),
			logTwo.GetPrecision())


	if tenToPowerPrecision > 0 {
		scale:=
			big.NewInt(0).Exp(
				bigTen,
				big.NewInt(int64(tenToPowerPrecision)),
				nil)

		magnitude.Quo(magnitude, scale)
	}

	testNum := big.NewInt(0).Exp(bigTen,magnitude, nil)

	if testNum.Cmp(target) == 1 {
		magnitude.Sub(magnitude,big.NewInt(1))
	}

	return magnitude, err
}

/*
func (bIntMath BigIntMath) GetMagnitude(initialValue *big.Int) (magnitude *big.Int, err error) {

	magnitude = big.NewInt(0)
	err = nil

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

	// Conversion constant 3.3219281
	conversionConst := big.NewInt(int64(33219281))

	numBitLen := big.NewInt (int64(target.BitLen()))

	factor, factorPrecision, _ := BigIntMathDivide{}.BigIntFracQuotient(
		numBitLen, 0,  conversionConst, 7, 10)

	base10 := big.NewInt(10)
	bigOne := big.NewInt(1)

	// Truncate Factor
	if factorPrecision > 0 {
		scale := big.NewInt(0).Exp(base10, big.NewInt(int64(factorPrecision)), nil)
		factor.Quo(factor, scale)
	}

	lastMag := big.NewInt(0).Sub(factor, big.NewInt(1))

	for true {

		val := big.NewInt(0).Exp(base10, factor, nil)

		if val.Cmp(target) == 1 {
			break
		}

		lastMag.Set(factor)
		factor.Add(factor, bigOne)
	}


	return lastMag, nil
}



func (bIntMath BigIntMath) GetMagnitude(target *big.Int) (*big.Int, error) {

	compareResult := target.Cmp(big.NewInt(0))

	if compareResult == -1 {
		ePrefix := "BigIntMath.GetMagnitude() "
		return big.NewInt(0),
			fmt.Errorf(ePrefix+"Error: Input parameter 'target' is a negative number. "+
				"target='%v' ", target.Text(10))
	}

	if compareResult == 0 {
		return big.NewInt(0), nil
	}

	// Conversion constant 3.3219281
	conversionConst := BigIntNum{}.NewInt64Exponent(33219281, -7)
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
*/