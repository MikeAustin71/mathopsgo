package mathops

import (
	"fmt"
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

// BigIntNum
//
// Wraps a *big.Int integer with its associated precision and sign
// Values. While the numeric value is stored as an integer of type
// *big.Int, the BigIntNum type is capable of storing decimal
// fractions.
//
// All methods associated with this type all assume that
// the *big.Int value stored by the BigIntNum Type is configured
// in base 10.
//
//	INumMgr
//	========
//
//	The BigIntNum Type implements the INumMgr interface.
//
//	Source Code Repository
//	=======================
//
// https://github.com/MikeAustin71/mathopsgo.git
//
//	Local File: bigintnum.go
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

var _ INumMgr = (*BigIntFixedDecimal)(nil)

// Ceiling
//
// Returns the ceiling integer value of the current BigIntNum
// instance.
//
// Ceiling is defined as: The least, or lowest value integer,
// which is greater than or equal to the numeric value of the
// current BigIntNum.
//
//	 Reference Wikipedia
//	 ===================
//
//	 https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//		 Examples
//		 ========
//
//		     Initial          Ceiling
//		      Value            Value
//		     -------          -------
//		        5.95             6
//		        5.05             6
//		        5                5
//		       -5.05            -5
//		        2.4              3
//		        2.9              3
//		       -2.7             -2
//		       -2               -2
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

// ChangeSign
//
// Changes the sign of the current BigIntNum value.
//
// If the value of BigIntNum is zero, the sign will remain unchanged
// and this method will return with no action taken.
//
// If the sign of the current BigIntNum value is positive (+), the
// sign will be changed to negative (-). Likewise, if the current
// sign is negative (-), the sign will be changed to positive (+).
//
//	NOTE
//	====
//
//	This method will first test the current instance of BigIntNum
//	to determine if that instance is valid, or not. If the current
//	BigIntNum fails the validity test, an error will be returned.
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
// fractional quotient. Precision is defined as the number of fractional digits to the right
// of the decimal place. Be advised that these calculations can support very large precision
// values.
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
// If 'modulo' equals zero ('0'), it signals the current BigIntNum numerical value is 'even';
// that is, it is evenly divisible by two.
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// floating point 'modulo'. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
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

// FormatCurrencyStr
//
//	Formats the current BigIntNum numeric value as a currency string.
//
//	If the Currency Symbol was not previously set for this BigIntNum, the currency symbol
//	is defaulted to the USA standard dollar sign, ('$'). To use other currency symbols, see
//	method BigIntNum.SetCurrencySymbol(). For a list of Major Currency Unicode Symbols, see
//	constants located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
//	If the Decimal Separator was not previously set for this BigIntNum, the Decimal Separator
//	is defaulted to the USA standard period ('.'). To use another character for Decimal
//	Separator, see method BigIntNum.SetDecimalSeparator().
//
//	If the Thousands Separator was not previously set for this BigIntNum, the Thousands
//	Separator is defaulted to the USA standard comma (','). To use another character for
//	Thousands Separator, see method BigIntNum.SetThousandsSeparator().
//
//	Input Parameters
//	================
//
//	negValMode NegativeValueFmtMode -	Specifies the display mode for negative values:
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

// FormatNumStr
//
//	Formats the numeric value of the current BigIntNum
//	instance as number string consisting of integer digits to the left
//	of the decimal place and fractional digits to the right of the decimal
//	point, if such fractional digits exist. The resulting number string
//	will NOT contain a currency symbol or thousands separators.
//
//	If the Decimal Separator was not previously set for this BigIntNum,
//	the Decimal Separator is defaulted to the USA standard period ('.').
//	To use another character for Decimal Separator, see method
//	BigIntNum.SetDecimalSeparator().
//
//	Output Examples: 123456.789 or -123456.789
//
//	Input Parameters
//	================
//
//	negValMode        NegativeValueFmtMode
//	  Specifies the display mode for negative values:
//
//	   LEADMINUSNEGVALFMTMODE    - Negative values formatted with
//	                               a leading minus sign.
//	                               Example: -123456.78
//
//	   PARENTHESESNEGVALFMTMODE  - Negative values formatted with
//	                               surrounding parentheses.
//	                               Example: (123456.78)
//
//	   ABSOLUTEPURENUMSTRFMTMODE - Formats a pure number string with
//	                               absolute (positive) integer value
//	                               and no decimal place separator.
//	                               Example: (12345678)
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

// GetActualNumberOfDigits
//
//	Returns the number of numeric digits in the absolute value of this
//	BigIntNum instance. In addition, a boolean value is returned
//	indicating whether the absolute value is zero.
//
//	Examples
//	========
//
//	Numeric Value      Actual Number of Digits
//
//	        123.45                 5
//	  1,234,567                    7
//	 -1,234,567                    7
//	          0                    1
//	          0.00                 1
//	        012.34                 4
//	          0.1234               4
//	         -0.1234               4
//	          0.123400             4
//	          0.0123400            4
//	  1,234,567.800                8
//	          5                    1
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

// GetAbsoluteNumStr
//
//	Returns the absolute integer value (positive value) of the *big.Int
//	value encapsulated by this BigIntNum. No decimal place is included.
//
//	If fractional digits exist in the BigIntNum, they are included in
//	the number string, but no decimal place is inserted in the returned
//	number string.
//
//	Examples
//	========
//
//	   Value       Returned Number String
//	   123.456           "123456"
//	  -123.456           "123456"
//	   123456            "123456"
//	  -123456            "123456"
//
//	If an error is encountered, an empty string is returned.
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

// GetPrecisionInt
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
//					1.234    	GetPrecisionInt() = 3
//							5			GetPrecisionInt() = 0
//				0.12345  		GetPrecisionInt() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
func (bNum *BigIntNum) GetPrecisionInt() (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.GetPrecisionInt",
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
// Returns true if the current BigIntNum value is an integer which
// is evenly divisible by 2.
//
// Even Number Definitions:
//
//			https://www.mathsisfun.com/definitions/even-number.html
//
//		 https://en.wikipedia.org/wiki/Parity_(mathematics)
//
//			"In mathematics, parity is the property of an integer of whether
//			it is even or odd. An integer is even if it is divisible by 2, and
//			odd if it is not.[1] For example, −4, 0, and 82 are even numbers,
//			while −3, 5, 23, and 69 are odd numbers.
//
//			The above definition of parity applies only to integer numbers, hence
//			it cannot be applied to numbers with decimals or fractions like 1/2 or
//			4.6978."
//
//	   NOTE: Fractions like '4.6978' do not qualify as 'even' numbers.
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

// IsValid - returns an error value signaling whether the current
// BigIntNum object is valid.
//
// If the current BigIntNum instance is invalid, an error is
// returned.
//
// If the current BigIntNum instance is valid, an error value of
// 'nil' is returned.
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
		"BigIntNum.IsZero",
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
// Performs a modulo operation where the current BigIntNum numeric
// value is the dividend and the divisor is the input parameter,
// 'divisor'.  The modulo operation finds the remainder after
// division of one number by another (sometimes called modulus).
// (Wikipedia: https://en.wikipedia.org/wiki/Modulo_operation)
//
//	dividend = bNum
//	dividend % divisor = modulo
//
// The result of this modulo operation is returned as a BigIntNum,
// 'modulo'. 'modulo' may consist of an integer or a floating
// point value consisting of integer and fractional digits.
//
// Input parameter 'maxPrecision' is used to control the maximum
// precision of the resulting floating point 'modulo'. Precision
// is defined as the number of fractional digits to the right of
// the decimal place. Be advised that these calculations can
// support very large precision values.
//
// The returned BigIntNum instance, 'modulo', will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the current BigIntNum instance (bNum).
func (bNum *BigIntNum) Mod(
	divisor BigIntNum,
	maxPrecision uint) (modulo BigIntNum, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.Mod",
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
		"BigIntNum.Multiply",
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
		"BigIntNum.MultiplyByTenToPowerAdd",
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

// NewBigInt
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated precision of type uint.
//
//	Precision Example
//	=================
//
//	The 'precision' input parameter specifies the number of digits
//	to the right of the decimal place. The Numeric value is equal
//	to 'bigI' x 10^(precision x -1). This effectively locates the
//	decimal place by counting from the extreme right of the integer
//	number, 'precision' places to the left. See the example below.
//
//		   Integer Value    precision    Numeric Value
//		     123456             3           123.456
//	                 123456 x 10^-3 =  123.456
//
//	Input Parameters
//	================
//
//	bigInt                   *big.Int
//	  'bigInt' is a type *big.Int and represents the integer
//	  value of the number; that is, the numeric value without
//	  decimal digits.
//
//	precision                uint
//	  This unsigned integer (always a positive value) identifies
//	  the location of the decimal place in the integer value
//	  parameter 'bigI'. The decimal place location is calculated
//	  by starting with the right most digit in the integer number
//	  ('bigI') and counting	left, 'precision' places.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  The new instance of BigIntNum will be returned through
//	  this parameter.
//
//	error
//	  If no errors are encountered during execution, this method
//	  will return an error value of 'nil'.
//
//	Numeric Separators
//	==================
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewBigInt(
	bigInt *big.Int,
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

	bNum2 := BigIntNum{}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		bigInt,
		precision,
		numSepsDto,
		ePrefix.XCpy(fmt.Sprintf("bigInt= '%v'  precision= '%v'",
			bigInt.Text(10), precision)))

	return bNum2, err
}

// NewBigIntNumSeps
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated precision of type uint.
//
//	Precision Example
//	=================
//
//	The 'precision' input parameter specifies the number of digits
//	to the right of the decimal place. The Numeric value is equal
//	to 'bigI' x 10^(precision x -1). This effectively locates the
//	decimal place by counting from the extreme right of the integer
//	number, 'precision' places to the left. See the example below.
//
//		   Integer Value    precision    Numeric Value
//		     123456             3           123.456
//	                 123456 x 10^-3 =  123.456
//
//	Input Parameters
//	================
//
//	bigInt                   *big.Int
//	  'bigI' is a type *big.Int and represents the integer value
//	  of the number; that is, the numeric value without decimal
//	   digits.
//
//	precision                uint
//	  This unsigned integer (always a positive value) identifies
//	  the location of the decimal place in the integer value
//	  parameter 'bigI'. The decimal place location is calculated
//	  by starting with the right most digit in the integer number
//	  ('bigI') and counting	left, 'precision' places.
//
//	numSepsDto               NumericSeparatorDto
//	  Input parameter, 'numSeps' consits of a NumericSeparatorDto
//	  instance. A NumericSeparatorDto contains symbols or characters
//	  for the decimal separator, thousands separator and currency
//	  symbol. These separators are used when parsing number strings
//	  into numeric values or displaying numeric values in number
//	  strings.
//
//	  If any of the 'numSeps' Numeric Separator Components are
//	  invalid, those components will be automatically reset to USA
//	  default values.
//
//	  The returned value ('BigIntNum') will be configured with
//	  'numSeps' Numeric Separators.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  The new instance of BigIntNum will be returned through
//	  this parameter.
//
//	error
//	  If no errors are encountered during execution, this method
//	  will return an error value of 'nil'.
func (bNum *BigIntNum) NewBigIntNumSeps(
	bigInt *big.Int,
	precision uint,
	numSepsDto NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigIntNumSeps",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bigInt == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigI'",
			}
	}

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		bigInt,
		precision,
		numSepsDto,
		ePrefix.XCpy(fmt.Sprintf("bigI= '%v'  precision= '%v'",
			bigInt.Text(10), precision)))

	return bNum2, err
}

