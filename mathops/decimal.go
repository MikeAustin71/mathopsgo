package mathops

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
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

type Decimal struct {
	numStrDto             NumStrDto
}

// Add - Adds the value of the current Decimal to that of
// the incoming Decimal and returns in the result in a
// Decimal Type.
func (dec *Decimal) Add(d2 Decimal) (Decimal, error) {

	ePrefix := "Decimal.Add() "
	var err error

	err = dec.IsDecimalValid()

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "The current host Decimal object is INVALID! " +
			"Error='%v' ", err.Error())
	}

	err = d2.IsDecimalValid()

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Incoming Decimal (d2) is INVALID! " +
			"Error= '%v' ", err.Error())
	}

	base10 := big.NewInt(10)
	var d4 Decimal

	var s3Text string
	var newPrecision uint
	var nDto NumStrDto
	
	x, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error= '%v' ", err.Error())
	}
	
	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)


	x, err = d2.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error= '%v' ", err.Error())
	}
	
	
	d2SignedAllDigitsBigInt := big.NewInt(0).Set(x)
	
	// precision is uint
	decPrecision := dec.numStrDto.GetPrecision()
	
	d2Precision := d2.numStrDto.GetPrecision()
	
	
	if decPrecision == d2Precision {

		s3Val := big.NewInt(0).Add(decSignedAllDigitsBigInt, d2SignedAllDigitsBigInt)

		s3Text = s3Val.Text(10)

		newPrecision = decPrecision

	} else	if d2Precision > decPrecision {

		deltaPrecision := big.NewInt(int64(d2.GetPrecision() - dec.GetPrecision()))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		s1 := big.NewInt(0).Mul(decSignedAllDigitsBigInt, deltaPrecisionScale)
		s3 := big.NewInt(0).Add(s1, d2SignedAllDigitsBigInt)

		s3Text = s3.Text(10)

		newPrecision = d2Precision

	} else {

		// must be dec.precision > d2.precision
		idp := int(decPrecision) - int(d2Precision)
		i64DeltaPrecision := int64(idp)
		deltaPrecision := big.NewInt(i64DeltaPrecision)
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		s1 := big.NewInt(0).Mul(d2SignedAllDigitsBigInt, deltaPrecisionScale)
		s3 := big.NewInt(0).Add(s1, decSignedAllDigitsBigInt)

		s3Text = s3.Text(10)

		newPrecision = dec.numStrDto.GetPrecision()

	}

	// s3Text is now a pure number string with no decimal point.
	nDto, err = NumStrDto{}.NewPtr().ShiftPrecisionLeft(s3Text, newPrecision)

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewPtr()." +
			"ShiftPrecisionLeft(s3Text, newPrecision) " +
			"s3Text='%v' newPrecision='%v' Error='%v'",
				s3Text, newPrecision, err.Error())
	}

	d4, err = dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned from dec.MakeDecimalFromNumStrDto(nDto) " +
				"nDto.GetNumStr()='%v'  nDto.GetPrecision()='%v'  Error='%v'",
				nDto.GetNumStr(), nDto.GetPrecision(), err.Error())
	}

	return d4, nil
}

// AddToThis - adds the value of the incoming Decimal to that
// of the current Decimal object. The new total is retained
// in the current Decimal object.
func (dec *Decimal) AddToThis(d2 Decimal) error {

	dec3, err := dec.Add(d2)

	if err != nil {
		return fmt.Errorf("AddToThis() - Error received from Add(). Error= %v", err)
	}

	dec.CopyIn(dec3)

	return nil
}

// AddToThisArray - Receives an array of Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisArray(decs []Decimal) error {

	for _, dx := range decs {

		dec3, err := dec.Add(dx)

		if err != nil {
			return fmt.Errorf("AddToThisArray() - Error received from Add(). Error= %v", err)
		}

		dec.CopyIn(dec3)

	}

	return nil
}

// AddToThisMultiple - Receives multiple Decimal objects and
// adds them to the current Decimal value.
func (dec *Decimal) AddToThisMultiple(decs ...Decimal) error {

	for _, dx := range decs {

		dec3, err := dec.Add(dx)

		if err != nil {
			return fmt.Errorf("AddToThisMultiple() - Error received from Add(). Error= %v", err)
		}

		dec.CopyIn(dec3)

	}

	return nil
}

// AllDigitsNumStr - parses the incoming string and returns
// a pure number string consisting of all numeric digits. No
// sign characters, decimals or thousands separators are returned.
// The returned value is the absolute numeric value extracted from
// the 'numStrDto' input parameter.
func (dec *Decimal) AllDigitsNumStr(numStr string) (string, error) {

	nDto, err := NumStrDto{}.NewPtr().ParseNumStr(numStr)

	if err != nil {
		return "", fmt.Errorf("AllDigitsNumStr() - nDto.ParseNumStr(numStrDto) returned an error. numStrDto= '%v' Error= %v", numStr, err)
	}

	return string(nDto.GetAbsAllNumRunes()), nil
}

// NumStrToDecimal - Creates a Decimal type from a number
// string.
func (dec *Decimal) NumStrToDecimal(str string) (Decimal, error) {

	dec.SetEmptySeparatorsToDefault()

	n1 := NumStrDto{}.New()
	n1.SetThousandsSeparator(dec.GetThousandsSeparator())
	n1.SetDecimalSeparator(dec.GetDecimalSeparator())
	n1.SetCurrencySymbol(dec.GetCurrencySymbol())

	nDto, err := n1.ParseNumStr(str)

	if err != nil {
		return Decimal{}, fmt.Errorf("NumStrToDecimal() received errror from nDto.ParseNumStr(str). str='%v' Error= %v", str, err)
	}

	d2, err := dec.MakeDecimalFromNumStrDto(nDto)

	return d2, nil

}

// CopyIn - Receives an incoming Decimal object
// as an input parameter and sets the Current Decimal
// equal to that of the incoming Decimal object.
func (dec *Decimal) CopyIn(d2 Decimal) {

	dec.Empty()
	dec.numStrDto = d2.numStrDto.CopyOut()
	return 
}

