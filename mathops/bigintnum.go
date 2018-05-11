package mathops

import (
	"math/big"
	"fmt"
	"strconv"
	"errors"
)

// BigIntNum - wraps a *big.Int integer and its associated
// precision and sign Value. While the numeric value is
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
	bigInt      				*big.Int
	absBigInt   				*big.Int
	precision   				uint     		// Number of digits to the right of the decimal point.
	scaleFactor 				*big.Int 		// Scale Factor =  10^(precision * -1)
	sign        				int      		// Valid values are -1 or +1. Indicates the sign of the
																	// 		the 'bigInt' integer.
	decimalSeparator 		rune				// Character used to separate integer and fractional digits ('.')
	thousandsSeparator 	rune 				// Character used to separate thousands (1,000,000,000
	currencySymbol 			rune				// Currency Symbol
}

// Ceiling - Ceiling: The least, or lowest value integer, which is greater than
// or equal to the numeric value of the current BigIntNum. Reference Wikipedia:
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
func (bNum *BigIntNum) Ceiling() BigIntNum {

	if bNum.IsZero() {
		return BigIntNum{}.NewBigInt(big.NewInt(0), 0)
	}

	if bNum.precision == 0 {
		return bNum.CopyOut()
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision)), nil)

	absQuotient:= big.NewInt(0).Quo(bNum.absBigInt, scaleVal)

	// absQuotient IS NOT EQUAL TO bNum.absBigInt

	if bNum.sign > 0 {
		// bNum is positive
		absQuotient = big.NewInt(0).Add(absQuotient, big.NewInt(1))
		return BigIntNum{}.NewBigInt(absQuotient, 0)
	}

	// bNum is negative
	return BigIntNum{}.NewBigInt(big.NewInt(0).Neg(absQuotient), 0)
}

// CmpBigInt - Compares the value of the *big.Int integer to that
// contained in an incoming BigIntNum.
//
// Return Values:
// bNum == bigIntNum 				Return  0
// bNum > bigIntNum					Return +1
// bNum < bigIntNum					Return -1
func (bNum *BigIntNum) CmpBigInt(bigIntNum BigIntNum) int {

	return bNum.bigInt.Cmp(bigIntNum.bigInt)

}

// CopyIn - Receives an incoming BigIntNum type and
// copies the value into the current BigIntNum instance.
//
func (bNum *BigIntNum) CopyIn(bigN BigIntNum) {

	bNum.bigInt = big.NewInt(0).Set(bigN.bigInt)
	bNum.absBigInt = big.NewInt(0).Set(bigN.absBigInt)
	bNum.precision = bigN.precision
	bNum.scaleFactor = big.NewInt(0).Set(bigN.scaleFactor)
	bNum.sign = bigN.sign
}

// CopyOut - Makes a deep copy of the current BigIntNum instance
// and returns it as a new BigIntNum instance.
//
func (bNum *BigIntNum) CopyOut() BigIntNum {

	b2 := BigIntNum{}.NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)

	return b2
}

// Empty - Resets the BigIntNum data fields to their
// uninitialized or zero state.
//
func (bNum *BigIntNum) Empty() {

	bNum.bigInt = big.NewInt(0)
	bNum.absBigInt = big.NewInt(0)
	bNum.scaleFactor = big.NewInt(1)
	bNum.sign = 1
	bNum.precision = 0

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	if bNum.currencySymbol == 0 {
		bNum.currencySymbol = '$'
	}

}

// Equal - Compares two BigIntNum instances and returns 'true'
// if the two instances are equal in value.
//
// If they are not Equal, the method returns 'false'.
//
func (bNum *BigIntNum) Equal(b2 BigIntNum) bool {

	if bNum.bigInt.Cmp(b2.bigInt) != 0 {
		return false
	}

	if bNum.absBigInt.Cmp(b2.absBigInt) != 0 {
		return false
	}

	if bNum.scaleFactor.Cmp(b2.scaleFactor) != 0 {
		return false
	}

	if bNum.sign != b2.sign {
		return false
	}

	if bNum.precision != b2.precision {
		return false
	}

	return true
}

// Floor - returns the greatest integer less than or equal to
// the numeric value of the current BigIntNum. Reference Wikipedia,
// https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
// 						Initial 			Floor
//  					 Value				Value
// 						-------      -------
//  						5.95					5
//  						5.05					5
//  						5							5
// 						 -5.05			 	 -6
//  						2.4				  	2
//  						2.9					 	2
// 						 -2.7				 	 -3
// 						 -2					 	 -2
//
func (bNum *BigIntNum) Floor() BigIntNum {

	if bNum.IsZero() {
		return BigIntNum{}.NewBigInt(big.NewInt(0), 0)
	}

	if bNum.precision == 0 {
		return bNum.CopyOut()
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision)), nil)

	absQuotient := big.NewInt(0).Quo(bNum.absBigInt, scaleVal)

	// absQuotient IS NOT EQUAL TO bNum.absBigInt

	if bNum.sign > 0 {
		// bNum is positive
		return BigIntNum{}.NewBigInt(absQuotient, 0)
	}

	// bNum is negative
	absQuotient = big.NewInt(0).Add(absQuotient, big.NewInt(1))

	return BigIntNum{}.NewBigInt(
											big.NewInt(0).Neg(absQuotient),0)
}

