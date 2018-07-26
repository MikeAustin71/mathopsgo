package mathops

import (
	"fmt"
	"math/big"
	"errors"
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// default numeric separators (decimal separator, thousands separator and
// currency symbol).
//
// The returned BigIntNum 'result' will contain default numeric separators (decimal
// separator, thousands separator and currency symbol).
//
func (bSubtract BigIntMathSubtract) SubtractBigInts(
									minuend *big.Int,
										minPrecision uint,
											subtrahend *big.Int,
												subPrecision uint) BigIntNum {


	bPair := BigIntPair{}.NewBase(
												minuend,
													minPrecision,
														subtrahend,
															subPrecision)

	return bSubtract.subtractPairNoNumSeps(bPair)
}

// SubtractBigIntNums - Receives two 'BigIntNum' instances and proceeds to subtract
// b2 from b2.
//
// 								'minuend' - 'subtrahend' = difference or result
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntNum. The returned BigIntNum 'result' will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from the
// input parameter, 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNums(minuend, subtrahend BigIntNum) BigIntNum {

	numSeps := minuend.GetNumericSeparatorsDto()

	bPair := BigIntPair{}.NewBigIntNum(minuend, subtrahend)

	bigIntNumResult := bSubtract.subtractPairNoNumSeps(bPair)

	bigIntNumResult.SetNumericSeparatorsDto(numSeps)

	return bigIntNumResult
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
// Type BigIntNum. The returned BigIntNum subtraction 'result' will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied
// from the input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNumArray(
																				minuend BigIntNum,
																					subtrahends []BigIntNum) BigIntNum {

  numSeps := minuend.GetNumericSeparatorsDto()

	finalResult := minuend.CopyOut()

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult
	}

	for i:=0; i < lenSubtrahends; i++ {

		bPair := BigIntPair{}.NewBigIntNum(finalResult, subtrahends[i])

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)

	}

	finalResult.SetNumericSeparatorsDto(numSeps)

	return finalResult
}

// SubtractBigIntNumOutputToArray - The first input parameter to this method
// is a BigIntNum Type labeled, 'minuend'.  The second input parameter is an
// array of BigIntNum types labeled 'subtrahends'. The 'minuend' is subtracted
// from each element of the 'subtrahends' array with the result output to another
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
// Each of the BigIntNum instances included in the array of BigIntNum subtraction
// results returned by this method, will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNumOutputToArray(
																				minuend BigIntNum,
																					subtrahends []BigIntNum) []BigIntNum {

  numSeps := minuend.GetNumericSeparatorsDto()

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []BigIntNum{}
	}

	resultsArray := make([]BigIntNum, lenSubtrahends)
	
	for i:=0; i < lenSubtrahends; i++ {

		bPair := BigIntPair{}.NewBigIntNum(minuend, subtrahends[i])

		resultsArray[i] = bSubtract.subtractPairNoNumSeps(bPair)
		resultsArray[i].SetNumericSeparatorsDto(numSeps)
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
// Type BigIntNum.
//
// The subtraction result returned by this method as a Type BigIntNum will
// contain numeric separators (decimal separator, thousands separator and currency
// symbol) copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractBigIntNumSeries(
																				minuend BigIntNum,
																					subtrahends ... BigIntNum) BigIntNum {

  numSeps := minuend.GetNumericSeparatorsDto()

	finalResult := minuend.CopyOut()

	for _, subtrahend := range subtrahends {

		bPair := BigIntPair{}.NewBigIntNum(finalResult, subtrahend)

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	finalResult.SetNumericSeparatorsDto(numSeps)

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
// Type BigIntNum. This resulting BigIntNum instance will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from input
// parameter 'decMinuend'.
//
func (bSubtract BigIntMathSubtract) SubtractDecimal(
																			decMinuend Decimal,
																				decSubtrahend Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractDecimal() "

	numSeps := decMinuend.GetNumericSeparatorsDto()

	bPair := BigIntPair{}.NewBigIntNum(decMinuend.bigINum, decSubtrahend.bigINum)

	finalResult := bSubtract.subtractPairNoNumSeps(bPair)

	err := finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps) " +
				"Error='%v' \n", err.Error())
	}

	return finalResult, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractDecimalArray(
				minuend Decimal,
					subtrahends []Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractDecimalArray() "

	numSeps := minuend.GetNumericSeparatorsDto()

	finalResult := minuend.GetBigIntNum()

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	for i:=0; i < lenSubtrahends; i++ {

		bPair := BigIntPair{}.NewBigIntNum(
														finalResult, subtrahends[i].bigINum.CopyOut())

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err := finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v'", err.Error())

	}

	return finalResult, nil
}

