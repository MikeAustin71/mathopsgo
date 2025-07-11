package mathops

import "sync"

type scinotationnumElectron struct {
  lock sync.Mutex
}

// SetExponentCharIfEmpty - If the exponent rune is empty,
// this method will set the exponentChar value to 'e'.
func (sciNotElectron *scinotationnumElectron) setExponentCharIfEmpty(
  sciNotan *SciNotationNum) {

  sciNotElectron.lock.Lock()

  defer sciNotElectron.lock.Unlock()

  if sciNotan == nil {
    return
  }

  if sciNotan.exponentChar == 0 {
    sciNotan.exponentChar = 'e'
  }

  return
}

// setDecimalSeparatorChar
//
//	Sets the value of the Scientific Notation decimal separator used
//	to separate integer and fractional digits in the significand.
//
//	In a default scientific notation display, '2.652e+8', the decimal
//	separator character is presented as a period ('.') separating
//	integer and fractional digits in the significand ('2.652').
//	However, the user has the option to customize this decimal
//	separator character by calling this method.
func (sciNotElectron *scinotationnumElectron) setDecimalSeparatorChar(
  sciNotan *SciNotationNum,
  decimalChar rune) {

  sciNotElectron.lock.Lock()

  defer sciNotElectron.lock.Unlock()

  if sciNotan == nil {
    return
  }

  if decimalChar == 0 {
    return
  }

  sciNotan.decimalSeparator = decimalChar

  _ = sciNotan.significand.SetDecimalSeparator(decimalChar)

  return
}

// setDecimalSeparatorIfEmpty
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
func (sciNotElectron *scinotationnumElectron) setDecimalSeparatorIfEmpty(
  sciNotan *SciNotationNum) {

  sciNotElectron.lock.Lock()

  defer sciNotElectron.lock.Unlock()

  if sciNotan == nil {
    return
  }

  if sciNotan.decimalSeparator == 0 {
    sciNotan.decimalSeparator = '.'
    _ = sciNotan.significand.SetDecimalSeparator('.')
  }

  return
}

// setMantissaLength
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
func (sciNotElectron *scinotationnumElectron) setMantissaLength(
  sciNotan *SciNotationNum,
  mantissaLen uint) {

  sciNotElectron.lock.Lock()

  defer sciNotElectron.lock.Unlock()

  if sciNotan == nil {
    return
  }

  if mantissaLen == 0 {
    mantissaLen = 2
  }

  sciNotan.mantissaLength = mantissaLen

  return
}

// setMantissaLengthIfEmpty
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
func (sciNotElectron *scinotationnumElectron) setMantissaLengthIfEmpty(
  sciNotan *SciNotationNum) {

  sciNotElectron.lock.Lock()

  defer sciNotElectron.lock.Unlock()

  if sciNotan == nil {
    return
  }

  var significandPrecisionUint uint
  var err error

  significandPrecisionUint, err = sciNotan.significand.GetPrecisionUint()

  if err != nil {

    significandPrecisionUint = 0
  }

  if sciNotan.mantissaLength == 0 {

    if significandPrecisionUint == 0 {
      sciNotan.mantissaLength = 2

    } else {

      sciNotan.mantissaLength = significandPrecisionUint
    }

  }

  return
}