// NewBigIntBigPrecision
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated *big.Int precision.
//
// The 'precision' parameter specifies the number of digits to the
// right of the decimal place. The Numeric value is equal to:
//
//	bigI x 10^(precision x -1)
//
// This effectively locates the decimal place by counting from the
// extreme right of the integer number, 'precision' places to the
// left. See the example below.
//
//	Input Parameters
//	================
//
//	bigI          *big.Int
//
//	'bigI' is a type *big.Int and represents the integer value of
//	the number; that is, the numeric value without decimal digits.
//
//
//	precision     *big.In
//
//	This integer value (always a positive value) identifies the
//	location of the decimal place in the integer value 'bigI'. The
//	decimal place location is calculated by starting with the
//	right most digit in the integer number and counting left,
//	'precision' places.
//
//	  Integer Value    precision    Numeric Value
//
//		    123456					 3					  123.456
//	             123456 x 10^-3  =    123.456
//
//	If precision is greater than the maximum value of an unsigned
//	integer (+4,294,967,295,	which equals 2^32 − 1), an error will
//	be triggered. Also, if the 'precision' value is less than zero,
//	an error will be triggered.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators specify the symbols or characters (runes)
//	used for the decimal separator, thousands separator and
//	currency symbol. These separators are used when displaying
//	numeric values in number strings.
//
//	The new BigIntNum instance returned by this method will contain
//	USA default numeric separators (decimal separator, thousands
//	separator and currency symbol).
func (bNum *BigIntNum) NewBigIntBigPrecision(
	bigInt *big.Int, precision *big.Int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigIntBigPrecision",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bINum2 := BigIntNum{}

	err = new(bigIntNumUtility).setBigIntBigPrecision(
		&bINum2,
		bigInt,
		precision,
		ePrefix)

	return bINum2, err
}

