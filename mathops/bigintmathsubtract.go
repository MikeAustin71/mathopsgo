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
// Input Parameters
// ================
//
//	minuend *big.Int					- The number from which the subtrahend will be subtracted
//	minPrecision uint,				- The 'minuend' precision or numeric digits after
//																	the decimal point.
//	subtrahend *big.Int,			- The number to be subtracted from the 'minuend'.
//	subPrecision uint  				- The 'subtrahend' precision or numeric digits after
//																	the decimal point.
//
// Return Values
// =============
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractBigInts(
									minuend *big.Int,
										minPrecision uint,
											subtrahend *big.Int,
												subPrecision uint) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBase(
		minuend,
		minPrecision,
		subtrahend,
		subPrecision)

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
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
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
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
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

// SubtractBigIntNumOutputToArray - The first input parameter to this method
// is a BigIntNum Type labeled, 'minuend'.  The second element is an array
// of BigIntNum types labeled 'subtrahends'. The 'minuend' is subtracted from
// each element of the 'subtrahends' array with the result output to another
// 'results' array of BigIntNum types which is then returned to the calling
// function.
//
// Example
// =======
// 										    subtrahends										 Output
// Minuend   				    	  Array											   Array
//
//		10			-					subtrahends[0] = 2			=				  outputarray[0] =  8
//		10			-					subtrahends[1] = 3			=				  outputarray[1] =  7
//		10			-					subtrahends[2] = 4			=				  outputarray[2] =  6
//		10			-					subtrahends[3] = 5			=				  outputarray[3] =  5
//		10			-					subtrahends[4] = 6			=				  outputarray[4] =  4
//		10			-					subtrahends[5] = 9			=				  outputarray[5] =  1
//
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNumOutputToArray(
																				minuend BigIntNum,
																					subtrahends []BigIntNum) []BigIntNum {

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []BigIntNum{}
	}

	resultsArray := make([]BigIntNum, lenSubtrahends)
	
	for i:=0; i < lenSubtrahends; i++ {

		bPair := BigIntPair{}.NewBigIntNum(minuend, subtrahends[i])

		result := bSubtract.SubtractPair(bPair)

		resultsArray[i] = result.Result.CopyOut()

	}

	return resultsArray
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
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
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
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
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
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
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

// SubtractDecimalOutputToArray - The first input parameter to this method
// is a Decimal Type labeled, 'minuend'.  The second element is an 
// array of Decimal types labeled 'subtrahends'. The 'minuend' is subtracted
// from each element of the 'subtrahends' array with the result output to
// another 'results' array of Decimal types which is then returned to the
// calling function.
//
// Example
// =======
// 										    subtrahends										 Output
// Minuend   				    	   Array											 Array
//
//		10			-					subtrahends[0] = 2			=				  outputarray[0] =  8
//		10			-					subtrahends[1] = 3			=				  outputarray[1] =  7
//		10			-					subtrahends[2] = 4			=				  outputarray[2] =  6
//		10			-					subtrahends[3] = 5			=				  outputarray[3] =  5
//		10			-					subtrahends[4] = 6			=				  outputarray[4] =  4
//		10			-					subtrahends[5] = 9			=				  outputarray[5] =  1
//
//
func (bSubtract BigIntMathSubtract) SubtractDecimalOutputToArray(
												minuend Decimal,
														subtrahends []Decimal) ([]Decimal, error) {

  ePrefix := "BigIntMathSubtract.SubtractDecimalOutputToArray() "

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []Decimal{}, nil
	}

	resultsArray := make([]Decimal, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		bPair, err := BigIntPair{}.NewDecimal(minuend, subtrahends[i])

		if err != nil {
			return []Decimal{},
			 fmt.Errorf(ePrefix +
			 	"Error returned by BigIntPair{}.NewDecimal(minuend, subtrahends[i]) " +
			 	"minuend='%v' subtrahends[%v]='%v' Error='%v'. ",
			 		minuend.GetNumStr(), i, subtrahends[i].GetNumStr(), err.Error())
		}

		result := bSubtract.SubtractPair(bPair)

		resultsArray[i], err = result.Result.GetDecimal()

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix +
					"Error returned by result.Result.GetDecimal() " +
					"i='%v' result.Result='%v' Error='%v'. ",
					i, result.Result.GetNumStr(), err.Error())
		}
	}

	return resultsArray, nil
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
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
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
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
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
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
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
					"Error returned by BigIntNum{}.NewIntAry(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINum)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult, nil
}

