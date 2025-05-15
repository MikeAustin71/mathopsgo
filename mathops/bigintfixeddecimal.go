package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
)

// BigIntFixedDecimal
//
// A light data transfer structure used to represent a numeric value
// with a fixed number of decimal digits. Used primarily for
// intensive or complex calculations.
type BigIntFixedDecimal struct {
	integerNum *big.Int // All the numeric digits, both integer and fractional,
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

	// Numeric Separators
	decimalSeparator rune // Character used to separate integer and fractional digits ('.')

	thousandsSeparator rune // Character used to separate thousands (1,000,000,000

	currencySymbol rune // Currency Symbol

}

var _ INumMgr = (*BigIntFixedDecimal)(nil)

// Ceiling
//
// Returns the ceiling integer value for the current
// BigIntFixedDecimal instance.
//
// Ceiling is defined as: The least, or lowest value integer, which
// is greater than or equal to the numeric value of the current
// BigIntFixedDecimal.
//
//	Reference Wikipedia
//	===================
//
//	  https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//	Examples
//	========
//
//	  Initial    Ceiling
//	  Value       Value
//	  -------    -------
//	   5.95         6
//	   5.05         6
//	   5            5
//	  -5.05        -5
//	   2.4          3
//	   2.9          3
//	  -2.7         -2
//	  -2           -2
func (bigIFd *BigIntFixedDecimal) Ceiling() (BigIntFixedDecimal, error) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	cmpZeroResult := bigIFd.integerNum.Cmp(big.NewInt(0))

	if cmpZeroResult == 0 {
		return new(BigIntFixedDecimal).NewZero(0)
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
	return new(BigIntFixedDecimal).New(ceiling, 0)
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
		return
	}

	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		return
	}

	bigIFd.integerNum.Neg(bigIFd.integerNum)
}

// Cmp - Compares the numeric values of two BigIntFixedDecimal
// instances.
//
// If the current BigIntFixedDecimal value is greater than input
// parameter 'fd2', the method returns '1'.
//
// If the current BigIntFixedDecimal value is equal to that of the
// input parameter 'fd2', the method returns '0'.
//
// If the current BigIntFixedDecimal value is less than the input
// parameter 'fd2', the method returns '-1'.
//
// Examples
// ========
//
//						BigIntFixedDecimal         'fd2'        Return
//									Value                Value        Value
//						------------------         ------       -------
//	                5                     2             1
//	                5.2                   5.1           1
//	                5.2                   5.2           0
//	           837123.4              837123.5          -1
//	                0                     0.1          -1
//	               35.123456             40.5          -1
//	               35.123456              2.5           1
func (bigIFd *BigIntFixedDecimal) Cmp(fd2 BigIntFixedDecimal) (int, error) {

	ePrefix := "BigIntFixedDecimal.Cmp"

	err := bigIFd.IsValid(ePrefix + " Testing 'bigIFd'")

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err := bigIFd.IsValid(ePrefix + \" Testing 'bigIFd'\")",
				ErrContext: "Current BigIntFixedDecimal instance 'bigIFd' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	err = fd2.IsValid(ePrefix + " Testing 'fd2'")

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = fd2.IsValid(ePrefix + \" Testing 'fd2'\")",
				ErrContext: "Input parameter 'fd2' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	if bigIFd.precision == fd2.precision {
		return bigIFd.integerNum.Cmp(fd2.integerNum), nil
	}

	bigTen := big.NewInt(10)

	if fd2.precision > bigIFd.precision {

		delta := int64(fd2.precision - bigIFd.precision)

		fdValue := big.NewInt(0).Set(bigIFd.integerNum)

		scale := big.NewInt(0).Exp(bigTen, big.NewInt(delta), nil)

		fdValue.Mul(fdValue, scale)

		return fdValue.Cmp(fd2.integerNum), nil

	}

	// MUST BE bigIFd.precision > fd2.precision
	delta := int64(bigIFd.precision - fd2.precision)

	fd2Value := big.NewInt(0).Set(fd2.integerNum)

	scale := big.NewInt(0).Exp(bigTen, big.NewInt(delta), nil)

	fd2Value.Mul(fd2Value, scale)

	return bigIFd.integerNum.Cmp(fd2Value), nil
}

// CmpZero Compares the current BigIntFixedDecimal to Zero and
// returns an integer flag as follows:
//
// +1 = BigIntFixedDecimal > 0
//
//	0 = BigIntFixedDecimal == 0
//
// -1 = BigINtFixedDecimal < 0
func (bigIFd *BigIntFixedDecimal) CmpZero() int {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	return bigIFd.integerNum.Cmp(big.NewInt(0))
}

// CopyIn - Receives a BigIntFixedDecimal type and copies the
// value to the current BigIntFixedDecimal instance.
func (bigIFd *BigIntFixedDecimal) CopyIn(fd BigIntFixedDecimal) error {

	ePrefix := "BigIntFixedDecimal.CopyIn"

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	err := fd.IsValid(" Testing input parameter 'fd'")

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "err := fd.IsValid(\" Testing input parameter 'fd'\")",
			ErrContext: "Input parameter 'fd' is INVALID!",
			ErrMessage: err.Error(),
		}
	}

	intVal, err := fd.GetInteger()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "intVal, err := fd.GetInteger()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bigIFd.integerNum = big.NewInt(0).Set(intVal)

	bigIFd.precision, err = fd.GetPrecision()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "bigIFd.precision, err = fd.GetPrecision()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// CopyInPtr - Receives a pointer to a BigIntFixedDecimal type and
