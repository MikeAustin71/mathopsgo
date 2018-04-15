package mathops

import (
	"math/big"
	"fmt"
	"strconv"
)

// BigIntNum - wraps a *big.Int integer and its associated
// precision and Sign Value. While the numeric value is
// stored as an integer of type *big.Int, the BigIntNum
// type is capable of storing decimal fractions.
//
// All methods associated with this type all assume that
// the *big.Int value stored by the BigIntNum Type is configured
// in base 10.
//
// The BigIntNum Type implements the INumMgr interface.
//
type BigIntNum struct {
	BigInt			*big.Int
	AbsBigInt		*big.Int
	Precision 	uint			// Number of digits to the right of the decimal point.
	ScaleFactor *big.Int	// Scale Factor =  10^(precision * -1)
	Sign				int				// Valid values are -1 or +1. Indicates the sign of the
	// 	the 'BigInt' integer.
}

// CopyIn - Receives an incoming BigIntNum type and
// copies the value into the current BigIntNum instance.
//
func (bNum *BigIntNum) CopyIn(bigN BigIntNum) {

	bNum.BigInt = big.NewInt(0).Set(bigN.BigInt)
	bNum.AbsBigInt = big.NewInt(0).Set(bigN.AbsBigInt)
	bNum.Precision = bigN.Precision
	bNum.ScaleFactor = big.NewInt(0).Set(bigN.ScaleFactor)
	bNum.Sign = bigN.Sign
}

// CopyOut - Makes a deep copy of the current BigIntNum instance
// and returns it as a new BigIntNum instance.
//
func (bNum *BigIntNum) CopyOut() BigIntNum {

	b2 := BigIntNum{}.NewBigInt(big.NewInt(0).Set(bNum.BigInt), bNum.Precision)

	return b2
}

// Empty - Resets the BigIntNum data fields to their
// uninitialized or zero state.
//
func (bNum *BigIntNum) Empty() {

	bNum.BigInt = big.NewInt(0)
	bNum.AbsBigInt = big.NewInt(0)
	bNum.ScaleFactor = big.NewInt(1)
	bNum.Sign = 1
	bNum.Precision = 0

}

// Equal - Compares two BigIntNum instances and returns 'true'
// if the two instances are equal in value.
//
// If they are not Equal, the method returns 'false'.
//
func (bNum *BigIntNum) Equal(b2 BigIntNum) bool {

	if bNum.BigInt.Cmp(b2.BigInt) != 0 {
		return false
	}

	if bNum.AbsBigInt.Cmp(b2.AbsBigInt) != 0 {
		return false
	}

	if bNum.ScaleFactor.Cmp(b2.ScaleFactor) != 0 {
		return false
	}

	if bNum.Sign != b2.Sign {
		return false
	}

	if bNum.Precision != b2.Precision {
		return false
	}

	return true
}


// GetBigInt - return the numeric value as an integer
// of type *big.int
func (bNum *BigIntNum) GetBigInt() (*big.Int, error) {

	return bNum.BigInt, nil
}

// GetDecimal - Converts the current BigIntNum value to a Decimal
// instance. The resulting number value includes the decimal point
// and decimal digits if they exist.
//
func (bNum *BigIntNum) GetDecimal() (Decimal, error) {

	ePrefix := "BigIntNum.GetDecimal() "

	dec, err := Decimal{}.NewBigInt(bNum.BigInt, bNum.Precision)

	if err != nil {
		return Decimal{},
			fmt.Errorf (ePrefix +
				"Error returned by Decimal{}.NewBigInt(bNum.BigInt, bNum.Precision) " +
				"bNum.BigInt='%v' bNum.Precision='%v' Error='%v'",
				bNum.BigInt.Text(10), bNum.Precision, err.Error())
	}

	return dec, nil
}

