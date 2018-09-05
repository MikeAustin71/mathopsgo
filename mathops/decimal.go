package mathops

import (
	"errors"
	"fmt"
	"math/big"
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

	err = dec.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"The current host Decimal object is INVALID! "+
				"Error='%v' \n", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	err = d2.IsValid(ePrefix)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+"Input Parmeter Decimal (d2) is INVALID! "+
				"Error= '%v' \n", err.Error())
	}

	bINumResult := BigIntMathAdd{}.AddBigIntNums(dec.bigINum, d2.bigINum)

	d3 := Decimal{}.NewBigIntNum(bINumResult)

	err = d3.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"The Decimal Type resulting from Addition is INVALID! "+
				"Error='%v' \n", err.Error())
	}

	err = d3.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by d3.SetNumericSeparatorsDto(numSeps). "+
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
		return fmt.Errorf(ePrefix+
			"Error returned by dec.bigINum.SetNumericSeparatorsDto(numSeps) "+
			"Error='%v' \n", err.Error())
	}

	return nil
}

// AddToThisArray - Receives an array of Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisArray(decs []Decimal) error {

	ePrefix := "Decimal.AddToThisArray() "
	var err error

	err = dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: The current Decimal instance (dec) is INVALID! "+
			"Error='%v' ", err.Error())
	}

	if len(decs) == 0 {
		return errors.New(ePrefix +
			"Error: Input Array 'decs' is EMPTY!")
	}

	numSeps := dec.GetNumericSeparatorsDto()

	bINumResult := dec.bigINum.CopyOut()

	for i, dx := range decs {

		err = dx.IsValid(ePrefix)

		if err != nil {

			return fmt.Errorf(ePrefix+
				"Error: Array element decs[%v] is INVALID! Error='%v'",
				i, err.Error())
		}

		bINumResult = BigIntMathAdd{}.AddBigIntNums(bINumResult, dx.bigINum)

	}

	err = bINumResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by bINumResult.SetNumericSeparatorsDto(numSeps) "+
			"Error='%v' \n", err.Error())
	}

	dec.SetBigIntNum(bINumResult)

	err = dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: Decimal type resulting from Array Addition is INVALID! "+
			"Error='%v' ", err.Error())
	}

	return nil
}

// AddToThisSeries - Receives multiple Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisSeries(decs ...Decimal) error {

	ePrefix := "Decimal.AddToThisSeries() "
	var err error

	err = dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: The current Decimal instance (dec) is INVALID! "+
			"Error='%v' ", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	bINumResult := dec.bigINum.CopyOut()

	for i, dx := range decs {

		err = dx.IsValid(ePrefix)

		if err != nil {

			return fmt.Errorf(ePrefix+
				"Error: Series element decs[%v] is INVALID! Error='%v'",
				i, err.Error())
		}

		bINumResult = BigIntMathAdd{}.AddBigIntNums(bINumResult, dx.bigINum)
	}

	bINumResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by bINumResult.SetNumericSeparatorsDto(numSeps) "+
			"Error='%v' \n", err.Error())
	}

	dec.SetBigIntNum(bINumResult)

	err = dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: Decimal type resulting from Array Addition is INVALID! "+
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
			fmt.Errorf(ePrefix+
				"Error returned by BigIntNum{}.NewNumStr(numStr) "+
				"numStr='%v' Error='%v' \n", numStr, err.Error())
	}

	return bigIntNum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE), nil
}

// NumStrToDecimal - Creates a Decimal type from a number
// string.
//
// The returned Decimal contains teh same numeric separators (decimal separator,
// thousands separator and currency symbol) as the current Decimal instance.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) NumStrToDecimal(numStr string) (Decimal, error) {

	ePrefix := "Decimal.NumStrToDecimal() "

	d2 := Decimal{}

	var err error

	numSeps := dec.GetNumericSeparatorsDto()

	d2.bigINum, err = BigIntNum{}.NewNumStrWithNumSeps(numStr, numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by BigIntNum{}.NewNumStr(numStr). "+
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

// Cmp - Performs a comparison of two Decimal numeric values
// and returns an integer value indicating the the relationship
// between the two numeric values (i.e. greater than, less than,
// or equal).
//
// Return Values:
// dec == dec2 							Return  0
// dec > 	dec2							Return +1
// dec < 	dec2							Return -1
//
func (dec *Decimal) Cmp(dec2 Decimal) int {

	return dec.bigINum.Cmp(dec2.bigINum)
}

// CubeRoot - Returns a Decimal instance with a numeric value equal to the
// cube root of the current Decimal numeric value. The current Decimal instance
// is the radicand.
//
// Returns:
// ========
// The calculation result is returned as a Decimal instance. The returned Decimal instance
// will contain	numeric separators (decimal separator, thousands separator and currency symbol)
// copied from the current Decimal instance (dec).
//
func (dec *Decimal) CubeRoot(maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.CubeRoot() "

	err := dec.IsValid(ePrefix + "Current Decimal instance is INVALID! ")

	if err != nil {
		return Decimal{}, err
	}

	bINumThree := BigIntNum{}.NewThree(0)

	decCubeRoot := Decimal{}.New()

	decCubeRoot.bigINum, err =
		BigIntMathNthRoot{}.GetNthRoot(dec.bigINum, bINumThree, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by BigIntMathNthRoot{}.NthRoot(...). "+
				"Error='%v'\n", err.Error())
	}

	err = decCubeRoot.IsValid(ePrefix + "decCubeRoot INVALID! ")

	if err != nil {
		return Decimal{}.New(), err
	}

	return decCubeRoot, nil
}

// Divide - Divides the current decimal value by the input parameter 'divisor'
// and returns the quotient as a new Decimal instance.
//
// The calculation result is returned as a Decimal instance. The returned Decimal
// instance will contain	numeric separators (decimal separator, thousands separator
// and currency symbol) copied from the current Decimal instance (dec).
//
func (dec *Decimal) Divide(divisor Decimal, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.Divide() "
	var err error

	err = dec.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: The current Decimal Instance (dec) is INVALID! "+
				"Error='%v' \n", err.Error())
	}

	err = divisor.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: Input parameter 'divisor' is INVALID! "+
				"Error='%v' \n", err.Error())
	}

	d2Quotient := Decimal{}.New()

	d2Quotient.bigINum, err =
		BigIntMathDivide{}.BigIntNumFracQuotient(
			dec.bigINum, divisor.bigINum, maxPrecision)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+
				"Error returned by BigIntMathDivide{}.BigIntNumFracQuotient() "+
				"Error='%v' ", err.Error())
	}

	err = d2Quotient.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error: The quotient Decimal Instance (d2) resulting from Division is INVALID! "+
				"Error='%v' \n", err.Error())
	}

	return d2Quotient, nil
}

// Equal - Returns true if the input Decimal instance is equal
// in all respects to the current Decimal instance.
//
// Note: this method will return false if two decimal instances
// being compared have equal numeric values but different precisions
// or different numeric separators.
//
// To perform an equivalency test of numeric values, see method
// Decimal{}.EqualValue() below.
//
func (dec *Decimal) Equal(dec2 Decimal) bool {

	return dec.bigINum.Equal(dec2.bigINum)
}

// EqualValue - Returns 'true' if the input Decimal instance holds
// a numeric value equal to that of the current Decimal
// instance.
//
func (dec *Decimal) EqualValue(dec2 Decimal) bool {

	return dec.bigINum.EqualValue(dec2.bigINum)
}

// Empty - Sets all values of the current Decimal's
// fields to their 'zero' values.
func (dec *Decimal) Empty() {
	dec.bigINum = BigIntNum{}.NewZero(0)
	dec.SetNumericSeparatorsToUSADefault()
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

	err := dec.IsValid(ePrefix)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v'", err.Error())
	}

	return dec.bigINum.FormatNumStr(ABSOLUTEPURENUMSTRFMTMODE), nil
}

