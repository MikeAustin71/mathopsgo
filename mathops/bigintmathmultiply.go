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
//	multiplierPrecision uint,		- The 'multiplier' precision or numeric digits after
//																	the decimal point.
//	multiplicand *big.Int,			- The number to be multiplied by the 'multiplier'.
//	multiplicandPrecision uint  - The 'multiplicand' precision or numeric digits after
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
// 											Result.bigInt = product
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


// MultiplyBigIntNums - Receives two BigIntNum types as input parameters and then
// performs a 'multiplication' operation on these BigIntNum's.
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
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyBigIntNums(
								multiplier BigIntNum,
									multiplicand BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(multiplier, multiplicand)

	return bMultiply.MultiplyPair(bPair)

}

// MultiplyBigIntNumArray - Receives one BigIntNum which is classified as the 'multiplier'.
// The second input parameter is an array of BigIntNum Types labeled, 'multiplicands'. The
// first element of the 'multiplicands' array is multiplied by the 'multiplier' to produce
// a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next element
// in the multiplicands array. This process is continued through the last element in the array
// and the combined, final 'product' is returned as a Type 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result BigIntNum
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyBigIntNumArray(
									multiplier BigIntNum,
											multiplicands []BigIntNum) BigIntBasicMathResult {

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result = multiplier.CopyOut()
	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return finalResult
	}

	for i:=0; i < lenMultiplicands; i++ {

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, multiplicands[i])

		result := bMultiply.MultiplyPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult
}

// MultiplyBigIntNumOutputToArray - Receives one input parameter of Type BigIntNum which
// is classified as the 'multiplier'. The second input parameter is an array of BigIntNum
// Types labeled, 'multiplicands'.
//
// Each element of the 'multiplicands' array is multiplied by the 'multiplier'. The result or
// 'product' is then stored in a results array which is returned to the calling function.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier', 3 is the 'multiplicand'
// and 15 is the 'product' or result.
//
// Example
// =======
// 										Multiplicands												Output
// Multiplier				    	Array														Array
//
//		3			x				multiplicands[0] = 2			=				  outputarray[0] =  6
//		3			x				multiplicands[1] = 3			=				  outputarray[1] =  9
//		3			x				multiplicands[2] = 4			=				  outputarray[2] = 12
//		3			x				multiplicands[3] = 5			=				  outputarray[3] = 15
//		3			x				multiplicands[4] = 6			=				  outputarray[4] = 18
//		3			x				multiplicands[5] = 7			=				  outputarray[5] = 21
//
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' in an Array of 'BigIntNums' ([] BigIntNums).
//
func (bMultiply BigIntMathMultiply) MultiplyBigIntNumOutputToArray(
																		multiplier BigIntNum,
																			multiplicands []BigIntNum) [] BigIntNum {

	bINumInterimResult := multiplier.CopyOut()

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return []BigIntNum{}
	}

	resultArray := make([]BigIntNum, lenMultiplicands)

	for i:=0; i < lenMultiplicands; i++ {

		bPair := BigIntPair{}.NewBigIntNum(bINumInterimResult, multiplicands[i])

		result := bMultiply.MultiplyPair(bPair)

		resultArray[i] = result.Result.CopyOut()

	}

	return resultArray
}

// MultiplyBigIntNumSeries - Receives one input parameter of Type BigIntNum which is classified
//  as the 'multiplier'. The second input parameter is a series of BigIntNum Types labeled,
// 'multiplicands'. The first element of the 'multiplicands' series is multiplied by the 'multiplier'
// to produce a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next
// element in the multiplicands series. This process is continued through the last element in the
// series. Afterwards, the combined final 'product' is returned as a Type 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result BigIntNum
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyBigIntNumSeries(
																multiplier BigIntNum,
																	multiplicands ... BigIntNum) BigIntBasicMathResult {

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result = multiplier.CopyOut()
	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return finalResult
	}

	for _, multiplicand := range multiplicands {

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, multiplicand)

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
// 											Result.bigInt = product
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

