package mathops

import (
	"fmt"
	"math/big"
			"errors"
)

/*
	Decimal
	=======

	The source code repository for decimal.go is located at:
			https://github.com/MikeAustin71/mathopsgo.git

	The source file decimal.go is located in directory:
		MikeAustin71/mathopsgo/mathops/decimal.go

*/

// Decimal - This type is used to perform math operations
// which achieve a high degree of accuracy when dealing
// with fractional numbers containing digits to the right
// of the decimal place.
//
// While storage operations are provided by a type NumStrDto,
// math operations are performed using types *big.Int and
// *big.Rat.
//
// The Decimal Type implements the INumMgr interface.
//
type Decimal struct {
	bigINum BigIntNum
}

// Add - Adds the value of the current Decimal to that of
// the incoming Decimal and returns in the result in a
// Decimal Type.
//
// Note that Numeric separators remain unchanged and are
// are set to the values of the current Decimal instance.
//
func (dec *Decimal) Add(d2 Decimal) (Decimal, error) {

	ePrefix := "Decimal.Add() "
	var err error

	err = dec.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
		fmt.Errorf(ePrefix + "The current host Decimal object is INVALID! " +
			"Error='%v' \n", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	err = d2.IsDecimalValid()

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Input Parmeter Decimal (d2) is INVALID! " +
			"Error= '%v' \n", err.Error())
	}

	bINumResult := BigIntMathAdd{}.AddBigIntNums(dec.bigINum, d2.bigINum)

	d3 := Decimal{}.NewBigIntNum(bINumResult)

	err = d3.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix + "The Decimal Type resulting from Addition is INVALID! " +
				"Error='%v' \n", err.Error())
	}

	err = d3.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by d3.SetNumericSeparatorsDto(numSeps). " +
				"Error='%v' \n", err.Error())
	}

	return d3, nil
}

// AddToThis - adds the value of the incoming Decimal to that
// of the current Decimal object. The new total is retained
// in the current Decimal object.
//
func (dec *Decimal) AddToThis(d2 Decimal) error {

	numSeps := dec.GetNumericSeparatorsDto()

	dec.bigINum = BigIntMathAdd{}.AddBigIntNums(dec.bigINum, d2.bigINum)

	err := dec.bigINum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		ePrefix := "Decimal.AddToThis() "
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetNumericSeparatorsDto(numSeps) " +
			"Error='%v' \n", err.Error())
	}

	return nil
}

// AddToThisArray - Receives an array of Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisArray(decs []Decimal) error {

	ePrefix := "Decimal.AddToThisArray() "
	var err error

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: The current Decimal instance (dec) is INVALID! " +
			"Error='%v' ", err.Error())
	}

	if len(decs) == 0 {
		return errors.New(ePrefix +
			"Error: Input Array 'decs' is EMPTY!")
	}

	numSeps := dec.GetNumericSeparatorsDto()


	bINumResult := dec.bigINum.CopyOut()

	for i, dx := range decs {

		err = dx.IsDecimalValid()

		if err != nil {

			return fmt.Errorf(ePrefix +
				"Error: Array element decs[%v] is INVALID! Error='%v'",
					i, err.Error())
		}

		bINumResult = BigIntMathAdd{}.AddBigIntNums(bINumResult, dx.bigINum)

	}

	bINumResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by bINumResult.SetNumericSeparatorsDto(numSeps) " +
			"Error='%v' \n", err.Error())
	}

	dec.SetBigIntNum(bINumResult)

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: Decimal type resulting from Array Addition is INVALID! " +
			"Error='%v' ", err.Error())
	}

	return nil
}

// AddToThisMultiple - Receives multiple Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisMultiple(decs ...Decimal) error {

	ePrefix := "Decimal.AddToThisMultiple() "
	var err error

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: The current Decimal instance (dec) is INVALID! " +
			"Error='%v' ", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	bINumResult := dec.bigINum.CopyOut()


	for i, dx := range decs {

		err = dx.IsDecimalValid()

		if err != nil {

			return fmt.Errorf(ePrefix +
				"Error: Series element decs[%v] is INVALID! Error='%v'",
				i, err.Error())
		}

		bINumResult = BigIntMathAdd{}.AddBigIntNums(bINumResult, dx.bigINum)
	}

	bINumResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by bINumResult.SetNumericSeparatorsDto(numSeps) " +
			"Error='%v' \n", err.Error())
	}

	dec.SetBigIntNum(bINumResult)

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error: Decimal type resulting from Array Addition is INVALID! " +
			"Error='%v' ", err.Error())
	}

	return nil
}

// AllDigitsNumStr - parses the incoming string and returns
// a pure number string consisting of all numeric digits. No
// sign characters, decimals or thousands separators are returned.
// The returned value is the absolute numeric value extracted from
// the 'numStr' input parameter.
func (dec *Decimal) AllDigitsNumStr(numStr string) (string, error) {


	bigIntNum, err := BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		ePrefix := "Decimal.AllDigitsNumStr() "
		return "",
		fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStr(numStr) " +
			"numStr='%v' Error='%v' \n", numStr, err.Error())
	}

	return bigIntNum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE), nil
}

// NumStrToDecimal - Creates a Decimal type from a number
// string.
func (dec *Decimal) NumStrToDecimal(numStr string) (Decimal, error) {

	d2 := Decimal{}

	var err error

	d2.bigINum, err = BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		ePrefix := "Decimal.NumStrToDecimal() "
		return Decimal{}.NewZero(0),
		fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStr(numStr). " +
			"numStr='%v' Error='%v' ", numStr, err.Error())
	}


	return d2, nil
}

// CopyIn - Receives an incoming Decimal object
// as an input parameter and sets the Current Decimal
// equal to that of the incoming Decimal object.
func (dec *Decimal) CopyIn(d2 Decimal) {

	dec.Empty()
	dec.bigINum = d2.bigINum.CopyOut()
	return 
}

// CopyOut - Returns a deep copy of the current Decimal
// instance.
//
func (dec *Decimal) CopyOut() Decimal {
	d2 := Decimal{}.New()
	d2.SetBigIntNum(dec.bigINum.CopyOut())
	return d2
}

// Divide - Divides the current decimal value by the input
// parameter 'divisor' and returns the quotient as a new
// Decimal instance.
//
func (dec *Decimal) Divide(divisor Decimal, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.Divide() "
	var err error

	err = dec.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
			"Error: The current Decimal Instance (dec) is INVALID! " +
				"Error='%v' \n", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	err = divisor.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error: Input parameter 'divisor' is INVALID! " +
				"Error='%v' \n", err.Error())
	}

  bINumQuotient, err :=
  		BigIntMathDivide{}.BigIntNumFracQuotient(
  				dec.bigINum, divisor.bigINum, maxPrecision)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by BigIntMathDivide{}.BigIntNumFracQuotient() " +
				"Error='%v' ", err.Error())
	}

	d2Quotient := Decimal{}.NewBigIntNum(bINumQuotient)

	err = d2Quotient.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error: The quotient Decimal Instance (d2) resulting from Division is INVALID! " +
				"Error='%v' \n", err.Error())
	}

	err = d2Quotient.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
		fmt.Errorf(ePrefix +
			"Error returned by d2Quotient.SetNumericSeparatorsDto(numSeps) " +
			"Error='%v' \n", err.Error())
	}

	return d2Quotient, nil
}

// Equal - Returns true if the input Decimal instance is equal
// in all respects to the current Decimal instance.
func (dec *Decimal) Equal(dec2 Decimal) bool {

	return dec.bigINum.Equal(dec2.bigINum)
}

// Empty - Sets all values of the current Decimal's
// fields to their 'zero' values.
func (dec *Decimal) Empty() {
	dec.bigINum =  BigIntNum{}.NewZero(0)
	dec.SetEmptySeparatorsToDefault()
}