func (dec *Decimal) CopyOut() Decimal {
	d2 := Decimal{}.New()
	d2.numStrDto = dec.numStrDto.CopyOut()
	return d2
}

// Divide - Divides the current decimal value by the input
// parameter 'divisor'.
func (dec *Decimal) Divide(divisor Decimal, precision int) (Decimal, error) {

	ePrefix := "Decimal.Divide() "
	
	base10 := big.NewInt(10)
	sDividend1 := big.NewInt(0)
	sDivisor2 := big.NewInt(0)
	x, err := dec.numStrDto.GetSignedBigInt()
	
	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
			"Error='%v' ", err.Error())
	}

	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)
	decPrecision := dec.numStrDto.GetPrecision()

	x, err = divisor.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by divisor.numStrDto.GetSignedBigInt() " +
				"Error='%v' ", err.Error())
	}
	
	divisorSignedAllDigitsBigInt:= big.NewInt(0).Set(x)
	divisorPrecision := divisor.numStrDto.GetPrecision()
	
	
	if divisorPrecision > decPrecision {
		deltaPrecision := big.NewInt(int64(divisor.GetPrecision() - dec.GetPrecision()))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		sDividend1 = big.NewInt(0).Mul(decSignedAllDigitsBigInt, deltaPrecisionScale)
		sDivisor2 = big.NewInt(0).Set(divisorSignedAllDigitsBigInt)

	} else if decPrecision > divisorPrecision {
		deltaPrecision := big.NewInt(int64(dec.GetPrecision() - divisor.GetPrecision()))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		sDividend1 = big.NewInt(0).Set(decSignedAllDigitsBigInt)
		sDivisor2 = big.NewInt(0).Mul(divisorSignedAllDigitsBigInt, deltaPrecisionScale)

	} else {
		sDividend1 = big.NewInt(0).Set(decSignedAllDigitsBigInt)
		sDivisor2 = big.NewInt(0).Set(divisorSignedAllDigitsBigInt)

	}

	rDividend := big.NewRat(1, 1).SetInt(sDividend1)
	rDivisor := big.NewRat(1, 1).SetInt(sDivisor2)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	numStr := rQuotient.FloatString(precision)

	return dec.NumStrToDecimal(numStr)
}

// Empty - Sets all values of the current Decimal's
// fields to their 'zero' values.
func (dec *Decimal) Empty() {
	dec.numStrDto =  NumStrDto{}.NewPtr().GetZeroNumStr(0)
	dec.numStrDto.SetSignValue(1)
	dec.SetEmptySeparatorsToDefault()
}

// GetAbsoluteValue - returns the absolute value of the
// decimal expressed as a string. If the decimal value is
// '-123.456', this method will return '123.456'.
func (dec *Decimal) GetAbsoluteValue() (Decimal, error) {
	ePrefix := "Decimal.GetAbsoluteValue() "
	var err error

	err = dec.IsDecimalValid()

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "This Decimal data is INVALID! Please re-initialize. " +
			"Error='%v'", err.Error())
	}

	d2 := dec.CopyOut()

	if d2.numStrDto.GetSign() == 1 {
		return d2, nil
	}

	d2.numStrDto.SetSignValue(1)

	x, err := d2.numStrDto.GetSignedBigInt()
	
	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by d2.numStrDto.GetSignedBigInt() " +
				"Error='%v'", err.Error())
	}
	
	d2SignedAllDigitsBigInt := big.NewInt(0).Set(x)

	rDividend := big.NewRat(1, 1).SetInt(d2SignedAllDigitsBigInt)

	d2ScaleFactor, err := d2.numStrDto.GetScaleFactor()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by d2.numStrDto.GetScaleFactor() " +
				"Error=%v", err.Error())
	}

	rDivisor := big.NewRat(1, 1).SetInt(d2ScaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)


	d2.numStrDto, err = 
			NumStrDto{}.NewNumStr(rQuotient.FloatString(int(d2.numStrDto.GetPrecision())))

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
			"Error returned by NewNumStr(rQuotient.FloatString(int(d2.precision))) " +
			"Error=%v", err.Error())
	}

	return d2, nil
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

	decSignedAllDigitsBigInt, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return "",
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt(). " +
				"Error='%v'", err.Error())
	}
		
	absVal := big.NewInt(0).Abs(decSignedAllDigitsBigInt)

	return absVal.String(), nil
}

// GetBigFloat - returns big Float representation of the Decimal Value.
func (dec *Decimal) GetBigFloat() (*big.Float, error) {
	ePrefix := "Decimal.GetBigFloat() "

	err := dec.IsDecimalValid()

	if err!=nil {
		return big.NewFloat(0.0),
		fmt.Errorf(ePrefix +
			"This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v'", err.Error())
	}

	bf, status := big.NewFloat(0.0).SetString(dec.numStrDto.GetNumStr())

	if !status {
		return big.NewFloat(0.0),
		fmt.Errorf(ePrefix + "SetString() Failed. NumStr= %v", dec.numStrDto.GetNumStr())
	}

	return bf, nil
}

