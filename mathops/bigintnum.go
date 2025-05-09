package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strconv"
)

// BigIntNum - wraps a *big.Int integer and its associated
// precision and sign Value. While the numeric value is
// stored as an integer of type *big.Int, the BigIntNum
// type is capable of storing decimal fractions.
//
//	All methods associated with this type all assume that
//	the *big.Int value stored by the BigIntNum Type is configured
//
// in base 10.
//
// INumMgr
// ========
//
// The BigIntNum Type implements the INumMgr interface.
//
// Source Code Repository:
// =======================
// https://github.com/MikeAustin71/mathopsgo.git
//
// Local File:
// ===========
// MikeAustin71\mathopsgo\mathops\bigintnum.go
type BigIntNum struct {
	bigInt                 *big.Int
	absBigInt              *big.Int
	precision              uint     // Number of digits to the right of the decimal place.
	scaleFactor            *big.Int // Scale Factor =  10^(precision)
	numberOfExpectedDigits *big.Int // Number of digits in the 'absBigInt' value
	sign                   int      // Valid values are -1 or +1. Indicates the sign of the
	// Numeric Separators
	decimalSeparator   rune // Character used to separate integer and fractional digits ('.')
	thousandsSeparator rune // Character used to separate thousands (1,000,000,000
	currencySymbol     rune // Currency Symbol
}

// Ceiling - Returns the ceiling integer value of the current BigIntNum
// instance.
//
// Ceiling is defined as: The least, or lowest value integer, which is greater
// than or equal to the numeric value of the current BigIntNum.
// Reference Wikipedia:
//
//	https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
// Examples
// ========
//
//							Initial 		 Ceiling
//	 					 Value				Value
//							-------      -------
//	 						5.95					6
//	 						5.05					6
//	 						5							5
//						 -5.05			 	 -5
//	 						2.4				  	3
//	 						2.9					 	3
//						 -2.7				 	 -2
//						 -2						 -2
func (bNum *BigIntNum) Ceiling() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumElectron).bigIntNumCeiling(
		bNum, ePrefix)
}

// ChangeSign - Changes the sign of the current BigIntNum value.
//
// If the value of BigIntNum is zero, the sign will remain unchanged
// and this method will return with no action taken.
//
// If the sign of the current BigIntNum value is positive (+), the sign
// will be changed to negative (-). Likewise, if the current sign is
// negative (-), the sign will be changed to positive (+).
//
// NOTE:
// This method will first test the current instance of BigIntNum
// to determine if that instance is valid, or not. If the current
// BigIntNum fails the validity test, an error will be returned.
func (bNum *BigIntNum) ChangeSign() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.ChangeSign",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Testing 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumUtility).bigIntNumChangeSign(
		bNum,
		ePrefix)
}

// Cmp - Performs a comparison of two BigIntNum numeric values
// and returns an integer value indicating the relationship
// between the two numeric values (i.e. greater than, less than,
// or equal).
//
// Note: Unlike method CmpBigInt() below, this method does more than
// just compare the root *big.Int. In making the comparision, this
// method takes into account, numeric sign values and precision. Therefore,
// this method effectively compares numeric values. As such, this method
// provides a true and comprehensive picture of the relationship between
// two BigIntNum values.
//
// Return Values:
// bNum == bigIntNum 				Return  0
// bNum > bigIntNum					Return +1
// bNum < bigIntNum					Return -1
func (bNum *BigIntNum) Cmp(bigIntNum BigIntNum) (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Cmp",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).bigIntNumCmp(
		bNum,
		&bigIntNum,
		ePrefix)
}

// CmpBigInt - Compares the value of the *big.Int integer to that
// contained in an incoming BigIntNum.
//
// For a true comparison of BigIntNum values see Method 'Cmp', above.
//
// Return Values:
// bNum == bigIntNum 				Return  0
// bNum > bigIntNum					Return +1
// bNum < bigIntNum					Return -1
func (bNum *BigIntNum) CmpBigInt(bigIntNum BigIntNum) (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.CmpBigInt",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).cmpBigInt(
		bNum,
		&bigIntNum,
		ePrefix)
}

// CopyIn - Receives an incoming BigIntNum type and
// copies the value into the current BigIntNum instance.
func (bNum *BigIntNum) CopyIn(bigN *BigIntNum) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.CopyIn",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumUtility).bigIntNumCopyIn(
		bNum,
		bigN,
		ePrefix)
}

// CopyOut - Makes a deep copy of the current BigIntNum instance
// and returns it as a new BigIntNum instance.
func (bNum *BigIntNum) CopyOut() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.CopyOut",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)
}

// Decrement - Subtracts a value of +1 (plus one) from the numeric
// value of the current BigIntNum instance.
//
// The numeric separators (decimal separator, thousands separator
// and currency symbol) from the original BigIntNum will remain
// unchanged.
func (bNum *BigIntNum) Decrement() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Decrement",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumProton).bigIntNumDecrement(
		bNum,
		ePrefix)
}

// Divide - Performs a division operation. The current BigIntNum instance is the 'dividend'
// is divided by the input parameter, 'divisor'. The result of this division operation is
// the 'fracQuotient' which is returned as a BigIntNum type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//		bNum = dividend
//	 -----------------------------
//		dividend / divisor = quotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// quotient. Precision is defined as the number of fractional digits to the right of
// the decimal place. Be advised that these calculations can support very large precision
// values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
func (bNum *BigIntNum) Divide(
	divisor BigIntNum,
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Divide",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumDivide(
		bNum,
		&divisor,
		maxPrecision,
		ePrefix)
}

// DivideByFive - Divides the numerical value of the current BigIntNum by five ('5'). The
// result of this division operation is the 'fracQuotient' which is returned as a BigIntNum
// type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 5 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
func (bNum *BigIntNum) DivideByFive(
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.DivideByFive",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumDivideByFive(
		bNum,
		maxPrecision,
		ePrefix)
}

// DivideByTen - Divides the numerical value of the current BigIntNum by ten ('10'). The
// result of this division operation is the 'fracQuotient' which is returned as a BigIntNum
// type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 10 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
func (bNum *BigIntNum) DivideByTen(
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.DivideByTen",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumDivideByTen(
		bNum,
		maxPrecision,
		ePrefix)
}

// DivideByTenToPower - Divides the numerical value of the current BigIntNum
// instance by 10 raised to the power of the input parameter, 'exponent'.
//
//	bNum = bNum / (10^exponent)
//
// The original value of the current BigIntNum will be destroyed and overwritten
// by this method.
//
// The final BigIntNum will retain the original numeric separators (decimal separator,
// thousands separator and currency symbol).
func (bNum *BigIntNum) DivideByTenToPower(exponent uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.DivideByTenToPower",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumProton).bigIntNumDivideByTenToPower(
		bNum,
		exponent,
		ePrefix)
}

// DivideByThree - Divides the numerical value of the current BigIntNum by three ('3').
// The result of this division operation is the 'fracQuotient' which is returned as a
// BigIntNum type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 3 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
func (bNum *BigIntNum) DivideByThree(
	maxPrecision uint) (fracQuotient BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.DivideByThree",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumDivideByThree(
		bNum,
		maxPrecision,
		ePrefix)
}

// DivideByTwo - Divides the numerical value of the current BigIntNum by two ('2'). The
// result of this division operation is the 'fracQuotient' which is returned as a BigIntNum
// type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 2 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
func (bNum *BigIntNum) DivideByTwo(
	maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.DivideByTwo",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumDivideByTwo(
		bNum,
		maxPrecision,
		ePrefix)
}

// DivideByTwoQuoMod - Divides the numerical value of the current BigIntNum by two ('2').
// The result of the division operation is returned as an integer quotient, 'intQuotient',
// and a floating point modulo or remainder, 'modulo'.
//
//	bNum / 2 = integer quotient and floating point modulo
//
// If 'modulo' equals zero ('0'), it signals the the current BigIntNum numerical value is
// 'even'; that is, it is evenly divisible by two.
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// floating point 'modulo'. Precision is defined as the the number of fractional digits to
// the right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// Both returned BigIntNum 'intQuotient' and 'modulo' BigIntNum types will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied from the
// current BigIntNum instance (bNum).
func (bNum *BigIntNum) DivideByTwoQuoMod(
	maxPrecision uint) (intQuotient BigIntNum, modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.DivideByTwoQuoMod",
		"")

	if err != nil {
		return intQuotient, modulo, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return intQuotient, modulo, err
	}

	return new(bigIntNumProton).bigIntNumDivideByTwoQuoMod(
		bNum,
		maxPrecision,
		ePrefix)
}

// Empty - Resets the BigIntNum data fields to their
// uninitialized or zero state.
func (bNum *BigIntNum) Empty() {

	new(bigIntNumElectron).empty(bNum)

}

// Equal - Compares two BigIntNum instances and returns 'true'
// if the two instances are equal in all respects.
//
// Be careful, two BigIntNum instances could have equal
// values with different precisions. In that case this
// method would return 'false'. To test for equivalent
// values, see method BigIntNum.EqualValue(), below.
//
// If they are not Equal, the method returns 'false'.
func (bNum *BigIntNum) Equal(b2 BigIntNum) (bool, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Equal",
		"")

	if err != nil {
		return false, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return false, err
	}

	return new(bigIntNumElectron).bigIntNumEqual(
		bNum,
		&b2,
		ePrefix)
}

// EqualValue - Compares the values of the current BigIntNum instance
// and the input parameter BigIntNum, 'b2'. If the two numeric values
// are equal, this method returns 'true'.
func (bNum *BigIntNum) EqualValue(b2 BigIntNum) (bool, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.EqualValue",
		"")

	if err != nil {
		return false, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return false, err
	}

	return new(bigIntNumElectron).bigIntNumEqualValue(
		bNum,
		&b2,
		ePrefix)
}

// ExtendPrecision - Extends the current precision.
//
// Precision is the number of fractional digits to the right
// of the decimal place. This method will extend the number of
// digits to the right of the decimal place by adding trailing
// zeros to the current numeric value of this 'BigIntNum' instance.
// The number of trailing zeros to be added is determined by the
// input parameter, 'deltaPrecision'.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) ExtendPrecision(deltaPrecision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.ExtendPrecision",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumProton).bigIntNumExtendPrecision(
		bNum,
		deltaPrecision,
		ePrefix)
}

// Floor - returns the greatest integer less than or equal to
// the numeric value of the current BigIntNum. Reference Wikipedia,
// https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//							Initial 			Floor
//	 					 Value				Value
//							-------      -------
//	 						5.95					5
//	 						5.05					5
//	 						5							5
//						 -5.05			 	 -6
//	 						2.4				  	2
//	 						2.9					 	2
//						 -2.7				 	 -3
//						 -2					 	 -2
func (bNum *BigIntNum) Floor() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Floor",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumFloor(
		bNum,
		ePrefix)
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
//
//	LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//													 		a leading minus sign.
//															Example: -$123,456.78
//
//	PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//															surrounding parentheses.
//															Example: ($123,456.78)
//
//
//	ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//															absolute (positive) integer value
//															and no decimal place separator.
//															Example: ($12,345,678)
func (bNum *BigIntNum) FormatCurrencyStr(
	negValMode NegativeValueFmtMode) (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.FormatCurrencyStr()",
		"")

	if err != nil {
		return "", err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {

		return "", err
	}

	return new(bigIntNumMolecule).formatCurrencyStr(
		bNum,
		negValMode,
		ePrefix)
}

