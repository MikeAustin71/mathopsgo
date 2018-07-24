package mathops

import (
	"fmt"
	"math/big"
	"errors"
)

type BigIntMathAdd struct {
	Input  BigIntPair
	Result BigIntNum
}

// AddBigIntNums - Adds two BigIntNums and returns the result in a new
// BigIntNum instance
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'b1'.
//
func (bAdd BigIntMathAdd) AddBigIntNums(b1, b2 BigIntNum) BigIntNum {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bAdd.AddPair(bPair)
}

// AddBigIntNumArray - Adds an Array of 'BigIntNum' types and returns the result
// as type 'BigIntNum'
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from the first element of the bNums array (bNums[0]).
//
func (bAdd BigIntMathAdd) AddBigIntNumArray(bNums []BigIntNum ) BigIntNum {

	finalResult := BigIntNum{}.NewZero(0)
	lenBNums := len(bNums)

	if lenBNums == 0 {
		return finalResult
	}

	numSeps := bNums[0].GetNumericSeparatorsDto()

	for i:=0; i < lenBNums; i++ {

		if i == 0 {
			finalResult = bNums[i].CopyOut()
			continue
		}

		finalResult =  bAdd.AddBigIntNums(finalResult, bNums[i])
	}

	finalResult.SetNumericSeparatorsDto(numSeps)

	return finalResult
}


// AddBigIntNumOutputToArray - The first input parameter to this method
// is a BigIntNum Type labeled, 'addend'.  The second element is an 
// array of BigIntNum types labeled 'bNums'. The 'addend' is added to 
// each element of the 'bNums' array with the result output to another
// array of BigIntNum types ([]BigIntNum) which is returned to the calling
// function.
//
// Example
// =======
// 										Multiplicands												Output
//  Addend   				    	Array														Array
//
//		3			+					bNums[0] = 2			=				  outputarray[0] =  5
//		3			+					bNums[1] = 3			=				  outputarray[1] =  6
//		3			+					bNums[2] = 4			=				  outputarray[2] =  7
//		3			+					bNums[3] = 5			=				  outputarray[3] =  8
//		3			+					bNums[4] = 6			=				  outputarray[4] =  9
//		3			+					bNums[5] = 9			=				  outputarray[5] = 12
//
// Each element of the the []BigIntNum 'result' array returned by this addition
// operation will contain numeric separators (decimal separator, thousands
// separator and currency symbol) which were copied from input parameter
// 'addend'.
//
func (bAdd BigIntMathAdd) AddBigIntNumOutputToArray(
													addend BigIntNum, 
														bNums []BigIntNum ) []BigIntNum {

	lenBNums := len(bNums)

	if lenBNums == 0 {
		return []BigIntNum{}
	}

	numSeps := addend.GetNumericSeparatorsDto()

	resultArray := make([]BigIntNum, lenBNums)

	for i:=0; i < lenBNums; i++ {

		bPair := BigIntPair{}.NewBigIntNum(addend, bNums[i])

		result := bAdd.addPairNoNumSeps(bPair)

		result.SetNumericSeparatorsDto(numSeps)

		resultArray[i] = result.CopyOut()

	}

	return resultArray
}

// AddBigIntNumSeries - Adds a series of BigIntNum types and returns the total in a
// BigIntNum instance.
//
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from the first element in input series 'bNums'.
//
func (bAdd BigIntMathAdd) AddBigIntNumSeries(bNums ... BigIntNum) BigIntNum {

	finalResult := BigIntNum{}.New()

	numSeps := NumericSeparatorDto{}

	for i, bNum := range bNums {

		if i == 0 {
			finalResult = bNum.CopyOut()
			numSeps = finalResult.GetNumericSeparatorsDto()
			continue
		}

		finalResult = bAdd.AddBigIntNums(finalResult, bNum)
	}

	finalResult.SetNumericSeparatorsDto(numSeps)

	return finalResult
}