// GetBigFloat - returns big Float representation of the Decimal Value.
func (dec *Decimal) GetBigFloat() (*big.Float, error) {
	ePrefix := "Decimal.GetBigFloat() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return big.NewFloat(0.0),
			fmt.Errorf(ePrefix+
				"The current Decimal object is INVALID! Please re-initialize. "+
				"Error='%v'", err.Error())
	}

	return dec.bigINum.GetBigFloat(), nil
}

// GetBigFloatString - returns a signed number string which is accurate out
// to a large number of decimal places.
//
func (dec *Decimal) GetBigFloatString(precision uint) (string, error) {

	ePrefix := "Decimal.GetBigFloatString() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
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

	err := dec.IsValid(ePrefix)

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v' ", err.Error())
	}

	bInt, err := dec.bigINum.GetBigInt()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix+"Error returned by dec.bigINum.GetBigInt() "+
				"Error='%v' ", err.Error())
	}

	return bInt, nil
}

// GetBigIntNum - Converts the current Decimal numeric value to
// an instance of type 'BigIntNum' and returns it to the calling
// function.
//
// The returned BigIntNum type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal
// instance.
//
// This method performs a validity test on the current Decimal instance.
//
func (dec *Decimal) GetBigIntNum() (BigIntNum, error) {

	ePrefix := "Decimal.GetBigIntNum() "

	err := dec.IsValid(ePrefix + "Decimal INVALID! ")

	if err != nil {
		return BigIntNum{}, err
	}

	return dec.bigINum.CopyOut(), nil
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

// GetDecimal - Returns a deep copy of the current Decimal instance.
//
// The returned BigIntNum type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal
// instance.
//
// This method performs a validity test on the current Decimal instance.
//
func (dec *Decimal) GetDecimal() (Decimal, error) {

	ePrefix := "Decimal.GetDecimal() "

	err := dec.IsValid(ePrefix + "Decimal INVALID! ")

	if err != nil {
		return Decimal{}, err
	}

	return dec.CopyOut(), nil
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

	err := dec.IsValid(ePrefix)

	if err != nil {
		return float32(0), big.Accuracy(0),
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v'", err.Error())
	}

	bf, status := big.NewFloat(0.0).SetString(dec.bigINum.GetNumStr())

	if !status {
		return float32(0.0), big.Accuracy(0),
			fmt.Errorf(ePrefix+"SetString() Failed. NumStr= %v", dec.bigINum.GetNumStr())
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

	err := dec.IsValid(ePrefix)

	if err != nil {
		return float64(0), big.Accuracy(0),
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v' ", err.Error())
	}

	bf, status := big.NewFloat(0.0).SetString(dec.bigINum.GetNumStr())

	if !status {
		return float64(0.0), big.Accuracy(0),
			fmt.Errorf(ePrefix+"SetString() Failed. NumStr= %v", dec.bigINum.GetNumStr())
	}

	f64, accuracy := bf.Float64()

	return f64, accuracy, nil
}

// GetIntAryElements - Returns an IntAry structure initialized
// to the value of the current 'Decimal' object.
//
// The returned IntAry contains numeric separators (decimal
// separator, thousands separator and currency symbol) copied
// from the current Decimal instance.
//
func (dec *Decimal) GetIntAry() (IntAry, error) {

	ePrefix := "Decimal.GetIntAryElements() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return IntAry{}.New(),
			fmt.Errorf(ePrefix+"This Decimal instance is INVALID! "+
				"Error= %v", err.Error())
	}

	ia, err := dec.bigINum.GetIntAry()

	if err != nil {
		return IntAry{}.New(),
			fmt.Errorf(ePrefix+"Error received from "+
				"dec.bigINum.GetIntAryElements(). Error= %v",
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
//
// The returned NumStrDto type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal
// instance.
//
func (dec *Decimal) GetNumStrDto() (NumStrDto, error) {
	ePrefix := "Decimal.GetNumStrDto() "

	nDto, err := dec.bigINum.GetNumStrDto()

	if err != nil {
		return NumStrDto{},
			fmt.Errorf(ePrefix+"- Error returned from dec.bigINum.GetNumStrDto(). "+
				"dec.bigINum= '%v' Error= %v", dec.bigINum.GetNumStr(), err.Error())
	}

	return nDto, nil
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

	err := dec.IsValid(ePrefix)

	if err != nil {
		return big.NewRat(1, 1),
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v' ", err.Error())
	}

	x, err := dec.bigINum.GetBigInt()

	if err != nil {
		return big.NewRat(1, 1),
			fmt.Errorf(ePrefix+"Error returned by dec.bigINum.GetBigInt() "+
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

// GetSciNotationNumber - Converts the numeric value of the current
// Decimal instance into scientific notation and returns this value
// as an instance of type SciNotationNum.
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//										scientific notation string. If the value of 'mantissaLen'
//										is less than two ('2'), this method will automatically set
//										the 'mantissaLen' to a default value of two ('2').
//
// 										Example Scientific Notation:
// 										----------------------------
//
//  										scientific notation string: '2.652e+8'
//
//  										significand = '2.652'
//  										significand integer digit = '2'
//											mantissa		= significand factional digits = '.652'
//  										exponent    = '8'  (10^8)
//
func (dec *Decimal) GetSciNotationNumber(mantissaLen uint) (SciNotationNum, error) {

	sciNotationNum, err := dec.bigINum.GetSciNotationNumber(mantissaLen)

	if err != nil {
		ePrefix := "Decimal.GetSciNotationNumber() "
		return SciNotationNum{}.New(),
			fmt.Errorf(ePrefix+
				"Error returned by dec.bigINum.GetSciNotationNumber(mantissaLen). "+
				"Error='%v' ", err.Error())
	}

	return sciNotationNum, nil
}

// GetSciNotationStr - Returns a string expressing the current Decimal
// numerical value as scientific notation.
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//										scientific notation string. If the value of 'mantissaLen'
//										is less than two ('2'), this method will automatically set
//										the 'mantissaLen' to a default value of two ('2').
//
// 										Example Scientific Notation:
// 										----------------------------
//
//  										scientific notation string: '2.652e+8'
//
//  										significand = '2.652'
//  										significand integer digit = '2'
//											mantissa		= significand factional digits = '.652'
//  										exponent    = '8'  (10^8)
//
func (dec *Decimal) GetSciNotationStr(mantissaLen uint) (string, error) {

	ePrefix := "BigIntNum.GetSciNotationStr() "

	sciNotn, err := dec.bigINum.GetSciNotationNumber(mantissaLen)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"Error returned by dec.bigINum.GetSciNotationNumber(mantissaLen). "+
				"Error='%v'", err.Error())
	}

	result, err := sciNotn.GetSciNotationStr(mantissaLen)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"Error returned by sciNotn.GetSciNotationStr(mantissaLen). "+
				"Error='%v'", err.Error())
	}

	return result, nil
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
			fmt.Errorf(ePrefix+"Error returned by dec.bigINum.GetNumStrDto(). "+
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

	bInt, err := dec.bigINum.GetBigInt()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix+"Error returned by dec.bigINum.GetBigInt(). "+
				"Error='%v' ", err.Error())
	}

	return big.NewInt(0).Set(bInt), nil
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
// Decimal.IsValid(ePrefix) which returns an
// 'error' type.
//
func (dec *Decimal) GetIsValid() bool {

	ePrefix := "Decimal.GetIsValid() "
	err := dec.IsValid(ePrefix)

	if err != nil {
		return false
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
			fmt.Errorf(ePrefix+
				"Error returned by dec.bigINum.GetInverse(uint(maxPrecision)) "+
				"Error='%v' \n", err.Error())
	}

	d2, err := bin2.GetDecimal()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+
				"Error returned by bin2.GetDecimal() "+
				"Error='%v' \n", err.Error())
	}

	return d2, nil
}

// IsEvenNumber - If the numeric value of the current Decimal instance
// is evenly divisible by two, with no remainder, it is classified as
// an even number and this method will return 'true'.
//
func (dec Decimal) IsEvenNumber() (bool, error) {

	return dec.bigINum.IsEvenNumber()
}

// IsFraction - returns a boolean value. If 'true',
// it signals that the Decimal has digits to the
// right of the decimal place. If 'false', it
// signals that the decimal value is an integer with
// no digits to the right of the decimal place.
func (dec *Decimal) IsFraction() (bool, error) {

	ePrefix := "Decimal.IsFraction() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return false,
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v' ", err.Error())
	}

	if dec.bigINum.precision != 0 {
		return true, nil
	}

	return false, nil
}