// GetBigFloatString - returns a signed number string which is accurate out
// to a very large number of decimal places (40+ decimal places). In contrast,
// the *big.Float returned by GetBigFloat() is only accurate to about
// 16-17 decimal places. func (nDto *Decimal).
//
// The returned string result is trimmed to eliminate leading and trailing spaces.
func (dec *Decimal) GetBigFloatString(precision uint) (string, error) {

	ePrefix := "Decimal.GetBigFloatString() "

	err := dec.IsDecimalValid()

	if err != nil {
		return "",
		fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v' ", err.Error())
	}

	x, err := dec.numStrDto.GetSignedBigInt()
		
	if err != nil {
		return "",
			fmt.Errorf(ePrefix + 
				"Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error='%v' ", err.Error())
	}

	decSignedAllDigitsBigInt:= big.NewInt(0).Set(x)
		
	rDividend := big.NewRat(1, 1).SetInt(decSignedAllDigitsBigInt)
	decScaleFactor, err := dec.numStrDto.GetScaleFactor()

	if err != nil {
		return "",
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetScaleFactor()" +
				"Error='%v' ", err.Error())
	}

	rDivisor := big.NewRat(1, 1).SetInt(decScaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	fmtResult := strings.TrimLeft(strings.TrimRight(rQuotient.FloatString(int(precision)), " "), " ")

	return fmtResult, nil
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

	return dec.numStrDto.GetCurrencySymbol()
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
// GetThouStr() = $1,000,000.23
//
// numstr = -1000000.23
// GetThouStr() = -$1,000,000.23
//
// Note: If the current Decimal is invalid, this method
// returns an empty string.
//
func (dec *Decimal) GetCurrencyStr() string {
	return dec.numStrDto.GetCurrencyStr()
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
	return dec.numStrDto.GetCurrencyParen()
}


// GetDecimalSeparator - returns the Decimal's current
// value for Decimal Separator (i.e. '.')
func (dec *Decimal) GetDecimalSeparator() rune {

	return dec.numStrDto.GetDecimalSeparator()
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

	bf, status := big.NewFloat(0.0).SetString(dec.numStrDto.GetNumStr())

	if !status {
		return float32(0.0), big.Accuracy(0),
			fmt.Errorf(ePrefix + "SetString() Failed. NumStr= %v", dec.numStrDto.GetNumStr())
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

	bf, status := big.NewFloat(0.0).SetString(dec.numStrDto.GetNumStr())

	if !status {
		return float64(0.0), big.Accuracy(0),
		fmt.Errorf(ePrefix + "SetString() Failed. NumStr= %v", dec.numStrDto.GetNumStr())
	}

	f64, accuracy := bf.Float64()

	return f64, accuracy, nil
}

// GetIntAry - Returns an IntAry structure initialized
// to the value of the current 'Decimal' object.
func (dec *Decimal) GetIntAry() (IntAry, error) {
	
	ePrefix := "Decimal.GetIntAry() "
	x, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return IntAry{}.New(),
			fmt.Errorf(ePrefix + "- Received error from " +
				"dec.numStrDto.GetSignedBigInt() " +
				"Error= %v", err.Error())
	}

	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)
	
	decPrecision := dec.numStrDto.GetPrecision()
	
	ia, err := IntAry{}.NewBigInt(decSignedAllDigitsBigInt, decPrecision)

	if err != nil {
		return IntAry{}.New(),
		fmt.Errorf("Decimal.GetIntAry() - Received error from " +
			"IntAry{}.NewBigInt(decSignedAllDigitsBigInt, decPrecision). " +
			" decSignedAllDigitsBigInt='%v' decPrecision='%v' Error= %v",
			decSignedAllDigitsBigInt.Text(10), decPrecision, err.Error())
	}

	return ia, nil
}

// GetNumStr - Returns the internal value of the Decimal
// expressed as a signed numeric string. Precision, or
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

	return dec.numStrDto.GetNumStr()

}

// GetNumPren - Returns the internal value of the
// Decimal expressed as number string. Precision
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
	return dec.numStrDto.GetNumParen()
}


// GetNumStrDto - returns a NumStrDto structure initialized
// to the value of the current Decimal object.
func (dec *Decimal) GetNumStrDto() (NumStrDto, error) {
	ePrefix := "Decimal.GetNumStrDto() "

	err := dec.numStrDto.IsNumStrDtoValid(ePrefix + "-")

	if err != nil {
		return NumStrDto{},
			fmt.Errorf(ePrefix + "- Error returned from dec.numStrDto.IsNumStrDtoValid(ePrefix). " +
				"dec.numStrDto= '%v' Error= %v", dec.numStrDto.GetNumStr(), err)
	}

	return dec.numStrDto.CopyOut(), nil
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
func (dec *Decimal) GetNthRoot(nthRoot, maxPrecision uint) (Decimal, error) {

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
		fmt.Errorf(ePrefix + "- Received error from dec.GetIntAry(). Error= %v", err)
	}

	nrt := NthRootOp{}

	iaNth, err := nrt.GetNthRootIntAry(&ia, nthRoot, maxPrecision)

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
func (dec *Decimal) GetPrecision() int {

	return int(dec.numStrDto.GetPrecision())

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

	x, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return big.NewRat(1, 1),
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error='%v' ", err.Error())
	}

	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)
	
	rDividend := big.NewRat(1, 1).SetInt(decSignedAllDigitsBigInt)

	decScaleFactor, err := dec.numStrDto.GetScaleFactor()

	if err != nil {
		return big.NewRat(1, 1),
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetScaleFactor() " +
				"Error='%v' ", err.Error())
	}

	rDivisor := big.NewRat(1, 1).SetInt(decScaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	return rQuotient, nil
}

// GetSign - Returns the sign of the current
// Decimal Value. Return values are one of two
// integers: +1 or -1.
func (dec *Decimal) GetSign() int {

	return dec.numStrDto.GetSign()
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

	nRunes := dec.numStrDto.GetAbsFracRunes()

	lnRunes := len(nRunes)

	if lnRunes == 0 {
		return 0
	}

	lastDecIdx := 0

	for i := 0; i < lnRunes; i++ {

		if nRunes[i] >= '1' && nRunes[i] <= '9' {
			lastDecIdx = i
		}
	}

	return uint(lastDecIdx + 1)
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

	ePrefix := "Decimal.GetScaleVal() "

	err := dec.IsDecimalValid()

	if err != nil {
		return big.NewInt(0),
		fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
			"Error='%v' ", err.Error())
	}

	decScaleFactor, err := dec.numStrDto.GetScaleFactor()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetScaleFactor()" +
				"Error='%v' ", err.Error())

	}

	return decScaleFactor, nil
}