// copies the value to the current BigIntFixedDecimal instance.
func (bigIFd *BigIntFixedDecimal) CopyInPtr(fd *BigIntFixedDecimal) error {

	ePrefix := "BigIntFixedDecimal.CopyInPtr"

	err := fd.IsValid(ePrefix + " Testing input parameter 'fd'")

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "err := fd.IsValid(ePrefix + \" Testing input parameter 'fd'\")",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	intVal, err := fd.GetInteger()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "intVal, err := fd.GetInteger()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bigIFd.integerNum = big.NewInt(0).Set(intVal)

	bigIFd.precision, err = fd.GetPrecision()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "bigIFd.precision, err = fd.GetPrecision()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// CopyOut - Returns a new BigIntFixedDecimal instance which is
// a deep copy of the current BigIntFixedDecimal instance.
func (bigIFd *BigIntFixedDecimal) CopyOut() BigIntFixedDecimal {

	if bigIFd.integerNum == nil {
		bigIFd.SetNumericValue(big.NewInt(0), bigIFd.precision)
	}

	return new(BigIntFixedDecimal).New(bigIFd.integerNum, bigIFd.precision)
}

// DivideByTenToPower - Divides the numeric value of the current
// BigIntFixedDecimal by 10 to the power of 'exponent'.
//
//	result = BigIntFixedDecimal / 10^exponent
//
// Input Parameter
// ===============
//
// exponent	uint	- The value of the current BigIntFixedDecimal
//
//	instance will be divided by ten raised to
//	the power of 'exponent'.
//
// This method will destroy and overwrite the previous value of
// the current BigIntFixedDecimal instance with the results of
// this calculation.
func (bigIFd *BigIntFixedDecimal) DivideByTenToPower(exponent uint) error {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	ePrefix := "BigIntFixedDecimal.DivideByTenToPower"

	var err error

	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		return err
	}

	scale :=
		big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(exponent)), nil)

	factor := new(BigIntFixedDecimal).New(scale, 0)

	result, err :=
		BigIntMathDivide{}.FixedDecimalFracQuotient(
			bigIFd.CopyOut(), factor, bigIFd.precision+exponent)

	if err != nil {
		return fmt.Errorf("%v"+
			"Error returned from BigIntMathDivide{}.FixedDecimalFracQuotient(bigIFd.CopyOut(), factor, bigIFd.precision+exponent)\n"+
			"Error= %v\n",
			ePrefix,
			err.Error())
	}

	err = bigIFd.CopyIn(result)

	if err != nil {
		return fmt.Errorf("%v"+
			"Error returned by bigIFd.CopyIn(result)\n"+
			"result= %v\n Error= %v\n",
			result,
			err.Error())
	}

	return nil
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
//	quotient =  BigIntFixedDecimal / 2^(exponent)
//
// In this example BigIntFixedDecimal= 33,123.456, so 33,123.456/ 2^8:
//
//			(1) The fractional quotient of 33,123.456/256 (or 2^8) is 129.3885.
//
//	   (2) This method will use a right shift technique on the integer value
//	       33123456 / 2^(8) to generate an integer quotient of 129388.
//
// Consider the example BigIntFixedDecimal = 33123456 (no decimal fraction):
//
//	Dividing 33123456 / 2^8 = fractional quotient = 129388.5
//	Be careful when using this method.
//
// **************************************************************************
// (1) 	Be sure to consider the outcomes when sending a decimal fraction to
//
//	this method.
//
// (2)	Results returned by this method will always have precision = 0
//
// **************************************************************************
//
// Note:	This method will destroy and overwrite the previous value
//
//	of the current BigIntFixedDecimal instance with the results
//	of this calculation.
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
//					 BigIntFixedDecimal           Floor
//								Value                   Value
//					 ------------------         ----------
//	             0													0
//	             4                         4
//								3.2												3
//								2.9	                      2
//							 -2.7											 -3
//							 -2                        -2
func (bigIFd *BigIntFixedDecimal) Floor() BigIntFixedDecimal {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	cmpZeroResult := bigIFd.integerNum.Cmp(big.NewInt(0))

	if cmpZeroResult == 0 {
		return new(BigIntFixedDecimal).NewZero(0)
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
	return new(BigIntFixedDecimal).New(floor, 0)
}

// FormatNumStr - converts the numeric value of the current BigIntFixedDecimal
// instance to a number string. The returned number string will consist of a
// string of numeric digits. If the number contains fractional digits, the
// decimal separator period (.) will be used to separate integer and fractional
// digits within the string. There are no thousands separator present in the
// returned string.
//
// The input parameter 'negValMode' is of type NegativeValueFmtMode.
// NegativeValueFmtMode encompasses a series of constants that are used
// to format negative values in a number string.
//
// Valid NegativeValueFormatMode's are defined as follows:
//
// LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//
//	 		a leading minus sign.
//			Example: -123456.78
//
// PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//
//	surrounding parentheses.
//	Example: (123456.78)
//
// ABSOLUTEPURENUMSTRFMTMODE - Formats a pure integerNum string with
//
//	 absolute (positive) integer value
//	 and no decimal point separator.
//	Example: (12345678)
func (bigIFd *BigIntFixedDecimal) FormatNumStr(negValMode NegativeValueFmtMode) string {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
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

// GetBigInt
//
// Converts the current instance of BigIntFixedDecimal to
// a numeric value and returns that value as a type
// '*big.Int'.
func (bigIFd *BigIntFixedDecimal) GetBigInt() (*big.Int, error) {

	ePrefix := "BigIntFixedDecimal.GetBigInt"

	err := bigIFd.IsValid(ePrefix + " Testing 'bigIFd'")

	if err != nil {

		return big.NewInt(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err := bigIFd.IsValid(ePrefix + \" Testing 'bigIFd'\")",
				ErrContext: "Testing validity of current BigIntFixedDecimal instance 'bigIFd'.\n" +
					"'bigIFd' FAILED validation testing. 'bigIFd' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	return bigIFd.integerNum, nil
}

// GetBigIntNum - Returns a new BigIntNum object initialized
// to the value of the current BigIntFixedDecimal instance.
func (bigIFd *BigIntFixedDecimal) GetBigIntNum() (BigIntNum, error) {

	ePrefix := "BigIntFixedDecimal.GetBigIntNum"

	err := bigIFd.IsValid(ePrefix + " Testing 'bigIFd'")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err := bigIFd.IsValid(ePrefix + \" Testing 'bigIFd'\")",
				ErrContext: "Testing validity of current BigIntFixedDecimal instance 'bigIFd'.",
				ErrMessage: err.Error(),
			}
	}

	bigINum, err := new(BigIntNum).NewBigInt(bigIFd.integerNum, bigIFd.precision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bigINum, err := new(BigIntNum).NewBigInt(bigIFd.integerNum, bigIFd.precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bigINum, nil
}

// GetDecimal
//
// Converts the current instance of BigIntFixedDecimal to a numeric
// value and returns that value as a type 'Decimal'.
func (bigIFd *BigIntFixedDecimal) GetDecimal() (Decimal, error) {

	ePrefix := "BigIntFixedDecimal.GetDecimal"

	err := bigIFd.IsValid(ePrefix + " Testing 'bigIFd'")

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err := bigIFd.IsValid(ePrefix + \" Testing 'bigIFd'\")",
				ErrContext: "Testing validity of current BigIntFixedDecimal instance 'bigIFd'.",
				ErrMessage: err.Error(),
			}
	}

	bigINum, err := new(BigIntNum).NewBigInt(bigIFd.integerNum, bigIFd.precision)

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "bigINum, err := new(BigIntNum).NewBigInt(bigIFd.integerNum, bigIFd.precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	decNum, err := bigINum.GetDecimal()

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "decNum, err := bigINum.GetDecimal()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return decNum, nil
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
//
//	Numeric Value        Integer Value
//
// ------------------    -------------
//
//	582.12345            58212345
func (bigIFd *BigIntFixedDecimal) GetInteger() (*big.Int, error) {

	ePrefix := "BigIntFixedDecimal.GetInteger()"

	if bigIFd.integerNum == nil {

		bigIFd.integerNum = big.NewInt(0)

		return big.NewInt(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: "if bigIFd.integerNum == nil {",
				ErrMessage: "ERROR: bigIFd.integerNum, an instance of type\n" +
					"BigIntFixedDecimal, is 'nil'!",
			}
	}

	retVal := big.NewInt(0).Set(bigIFd.integerNum)

	return retVal, nil
}

// GetIntegerFractionalParts - Returns two BigIntFixedDecimals comprising the integer
// and fractional parts of the current BigIntFixedDecimal numeric value.
//
// Examples:
// =========
//
// BigIntFixedDecimal              Returned       Integer    		 Returned					 Fraction
//
//	Numeric Value              Integer Value    Precision		Fractional Value	   Precision
//
// ------------------           -------------   ----------		----------------     ---------
//
//	859649.123456789								859649					0						0.123456789            9
//
// -859649.123456789							 -859649				 	0					 -0.123456789            9
func (bigIFd *BigIntFixedDecimal) GetIntegerFractionalParts() (integer BigIntFixedDecimal, fraction BigIntFixedDecimal) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	integer = new(BigIntFixedDecimal).NewZero(0)
	fraction = new(BigIntFixedDecimal).NewZero(0)

	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		return integer, fraction
	}

	scale := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(0).SetUint64(uint64(bigIFd.precision)), nil)

	scratch := big.NewInt(0)

	intRadicand, fracRadicand := big.NewInt(0).QuoRem(bigIFd.integerNum, scale, scratch)

	integer = new(BigIntFixedDecimal).New(intRadicand, 0)

	fraction = new(BigIntFixedDecimal).New(fracRadicand, bigIFd.precision)

	return integer, fraction
}

// GetIntAry - Returns a new IntAry instance initialized to the
// value of the current BigIntFixedDecimal object.
func (bigIFd *BigIntFixedDecimal) GetIntAry() (IntAry, error) {

	ePrefix := "BigIntFixedDecimal.GetIntAry() "

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	ia, err := new(IntAry).NewBigInt(bigIFd.integerNum, int(bigIFd.precision))

	if err != nil {
		return new(IntAry).New(),
			fmt.Errorf(ePrefix+"Error returned: %v", err.Error())
	}

	return ia, nil
}

// GetMagnitude - Returns the magnitude of the current BigIntFixedDecimal
// as a type *big.Int integer value.
//
// Magnitude as used here is defined as the power of 10 which generates a
// value less than or equal to the current BigIntFixedDecimal value
//
//	10^magnitude  <= BigIntFixedDecimal
//
// Examples
// ========
//
//	        BigIntFixedDecimal
//				   			Value									magnitude
//					------------------					---------
//
//				  		 963,256										5
//										 2										0
//										32										1
//				 8,456,123,921					  				9
//	                  2.2										0
//	      8,456,123,912.123									9
//	           -643,212.123									5
//	                324.123456							2
func (bigIFd *BigIntFixedDecimal) GetMagnitude() (*big.Int, error) {

	bigZero := big.NewInt(0)

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
		return bigZero, nil
	}

	target := big.NewInt(0).Set(bigIFd.integerNum)
	bigTen := big.NewInt(10)

	if bigIFd.precision > 0 {
		target.Quo(
			target,
			big.NewInt(0).Exp(bigTen, big.NewInt(int64(bigIFd.precision)),
				nil))
	}

	magnitude, err := BigIntMath{}.GetMagnitude(target)

	if err != nil {
		ePrefix := "BigIntFixedDecimal.GetMagnitude() "
		return bigZero,
			fmt.Errorf(ePrefix+
				"Error returned "+
				"Error='%v' ", err.Error())
	}

	return magnitude, nil
}

// GetNumericValue - Returns the 'integerNum' and 'precision' values for the
// current BigIntFixedDecimal instance.
func (bigIFd *BigIntFixedDecimal) GetNumericValue() (*big.Int, uint) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	return big.NewInt(0).Set(bigIFd.integerNum), bigIFd.precision
}

// GetNumericSeparatorsDto
//
// Returns a structure containing the character or rune values
// for decimal place separator, thousands separator and
// currency symbol.
func (bigIFd *BigIntFixedDecimal) GetNumericSeparatorsDto() (NumericSeparatorDto, error) {

	numSeps := NumericSeparatorDto{}

	numSeps.DecimalSeparator = bigIFd.decimalSeparator

	numSeps.ThousandsSeparator = bigIFd.thousandsSeparator

	numSeps.CurrencySymbol = bigIFd.currencySymbol

	return numSeps, nil

}

// GetNumStr - Converts the current BigIntFixedDecimal value to
// string of numbers which includes the decimal place and decimal
// digits if they exist.
func (bigIFd *BigIntFixedDecimal) GetNumStr() (string, error) {

	return bigIFd.FormatNumStr(LEADMINUSNEGVALFMTMODE), nil
}

// GetNumStrDto
//
// Converts the current BigIntFixedDecimal value to a NumStrDto instance.
// The resulting number string includes the decimal place and
// decimal digits if they exist.
//
// The returned NumStrDto type contains numeric separators (decimal
// separator, thousands separator, and currency symbol) copied from
// the current BigIntNum instance.
//
// This method performs a validity test on the current BigIntNum instance.
func (bigIFd *BigIntFixedDecimal) GetNumStrDto() (NumStrDto, error) {

	ePrefix := "BigIntFixedDecimal.GetNumStrDto()"

	err := bigIFd.IsValid(ePrefix + "Testing current instance of BigIntFixedDecimal (bigIFd).")

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = nDto.SetNumericSeparatorsDto(numSepDto)",
				ErrContext: "The 'bigIFd' FAILED the Validation Test.",
				ErrMessage: "The current instance of BigIntFixedDecimal ('bigIFd') is INVALID!",
			}
	}

	nDto, err := new(NumStrDto).NewBigInt(big.NewInt(0).Set(bigIFd.integerNum), bigIFd.precision)

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "nDto, err := new(NumStrDto).NewBigInt(big.NewInt(0).\n" +
					"Set(bigIFd.integerNum), bigIFd.precision)",
				ErrContext: fmt.Sprintf("bigIFd.integerNum='%v'\nbigIFd.precision='%v'",
					bigIFd.integerNum.Text(10), bigIFd.precision),
				ErrMessage: err.Error(),
			}
	}

	nDto.decimalSeparator = bigIFd.decimalSeparator

	nDto.thousandsSeparator = bigIFd.thousandsSeparator

	nDto.currencySymbol = bigIFd.currencySymbol

	return nDto, nil
}