// IsValid - Performs an internal diagnostic on the current
// Decimal instance and returns an 'error' if the instance is INVALID.
//
func (dec *Decimal) IsValid(errName string) error {

	if errName == "" {
		errName = "IntAry.IsValid() "
	}

	err := dec.bigINum.IsValid(errName)

	if err != nil {
		return err
	}

	return nil
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

	err = ia.IsValid(ePrefix + "- ")

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix+"IntAry Invalid - Error: %v\n", err.Error())
	}

	d2 := Decimal{}
	d2.bigINum, err = ia.GetBigIntNum()

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix+"Error returned by ia.GetBigIntNum() "+
				"Error: %v\n", err.Error())
	}

	return d2, nil
}

// Mod - performs a modulo operation where the current BigIntNum numeric value is the
// dividend and the divisor is the input parameter, 'divisor'.  The modulo operation finds
// the remainder after division of one number by another (sometimes called modulus).
// (Wikipedia: https://en.wikipedia.org/wiki/Modulo_operation)
//
// 	 									dividend = bNum
//   									dividend % divisor = modulo
//
// The result of this modulo operation is returned as a BigIntNum, 'modulo'. 'modulo' may
// consist of an integer or a floating point value consisting of integer and fractional
// digits.
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// floating point 'modulo'. Precision is defined as the the number of fractional digits to
// the right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// The returned BigIntNum instance, 'modulo', will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from the current BigIntNum
// instance (bNum).
//
func (dec *Decimal) Mod(divisor Decimal,
	maxPrecision uint) (modulo Decimal, err error) {

	ePrefix := "Decimal.Mod() "
	modulo = Decimal{}.New()
	var errx error

	modulo.bigINum, errx = BigIntMathDivide{}.BigIntNumModulo(
		dec.bigINum.CopyOut(), divisor.bigINum, maxPrecision)

	if errx != nil {
		modulo = Decimal{}.NewZero(0)
		err = fmt.Errorf(ePrefix+""+
			"Error returned by BigIntMathDivide{}.BigIntNumModulo(dec, divisor, maxPrecision). "+
			"dec='%v' divisor='%v' maxPrecision='%v' Error='%v'\n",
			dec.GetNumStr(), divisor.GetNumStr(), maxPrecision, errx.Error())

		return modulo, err
	}

	err = nil

	return modulo, err
}

// Multiply - Multiplies the numeric value of the current Decimal instance
// (multiplier) by input parameter Decimal type 'multiplicand' and returns
// the product as a new Decimal instance.
//
// Before the multiplication operation is initiated, this method performs a
// validity check on both the current Decimal instance and the input
// parameter.
//
// The returned 'product' Decimal instance will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from
// the current Decimal instance.
//
func (dec *Decimal) Multiply(multiplicand Decimal) (product Decimal, err error) {

	ePrefix := "Decimal.Multiply() "
	product = Decimal{}.New()

	err = dec.IsValid(ePrefix + "The current Decimal instance (dec) is INVALID! ")

	if err != nil {
		return product, err
	}

	err = multiplicand.IsValid(ePrefix)

	if err != nil {
		return product, err
	}

	numSeps := dec.GetNumericSeparatorsDto()

	product.bigINum =
		BigIntMathMultiply{}.MultiplyBigIntNums(
			dec.bigINum, multiplicand.bigINum)

	errx := product.bigINum.SetNumericSeparatorsDto(numSeps)

	if errx != nil {
		product = Decimal{}.NewZero(0)
		err = fmt.Errorf(ePrefix+
			"Error returned by product.bigINum.SetNumericSeparatorsDto(numSeps). "+
			"Error='%v' ", errx.Error())

		return product, err
	}

	err = nil

	return product, err
}

// New() - Creates and returns a Decimal type. The Decimal numeric value is
// initialized to zero.
//
// The returned new Decimal instance will contain USA default numeric separators
// (decimal separator, thousands separator and currency symbol).
//
// Example Usage:
//   d := Decimal{}.New()
//
// This is the recommended procedure for creating
// a Decimal type.
func (dec Decimal) New() Decimal {

	d := Decimal{}
	d.Empty()

	return d

}