// GetAbsoluteBigIntValue - returns the absolute value of the
// decimal expressed as a string. If the decimal value is
// '-123.456', this method will return '123.456'.
func (dec *Decimal) GetAbsoluteValue() Decimal {

	bi2 := dec.bigINum.GetAbsoluteBigIntNumValue()

	return Decimal{}.NewBigIntNum(bi2)
}

// GetAbsoluteAllDigitsStr - Returns the absolute value of the Decimal integer.
// Fractions are not returned, only the string of signed numeric digits which
// constitutes the entire number. In other words, if the value of the decimal
// is '-123.456', this method will return '123456.
func (dec *Decimal) GetAbsoluteAllDigitsStr() (string, error) {
	ePrefix := "Decimal.GetAbsoluteAllDigitsStr() "

	err := dec.IsDecimalValid()

	if err != nil {
		return "",
		fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v'", err.Error())
	}

	return dec.bigINum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE), nil
}

// GetBigFloat - returns big Float representation of the Decimal Value.
func (dec *Decimal) GetBigFloat() (*big.Float, error) {
	ePrefix := "Decimal.GetBigFloat() "

	err := dec.IsDecimalValid()

	if err!=nil {
		return big.NewFloat(0.0),
		fmt.Errorf(ePrefix +
			"The current Decimal object is INVALID! Please re-initialize. " +
			"Error='%v'", err.Error())
	}


	return dec.bigINum.GetBigFloat(), nil
}

// GetBigFloatString - returns a signed number string which is accurate out
// to a large number of decimal places.
//
func (dec *Decimal) GetBigFloatString(precision uint) (string, error) {

	ePrefix := "Decimal.GetBigFloatString() "

	err := dec.IsDecimalValid()

	if err != nil {
		return "",
		fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v' ", err.Error())
	}

	biNum2 := dec.bigINum.CopyOut()

	biNum2.SetPrecision(precision)

	return biNum2.GetNumStr(), nil
}

// GetBigInt - returns the Decimal value expressed as an
// integer value using type *big.Int. No factional values are included.
// For example, the value '-123.456' would be returned as the integer
// value '-123456'.  To compute the precise value of the Decimal, this
// integer value would need to be divided by the 'precision Value'. See
// GetScaleVal() below.
func (dec *Decimal) GetBigInt() (*big.Int, error) {

	ePrefix := "Decimal.GetBigInt() "

	err := dec.IsDecimalValid()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v' ", err.Error())
	}

	bInt, err := dec.bigINum.GetBigInt()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix + "Error returned by dec.bigINum.GetBigInt() " +
				"Error='%v' ", err.Error())
	}

	return bInt, nil
}

// GetBigIntNum - Converts the current Decimal numeric value to
// an instance of type 'BigIntNum' and returns it to the calling
// function.
//
func (dec *Decimal) GetBigIntNum() BigIntNum {

	return dec.bigINum.CopyOut()
}

// GetBigRat - Returns the current Decimal's numeric value expressed
// as a rational number of type *big.Rat.
//
func (dec *Decimal) GetBigRat() *big.Rat {
	return dec.bigINum.GetBigRat()
}

// GetCurrencySymbol - Returns the Decimal's current
// value for Currency Symbol.
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// In the USA, the currency symbol is the dollar sign
// ('$').
//
func (dec *Decimal) GetCurrencySymbol() rune {

	return dec.bigINum.GetCurrencySymbol()
}

// GetCurrencyStr - Returns the Decimal's numeric value expressed
// as number string delimited with the Decimal's Thousands Separator
// and prefixed with the designated Currency Symbol characters.
//
// Note: The file mathopsconstants.go file contains Unicode characters
// for most of the world's major currencies. This file is located at:
// 			MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. These Separators and Currency
// Symbol are variable and may be controlled by the user.
//
// If the numeric value is negative, a leading minus sign will be prefixed
// to the currency display.
//
// Example:
// numstr = 1000000.23
// GetCurrencyStr() = $1,000,000.23
//
// numstr = -1000000.23
// GetCurrencyStr() = -$1,000,000.23
//
// Note: If the current Decimal is invalid, this method
// returns an empty string.
//
func (dec *Decimal) GetCurrencyStr() string {
	return dec.bigINum.FormatCurrencyStr(LEADMINUSNEGVALFMTMODE)
}

// GetCurrencyParen - Returns the Decimal's numeric value expressed
// as number string delimited with the Decimal's Thousands Separator
// and prefixed with the designated Currency Symbol characters.
//
// Note: The file mathopsconstants.go file contains Unicode characters
// for most of the world's major currencies. This file is located at:
// 			MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. These Separators and Currency
// Symbol are variable and may be controlled by the user.
//
// If the numeric value is negative, the resulting number string is
// surrounded by parentheses.
//
// Example:
// numstr = 1000000.23
// GetCurrencyParen() = $1,000,000.23
//
// numstr = -1000000.23
// GetCurrencyParen() = ($1,000,000.23)
//
// Note: If the current Decimal is invalid, this method
// returns an empty string.
//
func (dec *Decimal) GetCurrencyParen() string {
	return dec.bigINum.FormatCurrencyStr(PARENTHESESNEGVALFMTMODE)
}


// GetDecimalSeparator - returns the Decimal's current
// value for Decimal Separator (i.e. '.')
func (dec *Decimal) GetDecimalSeparator() rune {

	return dec.bigINum.GetDecimalSeparator()
}

// GetFloat32 - Returns the current value of the Decimal as a
// float32. There may be a loss of accuracy during this conversion.
// The return parameter big.Accuracy will describe the level of
// accuracy provided.
//
// See big.Accuracy:
//   Below Accuracy = -1
//   Exact Accuracy = 0
//   Above Accuracy = +1
func (dec *Decimal) GetFloat32() (float32, big.Accuracy, error) {

	ePrefix := "Decimal.GetFloat32() "

	err := dec.IsDecimalValid()

	if err != nil  {
		return float32(0), big.Accuracy(0),
		fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v'", err.Error())
	}

	bf, status := big.NewFloat(0.0).SetString(dec.bigINum.GetNumStr())

	if !status {
		return float32(0.0), big.Accuracy(0),
			fmt.Errorf(ePrefix + "SetString() Failed. NumStr= %v", dec.bigINum.GetNumStr())
	}

	f32, accuracy := bf.Float32()

	return f32, accuracy, nil

}

// GetFloat64 - Returns the current value of the Decimal as a
// float64. There may be a loss of accuracy during this conversion.
// The return parameter big.Accuracy will describe the level of
// accuracy provided.
//
// See big.Accuracy:
//   Below Accuracy = -1
//   Exact Accuracy = 0
//   Above Accuracy = +1
//
func (dec *Decimal) GetFloat64() (float64, big.Accuracy, error) {

	ePrefix := "Decimal.GetFloat64() "

	err := dec.IsDecimalValid()

	if err != nil {
		return float64(0), big.Accuracy(0),
		fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v' ", err.Error())
	}

	bf, status := big.NewFloat(0.0).SetString(dec.bigINum.GetNumStr())

	if !status {
		return float64(0.0), big.Accuracy(0),
		fmt.Errorf(ePrefix + "SetString() Failed. NumStr= %v", dec.bigINum.GetNumStr())
	}

	f64, accuracy := bf.Float64()

	return f64, accuracy, nil
}

// GetIntAry - Returns an IntAry structure initialized
// to the value of the current 'Decimal' object.
func (dec *Decimal) GetIntAry() (IntAry, error) {
	
	ePrefix := "Decimal.GetIntAry() "

	err := dec.IsDecimalValid()

	if err != nil {
		return IntAry{}.New(),
			fmt.Errorf(ePrefix + "This Decimal instance is INVALID! " +
				"Error= %v", err.Error())
	}

	ia, err := dec.bigINum.GetIntAry()

	if err != nil {
		return IntAry{}.New(),
		fmt.Errorf( ePrefix + "Error received from " +
			"dec.bigINum.GetIntAry(). Error= %v",
			err.Error())
	}

	return ia, nil
}