// GetSignedAllDigitsStr - Returns the Decimal's internal
// Signed All Digits Integer Value expressed as a string.
// No Fractional Digits are included, this is a signed
// integer number string. Example: The value '-123.456'
// would be returned as '-123456'
func (dec *Decimal) GetSignedAllDigitsStr() (string, error) {

	ePrefix := "Decimal.GetSignedAllDigitsStr() "

	err := dec.IsDecimalValid()

	if err != nil {
		return "",
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v' ", err.Error())
	}

	result := ""
	
	if dec.numStrDto.signVal < 0 {
		result += "-"
	}
	
	result += string(dec.numStrDto.absAllNumRunes)
	
	return result, nil
}

// GetSignedAllDigitsVal - returns the Decimal value expressed as an
// integer value using type *big.Int. No factional values are included.
// For example, the value '-123.456' would be returned as the integer
// value '-123456'.  To compute the precise value of the Decimal, this
// integer value would need to be divided by the 'Precision Value'. See
// GetScaleVal() below.
func (dec *Decimal) GetSignedAllDigitsVal() (*big.Int, error) {

	ePrefix := "Decimal.GetSignedAllDigitsVal() "

	err := dec.IsDecimalValid()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v' ", err.Error())
	}

	x, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return big.NewInt(0),
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error='%v' ", err.Error())
	}
	
	return big.NewInt(0).Set(x), nil
}

// GetSquareRoot - Returns a Decimal object equal to the square root
// of the current Decimal value.
//
// Note: If the current Decimal value is a negative value, an error will
// be generated. You cannot take the square root of a negative number.
func (dec *Decimal) GetSquareRoot(maxPrecision uint) (Decimal, error) {

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

	return dec.numStrDto.GetThousandsSeparator()
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
	return dec.numStrDto.GetThouStr()
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
	return dec.numStrDto.GetThouParen()
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
// maxPrecision int - determines the number of digits
// 			to the right of the decimal point in the result.
// 			if maxPrecision is less than zero, a default precision
//      of 500 digits to the right of the decimal point is
//      applied.
//
func (dec *Decimal) Inverse(maxPrecision int) (Decimal, error) {

	ePrefix := "Decimal.Inverse() "

	err := dec.IsDecimalValid()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v' ", err.Error())
	}


	if maxPrecision < 0 {
		maxPrecision = 500
	}

	x, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error='%v' ", err.Error())
	}

	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)
	
	rDividend := big.NewRat(1, 1).SetInt(decSignedAllDigitsBigInt)

	decScaleFactor, err := dec.numStrDto.GetScaleFactor()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Received error from dec.numStrDto.GetScaleFactor(). Error= %v", err)
	}


	rDivisor := big.NewRat(1, 1).SetInt(decScaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)
	rInverse := big.NewRat(1, 1).Inv(rQuotient)

	numStr := rInverse.FloatString(maxPrecision)

	d2, err := dec.NumStrToDecimal(numStr)

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix +
			"Received error from nDto.NumStrToDecimal(numStrDto). Error= %v", err)
	}

	rP := d2.GetRelevantPrecision()

	err = d2.SetPrecisionTrunc(rP)

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix +
			"Received error from d2.SetPrecisionTrunc(rP). Error= %v", err)
	}

	d2.SetSeparators(dec.GetDecimalSeparator(), dec.GetThousandsSeparator(), dec.GetCurrencySymbol())

	err = d2.IsDecimalValid()

	if err != nil {
		return Decimal{},
		 fmt.Errorf(ePrefix + "Newly created Decimal object is INVALID! " +
		 	"Error='%v' ", err.Error())

	}

	return d2, nil
}

// IsDecimalValid - Performs an internal diagnostic on the current
// Decimal instance and returns an 'error' if the instance is INVALID.
//
func (dec *Decimal) IsDecimalValid() error {

	ePrefix := "Decimal.IsDecimalValid() "

	err := dec.numStrDto.IsNumStrDtoValid(ePrefix + "- ")

	if err != nil {
		return err
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

	if dec.numStrDto.precision != 0 {
		return true, nil
	}

	return false, nil
}

// MakeDecimalBigIntPrecision - This method receives a *big.Int and a precision value which
// are used to construct and return a Decimal Type. The value of 'precision' determines the
// number of Big Int digits which will be placed to the right of the decimal place.
func (dec *Decimal) MakeDecimalBigIntPrecision(iBig *big.Int, precision uint) (Decimal, error) {

	ePrefix := "Decimal.MakeDecimalBigIntPrecision() "

	base10 := big.NewInt(10)

	exponent := big.NewInt(int64(precision))
	scaleFactor := big.NewInt(0).Exp(base10, exponent, nil)
	signedAllDigitsBigInt := big.NewInt(0).Set(iBig)

	rDividend := big.NewRat(1, 1).SetInt(signedAllDigitsBigInt)
	rDivisor := big.NewRat(1, 1).SetInt(scaleFactor)
	rQuotient := big.NewRat(1, 1).Quo(rDividend, rDivisor)

	var err error

	numStrDto, err := NumStrDto{}.NewNumStr(rQuotient.FloatString(int(precision)))

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix +
			"Error returned by NumStrDto{}.NewNumStr(rQuotient.FloatString(int(precision))) " +
			"Error='%v'", err.Error())
	}


	numStrDto.SetSeparators(dec.GetDecimalSeparator(),
													dec.GetThousandsSeparator(),
													dec.GetCurrencySymbol())

	d2, err := dec.MakeDecimalFromNumStrDto(numStrDto)

	if err != nil {

		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by d2.numStrDto.IsNumStrDtoValid() " +
				"Error='%v'", err.Error())
	}

	return d2, nil
}