// NewBigIntBigPrecisionNumSeps
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated precision (also of type *big.Int).
//
// The 'precision' parameter specifies the number of digits to
// the right of the decimal place. The Numeric value is equal to
// bigI x 10^(precision x -1). This effectively locates the decimal
// place by counting from the extreme right of the integer number,
// 'precision' places to the left. See the example below.
//
//	Precision Example:
//	==================
//
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
// Numeric Seprators
// =================
//
// The returned BigIntNum instance will be configured with the
// Numeric Separators provided by input parameter 'numSeps'.
// Numeric Separatpors consist of decimal separators, thousands
// separators and a currency symbol.
//
//	Input Parameters
//	================
//
//	bigI 				*big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without decimal digits.
//
//
//	precision		*big.Int
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
//
//	numSeps			NumericSeparatorDto
//
//	The returned instance of BigIntNum will be configured with the
//	Numeric Separators contained in this input parameter, 'numSeps'.
//	Numeric Separatpors consist of decimal separators, thousands
//	separators and a currency symbol.
//
//	Return Parameters
//	=================
//
//	BigIntNum - a type BigIntNum numeric value
//
//	error			- If not 'nil', this prameter will
//							transmit any processing errors
//							encountered.
func (bNum *BigIntNum) NewBigIntBigPrecisionNumSeps(
	bigInt *big.Int,
	precision *big.Int,
	numSeps NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigIntBigPrecision",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bNum2 := BigIntNum{}

	err = new(bigIntNumUtility).setBigIntBigPrecisionNumSeps(
		&bNum2,
		bigInt,
		precision,
		numSeps,
		ePrefix.XCpy(fmt.Sprintf("bigInt= '%v'  precision= '%v'",
			bigInt.Text(10), precision)))

	return bNum2, err
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
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewBigIntExponent(
	bigI *big.Int, exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewBigIntExponent",
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
//	Input Parameters
//	================
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
//
// BE ADVISED
// ==========
//
// This method will set numeric separators to USA defaults.
//
//	decimal separator = '.'
//	thousands separator = ','
//	currency separator = '$'
func (bNum *BigIntNum) NewFromIntFracStrings(
	intStr string, fracStr string, signVal int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewFromIntFracStrings",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	b2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		&b2,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumNeutron).
		setIntFracStrings(&b2, intStr, fracStr, signVal,
			ePrefix.XCpy("Setting b2"))

	return b2, err
}

// NewFloat32 - Returns a new BigIntNum instance using a float32 floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
//	Input Parameters
//	================
//
//	f32							float32
//
//	This float32 value will be converted to a new instance of
//	BigIntNum.
//
//	maxPrecision		uint
//
//	The maximum precision for the result BigIntNum after
//	conversion of input parameter f64. Precision will never
//	be greater than 'maxPrecision'; however, actual precision
//	may be less than 'maxPrecision'.
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewFloat32(
	f32 float32, maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewFloat32",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bINum, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	ratNum := big.NewRat(1, 1).
		SetFloat64(float64(f32))

	err = new(bigIntNumMolecule).setBigRat(
		&bINum,
		ratNum,
		maxPrecision,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setBigRat(" +
					"    &bINum, ratNum, maxPrecision, ePrefix)",
				ErrContext: fmt.Sprintf("ratNum= '%v' maxPrecision= '%v'", ratNum.String(), maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return bINum, nil
}

// NewFloat64
// Returns a new BigIntNum instance using a float64 floating point
// input parameter.  The precision of the number is specified by the input
// parameter, 'decimalPlaces'.
//
//	 Input Parameters
//	 ================
//
//		f64						float64
//
//		This float64 value will be converted into an instance of
//		BigIntNum.
//
//		maxPrecision	uint
//
//		The maximum precision for the result BigIntNum after conversion
//		of input parameter f64. Precision will never be greater than
//		'maxPrecision'; however, actual precision may be less than
//		'maxPrecision'.
//
//		NOTE
//		====
//
//		The returned new instance of BigIntNum will contain default
//		USA numeric separators (decimal separator, thousands seprator,
//		and currency symbol)
func (bNum *BigIntNum) NewFloat64(f64 float64, maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewFloat64",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bINum, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	ratNum := big.NewRat(1, 1).
		SetFloat64(f64)

	err = new(bigIntNumMolecule).setBigRat(
		&bINum,
		ratNum,
		maxPrecision,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setBigRat(" +
					"    &bINum, ratNum, maxPrecision, ePrefix)",
				ErrContext: fmt.Sprintf("ratNum= '%v' maxPrecision= '%v'", ratNum.String(), maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return bINum, nil
}

// NewInt
//
// Creates a new BigIntNum instance initialized to the value
// of input parameter 'intNum' which is passed as type 'int'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ======
//
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
// =========
//
//	  intNum			precision				BigIntNum Result
//		123456					4								12.3456
//	  123456          0								123456
//	  123456          1								12345.6
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewInt(intNum int, precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewInt",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	b2, err := new(bigIntNumMechanics).newBigInt(
		big.NewInt(int64(intNum)), precision, ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by BigIntNum.NewBigInt()\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		&b2, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(\n" +
					"    b2, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return b2, nil
}

// NewIntExponent
// This method returns a new BigIntNum instance in which the
// numeric value is set using an integer multiplied by 10
// raised to the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'intNum' is of type int.
//
// Input parameter 'exponent' is of type int. Remember, for
// purposes of this method, 'exponent' will only affect digits
// to the right of the decimal place.
//
// Usage:
// ======
//
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
// =========
//
//	  intNum			 exponent			  BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewIntExponent(intNum int, exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewIntExponent",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bigI := big.NewInt(int64(intNum))

	bIntNum, err := new(bigIntNumMechanics).newZero(0, ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" bIntNum, err := new(BigIntNum).NewZero(0)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		&bIntNum, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(\n" +
					"    &bIntNum, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumMolecule).setBigIntExponent(
		&bIntNum, bigI, exponent, ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by:\n"+
				" err = b.SetBigIntExponent(bigI, exponent)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return bIntNum, nil
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
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewInt32(int32Num int32, precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewInt32",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bIntNum, err := new(bigIntNumMechanics).newBigInt(
		big.NewInt(int64(int32Num)), precision, ePrefix)

	if err != nil {

		return BigIntNum{}, &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bIntNum, err := new(bigIntNumMechanics).newBigInt(\n" +
				"   big.NewInt(int64(int32Num)), precision, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return bIntNum, nil
}

// NewInt32Exponent
//
// This method returns a new BigIntNum instance in which the
// numeric value is set using an integer multiplied by 10
// raised to the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'int32Num' is of type int32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ======
//
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
// =========
//
//	 int32Num		 exponent			  BigIntNum Result
//		123456				-3							123.456
//		123456				 3							123456.000
//	  123456				 0              123456
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewInt32Exponent(int32Num int32, exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewInt32Exponent",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bigI := big.NewInt(int64(int32Num))

	bIntNum, err := new(BigIntNum).NewZero(0)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bIntNum, err := new(BigIntNum).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumMolecule).setBigIntExponent(
		&bIntNum, bigI, exponent, ePrefix)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setBigIntExponent(\n" +
					"    &bIntNum, bigI, exponent, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bIntNum, nil
}

// NewInt64
//
// Creates a new BigIntNum instance initialized to the value
// of input parameter 'int64Num' which is passed as type 'int64'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ======
//
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
// =========
//
//	  int64Num			precision			BigIntNum Result
//		123456					4							 12.3456
//	  123456          0              123456
//	  123456          1              12345.6
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewInt64(int64Num int64, precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewInt64",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bIN2, err := new(bigIntNumMechanics).newBigInt(
		big.NewInt(int64Num), precision, ePrefix)

	if err != nil {

		return BigIntNum{}, &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bIN2, err := new(bigIntNumMechanics).newBigInt(\n" +
				"    big.NewInt(int64Num), precision, ePrefix))",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return bIN2, nil
}

// NewInt64Exponent
//
// This method returns a new BigIntNum instance in which the
// numeric value is set using an integer multiplied by 10 raised
// to the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ======
//
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
// =========
//
//	 int64Num		 exponent			  BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewInt64Exponent(
	int64Num int64,
	exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewInt64Exponent",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bigI := big.NewInt(int64Num)

	bIntNum, err := new(bigIntNumMechanics).newZero(0, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bIntNum, err := new(bigIntNumMechanics).newZero(\n" +
					"    0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumMolecule).setBigIntExponent(
		&bIntNum, bigI, exponent, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setBigIntExponent(\n" +
					"    &bIntNum, bigI, exponent, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bIntNum, nil
}

// NewIntAry
//
// Creates a new BigIntNum instance from an input parameter
// IntAry.
//
// Be careful, IntAry's can accommodate very, very large numbers.
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewIntAry(ia IntAry) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewIntAry",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = ia.IsValid(ePrefix.XCpy("Testing ia (IntAry)").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = ia.IsValid(ePrefix.XCpy(\n" +
					"    \"Testing ia (IntAry)\").String())",
				ErrContext: "'ia' is an input parameter for BigIntNum.NewAry(ia IntAry)",
				ErrMessage: err.Error(),
			}
	}

	bInt, err := ia.GetBigInt()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bInt, err := ia.GetBigInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	precision, err := ia.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "precision, err := ia.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bIntNum, err := new(bigIntNumMechanics).newZero(0, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bIntNum, err := new(bigIntNumMechanics).\n" +
					"    newZero(0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumNanobot).setBigInt(
		&bIntNum,
		bInt,
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumNanobot).setBigInt(\n" +
					"   &bIntNum, bInt, precision, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bIntNum, nil
}

// NewIntFracStr
//
// Creates a new BigIntNum instance based on a numeric value
// represented by separate integer and fractional components.
//
// Input parameters 'intStr' and 'fracStr' are strings
// representing the integer and fractional components. They
// are combined by this method to create a numeric value which
// is assigned to the current BigIntNum instance.
//
// Input parameter 'signVal' must be set to one of two values:
// +1 or -1. This value is used to signal the sign of the
// resulting numeric value. +1 generates a positive number and
// -1 generates a negative number. If input parameters 'inStr'
// or 'fracStr' contain a leading minus or plus sign character,
// it will be ignored. The sign of the resulting numeric value
// is controlled strictly by input parameter, 'signVal'.
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewIntFracStr(intStr, fracStr string, signVal int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewIntFracStr",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bIntNum, err := new(bigIntNumMechanics).newZero(
		0, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bIntNum, err := new(bigIntNumMechanics).newZero(\n" +
					"    0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumNeutron).setIntFracStrings(
		&bIntNum,
		intStr,
		fracStr,
		signVal,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumNeutron).setIntFracStrings(\n" +
					"    &bIntNum, intStr, fracStr, signVal, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bIntNum, nil
}

// NewINumMgr
//
// Receives an object which implements the INumMgr interface.
// The method then proceeds to create a new BigIntNum instance
// equivalent in numeric value to the input parameter, 'numMgr'.
// The BigIntNum instance is then returned to the calling
// function.
//
// Currently, the following 'mathops' Types implement the
// INumMgr interface:
//
//	Decimal,
//	IntAry,
//	NumStrDto,
//	BigIntNum
//	BigIntFixedDecimal
//
// Note: 'numMgr' must be a pointer to a type. This method will
// not accept 'numMgr' as a value. The pointer to the type is
// needed in or order to call methods on 'numMgr'.
//
// Example:
//
//	dec, err := new(Decimal).NewNumStr(nStr)
//	bINum, err := new(BigIntNum).NewINumMgr(&dec)
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain the
//	numeric separators passed by input parameter 'numMgr'.
//	Numeric separators consist of the decimal separator,
//	thousands seprator,	and currency symbol)
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

	err = new(bigIntNumMolecule).setINumMgr(
		&bINum, numMgr, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setINumMgr(\n" +
					"    &bINum, numMgr, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bINum, nil
}

// NewNumStr
//
//		 Receives a number string as input, converts that string to
//		 numeric value and returns a new BigIntNum instance encapsulating
//		 that numeric value.
//
//		 This method assumes that the input parameter 'numStr' is a
//		 string of numeric digits which will be delimited by default
//		 USA numeric separators. Default USA numeric separators are
//		 defined as:
//
//		         decimal separator = '.'
//		       thousands separator = ','
//		           currency symbol = '$'
//
//		 If the subject 'numStr' employs other national or cultural
//		 numeric separators, see the method:
//
//		   BigIntNum.NewNumStrWithNumSeps()
//
//
//		 Number Strings
//		 ==============
//
//		 Number strings are strings of numeric digits. These digits
//		 must be formatted in a way that facilitates conversion to a
//		 corresponding numeric value.
//
//		 Number String Negative Values
//		 =============================
//
//		 The 'numStr' number string parameter passed to this method must
//		 consist of a string of numeric digits representing a numeric
//		 value. A leading minus sign (-), or surrounding parentheses
//		 '()', may be included in this number string to indicate a
//		 negative numeric value.
//
//		 Fractional Digits in Number Strings
//		 ===================================
//
//		 The 'numStr' number string of numeric digits may also include
//		 a delimiting decimal separator to identify fractional digits to
//		 the right of the decimal separator. This method uses the
//		 default USA Decimal Separator ('.') to parse 'numStr' and
//		 identify any existing fractional digits.
//
//		 Numeric Separators
//		 ==================
//
//		 Numeric Separators define the Decimal Separator character,
//		 Thousands Separator character, and Currency Symbol character.
//		 These separator characters serve two purposes. First they are
//		 used to format and display numeric values as number strings.
//		 Second, they are also used to parse number strings and convert
//		 them into numeric values.
//
//		 For this method, the number string, 'numStr', will be parsed
//		 and converted to a numeric value based on the default USA
//		 Numeric Separators. Likewise, the returned instance of
//		 BigIntNum will be formatted with default USA Numeric
//		 Separators. If the 'numStr' input parameter employs other
//		 national or cultural numeric separators, see the method:
//
//		      BigIntNum.NewNumStrWithNumSeps()
//
//	  Input Parameters
//	  ================
//
//	  numStr                   string
//	    A string of numeric digits formatted as outlined above.
//	    Using default USA Numeric Separators, this method will
//	    parse the 'numStr' number string, convert it to a numeric
//	    value and return that value as an instanace of BigIntNum.
//
//		Return Values
//		=============
//
//		BigIntNum
//		  This type returns the converted numeric value of input
//		  parameter 'numStr'.
//
//		error
//		  If no errors are encountered during method execution, this
//		  returned error parameter is set to 'nil'.
func (bNum *BigIntNum) NewNumStr(
	numStr string) (BigIntNum, error) {

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

	// Sets numeric separators to USA Defaults
	bigINum, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigINum, err := new(bigIntNumMechanics).newZero(\n" +
					"0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSeps := NumericSeparatorDto{}
	numSeps.SetUSADefaults()

	err = new(bigIntNumMolecule).setNumStr(
		&bigINum,
		numStr,
		&numSeps,
		&numSeps,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setNumStr(\n" +
					"    &bigINum, numStr, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bigINum, nil
}

// NewNumStrWithNumSeps
//
//	Receives a number string as input, converts that string to a
//	numeric value and returns that value in new BigIntNum
//	instance. The input parameter 'numStrNumSeps' contains numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) which will be used to parse the number string and format
//	the returned BigIntNum value.
//
//	Number Strings
//	==============
//
//	Number strings are strings of numeric digits. These digits
//	must be formatted in a way that facilitates conversion to a
//	corresponding numeric value.
//
//	Number String Negative Values
//	=============================
//
//	The 'numStr' number string parameter passed to this method must
//	consist of a string of numeric digits representing a numeric
//	value. A leading minus sign (-), or surrounding parentheses
//	'()', may be included in this number string to indicate a
//	negative numeric value.
//
//	Fractional Digits in Number Strings
//	===================================
//
//	The 'numStr' number string of numeric digits may also include
//	a delimiting decimal separator to identify fractional digits to
//	the right of the decimal separator. In the USA, the default
//	decimal separator is the period character ('.'). The actual
//	decimal separator character used to parse the 'numStr' number
//	string is determined by the Numeric Separators parameter,
//	'numStrNumSeps'.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and convert
//	them into numeric values. Number string 'numStr' will be parsed
//	and converted to a numeric value based on the Numeric
//	Separators provided by input parameter, 'numStrNumSeps'.
//
//	Numeric Separator characters are typically encapsulated in a
//	type NumericSeparatorDto.
//
//	Input parameter 'numStrNumSeps' is of type IGetNumSeparators.
//	This interface type allows users to submit one of two types
//	for this parameter, a type NumericSeparatorDto instance or
//	a type NumericSeparatorPairDto.
//
//	If a type NumericSeparatorDto is submitted, the encapsulated
//	Numeric Separators will be used both to parse the number string
//	passed through input parameter 'numStr' and to format the
//	returned BigIntNum numeric value.
//
//	If a type NumericSeparatorPairDto is submitted for input
//	parameter 'numStrNumSeps', separate sets of Numeric Separators
//	will be used to parse the number string ('numStr') and format
//	the returned 'BigIntNum' numeric value. Type
//	NumericSeparatorPairDto contains two separate, embedded
//	instances of NumericSeparatorDto. The 'input'
//	NumericSeparatorDto will be used to parse the number string
//	while the 'output' NumericSeparatorDto will be used to format
//	the returned BigIntNum numeric value. The presence of two
//	separate sets of Numeric Separators allows users to parse a
//	number string formatted in one national number system while
//	formatting the returned value in a different national number
//	systems.
//
//	  Example:
//	    Input Format  = USA
//	    Output Format = European Union
//
//	If parameter 'numStrNumSeps' proves to be invalid, an error
//	will be returned.
//
//	Example Method Call
//	===================
//
//	** Pass numSeps as a pointer **
//	bIntNum, err := new(BigIntNum).NewNumStrWithNumSeps(outStr, &numSeps)
//
//	Input Parameters
//	================
//
//	numStr                   string
//	  This string value should be formatted as a string of
//	  numeric digits as outlined above. Using the Numeric
//	  Separators provided by input parameter 'numStrNumSeps',
//	  this method will parse the 'numStr' number string and
//	  convert it to a numeric value which will be returned as a
//	  BigIntNum.
//
//	numStrNumSeps            IGetNumSeparators
//	  The IGetNumSeparators interface type gives users the option
//	  of submitting one of two different concrete types.
//
//	  Users may choose to submit a type NumericSeparatorDto
//	  consisting of one set of Numeric Separators. These Numeric
//	  Separators will be used to both parse the number string
//	  provided by input parameter 'numStr' and format the returned
//	  BigIntNum type containing the converted numeric value.
//
//	  The second alternatives allows the user to submit a type
//	  NumericSeparatorPairDto for this parameter. This type
//	  encapsulates two separate instances of NumericSeparatorDto.
//	  The 'input' NumericSeparatorDto instance will be used to
//	  parse number string 'numStr' while the 'output' instance
//	  will be used to format the numeric value returned as a type
//	  BigIntNum.
//
//	Return Values
//	=============
//
//	BigIntNum
//	  This type returns the converted numeric value of input
//	  parameter 'numStr'.
//
//	error
//	  If no errors are encountered during method execution, this
//	  returned error parameter is set to 'nil'.
func (bNum *BigIntNum) NewNumStrWithNumSeps(
	numStr string, numStrNumSeps IGetNumSeparators) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewNumStrWithNumSeps()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if numStrNumSeps == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrNumSeps'",
			}
	}

	var inputNumSeps, outputNumSeps *NumericSeparatorDto

	inputNumSeps, err = numStrNumSeps.GetInputSeparators()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "inputNumSeps, err = numStrNumSeps.GetInputSeparators()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	outputNumSeps, err = numStrNumSeps.GetOutputSeparators()

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "outputNumSeps, err = numStrNumSeps.GetOutputSeparators()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	bINum2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINum2, err := new(bigIntNumMechanics).newZero(\n" +
					"    0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumMolecule).setNumStr(
		&bINum2,
		numStr,
		inputNumSeps,
		outputNumSeps,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setNumStr(\n" +
					"    &bINum2, numStr, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bINum2, nil
}

// NewNumStrMaxPrecision
//
//		Receives a number string as input, converts that string to a
//		numeric value and returns that value in new BigIntNum
//		instance. The input parameter 'numStrNumSeps' contains numeric
//		separators (decimal separator, thousands separator and currency
//		symbol) which will be used to parse the number string and format
//		the returned BigIntNum value.
//
//	 If the resulting precision exceeds input parameter
//	 'maxPrecision', the returned BigIntNum result will be rounded to
//	 'maxPrecision' decimal places.
//
//
//		Number Strings
//		==============
//
//		Number strings are strings of numeric digits. These digits
//		must be formatted in a way that facilitates conversion to a
//		corresponding numeric value.
//
//		Number String Negative Values
//		=============================
//
//		The 'numStr' number string parameter passed to this method must
//		consist of a string of numeric digits representing a numeric
//		value. A leading minus sign (-), or surrounding parentheses
//		'()', may be included in this number string to indicate a
//		negative numeric value.
//
//		Fractional Digits in Number Strings
//		===================================
//
//		The 'numStr' number string of numeric digits may also include
//		a delimiting decimal separator to identify fractional digits to
//		the right of the decimal separator. In the USA, the default
//		decimal separator is the period character ('.'). The actual
//		decimal separator character used to parse the 'numStr' number
//		string is determined by the Numeric Separators parameter,
//		'numStrNumSeps'.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//		them into numeric values. Number string 'numStr' will be parsed
//		and converted to a numeric value based on the Numeric
//		Separators provided by input parameter, 'numStrNumSeps'.
//
//		Numeric Separator characters are typically encapsulated in a
//		type NumericSeparatorDto.
//
//		Input parameter 'numStrNumSeps' is of type IGetNumSeparators.
//		This interface type allows users to submit one of two types
//		for this parameter, a type NumericSeparatorDto instance or
//		a type NumericSeparatorPairDto.
//
//		If a type NumericSeparatorDto is submitted, the encapsulated
//		Numeric Separators will be used both to parse the number string
//		passed through input parameter 'numStr' and to format the
//		returned BigIntNum numeric value.
//
//		If a type NumericSeparatorPairDto is submitted for input
//		parameter 'numStrNumSeps', separate sets of Numeric Separators
//		will be used to parse the number string ('numStr') and format
//		the returned 'BigIntNum' numeric value. Type
//		NumericSeparatorPairDto contains two separate, embedded
//		instances of NumericSeparatorDto. The 'input'
//		NumericSeparatorDto will be used to parse the number string
//		while the 'output' NumericSeparatorDto will be used to format
//		the returned BigIntNum numeric value. The presence of two
//		separate sets of Numeric Separators allows users to parse a
//		number string formatted in one national number system while
//		formatting the returned value in a different national number
//		systems.
//
//		  Example:
//		    Input Format  = USA
//		    Output Format = European Union
//
//		If parameter 'numStrNumSeps' proves to be invalid, an error
//		will be returned.
//
//		Input Parameters
//		================
//
//		numStr                   string
//		  This string value should be formatted as a string of
//		  numeric digits as outlined above. Using the Numeric
//		  Separators provided by input parameter 'numStrNumSeps',
//		  this method will parse the 'numStr' number string and
//		  convert it to a numeric value which will be returned as a
//		  BigIntNum.
//
//		numStrNumSeps            IGetNumSeparators
//		  The IGetNumSeparators interface type gives users the option
//		  of submitting one of two different concrete types.
//
//		  User may choose to submit a type NumericSeparatorDto
//		  consisting of one set of Numeric Separators. These Numeric
//		  Separators will be used to both parse the number strings
//		  provided by input parameter 'numstr' and format the returned
//		  BigIntNum type containing the converted numeric value.
//
//		  The second alternatives allows the user to submit a type
//		  NumericSeparatorPairDto for this parameter. This type
//		  encapsulates two separate instances of NumericSeparatorDto.
//		  The 'input' NumericSeparatorDto instance will be used to
//		  parse number string 'numStr' while the 'output' instance
//		  will be used to format the numeric value returned as a type
//		  BigIntNum.
func (bNum *BigIntNum) NewNumStrMaxPrecision(
	numStr string,
	maxPrecision uint,
	numStrNumSeps IGetNumSeparators) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewNumStrMaxPrecision()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bINum2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setNumStr(\n" +
					"    &bINum2, numStr, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var inputNumSeps, outputNumSeps *NumericSeparatorDto

	inputNumSeps, err = numStrNumSeps.GetInputSeparators()

	outputNumSeps, err = numStrNumSeps.GetOutputSeparators()

	err = new(bigIntNumMolecule).setNumStr(
		&bINum2,
		numStr,
		inputNumSeps,
		outputNumSeps,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).setNumStr(\n" +
					"    &bINum2, numStr, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if bINum2.precision > maxPrecision {

		err = new(bigIntNumBoson).roundToDecimalPlace(
			&bINum2, maxPrecision, ePrefix)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(bigIntNumBoson).roundToDecimalPlace(\n" +
						"    &bINum2, maxPrecision, ePrefix)",
					ErrContext: "if bINum2.precision > maxPrecision {",
					ErrMessage: err.Error(),
				}
		}
	}

	return bINum2, nil
}

// NewNumStrDto
//
//	 Receives a NumStrDto instance as input and returns
//	 a new BigIntNum instance.
//
//	 NOTE
//	 ====
//
//	 The returned new instance of BigIntNum will contain the numeric
//	 separators contained input parameter 'nDto'. Numeric separators
//	 consist of decimal separator, thousands seprator,	and currency
//	 symbol.
//
//		Input Parameters
//		================
//
//		nDto                     NewNumStrDto
//		  The NewNumStrDto type contains all the member elements
//		  necessary to convert the internal number string to a valid
//	   numeric value.
//
//		Return Values
//		=============
//
//		BigIntNum
//		  This type returns the converted numeric value of input
//		  parameter 'nDto'.
//
//		error
//		  If no errors are encountered during method execution, this
//		  returned error parameter is set to 'nil'.
func (bNum *BigIntNum) NewNumStrDto(nDto NumStrDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewNumStrDto()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bINum2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix.XCpy("Setting bINum2"))

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumMolecule).setNumStrDto(
		&bINum2,
		nDto,
		ePrefix.XCpy("Setting 'bINum2' = nDto"))

	return bINum2, err
}

// NewOne
//
// Returns a BigIntNum Type with a value equal to '1' (one).
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
//			0								1
//			1								1.0
//			2								1.00
//			3								1.000
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewOne(precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewOne()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		big.NewInt(1),
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// NewTwo
//
// Returns a BigIntNum Type with a value equal to  '2' (two).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter
// 'precision'.
//
// To create an integer with a value equal to '1', set
// 'precision' equal to zero (0).
//
// Examples:
// =========
//
//	precision
//	  Value		Result
//			0				2
//			1				2.0
//			2				2.00
//			3				2.000
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewTwo(precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewTwo()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		big.NewInt(2),
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// NewThree
//
// Returns a BigIntNum Type with a value equal to  '3' (three).
// The number of zeros created after the decimal placeholder
// (fractional digits) is determined by the input parameter
// 'precision'.
//
// To create an integer with a value equal to '1', set 'precision'
// equal to zero (0).
//
// Examples:
// =========
//
//	precision
//	  Value 					Result
//			0								3
//			2								3.00
//			3								3.000
//
//		NOTE
//		====
//
//		The returned new instance of BigIntNum will contain default
//		USA numeric separators (decimal separator, thousands seprator,
//		and currency symbol)
func (bNum *BigIntNum) NewThree(precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewThree()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		big.NewInt(3),
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// NewFive
//
// Returns a BigIntNum with integer value of  '5' (five). The
// number of zeros created after the decimal placeholder (fractional
// digits) is determined by the input parameter 'precision'.
//
// To create an integer with a value equal to '10', set 'precision'
// equal to zero (0).
//
//	precision
//		Value 		Result
//			0					 5
//			1					 5.0
//			2					 5.00
//			3					 5.000
//
//		NOTE
//		====
//
//		The returned new instance of BigIntNum will contain default
//		USA numeric separators (decimal separator, thousands seprator,
//		and currency symbol)
func (bNum *BigIntNum) NewFive(precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewFive()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		big.NewInt(5),
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// NewTen
//
// Returns a BigIntNum with integer value of  '10' (ten). The
// number of zeros created after the decimal placeholder (fractional
// digits) is determined by the input parameter 'precision'.
//
// To create an integer with a value equal to '10', set 'precision'
// equal to zero (0).
//
//	precision
//		Value			Result
//			0					10
//			1					10.0
//			2					10.00
//			3					10.000
//
//		NOTE
//		====
//
//		The returned new instance of BigIntNum will contain default
//		USA numeric separators (decimal separator, thousands seprator,
//		and currency symbol)
func (bNum *BigIntNum) NewTen(precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewTen()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		big.NewInt(10),
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// NewUint
//
// Creates a new BigIntNum instance initialized to the value
// of input parameter 'uintNum' which is passed as type 'uint'.
//
// Input parameter 'precision' indicates the number of digits to
// be formatted to the right of the decimal place and is passed
// as type 'uint'.
//
// Usage:
// ======
//
// This method is designed to be used in conjunction with the
// new(BigIntNum) syntax thereby allowing BigIntNum type creation
// and initialization in one step.
//
//	uintNum := uint(123456)
//	precision := uint(3)
//	bINum := new(BigIntNum).NewUint(uintNum, precision)
//	bINum is now equal to 123.456
//
// Examples:
// =========
//
//	  uintNum			precision			BigIntNum Result
//		123456					4							12.3456
//	  123456          0             123456
//	  123456          1             12345.6
//
//		NOTE
//		====
//
//		The returned new instance of BigIntNum will contain default
//		USA numeric separators (decimal separator, thousands seprator,
//		and currency symbol)
func (bNum *BigIntNum) NewUint(uintNum uint, precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewUint()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		big.NewInt(0).SetUint64(uint64(uintNum)),
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// NewUintExponent
//
// This method returns a new BigIntNum instance in which the
// numeric value is set using an integer multiplied by 10 raised
// to the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uintNum' is of type uint.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ======
//
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//		biNum := new(BigIntNum).NewUintExponent(123456, -3)
//	 -- biNum is now equal to "123.456", precision = 3
//
//		biNum := new(BigIntNum).NewUintExponent(123456, 3)
//	 -- biNum is now equal to "123456.000", precision = 3
//
// Examples:
// =========
//
//	  uintNum			exponent			BigIntNum Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//		 123456          0              123456
//
//		NOTE
//		====
//
//		The returned new instance of BigIntNum will contain default
//		USA numeric separators (decimal separator, thousands seprator,
//		and currency symbol)
func (bNum *BigIntNum) NewUintExponent(uintNum uint, exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewUint()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	baseBInt := big.NewInt(int64(uintNum))

	// Sets Numeric Separators to default USA format
	bIntNum2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix.XCpy("Setting bIntNum2=0"))

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumMolecule).
		setBigIntExponent(
			&bIntNum2,
			baseBInt,
			exponent,
			ePrefix.XCpy(fmt.Sprintf("Setting 'bIntNum2' baseBInt= '%v' exponent= '%v'",
				baseBInt.Text(10), exponent)))

	return bIntNum2, err
}

// NewUint32
//
// Creates a new BigIntNum instance initialized to the value
// of input parameter 'uint32Num' which is passed as type 'uint32'.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'.
//
//	Usage
//	=====
//
// This method is designed to be used in conjunction with the BigIntNum{}
// syntax thereby allowing BigIntNum type creation and initialization in
// one step.
//
//					uint32Num := uint32(123456)
//					precision := uint(3)
//					bINum := BigIntNum{}.NewUint32(uint32Num, precision)
//					bINum is now equal to 123.456
//
//	Examples
//	========
//
// uint32Num  precision		Result
//
//	 123456       4       12.3456
//	 123456       0       123456
//	 123456       1       12345.6
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewUint32(
	uint32Num uint32, precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewUint32()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bigInt := big.NewInt(0).SetUint64(uint64(uint32Num))

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		bigInt,
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// NewUint32Exponent
//
// This method returns a new BigIntNum instance in which the
// numeric value is set using an integer multiplied by 10 raised
// to the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uint32Num' is of type uint32.
//
// Input parameter 'exponent' is of type int.
//
//	 Usage
//	 =====
//
//	 This method is designed to be used in conjunction with the BigIntNum{}
//	 syntax thereby allowing BigIntNum type creation and initialization in
//	 one step.
//
//			biNum := BigIntNum{}.NewUint32Exponent(123456, -3)
//		 -- biNum is now equal to "123.456", precision = 3
//
//			biNum := BigIntNum{}.NewUint32Exponent(123456, 3)
//		 -- biNum is now equal to "123456.000", precision = 3
//
//		Examples
//		========
//
//	 uint32Num exponent Result
//	  123456     -3     123.456
//	  123456      3     123456.000
//	  123456      0     123456
//
//	 NOTE
//	 ====
//
//	 The returned new instance of BigIntNum will contain default
//	 USA numeric separators (decimal separator, thousands seprator,
//	 and currency symbol)
func (bNum *BigIntNum) NewUint32Exponent(
	uint32Num uint32, exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewUint32Exponent",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	baseBInt := big.NewInt(int64(uint32Num))

	// Sets Numeric Separators to default USA format
	bIntNum2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix.XCpy("Setting bIntNum2=0"))

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumMolecule).
		setBigIntExponent(
			&bIntNum2,
			baseBInt,
			exponent,
			ePrefix.XCpy(fmt.Sprintf("Setting 'bIntNum2'; baseBInt= '%v' exponent= '%v'",
				baseBInt.Text(10), exponent)))

	return bIntNum2, err
}

// NewUint64
//
//	Creates a new BigIntNum instance initialized to the value of
//	input parameter 'uint64Num' which is passed as type 'uint64'.
//
//	Input parameter 'precision' indicates the number of digits to
//	be formatted to the right of the decimal place and is passed
//	as type 'uint'
//
//	Usage
//	=====
//
//	This method is designed to be used in conjunction with the
//	'new' key word.
//
//	  uint64Num := uint64(123456)
//	  precision := uint(3)
//	  bINum := new(BigIntNum.NewUint64(uint64Num, precision)
//	  bINum is now equal to 123.456
//
//	Examples
//	=========
//
//	uint64Num  precision  Result
//	  123456       4      12.3456
//	  123456       0      123456
//	  123456       1      12345.6
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol)
func (bNum *BigIntNum) NewUint64(
	uint64Num uint64, precision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewUint64()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	newVal := big.NewInt(0).SetUint64(uint64Num)

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		newVal,
		precision,
		numSepsDto,
		ePrefix.XCpy(fmt.Sprintf("Setting 'bIntNum2' newVal= '%v' precision= '%v'",
			newVal.Text(10), precision)))

	return bNum2, err

}

// NewUint64Exponent
//
// This method returns a new BigIntNum instance in which the
// numeric value is set using an integer multiplied by 10 raised
// to the power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
//
// Input parameter 'uint64Num' is of type uint64.
//
// Input parameter 'exponent' is of type int.
//
//	 Usage
//	 ======
//
//	 This method is designed to be used in conjunction with the BigIntNum{}
//	 syntax thereby allowing BigIntNum type creation and initialization in
//	 one step.
//
//			biNum := BigIntNum{}.NewUint64Exponent(123456, -3)
//		 -- biNum is now equal to "123.456", precision = 3
//
//			biNum := BigIntNum{}.NewUint64Exponent(123456, 3)
//		 -- biNum is now equal to "123456.000", precision = 3
//
//	 Examples
//	 =========
//
//	 uint64Num  exponent  Result
//	  123456       -3     123.456
//	  123456        3     123456.000
//	  123456        0     123456
//
//	 NOTE
//	 ====
//
//	 The returned new instance of BigIntNum will contain default
//	 USA numeric separators (decimal separator, thousands seprator,
//	 and currency symbol)
func (bNum *BigIntNum) NewUint64Exponent(uint64Num uint64, exponent int) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewUint64Exponent()",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	baseBInt := big.NewInt(0).SetUint64(uint64Num)

	bIntNum2, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix.XCpy("Setting bIntNum2=0"))

	if err != nil {
		return BigIntNum{}, err
	}

	err = new(bigIntNumMolecule).
		setBigIntExponent(
			&bIntNum2,
			baseBInt,
			exponent,
			ePrefix.XCpy(fmt.Sprintf("Setting 'bIntNum2' baseBInt= '%v' precision= '%v'",
				baseBInt.Text(10), exponent)))

	return bIntNum2, err
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
//	  value    Result
//	    0        0
//	    2        0.00
//	    3        0.000
//
//	NOTE
//	====
//
//	The returned new instance of BigIntNum will contain default
//	USA numeric separators (decimal separator, thousands seprator,
//	and currency symbol).
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

	numSepsDto := NumericSeparatorDto{}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2 := BigIntNum{}

	err = new(bigIntNumNanobot).setBigIntNumSeps(
		&bNum2,
		big.NewInt(0),
		precision,
		numSepsDto,
		ePrefix)

	return bNum2, err
}

// Reset
//
// Resets the current BigIntNum to a new valid BigIntNum using
// the BigIntNum components BigIntNum.bigInt and
// BigIntNum.precision.
//
// This method is usually called after method BigIntNum.IsValid()
// returns an err.
//
// Calling this method causes Numeric Separators to be reset
// to USA Defaults.
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

// RoundToDecPlace
//
// Rounds the current BigIntNum instance to a specified number of
// decimal places.
//
// 'precision' equals the number of digits to the right of the
// decimal place.
//
// Example:
//
//	integer= 123456; precision = 3; Numeric Value= 123.456
//
// If the value of BigIntNum.bigInt is zero ('0'), that zero value
// will remain unaltered. However, the BigIntNum.precision value
// will be set equal to input parameter, 'precision'.
//
// If the number of decimal places specified for rounding
// ('precision') is equal to the current BigIntNum.precision, no
// action is taken.
//
// If the number of decimal places specified for rounding
// ('precision') is greater than the current BigIntNum.precision
// value, trailing zeros are added to the current BigIntNum.bigInt
// value and BigIntNum.precision is set equal to input parameter,
// 'precision'.
//
// Finally, if the number of decimal places specified for rounding
// ('precision') is less than the current BigIntNum.precision
// value, the fractional digits will be rounded in accordance with
// the input parameter, 'precision'.
//
//	Examples
//	========
//
//	  Original       precision       Resulting
//	   Value       input parameter     Value
//	 654.123456          9            654.123456000
//	 654.123456          4            654.1235
//
//	-654.123456          9           -654.123456000
//	-654.123456          4           -654.1235
//
//	  0                  3              0.000
//	  0.000000           0              0
//
// Existing numeric separators (decimal separator, thousands
// separator and currency symbol) remain unchanged and are not
// altered by this method.
func (bNum *BigIntNum) RoundToDecPlace(precision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.RoundToDecPlace",
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

	return new(bigIntNumBoson).roundToDecimalPlace(
		bNum, precision, ePrefix)
}

// SetBigInt
//
//	Sets the value of the current BigIntNum instance using
//	the input parameters *big.Int integer and precision.
//
//	The 'precision' parameter specifies the number of digits to the right
//	of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
//	This effectively locates the decimal place by counting from the extreme right
//	of the integer number, 'precision' places to the left. See the example below.
//
//	Input Parameters
//	================
//
//	bigI			*big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without
//	decimal digits.
//
//
//	precision		uint
//
//	This unsigned integer (always a positive value) identifies
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places. Example:
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
//	Existing numeric separators (decimal separator, thousands separator
//	and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) SetBigInt(bigI *big.Int, precision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetBigInt",
		"")

	if err != nil {
		return err
	}

	if bigI == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigI'",
		}
	}

	return new(bigIntNumNanobot).setBigInt(
		bNum,
		bigI,
		precision,
		ePrefix.XCpy(
			fmt.Sprintf("Setting 'bNum' 'bigI'='%v' precision= '%v'",
				bigI.Text(10), precision)))
}

// SetBigIntNumSeps
//
// Sets the value of the current BigIntNum instance using
// the input parameters *big.Int integer and precision.
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
//	Input Parameters
//	================
//
//	bigI			*big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without
//	decimal digits.
//
//
//	precision		uint
//
//	This unsigned integer (always a positive value) identifies
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places. Example:
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
//
//	numSeps				NumericSeparatorDto
//
//	Input parameter 'bNum' (type BigIntNum) will be reconfigured
//	with numeric separators provided by input prameter 'numSeps',
//	an instance of NumericSeparatorDto. Type NumericSeparatorDto
//	contains the decimal separator, thousands separator and currency
//	symbol.
//
//		type NumericSeparatorDto struct {
//			DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//			ThousandsSeparator rune // Character used to separate thousands (1,000,000,000)
//			CurrencySymbol     rune // Currency Symbol
//		}
func (bNum *BigIntNum) SetBigIntNumSeps(
	bigI *big.Int,
	precision uint,
	numSeps NumericSeparatorDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetBigIntNumSeps",
		"")

	if err != nil {
		return err
	}

	if bigI == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigI'",
		}
	}

	return new(bigIntNumMolecule).setBigIntNumSeps(
		bNum,
		bigI,
		precision,
		numSeps,
		ePrefix.XCpy(
			fmt.Sprintf("Setting 'bNum' 'bigI'='%v' precision= '%v'",
				bigI.Text(10), precision)))
}

// SetBigIntBigPrecision
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated *big.Int precision.
//
// The 'precision' parameter specifies the number of digits to the
// right of the decimal place. The Numeric value is equal to:
//
//	bigInt x 10^(precision x -1)
//
// This effectively locates the decimal place by counting from the
// extreme right of the integer number, 'precision' places to the
// left. See the example below.
//
//	Input Parameters
//	================
//
//	bigInt          *big.Int
//
//	'bigInt' is a type *big.Int and represents the integer value
//	 of the number; that is, the numeric value without decimal
//	 digits.
//
//
//	precision       *big.Int
//
//	This integer value (always a positive value) identifies the
//	location of the decimal place in the integer value 'bigInt'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting left,
//	'precision' places.
//
//	  Integer Value    precision    Numeric Value
//
//		    123456					 3					  123.456
//	             123456 x 10^-3  =    123.456
//
//	If precision is greater than the maximum value of an unsigned
//	integer (+4,294,967,295,	which equals 2^32 − 1), an error will
//	be triggered. Also, if the 'precision' value is less than zero,
//	an error will be triggered.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators specify the symbols or characters (runes)
//	used for the decimal separator, thousands separator and
//	currency symbol. These separators are used when displaying
//	numeric values in number strings.
//
//	The Numeric Separators previous configured in the current
//	instance of BigIntNum will remain unchanged and will NOT
//	be altered by this method. However, if any of the current
//	Numeric Separators are invalid, thy will be automatically
//	reset to USA default values.
func (bNum *BigIntNum) SetBigIntBigPrecision(
	bigInt *big.Int,
	precision *big.Int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetBigInt",
		"")

	if err != nil {
		return err
	}

	numSepsDto := NumericSeparatorDto{
		DecimalSeparator:   bNum.decimalSeparator,
		ThousandsSeparator: bNum.thousandsSeparator,
		CurrencySymbol:     bNum.currencySymbol,
	}

	numSepsDto.SetDefaultsIfEmpty()

	return new(bigIntNumUtility).setBigIntBigPrecisionNumSeps(
		bNum,
		bigInt,
		precision,
		numSepsDto,
		ePrefix.XCpy(
			fmt.Sprintf("Setting 'bNum' 'bigI'='%v' precision= '%v'",
				bigInt.Text(10), precision)))
}

// SetBigIntBigPrecisionNumSeps
//
// Reconfigures the current BigIntNum instance using a *big.Int
// type and its associated precision (also of type *big.Int).
//
// The 'precision' parameter specifies the number of digits to
// the right of the decimal place. The Numeric value is equal to
// bigI x 10^(precision x -1). This effectively locates the decimal
// place by counting from the extreme right of the integer number,
// 'precision' places to the left. See the example below.
//
//	Precision Example:
//	==================
//
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
//	Numeric Seprators
//	=================
//
//	The returned BigIntNum instance will be configured with the
//	Numeric Separators provided by input parameter 'numSeps'.
//	Numeric Separatpors consist of decimal separators, thousands
//	separators and a currency symbol.
//
//	Input Parameters
//	================
//
//	bigI 				             *big.Int
//	  'bigI' is a type *big.Int and represents the integer
//	  value of the number; that is, the numeric value without
//	  decimal digits.
//
//
//	precision		              *big.Int
//	  This integer value (always a positive value) identifies
//	  the location of the decimal place in the integer value 'bigI'.
//	  The decimal place location is calculated by starting with the
//	  right most digit in the integer number and counting	left,
//	  'precision' places. If precision is greater than the maximum
//	  value of an unsigned integer (+4,294,967,295,	which equals
//	  2^32 − 1), an error will be triggered. Also, if the 'precision'
//	  value is less than zero, an error will be triggered.
//
//
//	numSeps			             NumericSeparatorDto
//	  The returned instance of BigIntNum will be configured with the
//	  Numeric Separators contained in this input parameter, 'numSeps'.
//	  Numeric Separatpors consist of decimal separators, thousands
//	  separators and a currency symbol.
//
//	Return Parameters
//	=================
//
//	error
//	  If no errors are encountered during execution, this method
//	  will return an error value of 'nil'.
func (bNum *BigIntNum) SetBigIntBigPrecisionNumSeps(
	bigInt *big.Int,
	precision *big.Int,
	numSeps NumericSeparatorDto) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetBigIntBigPrecisionNumSeps",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	bNum2 := BigIntNum{}

	err = new(bigIntNumUtility).setBigIntBigPrecisionNumSeps(
		&bNum2,
		bigInt,
		precision,
		numSeps,
		ePrefix.XCpy(fmt.Sprintf("bigInt= '%v'  precision= '%v'",
			bigInt.Text(10), precision)))

	return bNum2, err
}

// SetBigIntExponent
//
// Sets the numeric value using an integer multiplied by 10 raised
// to the power of the 'exponent' parameter.
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
//	 bigI     exponent     BigIntNum Result
//	123456      -3             123.456
//
// If exponent is greater than 0, bigI is multiplied by 10 raised to the
// power of exponent and precision is set equal to exponent.
//
//	 bigI      exponent     BigIntNum Result
//	123456         3            123456.000
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this
// method.
func (bNum *BigIntNum) SetBigIntExponent(
	bigI *big.Int, exponent int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetBigIntExponent()",
		"")

	if err != nil {
		return err
	}

	if bigI == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigI'",
		}
	}

	return new(bigIntNumMolecule).
		setBigIntExponent(bNum,
			bigI,
			exponent,
			ePrefix.XCpy(fmt.Sprintf("Setting 'bNum' bigI= '%v' exponent= '%v'",
				bigI.Text(10), exponent)))
}

