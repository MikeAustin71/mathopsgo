package mathops

import (
	"fmt"
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

type BigIntNumReadOnly struct {
	bigIntNum BigIntNum
}

// GetIntegerValue - Returns all the numeric digits in the
// base BigIntNum as a *big.Int integer value
func (birO *BigIntNumReadOnly) GetIntegerValue() *big.Int {

	birOIntValue, err := birO.bigIntNum.GetIntegerValue()

	if err != nil {
		birOIntValue = big.NewInt(0)
	}

	return birOIntValue
}

// GetPrecisionUint - Returns the precision of the underlying
// BigIntNum as a type uint.
func (birO *BigIntNumReadOnly) GetPrecisionUint() uint {

	birOPrecisionUint, err := birO.bigIntNum.GetPrecisionUint()

	if err != nil {
		birOPrecisionUint = 0
	}

	return birOPrecisionUint
}

// GetBigIntNum - Returns a deep copy of the underlying
// BigIntNum.
func (birO *BigIntNumReadOnly) GetBigIntNum() (BigIntNum, error) {

	ePrefix := "BigIntNumReadOnly.GetBigIntNum()"

	bIntNum, err := birO.bigIntNum.CopyOut()

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" bIntNum, err := birO.bigIntNum.CopyOut()\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

	}

	return bIntNum, nil
}

// GetFixedDecimal - Returns a deep copy of the underlying BigIntNum
// numeric value as a type BigIntFixedDecimal
func (birO *BigIntNumReadOnly) GetFixedDecimal() BigIntFixedDecimal {

	bigIFixDec, err := birO.bigIntNum.GetBigIntFixedDecimal()

	if err != nil {
		bigIFixDec = new(BigIntFixedDecimal).NewInt(0, 0)
	}

	return bigIFixDec
}

// NewBigIntNum - Receives a BigIntNum parameter and returns a new BigIntNumReadOnly
// instance.
func (birO *BigIntNumReadOnly) NewBigIntNum(biNum BigIntNum) (BigIntNumReadOnly, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNumReadOnly.NewBigIntNum",
		"")

	if err != nil {
		return BigIntNumReadOnly{}, err
	}

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

	err = birO2.bigIntNum.CopyIn(&biNum)

	if err != nil {

		return BigIntNumReadOnly{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = birO2.bigIntNum.CopyIn(&biNum)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

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

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNumReadOnly.NewNumStr",
		"")

	if err != nil {

		return BigIntNumReadOnly{}, err
	}

	readOnly, err := new(BigIntNum).NewNumStr(numStr)

	if err != nil {

		return BigIntNumReadOnly{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "readOnly, err := new(BigIntNum).NewNumStr(numStr)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	biRo := BigIntNumReadOnly{}

	biRo.bigIntNum, err = new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNumReadOnly{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "biRo.bigIntNum, err = new(BigIntNum).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = biRo.bigIntNum.CopyIn(&readOnly)

	if err != nil {

		return BigIntNumReadOnly{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = biRo.bigIntNum.CopyIn(&readOnly)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return biRo, nil
}
