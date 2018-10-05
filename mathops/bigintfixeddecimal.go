package mathops

import "math/big"

// BigIntFixedDecimal - A light data transfer structure used to represent
// a numeric value with a fixed number of decimal digits.
//
type BigIntFixedDecimal struct {
	integerNum *big.Int // All of the numeric digits, both integer and fractional,
	// necessary to define a fixed length floating point number.
	// The number of digits to the right of the decimal place
	// is specified by the data field, BigIntFixedDecimal.precision.

	precision uint // Specifies the number of digits to the right of the decimal
	// place in the series of numeric digits represented by
	// BigIntFixedDecimal.precision.

	// Example: To represent the floating point number 52.459
	// a BigIntDecimal Structure would be configured as follows:
	// 			BigIntFixedDecimal.integerNum	= 52459
	// 			BigIntFixedDecimal.precision	= 3
}

// CopyIn - Receives a BigIntFixedDecimal type and copies the
// value to the current BigIntFixedDecimal instance.
//
func (bigIFd *BigIntFixedDecimal) CopyIn(fd BigIntFixedDecimal) {

	intVal := fd.GetInteger()

	if intVal == nil {
		intVal = big.NewInt(0)
	}

	bigIFd.integerNum = big.NewInt(0).Set(intVal)
	bigIFd.precision = fd.GetPrecision()

}

// CopyIn - Receives a pointer to a BigIntFixedDecimal type and
// copies the value to the current BigIntFixedDecimal instance.
//
func (bigIFd *BigIntFixedDecimal) CopyInPtr(fd *BigIntFixedDecimal) {

	intVal := fd.GetInteger()

	if intVal == nil {
		intVal = big.NewInt(0)
	}

	bigIFd.integerNum = big.NewInt(0).Set(intVal)
	bigIFd.precision = fd.GetPrecision()

}

// CopyOut - Returns a new BigIntFixedDecimal instance which is
// a deep copy of the current BigIntFixedDecimal instance.
//
func (bigIFd *BigIntFixedDecimal) CopyOut() BigIntFixedDecimal {

	return BigIntFixedDecimal{}.New(bigIFd.integerNum, bigIFd.precision)
}

