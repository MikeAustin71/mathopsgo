package mathops

import (
	"fmt"
	"math/big"
)

// BigIntMathSubtract - Contains methods used to perform subtraction 
// operations on *big.Int numeric types.
//
// 			minuend − subtrahend = difference
//
type BigIntMathSubtract struct {
	Input		BigIntPair	// BigIntPair.Big1 = minuend  BigIntPair.Big2 = subtrahend
	Result 	BigIntNum		// The result of the subtraction otherwise known as the 'difference'
}


// SubtractBigInts - Performs the subtraction operation. This method receives
// two *big.Int types and subtracts 'minuend' from 'subtrahend'.
//
// The result is returned as a type BigIntBasicMathResult
//
//  'subtrahend' - 'minuend' = result
//
func (bSubtract BigIntMathSubtract) SubtractBigInts(
	subtrahend *big.Int,
	subPrecision uint,
	minuend *big.Int,
	minPrecision uint) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBase(
		subtrahend,
		subPrecision,
		minuend,
		minPrecision)

	return bSubtract.SubtractPair(bPair)
}



// SubtractBigIntNums - Receives one BigIntNum which is classified as the 'minuend'.
// The second input parameter is an array of BigIntNum Types. The array of 'subtrahends'
// is subtracted from the 'minuends'
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'subtrahend' - 'minuend' = result
//								b1 = 'subtrahend'
//								b2 = 'minuend'
//
// In this case the 'subtrahend' is an array of BigIntNum Types.
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNumArray(
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

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult
}


// SubtractBigIntNums - Receives two 'BigIntNum' instances and proceeds to subtract
// b2 from b2.
//
// 								b1 - b2 = difference
//								'subtrahend' - 'minuend' = result
//								b1 = 'subtrahend'
//								b2 = 'minuend'
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNums(b1, b2 BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bSubtract.SubtractPair(bPair)
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
func (bSubtract BigIntMathSubtract) SubtractDecimal(
	dec1 Decimal,
	dec2 Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractDecimal() "

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}


// SubtractIntAry - Receives two IntAry instances and proceeds to subtract
// ia2 from ia1.
//
// ia1 - ia2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractIntAry(
	ia1 IntAry,
	ia2 IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractIntAry() "

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
				"Error='%v' ", err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}


// SubtractNumStr - Receives two number strings and proceeds to subtract
// n2 from n1.
//
// n1 - n2 = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractNumStr(
	n1 string,
	n2 string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1, n2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1, n2). " +
				"Error='%v' ", err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}

// SubtractNumStrDto - Receives two NumStrDto instances and proceeds to subtract
// n2Dto from n1Dto.
//
// n1Dto - n2Dto = result
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractNumStrDto(
	n1Dto NumStrDto,
	nDto NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(n1Dto, nDto)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(n1Dto, nDto). " +
				"Error='%v' ", err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}


// SubtractPair - Performs the subtraction operation. This method receives a type
// 'BigIntPair' and proceeds to subtract b2.BigInt from b1.BigInt.
//
// The result is returned as type BigIntBasicMathResult
//
func (bSubtract BigIntMathSubtract) SubtractPair(bPair BigIntPair) BigIntBasicMathResult {

	bPair.MakePrecisionsEqual()

	b3 := big.NewInt(0).Sub(bPair.Big1.BigInt, bPair.Big2.BigInt)

	bResult := BigIntBasicMathResult{}
	bResult.Input = bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(b3, bPair.Big2.Precision)

	return bResult
}