// GetPrecision - Returns the 'precision' value for the current
// BigIntFixedDecimal instance. 'precision' specifies the number
// of digits to the right of the decimal place in the
// BigIntFixedDecimal.integerNum.
func (bigIFd *BigIntFixedDecimal) GetPrecision() (uint, error) {

	return bigIFd.precision, nil
}

// GetPrecisionUint - Returns the 'precision' value for the current
// BigIntFixedDecimal instance. 'precision' specifies the number
// of digits to the right of the decimal place in the
// BigIntFixedDecimal.integerNum.
//
// This method is required in order to implement the INumMgr
// interface.
func (bigIFd *BigIntFixedDecimal) GetPrecisionUint() (uint, error) {

	return bigIFd.precision, nil
}

// GetPrecisionBigInt - Returns the 'precision' value for the current
// BigIntFixedDecimal instance. 'precision' specifies the number
// of digits to the right of the decimal place in the
// BigIntFixedDecimal.integerNum. For this method, precision is
// returned as a *big.Int type.
func (bigIFd *BigIntFixedDecimal) GetPrecisionBigInt() *big.Int {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	return big.NewInt(0).SetUint64(uint64(bigIFd.precision))
}

// GetSign - Returns the numeric sign associated
// with the current numeric value encapsulated by
// the current BigIntFixedDecimal instance.
//
// GetSign() returns:
//
//	-1 if x < 0;
//	 0 if x == 0;
//	+1 if x > 0.
func (bigIFd *BigIntFixedDecimal) GetSign() (int, error) {

	return bigIFd.integerNum.Sign(), nil
}