// GetIntAry - Converts the current BigIntNum value to an IntAry
// instance. The resulting number value includes the decimal point
// and decimal digits if they exist.
//
func (bNum *BigIntNum) GetIntAry() (IntAry, error) {

	ePrefix := "BigIntNum.GetIntAry() "

	ia, err := IntAry{}.NewBigInt(bNum.BigInt, bNum.Precision)

	if err != nil {
		return IntAry{},
			fmt.Errorf (ePrefix +
				"Error returned by IntAry{}.NewBigInt(bNum.BigInt, bNum.Precision) " +
				"bNum.BigInt='%v' bNum.Precision='%v' Error='%v'",
				bNum.BigInt.Text(10), bNum.Precision, err.Error())
	}

	return ia, nil
}

// GetNumStr - Converts the current BigIntNum value to string of
// numbers which includes the decimal point and decimal digits
// if they exist.
//
func (bNum *BigIntNum) GetNumStr() (string) {

	nDto, err := NumStrDto{}.NewBigInt(bNum.BigInt, bNum.Precision)

	if err != nil {
		return ""
	}

	return nDto.GetNumStr()
}

// GetNumStrErr - Converts the current BigIntNum value to string of
// numbers which includes the decimal point and decimal digits
// if they exist. If an error occurs, the returned error object
// is not nil.
//
func (bNum *BigIntNum) GetNumStrErr() (string, error) {

	ePrefix := "BigIntNum.GetNumStrErr() "

	nDto, err := NumStrDto{}.NewBigInt(bNum.BigInt, bNum.Precision)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix +
				"Error returned by NumStrDto{}.NewBigInt(bNum.BigInt, bNum.Precision). " +
				"Error='%v' ", err.Error())

	}

	return nDto.GetNumStr(), nil
}

// GetNumStrDto - Converts the current BigIntNum value to a NumStrDto
// instance. The resulting number string includes the decimal point
// and decimal digits if they exist.
//
func (bNum *BigIntNum) GetNumStrDto() (NumStrDto, error) {

	ePrefix := "BigIntNum.GetNumStrDto() "

	nDto, err := NumStrDto{}.NewBigInt(bNum.BigInt, bNum.Precision)

	if err != nil {
		return NumStrDto{},
			fmt.Errorf (ePrefix +
				"Error returned by NumStrDto{}.NewBigInt(bNum.BigInt, bNum.Precision) " +
				"bNum.BigInt='%v' bNum.Precision='%v' Error='%v'",
				bNum.BigInt.Text(10), bNum.Precision, err.Error())
	}

	return nDto, nil
}

// GetPrecision - Returns precision as an integer of
// type 'int'.
//
// Precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
// 						1.234    	GetPrecision() = 3
// 								5			GetPrecision() = 0
// 					0.12345  		GetPrecision() = 5
//
//		Number String				Precision				Fractional Number
//			123456								3								123.456
//
//
func (bNum BigIntNum) GetPrecision() int {
	return int(bNum.Precision)
}

// GetPrecisionUint - Returns precision as an unsigned
// integer (uint).
//
// Precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
// 						1.234    	GetPrecision() = 3
// 								5			GetPrecision() = 0
// 					0.12345  		GetPrecision() = 5
//
//		Number String				Precision				Fractional Number
//			123456								3								123.456
//
//
func (bNum BigIntNum) GetPrecisionUint() uint {
	return bNum.Precision
}

// New - returns a new BigIntNum instance initialized to
// zero.
//
func (bNum BigIntNum) New() BigIntNum {
	b := BigIntNum{}
	b.Empty()
	return b
}

// NewBigInt - Creates a new BigIntNum instance using a *big.Int type and its
// associated precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal point. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal point by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
// Input Parameters
// bigI *big.Int	- 'bigI' is a type *big.Int and represents the integer
//									value of the number; that is, the numeric value with
//									out decimal digits.
//
// precision uint	- This unsigned integer (always a positive value) identifies
// 									the location of the decimal point in the integer value 'bigI'.
// 									The decimal point location is calculated by starting with the
// 									right most digit in the integer number and counting	left,
// 									'precision' places. Example:
//											Integer Value		Precision			Numeric Value
//											  123456					 3					  123.456
//
func (bNum BigIntNum) NewBigInt(bigI *big.Int, precision uint) BigIntNum {
	b := BigIntNum{}
	b.SetBigInt(bigI, precision)
	return b
}