// NewWithNumSeps() - Creates and returns a new Decimal instance.
// The returned new Decimal instance will contain numeric separators
// (decimal separator, thousands separator and currency symbol)
// copied from the input parameter, 'numSeps'.
//
// The Decimal value is initialized to zero.
//
// Example Usage:
//   d := Decimal{}.NewWithNumSeps(numSeps)
//
func (dec Decimal) NewWithNumSeps(numSeps NumericSeparatorDto) Decimal {

	d := Decimal{}
	d.Empty()

	numSeps.SetDefaultsIfEmpty()

	d.SetNumericSeparatorsDto(numSeps)

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
func (dec Decimal) NewBigInt(bigI *big.Int, precision uint) Decimal {

	d2 := Decimal{}.New()

	d2.SetBigInt(bigI, precision)

	return d2
}

// NewBigIntNum - Returns a Decimal instance based on input
// parameter,
func (dec Decimal) NewBigIntNum(bigINum BigIntNum) Decimal {

	d2 := Decimal{}
	d2.bigINum = bigINum.CopyOut()

	return d2
}

// NewInt - Returns a Decimal type based on input parameters 'intNum'
// and 'precision'.
//
// Input parameter 'intNum' is of type int.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
// 				intNum := int(123456)
// 				precision := uint(3)
// 				dec := Decimal{}.NewInt(intNum, precision)
//        dec is now equal to 123.456
//
// Examples:
// ---------
//   intNum				precision			Decimal Result
//	 123456		 		   4							12.3456
//   123456          0              123456
//   123456          1              12345.6
//
func (dec Decimal) NewInt(intNum int, precision uint) Decimal {

	d2 := Decimal{}.New()

	d2.SetInt(intNum, precision)
	d2.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2
}

// NewIntExponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('intNum') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
// 				numeric value = integer X 10^exponent
//
// Input parameter 'intNum' is of type int.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//	decNum := Decimal{}.NewIntExponent(123456, -3)
//  -- decNum is now equal to "123.456", precision = 3
//
//	decNum := Decimal{}.NewIntExponent(123456, 3)
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//   intNum			 exponent			  	Decimal Result
//	 123456		 		  -3							123.456
//	 123456		 		   3							123456.000
//   123456          0              123456
//
func (dec Decimal) NewIntExponent(intNum int, exponent int) Decimal {

	d2 := Decimal{}.New()

	d2.bigINum.SetBigIntExponent(big.NewInt(int64(intNum)), exponent)

	return d2
}

// NewInt32 - Returns a new Decimal instance based on input parameters
// 'int32Num' and 'precision'.
//
// Input parameter 'int32Num' is of type int32.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
// 				int32Num := int32(123456)
// 				precision := uint(3)
// 				dec := Decimal{}.NewInt32(int32Num, precision)
//        dec is now equal to 123.456
//
// Examples:
// ---------
//  int32Num			precision			Decimal Result
//	 123456		 		   4							12.3456
//   123456          0              123456
//   123456          1              12345.6
//
func (dec Decimal) NewInt32(int32Num int32, precision uint) Decimal {

	d2 := Decimal{}.New()

	d2.SetInt64(int64(int32Num), precision)

	return d2
}

// NewInt32Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('int32Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
// 				numeric value = integer X 10^exponent
//
// Input parameter 'int32Num' is of type int32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction
// with the Decimal{} syntax thereby allowing Decimal
// type creation and initialization in one step.
//
//	decNum := Decimal{}.NewInt32Exponent(123456, -3)
//  -- decNum is now equal to "123.456", precision = 3
//
//	decNum := Decimal{}.NewInt32Exponent(123456, 3)
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//   int32Num		 exponent			  	Decimal Result
//	 123456		 		  -3							123.456
//	 123456		 		   3							123456.000
//   123456          0              123456
//
func (dec Decimal) NewInt32Exponent(int32Num int32, exponent int) Decimal {

	d2 := Decimal{}.New()

	bigI := big.NewInt(int64(int32Num))

	d2.bigINum.SetBigIntExponent(bigI, exponent)

	return d2
}

// NewInt64 - Returns a new Decimal instance based on input parameters
// 'int64Num' and 'precision'.
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'precision' indicates the number
// of digits to be formatted to the right of the decimal
// place. Input parameter 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction
// with the Decimal{} syntax thereby allowing Decimal
// type creation and initialization in one step.
//
// 				int64Num := int64(123456)
// 				precision := uint(3)
// 				dec := Decimal{}.NewInt64(int64Num, precision)
//        dec is now equal to 123.456
//
// Examples:
// ---------
//   int64Num			precision			 Decimal Result
//	 123456		 		   4							12.3456
//   123456          0              123456
//   123456          1              12345.6
//
func (dec Decimal) NewInt64(int64Num int64, precision uint) Decimal {

	d2 := Decimal{}.New()

	d2.SetInt64(int64Num, precision)

	return d2
}

// NewInt64Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('int64Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
// 				numeric value = integer X 10^exponent
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//	decNum := Decimal{}.NewInt64Exponent(123456, -3)
//  -- decNum is now equal to "123.456", precision = 3
//
//	decNum := Decimal{}.NewInt64Exponent(123456, 3)
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//   int64Num		 exponent			  	Decimal Result
//	 123456		 		  -3							123.456
//	 123456		 		   3							123456.000
//   123456          0              123456
//
func (dec Decimal) NewInt64Exponent(int64Num int64, exponent int) Decimal {

	d2 := Decimal{}.New()

	bigI := big.NewInt(int64Num)

	d2.bigINum.SetBigIntExponent(bigI, exponent)

	d2.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2

}

// NewFive - Returns a Decimal Type with a value equal to '5' (five).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '5', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//   value 					Result
// 		0								5
//		1								5.0
//    2								5.00
// 		3								5.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
//
func (dec Decimal) NewFive(precision uint) Decimal {

	bINum := BigIntNum{}.NewFive(precision)
	d2 := Decimal{}.New()
	d2.SetBigIntNum(bINum)

	return d2
}

// NewFloat32 - Creates a new Decimal instance based on a float32
// input.
func (dec Decimal) NewFloat32(f32 float32) (Decimal, error) {

	d2 := Decimal{}.New()
	err := d2.SetFloat32(f32)

	if err != nil {
		return Decimal{},
			fmt.Errorf("Error returned by d2.SetFloat32(f32). "+
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
			fmt.Errorf("Error returned by d2.SetFloat64(f64). "+
				"f64='%v' Error='%v' ", f64, err.Error())
	}

	err = d2.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{},
			fmt.Errorf("Error returned by d2.SetFloat64(f64). "+
				"Error='%v' \n", err.Error())

	}

	return d2, nil
}

// NewNumStrsMultiple - Used to create and return an array of Decimal Types.
// Input parameters are a series of number strings.
func (dec Decimal) NewNumStrsMultiple(numStrs ...string) ([]Decimal, error) {

	ePrefix := "Decimal.NewNumStrsMultiple() "

	lenNumStrs := len(numStrs)

	decAry := make([]Decimal, lenNumStrs, lenNumStrs+100)

	for i, numStr := range numStrs {

		dec := Decimal{}.New()

		err := dec.SetNumStr(numStr)

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix+"Error returned by dec.SetNumStr(bigINum). "+
					"bigINum='%v' Index='%v' Error='%v'",
					numStr, i, err.Error())
		}

		decAry[i] = dec.CopyOut()

	}

	return decAry, nil
}

// NewNumStrArray - Used to create and return an array of Decimal Types.
// The input parameter is an array of number strings.
//
func (dec Decimal) NewNumStrArray(numStrs []string) ([]Decimal, error) {

	ePrefix := "Decimal.NewNumStrArray() "

	lenNumStrs := len(numStrs)

	decAry := make([]Decimal, lenNumStrs, lenNumStrs+100)

	for i, numStr := range numStrs {

		dec := Decimal{}.New()

		err := dec.SetNumStr(numStr)

		if err != nil {
			return []Decimal{},
				fmt.Errorf(ePrefix+"Error returned by dec.SetNumStr(bigINum). "+
					"bigINum='%v' Index='%v' Error='%v'",
					numStr, i, err.Error())
		}

		decAry[i] = dec.CopyOut()

	}

	return decAry, nil
}

// NewNumStr - Returns a Decimal type based on a number string
// input parameter.
//
// This method assumes that the input parameter 'numStr' is a string
// of numeric digits which may be delimited by default USA numeric
// separators. Default USA numeric separators are defined as:
//  	decimal separator = '.'
//    thousands separator = ','
// 		currency symbol = '$'
//
// If the subject 'numStr' employs other national or cultural numeric
// separators, see method Decimal.NewNumStrWithNumSeps(), below.
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
		ePrefix := "Decimal.NewNumStr() "
		return Decimal{},
			fmt.Errorf(ePrefix+"Error returned by "+
				"d2.SetNumStr(numStr). numStr='%v' Error='%v'",
				numStr, err.Error())
	}

	return d2, nil
}

// NewNumStrWithNumSeps - Receives a number string as input and returns a
// new Decimal instance. The input parameter 'numSeps' contains numeric
// separators (decimal separator, thousands separator and currency symbol)
// which will be used to parse the number string.
//
// In addition, the numeric separators contained in input parameter 'numSeps'
// will be copied to the returned Decimal instance.
//
func (dec Decimal) NewNumStrWithNumSeps(
	numStr string,
	numSeps NumericSeparatorDto) (Decimal, error) {

	ePrefix := "Decimal.NewNumStrWithNumSeps() "

	numSeps.SetDefaultsIfEmpty()

	d2 := Decimal{}.New()

	err := d2.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+"Error returned by "+
				"d2.SetNumericSeparatorsDto(numSeps). Error='%v'",
				err.Error())
	}

	err = d2.SetNumStr(numStr)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+"Error returned by "+
				"d2.SetNumStr(numStr). numStr='%v' Error='%v'",
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

	err := numDto.IsValid("")

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+"Input parameter 'numDto' is INVALID! "+
				"Error='%v' ", err.Error())
	}

	d2 := Decimal{}.New()

	err = d2.SetNumStrDto(numDto)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+
				"Error returned by d2.SetNumStrDto(numDto) "+
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
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example: Decimal{}.NewNumStrPrecision('123456', 3, false) = 123.456
//
func (dec Decimal) NewNumStrPrecision(numStr string, precision uint, roundResult bool) (Decimal, error) {

	ePrefix := "Decimal.NewNumStrPrecision() "

	d2, err := Decimal{}.NewPtr().NumStrPrecisionToDecimal(numStr, precision, roundResult)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+"Error returned by NumStrPrecisionToDecimal(numStr, precision, "+
				"roundResult) numStr='%v' precision='%v' roundResult='%v' Error='%v' ",
				numStr, precision, roundResult, err.Error())
	}

	return d2, nil
}

