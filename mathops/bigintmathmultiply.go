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


// MultiplyBigInts - Receives two *big.Int numbers and their associated precision
// specifications. This method then proceeds to perform a multiplication operation
// by multiplying the 'multiplier' by the 'multiplicand'
//
// In the multiplication operation, the number to be multiplied is called the
// "multiplicand", while the number of times the multiplicand is to be multiplied
// comes from the "multiplier". Usually the multiplier is placed first and the
// multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// Input Parameters
// ================
//
//	multiplier *big.Int					- The number to be multiplied by 'multiplicand'
//	multiplierPrecision uint,		- The 'multiplier' Precision or numeric digits after
//																	the decimal point.
//	multiplicand *big.Int,			- The number to be multiplied by the 'multiplier'.
//	multiplicandPrecision uint  - The 'multiplicand' Precision or numeric digits after
//																	the decimal point.
//
// Return Values
// =============
//
// This method performs the multiplication operation and returns the result or
// 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result BigIntNum
// 											Result.BigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyBigInts(
									multiplier *big.Int,
										multiplierPrecision uint,
											multiplicand *big.Int,
												multiplicandPrecision uint) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBase(
												multiplier,
													multiplierPrecision,
														multiplicand,
															multiplicandPrecision)

	return bMultiply.MultiplyPair(bPair)

}


// MultiplyBigIntNums - Receives two BigIntNum types and performs a 'multiplication'
// operation on the BigIntNum's.
//
// In the multiplication operation, the number to be multiplied is called the
// "multiplicand", while the number of times the multiplicand is to be multiplied
// comes from the "multiplier". Usually the multiplier is placed first and the
// multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result BigIntNum
// 											Result.BigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyBigIntNums(
								multiplier BigIntNum,
									multiplicand BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(multiplier, multiplicand)

	return bMultiply.MultiplyPair(bPair)

}

// MultiplyBigIntNumArray - Receives one BigIntNum which is classified as the 'minuend'.
// The second input parameter is an array of BigIntNum Types labeled, 'subtrahends'.
// The array of 'subtrahends' is subtracted from the 'minuend'.
//
// In the multiplication operation, the number to be multiplied is called the
// "multiplicand", while the number of times the multiplicand is to be multiplied
// comes from the "multiplier". Usually the multiplier is placed first and the
// multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result BigIntNum
// 											Result.BigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyBigIntNumArray(
	minuend BigIntNum,
	subtrahends []BigIntNum) BigIntBasicMathResult {

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result = minuend.CopyOut()

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return finalResult
	}

	for i:=0; i < lenSubtrahends; i++ {

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, subtrahends[i])

		result := bMultiply.MultiplyPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult
}


// MultiplyDecimal - Receives two Decimal instances and multiplies their
// numeric values. The result or 'product' is returned as a 'BigIntBasicMathResult'
// type.
//
// In the multiplication operation, the number to be multiplied is called the
// "multiplicand", while the number of times the multiplicand is to be multiplied
// comes from the "multiplier". Usually the multiplier is placed first and the
// multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result BigIntNum
// 											Result.BigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyDecimal(
														multiplier,
														multiplicand Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyDecimal() "

	bPair, err := BigIntPair{}.NewDecimal(multiplier, multiplicand)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(multiplier, multiplicand). " +
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