// SetBigFloat
//
// Sets the value of the current BigIntNum using a *big.Float
// floating point input parameter.  The precision of the number
// is specified by the input parameter, 'maxPrecision'.
//
//	 Input Parameters
//	 ================
//
//		bigFloat *big.Float
//
//		This *big.Float value will be converted saved as the
//		current instance of BigIntNum
//
//		maxPrecision uint
//
//		The maximum precision for the resulting BigIntNum value
//		after conversion of input parameter 'bigFloat'. Final
//		precision will never be greater than 'maxPrecision';
//		however, actual precision may be less than
//		'maxPrecision'.
//
//	 Background
//	 ==========
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
		"BigIntNum.SetBigFloat",
		"")

	if err != nil {
		return err
	}

	if bigFloat == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigI'",
		}
	}

	return new(bigIntNumNanobot).setBigFloat(
		bNum,
		bigFloat,
		maxPrecision,
		ePrefix.XCpy("Setting 'bNum'"))
}

// SetBigRat
//
// Sets the value of the current BigIntNum instance to that of
// input parameter 'ratNum', a rational number of type *big.Rat.
//
//		Input Parmeters
//		===============
//
//	 ratNum         *big.Rat
//
//		The value of ratNum will be used to configure the current
//		BigIntNum instance and reset its value.
//
//	 maxPrecision   uint
//
//		The maximum precision for the resulting BigIntNum value
//		after it is reset to the value of input parameter 'ratNum'.
//		Precision will never be greater than 'maxPrecision'; however,
//		actual precision may be less than 'maxPrecision'.
//
//		Existing numeric separators (decimal separator, thousands separator
//		and currency symbol) in the current BigIntNum instance will remain
//		unchanged and will not be altered by this method.
//
//	 NOTE
//	 ====
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) contained in the original BigIntNum ('bNum')
// will remain unchanged and will not be altered by this method.
func (bNum *BigIntNum) SetBigRat(
	ratNum *big.Rat, maxPrecision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetBigRat",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumMolecule).setBigRat(
		bNum,
		ratNum,
		maxPrecision,
		ePrefix.XCpy(fmt.Sprintf("Setting 'bNum'; ratNum= '%v'  maxPrecision= '%v'",
			ratNum.String(), maxPrecision)))
}