// NewOne - Returns a Decimal Type with a value equal to '1' (one).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '1', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//   value 					Result
// 		0								1
//		1								1.0
//    2								1.00
// 		3								1.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
//
func (dec Decimal) NewOne(precision uint) Decimal {

	bINum := BigIntNum{}.NewOne(precision)
	d2 := Decimal{}.New()
	d2.SetNumericSeparatorsToDefaultIfEmpty()
	d2.SetBigIntNum(bINum)

	return d2
}

// NewRationalNum - Creates a new Decimal instance based on input parameters consisting
// of a Rational Number (*big.Rat) and the specified 'precision' to be implemented in
// the resulting Decimal number value.
//
// For information on *big.Rat see https://golang.org/pkg/math/big/
//
func (dec Decimal) NewRationalNum(bigRat *big.Rat, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.NewRationalNum() "

	numSeps := dec.GetNumericSeparatorsDto()

	d2 := Decimal{}.NewZero(0)

	err := d2.bigINum.SetBigRat(bigRat, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by dec.bigINum.SetBigRat(bigRat, maxPrecision). "+
				"Error='%v' \n", err.Error())
	}

	err = d2.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by d2.SetNumericSeparatorsDto(numSeps). "+
				"Error='%v' \n", err.Error())
	}

	return d2, nil
}

// NewTen - Returns a Decimal Type with a value equal to '10' (ten).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '10', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//   value 					Result
// 		0							 10
//		1							 10.0
//    2							 10.00
// 		3							 10.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
//
func (dec Decimal) NewTen(precision uint) Decimal {

	bINum := BigIntNum{}.NewTen(precision)
	d2 := Decimal{}.New()
	d2.SetBigIntNum(bINum)

	return d2
}

// NewThree - Returns a Decimal Type with a value equal to '3' (three).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '3', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//   value 					Result
// 		0								3
//		1								3.0
//    2								3.00
// 		3								3.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
//
func (dec Decimal) NewThree(precision uint) Decimal {

	bINum := BigIntNum{}.NewThree(precision)
	d2 := Decimal{}.New()
	d2.SetBigIntNum(bINum)

	return d2
}

// NewTwo - Returns a Decimal Type with a value equal to '2' (two).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '2', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//   value 					Result
// 		0								2
//		1								2.0
//    2								2.00
// 		3								2.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
//
func (dec Decimal) NewTwo(precision uint) Decimal {

	bINum := BigIntNum{}.NewTwo(precision)
	d2 := Decimal{}.New()
	d2.SetBigIntNum(bINum)

	return d2
}

// NewUint - Returns a new Decimal instance based on input parameters
// 'uintNum' and 'precision'.
//
// Input parameter 'uintNum' is of type uint.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
// 				uintNum := uint(123456)
// 				precision := uint(3)
// 				dec := Decimal{}.NewUint(int32Num, precision)
//        dec is now equal to 123.456
//
// Examples:
// ---------
//  uintNum			 precision	  	  Decimal Result
//	 123456		 		   4							12.3456
//   123456          0              123456
//   123456          1              12345.6
//
func (dec Decimal) NewUint(uintNum uint, precision uint) Decimal {

	d2 := Decimal{}.New()
	d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uintNum)), precision)
	d2.bigINum.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2
}

// NewUintExponent - This method returns a new Decimal instance in which
// the numeric value is set using an integer ('uintNum') multiplied by 10
// raised to the power of the input parameter, 'exponent'.
//
// 				numeric value = integer X 10^exponent
//
// Input parameter 'uintNum' is of type uint.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//	decNum := Decimal{}.NewUintExponent(123456, -3)
//  -- decNum is now equal to "123.456", precision = 3
//
//	decNum := Decimal{}.NewUintExponent(123456, 3)
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//  uintNum		    exponent		  	Decimal Result
//	 123456		 		  -3							123.456
//	 123456		 		   3							123456.000
//   123456          0              123456
//
func (dec Decimal) NewUintExponent(uintNum uint, exponent int) Decimal {

	d2 := Decimal{}.New()

	d2.bigINum.SetBigIntExponent(
		big.NewInt(0).SetUint64(uint64(uintNum)),
		exponent)

	d2.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2
}

// NewUint32 - Returns a new Decimal instance based on input parameters
// 'uint32Num' and 'precision'.
//
// Input parameter 'uint32Num' is of type uint32.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
// 				uint32Num := uint32(123456)
// 				precision := uint(3)
// 				dec := Decimal{}.NewUint32(int32Num, precision)
//        dec is now equal to 123.456
//
// Examples:
// ---------
//  uint32Num			precision			Decimal Result
//	 123456		 		   4							12.3456
//   123456          0              123456
//   123456          1              12345.6
//
func (dec Decimal) NewUint32(uint32Num uint32, precision uint) Decimal {

	d2 := Decimal{}.New()
	d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uint32Num)), precision)
	d2.bigINum.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2
}

// NewUint32Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('uint32Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
// 				numeric value = integer X 10^exponent
//
// Input parameter 'uint32Num' is of type uint32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//	decNum := Decimal{}.NewUint32Exponent(123456, -3)
//  -- decNum is now equal to "123.456", precision = 3
//
//	decNum := Decimal{}.NewUint32Exponent(123456, 3)
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//  uint32Num		 exponent			  	Decimal Result
//	 123456		 		  -3							123.456
//	 123456		 		   3							123456.000
//   123456          0              123456
//
func (dec Decimal) NewUint32Exponent(uint32Num uint32, exponent int) Decimal {

	d2 := Decimal{}.New()

	d2.bigINum.SetBigIntExponent(
		big.NewInt(0).SetUint64(uint64(uint32Num)),
		exponent)

	d2.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2
}

// NewUint64 - Returns a new Decimal instance based on input parameters
// 'uint64Num' and 'precision'.
//
// Input parameter 'uint64Num' is of type uint64.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place. Input parameter
// 'precision' is of type uint.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
// 				uint64Num := uint64(123456)
// 				precision := uint(3)
// 				dec := Decimal{}.NewUint64(int32Num, precision)
//        dec is now equal to 123.456
//
// Examples:
// ---------
//  uint64Num			precision			Decimal Result
//	 123456		 		   4							12.3456
//   123456          0              123456
//   123456          1              12345.6
//
func (dec Decimal) NewUint64(uint64Num uint64, precision uint) Decimal {

	d2 := Decimal{}.New()
	d2.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64Num), precision)
	d2.bigINum.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2
}

// NewUint64Exponent - This method returns a new Decimal instance in which the
// numeric value is set using an integer ('uint64Num') multiplied by 10 raised
// to the power of the input parameter, 'exponent'.
//
// 				numeric value = integer X 10^exponent
//
// Input parameter 'uint64Num' is of type uint64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the Decimal{}
// syntax thereby allowing Decimal type creation and initialization in
// one step.
//
//	decNum := Decimal{}.NewUint64Exponent(123456, -3)
//  -- decNum is now equal to "123.456", precision = 3
//
//	decNum := Decimal{}.NewUint64Exponent(123456, 3)
//  -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//  uint64Num		 exponent			  	Decimal Result
//	 123456		 		  -3							123.456
//	 123456		 		   3							123456.000
//   123456          0              123456
//
func (dec Decimal) NewUint64Exponent(uint64Num uint64, exponent int) Decimal {

	d2 := Decimal{}.New()

	d2.bigINum.SetBigIntExponent(
		big.NewInt(0).SetUint64(uint64Num),
		exponent)

	d2.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto())

	return d2
}