// MultiplyDecimalArray - Receives one Decimal which is classified as the 'multiplier'.
// The second input parameter is an array of Decimal Types labeled, 'multiplicands'. The
// first element of the 'multiplicands' array is multiplied by the 'multiplier' to produce
// a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next element
// in the multiplicands array. This process is continued through the last element in the array
// when the combined, final 'product' is returned as a Type 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result Decimal
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyDecimalArray(
															multiplier Decimal,
																multiplicands []Decimal) (BigIntBasicMathResult, error) {

  ePrefix := "BigIntMathMultiply.MultiplyDecimalArray() "
  var err error
	finalResult := BigIntBasicMathResult{}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return finalResult, nil
	}

	finalResult.Result, err = BigIntNum{}.NewDecimal(multiplier)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewDecimal(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewDecimal(multiplicands[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(multiplicands[i]) " +
					"multiplicands[%v]='%v' Error='%v'. ",
					 i, multiplicands[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, multiplicandBINum )

		finalResult = bMultiply.MultiplyPair(bPair)

	}

	return finalResult, nil
}

// MultiplyDecimalOutputToArray - Receives one input parameter of Type Decimal which
// is classified as the 'multiplier'. The second input parameter is an array of Decimal
// Types labeled, 'multiplicands'.
//
// Each element of the 'multiplicands' array is multiplied by the 'multiplier'. The result or
// 'product' is then stored in a results array which is returned to the calling function.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier', 3 is the 'multiplicand'
// and 15 is the 'product' or result.
//
// Example
// =======
// 										Multiplicands												Output
// Multiplier				    	Array														Array
//
//		3			x				multiplicands[0] = 2			=				  outputarray[0] =  6
//		3			x				multiplicands[1] = 3			=				  outputarray[1] =  9
//		3			x				multiplicands[2] = 4			=				  outputarray[2] = 12
//		3			x				multiplicands[3] = 5			=				  outputarray[3] = 15
//		3			x				multiplicands[4] = 6			=				  outputarray[4] = 18
//		3			x				multiplicands[5] = 7			=				  outputarray[5] = 21
//
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' in an Array of 'Decimals' ([] Decimals).
//
func (bMultiply BigIntMathMultiply) MultiplyDecimalOutputToArray(
																		multiplier Decimal,
																			multiplicands []Decimal) ([]Decimal, error) {

	ePrefix := "BigIntMathMultiply.MultiplyDecimalOutputToArray() "

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return []Decimal{}, nil
	}

	multiplierBINum, err := BigIntNum{}.NewDecimal(multiplier)

	if err != nil {
		return []Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewDecimal(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	resultArray := make([]Decimal, lenMultiplicands)

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewDecimal(multiplicands[i])

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(multiplicands[i]) " +
					" multiplierBINum='%v' multiplicands[%v]='%v' Error='%v'. ",
					multiplierBINum.GetNumStr(), i, multiplicands[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(multiplierBINum, multiplicandBINum )

		finalResult := bMultiply.MultiplyPair(bPair)

		resultArray[i], err = finalResult.Result.GetDecimal()

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix +
					"Error returned by finalResult.Result.GetDecimal() " +
					"i='%v' Error='%v'. ", i,  err.Error())
		}
	}

	return resultArray, nil
}

// MultiplyDecimalSeries - Receives one input parameter of Type Decimal which is classified
//  as the 'multiplier'. The second input parameter is a comma delimited series of Decimal
// Types labeled, 'multiplicands'. The first element of the 'multiplicands' series is multiplied
// by the 'multiplier' to produce a 'product'. That 'product' replaces the 'multiplier' and is
// multiplied by the next element in the multiplicands series. This process is continued through
// the last element in the series. Afterwards, the combined final 'product' is returned as a Type
// 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result Decimal
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyDecimalSeries(
													multiplier Decimal,
														multiplicands ... Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyDecimalSeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()

	finalResult.Result, err = BigIntNum{}.NewDecimal(multiplier)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewDecimal(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	for _, multiplicand := range multiplicands {

		multiplicandBINum, err := BigIntNum{}.NewDecimal(multiplicand)

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(multiplicand) " +
					" multiplicand='%v' Error='%v'. ",
					multiplicand.GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result.CopyOut(), multiplicandBINum)

		finalResult = bMultiply.MultiplyPair(bPair)

	}

	return finalResult, nil
}