// MakeDecimalFromNumStrDto - generates a Decimal Type based on string information
// provided by the 'nDto' NumStrDto parameter.
func (dec *Decimal) MakeDecimalFromNumStrDto(nDto NumStrDto) (Decimal, error) {

	ePrefix := "Decimal.MakeDecimalFromNumStrDto() "

	if len(nDto.GetAbsAllNumRunes()) == 0 {

		d2 := Decimal{}.New()

		return d2, nil
	}

	d2 := Decimal{}.New()
	d2.numStrDto = nDto.CopyOut()

	d2.SetCurrencySymbol(nDto.GetCurrencySymbol())
	d2.SetDecimalSeparator(nDto.GetDecimalSeparator())
	d2.SetThousandsSeparator(nDto.GetThousandsSeparator())


	err := d2.IsDecimalValid()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error creating new Decimal object (d2). " +
				"Error='%v' ", err.Error())
	}

	return d2, nil
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
			fmt.Errorf(ePrefix + "IntAry Invalid - Error: %v", err.Error())
	}

	iaStats := ia.GetIntAryStats()

	if iaStats.IsZeroValue {

		d2 := Decimal{}.New()

		return d2, nil
	}

	d2 := Decimal{}.New()
	d2.numStrDto, err = ia.GetNumStrDto()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "- Error returned by ia.GetNumStrDto(). Error= %v", err.Error())
	}

	err = d2.IsDecimalValid()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error creating new Decimal object (d2). " +
				"Error='%v' ", err.Error())
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

	x, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error='%v' ", err.Error())
	}

	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)
	decPrecision := dec.numStrDto.precision
	
	x, err = d2.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by d2.numStrDto.GetSignedBigInt() " +
				"Error='%v' ", err.Error())
	}

	d2SignedAllDigitsBigInt := big.NewInt(0).Set(x)
	d2Precision := d2.numStrDto.precision
	
	s3Val := big.NewInt(0).Mul(decSignedAllDigitsBigInt, d2SignedAllDigitsBigInt)
	s3Text := s3Val.Text(10)

	newPrecision := uint( decPrecision + d2Precision)

	// s3Text is now a pure number string with no decimal point.
	nDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(s3Text, newPrecision)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewPtr()." +
				"ShiftPrecisionLeft(s3Text, newPrecision) " +
				"s3Text='%v' newPrecision='%v' Error='%v'",
				s3Text, newPrecision, err.Error())
	}

	d4, err := dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned from dec.MakeDecimalFromNumStrDto(nDto) " +
				"nDto.GetNumStr()='%v'  nDto.GetPrecision()='%v'  Error='%v'",
				nDto.GetNumStr(), nDto.GetPrecision(), err.Error())
	}

	return d4, nil
}

// MulTotal - Multiplies the value of the incoming
// Decimal type by the value of the current or receiving
// Decimal type. The result of the multiplication is
// stored in the current Decimal type.
func (dec *Decimal) MulTotal(d2 Decimal) error {

	d3, err := dec.Mul(d2)

	if err != nil {
		return fmt.Errorf("MulTotal() Error from dec.Mul(d2). Error= %v", err)
	}

	dec.CopyIn(d3)

	return nil
}

// New - Creates and returns a Decimal type.
// The Decimal value is initialized to zero.
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
func (dec Decimal) NewFloat64(f64 float64, precision int) (Decimal, error) {

	d2 := Decimal{}.New()
	err := d2.SetFloat64(f64, precision)

	if err != nil {
		return Decimal{},
		fmt.Errorf("Error returned by d2.SetFloat64(f64). " +
			"f64='%v' Error='%v' ", f64, err.Error())
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
			fmt.Errorf(ePrefix + "Error returned by dec.SetNumStr(numStrDto). " +
				"numStrDto='%v' Index='%v' Error='%v'",
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
			fmt.Errorf(ePrefix + "Error returned by dec.SetNumStr(numStrDto). " +
				"numStrDto='%v' Index='%v' Error='%v'",
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
				"d2.SetNumStr(numStrDto). numStrDto='%v' Error='%v'",
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
// with Decimal{} thereby allowing Decimal creation
// and initialization in one step.
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
func (dec *Decimal) NewRationalNum(bigRat *big.Rat, precision int ) (Decimal, error) {
	ePrefix := "Decimal.NewRationalNum() "

	n1, err := NumStrDto{}.NewRational(bigRat, precision)

	d2, err := dec.MakeDecimalFromNumStrDto(n1)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned from Decimal.MakeDecimalFromNumStrDto(n1). " +
				"dec.numStrDto= '%v' Error= %v", dec.numStrDto.GetNumStr(), err)
	}

	return d2, nil
}

// NumStrPrecisionToDecimal - receives a number string and a
// precision value as parameters. This method creates a Decimal
// Type containing the converted numeric and returns it.
// For example, if passed the string ('str') '123456' and a precision
// value of '3', the resulting Decimal value would be 123.456.
//
// Example Usage:
// d := Decimal{}.New()
// d2, err := d.NumStrPrecisionToDecimal("123456", 3, false)
// d2 is Now Equal to 123.456
//
func (dec *Decimal) NumStrPrecisionToDecimal(
					str string,
						requiredPrecision uint,
								roundResult bool) (Decimal, error) {

	ePrefix := "Decimal.NumStrPrecisionToDecimal() "

	dec.SetEmptySeparatorsToDefault()

	var err error
	d2 := Decimal{}

	n1, err := NumStrDto{}.NewNumStr(str)

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewNumStr(str) " +
			"str='%v' Error='%v'",
				str, err.Error())
	}

	err = n1.SetThisPrecision(requiredPrecision, roundResult)

	if err != nil {
		return Decimal{},
		fmt.Errorf(ePrefix +
			"Error returned by n1.SetThisPrecision(requiredPrecision, roundResult) " +
			"requiredPrecision='%v' roundResult='%v' Error='%v' ",
			requiredPrecision, roundResult, err.Error())
	}

	n1.SetSeparators(dec.GetDecimalSeparator(), dec.GetThousandsSeparator(), dec.GetCurrencySymbol())

	d2, err = dec.MakeDecimalFromNumStrDto(n1)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error received from dec.MakeDecimalFromNumStrDto(n1). " +
				"n1.NumStr='%v' requiredPrecision='%v' Error= %v",
				n1.GetNumStr(), requiredPrecision, err.Error())
	}

	return d2, nil
}


