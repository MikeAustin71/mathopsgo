package mathops

import (
	"fmt"
	"math/big"
)

type BigIntMathPower struct {
	Base BigIntNum
	Exponent  BigIntNum
	Result BigIntNum
}


func (bIPwr BigIntMathPower) Pwr(base, exponent BigIntNum) (BigIntNum, error) {
	return bIPwr.raiseToPositiveIntegerPower(base, exponent)
}


func (bIPwr BigIntMathPower) raiseToPositiveIntegerPower(base, exponent BigIntNum) (BigIntNum, error) {
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