// SubtractIntAryOutputToArray - The first input parameter to this method
// is a IntAry Type labeled, 'minuend'.  The second element is an 
// array of IntAry types labeled 'subtrahends'. The 'minuend' is subtracted
// from each element of the 'subtrahends' array with the result output to
// another 'results' array of IntAry types which is then returned to the
// calling function.
//
// Example
// =======
// 										    subtrahends										 Output
//  Minuend   				    	Array													Array
//
//		10			-					subtrahends[0] = 2			=				  outputarray[0] =  8
//		10			-					subtrahends[1] = 3			=				  outputarray[1] =  7
//		10			-					subtrahends[2] = 4			=				  outputarray[2] =  6
//		10			-					subtrahends[3] = 5			=				  outputarray[3] =  5
//		10			-					subtrahends[4] = 6			=				  outputarray[4] =  4
//		10			-					subtrahends[5] = 9			=				  outputarray[5] =  1
//
//
func (bSubtract BigIntMathSubtract) SubtractIntAryOutputToArray(
														minuend IntAry,
															subtrahends []IntAry) ([]IntAry, error) {

	ePrefix := "BigIntMathSubtract.SubtractIntAryOutputToArray() "

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []IntAry{}, nil
	}

	resultsArray := make([]IntAry, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		bPair, err := BigIntPair{}.NewIntAry(minuend, subtrahends[i])

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewIntAry(minuend, subtrahends[i]) " +
					"minuend='%v' subtrahends[%v]='%v' Error='%v'. ",
					minuend.GetNumStr(), i, subtrahends[i].GetNumStr(), err.Error())
		}

		result := bSubtract.SubtractPair(bPair)

		resultsArray[i], err = result.Result.GetIntAry()

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by result.Result.GetIntAry() " +
					"i='%v' result.Result='%v' Error='%v'. ",
					i, result.Result.GetNumStr(), err.Error())
		}
	}

	return resultsArray, nil
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
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// The 'difference' or 'result' is returned as a Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
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


// SubtractINumMgr - Receives two objects which implement the INumMgr Interface
// and subtracts their numeric values.
//
// The 'subtrahend' numeric value is subtracted from the 'minuend' numeric value.
//
// In the subtraction operation:
//
//								'minuend' - 'subtrahend' = difference or result
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractINumMgr(
						minuend,
							subtrahend INumMgr) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractINumMgr() "

	bPair, err := BigIntPair{}.NewINumMgr(minuend, subtrahend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by NBigIntPair{}.NewINumMgr(minuend, subtrahend). " +
				"minuend.GetNumStr()='%v', subtrahend.GetNumStr()='%v' Error='%v' ",
				minuend.GetNumStr(), subtrahend.GetNumStr(), err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}

// SubtractINumMgrArray - Receives two input parameters. The first parameter
// is an INumMgr instance which is classified as the 'minuend'. The second
// parameter is an array of INumMgr instances which resents the 'subtrahends'.
// A subtraction operation is performed on the 'minuend' and the 'subtrahends'
// array. The numeric value of the 'subtrahends' is subtracted from the 'minuend'.
//
// In the subtraction operation:
//
//								'minuend' - 'subtrahend' = difference or result
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry. This allows the user to mix different types in
// a single array and add their numeric values.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractINumMgrArray(
									minuend INumMgr,
										subtrahends []INumMgr) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractINumMgrArray() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewINumMgr(minuend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewINumMgr(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return finalResult, nil
	}

	for i:= 0; i < lenSubtrahends; i++ {

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, subtrahends[i])

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &subtrahends[i]). " +
					" i='%v' subtrahends[i].GetNumStr()='%v' Error='%v' ", i, subtrahends[i].GetNumStr(), err.Error())
		}

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}


// SubtractINumMgrSeries - Receives two input parameters. The first parameter
// is an INumMgr instance which is classified as the 'minuend'. The second
// parameter is a series of INumMgr instances which resents the 'subtrahends'.
// A subtraction operation is performed on the 'minuend' and the 'subtrahends'.
// The numeric value of the 'subtrahends' is subtracted from the 'minuend'.
//
// In the subtraction operation:
//
//					'minuend' - 'subtrahend' = difference or result
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry. This allows the user to mix different types in
// a single array and add their numeric values.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractINumMgrSeries(
					minuend INumMgr,
						subtrahends ... INumMgr) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractINumMgrSeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()

	finalResult.Result, err = BigIntNum{}.NewINumMgr(minuend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewINumMgr(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	for i, subtrahend := range subtrahends {

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, subtrahend)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &subtrahend). " +
					" i='%v' subtrahend.GetNumStr()='%v' Error='%v' ",
					i, subtrahend.GetNumStr(), err.Error())
		}

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// SubtractNumStr - Receives two number strings and proceeds to subtract
// n2 from n1.
//
// The strings passed to this method are 'number' strings in that they consist
// of a string of numeric digits which may include a period ('.') or decimal point
// used to separate fractional numeric digits.
//
// In the subtraction operation:
// 								n1 - n2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								n1 = 'minuend'
//								n2 = 'subtrahend'
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
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