// FormatCurrencyStr - Formats the current BigIntNum numeric value as a currency string.
//
// If the Currency Symbol was not previously set for this BigIntNum, the currency symbol
// is defaulted to the USA standard dollar sign, ('$'). To use other currency symbols, see
// method BigIntNum.SetCurrencySymbol(). For a list of Major Currency Unicode Symbols, see
// constants located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// If the Decimal Separator was not previously set for this BigIntNum, the Decimal Separator
// is defaulted to the USA standard period ('.'). To use another character for Decimal
// Separator, see method BigIntNum.SetDecimalSeparator().
//
// If the Thousands Separator was not previously set for this BigIntNum, the Thousands
// Separator is defaulted to the USA standard comma (','). To use another character for
// Thousands Separator, see method BigIntNum.SetThousandsSeparator().
//
// Input Parameters
// ================
//
// negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
//																		LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//																														 		a leading minus sign.
//																																Example: -$123,456.78
//
//																		PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//																																surrounding parentheses.
//																																Example: ($123,456.78)
//
func (bNum *BigIntNum) FormatCurrencyStr(negValMode NegativeValueFmtMode) string {

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	if bNum.currencySymbol == 0 {
		bNum.currencySymbol = '$'
	}

	outRunes := make([]rune, 0, 300)

	scratchNum := big.NewInt(0).Set(bNum.absBigInt)
	baseZero := big.NewInt(0)

	if scratchNum.Cmp(baseZero) == 0 {
		bNum.sign = 1

		outRunes = append(outRunes, bNum.currencySymbol)

		outRunes = append(outRunes, '0')

		if bNum.precision > 0 {
			outRunes = append(outRunes, bNum.decimalSeparator)

			for h := 0; h < int(bNum.precision); h++ {
				outRunes = append(outRunes, '0')
			}

		}

		return string(outRunes)
	}

	startIdx := 0
	modulo := big.NewInt(0)
	baseTen := big.NewInt(10)
	digitCnt := 0
	thouCnt := -1

	if bNum.precision == 0 {
		thouCnt = 0
	}


	for scratchNum.Cmp(baseZero) == 1 {

		if startIdx==0 &&
			bNum.sign == -1  &&
			negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, ')')
		}

		modX := big.NewInt(0)
		scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, baseTen, modX)
		outRunes = append(outRunes, rune(modulo.Int64() + int64(48)))
		digitCnt++
		startIdx++

		if thouCnt > -1 {
			thouCnt++
		}

		if scratchNum.Cmp(baseZero) == 1 &&
			thouCnt==3 {

			outRunes = append(outRunes, bNum.thousandsSeparator)
			startIdx++
			thouCnt = 0
		}

		if bNum.precision > 0 &&
			int(bNum.precision) == startIdx {

			outRunes = append(outRunes, bNum.decimalSeparator)
			startIdx++
			thouCnt = 0
		}

	}

	if int(bNum.precision) >= digitCnt {

		delta := int(bNum.precision) - digitCnt + 1

		for k:=0; k < delta; k++ {
			outRunes = append(outRunes, '0')
			startIdx++

			if bNum.precision > 0 &&
				int(bNum.precision) == startIdx {

				outRunes = append(outRunes, bNum.decimalSeparator)
				startIdx++
			}
		}
	}

	startIdx--

	// append Currency Symbol
	outRunes = append(outRunes, bNum.currencySymbol)
	startIdx++


	// adjust for negative sign value
	if bNum.sign == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes = append(outRunes, '-')
			startIdx++

		}	else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			outRunes = append(outRunes, '(')
			startIdx+=2
		}
	}

	sortLimit := startIdx / 2
	tRune := rune(0)
	yCnt := 0

	for i:= startIdx; i > sortLimit ; i-- {
		tRune = outRunes[yCnt]
		outRunes[yCnt] = outRunes[i]
		outRunes[i] = tRune
		yCnt++
	}

	return string(outRunes)
}

// FormatNumStr - Formats the numeric value of the current BigIntNum
// instance as number string consisting of integer digits to the left
// of the decimal point and fractional digits to the right of the decimal
// point, if such fractional digits exist. The resulting number string
// will NOT contain a currency symbol or thousands separators.
//
// If the Decimal Separator was not previously set for this BigIntNum,
// the Decimal Separator is defaulted to the USA standard period ('.').
// To use another character for Decimal Separator, see method
// BigIntNum.SetDecimalSeparator().
//
// Output Examples: 123456.789 or -123456.789
//
// Input Parameters
// ================
//
// negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
//																		LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//																														 		a leading minus sign.
//																																Example: -123456.78
//
//																		PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//																																surrounding parentheses.
//																																Example: (123456.78)
//
func (bNum *BigIntNum) FormatNumStr(negValMode NegativeValueFmtMode) string {

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	outRunes := make([]rune, 0, 300)

	scratchNum := big.NewInt(0).Set(bNum.absBigInt)
	baseZero := big.NewInt(0)

	if scratchNum.Cmp(baseZero) == 0 {
		bNum.sign = 1

		outRunes = append(outRunes, '0')

		if bNum.precision > 0 {
			outRunes = append(outRunes, bNum.decimalSeparator)

			for h := 0; h < int(bNum.precision); h++ {
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

		if startIdx==0 &&
			bNum.sign == -1  &&
			negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, ')')
		}

		modX := big.NewInt(0)
		scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, baseTen, modX)
		outRunes = append(outRunes, rune(modulo.Int64() + int64(48)))
		digitCnt++
		startIdx++

		if bNum.precision > 0 &&
			int(bNum.precision) == startIdx {

			outRunes = append(outRunes, bNum.decimalSeparator)
			startIdx++
		}

	}

	if int(bNum.precision) >= digitCnt {

		delta := int(bNum.precision) - digitCnt + 1

		for k:=0; k < delta; k++ {
			outRunes = append(outRunes, '0')
			startIdx++

			if bNum.precision > 0 &&
				int(bNum.precision) == startIdx {

				outRunes = append(outRunes, bNum.decimalSeparator)
				startIdx++
			}
		}
	}

	startIdx--

	// adjust for negative sign value
	if bNum.sign == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes = append(outRunes, '-')
			startIdx++

		}	else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			outRunes = append(outRunes, '(')
			startIdx+=2
		}
	}

	sortLimit := startIdx / 2
	tRune := rune(0)
	yCnt := 0

	for i:= startIdx; i > sortLimit ; i-- {
		tRune = outRunes[yCnt]
		outRunes[yCnt] = outRunes[i]
		outRunes[i] = tRune
		yCnt++
	}

	return string(outRunes)
}

