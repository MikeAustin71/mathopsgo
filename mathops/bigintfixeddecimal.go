package mathops

import (
	"errors"
	"fmt"
	"math/big"
)



// FixedDecimalReadOnly - encapsulates a Read Only
// BigIntFixedDecimal objects. This type is designed
// to be used with constant numeric values that should
// not be changed.
//
// A series of related functions provides the means to
// read the numeric value encapsulated by the FixedDecimalReadOnly
// type.
//
type FixedDecimalReadOnly struct {

	fixedDecimal BigIntFixedDecimal

}

// GetBigIntNum - returns the numeric value of the underlying BigIntFixedDecimal
// as a type BigIntNum.
//
func (fDecRO *FixedDecimalReadOnly) GetBigIntNum() BigIntNum {

	fDecRO.fixedDecimal.IsValid()

	return fDecRO.fixedDecimal.GetBigIntNum()
}

// GetIntAry - Returns the numeric value of the underlying BigIntFixedDecimal
// as a type IntAry.
//
func (fDecRO *FixedDecimalReadOnly) GetIntAry() (IntAry,error) {

	fDecRO.fixedDecimal.IsValid()

	return fDecRO.fixedDecimal.GetIntAry()
}

// GetFixedDecimal - Returns a deep copy of the underlying
// BigIntFixedDecimal value.
func (fDecRO *FixedDecimalReadOnly) GetFixedDecimal() BigIntFixedDecimal {

	fDecRO.fixedDecimal.IsValid()

	return fDecRO.fixedDecimal.CopyOut()
}

// GetInteger - Returns the *big.Int integer value from the
// underlying BigIntFixedDecimal
func (fDecRO *FixedDecimalReadOnly) GetInteger() *big.Int {

	fDecRO.fixedDecimal.IsValid()

	return fDecRO.fixedDecimal.GetInteger()

}

// GetNumStr - Converts the underlying BigIntFixedDecimal
// value to a string of numbers which includes the decimal
// place and decimal digits if they exist. Note that
// the period ('.') is the decimal separator character
// always used in the returned number string to separate
// integer and fractional digits.
//
func (fDecRO *FixedDecimalReadOnly) GetNumStr() string {

	fDecRO.fixedDecimal.IsValid()

	return fDecRO.fixedDecimal.GetNumStr()

}

// GetPrecision - returns a unsigned integer specifying
// the number of digits to the right of the decimal place.
//
func (fDecRO *FixedDecimalReadOnly) GetPrecision() uint {

	fDecRO.fixedDecimal.IsValid()

	return fDecRO.fixedDecimal.GetPrecision()

}

// IsValid - Performs diagnostic and remedial actions on
// the underlying BigIntFixedDecimal instance. Returns
// 'false' if the instance is uninitialized.
//
func (fDecRO *FixedDecimalReadOnly) IsValid() bool {

	return fDecRO.fixedDecimal.IsValid()
}

func (fDecRO *FixedDecimalReadOnly) IsZero() bool {

	return fDecRO.fixedDecimal.IsZero()
}

// NewZero - Creates and returns a new FixedDecimalReadOnly
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
//
func (fDecRO FixedDecimalReadOnly) NewNumStr(numStr string) (FixedDecimalReadOnly, error) {
	ePrefix := "FixedDecimalReadOnly.NewNumStr() "

	fo2 := FixedDecimalReadOnly{}

	fo2.fixedDecimal = BigIntFixedDecimal{}.NewZero(0)

	err := fo2.fixedDecimal.SetNumStr(numStr)

	if err != nil {
		return FixedDecimalReadOnly{}.NewZero(0),
			fmt.Errorf(ePrefix )
	}

	return fo2, nil
}

// NewFixedDecimal - Receives a BigIntFixedDecimal instance
// as an input parameter and returns a new FixedDecimalReadOnly
// object.
//
func (fDecRO FixedDecimalReadOnly) NewFixedDecimal(
	fixedDecimal BigIntFixedDecimal) FixedDecimalReadOnly {

	fixedDecimal.IsValid()

	f2 := FixedDecimalReadOnly{}.NewZero(0)
	f2.fixedDecimal.CopyIn(fixedDecimal)

	return f2
}