// SetBigRatNumSeps
//
// Sets the value of the current BigIntNum instance to that of
//
//		input parameter 'ratNum', a rational number of type *big.Rat.
//
//		Input Parmeters
//		===============
//
//	 ratNum         *big.Rat
//
//			The value of ratNum will be used to configure the current
//			BigIntNum instance and reset its value.
//
//	 maxPrecision   uint
//
//	   The maximum precision for the resulting BigIntNum value
//	   after it is reset to the value of input parameter 'ratNum'.
//	   Precision will never be greater than 'maxPrecision'; however,
//	   actual precision may be less than 'maxPrecision'.
//
//	 numSeps        NumericSeparatorDto
//
//	   The current BigIntNum instance will be reconfigured with
//	   numeric separators provided by 'numSeps' an instance of
//	   NumericSeparatorDto. Type NumericSeparatorDto contains
//	   the decimal separator, thousands separator and currency
//	   symbol.
//
//	   type NumericSeparatorDto struct {
//	     DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//	     ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
//	     CurrencySymbol     rune // Currency Symbol
//	   }
func (bNum *BigIntNum) SetBigRatNumSeps(
	ratNum *big.Rat,
	maxPrecision uint,
	numSeps NumericSeparatorDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetBigRat",
		"")

	if err != nil {
		return err
	}

	err = new(bigIntNumMolecule).setBigRat(
		bNum,
		ratNum,
		maxPrecision,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumMolecule).setBigRat(\n" +
				"    bNum, ratNum, maxPrecision, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numSeps.SetDefaultsIfEmpty()

	return new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix.XCpy("Setting 'bNum' Numeric Separators"))
}

