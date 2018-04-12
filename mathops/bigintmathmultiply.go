package mathops

import (
	"math/big"
	"fmt"
)

// BigIntMathDto - This type is used to perform math operations
// using the *big.Int Type.
//
// If you are unfamiliar with the *big.Int type, reference:
// 						https://golang.org/pkg/math/big/
//
type BigIntMathMultiply struct {
	Input  BigIntPair
	Result BigIntNum
}

// New - Creates a BigIntMathMultiply intance with data
// variables initialized to zero.
//
func (bMultiply BigIntMathMultiply) New() BigIntMathMultiply {

	b2Math := BigIntMathMultiply{}

	b2Math.Input = BigIntPair{}.New()

	baseZero := big.NewInt(0)

	b2Math.Result = BigIntNum{}.NewBigInt(baseZero, 0)

	return b2Math
}

// NewNewBigIntPair - Creates a new BigIntMathMultiply based on input parameter
// type, 'BigIntPair'
//
func (bMultiply BigIntMathMultiply) NewBigIntPairResult(bPair BigIntPair) BigIntMathMultiply {

	b2Math := BigIntMathMultiply{}

	b2Math.Input = bPair.CopyOut()

	return b2Math
}


// MultiplyBigIntNums - Receives two BigIntNum types and proceeds to multiply the
// two numeric values.
//
// 				b1 X b2 = result
//
// The result is returned as a BigIntMathResult type.
//
func (bMultiply BigIntMathMultiply) MultiplyBigIntNums(b1, b2 BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bMultiply.MultiplyPair(bPair)

}

// MultiplyBigInts - Receives two *big.Int numbers and their associated precision
// specifications. The method the proceeds to mulitpy b1 times b2.
//
//						b1 x b2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bMultiply BigIntMathMultiply) MultiplyBigInts(
	b1 *big.Int,
	b1Precision uint,
	b2 *big.Int,
	b2Precision uint) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBase(b1, b1Precision, b2, b2Precision)

	return bMultiply.MultiplyPair(bPair)

}

// MultiplyDecimal - Receives two Decimal instances and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
func (bMultiply BigIntMathMultiply) MultiplyDecimal(
														dec1,
														dec2 Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyDecimal() "

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	return bMultiply.MultiplyPair(bPair), nil
}

// MultiplyIntAry - Receives two IntAry instances and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
func (bMultiply BigIntMathMultiply) MultiplyIntAry(
														ia1,
														ia2 IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyIntAry() "

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
				"Error='%v' ", err.Error())
	}

	return bMultiply.MultiplyPair(bPair), nil
}

// MultiplyNumStr - Receives two number strings and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
func (bMultiply BigIntMathMultiply) MultiplyNumStr(
							n1NumStr,
								n2NumStr string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1NumStr, n2NumStr)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1NumStr, n2NumStr). " +
				"Error='%v' ", err.Error())
	}

	return bMultiply.MultiplyPair(bPair), nil
}

// MultiplyNumStrDto - Receives two NumStrDto instances and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
func (bMultiply BigIntMathMultiply) MultiplyNumStrDto(
							n1Dto,
							n2Dto NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(n1Dto, n2Dto)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(n1Dto, n2Dto). " +
				"Error='%v' ", err.Error())
	}

	return bMultiply.MultiplyPair(bPair), nil
}

// MultiplyPair - Receives a BigIntPair instance and proceeds to multiply
// b1.BigIntNum by b2.BigIntNum.
//
// b1.BigIntNum x b2.BigIntNum = Result
//
// The result is returned as a BigIntBasicMathResult type.
//
func (bMultiply BigIntMathMultiply) MultiplyPair(bPair BigIntPair) BigIntBasicMathResult {



	b3 := big.NewInt(0).Mul(bPair.Big1.BigInt, bPair.Big2.BigInt)

	bResult := BigIntBasicMathResult{}
	bResult.Input = bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(b3, bPair.Big1.Precision + bPair.Big2.Precision)

	return bResult

}