// FormatThousandsStr - Returns the number string delimited with the
// BigIntNum ThousandsSeparator character plus the Decimal Separator
// character if applicable. See methods BigIntNum.SetThousandsSeparator()
// and BigIntNum.SetDecimalSeparator().
//
// If the Decimal Separator was not previously set for this BigIntNum,
// the Decimal Separator is defaulted to the USA standard period ('.').
// To use another character for Decimal Separator, see method
// BigIntNum.SetDecimalSeparator().
//
// If the Thousands Separator was not previously set for this BigIntNum,
// the Thousands Separator is defaulted to the USA standard comma (',').
// To use another character for Thousands Separator, see method
// BigIntNum.SetThousandsSeparator().
//
// Example:
// numstr = 1000000.234 converted to 1,000,000.234
//
// Input Parameters
// ================
//
// Input Parameters
// ================
//
// negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
//																		LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//																														 		a leading minus sign.
//																																Example: -123,456.78
//
//																		PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//																																surrounding parentheses.
//																																Example: (123,456.78)
//
func (bNum *BigIntNum) FormatThousandsStr(negValMode NegativeValueFmtMode) string {

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	outRunes := make([]rune, 0, 300)

	scratchNum := big.NewInt(0).Set(bNum.absBigInt)
	baseZero := big.NewInt(0)

	if scratchNum.Cmp(baseZero) == 0 {
		bNum.sign = 1

		outRunes = append(outRunes, '0')

		if bNum.precision > 0 {
			outRunes = append(outRunes, bNum.decimalSeparator)

			for h := 0; h < int(bNum.precision); h++ {
				outRunes = append(outRunes, '0')
			}

		}

		return string(outRunes)
	}

	startIdx := 0
	modulo := big.NewInt(0)
	baseTen := big.NewInt(10)
	digitCnt := 0
	thouCnt := -1

	if bNum.precision == 0 {
		thouCnt = 0
	}


	for scratchNum.Cmp(baseZero) == 1 {

		if startIdx==0 &&
			bNum.sign == -1  &&
			negValMode == PARENTHESESNEGVALFMTMODE {

			outRunes = append(outRunes, ')')
		}

		modX := big.NewInt(0)
		scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, baseTen, modX)
		outRunes = append(outRunes, rune(modulo.Int64() + int64(48)))
		digitCnt++
		startIdx++

		if thouCnt > -1 {
			thouCnt++
		}

		if scratchNum.Cmp(baseZero) == 1 &&
			thouCnt==3 {

			outRunes = append(outRunes, bNum.thousandsSeparator)
			startIdx++
			thouCnt = 0
		}

		if bNum.precision > 0 &&
			int(bNum.precision) == startIdx {

			outRunes = append(outRunes, bNum.decimalSeparator)
			startIdx++
			thouCnt = 0
		}

	}

	if int(bNum.precision) >= digitCnt {

		delta := int(bNum.precision) - digitCnt + 1

		for k:=0; k < delta; k++ {
			outRunes = append(outRunes, '0')
			startIdx++

			if bNum.precision > 0 &&
				int(bNum.precision) == startIdx {

				outRunes = append(outRunes, bNum.decimalSeparator)
				startIdx++
			}
		}
	}

	startIdx--

	// adjust for negative sign value
	if bNum.sign == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes = append(outRunes, '-')
			startIdx++

		}	else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			outRunes = append(outRunes, '(')
			startIdx+=2
		}
	}

	sortLimit := startIdx / 2
	tRune := rune(0)
	yCnt := 0

	for i:= startIdx; i > sortLimit ; i-- {
		tRune = outRunes[yCnt]
		outRunes[yCnt] = outRunes[i]
		outRunes[i] = tRune
		yCnt++
	}

	return string(outRunes)

}

// GetAbsoluteNumStr - Returns the absolute integer value (positive value) of the
// *big.Int value encapsulated by this BigIntNum. No decimal point is included.
//
// If an error is encountered, an empty string is returned.
//
func (bNum *BigIntNum) GetAbsoluteNumStr() string {

	nDto, err := NumStrDto{}.NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)

	if err != nil {
		return ""
	}

	return nDto.GetAbsNumStr()
}