// MultiplyIntAry - Receives two IntAry instances and multiplies their
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
// 											Result.bigInt = product
//					}
//
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

// MultiplyIntAryArray - Receives one IntAry which is classified as the 'multiplier'.
// The second input parameter is an array of IntAry Types labeled, 'multiplicands'. The
// first element of the 'multiplicands' array is multiplied by the 'multiplier' to produce
// a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next element
// in the multiplicands array. This process is continued through the last element in the array
// when the combined, final 'product' is returned as a Type 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result IntAry
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyIntAryArray(
													multiplier IntAry,
														multiplicands []IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyIntAryArray() "

	var err error

	finalResult := BigIntBasicMathResult{}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return finalResult, nil
	}

	finalResult.Result, err = BigIntNum{}.NewIntAry(multiplier)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewIntAry(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewIntAry(multiplicands[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewIntAry(multiplicands[i]) " +
					"multiplicands[%v]='%v' Error='%v'. ",
					i, multiplicands[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, multiplicandBINum )

		finalResult = bMultiply.MultiplyPair(bPair)

	}

	return finalResult, nil
}

// MultiplyIntAryOutputToArray - Receives one input parameter of Type IntAry which
// is classified as the 'multiplier'. The second input parameter is an array of IntAry
// Types labeled, 'multiplicands'.
//
// Each element of the 'multiplicands' array is multiplied by the 'multiplier'. The result or
// 'product' is then stored in a results array which is returned to the calling function.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier', 3 is the 'multiplicand'
// and 15 is the 'product' or result.
//
// Example
// =======
// 										Multiplicands												Output
// Multiplier				    	Array														Array
//
//		3			x				multiplicands[0] = 2			=				  outputarray[0] =  6
//		3			x				multiplicands[1] = 3			=				  outputarray[1] =  9
//		3			x				multiplicands[2] = 4			=				  outputarray[2] = 12
//		3			x				multiplicands[3] = 5			=				  outputarray[3] = 15
//		3			x				multiplicands[4] = 6			=				  outputarray[4] = 18
//		3			x				multiplicands[5] = 7			=				  outputarray[5] = 21
//
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' in an Array of 'IntArys' ([] IntArys).
//
func (bMultiply BigIntMathMultiply) MultiplyIntAryOutputToArray(
																		multiplier IntAry,
																			multiplicands []IntAry) ([]IntAry, error) {

	ePrefix := "BigIntMathMultiply.MultiplyIntAryOutputToArray() "

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return []IntAry{}, nil
	}

	multiplierBINum, err := BigIntNum{}.NewIntAry(multiplier)

	if err != nil {
		return []IntAry{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewIntAry(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	resultArray := make([]IntAry, lenMultiplicands)

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewIntAry(multiplicands[i])

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewIntAry(multiplicands[i]) " +
					" multiplierBINum='%v' multiplicands[%v]='%v' Error='%v'. ",
					multiplierBINum.GetNumStr(), i, multiplicands[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(multiplierBINum, multiplicandBINum )

		finalResult := bMultiply.MultiplyPair(bPair)

		resultArray[i], err = finalResult.Result.GetIntAry()

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by finalResult.Result.GetIntAry() " +
					"i='%v' Error='%v'. ", i,  err.Error())
		}
	}

	return resultArray, nil
}

// MultiplyIntArySeries - Receives one input parameter of Type IntAry which is classified
//  as the 'multiplier'. The second input parameter is a comma separated series of IntAry Types
// labeled, 'multiplicands'. The first element of the 'multiplicands' series is multiplied by
// the 'multiplier' to produce a 'product'. That 'product' replaces the 'multiplier' and is
// multiplied by the next element in the multiplicands series. This process is continued through
// the last element in the series. Afterwards, the combined final 'product' is returned as a
// Type 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result IntAry
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyIntArySeries(
									multiplier IntAry,
													multiplicands ... IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyIntArySeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()

	finalResult.Result, err = BigIntNum{}.NewIntAry(multiplier)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewIntAry(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	for _, multiplicand := range multiplicands {

		multiplicandBINum, err := BigIntNum{}.NewIntAry(multiplicand)

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewIntAry(multiplicand) " +
					" multiplicand='%v' Error='%v'. ",
					multiplicand.GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result.CopyOut(), multiplicandBINum)

		finalResult = bMultiply.MultiplyPair(bPair)

	}

	return finalResult, nil
}


// MultiplyNumStr - Receives two number strings and multiplies their numeric
// values.
//
// 							n1NumStr x n2NumStr = product
//
// The result or 'product' is returned as a 'BigIntBasicMathResult'
// type.
//
// The strings passed to this method are 'number' strings in that they consist
// of a string of numeric digits which may include a period ('.') or decimal point
// used to separate fractional numeric digits.
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
// 											Result.bigInt = product
//					}
//
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

// MultiplyNumStrArray - Receives one number string which is classified as the 'multiplier'.
// The second input parameter is an array of number strings labeled, 'multiplicands'. The
// first element of the 'multiplicands' array is multiplied by the 'multiplier' to produce
// a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next element
// in the multiplicands array. This process is continued through the last element in the array
// when the combined, final 'product' is returned as a Type 'BigIntBasicMathResult'.
//
// The strings passed to this method are 'number' strings in that they consist of a string
// of numeric digits which may include a period ('.') or decimal point used to separate
// fractional numeric digits.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result Decimal
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyNumStrArray(
													multiplier string,
															multiplicands []string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyNumStrArray() "

	var err error

	finalResult := BigIntBasicMathResult{}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return finalResult, nil
	}

	finalResult.Result, err = BigIntNum{}.NewNumStr(multiplier)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewDecimal(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier, err.Error())
	}

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewNumStr(multiplicands[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(multiplicands[i]) " +
					"multiplicands[%v]='%v' Error='%v'. ",
					i, multiplicands[i], err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, multiplicandBINum )

		finalResult = bMultiply.MultiplyPair(bPair)
	}

	return finalResult, nil
}

