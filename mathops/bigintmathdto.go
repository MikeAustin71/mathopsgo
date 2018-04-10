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
type BigIntMathDto struct {
	Input  BigIntPair
	Result BigIntNum
}

// New - Creates a BigIntMathDto intance with data
// variables initialized to zero.
//
func (bMath BigIntMathDto) New() BigIntMathDto {

	b2Math := BigIntMathDto{}

	b2Math.Input = BigIntPair{}.New()

	baseZero := big.NewInt(0)

	b2Math.Result = BigIntNum{}.NewBigInt(baseZero, 0)

	return b2Math
}

// NewNewBigIntPair - Creates a new BigIntMathDto based on input parameter
// type, 'BigIntPair'
//
func (bMath BigIntMathDto) NewBigIntPairResult(bPair BigIntPair) BigIntMathDto {

	b2Math := BigIntMathDto{}

	b2Math.Input = bPair.CopyOut()

	return b2Math
}

// AddBigIntNums - Adds two BigIntNums and returns the result in a new
// BigIntBasicMathResult instance
//
func (bMath BigIntMathDto) AddBigIntNums(b1, b2 BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bMath.AddPair(bPair)
}

func (bMath BigIntMathDto) AddBigInts(
				b1 *big.Int,
				precision1 uint,
				b2 *big.Int,
				precision2 uint ) BigIntBasicMathResult {

	b1Pair := BigIntPair{}.NewBase(b1, precision1, b2, precision2)

	return bMath.AddPair(b1Pair)
}

// AddDecimal - Receives two Decimal instances and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bMath BigIntMathDto) AddDecimal(dec1, dec2 Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.AddNumStrDto() "

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	return bMath.AddPair(bPair), nil
}

// AddIntAry - Receives two IntAry instances and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bMath BigIntMathDto) AddIntAry(ia1, ia2 IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.AddNumStrDto() "

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
				"Error='%v' ", err.Error())
	}

	return bMath.AddPair(bPair), nil
}

// AddNumStr - Receives two number strings and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bMath BigIntMathDto) AddNumStr(n1NumStr, n2NumStr string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.AddNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1NumStr, n2NumStr)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1NumStr, n2NumStr). " +
				"Error='%v' ", err.Error())
	}

	return bMath.AddPair(bPair), nil
}

// AddNumStrDto - Receives two NumStrDto instances and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bMath BigIntMathDto) AddNumStrDto(n1Dto, n2Dto NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.AddNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(n1Dto, n2Dto)

	if err != nil {
		return BigIntBasicMathResult{},
		fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(n1Dto, n2Dto). " +
			"Error='%v' ", err.Error())
	}

	return bMath.AddPair(bPair), nil

}

// AddPair - Receives a BigIntPair and proceeds to add b1.BigIntNum to
// b2.BigIntNum.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bMath BigIntMathDto) AddPair(bPair BigIntPair) BigIntBasicMathResult {

	bPair.MakePrecisionsEqual()

	b3 := big.NewInt(0).Add(bPair.Big1.BigInt, bPair.Big2.BigInt)

	bResult := BigIntBasicMathResult{}
	bResult.Input = bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(b3, bPair.Big2.Precision)

	return bResult
}

// MultiplyBigIntNums - Receives two BigIntNum types and proceeds to multiply the
// two numeric values.
//
// 				b1 X b2 = result
//
// The result is returned as a BigIntMathResult type.
//
func (bMath BigIntMathDto) MultiplyBigIntNums(b1, b2 BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bMath.MultiplyPair(bPair)

}

// MultiplyBigInts - Receives two *big.Int numbers and their associated precision
// specifications. The method the proceeds to mulitpy b1 times b2.
//
//						b1 x b2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bMath BigIntMathDto) MultiplyBigInts(
	b1 *big.Int,
	b1Precision uint,
	b2 *big.Int,
	b2Precision uint) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBase(b1, b1Precision, b2, b2Precision)

	return bMath.MultiplyPair(bPair)

}

// MultiplyDecimal - Receives two Decimal instances and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
func (bMath BigIntMathDto) MultiplyDecimal(
														dec1,
														dec2 Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.MultiplyDecimal() "

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	return bMath.MultiplyPair(bPair), nil
}

// MultiplyIntAry - Receives two IntAry instances and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
func (bMath BigIntMathDto) MultiplyIntAry(
														ia1,
														ia2 IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.MultiplyIntAry() "

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
				"Error='%v' ", err.Error())
	}

	return bMath.MultiplyPair(bPair), nil
}