// NewBigIntExponent - New BigInt Exponent returns a new
// BigIntNum instance in which the numeric value is
// set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
// 				numeric value = integer X 10^exponent
//
// If exponent is less than +1, precision is set equal to exponent and
// bigI is unchanged.
//
// If exponent is greater than 0, bigI is multiplied by 10 raised to the power
// of 'exponent' and precision is set equal to zero.
//
// Examples:
//
//	biNum :=
// 			BigIntNum{}.NewBigIntExponent(big.NewInt(int64(123456)), -3) = "123.456"  precision = 3
//
//	biNum :=
// 			BigIntNum{}.NewBigIntExponent(big.NewInt(int64(123456)), 3) = "123456000" precision = 0
//
func (bNum BigIntNum) NewBigIntExponent(bigI *big.Int, exponent int) BigIntNum {

	b := BigIntNum{}
	b.SetBigIntExponent(bigI, exponent)
	return b
}

// NewBigFloat - Returns a new BigIntNum instance using a *big.Float floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
// Input Parameters
// ================
//
// bigFloat *big.Float	- This *big.Float value will be converted into an instance of
//												BigIntNum.
//
// decimalPlaces int		- If greater than -1, this value designates the number
// 												of decimal places which will be extracted from the
//												*big.Float value. If 'decimalPlaces' = -1, the number
//												of decimal places will be inferred. -1 uses the smallest
// 												number of digits necessary return 'bigFloat' exactly.
//												Any 'decimalPlaces' value less than zero will be
//												automatically converted to -1. Be careful, -1 could
//												generate a very, very large number of decimal places.
//
func (bNum BigIntNum) NewBigFloat(bigFloat *big.Float, decimalPlaces int) (BigIntNum, error) {
	ePrefix := "BigIntNumNewFloat64() "

	b := BigIntNum{}
	err := b.SetBigFloat(bigFloat, decimalPlaces)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by b.SetBigFloat(bigFloat, decimalPlaces). " +
				"Error='%v' ", err.Error())
	}

	return b, nil
}

// NewDecimal - Receives a 'Decimal' type as input and returns a BigIntNum.
//
func (bNum BigIntNum) NewDecimal(decNum Decimal) (BigIntNum, error) {
	ePrefix := "BigIntNum.NewIntAry() "

	err := decNum.IsDecimalValid()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error: Input Parameter 'decNum' is INVALID!. Error returned by " +
				"decNum.IsDecimalValid(). Error='%v'", err.Error())
	}

	bInt, err := decNum.GetSignedBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by decNum.GetBigInt(). " +
				"Error='%v'", err.Error())
	}

	precision := uint(decNum.GetPrecision())

	b := BigIntNum{}
	b.SetBigInt(bInt, precision)
	return b, nil
}

// NewFloat32 - Returns a new BigIntNum instance using a float32 floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
// Input Parameters
// ================
//
// f32 float32				- This float32 value will be converted into an instance of
//											BigIntNum.
//
// decimalPlaces int	- If greater than -1, this value designates the number
// 											of decimal places which will be extracted from the
//											float32 value. If 'decimalPlaces' = -1, the number
//											of decimal places will be inferred. -1 uses the smallest
// 											number of digits necessary return 'f32' exactly.
//											Any 'decimalPlaces' value less than zero will be
//											automatically converted to -1.
//
func (bNum BigIntNum) NewFloat32(f32 float32, decimalPlaces int) (BigIntNum, error) {
	ePrefix := "BigIntNumNewFloat32() "

	b := BigIntNum{}
	err := b.SetFloat32(f32, decimalPlaces)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by b.SetFloat32(f32, decimalPlaces). " +
				"Error='%v' ", err.Error())
	}

	return b, nil
}

// NewFloat64 - Returns a new BigIntNum instance using a float64 floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
// Input Parameters
// ================
//
// f64 float64				- This float64 value will be converted into an instance of
//											BigIntNum.
//
// decimalPlaces int	- If greater than -1, this value designates the number
// 											of decimal places which will be extracted from the
//											float64 value. If 'decimalPlaces' = -1, the number
//											of decimal places will be inferred. -1 uses the smallest
// 											number of digits necessary return 'f64' exactly.
//											Any 'decimalPlaces' value less than zero will be
//											automatically converted to -1.
//
func (bNum BigIntNum) NewFloat64(f64 float64, decimalPlaces int) (BigIntNum, error) {
	ePrefix := "BigIntNumNewFloat64() "

	b := BigIntNum{}
	err := b.SetFloat64(f64, decimalPlaces)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by b.SetFloat64(f64, decimalPlaces). " +
				"Error='%v' ", err.Error())
	}

	return b, nil
}

