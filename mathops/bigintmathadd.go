package mathops

import (
	"fmt"
	"math/big"
)

type BigIntMathAdd struct {
	Input  BigIntPair
	Result BigIntNum
}

// AddBigIntNums - Adds two BigIntNums and returns the result in a new
// BigIntBasicMathResult instance
//
func (bAdd BigIntMathAdd) AddBigIntNums(b1, b2 BigIntNum) BigIntBasicMathResult {

	bPair := BigIntPair{}.NewBigIntNum(b1, b2)

	return bAdd.AddPair(bPair)
}

// AddBigIntNumArray - Adds an Array of 'BigIntNum' types and returns the result
// as type 'BigIntBasicMathResult'
//
func (bAdd BigIntMathAdd) AddBigIntNumArray(bNums []BigIntNum ) BigIntBasicMathResult {

	finalResult := BigIntBasicMathResult{}.New()
	lenBNums := len(bNums)

	if lenBNums == 0 {
		return finalResult
	}

	for i:=0; i < lenBNums; i++ {

		if i == 0 {
			finalResult.Result = bNums[i].CopyOut()
			continue
		}

		result := bAdd.AddBigIntNums(finalResult.Result, bNums[i])

		finalResult.Result = result.Result.CopyOut()

	}

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
//
func (bAdd BigIntMathAdd) AddBigIntNumOutputToArray(
													addend BigIntNum, 
														bNums []BigIntNum ) []BigIntNum {

	lenBNums := len(bNums)

	if lenBNums == 0 {
		return []BigIntNum{}
	}

	resultArray := make([]BigIntNum, lenBNums)

	for i:=0; i < lenBNums; i++ {

		bPair := BigIntPair{}.NewBigIntNum(addend, bNums[i])

		result := bAdd.AddPair(bPair)

		resultArray[i] = result.Result.CopyOut()

	}

	return resultArray
}

// AddBigIntNumSeries - Adds a series of BigIntNum types and returns the total in a
// BigIntBasicMathResult instance.
//
func (bAdd BigIntMathAdd) AddBigIntNumSeries(bNums ... BigIntNum) BigIntBasicMathResult {

	finalResult := BigIntBasicMathResult{}.New()

	for i, bNum := range bNums {

		if i == 0 {
			finalResult.Result = bNum.CopyOut()
			continue
		}

		result := bAdd.AddBigIntNums(finalResult.Result, bNum)

		finalResult.Result = result.Result.CopyOut()

	}

	return finalResult
}



// AddBigInts - Adds two *big.Int numbers. Each *big.Int number
// is passed to the method with an associated decimal place precision
// specification.
//
func (bAdd BigIntMathAdd) AddBigInts(
	b1 *big.Int,
	precision1 uint,
	b2 *big.Int,
	precision2 uint ) BigIntBasicMathResult {

	b1Pair := BigIntPair{}.NewBase(b1, precision1, b2, precision2)

	return bAdd.AddPair(b1Pair)
}

// AddDecimal - Receives two Decimal instances and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bAdd BigIntMathAdd) AddDecimal(dec1, dec2 Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto() "

	bPair, err := BigIntPair{}.NewDecimal(dec1, dec2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	return bAdd.AddPair(bPair), nil
}

// AddDecimalArray - Adds an array of 'Decimal' types and returns the combined total
// as an instance of Type, 'BigIntBasicMathResult'.
//
func (bAdd BigIntMathAdd) AddDecimalArray(decs []Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddDecimalArray() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	lenDecs := len(decs)

	if lenDecs == 0 {
		return finalResult, nil
	}

	for i:= 0; i < lenDecs; i++ {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewDecimal(decs[i])

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix +
						"Error returned by BigIntNum{}.NewDecimal(decs[i]). " +
						" i='%v' decs[i].GetNumStr()='%v' Error='%v' ",
						i, decs[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, &decs[i])

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &decs[i]). " +
					" i='%v' dec[i].GetNumStr()='%v' Error='%v' ", i, decs[i].GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
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
func (bAdd BigIntMathAdd) AddDecimalOutputToArray(
														addend Decimal, 
																decs []Decimal) ([]Decimal, error) {

	ePrefix := "BigIntMathAdd.AddDecimalOutputToArray() "

	lenDecs := len(decs)

	if lenDecs == 0 {
		return []Decimal{}, nil
	}

	resultsArray := make([]Decimal, lenDecs)
	
	for i:= 0; i < lenDecs; i++ {

		bPair, err := BigIntPair{}.NewDecimal(addend, decs[i])

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(addend, decs[i]). " +
					" i='%v' dec[i].GetNumStr()='%v' Error='%v' ", i, decs[i].GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		resultsArray[i], err = result.Result.GetDecimal()

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
// as an instance of Type, 'BigIntBasicMathResult'.
//
func (bAdd BigIntMathAdd) AddDecimalSeries(decs ... Decimal) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddDecimalSeries() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	for i, dec := range decs {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewDecimal(dec)

			if err != nil {
				return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntNum{}.NewDecimal(dec). " +
					" i='%v' dec.GetNumStr()='%v' Error='%v' ",
					i, dec.GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, &dec)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &dec). " +
					" i='%v' dec.GetNumStr()='%v' Error='%v' ", i, dec.GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// AddIntAry - Receives two IntAry instances and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bAdd BigIntMathAdd) AddIntAry(ia1, ia2 IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto() "

	bPair, err := BigIntPair{}.NewIntAry(ia1, ia2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewIntAry(ia1, ia2). " +
				"Error='%v' ", err.Error())
	}

	return bAdd.AddPair(bPair), nil
}

// AddIntAryArray - Receives an array of IntAry objects and totals their numeric values.
// The total numeric value is returned in a BigIntBasicMathResult instance.
//
func (bAdd BigIntMathAdd) AddIntAryArray(iarys []IntAry) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddIntAryArray() "

	finalResult := BigIntBasicMathResult{}.New()

	var err error

	lenIaArray := len(iarys)

	for i := 0; i < lenIaArray; i++ {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewIntAry(iarys[i])

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(iarys[i]). " +
						" i='%v' NumStr='%v' Error='%v' ", i, iarys[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, &iarys[i])

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &iarys[i]). " +
					" i='%v' iarys[i].GetNumStr()='%v' Error='%v' ",
					i, iarys[i].GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
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
//
func (bAdd BigIntMathAdd) AddIntAryOutputToArray(
														addend IntAry,
															iarys []IntAry) ([]IntAry, error) {

	ePrefix := "BigIntMathAdd.AddIntAryOutputToArray() "

	lenIaArray := len(iarys)

	if lenIaArray == 0 {
		return []IntAry{}, nil
	}
	
	resultsArray := make([]IntAry, lenIaArray)
	
	for i := 0; i < lenIaArray; i++ {

		bPair, err := BigIntPair{}.NewIntAry(addend, iarys[i])

		if err != nil {
			return []IntAry{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewIntAry(addend, iarys[i]). " +
					" i='%v' iarys[i].GetNumStr()='%v' Error='%v' ",
					i, iarys[i].GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		resultsArray[i], err = result.Result.GetIntAry()
	}

	return resultsArray, nil
}

// AddIntArySeries - Adds a series of IntAry objects and returns the total in a
// BigIntBasicMathResult.
//
func (bAdd BigIntMathAdd) AddIntArySeries(iarys ... IntAry) (BigIntBasicMathResult, error) {
	ePrefix := "BigIntMathAdd.AddIntArySeries() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	for i, ia := range iarys {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewIntAry(ia)

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(ia). " +
						" i='%v' NumStr='%v' Error='%v' ", i, ia.GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, &ia)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &ia). " +
					" i='%v' ia.GetNumStr()='%v' Error='%v' ",
					i, ia.GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// AddINumMgr - Receives two objects which implement the INumMgr Interface
// and adds their numeric values.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry.
//
// The result is returned as an instance of Type, 'BigIntBasicMathResult'.
//
func (bAdd BigIntMathAdd) AddINumMgr(num1, num2 INumMgr) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddINumMgr() "

	bPair, err := BigIntPair{}.NewINumMgr(num1, num2)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix +
				"Error returned by NBigIntPair{}.NewINumMgr(num1, num2). " +
				"num1.GetNumStr()='%v', num2.GetNumStr()='%v' Error='%v' ",
				num1.GetNumStr(), num2.GetNumStr(), err.Error())
	}

	return bAdd.AddPair(bPair), nil
}