// GetNumericSeparatorsDto - returns a NumericSeparatorDto structure
// containing the current characters (runes) used to specify
// decimal point separator, thousands separator and currency symbol.
//
func (dec *Decimal) GetNumericSeparatorsDto() NumericSeparatorDto {

	return dec.bigINum.GetNumericSeparatorsDto()
}

// GetNumStr - Returns the internal value of the Decimal
// expressed as a signed numeric string. precision, or
// placement of the decimal point, is controlled by
// the Decimal's precision setting.
//
// Example Output:
// ===============
//
//  123
//  123.4
//  123456789
//  123456789.44
// -123
// -123.4
// -123456789
// -123456789.44
//
func (dec *Decimal) GetNumStr() string {

	return dec.bigINum.FormatNumStr(LEADMINUSNEGVALFMTMODE)

}

// GetNumPren - Returns the internal value of the
// Decimal expressed as number string. precision
// or placement of the decimal point is controlled
// by the Decimal's 'precision' setting.
//
// If the numeric value is less than zero, a negative
// number, the number string is surrounded in parentheses.
//
// Example Output:
// ===============
//
//  123
//  123.4
// (123)
// (123.4)
//
func (dec *Decimal) GetNumParen() string {
	return dec.bigINum.FormatNumStr(PARENTHESESNEGVALFMTMODE)
}


// GetNumStrDto - returns a NumStrDto structure initialized
// to the value of the current Decimal object.
func (dec *Decimal) GetNumStrDto() (NumStrDto, error) {
	ePrefix := "Decimal.GetNumStrDto() "

	nDto, err := dec.bigINum.GetNumStrDto()

	if err != nil {
		return NumStrDto{},
			fmt.Errorf(ePrefix + "- Error returned from dec.bigINum.GetNumStrDto(). " +
				"dec.bigINum= '%v' Error= %v", dec.bigINum.GetNumStr(), err.Error())
	}

	return nDto, nil
}

// GetNthRoot - Calculates the nth root of the current Decimal value.
//
// Input parameters:
//
//  nthRoot uint - Nth root specifies the root which will be calculated for parameter, 'num'.
// 									Example, square root, cube root, 4th root, 9th root etc.
//
// maxPrecision uint -  Specifies the number of decimals to the right of the decimal place to
// 											which the Nth root will be calculated.
//
// The calculation result is returned as an IntAry object.
//
// Note: A negative Decimal value with an even nthRoot will generate an error.
func (dec *Decimal) GetNthRoot(nthRoot, maxPrecision int) (Decimal, error) {

	ePrefix := "Decimal.GetNthRoot() "

	err := dec.IsDecimalValid()

	if err != nil {
		return Decimal{}.New(),
		fmt.Errorf(ePrefix +
			"- The current Decimal object is INVALID! Please re-initialize." +
			"Error='%v' ", err.Error())
	}

	ia, err := dec.GetIntAry()

	if err != nil {
		return Decimal{}.New(),
		fmt.Errorf(ePrefix + "- Received error from dec.GetIntAry(). Error= %v \n",
			err.Error())
	}

	iaNthRoot, err := IntAry{}.NewInt(nthRoot, 0)

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix +
				"- Error returned by IntAry{}.NewInt(nthRoot, 0). Error= %v \n",
				err.Error())
	}


	nrt := NthRootOp{}

	iaNth, err := nrt.GetNthRootIntAry(&ia, &iaNthRoot, maxPrecision)

	if err != nil {
		return Decimal{}.New(),
		fmt.Errorf(ePrefix +
					"- Received error from  NthRootOp.GetNthRootIntAry(). Error= %v", err)
	}

	dOut, err := dec.MakeDecimalFromIntAry(&iaNth)

	if err != nil {
		return Decimal{}.New(),
		fmt.Errorf(ePrefix + "- Received error from  dec.MakeDecimalFromIntAry(&iaNth). Error= %v", err.Error())
	}

	return dOut, nil
}

// GetPrecision - returns the Decimal's current precision
// value. The Decimal structure maintains precision as an
// unsigned integer.
//
// The value of 'precision' returned by this method will
// always be >= zero (greater than or equal to zero '0').
//
// precision is defined as the number of numeric digits to
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
//		Number String				precision				Fractional Number
//			123456								3								123.456
//
func (dec *Decimal) GetPrecision() int {

	return dec.bigINum.GetPrecision()

}

// GetPrecisionUint - Returns precision as an
// unsigned integer.
//
// precision is defined as the number of numeric digits to
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
//		Number String				precision				Fractional Number
//			123456								3								123.456
//
//
func (dec *Decimal) GetPrecisionUint() uint {

	return uint(dec.bigINum.GetPrecision())
}

// GetRational - returns a big Rational number type which
// is capable of very high accuracy.
//
// The returned *big.Rat number is initialized to the
// current value of the Decimal object.
func (dec *Decimal) GetRational() (*big.Rat, error) {

	ePrefix := "Decimal.GetRational() "

	err := dec.IsDecimalValid()

	if err != nil {
		return big.NewRat(1, 1),
		fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v' ", err.Error())
	}

	x, err := dec.bigINum.GetBigInt()

	if err != nil {
		return big.NewRat(1, 1),
			fmt.Errorf(ePrefix + "Error returned by dec.bigINum.GetBigInt() " +
				"Error='%v' ", err.Error())
	}

	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)
	
	rDividend := big.NewRat(1, 1).SetInt(decSignedAllDigitsBigInt)

	decScaleFactor := dec.bigINum.GetScaleFactor()

	rDivisor := big.NewRat(1, 1).SetInt(decScaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	return rQuotient, nil
}

// GetSign - Returns the sign of the current
// Decimal Value. Return values are one of two
// integers: +1 or -1.
func (dec *Decimal) GetSign() int {

	return dec.bigINum.GetSign()
}

// GetRelevantPrecision - Returns an unsigned integer representing
// the number of decimal places to the right of the decimal which
// are occupied by non-zero values.
//
// Example: Value = 1.640700000.  The number of relevant decimal
// places to the right of the decimal is '4'. In the case of an
// integer number the relevant precision is zero ('0') because
// there are no digits to the right of the decimal.
func (dec *Decimal) GetRelevantPrecision() uint {

	bI2 := dec.bigINum.CopyOut()

	bI2.TrimTrailingFracZeros()

	return bI2.GetPrecisionUint()
}

// GetScaleVal - Returns the scale value associated with this decimal value. The
// scale value is expressed as 10 to an exponent. Example 10^2 == 100.
//
// Scale Value, or Scale Factor, is defined by 10 raised to the power
// of Decimal precision.
//
// The return scale value is of type big integer (*big.Int)
//
func (dec *Decimal) GetScaleVal() (*big.Int, error) {

	return dec.bigINum.GetScaleFactor(), nil
}

// GetSignedAllDigitsStr - Returns the Decimal's internal
// Signed All Digits Integer Value expressed as a string.
// No Fractional Digits are included, this is a signed
// integer number string. Example: The value '-123.456'
// would be returned as '-123456'
func (dec *Decimal) GetSignedAllDigitsStr() (string, error) {

	ePrefix := "Decimal.GetSignedAllDigitsStr() "

	nDto, err := dec.bigINum.GetNumStrDto()

	if err != nil {
		return "",
		fmt.Errorf(ePrefix + "Error returned by dec.bigINum.GetNumStrDto(). " +
			"Error='%v' \n", err.Error())
	}

	result := ""
	
	if dec.bigINum.GetSign() < 0 {
		result += "-"
	}
	
	result += string(nDto.absAllNumRunes)
	
	return result, nil
}