// NewZero - Creates a New Decimal Instance with a value of zero. Input
// parameter 'precision' indicates the number of zeros formatted to the
// right of the decimal place.
//
// Examples:
// =========
//
// 'precision'
//   value 					Result
// 		0								0
//		1								0.0
//    2								0.00
// 		3								0.000
//
// The new Decimal instance returned by this method will contain USA default
// numeric separators (decimal separator, thousands separator and currency symbol).
//
func (dec Decimal) NewZero(precision uint) Decimal {

	biNum := BigIntNum{}.NewZero(precision)

	d2 := Decimal{}.New()
	d2.SetNumericSeparatorsToDefaultIfEmpty()
	d2.SetBigIntNum(biNum)

	return d2
}

// NthRoot - Calculates the nth root of the current Decimal value. The numeric value of
// the current Decimal instance constitutes the radicand.
//
// Input Parameters:
// =================
//
//  nthRoot Decimal - Nth root specifies the root which will be calculated using the current
// 										Decimal instance as the radicand.
// 										Example, square root, cube root, 4th root, 9th root etc.
//
// maxPrecision uint -  Specifies the maximum number of decimals to the right of the decimal
// 											place to which the Nth root will be calculated. If the internal
// 											calculation exceeds the limit the nth root result will be rounded
//											to 'maxPrecision' decimal places.
//
// Returns:
// ========
// The nth root calculation result is returned as a Decimal instance. The returned Decimal
// instance will contain	numeric separators (decimal separator, thousands separator and
// currency symbol) copied from the current Decimal instance (radicand).
//
func (dec *Decimal) NthRoot(nthRoot Decimal, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.NthRoot() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix+
				"- The current Decimal object is INVALID!"+
				"Error='%v' ", err.Error())
	}

	err = nthRoot.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.New(),
			fmt.Errorf(ePrefix+
				"- The nthRoot Input Parameter is INVALID!"+
				"Error='%v' ", err.Error())
	}

	// If the radicand is zero, the result will always be zero
	if dec.IsZero() {
		return Decimal{}.NewZero(0), nil
	}

	// If nth root is zero, the result is always one.
	if nthRoot.IsZero() {
		return Decimal{}.NewOne(0), nil
	}

	nthRootIsEven, err := nthRoot.IsEvenNumber()

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by nthRoot.IsEvenNumber(). "+
				"Error='%v'\n", err.Error())
	}

	if dec.GetSign() == -1 && nthRootIsEven {

		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "INVALID ENTRY! Cannot calculate nth root of a negative radicand " +
				"when nthRoot is even.")
	}

	decNthRoot := Decimal{}.New()

	decNthRoot.bigINum, err =
		BigIntMathNthRoot{}.GetNthRoot(dec.bigINum, nthRoot.bigINum, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by BigIntMathNthRoot{}.NthRoot(dec, nthRoot, maxPrecision). "+
				"dec='%v' nthRoot='%v' maxPrecision='%v' Error='%v'\n",
				dec.GetNumStr(), nthRoot.GetNumStr(), maxPrecision, err.Error())
	}

	err = decNthRoot.IsValid(ePrefix + "decNthRoot result INVALID!")

	if err != nil {
		return Decimal{}.NewZero(0), err
	}

	return decNthRoot, nil
}

// NumStrPrecisionToDecimal - receives a number string and a
// precision value as parameters. This method creates a Decimal
// Type containing the converted numeric value and returns it.
// For example, if passed the string ('str') '123456' and a precision
// value of '3', the resulting Decimal value would be 123.456.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
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
			fmt.Errorf(ePrefix+"Error returned by BigIntNum{}.NewNumStr(numStr) "+
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
			fmt.Errorf(ePrefix+"Error returned by d2.bigINum.SetNumericSeparatorsDto(numSeps) "+
				"Error='%v'", err.Error())
	}

	return d2, nil
}

// Pow - Raises the numeric value of the current Decimal instance to the power of
// input parameter Decimal Type 'exponent'. The result is returned as a new Decimal
// instance.
//
// Note that this method can process positive, negative, integer and fractional
// exponents.
//
// exponent Decimal -		The numerical value of the current Decimal instance
//                      will be raised to the power of 'exponent'.
//                        			 result = dec^exponent
//
// maxPrecision uint - 	Determines the maximum number of digits
// 											to the right of the	decimal point returned
// 											in the result. The actual precision may be
//											less than 'maxPrecision'.
//
// The returned Decimal type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal instance.
//
func (dec *Decimal) Pow(exponent Decimal, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.Pow() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v'", err.Error())
	}

	d3 := Decimal{}.New()

	d3.bigINum, err = BigIntMathPower{}.Pwr(dec.bigINum, exponent.bigINum, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by BigIntMathPower{}.Pwr(dec.bigINum, biNumExponent, maxPrecision). "+
				"Error='%v'", err.Error())
	}

	err = d3.IsValid(ePrefix + "d3 INVALID! ")

	if err != nil {
		return Decimal{}.New(), err
	}

	return d3, nil
}

// PowInt - raises the current Decimal to the power of an integer 'exponent'.
// The result is returned as a Decimal type.
//
// Input Parameters:
//
// exponent 		int	 -	The numerical value of the current Decimal instance
//                      will be raised to the power of 'exponent'.
//                        			 result = dec^exponent
//
// maxPrecision uint - 	Determines the maximum number of digits
// 											to the right of the	decimal point returned
// 											in the result. The actual precision may be
//											less than 'maxPrecision'.
//
// The returned Decimal type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current Decimal instance.
//
func (dec *Decimal) PowInt(exponent int, maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.PowInt() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"This Decimal object is INVALID! Please re-initialize. "+
				"Error='%v'", err.Error())
	}

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	biNumExponent := BigIntNum{}.NewBigInt(big.NewInt(int64(exponent)), 0)

	bINumResult, err := BigIntMathPower{}.Pwr(dec.bigINum, biNumExponent, maxPrecision)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by BigIntMathPower{}.Pwr(dec.bigINum, biNumExponent, maxPrecision). "+
				"Error='%v'", err.Error())
	}

	d3 := Decimal{}.NewBigIntNum(bINumResult)

	err = d3.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+
				"Error returned by d3.SetNumericSeparatorsDto(dec.GetNumericSeparatorsDto()). "+
				"Error='%v'", err.Error())
	}

	return d3, nil
}

// SetBigIntNum - Sets the numeric value of the current Decimal
// instance to that of input parameter, 'bigINum'
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// The numeric separators associated with the incoming 'bigINum' are
// not copied to the current Decimal instance.
//
func (dec *Decimal) SetBigIntNum(bigINum BigIntNum) {

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	dec.bigINum = bigINum.CopyOut()

	dec.bigINum.SetNumericSeparatorsDto(numSeps)

}

// SetBigInt - Sets the value of the current Decimal instance to the
// input parameter 'iBig' scaled to the value of precision. In other
// words, if 'iBig' is set to a value of '123456' and precision is
// set to '3', the current Decimal will be set to a value of '123.456'.
//
// Using the same example, a big int value of '123456' with a precision of
// zero ('0') will yield an integer of '123456'.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// iBig := big.NewInt(int64(123))
// d.SetBigInt(iBig, 1)
// This yields a numeric value of d = 12.3
//
func (dec *Decimal) SetBigInt(iBig *big.Int, precision uint) {

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	dec.bigINum = BigIntNum{}.NewBigInt(iBig, precision)

	dec.bigINum.SetNumericSeparatorsDto(numSeps)

	return
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

// SetFloat32 - Sets the value of the current decimal to
// that of the passed-in float32 parameter.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum()
// f32:= float32(123.456)
// d.SetFloat32(f32)
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SetFloat32(f32 float32) error {

	ePrefix := "Decimal.SetFloat32() "

	dec.SetNumericSeparatorsToDefaultIfEmpty()

	bigFloat := big.NewFloat(float64(f32))

	precision := bigFloat.Prec()

	err := dec.bigINum.SetBigFloat(bigFloat, precision)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by dec.bigINum.SetBigFloat(bigFloat, precision). "+
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
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SetFloat64(f64 float64) error {

	ePrefix := "Decimal.SetFloat64() "

	dec.SetNumericSeparatorsToDefaultIfEmpty()

	bigFloat := big.NewFloat(f64)

	precision := bigFloat.Prec()

	err := dec.bigINum.SetBigFloat(bigFloat, precision)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by dec.bigINum.SetBigFloat(bigFloat, precision). "+
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
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SetFloatBig(bigFloat *big.Float) error {

	ePrefix := "Decimal.SetFloatBig() "

	dec.SetNumericSeparatorsToDefaultIfEmpty()

	err := dec.bigINum.SetBigFloat(bigFloat, bigFloat.Prec())

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by dec.bigINum.SetBigFloat(bigFloat, precision). "+
			"Error= %v\n", err.Error())
	}

	return nil
}

// SetInt - Sets the value of the current Decimal to the input parameter 'iNum'
// scaled to the value of input parameter 'precision'. In other words, if 'iNum'
// is set to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(9876)
// The numeric value of 'd' is now, 9876.
//
// d.SetInt(123456, 3) yields a new value for 'd' of 123.456.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SetInt(iNum int, precision uint) {

	iBig := big.NewInt(int64(iNum))

	dec.bigINum.SetBigInt(iBig, precision)

	return
}