// SubtractDecimalOutputToArray - The first input parameter to this method
// is a Decimal Type labeled, 'minuend'.  The second input parameter is an 
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
// The array element of the []Decimal 'result' returned by this subtraction
// operation will contain numeric separators (decimal separator, thousands
// separator and currency symbol) copied from input parameter 'minuend'.
//
// The 'result' array ([]Decimal) returned by this subtraction operation will
// contain array elements with numeric separators (decimal separator, thousands
// separator and currency symbol) which have been copied from input parameter
// 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractDecimalOutputToArray(
												minuend Decimal,
														subtrahends []Decimal) ([]Decimal, error) {

  ePrefix := "BigIntMathSubtract.SubtractDecimalOutputToArray() "

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []Decimal{},
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	bINumMinuend := minuend.GetBigIntNum()

	resultsArray := make([]Decimal, lenSubtrahends)

	var err error

	for i:=0; i < lenSubtrahends; i++ {

		bPair := BigIntPair{}.NewBigIntNum(bINumMinuend, subtrahends[i].GetBigIntNum())

		result := bSubtract.subtractPairNoNumSeps(bPair)

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by result.SetNumericSeparatorsDto(numSeps). " +
				"index='%v' Error='%v'", i, err.Error())
		}

		resultsArray[i], err = result.GetDecimal()

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix +
					"Error returned by result.Result.GetDecimal() " +
					"i='%v' result.Result='%v' Error='%v'. ",
					i, result.GetNumStr(), err.Error())
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractDecimalSeries(
										minuend Decimal,
											subtrahends ... Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractDecimalSeries() "

	numSeps := minuend.GetNumericSeparatorsDto()

	finalResult := minuend.GetBigIntNum()

	if len(subtrahends) == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends series is Empty!")
	}

	for _, subtrahend := range subtrahends {

		bPair := BigIntPair{}.NewBigIntNum(finalResult, subtrahend.GetBigIntNum())

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err := finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
		fmt.Errorf(ePrefix +
			"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' \n", err.Error())
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'iaMinuend'.
//
func (bSubtract BigIntMathSubtract) SubtractIntAry(
									iaMinuend IntAry,
										iaSubtrahend IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractIntAry() "

	// Method NewIntAry will test validity of iaMinuend
	bINumMinuend, err := BigIntNum{}.NewIntAry(iaMinuend)

	if err != nil {
		return BigIntNum{}.New(),
		 fmt.Errorf(ePrefix + "Error returned from BigIntNum{}.NewIntAry(iaMinuend). " +
		 	"Error='%v' ", err.Error())
	}

	// Method NewIntAry will test validity of iaSubtrahend
	bINumSubtrahend, err := BigIntNum{}.NewIntAry(iaSubtrahend)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix + "Error returned from BigIntNum{}.NewIntAry(iaSubtrahend). " +
				"Error='%v' ", err.Error())
	}

	numSeps := iaMinuend.GetNumericSeparatorsDto()

	bPair := BigIntPair{}.NewBigIntNum(bINumMinuend, bINumSubtrahend)

	finalResult := bSubtract.subtractPairNoNumSeps(bPair)

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' \n", err.Error())
	}

	return finalResult, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractIntAryArray(
				minuend IntAry,
					subtrahends []IntAry) (BigIntNum, error) {


	ePrefix := "BigIntMathSubtract.SubtractInAryArray() "


	// NewIntAry method will test validity of 'minuend'
	finalResult, err := BigIntNum{}.NewIntAry(minuend)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewIntAry(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	for i:=0; i < lenSubtrahends; i++ {

		// Method NewIntAry will test validity of subtrahends[i]
		bigINum, err := BigIntNum{}.NewIntAry(subtrahends[i])

		if err != nil {
			return BigIntNum{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewIntAry(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult, bigINum)

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' \n", err.Error())
	}

	return finalResult, nil
}

// SubtractIntAryOutputToArray - The first input parameter to this method
// is a IntAry Type labeled, 'minuend'.  The second input parameter is an
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
// Each element in the []IntAry result returned by this subtraction operation will
// contain numeric separators (decimal separator, thousands separator and currency
// symbol) which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractIntAryOutputToArray(
														minuend IntAry,
															subtrahends []IntAry) ([]IntAry, error) {

	ePrefix := "BigIntMathSubtract.SubtractIntAryOutputToArray() "

	// NewIntAry will test validity of 'minuend'
	bINumMinuend, err := BigIntNum{}.NewIntAry(minuend)

	if err != nil {
		return []IntAry{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(minuend). " +
			"Error='%v' ", err.Error())
	}

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []IntAry{},
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	resultsArray := make([]IntAry, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		// Method NewIntAry will test validity of subtrahends[i]
		bINumSubtrahend, err := BigIntNum{}.NewIntAry(subtrahends[i])

		if err != nil {
			return []IntAry{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(subtrahends[%v]). " +
				"Error='%v' \n", i, err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(bINumMinuend, bINumSubtrahend)

		result := bSubtract.subtractPairNoNumSeps(bPair)

		resultsArray[i], err = result.GetIntAry()

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by result.Result.GetIntAry() " +
					"i='%v' result.Result='%v' Error='%v'. ",
					i, result.GetNumStr(), err.Error())
		}

		err = resultsArray[i].SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by resultsArray[i].SetNumericSeparatorsDto(numSeps). " +
					"Error='%v' \n", err.Error())
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
// Type BigIntNum.
//
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractIntArySeries(
																			minuend IntAry,
																				subtrahends ... IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractIntArySeries() "

	// Method New IntAry will test the validity of minuend
	finalResult, err := BigIntNum{}.NewIntAry(minuend)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewIntAry(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	if len(subtrahends) == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends series is Empty!")
	}

	for i, subtrahend := range subtrahends {

	  // Method NewIntAry will test the validity of IntAry subtrahend
		bigINumSubtrahend, err := BigIntNum{}.NewIntAry(subtrahend)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewIntAry(subtrahend). " +
					"index='%v' subtrahend='%v' Error='%v'",
					i, subtrahend.GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult, bigINumSubtrahend)

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' \n", err.Error())
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractINumMgr(
						minuend,
							subtrahend INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractINumMgr() "



	bPair, err := BigIntPair{}.NewINumMgr(minuend, subtrahend)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by NBigIntPair{}.NewINumMgr(minuend, subtrahend). " +
				"minuend.GetNumStr()='%v', subtrahend.GetNumStr()='%v' Error='%v' ",
				minuend.GetNumStr(), subtrahend.GetNumStr(), err.Error())
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	finalResult := bSubtract.subtractPairNoNumSeps(bPair)

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps) " +
				"Error='%v' \n", err.Error())
	}

	return finalResult, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractINumMgrArray(
									minuend INumMgr,
										subtrahends []INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractINumMgrArray() "

	finalResult, err := BigIntNum{}.NewINumMgr(minuend)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewINumMgr(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	for i:= 0; i < lenSubtrahends; i++ {

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, subtrahends[i])

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &subtrahends[i]). " +
					" i='%v' subtrahends[i].GetNumStr()='%v' Error='%v' ", i, subtrahends[i].GetNumStr(), err.Error())
		}

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
		fmt.Errorf(ePrefix + "Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' \n", err.Error())

	}

	return finalResult, nil
}