// FormatNumStr - converts the numeric value of the current BigIntFixedDecimal
// instance to a number string. The returned number string will consist of a
// string of numeric digits. If the number contains fractional digits, the
// decimal separator period (.) will be used to separate integer and fractional
// digits within the string. There are no thousands separator present in the
// the returned string.
//
// The input parameter 'negValMode' is of type NegativeValueFmtMode.
// NegativeValueFmtMode encompasses a series of constants that are used
// to format negative values in a number string.
//
// Valid NegativeValueFormatMode's are defined as follows:
//
// LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//													 		a leading minus sign.
//															Example: -123456.78
//
//
// PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//														surrounding parentheses.
//														Example: (123456.78)
//
// ABSOLUTEPURENUMSTRFMTMODE - Formats a pure integerNum string with
//														 absolute (positive) integer value
//														 and no decimal point separator.
//														Example: (12345678)
//
func (bigIFd *BigIntFixedDecimal) FormatNumStr(negValMode NegativeValueFmtMode) string {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	decimalSeparator := '.'
	baseZero := big.NewInt(0)
	bigIFdSign := 1

	if bigIFd.integerNum.Cmp(baseZero) == -1 {
		bigIFdSign = -1
	}

	absBigInt := big.NewInt(0).Set(bigIFd.integerNum)

	if absBigInt.Cmp(baseZero) == -1 {
		absBigInt.Mul(absBigInt, big.NewInt(-1))
	}

	outRunes := make([]rune, 0, 300)

	scratchNum := big.NewInt(0).Set(absBigInt)

	if scratchNum.Cmp(baseZero) == 0 {
		bigIFdSign = 1

		outRunes = append(outRunes, '0')

		if bigIFd.precision > 0 {

			if negValMode != ABSOLUTEPURENUMSTRFMTMODE {
				outRunes = append(outRunes, decimalSeparator)
			}

			cnt := int(bigIFd.precision)

			if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
				cnt--
			}

			for h := 0; h < cnt; h++ {
				outRunes = append(outRunes, '0')
			}

		}

		return string(outRunes)
	}

	startIdx := 0
	modulo := big.NewInt(0)
	baseTen := big.NewInt(10)
	digitCnt := 0

	for scratchNum.Cmp(baseZero) == 1 {

		if startIdx == 0 &&
			bigIFdSign == -1 &&
			negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, ')')
		}

		modX := big.NewInt(0)
		scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, baseTen, modX)
		outRunes = append(outRunes, rune(modulo.Int64()+int64(48)))
		digitCnt++
		startIdx++

		if bigIFd.precision > 0 &&
			int(bigIFd.precision) == startIdx &&
			negValMode != ABSOLUTEPURENUMSTRFMTMODE {

			outRunes = append(outRunes, decimalSeparator)
			startIdx++
		}

	}

	if int(bigIFd.precision) >= digitCnt {

		delta := int(bigIFd.precision) - digitCnt + 1

		if negValMode == ABSOLUTEPURENUMSTRFMTMODE {
			delta--
		}

		for k := 0; k < delta; k++ {
			outRunes = append(outRunes, '0')
			startIdx++

			if bigIFd.precision > 0 &&
				int(bigIFd.precision) == startIdx &&
				negValMode != ABSOLUTEPURENUMSTRFMTMODE {

				outRunes = append(outRunes, decimalSeparator)
				startIdx++
			}
		}
	}

	startIdx--

	// adjust for negative sign value
	if bigIFdSign == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes = append(outRunes, '-')
			startIdx++

		} else if negValMode == PARENTHESESNEGVALFMTMODE {
			outRunes = append(outRunes, '(')
			startIdx += 2
		}

		/*
				MUST BE negValMode == ABSOLUTEPURENUMSTRFMTMODE
			  Do NOT Display Sign Character

		*/
	}

	sortLimit := startIdx / 2
	tRune := rune(0)
	yCnt := 0

	for i := startIdx; i > sortLimit; i-- {
		tRune = outRunes[yCnt]
		outRunes[yCnt] = outRunes[i]
		outRunes[i] = tRune
		yCnt++
	}

	return string(outRunes)
}

// GetInteger - Returns the 'integerNum' for the current
// BigIntFixedDecimal instance. The returned *big.Int type
// contains all the numeric digits which comprise the fixed
// decimal numerical value represented by this BigIntFixedDecimal
// instance.
//
func (bigIFd *BigIntFixedDecimal) GetInteger() *big.Int {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	return big.NewInt(0).Set(bigIFd.integerNum)
}

// GetNumericValue - Returns the 'integerNum' and 'precision' values for the
// current BigIntFixedDecimal instance.
func (bigIFd *BigIntFixedDecimal) GetNumericValue() (*big.Int, uint) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	return big.NewInt(0).Set(bigIFd.integerNum), bigIFd.precision
}

// GetNumStr - Converts the current BigIntFixedDecimal value to
// string of numbers which includes the decimal place and decimal
// digits if they exist.
//
func (bigIFd *BigIntFixedDecimal) GetNumStr() string {

	return bigIFd.FormatNumStr(LEADMINUSNEGVALFMTMODE)

}

// GetPrecision - Returns the 'precision' value for the current
// BigIntFixedDecimal instance. 'precision' specifies the number
// of digits to the right of the decimal place in the
// BigIntFixedDecimal.integerNum.
func (bigIFd *BigIntFixedDecimal) GetPrecision() uint {

	return bigIFd.precision
}

// IsValid - IsValid test the validity of the internal
// data field BigIntFixedDecimal.integerNum. If this
// data field is 'nil', BigIntFixedDecimal.integerNum
// is set to zero and the function returns 'false'.
//
// Otherwise, the function returns 'true'.
//
func (bigIFd *BigIntFixedDecimal) IsValid() bool {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		return false
	}

	return true
}

// New - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// Input Parameters
// ================
//
// integer	*big.Int	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) New(integer *big.Int, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(integer, precision)

	return num
}

// NewInt - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// Input Parameters
// ================
//
// integer	int				- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewInt(integer int, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(int64(integer)), precision)

	return num
}