// SetInt32 - Sets the value of the current Decimal to the input parameter
// 'int32Num' scaled to the value of input parameter, 'precision'. In other
// words, if 'int64Num' is set to a value of '123456' and precision is set
// to '3', the current Decimal will be set to a numeric value of '123.456'.
//
// Using the same example, an int32 value of '123456' and a precision
// value of zero ('0') will yield an integer value of '123456'.
//
// Input parameter 'int32Num' is an int32 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) yields numeric value of decimal instance 'd'
//     equals 956789.
//
// int32Num:= int32(123456)
// d.SetInt32(int32Num, 3) sets the numeric value of Decimal instance 'd' to 123.456.
//
func (dec *Decimal) SetInt32(int32Num int32, precision uint) {

	dec.bigINum.SetBigInt(big.NewInt(int64(int32Num)), precision)

	return
}

// SetInt64 - Sets the value of the current Decimal to the input parameter
// 'i64' scaled to the value of precision. In other words, if 'i64' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a numeric value of '123.456'.
//
// Using the same example, an int64 value of '123456' and a precision
// value of zero ('0') will yield an integer value of '123456'.
//
// Input parameter 'i64' is an int64 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//     equals 956789.
//
// i64:= int64(123456)
// d.SetInt64(i64, 3) sets the numeric value of Decimal instance 'd' to 123.456.
//
func (dec *Decimal) SetInt64(i64 int64, precision uint) {

	dec.bigINum.SetBigInt(big.NewInt(i64), precision)

	return
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
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
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

	err = binIntNum.SetIntFracStrings(intNum, fracNum, signVal)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by BigIntNum{}.NewNumStr(intNum) "+
			"intNum='%v' Error='%v' \n", intNum, err.Error())
	}

	dec.bigINum.CopyIn(binIntNum)

	return nil
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
		return fmt.Errorf(ePrefix+
			"Error returned by dec.bigINum.SetSeparatorDto(customSeparators) "+
			"Error='%v' ", err.Error())
	}

	return nil
}

// SetNumericSeparators - Used to assign values for the Decimal and Thousands separators as well
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
func (dec *Decimal) SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol rune) {

	dec.bigINum.SetNumericSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

}

// SetNumericSeparatorsToDefaultIfEmpty - If numeric separators are
// set to zero or nil, this method will set those numeric
// separators to the USA defaults. This means that the
// Decimal separator is set to ('.'), the Thousands separator
// is set to (',') and the currency symbol is set to '$'.
//
// If the numeric separators were previously set to a value
// other than zero or nil, that value is not altered by this
// method.
//
// Effectively, this method ensures that numeric separators
// are set to valid values.
//
func (dec *Decimal) SetNumericSeparatorsToDefaultIfEmpty() {

	dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

}

// SetNumericSeparatorsToUSADefault - Sets Numeric separators:
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
func (dec *Decimal) SetNumericSeparatorsToUSADefault() {

	dec.SetDecimalSeparator('.')
	dec.SetThousandsSeparator(',')
	dec.SetCurrencySymbol('$')
}

// SetNumStrPrecision - Sets the Decimal's value to a number string and
// applies the appropriate 'precision' in order to determine placement
// of the decimal point. For example, if the number string ('str') is passed
// in as "123456" with a 'precision value of '3', the result is a value of
// 123.456.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example Usage:
// d:= Decimal{}.NewBigIntNum()
// d.SetNumStrPrecision("123456", 3)
// Resulting Decimal Value = 123.456
func (dec *Decimal) SetNumStrPrecision(str string, precision uint, roundResult bool) error {

	ePrefix := "Decimal.SetNumStrPrecision() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"This Decimal object (dec) is INVALID! Please Re-initialize. "+
			"Error='%v' \n", err.Error())
	}

	numSeps := dec.GetNumericSeparatorsDto()

	d2, err := dec.NewNumStrPrecision(str, precision, roundResult)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"NumStrPrecisionToDecimal(str) failed. str=%v. Error= %v",
			str, err.Error())
	}

	err = d2.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by d2.SetNumericSeparatorsDto(numSeps). Error= '%v' \n",
			err.Error())

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
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SetNumStr(str string) error {

	ePrefix := "Decimal.SetNumStr() "

	numSeps := dec.GetNumericSeparatorsDto()

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error: NumStrToDecimal(str) failed. str=%v. Error= %v.",
			str, err.Error())
	}

	err = d2.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by d2.SetNumericSeparatorsDto(numSeps). Error= %v.",
			err.Error())
	}

	dec.CopyIn(d2)

	return nil
}

// SetNumStrDto - Sets the value of the current Decimal type
// to the value represented by the incoming NumStrDto parameter.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SetNumStrDto(nDto NumStrDto) error {

	ePrefix := "Decimal.SetNumStrDto() "

	numSeps := dec.GetNumericSeparatorsDto()

	bIntNum, err := BigIntNum{}.NewNumStrDto(nDto)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by BigIntNum{}.NewNumStrDto(nDto). "+
			"nDto='%v' Error='%v' \n", nDto.GetNumStr(), err.Error())
	}

	err = bIntNum.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by bIntNum.SetNumericSeparatorsDto(numSeps). "+
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

	err := dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"This Decimal object (dec) is INVALID! Please Re-initialize. "+
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
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SetPrecisionTrunc(precision uint) error {

	ePrefix := "Decimal.SetPrecisionTrunc() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"This Decimal object (dec) is INVALID! Please Re-initialize. "+
			"Error='%v' ", err.Error())
	}

	dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

	dec.bigINum.TruncToDecPlace(precision)

	return nil
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

// SetUint - Sets the value of the current Decimal instance to the input
// parameter 'uintNum' scaled to the value of input parameter 'precision'.
// In other words, if 'uintNum' is set to a value of '123456' and precision
// is set to '3', the current Decimal will be set to a numeric value of
// '123.456'.
//
// Using the same example, an uint32 value of '123456' and a precision value
// of zero ('0') will yield a numeric value of '123456'.
//
// Input parameter 'uintNum' is an uint type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//     equals 956789.
//
// uintNum := uint(123456)
// d.SetUint(uintNum, 3) sets the numeric value of Decimal instance 'd' to 123.456.
//
func (dec *Decimal) SetUint(uintNum uint, precision uint) {

	dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uintNum)), precision)

}