// SubtractINumMgrOutputToArray - The first input parameter to this method
// is an object which implements the INumMgr interface ('minuend').  The
// second input parameter is an array of INumMgr interface types labeled
// 'subtrahends'. The 'minuend' is subtracted from each element of the
// 'subtrahends' array with the result output to another 'results' array
// of INumMgr interface types which is then returned to the calling function.
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
// Note: The underlying type for the returned results array is 'BigIntNum', a
// type which implements the INumMgr Interface.
//
//
// Each element of the []INumMgr 'result' array returned by this subtraction
// operation will contain numeric separators (decimal separator, thousands
// separator and currency symbol) which were copied from input parameter
// 'minuend'.
//

func (bSubtract BigIntMathSubtract) SubtractINumMgrOutputToArray(
										minuend INumMgr,
											subtrahends []INumMgr) ([]INumMgr, error) {

	ePrefix := "BigIntMathSubtract.SubtractINumMgrOutputToArray() "

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []INumMgr{},
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	resultsArray := make([]INumMgr, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		bPair, err := BigIntPair{}.NewINumMgr(minuend, subtrahends[i])

		if err != nil {
			return []INumMgr{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(minuend, subtrahends[i]) " +
					"minuend='%v' subtrahends[%v]='%v' Error='%v'. ",
					minuend.GetNumStr(), i, subtrahends[i].GetNumStr(), err.Error())
		}

		result := bSubtract.subtractPairNoNumSeps(bPair)

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []INumMgr{},
				fmt.Errorf(ePrefix +
					"Error returned by result.SetNumericSeparatorsDto(numSeps). " +
					"Index='%v' Error='%v' \n", i, err.Error())
		}

		resultsArray[i] = &result

	}

	return resultsArray, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractINumMgrSeries(
					minuend INumMgr,
						subtrahends ... INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractINumMgrSeries() "

	finalResult, err := BigIntNum{}.NewINumMgr(minuend)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewINumMgr(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	if len(subtrahends) == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends series is Empty!")
	}

	for i, subtrahend := range subtrahends {

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, subtrahend)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &subtrahend). " +
					" i='%v' subtrahend.GetNumStr()='%v' Error='%v' ",
					i, subtrahend.GetNumStr(), err.Error())
		}

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v'\n", err.Error())
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// set to default values.
//
func (bSubtract BigIntMathSubtract) SubtractNumStr(
													n1 string,
														n2 string) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1, n2)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1, n2). " +
				"Error='%v' ", err.Error())
	}

	finalResult := bSubtract.subtractPairNoNumSeps(bPair)

	finalResult.SetNumericSeparatorsToDefaultIfEmpty()

	return finalResult, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// set to default values.