// NewInt32 - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// Input Parameters
// ================
//
// integer		 int32	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewInt32(integer int32, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(int64(integer)), precision)

	return num
}

// NewInt64 - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// Input Parameters
// ================
//
// integer		 int32	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewInt64(integer int64, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(integer), precision)

	return num
}

// NewUInt - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// Input Parameters
// ================
//
// integer		  uint	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewUInt(integer uint, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0).SetUint64(uint64(integer)), precision)

	return num
}

// NewUInt32 - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// Input Parameters
// ================
//
// integer		  uint32	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	  - Specifies the number of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewUInt32(integer uint32, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0).SetUint64(uint64(integer)), precision)

	return num
}

// NewUInt64 - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// Input Parameters
// ================
//
// integer		  uint32	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	  - Specifies the number of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewUInt64(integer uint64, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0).SetUint64(integer), precision)

	return num
}

// NewZero - Creates and returns a new BigIntFixedDecimal type with a zero value. The
// number of digits to the right of the decimal place is specified by input parameter,
// precision.
//
// Input Parameters
// ================
//
// precision		uint	- Specifies the number of digits to the right of the decimal point.
//
//
func (bigIFd BigIntFixedDecimal) NewZero(precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0), precision)

	return num
}

// RoundToDecPlace - Rounds the numeric value of the current BigIntFixedDecimal
// instance to a specified number of decimal places.
//
// 'precision' equals the number of digits to the right of the decimal
// place.
//
// Example:
// 	integer= 123456; precision = 3; Numeric Value= 123.456
//
// If the value of BigIntFixedDecimal.bigInt is zero ('0'), that zero value will
// remain unaltered. However, the BigIntNum.precision value will be set equal to
// input parameter, 'precision'.
//
// If the number of decimal places specified for rounding ('precision") is
// equal to the current BigIntFixedDecimal.precision, no action is taken.
//
// If the number of decimal places specified for rounding ('precision') is
// greater than the current BigIntFixedDecimal.precision value, trailing
// zeros are added to the current BigIntFixedDecimal.bigInt value and BigIntNum.precision is set equal
// to input parameter, 'precision'.
//
// Finally, if the number of decimal places specified for rounding ('precision') is
// less than the current BigIntNum.precision value, the fractional digits will be
// rounded in accordance with the input parameter, 'precision'.
//
// Examples:
//
// 	 Original       				'precision'				Resulting
//    Value								input parameter			  Value
//  --------------				---------------     -------------
//	654.123456									9							 654.123456000
//	654.123456									4							 654.1235
// -654.123456									9							-654.123456000
// -654.123456									4							-654.1235
//		0													3								 0.000
//    0.000000									0								 0
//
//
func (bigIFd BigIntFixedDecimal) RoundToDecPlace(precision uint) {

	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), precision)
		return
	}

	if bigIFd.precision == precision {
		// Nothing to do. Specified 'precision' is already implemented.
		return
	}

	cmpToZeroResult := bigIFd.integerNum.Cmp(big.NewInt(0))

	// bigInt == zero, set precision an return
	if cmpToZeroResult == 0 {
		bigIFd.CopyIn(BigIntFixedDecimal{}.NewZero(precision))
		return
	}

	scale := big.NewInt(0)

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bigIFd.precision < precision {
		deltaPrecision := precision - bigIFd.precision
		scale =
			big.NewInt(0).Mul(
				big.NewInt(10),
				big.NewInt(int64(deltaPrecision)))
		bigIFd.integerNum.Exp(bigIFd.integerNum, scale, nil)
		bigIFd.precision += precision
		return
	}

	// Must be: bigIFd.precision >  precision

	bigNumRound5 :=
		BigIntFixedDecimal{}.NewInt(5, uint(precision+1))

	if cmpToZeroResult == -1 {
		bigNumRound5.integerNum.Mul(
			bigNumRound5.integerNum,
			big.NewInt(-1))
	}

	result := BigIntMathAdd{}.FixedDecimalAdd(bigIFd, bigNumRound5)

	// 10^deltaPrecision
	scale.Exp(big.NewInt(10),
		big.NewInt(int64(bigIFd.precision-precision)), nil)

	result.integerNum.Quo(result.integerNum, scale)

	bigIFd.SetNumericValue(result.integerNum, precision)
}