// GetBigInt - Returns the numeric value of the current Decimal
// instance as a signed *big.Int.
func (dec *Decimal) GetSignedBigInt() (*big.Int, error) {
	ePrefix := "Decimal.GetBigInt() "

	bInt, err :=	dec.bigINum.GetBigInt()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix + "Error returned by dec.bigINum.GetBigInt(). " +
				"Error='%v' ", err.Error())
	}

	return big.NewInt(0).Set(bInt), nil
}

// GetSquareRoot - Returns a Decimal object equal to the square root
// of the current Decimal value.
//
// Note: If the current Decimal value is a negative value, an error will
// be generated. You cannot take the square root of a negative number.
func (dec *Decimal) GetSquareRoot(maxPrecision int) (Decimal, error) {

	ePrefix := "Decimal.GetSquareRoot() "

	err := dec.IsDecimalValid()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v' ", err.Error())
	}


	ia, err := dec.GetIntAry()

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix + "- Received error from dec.GetIntAry(). Error= %v", err)
	}

	nrt := NthRootOp{}

	iaSq, err := nrt.GetSquareRootIntAry(&ia, maxPrecision)

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix +
					"- Received error from  NthRootOp.GetSquareRootIntAry(). Error= %v", err)
	}

	dOut, err := dec.MakeDecimalFromIntAry(&iaSq)

	if err != nil {
		return Decimal{}.New(),
		fmt.Errorf(ePrefix +
			"- Received error from  dec.MakeDecimalFromIntAry(&iaSq). Error= %v", err)
	}

	return dOut, nil
}

// GetThisPointer - Returns a pointer to the current Decimal instance
//
func (dec *Decimal) GetThisPointer() *Decimal {

	return dec
}

// GetThousandsSeparator - Gets the current value of
// the Thousands Separator for the current Decimal
// object.
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// For U.S.A. - The thousands separator is a the comma (',')
//
func (dec *Decimal) GetThousandsSeparator() rune {

	return dec.bigINum.GetThousandsSeparator()
}

// GetThouStr - Returns a number string which represents the Decimal's
// numeric value. Thousands are separated by the Decimal's Thousands
// Separator. In the USA, the Thousands Separator is a comma character
// (',').
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// Negative numbers are preceded by a minus sign.
//
// Example Output
// ==============
//
//  123,456,789
//  123,456,789.12
// -123,456,789
// -123,456,789.12
//
func (dec *Decimal) GetThouStr() string {
	return dec.bigINum.FormatThousandsStr(LEADMINUSNEGVALFMTMODE)
}

// GetThouParen - Returns the Decimal's numeric value formatted
// as a number string with thousands separated by the Decimal's
// Thousands Separator.  In the USA, the Thousands Separator is
// a comma character (',').
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// Negative numbers are surrounded with parentheses.
//
// Example Output
// ==============
//
//  123,456,789
//  123,456,789.12
// (123,456,789)
// (123,456,789.12)
//
func (dec *Decimal) GetThouParen() string {
	return dec.bigINum.FormatThousandsStr(PARENTHESESNEGVALFMTMODE)
}

// GetIsValid - returns a boolean indicating
// the current state of the Decimal information.
// If the current Decimal object is VALID, the
// method returns 'true'.
//
// If the current Decimal object is INVALID, the
// method returns 'false'.
//
// Notice that this method relies on
// Decimal.IsDecimalValid() which returns an
// 'error' type.
//
func (dec *Decimal) GetIsValid() bool {

	err := dec.IsDecimalValid()

	if err != nil {
		return  false
	}

	return true
}

// Inverse - Returns the inverse of the current
// Decimal.  Inverse = 1/Current Decimal Value.
//
// Input Parameters:
// maxPrecision uint - determines the number of digits
// 			to the right of the decimal point in the result.
// 			if maxPrecision is less than zero, an error will
// 			be triggered
//
func (dec *Decimal) Inverse(maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.Inverse() "

	bin2, err := dec.bigINum.GetInverse(maxPrecision)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by dec.bigINum.GetInverse(uint(maxPrecision)) " +
				"Error='%v' \n", err.Error())
	}

	d2, err := bin2.GetDecimal()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by bin2.GetDecimal() " +
				"Error='%v' \n", err.Error())
	}

	return d2, nil
}

// IsDecimalValid - Performs an internal diagnostic on the current
// Decimal instance and returns an 'error' if the instance is INVALID.
//
func (dec *Decimal) IsDecimalValid() error {

	if !dec.bigINum.IsValid() {
		return errors.New("Decimal) IsDecimalValid() Error: " +
			"This Decimal instance is INVALID! \n")
	}

	return nil
}

// IsFraction - returns a boolean value. If 'true',
// it signals that the Decimal has digits to the
// right of the decimal place. If 'false', it
// signals that the decimal value is an integer with
// no digits to the right of the decimal place.
func (dec *Decimal) IsFraction() (bool, error) {

	ePrefix := "Decimal.IsFraction() "

	err := dec.IsDecimalValid()

	if err != nil {
		return false,
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v' ", err.Error())
	}

	if dec.bigINum.precision != 0 {
		return true, nil
	}

	return false, nil
}

// IsZero - Returns true if the numeric value of the current
// 'Decimal' instance is zero.
func (dec *Decimal) IsZero() bool {
	return dec.bigINum.IsZero()
}

// MakeDecimalFromNumStrDto - generates a Decimal Type based on string information
// provided by the 'ia' *IntAry input parameter.
func (dec *Decimal) MakeDecimalFromIntAry(ia *IntAry) (Decimal, error) {

	ePrefix := "Decimal.MakeDecimalFromIntAry() "

	var err error

	ia.SetInternalFlags()

	err = ia.IsIntAryValid(ePrefix + "- ")

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix + "IntAry Invalid - Error: %v\n", err.Error())
	}

	d2 := Decimal{}
	d2.bigINum, err = ia.GetBigIntNum()

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix + "Error returned by ia.GetBigIntNum() " +
				"Error: %v\n", err.Error())
	}


	return d2, nil
}

// Mul - Multiplies the incoming Decimal value, by the
// current Decimal Value and returns the result in a
// Decimal Type.
func (dec *Decimal) Mul(d2 Decimal) (Decimal, error) {

	ePrefix := "Decimal.Mul() "

	err := dec.IsDecimalValid()

	if err != nil  {
		return Decimal{},
		fmt.Errorf(ePrefix + "This Decimal object (dec) is INVALID! " +
			"Error='%v' ", err.Error())
	}

	err = d2.IsDecimalValid()

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Incoming Decimal object (d2) is INVALID! " +
			"Error='%v' ", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	bigIntNum := BigIntMathMultiply{}.MultiplyBigIntNums(dec.bigINum, d2.bigINum)

	d4 := Decimal{}.NewBigIntNum(bigIntNum)

	err = d4.bigINum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
		fmt.Errorf(ePrefix +
			"Error returned by d4.bigINum.SetNumericSeparatorsDto(numSeps) " +
			"Error='%v' \n", err.Error())
	}

	return d4, nil
}

// MulThis - Multiplies the value of the incoming
// Decimal type by the value of the current or receiving
// Decimal type. The result of the multiplication is
// stored in the current Decimal type.
//
// The original value of this current Decimal Type is
// therefore destroyed and overwritten.
//
func (dec *Decimal) MulThis(d2 Decimal) error {

	ePrefix := "Decimal.MulThis() "
	err := dec.IsDecimalValid()

	if err != nil  {
		return 	fmt.Errorf(ePrefix + "This Decimal object (dec) is INVALID! " +
				"Error='%v' ", err.Error())
	}

	err = d2.IsDecimalValid()

	if err != nil {
		return 	fmt.Errorf(ePrefix + "Incoming Decimal object (d2) is INVALID! " +
				"Error='%v' ", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	dec.bigINum = BigIntMathMultiply{}.MultiplyBigIntNums(dec.bigINum, d2.bigINum)

	err = dec.bigINum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetNumericSeparatorsDto(numSeps) . " +
			"Error= %v", err.Error())
	}

	return nil
}