// NewIntExponent - New Int Exponent returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
// 				numeric value = integer X 10^exponent
//
// If exponent is less than +1, precision is set equal to exponent and
// 'iNum' is unchanged.
//
// If exponent is greater than 0, 'iNum' is multiplied by 10 raised to the power
// of 'exponent' and precision is set equal to zero.
//
// Examples:
//
//	biNum := BigIntNum{}.NewIntExponent(123456, -3) = "123.456" precision = 3
//
//	biNum := BigIntNum{}.NewIntExponent(123456, 3) = "123456000" precision = 0
//
func (bNum BigIntNum) NewIntExponent(iNum int, exponent int) BigIntNum {

	bigI := big.NewInt(int64(iNum))
	b := BigIntNum{}
	b.SetBigIntExponent(bigI, exponent)
	return b
}

// NewInt32Exponent - New Int32 Exponent returns a new BigIntNum instance in
// which the numeric value is set using an integer multiplied by 10 raised
// to the power of the 'exponent' parameter.
//
// 				numeric value = integer X 10^exponent
//
// If exponent is less than +1, precision is set equal to exponent and
// 'i32' is unchanged.
//
// If exponent is greater than 0, 'i32' is multiplied by 10 raised to the power
// of 'exponent' and precision is set equal to zero.
//
// Examples:
//
//	biNum := BigIntNum{}.NewInt32Exponent(123456, -3) = "123.456" precision = 3
//
//	biNum := BigIntNum{}.NewInt32Exponent(123456, 3) = "123456000" precision = 0
//
func (bNum BigIntNum) NewInt32Exponent(i32 int32, exponent int) BigIntNum {

	bigI := big.NewInt(int64(i32))
	b := BigIntNum{}
	b.SetBigIntExponent(bigI, exponent)
	return b
}

// NewInt64Exponent - New Int64 Exponent returns a new BigIntNum instance in
// which the numeric value is set using an integer multiplied by 10 raised
// to the power of the 'exponent' parameter.
//
// 				numeric value = integer X 10^exponent
//
// If exponent is less than +1, precision is set equal to exponent and
// 'i64' is unchanged.
//
// If exponent is greater than 0, 'i64' is multiplied by 10 raised to the power
// of 'exponent' and precision is set equal to zero.
//
// Examples:
//
//	biNum := BigIntNum{}.NewInt64Exponent(123456, -3) = "123.456" precision = 3
//
//	biNum := BigIntNum{}.NewInt64Exponent(123456, 3) = "123456000" precision = 0
//
func (bNum BigIntNum) NewInt64Exponent(i64 int64, exponent int) BigIntNum {

	bigI := big.NewInt(i64)
	b := BigIntNum{}
	b.SetBigIntExponent(bigI, exponent)
	return b
}


// NewIntAry - Creates a new BigIntNum instance from an input parameter
// IntAry.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
func (bNum BigIntNum) NewIntAry(ia IntAry) (BigIntNum, error) {
	ePrefix := "BigIntNum.NewIntAry() "

	err := ia.IsIntAryValid("")

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error: Input Parameter 'ia' is INVALID!. Error returned by " +
				"ia.IsIntAryValid(\"\"). Error='%v'", err.Error())
	}

	bInt,_ := ia.GetBigInt()

	precision := ia.GetPrecisionUint()

	b := BigIntNum{}
	b.SetBigInt(bInt, precision)
	return b, nil
}