// GetAbsoluteNumStrErr - Returns the absolute integer value (positive value) of the
// *big.Int value encapsulated by this BigIntNum. No decimal point is included.
//
// This method is identical to GetAbsoluteNumStr() above, except that this method
// returns an error value.
//
func (bNum *BigIntNum) GetAbsoluteNumStrErr() (string, error) {
	nDto, err := NumStrDto{}.NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)

	if err != nil {
		return "",
		fmt.Errorf("Error returned by NumStrDto{}.NewBigInt(" +
			"big.NewInt(0).Set(bNum.bigInt), bNum.precision). Error='%v'",
				err.Error())
	}

	return nDto.GetAbsNumStr(),nil
}

// GetAbsoluteValue - returns the absolute value of the
// *big.Int value encapsulated by the current BigIntNum.
func (bNum *BigIntNum) GetAbsoluteValue() *big.Int {
	
	return big.NewInt(0).Set(bNum.absBigInt)
}

// GetBigInt - return the numeric value as an integer
// of type *big.int.
//
func (bNum *BigIntNum) GetBigInt() (*big.Int, error) {

	return big.NewInt(0).Set(bNum.bigInt), nil
}

// GetCurrencySymbol - Returns the character currently designated
// as the currency symbol for this BigIntNum instance.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// Example: $123.45
//
func (bNum *BigIntNum) GetCurrencySymbol() rune {

	if bNum.currencySymbol == 0 {
		bNum.currencySymbol = '$'
	}

	return bNum.currencySymbol

}

// GetDecimal - Converts the current BigIntNum value to a Type Decimal
// instance. The resulting number value includes the decimal point
// and decimal digits if they exist.
//
func (bNum *BigIntNum) GetDecimal() (Decimal, error) {

	ePrefix := "BigIntNum.GetDecimal() "

	dec, err := Decimal{}.NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)

	if err != nil {
		return Decimal{},
			fmt.Errorf (ePrefix +
				"Error returned by Decimal{}.NewBigInt(bNum.bigInt, bNum.precision) " +
				"bNum.bigInt='%v' bNum.precision='%v' Error='%v'",
				bNum.bigInt.Text(10), bNum.precision, err.Error())
	}

	bNum.setDefaultSeparators()

	dec.SetSeparators(bNum.decimalSeparator, bNum.thousandsSeparator, bNum.currencySymbol)

	return dec, nil
}


// GetDecimalSeparator - returns the character designated
// as the decimal separator for the current NumStrDto instance.
//
// In the USA, the decimal separator is the period character ('.').
//
// Example:		123.456
//
func (bNum *BigIntNum) GetDecimalSeparator() rune {

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	return bNum.decimalSeparator

}

// GetFractionalPart - Returns the fractional digits of the
// current BigIntNum instance as a new BigIntNum instance
// containing those correctly formatted fractional digits.
//
// Examples
// ========
//
// 			 Current
// 			BigIntNum				 		Return
//  			Value						  Value
// 			----------				---------
//
//  			123.456						 0.456
// 			 -123.456						-0.456
// 			  123								 0
// 			 -123								 0
//
func (bNum *BigIntNum) GetFractionalPart() BigIntNum{

	if bNum.IsZero() {
		return BigIntNum{}.NewBigInt(big.NewInt(0), 0)
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
								big.NewInt(int64(bNum.precision)), nil)

	modulo := big.NewInt(0).Rem(bNum.bigInt, scaleVal)

	return BigIntNum{}.NewBigInt(modulo, bNum.precision)
}

// GetIntAry - Converts the current BigIntNum value to an IntAry
// instance. The resulting number value includes the decimal point
// and fractional digits if they exist.
//
// Note that the BigIntNum settings for 'decimalSeparator', 'thousandsSeparator'
// and 'currencySymbol' are transferred to the new IntAry instance returned to the
// calling function.
//
func (bNum *BigIntNum) GetIntAry() (IntAry, error) {

	ePrefix := "BigIntNum.GetIntAry() "

	ia, err := IntAry{}.NewBigInt(big.NewInt(0).Set(bNum.bigInt), int(bNum.precision))

	if err != nil {
		return IntAry{},
			fmt.Errorf (ePrefix +
				"Error returned by IntAry{}.NewBigInt(bNum.bigInt, bNum.precision) " +
				"bNum.bigInt='%v' bNum.precision='%v' Error='%v'",
				bNum.bigInt.Text(10), bNum.precision, err.Error())
	}

	bNum.setDefaultSeparators()
	ia.SetSeparators(bNum.decimalSeparator, bNum.thousandsSeparator, bNum.currencySymbol)

	return ia, nil
}

// GetIntegerPart - returns a BigIntNum equal to the integer
// value of the current BigIntNum.
// Examples:
//
// 			 Current
// 			BigIntNum				 		Return
//  			Value						  Value
// 			----------				---------
//
//  			123.456						 123
// 			 -123.456						-123
// 			  123								 123
// 			 -123								-123
//
func (bNum *BigIntNum) GetIntegerPart() BigIntNum {

	if bNum.IsZero() {
		return BigIntNum{}.NewBigInt(big.NewInt(0), 0)
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
								big.NewInt(int64(bNum.precision)), nil)

	quotient := big.NewInt(0).Quo(bNum.bigInt, scaleVal)

	return BigIntNum{}.NewBigInt(quotient, 0)
}

// GetNumStr - Converts the current BigIntNum value to string of
// numbers which includes the decimal point and decimal digits
// if they exist.
//
func (bNum *BigIntNum) GetNumStr() (string) {

	return bNum.FormatNumStr(LEADMINUSNEGVALFMTMODE)

}