// NewBigIntNum - Creates and returns a Decimal type.
// The Decimal value is initialized to zero.
//
// Example Usage:
//   d := Decimal{}.NewBigIntNum()
//
// This is the recommended procedure for creating
// a Decimal type.
func (dec Decimal) New() Decimal {

	d := Decimal{}
	d.Empty()

	return d

}

// NewPtr - Creates and returns a pointer to
// a new Decimal type. Can be used for initializing
// a Decimal Type and calling pointer methods in
// one statement.
//
// Example:
// d1, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(inStr, precision,true)
//
func (dec Decimal) NewPtr() *Decimal {
	d := Decimal{}
	d.Empty()
	return &d
}

// NewBigInt - Returns a Decimal type based on Big Int and
// precision input parameters. If an error is encountered, it
// will trigger a panic condition.
//
// The 'NewBigInt' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewBigInt(bigI, precision)
// bigI := big.NewInt(123456)
// Decimal{}.NewBigInt(bigI, 3) = 123.456
//
func (dec Decimal) NewBigInt(bigI *big.Int, precision uint) (Decimal, error) {

	d2 := Decimal{}.New()

	err := d2.SetBigInt(bigI, precision)

	if err != nil {
		return Decimal{},
		fmt.Errorf("Decimal.NewBigInt() Error returned by " +
			"d2.SetBigInt(bigI, precision) " +
			"bigI='%v' precision='%v' Error='%v'",
				bigI.String(), precision, err.Error())
	}

	return d2, nil
}

// NewBigIntNum - Returns a Decimal instance based on input
// parameter,
func (dec Decimal) NewBigIntNum(bigINum BigIntNum) Decimal {

	d2 := Decimal{}
	d2.bigINum = bigINum.CopyOut()

	return d2
}

// NewInt - Returns a Decimal type based on int and precision input
// parameters. If an error is encountered, it will trigger a panic
// condition.
//
// The 'NewInt' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation and
// initialization in one step.
//
// Example: Decimal{}.NewInt(123456, 3) = 123.456
//
func (dec Decimal) NewInt(iNum int, precision uint) (Decimal, error) {
	d2 := Decimal{}.New()

	err := d2.SetInt(iNum, precision)

	if err != nil {
		return Decimal{},
		fmt.Errorf("Decimal.NewInt() Error returned by " +
			"d2.SetInt(iNum, precision). iNum='%v' precision='%v' Error='%v'",
				iNum, precision, err.Error())
	}

	return d2, nil
}

// NewI64 - Returns a Decimal type based on int64 and precision
// input parameters. If an error is encountered, it will trigger
// a panic condition.
//
// The 'NewI64' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewI64(i64, precision)
//	Decimal{}.NewI64(123456, 3) = 123.456
//
func (dec Decimal) NewI64(i64 int64, precision uint) (Decimal, error) {
	d2 := Decimal{}.New()
	err := d2.SetInt64(i64, precision)

	if err != nil {
		return Decimal{},
			fmt.Errorf("Decimal.NewInt() Error returned by " +
				"d2.SetInt64(i64, precision). i64='%v' precision='%v' Error='%v'",
				i64, precision, err.Error())
	}

	return d2, nil
}

// NewFloat32 - Creates a new Decimal instance based on a float32
// input.
func (dec Decimal) NewFloat32(f32 float32) (Decimal, error) {

	d2 := Decimal{}.New()
	err := d2.SetFloat32(f32)

	if err != nil {
		return Decimal{},
		fmt.Errorf("Error returned by d2.SetFloat32(f32). " +
			"f32='%v' Error='%v' ", f32, err.Error())
	}

	return d2, nil
}

// NewFloat64 - Creates a new Decimal instance based on a float64
// input.
func (dec Decimal) NewFloat64(f64 float64, maxPrecision uint) (Decimal, error) {

	numSeps := dec.GetNumericSeparatorsDto()

	d2 := Decimal{}.New()
	bigFloat := big.NewFloat(f64)

	err := d2.bigINum.SetBigFloat(bigFloat, maxPrecision)

	if err != nil {
		return Decimal{},
		fmt.Errorf("Error returned by d2.SetFloat64(f64). " +
			"f64='%v' Error='%v' ", f64, err.Error())
	}

	err = d2.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{},
			fmt.Errorf("Error returned by d2.SetFloat64(f64). " +
				"Error='%v' \n", err.Error())

	}

	return d2, nil
}

// NewNumStrsMultiple - Used to create and return an array of Decimal Types.
// Input parameters are a series of number strings.
func (dec Decimal) NewNumStrsMultiple(numStrs ...string) ([] Decimal, error) {

	ePrefix := "Decimal.NewNumStrsMultiple() "

	lenNumStrs := len(numStrs)

	decAry := make([] Decimal, lenNumStrs, lenNumStrs + 100)

	for i, numStr := range numStrs {

		dec := Decimal{}.New()

		err := dec.SetNumStr(numStr)

		if err != nil {
			return []Decimal{},
			fmt.Errorf(ePrefix + "Error returned by dec.SetNumStr(bigINum). " +
				"bigINum='%v' Index='%v' Error='%v'",
					numStr, i, err.Error()	)
		}

		decAry[i] = dec.CopyOut()

	}

	return decAry, nil
}

// NewNumStrArray - Used to create and return an array of Decimal Types.
// The input parameter is an array of number strings.
//
func (dec Decimal) NewNumStrArray(numStrs []string) ([] Decimal, error) {

	ePrefix := "Decimal.NewNumStrArray() "

	lenNumStrs := len(numStrs)

	decAry := make([] Decimal, lenNumStrs, lenNumStrs + 100)

	for i, numStr := range numStrs {

		dec := Decimal{}.New()

		err := dec.SetNumStr(numStr)

		if err != nil {
			return []Decimal{},
			fmt.Errorf(ePrefix + "Error returned by dec.SetNumStr(bigINum). " +
				"bigINum='%v' Index='%v' Error='%v'",
					numStr, i, err.Error()	)
		}

		decAry[i] = dec.CopyOut()

	}

	return decAry, nil
}

// NewNumStr - Returns a Decimal type based on a number string
// input parameter. If an error is encountered, it will trigger
// a panic condition.
//
// The 'NewNumStr' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewNumStr("123.456")
//
func (dec Decimal) NewNumStr(numStr string) (Decimal, error) {
	d2 := Decimal{}.New()

	err := d2.SetNumStr(numStr)

	if err != nil {
		return Decimal{},
			fmt.Errorf("Decimal.NewInt() Error returned by " +
				"d2.SetNumStr(bigINum). bigINum='%v' Error='%v'",
				numStr, err.Error())
	}

	return d2, nil
}

// NewNumStrDto - Returns a Decimal type based on a NumStrDto
// input parameter. If an error is encountered, it will trigger
// a panic condition.
//
// The 'NewNumStrDto' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
//
// Example: Decimal{}.NewNumStrDto(numDto)
//
func (dec Decimal) NewNumStrDto(numDto NumStrDto) (Decimal, error) {

	ePrefix := "Decimal.NewNumStrDto() "

	err := numDto.IsNumStrDtoValid("")

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Input parameter 'numDto' is INVALID! " +
			"Error='%v' ", err.Error())
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStrDto(numDto)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by d2.SetNumStrDto(numDto) " +
				" numDto.GetNumStr()='%v' Error='%v'",
				numDto.GetNumStr(), err.Error())
	}

	return d2, nil
}