// SetUint32 - Sets the value of the current Decimal instance to the input
// parameter 'uint32Num' scaled to the value of input parameter 'precision'.
// In other words, if 'uint32Num' is set to a value of '123456' and precision
// is set to '3', the current Decimal will be set to a numeric value of
// '123.456'.
//
// Using the same example, an uint32 value of '123456' and a precision value
// of zero ('0') will yield a numeric value of '123456'.
//
// Input parameter 'uint32Num' is an uint32 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//     equals 956789.
//
// uint32Num := uint32(123456)
// d.SetUint32(uint32Num, 3) sets the numeric value of Decimal instance 'd' to 123.456.
//
func (dec *Decimal) SetUint32(uint32Num uint32, precision uint) {

	dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64(uint32Num)), precision)

}

// SetUint64 - Sets the value of the current Decimal instance to the input
// parameter 'uint64Num' scaled to the value of input parameter 'precision'.
// In other words, if 'uint64Num' is set to a value of '123456' and precision
// is set to '3', the current Decimal will be set to a numeric value of
// '123.456'.
//
// Using the same example, an uint64 value of '123456' and a precision value
// of zero ('0') will yield a numeric value of '123456'.
//
// Input parameter 'uint64Num' is an uint64 type.
//
// Input parameter 'precision' is an uint type which indicates the number
// of digits to be formatted to the right of the decimal place.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
// Example usage:
// d:= Decimal{}.NewBigIntNum(956789) = numeric value of decimal instance 'd'
//     equals 956789.
//
// uint64Num := uint64(123456)
// d.SetUint64(uint64Num, 3) sets the numeric value of Decimal instance 'd' to 123.456.
//
func (dec *Decimal) SetUint64(uint64Num uint64, precision uint) {

	dec.bigINum.SetBigInt(big.NewInt(0).SetUint64(uint64Num), precision)

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
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) ShiftPrecisionLeft(shiftLeftPlaces uint) error {

	ePrefix := "Decimal.ShiftPrecisionLeft() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"This Decimal object (dec) is INVALID! Please Re-initialize. "+
			"Error='%v' ", err.Error())
	}

	dec.bigINum.ShiftPrecisionLeft(shiftLeftPlaces)

	err = dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"After ShiftLeftPrecision() This Decimal instance is INVALID! "+
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
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) ShiftPrecisionRight(shiftRightPlaces uint) error {

	ePrefix := "Decimal.ShiftPrecisionRight() "

	err := dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"This Decimal object (dec) is INVALID! Please Re-initialize. "+
			"Error='%v' ", err.Error())
	}

	dec.bigINum.ShiftPrecisionRight(shiftRightPlaces)

	err = dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"After ShiftPrecisionRight() This Decimal instance is INVALID! "+
			"Error='%v'", err.Error())
	}

	return nil
}

// SquareRoot - Returns a Decimal instance with a numeric value equal to the
// square root of the current Decimal numeric value. The current Decimal instance
// is the radicand.
//
// Note: If the current Decimal value is a negative value, an error will be generated.
// You cannot take the square root of a negative number.
//
// Returns:
// ========
// The calculation result is returned as a Decimal instance. The returned Decimal instance
// will contain	numeric separators (decimal separator, thousands separator and currency symbol)
// copied from the current Decimal instance (dec).
//
func (dec *Decimal) SquareRoot(maxPrecision uint) (Decimal, error) {

	ePrefix := "Decimal.SquareRoot() "

	err := dec.IsValid(ePrefix + "Current Decimal instance is INVALID! ")

	if err != nil {
		return Decimal{}, err
	}

	if dec.GetSign() == -1 {

		return Decimal{}.NewZero(0),
			errors.New(ePrefix + "INVALID ENTRY! Cannot calculate nth root of a negative radicand " +
				"when nthRoot is even.")
	}

	bINumTwo := BigIntNum{}.NewTwo(0)

	decSqRoot := Decimal{}.New()

	decSqRoot.bigINum, err =
		BigIntMathNthRoot{}.GetNthRoot(dec.bigINum, bINumTwo, uint(maxPrecision))

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error returned by BigIntMathNthRoot{}.NthRoot(...). "+
				"Error='%v'\n", err.Error())
	}

	err = decSqRoot.IsValid(ePrefix + "decSqRoot INVALID! ")

	if err != nil {
		return Decimal{}.New(), err
	}

	return decSqRoot, nil
}

// Subtract - Subtracts the incoming Decimal from the current
// Decimal and returns the result as Decimal Type.
//
// The returned Decimal Type contains the same numeric separators
// (decimal separator, thousands separator and currency symbol)
// as those of the current Decimal instance. The numeric separators
// are copied form the current Decimal instance to the returned
// Decimal instance.
//
func (dec *Decimal) Subtract(d2 Decimal) (Decimal, error) {

	ePrefix := "Decimal.Subtract() "
	var err error

	err = dec.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error: The current Decimal Instance (dec) is INVALID! "+
				"Error='%v' ", err.Error())
	}

	err = d2.IsValid(ePrefix)

	if err != nil {
		return Decimal{}.NewZero(0),
			fmt.Errorf(ePrefix+"Error: The Input Parameter (d2) is INVALID! "+
				"Error='%v' ", err.Error())
	}

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	bINumResult := BigIntMathSubtract{}.SubtractBigIntNums(dec.bigINum, d2.bigINum)

	d3 := Decimal{}.NewBigIntNum(bINumResult)

	err = d3.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+
				"Error returned by d3.SetNumericSeparatorsDto(numSeps) "+
				"Error='%v' \n", err.Error())
	}

	err = d3.IsValid(ePrefix)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix+
				"This Decimal Type resulting from subtraction is INVALID! "+
				"Error='%v' ", err.Error())
	}

	return d3, nil
}

// SubtractFromThis - Subtracts the value of the incoming Decimal type
// from the current Decimal type. The updated value is stored and retained
// in the current Decimal instance.
//
// The numeric separators (decimal separator, thousands separator and
// currency symbol) for the current Decimal instance remain unchanged
// and are not modified by this method.
//
func (dec *Decimal) SubtractFromThis(d2 Decimal) error {

	ePrefix := "Decimal.Subtract() "
	var err error

	err = dec.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: The current Decimal Instance (dec) is INVALID! "+
			"Error='%v' ", err.Error())
	}

	err = d2.IsValid(ePrefix)

	if err != nil {
		return fmt.Errorf(ePrefix+"Error: The Input Parameter (d2) is INVALID! "+
			"Error='%v' ", err.Error())
	}

	numSeps := dec.bigINum.GetNumericSeparatorsDto()

	bINumResult := BigIntMathSubtract{}.SubtractBigIntNums(dec.bigINum, d2.bigINum)

	err = bINumResult.SetNumericSeparatorsDto(numSeps)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by bINumResult.SetNumericSeparatorsDto(numSeps) "+
			"Error='%v' \n", err.Error())
	}

	dec.bigINum = bINumResult.CopyOut()

	return nil
}

// SubtractFromThisMultiple - Subtracts the value of multiple incoming
// Decimal instances from the current Decimal type. The updated value
// is stored and retained in the current Decimal instance.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SubtractFromThisMultiple(decs ...Decimal) error {

	dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

	bINumResult := dec.bigINum.CopyOut()

	for _, dx := range decs {

		bINumResult = BigIntMathSubtract{}.SubtractBigIntNums(bINumResult, dx.bigINum)
	}

	dec.bigINum = bINumResult.CopyOut()

	return nil
}

// SubtractFromThisArray - Subtracts the values of an Array of incoming
// Decimal instances from the current Decimal type. The updated value
// is stored and retained in the current Decimal instance.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
//
func (dec *Decimal) SubtractFromThisArray(decs []Decimal) error {

	dec.bigINum.SetNumericSeparatorsToDefaultIfEmpty()

	bINumResult := dec.bigINum.CopyOut()

	for _, dx := range decs {

		bINumResult = BigIntMathSubtract{}.SubtractBigIntNums(bINumResult, dx.bigINum)

	}

	dec.bigINum = bINumResult.CopyOut()

	return nil
}