// FormatNumStr - Formats the numeric value of the current BigIntNum
// instance as number string consisting of integer digits to the left
// of the decimal place and fractional digits to the right of the decimal
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
//
//	LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//													 		a leading minus sign.
//															Example: -123456.78
//
//	PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//															surrounding parentheses.
//															Example: (123456.78)
//
//	ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//															absolute (positive) integer value
//															and no decimal place separator.
//															Example: (12345678)
func (bNum *BigIntNum) FormatNumStr(
	negValMode NegativeValueFmtMode) (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.FormatNumStr",
		"")

	if err != nil {
		return "", err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Testing Validity of 'bNum'"))

	if err != nil {

		return "", err
	}

	return new(bigIntNumMolecule).formatBigIntNumStr(
		bNum,
		negValMode,
		ePrefix)
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
// numStr = 1000000.234 converted to 1,000,000.234
//
// Input Parameters
// ================
//
// Input Parameters
// ================
//
// negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
//
//	LEADMINUSNEGVALFMTMODE 		-	Negative values formatted with
//													 		a leading minus sign.
//															Example: -123,456.78
//
//	PARENTHESESNEGVALFMTMODE	-	Negative values formatted with
//															surrounding parentheses.
//															Example: (123,456.78)
//
//
//	ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//															absolute (positive) integer value
//															and no decimal place separator.
//															Example: (12,345,678)
func (bNum *BigIntNum) FormatThousandsStr(
	negValMode NegativeValueFmtMode) (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.FormatThousandsStr",
		"")

	if err != nil {
		return "", err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {

		return "", err
	}

	return new(bigIntNumMolecule).formatThousandsStr(
		bNum,
		negValMode,
		ePrefix)

}

// GetActualNumberOfDigits - Returns the number of numeric digits
// in the absolute value of this BigIntNum instance. In addition,
// a boolean value is returned indicating whether the absolute value
// is zero.
//
// Examples
// ========
//
//	      123.45														5
//	1,234,567                               7
//
// -1,234,567                               7
//
//					 0															1
//	         0.00                           1
//	       012.34                           4
//	         0.1234													4
//	       - 0.1234													4
//	         0.123400												4
//	         0.0123400											4
//	 1,234,567.800												  8
//	         5                              1
func (bNum *BigIntNum) GetActualNumberOfDigits() (
	numberOfDigits *big.Int, isZeroValue bool, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetActualNumberOfDigits",
		"")

	numberOfDigits = big.NewInt(0)

	if err != nil {
		return numberOfDigits, false, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Testing Validity of 'bNum'"))

	if err != nil {

		return numberOfDigits, false, err
	}

	return new(bigIntNumMolecule).getActualNumberOfDigits(
		bNum,
		ePrefix)
}

// GetAbsoluteNumStr - Returns the absolute integer value (positive value) of the
// *big.Int value encapsulated by this BigIntNum. No decimal place is included.
//
// If an error is encountered, an empty string is returned.
func (bNum *BigIntNum) GetAbsoluteNumStr() (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetAbsoluteNumStr",
		"")

	if err != nil {
		return "", err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return "", err
	}

	return new(bigIntNumProton).bigIntNumGetAbsoluteNumStr(
		bNum, ePrefix)
}

// GetAbsoluteBigIntNumValue - Returns the absolute numeric value
// of this BigIntNum instance as a new BigIntNum Type.
//
// If the current BigIntNum value is'-123.456', this method will
// return '123.456'.
//
// If the current BigIntNum value is'123.456', this method will
// return '123.456'.
func (bNum *BigIntNum) GetAbsoluteBigIntNumValue() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetAbsoluteBigIntNumValue",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumGetAbsoluteBigIntNumValue(
		bNum,
		ePrefix)
}

// GetAbsoluteBigIntValue - returns the absolute value of the
// *big.Int value encapsulated by the current BigIntNum.
func (bNum *BigIntNum) GetAbsoluteBigIntValue() (*big.Int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetAbsoluteBigIntNumValue",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewInt(0), err
	}

	return new(bigIntNumProton).bigIntNumGetAbsoluteBigIntValue(
		bNum,
		ePrefix)
}

// GetBigFloat - Returns the numeric value of the current
// BigIntNum as *big.Float type.
func (bNum *BigIntNum) GetBigFloat() (*big.Float, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetBigFloat",
		"")

	if err != nil {
		return big.NewFloat(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewFloat(0), err
	}

	return new(bigIntNumProton).bigIntNumGetBigFloat(
		bNum, ePrefix)
}

// GetBigInt - return the numeric value as an integer
// of type *big.int.
func (bNum *BigIntNum) GetBigInt() (*big.Int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetBigInt",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewInt(0), err
	}

	return new(bigIntNumProton).bigIntNumGetBigInt(
		bNum,
		ePrefix)
}

// GetBigIntNum
// Return a copy of the value represented by the
// current instance of BigIntNum.
//
// This method is required by the INumMgr interface
func (bNum *BigIntNum) GetBigIntNum() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)
}

// GetBigIntFixedDecimal - returns a BigIntFixedDecimal instance
// which contains a copy of the current BigIntNum values.
func (bNum *BigIntNum) GetBigIntFixedDecimal() (BigIntFixedDecimal, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetBigIntFixedDecimal",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	return new(bigIntNumProton).bigIntNumGetBigIntFixedDecimal(
		bNum,
		ePrefix)
}

// GetBigRat - Returns the numeric value of the current
// BigIntNum as a *big.Rat type.
func (bNum *BigIntNum) GetBigRat() (*big.Rat, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetBigRat",
		"")

	if err != nil {
		return big.NewRat(1, 1), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewRat(1, 1), err
	}

	return new(bigIntNumProton).bigIntNumGetBigRat(
		bNum, ePrefix)
}

// GetCurrencySymbol - Returns the character currently designated
// as the currency symbol for this BigIntNum instance.
//
// If the current instance of BigIntNum is determined to be
// invalid, an error will be returned.
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// USA Example: $123.45
func (bNum *BigIntNum) GetCurrencySymbol() (rune, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetCurrencySymbol",
		"")

	if err != nil {
		return '0', err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return '0', err
	}

	var nSepSymbol NumSepSymbolCode

	nSepSymbol = CURRENCYSYMBOL

	return new(bigIntNumProton).bigIntNumGetNumSepSymbol(
		bNum, nSepSymbol, ePrefix)

}

// GetDecimal - Converts the current BigIntNum value to a Type Decimal
// instance. The resulting number value includes the decimal place
// and decimal digits if they exist.
//
// The returned Decimal instance contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum
// instance, 'bNum'.
//
// This method performs a validity test on the current BigIntNum instance.
func (bNum *BigIntNum) GetDecimal() (Decimal, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetDecimal",
		"")

	if err != nil {
		return Decimal{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return Decimal{}, err
	}

	return new(bigIntNumProton).decimalGetDecimal(
		bNum, ePrefix)
}

// GetDecimalSeparator - returns the character designated
// as the decimal separator for the current NumStrDto instance.
//
// In the USA, the decimal separator is the period character ('.').
//
// Example:		123.456
func (bNum *BigIntNum) GetDecimalSeparator() (rune, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetDecimalSeparator",
		"")

	if err != nil {
		return '0', err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return '0', err
	}

	var nSepSymbol NumSepSymbolCode

	nSepSymbol = DECIMALSYMBOL

	return new(bigIntNumProton).bigIntNumGetNumSepSymbol(
		bNum, nSepSymbol, ePrefix)
}

// GetExpectedNumberOfDigits - Returns the number of expected numeric
// digits associated with this BigIntNum instance. The returned value
// is stored in data field, BigIntNum.numberOfExpectedDigits. The value
// is set by calling method BigIntNum.SetExpectedNumberOfDigits().
//
// This value is useful in tracking leading zeros.
func (bNum *BigIntNum) GetExpectedNumberOfDigits() (*big.Int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetExpectedNumberOfDigits",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewInt(0), err
	}

	return bNum.numberOfExpectedDigits, nil
}

// GetFractionalPart - Returns the fractional digits of the
// current BigIntNum instance as a new BigIntNum instance
// containing those correctly formatted fractional digits.
//
// Examples
// ========
//
//				 Current
//				BigIntNum				 		Return
//	 			Value						  Value
//				----------				---------
//
//	 			123.456						 0.456
//				 -123.456						-0.456
//				  123								 0
//				 -123								 0
func (bNum *BigIntNum) GetFractionalPart() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetFractionalPart",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumGetFractionalPart(
		bNum,
		ePrefix)
}

// GetInt - Returns a type 'int' containing the 32-bit integer
// value of the current BigIntNum instance.
//
// If the current BigIntNum value is greater than the maximum
// 'int' value, the maximum 32-bit integer value is returned
// in addition to an 'error'.
//
// If the current BigIntNum value is less than the minimum 'int'
// value, the minimum 32-bit integer value is returned along with
// an 'error'.
func (bNum *BigIntNum) GetInt() (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetInt",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).bigIntNumGetInt(
		bNum,
		ePrefix)
}

// GetIntAry - Converts the current BigIntNum value to an IntAry
// instance. The resulting number value includes the decimal place
// and fractional digits if they exist.
//
// Note that the BigIntNum settings for 'decimalSeparator', 'thousandsSeparator'
// and 'currencySymbol' are transferred to the new IntAry instance returned to the
// calling function.
//
// The returned IntAry type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum
// instance.
//
// This method performs a validity test on the current BigIntNum instance.
func (bNum *BigIntNum) GetIntAry() (IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetIntAry",
		"")

	if err != nil {
		return IntAry{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return IntAry{}, err
	}

	return new(bigIntNumProton).bigIntNumGetIntAry(
		bNum, ePrefix)
}

// GetIntegerPart - returns a BigIntNum equal to the integer
// value of the current BigIntNum.
// Examples:
//
//					 Current
//					BigIntNum				 		Return
//		 			Value						  Value
//					----------				---------
//
//	         123.456						 123
//					 -123.456						-123
//					  123								 123
//					 -123								-123
func (bNum *BigIntNum) GetIntegerPart() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetIntegerPart",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumAtom).getIntegerPart(
		bNum,
		ePrefix)
}

// GetIntegerValue
// Returns the internal *big.Int number for the current
// BigIntNum instance.
func (bNum *BigIntNum) GetIntegerValue() (*big.Int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetIntegerValue",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewInt(0), err
	}

	return new(bigIntNumProton).bigIntNumGetIntegerValue(
		bNum, ePrefix)
}

// GetInverse - Returns the value of one (1) divided by the current
// BigIntNum instance as a new BigIntNum Type.
func (bNum *BigIntNum) GetInverse(maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetInverse",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumProton).bigIntNumGetInverse(
		bNum, maxPrecision, ePrefix)
}

