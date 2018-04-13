package mathops

import (
	"math/big"
	"fmt"
	"strconv"
)

// BigIntNum - wraps a *big.Int integer and its associated
// precision and Sign Value.
//
// Note: The methods associated with this type all assume that
// the *big.Int are configured in base 10.
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

// NewBigIntExp - New BigInt Exponent returns a new
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
func (bNum BigIntNum) NewBigIntExp(bigI *big.Int, exponent int) BigIntNum {

	b := BigIntNum{}
	b.SetBigIntExponent(bigI, exponent)
	return b
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
			fmt.Errorf(ePrefix + "Error returned by decNum.GetSignedBigInt(). " +
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

// NewInt32Exp - New Int32 Exponent returns a new BigIntNum instance in
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
//	biNum := BigIntNum{}.NewInt32Exp(123456, -3) = "123.456" precision = 3 
//
//	biNum := BigIntNum{}.NewInt32Exp(123456, 3) = "123456000" precision = 0
//
func (bNum BigIntNum) NewInt32Exp(i32 int32, exponent int) BigIntNum {

	bigI := big.NewInt(int64(i32))
	b := BigIntNum{}
	b.SetBigIntExponent(bigI, exponent)
	return b
}

// NewInt64Exp - New Int64 Exponent returns a new BigIntNum instance in
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
//	biNum := BigIntNum{}.NewInt64Exp(123456, -3) = "123.456" precision = 3 
//
//	biNum := BigIntNum{}.NewInt64Exp(123456, 3) = "123456000" precision = 0
//
func (bNum BigIntNum) NewInt64Exp(i64 int64, exponent int) BigIntNum {

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

	bInt := ia.GetBigInt()
	precision := uint(ia.GetPrecision())

	b := BigIntNum{}
	b.SetBigInt(bInt, precision)
	return b, nil
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

	bigI, err := nDto.GetSignedBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by nDto.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}

	b := BigIntNum{}

	b.SetBigInt(bigI, nDto.GetPrecision())

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

	bigI, err := nDto.GetSignedBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by nDto.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}

	b := BigIntNum{}

	b.SetBigInt(bigI, nDto.GetPrecision())

	return b, nil
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

	bigI, err := nDto.GetSignedBigInt()

	if err != nil {
		fmt.Errorf(ePrefix + "Error returned by nDto.GetSignedBigInt(). " +
			"Error='%v'", err.Error())
	}

	bNum.SetBigInt(bigI, nDto.GetPrecision())

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

	bigI, err := nDto.GetSignedBigInt()

	if err != nil {
			fmt.Errorf(ePrefix + "Error returned by nDto.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}

	bNum.SetBigInt(bigI, nDto.GetPrecision())

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

	bigI, err := nDto.GetSignedBigInt()

	if err != nil {
			fmt.Errorf(ePrefix + "Error returned by nDto.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}

	bNum.SetBigInt(bigI, nDto.GetPrecision())

	return nil
} 


// GetNumStr - Converts the BigIntNum value to string of numbers
// which includes the decimal point and decimal digits if they exist.
//
func (bNum *BigIntNum) GetNumStr() (string, error) {

	ePrefix := "BigIntNum.GetNumStr() "

	nDto, err := NumStrDto{}.NewBigInt(bNum.BigInt, bNum.Precision)

	if err != nil {
		return "",
		fmt.Errorf (ePrefix +
			"Error returned by NumStrDto{}.NewBigInt(bNum.BigInt, bNum.Precision) " +
			"bNum.BigInt='%v' bNum.Precision='%v' Error='%v'",
				bNum.BigInt.Text(10), bNum.Precision, err.Error())
	}

	return nDto.GetNumStr(), nil
}

// BigIntPair - contains a pair of 'BitIntNum' types. This structure
// is used to set up calculations involving *big.Int types.
type BigIntPair struct {
	Big1 							BigIntNum
	Big1Compare 			int
	Big1AbsCompare 		int
	Precision1Compare int
	Big2							BigIntNum
}

// CopyIn - Copies the values provided by incoming BigIntPair
// parameter into the current BigIntPair instance.
func (bPair *BigIntPair) CopyIn(bd2 BigIntPair) {

	bPair.Big1 = bd2.Big1.CopyOut()
	bPair.Big2 = bd2.Big2.CopyOut()
	bPair.Big1Compare = bd2.Big1Compare
	bPair.Big1AbsCompare = bd2.Big1AbsCompare
	bPair.Precision1Compare = bd2.Precision1Compare

}

// CopyOut - Makes a deep copy of the current BigIntPair
// instance and returns it as a new BigIntPair object.
//
func (bPair *BigIntPair) CopyOut() BigIntPair {

	bd2 := BigIntPair{}.NewBigIntNum(bPair.Big1, bPair.Big2)

	return bd2
}

// Empty - Sets all data fields for the current BigIntPair instance
// to their uninitialized or zero states.
func (bPair *BigIntPair) Empty() {
	bPair.Big1.Empty()
	bPair.Big2.Empty()
	bPair.Big1Compare = 0
	bPair.Big1AbsCompare = 0
	bPair.Precision1Compare = 0

}

// MakePrecisionsEqual - Analyzes the two component BigIntNum's, b1 and b2,
// and then converts the one with the smallest precision to a new value
// equivalent in precision to the other BigIntNum. When completed, this
// method insures that both component BigIntNum's are both formatted to
// the largest precision.
//
func(bPair *BigIntPair) MakePrecisionsEqual() {

	if bPair.Big1.Precision == bPair.Big2.Precision {
		// Nothing to do. Precisions are equal.
		return
	}

	base10 := big.NewInt(10)

	if bPair.Big1.Precision > bPair.Big2.Precision {
		deltaPrecision := big.NewInt(int64(bPair.Big1.Precision - bPair.Big2.Precision))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		newB2Int := big.NewInt(0).Mul(bPair.Big2.BigInt, deltaPrecisionScale)
		newB2Num := BigIntNum{}.NewBigInt(newB2Int, bPair.Big1.Precision)
		newBPair := BigIntPair{}.NewBigIntNum(bPair.Big1, newB2Num)
		bPair.CopyIn(newBPair)
		return

	}

	// Must be bPair.Big2.Precision > bPair.Big1.Precision
	deltaPrecision := big.NewInt(int64(bPair.Big2.Precision - bPair.Big1.Precision))
	deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
	newB1Int := big.NewInt(0).Mul(bPair.Big1.BigInt, deltaPrecisionScale)
	newB1Num := BigIntNum{}.NewBigInt(newB1Int, bPair.Big2.Precision)
	newBPair := BigIntPair{}.NewBigIntNum(newB1Num, bPair.Big2)

	bPair.CopyIn(newBPair)

	return
}

// New - Creates an Empty BigIntPair instance. Both
// 'Big1' and 'Big2' are set to zero.  Both Precision
// values are also set to zero.
func (bPair BigIntPair) New() BigIntPair {
		base1Zero := big.NewInt(0)
		base2Zero := big.NewInt(0)
		b2Pair := BigIntPair{}.NewBase(base1Zero, 0, base2Zero, 0)
		return b2Pair
}

// NewBase - Creates a BigIntPair instance using two sets of
// *big.Int integers and their associated precision specifications.
func (bPair BigIntPair) NewBase(
						b1 *big.Int,
						b1Precision uint,
						b2 *big.Int,
						b2Precision uint) BigIntPair {


	b1BigIntNum := BigIntNum{}.NewBigInt(b1, b1Precision)
	b2BigIntNum := BigIntNum{}.NewBigInt(b2, b2Precision)

	return BigIntPair{}.NewBigIntNum(b1BigIntNum, b2BigIntNum)

}


// NewBigIntNum - Creates a new BigIntPair instance from input parameters
// consisting of two 'BigIntNum' types.
func (bPair BigIntPair) NewBigIntNum(b1, b2 BigIntNum ) BigIntPair {

	bd2 := BigIntPair{}

	bd2.SetBigIntPair(b1, b2)

	return bd2
}


// NewIntAry - Creates a new BigIntPair instance from two
// Decimal instances passed as input parameters.
//
func (bPair BigIntPair) NewDecimal(dec1, dec2 Decimal) (BigIntPair, error) {

	ePrefix := "BigIntPair.NewDecimal() "

	b1Num, err := BigIntNum{}.NewDecimal(dec1)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewDecimal(dec1). " +
				"Error='%v' ", err.Error())
	}

	b2Num, err := BigIntNum{}.NewDecimal(dec2)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewDecimal(dec2). " +
				"Error='%v' ", err.Error())
	}

	bd2 := BigIntPair{}

	bd2.SetBigIntPair(b1Num, b2Num)

	return bd2, nil
}