// MultiplyNumStrOutputToArray - Receives one input parameter of Type string which
// is classified as the 'multiplier'. The second input parameter is an array of strings
// labeled, 'multiplicands'.
//
// Each element of the 'multiplicands' array is multiplied by the 'multiplier'. The result or
// 'product' is then stored in a results array which is returned to the calling function.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier', 3 is the 'multiplicand'
// and 15 is the 'product' or result.
//
// Example
// =======
// 										Multiplicands												Output
// Multiplier				    	Array														Array
//
//		3			x				multiplicands[0] = 2			=				  outputarray[0] =  6
//		3			x				multiplicands[1] = 3			=				  outputarray[1] =  9
//		3			x				multiplicands[2] = 4			=				  outputarray[2] = 12
//		3			x				multiplicands[3] = 5			=				  outputarray[3] = 15
//		3			x				multiplicands[4] = 6			=				  outputarray[4] = 18
//		3			x				multiplicands[5] = 7			=				  outputarray[5] = 21
//
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' in an Array of 'NumStrs' ([] NumStrs).
//
func (bMultiply BigIntMathMultiply) MultiplyNumStrOutputToArray(
													multiplier string,
														multiplicands []string) ([]string, error) {

	ePrefix := "BigIntMathMultiply.MultiplyNumStrOutputToArray() "

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return []string{}, nil
	}

	multiplierBINum, err := BigIntNum{}.NewNumStr(multiplier)

	if err != nil {
		return []string{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier, err.Error())
	}

	resultArray := make([]string, lenMultiplicands)

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewNumStr(multiplicands[i])

		if err != nil {
			return []string{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStr(multiplicands[i]) " +
					" multiplierBINum='%v' multiplicands[%v]='%v' Error='%v'. ",
					multiplierBINum.GetNumStr(), i, multiplicands[i], err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(multiplierBINum, multiplicandBINum )

		finalResult := bMultiply.MultiplyPair(bPair)

		resultArray[i], err = finalResult.Result.GetNumStrErr()

		if err != nil {
			return []string{},
				fmt.Errorf(ePrefix +
					"Error returned by finalResult.Result.GetNumStrErr() " +
					"i='%v' Error='%v'. ", i,  err.Error())
		}
	}

	return resultArray, nil
}