// NewInt - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'intValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue.
//
func (fDecRO FixedDecimalReadOnly) NewInt(
	intValue int,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = BigIntFixedDecimal{}.NewInt(intValue, precision)

	return f2

}

// NewInt32 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'int32Value' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue.
//
func (fDecRO FixedDecimalReadOnly) NewInt32(
	int32Value int32,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = BigIntFixedDecimal{}.NewInt32(int32Value, precision)

	return f2

}

// NewInt64 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'int64Value' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue.
//
func (fDecRO FixedDecimalReadOnly) NewInt64(
	int64Value int64,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = BigIntFixedDecimal{}.NewInt64(int64Value, precision)

	return f2

}

// NewUInt - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'uintValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue.
//
func (fDecRO FixedDecimalReadOnly) NewUInt(
	uintValue,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = BigIntFixedDecimal{}.NewUInt(uintValue, precision)

	return f2

}

// NewUInt32 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'uintValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue.
//
func (fDecRO FixedDecimalReadOnly) NewUInt32(
	uint32Value uint32,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = BigIntFixedDecimal{}.NewUInt32(uint32Value, precision)

	return f2

}

// NewUInt64 - Creates and returns a new FixedDecimalReadOnly instance initialized
// to by the input parameters 'uintValue' and 'precision'. 'precision' specifies
// the number of digits to the right of the decimal place in 'intValue.
//
func (fDecRO FixedDecimalReadOnly) NewUInt64(
	uint64Value uint64,
	precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = BigIntFixedDecimal{}.NewUInt64(uint64Value, precision)

	return f2
}

// NewZero - Creates and returns a new FixedDecimalReadOnly instance initialized
// to zero. The input parameter 'precision' specifies the number of zero digits
// to the right of the decimal place.
//
func (fDecRO FixedDecimalReadOnly) NewZero(precision uint) FixedDecimalReadOnly {

	f2 := FixedDecimalReadOnly{}

	f2.fixedDecimal = BigIntFixedDecimal{}.NewZero(0)

	return f2

}


// BigIntFixedDecimal - A light data transfer structure used to represent
// a numeric value with a fixed number of decimal digits. Used primarily
// for intensive and complex calculations.
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


// Ceiling - Returns the ceiling integer value for the current
// BigIntFixedDecimal instance.
//
// Ceiling is defined as: The least, or lowest value integer,
// which is greater than or equal to the numeric value of the
// current BigIntFixedDecimal.
//
// Reference Wikipedia:
// 				https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
// Examples
// ========
//
// 						Initial 		 Ceiling
//  					 Value				Value
// 						-------      -------
//  						5.95					6
//  						5.05					6
//  						5							5
// 						 -5.05			 	 -5
//  						2.4				  	3
//  						2.9					 	3
// 						 -2.7				 	 -2
// 						 -2					 	 -2
//
func (bigIFd *BigIntFixedDecimal) Ceiling() BigIntFixedDecimal {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	cmpZeroResult := bigIFd.integerNum.Cmp(big.NewInt(0))

	if cmpZeroResult == 0 {
		return BigIntFixedDecimal{}.NewZero(0)
	}

	ceiling := big.NewInt(0).Set(bigIFd.integerNum)

	if bigIFd.precision > 0 {

		scale := big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(bigIFd.precision)),
			nil)

		ceiling.Quo(ceiling, scale)

		if cmpZeroResult == 1 {
			// signVal must be plus
			ceiling.Add(ceiling, big.NewInt(1))
		}

	}

	// else bigIFd.precision must be zero
	return BigIntFixedDecimal{}.New(ceiling, 0)
}

// ChangeSign - This method will change the sign of the
// current BigIntFixedDecimal numeric value. If the value
// is negative, this method will change the sign to
// positive. Likewise, if the sign is currently positive,
// calling this method will change the sign to negative.
func (bigIFd *BigIntFixedDecimal) ChangeSign() {
	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	if bigIFd.integerNum.Cmp(big.NewInt(0))==0 {
		return
	}

	bigIFd.integerNum.Neg(bigIFd.integerNum)
}