// NewIntAry - Creates a new BigIntPair instance from two
// IntAry instances passed as input parameters.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
func (bPair BigIntPair) NewIntAry(ia1, ia2 IntAry) (BigIntPair, error) {

	ePrefix := "BigIntPair.NewIntAry() "

	b1Num, err := BigIntNum{}.NewIntAry(ia1)

	if err != nil {
		return BigIntPair{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(ia1). " +
			"Error='%v' ", err.Error())
	}

	b2Num, err := BigIntNum{}.NewIntAry(ia2)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewIntAry(ia2). " +
				"Error='%v' ", err.Error())
	}

	bd2 := BigIntPair{}

	bd2.SetBigIntPair(b1Num, b2Num)

	return bd2, nil
}

// NewNumStr - Creates a new BigIntPair instance from two number strings
// passed as input parameters.
//
func (bPair BigIntPair) NewNumStr(n1NumStr, n2NumStr string) (BigIntPair, error) {
	ePrefix := "BigIntPair.NewNumStrDto() "
	b1Num, err := BigIntNum{}.NewNumStr(n1NumStr)

	if err != nil {
		return BigIntPair{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(n1NumStr). " +
			"numStr='%v' Error='%v' ", n1NumStr, err.Error())
	}

	b2Num, err := BigIntNum{}.NewNumStr(n2NumStr)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(n2NumStr). " +
				"numStr='%v' Error='%v' ", n2NumStr, err.Error())
	}

	b2Pair := BigIntPair{}

	b2Pair.SetBigIntPair(b1Num, b2Num)

	return b2Pair, nil

}