// AddBigInts - Adds two *big.Int numbers. Each *big.Int number
// is passed to the method with an associated decimal place precision
// specification.
//
// The BigIntNum 'result' returned by this addition operation will contain
// default numeric separators (decimal separator, thousands separator and
// currency symbol).
//
func (bAdd BigIntMathAdd) AddBigInts(
														b1 *big.Int,
															precision1 uint,
																b2 *big.Int,
																	precision2 uint ) BigIntNum {

	b1Pair := BigIntPair{}.NewBase(b1, precision1, b2, precision2)

	finalResult := bAdd.addPairNoNumSeps(b1Pair)

	finalResult.SetNumericSeparatorsToDefaultIfEmpty()

	return finalResult
}

// AddDecimal - Receives two Decimal instances and adds their numeric values.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'dec1'.
//

func (bAdd BigIntMathAdd) AddDecimal(dec1, dec2 Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto() "

	numSeps := dec1.GetNumericSeparatorsDto()

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	finalResult := bAdd.addPairNoNumSeps(bPair)

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
		fmt.Errorf(ePrefix +
			"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' ", err.Error())
	}

	return finalResult, nil
}

// AddDecimalArray - Adds an array of 'Decimal' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from the first element of the input []Decimal 'decs' (decs[0]).
//
func (bAdd BigIntMathAdd) AddDecimalArray(decs []Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddDecimalArray() "

	finalResult := BigIntNum{}.New()
	var err error

	lenDecs := len(decs)

	if lenDecs == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: decs array is Empty!")
	}

	numSeps := decs[0].GetNumericSeparatorsDto()

	for i:= 0; i < lenDecs; i++ {

		if i == 0 {

			finalResult, err = BigIntNum{}.NewDecimal(decs[i])

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix +
						"Error returned by BigIntNum{}.NewDecimal(decs[i]). " +
						" i='%v' decs[i].GetNumStr()='%v' Error='%v' ",
						i, decs[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, &decs[i])

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &decs[i]). " +
					" i='%v' dec[i].GetNumStr()='%v' Error='%v' ", i, decs[i].GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)
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

// AddDecimalOutputToArray - The first input parameter to this method
// is a Decimal Type labeled, 'addend'.  The second element is an 
// array of Decimal types labeled 'decs'. The 'addend' is added to 
// each element of the 'decs' array with the result output to another
// array of Decimal types which is returned to the calling function.
//
// Example
// =======
// 										    decs										 Output
//  Addend   				    	Array											Array
//
//		3			+					decs[0] = 2			=				  outputarray[0] =  5
//		3			+					decs[1] = 3			=				  outputarray[1] =  6
//		3			+					decs[2] = 4			=				  outputarray[2] =  7
//		3			+					decs[3] = 5			=				  outputarray[3] =  8
//		3			+					decs[4] = 6			=				  outputarray[4] =  9
//		3			+					decs[5] = 9			=				  outputarray[5] = 12
//
//
// Each element in the []Decimal array 'result' returned by this addition
// operation will contain numeric separators (decimal separator, thousands
// separator and currency symbol) which were copied from input parameter
// 'addend'.
//
func (bAdd BigIntMathAdd) AddDecimalOutputToArray(
														addend Decimal, 
																decs []Decimal) ([]Decimal, error) {

	ePrefix := "BigIntMathAdd.AddDecimalOutputToArray() "

	lenDecs := len(decs)

	if lenDecs == 0 {
		return []Decimal{},
			errors.New(ePrefix + "Error: decs array is Empty!")
	}

	numSeps := addend.GetNumericSeparatorsDto()

	resultsArray := make([]Decimal, lenDecs)
	
	for i:= 0; i < lenDecs; i++ {

		bPair, err := BigIntPair{}.NewDecimal(addend, decs[i])

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(addend, decs[i]). " +
					" i='%v' dec[i].GetNumStr()='%v' Error='%v' ", i, decs[i].GetNumStr(), err.Error())
		}

		result := bAdd.addPairNoNumSeps(bPair)

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix +
					"Error returned by result.SetNumericSeparatorsDto(numSeps). " +
					" index='%v' Error='%v' ", i, err.Error())
		}

		resultsArray[i], err = result.GetDecimal()

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix +
					"Error returned by result.Result.GetDecimal(). " +
					" i='%v' decs[i].GetNumStr()='%v' Error='%v' ",
					i, decs[i].GetNumStr(), err.Error())
		}

	}

	return resultsArray, nil
}