//
func (bSubtract BigIntMathSubtract) SubtractNumStrArray(
				minuend string,
					subtrahends []string) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrArray() "

	finalResult, err := BigIntNum{}.NewNumStr(minuend)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(minuend). " +
				"minuend='%v' Error='%v'", minuend, err.Error())
	}

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	for i:=0; i < lenSubtrahends; i++ {

		bigINum, err := BigIntNum{}.NewNumStr(subtrahends[i])

		if err != nil {
			return BigIntNum{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStr(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i], err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult, bigINum)

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	finalResult.SetNumericSeparatorsToDefaultIfEmpty()

	return finalResult, nil
}


// SubtractNumStrOutputToArray - The first input parameter to this method
// is a string Type labeled, 'minuend'.  The second input parameter is an
// array of strings labeled 'subtrahends'. The 'minuend' is subtracted
// from each element of the 'subtrahends' array with the result output to
// another 'results' array of strings which is then returned to the calling
// function.
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
func (bSubtract BigIntMathSubtract) SubtractNumStrOutputToArray(
														minuend string,
															subtrahends []string) ([]string, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrOutputToArray() "

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		return []string{},
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	resultsArray := make([]string, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		bPair, err := BigIntPair{}.NewNumStr(minuend, subtrahends[i])

		if err != nil {
			return []string{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewNumStr(minuend, subtrahends[i]) " +
					"minuend='%v' subtrahends[%v]='%v' Error='%v'. ",
					minuend, i, subtrahends[i], err.Error())
		}

		result := bSubtract.subtractPairNoNumSeps(bPair)

		resultsArray[i] = result.GetNumStr()

		if err != nil {
			return []string{},
				fmt.Errorf(ePrefix +
					"Error returned by result.Result.GetNumStr() " +
					"i='%v' result.Result='%v' Error='%v'. ",
					i, result.GetNumStr(), err.Error())
		}
	}

	return resultsArray, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which are set to their default values.
//
func (bSubtract BigIntMathSubtract) SubtractNumStrSeries(
															minuend string,
																subtrahends ... string) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrSeries() "

	finalResult, err := BigIntNum{}.NewNumStr(minuend)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(minuend). " +
				"minuend='%v' Error='%v'", minuend, err.Error())
	}

	if len(subtrahends) == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends series is Empty!")
	}

	for i, subtrahend := range subtrahends {

		bigINumSubtrahend, err := BigIntNum{}.NewNumStr(subtrahend)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStr(subtrahend). " +
					"index='%v' subtrahend='%v' Error='%v'",
					i, subtrahend, err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult, bigINumSubtrahend)

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	finalResult.SetNumericSeparatorsToDefaultIfEmpty()

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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'nDtoMinuend'.
//

