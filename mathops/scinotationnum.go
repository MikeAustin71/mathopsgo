package mathops

import (
	"fmt"

	ePref "github.com/MikeAustin71/errpref"
)

// SciNotationNum
//
//	Example Scientific Notation
//
//	        2.652e+8
//
//	significand     = '2.652'
//	significand integer digits = '2'
//	mantissa        = significand factional digits = '.652'
//	exponent        = '8' (10^8)
//	mantissaLength  = length of fractional digits displayed in scientific notation.
type SciNotationNum struct {
	significand BigIntNum // The significand consists of the leading integer and
	//	fractional digits of the scientific notation.
	exponent         BigIntNum // The exponent portion of the scientific notation string
	exponentChar     rune      // 	defaults to 'e'. May be customized to 'E'
	decimalSeparator rune      // The decimal separator used to separate integer and
	// 	fractional digits in the significand. The default is
	// 	the standard USA decimal separator, the decimal point ('.').
	mantissaLength uint // The length of the fractional digits in
	// 	the significand which will be displayed
	// 	when SciNotationNum.GetSciNotationStr()
	// 	is called.
	exponentUsesLeadingPlus bool // If true, positive exponent values are
	// 	prefixed with a leading plus (+) sign.
	//  '2.652e+8'
}

// GetDecimalSeparator
//
//	Returns the current decimal separator character as a string.
//
//	In a default scientific notation display, '2.652e+8', the decimal
//	separator character is presented as a period ('.') separating
//	integer and fractional digits in the significand ('2.652').
//	However, the user has the option to customize this decimal
//	separator character through method:
//	  SciNotationNum.SetDecimalSeparatorChar().
func (sciNotan *SciNotationNum) GetDecimalSeparator() rune {

	new(scinotationnumElectron).setDecimalSeparatorIfEmpty(sciNotan)

	return sciNotan.decimalSeparator
}

// GetExponent
//
//	Returns the exponent element of the scientific notation as a
//	BigIntNum.
//
//	Example Scientific Notation
//	===========================
//
//	  2.652e+8
//
//	significand     = '2.652'
//	significand integer digits = '2'
//	mantissa        = significand factional digits = '.652'
//	exponent        = '8'  (10^8)
//	mantissaLength  = length of fractional digits displayed in scientific notation.
func (sciNotan *SciNotationNum) GetExponent() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"SciNotationNum.GetExponent",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	sciNotExponent, err := sciNotan.exponent.CopyOut()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "sciNotExponent, err := sciNotan.exponent.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return sciNotExponent, nil
}

// GetExponentChar
//
//	Returns the current exponent char (character) as a string.
//
//	In a default scientific notation display, '2.652e+8', the exponent
//	char is presented as an 'e'. However, the user has the option to
//	customize this exponent character through method:
//	  SciNotationNum.SetExponentChar().
//
//	This method returns the current exponent character which will be
//	used in formatting scientific notation strings.
func (sciNotan *SciNotationNum) GetExponentChar() string {

	new(scinotationnumElectron).setExponentCharIfEmpty(sciNotan)

	return string(sciNotan.exponentChar)
}

// GetExponentUsesLeadingPlus
//
//	Returns the internal 'exponentUsesLeadingPlus' flag. By default,
//	this is set to 'true'.  When set to 'true' the scientific notation
//	is displayed with positive exponents prefixed with a plus sign (+).
//
//	Default Example
//	---------------
//	2.652e+8
func (sciNotan *SciNotationNum) GetExponentUsesLeadingPlus() bool {

	return sciNotan.exponentUsesLeadingPlus
}

// GetNumStr - Returns a scientific notation string representing
// the underlying numeric value. The number of decimals in the
// mantissa will default to the current value of sciNotan.mantissaLength.
func (sciNotan *SciNotationNum) GetNumStr() (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"SciNotationNum.GetNumStr",
		"")

	if err != nil {
		return "", err
	}

	new(scinotationnumElectron).setMantissaLengthIfEmpty(sciNotan)

	mantissaLength := sciNotan.mantissaLength

	return new(scinotationnumNanobot).getSciNotationStr(
		sciNotan, mantissaLength, ePrefix)
}

// GetSciNotationStr
//
//	Returns a string containing the scientific notation display.  This
//	method differs from method SciNotationNum.GetNumStr() in that this
//	method requires the user to provide input parameter 'mantissaLen'
//	which sets the number of decimal places in the significand of the
//	returned science notation string.
//
//	Default Example
//	---------------
//	"2.652e+8"
func (sciNotan *SciNotationNum) GetSciNotationStr(mantissaLen uint) (string, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"SciNotationNum.GetSciNotationStr",
		"")

	if err != nil {
		return "", err
	}

	return new(scinotationnumNanobot).getSciNotationStr(
		sciNotan, mantissaLen, ePrefix)
}