// Pow - raises the current Decimal to the power of
// an integer 'exponent'. The result is returned as
// a Decimal type.
//
// Input Parameters:
// maxPrecision int - Determines the number of digits to the right of the
// 			decimal point returned in the result.
// 			** If the value of maxPrecision is minus 1 (-1):
//  					-- 	For Positive exponents, the actual number of digits calculated to
//        					the right of the decimal point will be returned.
//
//      			--	For Negative exponents, a default value of 500 digits to the right
//                  of the decimal point will be returned.
//
func (dec *Decimal) Pow(exponent int, maxPrecision int) (Decimal, error) {

	ePrefix := "Decimal.Pow() "

	err := dec.IsDecimalValid()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "This Decimal object is INVALID! Please re-initialize. " +
				"Error='%v'", err.Error())
	}

	expSign := 1

	if exponent < 0 {
		expSign = -1
		exponent = exponent * -1
	}

	decSignedAllDigitsBigInt, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by dec.numStrDto.GetSignedBigInt() " +
				"Error='%v'", err.Error())
	}
	
	
	s1Val :=  big.NewInt(0).Set(decSignedAllDigitsBigInt)
	ex := big.NewInt(int64(exponent))

	newPrecision := uint(dec.GetPrecision()) * uint(exponent)

	s3Val := big.NewInt(0).Exp(s1Val, ex, nil)

	s3Text := s3Val.Text(10)

	// s3Text is now a pure number string with no decimal point.
	nDto, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(s3Text, newPrecision)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewPtr()." +
				"ShiftPrecisionLeft(s3Text, newPrecision) " +
				"s3Text='%v' newPrecision='%v' Error='%v'",
				s3Text, newPrecision, err.Error())
	}

	d2, err := dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned from dec.MakeDecimalFromNumStrDto(nDto) " +
				"nDto.GetNumStr()='%v'  nDto.GetPrecision()='%v'  Error='%v'",
				nDto.GetNumStr(), nDto.GetPrecision(), err.Error())
	}

	if expSign == 1 {
		if maxPrecision > -1 {
			d2.SetPrecisionRound(uint(maxPrecision))
		}

		return d2, nil
	}

	// This is a negative exponent. Take the inverse of the result.
	d3, err := d2.Inverse(maxPrecision)

	if err != nil {
		return Decimal{}.New(), fmt.Errorf("Pow() Error from d2.Inverse(). Error= %v.", err.Error())
	}

	return d3, nil
}

// Sets the value of the current Decimal to the input parameter 'iBig'
// scaled to the value of precision. In other words, if 'iBig' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Using the same example, a big int value of '123456' with a precision of
// zero ('0') will yield an integer of '123456'.
//
// Example usage:
// d:= Decimal{}.New()
// iBig := big.NewInt(int64(123))
// d.SetBigInt(iBig, 1)
//
func (dec *Decimal) SetBigInt(iBig *big.Int, precision uint) error {

	d2, err := dec.MakeDecimalBigIntPrecision(iBig, precision)

	if err != nil {
		return fmt.Errorf("Error from MakeDecimalBigIntPrecision(). Error = %v", err)
	}

	dec.CopyIn(d2)

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

	dec.numStrDto.SetCurrencySymbol(currencySymbol)

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

	dec.numStrDto.SetDecimalSeparator(decimalSeparator)

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
// d:= Decimal{}.New()
// f32:= float32(123.456)
// d.SetFloat32(f32)
//
func (dec *Decimal) SetFloat32(f32 float32) error {

	ePrefix := "Decimal.SetFloat32() "

	dec.SetEmptySeparatorsToDefault()

	str := strconv.FormatFloat(float64(f32), 'f', -1, 32)

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by nDto.NumStrToDecimal(str). str= '%v'. " +
			"Error= %v", str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetFloat64 - Sets the value of the current decimal to
// that of the passed-in float64 parameter.
//
// Example usage:
// d:= Decimal{}.New()
// f64:= float64(123.456)
// d.SetFloat32(f64)
// Number String = "123.456"
func (dec *Decimal) SetFloat64(f64 float64, precision int) error {

	ePrefix := "Decimal.SetFloat64() "

	dec.SetEmptySeparatorsToDefault()

	str := strconv.FormatFloat(f64, 'f', precision, 64)

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error from nDto.NumStrToDecimal(str). str= '%v'. Error= %v",
				str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// SetFloatBig - Sets the value of the current Decimal to the
// passed-in *big.Float parameter.
//
// Example usage:
// d:= Decimal{}.New()
// bigFloat:= big.NewFloat(float64(123.456))
// d.SetBigFloat(bigFloat)
// Number String = "123.456"
//
func (dec *Decimal) SetFloatBig(bigFloat *big.Float) error {

	ePrefix := "Decimal.SetFloatBig() "

	dec.SetEmptySeparatorsToDefault()

	str := bigFloat.Text('f', -1)

	d2, err := dec.NumStrToDecimal(str)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error from nDto.NumStrToDecimal(str). str= '%v'. Error= %v",
				str, err)
	}

	dec.CopyIn(d2)

	return nil
}

// Sets the value of the current Decimal to the input parameter 'iNum'
// scaled to the value of precision. In other words, if 'iNum' is set
// to a value of '123456' and precision is set to '3', the current
// Decimal will be set to a value of '123.456'.
//
// Example usage:
// d:= Decimal{}.New()
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
// The parameter 'signVal will determine the Sign Value for the returned
// Decimal type. It should be set to either +1 or -1.
//
func (dec *Decimal) SetIntFracStrings(signVal int, intNum, fracNum string) error {

	ePrefix := "Decimal.SetIntFracStrings() "

	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	dec.SetEmptySeparatorsToDefault()

	intDto, err := NumStrDto{}.NewNumStr(intNum)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by NumStrDto{}.NewNumStr(intNum). " +
			"intNum='%v' Error='%v' ",intNum, err.Error())
	}

	fracDto, err := NumStrDto{}.NewNumStr(fracNum)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by NumStrDto{}.NewNumStr(fracNum). " +
			"fracNum='%v' Error='%v' ",fracNum, err.Error())
	}

	var nxNumStr string

	if signVal < 0 {
		nxNumStr += "-"
	}

	nxNumStr += string(intDto.absAllNumRunes)

	nxNumStr += string(dec.GetDecimalSeparator())

	nxNumStr += string(fracDto.absAllNumRunes)

	n0 := NumStrDto{}.New()

	n0.SetSeparators(dec.GetDecimalSeparator(),dec.GetThousandsSeparator(), dec.GetCurrencySymbol())

	n1, err := n0.ParseNumStr(nxNumStr)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"- Error returned from n0.ParseNumStr(nxNumStr). intNum= '%v'  fracNum='%v' Error= %v",
				intNum, fracNum, err.Error())
	}

	d2, err := dec.MakeDecimalFromNumStrDto(n1)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"- Error returned from dec.MakeDecimalFromNumStrDto(n1). " +
			"n1.GetNumStr() = '%v' Error= %v",
			n1.GetNumStr(), err)
	}

	dec.CopyIn(d2)

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
// d:= Decimal{}.New()
// i64:= int64(123456)
// d.SetInt64(i64, 3)
func (dec *Decimal) SetInt64(i64 int64, precision uint) error {

	return dec.SetBigInt(big.NewInt(i64), precision)

}