// Inverse - Converts the current BigIntFixedDecimal
// to the inverseBigIntNum of its numeric value.
//
// Example:
// Current Value = '2'         Inverse= '1/2'
//
// Note: When called this method will destroy and
// overwrite the previous numeric value with the
// inverseBigIntNum value.
//
// Input Parameter
// ===============
//
// maxPrecision uint	- Defines the maximum precision for the
//
//	                     inverseBigIntNum value computed by this method.
//	                     As used here, 'maxPrecision' specifies
//	                     the maximum number of numeric digits to
//												the right of the decimal place.
func (bigIFd *BigIntFixedDecimal) Inverse(maxPrecision uint) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
		return
	}

	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		return
	}

	bigOne := big.NewInt(1)

	inverseBigInt, inversePrecision, _ :=
		BigIntMathDivide{}.BigIntFracQuotient(
			bigOne,
			big.NewInt(0),
			bigIFd.integerNum,
			big.NewInt(0).SetUint64(uint64(bigIFd.precision)),
			big.NewInt(0).SetUint64(uint64(maxPrecision)))

	bigIFd.SetNumericValue(inverseBigInt, uint(inversePrecision.Uint64()))

	return
}

// IsInteger - Returns true if the numeric value of
// the current BigIntFixedDecimal is an integer value.
func (bigIFd *BigIntFixedDecimal) IsInteger() bool {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
		return true
	}

	if bigIFd.precision == 0 {
		return true
	}

	return false
}