// AddDecimalSeries - Adds a series of 'Decimal' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
// The BigIntNum 'result' returned by this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) which
// were copied from input series element 'decs[0]'.
//
func (bAdd BigIntMathAdd) AddDecimalSeries(decs ... Decimal) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddDecimalSeries() "

	finalResult := BigIntNum{}.New()
	var err error

	if len(decs) == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: decs series is Empty!")
	}

	numSeps := NumericSeparatorDto{}

	for i, dec := range decs {

		if i == 0 {

			finalResult, err = BigIntNum{}.NewDecimal(dec)

			if err != nil {
				return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(dec). " +
					" i='%v' dec.GetNumStr()='%v' Error='%v' ",
					i, dec.GetNumStr(), err.Error())
			}

			numSeps = dec.GetNumericSeparatorsDto()

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, &dec)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &dec). " +
					" i='%v' dec.GetNumStr()='%v' Error='%v' ", i, dec.GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)

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

// AddIntAry - Receives two IntAry instances and adds their numeric values.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) which
// were copied from input parameter 'ia1'.
//
func (bAdd BigIntMathAdd) AddIntAry(ia1, ia2 IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto() "

	err := ia1.IsIntAryValid(ePrefix + "ia1 Invalid: ")

	if err != nil {
		return BigIntNum{}.New(),
			err
	}

	err = ia2.IsIntAryValid(ePrefix + "ia2 Invalid: ")

	if err != nil {
		return BigIntNum{}.New(),
			err
	}

	numSeps := ia1.GetNumericSeparatorsDto()

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
				"Error='%v' ", err.Error())
	}

	finalResult := bAdd.addPairNoNumSeps(bPair)

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' ", err.Error())
	}

	return finalResult, nil
}

// AddIntAryArray - Receives an array of IntAry objects and totals their numeric values.
// The total numeric value is returned in a BigIntNum instance.
//
// The BigIntNum 'result' returned by this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) which
// were copied from the first element of the input array ([]IntAry) 'iarys'.
//
func (bAdd BigIntMathAdd) AddIntAryArray(iarys []IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddIntAryArray() "

	finalResult := BigIntNum{}.New()

	lenIaArray := len(iarys)

	if lenIaArray == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: iarys array is Empty!")
	}

	numSeps := iarys[0].GetNumericSeparatorsDto()

	var err error

	for i := 0; i < lenIaArray; i++ {

		err = iarys[i].IsIntAryValid(ePrefix +
			fmt.Sprintf("iarys[%v] is INVALID!", i))

		if err != nil {
			return BigIntNum{}.New(), err
		}

		if i == 0 {

			finalResult, err = BigIntNum{}.NewIntAry(iarys[i])

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(iarys[i]). " +
						" i='%v' NumStr='%v' Error='%v' ", i, iarys[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, &iarys[i])

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &iarys[i]). " +
					" i='%v' iarys[i].GetNumStr()='%v' Error='%v' ",
					i, iarys[i].GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)
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

// AddIntAryOutputToArray - The first input parameter to this method
// is an IntAry Type labeled, 'addend'.  The second element is an 
// array of IntAry types labeled 'iarys'. The 'addend' is added to 
// each element of the 'iarys' array with the result output to another
// array of IntAry types which is returned to the calling function.
//
// Example
// =======
// 										    decs										 Output
//  Addend   				    	Array											Array
//
//		3			+					iarys[0] = 2			=				  outputarray[0] =  5
//		3			+					iarys[1] = 3			=				  outputarray[1] =  6
//		3			+					iarys[2] = 4			=				  outputarray[2] =  7
//		3			+					iarys[3] = 5			=				  outputarray[3] =  8
//		3			+					iarys[4] = 6			=				  outputarray[4] =  9
//		3			+					iarys[5] = 9			=				  outputarray[5] = 12
//
// Each element of the []IntAry returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency symbol)
// which were copied from input parameter 'addend'.
//
func (bAdd BigIntMathAdd) AddIntAryOutputToArray(
														addend IntAry,
															iarys []IntAry) ([]IntAry, error) {

	ePrefix := "BigIntMathAdd.AddIntAryOutputToArray() "

	err := addend.IsIntAryValid(ePrefix + "'addend' INVALID! ")

	if err != nil {
		return []IntAry{}, err
	}

	lenIaArray := len(iarys)

	if lenIaArray == 0 {
		return []IntAry{},
			errors.New(ePrefix + "Error: iarys array is Empty!")
	}

	numSeps := addend.GetNumericSeparatorsDto()

	resultsArray := make([]IntAry, lenIaArray)
	
	for i := 0; i < lenIaArray; i++ {

		err = iarys[i].IsIntAryValid(ePrefix +
			fmt.Sprintf("iarys[%v] INVALID!", i))

		if err != nil {
			return []IntAry{}, err
		}

		bPair, err := BigIntPair{}.NewIntAry(addend, iarys[i])

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewIntAry(addend, iarys[i]). " +
					" i='%v' iarys[i].GetNumStr()='%v' Error='%v' ",
					i, iarys[i].GetNumStr(), err.Error())
		}

		result := bAdd.addPairNoNumSeps(bPair)

		err = result.SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by result.SetNumericSeparatorsDto(numSeps). " +
					"index='%v' Error='%v' ",	i, err.Error())
		}

		resultsArray[i], err = result.GetIntAry()
	}

	return resultsArray, nil
}