// NewNumStrPrecision - Returns a Decimal type based on a number string
// and a precision value received as input parameters.
//
// The 'NewNumStrPrecision' method is designed to used in conjunction
// with Decimal{} thereby allowing Decimal creation and initialization
// in one step.
//
// Example: Decimal{}.NewNumStrPrecision('123456', 3, false) = 123.456
//
func (dec Decimal) NewNumStrPrecision(numStr string, precision uint, roundResult bool) (Decimal, error) {

	ePrefix := "Decimal.NewNumStrPrecision() "

	d2, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(numStr, precision, roundResult)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by NumStrPrecisionToDecimal(numStr, precision, " +
				"roundResult) numStr='%v' precision='%v' roundResult='%v' Error='%v' ",
					numStr, precision, roundResult, err.Error())
	}

	return d2, nil
}


// NewRationalNum - Creates a new Decimal instance based on input parameters consisting
// of a Rational Number (*big.Rat) and the specified 'precision' to be implemented in
// the resulting Decimal number value.
//
// For information on *big.Rat see https://golang.org/pkg/math/big/
//
func (dec *Decimal) NewRationalNum(bigRat *big.Rat, maxPrecision uint ) (Decimal, error) {

	ePrefix := "Decimal.NewRationalNum() "

	numSeps := dec.GetNumericSeparatorsDto()

	d2 := Decimal{}.NewZero(0)

	err := d2.bigINum.SetBigRat(bigRat, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
		fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetBigRat(bigRat, maxPrecision). " +
			"Error='%v' \n", err.Error())
	}

	err = d2.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
		fmt.Errorf(ePrefix +
			"Error returned by d2.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' \n", err.Error())
	}


	return d2, nil
}

// NewZero - Creates a New Decimal Instance with a
// value of zero. Input parameter 'precision' indicates
// the number of zeros formatted to the right of the
// decimal point.
//
func (dec Decimal) NewZero(precision uint) Decimal {
	d2 := Decimal{}

	d2.SetBigInt(big.NewInt(0), precision)

	return d2
}

// NumStrPrecisionToDecimal - receives a number string and a
// precision value as parameters. This method creates a Decimal
// Type containing the converted numeric value and returns it.
// For example, if passed the string ('str') '123456' and a precision
// value of '3', the resulting Decimal value would be 123.456.
//
// Example Usage:
// d := Decimal{}.NewBigIntNum()
// d2, err := d.NumStrPrecisionToDecimal("123456", 3, false)
// d2 is Now Equal to 123.456
//
func (dec *Decimal) NumStrPrecisionToDecimal(
					numStr string,
						requiredPrecision uint, roundResult bool) (Decimal, error) {

	ePrefix := "Decimal.NumStrPrecisionToDecimal() "
	var err error

	numSeps := dec.GetNumericSeparatorsDto()

	d2 := Decimal{}

	d2.bigINum, err = BigIntNum{}.NewNumStr(numStr)

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Error returned by BigIntNum{}.NewNumStr(numStr) " +
			"numStr='%v' Error='%v'",
			numStr, err.Error())
	}

	if roundResult == true {

		d2.bigINum.SetPrecision(requiredPrecision)

	} else {

		d2.bigINum.TruncToDecPlace(requiredPrecision)

	}

	err = d2.bigINum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by d2.bigINum.SetNumericSeparatorsDto(numSeps) " +
				"Error='%v'", err.Error())
	}


	return d2, nil
}

// Pow - Raises the current Decimal to the power of input parameter 'exponent'
// which is a Decimal Type. The result is returned as a Decimal type.
//
// Note that this method can process positive, negative, integer and fractional
// exponents.
//
// maxPrecision uint - 	Determines the maximum number of digits
// 											to the right of the	decimal point returned
// 											in the result. The actual precision may be
//											less than 'maxPrecision'.
//
func (dec *Decimal) Pow(exponent Decimal, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.Pow() "

	err := dec.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v'", err.Error())
	}

	bINumResult, err := BigIntMathPower{}.Pwr(dec.bigINum, exponent.bigINum, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntMathPower{}.Pwr(dec.bigINum, biNumExponent, maxPrecision). " +
				"Error='%v'", err.Error())
	}

	d3 := Decimal{}.NewBigIntNum(bINumResult)

	err = d3.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by d3.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto()). " +
				"Error='%v'", err.Error())
	}

	return d3, nil
}

// PowInt - raises the current Decimal to the power of
// an integer 'exponent'. The result is returned as
// a Decimal type.
//
// Input Parameters:
//
// maxPrecision uint - 	Determines the maximum number of digits
// 											to the right of the	decimal point returned
// 											in the result. The actual precision may be
//											less than 'maxPrecision'.
//
func (dec *Decimal) PowInt(exponent int, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.PowInt() "

	err := dec.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v'", err.Error())
	}

	biNumExponent := BigIntNum{}.NewBigInt(big.NewInt(int64(exponent)), 0)

	bINumResult, err := BigIntMathPower{}.Pwr(dec.bigINum, biNumExponent, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by BigIntMathPower{}.Pwr(dec.bigINum, biNumExponent, maxPrecision). " +
				"Error='%v'", err.Error())
	}

	d3 := Decimal{}.NewBigIntNum(bINumResult)

	err = d3.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix +
				"Error returned by d3.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto()). " +
				"Error='%v'", err.Error())
	}

	return d3, nil
}

// SetBigIntNum - Sets the numeric value of the current Decimal
// instance to that of input parameter, 'bigINum'
//
func (dec *Decimal) SetBigIntNum(bigINum BigIntNum) {
	dec.bigINum = bigINum.CopyOut()
}

// SetBigInt - Sets the value of the current Decimal instance to the
// input parameter 'iBig' scaled to the value of precision. In other
// words, if 'iBig' is set to a value of '123456' and precision is
// set to '3', the current Decimal will be set to a value of '123.456'.
//
// Using the same example, a big int value of '123456' with a precision of
// zero ('0') will yield an integer of '123456'.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// iBig := big.NewInt(int64(123))
// d.SetBigInt(iBig, 1)
//
func (dec *Decimal) SetBigInt(iBig *big.Int, precision uint) error {


	dec.bigINum = BigIntNum{}.NewBigInt(iBig, precision)

	return nil

}

// SetCurrencySymbol - sets the character which serves as the
// currency symbol for this Decimal value. Currency defaults
// to '$'.
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// For a listing of Major World Currency Symbols in Unicode format,
// see array 'NumStrCurrencySymbols' in source file:
//   MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
func (dec *Decimal) SetCurrencySymbol(currencySymbol rune) error {

	dec.bigINum.SetCurrencySymbol(currencySymbol)

	return nil
}

// SetDecimalSeparator - sets the character which separates
// the number into integer and fractional components. This
// defaults to the period '.'
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
// In the USA the Decimal Separator character is a period ('.')
//
func (dec *Decimal) SetDecimalSeparator(decimalSeparator rune) error {

	dec.bigINum.SetDecimalSeparator(decimalSeparator)

	return nil
}

// SetEmptySeparatorsToDefault - Ensures that separators are set to a valid value.
// If separator runes are zero, this methods sets the default values for
// decimal separator, thousands separator and currency symbol.
func (dec *Decimal) SetEmptySeparatorsToDefault() {

	if dec.GetDecimalSeparator() == 0 {
		dec.SetDecimalSeparator('.')
	}

	if dec.GetThousandsSeparator() == 0 {
		dec.SetThousandsSeparator(',')
	}

	if dec.GetCurrencySymbol() == 0 {
		dec.SetCurrencySymbol('$')
	}

}