// IsEven - returns 'true' if the numeric value
// of the current BigIntFixedDecimal is even.
//
// ------------------------------------------------------
// "In mathematics, parity is the property of an
// integer's inclusion in one of two categories:
// even or odd. An integer is even if it is evenly
// divisible by two and odd if it is not even."
//
// "Examples of even numbers include −4, 0, 82 and 178."
// In particular, zero is an even number."
// ------------------------------------------------------
// https://en.wikipedia.org/wiki/Parity_(mathematics)
//
// Also, see 	https://www.mathsisfun.com/definitions/even-number.html
func (bigIFd *BigIntFixedDecimal) IsEven() bool {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
		return true
	}

	if bigIFd.precision > 0 {
		//BigIntFixedDecimal is NOT an integer
		return false
	}

	bigZero := big.NewInt(0)

	if bigIFd.integerNum.Cmp(bigZero) == 0 {
		return true
	}

	remainder := big.NewInt(0).Rem(bigIFd.integerNum, big.NewInt(2))

	if remainder.Cmp(bigZero) == 0 {
		return true
	}

	return false
}

// IsValid
//
// This method performs validity testing on the
// current instance of BigIntFixedDecimal.
//
// If this BigIntFixedDecimal instance fails the
// validity tests, an error will be returned.
//
// If this BigIntFixedDecimal instance passes
// all validity tests, an error value of 'nil'
// will be returned.
func (bigIFd *BigIntFixedDecimal) IsValid(callingMethodName string) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	if len(callingMethodName) > 0 {
		callingMethodName = "BigIntFixedDecimal.IsValid" + "\n" + callingMethodName
	} else {
		callingMethodName = "BigIntFixedDecimal.IsValid"
	}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		callingMethodName,
		"")

	if err != nil {
		return err
	}

	return new(bigIntFixedDecAtom).isBigIntFxDecValid(
		bigIFd,
		ePrefix)
}

// IsZero
//
// Returns true only if the current BigIntFixedDecimal numeric
// value is equal to zero. If the current BigIntFixedDecimal
// is invalid, an error will be returned.
func (bigIFd *BigIntFixedDecimal) IsZero() (bool, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntFixedDecimal.IsZero",
		"")

	if err != nil {
		return false, err
	}

	err = new(bigIntFixedDecAtom).isBigIntFxDecValid(
		bigIFd,
		ePrefix.XCpy("Testing Validity of 'bigIFd'"))

	if err != nil {
		return false, err
	}

	return new(bigIntFixedDecMolecule).isBigIntFxDecZero(
		bigIFd,
		ePrefix.XCpy("Is 'bigIFd' Zero"))
}

// MultiplyByTenToPower - Multiplies the numeric value of the current
// BigIntFixedDecimal by 10 to the power of 'exponent'.
//
//	result = BigIntFixedDecimal x 10^exponent
//
// Examples:
// =========
//
//		BigIntFixedDecimal
//				Value									Exponent							   Result
//	 ------------------        --------              -----------------
//	  105.6752										 0										      105.6752
//	  105.6752                    1									       1056.752
//	  105.6752                    2                        10567.52
//	  105.6752                    3                       105675.2
//	  105.6752                    8                  10567520000
//
// Input Parameter
// ===============
//
// exponent	uint	- The value of the current BigIntFixedDecimal
//
//	instance will be multiplied by ten raised to
//	the power of 'exponent'.
//
// Note:	(1)	This method will delete trailing fractional zeros from
//
//	 			the returned product.
//
//		(2)	This method will destroy and overwrite the previous value
//				of the current BigIntFixedDecimal instance with the results
//				of this calculation.
func (bigIFd *BigIntFixedDecimal) MultiplyByTenToPower(exponent uint) error {

	ePrefix := "BigIntFixedDecimal.MultiplyByTenToPower()"

	err := bigIFd.IsValid(ePrefix + " Testing 'bigIFd' current instance of 'BigIntFixedDecimal'")

	if err != nil {
		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "ePrefix + \" Testing 'bigIFd' current instance of 'BigIntFixedDecimal'",
			ErrContext: "Testing validity of 'bigIFd', an instance of 'BigIntFixedDecimal'",
			ErrMessage: "'bigIFd' FAILED Validation Testing and is therefore INVALID",
		}
	}

	if bigIFd.integerNum.Cmp(big.NewInt(0)) == 0 {
		bigIFd.precision = 0
		return nil
	}

	scale :=
		big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(exponent)), nil)

	factor := new(BigIntFixedDecimal).New(scale, 0)

	result := new(BigIntMathMultiply).FixedDecimalMultiply(bigIFd.CopyOut(), factor)

	result.TrimTrailingFracZeros()

	err = bigIFd.CopyIn(result)

	if err != nil {
		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "err := bigIFd.CopyIn(result)",
			ErrContext: "Testing validity of 'bigIFd', an instance of 'BigIntFixedDecimal'",
			ErrMessage: "'bigIFd' FAILED Validation Testing and is therefore INVALID",
		}
	}

	return nil
}