// MultiplyNumStrSeries - Receives one input parameter of Type string which is classified
//  as the 'multiplier'. The second input parameter is a comma delimited series of strings labeled,
// 'multiplicands'. The first element of the 'multiplicands' series is multiplied by the 'multiplier'
// to produce a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next
// element in the multiplicands series. This process is continued through the last element in the
// series. Afterwards, the combined final 'product' is returned as a Type 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result Decimal
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyNumStrSeries(
												multiplier string,
													multiplicands ... string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyNumStrSeries() "
	var err error

	finalResult := BigIntBasicMathResult{}.New()

	finalResult.Result, err = BigIntNum{}.NewNumStr(multiplier)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStr(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier, err.Error())
	}

	for _, multiplicand := range multiplicands {

		multiplicandBINum, err := BigIntNum{}.NewNumStr(multiplicand)

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStr(multiplicand) " +
					" multiplicand='%v' Error='%v'. ",
					multiplicand, err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result.CopyOut(), multiplicandBINum)

		finalResult = bMultiply.MultiplyPair(bPair)

	}

	return finalResult, nil
}

// MultiplyNumStrDto - Receives two NumStrDto instances and multiplies their
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
// 											Result.bigInt = product
//					}
//
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

// MultiplyNumStrDtoArray - Receives one NumStrDto which is classified as the 'multiplier'.
// The second input parameter is an array of NumStrDto Types labeled, 'multiplicands'. The
// first element of the 'multiplicands' array is multiplied by the 'multiplier' to produce
// a 'product'. That 'product' replaces the 'multiplier' and is multiplied by the next element
// in the multiplicands array. This process is continued through the last element in the array
// when the combined, final 'product' is returned as a Type 'BigIntBasicMathResult'.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier',
// 3 is the 'multiplicand' and 15 is the 'product' or result.
//
//							multiplier x multiplicand = product or result
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' as a BigIntBasicMathResult type.
//
// 					type BigIntBasicMathResult struct {
// 								Input BigIntPair
//											Input.Big1		= multiplier
//											Input.Big2		= multiplicand
//
// 								Result NumStrDto
// 											Result.bigInt = product
//					}
//
//
func (bMultiply BigIntMathMultiply) MultiplyNumStrDtoArray(
	multiplier NumStrDto,
	multiplicands []NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathMultiply.MultiplyNumStrDtoArray() "
	var err error

	finalResult := BigIntBasicMathResult{}

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return finalResult, nil
	}

	finalResult.Result, err = BigIntNum{}.NewNumStrDto(multiplier)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStrDto(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewNumStrDto(multiplicands[i])

		if err != nil {
			return BigIntBasicMathResult{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStrDto(multiplicands[i]) " +
					"multiplicands[%v]='%v' Error='%v'. ",
					i, multiplicands[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(finalResult.Result, multiplicandBINum )

		finalResult = bMultiply.MultiplyPair(bPair)

	}

	return finalResult, nil
}

// MultiplyNumStrDtoOutputToArray - Receives one input parameter of Type NumStrDto which
// is classified as the 'multiplier'. The second input parameter is an array of NumStrDto
// Types labeled, 'multiplicands'.
//
// Each element of the 'multiplicands' array is multiplied by the 'multiplier'. The result or
// 'product' is then stored in a results array which is returned to the calling function.
//
// In the multiplication operation, the number to be multiplied is called the "multiplicand",
// while the number of times the multiplicand is to be multiplied comes from the "multiplier".
// Usually the multiplier is placed first and the multiplicand is placed second.
//
// For example, in the problem 5 x 3 equals 15, the 5 is the 'multiplier', 3 is the 'multiplicand'
// and 15 is the 'product' or result.
//
// Example
// =======
// 										Multiplicands												Output
// Multiplier				    	Array														Array
//
//		3			x				multiplicands[0] = 2			=				  outputarray[0] =  6
//		3			x				multiplicands[1] = 3			=				  outputarray[1] =  9
//		3			x				multiplicands[2] = 4			=				  outputarray[2] = 12
//		3			x				multiplicands[3] = 5			=				  outputarray[3] = 15
//		3			x				multiplicands[4] = 6			=				  outputarray[4] = 18
//		3			x				multiplicands[5] = 7			=				  outputarray[5] = 21
//
//
// This method performs the multiplication operation described above and afterwards returns the
// result or 'product' in an Array of 'NumStrDtos' ([] NumStrDtos).
//
func (bMultiply BigIntMathMultiply) MultiplyNumStrDtoOutputToArray(
	multiplier NumStrDto,
	multiplicands []NumStrDto) ([]NumStrDto, error) {

	ePrefix := "BigIntMathMultiply.MultiplyNumStrDtoOutputToArray() "

	lenMultiplicands := len(multiplicands)

	if lenMultiplicands == 0 {
		return []NumStrDto{}, nil
	}

	multiplierBINum, err := BigIntNum{}.NewNumStrDto(multiplier)

	if err != nil {
		return []NumStrDto{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntNum{}.NewNumStrDto(multiplier) " +
				" multiplier='%v' Error='%v'. ",
				multiplier.GetNumStr(), err.Error())
	}

	resultArray := make([]NumStrDto, lenMultiplicands)

	for i:=0; i < lenMultiplicands; i++ {

		multiplicandBINum, err := BigIntNum{}.NewNumStrDto(multiplicands[i])

		if err != nil {
			return []NumStrDto{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewNumStrDto(multiplicands[i]) " +
					" multiplierBINum='%v' multiplicands[%v]='%v' Error='%v'. ",
					multiplierBINum.GetNumStr(), i, multiplicands[i].GetNumStr(), err.Error())
		}

		bPair := BigIntPair{}.NewBigIntNum(multiplierBINum, multiplicandBINum )

		finalResult := bMultiply.MultiplyPair(bPair)

		resultArray[i], err = finalResult.Result.GetNumStrDto()

		if err != nil {
			return []NumStrDto{},
				fmt.Errorf(ePrefix +
					"Error returned by finalResult.Result.GetNumStrDto() " +
					"i='%v' Error='%v'. ", i,  err.Error())
		}
	}

	return resultArray, nil
}

// MultiplyPair - Receives a BigIntPair instance and proceeds to multiply
// b1.BigIntNum by b2.BigIntNum.
//
// b1.BigIntNum x b2.BigIntNum = Result
//
// The result is returned as a BigIntBasicMathResult type.
//
func (bMultiply BigIntMathMultiply) MultiplyPair(bPair BigIntPair) BigIntBasicMathResult {

	b3 := big.NewInt(0).Mul(bPair.GetBig1BigInt(), bPair.GetBig2BigInt())

	bResult := BigIntBasicMathResult{}
	bResult.Input = bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(
								b3,
								bPair.Big1.GetPrecisionUint() + bPair.Big2.GetPrecisionUint())

	return bResult

}