// GetNumberOfDigits
//
// Returns the number of digits in the numeric value of the
// current BigIntNum instance. The count only includes numeric
// digits and as such, EXCLUDES number signs ('-' or '+' ),
// thousands separators (',') and decimal separators.
//
// Examples:
// =========
//
//	Result=
//
// Numeric String           Number of
//
//	Value                Numeric Digits
//
// =============           ==============
//
//	         123.45							5
//		  1,234,567                 7
//	  -1,234,567.8				 			  8
//		          0									1
//	           0.00              1
//	         012.34              4
//	           0.1234						4
//	           0.123400					6
//	           0.0123400					6
//	   1,234,567.800						 10
//	           5                 1
//
// Note:
//
//	The returned integer number will always be a positive number.
//	Also, GetActualNumberOfDigits() will be faster for larger
//	numbers.
func (bNum *BigIntNum) GetNumberOfDigits() (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetNumberOfDigits",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Testing Validity of 'bNum'"))

	if err != nil {

		return 0, err
	}

	return new(bigIntNumUtility).getBigIntNumOfDigits(
		bNum,
		ePrefix)
}

// GetNumericSeparatorsDto
//
// Returns a structure containing the character or rune values
// for decimal place separator, thousands separator and
// currency symbol.
func (bNum *BigIntNum) GetNumericSeparatorsDto() (NumericSeparatorDto, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetNumericSeparatorsDto",
		"")

	if err != nil {
		return NumericSeparatorDto{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Testing Validity of 'bNum'"))

	if err != nil {

		return NumericSeparatorDto{}, err
	}

	return new(bigIntNumAtom).getNumericSeparatorsDto(
		bNum,
		ePrefix)
}

// GetNumStr
//
// Converts the current BigIntNum value to string of numbers
// which includes the decimal place and decimal digits, if
// they exist.
func (bNum *BigIntNum) GetNumStr() (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetNumStr",
		"")

	if err != nil {
		return "", err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return "", err
	}

	return new(bigIntNumAtom).getBigIntNumStr(
		bNum,
		ePrefix)
}

// GetNumStrDto
//
// Converts the current BigIntNum value to a NumStrDto instance.
// The resulting number string includes the decimal place and
// decimal digits if they exist.
//
// The returned NumStrDto type contains numeric separators (decimal
// separator, thousands separator, and currency symbol) copied from
// the current BigIntNum instance.
//
// This method performs a validity test on the current BigIntNum instance.
func (bNum *BigIntNum) GetNumStrDto() (NumStrDto, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetNumStrDto",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return NumStrDto{}, err
	}

	return new(bigIntNumProton).bigIntNumGetNumStrDto(
		bNum, ePrefix)
}

// GetPrecision
//
// Returns the precision associated with the current
// instance of BigIntNum as an integer of type 'int'.
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal place in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//					1.234    	GetPrecision() = 3
//							5			GetPrecision() = 0
//				0.12345  		GetPrecision() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
func (bNum *BigIntNum) GetPrecision() (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetPrecision",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).bigIntNumGetPrecision(
		bNum, ePrefix)
}

// GetPrecisionBigInt - Returns the 'precision' of the current
// BigIntNum instance as a *big.Int Type.
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal place in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//		1.234    	GetPrecisionBigInt() = 3
//				5			GetPrecisionBigInt() = 0
//	0.12345  		GetPrecisionBigInt() = 5
func (bNum *BigIntNum) GetPrecisionBigInt() (*big.Int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetPrecisionBigInt",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewInt(0), err
	}

	return new(bigIntNumProton).bigIntNumGetPrecisionBigInt(
		bNum, ePrefix)

}

// GetPrecisionUint - Returns precision as an unsigned
// integer (uint).
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal place in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//					1.234    	GetPrecisionUint() = 3
//							5			GetPrecisionUint() = 0
//				0.12345  		GetPrecisionUint() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
func (bNum *BigIntNum) GetPrecisionUint() (uint, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetPrecisionUint",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).bigIntNumGetPrecisionUint(
		bNum, ePrefix)
}

// GetScaleFactor
//
// Returns the scale value of the current instance of
// BigIntNum.  Scale value is a function of 'precision' or
// the number of digits to the right of the decimal place.
// Therefore, scale factor is defined by 10 raised to the
// power BigIntNum precision.
//
// Example:
// precision = 0 		Scale Factor = 10^0   	Scale Factor =    1
// precision = 1		Scale Factor = 10^1			Scale Factor =   10
// precision = 2		Scale Factor = 10^2			Scale Factor =  100
// precision = 3    Scale Factor = 10^3			Scale Factor = 1000
func (bNum *BigIntNum) GetScaleFactor() (*big.Int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetScaleFactor",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewInt(0), err
	}

	return new(bigIntNumProton).bigIntNumGetScaleFactor(
		bNum, ePrefix)
}

// GetSign - Returns the numeric sign associated
// with the current numeric value encapsulated by
// this BigIntNum.
//
// GetSign() returns:
//
//	-1 if x < 0;
//	 0 if x == 0;
//	+1 if x > 0.
func (bNum *BigIntNum) GetSign() (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetSign",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).bigIntNumGetSign(
		bNum, ePrefix)
}

// GetSignedBigInt
// Returns the integer value of the current BigIntNum
// as a signed *big.Int Type.
func (bNum *BigIntNum) GetSignedBigInt() (*big.Int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetSignedBigInt",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return big.NewInt(0), err
	}

	return new(bigIntNumProton).bigIntNumGetSignedBigInt(
		bNum, ePrefix)
}

// GetSciNotationNumber
//
// Converts the numeric value of the current BigIntNum instance
// into scientific notation and returns this value as an instance
// of type SciNotationNum.
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//
//											scientific notation string. If the value of 'mantissaLen'
//											is less than two ('2'), this method will automatically set
//											the 'mantissaLen' to a default value of two ('2').
//
//											Example Scientific Notation:
//											----------------------------
//
//	 										scientific notation string: '2.652e+8'
//
//	 										significand = '2.652'
//	 										significand integer digit = '2'
//												mantissa		= significand factional digits = '.652'
//	 										exponent    = '8'  (10^8)
func (bNum *BigIntNum) GetSciNotationNumber(mantissaLen uint) (SciNotationNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetSciNotationNumber",
		"")

	if err != nil {
		return SciNotationNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return SciNotationNum{}, err
	}

	return new(bigIntNumProton).bigIntNumGetSciNotationNumber(
		bNum, mantissaLen, ePrefix)
}

// GetSciNotationStr
//
// Returns a string expressing the current BigIntNum numerical
// value as scientific notation.
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//
//											scientific notation string. If the value of 'mantissaLen'
//											is less than two ('2'), this method will automatically set
//											the 'mantissaLen' to a default value of two ('2').
//
//											Example Scientific Notation:
//											----------------------------
//
//	 										scientific notation string: '2.652e+8'
//
//	 										significand = '2.652'
//	 										significand integer digit = '2'
//												mantissa		= significand factional digits = '.652'
//	 										exponent    = '8'  (10^8)
func (bNum *BigIntNum) GetSciNotationStr(mantissaLen uint) (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetSciNotationStr",
		"")

	if err != nil {
		return "", err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return "", err
	}

	return new(bigIntNumAtom).getSciNotationStr(
		bNum, mantissaLen, ePrefix)
}

// GetThousandsSeparator - returns a rune which represents
// the character currently used to separate thousands in
// the display of the current BigIntNum instance.
//
// In the USA, the thousands separator is a comma character.
//
// Example: 1,000,000,000
func (bNum *BigIntNum) GetThousandsSeparator() (rune, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetThousandsSeparator",
		"")

	if err != nil {
		return '0', err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return '0', err
	}

	var nSepSymbol NumSepSymbolCode

	nSepSymbol = THOUSANDSYMBOL

	return new(bigIntNumProton).bigIntNumGetNumSepSymbol(
		bNum, nSepSymbol, ePrefix)
}

// GetUInt
//
// Returns a type 'uint' containing the 32-bit unsigned
// integer value of the current BigIntNum instance.
//
// If the current BigIntNum value is greater than the maximum
// 'uint' value, the maximum 32-bit unsigned integer value is returned
// in addition to an 'error'.
//
// If the current BigIntNum value is less than the minimum 'uint'
// value, the minimum 32-bit integer value of zero is returned along
// with an 'error'.
func (bNum *BigIntNum) GetUInt() (uint, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetUInt",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).bigIntNumGetUInt(
		bNum, ePrefix)
}

// GetUInt64
//
// Returns the integer value of BigIntNum.bigInt as a  64-bit
// unsigned integer. If the value of BigIntNum.bigInt exceeds
// that of the maximum unsigned 64-bit integer value, an error
// is returned.
func (bNum *BigIntNum) GetUInt64() (uint64, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetUInt64",
		"")

	if err != nil {
		return 0, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return 0, err
	}

	return new(bigIntNumProton).bigIntNumGetUInt64(bNum, ePrefix)
}

// Inverse
//
// Returns the inverseBigIntNum of the current BigIntNum value.
// The inverseBigIntNum of the value is equal to one ('1') divided by the
// numeric value of the current BigIntNum.
//
// The BigIntNum return value for this operation will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
func (bNum *BigIntNum) Inverse(maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetIntegerValue",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).inverseBigIntNum(
		bNum,
		maxPrecision,
		ePrefix)
}

// IsEvenNumber
//
// Returns true if the current BigIntNum value is evenly
// divisible by 2.
//
// Even Number Definitions:
//
//	https://www.mathsisfun.com/definitions/even-number.html
//
// "In mathematics, parity is the property of an
// integer's inclusion in one of two categories:
// even or odd. An integer is even if it is evenly
// divisible by two and odd if it is not even."
//
// "Examples of even numbers include −4, 0, 82 and 178."
// In particular, zero is an even number."
//
// https://en.wikipedia.org/wiki/Parity_(mathematics)
func (bNum *BigIntNum) IsEvenNumber() (bool, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.IsEvenNumber",
		"")

	if err != nil {
		return false, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return false, err
	}

	return new(bigIntNumNeutron).isEvenBigIntNumber(
		bNum, ePrefix)
}

// Increment
//
// Adds a value of +1 (plus one) to the numeric value of the
// current BigIntNum instance.
//
// The numeric separators (decimal separator, thousands separator
// and currency symbol) from the original BigIntNum will remain
// unchanged.
func (bNum *BigIntNum) Increment() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Increment",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumNeutron).incrementBigIntNum(
		bNum, ePrefix)
}

// IsValid - returns a boolean value signaling whether the
// current BigIntNum object is valid.
func (bNum *BigIntNum) IsValid(callingMethodName string) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	if len(callingMethodName) > 0 {
		callingMethodName = "BigIntNum.IsValid" + "\n" + callingMethodName
	} else {
		callingMethodName = "BigIntNum.IsValid"
	}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		callingMethodName,
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)
}

// IsZero - Returns a boolean signaling whether the current
// BigIntNum value is zero.
func (bNum *BigIntNum) IsZero() (bool, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return false, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return false, err
	}

	return new(bigIntNumMolecule).isBIntNumZero(
		bNum,
		ePrefix)
}

// Mod
//
// Performs a modulo operation where the current BigIntNum numeric value is the
// dividend and the divisor is the input parameter, 'divisor'.  The modulo operation finds
// the remainder after division of one number by another (sometimes called modulus).
// (Wikipedia: https://en.wikipedia.org/wiki/Modulo_operation)
//
//		 									dividend = bNum
//	  									dividend % divisor = modulo
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
func (bNum *BigIntNum) Mod(
	divisor BigIntNum,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return modulo, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return modulo, err
	}

	return new(bigIntNumNeutron).modBigIntNum(
		bNum,
		&divisor,
		maxPrecision,
		ePrefix)
}

