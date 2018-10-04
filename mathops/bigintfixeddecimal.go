package mathops

import "math/big"

// BigIntFixedDecimal - A light data transfer structure used to represent
// a numeric value with a fixed integerNum of decimal digits.
//
type BigIntFixedDecimal struct {

	integerNum *big.Int // All of the numeric digits, both integer and fractional,
	// necessary to define a fixed length floating point number.
	// The number of digits to the right of the decimal place
	// is specified by the data field, BigIntFixedDecimal.precision.

	precision  uint			// Specifies the number of digits to the right of the decimal
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
func (bigIFd *BigIntFixedDecimal) CopyOut() (BigIntFixedDecimal) {

	return BigIntFixedDecimal{}.New(bigIFd.integerNum, bigIFd.precision)
}


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
func (bigIFd *BigIntFixedDecimal) GetInteger() (*big.Int) {

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
func (bigIFd *BigIntFixedDecimal) GetPrecision() (uint) {

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
// precision		uint	- Specifies the integerNum of digits to the right of the decimal point
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
// precision		uint	- Specifies the integerNum of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewInt(integer int, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(int64(integer)) , precision)

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
// precision		uint	- Specifies the integerNum of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewInt32(integer int32, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(int64(integer)) , precision)

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
// precision		uint	- Specifies the integerNum of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewInt64(integer int64, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(integer) , precision)

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
// precision		uint	- Specifies the integerNum of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewUInt(integer uint, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0).SetUint64(uint64(integer)) , precision)

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
// precision		uint	  - Specifies the integerNum of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewUInt32(integer uint32, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0).SetUint64(uint64(integer)) , precision)

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
// precision		uint	  - Specifies the integerNum of digits to the right of the decimal point
// 											in input parameter, 'integer'.
//
func (bigIFd BigIntFixedDecimal) NewUInt64(integer uint64, precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0).SetUint64(integer) , precision)

	return num
}



// NewZero - Creates and returns a new BigIntFixedDecimal type with a zero value. The
// number of digits to the right of the decimal place is specified by input parameter,
// precision.
//
// Input Parameters
// ================
//
// precision		uint	- Specifies the integerNum of digits to the right of the decimal point.
//
//
func (bigIFd BigIntFixedDecimal) NewZero(precision uint) BigIntFixedDecimal {

	num := BigIntFixedDecimal{}

	num.SetNumericValue(big.NewInt(0), precision)

	return num
}

// SetNumericValue - Sets the 'integerNum' or integer value for the current
// BigIntFixedDecimal instance.
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
// a numeric value with a fixed integerNum of fractional digits to the right of the
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
// instance. 'precision' specifies the integerNum of fractional digits to the right
// of the decimal place.
//
func (bigIFd *BigIntFixedDecimal) SetPrecisionValue(precision uint) {

	bigIFd.precision = precision

}