// AddIntArySeries - Adds a series of IntAry objects and returns the total in a
// BigIntNum.
//
// The BigIntNum result of this addition operation will contain numeric separators
// (decimal separator, thousands separator and currency symbol) which were copied
// from the first element input series 'iarys'.
//
func (bAdd BigIntMathAdd) AddIntArySeries(iarys ... IntAry) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddIntArySeries() "

	finalResult := BigIntNum{}.New()

	if len(iarys) == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: 'iarys' series is Empty!")
	}

	var err error
	numSeps := NumericSeparatorDto{}

	for i, ia := range iarys {

		err = ia.IsIntAryValid(ePrefix +
			fmt.Sprintf("ia index='%v' INVALID!", i))

		if err != nil {
			return BigIntNum{}.New(), err
		}

		if i == 0 {

			numSeps = ia.GetNumericSeparatorsDto()

			finalResult, err = BigIntNum{}.NewIntAry(ia)

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(ia). " +
						" i='%v' NumStr='%v' Error='%v' ", i, ia.GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, &ia)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &ia). " +
					" i='%v' ia.GetNumStr()='%v' Error='%v' ",
					i, ia.GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)
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

// AddINumMgr - Receives two objects which implement the INumMgr Interface
// and adds their numeric values.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry.
//
// The result is returned as an instance of Type, 'BigIntNum'.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) which were copied from the input parameter 'num1'.
//
func (bAdd BigIntMathAdd) AddINumMgr(num1, num2 INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddINumMgr() "

	numSeps := num1.GetNumericSeparatorsDto()

	bPair, err := BigIntPair{}.NewINumMgr(num1, num2)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewINumMgr(num1, num2). " +
				"num1.GetNumStr()='%v', num2.GetNumStr()='%v' Error='%v' ",
				num1.GetNumStr(), num2.GetNumStr(), err.Error())
	}

	finalResult := bAdd.addPairNoNumSeps(bPair)

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' ", err.Error())
	}

	return finalResult, nil
}

// AddINumMgrArray - Adds an array of objects which implement the 'INumMgr'
// interface. The combined total of the numeric values from these objects
// is returned as an instance of Type, 'BigIntNum'.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry. This allows the user to mix different types in
// a single array and add their numeric values.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) which were copied from the first element of the input parameter
// array, nums (nums[0]).
//
func (bAdd BigIntMathAdd) AddINumMgrArray(nums []INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrArray() "

	finalResult := BigIntNum{}.New()
	var err error

	lenDecs := len(nums)

	if lenDecs == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: 'nums' array is Empty!")
	}

	numSeps := nums[0].GetNumericSeparatorsDto()

	for i:= 0; i < lenDecs; i++ {

		if i == 0 {

			finalResult, err = BigIntNum{}.NewINumMgr(nums[i])

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix +
						"Error returned by BigIntNum{}.NewINumMgr(nums[i]). " +
						" i='%v' nums[i].GetNumStr()='%v' Error='%v' ",
						i, nums[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, nums[i])

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &nums[i]). " +
					" i='%v' nums[i].GetNumStr()='%v' Error='%v' ", i, nums[i].GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' ", err.Error())
	}

	return finalResult, nil
}