// Multiply
//
// Multiplies the numerical value of the current BigIntNum instance
// ('multiplier') times input parameter 'multiplicand'. The 'product' of this
// multiplication operation is returned as a BigIntNum.
//
//	multiplier = bNum
//	multiplier X multiplicand = product
//
// The BigIntNum instance returned by this method, 'product', will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from the current BigIntNum instance.
func (bNum *BigIntNum) Multiply(multiplicand BigIntNum) (product BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return product, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return product, err
	}

	return new(bigIntNumNeutron).multiplyBigIntNum(
		bNum,
		&multiplicand,
		ePrefix)
}

// MultiplyByFive - Multiplies the numerical value of the current BigIntNum
// instance times five (5). The product is returned as a BigIntNum.
//
//	product = bNum X 5
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
func (bNum *BigIntNum) MultiplyByFive() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.MultiplyByFive",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).multiplyByFiveBigIntNum(
		bNum, ePrefix)
}

// MultiplyByTen - Multiplies the numerical value of the current BigIntNum
// instance times ten (10). The product is returned as a BigIntNum.
//
//	product = bNum X 10
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
func (bNum *BigIntNum) MultiplyByTen() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.MultiplyByTen",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).multiplyByTenBigIntNum(bNum, ePrefix)
}

// MultiplyByTenToPower
//
// Multiplies the numerical value of the current BigIntNum/ instance times
// ten to the power of 'exponent' (10^exponent). The product is returned
// as the new value for the current BigIntNum. The original value of the
// BigIntNum instance will be overwritten and destroyed.
//
//	bNum = bNum X 10^exponent
//
// The BigIntNum instance generated by this method will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from the original BigIntNum instance.
func (bNum *BigIntNum) MultiplyByTenToPower(exponent uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.MultiplyByTenToPower",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumNeutron).multiplyByTenToPowerBigIntNum(
		bNum, exponent, ePrefix)
}

// MultiplyByTenToPowerAdd
//
// Performs three operations on the current BigIntNum instance.
//
// (1) 	First, the method multiplies the numerical value of the current BigIntNum
//
//	instance times ten to the power of 'exponent' (10^exponent).
//
//						bNum1 = bNum X 10^exponent
//
// (2)  Second, the method adds input parameter 'addend' to the product generated
//
//	by operation (1), above.
//
//						bNum2 = bNum1 + 'addend'
//
// (3)  Third and finally, the original value of the current BigIntNum instance
//
//	will be overwritten and replaced by the 'bNum2' value generated in
//	operation (2), above.
//
// The BigIntNum instance generated by this method will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from the
// original BigIntNum instance.
func (bNum *BigIntNum) MultiplyByTenToPowerAdd(
	exponent uint, addend BigIntNum) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	return new(bigIntNumNeutron).multiplyByTenToPowerAddBigIntNum(
		bNum, exponent, &addend, ePrefix)
}

// MultiplyByThree
//
// Multiplies the numerical value of the current BigIntNum
// instance times three (3). The product is returned as a BigIntNum.
//
//	product = bNum X 3
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
func (bNum *BigIntNum) MultiplyByThree() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.MultiplyByThree",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).multiplyByThreeBigIntNum(
		bNum, ePrefix)
}

// MultiplyByTwo
//
// Multiplies the numerical value of the current BigIntNum
// instance times two (2). The product is returned as a BigIntNum.
//
//	product = bNum X 2
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
func (bNum *BigIntNum) MultiplyByTwo() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.MultiplyByTwo",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).multiplyByTwoBigIntNum(
		bNum, ePrefix)
}

// New - returns a new BigIntNum instance initialized to zero.
//
// The BigIntNum instance returned by this method will contain USA
// default numeric separators (decimal separator, thousands separator
// and currency symbol).
func (bNum *BigIntNum) New() BigIntNum {

	return new(bigIntNumMechanics).new()
}

// NewWithNumSeps
//
// Returns a new BigIntNum instance initialized to zero.
//
// Input parameter 'numSeps' will be used to seed the new
// BigIntNum instance with numeric separators (decimal
// separator, thousands separator and currency symbol).
//
// If input parameter 'numSeps' is determined to be empty,
// the new returned instance of BigIntNum, will be
// automatically configured with USA default numeric
// separators (decimal separator, thousands separator
// and currency symbol).
func (bNum *BigIntNum) NewWithNumSeps(
	numSeps NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewWithNumSeps",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNanobot).newWithNumSeps(numSeps, ePrefix)
}

// NewBigInt - Creates a new BigIntNum instance using a *big.Int type and its
// associated precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
// # Input Parameters
//
// bigI *big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without decimal
//	digits.
//
// precision int
//
//	This unsigned integer (always a positive value) identifies
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places.
//
//	Example:
//
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bNum *BigIntNum) NewBigInt(
	bigI *big.Int,
	precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigInt",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumMechanics).newBigInt(
		bigI, precision, ePrefix)
}

// NewBigIntPrecision
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated precision (also of type *big.Int).
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
//	Precision Example:
//	==================
//
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
// Input Parameters
// ================
//
// bigI 			*big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without decimal digits.
//
// precision  *big.Int
//
//	This integer value (always a positive value) identifies
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places. If precision is greater than the maximum
//	value of an unsigned integer (+4,294,967,295,	which equals
//	2^32 − 1), an error will be triggered. Also, if the 'precision'
//	value is less than zero, an error will be triggered.
//
// Return Parameters
// =================
//
//	BigIntNum - a type BigIntNum numeric value
//
//	error			- If not 'nil', this prameter will
//							transmit any processing errors
//							encountered.
//
//
//		The new BigIntNum instance returned by this method will contain USA default
//		numeric separators (decimal separator, thousands separator and currency
//		symbol). To reconfigure the numeric separators reference method:
//							BigIntNum.SetNumericSeparators()
func (bNum *BigIntNum) NewBigIntPrecision(
	bigInt *big.Int, precision *big.Int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigIntPrecision",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).newBigIntNumWithPrecision(
		bigInt, precision, ePrefix)
}

// NewBigIntExponent
//
// New bigInt Exponent returns a new BigIntNum instance
// in which the numeric value is set using an integer
// multiplied by 10 raised to the power of the 'exponent'
// parameter.
//
//	numeric value = integer X 10^exponent
//
//						OR
//
//	BigIntNum (return value) = bigI X 10^exponent
//
// If exponent is less than +1, precision is set equal to
// exponent and bigI is unchanged.
//
// If exponent is greater than 0, bigI is multiplied by 10
// raised to the power of 'exponent', and precision is set
// equal to zero.
//
// Examples:
//
//	biNum :=
//			new(BigIntNum).
//				NewBigIntExponent(big.NewInt(int64(123456)), -3) =
//								"123.456"  precision = 3
//
//	biNum :=
//			BigIntNum{}.NewBigIntExponent(big.NewInt(int64(123456)), 3) = "123456.000" precision = 3
func (bNum *BigIntNum) NewBigIntExponent(
	bigI *big.Int, exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigIntPrecision",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).newBigIntExponent(
		bigI, exponent, ePrefix)
}

// NewBigFloat
//
// Returns a new BigIntNum instance using a *big.Float floating
// point input parameter.  The precision of the resulting BigIint
// numeric value is specified by the input parameter,
// 'maxPrecision'.
//
// Input Parameters
// ================
//
// bigFloat *big.Float
//
//	This *big.Float value will be converted into an instance
//	of BigIntNum.
//
// maxPrecision uint
//
//	The maximum precision for the resulting BigIntNum after
//	conversion of input parameter 'bigFloat'. Resulting precision
//	will never be greater than 'maxPrecision'; however, actual
//	precision may be less than 'maxPrecision'.
func (bNum *BigIntNum) NewBigFloat(
	bigFloat *big.Float,
	maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigFloat",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).newBigFloat(
		bNum,
		bigFloat,
		maxPrecision,
		ePrefix)
}

// NewDecimal
//
// Receives an input parameter 'decNum' of type Decimal and returns a BigIntNum
// instance configured with the numeric value passed by parameter 'decNum'.
//
// Input parameter 'decNum' will be subjected to validation testing. If 'decNum'
// fails these validation tests, an error will be returned.
func (bNum *BigIntNum) NewDecimal(decNum Decimal) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewDecimal",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).newDecimal(
		bNum, decNum, ePrefix)
}

// NewBigIntFixedDecimal
//
// Creates and returns a new BigIntNum instance based
// on input parameter 'fd' of type BigIntFixedDecimal.
//
// The 'fd' numeric value will be converted to type
// a BigIntNum which is then returned by this method.
//
// If input parameter 'fd' proves invalid, an error
// will be returned.
func (bNum *BigIntNum) NewBigIntFixedDecimal(
	fd BigIntFixedDecimal) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigIntFixedDecimal",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumNeutron).newBigIntFixedDecimal(
		fd, ePrefix)
}

// NewFromIntFracStrings - Creates a new BigIntNum instance based on a numeric
// value represented by separate integer and fractional components.
//
// Input parameters 'intStr' and 'fracStr' are strings representing the integer and
// fractional elements of the numeric value. These elements are combined by this
// method to create a numeric value which is then assigned to the new BigIntNum
// instance.
//
// Input parameter 'signVal' must be set to one of two values: +1 or -1. This value is
// used to signal the sign of the resulting numeric value. +1 generates a positive number
// and -1 generates a negative number.
func (bNum *BigIntNum) NewFromIntFracStrings(
	intStr string, fracStr string, signVal int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewFromIntFracStrings()"

	b2, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b2, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = b2.SetIntFracStrings(intStr, fracStr, signVal)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" err := b2.SetIntFracStrings(intStr, fracStr, signVal)\n"+
				"Error='%v' \n",
				ePrefix,
				err.Error())
	}

	return b2, err
}

// NewFloat32 - Returns a new BigIntNum instance using a float32 floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
// Input Parameters
// ================
//
// f32 float32				- This float32 value will be converted into an instance of
//
//	BigIntNum.
//
// maxPrecision uint  - The maximum precision for the result BigIntNum after conversion
//
//	of input parameter f64. Precision will never be greater than
//	'maxPrecision'; however, actual precision may be less than
//	'maxPrecision'.
func (bNum *BigIntNum) NewFloat32(f32 float32, maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntNumNewFloat32()"

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = b.SetFloat32(f32, maxPrecision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by b.SetFloat32(f32, maxPrecision).\n"+
				"f32= '%v'\n"+
				"maxPrecision='%v'\n",
				"Error= %v\n",
				ePrefix,
				fmt.Sprintf("%v", f32),
				maxPrecision,
				err.Error())
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
//
//	BigIntNum.
//
// maxPrecision uint  - The maximum precision for the result BigIntNum after conversion
//
//	of input parameter f64. Precision will never be greater than
//	'maxPrecision'; however, actual precision may be less than
//	'maxPrecision'.
func (bNum *BigIntNum) NewFloat64(f64 float64, maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntNumNewFloat64() "

	b := BigIntNum{}

	b.Empty()

	err := b.SetFloat64(f64, maxPrecision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by b.SetFloat64(f64, maxPrecision).\n"+
				"f64= '%v'\n"+
				"maxPrecision='%v'\n"+
				"Error= %v\n",
				ePrefix,
				strconv.FormatFloat(f64, 'f', -1, 64),
				maxPrecision,
				err.Error())
	}

	return b, nil
}