// SubtractNumStrArray - Receives one string parameter which is classified as
// the 'minuend'. The second input parameter is an array of strings labeled,
// 'subtrahends'.
//
// The strings passed to this method are 'number' strings in that they consist
// of a string of numeric digits which may include a period ('.') or decimal point
// used to separate fractional numeric digits.
//
// The array of 'subtrahends' is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is an array of string Types.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractNumStrArray(
				minuend string,
					subtrahends []string) (BigIntBasicMathResult, error) {

	var err error
	ePrefix := "BigIntMathSubtract.SubtractNumStrArray() "
	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewNumStr(minuend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(minuend). " +
				"minuend='%v' Error='%v'", minuend, err.Error())
	}

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return finalResult, nil
	}

	for i:=0; i < lenSubtrahends; i++ {

		bigINum, err := BigIntNum{}.NewNumStr(subtrahends[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStr(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i], err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINum)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult, nil
}

// SubtractNumStrSeries - Receives one 'string' Type which is classified as
// the 'minuend'. The second input parameter, 'subtrahends' is a series of
// 'strings' .
//
// The strings passed to this method are 'number' strings in that they consist
// of a string of numeric digits which may include a period ('.') or decimal point
// used to separate fractional numeric digits.
//
// The 'subtrahends' series is subtracted from the 'minuend'.
//
// In the subtraction operation:
// 								b1 - b2 = difference or result
//								'minuend' - 'subtrahend' = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is a series of strings.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractNumStrSeries(
				minuend string,
					subtrahends ... string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrSeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewNumStr(minuend)

	if err != nil {
		return BigIntBasicMathResult{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(minuend). " +
				"minuend='%v' Error='%v'", minuend, err.Error())
	}

	for i, subtrahend := range subtrahends {

		bigINumSubtrahend, err := BigIntNum{}.NewNumStr(subtrahend)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStr(subtrahend). " +
					"index='%v' subtrahend='%v' Error='%v'",
					i, subtrahend, err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINumSubtrahend)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}


// SubtractNumStrDto - Receives two 'NumStrDto' instances and proceeds to subtract
// 'nDtoSubtrahend' from 'nDtoMinuend'.
//
// 			nDtoMinuend - nDtoSubtrahend = difference or result
//
// In the subtraction operation:
//							  'minuend' - 'subtrahend' = difference or result
//								nDtoMinuend 		= 'minuend'
//								nDtoSubtrahend 	= 'subtrahend'
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractNumStrDto(
					nDtoMinuend NumStrDto,
						nDtoSubtrahend NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(nDtoMinuend, nDtoSubtrahend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(nDtoMinuend, nDtoSubtrahend). " +
				"Error='%v' ", err.Error())
	}

	return bSubtract.SubtractPair(bPair), nil
}

// SubtractNumStrDtoArray - Receives one NumStrDto parameter which is classified as
// the 'minuend'. The second input parameter is an array of NumStrDto Types
// labeled, 'subtrahends'.
//
// The array of 'subtrahends' is subtracted from the 'minuend'.
//
// In the subtraction operation:
//								'minuend' - 'subtrahend' = difference or result
// 								b1 - b2 = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is an array of NumStrDto Types.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractNumStrDtoArray(
				minuend NumStrDto,
					subtrahends []NumStrDto) (BigIntBasicMathResult, error) {

	var err error
	ePrefix := "BigIntMathSubtract.SubtractNumStrDtoArray() "
	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewNumStrDto(minuend)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStrDto(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return finalResult, nil
	}

	for i:=0; i < lenSubtrahends; i++ {

		bigINum, err := BigIntNum{}.NewNumStrDto(subtrahends[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStrDto(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINum)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult, nil
}

// SubtractNumStrDtoSeries - Receives one NumStrDto Type which is classified as
// the 'minuend'. The second input parameter, 'subtrahends' is a series of
// Type NumStrDto .
//
// The 'subtrahends' series is subtracted from the 'minuend'.
//
// In the subtraction operation:
//								'minuend' - 'subtrahend' = difference or result
// 								b1 - b2 = difference or result
//								b1 = 'minuend'
//								b2 = 'subtrahend'
//
// In this method, the 'subtrahend' is a series of NumStrDto Types.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= (last) subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractNumStrDtoSeries(
				minuend NumStrDto,
					subtrahends ... NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrDtoSeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result, err = BigIntNum{}.NewNumStrDto(minuend)

	if err != nil {
		return BigIntBasicMathResult{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStrDto(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	for i, subtrahend := range subtrahends {

		bigINumSubtrahend, err := BigIntNum{}.NewNumStrDto(subtrahend)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStrDto(subtrahend). " +
					"index='%v' subtrahend='%v' Error='%v'",
					i, subtrahend.GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, bigINumSubtrahend)

		result := bSubtract.SubtractPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// SubtractPair - Performs the subtraction operation. This method receives a type
// 'BigIntPair' and proceeds to subtract b2.bigInt from b1.bigInt.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntBasicMathResult.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= minuend
//											Input.Big2		= subtrahend
//
// 								Result BigIntNum
// 											Result.bigInt = difference or result
//					}
//
func (bSubtract BigIntMathSubtract) SubtractPair(bPair BigIntPair) BigIntBasicMathResult {

	bPair.MakePrecisionsEqual()
	
	b3 := big.NewInt(0).Sub(bPair.GetBig1BigInt(), bPair.GetBig2BigInt())

	bResult := BigIntBasicMathResult{}
	bResult.Input = bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(b3, bPair.Big2.GetPrecisionUint())

	return bResult
}