// GetSignificand
//
//	Returns the significand component of the scientific notation as a
//	BigIntNum.
//
//	Example Scientific Notation
//	===========================
//
//	2.652e+8
//
//	significand     = '2.652'
//	significand integer digits = '2'
//	mantissa        = significand factional digits = '.652'
//	exponent        = '8'  (10^8)
//	mantissaLength  = length of fractional digits displayed in
//	                  scientific notation.
func (sciNotan *SciNotationNum) GetSignificand() (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"NumStrDto.CopyIn",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	significand2, err := sciNotan.significand.CopyOut()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "significand2, err := sciNotan.significand.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return significand2, nil
}

// New
//
//	Creates and returns an empty SciNotationNum structure. It is a
//	good idea to call this method in order to initialize default
//	settings.
func (sciNotan *SciNotationNum) New() SciNotationNum {

	return new(scinotationnumNanobot).new()
}

// NewNumStr - Creates and returns a new SciNotationNum instance initialized from
// an input string.
//
// Input parameter 'sciNotationStr' should be properly formatted as a valid scientific
// notation string. Invalid input strings will trigger an error.
func (sciNotan *SciNotationNum) NewNumStr(sciNotationStr string) (SciNotationNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"SciNotationNum.NewNumStr",
		"")

	if err != nil {
		return SciNotationNum{}, err
	}

	s2 := new(scinotationnumNanobot).new()

	err = new(scinotationnumMolecule).setNumStr(&s2, sciNotationStr, ePrefix)

	if err != nil {

		return SciNotationNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(scinotationnumMolecule).\n" +
					"setNumStr(&s2, sciNotationStr, ePrefix)",
				ErrContext: fmt.Sprintf("sciNotationStr= '%v", sciNotationStr),
				ErrMessage: err.Error(),
			}
	}

	return s2, nil
}

// SetDecimalSeparatorIfEmpty
//
//	Sets the default decimal separator if the
//	SciNotationNum.decimalSeparator rune is empty or equal to zero.
//
//	In the example scientific notation '2.652e+8', the decimal
//	separator is the period or decimal point ('.'). If the
//	SciNotationNum.decimalSeparator rune is empty, this method will
//	set the decimal separator to the USA default decimal separator,
//	the period or decimal point ('.').
//
//	The decimal separator character may be customized to support characters
//	used by other cultures or nations.
//
//	See method SciNotationNum.SetDecimalSeparatorChar() below.
func (sciNotan *SciNotationNum) SetDecimalSeparatorIfEmpty() {

	new(scinotationnumElectron).setDecimalSeparatorIfEmpty(
		sciNotan)

	return
}

// SetDecimalSeparatorChar
//
//	Sets the value of the Scientific Notation decimal separator used
//	to separate integer and fractional digits in the significand.
//
//	In a default scientific notation display, '2.652e+8', the decimal
//	separator character is presented as a period ('.') separating
//	integer and fractional digits in the significand ('2.652').
//	However, the user has the option to customize this decimal
//	separator character by calling this method.
func (sciNotan *SciNotationNum) SetDecimalSeparatorChar(decimalChar rune) {

	new(scinotationnumElectron).setDecimalSeparatorChar(sciNotan, decimalChar)

	return
}

// SetExponentCharIfEmpty
//
//	If the current instance of SciNotationNum has an empty (0 value)
//	exponent rune, this method will set the exponentChar value to 'e'.
func (sciNotan *SciNotationNum) SetExponentCharIfEmpty() {

	new(scinotationnumElectron).setExponentCharIfEmpty(sciNotan)

	return
}

// SetExponentChar
//
//	The exponent char used in displaying scientific notation may be
//	customized through this method.
//
//	By default, scientific notation is displayed in the following
//	format using the 'e' character as the exponentChar: 2.652e+8
//
//	The user may find it preferable to display the scientific notation
//	using capital 'E': 2.652E+8
//
//	This method is used to customize the exponent character used when
//	calling SciNotationNum.GetSciNotationStr()
func (sciNotan *SciNotationNum) SetExponentChar(exponentChar rune) {

	sciNotan.exponentChar = exponentChar
}

