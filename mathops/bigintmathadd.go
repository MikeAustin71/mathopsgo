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

// AddBigIntNumSeries - Adds a series of BigIntNum types and returns the total in a
// BigIntBasicMathResult instance.
//
func (bAdd BigIntMathAdd) AddBigIntNumSeries(bNums ... BigIntNum) BigIntBasicMathResult {

	total := BigIntNum{}.New()

	for i, bNum := range bNums {

		if i == 0 {
			total = bNum.CopyOut()
			continue
		}

		result := bAdd.AddBigIntNums(total, bNum)

		total = result.Result.CopyOut()

	}

	finalResult := BigIntBasicMathResult{}.New()
	finalResult.Result = total.CopyOut()

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
			fmt.Errorf(ePrefix + "Error returned by BigIntPair{}.NewDecimal(dec1, dec2). " +
				"Error='%v' ", err.Error())
	}

	return bAdd.AddPair(bPair), nil
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
