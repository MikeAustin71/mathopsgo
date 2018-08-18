package mathops

import (
	"math/big"
	"fmt"
)

type FactorialDto struct {
	UpperLimit uint64
	LowerLimit uint64
}



type NFactorial struct {
	NFac 				uint64
	LowerLimit 	uint64
}

func (nFac NFactorial) GetFactorialArray(nFactorial int) []int{

	limit := nFactorial

	result := make([]int, limit)

	for i:= 0; i < limit; i++ {

		result[i] = nFactorial
		nFactorial--

	}

	return result
}

// CalcFactorialValueInt - Computes the value of nFactorial! and returns that value as a
// BigIntNum.
//
// Input Parameters:
// =================
//
// nFactorial	int		- The starting value in the factorial calculation. 'nFactorial' MUST BE
//                    a positive integer number.
//
//
// lowerLimit int		- The lower boundary for the factorial calculation. 'lowerLimit' MUST BE
//										a positive integer number.
//
// Examples:
// =========
//
//				1.	nFactorial = 7  and lowerLimit = 3
//
// 						The input parameter 'lowerLimit' specifies the lower boundary for the calculation.
// 						'nFactorial' = 7 and 'lowerLimit' = 3  will yield a calculation of:
//														7x6x5x4 = 840
//
//				2.	nFactorial = 7 and lowerLimit = 7
// 						Equivalent of 0! and 0! = 1
//
//        3.	nFactorial = 7 and lowerLimit = 1
//														7x6x5x4x3x2 = 5040
//
func (nFac NFactorial) CalcFactorialValueInt(nFactorial, lowerLimit int) (BigIntNum, error) {

	ePrefix := "NFactorial.CalcFactorialValueInt() "

	if nFactorial < 0 {
		return BigIntNum{}.NewZero(0),
		 fmt.Errorf(ePrefix +
		 	"Error: Input parameter 'nFactorial' is less than zero! " +
		 	"'nFactorial' must be a positive integer. 'nFactorial'='%v' ",
		 	nFactorial)
	}

	if lowerLimit < 0 {
		return BigIntNum{}.NewZero(0),
		 fmt.Errorf(ePrefix +
		 	"Error: Input parameter 'lowerLimit' is less than zero! " +
		 	"'lowerLimit' must be a positive integer. 'lowerLimit'='%v' ",
			 lowerLimit)
	}

	nFacBigInt := big.NewInt(int64(nFactorial))

	lowerLimitBigInt := big.NewInt(int64(lowerLimit))

	return nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt)
}

// CalcFactorialValueInt32 - Computes the value of nFactorial! and returns that value as a
// BigIntNum.
//
// Input Parameters:
// =================
//
// nFactorial	int32		- The starting value in the factorial calculation. 'nFactorial' MUST BE
//                    	a positive integer number.
//
//
// lowerLimit int32		- The lower boundary for the factorial calculation. 'lowerLimit' MUST BE
//											a positive integer number.
//
// Examples:
// =========
//
//				1.	nFactorial = 7  and lowerLimit = 3
//
// 						The input parameter 'lowerLimit' specifies the lower boundary for the calculation.
// 						'nFactorial' = 7 and 'lowerLimit' = 3  will yield a calculation of:
//														7x6x5x4 = 840
//
//				2.	nFactorial = 7 and lowerLimit = 7
// 						Equivalent of 0! and 0! = 1
//
//        3.	nFactorial = 7 and lowerLimit = 1
//														7x6x5x4x3x2 = 5040
//
func (nFac NFactorial) CalcFactorialValueInt32(nFactorial, lowerLimit int32) (BigIntNum, error) {

	ePrefix := "NFactorial.CalcFactorialValueInt32() "

	if nFactorial < 0 {
		return BigIntNum{}.NewZero(0),
		 fmt.Errorf(ePrefix +
		 	"Error: Input parameter 'nFactorial' is less than zero! " +
		 	"'nFactorial' must be a positive integer. 'nFactorial'='%v' ",
		 	nFactorial)
	}

	if lowerLimit < 0 {
		return BigIntNum{}.NewZero(0),
		 fmt.Errorf(ePrefix +
		 	"Error: Input parameter 'lowerLimit' is less than zero! " +
		 	"'lowerLimit' must be a positive integer. 'lowerLimit'='%v' ",
			 lowerLimit)
	}

	nFacBigInt := big.NewInt(int64(nFactorial))

	lowerLimitBigInt := big.NewInt(int64(lowerLimit))

	return nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt)
}

// CalcFactorialValueInt64 - Computes the value of nFactorial! and returns that value as a
// BigIntNum.
//
// Input Parameters:
// =================
//
// nFactorial	int64		- The starting value in the factorial calculation. 'nFactorial' MUST BE
//                    	a positive integer number.
//
//
// lowerLimit int64		- The lower boundary for the factorial calculation. 'lowerLimit' MUST BE
//											a positive integer number.
//
// Examples:
// =========
//
//				1.	nFactorial = 7  and lowerLimit = 3
//
// 						The input parameter 'lowerLimit' specifies the lower boundary for the calculation.
// 						'nFactorial' = 7 and 'lowerLimit' = 3  will yield a calculation of:
//														7x6x5x4 = 840
//
//				2.	nFactorial = 7 and lowerLimit = 7
// 						Equivalent of 0! and 0! = 1
//
//        3.	nFactorial = 7 and lowerLimit = 1
//														7x6x5x4x3x2 = 5040
//
func (nFac NFactorial) CalcFactorialValueInt64(nFactorial, lowerLimit int64) (BigIntNum, error) {

	ePrefix := "NFactorial.CalcFactorialValueInt64() "

	if nFactorial < 0 {
		return BigIntNum{}.NewZero(0),
		 fmt.Errorf(ePrefix +
		 	"Error: Input parameter 'nFactorial' is less than zero! " +
		 	"'nFactorial' must be a positive integer. 'nFactorial'='%v' ",
		 	nFactorial)
	}

	if lowerLimit < 0 {
		return BigIntNum{}.NewZero(0),
		 fmt.Errorf(ePrefix +
		 	"Error: Input parameter 'lowerLimit' is less than zero! " +
		 	"'lowerLimit' must be a positive integer. 'lowerLimit'='%v' ",
			 lowerLimit)
	}

	nFacBigInt := big.NewInt(nFactorial)

	lowerLimitBigInt := big.NewInt(lowerLimit)

	return nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt)
}