// SetExponentUsesLeadingPlus
//
//	When the input parameter 'useLeadingPlus' is set to true, a plus
//	sign will be displayed in the scientific notation output string.
//	  Example: 2.652e+8
//
//	If 'useLeadingPlus' is set to false, the plus sign will NOT be
//	displayed in front of the exponent.
//	  Example: 2.652e8
func (sciNotan *SciNotationNum) SetExponentUsesLeadingPlus(useLeadingPlus bool) {

	sciNotan.exponentUsesLeadingPlus = useLeadingPlus
}

// SetMantissaLength
//
//	This method sets the length of the mantissa or fractional digits
//	which will be displayed in the significand when
//	SciNotationNum.GetSciNotationStr() is called.
//
//	If input parameter mantissaLen is set equal to zero, this method
//	will automatically set the value to two ('2').
//
//	Example Scientific Notation
//	===========================
//
//	2.652e+8
//
//	significand = '2.652'
//	significand integer digit = '2'
//	mantissa		= significand factional digits = '.652'
//	exponent    = '8'  (10^8)
func (sciNotan *SciNotationNum) SetMantissaLength(mantissaLen uint) {

	new(scinotationnumElectron).setMantissaLength(sciNotan, mantissaLen)

	return
}

// SetMantissaLengthIfEmpty
//
//	If mantissa length is zero, this method attempts to set mantissa
//	length equal to the precision of 'significand'.
//
//	Mantissa is defined as the length or number of fractional digits
//	which will be displayed in the significand when
//	SciNotationNum.GetSciNotationStr() is called to produce scientific
//	notation formatted as a string.
//
//	In the example scientific notation '2.652e+8', the mantissa is
//	'.652'
func (sciNotan *SciNotationNum) SetMantissaLengthIfEmpty() {

	new(scinotationnumElectron).setMantissaLengthIfEmpty(sciNotan)

	return
}

// SetBigIntNumElements
//
//	Sets the components of the current SciNotationNum instance based
//	on two BigIntNum input parameters.
//
//	Input Parameters
//	================
//
//	significand              BigIntNum
//	  In the example scientific notation '2.652e+8',	the significand
//	  is represented by '2.652'.
//
//	exponent                 BigIntNum
//	  In the example '2.652e+8' the exponent component is represented
//	  by the integer value, '8'. Note: If exponent is NOT an integer
//	  value and contains fractional digits, an error will be
//	  triggered.
func (sciNotan *SciNotationNum) SetBigIntNumElements(
	significand, exponent BigIntNum) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"SciNotationNum.SetBigIntNumElements",
		"")

	if err != nil {
		return err
	}

	return new(scinotationnumNanobot).setBigIntNumElements(
		sciNotan, &significand, true, &exponent, true, ePrefix)
}

// SetIntAryElements
//
//	Sets the components of the current SciNotationNum instance based
//	on two IntAry input parameters.
//
//	Input Parameters
//	================
//
//	significand              IntAry
//	  In the example scientific notation '2.652e+8', the significand
//	  is represented by '2.652'.
//
//	exponent                 IntAry
//	  In the example '2.652e+8' the exponent component is represented
//	  by the integer value, '8'. Note: If exponent is NOT an integer
//	  value and contains fractional digits, an error will be
//	  returned.
func (sciNotan *SciNotationNum) SetIntAryElements(
	significand, exponent IntAry) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"SciNotationNum.SetIntAryElements",
		"")

	if err != nil {
		return err
	}

	return new(scinotationnumNanobot).setIntAryElements(
		sciNotan, &significand, true, &exponent, true, ePrefix)
}

// SetNumStr
//
//	Receives a properly formatted scientific notation string as an
//	input parameter and converts to the data fields of the current
//	SciNotationNum instance.
//
//	Examples of properly formatted scientific notation strings are
//	provided below:
//
//	  2.652e+8
//	  2.652E+8
//	  2.652e8
//	  2.652E8
//
//	The use of the decimal point may be customized by first setting
//	the desired decimal separator using method,
//	SciNotationNum.SetDecimalSeparatorChar().
//
//	Also, note that exponent digits must be integer numbers. Used of fractional
//	digits in the exponent will trigger an error. Example:
//
//	2.652E9.24 = ERROR fractional digits in exponent!
func (sciNotan *SciNotationNum) SetNumStr(sciNotationStr string) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"SciNotationNum.SetNumStr",
		"")

	if err != nil {
		return err
	}

	return new(scinotationnumMolecule).setNumStr(
		sciNotan, sciNotationStr, ePrefix)
}