// SetNumericValue - Sets the BigIntFixedDecimal.integerNum or integer value
// for the current BigIntFixedDecimal instance.
//
func (bigIFd *BigIntFixedDecimal) SetIntegerValue(integer *big.Int) {

	if integer == nil {
		bigIFd.integerNum = big.NewInt(0)
	} else {
		bigIFd.integerNum = big.NewInt(0).Set(integer)
	}

}

// SetNumericValue - Sets the 'integerNum' and 'precision' values for the current
// BigIntFixedDecimal instance. Taken together, 'integerNum' and 'precision' describe
// a numeric value with a fixed number of fractional digits to the right of the
// decimal place.
//
func (bigIFd *BigIntFixedDecimal) SetNumericValue(integer *big.Int, precision uint) {

	if integer == nil {
		bigIFd.integerNum = big.NewInt(0)
	} else {
		bigIFd.integerNum = big.NewInt(0).Set(integer)
	}

	bigIFd.precision = precision

}

// SetPrecision - Sets the 'precision' value for the current BigIntFixedDecimal
// instance. 'precision' specifies the number of fractional digits to the right
// of the decimal place.
//
func (bigIFd *BigIntFixedDecimal) SetPrecisionValue(precision uint) {

	bigIFd.precision = precision

}

// TruncToDecPlace - Truncates the current BigIntFixedDecimal to the
// number of decimal places specified by input parameter 'precision'.
// No rounding occurs, the trailing digits are simply truncated or
// deleted in order to achieve the specified number of decimal places.
//
// 'precision' equals the number of digits to the right of the decimal
// place.
//
// If the value of BigIntFixedDecimal.integerNum is zero ('0'), that
// zero value will remain unaltered. However, BigIntFixedDecimal.precision
// will be set equal to input parameter, 'precision'.
//
// If the number of decimal places specified for truncation ('precision") is
// equal to the current BigIntFixedDecimal.precision, no action is taken and
// the original BigIntFixedDecimal numeric value remains unchanged.
//
// If the number of decimal places specified for truncation ('precision') is
// greater than the current BigIntFixedDecimal.precision, trailing zeros
// are added to the current BigIntFixedDecimal.integerNum value and
// BigIntFixedDecimal.precision is set equal to input parameter, 'precision'.
//
// If 'precision' is less than the current BigIntFixedDecimal.precision
// value, the current BigIntFixedDecimal numeric value is truncated to
// the specified 'precision' value and NO rounding occurs.
//
// Examples:
//
// 	 Original       			'newPrecision'				Resulting
//    Value								input parameter			  Value
//  --------------				---------------     -------------
//	654.123456									9							 654.123456000
//	654.123456									4							 654.1234 (no rounding)
// -654.123456									9							-654.123456000
// -654.123456									4							-654.1234 (no rounding)
//		0													3								 0.000
//    0.000000									0								 0
//
func (bigIFd *BigIntFixedDecimal) TruncToDecPlace(precision uint) {

	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), bigIFd.precision)
	}

	if bigIFd.precision == precision {
		// Nothing to do. Specified 'precision' is already implemented.
		return
	}

	// bigInt == zero, set precision an return
	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		bigIFd.precision = precision
		return
	}

	scale := big.NewInt(0)
	big10 := big.NewInt(10)
	delta := uint(0)

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bigIFd.precision < precision {
		delta = precision - bigIFd.precision
		scale.Exp(
			big10,
			big.NewInt(int64(delta)),
			nil)
		bigIFd.integerNum.Mul(bigIFd.integerNum, scale)
		bigIFd.precision += delta
		return
	}

	// Must be bigIFd.precision > precision
	delta = bigIFd.precision - precision
	scale.Exp(big10, big.NewInt(int64(delta)), nil)
	bigIFd.integerNum.Quo(bigIFd.integerNum, scale)
	bigIFd.precision = precision
	
}