// CopyIn - Receives a BigIntFixedDecimal type and copies the
// value to the current BigIntFixedDecimal instance.
//
func (bigIFd *BigIntFixedDecimal) CopyIn(fd BigIntFixedDecimal) {


	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), bigIFd.precision)
	}

	fd.IsValid()

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

	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), bigIFd.precision)
	}

	fd.IsValid()

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

	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), bigIFd.precision)
	}

	return BigIntFixedDecimal{}.New(bigIFd.integerNum, bigIFd.precision)
}

// DivideByTenToPower - Divides the numeric value of the current
// BigIntFixedDecimal by 10 to the power of 'exponent'.
//
//       result = BigIntFixedDecimal / 10^exponent
//
//
// Input Parameter
// ===============
//
// exponent	uint	- The value of the current BigIntFixedDecimal
//									instance will be divided by ten raised to
// 									the power of 'exponent'.
//
// This method will destroy and overwrite the previous value of
// the current BigIntFixedDecimal instance with the results of
// this calculation.
//
func (bigIFd *BigIntFixedDecimal) DivideByTenToPower(exponent uint) {

	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), bigIFd.precision)
	}

	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		return
	}

	scale :=
		big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(exponent)), nil)

	factor := BigIntFixedDecimal{}.New(scale, 0)

	result, _ :=
		BigIntMathDivide{}.FixedDecimalFracQuotient(
			bigIFd.CopyOut(), factor, bigIFd.precision+exponent)

	bigIFd.CopyIn(result)
}

// DivideByTwoToPower - Performs integer division by two using a 'right-shift'
// technique. Remainders from this division operation are discarded, only the
// integer quotient is returned. When the calculation is completed, the value
// of the integer quotient will replace the old value of the current BigIntFixedDecimal
// instance.
//
// Example:
// ========
//
//            quotient =  BigIntFixedDecimal / 2^(exponent)
//
// In this example BigIntFixedDecimal= 33,123.456, so 33,123.456/ 2^8:
//
//		(1) The fractional quotient of 33,123.456/256 (or 2^8) is 129.3885.
//
//    (2) This method will use a right shift technique on the integer value
//        33123456 / 2^(8) to generate an integer quotient of 129388.
//
// Consider the example BigIntFixedDecimal = 33123456 (no decimal fraction):
//
//    Dividing 33123456 / 2^8 = fractional quotient = 129388.5
//    Be careful when using this method.
//
// **************************************************************************
// (1) 	Be sure to consider the outcomes when sending a decimal fraction to
// 			this method.
//
// (2)	Results returned by this method will always have precision = 0
//
// **************************************************************************
//
// Note:	This method will destroy and overwrite the previous value
// 				of the current BigIntFixedDecimal instance with the results
// 				of this calculation.
//
func (bigIFd *BigIntFixedDecimal) DivideByTwoToPower(exponent uint) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
		return
	}

	bigIFd.integerNum =
		BigIntMathDivide{}.BigIntDividedByTwoToPower(
			bigIFd.integerNum,
			exponent)

	bigIFd.precision = 0
}


// Empty - Reinitialize the current BigIntFixedDecimal instance
// to a zero value with zero precision.
//
func (bigIFd *BigIntFixedDecimal) Empty() {

	bigIFd.integerNum = big.NewInt(0)
	bigIFd.precision = 0

}