// SetFloat32 - Sets the value of the current decimal to
// that of the passed-in float32 parameter.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// f32:= float32(123.456)
// d.SetFloat32(f32)
//
func (dec *Decimal) SetFloat32(f32 float32) error {

	ePrefix := "Decimal.SetFloat32() "

	dec.SetEmptySeparatorsToDefault()

	bigFloat := big.NewFloat(float64(f32))

	precision := bigFloat.Prec()

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	err := dec.bigINum.SetBigFloat(bigFloat, precision)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetBigFloat(bigFloat, precision). " +
			"Error= %v\n", err.Error())
	}

	err = dec.bigINum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetNumericSeparatorsDto(numSeps)). " +
			"Error= %v\n", err.Error())
	}


	return nil
}

// SetFloat64 - Sets the value of the current decimal to
// that of the passed-in float64 parameter.
//
// Example usage:
// d:= Decimal{}.NewZero(0)
// f64:= float64(123.456)
// d.SetFloat32(f64)
// Number String = "123.456"

func (dec *Decimal) SetFloat64(f64 float64) error {

	ePrefix := "Decimal.SetFloat64() "

	dec.SetEmptySeparatorsToDefault()

	bigFloat := big.NewFloat(f64)

	precision := bigFloat.Prec()

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	err := dec.bigINum.SetBigFloat(bigFloat, precision)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetBigFloat(bigFloat, precision). " +
			"Error= %v\n", err.Error())
	}

	err = dec.bigINum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetNumericSeparatorsDto(numSeps)). " +
			"Error= %v\n", err.Error())
	}

	return nil
}

// SetFloatBig - Sets the value of the current Decimal to the
// passed-in *big.Float parameter.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// bigFloat:= big.NewFloat(float64(123.456))
// d.SetBigFloat(bigFloat)
// Number String = "123.456"
//
func (dec *Decimal) SetFloatBig(bigFloat *big.Float) error {

	ePrefix := "Decimal.SetFloatBig() "

	dec.SetEmptySeparatorsToDefault()

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	err := dec.bigINum.SetBigFloat(bigFloat, bigFloat.Prec())

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetBigFloat(bigFloat, precision). " +
			"Error= %v\n", err.Error())
	}

	err = dec.bigINum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetNumericSeparatorsDto(numSeps)). " +
			"Error= %v\n", err.Error())
	}

	return nil
}

// Sets the value of the current Decimal to the input parameter 'iNum'
// scaled to the value of precision. In other words, if 'iNum' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// d.SetInt(123456, 3)
func (dec *Decimal) SetInt(iNum int, precision uint) error {

	ePrefix := "Decimal.SetInt() "

	iBig := big.NewInt(int64(iNum))

	err := dec.SetBigInt(iBig, precision)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.SetBigInt(iBig, precision). " +
			"Error='%v' ", err.Error())
	}

	return nil
}

// SetIntFracStrings - Sets the value of a decimal based on separate
// integer and fraction strings passed as input parameters. 'intNum'
// parameter will constitute the integer component of the Decimal value
// while 'fracNum' will be converted to the fractional component of the
// new Decimal value. 'fracNum' represents all the numeric digits to
// the right of the decimal place, while 'intNum' represents all the
// integer digits to the left of the decimal place.
//
// The parameter 'signVal will determine the sign Value for the returned
// Decimal type. It should be set to either +1 or -1.
//
func (dec *Decimal) SetIntFracStrings(intNum, fracNum string, signVal int) error {

	ePrefix := "Decimal.SetIntFracStrings() "

	var err error

	numSeps := dec.GetNumericSeparatorsDto()

	binIntNum := BigIntNum{}

	err = binIntNum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by binIntNum.SetNumericSeparatorsDto(numSeps). " +
				err.Error())
	}

	err = binIntNum.SetFromIntFracStrings(intNum, fracNum, signVal)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStr(intNum) " +
			"intNum='%v' Error='%v' \n", intNum, err.Error())
	}

	dec.bigINum.CopyIn(binIntNum)

	return nil
}

// Sets the value of the current Decimal to the input parameter 'i64'
// scaled to the value of precision. In other words, if 'i64' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Using the same example, an int64 value of '123456' and a precision
// value of zero ('0') will yield an integer value of '123456'.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// i64:= int64(123456)
// d.SetInt64(i64, 3)
func (dec *Decimal) SetInt64(i64 int64, precision uint) error {

	return dec.SetBigInt(big.NewInt(i64), precision)

}

// SetNumericSeparatorsDto - Sets the values of numeric separators:
// 		decimal point separator
//		thousands separator
//		currency symbol
//
// based on values transmitted through input parameter 'customSeparators'.
//
// If any of the values contained in input parameter 'customSeparators' is set
// to zero, an error will be returned.
//
func (dec *Decimal) SetNumericSeparatorsDto(customSeparators NumericSeparatorDto) error {

	ePrefix := "Decimal.SetNumericSeparatorsDto() "

	err := dec.bigINum.SetNumericSeparatorsDto(customSeparators)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by dec.bigINum.SetSeparatorDto(customSeparators) " +
			"Error='%v' ", err.Error())
	}

	return nil
}

// SetNumStrPrecision - Sets the Decimal's value to a number string and
// applies the appropriate 'precision' in order to determine placement
// of the decimal point. For example, if the number string ('str') is passed
// in as "123456" with a 'precision value of '3', the result is a value of
// 123.456.
//
// Example Usage:
// d:= Decimal{}.NewBigIntNum()
// d.SetNumStrPrecision("123456", 3)
// Resulting Decimal Value = 123.456
func (dec *Decimal) SetNumStrPrecision(str string, precision uint, roundResult bool) error {

	ePrefix := "Decimal.SetNumStrPrecision() "

	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	dec.SetEmptySeparatorsToDefault()

	d2, err := dec.NewNumStrPrecision(str, precision, roundResult)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"NumStrPrecisionToDecimal(str) failed. str=%v. Error= %v",
			str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetNumStr - Set's the Decimal's value to the input
// parameter 'str'. For example if 'str' is set equal
// to '123.456', this method will set the Decimal's
// value to 123.456.
//
// Example Usage:
// d := Decimal{}.NewBigIntNum()
// d.SetNumStr("123.456")
// Decimal Value = 123.456
func (dec *Decimal) SetNumStr(str string) error {

	dec.SetEmptySeparatorsToDefault()

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf("SetNumStr() Error. NumStrToDecimal(str) failed. str=%v. Error= %v.", str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetNumStrDto - Sets the value of the current Decimal type
// to the value represented by the incoming NumStrDto parameter.
func (dec *Decimal) SetNumStrDto(nDto NumStrDto) error {

	ePrefix := "Decimal.SetNumStrDto() "

	numSeps := dec.GetNumericSeparatorsDto()

	bIntNum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by BigIntNum{}.NewNumStrDto(nDto). " +
			"nDto='%v' Error='%v' \n", nDto.GetNumStr(), err.Error())
	}

	err = bIntNum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by bIntNum.SetNumericSeparatorsDto(numSeps). " +
			"Error='%v' \n", err.Error())
	}

	dec.bigINum = bIntNum.CopyOut()

	return nil
}

// SetPrecisionRound - Sets the precision or
// scale of the Decimal value. precision determines
// the number of digits displayed to the right of
// the decimal place. Note that precision is
// processed as an unsigned integer.
//
// When reducing precision, existing digits
// are ROUNDED! When increasing precision,
// additional zeros ('0') are added to the right
// of the decimal place.
func (dec *Decimal) SetPrecisionRound(precision uint) error {
	ePrefix := "Decimal.SetPrecisionRound() "

	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	dec.bigINum.SetPrecision(precision)

	return nil
}

// SetPrecisionTrunc - Sets the precision or
// scale of the Decimal value. precision determines
// the number of digits displayed to the right of
// the decimal place. Note that precision is
// processed as an unsigned integer.
//
// When reducing precision, existing digits
// are TRUNCATED! When increasing precision,
// additional zeros ('0') are added to the right
// of the decimal place.
func (dec *Decimal) SetPrecisionTrunc(precision uint) error {

	ePrefix := "Decimal.SetPrecisionTrunc() "

	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	dec.bigINum.TruncToDecPlace(precision)

	return nil
}

