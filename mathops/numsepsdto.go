package mathops

import (
  "fmt"

  ePref "github.com/MikeAustin71/errpref"
)

type NumSepsProfileSelection struct {
  SourceObjectName         string
  OutputNumSepsName        string
  UseDefaultNumSeps        bool
  SetDefaultNumSepsIfEmpty bool
  ValidateNumSeps          bool
  OverrideNumSeps          NumericSeparatorDto
}

// NumericSeparatorDto
// Numeric Separator Data Transfer Object
// Used to transmit symbols used for decimal point,
// thousands separator and currency symbol.
type NumericSeparatorDto struct {
  DecimalSeparator   rune // Character used to separate integer and fractional digits ('.')
  ThousandsSeparator rune // Character used to separate thousands (1,000,000,000
  CurrencySymbol     rune // Currency Symbol
}

// CopyIn
//
// Receives an external instance of NumericSeparatorDto and
// copies the member variables into the current instance of
// NumericSeparatorDto
//
//	Input Parameters
//	================
//
//	numSepsDtoSrc            *NumericSeparatorDto
//	  An external instance of NumericSeparatorDto which
//	  serves as the source for member element data copied
//	  from to the current instance of NumericSeparatorDto.
//	  Upon successful completion of this copy operation,
//	  the current instance of NumericSeparatorDto will
//	  constitute an identical copy of 'numSepsDtoSrc'.
//
//	setDefaultsIfEmpty       bool
//	  If this boolean parameter is set to 'true' any
//	  invalid values copied to the current instance of
//	  NumericSeparatorDto will be automatically reset
//	  to valid USA defaults.
//
//	  If this setDefaultsIfEmpty is set to 'false', any
//	  invalid values identified in 'numSepsDtoSrc' will
//	  trigger an error return.
func (numSep *NumericSeparatorDto) CopyIn(
  numSepsDtoSrc *NumericSeparatorDto,
  setDefaultsIfEmpty bool) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumericSeparatorDto.CopyIn",
    "")

  if err != nil {
    return err
  }

  return new(numSepsDtoMechanics).copyNumSepsDto(
    numSep, // Destination
    numSepsDtoSrc,
    setDefaultsIfEmpty,
    ePrefix.XCpy("Copy 'numSepsDtoSrc' Into 'numSep'"))
}

// CopyOut
//
// Returns a deep copy of the current NumericSeparatorDto instance.
//
//	Input Parameters
//	================
//
//	setDefaultsIfEmpty                bool
//	  If this boolean parameter is set to 'true' any existging
//	  invalid values in the returned NumericSeparatorDto copy
//	  will be automatically reset to valid USA defaults.
//
//	  If this setDefaultsIfEmpty is set to 'false', any invalid
//	  values in the returned NumericSeparatorDto copy will
//	  trigger an error return.
func (numSep *NumericSeparatorDto) CopyOut(setDefaultsIfEmpty bool) (NumericSeparatorDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumericSeparatorDto.CopyOut",
    "")

  if err != nil {
    return NumericSeparatorDto{}, err
  }

  var numSepsDest = new(NumericSeparatorDto)

  err = new(numSepsDtoMechanics).copyNumSepsDto(
    numSepsDest,
    numSep,
    setDefaultsIfEmpty,
    ePrefix.XCpy("Copy current 'numSep' into 'numSepsDest'"))

  return *numSepsDest, err
}

// Empty
//
//	Resets all the member elements of the NumericSeparatorDto
//	struct to their initital or 'zero' values.
func (numSep *NumericSeparatorDto) Empty() {

  numSep.DecimalSeparator = 0
  numSep.ThousandsSeparator = 0
  numSep.CurrencySymbol = 0

  return
}

// Equal - Compares two NumericSeparatorDto's and returns 'true' if they
// are equivalent.
func (numSep *NumericSeparatorDto) Equal(numSep2 NumericSeparatorDto) bool {

  if numSep.DecimalSeparator != numSep2.DecimalSeparator {
    return false
  }

  if numSep.ThousandsSeparator != numSep2.ThousandsSeparator {
    return false
  }

  if numSep.CurrencySymbol != numSep2.CurrencySymbol {
    return false
  }

  return true
}

// GetInputSeparators
//
//	Implements the IGetNumSeparators interface. This method is
//	called when the encapsulated Numeric Separators for the
//	current instance of NumericSeparatorDto are used as input
//	Numeric Separators. Input Numeric Separators are primarily
//	used to parse number strings. The parsing number strings
//	operation is used to convert number strings to numeric
//	values.
//
//	Input Parameters
//	================
//
//	None
//
//	Output Parameters
//	=================
//
//	*NumericSeparatorDto
//	  This returned instance of NumericSeparatorDto is designed to
//	  used as input Numeric Separators when parsing number strings
//	  and converting those strings to a numeric value.
//
//	error
//	  If no errors are encountered, this returned error parameter
//	  will be set to 'nil'.
func (numSep *NumericSeparatorDto) GetInputSeparators() (*NumericSeparatorDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumericSeparatorDto.GetInputSeparators",
    "")

  if err != nil {
    return &NumericSeparatorDto{}, err
  }

  err = new(numSepsDtoElectron).isValidNumStrDto(
    numSep,
    ePrefix.XCpy("Validating current 'numSep' instance"))

  if err != nil {

    return &NumericSeparatorDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "The current instance of NumericSeparatorDto is invalid!",
        ErrMessage: err.Error(),
      }
  }

  inputNumSeps := &NumericSeparatorDto{}

  inputNumSeps.CurrencySymbol = numSep.CurrencySymbol
  inputNumSeps.DecimalSeparator = numSep.DecimalSeparator
  inputNumSeps.ThousandsSeparator = numSep.ThousandsSeparator

  return inputNumSeps, nil
}