// Floor - Returns the floor integer value for the current
// BigIntFixedDecimal.
//
// In mathematics and computer science, the floor function
// is the function that takes as input a real number x dnd
// gives as output the greatest integer less than or equal
// to x.
//
// Source:
// https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
// Examples
// ========
//
//				 BigIntFixedDecimal           Floor
//							Value                   Value
//				 ------------------         ----------
//              0													0
//              4                         4
//							3.2												3
//							2.9	                      2
//						 -2.7											 -3
//						 -2                        -2
//
func (bigIFd *BigIntFixedDecimal) Floor() BigIntFixedDecimal {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	cmpZeroResult := bigIFd.integerNum.Cmp(big.NewInt(0))

	if cmpZeroResult == 0 {
		return BigIntFixedDecimal{}.NewZero(0)
	}

	floor := big.NewInt(0).Set(bigIFd.integerNum)

	if bigIFd.precision > 0 {

		scale := big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(bigIFd.precision)),
			nil)

		floor.Quo(floor, scale)

		if cmpZeroResult == -1 {
			// signVal must be -1
			floor.Add(floor, big.NewInt(-1))
		}

	}

	// else bigIFd.precision must be zero
	return BigIntFixedDecimal{}.New(floor, 0)
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
		absBigInt.Neg(absBigInt)
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

// GetBigIntNum - Returns a new BigIntNum object initialized
// to the value of the current BigIntFixedDecimal instance.
//
func (bigIFd *BigIntFixedDecimal) GetBigIntNum() BigIntNum {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	return BigIntNum{}.NewBigInt(bigIFd.integerNum, bigIFd.precision)
}

// GetInteger - Returns the 'integerNum' for the current
// BigIntFixedDecimal instance. The returned *big.Int type
// contains all the numeric digits which comprise the fixed
// decimal numerical value represented by this BigIntFixedDecimal
// instance. If the value includes fractional digits, these
// too are included in the returned integer value.
//
// Example:
// ========
//
// BigIntFixedDecimal      Returned
//  Numeric Value        Integer Value
// ------------------    -------------
//		582.12345            58212345
//
func (bigIFd *BigIntFixedDecimal) GetInteger() *big.Int {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	return big.NewInt(0).Set(bigIFd.integerNum)
}

// GetIntegerFractionalParts - Returns two BigIntFixedDecimals comprising the integer
// and fractional parts of the current BigIntFixedDecimal numeric value.
//
// Examples:
// =========
//
// BigIntFixedDecimal              Returned         Returned
//   Numeric Value              Integer Value    Fractional Value
// ------------------           -------------    ----------------
//
//  859649.123456789								859649					0.123456789
// -859649.123456789							 -859649				 -0.123456789
//
func (bigIFd *BigIntFixedDecimal) GetIntegerFractionalParts() (integer BigIntFixedDecimal, fraction BigIntFixedDecimal) {

		integer = bigIFd.CopyOut()

		fraction = bigIFd.CopyOut()

		if integer.integerNum.Cmp(big.NewInt(0)) == 0 {
			integer.precision = 0
			return integer, fraction
		}

		scale :=
			big.NewInt(0).Exp(
				big.NewInt(10),
				big.NewInt(int64(bigIFd.precision)),
				nil)

		integer.integerNum.Quo(integer.integerNum,scale)
		integer.precision = 0

		scratch := integer.CopyOut()

		scratch.integerNum.Mul(scratch.integerNum, scale)

		fraction.integerNum.Sub(fraction.integerNum, scratch.integerNum)

		return integer, fraction
}

// GetIntAry - Returns a new IntAry instance initialized to the
// value of the current BigIntFixedDecimal object.
func (bigIFd *BigIntFixedDecimal) GetIntAry() (IntAry, error) {

	ePrefix := "BigIntFixedDecimal.GetIntAry() "

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	ia, err := IntAry{}.NewBigInt(bigIFd.integerNum, int(bigIFd.precision))

	if err != nil {
		return IntAry{}.NewZero(0),
		 fmt.Errorf(ePrefix + "Error returned: %v", err.Error())
	}

	return ia, nil
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

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

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

func (bigIFd *BigIntFixedDecimal) IsZero() bool {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	if bigIFd.integerNum.Cmp(big.NewInt(0))==0 {
		return true
	}

	return false
}

// MultiplyByTenToPower - Multiplies the numeric value of the current
// BigIntFixedDecimal by 10 to the power of 'exponent'.
//
//       result = BigIntFixedDecimal x 10^exponent
//
// Examples:
// =========
// 	BigIntFixedDecimal
//			Value									Exponent							   Result
//  ------------------        --------              -----------------
//   105.6752										 0										      105.6752
//   105.6752                    1									       1056.752
//   105.6752                    2                        10567.52
//   105.6752                    3                       105675.2
//   105.6752                    8                  10567520000
//
// Input Parameter
// ===============
//
// exponent	uint	- The value of the current BigIntFixedDecimal
//									instance will be multiplied by ten raised to
// 									the power of 'exponent'.
//
//
// Note:	(1)	This method will delete trailing fractional zeros from
// 			 			the returned product.
//
// 				(2)	This method will destroy and overwrite the previous value
// 						of the current BigIntFixedDecimal instance with the results
// 						of this calculation.
//
func (bigIFd *BigIntFixedDecimal) MultiplyByTenToPower(exponent uint) {

	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), bigIFd.precision)
	}

	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		bigIFd.precision = 0
		return
	}

	scale :=
		big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(exponent)), nil)

	factor := BigIntFixedDecimal{}.New(scale, 0)

	result := BigIntMathMultiply{}.FixedDecimalMultiply(bigIFd.CopyOut(), factor)

	result.TrimTrailingFracZeros()

	bigIFd.CopyIn(result)

}