// NewInt - Creates a new BigIntNum instance initialized to the value
// of input parameter 'intNum' which is passed as type 'int'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//					intNum := int(123456)
//					precision := uint(3)
//					bINum := BigIntNum{}.NewInt(intNum, precision)
//	       bINum is now equal to 123.456
//
// Examples:
// ---------
//
//	  intNum				precision			BigIntNum Result
//		 123456		 		   4							12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (bNum *BigIntNum) NewInt(intNum int, precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewInt()"

	b2, err := new(BigIntNum).NewBigInt(big.NewInt(int64(intNum)), precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by BigIntNum.NewBigInt()\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	numDto, err := bNum.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "numDto, err := bNum.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = b2.SetNumericSeparatorsDto(numDto)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "err = b2.SetNumericSeparatorsDto(numDto)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return b2, err
}

// NewIntExponent - This method returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'intNum' is of type int.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := BigIntNum{}.NewIntExponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := BigIntNum{}.NewIntExponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  intNum			 exponent			  BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (bNum *BigIntNum) NewIntExponent(intNum int, exponent int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewIntExponent()"

	bigI := big.NewInt(int64(intNum))

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = b.SetBigIntExponent(bigI, exponent)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" err = b.SetBigIntExponent(bigI, exponent)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return b, nil
}

// NewInt32 - Creates a new BigIntNum instance initialized to the value
// of input parameter 'int32Num' which is passed as type 'int32'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//					num := int32(123456)
//					precision := uint(3)
//					bINum := BigIntNum{}.NewInt32(num, precision)
//	       bINum is now equal to 123.456
//
// Examples:
// ---------
//
//	  int32Num			precision			BigIntNum Result
//		123456					4								12.3456
//	  123456          0								123456
//	  123456          1								12345.6
func (bNum *BigIntNum) NewInt32(int32Num int32, precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewInt32()"

	bIn2, err := new(BigIntNum).NewBigInt(big.NewInt(int64(int32Num)), precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" bIn2, err := new(BigIntNum).NewBigInt(big.NewInt(int64(int32Num)), precision)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return bIn2, nil
}

// NewInt32Exponent -This method returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'int32Num' is of type int32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := BigIntNum{}.NewInt32Exponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := BigIntNum{}.NewInt32Exponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	 int32Num		 exponent			  BigIntNum Result
//		123456				-3							123.456
//		123456				 3							123456.000
//	  123456				 0              123456
func (bNum *BigIntNum) NewInt32Exponent(int32Num int32, exponent int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewInt32Exponent()"

	bigI := big.NewInt(int64(int32Num))

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = b.SetBigIntExponent(bigI, exponent)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" err = b.SetBigIntExponent(bigI, exponent)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return b, nil
}

// NewInt64 - Creates a new BigIntNum instance initialized to the value
// of input parameter 'int64Num' which is passed as type 'int64'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//					int64Num := int64(123456)
//					precision := uint(3)
//					bINum := BigIntNum{}.NewInt64(int64Num, precision)
//	       bINum is now equal to 123.456
//
// Examples:
// ---------
//
//	  int64Num			precision			BigIntNum Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (bNum *BigIntNum) NewInt64(int64Num int64, precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewInt64()"

	bIN2, err := new(BigIntNum).NewBigInt(big.NewInt(int64Num), precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return bIN2, nil
}

// NewInt64Exponent -This method returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := BigIntNum{}.NewInt64Exponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := BigIntNum{}.NewInt64Exponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	 int64Num		 exponent			  BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (bNum *BigIntNum) NewInt64Exponent(int64Num int64, exponent int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewInt64()"

	bigI := big.NewInt(int64Num)

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = b.SetBigIntExponent(bigI, exponent)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" err = b.SetBigIntExponent(bigI, exponent)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return b, nil
}

// NewIntAry - Creates a new BigIntNum instance from an input parameter
// IntAry.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
// The new BigIntNum instance returned by this method will contain default
// numeric separators (decimal separator, thousands separator and currency
// symbol).
func (bNum *BigIntNum) NewIntAry(ia IntAry) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewIntAry()"

	err := ia.IsValid(ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	bInt, err := ia.GetBigInt()

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" bInt, err := ia.GetBigInt()\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	precision, err := ia.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "precision, err := ia.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix,
				ReturnFunc: "precision, err := ia.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var eXPrefix *ePref.ErrPrefixDto

	eXPrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumNanobot).setBigInt(
		&b,
		bInt,
		precision,
		eXPrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix,
				ReturnFunc: "err = new(bigIntNumNanobot).setBigInt(\n" +
					"&b, bInt, precision, eXPrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return b, nil
}

// NewIntFracStr - Creates a new BigIntNum instance based on a numeric value represented
// by separate integer and fractional components.
//
// Input parameters 'intStr' and 'fracStr' are strings representing the integer and
// fractional components. They are combined by this method to create a numeric value
// which is assigned to the current BigIntNum instance.
//
// Input parameter 'signVal' must be set to one of two values: +1 or -1. This value is
// used to signal the sign of the resulting numeric value. +1 generates a positive number
// and -1 generates a negative number. If input parameters 'inStr' or 'fracStr' contain
// a leading minus or plus sign character, it will be ignored. The sign of the resulting
// numeric value is controlled strictly by input parameter, 'signVal'.
func (bNum *BigIntNum) NewIntFracStr(intStr, fracStr string, signVal int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewIntFracStr()"

	bIntNum, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" bIntNum, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = bIntNum.SetIntFracStrings(intStr, fracStr, signVal)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" err = bIntNum.SetIntFracStrings(intStr, fracStr, signVal)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return bIntNum, nil
}

// NewINumMgr - Receives an object which implements the INumMgr interface.
// The method then proceeds to create a new BigIntNum instance equivalent
// in numeric value to the input parameter, 'numMgr'. The BigIntNum instance
// is then returned to the calling function.
//
// Currently, the following 'mathops' Types implement the INumMgr interface:
//
//	Decimal,
//	IntAry,
//	NumStrDto,
//	BigIntNum
//
// Note: 'numMgr' must be a pointer to a type. This method will not accept
// 'numMgr' as a value. The pointer to the type is needed in or order to
// call methods on 'numMgr'.
//
// Example 1:
//
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
func (bNum *BigIntNum) NewINumMgr(numMgr INumMgr) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewINumMgr()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy(" Testing 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	bINum, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINum, err := new(bigIntNumMechanics).newZero(0,ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = bINum.SetINumMgr(numMgr)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" err = bINum.SetINumMgr(numMgr)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return bINum, nil
}

// NewNumStr - Receives a number string as input and returns
// a new BigIntNum instance.
//
// This method assumes that the input parameter 'numStr' is a string
// of numeric digits which may be delimited by default USA numeric
// separators. Default USA numeric separators are defined as:
//
//	 	decimal separator = '.'
//	   thousands separator = ','
//			currency symbol = '$'
//
// If the subject 'numStr' employs other national or cultural numeric
// separators, see method BigIntNum.NewNumStrWithNumSeps(), below.
func (bNum *BigIntNum) NewNumStr(numStr string) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewNumStr()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy(" Testing 'bNum'"))

	if err != nil {
		return BigIntNum{}, err
	}

	b, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	err = b.SetNumStr(numStr)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "precision, err := ia.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return b, nil
}

// NewNumStrWithNumSeps - Receives a number string as input and returns a
// new BigIntNum instance. The input parameter 'numSeps' contains numeric
// separators (decimal separator, thousands separator and currency symbol)
// which will be used to parse the number string.
//
// In addition, the numeric separators contained in input parameter 'numSeps'
// will be copied to the returned BigIntNum instance.
func (bNum *BigIntNum) NewNumStrWithNumSeps(
	numStr string,
	numSeps NumericSeparatorDto) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewNumStrWithNumSeps() "

	numSeps.SetDefaultsIfEmpty()

	b2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" b2, err := new(bigIntNumMechanics).\n"+
				"   newZero( 0, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		&b2,
		numSeps,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" err = new(bigIntNumAtom).setNumericSeparatorsDto(\n"+
				"   &b2, numSeps, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = b2.SetNumStr(numStr)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" err = b2.SetNumStr(numStr)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return b2, nil
}

// NewNumStrMaxPrecision - Receives a number string as input and returns
// a new BigIntNum instance. If the resulting precision exceeds
// input parameter 'maxPrecision', the returned BigIntNum result
// will be rounded to 'maxPrecision' decimal places.
func (bNum *BigIntNum) NewNumStrMaxPrecision(
	numStr string,
	maxPrecision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewNumStr()"

	b, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" b, err := new(bigIntNumMechanics).\n"+
				"  newZero(0, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = b.SetNumStr(numStr)

	err = new(bigIntNumMolecule).setNumStr(
		&b,
		numStr,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" b, err := new(bigIntNumMechanics).\n"+
				"  newZero(0, ePrefix)\n"+
				"numStr='%v'\n"+
				"Error= %v\n",
				ePrefix,
				numStr,
				err.Error())
	}

	if b.precision > maxPrecision {

		err = b.RoundToDecPlace(maxPrecision)

		if err != nil {

			return BigIntNum{},
				fmt.Errorf("%v\n"+
					"Error returned by: \n"+
					" err = b.RoundToDecPlace(maxPrecision)\n"+
					"Error= %v\n",
					ePrefix,
					err.Error())
		}
	}

	return b, nil
}

// NewNumStrDto - Receives a NumStrDto instance as input and returns
// a new BigIntNum instance.
func (bNum *BigIntNum) NewNumStrDto(nDto NumStrDto) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewNumStrDto() "

	err := nDto.IsValid(ePrefix + "'nDto' INVALID! ")

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix+"Error returned from nDto.IsValid(\"\"). "+
				"NumStr='%v' Error='%v'", nDto.GetNumStr(), err.Error())
	}

	bigI, err := nDto.GetBigInt()

	if err != nil {
		return BigIntNum{},
			fmt.Errorf(ePrefix+"Error returned by nDto.GetBigInt(). "+
				"Error='%v'", err.Error())
	}

	b, err := new(BigIntNum).NewZero(0)

	err = new(bigIntNumNanobot).setBigInt(
		&b,
		bigI,
		uint(nDto.GetPrecision()),
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	return b, nil
}

// NewOne - Returns a BigIntNum Type with a value equal to '1' (one).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '1', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								1
//			1								1.0
//			2								1.00
//			3								1.000
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bNum *BigIntNum) NewOne(precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewOne()"

	return new(bigIntNumMolecule).newOne(
		precision,
		ePrefix)
}