// SetNumStrPrecision - Sets the Decimal's value to a number string and
// applies the appropriate 'precision' in order to determine placement
// of the decimal point. For example, if the number string ('str') is passed
// in as "123456" with a 'precision value of '3', the result is a value of
// 123.456.
//
// Example Usage:
// d:= Decimal{}.New()
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
// d := Decimal{}.New()
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

	if len(nDto.absAllNumRunes) == 0 || nDto.GetNumStr() == "" {

		err := nDto.SetNumStr("0")

		if err != nil {
			return fmt.Errorf(ePrefix + "Error returned by nDto.SetNumStr(\"0\"). " +
				"Error='%v' ", err.Error())
		}

	}

	dec.SetEmptySeparatorsToDefault()

	d2, err := dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"- Error received from dec.MakeDecimalFromNumStrDto(nDto). " +
			"nDto.GetNumStr()= '%v' Error= %v",
			nDto.GetNumStr(), err.Error() )
	}

	dec.CopyIn(d2)

	return nil
}

// SetPrecisionRound - Sets the precision or
// scale of the Decimal value. Precision determines
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


	n1, err := NumStrDto{}.NewPtr().SetPrecision(dec.numStrDto.GetNumStr(), precision, true)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"- Received error from NumStrDto.SetPrecision(dec.numStrDto, precision, true). " +
			"dec.numStrDto='%v' precision='%v' Error= %v",
			dec.numStrDto, precision, err)

	}

	d2, err := dec.MakeDecimalFromNumStrDto(n1)

	if err != nil {
		return fmt.Errorf(ePrefix + "- Received error from dec.MakeDecimalFromNumStrDto(n1, precision). " +
			"dec.numStrDto='%v' precision='%v' Error= %v", dec.numStrDto.GetNumStr(), precision, err)

	}

	dec.CopyIn(d2)

	return nil
}

// SetPrecisionTrunc - Sets the precision or
// scale of the Decimal value. Precision determines
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

	if uint(dec.GetPrecision()) == precision {
		return nil
	}

	n1, err := NumStrDto{}.NewPtr().SetPrecision(dec.numStrDto.GetNumStr(), precision, false)

	if err != nil {
		return fmt.Errorf(ePrefix + "- Received error from NumStrDto.SetPrecision" +
			"(dec.numStrDto, precision, true). dec.numStrDto='%v' precision='%v' Error= %v",
			dec.numStrDto.GetNumStr(), precision, err)
	}

	d2, err := dec.MakeDecimalFromNumStrDto(n1)

	if err != nil {
		return fmt.Errorf(ePrefix +
			"- Received error from dec.MakeDecimalFromNumStrDto(n1, precision). " +
			"dec.numStrDto='%v' precision='%v' Error= %v", dec.numStrDto.GetNumStr(), precision, err)
	}

	dec.CopyIn(d2)

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

	dec.numStrDto.SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol)

}

// SetSignValue - Sets the sign of the numeric value
// for the current Decimal instance. Only two values
// are allowed: +1 and -1.
//
// If any other value is passed an error is thrown.
//
func (dec *Decimal) SetSign(newSignVal int) error {

	err := dec.numStrDto.SetSignValue(newSignVal)

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

	dec.numStrDto.SetThousandsSeparator(thousandsSeparator)

	return nil
}

// ShiftPrecisionLeft - shifts precision of the current Decimal instance
// to the left by 'shiftLeftPlaces' places. This is a 'relative' shift-left
// operation. The shift is performed with the current decimal point position
// as the starting point.
//
// This method performs a relative shift left of the decimal point position.
// See Examples below.
//
// This operation is equivalent to:
// 							result = signed number / 10^shiftPrecision
// 									or signed number divided by 10 raised to the power of shiftPrecision.
//
// Input Parameters
// ================
//
//	shiftLeftPlaces int	- The number of positions the decimal point will be shifted left
//												from its current position.
//
// Examples
// ========
//
// signedNumStr			precision			Result
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
func (dec *Decimal) ShiftPrecisionLeft(shiftLeftPlaces int) error {

	ePrefix := "Decimal.ShiftPrecisionLeft() "

	if shiftLeftPlaces < 0 {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter 'shiftLeftPlaces' is LESS THAN ZERO! " +
			"shiftLeftPlaces='%v' ", shiftLeftPlaces)
	}

	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	n1, err := NumStrDto{}.NewPtr().ShiftPrecisionLeft(
													dec.numStrDto.GetNumStr(),
														uint(shiftLeftPlaces))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by NumStrDto{}.NewPtr().ShiftPrecisionLeft(dec.numStrDto, " +
			"uint(shiftLeftPlaces) dec.numStrDto='%v' shiftLeftPlaces='%v' Error='%v'",
				dec.numStrDto.GetNumStr(), shiftLeftPlaces, err.Error())
	}

	dec.numStrDto = n1.CopyOut()

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal instance is INVALID! Please Re-Initialize. " +
			"Error='%v'", err.Error())
	}

	return nil
}

