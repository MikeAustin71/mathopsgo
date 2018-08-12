package mathops

// Example Scientific Notation
// ===========================
//
//  2.652e+8
//
//  significand 		= '2.652'
//	mantissa				= factional digits = '.652'
//  exponent    		= '8'  (10^8)
//	mantissaLength	= length of fractional digits displayed in scientific notation.
//
type SciNotationNum struct {

		significand BigIntNum		// The significand consists of the leading integer and
														//	fractional digits of the scientific notation.
		exponent BigIntNum			// The exponent portion of the scientific notation string
		exponentChar rune				// defaults to 'e'. May be customized to 'E'
		mantissaLength uint			// The length of the fractional digits in
														// 	the significand which will be displayed
														// 	when SciNotationNum.GetSciNotationStr()
														// 	is called.
	  exponentUsesLeadingPlus	bool	// If true, positive exponent values are
	  															// 	prefixed with a leading plus (+) sign.
	  															//  '2.652e+8'
}

// New() - Creates and returns an empty SciNotationNum
// structure. It is a good idea to call this method
// in order to initialize default settings.
//
func (sciNotan SciNotationNum) New() SciNotationNum {

	s2 := SciNotationNum{}

	s2.SetExponentCharIfEmpty()

	s2.significand = BigIntNum{}.NewZero(0)
	s2.exponent = BigIntNum{}.NewZero(0)
	s2.exponentUsesLeadingPlus = true

	return s2
}

// GetExponentChar - Returns the current exponent char
// as a string.
//
// In a default scientific notation display, '2.652e+8',
// the exponent char is presented as an 'e'. However, the
// user has the option to customize this exponent character
// through method SciNotationNum.SetExponentChar().
//
// This method returns the current exponent character which
// will be used in formatting scientific notation strings.
//
func (sciNotan *SciNotationNum) GetExponentChar() string {

	sciNotan.SetExponentCharIfEmpty()

	return string(sciNotan.exponentChar)
}

// GetExponentUsesLeadingPlus - returns the internal 'exponentUsesLeadingPlus'
// flag. By default this is set to 'true'.  When set to 'true' the scientific
// notation is displayed with positive exponents prefixed with a plus sign (+).
//
// Default Example
// ---------------
// 2.652e+8
//
func (sciNotan *SciNotationNum) GetExponentUsesLeadingPlus() bool {

	return sciNotan.exponentUsesLeadingPlus
}

// GetSciNotationStr - Returns a string containing the scientific
// notation display.
func (sciNotan *SciNotationNum) GetSciNotationStr() string {

	outStr := ""

	if sciNotan.significand.GetPrecisionUint() != sciNotan.mantissaLength {
		bINum := sciNotan.significand.CopyOut()

		bINum.SetPrecision(sciNotan.mantissaLength)

		outStr += bINum.GetNumStr()

	} else {

		outStr += sciNotan.significand.GetNumStr()
	}


	outStr += sciNotan.GetExponentChar()

	if sciNotan.exponent.GetSign() == 1 &&
		sciNotan.exponentUsesLeadingPlus {
		outStr += "+"
	}

	outStr += sciNotan.exponent.GetNumStr()

	return outStr
}

// NewNumStr - Creates and returns a new SciNotationNum instance initialized from
// an input string.
//
// Input parameter 'sciNotationStr' should be properly formatted as a valid scientific
// notation string. Invalid input strings will trigger an error.
//
func (sciNotan SciNotationNum) NewNumStr(sciNotationStr string) (SciNotationNum, error) {

	return SciNotationNum{}.New(), nil
}


// SetExponentCharIfEmpty - If the exponent rune is empty,
// this method will set the exponentChar value to 'e'.
//
func (sciNotan *SciNotationNum) SetExponentCharIfEmpty() {

	if sciNotan.exponentChar == 0 {
		sciNotan.exponentChar = 'e'
	}

}

// SetExponentChar - The exponent char used in displaying scientific notation
// may be customized through this method.
//
// By default, scientific notation is displayed in the following format using the
// 'e' character as the exponentChar: 2.652e+8
//
// The user may find it preferable to display the scientific notation using capital
// 'E': 2.652E+8
//
// This method is used to customized the exponent character used when calling
// SciNotationNum.GetSciNotationStr()
func (sciNotan *SciNotationNum) SetExponentChar(exponentChar rune) {
	sciNotan.exponentChar = exponentChar
}

func (sciNotan *SciNotationNum) SetExponentUsesLeadingPlus(useLeadingPlus bool) {

	sciNotan.exponentUsesLeadingPlus = useLeadingPlus

}

// SetMantissaLength - This method sets the length of the mantissa or
// fractional digits which will be displayed in the the significand
// when SciNotationNum.GetSciNotationStr() is called.
//
// Example Scientific Notation
// ===========================
//
//  2.652e+8
//
//  significand = '2.652'
//	mantissa		= factional digits = '.652'
//  exponent    = '8'  (10^8)
//
func (sciNotan *SciNotationNum) SetMantissaLength(mantissaLen uint) {

	sciNotan.mantissaLength = mantissaLen

}

func (sciNotan *SciNotationNum) SetNumStr(sciNotationStr string) error {

	return nil
}