// AddINumMgrOutputToArray - The first input parameter to this method
// is an object that implements the  INumMgr interface and is labeled,
// 'addend'.  The second element is an array of objects which implement
// the INumMgr interface. This array is labeled labeled, 'numMgrs'. The
// 'addend' is added to each element of the 'numMgrs' array with the result
// output to another INumMgr array which is returned to the calling function.
//
// Example
// =======
// 										    numMgrs										 Output
//  Addend   				    	Array											Array
//
//		3			+					numMgrs[0] = 2			=				  outputarray[0] =  5
//		3			+					numMgrs[1] = 3			=				  outputarray[1] =  6
//		3			+					numMgrs[2] = 4			=				  outputarray[2] =  7
//		3			+					numMgrs[3] = 5			=				  outputarray[3] =  8
//		3			+					numMgrs[4] = 6			=				  outputarray[4] =  9
//		3			+					numMgrs[5] = 9			=				  outputarray[5] = 12
//
// Each element of the returned []INumMgr array will contain numeric separators
// (decimal separator, thousands separator and currency symbol) which were
// copied from the input parameter 'addend'.
//
func (bAdd BigIntMathAdd) AddINumMgrOutputToArray(
																	addend INumMgr,
																		numMgrs []INumMgr) ([]INumMgr, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrOutputToArray() "

	lenDecs := len(numMgrs)

	if lenDecs == 0 {
		return []INumMgr{},
			errors.New(ePrefix + "Error: 'numMgrs' array is Empty!")
	}

	numSeps := addend.GetNumericSeparatorsDto()

	resultsArray := make([]INumMgr, lenDecs) 
		
	for i:= 0; i < lenDecs; i++ {

		bPair, err := BigIntPair{}.NewINumMgr(addend, numMgrs[i])

		if err != nil {
			return []INumMgr{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(addend, numMgrs[i]). " +
					" i='%v' numMgrs[i].GetNumStr()='%v' Error='%v' ", i, numMgrs[i].GetNumStr(), err.Error())
		}

		bIntNum := bAdd.addPairNoNumSeps(bPair)

		err = bIntNum.SetNumericSeparatorsDto(numSeps)

		if err != nil {
			return []INumMgr{},
				fmt.Errorf(ePrefix +
					"Error returned by bIntNum.SetNumericSeparatorsDto(numSeps). " +
					"Error='%v' ", err.Error())
		}


		resultsArray[i] = &bIntNum
	}

	return resultsArray, nil
}


// AddINumMgrSeries - Adds a series of objects which implement the 'INumMgr'
// interface. The combined total of the numeric values from these objects
// is returned as an instance of Type, 'BigIntNum'.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry. This allows the user to mix different types in
// a single array and add their numeric values.
//
// The returned BigIntNum result of this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) which were copied from the first element of the input parameter
// 'nums'.
//
func (bAdd BigIntMathAdd) AddINumMgrSeries(nums ... INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrSeries() "

	finalResult := BigIntNum{}.New()

	if len(nums) == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: 'nums' series is Empty!")
	}


	var err error

	numSeps := NumericSeparatorDto{}

	for i, num := range nums {

		if i == 0 {

			numSeps = num.GetNumericSeparatorsDto()

			finalResult, err = BigIntNum{}.NewINumMgr(num)

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix +
						"Error returned by BigIntNum{}.NewINumMgr(num). " +
						" i='%v' num.GetNumStr()='%v' Error='%v' ",
						i, num.GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, num)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &num). " +
					" i='%v' num.GetNumStr()='%v' Error='%v' ",
					i, num.GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' ", err.Error())
	}

	return finalResult, nil
}