// GetOutputSeparators
//
//	Implements the IGetNumSeparators interface. This method is
//	called when the encapsulated Numeric Separators for the
//	current instance of NumericSeparatorDto are used as output
//	Numeric Separators. Output Numeric Separators are primarily
//	used format number types returned from functions. When these
//	number types are later converted to number strings for display
//	purposes, the output formatting for decimal separators,
//	thousands separators and currency symbols will be controlled
//	by these output Numeric Separators.
//
//	Input Parameters
//	================
//
//	None
//
//	Output Parameters
//	=================
//
//	*NumericSeparatorDto
//	  This returned instance of NumericSeparatorDto is designed to
//	  used in formatting numeric types returned by other methods.
//
//	error
//	  If no errors are encountered, this returned error parameter
//	  will be set to 'nil'.
func (numSep *NumericSeparatorDto) GetOutputSeparators() (*NumericSeparatorDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumericSeparatorDto.GetOutputSeparators",
    "")

  if err != nil {
    return &NumericSeparatorDto{}, err
  }

  err = new(numSepsDtoElectron).isValidNumStrDto(
    numSep,
    ePrefix.XCpy("Validating current 'numSep' instance"))

  if err != nil {

    return &NumericSeparatorDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "The current instance of NumericSeparatorDto is invalid!",
        ErrMessage: err.Error(),
      }
  }

  outputNumSeps := &NumericSeparatorDto{}

  outputNumSeps.CurrencySymbol = numSep.CurrencySymbol
  outputNumSeps.DecimalSeparator = numSep.DecimalSeparator
  outputNumSeps.ThousandsSeparator = numSep.ThousandsSeparator

  return outputNumSeps, nil
}

// IsValid
//
// This method will test the current instance of
// NumericSeparatorDto to determine if all member variables are
// valid. If any member elements are determined to be invalid,
// an error will be returned.
//
// If all member elements of the current NumericSeparatorDto
// instance are determined to be valid, the method will return
// a 'nil' value.
func (numSep *NumericSeparatorDto) IsValid(callingMethodName string) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  if len(callingMethodName) > 0 {
    callingMethodName = "NumericSeparatorDto.IsValid" + "\n" + callingMethodName
  } else {
    callingMethodName = "NumericSeparatorDto.IsValid"
  }

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    callingMethodName,
    "")

  if err != nil {
    return err
  }

  return new(numSepsDtoElectron).isValidNumStrDto(
    numSep,
    ePrefix)
}

// IsValidOrSetDefaults
//
//  This method will test the current instance of
//  NumericSeparatorDto to determine if all member variables are
//  valid.
//
//  If any member elements are determined to be invalid,
//  this method will populate the current instance of
//  NumericSeparatorDto with default USA Numeric Separators.
//
//  USA default values are listed as follows:
//
//    Decimal Separator   = '.' (period)
//    Thousands Separator = ',' (comma)
//    Currency Symbol     = '$' (dollar sign)

func (numSep *NumericSeparatorDto) IsValidOrSetDefaults(callingMethodName string) (wasRestToDefault bool, err error) {
  var ePrefix *ePref.ErrPrefixDto

  if len(callingMethodName) > 0 {
    callingMethodName = "NumericSeparatorDto.IsValid" + "\n" + callingMethodName
  } else {
    callingMethodName = "NumericSeparatorDto.IsValid"
  }

  wasRestToDefault = false

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    callingMethodName,
    "")

  if err != nil {
    return wasRestToDefault, err
  }

  err = new(numSepsDtoElectron).isValidNumStrDto(
    numSep,
    ePrefix)

  if err != nil {

    new(numSepsDtoElectron).setNumSepDtoDefaultsIfEmpty(
      numSep)

    wasRestToDefault = true
  }

  return wasRestToDefault, err
}

// New - Returns a new instance of NumericSeparatorDto. The
// rune values are automatically set to USA defaults.
func (numSep *NumericSeparatorDto) New() NumericSeparatorDto {

  n2 := NumericSeparatorDto{}

  new(numSepsDtoElectron).setNumSepDtoDefaultsIfEmpty(
    &n2)

  return n2
}

// NewUSADefaults
//
//	Returns a new instance of NumericSeparatorDto. The rune values
//	are automatically set to USA defaults.
func (numSep *NumericSeparatorDto) NewUSADefaults() NumericSeparatorDto {

  return new(numSepsDtoElectron).newUSADefaults()
}

// SetDefaultsIfEmpty - If any of the NumericSeparatorDTo rune values
// are zero, this method will set those elements to USA default values.
func (numSep *NumericSeparatorDto) SetDefaultsIfEmpty() {

  new(numSepsDtoElectron).setNumSepDtoDefaultsIfEmpty(
    numSep)
}

// SetUSADefaults
//
//	This method will arbitrarily set all member variables of
//	the current NumericSeparatorDto instance to USA defaults.
//	USA default values are listed as follows:
//
//	  Decimal Separator   = '.' (period)
//	  Thousands Separator = ',' (comma)
//	  Currency Symbol     = '$' (dollar sign)
func (numSep *NumericSeparatorDto) SetUSADefaults() {

  new(numSepsDtoElectron).setNumSepDtoDefaultsIfEmpty(
    numSep)
}

// String - Provides a formatted listing of the contents from the current
// NumericSeparatorDto instance.
func (numSep *NumericSeparatorDto) String() string {
  return fmt.Sprintf("Decimal Separator: %q  Thousands Separator: %q  Currency Symbol: %q",
    numSep.DecimalSeparator, numSep.ThousandsSeparator, numSep.CurrencySymbol)
}