// MultiplyNumStr - Receives two number strings and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
func (bMath BigIntMathDto) MultiplyNumStr(
							n1NumStr,
								n2NumStr string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.MultiplyNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1NumStr, n2NumStr)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1NumStr, n2NumStr). " +
				"Error='%v' ", err.Error())
	}

	return bMath.MultiplyPair(bPair), nil
}

// MultiplyNumStrDto - Receives two NumStrDto instances and multiplies their
// numeric values. The result is returned as a 'BigIntBasicMathResult'
// type.
//
func (bMath BigIntMathDto) MultiplyNumStrDto(
							n1Dto,
							n2Dto NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.MultiplyNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(n1Dto, n2Dto)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(n1Dto, n2Dto). " +
				"Error='%v' ", err.Error())
	}

	return bMath.MultiplyPair(bPair), nil
}

// MultiplyPair - Receives a BigIntPair instance and proceeds to multiply
// b1.BigIntNum by b2.BigIntNum.
//
// b1.BigIntNum x b2.BigIntNum = Result
//
// The result is returned as a BigIntBasicMathResult type.
//
func (bMath BigIntMathDto) MultiplyPair(bPair BigIntPair) BigIntBasicMathResult {



	b3 := big.NewInt(0).Mul(bPair.Big1.BigInt, bPair.Big2.BigInt)

	bResult := BigIntBasicMathResult{}
	bResult.Input = bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(b3, bPair.Big1.Precision + bPair.Big2.Precision)

	return bResult

}

// SubtractBigInts - Performs the subtraction operation. This method receives
// two *big.Int types and subtracts 'minuend' from 'subtrahend'.
//
// The result is returned as a type BigIntBasicMathResult
//
//  'subtrahend' - 'minuend' = result
//
func (bMath BigIntMathDto) SubtractBigInts(
														subtrahend *big.Int,
															subPrecision uint,
																minuend *big.Int,
																	minPrecision uint) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBase(
												subtrahend,
													subPrecision,
														minuend,
															minPrecision)

	return bMath.SubtractPair(bPair)
}


// SubtractBigIntNums - Receives two 'BigIntNum' instances and proceeds to subtract
// b2 from b1.
//
// b1 - b2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bMath BigIntMathDto) SubtractBigIntNums(b1, b2 BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bMath.SubtractPair(bPair)
}


// SubtractDecimal - Receives two Decimal instances and proceeds to subtract
// dec2 from dec1.
//
// 											dec1 - dec2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
func (bMath BigIntMathDto) SubtractDecimal(
														dec1 Decimal,
															dec2 Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.SubtractDecimal() "

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	return bMath.SubtractPair(bPair), nil
}


// SubtractIntAry - Receives two IntAry instances and proceeds to subtract
// ia2 from ia1.
//
// ia1 - ia2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bMath BigIntMathDto) SubtractIntAry(
														ia1 IntAry,
															ia2 IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.SubtractIntAry() "

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
				"Error='%v' ", err.Error())
	}

	return bMath.SubtractPair(bPair), nil
}


// SubtractNumStr - Receives two number strings and proceeds to subtract
// n2 from n1.
//
// n1 - n2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bMath BigIntMathDto) SubtractNumStr(
																	n1 string,
																		n2 string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.SubtractNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1, n2)

	if err != nil {
		return BigIntBasicMathResult{},
		fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1, n2). " +
			"Error='%v' ", err.Error())
	}

	return bMath.SubtractPair(bPair), nil
}

// SubtractNumStrDto - Receives two NumStrDto instances and proceeds to subtract
// n2Dto from n1Dto.
//
// n1Dto - n2Dto = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bMath BigIntMathDto) SubtractNumStrDto(
															n1Dto NumStrDto,
																nDto NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathDto.SubtractNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(n1Dto, nDto)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(n1Dto, nDto). " +
				"Error='%v' ", err.Error())
	}

	return bMath.SubtractPair(bPair), nil
}


// SubtractPair - Performs the subtraction operation. This method receives a type
// 'BigIntPair' and proceeds to subtract b2.BigInt from b1.BigInt.
//
// The result is returned as type BigIntBasicMathResult
//
func (bMath BigIntMathDto) SubtractPair(bPair BigIntPair) BigIntBasicMathResult {

	bPair.MakePrecisionsEqual()

	b3 := big.NewInt(0).Sub(bPair.Big1.BigInt, bPair.Big2.BigInt)

	bResult := BigIntBasicMathResult{}
	bResult.Input = bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(b3, bPair.Big2.Precision)

	return bResult
}