// SetSeparators - Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current number string.
//
// Note: If zero values are submitted as input, the values will default to USA standards.
//
// USA Examples:
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
//
func (dec *Decimal) SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol rune) {

	dec.bigINum.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

}


// SetSeparatorsToUSADefault - Sets Numeric separators:
// 			Decimal Point Separator
// 			Thousands Separator
//			Currency Symbol
//
// to United States of America (USA) defaults.
//
// Call specific methods to set numeric separators for other countries or
// cultures:
// 		dec.SetDecimalSeparator()
// 		dec.SetThousandsSeparator()
// 		dec.SetCurrencySymbol()
//
func (dec *Decimal) SetSeparatorsToUSADefault() {
	dec.SetDecimalSeparator('.')
	dec.SetThousandsSeparator(',')
	dec.SetCurrencySymbol('$')
}

// SetSignValue - Sets the sign of the numeric value
// for the current Decimal instance. Only two values
// are allowed: +1 and -1.
//
// If any other value is passed an error is thrown.
//
func (dec *Decimal) SetSign(newSignVal int) error {

	err := dec.bigINum.SetSignValue(newSignVal)

	if err != nil {
		return fmt.Errorf("Decimal.SetSign() Error - %v", err.Error())
	}

	return nil
}

// SetThousandsSeparator - sets the character which serves
// as the 'thousands' separator.
//
// In the USA, the Thousands Separator character is the comma (',').
// The Decimal Thousands Separator value defaults to the comma (',').
//
// Characters for Thousands Separators, Decimal Separators and Currency
// Symbols vary by country and culture. The Decimal Type allows the user
// to control the characters used for Thousands Separators, Decimal Separators
// and Currency Symbols.
//
func (dec *Decimal) SetThousandsSeparator(thousandsSeparator rune) error {

	dec.bigINum.SetThousandsSeparator(thousandsSeparator)

	return nil
}

// ShiftPrecisionLeft - Shifts precision of the current Decimal instance
// numeric value to the left by 'shiftLeftPlaces' decimal places. This
// is a 'relative' shift-left operation. The shift left operation is
// therefore performed with the current decimal point position as the
// starting point.
//
// This operation is equivalent to:	result = Decimal value / 10^shiftLeftPlaces
// or signed number divided by 10 raised to the power of shiftLeftPlaces.
//
// This method performs a relative shift left of the decimal point position.
// Be careful, this is NOT Shift Number Left operation. This is Shift Precision
// Left which means that the decimal point will be shifted left.
//
// See Examples below.
//
// Input Parameters
// ================
//
//	shiftLeftPlaces int	- The number of positions the decimal point will be
// 												shifted left from its current position.
//
// Examples:
// =========
//                  shift-left
// signed Number		  places				Result
//  "123456.789"				3						"123.456789"
//  "123456.789"				2						"1234.56789"
//  "123456.789"        6					  "0.123456789"
//  "123456789"	 			  6						"123.456789"
//  "123"               5	          "0.00123"
//  "0"								  3						"0"
//  "123456.789"				0						"123456.789"		- Zero has no effect on original number string
// "-123456.789"        0          "-123456.789"
// "-123456.789"        3          "-123.456789"
// "-123456789"			    6					 "-123.456789"
//
func (dec *Decimal) ShiftPrecisionLeft(shiftLeftPlaces uint) error {

	ePrefix := "Decimal.ShiftPrecisionLeft() "

	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	dec.bigINum.ShiftPrecisionLeft(shiftLeftPlaces)

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"After ShiftLeftPrecision() This Decimal instance is INVALID! " +
			"Error='%v'", err.Error())
	}

	return nil
}

// ShiftPrecisionRight - Shifts precision of the current Decimal instance
// numeric value to the right by 'shiftRightPlaces' decimal places. This
// is a 'relative' shift-right operation. The shift right operation is
// therefore performed with the current decimal point position as the
// starting point.
//
// This is equivalent to: result = Decimal value X 10^shiftRightPrecision or
// Decimal numeric value multiplied by 10 raised to the power of
// shiftRightPrecision.
//
// This method performs a relative shift right of the decimal point position.
// Be careful, this is NOT a Shift Number Right operation. This is Shift Precision
// Right which means that the decimal point will be shifted right.
//
// See Examples below.
//
// Input Parameters
// ================
//
//	shiftRightPlaces int	- The number of positions the decimal point will be
// 													shifted right from its current position.
//
// Examples:
// =========
//                  shift-right
// signed Number		  places				Result
//  "123456.789"				3						"123456789"
//  "123456.789"				2						"12345678.9"
//  "123456.789"        6					  "123456789000"
//  "123456789"	 			  6						"123456789000000"
//  "123"               5	          "12300000"
//  "0"								  3						"0"
//  "123456.789"				0						"123456.789"		- Zero has no effect on original number string
// "-123456.789"        0          "-123456.789"
// "-123456.789"        3          "-123456789"
// "-123456789"			    6					 "-123456789000000"
//
func (dec *Decimal) ShiftPrecisionRight(shiftRightPlaces uint) error {

	ePrefix := "Decimal.ShiftPrecisionRight() "


	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	dec.bigINum.ShiftPrecisionRight(shiftRightPlaces)

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"After ShiftPrecisionRight() This Decimal instance is INVALID! " +
			"Error='%v'", err.Error())
	}

	return nil
}

// Subtract - Subtracts the incoming Decimal from the current
// Decimal and returns the result as Decimal Type.
func (dec *Decimal) Subtract(d2 Decimal) (Decimal, error) {

	ePrefix := "Decimal.Subtract() "
	var err error

	err = dec.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
		fmt.Errorf(ePrefix + "Error: The current Decimal Instance (dec) is INVALID! " +
			"Error='%v' ", err.Error())
	}

	err = d2.IsDecimalValid()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix + "Error: The Input Parameter (d2) is INVALID! " +
				"Error='%v' ", err.Error())
	}

	bINumResult := BigIntMathSubtract{}.SubtractBigIntNums(dec.bigINum, d2.bigINum)

	d3 := Decimal{}.NewBigIntNum(bINumResult)

	err = d3.IsDecimalValid()

	if err != nil {
		return Decimal{},
		  fmt.Errorf(ePrefix +
			"This Decimal Type resulting from subtraction is INVALID! " +
			"Error='%v' ", err.Error())
	}


	return d3, nil
}

// SubtractFromThis - Subtracts the value of the incoming Decimal type
// from the current Decimal type. The updated value is stored and retained
// in the current Decimal instance.
func (dec *Decimal) SubtractFromThis(d2 Decimal) error {

	d3, err := dec.Subtract(d2)

	if err != nil {
		return err
	}

	dec.CopyIn(d3)

	return nil
}

// SubtractFromThisMultiple - Subtracts the value of multiple incoming
// Decimal instances from the current Decimal type. The updated value
// is stored and retained in the current Decimal instance.
//
func (dec *Decimal) SubtractFromThisMultiple(decs ...Decimal) error {

	for _, dx := range decs {

		dec3, err := dec.Subtract(dx)

		if err != nil {
			return fmt.Errorf("SubtractFromThisMultiple() - Error received from Add(). Error= %v", err)
		}

		dec.CopyIn(dec3)

	}

	return nil
}

// SubtractFromThisArray - Subtracts the values of an Array of incoming
// Decimal instances from the current Decimal type. The updated value
// is stored and retained in the current Decimal instance.
//
func (dec *Decimal) SubtractFromThisArray(decs []Decimal) error {

	for _, dx := range decs {

		dec3, err := dec.Subtract(dx)

		if err != nil {
			return fmt.Errorf("SubtractFromThisArray() - Error received from Add(). Error= %v", err)
		}

		dec.CopyIn(dec3)

	}

	return nil
}