// NewINumMgr - Receives an object which implements the INumMgr interface.
// The method then proceeds to create a new BigIntNum instance equivalent
// in numeric value to the input parameter, 'numMgr'. The BigIntNum instance
// is then returned to the calling function.
//
// Currently, the following 'mathops' Types implement the INumMgr interface:
// 					Decimal,
//					IntAry,
//					NumStrDto,
//					BigIntNum
//
// Note: 'numMgr' must be a pointer to a type. This method will not accept
// 'numMgr' as a value. The pointer to the type is needed in or order to
// call methods on 'numMgr'.
//
// Example:
//	dec, err := Decimal{}.NewNumStr(nStr)
//	bINum, err := BigIntNum{}.NewINumMgr(&dec)
//
func (bNum BigIntNum) NewINumMgr(numMgr INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewINumMgr() "

	bINum := BigIntNum{}.New()

	err := bINum.SetINumMgr(numMgr)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by bINum.SetINumMgr(numMgr). " +
				"Error='%v' ", err.Error())
	}

	return bINum, nil
}

// NewNumStr - Receives a number string as input and returns
// a new BigIntNum instance.
//
func (bNum BigIntNum) NewNumStr(numStr string) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewNumStr() "

	nDto, err := NumStrDto{}.NewNumStr(numStr)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewNumStr(numStr). " +
				"numStr='%v' Error='%v'", numStr, err.Error())
	}

	bigI, err := nDto.GetBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by nDto.GetBigInt(). " +
				"Error='%v'", err.Error())
	}

	b := BigIntNum{}

	b.SetBigInt(bigI, uint(nDto.GetPrecision()))

	return b, nil
}

// NewNumStrDto - Receives a NumStrDto instance as input and returns
// a new BigIntNum instance.
//
func (bNum BigIntNum) NewNumStrDto(nDto NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewNumStrDto() "

	err := nDto.IsNumStrDtoValid("")

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned from nDto.IsNumStrDtoValid(\"\"). " +
				"NumStr='%v' Error='%v'", nDto.GetNumStr(), err.Error())
	}

	bigI, err := nDto.GetBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by nDto.GetBigInt(). " +
				"Error='%v'", err.Error())
	}

	b := BigIntNum{}

	b.SetBigInt(bigI, uint(nDto.GetPrecision()))

	return b, nil
}


// RoundToDecPlace - Rounds the current BigIntNum instance to a specified
// number of decimal places.
//
// If the number of decimal places specified for round is greater than
// or equal to the current number of decimal places, no action is taken.
//
// 'precision' must be less than the current BigIntNum.Precision before
// the rounding operation will engage.
//
func (bNum *BigIntNum) RoundToDecPlace(precision uint) {

	if bNum.Precision <= precision {
		// Nothing to do. We can only round to a precision
		// which is less than the current precision.
		return
	}

	if bNum.BigInt.Cmp(big.NewInt(0)) == 0 {
		bNum.Precision = precision
		return
	}

	base5 := big.NewInt(5)

	bigNumRound5 := BigIntNum{}.NewBigInt(base5, uint(precision + 1))

	baseIRound := big.NewInt(0).Set(bNum.AbsBigInt)

	bigNumBase := BigIntNum{}.NewBigInt(baseIRound, bNum.Precision)

	result := BigIntMathAdd{}.AddBigIntNums(bigNumBase, bigNumRound5)

	base10 := big.NewInt(10)

	newBigInt := big.NewInt(0).Div(result.Result.BigInt, base10)

	if bNum.Sign < 0 {
		newBigInt = big.NewInt(0).Neg(newBigInt)
	}

	bNum.SetBigInt(newBigInt, precision)

}