// NewTwo - Returns a BigIntNum Type with a value equal to  '2' (two).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '1', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								2
//	   1								2.0
//	   2								2.00
//			3								2.000
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bNum *BigIntNum) NewTwo(precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewTwo()"

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b.SetBigInt(newVal, precision)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	if precision == 0 {

		err = new(bigIntNumNanobot).setBigInt(
			&b,
			big.NewInt(2),
			0,
			ePrefix)

		if err != nil {

			return BigIntNum{}, err
		}

		return b, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	newVal := big.NewInt(0).Mul(big.NewInt(2), scaleVal)

	err = new(bigIntNumNanobot).setBigInt(
		&b,
		newVal,
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	return b, nil
}

// NewThree - Returns a BigIntNum Type with a value equal to  '3' (three).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '1', set 'precision' equal
// to zero (0).
//
// Examples:
// =========
//
// 'precision'
//
//	  value 					Result
//			0								3
//	   2								3.00
//			3								3.000
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bNum *BigIntNum) NewThree(precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewThree()"

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix, err.Error())
	}

	if precision == 0 {

		err = new(bigIntNumNanobot).setBigInt(
			&b,
			big.NewInt(3),
			0,
			ePrefix)

		if err != nil {

			return BigIntNum{}, err
		}

		return b, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	newVal := big.NewInt(0).Mul(big.NewInt(3), scaleVal)

	err = b.SetBigInt(newVal, precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" b.SetBigInt(newVal, precision)\n"+
				"Error= %v\n",
				ePrefix, err.Error())
	}

	return b, nil
}

// NewFive - Returns a BigIntNum with integer value of  '5' (five).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '10', set 'precision' equal
// to zero (0).
//
// 'precision'
//
//	  value 					Result
//			0								 5
//			1								 5.0
//	   2								 5.00
//			3								 5.000
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bNum *BigIntNum) NewFive(precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewFive()"

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix, err.Error())
	}

	if precision == 0 {

		err = b.SetBigInt(big.NewInt(5), 0)

		if err != nil {

			return BigIntNum{},
				fmt.Errorf("%v\n"+
					"Error returned by b.SetBigInt(big.NewInt(5), 0)\n"+
					"Error= %v\n",
					ePrefix, err.Error())
		}

		return b, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	newVal := big.NewInt(0).Mul(big.NewInt(5), scaleVal)

	err = b.SetBigInt(newVal, precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by b.SetBigInt(newVal, precision)\n"+
				"Error= %v\n",
				ePrefix, err.Error())
	}

	return b, nil
}

// NewTen - Returns a BigIntNum with integer value of  '10' (ten).
// The number of zeros created after the decimal place holder
// (fractional digits) is determined by the input parameter 'precision'.
// To create an integer with a value equal to '10', set 'precision' equal
// to zero (0).
//
// 'precision'
//
//	  value 					Result
//			0								10
//			1								10.0
//	   2								10.00
//			3								10.000
//
// The new BigIntNum instance returned by this method will contain USA default numeric
// separators (decimal separator, thousands separator and currency symbol).
func (bNum *BigIntNum) NewTen(precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewTen()"

	var err error

	b, err := new(BigIntNum).NewZero(0)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by  new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix, err.Error())
	}

	if precision == 0 {

		err = b.SetBigInt(big.NewInt(10), 0)

		if err != nil {

			return BigIntNum{},
				fmt.Errorf("%v\n"+
					"Error returned by b.SetBigInt(big.NewInt(10), 0)\n"+
					"Error= %v\n",
					ePrefix, err.Error())
		}

		return b, nil

	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10), big.NewInt(int64(precision)), nil)

	newVal := big.NewInt(0).Mul(big.NewInt(10), scaleVal)

	err = b.SetBigInt(newVal, precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by b.SetBigInt(newVal, precision)\n"+
				"Error= %v\n",
				ePrefix, err.Error())
	}

	return b, nil
}

// NewUint - Creates a new BigIntNum instance initialized to the value
// of input parameter 'uintNum' which is passed as type 'uint'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//					uintNum := uint(123456)
//					precision := uint(3)
//					bINum := BigIntNum{}.NewUint(uintNum, precision)
//	       bINum is now equal to 123.456
//
// Examples:
// ---------
//
//	  uintNum			precision			BigIntNum Result
//		123456					4							12.3456
//	  123456          0             123456
//	  123456          1             12345.6
func (bNum *BigIntNum) NewUint(uintNum uint, precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewUint()"

	bIN2, err := new(BigIntNum).NewBigInt(big.NewInt(0).SetUint64(uint64(uintNum)), precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" new(BigIntNum).NewBigInt(big.NewInt(0).SetUint64(uint64(uintNum)), precision)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

	}

	return bIN2, nil
}

// NewUintExponent - This method returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uintNum' is of type uint.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := BigIntNum{}.NewUintExponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := BigIntNum{}.NewUintExponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  uintNum			exponent			BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (bNum *BigIntNum) NewUintExponent(uintNum uint, exponent int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewUintExponent()"

	baseBInt := big.NewInt(int64(uintNum))

	b2 := new(BigIntNum).New()

	err := b2.SetBigIntExponent(baseBInt, exponent)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" b2.SetBigIntExponent(baseBInt, exponent)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

	}

	return b2, nil
}

// NewUint32 - Creates a new BigIntNum instance initialized to the value
// of input parameter 'uint32Num' which is passed as type 'uint32'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//					uint32Num := uint32(123456)
//					precision := uint(3)
//					bINum := BigIntNum{}.NewUint32(uint32Num, precision)
//	       	bINum is now equal to 123.456
//
// Examples:
// ---------
//
//	  uint32Num		precision			BigIntNum Result
//		123456					4							12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (bNum *BigIntNum) NewUint32(uint32Num uint32, precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewUint32()"

	bIN2, err := new(BigIntNum).NewBigInt(big.NewInt(0).SetUint64(uint64(uint32Num)), precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" new(BigIntNum).NewBigInt(big.NewInt(0).SetUint64(uint64(uint32Num)), precision)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return bIN2, nil
}

// NewUint32Exponent - This method returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uint32Num' is of type uint32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := BigIntNum{}.NewUint32Exponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := BigIntNum{}.NewUint32Exponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  uint32Num		exponent			BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
func (bNum *BigIntNum) NewUint32Exponent(uint32Num uint32, exponent int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewUint32Exponent()"

	baseBInt := big.NewInt(int64(uint32Num))

	b2 := new(BigIntNum).New()

	err := b2.SetBigIntExponent(baseBInt, exponent)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" b2.SetBigIntExponent(baseBInt, exponent)\n"+
				"Error= %v",
				ePrefix,
				err.Error())
	}

	return b2
}

// NewUint64 - Creates a new BigIntNum instance initialized to the value
// of input parameter 'uint64Num' which is passed as type 'uint64'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//					uint64Num := uint64(123456)
//					precision := uint(3)
//					bINum := BigIntNum{}.NewUint64(uint64Num, precision)
//	       bINum is now equal to 123.456
//
// Examples:
// ---------
//
//	  uint64Num		precision			BigIntNum Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (bNum *BigIntNum) NewUint64(
	uint64Num uint64, precision uint) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewUint64()"

	bInt2, err := new(BigIntNum).NewBigInt(big.NewInt(0).SetUint64(uint64Num), precision)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" bInt2, err := new(BigIntNum).NewBigInt(big.NewInt(0).SetUint64(uint64Num), precision)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

	}

	return bInt2, nil
}

// NewZero
//
// Returns a BigIntNum instance with a value equal to zero.
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter
// 'precision'.
//
// To create an integer with a value equal to '0', set
// 'precision' equal to zero (0).
//
//	precision
//	  value 					Result
//			0								0
//			2								0.00
//			3								0.000
func (bNum *BigIntNum) NewZero(precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewZero",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntNumMechanics).newZero(
		precision,
		ePrefix)
}

// NewUint64Exponent - This method returns a new BigIntNum instance in which
// the numeric value is set using an integer multiplied by 10 raised to
// the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uint64Num' is of type uint64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := BigIntNum{}.NewUint64Exponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := BigIntNum{}.NewUint64Exponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  uint64Num		exponent			BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (bNum *BigIntNum) NewUint64Exponent(uint64Num uint64, exponent int) (BigIntNum, error) {

	ePrefix := "BigIntNum.NewUint64Exponent()"

	baseBInt := big.NewInt(0).SetUint64(uint64Num)

	b2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" b2, err := new(bigIntNumMechanics).\n"+
				"   newZero(0, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = new(bigIntNumMolecule).
		setBigIntExponent(&b2, baseBInt, exponent, ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" err = new(bigIntNumMolecule).\n"+
				"  setBigIntExponent(&b2,baseBInt, exponent, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return b2, nil
}

// Reset - Resets the current BigIntNum to a new
// valid BigIntNum using the BigIntNum components
// BigIntNum.bigInt and BigIntNum.precision. This
// method is usually called after method bNum.IsValid()
// returns false.
func (bNum *BigIntNum) Reset() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Reset",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumElectron).resetBigIntNum(
		bNum, ePrefix)
}

// RoundToDecPlace - Rounds the current BigIntNum instance to a specified
// number of decimal places.
//
// 'precision' equals the number of digits to the right of the decimal
// place.
//
// Example:
//
//	integer= 123456; precision = 3; Numeric Value= 123.456
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
//		 Original       				'precision'				Resulting
//	   Value								input parameter			  Value
//	 --------------				---------------     -------------
//		654.123456									9							 654.123456000
//		654.123456									4							 654.1235
//
// -654.123456									9							-654.123456000
// -654.123456									4							-654.1235
//
//		 0												3								 0.000
//	   0.000000									0								 0
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) RoundToDecPlace(precision uint) error {

	ePrefix := "BigIntNum.RoundToDecPlace()"

	err := new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {
		return err
	}

	if bNum.precision == precision {
		// Nothing to do. Specified 'precision' is already implemented.
		return nil
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		bNum,
		ePrefix)

	if err != nil {
		return err
	}

	numSeps := bNum.GetNumericSeparatorsDto()

	// bigInt == zero, set precision an return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		bNum2, err := new(bigIntNumMechanics).newBigInt(
			big.NewInt(0),
			precision,
			ePrefix)

		if err != nil {
			return err
		}

		err = new(bigIntNumUtility).bigIntNumCopyIn(
			bNum,
			&bNum2,
			ePrefix)

		if err != nil {
			return err
		}

		err = new(bigIntNumAtom).setNumericSeparatorsDto(
			bNum,
			numSeps,
			ePrefix)

		return err
	}

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bNum.precision < precision {

		deltaPrecision := precision - bNum.precision

		bNum.ExtendPrecision(deltaPrecision)

		err = new(bigIntNumAtom).setNumericSeparatorsDto(
			bNum,
			numSeps,
			ePrefix)

		if err != nil {
			return err
		}

		return nil
	}

	// Must be: bNum.precision >  precision

	//bigNumRound5 :=
	//	BigIntNum{}.NewBigInt(big.NewInt(5), uint(precision+1))

	bigNumRound5, err := new(bigIntNumMechanics).newBigInt(
		big.NewInt(5),
		precision+1,
		ePrefix)

	if err != nil {
		return err
	}

	bigNumBase, err := new(bigIntNumMechanics).newBigInt(
		bNum.absBigInt,
		bNum.precision,
		ePrefix)

	if err != nil {
		return err
	}

	result, err := BigIntMathAdd{}.AddBigIntNums(bigNumBase, bigNumRound5)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned by: \n"+
			" result, err := BigIntMathAdd{}.AddBigIntNums(bigNumBase, bigNumRound5)\n"+
			"Error= %v\n",
			ePrefix,
			err.Error())

	}

	// 10^deltaPrecision
	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision-precision)), nil)

	result.bigInt = big.NewInt(0).Quo(result.bigInt, scaleVal)

	if bNum.sign < 0 {
		result.bigInt = big.NewInt(0).Neg(result.bigInt)
	}

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		result.bigInt,
		precision,
		ePrefix)

	return err
}