// MultiplyByTwoToPower - Multiplies the numeric value of the current
// BigIntFixedDecimal by 2 to the power of 'exponent'.
//
//	product = BigIntFixedDecimal x 2^exponent
//
// When the calculation is completed, the value of 'product' will
// replace the old value of the current BigIntFixedDecimal instance.
//
// Examples:
// =========
//
//	 BigIntFixedDecimal 			 	exponent		  product
//	interNum		Precision
//
// -------------------------------------------------------------
//
//	12345						5								15				4045.2096
//								(0.12345 x 2^15 = 4045.2096)
//
// -------------------------------------------------------------
//
//	571						1								 8			 14617.6
//	            (57.1 x 2^8 = 14617.6)
//
// -------------------------------------------------------------
//
// Note:	(1)	This method will delete trailing fractional zeros from
//
//	 			the returned product.
//
//		(2)	This method will destroy and overwrite the previous value
//				of the current BigIntFixedDecimal instance with the results
//				of this calculation.
func (bigIFd *BigIntFixedDecimal) MultiplyByTwoToPower(exponent uint) error {

	ePrefix := "BigIntFixedDecimal.MultiplyByTwoToPower"

	var err error

	err = bigIFd.IsValid(ePrefix + " Testing 'bigIFd'")

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "err = bigIFd.IsValid(ePrefix + \" Testing 'bigIFd'\")",
			ErrContext: "Testing validity of current BigIntFixedDecimal instance 'bigIFd'.\n" +
				"'bigIFd' FAILED validity testing and is INVALID!",
			ErrMessage: err.Error(),
		}
	}

	precision := big.NewInt(0)

	bigIFd.integerNum, precision, err =
		new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(
			bigIFd.integerNum,
			big.NewInt(0).SetUint64(uint64(bigIFd.precision)),
			exponent)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix,
			ReturnFunc: "bigIFd.integerNum, precision, err =\n" +
				"    new(BigIntMathMultiply).BigIntMultiplyByTwoToPower(\n" +
				"    bigIFd.integerNum,\n" +
				"    big.NewInt(0).SetUint64(uint64(bigIFd.precision))",
			ErrContext: "Testing validity of current BigIntFixedDecimal instance 'bigIFd'.",
			ErrMessage: err.Error(),
		}
	}

	bigIFd.precision = uint(precision.Uint64())

	return nil
}

// New - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//	decimal separator = '.'
//	thousands separator = ','
//	currency separator = '$'
//
// Input Parameters
// ================
//
// integer	*big.Int	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
//
//	in input parameter, 'integer'.
func (bigIFd *BigIntFixedDecimal) New(integer *big.Int, precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(integer, precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
}

// NewBigIntPrecision - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
//
// Input Parameters
// ================
//
// bigInt			*big.Int	- Specifies the sequence of numerical digits in the numeric value.
//
// precision	*big.Int	- Specifies the number of digits to the right of the decimal point
//
//													in input parameter, 'integer'. If 'precision' is greater than the
//	                       maximum value of an unsigned integer (+4,294,967,295, which equals
//													2^32 − 1), an error will be triggered. Also, if 'precision' is less
//	                       than zero, an error will be triggered.
func (bigIFd *BigIntFixedDecimal) NewBigIntPrecision(
	bigInt, precision *big.Int) (BigIntFixedDecimal, error) {

	ePrefix := "BigIntFixedDecimal.NewBigIntPrecision() "

	if precision.Cmp(big.NewInt(0)) == -1 {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "result, err := new(BigIntMathMultiply).MultiplyBigIntNumByThree(bNum2)",
				ErrContext: fmt.Sprintf("precision= '%v'", precision.Text(10)),
				ErrMessage: "Error: Input parameter 'precision' LESS THAN ZERO!",
			}
	}

	maxUint32 := big.NewInt(0).SetUint64(uint64(math.MaxUint32))

	if precision.Cmp(maxUint32) == 1 {

		return BigIntFixedDecimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "",
				ErrContext: " if precision.Cmp(maxUint32) == 1 {",
				ErrMessage: "Error: Input parameter 'precision' exceeds maximum limit of '4,294,967,295'!\n" +
					fmt.Sprintf("precision= '%v'", precision.Text(10)),
			}
	}

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(bigInt, uint(precision.Uint64()))

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num, nil
}