func (bSubtract BigIntMathSubtract) SubtractNumStrDto(
									nDtoMinuend NumStrDto,
										nDtoSubtrahend NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrDto() "

	err := nDtoMinuend.IsValid(ePrefix + "'nDtoMinuend' INVALID! ")

	if err != nil {
		return BigIntNum{}.New(), err
	}

	err = nDtoSubtrahend.IsValid(ePrefix + "'nDtoSubtrahend' INVALID! ")

	if err != nil {
		return BigIntNum{}.New(), err
	}

	bPair, err := BigIntPair{}.NewNumStrDto(nDtoMinuend, nDtoSubtrahend)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewNumStrDto(nDtoMinuend, nDtoSubtrahend). " +
				"Error='%v' ", err.Error())
	}

	numSeps := nDtoMinuend.GetNumericSeparatorsDto()

	finalResult := bSubtract.subtractPairNoNumSeps(bPair)

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' ", err.Error())
	}

	return finalResult, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractNumStrDtoArray(
				minuend NumStrDto,
					subtrahends []NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrDtoArray() "

	err := minuend.IsValid(ePrefix + "'minuend' INVALID! ")

	if err != nil {
		return BigIntNum{}.New(), err
	}

	finalResult, err := BigIntNum{}.NewNumStrDto(minuend)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStrDto(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	lenSubtrahends := len(subtrahends)

	if lenSubtrahends == 0 {
		finalResult.SetNumericSeparatorsDto(numSeps)
		return finalResult,
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	for i:=0; i < lenSubtrahends; i++ {

		err := subtrahends[i].IsValid(ePrefix +
								fmt.Sprintf("subtrahends[%v] INVALID!", i))

		bigINum, err := BigIntNum{}.NewNumStrDto(subtrahends[i])

		if err != nil {
			return BigIntNum{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStrDto(subtrahends[i]). " +
					"i='%v' subtrahend='%v' Error='%v'",
					i, subtrahends[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult, bigINum)

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' \n", err.Error())
	}

	return finalResult, nil
}

// SubtractNumStrDtoOutputToArray - The first input parameter to this method
// is a NumStrDto Type labeled, 'minuend'.  The second input parameter is an 
// array of NumStrDto types labeled 'subtrahends'. The 'minuend' is subtracted
// from each element of the 'subtrahends' array with the result output to
// another 'results' array of NumStrDto types which is then returned to the
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
// Each array element in the []NumStrDto 'result' returned by this subtraction
// operation will contain numeric separators (decimal separator, thousands
// separator and currency symbol) which were copied from input parameter
// 'minuend'.
//

func (bSubtract BigIntMathSubtract) SubtractNumStrDtoOutputToArray(
												minuend NumStrDto,
														subtrahends []NumStrDto) ([]NumStrDto, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrDtoOutputToArray() "

	lenSubtrahends := len(subtrahends)

	err := minuend.IsValid(ePrefix + "'minuend' INVALID! ")

	if err != nil {
		return []NumStrDto{}, err
	}

	if lenSubtrahends == 0 {
		return []NumStrDto{},
			errors.New(ePrefix + "Error: subtrahends array is Empty!")
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	resultsArray := make([]NumStrDto, lenSubtrahends)

	for i:=0; i < lenSubtrahends; i++ {

		err := subtrahends[i].IsValid(ePrefix +
			fmt.Sprintf("subtrahends[%v] INVALID! ", i))

		bPair, err := BigIntPair{}.NewNumStrDto(minuend, subtrahends[i])

		if err != nil {
			return []NumStrDto{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewNumStrDto(minuend, subtrahends[i]) " +
					"minuend='%v' subtrahends[%v]='%v' Error='%v'. ",
					minuend.GetNumStr(), i, subtrahends[i].GetNumStr(), err.Error())
		}

		result := bSubtract.subtractPairNoNumSeps(bPair)

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []NumStrDto{},
			fmt.Errorf(ePrefix +
				"Error returned by result.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' \n", err.Error())
		}

		resultsArray[i], err = result.GetNumStrDto()

		if err != nil {
			return []NumStrDto{},
				fmt.Errorf(ePrefix +
					"Error returned by result.Result.GetNumStrDto() " +
					"i='%v' result.Result='%v' Error='%v'. ",
					i, result.GetNumStr(), err.Error())
		}
	}

	return resultsArray, nil
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
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'minuend'.
//
func (bSubtract BigIntMathSubtract) SubtractNumStrDtoSeries(
																			minuend NumStrDto,
																				subtrahends ... NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathSubtract.SubtractNumStrDtoSeries() "

	err := minuend.IsValid(ePrefix + "'minuend' INVALID! ")

	if err != nil {
		return BigIntNum{}.New(), err
	}

	finalResult, err := BigIntNum{}.NewNumStrDto(minuend)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStrDto(minuend). " +
				"minuend='%v' Error='%v'", minuend.GetNumStr(), err.Error())
	}

	numSeps := minuend.GetNumericSeparatorsDto()

	if len(subtrahends) == 0 {

		finalResult.SetNumericSeparatorsDto(numSeps)

		return finalResult,
			errors.New(ePrefix + "Error: subtrahends series is Empty!")
	}

	for i, subtrahend := range subtrahends {

		err := subtrahend.IsValid(ePrefix +
			fmt.Sprintf("subtrahend index='%v' INVALID! ", i))

		bigINumSubtrahend, err := BigIntNum{}.NewNumStrDto(subtrahend)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStrDto(subtrahend). " +
					"index='%v' subtrahend='%v' Error='%v'",
					i, subtrahend.GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult, bigINumSubtrahend)

		finalResult = bSubtract.subtractPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	return finalResult, nil
}


// subtractPairNoNumSeps - Performs the subtraction operation. This method receives a type
// 'BigIntPair' and proceeds to subtract b2.bigInt from b1.bigInt.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter bPair.Big1, the minuend.
//
func (bSubtract BigIntMathSubtract) SubtractPair(bPair BigIntPair) BigIntNum {

	numSeps := bPair.Big1.GetNumericSeparatorsDto()

	finalResult := bSubtract.subtractPairNoNumSeps(bPair)

	finalResult.SetNumericSeparatorsDto(numSeps)

	return finalResult
}

// subtractPairNoNumSeps - Performs the subtraction operation. This method receives a type
// 'BigIntPair' and proceeds to subtract b2.bigInt from b1.bigInt.
//
// After the subtraction operation, the 'difference' or 'result' is returned as a
// Type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// default numeric separators (decimal separator, thousands separator and currency
// symbol).
//

func (bSubtract BigIntMathSubtract) subtractPairNoNumSeps(bPair BigIntPair) BigIntNum {

	bPair.MakePrecisionsEqual()
	
	b3 := big.NewInt(0).Sub(bPair.GetBig1BigInt(), bPair.GetBig2BigInt())

	bResult := BigIntNum{}.NewBigInt(b3, bPair.Big2.GetPrecisionUint())

	bResult.SetNumericSeparatorsToDefaultIfEmpty()

	return bResult
}

