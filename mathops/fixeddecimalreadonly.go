package mathops

import (
	"fmt"
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

// FixedDecimalReadOnly - encapsulates a Read Only
// BigIntFixedDecimal objects. This type is designed
// to be used with constant numeric values that should
// not be changed.
//
// A series of related functions provides the means to
// read the numeric value encapsulated by the FixedDecimalReadOnly
// type.
type FixedDecimalReadOnly struct {
	fixedDecimal BigIntFixedDecimal
}

// GetBigIntNum - returns the numeric value of the underlying BigIntFixedDecimal
// as a type BigIntNum.
func (fDecRO *FixedDecimalReadOnly) GetBigIntNum() BigIntNum {

	bigIntNum, _ := fDecRO.fixedDecimal.GetBigIntNum()

	return bigIntNum
}

// GetIntAry - Returns the numeric value of the underlying BigIntFixedDecimal
// as a type IntAry.
func (fDecRO *FixedDecimalReadOnly) GetIntAry() (IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"FixedDecimalReadOnly.GetIntAry",
		"")

	if err != nil {
		return IntAry{}, err
	}

	err = fDecRO.fixedDecimal.IsValid(ePrefix.XCpy("Validating fDecRO.fixedDecimal").String())

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = fDecRO.fixedDecimal.IsValid(ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	intAry, err := fDecRO.fixedDecimal.GetIntAry()

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "intAry, err := fDecRO.fixedDecimal.GetIntAry()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return intAry, nil
}

// GetFixedDecimal - Returns a deep copy of the underlying
// BigIntFixedDecimal value.
func (fDecRO *FixedDecimalReadOnly) GetFixedDecimal() BigIntFixedDecimal {

	fDecRoCopy, _ := fDecRO.fixedDecimal.CopyOut()

	return fDecRoCopy
}

// GetInteger - Returns the *big.Int integer value from the
// underlying BigIntFixedDecimal
func (fDecRO *FixedDecimalReadOnly) GetInteger() *big.Int {

	bigInt, _ := fDecRO.fixedDecimal.GetIntegerValue()

	return bigInt
}

// GetNumStr - Converts the underlying BigIntFixedDecimal
// value to a string of numbers which includes the decimal
// place and decimal digits if they exist. Note that
// the period ('.') is the decimal separator character
// always used in the returned number string to separate
// integer and fractional digits.
func (fDecRO *FixedDecimalReadOnly) GetNumStr() string {

	numStr, _ := fDecRO.fixedDecimal.GetNumStr()

	return numStr
}

// GetPrecisionUint - returns an unsigned integer specifying
// the number of digits to the right of the decimal place.
func (fDecRO *FixedDecimalReadOnly) GetPrecisionUint() uint {

	precisionUint, _ := fDecRO.fixedDecimal.GetPrecisionUint()

	return precisionUint
}

// GetPrecisionBigInt - returns an unsigned integer specifying
// the number of digits to the right of the decimal place.
func (fDecRO *FixedDecimalReadOnly) GetPrecisionBigInt() *big.Int {

	precisionBigInt, _ := fDecRO.fixedDecimal.GetPrecisionBigInt()

	return precisionBigInt
}

func (fDecRO *FixedDecimalReadOnly) GetBigIntPrecision() (integer, precision *big.Int) {

	integer = fDecRO.GetInteger()

	if integer == nil {
		integer = big.NewInt(0)
	}

	precision = fDecRO.GetPrecisionBigInt()

	if precision == nil {
		precision = big.NewInt(0)
	}

	return integer, precision
}

// IsValid - Performs diagnostic and remedial actions on
// the underlying BigIntFixedDecimal instance. Returns
// 'false' if the instance is uninitialized.
func (fDecRO *FixedDecimalReadOnly) IsValid() bool {

	err := fDecRO.fixedDecimal.IsValid("fDecRO.fixedDecimal")

	if err != nil {
		return false
	}

	return true
}

func (fDecRO *FixedDecimalReadOnly) IsZero() bool {

	isZero, err := fDecRO.fixedDecimal.IsZero()

	if err != nil {
		return false
	}

	return isZero
}

// NewNumStr - Creates and returns a new FixedDecimalReadOnly
// instance initialized to the value of the number string
// input parameter, 'numStr'.
//
// A number string is a string of numeric digits which may,
// or may not, be prefixed with a minus sign ('-') indicating
// a negative number. If the numeric string of digits is prefixed
// by a left parenthesis ('(') and suffixed by a corresponding
// right parenthesis (')'), this also indicates a negative value.
//
// The numeric string of digits may also contain a period
// ('.') which is treated as a decimal separator and used to
// separate integer and fractional digits within the number
// string.
//
// The only decimal separator recognized by this method is the
// period ('.').
func (fDecRO *FixedDecimalReadOnly) NewNumStr(numStr string) (FixedDecimalReadOnly, error) {
	ePrefix := "FixedDecimalReadOnly.NewNumStr() "

	fo2 := FixedDecimalReadOnly{}

	fo2.fixedDecimal = new(BigIntFixedDecimal).NewZero(0)

	err := fo2.fixedDecimal.SetNumStr(numStr, '.')

	if err != nil {

		return fo2,
			fmt.Errorf(ePrefix)
	}

	return fo2, nil
}

// NewFixedDecimal - Receives a BigIntFixedDecimal instance
// as an input parameter and returns a new FixedDecimalReadOnly
// object.
func (fDecRO *FixedDecimalReadOnly) NewFixedDecimal(
	fixedDecimal BigIntFixedDecimal) (FixedDecimalReadOnly, error) {

	ePrefix := "FixedDecimalReadOnly.NewFixedDecimal()"

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewZero(0)

	var err error

	errX := fixedDecimal.IsValid(ePrefix)

	if errX != nil {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input Parameter 'fixedDecimal' is invalid\n"+
			"Error: %v\n", errX.Error(),
			ePrefix)

		return f2, err
	}

	err = f2.fixedDecimal.CopyIn(fixedDecimal)

	if err != nil {
		return f2,
			fmt.Errorf("%v\n"+
				"Error returned by call to f2.fixedDecimal.CopyIn(fixedDecimal).\n"+
				"Error='%v'\n",
				ePrefix, err.Error())
	}

	return f2, err
}