// SetCurrencySymbol
//
// Assigns the input parameter rune as the currency symbol for
// the current instance of BigIntNum. The currency symbol is used
// when generating number strings for display.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
//	Example: $123
//
// Note: If a zero value is submitted for input parameter
// 'currencySymbol', an error will be returned.
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
func (bNum *BigIntNum) SetCurrencySymbol(currencySymbol rune) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetCurrencySymbol",
		"")

	if err != nil {
		return err
	}

	var nSepSymbolCode NumSepSymbolCode

	nSepSymbolCode = CURRENCYSYMBOL

	return new(bigIntNumBoson).setNumSepSymbol(
		bNum,
		nSepSymbolCode,
		currencySymbol,
		ePrefix.XCpy("Setting 'bNum' currency symbol"))
}

// SetDecimalSeparator
//
// Configures the input parameter rune as the Decimal Separator for
// the current instance of BigIntNum.
//
// The Decimal Separator is used to separate the integer and
// fractional elements of a number string.
//
// In the USA, the Decimal Separator is a period character ('.').
//
//	Example: 123.45
//
// Note: If a zero value is submitted for input parameter
// 'decimalSeparator', an error will be returned.
func (bNum *BigIntNum) SetDecimalSeparator(
	decimalSeparator rune) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetDecimalSeparator",
		"")

	if err != nil {
		return err
	}

	var nSepSymbolCode NumSepSymbolCode

	nSepSymbolCode = DECIMALSYMBOL

	return new(bigIntNumBoson).setNumSepSymbol(
		bNum,
		nSepSymbolCode,
		decimalSeparator,
		ePrefix.XCpy("Setting 'bNum' Decimal Separator"))
}