// AddNumStr - Receives two number strings and adds their numeric values.
//
// The two input parameters, 'n1NumStr' and 'n2NumStr' must be formatted
// as strings of numeric digits, or number strings. Number strings may have
// a leading minus sign ('-') to indicate the numeric sign value. In addition,
// the string of numeric digits may include a decimal point ('.') to separate
// fractional digits.
//
// The result is returned as type BigIntNum.
//
// The returned BigIntNum result of this addition operation will contain
// default numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bAdd BigIntMathAdd) AddNumStr(n1NumStr, n2NumStr string) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1NumStr, n2NumStr)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1NumStr, n2NumStr). " +
				"Error='%v' ", err.Error())
	}

	return bAdd.addPairNoNumSeps(bPair), nil
}

// AddNumStrArray - Adds a series of number strings and returns the combined total
// as an instance of Type 'BigIntNum'.
//
// All the elements of the 'numStrs' array must be formatted as strings of numeric
// digits or number strings. Number strings may have a leading minus sign ('-')
// to indicate the numeric sign value. In addition, the string of numeric digits
// may include a decimal point ('.') to separate fraction digits.
//
// The returned BigIntNum result of this addition operation will contain
// default numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bAdd BigIntMathAdd) AddNumStrArray(numStrs []string) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrArray() "

	finalResult := BigIntNum{}.New()
	var err error

	lenNumStrs := len(numStrs)

	if lenNumStrs == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: 'numStrs' array is Empty!")
	}

	for i:= 0; i < lenNumStrs; i++ {

		if i == 0 {

			finalResult, err = BigIntNum{}.NewNumStr(numStrs[i])

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
						" i='%v' NumStr='%v' Error='%v' ", i, numStrs[i], err.Error())
			}

			continue
		}

		b2Num, err := BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
					" i='%v' NumStr='%v' Error='%v' ", i, numStrs[i], err.Error())
		}

		finalResult = bAdd.AddBigIntNums(finalResult, b2Num)
	}

	return finalResult, nil
}

// AddNumStrOutputToArray - The first input parameter to this method
// is a string Type labeled, 'addend'.  The second element is an 
// array of string types labeled 'numStrs'. The 'addend' is added to 
// each element of the 'numStrs' array with the result output to another
// array of string types which is returned to the calling function.
//
// Input parameters 'addend' and all elements of the 'numStrs' array must
// be formatted as strings of numeric digits or number strings. Number strings
// may have a leading minus sign ('-') to indicate the numeric sign value. In
// addition, the string of numeric digits may include a decimal point ('.') to
// separate fraction digits.
//
// Example
// =======
// 										    numStrs										 Output
//  Addend   				    	Array											Array
//
//		3			+					numStrs[0] = 2			=				  outputarray[0] =  5
//		3			+					numStrs[1] = 3			=				  outputarray[1] =  6
//		3			+					numStrs[2] = 4			=				  outputarray[2] =  7
//		3			+					numStrs[3] = 5			=				  outputarray[3] =  8
//		3			+					numStrs[4] = 6			=				  outputarray[4] =  9
//		3			+					numStrs[5] = 9			=				  outputarray[5] = 12
//
//
func (bAdd BigIntMathAdd) AddNumStrOutputToArray(
													addend string, 
														numStrs []string) ([]string, error) {

	ePrefix := "BigIntMathAdd.AddNumStrOutputToArray() "

	lenNumStrs := len(numStrs)

	if lenNumStrs == 0 {
		return []string{},
			errors.New(ePrefix + "Error: 'numStrs' array is Empty!")
	}

	bINumAddend, err := BigIntNum{}.NewNumStr(addend)

	if err != nil {
		return []string{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(addend). " +
				" addend='%v' Error='%v' ", addend, err.Error())
	}
	
	resultsArray := make([]string, lenNumStrs)
	
	for i:= 0; i < lenNumStrs; i++ {

		b2Num, err := BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			return []string{},
				fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
					" i='%v' NumStr='%v' Error='%v' ", i, numStrs[i], err.Error())
		}

		result := bAdd.addPairNoNumSeps(BigIntPair{}.NewBigIntNum(bINumAddend, b2Num))

		resultsArray[i] = result.GetNumStr()
		
	}

	return resultsArray, nil
}