// SetBigInt - Sets the value of the current BigIntNum instance using
// the input parameters *big.Int integer and precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
// Input Parameters
// bigI *big.Int	- 'bigI' is a type *big.Int and represents the integer
//
//	value of the number; that is, the numeric value with
//	out decimal digits.
//
// precision uint	- This unsigned integer (always a positive value) identifies
//
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places. Example:
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) SetBigInt(bigI *big.Int, precision uint) error {

	ePrefix := "BigIntNum.SetBigInt()"

	return new(bigIntNumNanobot).setBigInt(
		bNum,
		bigI,
		precision,
		ePrefix)
}

// SetBigIntExponent - Sets the numeric value using an integer
// multiplied by 10 raised to the power of the 'exponent'
// parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'bigI' is of type *big.Int.
//
// Input parameter 'exponent' is of type int.
//
// If exponent is less than +1, precision is set equal to exponent and
// bigI is unchanged. Example:
//
//	   bigI				exponent			BigIntNum Result
//		 123456		 		  -3							123.456
//
// If exponent is greater than 0, bigI is multiplied by 10 raised to the
// power of exponent and precision is set equal to exponent.
//
//	   bigI				exponent			BigIntNum Result
//		 123456		 		   3							123456.000
func (bNum *BigIntNum) SetBigIntExponent(
	bigI *big.Int, exponent int) error {

	ePrefix := "BigIntNum.NewBigIntExponent()"

	if bigI == nil {

		return fmt.Errorf("%v\n"+
			"Error: Input parameter 'bigI' is a nil pointer!\n",
			ePrefix)
	}

	return new(bigIntNumMolecule).
		setBigIntExponent(bNum, bigI, exponent, ePrefix)
}

// SetBigFloat
//
// Sets the value of the current BigIntNum using a *big.Float
// floating point input parameter.  The precision of the number
// is specified by the input parameter, 'maxPrecision'.
//
// Input Parameters
// ================
//
// bigFloat *big.Float
//
//	This *big.Float value will be converted saved as the
//	current instance of BigIntNum
//
// maxPrecision uint
//
//		The maximum precision for the resulting BigIntNum value
//		after conversion of input parameter 'bigFloat'. Final
//		precision will never be greater than 'maxPrecision';
//	 	however, actual precision may be less than
//	 	'maxPrecision'.
//
// Background
// ==========
//
// As part of converting a BigFloat to a BigIntNum number,
// the Accuracy flag is analyzed to determine if rounding errors
// occurred in connection with this conversion. The internal
// Accuracy Flag is set as:
//
//			Below Accuracy == -1	(Returns an error)
//	   Exact Accuracy == 0		(No Error Returned)
//	   Above Accuracy == +1		(Returns an error)
//
// If Accuracy == 0, no error is issued by this method. However, if
// Accuracy == -1 or Accuracy == +1, an error will be returned. All
// conversions must be exact.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) contained in the original BigIntNum ('bNum')
// will remain unchanged and will not be altered by this method.
func (bNum *BigIntNum) SetBigFloat(
	bigFloat *big.Float, maxPrecision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumNanobot).setBigFloat(
		bNum,
		bigFloat,
		maxPrecision,
		ePrefix)
}

// SetBigRat
//
// Sets the value of the current BigIntNum instance to that of
//
//	input parameter 'ratNum', a rational number of type *big.Rat.
//
//	Input Parmeters
//	===============
//
// ratNum 			*big.Rat
//
//	The value of ratNum will be used to configure the current
//	BigIntNum instance and reset its value.
//
// maxPrecision uint
//
//	The maximum precision for the resulting BigIntNum value
//	after it is reset to the value of input parameter 'ratNum'.
//	Precision will never be greater than 'maxPrecision'; however,
//	actual precision may be less than 'maxPrecision'.
//
//	Existing numeric separators (decimal separator, thousands separator
//	and currency symbol) in the current BigIntNum instance will remain
//	unchanged and will not be altered by this method.
func (bNum *BigIntNum) SetBigRat(
	ratNum *big.Rat, maxPrecision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Ceiling",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumMolecule).setBigRat(
		bNum,
		ratNum,
		maxPrecision,
		ePrefix)
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
func (bNum *BigIntNum) SetDecimalSeparator(decimalSeparator rune) {

	if decimalSeparator == 0 {
		decimalSeparator = '.'
	}

	bNum.decimalSeparator = decimalSeparator
}

// SetIntFracStrings - Sets the value of the current BigIntNum instance based on
// a numeric value represented by separate integer and fractional components.
//
// Input parameters 'intStr' and 'fracStr' are strings representing the integer and
// fractional components. They are combined by this method to create a numeric value
// which is assigned to the current BigIntNum instance.
//
// Input parameter 'signVal' must be set to one of two values: +1 or -1. This value is
// used to signal the sign of the resulting numeric value. +1 generates a positive number
// and -1 generates a negative number. If input parameters 'inStr' or 'fracStr' contain
// a leading minus or plus sign character, it will be ignored. The sign of the resulting
// numeric value is controlled strictly by input parameter, 'signVal'.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) SetIntFracStrings(intStr, fracStr string, signVal int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetIntFracStrings",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix.XCpy("Validating 'bNum'"))

	if err != nil {
		return err
	}

	ePrefix := "BigIntNum.SetIntFracStrings() "

	cleanIntRuneAry := make([]rune, 0, 100)

	zeroChar := uint8('0')
	nineChar := uint8('9')

	lStr := len(intStr)

	if lStr == 0 {

		return fmt.Errorf("%v\n"+
			"Error: Input Parameter 'intStr' is zero Length!",
			ePrefix)

	}

	isFirstRune := true

	// Create pure number string from 'intStr'
	for i := 0; i < lStr; i++ {

		if intStr[i] >= zeroChar &&
			intStr[i] <= nineChar {

			if isFirstRune && signVal == -1 {
				cleanIntRuneAry = append(cleanIntRuneAry, '-')
			}

			isFirstRune = false

			cleanIntRuneAry = append(cleanIntRuneAry, rune(intStr[i]))
		}
	}

	if len(cleanIntRuneAry) == 0 {
		cleanIntRuneAry = append(cleanIntRuneAry, '0')
	}

	lStr = len(fracStr)

	if lStr > 0 {

		isFirstRune = true

		for j := 0; j < lStr; j++ {

			if fracStr[j] >= zeroChar &&
				fracStr[j] <= nineChar {

				if isFirstRune {
					cleanIntRuneAry = append(cleanIntRuneAry, bNum.GetDecimalSeparator())
					isFirstRune = false
				}

				cleanIntRuneAry = append(cleanIntRuneAry, rune(fracStr[j]))
			}

		}
	}

	err := bNum.SetNumStr(string(cleanIntRuneAry))

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by bNum.SetNumStr(string(cleanIntRuneAry)). "+
			"cleanIntRuneAry='%v' Error='%v' ", string(cleanIntRuneAry), err.Error())
	}

	return new(bigIntNumNanobot).setIntFracStrings(
		bNum, intStr, fracStr, signVal, ePrefix)
}

// SetFloat32 - Sets the value of a BigIntNum using a float32 floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
// Input Parameters
// ================
//
// f32 float32				- This float32 value will be converted into an instance of
//
//	BigIntNum.
//
// maxPrecision uint  - The maximum precision for the resulting BigIntNum after
//
//	conversion of input parameter 'ratNum'. Precision will
//	never be greater than 'maxPrecision'; however, actual
//	precision may be less than 'maxPrecision'.
func (bNum *BigIntNum) SetFloat32(f32 float32, maxPrecision uint) error {

	ePrefix := "BigIntNum.SetFloat32() "

	rat := big.NewRat(1, 1).SetFloat64(float64(f32))

	err := bNum.SetBigRat(rat, maxPrecision)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by bNum.SetBigRat(rat, maxPrecision). "+
			"Error='%v' \n", err.Error())
	}

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
//
//	BigIntNum.
//
// maxPrecision uint  - The maximum precision for the resulting BigIntNum after
//
//	conversion of input parameter 'f64'. Resulting precision
//	will never be greater than 'maxPrecision'; however, actual
//	precision may be less than 'maxPrecision'.
func (bNum *BigIntNum) SetFloat64(f64 float64, maxPrecision uint) error {

	ePrefix := "BigIntNum.SetFloat64() "

	rat := big.NewRat(1, 1).SetFloat64(f64)

	err := bNum.SetBigRat(rat, maxPrecision)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned by bNum.SetBigRat(rat, maxPrecision). "+
			"Error='%v' \n", err.Error())
	}

	return nil
}

// SetExpectedNumberOfDigits - Sets the number of expected digits associated with the
// Absolute Value of this 'BigIntNum.absBigInt'. The value is stored in the data
// field, 'BigIntNum.numberOfExpectedDigits'.
//
// Useful in tracking leading zeros.
func (bNum *BigIntNum) SetExpectedNumberOfDigits(numOfDigits *big.Int) error {

	ePrefix := "BigIntNum.SetExpectedNumberOfDigits()"

	return new(bigIntNumMolecule).setExpectedNumberOfDigits(
		bNum,
		numOfDigits,
		ePrefix)
}

// SetExpectedToActualNumberOfDigits - Sets the 'Expected' number of numeric
// digits associated with this BigIntNum, to the actual number of numeric digits
// in the BigIntNum value at the time when this method is called.
func (bNum *BigIntNum) SetExpectedToActualNumberOfDigits() error {

	ePrefix := "BigIntNum.SetExpectedToActualNumberOfDigits()"

	var err error

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			bNum.precision,
			ePrefix)

		if err != nil {
			return err
		}

	}

	actNumOfDigits, _, err := bNum.GetActualNumberOfDigits()

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned by: \n"+
			" actNumOfDigits, _, err := bNum.GetActualNumberOfDigits()\n"+
			"Error= %v\n",
			ePrefix,
			err.Error())

	}

	bNum.numberOfExpectedDigits = big.NewInt(0).Set(actNumOfDigits)

	return nil
}

