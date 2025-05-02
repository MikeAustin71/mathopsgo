package mathops

import (
	"fmt"
	"math/big"
)

type BigIntNumReadOnly struct {
	bigIntNum BigIntNum
}

// GetIntegerValue - Returns all the numeric digits in the
// base BigIntNum as a *big.Int integer value
func (birO *BigIntNumReadOnly) GetIntegerValue() *big.Int {

	return birO.bigIntNum.GetIntegerValue()
}

// GetPrecisionUint - Returns the precision of the underlying
// BigIntNum as a type uint.
func (birO *BigIntNumReadOnly) GetPrecisionUint() uint {

	return birO.bigIntNum.GetPrecisionUint()
}

// GetBigIntNum - Returns a deep copy of the underlying
// BigIntNum.
func (birO *BigIntNumReadOnly) GetBigIntNum() BigIntNum {

	return birO.bigIntNum.CopyOut()
}

// GetFixedDecimal - Returns a deep copy of the underlying BigIntNum
// numeric value as a type BigIntFixedDecimal
func (birO *BigIntNumReadOnly) GetFixedDecimal() BigIntFixedDecimal {

	return birO.bigIntNum.GetBigIntFixedDecimal()
}

// NewBigIntNum - Receives a BigIntNum parameter and returns a new BigIntNumReadOnly
// instance.
func (birO *BigIntNumReadOnly) NewBigIntNum(biNum BigIntNum) (BigIntNumReadOnly, error) {

	ePrefix := "BigIntNumReadOnly.NewBigIntNum()"

	var err error

	birO2 := BigIntNumReadOnly{}

	birO2.bigIntNum, err = new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNumReadOnly{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" birO2.bigIntNum, err = new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	birO2.bigIntNum.CopyIn(biNum)

	return birO2, nil
}

// NewFixedDecimal - Receives a BigIntFixedDecimal parameter and returns a new
// BigIntNumReadOnly instance.
func (birO *BigIntNumReadOnly) NewFixedDecimal(fixedDec BigIntFixedDecimal) (BigIntNumReadOnly, error) {

	ePrefix := "BigIntNumReadOnly.NewFixedDecimal()"

	var err error

	birO2 := BigIntNumReadOnly{}

	birO2.bigIntNum, err = new(BigIntNum).NewBigIntFixedDecimal(fixedDec)

	if err != nil {

		return BigIntNumReadOnly{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" birO2.bigIntNum, err = new(BigIntNum).NewBigIntFixedDecimal(fixedDec)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return birO2, nil
}

// NewNumStr - Receives a number string as input and returns
// a new BigIntNumReadOnly instance.
//
// This method assumes that the input parameter 'numStr' is a string
// of numeric digits which may be delimited by default USA numeric
// separators. Default USA numeric separators are defined as:
//
//	 	decimal separator = '.'
//	   thousands separator = ','
//			currency symbol = '$'
func (birO *BigIntNumReadOnly) NewNumStr(numStr string) (BigIntNumReadOnly, error) {

	ePrefix := "BigIntNumReadOnly.NewNumStr()"

	readOnly, err := new(BigIntNum).NewNumStr(numStr)

	if err != nil {

		return BigIntNumReadOnly{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" new(BigIntNum).NewNumStr(numStr)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	biRo := BigIntNumReadOnly{}

	biRo.bigIntNum, err = new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNumReadOnly{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" biRo.bigIntNum, err = new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	biRo.bigIntNum.CopyIn(readOnly)

	return biRo, nil
}