// AddNumStrSeries - Adds a series of number strings and returns
// the combined total as an instance of Type, 'BigIntNum'.
//
// All the elements of the 'numStrs' series must be formatted as strings of numeric
// digits, or number strings. Number strings may have a leading minus sign ('-')
// to indicate the numeric sign value. In addition, the string of numeric digits
// may include a decimal point ('.') to separate fractional digits.
//
// The returned BigIntNum result of this addition operation will contain
// default numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bAdd BigIntMathAdd) AddNumStrSeries(numStrs ... string) (BigIntNum, error) {
	ePrefix := "BigIntMathAdd.AddNumStrSeries() "

	finalResult := BigIntNum{}.New()

	if len(numStrs) == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: 'numStrs' series is Empty!")
	}

	var err error

	for i, numStr := range numStrs {

		if i == 0 {

			finalResult, err = BigIntNum{}.NewNumStr(numStr)

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStr). " +
						" i='%v' NumStr='%v' Error='%v' ", i, numStr, err.Error())
			}

			continue
		}

		b2Num, err := BigIntNum{}.NewNumStr(numStr)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStr). " +
					" i='%v' NumStr='%v' Error='%v' ", i, numStr, err.Error())
		}


		result := bAdd.AddBigIntNums(finalResult, b2Num)

		finalResult = result.CopyOut()
	}

	finalResult.SetNumericSeparatorsToDefaultIfEmpty()

	return finalResult, nil
}

// AddNumStrDto - Receives two NumStrDto instances and adds their numeric values.
//
// The result is returned as type BigIntNum.
//
// The returned BigIntNum result of this addition operation will contain
// default numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bAdd BigIntMathAdd) AddNumStrDto(n1Dto, n2Dto NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(n1Dto, n2Dto)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(n1Dto, n2Dto). " +
				"Error='%v' ", err.Error())
	}

	finalResult := bAdd.addPairNoNumSeps(bPair)

	finalResult.SetNumericSeparatorsToDefaultIfEmpty()

	return finalResult, nil
}

// AddNumStrDtoArray - Adds an array of 'NumStrDto' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
// The returned BigIntNum result of this addition operation will contain
// default numeric separators (decimal separator, thousands separator and currency
// symbol).
//
func (bAdd BigIntMathAdd) AddNumStrDtoArray(nDtos []NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddDecimalArray() "
	finalResult := BigIntNum{}.New()
	var err error

	lenDecs := len(nDtos)

	if lenDecs == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: 'nDtos' array is Empty!")
	}

	for i:= 0; i < lenDecs; i++ {

		if i == 0 {

			finalResult, err = BigIntNum{}.NewNumStrDto(nDtos[i])

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(nDtos[i]). " +
						" i='%v' nDtos[i].GetNumStr()='%v' Error='%v' ", i, nDtos[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, &nDtos[i])

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &nDtos[i]). " +
					" i='%v' dec[i].GetNumStr()='%v' Error='%v' ", i, nDtos[i].GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)
	}

	finalResult.SetNumericSeparatorsToDefaultIfEmpty()

	return finalResult, nil

}

// AddNumStrDtoOutputToArray - The first input parameter to this method
// is a NumStrDto Type labeled, 'addend'.  The second element is an 
// array of NumStrDto types labeled 'nDtos'. The 'addend' is added to 
// each element of the 'nDtos' array with the result output to another
// array of NumStrDto types which is returned to the calling function.
//
// Example
// =======
// 										    nDtos										 Output
//  Addend   				    	Array											Array
//
//		3			+					nDtos[0] = 2			=				  outputarray[0] =  5
//		3			+					nDtos[1] = 3			=				  outputarray[1] =  6
//		3			+					nDtos[2] = 4			=				  outputarray[2] =  7
//		3			+					nDtos[3] = 5			=				  outputarray[3] =  8
//		3			+					nDtos[4] = 6			=				  outputarray[4] =  9
//		3			+					nDtos[5] = 9			=				  outputarray[5] = 12
//
// Each element in the returned array of []NumStrDto resulting of this addition
// operation will contain numeric separators (decimal separator, thousands separator
// and currency symbol) copied from input parameter 'addend'.
//
func (bAdd BigIntMathAdd) AddNumStrDtoOutputToArray(
													addend NumStrDto,
															nDtos []NumStrDto) ([]NumStrDto, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDtoOutputToArray() "

	err := addend.IsNumStrDtoValid(ePrefix + "'addend' is INVALID! ")

	if err != nil {
		return []NumStrDto{}, err
	}

	lenDecs := len(nDtos)

	if lenDecs == 0 {
		return []NumStrDto{},
			errors.New(ePrefix + "Error: 'nDtos' array is Empty!")
	}

	numSeps := addend.GetNumericSeparatorsDto()

	resultsArray := make([]NumStrDto, lenDecs)

	for i:= 0; i < lenDecs; i++ {

		err := nDtos[i].IsNumStrDtoValid(ePrefix +
			fmt.Sprintf("nDtos[%v] INVALID! ", i))

		if err != nil {
			return []NumStrDto{}, err
		}


		bPair, err := BigIntPair{}.NewNumStrDto(addend, nDtos[i])

		if err != nil {
			return []NumStrDto{},
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(addend, nDtos[i]). " +
					" i='%v' dec[i].GetNumStr()='%v' Error='%v' ", i, nDtos[i].GetNumStr(), err.Error())
		}

		result := bAdd.addPairNoNumSeps(bPair)

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
					"Error returned by result.Result.GetNumStrDto(). " +
					" i='%v' nDtos[i].GetNumStr()='%v' Error='%v' ",
					i, nDtos[i].GetNumStr(), err.Error())
		}

	}

	return resultsArray, nil
}