// MultiplyByTwoToPower - Multiplies the numeric value of the current
// BigIntFixedDecimal by 2 to the power of 'exponent'.
//
//        product = BigIntFixedDecimal x 2^exponent
//
// When the calculation is completed, the value of 'product' will
// replace the old value of the current BigIntFixedDecimal instance.
//
// Examples:
// =========
//
//	 BigIntFixedDecimal 			 	exponent		  product
//	interNum		Precision
// -------------------------------------------------------------
//	12345						5								15				4045.2096
//								(0.12345 x 2^15 = 4045.2096)
// -------------------------------------------------------------
//
//    571						1								 8			 14617.6
//                (57.1 x 2^8 = 14617.6)
// -------------------------------------------------------------
//
// Note:	(1)	This method will delete trailing fractional zeros from
// 			 			the returned product.
//
// 				(2)	This method will destroy and overwrite the previous value
// 						of the current BigIntFixedDecimal instance with the results
// 						of this calculation.
//
func (bigIFd *BigIntFixedDecimal) MultiplyByTwoToPower(exponent uint) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
		return
	}

	bigIFd.integerNum, bigIFd.precision =
		BigIntMathMultiply{}.BigIntMultiplyByTwoToPower(
			bigIFd.integerNum,
			bigIFd.precision,
			exponent)

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


