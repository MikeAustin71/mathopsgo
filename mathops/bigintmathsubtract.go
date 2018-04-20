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


// SubtractBigInts - Performs the subtraction operation.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
func (bSubtract BigIntMathSubtract) SubtractBigInts(
	minuend *big.Int,
	minPrecision uint,
	subtrahend *big.Int,
	subPrecision uint) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBase(
		subtrahend,
		subPrecision,
		minuend,
		minPrecision)

	return bSubtract.SubtractPair(bPair)
}

// SubtractBigIntNums - Receives two 'BigIntNum' instances and proceeds to subtract
// b2 from b2.
//
// 								b1 - b2 = difference
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNums(b1, b2 BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bSubtract.SubtractPair(bPair)
}

// SubtractBigIntNumArray - Receives one BigIntNum which is classified as the 'minuend'.
// The second input parameter is an array of BigIntNum Types labeled, 'subtrahends'.
// The array of 'subtrahends' is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is an array of BigIntNum Types.
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

// SubtractBigIntNumSeries - Receives one BigIntNum which is classified as
// the 'minuend'. The second input parameter, 'subtrahends' is a series of
// Type BigIntNum .
//
// The 'subtrahends' series is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is a series of BigIntNum Types.
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNumSeries(
																				minuend BigIntNum,
																					subtrahends ... BigIntNum) BigIntBasicMathResult {

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result = minuend.CopyOut()

	for _, subtrahend := range subtrahends {

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, subtrahend)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult
}

// SubtractDecimal - Performs the subtraction operation on two Decimal Types.
//
// 				decMinuend - decSubtrahend = difference
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// The result of this subtraction operation is returned as a
// type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractDecimal(
																			decMinuend Decimal,
																				decSubtrahend Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractDecimal() "

	bPair, err := BigIntPair{}.NewDecimal(decMinuend, decSubtrahend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(decMinuend, decSubtrahend). " +
				"Error='%v' ", err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}

// SubtractDecimalArray - Receives one Decimal parameter which is classified as
// the 'minuend'. The second input parameter is an array of Decimal Types
// labeled, 'subtrahends'.
//
// The array of 'subtrahends' is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is an array of Decimal Types.
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractDecimalArray(
				minuend Decimal,
					subtrahends []Decimal) (BigIntBasicMathResult, error) {

	var err error
	ePrefix := "BigIntMathSubtract.SubtractDecimalArray() "
	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewDecimal(minuend)

	if err != nil {
		return BigIntBasicMathResult{},
		fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewDecimal(minuend). " +
			"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}



	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return finalResult, nil
	}

	for i:=0; i < lenSubtrahends; i++ {

		bigINum, err := BigIntNum{}.NewDecimal(subtrahends[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINum)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult, nil
}

// SubtractDecimalSeries - Receives one Decimal Type which is classified as
// the 'minuend'. The second input parameter, 'subtrahends' is a series of
// Type Decimal .
//
// The 'subtrahends' series is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is a series of Decimal Types.
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractDecimalSeries(
										minuend Decimal,
											subtrahends ... Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractDecimalSeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewDecimal(minuend)

	if err != nil {
		return BigIntBasicMathResult{}.New(),
		fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewDecimal(minuend). " +
			"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	for i, subtrahend := range subtrahends {

		bigINumSubtrahend, err := BigIntNum{}.NewDecimal(subtrahend)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(subtrahend). " +
					"index='%v' subtrahend='%v' Error='%v'",
					i, subtrahend.GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINumSubtrahend)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// SubtractIntAry - Performs the subtraction operation on two
//
// IntAry Types.
// 				iaMinuend - iaSubtrahend = difference
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// The result of this subtraction operation is returned as a
// type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractIntAry(
									iaMinuend IntAry,
										iaSubtrahend IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractIntAry() "

	bPair, err := BigIntPair{}.NewIntAry(iaMinuend, iaSubtrahend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewIntAry(iaMinuend, iaSubtrahend). " +
				"Error='%v' ", err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}

// SubtractIntAryArray - Receives one IntAry parameter which is classified as
// the 'minuend'. The second input parameter is an array of IntAry Types
// labeled, 'subtrahends'.
//
// The array of 'subtrahends' is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is an array of IntAry Types.
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractIntAryArray(
				minuend IntAry,
					subtrahends []IntAry) (BigIntBasicMathResult, error) {

	var err error
	ePrefix := "BigIntMathSubtract.SubtractInAryArray() "
	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewIntAry(minuend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewIntAry(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}



	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return finalResult, nil
	}

	for i:=0; i < lenSubtrahends; i++ {

		bigINum, err := BigIntNum{}.NewIntAry(subtrahends[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINum)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult, nil
}

// SubtractIntArySeries - Receives one IntAry Type which is classified as
// the 'minuend'. The second input parameter, 'subtrahends' is a series of
// Type IntAry .
//
// The 'subtrahends' series is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is a series of IntAry Types.
//
// The result is returned as a type 'BigIntBasicMathResult'.
//
func (bSubtract BigIntMathSubtract) SubtractIntArySeries(
					minuend IntAry,
						subtrahends ... IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractIntArySeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewIntAry(minuend)

	if err != nil {
		return BigIntBasicMathResult{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewIntAry(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	for i, subtrahend := range subtrahends {

		bigINumSubtrahend, err := BigIntNum{}.NewIntAry(subtrahend)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewIntAry(subtrahend). " +
					"index='%v' subtrahend='%v' Error='%v'",
					i, subtrahend.GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINumSubtrahend)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
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