// NewInt - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'intValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue'.
func (fDecRO *FixedDecimalReadOnly) NewInt(
	intValue int,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewInt(intValue, precision)

	return f2

}

// NewInt32 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'int32Value' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue'.
func (fDecRO *FixedDecimalReadOnly) NewInt32(
	int32Value int32,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewInt32(int32Value, precision)

	return f2

}

// NewInt64 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'int64Value' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue'.
func (fDecRO *FixedDecimalReadOnly) NewInt64(
	int64Value int64,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewInt64(int64Value, precision)

	return f2

}

// NewUInt - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'uintValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue'.
func (fDecRO *FixedDecimalReadOnly) NewUInt(
	uintValue,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewUInt(uintValue, precision)

	return f2

}

// NewUInt32 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'uintValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue'.
func (fDecRO *FixedDecimalReadOnly) NewUInt32(
	uint32Value uint32,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewUInt32(uint32Value, precision)

	return f2

}

// NewUInt64 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'uintValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue'.
func (fDecRO *FixedDecimalReadOnly) NewUInt64(
	uint64Value uint64,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewUInt64(uint64Value, precision)

	return f2
}

// NewZero - Creates and returns a new FixedDecimalReadOnly instance initialized
// to zero. The input parameter 'precision' specifies the number of zero digits
// to the right of the decimal place.
func (fDecRO *FixedDecimalReadOnly) NewZero() FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = new(BigIntFixedDecimal).NewZero(0)

	return f2

}