// GetNumStrDto - Converts the current BigIntNum value to a NumStrDto
// instance. The resulting number string includes the decimal point
// and decimal digits if they exist.
//
func (bNum *BigIntNum) GetNumStrDto() (NumStrDto, error) {

	ePrefix := "BigIntNum.GetNumStrDto() "

	nDto, err := NumStrDto{}.NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)

	if err != nil {
		return NumStrDto{},
			fmt.Errorf (ePrefix +
				"Error returned by NumStrDto{}.NewBigInt(bNum.bigInt, bNum.precision) " +
				"bNum.bigInt='%v' bNum.precision='%v' Error='%v'",
				bNum.bigInt.Text(10), bNum.precision, err.Error())
	}

	bNum.setDefaultSeparators()

	nDto.SetSeparators(bNum.decimalSeparator, bNum.thousandsSeparator, bNum.currencySymbol)

	return nDto, nil
}

// GetPrecision - Returns precision as an integer of
// type 'int'.
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
func (bNum *BigIntNum) GetPrecision() int {
	return int(bNum.precision)
}

// GetPrecisionUint - Returns precision as an unsigned
// integer (uint).
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
func (bNum *BigIntNum) GetPrecisionUint() uint {
	return bNum.precision
}

// GetSign - Returns the numeric sign associated
// with the current numeric value encapsulated by
// this BigIntNum.
func (bNum *BigIntNum) GetSign() int {
	return bNum.sign
}

// GetScaleFactor - Returns the scale value of the current 
// BigIntNum.  Scale value is a function of 'precision' or 
// the number of digits to the right of the decimal place. 
//
// Example:
// precision = 0 		Scale Factor = 10^0   	Scale Factor =    1
// precision = 1		Scale Factor = 10^1			Scale Factor =   10
// precision = 2		Scale Factor = 10^2			Scale Factor =  100
// precision = 3    Scale Factor = 10^3			Scale Factor = 1000
//
func (bNum *BigIntNum) GetScaleFactor() *big.Int {
	return big.NewInt(0).Set(bNum.scaleFactor)
}

// GetThisPointer - Returns a pointer to the current
// instance of this BigIntNum.
func (bNum *BigIntNum) GetThisPointer() *BigIntNum {
	return bNum
}

// GetThousandsSeparator - returns a rune which represents
// the character currently used to separate thousands in
// the display of the current BigIntNum instance.
//
// In the USA, the thousands separator is a comma character.
//
// Example: 1,000,000,000
//
func (bNum *BigIntNum) GetThousandsSeparator() rune {

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	return bNum.thousandsSeparator
}

// IsZero - Returns a boolean signaling whether the current
// BigIntNum value is zero.
func (bNum *BigIntNum) IsZero() bool {

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		return true
	}

	return false
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
// precision int	- This unsigned integer (always a positive value) identifies
// 									the location of the decimal point in the integer value 'bigI'.
// 									The decimal point location is calculated by starting with the
// 									right most digit in the integer number and counting	left,
// 									'precision' places.
//
// 									Example:
//
//											Integer Value		precision			Numeric Value
//											  123456					 3					  123.456
//
func (bNum BigIntNum) NewBigInt(bigI *big.Int, precision uint) BigIntNum {
	b := BigIntNum{}
	b.Empty()
	b.SetBigInt(bigI, precision)
	return b
}

// NewBigIntExponent - New bigInt Exponent returns a new
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
	b.Empty()
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
	b.Empty()

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
	b.Empty()
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
	b.Empty()
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
	b.Empty()
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
	b.Empty()
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
	b.Empty()
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
	b.Empty()
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
	b.Empty()
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
// Example 1:
//	dec, err := Decimal{}.NewNumStr(nStr)
//	bINum, err := BigIntNum{}.NewINumMgr(&dec)
//
// Example 2:
// dec, err := Decimal{}.NewNumStr(nStr)
// bINum, err := BigIntNum{}.NewINumMgr(dec.GetThisPointer())
//
// Example 3:
// dec := Decimal{}.NewPtr()
// err := dec.SetNumStr(nStr)
// bINum, err := BigIntNum{}.NewINumMgr(dec)
//
func (bNum BigIntNum) NewINumMgr(numMgr INumMgr) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewINumMgr() "

	bINum := BigIntNum{}.New()
	bINum.Empty()

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

	b := BigIntNum{}
	err := b.SetNumStr(numStr)

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by b.SetNumStr(numStr). " +
				"numStr='%v' Error='%v'",
						numStr, err.Error())
	}

	return b, nil
}

// NewNumStr - Receives a number string as input and returns
// a new BigIntNum instance.
//
func (bNum BigIntNum) NewNumStrMaxPrecision(
													numStr string,
															maxPrecision uint) (BigIntNum, error) {

	b := BigIntNum{}

	err:= b.SetNumStr(numStr)

	if err != nil {

		ePrefix := "BigIntNum.NewNumStr() "

		return BigIntNum{},
			fmt.Errorf(ePrefix + "Error returned by b.SetNumStr(numStr). " +
				"numStr='%v' Error='%v' ",
					numStr, err.Error())
	}

	if b.precision > maxPrecision {
		b.RoundToDecPlace(maxPrecision)
	}

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
	b.Empty()
	b.SetBigInt(bigI, uint(nDto.GetPrecision()))

	return b, nil
}

// NewOne - Returns a BigIntNum with an integer value of '1' (one).
func (bNum *BigIntNum) NewOne() BigIntNum {
	b := BigIntNum{}
	b.Empty()
	b.SetBigInt(big.NewInt(1), 0)

	return b
}