// NewInt - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
//
// Input Parameters
// ================
//
// integer	int				- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
//
//	in input parameter, 'integer'.
func (bigIFd *BigIntFixedDecimal) NewInt(integer int, precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(big.NewInt(int64(integer)), precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
}

// NewInt32 - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
//
// Input Parameters
// ================
//
// integer		 int32	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
//
//	in input parameter, 'integer'.
func (bigIFd *BigIntFixedDecimal) NewInt32(integer int32, precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(big.NewInt(int64(integer)), precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
}

// NewInt64 - Creates and returns a new BigIntFixedDecimal type based on input parameters,
// 'integer' and 'precision'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
//
// Input Parameters
// ================
//
// integer		 int32	- Specifies the sequence of numerical digits in the numeric value.
//
// precision		uint	- Specifies the number of digits to the right of the decimal point
//
//	in input parameter, 'integer'.
func (bigIFd *BigIntFixedDecimal) NewInt64(integer int64, precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(big.NewInt(integer), precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
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
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
func (bigIFd *BigIntFixedDecimal) NewNumStr(numStr string) (BigIntFixedDecimal, error) {

	ePrefix := "BigIntFixedDecimal.NewNumStr() "

	fixedDecimal := new(BigIntFixedDecimal)

	err := fixedDecimal.SetNumStr(numStr)

	if err != nil {
		return BigIntFixedDecimal{},
			fmt.Errorf(ePrefix+"Error returned by fixedDecimal.SetNumStr(numStr). "+
				"Error='%v' ", err.Error())
	}

	_ = fixedDecimal.SetNumericSeparatorsToDefaultIfEmpty()

	return *fixedDecimal, nil
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
//
//	in input parameter, 'integer'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
func (bigIFd *BigIntFixedDecimal) NewUInt(integer uint, precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(big.NewInt(0).SetUint64(uint64(integer)), precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
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
//
//	in input parameter, 'integer'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
func (bigIFd *BigIntFixedDecimal) NewUInt32(integer uint32, precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(big.NewInt(0).SetUint64(uint64(integer)), precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
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
//
//	in input parameter, 'integer'.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
func (bigIFd *BigIntFixedDecimal) NewUInt64(integer uint64, precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(big.NewInt(0).SetUint64(integer), precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
}

// NewZero - Creates and returns a new BigIntFixedDecimal type with a zero value. The
// number of digits to the right of the decimal place is specified by input parameter,
// precision.
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//			decimal separator = '.'
//			thousands separator = ','
//	   currency separator = '$'
//
// Input Parameters
// ================
//
// precision		uint	- Specifies the number of digits to the right of the decimal point.
func (bigIFd *BigIntFixedDecimal) NewZero(precision uint) BigIntFixedDecimal {

	num := new(BigIntFixedDecimal)

	num.SetNumericValue(big.NewInt(0), precision)

	_ = num.SetNumericSeparatorsToDefaultIfEmpty()

	return *num
}

// RoundToDecPlace - Rounds the numeric value of the current BigIntFixedDecimal
// instance to a specified number of decimal places.
//
// 'precision' equals the number of digits to the right of the decimal
// place.
//
// Example:
//
//	integer= 123456; precision = 3; Numeric Value= 123.456
//
// If the value of BigIntFixedDecimal.bigInt is zero ('0'), that zero value will
// remain unaltered. However, the BigIntNum.precision value will be set equal to
// input parameter, 'precision'.
//
// If the number of decimal places specified for rounding ('precision') is
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
//		 Original       				'precision'				Resulting
//	   Value								input parameter			  Value
//	 --------------					---------------     -------------
//		654.123456									9							 654.123456000
//		654.123456									4							 654.1235
//
//	 -654.123456									9							-654.123456000
//	 -654.123456									4							-654.1235
//
//		 0													3								 0.000
//		 0.000000										0								 0
//
// Note: This method does NOT trim or delete trailing fractional zero
// digits.
func (bigIFd *BigIntFixedDecimal) RoundToDecPlace(
	precision uint) error {

	ePrefix := "BigIntFixedDecimal.RoundToDecPlace()"

	if bigIFd.integerNum == nil {

		bigIFd.SetNumericValue(big.NewInt(0), precision)

		return nil
	}

	if bigIFd.precision == precision {
		// Nothing to do. Specified 'precision' is already implemented.
		return nil
	}

	cmpToZeroResult := bigIFd.integerNum.Cmp(big.NewInt(0))

	// bigInt == zero, set precision and return
	if cmpToZeroResult == 0 {

		err := bigIFd.CopyIn(new(BigIntFixedDecimal).NewZero(precision))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err := bigIFd.CopyIn(new(BigIntFixedDecimal).NewZero(precision))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

		return nil
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

		return nil
	}

	// Must be: bigIFd.precision >  precision

	bigNumRound5 :=
		new(BigIntFixedDecimal).NewInt(5, precision+1)

	if cmpToZeroResult == -1 {

		bigNumRound5.integerNum.Mul(
			bigNumRound5.integerNum,
			big.NewInt(-1))
	}

	result, err := new(BigIntMathAdd).FixedDecimalAdd(bigIFd.CopyOut(), bigNumRound5)

	if err != nil {
		return fmt.Errorf("%v\n"+
			"Error returned by BigIntMathAdd{}.FixedDecimalAdd()\n"+
			"Error= %v\n",
			ePrefix,
			err.Error())
	}

	// 10^deltaPrecision
	scale.Exp(big.NewInt(10),
		big.NewInt(int64(bigIFd.precision-precision)), nil)

	result.integerNum.Quo(result.integerNum, scale)

	bigIFd.SetNumericValue(result.integerNum, precision)

	return nil
}

// SetDecimalSeparator
//
// Assigns a rune or character to the internal data field,
// 'decimalSeparator'. The Decimal Separator is used to
// separate the integer and fractional elements of a number
// string.
//
// The BigIntFixedDecimal Type uses this character when
// generating number strings for display.
//
// In the USA, the Decimal Separator is a period character
// ('.').
//
// Note: If a zero value is submitted as input, the Decimal Separator
// will default to the USA standard period character ('.').
//
// Example: 123.456
func (bigIFd *BigIntFixedDecimal) SetDecimalSeparator(decimalSeparator rune) {

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	bigIFd.decimalSeparator = decimalSeparator
}

// SetCurrencySymbol
//
// Assigns the input parameter rune as the currency symbol to
// be used by the BigIntNum type when generating number
// strings for display.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
// Note: If a zero value is submitted as input, Currency Symbol
// will default to the USA dollar sign ('$').
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// Example: $123.45
func (bigIFd *BigIntFixedDecimal) SetCurrencySymbol(currencySymbol rune) {

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	bigIFd.currencySymbol = currencySymbol
}

// SetIntegerValue - Sets the BigIntFixedDecimal.integerNum or integer value
// for the current BigIntFixedDecimal instance.
func (bigIFd *BigIntFixedDecimal) SetIntegerValue(integer *big.Int) {

	if integer == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	} else {
		bigIFd.integerNum = big.NewInt(0).Set(integer)
		bigIFd.precision = 0
	}

}

// SetNumericSeparators
//
// Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current number string.
//
// Different nations and cultures use different symbols to delimit numerical
// values. In the USA and many other countries, a period character ('.') is
// used to delimit integer and fractional digits within a numeric value
// (123.45). Likewise, thousands may be delimited by a comma (','). Currency
// signs very by nationality. For instance, the USA, Canada and several other
// countries use the dollar sign ($) as a currency symbol.
//
// For a list of major world currency symbols see:
//
//		MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//	 http://www.xe.com/symbols.php
//
// Note: If zero values are submitted as input for separator values, those values will default
// to USA standards.
//
// USA Examples:
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
func (bigIFd *BigIntFixedDecimal) SetNumericSeparators(
	decimalSeparator rune,
	thousandsSeparator rune,
	currencySymbol rune) error {

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	bigIFd.decimalSeparator = decimalSeparator
	bigIFd.thousandsSeparator = thousandsSeparator
	bigIFd.currencySymbol = currencySymbol

	return nil
}

// SetNumericSeparatorsDto
//
// Sets the values of numeric separators:
//
//	decimal place separator
//	thousands separator
//	currency symbol
//
// These numeric separators are configured based on values
// transmitted through input parameter 'customSeparators'.
//
// If any of the values contained in input parameter 'customSeparators'
// is set to zero, an error will be returned.
func (bigIFd *BigIntFixedDecimal) SetNumericSeparatorsDto(
	customSeparators NumericSeparatorDto) error {

	ePrefix := "BigIntFixedDecimal.SetNumericSeparatorsDto()"

	if customSeparators.DecimalSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "",
			ErrContext: " if customSeparators.DecimalSeparator == 0 {",
			ErrMessage: "Error: Input parameter 'customSeparators.DecimalSeparator' is set to '0' - Invalid rune!",
		}
	}

	if customSeparators.ThousandsSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "",
			ErrContext: " if customSeparators.ThousandsSeparator == 0 {",
			ErrMessage: "Error: Input parameter 'customSeparators.ThousandsSeparator' is set to '0' - Invalid rune!",
		}
	}

	if customSeparators.CurrencySymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix,
			ReturnFunc: "",
			ErrContext: " if customSeparators.CurrencySymbol == 0 {",
			ErrMessage: "Error: Input parameter 'customSeparators.CurrencySymbol' is set to '0' - Invalid rune!",
		}

	}

	bigIFd.decimalSeparator = customSeparators.DecimalSeparator

	bigIFd.thousandsSeparator = customSeparators.ThousandsSeparator

	bigIFd.currencySymbol = customSeparators.CurrencySymbol

	return nil

}

// SetNumericSeparatorsToDefaultIfEmpty
//
// If numeric separators were previously set to zero or nil,
// this method will set those numeric separators to the USA
// defaults. This means that the Decimal separator is set
// to ('.'), the Thousands separator is set to (',') and the
// currency symbol is set to '$'.
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that all numeric separators
// are set to valid values.
func (bigIFd *BigIntFixedDecimal) SetNumericSeparatorsToDefaultIfEmpty() error {

	if bigIFd.decimalSeparator == 0 {
		bigIFd.decimalSeparator = '.'
	}

	if bigIFd.thousandsSeparator == 0 {
		bigIFd.thousandsSeparator = ','
	}

	if bigIFd.currencySymbol == 0 {
		bigIFd.currencySymbol = '$'
	}

	return nil
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
func (bigIFd *BigIntFixedDecimal) SetNumStr(numStr string) error {

	ePrefix := "BigIntFixedDecimal.SetNumStr() "

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	if len(numStr) == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStr' is an EMPTY string!\n",
			ePrefix)
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

		if baseRunes[i] == '-' && isStartNumericDigits == false {
			hasMinusSign = true
			continue
		}

		if baseRunes[i] == '(' && isStartNumericDigits == false {
			hasLeftParen = true
			continue
		}

		if baseRunes[i] == ')' &&
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

			newAbsBigInt.Mul(newAbsBigInt, baseTen)

			newAbsBigInt.Add(newAbsBigInt, big.NewInt(int64(baseRunes[i]-48)))

			numOfNumericDigits++

			if startFractionalDigits == true {
				newPrecision++
			}

			continue
		}

	}

	if numOfNumericDigits == 0 {
		return fmt.Errorf(ePrefix+
			"Error: No numeric digits were found in input parameter 'numStr'. "+
			"numStr='%v'", numStr)
	}

	if hasMinusSign == true || (hasLeftParen == true && hasRightParen == true) {

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
func (bigIFd *BigIntFixedDecimal) SetNumericValue(integer *big.Int, precision uint) {

	if integer == nil {
		bigIFd.integerNum = big.NewInt(0)
	} else {
		bigIFd.integerNum = big.NewInt(0).Set(integer)
	}

	bigIFd.precision = precision

}

// SetPrecisionValue - Sets the 'precision' value for the current BigIntFixedDecimal
// instance. 'precision' specifies the number of fractional digits to the right
// of the decimal place.
func (bigIFd *BigIntFixedDecimal) SetPrecisionValue(precision uint) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
	}

	bigIFd.precision = precision

}

// SetThousandsSeparator
//
// Sets the value of the character which will be used to separate
// thousands in the display of the NumStrDto number string. In the
// USA the typical thousands separator is the comma (',').
//
// If a zero value is submitted, the Thousands Separator will
// default to the comma character (USA Default).
//
// Example:
// 1,000,000
func (bigIFd *BigIntFixedDecimal) SetThousandsSeparator(thousandsSeparator rune) {

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	bigIFd.thousandsSeparator = thousandsSeparator

	return
}

// TrimTrailingFracZeros - This method will delete non-significant
// trailing zeros from the fractional digits of the current
// BigIntFixedDecimal numerical value.
//
// Examples:
//
//	Initial Value			Trimmed Value
//		456.123000 			 456.123
//			0.000					 0
//			7.0						 7
//	 -456.123000			-456.123
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
// If the number of decimal places specified for truncation ('precision') is
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
//		Original						'newPrecision'				Resulting
//		Value								input parameter					Value
//	 --------------				---------------     -------------
//		654.123456								9								654.123456000
//		654.123456								4								654.1234 (no rounding)
//
//	 -654.123456								9							 -654.123456000
//	 -654.123456								4							 -654.1234 (no rounding)
//
//			0												3									0.000
//			0.000000								0									0
func (bigIFd *BigIntFixedDecimal) TruncToDecPlace(precision uint) {

	if bigIFd.integerNum == nil {
		bigIFd.integerNum = big.NewInt(0)
		bigIFd.precision = 0
	}

	if bigIFd.precision == precision {
		// Nothing to do. Specified 'precision' is already implemented.
		return
	}

	// bigInt == zero, set precision and return
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