// AddINumMgrArray - Adds an array of objects which implement the 'INumMgr'
// interface. The combined total of the numeric values from these objects
// is returned as an instance of Type, 'BigIntBasicMathResult'.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry. This allows the user to mix different types in
// a single array and add their numeric values.
//
func (bAdd BigIntMathAdd) AddINumMgrArray(nums []INumMgr) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrArray() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	lenDecs := len(nums)

	if lenDecs == 0 {
		return finalResult, nil
	}

	for i:= 0; i < lenDecs; i++ {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewINumMgr(nums[i])

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix +
						"Error returned by BigIntNum{}.NewINumMgr(nums[i]). " +
						" i='%v' nums[i].GetNumStr()='%v' Error='%v' ",
						i, nums[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, nums[i])

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &nums[i]). " +
					" i='%v' nums[i].GetNumStr()='%v' Error='%v' ", i, nums[i].GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
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
//
func (bAdd BigIntMathAdd) AddINumMgrOutputToArray(
																	addend INumMgr,
																		numMgrs []INumMgr) ([]INumMgr, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrOutputToArray() "

	lenDecs := len(numMgrs)

	if lenDecs == 0 {
		return []INumMgr{}, nil
	}

	resultsArray := make([]INumMgr, lenDecs) 
		
	for i:= 0; i < lenDecs; i++ {

		bPair, err := BigIntPair{}.NewINumMgr(addend, numMgrs[i])

		if err != nil {
			return []INumMgr{},
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(addend, numMgrs[i]). " +
					" i='%v' numMgrs[i].GetNumStr()='%v' Error='%v' ", i, numMgrs[i].GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)
		
		bIntNum := result.Result.CopyOut()
		
		resultsArray[i] = &bIntNum
	}

	return resultsArray, nil
}