// NewOne - Returns a BigIntNum with integer value of  '10' (ten).
func (bNum *BigIntNum) NewTen() BigIntNum {

	b := BigIntNum{}
	b.Empty()
	b.SetBigInt(big.NewInt(10), 0)

	return b

}

// New Zero - Returns a BigIntNum with a value of zero. The number
// of zeros created after the decimal place holder (fractional digits)
// is determined by the input parameter 'precision'.
func (bNum *BigIntNum) NewZero(precision uint) BigIntNum {

	b := BigIntNum{}
	b.Empty()
	b.SetBigInt(big.NewInt(10), precision)

	return b

}

// RoundToDecPlace - Rounds the current BigIntNum instance to a specified
// number of decimal places.
//
// 'precision' equals the number of digits to the right of the decimal
// place.
//
// Example:
// 	integer= 123456; precision = 3; Numeric Value= 123.456
//
// If the value of BigIntNum.bigInt is zero ('0'), that zero value will
// remain unaltered. However, the BigIntNum.precision value will be set equal to
// input parameter, 'precision'.
//
// If the number of decimal places specified for rounding ('precision") is
// equal to the current BigIntNum.precision, no action is taken.
//
// If the number of decimal places specified for rounding ('precision') is
// greater than the current BigIntNum.precision value, trailing zeros are added to
// the current BigIntNum.bigInt value and BigIntNum.precision is set equal
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
func (bNum *BigIntNum) RoundToDecPlace(precision uint) {

	if bNum.precision == precision {
		// Nothing to do. Specified 'precision' is already implemented.
		return
	}

	// bigInt == zero, set precision an return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		bNum.CopyIn(BigIntNum{}.NewBigInt(big.NewInt(0), precision))
		return
	}

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bNum.precision < precision {
		deltaPrecision := precision - bNum.precision
		bNum.extendPrecision(deltaPrecision)
		return
	}

	// Must be: bNum.precision >  precision

	bigNumRound5 :=
		BigIntNum{}.NewBigInt(big.NewInt(5), uint(precision + 1))

	bigNumBase := BigIntNum{}.NewBigInt(bNum.absBigInt, bNum.precision)

	result := BigIntMathAdd{}.AddBigIntNums(bigNumBase, bigNumRound5)

	// 10^deltaPrecision
	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
								big.NewInt(int64(bNum.precision - precision)), nil)

	result.bigInt = big.NewInt(0).Quo(result.bigInt, scaleVal)

	if bNum.sign < 0 {
		result.bigInt = big.NewInt(0).Neg(result.bigInt)
	}

	bNum.SetBigInt(result.bigInt, precision)
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
//											Integer Value		precision			Numeric Value
//											  123456					 3					  123.456
//
func (bNum *BigIntNum) SetBigInt(bigI *big.Int, precision uint) {

	bNum.Empty()

	bNum.bigInt = big.NewInt(0).Set(bigI)
	bNum.precision = precision
	base10 := big.NewInt(0).SetInt64(int64(10))
	bigPrecision := big.NewInt(0).SetInt64(int64(bNum.precision))
	bNum.scaleFactor = big.NewInt(0).Exp(base10, bigPrecision, nil)
	baseZero := big.NewInt(0).SetInt64(0)
	result := bNum.bigInt.Cmp(baseZero)

	if result == -1 {
		bNum.sign = -1
		minusOne := big.NewInt(0).SetInt64(-1)
		bNum.absBigInt = big.NewInt(0).Mul(bNum.bigInt, minusOne)
	} else {
		bNum.sign = 1
		bNum.absBigInt = big.NewInt(0).Set(bNum.bigInt)
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

// SetCurrencySymbol - assigns the input parameter rune as the
// currency symbol to be used by the BigIntNum when generating
// number strings for display.
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
//
func (bNum *BigIntNum) SetCurrencySymbol(currencySymbol rune) {

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	bNum.currencySymbol = currencySymbol
}

// SetDecimalSeparator - Assigns a rune or character to the internal
// data field, 'decimalSeparator'. The Decimal Separator is used to
// separate the integer and fractional elements of a number string.
//
// The BigIntNum Type uses this character when generating number strings
// for display.
//
// In the USA, the Decimal Separator is a period character ('.').
//
// Note: If a zero value is submitted as input, the Decimal Separator
// will default to the USA standard period character ('.').
//
// Example: 123.456
//
func (bNum *BigIntNum) SetDecimalSeparator(decimalSeparator rune) {

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	bNum.decimalSeparator = decimalSeparator
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
// Example 1:
//	dec, err := Decimal{}.NewNumStr(nStr)
//	bINum := BigIntNum{}
//	err := bINum.SetINumMgr(&dec)
//
// Example 2:
// dec, err := Decimal{}.NewNumStr(nStr)
// bINum := BigIntNum{}
// err := bINum.SetINumMgr(dec.GetThisPointer())
//
// Example 3:
// dec := Decimal{}.NewPtr()
// err := dec.SetNumStr(nStr)
// bINum := BigIntNum{}
// err := bINum.SetINumMgr(dec)
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

// SetNumStr - Initializes the current BigIntNum instance
// of the numeric value of the number string input parameter.
// A number string is a string of numeric digits which may
// or may not be prefixed with a minus sign ('-'). The numeric
// string of digits may also contain a decimal separator such
// as a period ('.'). The decimal separator may be set by the
// user. See Method BigIntNum.SetDecimalSeparator(). The decimal
// separator is used to separate integer and fractional numeric
// digits within the number string.
//
func (bNum *BigIntNum) SetNumStr(numStr string) error {

	ePrefix := "BigIntNum) SetNumStr() "

	if len(numStr) == 0 {
		return errors.New(ePrefix + "Error: Input parameter 'numStr' is an EMPTY string!")
	}

	baseRunes := []rune(numStr)
	lBaseRunes := len(baseRunes)

	bNum.setDefaultSeparators()

	newSign := 1
	newPrecision := uint(0)
	newAbsBigInt := big.NewInt(0)
	baseTen := big.NewInt(10)
	isStartRunes := false
	isEndRunes := false
	isFractionalValue := false

	for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

		if baseRunes[i] == '+' ||
			baseRunes[i] == ' '  ||
			baseRunes[i] == bNum.thousandsSeparator ||
			baseRunes[i] == bNum.currencySymbol {

			continue

		}

		if baseRunes[i] == ',' && bNum.decimalSeparator != ',' {
			continue
		}

		if isStartRunes == true &&
			isEndRunes == false &&
			isFractionalValue &&
			baseRunes[i] == bNum.decimalSeparator {

			continue
		}

		if baseRunes[i] == '-' &&
			isStartRunes == false && isEndRunes == false &&
			i+1 < lBaseRunes &&
			((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
				baseRunes[i+1] == bNum.decimalSeparator) {

			newSign = -1
			isStartRunes = true
			continue

		} else if isEndRunes == false &&
			baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			newAbsBigInt = big.NewInt(0).Mul(newAbsBigInt, baseTen)

			newAbsBigInt = big.NewInt(0).Add(newAbsBigInt,
					big.NewInt(int64(baseRunes[i]-48)))

			isStartRunes = true

			if isFractionalValue {
				newPrecision++
			}

		} else if isEndRunes == false &&
			i+1 < lBaseRunes &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
			baseRunes[i] == bNum.decimalSeparator {

			isFractionalValue = true
			continue

		} else if isStartRunes && !isEndRunes {

			isEndRunes = true

		}

	}

	bNum.Empty()
	bNum.sign = newSign
	bNum.precision = newPrecision
	bNum.absBigInt = big.NewInt(0).Set(newAbsBigInt)

	if bNum.sign == 1 {
		bNum.bigInt = big.NewInt(0).Set(newAbsBigInt)
	} else {
		bNum.bigInt = big.NewInt(0).Neg(newAbsBigInt)
	}

	bNum.scaleFactor = big.NewInt(0).Exp(baseTen,
											big.NewInt(int64(newPrecision)),
											nil)

	return nil
}


// SetPrecision - Sets a new 'precision' value for the current
// BigIntNum instance. The new 'precision' is specified by the
// uint type input parameter, 'newPrecision'.
//
// Precision is defined as the number of numeric digits to right
// of the decimal place.
//
// If 'newPrecision' is equal to the current BigIntNum.precision value,
// no action is taken and the original BigIntNum numeric value remains
// unchanged.
//
// If 'newPrecision' is greater than the current BigIntNum.precision
// value, trailing zeros are added to the fractional digits to the
// right of the decimal place.
//
// If 'newPrecision' is less than the current BigIntNum precision
// value, the current BigIntNum numeric value is rounded to the
// specified 'newPrecision' value.
//
// Examples:
//
// 	 Original       			'newPrecision'				Resulting
//    Value								input parameter			  Value
//  --------------				---------------     -------------
//	654.123456									9							 654.123456000
//	654.123456									4							 654.1235
// -654.123456									9							-654.123456000
// -654.123456									4							-654.1235
//		0													3								 0.000
//    0.000000									0								 0
//
func (bNum *BigIntNum) SetPrecision(newPrecision uint) {

	if newPrecision == bNum.precision {
		return
	}

	if bNum.precision > newPrecision {
		bNum.RoundToDecPlace(newPrecision)
		return
	}

	deltaPrecision := newPrecision - bNum.precision
	// bNum.precision must be less than newPrecision
	bNum.extendPrecision(deltaPrecision)
	return

}

// SetThousandsSeparator - Sets the value of the character which will be
// used to separate thousands in the display of the NumStrDto number
// string. In the USA the typical thousands separator is the comma.
//
// If if a zero value is submitted, the Thousands Separator will default
// to the comma character.
//
// Example:
// 1,000,000
//
func (bNum *BigIntNum) SetThousandsSeparator(thousandsSeparator rune) {

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	bNum.thousandsSeparator = thousandsSeparator

}

// TrimTrailingFracZeros - This method will delete non-significant
// trailing zeros from the fractional digits of the current BigIntNum
// numerical value.
//
// Examples:
//  Initial Value			Trimmed Value
//		456.123000 			 456.123
//			0.000					 0
//			7.0						 7
//	 -456.123000			-456.123
//
func (bNum *BigIntNum) TrimTrailingFracZeros(){

	if bNum.precision == 0 {
		return
	}

	biBaseZero := big.NewInt(0)

	if bNum.bigInt.Cmp(biBaseZero) == 0 {
		bNum.precision = 0
		bNum.scaleFactor = big.NewInt(1)
		return
	}


	biBase10 := big.NewInt(10)

	mod10 := big.NewInt(0).Mod(bNum.bigInt, biBase10)
	doReset := false

	for mod10.Cmp(biBaseZero) == 0 && bNum.precision > 0 {
		bNum.bigInt = big.NewInt(0).Quo(bNum.bigInt, biBase10)
		bNum.precision --
		doReset = true
		mod10 = big.NewInt(0).Mod(bNum.bigInt, biBase10)
	}

	if doReset {
		if bNum.sign < 0 {
			bNum.absBigInt = big.NewInt(0).Neg(bNum.bigInt)
		} else {
			bNum.absBigInt = big.NewInt(0).Set(bNum.bigInt)
		}

		bigPrecision := big.NewInt(0).SetInt64(int64(bNum.precision))
		bNum.scaleFactor = big.NewInt(0).Exp(biBase10, bigPrecision, nil)

	}

	return
}

// SetSeparators - Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current number string.
//
// Different nations and cultures use different symbols to delimit numerical values. In the
// USA and many other countries, a period character ('.') is used to delimit integer and
// fractional digits within a numeric value (123.45). Likewise, thousands may be delimited
// by a comma (','). Currency signs very by nationality. For instance, the USA, Canada and
// several other countries use the dollar sign ($) as a currency symbol.
//
// For a list of major world currency symbols see:
// 	MikeAustin71\mathopsgo\mathops\mathopsconstants.go
//  http://www.xe.com/symbols.php
//
// Note: If zero values are submitted as input for separator values, those values will default
// to USA standards.
//
// USA Examples:
//
// Decimal Separator period ('.') 		= 123.456
// Thousands Separator comma (',') 		= 1,000,000,000
// Currency Symbol dollar sign ('$')	= $123
//
func (bNum *BigIntNum) SetSeparators(decimalSeparator, thousandsSeparator, currencySymbol rune) {

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	if thousandsSeparator == 0 {
		thousandsSeparator = ','
	}

	if currencySymbol == 0 {
		currencySymbol = '$'
	}

	bNum.decimalSeparator = decimalSeparator
	bNum.thousandsSeparator = thousandsSeparator
	bNum.currencySymbol = currencySymbol
}

// TruncToDecPlace - Truncates the current BigIntNum to the number
// of decimal places specified by input parameter 'precision'.
// No rounding occurs, the trailing digits are simply truncated or
// deleted in order to achieve the specified number of decimal places.
//
// 'precision' equals the number of digits to the right of the decimal
// place.
//
// If the value of BigIntNum.bigInt is zero ('0'), that zero value will
// remain unaltered. However, BigIntNum.precision will be set equal to
// input parameter, 'precision'.
//
// If the number of decimal places specified for truncation ('precision") is
// equal to the current BigIntNum.precision, no action is taken and the
// original BigIntNum numeric value remains unchanged.
//
// If the number of decimal places specified for truncation ('precision') is
// greater than the current BigIntNum.precision, trailing zeros are added to
// the current BigIntNum.bigInt value and BigIntNum.precision is set equal
// to input parameter, 'precision'.
//
// If 'precision' is less than the current BigIntNum.precision
// value, the current BigIntNum numeric value is truncated to
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
func (bNum *BigIntNum) TruncToDecPlace(precision uint) {

	if bNum.precision == precision {
		// Nothing to do. Specified 'precision' is already implemented.
		return
	}

	// bigInt == zero, set precision an return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		bNum.precision = precision
		return
	}

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bNum.precision < precision {

		bNum.extendPrecision(precision - bNum.precision)
		return
	}

	// Must be bNum.precision > precision
	base10 := big.NewInt(10)
	deltaPrecision := big.NewInt(int64(bNum.precision - precision))
	newBigInt := big.NewInt(0).Set(bNum.absBigInt)
	newScaleVal := big.NewInt(0).Exp(base10, deltaPrecision, nil)
	newBigInt = big.NewInt(0).Quo(newBigInt, newScaleVal)

	if bNum.sign < 1 {
		newBigInt = big.NewInt(0).Neg(newBigInt)
	}

	bNum.SetBigInt(newBigInt, precision)
}

// extendPrecision - Extends the current precision.
//
// Precision is the number of fractional digits to the right
// of the decimal place. This method will extend the number of
// digits to the right of the decimal place by adding trailing
// zeros to the current numeric value of this 'BigIntNum' instance.
// The number of trailing zeros to be added is determined by the
// input parameter, 'deltaPrecision'.
//
func (bNum *BigIntNum) extendPrecision(deltaPrecision uint) {

	if deltaPrecision == 0 {
		return
	}

	newPrecision := bNum.precision + deltaPrecision

	// bigInt == zero, set precision an return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		bNum.CopyIn(BigIntNum{}.NewBigInt(big.NewInt(0), newPrecision))
		return
	}

	base10 := big.NewInt(10)
	scaleVal := big.NewInt(0).Exp(base10,big.NewInt(int64(deltaPrecision)), nil)
	bigINum := big.NewInt(0).Set(bNum.bigInt)

	bigINum = big.NewInt(0).Mul(bigINum, scaleVal)

	bNum.CopyIn(BigIntNum{}.NewBigInt(bigINum, newPrecision))
}


// setDefaultSeparators - Sets default characters for
// decimal separator, thousands separator and currency
// symbol if those variables have not been previously
// assigned values.
//
func (bNum *BigIntNum) setDefaultSeparators() {

	if bNum.decimalSeparator == 0 {
		bNum.decimalSeparator = '.'
	}

	if bNum.thousandsSeparator == 0 {
		bNum.thousandsSeparator = ','
	}

	if bNum.currencySymbol == 0 {
		bNum.currencySymbol = '$'
	}
}