// SetIntFracStrings
//
// Sets the value of the current BigIntNum instance based on a numeric value
// represented by separate integer and fractional components.
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
// and currency symbol), contained in the current instance of BigIntNum,
// will remain unchanged and are not altered by this method.
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

	return new(bigIntNumNeutron).setIntFracStrings(
		bNum,
		intStr,
		fracStr,
		signVal,
		ePrefix.XCpy("Setting 'bNum'"))
}

// SetFloat32
//
// Sets the value of a BigIntNum using a float32 floating point
// input parameter.  The precision of the number is specified by
// the input parameter, 'decimalPlaces'.
//
//	 Input Parameters
//	 ================
//
//		f32            float32
//
//		This float32 value will be converted into an instance of
//		BigIntNum.
//
//
//		maxPrecision    uint
//
//		The maximum precision for the resulting BigIntNum after
//		conversion of input parameter 'ratNum'. Precision will
//		never be greater than 'maxPrecision'; however, actual
//		precision may be less than 'maxPrecision'.
//
//		NOTE
//		====
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) configured for the current instance of BigintNum remain unchanged and are not altered by this method.
func (bNum *BigIntNum) SetFloat32(f32 float32, maxPrecision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetFloat32",
		"")

	if err != nil {
		return err
	}

	ratNum := big.NewRat(1, 1).
		SetFloat64(float64(f32))

	return new(bigIntNumMolecule).setBigRat(
		bNum,
		ratNum,
		maxPrecision,
		ePrefix.XCpy("Setting 'bNum'"))
}

// SetFloat64
//
// Sets the value of a BigIntNum using a float64 floating point
// input parameter.  The precision of the number is specified
// by the input parameter 'decimalPlaces'.
//
//	 Input Parameters
//	 ================
//
//		f64              float64
//
//		This float64 value will be converted into an instance of
//		BigIntNum.
//
//		maxPrecision     uint
//
//		The maximum precision for the resulting BigIntNum after
//		conversion of input parameter 'f64'. Resulting precision
//		will never be greater than 'maxPrecision'; however, actual
//		precision may be less than 'maxPrecision'.
func (bNum *BigIntNum) SetFloat64(f64 float64, maxPrecision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetFloat64",
		"")

	if err != nil {
		return err
	}

	ratNum := big.NewRat(1, 1).SetFloat64(f64)

	return new(bigIntNumMolecule).setBigRat(
		bNum,
		ratNum,
		maxPrecision,
		ePrefix.XCpy(fmt.Sprintf("Setting 'bNum'; ratNum= '%v' maxPrecistion= '%v'",
			ratNum.String(), maxPrecision)))
}

// SetExpectedNumberOfDigits
//
// Sets the number of expected digits associated with the Absolute
// Value of this 'BigIntNum.absBigInt'. The value is stored in the
// data field, 'BigIntNum.numberOfExpectedDigits'.
//
// Useful in tracking leading zeros.
func (bNum *BigIntNum) SetExpectedNumberOfDigits(numOfDigits *big.Int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetExpectedNumberOfDigits",
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

	return new(bigIntNumMolecule).setExpectedNumberOfDigits(
		bNum,
		numOfDigits,
		ePrefix)
}

// SetExpectedToActualNumberOfDigits
//
//	Sets the 'Expected' number of numeric digits associated with this
//	BigIntNum, to the actual number of numeric digits in the BigIntNum
//	value at the time when this method is called.
func (bNum *BigIntNum) SetExpectedToActualNumberOfDigits() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetExpectedNumberOfDigits",
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

	actNumOfDigits, _, err := new(bigIntNumMolecule).
		getActualNumberOfDigits(
			bNum,
			ePrefix.XCpy("bNum Actual Num digits"))

	if err != nil {
		return err
	}

	bNum.numberOfExpectedDigits = big.NewInt(0).Set(actNumOfDigits)

	return nil
}

// SetINumMgr
//
// Receives an input parameter implementing the INumMgr interface
// and proceeds to set the current BigIntNum instance to its
// equivalent numeric value.
//
// Currently, the following 'mathops' Types implement the
// INumMgr interface:
//
//	Decimal,
//	IntAry,
//	NumStrDto,
//	BigIntNum
//
// 'numMgr' must be a pointer to a type. This method will not
// accept 'numMgr' as a value. The pointer to the type is needed
// in or order to call methods on 'numMgr'.
//
// This method will test the validity of input parameter,
// 'numMgr'.
//
// Example 1:
//
//	dec, err := new(Decimal).NewNumStr(nStr)
//	bINum := new(BigIntNum)
//	err := bINum.SetINumMgr(&dec)
//
// Example 2:
// dec, err := new(Decimal).NewNumStr(nStr)
// bINum := new(BigIntNum)
// err := bINum.SetINumMgr(dec.GetThisPointer())
//
// Example 3:
// dec := new(Decimal).NewPtr()
// err := dec.SetNumStr(nStr)
// bINum := new(BigIntNum)
// err := bINum.SetINumMgr(dec)
//
// Example 4:
// fd := new(BigIntFixedDecimal).NewZero()
func (bNum *BigIntNum) SetINumMgr(numMgr INumMgr) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetINumMgr",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumMolecule).setINumMgr(
		bNum, numMgr, ePrefix.XCpy("numMgr->bNum"))
}

// SetPrecision
//
// Sets a new 'precision' value for the current BigIntNum
// instance. The new 'precision' is specified by the uint type
// input parameter, 'newPrecision'.
//
// Precision is defined as the number of numeric digits to right
// of the decimal place.
//
// If 'newPrecision' is equal to the current BigIntNum.precision
// value, no action is taken and the original BigIntNum numeric
// value remains unchanged.
//
// If 'newPrecision' is greater than the current
// BigIntNum.precision value, trailing zeros are added to the
// fractional digits to the right of the decimal place.
//
// If 'newPrecision' is less than the current BigIntNum precision
// value, the current BigIntNum numeric value is rounded to the
// specified 'newPrecision' value.
//
//	 Examples
//	 ========
//
//	   Original     'newPrecision'    Resulting
//		   Value       input parameter     Value
//
//	   654.123456         9            654.123456000
//	   654.123456         4            654.1235
//
//	  -654.123456         9           -654.123456000
//	  -654.123456         4           -654.1235
//
//	     0                3              0.000
//	     0.000000         0              0
//
//	NOTE
//	====
//
// Existing numeric separators (decimal separator, thousands
// separator and currency symbol) remain unchanged and are not
// altered by this method.
func (bNum *BigIntNum) SetPrecision(newPrecision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetINumMgr",
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

	if newPrecision == bNum.precision {
		return nil
	}

	if bNum.precision > newPrecision {

		return new(bigIntNumBoson).roundToDecimalPlace(
			bNum,
			newPrecision,
			ePrefix.XCpy(fmt.Sprintf("Rounding 'bNum' to newPrecision= '%v'",
				newPrecision)))
	}

	deltaPrecision := newPrecision - bNum.precision

	// bNum.precision must be less than newPrecision

	return new(bigIntNumProton).bigIntNumExtendPrecision(
		bNum,
		deltaPrecision,
		ePrefix.XCpy(
			fmt.Sprintf("bNum.precision < newPrecision; deltaPrecision= '%v'",
				deltaPrecision)))
}

// SetNumericSeparators
//
// Used to assign values for the Decimal and Thousands separators as well
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
// USA Examples
// ============
//
// Decimal Separator period ('.')     = 123.456
// Thousands Separator comma (',')    = 1,000,000,000
// Currency Symbol dollar sign ('$')  = $123
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
//
// NumericSeparatorDto
//
//	type NumericSeparatorDto struct {
//	 DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
//	 ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
//	 CurrencySymbol     rune // Currency Symbol
//	}
func (bNum *BigIntNum) SetNumericSeparatorsDto(
	customSeparators NumericSeparatorDto) error {

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
		ePrefix.XCpy("Setting 'bNum' default NumSeps"))
}

// SetNumericSeparatorsToUSADefault
//
// Sets the following Numeric separators to the United States of
// America (USA) defaults.:
//
//	Decimal Point Separator = '.' (period)
//	Thousands Separator     = ',' (comma)
//	Currency Symbol         = '$' (dollar sign)
//
// Call specific methods to set numeric separators for other countries
// or cultures:
//
//	new(BigIntNum).SetNumericSeparators
//	new(BigIntNum).SetNumericSeparators
//	new(BigIntNum).SetDecimalSeparator()
//	new(BigIntNum).SetThousandsSeparator()
//	new(BigIntNum).SetCurrencySymbol()
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
		ePrefix.XCpy("Setting 'bNum' to USA NumSeps"))
}