// NewNumStr - Receives a number string as input and returns
// a new BigIntFixedDecimal instance.
//
// A number string is a string of numeric digits which may,
// or may not, be prefixed with a minus sign ('-') indicating
// a negative number. If the numeric string of digits is prefixed
// by a left parenthesis and suffixed by a corresponding right
// parenthesis, this also indicates a negative value.
//
// The numeric string of digits may also contain a period
// ('.') which is treated as a decimal separator and used to
// separate integer and fractional digits within the number
// string.
//
// The only decimal separator recognized by this method is the
// period ('.').
//
func (bigIFd BigIntFixedDecimal) NewNumStr(numStr string) (BigIntFixedDecimal, error) {

	ePrefix := "BigIntFixedDecimal.NewNumStr() "

	fixedDecimal := BigIntFixedDecimal{}

	err := fixedDecimal.SetNumStr(numStr)

	if err != nil {
		return BigIntFixedDecimal{},
			fmt.Errorf(ePrefix + "Error returned by fixedDecimal.SetNumStr(numStr). " +
				"Error='%v' ", err.Error())
	}

	return fixedDecimal, nil
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
// Note: This method does NOT trim or delete trailing fractional zero
// digits.
//
func (bigIFd *BigIntFixedDecimal) RoundToDecPlace(precision uint) {

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
	base10 := big.NewInt(10)

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bigIFd.precision < precision {

		deltaPrecision := precision - bigIFd.precision

		scale =
			big.NewInt(0).Exp(
				base10,
				big.NewInt(int64(deltaPrecision)),
				nil)

		bigIFd.integerNum.Mul(bigIFd.integerNum, scale)

		bigIFd.precision += deltaPrecision

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

	result := BigIntMathAdd{}.FixedDecimalAdd(bigIFd.CopyOut(), bigNumRound5)

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

// SetNumStr - Initializes the current BigIntFixedDecimal
// instance of the numeric value of the number string input
// parameter.
//
// A number string is a string of numeric digits which may
// or may not be prefixed with a minus sign ('-'). The numeric
// string of digits may also contain a decimal separator period
// ('.') which is used to separate integer and fractional digits
// within the number string.
//
// The only decimal separator recognized by this method is the
// period ('.').
//
func (bigIFd *BigIntFixedDecimal) SetNumStr(numStr string) error {

	ePrefix := "BigIntFixedDecimal.SetNumStr() "

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	if len(numStr) == 0 {
		return errors.New(ePrefix + "Error: Input parameter 'numStr' is an EMPTY string!")
	}

	decimalSeparator := '.'

	baseRunes := []rune(numStr)
	lBaseRunes := len(baseRunes)

	newPrecision := uint(0)
	newAbsBigInt := big.NewInt(0)
	baseTen := big.NewInt(10)
	hasLeftParen := false
	hasRightParen := false
	hasMinusSign := false
	startFractionalDigits := false
	isStartNumericDigits := false
	isEndNumericDigits := false
	numOfNumericDigits := uint(0)

	for i := 0; i < lBaseRunes; i++ {

		if isEndNumericDigits == true {
			continue
		}

		if baseRunes[i] == '-' && isStartNumericDigits==false {
			hasMinusSign = true
			continue
		}

		if baseRunes[i] == '(' && isStartNumericDigits==false {
			hasLeftParen = true
			continue
		}

		if baseRunes[i] == ')' 	&&
				isStartNumericDigits == true &&
					hasLeftParen == true {

			hasRightParen = true
			isEndNumericDigits = true

		}

		if baseRunes[i] == decimalSeparator {
			startFractionalDigits = true
			continue
		}

		if baseRunes[i] >= '0' &&
			baseRunes[i] <= '9' {

			isStartNumericDigits = true

			newAbsBigInt.Mul(newAbsBigInt,baseTen)

			newAbsBigInt.Add(newAbsBigInt, big.NewInt(int64(baseRunes[i]-48)))

			numOfNumericDigits++

			if startFractionalDigits == true {
				newPrecision ++
			}

			continue
		}

	}

	if numOfNumericDigits==0 {
		return fmt.Errorf(ePrefix +
			"Error: No numeric digits were found in input parameter 'numStr'. " +
			"numStr='%v'", numStr)
	}

	if hasMinusSign == true || (hasLeftParen==true && hasRightParen==true) {

		newAbsBigInt.Neg(newAbsBigInt)

	}

	bigIFd.integerNum.Set(newAbsBigInt)
	bigIFd.precision = newPrecision

	return nil
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

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	bigIFd.precision = precision

}

// TrimTrailingFracZeros - This method will delete non-significant
// trailing zeros from the fractional digits of the current
// BigIntFixedDecimal numerical value.
//
// Examples:
//					Initial Value			Trimmed Value
//						456.123000 			 456.123
//							0.000					 0
//							7.0						 7
//					 -456.123000			-456.123
//
func (bigIFd *BigIntFixedDecimal) TrimTrailingFracZeros() {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
		return
	}

	if bigIFd.precision == 0 {
		return
	}

	// bigIFd.precision must be GREATER THAN ZERO
	// Delete trailing fractional zeros

	scrap := big.NewInt(0)
	biBase10 := big.NewInt(10)
	biBaseZero := big.NewInt(0)
	newintegerNum, mod10 := big.NewInt(0).QuoRem(bigIFd.integerNum, biBase10, scrap)

	for mod10.Cmp(biBaseZero) == 0 && bigIFd.precision > 0 {
		bigIFd.integerNum.Set(newintegerNum)
		bigIFd.precision--
		newintegerNum, mod10 = big.NewInt(0).QuoRem(bigIFd.integerNum, biBase10, scrap)
	}

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