// CalcFactorialValueInt64 - Computes the value of nFactorial! and returns that value as a
// BigIntNum.
//
// Input Parameters:
// =================
//
// nFactorial	FactorialDto		- This structure contains a value of 'UpperLimit' or initial starting value
// 															of the the factorial calculation. In addition, the structure contains a
//															data field, 'LowerLimit', which specifies the lower boundary for the factorial
//															calculation.
//
// Examples:
// =========
//
//				1.	nFactorial.UpperLimit = 7  and nFactorial.LowerLimit = 3
//
// 						The input parameter 'LowerLimit' specifies the lower boundary for the calculation.
// 						'UpperLimit' = 7 and 'LowerLimit' = 3  will yield a calculation of:
//														7x6x5x4 = 840
//
//				2.	'UpperLimit' = 7 and 'LowerLimit' = 7
// 						Equivalent of 0! and 0! = 1
//
//        3.	nFactorial = 7 and lowerLimit = 1
//														7x6x5x4x3x2 = 5040
//
func (nFac NFactorial) CalcFactorialDtoValue(nFactorial FactorialDto) (BigIntNum, error) {

	ePrefix := "NFactorial.CalcFactorialDtoValue() "

	nFacBigInt := big.NewInt(0).SetUint64(nFactorial.UpperLimit)

	lowerLimitBigInt := big.NewInt(0).SetUint64(nFactorial.LowerLimit)

	result, err := nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt)

	if err != nil {
		return BigIntNum{}.NewZero(0),
		fmt.Errorf(ePrefix + "Error returned by nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt) " +
			"Error='%v'\n", err.Error())
	}

	return result, nil
}


func (nFac NFactorial) CalcFactorialValueBigIntNum(
																nFactorial, lowerLimit BigIntNum ) (BigIntNum, error) {

	return nFac.CalcFactorialValueBigInt(nFactorial.bigInt, lowerLimit.bigInt)
}

func (nFac NFactorial) CalcFactorialValueUint(nFactorial, lowerLimit uint) (BigIntNum, error) {

	nFacBigInt := big.NewInt(int64(nFactorial))

	lowerLimitBigInt := big.NewInt(int64(lowerLimit))

	return nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt)
}

func (nFac NFactorial) CalcFactorialValueUint32(nFactorial, lowerLimit uint32) (BigIntNum, error) {

	nFacBigInt := big.NewInt(int64(nFactorial))

	lowerLimitBigInt := big.NewInt(int64(lowerLimit))

	return nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt)
}


func (nFac NFactorial) CalcFactorialValueUint64(nFactorial, lowerLimit uint64) (BigIntNum, error) {

	nFacBigInt := big.NewInt(0).SetUint64(nFactorial)

	lowerLimitBigInt := big.NewInt(0).SetUint64(lowerLimit)

	return nFac.CalcFactorialValueBigInt(nFacBigInt, lowerLimitBigInt)
}


func (nFac NFactorial) CalcFactorialValueBigInt(nFactorial, lowerLimit *big.Int) (result BigIntNum, err error) {

	ePrefix := "NFactorial.CalcFactorialValueBigInt() "

	result = BigIntNum{}.NewZero(0)
	err = nil

	cmpResult := lowerLimit.Cmp(big.NewInt(0))

	// lower limit Must be at least '1'
	if  cmpResult == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: Lower Limit is less than '0'. lowerLimit='%v'", lowerLimit.Text(10))
		return result, err
	}

	if cmpResult == 0 {
		lowerLimit = big.NewInt(1)
	}


	cmpResult = nFactorial.Cmp(lowerLimit)

	if cmpResult == -1 {
		err = fmt.Errorf(ePrefix +
			"Error: 'nFactorial' is less than 'lowerLimit'. " +
			"nFactorial='%v' lowerLimit='%v'",nFactorial.Text(10), lowerLimit.Text(10))
		return result, err
	}

	// nFactorial = lowerLimit
	// This is equivalent to 0!
	// 0! = 1
	if cmpResult == 0 {
		err = nil
		result = BigIntNum{}.NewBigInt(big.NewInt(1), 0)
		return result, err
	}

	total := big.NewInt(0).Set(nFactorial)

	bigOne := big.NewInt(1)
	count := big.NewInt(0).Sub(total, bigOne)
	cmpResult = count.Cmp(lowerLimit)

	for cmpResult == 1 {
		total = big.NewInt(0).Mul(total, count)
		count = big.NewInt(0).Sub(count, bigOne)
		cmpResult = count.Cmp(lowerLimit)
	}

	result = BigIntNum{}.NewBigInt(total, 0)
	err = nil

	return result, err
}