// AddNumStrDtoSeries - Adds a series of 'NumStrDto' types and returns the combined total
// as an instance of Type, 'BigIntNum'.
//
// The returned BigIntNum resulting of this addition operation will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied
// from the first element of the input series 'nDtos'.
//
func (bAdd BigIntMathAdd) AddNumStrDtoSeries(nDtos ... NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDtoSeries() "

	finalResult := BigIntNum{}.New()

	if len(nDtos) == 0 {
		return finalResult,
			errors.New(ePrefix + "Error: 'nDtos' series is Empty!")
	}

	var err error
	numSeps := NumericSeparatorDto{}

	for i, nDto := range nDtos {

		err = nDto.IsNumStrDtoValid(ePrefix +
			fmt.Sprintf("nDto Index='%v' is INVALID! ",i))

		if err != nil {
			return BigIntNum{}.New(), err
		}

		if i == 0 {

			numSeps = nDto.GetNumericSeparatorsDto()

			finalResult, err = BigIntNum{}.NewNumStrDto(nDto)

			if err != nil {
				return BigIntNum{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(nDto). " +
						" i='%v' NumStr='%v' Error='%v' ", i, nDto.GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult, &nDto)

		if err != nil {
			return BigIntNum{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &nDto). " +
					" i='%v' nDto.GetNumStr()='%v' Error='%v' ",
					i, nDto.GetNumStr(), err.Error())
		}

		finalResult = bAdd.addPairNoNumSeps(bPair)
	}

	err = finalResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return BigIntNum{}.New(),
			fmt.Errorf(ePrefix +
				"Error returned by finalResult.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' ", err.Error())
	}

	return finalResult, nil
}

// AddPair - Receives a BigIntPair instance and proceeds to add b1.BigIntNum
// to b2.BigIntNum.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this addition operation will contain
// numeric separators (decimal separator, thousands separator and currency
// symbol) copied from b1.BigIntNum.
//
func (bAdd BigIntMathAdd) AddPair(bPair BigIntPair) BigIntNum {

	numSeps := bPair.Big1.GetNumericSeparatorsDto()

	finalResult := bAdd.addPairNoNumSeps(bPair)

	finalResult.SetNumericSeparatorsDto(numSeps)

	return finalResult
}

// addPairNoNumSeps - Receives a BigIntPair and proceeds to add b1.BigIntNum to
// b2.BigIntNum.
//
// The result is returned as type BigIntNum.
//
// The BigIntNum 'result' returned by this subtraction operation will contain
// default numeric separators (decimal separator, thousands separator and
// currency symbol).
//
func (bAdd BigIntMathAdd) addPairNoNumSeps(bPair BigIntPair) BigIntNum {

	bPair.MakePrecisionsEqual()

	b3 := big.NewInt(0).Add(bPair.GetBig1BigInt(), bPair.GetBig2BigInt())

	bResult := BigIntNum{}.NewBigInt(b3, bPair.Big2.GetPrecisionUint())

	return bResult
}