// SetNumStr
//
//	Initializes the current BigIntNum instance for the numeric value
//	of the number string input parameter 'numStr'.
//
//	A number string is a string of numeric digits. If the number
//	string is prefixed with a minus sign ('-') or surrounded in
//	parentheses, it is assumed to be a negative value. Otherwise,
//	the numeric value is assumed to be positive. Currency symbols
//	are ignored.
//
//	The numeric string of digits may also contain a decimal
//	separator defined by input parameter 'decimalSeparator'.
//	The decimal separator is used to separate integer and
//	fractional numeric digits within the number string.
//
//	Input parameter numStrNumSeps is an instance of
//	NumericSeparatorDto containing character symbols for
//	Decimal separators, Thousands separtors and the Currency
//	Symbol used in parsing the number string ('numStr').
//
//	Upon completion this method will configure the current
//	instance of BigIntNum with the numeric value represented
//	by input parameter 'numStr'.
//
//	The previously configured Numeric Separators for the current
//	BigIntNum instance will remain unchanged.
func (bNum *BigIntNum) SetNumStr(
	numStr string,
	numStrNumSeps NumericSeparatorDto) error {

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

	bNumNumSeps := NumericSeparatorDto{
		ThousandsSeparator: bNum.thousandsSeparator,
		DecimalSeparator:   bNum.decimalSeparator,
		CurrencySymbol:     bNum.currencySymbol,
	}

	return new(bigIntNumMolecule).setNumStr(
		bNum,
		numStr,
		&numStrNumSeps,
		&bNumNumSeps,
		ePrefix)

}

// SetNumStrDto
//
// This method configures the current instance of BigIntNum with
// the numeric value and numeric separators passed by input
// parameter, 'numStrDto'.
func (bNum *BigIntNum) SetNumStrDto(numStrDto NumStrDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.NewNumStrDto()",
		"")

	if err != nil {
		return err
	}

	return new(bigIntNumMolecule).setNumStrDto(
		bNum,
		numStrDto,
		ePrefix.XCpy("Setting 'bNum' == NumStrDto"))
}

// SetSignValue - Sets the sign value of the current BigIntNum
// to positive (+1) or negative (-1).
//
// If a value other than +1 or -1 is transmitted by input
// parameter 'signVal', an error will be returned.
func (bNum *BigIntNum) SetSignValue(signVal int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetSignValue()",
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

	if signVal == 1 || signVal == -1 {

		if bNum.sign == signVal {

			return nil
		}

		return new(bigIntNumUtility).bigIntNumChangeSign(
			bNum,
			ePrefix)
	}

	return &FuncReturnError{
		ErrPrefix:  ePrefix.String(),
		ReturnFunc: "",
		ErrContext: "Error: Input parameter 'signVal' must be +1 or -1.\n" +
			"signVal='%v'\nInput parameter 'signVal' is INVALID!",
		ErrMessage: "",
	}
}

// ShiftPrecisionLeft
//
// Shifts precision of the current BigIntNum numeric value to the
// left by 'shiftLeftPlaces' decimal places. This is a 'relative'
// shift-left operation. The shift left operation is therefore
// performed with the current decimal place position as the starting
// point.
//
// This operation is equivalent to:
//
//	result = Decimal value / 10^shiftLeftPlaces
//	                 or
//	signed number divided by 10 raised to the
//	power of shiftLeftPlaces.
//
// This method performs a relative shift left of the decimal place
// position. Be careful, this is NOT Shift Number Left operation.
// Instead, this is a Shift Precision Left operation which means
// that the decimal place will be shifted left.
//
// See Examples below.
//
//	 Input Parameters
//	 ================
//
//			shiftLeftPlaces int	- The number of positions the decimal place will be
//														shifted left from its current position.
//
//	 Examples
//	 ========
//
//			shift-left
//
//		   signed Number   places      Result
//
//		    "123456.789"      3      "123.456789"
//		    "123456.789"      2      "1234.56789"
//		    "123456.789"      6      "0.123456789"
//		    "123456789"       6      "123.456789"
//		    "123"             5      "0.00123"
//		    "0"               3      "0"
//		    "123456.789"      0      "123456.789" - zero has no effect on original number string
//
//		   "-123456.789"      0     "-123456.789"
//		   "-123456.789"      3     "-123.456789"
//		   "-123456789"       6     "-123.456789"
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNum *BigIntNum) ShiftPrecisionLeft(shiftLeftPlaces uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.ShiftPrecisionLeft()",
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

	return new(bigIntNumBoson).shiftPrecisionLeft(
		bNum,
		shiftLeftPlaces,
		ePrefix.XCpy(fmt.Sprintf("Shift 'bNum' Left. shiftLeftPlaces= '%v'",
			shiftLeftPlaces)))
}

// ShiftPrecisionRight
//
// Shifts precision of the current BigIntNum numeric value to
// the right by 'shiftRightPlaces' decimal places. This is a
// 'relative' shift-right operation. The shift right operation
// is therefore performed with the current decimal place
// position as the starting point.
//
// This is equivalent to:
//
//	result = Decimal value X 10^shiftRightPrecision
//	                      or
//	Decimal numeric value multiplied by 10 raised to
//	the power of shiftRightPrecision.
//
// This method performs a relative shift right of the decimal place position.
// Be careful, this is NOT a Shift Number Right operation. This is Shift Precision
// Right which means that the decimal place will be shifted right.
//
// See Examples below.
//
//	 Input Parameters
//	 ================
//
//		shiftRightPlaces int	- The number of positions the decimal place will be
//														shifted right from its current position.
//
//	 Examples
//	 ========
//
//		shift-right
//
//	  signed Number     places    Result
//
//	  "123456.789"        3       "123456789"
//	  "123456.789"        2       "12345678.9"
//	  "123456.7896"       7       "1234567896000"
//	  "123456789"         6       "123456789000000"
//	  "123"               5       "12300000"
//	  "0"                 3       "0"
//	  "123456.789"        0       "123456.789" - zero has no effect on original number string
//
//	 "-123456.789         0      "-123456.789"
//	 "-123456.789         3      "-123456789"
//	 "-123456789"         6      "-123456789000000"
//
//	 Numeric Separators
//	 ==================
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) for the current instance of BigIntNum will
// remain unchanged and are not altered by this method.
func (bNum *BigIntNum) ShiftPrecisionRight(shiftRightPlaces uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.ShiftPrecisionRight()",
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

	return new(bigIntNumBoson).shiftPrecisionRight(
		bNum,
		shiftRightPlaces,
		ePrefix.XCpy(fmt.Sprintf("Shift 'bNum' right. shiftRightPlaces= '%v'",
			shiftRightPlaces)))
}

// SetThousandsSeparator
//
// Configures the current instance of BigIntNum and sets the value
// of the character which will be used to separate thousands in the
// display of number strings.
//
// In the USA, the Thousands Separator is the comma character (',').
// Example: 1,000,000
//
// If a zero value is submitted for the input parameter
// 'thousandsSeparator', an error will be returned.
func (bNum *BigIntNum) SetThousandsSeparator(
	thousandsSeparator rune) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.SetThousandsSeparator",
		"")

	if err != nil {
		return err
	}

	var nSepSymbolCode NumSepSymbolCode

	nSepSymbolCode = THOUSANDSYMBOL

	return new(bigIntNumBoson).setNumSepSymbol(
		bNum,
		nSepSymbolCode,
		thousandsSeparator,
		ePrefix.XCpy("Setting 'bNum' Thousands Separator"))
}

// String
//
//		Implements the fmt.Stringer interface.
//
//	 The returned string is a signed floating point
//	 number string representing the true numeric value
//	 encapsulated by the current instance of BigIntNum.
//
//	 As such, it is identical to the string returned by
//	 method: BigIntNum.GetNumStr()
//
//	 BigIntNum.GetNumStr() is the preferred means of
//	 returned a signed number string because that method
//	 also returns an error value.
//
//	 If an error is encountered by this method, BigIntNum.String(),
//	 the returned string will be set equal to "ERROR".
func (bNum *BigIntNum) String() string {

	var bNumNumStr string

	bNumNumStr, err := bNum.GetNumStr()

	if err != nil {
		bNumNumStr = "ERROR"
	}

	return bNumNumStr
}

// TrimTrailingFracZeros
//
// This method will delete non-significant trailing zeros from
// the fractional digits of the current BigIntNum numerical
// value.
//
//	Examples
//	========
//
//	  Initial Value   Trimmed Value
//
//	  456.123000         456.123
//	    0.000              0
//	    7.0                7
//	 -456.123000        -456.123
func (bNum *BigIntNum) TrimTrailingFracZeros() error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.TrimTrailingFracZeros()",
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

	return new(bigIntNumBoson).trimTrailingFracZeros(
		bNum,
		ePrefix.XCpy("Trimming 'bNum' Trailing Frac Zeros"))
}

// TruncToDecPlace
//
// Truncates the current BigIntNum to the number of decimal places
// specified by input parameter 'precision'. No rounding occurs.
// The trailing digits are simply truncated or deleted in order
// to achieve the specified number of decimal places.
//
// 'precision' equals the number of digits to the right of the
// decimal place.
//
// If the value of BigIntNum.bigInt is zero ('0'), that zero value
// will remain unaltered. However, BigIntNum.precision will be set
// equal to input parameter, 'precision'.
//
// If the number of decimal places specified for truncation
// ('precision') is equal to the current BigIntNum.precision, no
// action is taken and the original BigIntNum numeric value
// remains unchanged.
//
// If the number of decimal places specified for truncation
// ('precision') is greater than the current BigIntNum.precision,
// trailing zeros are added to the current BigIntNum.bigInt value
// and BigIntNum.precision is set equal to input parameter,
// 'precision'.
//
// If 'precision' is less than the current BigIntNum.precision
// value, the current BigIntNum numeric value is truncated to
// the specified 'precision' value and NO rounding occurs.
//
//	Examples
//	========
//
//	   Original           'newPrecision'        Resulting
//	    Value             input parameter         Value
//
//	  654.123456                9              654.123456000
//	  654.123456                4              654.1234 (no rounding)
//
//	 -654.123456                9             -654.123456000
//	 -654.123456                4             -654.1234 (no rounding)
//
//	    0                       3                0.000
//	    0.000000                0                0
//
//	Numeric Separators
//	==================
//
// Existing numeric separators (decimal separator, thousands
// separator and currency symbol) in the current instance of
// BigIntNum will remain unchanged and are not altered by this
// method.
func (bNum *BigIntNum) TruncToDecPlace(precision uint) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntNum.TruncToDecPlace",
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

	return new(bigIntNumBoson).truncToDecPlace(
		bNum,
		precision,
		ePrefix.XCpy(fmt.Sprintf("Trunc 'bNum' Dec Places. precision= '%v'",
			precision)))
}