// AddINumMgrSeries - Adds a series of objects which implement the 'INumMgr'
// interface. The combined total of the numeric values from these objects
// is returned as an instance of Type, 'BigIntBasicMathResult'.
//
// The INumMgr interface is implemented by types, BigIntNum, Decimal,
// NumStrDto and IntAry. This allows the user to mix different types in
// a single array and add their numeric values.
//
func (bAdd BigIntMathAdd) AddINumMgrSeries(nums ... INumMgr) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddINumMgrSeries() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	for i, num := range nums {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewINumMgr(num)

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix +
						"Error returned by BigIntNum{}.NewINumMgr(num). " +
						" i='%v' num.GetNumStr()='%v' Error='%v' ",
						i, num.GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, num)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &num). " +
					" i='%v' num.GetNumStr()='%v' Error='%v' ",
					i, num.GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// AddNumStr - Receives two number strings and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bAdd BigIntMathAdd) AddNumStr(n1NumStr, n2NumStr string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddNumStr() "

	bPair, err := BigIntPair{}.NewNumStr(n1NumStr, n2NumStr)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStr(n1NumStr, n2NumStr). " +
				"Error='%v' ", err.Error())
	}

	return bAdd.AddPair(bPair), nil
}

// AddNumStrArray - Adds a series of number strings and returns the combined total
// as an instance of Type 'BigIntBasicMathResult'.
//
func (bAdd BigIntMathAdd) AddNumStrArray(numStrs []string) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddNumStrArray() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	lenNumStrs := len(numStrs)

	if lenNumStrs == 0 {
		return finalResult, nil
	}

	for i:= 0; i < lenNumStrs; i++ {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewNumStr(numStrs[i])

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
						" i='%v' NumStr='%v' Error='%v' ", i, numStrs[i], err.Error())
			}

			continue
		}

		b2Num, err := BigIntNum{}.NewNumStr(numStrs[i])

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStrs[i]). " +
					" i='%v' NumStr='%v' Error='%v' ", i, numStrs[i], err.Error())
		}


		result := bAdd.AddBigIntNums(finalResult.Result, b2Num)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// AddNumStrSeries - Adds a series of number strings and returns