// SetINumMgr - Receives an input parameter implementing
// the INumMgr interface and proceeds to set the current
// BigIntNum instance to its equivalent numeric value.
//
// Currently, the following 'mathops' Types implement the INumMgr interface:
//
//	Decimal,
//	IntAry,
//	NumStrDto,
//	BigIntNum
//
// 'numMgr' must be a pointer to a type. This method will not accept
// 'numMgr' as a value. The pointer to the type is needed in or order to
// call methods on 'numMgr'.
//
// This method will test the validity of input parameter, 'numMgr'.
//
// Example 1:
//
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
func (bNum *BigIntNum) SetINumMgr(numMgr INumMgr) error {

	ePrefix := "BigIntNum.SetINumMgr() "

	err := numMgr.IsValid(ePrefix + "numMgr INVALID! ")

	if err != nil {
		return err
	}

	bigInt, err := numMgr.GetBigInt()

	if err != nil {
		return fmt.Errorf(ePrefix+"Error returned by numMgr.GetBigInt(). "+
			"Error='%v'", err.Error())
	}

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		bigInt,
		numMgr.GetPrecisionUint(),
		ePrefix)

	return err
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
//		 Original       			'newPrecision'				Resulting
//	   Value								input parameter			  Value
//	 --------------				---------------     -------------
//		654.123456									9							 654.123456000
//		654.123456									4							 654.1235
//
// -654.123456									9							-654.123456000
// -654.123456									4							-654.1235
//
//			0													3								 0.000
//	   0.000000									0								 0
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) SetPrecision(newPrecision uint) error {

	ePrefix := "BigIntNum.SetPrecision()"

	var err error

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			bNum.precision,
			ePrefix)

		if err != nil {
			return err
		}

	}

	if newPrecision == bNum.precision {
		return nil
	}

	if bNum.precision > newPrecision {

		err = bNum.RoundToDecPlace(newPrecision)

		if err != nil {

			return fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" err = bNum.RoundToDecPlace(newPrecision)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

		}

		return nil
	}

	deltaPrecision := newPrecision - bNum.precision

	// bNum.precision must be less than newPrecision
	bNum.ExtendPrecision(deltaPrecision)

	return nil
}

// SetNumericSeparators - Used to assign values for the Decimal and Thousands separators as well
// as the Currency Symbol to be used in displaying the current number string.
//
// Different nations and cultures use different symbols to delimit numerical values. In the
// USA and many other countries, a period character ('.') is used to delimit integer and
// fractional digits within a numeric value (123.45). Likewise, thousands may be delimited
// by a comma (','). Currency signs very by nationality. For instance, the USA, Canada and
// several other countries use the dollar sign ($) as a currency symbol.
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
func (bNum *BigIntNum) SetNumericSeparators(
	decimalSeparator rune,
	thousandsSeparator rune,
	currencySymbol rune) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetNumericSeparators",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumAtom).setNumericSeparators(
		bNum,
		decimalSeparator,
		thousandsSeparator,
		currencySymbol,
		ePrefix)

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
func (bNum *BigIntNum) SetNumericSeparatorsDto(customSeparators NumericSeparatorDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetNumericSeparatorsDto",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		customSeparators,
		ePrefix)
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
func (bNum *BigIntNum) SetNumericSeparatorsToDefaultIfEmpty() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetNumeSetNumericSeparatorsToDefaultIfEmptyricSeparatorsDto",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		bNum,
		ePrefix)
}

// SetNumericSeparatorsToUSADefault - Sets Numeric separators:
//
//	Decimal Point Separator
//	Thousands Separator
//	Currency Symbol
//
// to the United States of America (USA) defaults.
//
// Call specific methods to set numeric separators for other countries or
// cultures:
//
//	bNum.SetDecimalSeparator()
//	bNum.SetThousandsSeparator()
//	bNum.SetCurrencySymbol()
func (bNum *BigIntNum) SetNumericSeparatorsToUSADefault() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetNumericSeparatorsToUSADefault",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumAtom).setNumericSeparatorsToUSADefault(
		bNum,
		ePrefix)
}

// SetNumStr - Initializes the current BigIntNum instance
// for the numeric value of the number string input parameter.
// A number string is a string of numeric digits which may
// or may not be prefixed with a minus sign ('-'). The numeric
// string of digits may also contain a decimal separator such
// as a period ('.'). The decimal separator may be set by the
// user. See Method BigIntNum.SetDecimalSeparator(). The decimal
// separator is used to separate integer and fractional numeric
// digits within the number string.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) SetNumStr(numStr string) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetNumStr",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumMolecule).setNumStr(
		bNum,
		numStr,
		ePrefix)

}

// SetSignValue - Sets the sign value of the current BigIntNum
// to positive (+1) or negative (-1).
//
// If a value other than +1 or -1 is transmitted as an input
// parameter, an error will be returned.
func (bNum *BigIntNum) SetSignValue(signVal int) error {

	ePrefix := "BigIntNum.SetSignValue()"

	var err error

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			bNum.precision,
			ePrefix)

		if err != nil {
			return err
		}

	}

	if signVal == 1 || signVal == -1 {

		if bNum.GetSign() == signVal {

			return nil
		}

		bNum.ChangeSign()

		return nil
	}

	return fmt.Errorf("%v\n"+
		"Error: Input parameter 'signVal' "+
		"must be +1 or -1.\n"+
		"signVal='%v'\n",
		ePrefix,
		signVal)

}

// ShiftPrecisionLeft - Shifts precision of the current BigIntNum
// numeric value to the left by 'shiftLeftPlaces' decimal places. This
// is a 'relative' shift-left operation. The shift left operation is
// therefore performed with the current decimal place position as the
// starting point.
//
// This operation is equivalent to:	result = Decimal value / 10^shiftLeftPlaces
// or signed number divided by 10 raised to the power of shiftLeftPlaces.
//
// This method performs a relative shift left of the decimal place position.
// Be careful, this is NOT Shift Number Left operation. This is Shift Precision
// Left which means that the decimal place will be shifted left.
//
// See Examples below.
//
// Input Parameters
// ================
//
//	shiftLeftPlaces int	- The number of positions the decimal place will be
//												shifted left from its current position.
//
// Examples:
// =========
//
//	shift-left
//
// signed Number		  places				Result
//
//	"123456.789"				3						"123.456789"
//	"123456.789"				2						"1234.56789"
//	"123456.789"        6					  "0.123456789"
//	"123456789"	 			  6						"123.456789"
//	"123"               5	          "0.00123"
//	"0"								  3						"0"
//	"123456.789"				0						"123456.789"		- zero has no effect on original number string
//
// "-123456.789"        0          "-123456.789"
// "-123456.789"        3          "-123.456789"
// "-123456789"			    6					 "-123.456789"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) ShiftPrecisionLeft(shiftLeftPlaces uint) error {

	ePrefix := "BigIntNum.ShiftPrecisionLeft()"

	var err error

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix+" Testing 'bNum'")

	if err != nil {
		return err
	}

	isbNumZero := false

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		isbNumZero = true
	}

	if shiftLeftPlaces == 0 || isbNumZero {

		return nil
	}

	newPrecision := bNum.precision + shiftLeftPlaces

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		bNum.bigInt,
		newPrecision,
		ePrefix)

	return err
}

// ShiftPrecisionRight - Shifts precision of the current BigIntNum
// numeric value to the right by 'shiftRightPlaces' decimal places. This
// is a 'relative' shift-right operation. The shift right operation is
// therefore performed with the current decimal place position as the
// starting point.
//
// This is equivalent to: result = Decimal value X 10^shiftRightPrecision or
// Decimal numeric value multiplied by 10 raised to the power of
// shiftRightPrecision.
//
// This method performs a relative shift right of the decimal place position.
// Be careful, this is NOT a Shift Number Right operation. This is Shift Precision
// Right which means that the decimal place will be shifted right.
//
// See Examples below.
//
// Input Parameters
// ================
//
//	shiftRightPlaces int	- The number of positions the decimal place will be
//													shifted right from its current position.
//
// Examples:
// =========
//
//	shift-right
//
// signed Number		  places				Result
//
//	"123456.789"				3						"123456789"
//	"123456.789"				2						"12345678.9"
//	"123456.789"        6					  "123456789000"
//	"123456789"	 			  6						"123456789000000"
//	"123"               5	          "12300000"
//	"0"								  3						"0"
//	"123456.789"				0						"123456.789"		- zero has no effect on original number string
//
// "-123456.789"        0          "-123456.789"
// "-123456.789"        3          "-123456789"
// "-123456789"			    6					 "-123456789000000"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) ShiftPrecisionRight(shiftRightPlaces uint) error {

	ePrefix := "BigIntNum.ShiftPrecisionRight()"

	var err error

	err = new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix+" Testing 'bNum'")

	if err != nil {
		return err
	}

	isbNumZero := false

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		isbNumZero = true
	}

	if shiftRightPlaces == 0 || isbNumZero {

		return nil

	}

	bigINanobot := new(bigIntNumNanobot)

	if shiftRightPlaces <= bNum.precision {

		newPrecision := bNum.precision - shiftRightPlaces

		err = bigINanobot.setBigInt(
			bNum,
			bNum.bigInt,
			newPrecision,
			ePrefix)

		return err
	}

	// shiftRightPlaces > bNum.precision

	newPrecision := shiftRightPlaces - bNum.precision

	bigITen := big.NewInt(10)

	exponent := big.NewInt(int64(newPrecision))

	scaleFactor := big.NewInt(0).Exp(bigITen, exponent, nil)

	newValue := big.NewInt(0).Mul(bNum.bigInt, scaleFactor)

	err = bigINanobot.setBigInt(
		bNum,
		newValue,
		0,
		ePrefix)

	return err
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
//
//	Initial Value			Trimmed Value
//		456.123000 			 456.123
//			0.000					 0
//			7.0						 7
//	 -456.123000			-456.123
func (bNum *BigIntNum) TrimTrailingFracZeros() error {

	ePrefix := "BigIntNum.TrimTrailingFracZeros()"

	err := new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {
		return err
	}

	if bNum.precision == 0 {
		return nil
	}

	biBaseZero := big.NewInt(0)

	if bNum.bigInt.Cmp(biBaseZero) == 0 {

		bNum.precision = 0

		bNum.scaleFactor = big.NewInt(1)

		return nil
	}

	// bNum.precision must be greater than zero
	biBase10 := big.NewInt(10)
	scrap := big.NewInt(0)
	newBigIntNum, mod10 := big.NewInt(0).QuoRem(bNum.bigInt, biBase10, scrap)
	doReset := false

	for mod10.Cmp(biBaseZero) == 0 && bNum.precision > 0 {

		bNum.bigInt.Set(newBigIntNum)

		bNum.precision--

		newBigIntNum, mod10 = big.NewInt(0).QuoRem(bNum.bigInt, biBase10, scrap)

		doReset = true
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
//		 Original       			'newPrecision'				Resulting
//	   Value								input parameter			  Value
//	 --------------				---------------     -------------
//		654.123456									9							 654.123456000
//		654.123456									4							 654.1234 (no rounding)
//
// -654.123456									9							-654.123456000
// -654.123456									4							-654.1234 (no rounding)
//
//			0													3								 0.000
//	   0.000000									0								 0
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) TruncToDecPlace(precision uint) error {

	ePrefix := "BigIntNum.TruncToDecPlace()"

	err := new(bigIntNumAtom).isBigIntNumValid(
		bNum,
		ePrefix)

	if err != nil {
		return err
	}

	if bNum.bigInt == nil {

		err = new(bigIntNumNanobot).setBigInt(
			bNum,
			big.NewInt(0),
			precision,
			ePrefix)

		if err != nil {
			return err
		}

	}

	if bNum.precision == precision {

		// Nothing to do. Specified 'precision' is already implemented.

		return nil
	}

	// bigInt == zero, set precision an return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		bNum.precision = precision

		return nil
	}

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bNum.precision < precision {

		bNum.ExtendPrecision(precision - bNum.precision)

		return nil
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

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		newBigInt,
		precision,
		ePrefix)

	return err
}