// SetBigInt - Sets the value of the current BigIntNum instance using
// the input parameters *big.Int integer and precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal point. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal point by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
// Input Parameters
// bigI *big.Int	- 'bigI' is a type *big.Int and represents the integer
//									value of the number; that is, the numeric value with
//									out decimal digits.
//
// precision uint	- This unsigned integer (always a positive value) identifies
// 									the location of the decimal point in the integer value 'bigI'.
// 									The decimal point location is calculated by starting with the
// 									right most digit in the integer number and counting	left,
// 									'precision' places. Example:
//											Integer Value		Precision			Numeric Value
//											  123456					 3					  123.456
//
func (bNum *BigIntNum) SetBigInt(bigI *big.Int, precision uint) {

	bNum.BigInt = big.NewInt(0).Set(bigI)
	bNum.Precision = precision
	base10 := big.NewInt(0).SetInt64(int64(10))
	bigPrecision := big.NewInt(0).SetInt64(int64(bNum.Precision))
	bNum.ScaleFactor = big.NewInt(0).Exp(base10, bigPrecision, nil)
	baseZero := big.NewInt(0).SetInt64(0)
	result := bNum.BigInt.Cmp(baseZero)

	if result == -1 {
		bNum.Sign = -1
		minusOne := big.NewInt(0).SetInt64(-1)
		bNum.AbsBigInt = big.NewInt(0).Mul(bNum.BigInt, minusOne)
	} else {
		bNum.Sign = 1
		bNum.AbsBigInt = big.NewInt(0).Set(bNum.BigInt)
	}

}

// SetBigIntExponent - Sets the numeric value using an integer
// multiplied by 10 raised to the power of the 'exponent'
// parameter.
//
// 				numeric value = integer X 10^exponent
//
// If exponent is less than +1, precision is set equal to exponent and
// bigI is unchanged.
//
// If exponent is greater than 0, bigI is multiplied by 10 raised to the
// power of exponent and precision is set equal to zero.
//
func (bNum *BigIntNum) SetBigIntExponent(bigI *big.Int, exponent int) {

	if exponent < 1 {
		precision := uint(exponent * -1)
		bNum.SetBigInt(bigI, precision)
		return
	}

	// exponent must be greater than zero.
	// scale left exponent places and set precision to zero

	big10 := big.NewInt(10)
	scale := big.NewInt(int64(exponent))
	scaleValue := big.NewInt(0).Exp(big10, scale, nil)
	newBigI := big.NewInt(0).Mul(bigI, scaleValue)

	bNum.SetBigInt(newBigI, uint(0))
	return
}

// SetBigFloat - Sets the value of a BigIntNum using a *big.Float floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
// Input Parameters
// ================
//
// bigFloat *big.Float	- This float32 value will be converted into an instance of
//												BigIntNum.
//
// decimalPlaces int		- If greater than -1, this value designates the number
// 												of decimal places which will be extracted from the
//												*big.Float value. If 'decimalPlaces' = -1, the number
//												of decimal places will be inferred automatically.
// 												-1 uses the smallest number of digits necessary return
// 												the *big.Float value exactly.	Any 'decimalPlaces' value
// 												less than zero will be automatically converted to -1. Be
//												careful, -1 could generate a very, very large number of
//												decimal places.
//
func (bNum *BigIntNum) SetBigFloat(bigFloat *big.Float, decimalPlaces int) error {

	ePrefix := "NumStrDto.NewBigFloat() "

	if decimalPlaces < 0 {
		decimalPlaces = -1
	}

	numStr := bigFloat.Text('f', decimalPlaces)

	nDto, err := NumStrDto{}.NewNumStr(numStr)

	if err != nil {
		return 	fmt.Errorf(ePrefix	+
			"Error returned by NumStrDto{}.NewNumStr(numStr). " +
			"numStr='%v'  Error='%v'",
			numStr, err.Error())
	}

	bigI, err := nDto.GetBigInt()

	if err != nil {
		fmt.Errorf(ePrefix + "Error returned by nDto.GetBigInt(). " +
			"Error='%v'", err.Error())
	}

	bNum.SetBigInt(bigI, uint(nDto.GetPrecision()))

	return nil
}

// SetFloat32 - Sets the value of a BigIntNum using a float32 floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
// Input Parameters
// ================
//
// f32 float32				- This float32 value will be converted into an instance of
//											BigIntNum.
//
// decimalPlaces int	- If greater than -1, this value designates the number
// 											of decimal places which will be extracted from the
//											float32 value. If 'decimalPlaces' = -1, the number
//											of decimal places will be inferred. -1 uses the smallest
// 											number of digits necessary return 'f32' exactly.
//											Any 'decimalPlaces' value less than zero will be
//											automatically converted to -1.
//
func (bNum *BigIntNum) SetFloat32(f32 float32, decimalPlaces int) error {

	ePrefix := "BigIntNum.SetFloat32() "

	if decimalPlaces < 0 {
		decimalPlaces = -1
	}

	numStr := strconv.FormatFloat(float64(f32), 'f', decimalPlaces, 32)

	nDto, err := NumStrDto{}.NewNumStr(numStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'. ",
			numStr, err.Error())
	}

	bigI, err := nDto.GetBigInt()

	if err != nil {
		fmt.Errorf(ePrefix + "Error returned by nDto.GetBigInt(). " +
			"Error='%v'", err.Error())
	}

	bNum.SetBigInt(bigI, uint(nDto.GetPrecision()))

	return nil
}