// ShiftPrecisionRight - Shifts precision of the current Decimal instance
// to the right by 'shiftRightPlaces' places. This is a 'relative' shift-right
// operation. The shift is performed with the current decimal point position
// as the starting point.
//
// This is equivalent to: result = signedNumStr X 10^precision or signedNumStr Multiplied
// by 10 raised to the power of precision.
//
// This method performs a relative shift right of the decimal point position.
// See Examples below.
//
// Input Parameters
// ================
//
//	shiftRightPlaces int	- The number of positions the decimal point will be shifted right
//													from its current position.
//
// Examples:
// signedNumStr			precision			Result
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
func (dec *Decimal) ShiftPrecisionRight(shiftRightPlaces int) error {

	ePrefix := "Decimal.ShiftPrecisionRight() "

	if shiftRightPlaces < 0 {
		return fmt.Errorf(ePrefix +
			"Error: Input parameter 'shiftRightPlaces' is LESS THAN ZERO! " +
			"shiftLeftPlaces='%v' ", shiftRightPlaces)
	}

	err := dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	n1, err := NumStrDto{}.NewPtr().ShiftPrecisionRight(
																	dec.numStrDto.GetNumStr(),
																		uint(shiftRightPlaces))

	if err != nil {
		return fmt.Errorf(ePrefix +
			"Error returned by NumStrDto{}.NewPtr().ShiftPrecisionLeft(dec.numStrDto, " +
			"uint(shiftLeftPlaces) dec.numStrDto='%v' shiftRightPlaces='%v' Error='%v'",
			dec.numStrDto.GetNumStr(), shiftRightPlaces, err.Error())
	}

	dec.numStrDto = n1.CopyOut()

	err = dec.IsDecimalValid()

	if err != nil {
		return fmt.Errorf(ePrefix +
			"This Decimal instance is INVALID! Please Re-Initialize. " +
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
		return Decimal{},
		  fmt.Errorf(ePrefix +
			"This Decimal object (dec) is INVALID! Please Re-initialize. " +
			"Error='%v' ", err.Error())
	}

	err = d2.IsDecimalValid()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"The input Decimal object (d2) is INVALID! Please Re-initialize. " +
				"Error='%v' ", err.Error())
	}

	base10 := big.NewInt(10)
	var d4 Decimal
	var newPrecision uint
	var s3Text string
	var nDto NumStrDto

	x, err := dec.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by dec.numStrDto.GetSignedBigInt(). " +
				"Error='%v' ", err.Error())
	}

	decSignedAllDigitsBigInt := big.NewInt(0).Set(x)
	decPrecision := dec.numStrDto.precision

	x, err = d2.numStrDto.GetSignedBigInt()

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix +
				"Error returned by d2.numStrDto.GetSignedBigInt(). " +
				"Error='%v' ", err.Error())
	}

	d2SignedAllDigitsBigInt := big.NewInt(0).Set(x)
	d2Precision := d2.numStrDto.precision
	
	
	
	if decPrecision == d2Precision {

		s1Val := big.NewInt(0).Set(decSignedAllDigitsBigInt)

		s2Val := big.NewInt(0).Set(d2SignedAllDigitsBigInt)

		s3Val := big.NewInt(0).Sub(s1Val, s2Val)

		s3Text = s3Val.String()

		newPrecision = uint(dec.GetPrecision())

	} else	if d2Precision > decPrecision {
		deltaPrecision := big.NewInt(int64(d2.GetPrecision() - dec.GetPrecision()))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)

		s1Val := big.NewInt(0).Mul(decSignedAllDigitsBigInt, deltaPrecisionScale)
		s2Val := big.NewInt(0).Set(d2SignedAllDigitsBigInt)
		s3Val := big.NewInt(0).Sub(s1Val, s2Val)
		s3Text = s3Val.Text(10)
		newPrecision = uint(d2.GetPrecision())

	} else {
		// must be decPrecision > d2Precision

		deltaPrecision := big.NewInt(int64(dec.GetPrecision() - d2.GetPrecision()))
		deltaPrecisionScale := big.NewInt(0).Exp(base10, deltaPrecision, nil)
		s1Val := big.NewInt(0).Set(decSignedAllDigitsBigInt)
		s2Val := big.NewInt(0).Mul(d2SignedAllDigitsBigInt, deltaPrecisionScale)

		s3Val := big.NewInt(0).Sub(s1Val, s2Val)

		s3Text = s3Val.Text(10)
		newPrecision = uint(dec.GetPrecision())

	}

	// s3Text is now a pure number string with no decimal point.
	nDto, err = NumStrDto{}.NewPtr().ShiftPrecisionLeft(s3Text, newPrecision)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned by NumStrDto{}.NewPtr()." +
				"ShiftPrecisionLeft(s3Text, newPrecision) " +
				"s3Text='%v' newPrecision='%v' Error='%v'",
				s3Text, newPrecision, err.Error())
	}

	d4, err = dec.MakeDecimalFromNumStrDto(nDto)

	if err != nil {
		return Decimal{},
			fmt.Errorf(ePrefix + "Error returned from dec.MakeDecimalFromNumStrDto(nDto) " +
				"nDto.GetNumStr()='%v'  nDto.GetPrecision()='%v'  Error='%v'",
				nDto.GetNumStr(), nDto.GetPrecision(), err.Error())
	}

	return d4, nil
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