// the combined total as an instance of Type, 'BigIntBasicMathResult'.
//
func (bAdd BigIntMathAdd) AddNumStrSeries(numStrs ... string) (BigIntBasicMathResult, error) {
	ePrefix := "BigIntMathAdd.AddNumStrSeries() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	for i, numStr := range numStrs {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewNumStr(numStr)

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStr). " +
						" i='%v' NumStr='%v' Error='%v' ", i, numStr, err.Error())
			}

			continue
		}

		b2Num, err := BigIntNum{}.NewNumStr(numStr)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStr). " +
					" i='%v' NumStr='%v' Error='%v' ", i, numStr, err.Error())
		}


		result := bAdd.AddBigIntNums(finalResult.Result, b2Num)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// AddNumStrDto - Receives two NumStrDto instances and adds their numeric values.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bAdd BigIntMathAdd) AddNumStrDto(n1Dto, n2Dto NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDto() "

	bPair, err := BigIntPair{}.NewNumStrDto(n1Dto, n2Dto)

	if err != nil {
		return BigIntBasicMathResult{},
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewNumStrDto(n1Dto, n2Dto). " +
				"Error='%v' ", err.Error())
	}

	return bAdd.AddPair(bPair), nil
}

// AddNumStrDtoArray - Adds an array of 'NumStrDto' types and returns the combined total
// as an instance of Type, 'BigIntBasicMathResult'.
//
func (bAdd BigIntMathAdd) AddNumStrDtoArray(nDtos []NumStrDto) (BigIntBasicMathResult, error) {
	ePrefix := "BigIntMathAdd.AddDecimalArray() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	lenDecs := len(nDtos)

	if lenDecs == 0 {
		return finalResult, nil
	}

	for i:= 0; i < lenDecs; i++ {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewNumStrDto(nDtos[i])

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(nDtos[i]). " +
						" i='%v' nDtos[i].GetNumStr()='%v' Error='%v' ", i, nDtos[i].GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, &nDtos[i])

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &nDtos[i]). " +
					" i='%v' dec[i].GetNumStr()='%v' Error='%v' ", i, nDtos[i].GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil

}

// AddDecimalSeries - Adds a series of 'NumStrDto' types and returns the combined total
// as an instance of Type, 'BigIntBasicMathResult'.
//
func (bAdd BigIntMathAdd) AddNumStrDtoSeries(nDtos ... NumStrDto) (BigIntBasicMathResult, error) {

	ePrefix := "BigIntMathAdd.AddNumStrDtoSeries() "

	finalResult := BigIntBasicMathResult{}.New()
	var err error

	for i, nDto := range nDtos {

		if i == 0 {

			finalResult.Result, err = BigIntNum{}.NewNumStrDto(nDto)

			if err != nil {
				return BigIntBasicMathResult{}.New(),
					fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(nDto). " +
						" i='%v' NumStr='%v' Error='%v' ", i, nDto.GetNumStr(), err.Error())
			}

			continue
		}

		bPair, err := BigIntPair{}.NewINumMgr(&finalResult.Result, &nDto)

		if err != nil {
			return BigIntBasicMathResult{}.New(),
				fmt.Errorf(ePrefix +
					"Error returned by BigIntPair{}.NewINumMgr(&finalResult.Result, &nDto). " +
					" i='%v' nDto.GetNumStr()='%v' Error='%v' ",
					i, nDto.GetNumStr(), err.Error())
		}

		result := bAdd.AddPair(bPair)

		finalResult.Result = result.Result.CopyOut()
	}

	return finalResult, nil
}

// AddPair - Receives a BigIntPair and proceeds to add b1.BigIntNum to
// b2.BigIntNum.
//
// The result is returned as type BigIntBasicMathResult.
//
func (bAdd BigIntMathAdd) AddPair(bPair BigIntPair) BigIntBasicMathResult {

	bPair.MakePrecisionsEqual()

	b3 := big.NewInt(0).Add(bPair.Big1.BigInt, bPair.Big2.BigInt)

	bResult := BigIntBasicMathResult{}
	bResult.Input= bPair.CopyOut()
	bResult.Result = BigIntNum{}.NewBigInt(b3, bPair.Big2.Precision)

	return bResult
}