// NewNumStrDto - Creates a new BigIntPair instance from two NumStrDto
// instances passed as input parameters.
//
func (bPair BigIntPair) NewNumStrDto(n1Dto, n2Dto NumStrDto) (BigIntPair, error) {

	ePrefix := "BigIntPair.NewNumStrDto() "
	b1Num, err := BigIntNum{}.NewNumStrDto(n1Dto)

	if err != nil {
		return BigIntPair{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(n1Dto). " +
			"numStr='%v' Error='%v' ", n1Dto.GetNumStr(), err.Error())
	}

	b2Num, err := BigIntNum{}.NewNumStrDto(n2Dto)

	if err != nil {
		return BigIntPair{},
			fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStrDto(n2Dto). " +
				"numStr='%v' Error='%v' ", n2Dto.GetNumStr(), err.Error())
	}

	b2Pair := BigIntPair{}

	b2Pair.SetBigIntPair(b1Num, b2Num)

	return b2Pair, nil

}

// SetBigIntPairSetBigIntPair -Sets the values of the current
// BigIntPair instance to the input values of b1 and b2
// respectively.
//
func (bPair *BigIntPair) SetBigIntPair(b1, b2 BigIntNum ) {
	bPair.Big1.CopyIn(b1)
	bPair.Big2.CopyIn(b2)

	bPair.Big1Compare = bPair.Big1.BigInt.Cmp(bPair.Big2.BigInt)
	bPair.Big1AbsCompare = bPair.Big1.AbsBigInt.Cmp(bPair.Big2.AbsBigInt)

	if bPair.Big1.Precision == bPair.Big2.Precision {
		bPair.Precision1Compare = 0
	} else if bPair.Big1.Precision > bPair.Big2.Precision {
		bPair.Precision1Compare = 1
	} else {
		// Must be bPair.Big1.Precision < bPair.Big2.Precision
		bPair.Precision1Compare = -1
	}

}


// BigIntBasicMathResult - Used to return the result
// of an Addition, Subtraction or Multiplication operation.
//
type BigIntBasicMathResult struct {
	Input BigIntPair
	Result BigIntNum
}

// New - Returns a new BigIntBasicMathResult with all values set to
// zero.
//
func (basicResult BigIntBasicMathResult) New() BigIntBasicMathResult {
	b2 := BigIntBasicMathResult{}

	b2.Input = BigIntPair{}.NewBase(big.NewInt(0), 0, big.NewInt(0), 0)
	b2.Result = BigIntNum{}.NewBigInt(big.NewInt(0), 0)

	return b2
}

// BigIntDivideModResult - Used to return the result
// of Big Int division with a Quotient and Modulo
// Values.
//
type BigIntDivideModResult struct {
  Dividend BigIntNum
  Divisor BigIntNum
  Quotient BigIntNum
  Modulo BigIntNum
}

type BigIntDivideFracResult struct {
	Dividend BigIntNum
	Divisor BigIntNum
	Quotient *big.Rat
}