// SetFloat64 - Sets the value of a BigIntNum using a float64 floating point
// input parameter.  The precision of the number is specified by the input
// parameter 'decimalPlaces'.
//
// Input Parameters
// ================
//
// f64 float64				- This float64 value will be converted into an instance of
//											BigIntNum.
//
// decimalPlaces int	- If greater than -1, this value designates the number
// 											of decimal places which will be extracted from the
//											float64 value. If 'decimalPlaces' = -1, the number
//											of decimal places will be inferred. -1 uses the smallest
// 											number of digits necessary return 'f64' exactly.
//											Any 'decimalPlaces' value less than zero will be
//											automatically converted to -1.
//
func (bNum *BigIntNum) SetFloat64(f64 float64, decimalPlaces int) error {

	ePrefix := "BigIntNum.SetFloat64() "

	if decimalPlaces < 0 {
		decimalPlaces = -1
	}

	numStr := strconv.FormatFloat(f64, 'f', decimalPlaces, 64)

	nDto, err := NumStrDto{}.NewNumStr(numStr)

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v'. ",
			numStr, err.Error())
	}

	bigI, err := nDto.GetBigInt()

	if err != nil {
		fmt.Errorf(ePrefix + "Error returned by nDto.GetBigInt(). " +
			"Error='%v'", err.Error())
	}

	bNum.SetBigInt(bigI, uint(nDto.GetPrecision()))

	return nil
}

// SetINumMgr - Receives an input parameter implementing
// the INumMgr interface and proceeds to set the current
// BigIntNum instance to its equivalent numeric value.
//
// Currently, the following 'mathops' Types implement the INumMgr interface:
// 					Decimal,
//					IntAry,
//					NumStrDto,
//					BigIntNum
//
// 'numMgr' must be a pointer to a type. This method will not accept
// 'numMgr' as a value. The pointer to the type is needed in or order to
// call methods on 'numMgr'.
//
// Example:
//
//	bINum := BigIntNum{}
//	dec, err := Decimal{}.NewNumStr(nStr)
//	err := bINum.SetINumMgr(&dec)
//
func (bNum *BigIntNum) SetINumMgr(numMgr INumMgr) error {

	ePrefix := "BigIntNum.SetINumMgr() "

	bigInt, err := numMgr.GetBigInt()

	if err != nil {
		return fmt.Errorf(ePrefix + "Error returned by numMgr.GetBigInt(). " +
			"Error='%v'", err.Error() )
	}

	bNum.SetBigInt(bigInt, numMgr.GetPrecisionUint())

	return nil
}

// TruncToDecPlace - Truncates the current BigIntNum to the number
// of decimal places specified by input parameter 'precision'.
// No rounding occurs, the trailing digits are simply truncated or
// deleted in order to achieve the specified number of decimal places.
//
// 'precision' equals the number of digits to the right of the decimal
// place.
//
func (bNum *BigIntNum) TruncToDecPlace(precision uint) {

	if bNum.Precision <= precision {
		// Nothing to do. We can only round to a precision
		// which is less than the current precision.
		return
	}

	if bNum.BigInt.Cmp(big.NewInt(0)) == 0 {
		bNum.Precision = precision
		return
	}

	base10 := big.NewInt(10)

	newBigInt := big.NewInt(0).Div(bNum.AbsBigInt, base10)

	if bNum.Sign < 1 {
		newBigInt = big.NewInt(0).Neg(newBigInt)
	}

	bNum.SetBigInt(newBigInt, precision)
}

