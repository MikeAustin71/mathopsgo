package mathops

import (
  "errors"
  "fmt"
  ePref "github.com/MikeAustin71/errpref"
  "math/big"
  "strconv"
)

/*
	The source code repository for numstrdto.go is located at:
			https://github.com/MikeAustin71/mathopsgo.git

	The source file decimal.go is located in directory:
		MikeAustin71/mathopsgo/mathops/numstrdto.go
*/

// NumStrDto - This Type contains data fields and methods used
// to manage, store and transport number strings.
//
// The NumStrDto Type implements the INumMgr interface.
type NumStrDto struct {
  signVal int // An integer value indicating the numeric sign of this number string.
  // 		Valid values are +1 or -1
  absAllNumRunes []rune // An array of runes containing all the numeric digits in a number with
  //		no preceding plus or minus sign character. Example: 123.456 =
  //		[]rune{'1','2','3','4','5','6'}
  precision          uint // The number of digits to the right of the decimal point.
  thousandsSeparator rune // Separates thousands in the integer number: '1,000,000,000'
  decimalSeparator   rune // Separates integer and fractional elements of a number. '123.456'
  currencySymbol     rune // Currency symbol used in currency string displays
}

// Add
//
//	Adds the value of input NumStrDto parameter 'n2Dto' to the
//	value of the current NumStrDto instance.
//
//	  nDto current instance + n2Dto = nDto current instance
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The Numeric Separators originally configured for the current
//	instance of NumStrDto will remain unchanged. No modifications
//	to Numeric Separators will be made by this method.
func (nDto *NumStrDto) Add(n2Dto NumStrDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.Add",
    "")

  if err != nil {
    return err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "Error: The current NumStrDto instance ('nDto') is INVALID!\n" +
        "'nDto' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
        "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(numStrDtoMechanics).addNumStrDto(
    numSeps, nDto, false, &n2Dto,
    true, ePrefix.XCpy("nDto + n2Dto->nDto"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(numStrDtoMechanics).addNumStrDto(\n" +
        "  numSeps, nDto, false, &n2Dto, true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// AddNumStrs
//
//	Adds the values represented by two NumStrDto objects and
//	returns the result as a new instance of NumStrDto.
//
//	 n1Dto + n2Dto = Returned NumStrDto
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators provided by the
//	current instance of 'nDto'. If these Numeric Separators are
//	determined to be invalid, an error will be returned.
func (nDto *NumStrDto) AddNumStrs(n1Dto NumStrDto, n2Dto NumStrDto) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.AddNumStrs",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
        ErrContext: "Error extracting Numeric Separators from current NumStrDto instance.",
        ErrMessage: err.Error(),
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{}, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
      ErrContext: "Error: The current instance of NumStrDto ('nDto') is INVALID!\n" +
        "Numeric Separators from 'nDto' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  return new(numStrDtoMechanics).addNumStrs(
    numSeps, &n1Dto, true, &n2Dto, true, ePrefix.XCpy("nDto"))
}

// CompareSignedValues
//
//	Compares the signed numeric values of two NumStrDto objects.
//
//	The term 'signed numeric values' as used here means that the
//	two NumStrDto objects being compared may be either positive
//	or negative numeric values.
//
//	Examples
//	========
//
//	   n1         n2          Result
//
//	-9691.23     91.245         -1
//	 9691.23     91.245          1
//	   -5        82             -1
//	    5         5              0
//
//	Input Parameters
//	================
//
//	n1Dto                    *NumStrDto
//	  A pointer to an instance of NumStrDto. The signed numeric
//	  value of this object will be compared to input parameter.
//	  'n2Dto'.
//
//	n2Dto                    *NumStrDto
//	  A pointer to an instance of NumStrDto. The signed numeric
//	  value of this object will be compared to input parameter.
//	  'n1Dto'.
//
//	Return Values
//	=============
//
//	int
//	  This returned integer will be set to one of three values:
//
//	  -1 = n1Dto is less than n2Dto
//
//	   0 = n1Dto is equal to n2Dto
//
//	   1 = n1Dto is greater than n2Dto
func (nDto *NumStrDto) CompareSignedValues(n1Dto *NumStrDto, n2Dto *NumStrDto) (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CompareSignedValues",
    "")

  if err != nil {
    return 0, err
  }

  // must be n2Dto.signVal == 1
  return new(numStrDtoMolecule).compareSignedValues(
    n1Dto, true, n2Dto, true, ePrefix.XCpy("n1Dto vs n2Dto"))
}

// CompareAbsoluteValues
//
//	Compares the absolute numeric values of two NumStrDto objects.
//	The signs (+ or -) of the two compared numeric values are
//	ignored. Only the absolute numeric values are compared.
//
//	Return Values:
//	-1 = n1Dto is less than n2Dto
//	 0 = n1Dto is equal to n2Dto
//	 1 = n1Dto is greater than n2Dto
//
//	Examples
//	========
//
//	   n1             n2            Result
//
//	-9691.23         91.245            1
//	 9691.23         91.245            1
//	   -5            82               -1
//	    5             5                0
func (nDto *NumStrDto) CompareAbsoluteValues(n1Dto, n2Dto *NumStrDto) (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CompareAbsoluteValues",
    "")

  if err != nil {
    return 0, err
  }

  return new(numStrDtoAtom).compareAbsoluteValues(
    n1Dto, true, n2Dto, true, ePrefix)
}

// CopyIn
//
//	Receives an incoming NumStrDto object and copies the
//	information to the current NumStrDto data fields.
func (nDto *NumStrDto) CopyIn(nInDto NumStrDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CopyIn",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoMolecule).copy(nDto, &nInDto, true, ePrefix)
}

// CopyInPtr
//
//	Receives a pointer to an incoming NumStrDto object and copies
//	the information to the current NumStrDto data fields.
func (nDto *NumStrDto) CopyInPtr(nInDto *NumStrDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CopyInPtr",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoMolecule).copy(nDto, nInDto, true, ePrefix)
}

// CopyOut
//
//	Creates a copy of the current NumStrDto member data elements
//	and copies them to a completely new instance of NumStrDto which
//	is returned to the calling function.
func (nDto *NumStrDto) CopyOut() (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CopyOut",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error: The current NumStrDto instance is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  numStrDtoOut := NumStrDto{}

  err = new(numStrDtoMolecule).copy(
    &numStrDtoOut,
    nDto,
    false,
    ePrefix)

  return numStrDtoOut, err
}

// CopyOutPtr
//
//	Creates a copy of the current NumStrDto member data elements
//	and copies them to a completely new instance of NumStrDto which
//	is returned to the calling function.
func (nDto *NumStrDto) CopyOutPtr() (*NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.CopyOut",
    "")

  if err != nil {
    return &NumStrDto{}, err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return &NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error: The current NumStrDto instance is INVALID!",
        ErrMessage: err.Error(),
      }
  }

  numStrDtoOut := new(NumStrDto)

  err = new(numStrDtoMolecule).copy(
    numStrDtoOut,
    nDto,
    false,
    ePrefix)

  return numStrDtoOut, err
}

// Divide
//
//	Divides the current NumStrDto by input parameter 'n2Dto'.
//
//	Maximum precision of the division result is controlled by the
//	input parameter, 'maximumPrecision'.
//
//	If 'maximumPrecision' is greater than or equal to zero ('0'),
//	the number of digits to the right of the decimal place will
//	not exceed 'maximumPrecision'.
//
//	'maximumPrecision' is set equal to minus one ('-1'), will be
//	set to a maximum of 1,024 digits to the right of the decimal
//	point. If 'maximumPrecision' is less than '-1', an error will
//	be returned.
//
//	'minimumPrecision' specifies the minimum precision of the final
//	result. If 'minimumPrecision' is less than zero, an error will
//	be returned.
func (nDto *NumStrDto) Divide(n2Dto NumStrDto, minimumPrecision, maximumPrecision int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.Divide",
    "")

  if err != nil {
    return err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "Error: The current NumStrDto instance 'nDto' is INVALID!\n" +
        "'nDto' FAILED validation tests.",
      ErrMessage: err.Error(),
    }
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
        "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  quotient, err := new(numStrDtoMechanics).divideNumStrs(
    numSeps, nDto, false, &n2Dto, true, minimumPrecision, maximumPrecision, ePrefix)

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "quotient, err := new(numStrDtoMechanics).divideNumStrs(\n" +
        "  numSeps, nDto,false, &n2Dto, true, minimumPrecision, maximumPrecision, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(numStrDtoMolecule).copy(nDto, &quotient, false, ePrefix)

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
        "  nDto, &quotient, false, ePrefix)",
      ErrContext: "Error occurred while copying final quotient result into\n" +
        "current NumStrDto instance 'nDto'.",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// Equal
//
//	Returns 'true' if the input parameter NumStrDto instance
//	'n2Dto' is equal in all respects to the current NumStrDto
//	instance ('nDto').
//
//	If the current instance of NumStrDto is invalid, this method
//	returns 'false'.
//
//	Likewise, if the incoming NumStrDto 'n2Dto' is invalid, this
//	method returns 'false'.
//
//	This method differs from NumStrDto.EqualTo in that this method
//	does not return an error.
func (nDto *NumStrDto) Equal(n2Dto NumStrDto) bool {

  return new(numStrDtoAtom).equal(nDto, &n2Dto)
}

// EqualTo
//
//	Returns 'true' if the input parameter NumStrDto instance
//	'n2Dto' is equal in all respects to the current NumStrDto
//	instance ('nDto').
//
//	If the current instance of NumStrDto is invalid, this method
//	returns 'false'.
//
//	Likewise, if the incoming NumStrDto 'n2Dto' is invalid, this
//	method returns 'false'.
//
//	This is identical to NumStrDto.Equal with the sole exceiption
//	being that this method returns an 'error'.
func (nDto *NumStrDto) EqualTo(n2Dto NumStrDto) (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.EqualTo",
    "")

  if err != nil {
    return false, err
  }

  return new(numStrDtoAtom).equalTo(nDto, &n2Dto, ePrefix)
}

// Empty - Sets all the fields in the NumStrDto
// to their initial or zero state.
func (nDto *NumStrDto) Empty() {

  new(numStrDtoElectron).emptyNumStrDto(nDto)

  return
}

// FindIntArraySignificantDigitLimits
//
//	Receives an array of integers and converts them to a number
//	string consisting of significant digits. Leading and trailing
//	zeros are eliminated.
//
//	See Method: FindNumStrSignificantDigitLimits()
func (nDto *NumStrDto) FindIntArraySignificantDigitLimits(
  intArray []int,
  precision uint,
  signVal int) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.FindIntArraySignificantDigitLimits",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  nStrDtoAtom := new(numStrDtoAtom)

  numSeps, err := nStrDtoAtom.getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto->numSeps"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "  nDto, ePrefix.XCpy(\"nDto->numSeps\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps' from 'nDto'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps' from 'nDto'\").String())",
        ErrContext: "Error: Numeric Separators copied from current NumStrDto ('nDto') is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  outNStrDto, err := nStrDtoAtom.findIntArraySignificantDigitLimits(
    numSeps, intArray, precision, signVal, ePrefix)

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "outNStrDto, err := nStrDtoAtom.\n" +
          "  findIntArraySignificantDigitLimits(\n" +
          "  numSeps, intArray, precision, signVal, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return outNStrDto, nil
}

// FindNumStrSignificantDigitLimits
//
//	Analyzes an array of characters which constitute a number
//	string are returns the significant digits.
//
//	Example
//	=======
//
//	absAllRunes  precision  signVal  Result
//
//	001236700        4         1     123.67
//	000006700        4         1       0.67
//	001230000        4         1     123.0
func (nDto *NumStrDto) FindNumStrSignificantDigitLimits(
  absAllRunes []rune, precision uint, signVal int) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.FindNumStrSignificantDigitLimits",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto->numSeps"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := \n" +
          "  new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "   nDto, ePrefix.XCpy(\"nDto->numSeps\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps' from 'nDto'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps' from 'nDto'\").String())",
        ErrContext: "Error: Numeric Separators copied from current NumStrDto ('nDto') is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return new(numStrDtoMuon).findNumStrSignificantDigitLimits(
    numSeps, absAllRunes, precision, signVal, ePrefix)
}

// FormatForMathOps
//
//	Receives two NumStrDto objects and converts their number
//	strings such that both have the same number of integer and
//	fractional digits while maintaining their original numeric
//	values. This transform will facilitate the performance of
//	string based math operations such as addition and subtraction.
//
//	The return values represent the formatted NumStrDto objects.
//	The first NumStrDto returned always contains the larger
//	absolute value. The second NumStrDto always contains the
//	absolute numeric value which is less than or equal to the first
//	NumStrDto object returned.
//
//	The third parameter returned by this method is an integer
//	('compare') value which will always be set to '1' or '0'.
//
//	'1' indicates that the absolute value of the first NumStrDto
//	object ('n1DtoOut') returned by this method is greater than the
//	second NumStrDto object ('n2DtoOut') returned by this method.
//
//	If the returned integer ('compare') value returned is zero, it
//	signals that the absolute values (not the signed values) of
//	both returned NumStrDto objects are equal.
//
//	If the absolute value of 'n1Dto' is less than 'n2Dto', return
//	value 'n1DtoOut' will be populated with 'n2Dto' values, return
//	value 'n2DtoOut' will be populated with 'n1Dto' values and return
//	parameter 'isOrderReversed' will be set to 'true'.
func (nDto *NumStrDto) FormatForMathOps(n1Dto, n2Dto NumStrDto) (n1DtoOut NumStrDto, n2DtoOut NumStrDto, compare int, isOrderReversed bool, err error) {

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.FormatForMathOps",
    "")

  if err != nil {
    return NumStrDto{}, NumStrDto{}, 0, false, err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return NumStrDto{}, NumStrDto{}, 0, false,
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error: The current NumStrDto instance ('nDto') is INVALID!\n" +
          "The Numeric Seprators encapsulated by 'nDto' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return NumStrDto{}, NumStrDto{}, 0, false,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return new(numStrDtoPhoton).formatForMathOps(
    numSeps, &n1Dto, true, &n2Dto, true, ePrefix.XCpy("n1Dto,n2Dto"))
}

// FormatCurrencyStr
//
//	Formats the current NumStrDto numeric value as a currency
//	string.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	negValMode               NegativeValueFmtMode
//	  Specifies the display mode for negative values:
//
//	  LEADMINUSNEGVALFMTMODE
//	    Negative values formatted with a leading minus sign.
//	    Example: -$123,456.78
//
//	  PARENTHESESNEGVALFMTMODE
//	    Negative values formatted with surrounding parentheses.
//	    Example: ($123,456.78)
//
//	Return Values
//	=============
//
//	string
//	  A number string formatted for currency containing the numeric
//	  value of the current NumStrDto instance.
//	    Example:  $123,456.78
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nDto *NumStrDto) FormatCurrencyStr(negValMode NegativeValueFmtMode) (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.FormatCurrencyStr",
    "")

  if err != nil {
    return "", err
  }

  return new(numStrDtoMechanics).formatCurrencyStr(
    nDto, true, negValMode, ePrefix)

}

// FormatNumStr
//
//	Formats the numeric value of the current NumStrDto as number
//	string consisting of integer digits to the left of the decimal
//	point and fractional digits to the right of the decimal point,
//	if such fractional digits exist. The resulting number string
//	will NOT contain a currency symbol or thousands separators.
//
//	Example: 123456.789
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	negValMode               NegativeValueFmtMode
//	  Specifies the display mode for negative values:
//
//	    LEADMINUSNEGVALFMTMODE
//	      Negative values formatted with a leading minus sign.
//	        Example: -123456.78
//
//	    PARENTHESESNEGVALFMTMODE
//	      Negative values formatted with surrounding parentheses.
//	        Example: (123456.78)
//
//	Return Values
//	=============
//
//	string
//	  A formatted number string containing the numeric value of the
//	  current NumStrDto instance.
//	    Example:  123456.78
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nDto *NumStrDto) FormatNumStr(negValMode NegativeValueFmtMode) (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsIntRunes",
    "")

  if err != nil {
    return "", err
  }

  return new(numStrDtoAtom).formatNumStr(
    nDto, true, negValMode, ePrefix)
}

// FormatThousandsStr
//
//	Extracts the numeric value from the current NumStrDto instance
//	and returns the equivalent number string delimited with the
//	nDto.thousandsSeparator character plus the decimal separator if
//	applicable.
//
//	Example
//	=======
//
//	  numstr = 1000000.234 converted to 1,000,000.234
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	  negValMode             NegativeValueFmtMode
//	    Specifies the display mode for negative values:
//
//	    LEADMINUSNEGVALFMTMODE
//	      Negative values formatted with a leading minus sign.
//	        Example: -123,456.78
//
//	    PARENTHESESNEGVALFMTMODE
//	      Negative values formatted with surrounding parentheses.
//	        Example: (123,456.78)
//
//	Return Values
//	=============
//
//	string
//	  A number string containing the numeric value of the current
//	  NumStrDto instance formatted with thousands separators.
//	    Example:  123,456.78
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nDto *NumStrDto) FormatThousandsStr(negValMode NegativeValueFmtMode) (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.FormatThousandsStr",
    "")

  if err != nil {
    return "", err
  }

  return new(numStrDtoMechanics).formatThousandsStr(
    nDto, true, negValMode, ePrefix)
}

// GetAbsoluteBigInt
//
//	Returns the absolute value of all numeric digits in the number
//	string (nDto.absAllNumRunes). As such, Fractional digits to the
//	right of the decimal are included in the consolidate integer
//	number.
//
//	In addition, this method returns the 'precision' value therby
//	identifying the number of fractional digits to the right of
//	the decimal place in the returned *big.Int numeric value.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	All the numeric digits in the number string are
//	therefore returned as a type *big.Int numeric value.
//
//	Input Parameters
//	================
//
//	  NONE
//
//	Return Values
//	=============
//
//	*big.Int
//	  A positive (+) big int number representing the absolute value
//	  of both the integer and  fractional digits extracted from the
//	  current NumStrDto instance, 'nDto'.
//
//	uint
//	  An unsigned integer value representing the number of
//	  fractional digits to the right of the decimal point in the
//	  numeric value returned by '*big.Int'.
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nDto *NumStrDto) GetAbsoluteBigInt() (*big.Int, uint, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsoluteBigInt",
    "")

  if err != nil {
    return big.NewInt(0), 0, err
  }

  return new(numStrDtoGluon).getAbsoluteBigInt(
    nDto, true, ePrefix)
}

// GetAbsAllNumRunes
//
//	Returns an array of runes representing all the integer and
//	fractional digits included in the current NumStrDto instance.
//	The rune array returned will consist of numeric digits with no
//	sign value prefixed. This effectively returns the absolute
//	value of all integer and fractional digits combined in one rune
//	array (there is no decimal point).
func (nDto *NumStrDto) GetAbsAllNumRunes() ([]rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsAllNumRunes",
    "")

  if err != nil {
    return []rune{}, err
  }

  return new(numStrDtoGluon).getAbsAllNumRunes(
    nDto, true, ePrefix)
}

// GetAbsFracRunes
//
//	Returns all the fractional digits to the right of the decimal
//	place in the current NumStrDto instance as an array of runes.
//	The rune array is not signed; that is, the rune array does not
//	contain a '+' or '-' character in the first array position. The
//	rune array is therefore said to represent the absolute value of
//	the fractional digits in the current NumStrDto numeric value.
func (nDto *NumStrDto) GetAbsFracRunes() ([]rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsFracRunes",
    "")

  if err != nil {
    return []rune{}, err
  }

  return new(numStrDtoGluon).getAbsFracRunes(
    nDto, true, ePrefix)
}

// GetAbsFracRunesLength
//
//	Returns the length of the fractional digits in the number
//	string.
func (nDto *NumStrDto) GetAbsFracRunesLength() (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsFracRunesLength",
    "")

  if err != nil {
    return 0, err
  }

  return new(numStrDtoGluon).getAbsFracRunesLength(
    nDto, true, ePrefix)
}

// GetAbsIntRunes
//
//	Returns all the integer digits included in the current
//	NumStrDto numeric value as an array of runes. The returned rune
//	array does not contain a sign value in the first position and
//	therefore represents the absolute or positive value of all the
//	integer digits. The integer digits of a NumStrDto numeric
//	includes all the digits to the left of the decimal point.
//
//	If the current NumStrDto consists of zero integers and
//	fractional digits (Example: '0.123456'), this method will
//	return a rune array consisting one array element with a '0'
//	value.
func (nDto *NumStrDto) GetAbsIntRunes() ([]rune, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsIntRunes",
    "")

  if err != nil {
    return []rune{}, err
  }

  return new(numStrDtoGluon).getAbsIntRunes(
    nDto, true, ePrefix)
}

// GetAbsPureNumStr
//
//	Returns all digits in the current NumStrDto numeric value as a
//	pure, unsigned number string. If fractional digits exists, they
//	are included in the string and NOT separated by a decimal
//	separator.
//
//	All the numeric digits in the numeric value are	returned as a
//	type string.
//
//	Examples
//	========
//
//	Numeric
//	Value       pureNumStr    precision    numberSign
//	------      ----------    ---------    ----------
//
//	 123.45      12345            2             1
//	 12345       12345            0             1
//	-123.45      11245            2            -1
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	  NONE
//
//	Return Values
//	=============
//
//	pureNumStr               string
//	  This returned string value contains the pure number string
//	  representaion of the absolue numeric value for the current
//	  NumStrDto instance.
//
//	precision                 uint
//	  An unsigned integer value representing the number of
//	  fractional digits to the right of the decimal point in the
//	  numeric value returned by 'pureNumStr'.
//
//	numberSign                int
//	  The returned number sign will designate the 'purNumStr' as
//	  either a positive value or a negative value.
//
//	  Possible return values:
//	    +1 = Positive numeric value (including zero)
//	                OR
//	    -1 = Negative numeric value
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nDto *NumStrDto) GetAbsPureNumStr() (
  pureNumStr string, precision uint, numberSign int, err error) {

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsPureNumStr",
    "")

  if err != nil {
    return pureNumStr, precision, numberSign, err
  }

  return new(numStrDtoGluon).getAbsPureNumStr(
    nDto, true, ePrefix)
}

// GetAbsIntRunesLength
//
//	Returns the length of the integer portion of the number string.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Examples
//	========
//
//	Numeric Value    Returned Length
//
//	  -123456              6
//	   123456              6
//	   123.456             6
//	   0.123456            6
func (nDto *NumStrDto) GetAbsIntRunesLength() (int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsIntRunesLength",
    "")

  if err != nil {
    return 0, err
  }

  return new(numStrDtoGluon).getAbsIntRunesLength(
    nDto, true, ePrefix)
}

// GetBigInt
//
//	Returns an integer of type *big.Int representing the signed
//	integer value of NumStrDto.numStrDto. Decimal numbers like
//	'-123.456' will be returned as signed integer values,
//	'-123456'.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	  NONE
//
//	Return Values
//	=============
//
//	signedBIntNum            *big.In
//	  Contains the signed numeric value of the current NumStrDto
//	  instance including both integer and fractional digits.
//	  representaion of the absolue numeric value for the current
//	  NumStrDto instance.
//
//	err                      error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nDto *NumStrDto) GetBigInt() (signedBIntNum *big.Int, err error) {

  signedBIntNum = big.NewInt(0)

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetBigInt",
    "")

  if err != nil {
    return signedBIntNum, err
  }

  signedBIntNum, _, err = new(numStrDtoBoson).getSignedBigIntPrecision(
    nDto, true, ePrefix)

  if err != nil {

    return signedBIntNum,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "signedBIntNum, _, err = new(numStrDtoBoson).\n" +
          "getSignedBigIntPrecision(nDto, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return signedBIntNum, nil
}

// GetBigIntPrecision
//
//	Returns an integer of type *big.Int representing the signed
//	integer value of NumStrDto.numStrDto. Decimal numbers like
//	'-123.456' will be returned as signed integer values,
//	'-123456'.
//
//	In addition, this method also returns the 'precision'
//	specification associated with the returned *big.Int numeric
//	value. 'precision' specifies the number of digits to the
//	right of the decimal point
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
//	returned.
//
//	Input Parameters
//	================
//
//	  NONE
//
//	Return Values
//	=============
//
//	signedBIntNum            *big.In
//	  Contains the signed numeric value of the current NumStrDto
//	  instance including both integer and fractional digits.
//	  representaion of the absolue numeric value for the current
//	  NumStrDto instance.
//
//	err                      error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (nDto *NumStrDto) GetBigIntPrecision() (signedBIntNum *big.Int, precision uint, err error) {

  signedBIntNum = big.NewInt(0)

  var ePrefix *ePref.ErrPrefixDto

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetBigIntPrecision",
    "")

  if err != nil {
    return signedBIntNum, precision, err
  }

  signedBIntNum, precision, err = new(numStrDtoBoson).getSignedBigIntPrecision(
    nDto, true, ePrefix)

  if err != nil {

    return signedBIntNum, precision,
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "signedBIntNum, precision, err = new(numStrDtoBoson).\n" +
          "getSignedBigIntPrecision(nDto, true, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return signedBIntNum, precision, nil
}

// GetBigIntNum - Converts the numeric value of the
// current NumStrDto to a 'BigIntNum' type and returns
// it to the calling function.
//
// The returned BigIntNum will contain numeric separators
// (decimal separator, thousands separator and currency
// symbol) copied from the current NumStrDto instance.
//
// Before returning the BigIntNum result, this method
// performs a validity test on the current NumStrDto instance.
func (nDto *NumStrDto) GetBigIntNum() (BigIntNum, error) {
  ePrefix := "NumStrDto.GetBigIntNum() "

  err := nDto.IsValid(ePrefix + " This NumStrDto INVALID! ")

  if err != nil {
    return new(BigIntNum).NewZero(0), err
  }

  numSeps := nDto.GetNumericSeparatorsDto()

  bInt, err := nDto.GetBigInt()

  if err != nil {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+
        "Error returned by nDto.GetBigInt() "+
        "Error='%v' ", err.Error())
  }

  bIntNum := new(BigIntNum).NewBigInt(bInt, nDto.precision)

  err = bIntNum.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return new(BigIntNum).NewZero(0),
      fmt.Errorf(ePrefix+
        "Error returned by bIntNum.SetNumericSeparatorsDto(numSeps) "+
        "Error='%v' \n", err.Error())
  }

  return bIntNum, nil
}

// GetCurrencySymbol - Returns the character currently designated
// as the currency symbol for this number string.
//
// In the USA, the currency symbol is the dollar sign ('$').
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// Example: $123.45
func (nDto *NumStrDto) GetCurrencySymbol() rune {

  if nDto.currencySymbol == 0 {
    nDto.currencySymbol = '$'
  }

  return nDto.currencySymbol

}

// GetCurrencyParen - Returns the number string delimited with the
// nDto.thousandsSeparator character and the currency symbol. If
// the value is negative, the number will be surrounded in parentheses.
//
// Example:
// numstr = 1000000.23
// GetThouStr() = $1,000,000.23
//
// numstr = -1000000.23
// GetThouStr() = ($1,000,000.23)
//
// Note: If the current NumStrDto is invalid, this method
// returns an empty string.
func (nDto *NumStrDto) GetCurrencyParen() string {

  outStr, err := nDto.FormatCurrencyStr(PARENTHESESNEGVALFMTMODE)

  if err != nil {
    return ""
  }

  return outStr

}

// GetCurrencyStr - Returns the number string delimited with the
// nDto.thousandsSeparator character and the currency symbol.
// If the value is negative, a leading minus sign will be prefixed
// to the currency display.
//
// Example:
// numstr = 1000000.23
// GetCurrencyStr() = $1,000,000.23
//
// numstr = -1000000.23
// GetCurrencyStr() = -$1,000,000.23
//
// Note: If the current NumStrDto is invalid, this method
// returns an empty string.
func (nDto *NumStrDto) GetCurrencyStr() string {

  outStr, err := nDto.FormatCurrencyStr(LEADMINUSNEGVALFMTMODE)

  if err != nil {
    return ""
  }

  return outStr
}

// GetDecimalSeparator - returns the character designated
// as the decimal separator for the current NumStrDto instance.
//
// In the USA, the decimal separator is the period character ('.').
//
// Example:		123.456
func (nDto *NumStrDto) GetDecimalSeparator() rune {

  if nDto.decimalSeparator == 0 {
    nDto.decimalSeparator = '.'
  }

  return nDto.decimalSeparator

}

// GetDecimal - Converts the current NumStrDto instance
// to a Type 'Decimal' and returns it to the calling
// function.
//
// The returned Decimal instance will contain numeric
// separators (decimal separator, thousands separator
// and currency symbol) copied from the current NumStrDto
// instance.
//
// Before returning the Decimal result, this method
// performs a validity test on the current NumStrDto instance.
func (nDto *NumStrDto) GetDecimal() (Decimal, error) {

  ePrefix := "NumStrDto.GetIntAryElements() "

  err := nDto.IsValid(ePrefix)

  if err != nil {
    return Decimal{}, err
  }

  numSeps := nDto.GetNumericSeparatorsDto()

  dec, err := new(Decimal).NewNumStrWithNumSeps(nDto.GetNumStr(), numSeps)

  if err != nil {
    return Decimal{},
      fmt.Errorf(ePrefix+
        "Error returned by new(Decimal).NewNumStrWithNumSeps(nDto.GetNumStr(), numSeps) "+
        "Error='%v' ", err.Error())
  }

  return dec, nil
}

// GetIntAry
//
//	Converts the current NumStrDto instance to a Type IntAry and
//	returns it to the calling function.
func (nDto *NumStrDto) GetIntAry() (IntAry, error) {
  ePrefix := "NumStrDto.GetIntAryElements() "

  err := nDto.IsValid(ePrefix)

  if err != nil {
    return IntAry{}, err
  }

  numSeps := nDto.GetNumericSeparatorsDto()

  ia, err := new(IntAry).NewNumStrWithNumSeps(nDto.GetNumStr(), numSeps)

  if err != nil {
    return IntAry{},
      fmt.Errorf(ePrefix+
        "Error returned by new(IntAry).NewNumStrWithNumSeps(nDto.GetNumStr(), numSeps). "+
        "nDto='%v' Error='%v'", nDto.GetNumStr(), err.Error())
  }

  return ia, nil
}

// GetNumericSeparatorsDto - Returns a structure containing the
// character or rune values for decimal point separator, thousands
// separator and currency symbol.
func (nDto *NumStrDto) GetNumericSeparatorsDto() (NumericSeparatorDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetNumericSeparatorsDto",
    "")

  if err != nil {
    return NumericSeparatorDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  return numSeps, nil
}

// GetNumParen - Returns the numeric value of the current NumStrDto
// instance as a signed number string. The resulting number string
// will NOT contain a currency symbol or thousands separators. It
// will contain a decimal separator and fractional digits if such
// fractional digits exist.
//
// Note: If the current NumStrDto is invalid, this method will return
// an empty string.
//
// If the sign of the numeric value is negative, the resulting number
// string will be surrounded in parentheses.
//
// Examples:
// numeric value						result
//
//	 123456.78							123456.78
//	-123456.78             (123456.78)
func (nDto *NumStrDto) GetNumParen() string {

  outStr, err := nDto.FormatNumStr(PARENTHESESNEGVALFMTMODE)

  if err != nil {
    return ""
  }

  return outStr

}

// GetNumStr - returns the numeric value of the current NumStrDto
// instance as a signed number string. The resulting number string
// will NOT contain a currency symbol or thousands separators. It
// will contain a decimal separator and fractional digits if such
// fractional digits exist.
//
// Note: If the current NumStrDto is invalid, this method will return
// an empty string.
//
// Examples:
//
//	 123456.78
//	-123456.78
func (nDto *NumStrDto) GetNumStr() (string, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetNumStr",
    "")

  if err != nil {
    return "", err
  }

  outStr, err := new(numStrDtoAtom).formatNumStr(
    nDto, true, LEADMINUSNEGVALFMTMODE, ePrefix)

  if err != nil {

    return "",
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "outStr, err :=new(numStrDtoAtom).formatNumStr(\n" +
          "  nDto, true, LEADMINUSNEGVALFMTMODE, ePrefix)",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return outStr, nil
}

// GetNumStrDto - Returns a deep copy of the current NumStrDto
// instance.
//
// The returned NumStrDto instance will contain numeric
// separators (decimal separator, thousands separator
// and currency symbol) copied from the current NumStrDto
// instance.
//
// Before returning the NumStrDto result, this method
// performs a validity test on the current NumStrDto instance.
//
// This method is necessary in order to fulfill the requirements
// of the INumMgr interface.
func (nDto *NumStrDto) GetNumStrDto() (NumStrDto, error) {

  ePrefix := "NumStrDto.GetNumStrDto() "

  err := nDto.IsValid(ePrefix + "NumStrDto INVALID! ")

  if err != nil {
    return new(NumStrDto).New(), err
  }

  return nDto.CopyOut(), nil
}

// GetPrecision - Returns the precision of the current
// NumStrDto Instance.
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// The value of 'precision' returned by this method will
// always be >= zero (greater than or equal to zero '0').
//
// Example:
//
//					1.234    	GetPrecision() = 3
//							5			GetPrecision() = 0
//				0.12345  		GetPrecision() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
func (nDto *NumStrDto) GetPrecision() int {
  return int(nDto.precision)
}

// GetPrecisionUint - Returns the precision of the
// current NumStrDto Instance as an unsigned integer
// (uint). precision represents the number of fractional
// digits to the right of the decimal point.
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal point in a string of numeric digits, go
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
func (nDto *NumStrDto) GetPrecisionUint() (uint, error) {
  return nDto.precision, nil
}

// GetRationalNumber - returns the sign value of the number string, plus the
// numeric value of the number string expressed as a Rational Number.
//
// This method will return an error if the NumStrDto fields are not properly
// initialized and populated.
//
// Returns
// =======
//
// sign value  						int 			- sign value of the number string
// big Rational Number		*big.Rat	- Number string expressed as a
//
//	rational number
//
// err										error			- In case of failure, an 'error' type
//
//	is returned. In case of success this
//	value is 'nil'
func (nDto *NumStrDto) GetRationalNumber() (int, *big.Rat, error) {

  ePrefix := "NumStrDto.GetRationalNumber() "

  err := nDto.IsValid("")

  if err != nil {
    return 0, big.NewRat(1, 1),
      fmt.Errorf(ePrefix+"This NumStrDto instance is INVALID! Error='%v'", err.Error())
  }

  ratZero := big.NewRat(0, 1)

  err = nDto.IsValid(ePrefix)

  if err != nil {
    return 0, ratZero, errors.New(ePrefix +
      "- Error: The existing NumStrDto is corrupted or improperly initialized. " +
      "Re-initialize the NumStrDto object and try again.")
  }

  signVal := nDto.signVal

  absInt, isOk := big.NewInt(0).SetString(string(nDto.absAllNumRunes), 10)

  if !isOk {
    return 0, ratZero, fmt.Errorf(ePrefix+
      "- Conversion of nDto.absAllNumRunes to big.Int Failed! "+
      "nDto.absIntRunes= '%v'", nDto.absAllNumRunes)
  }

  base10 := big.NewInt(10)

  bigPrecision := big.NewInt(int64(nDto.precision))

  scaleFactor := big.NewInt(0).Exp(base10, bigPrecision, nil)

  rationalNum := big.NewRat(1, 1).SetFrac(absInt, scaleFactor)

  return signVal, rationalNum, nil
}

// GetScaleFactor
//
//	Returns the scale factor for the number string encapsulated by
//	the current instance of NumStrDto.
//
//	Scale factor is defined by 10 raised to the power of
//	'precision' (nDto.precision).  nDto.precision is the number of
//	digits to the right of the decimal point.
//
//	This method will fail and return an error if the current instance
//	of NumStrDto is invalid.
func (nDto *NumStrDto) GetScaleFactor() (*big.Int, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetScaleFactor",
    "")

  if err != nil {
    return big.NewInt(0), err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return big.NewInt(0),
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error: The current NumStrDto instance ('nDto') is INVALID!\n" +
          "'nDto' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return new(numStrDtoGluon).getScaleFactor(
    nDto, false, ePrefix)
}

// GetSciNotationNumber - Converts the numeric value of the current
// NumStrDto instance into scientific notation and returns this value
// as an instance of type SciNotationNum.
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
func (nDto *NumStrDto) GetSciNotationNumber(mantissaLen uint) (SciNotationNum, error) {

  ePrefix := "NumStrDto.GetSciNotationNumber() "

  bINum, err := nDto.GetBigIntNum()

  if err != nil {
    return new(SciNotationNum).New(),
      fmt.Errorf(ePrefix+
        "Error returned by nDto.GetBigIntNum(). Error='%v'",
        err.Error())
  }

  sciNotation, err := bINum.GetSciNotationNumber(mantissaLen)

  if err != nil {
    return new(SciNotationNum).New(),
      fmt.Errorf(ePrefix+
        "Error returned by bINum.GetSciNotationNumber(mantissaLen). Error='%v'",
        err.Error())
  }

  return sciNotation, nil
}

// GetSciNotationStr - Returns a string expressing the current NumStrDto
// numerical value as scientific notation.
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
func (nDto *NumStrDto) GetSciNotationStr(mantissaLen uint) (string, error) {

  ePrefix := "NumStrDto.GetSciNotationStr() "

  if mantissaLen < 2 {
    mantissaLen = 2
  }

  sciNotation, err := nDto.GetSciNotationNumber(mantissaLen)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned by bNum.GetSciNotationNumber(mantissaLen). "+
        "Error='%v'", err.Error())
  }

  result, err := sciNotation.GetSciNotationStr(mantissaLen)

  if err != nil {
    return "",
      fmt.Errorf(ePrefix+
        "Error returned by sciNotation.GetSciNotationStr(mantissaLen). "+
        "Error='%v'", err.Error())
  }

  return result, nil
}

// GetSign - Returns the sign value for this NumStrDto
// numeric value. Return values will be either +1 or -1.
func (nDto *NumStrDto) GetSign() (int, error) {
  return nDto.signVal, nil
}

// GetThisPointer - Returns a pointer to the current NumStrDto instance.
func (nDto *NumStrDto) GetThisPointer() *NumStrDto {

  return nDto
}

// GetThouParen - Returns the number string delimited with the
// nDto.thousandsSeparator character. Negative values are
// surrounded in parentheses.
//
// Example:
// numstr = 1000000.234
// GetThouStr() = 1,000,000.234
//
// numstr = -1000000.234
// GetThouStr() = (1,000,000.234)
//
// Note: If the current NumStrDto is invalid, this method
// returns an empty string.
func (nDto *NumStrDto) GetThouParen() string {

  outStr, err := nDto.FormatThousandsStr(PARENTHESESNEGVALFMTMODE)

  if err != nil {
    return ""
  }

  return outStr
}

// GetThouStr - Returns the number string delimited with the
// nDto.thousandsSeparator character plus the Decimal Separator
// if applicable.
//
// Example:
// numstr = 1000000.234
// GetThouStr() = 1,000,000.234
//
// numstr = -1000000.234
// GetThouStr() = -1,000,000.234
//
// Note: If the current NumStrDto is invalid, this method
// returns an empty string.
func (nDto *NumStrDto) GetThouStr() string {

  outStr, err := nDto.FormatThousandsStr(LEADMINUSNEGVALFMTMODE)

  if err != nil {
    return ""
  }

  return outStr
}

// GetThousandsSeparator - returns a rune which represents
// the character currently used to separate thousands in
// the display of the current NumStrDto number string.
//
// In the USA, the thousands separator is a comma character.
//
// Example: 1,000,000,000
func (nDto *NumStrDto) GetThousandsSeparator() rune {

  return nDto.thousandsSeparator
}

// GetZeroNumStrDto
//
//	Returns a new NumStrDto initialized to zero value. If the
//	parameter numFracDigits is set to a value greater than zero,
//	then an equal number of zero characters will be added to the
//	right of the decimal point.
//
//	Examples
//	========
//
//	numFracDigits    Results NumStrOut
//
//	     0                "0"
//	     2                "0.00"
//	     4                "0.0000"
func (nDto *NumStrDto) GetZeroNumStrDto(numFracDigits uint) NumStrDto {

  numSeps := NumericSeparatorDto{}

  numSeps.DecimalSeparator = nDto.decimalSeparator
  numSeps.ThousandsSeparator = nDto.thousandsSeparator
  numSeps.CurrencySymbol = nDto.currencySymbol

  return new(numStrDtoMolecule).newZeroNumStrDto(numSeps, numFracDigits)
}

// HasNumericDigits - returns 'false' if the number
// string for the current NumStrDto instance is uninitialized
// and contains no numeric digits. In this case, the length
// of the number string is zero characters.
//
// If this method returns 'true' it signals that there is at
// least one numeric digit in the number string, even if that
// digit is zero.
func (nDto *NumStrDto) HasNumericDigits() bool {

  err := nDto.IsValid("NumStrDto.HasNumericDigits() ")

  if err != nil {
    return false
  }

  return true
}

// IsFractionalValue - Returns 'true' if the numeric value of the
// current NumStrDto object includes a fractional value; that is,
// the number has fractional digits to the right of the decimal
// point.
func (nDto *NumStrDto) IsFractionalValue() bool {

  if nDto.precision > 0 {
    return true
  }

  return false
}

// IsNumStrZeroValue
//
//	Receives an external instance of NumStrDto for analysis and
//	returns 'true' if all the digits in the number string are zero.
func (nDto *NumStrDto) IsNumStrZeroValue(numDto *NumStrDto) (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsAllNumRunes",
    "")

  if err != nil {
    return true, err
  }

  return new(numStrDtoElectron).isNumStrZeroValue(
    numDto, true, ePrefix.XCpy("numDto=0 ?"))
}

// IsValid
//
//	Performs a diagnostic review of the current NumStrDto instance
//	and returns 'nil' if the NumStrDto object is valid in all
//	respects.
//
//	If the NumStrDto instance is judged invalid, an error message is
//	returned.
func (nDto *NumStrDto) IsValid(callingFunction string) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    callingFunction,
    "NumStrDto.IsValid",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))
}

// IsZero
//
//	Analyzes the current instance of NumStrDto and returns 'true'
//	if all the digits in the number string are zero.
func (nDto *NumStrDto) IsZero() (bool, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.GetAbsAllNumRunes",
    "")

  if err != nil {
    return true, err
  }

  return new(numStrDtoElectron).isNumStrZeroValue(
    nDto, true, ePrefix.XCpy("Validating nDto"))
}

// Multiply
//
//	Multiplies the current NumStrDto by the input parameter
//	NumStrDto ('n2Dto') and stores the result in the current
//	NumStrDto instance ('nDto').
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	This method will not modify the Numeric Separators previously
//	configured for this current instance of NumStrDto ('nDto').
//	Numeric Separators will therefore remain unchanged.
func (nDto *NumStrDto) Multiply(n2Dto NumStrDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.Multiply",
    "")

  if err != nil {
    return err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "Error: The current NumStrDto instance ('nDto') is INVALID!\n" +
        "'nDto' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
        "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  n1Dto := NumStrDto{}

  err = new(numStrDtoMolecule).copy(
    &n1Dto,
    nDto,
    false,
    ePrefix.XCpy("nDto -> n1Dto"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
        "&n1Dto, nDto, false ePrefix)",
      ErrContext: "Error Copying current instance 'nDto' to 'n1Dto'",
      ErrMessage: err.Error(),
    }
  }

  nResult, err := new(numStrDtoTau).multiplyNumStrs(
    numSeps, &n1Dto, false, &n2Dto, true, ePrefix)

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "nResult, err := new(numStrDtoTau).multiplyNumStrs(\n" +
        "numSeps, &n1Dto, false, &n2Dto, true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  err = new(numStrDtoMolecule).copy(
    nDto,
    &nResult,
    false,
    ePrefix.XCpy("nResult -> nDto"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
        "nDto, &nResult, false ePrefix)",
      ErrContext: "Error Copying Final Result 'nResult' to 'nDto'",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// MultiplyNumStrs
//
//	Multiplies two NumStrDto instances and returns the result as a
//	separate NumStrDto instance.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators provided by the
//	current instance of NumStrDto ('nDto'). If these Numeric
//	Separators are determined to be invalid, an error will be
//	returned.
func (nDto *NumStrDto) MultiplyNumStrs(
  n1Dto NumStrDto, n2Dto NumStrDto) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.MultiplyNumStrs",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto->numSeps"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "  nDto, ePrefix.XCpy(\"nDto->numSeps\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps' from 'nDto'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps' from 'nDto'\").String())",
        ErrContext: "Error: Numeric Separators copied from current NumStrDto ('nDto') is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return new(numStrDtoTau).multiplyNumStrs(
    numSeps, &n1Dto, true, &n2Dto, true, ePrefix)
}

// NewBigFloat - Creates a new NumStrDto instance from a Big Float value
// (*big.Float) and a precision specification.
func (nDto *NumStrDto) NewBigFloat(
  bigFloat *big.Float, precision int) (NumStrDto, error) {

  ePrefix := "NumStrDto.NewBigFloat() "

  numStr := bigFloat.Text('f', precision)

  n2, err := new(NumStrDto).NewPtr().ParseNumStr(numStr)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"Error returned by n.ParseNumStr(numStr). "+
        "numStr='%v'  Error='%v'",
        numStr, err.Error())
  }

  return n2, nil

}

// NewBigInt - Creates a new NumStrDto instance from a signed big integer (*big.Int) and
// a precision specification.
func (nDto *NumStrDto) NewBigInt(signedBigInt *big.Int, precision uint) (NumStrDto, error) {

  ePrefix := "NumStrDto.NewBigInt() "

  n2, err := new(NumStrDto).ParseSignedBigInt(
    big.NewInt(0).Set(signedBigInt),
    precision)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"Error returned by ParseSignedBigInt(signedBigInt, precision). "+
        "signedBigInt='%v' precision='%v'  Error='%v'",
        signedBigInt.Text(10), precision, err.Error())
  }

  err = n2.IsValid(ePrefix + "'n2' INVALID! ")

  if err != nil {
    return NumStrDto{}, err
  }

  return n2, nil
}

// NewBigIntNum - Receives a BigIntNum and converts it to a NumStrDto
// instance which is returned to the calling function.
func (nDto *NumStrDto) NewBigIntNum(bINum BigIntNum) (NumStrDto, error) {
  ePrefix := "NumStrDto.NewBigIntNum() "
  n2, err := new(NumStrDto).ParseBigIntNum(bINum)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"Error returned by ParseBigIntNum(bINum). "+
        "bINum='%v'  Error='%v'",
        bINum.GetNumStr(), err.Error())
  }

  return n2, nil

}

// NewFloat32 - Creates a new NumStrDto instance from a float32
// and precision specification.
func (nDto *NumStrDto) NewFloat32(f32 float32, precision int) (NumStrDto, error) {

  ePrefix := "NumStrDto.NewFloat32() "

  numStr := strconv.FormatFloat(float64(f32), 'f', precision, 32)

  n2, err := new(NumStrDto).NewPtr().ParseNumStr(numStr)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"Error returned by n.ParseNumStr(numStr). "+
        "numStr='%v'  Error='%v'",
        numStr, err.Error())
  }

  return n2, nil

}

// NewFloat64 - Creates a new NumStrDto instance from a float64
// and precision specification.
func (nDto *NumStrDto) NewFloat64(f64 float64, precision int) (NumStrDto, error) {

  ePrefix := "NumStrDto.NewFloat64() "

  numStr := strconv.FormatFloat(f64, 'f', precision, 64)

  n2, err := new(NumStrDto).NewPtr().ParseNumStr(numStr)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"Error returned by n.ParseNumStr(numStr). "+
        "numStr='%v'  Error='%v'",
        numStr, err.Error())
  }

  return n2, nil
}

// NewInt
//
// Creates a new NumStrDto from an int and a precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place.
//
// The 'NewInt' method is designed to used in conjunction with
// NumStrDto{} syntax thereby allowing NumStrDto type creation and
// initialization in one step.
//
// Example: new(NumStrDto).NewInt(123456, 3) yields a new NumStrDto
// instance with a numeric value of 123.456.
func (nDto *NumStrDto) NewInt(intNum int, precision uint) NumStrDto {

  n2 := new(NumStrDto).NewInt64(int64(intNum), precision)

  return n2
}

// NewIntExponent - Returns a new NumStrDto instance. The numeric
// value is set using an integer multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//	numeric value = integer X 10^exponent
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
//		nDto := new(NumStrDto).NewIntExponent(123456, -3)
//	 -- nDto is now equal to "123.456", precision = 3
//
//		nDto := new(NumStrDto).NewIntExponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  intNum			 exponent			  	NumStrDto Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (nDto *NumStrDto) NewIntExponent(intNum int, exponent int) NumStrDto {

  return new(NumStrDto).NewInt64Exponent(int64(intNum), exponent)
}

// NewInt32 - Creates a new NumStrDto instance from an int32 and a
// precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place.
//
// The 'NewInt32' method is designed to used in conjunction with
// NumStrDto{} syntax thereby allowing NumStrDto type creation and
// initialization in one step.
//
// Example: new(NumStrDto).NewInt32(123456, 3) yields a new NumStrDto
// instance with a numeric value of 123.456.
func (nDto *NumStrDto) NewInt32(int32Num int32, precision uint) NumStrDto {

  n2 := new(NumStrDto).NewInt64(int64(int32Num), precision)

  return n2
}

// NewInt32Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an int32 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//	numeric value = int32 X 10^exponent
//
// For example, if exponent is -3, precision is set equal to 'int32Num'
// divided by 10^+3. Example:
//
//	  int32Num			exponent			NumStrDto Result
//		 123456		 		  -3							123.456
//
// If exponent is +3, int32Num is multiplied by 10 raised to the
// power of exponent and precision is set equal to exponent.
//
//	  int32Num			exponent			NumStrDto Result
//		 123456		 		   +3							123456.000
func (nDto *NumStrDto) NewInt32Exponent(int32Num int32, exponent int) NumStrDto {

  return new(NumStrDto).NewInt64Exponent(int64(int32Num), exponent)
}

// NewInt64
//
//	Creates a new NumStrDto instance from an int64 value and a
//	precision specification.
//
//	Input parameter 'precision' indicates the number of digits to
//	be formatted to the right of the decimal place.
//
//	Example
//	=======
//
//	          new(NumStrDto).NewInt64(123456, 3)
//	Yields a NumStrDto instance with a numeric value of 123.456.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The Numeric Separators originally configured for the current
//	instance of NumStrDto will be copied to the new, returned
//	instance of NumStrDto. If these Numeric Separators are
//	determined to be invalid, an error will be returned.
func (nDto *NumStrDto) NewInt64(i64 int64, precision uint) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetThisPrecision",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{}, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
      ErrContext: "Error: The current instance of NumStrDto ('nDto') is INVALID!\n" +
        "Numeric Separators from 'nDto' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  return new(numStrDtoMolecule).newInt64(
    numSeps, i64, precision, ePrefix)
}

// NewInt64Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an int64 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//	numeric value = int64 X 10^exponent
//
// Input parameter 'int64Num' is of type int64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//		nDto := new(NumStrDto).NewInt64Exponent(123456, -3)
//	 -- nDto is now equal to "123.456", precision = 3
//
//		nDto := new(NumStrDto).NewInt64Exponent(123456, 3)
//	 -- decNum is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  int64Num		 exponent			  	Decimal Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (nDto *NumStrDto) NewInt64Exponent(int64Num int64, exponent int) NumStrDto {

  numStr := strconv.FormatInt(int64Num, 10)

  if exponent > 0 {
    for i := 0; i < exponent; i++ {
      numStr += "0"
    }
  }

  if exponent < 0 {
    exponent = exponent * -1
  }

  var n2 NumStrDto

  if exponent == 0 {
    n2, _ = new(NumStrDto).NewNumStr(numStr)
  } else {
    n2, _ = nDto.ShiftPrecisionLeft(numStr, uint(exponent))
  }

  return n2
}

// NewUint - Creates a new NumStrDto instance from an uint and a
// precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'
//
// Usage:
// ------
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//					uintNum := uint(123456)
//					precision := uint(3)
//					nDto := new(NumStrDto).NewUint(uintNum, precision)
//	       nDto is now equal to 123.456
//
// Examples:
// ---------
//
//	  uintNum			precision			NumStrDto Result
//		 123456		 		   4							12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (nDto *NumStrDto) NewUint(uintNum uint, precision uint) NumStrDto {

  n2 := new(NumStrDto).NewUint64(uint64(uintNum), precision)

  return n2
}

// NewUintExponent - Returns a new NumStrDto instance. The numeric
// value is set using an uint value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//	numeric value = int64 X 10^exponent
//
// Usage:
// ------
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//		nDto := new(NumStrDto).NewUintExponent(123456, -3)
//	 -- nDto is now equal to "123.456", precision = 3
//
//		nDto := new(NumStrDto).NewUintExponent(123456, 3)
//	 -- nDto is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  uintNum			exponent			NumStrDto Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (nDto *NumStrDto) NewUintExponent(uintNum uint, exponent int) NumStrDto {

  return nDto.NewUint64Exponent(uint64(uintNum), exponent)
}

// NewUint32 - Creates a new NumStrDto instance from an uint32 and a
// precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place and is passed as type
// 'uint'.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//					uint32Num := uint32(123456)
//					precision := uint(3)
//					nDto := new(NumStrDto).NewUint32(uint32Num, precision)
//	       nDto is now equal to 123.456
//
// Examples:
// ---------
//
//	  uint32Num		precision			NumStrDto Result
//		 123456		 		   4							12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (nDto *NumStrDto) NewUint32(uint32Num uint32, precision uint) NumStrDto {

  n2 := new(NumStrDto).NewUint64(uint64(uint32Num), precision)

  return n2
}

// NewUint32Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an uint32 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//	numeric value = int64 X 10^exponent
//
// Input parameter 'uint32Num' is of type uint32.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//		nDto := new(NumStrDto).NewUint32Exponent(123456, -3)
//	 -- nDto is now equal to "123.456", precision = 3
//
//		nDto := new(NumStrDto).NewUint32Exponent(123456, 3)
//	 -- nDto is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  uint32Num		exponent			NumStrDto Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (nDto *NumStrDto) NewUint32Exponent(uint32Num uint32, exponent int) NumStrDto {

  return nDto.NewUint64Exponent(uint64(uint32Num), exponent)
}

// NewUint64 - Creates a new NumStrDto instance from an uint64 and a
// precision specification.
//
// Input parameter 'precision' indicates the number of digits to be
// formatted to the right of the decimal place.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//					uint64Num := uint64(123456)
//					precision := uint(3)
//					nDto := new(NumStrDto).NewUint64(uint64Num, precision)
//	       nDto is now equal to 123.456
//
// Examples:
// ---------
//
//	  uint64Num		precision			NumStrDto Result
//		 123456		 		   4							12.3456
//	  123456          0              123456
//	  123456          1              12345.6
func (nDto *NumStrDto) NewUint64(uint64Num uint64, precision uint) NumStrDto {

  ePrefix := "NumStrDto.NewUint64() "

  numStr := strconv.FormatUint(uint64Num, 10)

  n2, err := new(NumStrDto).NewPtr().ParseNumStr(numStr)
  // This should NEVER produce an error
  if err != nil {
    sError := fmt.Sprintf(ePrefix+
      "Fatal Error returned by new(NumStrDto).NewPtr().ParseNumStr(numStr) "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
    panic(sError)
  }

  n2.SetThisPrecision(precision, true)

  n2.SetNumericSeparatorsDto(nDto.GetNumericSeparatorsDto())

  return n2
}

// NewUint64Exponent - Returns a new NumStrDto instance. The numeric
// value is set using an uint64 value multiplied by 10 raised to the
// power of the 'exponent' parameter.
//
//	numeric value = int64 X 10^exponent
//
// Input parameter 'uint64Num' is of type uint64.
//
// Input parameter 'exponent' is of type int.
//
// Usage:
// ------
// This method is designed to be used in conjunction with the NumStrDto{}
// syntax thereby allowing NumStrDto type creation and initialization in
// one step.
//
//		nDto := new(NumStrDto).NewUint64Exponent(123456, -3)
//	 -- nDto is now equal to "123.456", precision = 3
//
//		nDto := new(NumStrDto).NewUint64Exponent(123456, 3)
//	 -- nDto is now equal to "123456.000", precision = 3
//
// Examples:
// ---------
//
//	  uint64Num		exponent			NumStrDto Result
//		 123456		 		  -3							123.456
//		 123456		 		   3							123456.000
//	  123456          0              123456
func (nDto *NumStrDto) NewUint64Exponent(uint64Num uint64, exponent int) NumStrDto {

  ePrefix := "NumStrDto.NewUint64Exponent() "
  numStr := strconv.FormatUint(uint64Num, 10)

  if exponent > 0 {
    for i := 0; i < exponent; i++ {
      numStr += "0"
    }
  }

  if exponent < 0 {
    exponent = exponent * -1
  }

  var n2 NumStrDto
  var err error

  if exponent == 0 {
    n2, err = new(NumStrDto).NewNumStr(numStr)
    // This should never produce an error.
    if err != nil {
      sErr := fmt.Sprintf(ePrefix+
        "Fatal Error returned by new(NumStrDto).NewNumStr(numStr). "+
        "numStr='%v' uint64Num='%v' Error='%v'",
        numStr, uint64Num, err.Error())
      panic(sErr)
    }

  } else {
    n2, err = nDto.ShiftPrecisionLeft(numStr, uint(exponent))
    // This should never produce an error.
    if err != nil {
      sErr := fmt.Sprintf(ePrefix+
        "Fatal Error returned by nDto.ShiftPrecisionLeft("+
        "numStr, uint(exponent)). "+
        "numStr='%v' uint64Num='%v' exponent='%v' Error='%v'",
        numStr, uint64Num, exponent, err.Error())
      panic(sErr)
    }

  }

  n2.SetNumericSeparatorsDto(nDto.GetNumericSeparatorsDto())

  return n2
}

// NewRational - Creates a new NumStrDto instance from a rational number and a precision
// specification.
//
// For information on Big Rational Numbers (*big.Rat), see https://golang.org/pkg/math/big/
func (nDto *NumStrDto) NewRational(bigRat *big.Rat, precision int) (NumStrDto, error) {

  ePrefix := "NumStrDto.NewRational() "
  numStr := bigRat.FloatString(precision)

  n2, err := new(NumStrDto).NewPtr().ParseNumStr(numStr)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"Error returned by n.ParseNumStr(numStr). "+
        "numStr='%v'  Error='%v'",
        numStr, err.Error())
  }

  return n2, nil
}

// NewNumStr - Used to create a populated NumStrDto instance.
// using a valid number string as an input parameter.
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
// separators, see method NumStrDto.NewNumStrWithNumSeps(), below.
//
// Usage Example:
//
//	n, err := new(NumStrDto).NewNumStr("123.456")
func (nDto *NumStrDto) NewNumStr(numStr string) (NumStrDto, error) {

  ePrefix := "NumStrDto.NewNumStr() "

  n := new(NumStrDto).New()

  n2, err := n.ParseNumStr(numStr)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"Error returned by n.ParseNumStr(numStrDto). "+
        "numStrDto='%v'  Error='%v'",
        numStr, err.Error())
  }

  return n2, nil
}

// NewNumStrWithNumSeps - Receives a number string as input and returns a
// new NumStrDto instance. The input parameter 'numSeps' contains numeric
// separators (decimal separator, thousands separator and currency symbol)
// which will be used to parse the number string.
//
// In addition, the numeric separators contained in input parameter 'numSeps'
// will be copied to the returned NumStrDto instance.
func (nDto *NumStrDto) NewNumStrWithNumSeps(
  numStr string,
  numSeps NumericSeparatorDto) (NumStrDto, error) {

  ePrefix := "IntAry.NewNumStrWithNumSeps() "

  n := new(NumStrDto).New()

  numSeps.SetDefaultsIfEmpty()

  err := n.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+
        "Error returned by  Ary.SetIntAryWithNumStr(numStr). "+
        "Error='%v' ", err.Error())
  }

  n2, err := n.ParseNumStr(numStr)

  if err != nil {
    return NumStrDto{},
      fmt.Errorf(ePrefix+
        "Error returned by n.ParseNumStr(numStr). "+
        "numStr='%v', Error='%v' ", numStr, err.Error())
  }

  return n2, nil
}

// New
//
//	Used to create a new instance of NumStrDto.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators copied from the current
//	instance of NumStrDto ('nDto'). If these Numeric Separators are
//	determined to be invalid, they will be automatically reset to
//	USA default values.
//
//	Final Result
//	============
//	This method returns a new instance of NumStrDto.
//
//	The numeric value of the returned NumStrDto object will be set
//	to zero ('0') with a precision of zero ('0').
func (nDto *NumStrDto) New() NumStrDto {

  numSeps := NumericSeparatorDto{}
  numSeps.DecimalSeparator = nDto.decimalSeparator
  numSeps.ThousandsSeparator = nDto.thousandsSeparator
  numSeps.CurrencySymbol = nDto.currencySymbol

  numSeps.SetDefaultsIfEmpty()

  return new(numStrDtoMolecule).newZeroNumStrDto(numSeps, 0)
}

// NewNumSeps
//
//		Used to create a new instance of NumStrDto.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and
//		convert them into numeric values.
//
//		The final NumStrDto result returned by this method will be
//		configured with the Numeric Separators provided by input
//	 parameter 'numSeps'. If 'numSeps' is determined to be invalid,
//	 an error will be returned.
//
//		Final Result
//		============
//		This method returns a new instance of NumStrDto.
//
//		The numeric value of the returned NumStrDto object will be set
//		to zero ('0') with a precision of zero ('0').
func (nDto *NumStrDto) NewNumSeps(numSeps NumericSeparatorDto) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.NewNumSeps",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
        ErrContext: "Error: Input parameter 'numSeps' is INVALID!\n" +
          "'numSeps' FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return new(numStrDtoMolecule).newZeroNumStrDto(numSeps, 0), nil
}

// NewPtr
//
//	Used to create a new instance of NumStrDto.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators copied from the current
//	instance of NumStrDto ('nDto'). If these Numeric Separators are
//	determined to be invalid, they will be automatically reset to
//	USA default values.
//
//	Final Result
//	============
//	This method returns a pointer to the newly created NumStrDto
//	instance.
//
//	The numeric value of the returned *NumStrDto will be set to
//	zero ('0') with a precision of zero ('0').
func (nDto *NumStrDto) NewPtr() *NumStrDto {

  numSeps := NumericSeparatorDto{}
  numSeps.DecimalSeparator = nDto.decimalSeparator
  numSeps.ThousandsSeparator = nDto.thousandsSeparator
  numSeps.CurrencySymbol = nDto.currencySymbol

  numSeps.SetDefaultsIfEmpty()

  newNumStrDto := new(numStrDtoMolecule).newZeroNumStrDto(numSeps, 0)

  return &newNumStrDto
}

// NewZero
//
//	Creates a new NumStrDto with a value of zero and a precision
//	specified by input parameter 'precision'.
//
//	Examples
//	========
//
//	precision      Results NumStrOut
//
//	     0                "0"
//	     2                "0.00"
//	     4                "0.0000"
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	If this value exceeds the maximum value for a 32-bit integer,
//	this value will be automatically reduced to the maximum
//	limit of 2,147,483,647 or	2^31 - 1.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators copied from the current
//	instance of NumStrDto ('nDto'). If these Numeric Separators are
//	determined to be invalid, they will be automatically reset to
//	USA default values.
func (nDto *NumStrDto) NewZero(precision uint) NumStrDto {

  numSeps := NumericSeparatorDto{}
  numSeps.DecimalSeparator = nDto.decimalSeparator
  numSeps.ThousandsSeparator = nDto.thousandsSeparator
  numSeps.CurrencySymbol = nDto.currencySymbol

  numSeps.SetDefaultsIfEmpty()

  return new(numStrDtoMolecule).newZeroNumStrDto(numSeps, precision)
}

// ParseBigIntNum
//
//	Receives a BigIntNum instance ('biNum') and coverts it to a
//	NumStrDto instance which is returned to the calling function.
//
//	If 'biNum' proves to be invalid, an error will be returned.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators copied from the input
//	parameter 'biNum'.
func (nDto *NumStrDto) ParseBigIntNum(biNum BigIntNum) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.ParseBigIntNum",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  err = new(bigIntNumAtom).isBigIntNumValid(
    &biNum,
    ePrefix.XCpy("Validating Input Parameter 'biNum'"))

  if err != nil {

    return NumStrDto{}, &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(bigIntNumAtom).isBigIntNumValid(\n" +
        "  biNum, ePrefix)",
      ErrContext: "Error: Input parameter 'biNum' (BigIntNum) is INVALID!\n" +
        "'biNum' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  numSeps, err := biNum.GetNumericSeparatorsDto()

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "numSeps, err := biNum.GetNumericSeparatorsDto()",
        ErrContext: "",
        ErrMessage: err.Error(),
      }
  }

  return new(numStrDtoQuark).parseBigIntNum(
    numSeps, &biNum, false, ePrefix)
}

// ParseSignedBigInt
//
//	Receives a signed *Big Int number and a precision parameter. It
//	then generates and returns a new instance of NumStrDto.
//
//	'precision'
//	===========
//
//	'precision' determines the number of digits to the right of the
//	decimal place.
//
//	   signedBigInt   precision     	    result
//	    946254            3               946.254
//	    946254            0               946254
//	   -946254            3              -946.254
//	   -946254            0              -946254
//
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	The maximum limit for a 'precision' uint value is
//	2,147,483,647 or	2^31 - 1. This is also the maximum
//	allowable limit for a signed 32-bit integer.
//
//	If the 'precision' value exceeds the maximum allowable limit,
//	an error will be returned.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators copied from the current
//	NumStrDto instance ('nDto'). If these Numeric Separators prove
//	to be invalid, they will be automatically reset to default USA
//	values.
func (nDto *NumStrDto) ParseSignedBigInt(signedBigInt *big.Int, precision uint) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.ParseSignedBigInt",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
        ErrContext: "Error extracting Numeric Separators from current instance of NumStrDto.",
        ErrMessage: err.Error(),
      }
  }

  numSeps.SetDefaultsIfEmpty()

  return new(numStrDtoQuark).parseSignedBigInt(
    numSeps, signedBigInt, precision, ePrefix)
}

// ParseSignedBigIntNumSeps
//
//	Receives a signed *Big Int number and a precision parameter. It
//	then generates and returns a new instance of NumStrDto.
//
//	'precision'
//	===========
//
//	'precision' determines the number of digits to the right of the
//	decimal place.
//
//	   signedBigInt    precision      result
//	    946254            3            946.254
//	    946254            0            946254
//	   -946254            3           -946.254
//	   -946254            0           -946254
//
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is an unsigned integer value.
//	The maximum limit for a 'precision' uint value is
//	2,147,483,647 or	2^31 - 1. This is also the maximum
//	allowable limit for a signed 32-bit integer.
//
//	If the 'precision' value exceeds the maximum allowable limit,
//	an error will be returned.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators copied from input
//	parameter 'numSeps'. If these Numeric Separators prove
//	to be invalid, an error will be returned.
func (nDto *NumStrDto) ParseSignedBigIntNumSeps(
  signedBigInt *big.Int,
  precision uint,
  numSeps NumericSeparatorDto) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.ParseSignedBigIntNumSeps",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps.SetDefaultsIfEmpty()

  return new(numStrDtoQuark).parseSignedBigInt(
    numSeps, signedBigInt, precision, ePrefix)
}

// ParseSignedBigIntPrecision
//
//	Receives a signed *Big Int number and a *Big Int precision
//	parameter. This method then generates and returns a new
//	instance of NumStrDto.
//
//	'precision'
//	===========
//
//	'precision' determines the number of digits to the right of the
//	decimal place.
//
//	   signedBigInt    precision      result
//	    946254            3            946.254
//	    946254            0            946254
//	   -946254            3           -946.254
//	   -946254            0           -946254
//
//
//	Maximum Precision Value
//	=======================
//
//	Input parameter 'precision' is a signed *Big Int value.
//	The maximum limit for a 'precision' value is positive number,
//	2,147,483,647 or	2^31 - 1. This is also the maximum
//	allowable limit for a signed 32-bit integer.
//
//	If the 'precision' value exceeds the maximum allowable limit,
//	an error will be returned.
//
//	Likewise, if 'precison' is less than zero, an error will be
//	returned.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The final NumStrDto result returned by this method will be
//	configured with the Numeric Separators copied from input
//	parameter 'numSeps'. If these Numeric Separators prove
//	to be invalid, an error will be returned.
func (nDto *NumStrDto) ParseSignedBigIntPrecision(
  signedBigInt *big.Int,
  precision *big.Int,
  numSeps NumericSeparatorDto) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.ParseSignedBigIntPrecision",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  return new(numStrDtoPhoton).parseSignedBigIntPrecision(
    numSeps, signedBigInt, precision, ePrefix)
}

// ParseNumStr
//
//	Receives a raw string and converts to a properly formatted
//	number string. The string is returned via a NumStrDto type.
//	Returned number strings may consist of a leading negative sign
//	('-') numeric digits and may include a decimal separator ('.').
//	The NumStrDto breaks the string down into sign, Integer and
//	Fractional components.
//
//	The numeric separators (decimal separator, thousands separator
//	and currency symbol) taken from the current NumStrDto instance
//	will be copied to the NumStrDto instance returned by this
//	method.
func (nDto *NumStrDto) ParseNumStr(str string) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.ParseNumStr",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {

    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix:  ePrefix.String(),
        ReturnFunc: "",
        ErrContext: "Error: The numeric separators for the current NumStrDto are INVALID!\n" +
          "The returned instance of NumericSeparatorDto ('numSeps') FAILED Validation Tests.",
        ErrMessage: err.Error(),
      }
  }

  return new(numStrDtoQuark).parseNumStr(
    numSeps,
    str,
    ePrefix)
}

// ScaleNumStr - Shifts the position of the decimal point left or right depending
// on the value of input parameter 'scaleMode'.
//
// Input Parameters
// ================
//
// signedNumStr					 string -		A valid Signed Number String
//
// shiftPrecision 	 			 uint -		The number of positions which the decimal point
//
//	will be shifted. If 'shiftPrecision' is Equal to
//	zero, no action will be taken, no error will be
//	issued and the original signedNumStr will be
//	returned.
//
// scaleMode	PrecisionScaleMode -	A constant with one of two Scale Mode values.
//
//	SCALEPRECISIONLEFT - 	Shifts the decimal point
//												from its current position
//												to the left.
//
//	SCALEPRECISIONRIGHT - Shifts the decimal point
//												from its current position
//												to the right.
//
// Note: 	See Methods NumStrDto.ShiftPrecisionRight() and NumStrDto.ShiftPrecisionLeft()
//
//	for additional information.
func (nDto *NumStrDto) ScaleNumStr(signedNumStr string,
  shiftPrecision uint,
  scaleMode PrecisionScaleMode) (NumStrDto, error) {

  ePrefix := "NumStrDto.ScaleNumStr() "

  n2Dto := NumStrDto{}

  var err error

  if scaleMode == SCALEPRECISIONLEFT {

    n2Dto, err = nDto.ShiftPrecisionLeft(signedNumStr, shiftPrecision)

    if err != nil {
      return NumStrDto{},
        fmt.Errorf(ePrefix+
          "Error returned from nDto.ShiftPrecisionLeft(signedNumStr, shiftPrecision) "+
          "signedNumStr='%v' shiftPrecision='%v' scaleMode='%v' Error='%v' ",
          signedNumStr, shiftPrecision, scaleMode.String(), err.Error())

    }

  } else if scaleMode == SCALEPRECISIONRIGHT {

    n2Dto, err = nDto.ShiftPrecisionRight(signedNumStr, shiftPrecision)

    if err != nil {
      return NumStrDto{},
        fmt.Errorf(ePrefix+
          "Error returned from nDto.ShiftPrecisionRight(signedNumStr, shiftPrecision) "+
          "signedNumStr='%v' shiftPrecision='%v' scaleMode='%v' Error='%v' ",
          signedNumStr, shiftPrecision, scaleMode.String(), err.Error())
    }

  } else {

    return NumStrDto{},
      fmt.Errorf(ePrefix+
        "Error! Scale Mode is INVALID! "+
        "Scale Mode is NOT Equal to SCALEPRECISIONLEFT or SCALEPRECISIONRIGHT. scaleMode='%v' ",
        scaleMode.String())

  }

  return n2Dto, nil
}

// SetCurrencySymbol
//
//	Assigns the input parameter rune as the currency symbol to be
//	used by the current instance of NumStrDto.
//
//	In the USA, the currency symbol is the dollar sign ('$').
//
//	If a zero value is submitted as input for the 'currencySymbol'
//	rune, an error will be returned.
//
//	For a list of Major Currency Unicode Symbols, see constants
//	located in: /mathops/mathopsconstants.go
//
//	USA Example
//	===========
//
//	  $123.45
func (nDto *NumStrDto) SetCurrencySymbol(currencySymbol rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetCurrencySymbol",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoElectron).setCurrencySymbol(
    nDto, currencySymbol, ePrefix)
}

// SetDecimalSeparator
//
//	 Assigns a rune or character to the internal data field,
//	 'decimalSeparator'. The Decimal Separator is used to separate
//	 the integer and fractional elements of a number string.
//
//		If a zero value is submitted as input for the 'decimalSeparator'
//		rune, an error will be returned.
//
//		In the USA, the decimal separator is the period character ('.').
//
//		USA Example
//		===========
//
//		123.45892
func (nDto *NumStrDto) SetDecimalSeparator(decimalSeparator rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetDecimalSeparator",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoElectron).setDecimalSeparator(
    nDto, decimalSeparator, ePrefix)
}

// SetThousandsSeparator
//
//	Sets the value of the character which will be
//
// used to separate thousands in the display of the NumStrDto number
// string. In the USA the typical thousands separator is the comma.
//
// If a zero value is submitted, the Thousands Separator will default
// to the comma character.
//
// Example:
// 1,000,000
func (nDto *NumStrDto) SetThousandsSeparator(thousandsSeparator rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetDecimalSeparator",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoElectron).setThousandsSeparator(
    nDto, thousandsSeparator, ePrefix)
}

// ShiftPrecisionLeft - Shifts the relative position of a decimal point within a number
// string. The position of the decimal point is shifted 'shiftPrecision' positions to
// the left of the current decimal point position.
//
// This is equivalent to: result = signedNumStr / 10^precision or signedNumStr divided
// by 10 raised to the power of precision.
//
// See the Examples section below.
//
// Input Parameters
// ================
//
//	signedNumStr		string		- A valid number string. The leading digit may optionally
//															be a '+' or '-' indicating numeric sign value. If '+'
//															or '-'	characters are not present in the first character
//															position, the number is assumed to represent a positive
//															numeric value ('+').
//
//	shiftPrecision		uint		- The number of digits by which the current decimal point
//															position in the number string, 'signedNumStr' will
//															be shifted to the left.
//
// Returns
// =======
//
//	NumStrDto				- If successful, the method returns the result of the Shift Left precision
//										operation in the form of a 'NumStrDto' instance
//
//	error						- If successful, the 'error' type is set to 'nil'. In case of an error,
//										the 'error' instance returned will hold the error message.
//
// Examples
// ========
//
//	                   Shift-Left
//	 signedNumStr			precision				Result
//		 "123456.789"				  3						"123.456789"
//		 "123456.789"				  2						"1234.56789"
//		 "123456.789"	   		  6					  "0.123456789"
//		 "123456789"					6						"123.456789"
//		 "123"								5						"0.00123"
//	  "0"									3						"0.000"
//		 "0.000"							2						"0.00000"
//	 "123456.789"					0						"123456.789"		- zero 'shiftPrecision' has no effect on
//																												original number string
//
// "-123456.789"          0          "-123.456789"
// "-123456.789"          3          "-123.456789"
// "-123456789"						6					 "-123.456789"
func (nDto *NumStrDto) ShiftPrecisionLeft(
  signedNumStr string,
  shiftLeftPrecision uint) (NumStrDto, error) {

  ePrefix := "NumStrDto.ShiftPrecisionLeft() "

  if len(signedNumStr) == 0 {
    return NumStrDto{}, errors.New(ePrefix +
      "Received zero length number string!")
  }

  // Set defaults for thousands separators,
  // decimal separators and currency Symbols
  if nDto.thousandsSeparator == 0 {
    nDto.thousandsSeparator = ','
  }

  if nDto.decimalSeparator == 0 {
    nDto.decimalSeparator = '.'
  }

  if nDto.currencySymbol == 0 {
    nDto.currencySymbol = '$'
  }

  n1, err := new(NumStrDto).NewPtr().ParseNumStr(signedNumStr)

  if err != nil {
    return NumStrDto{}, fmt.Errorf(ePrefix+
      "Received Error from NumStrDto.ParseNumStr(signedNumStr). "+
      "str= '%v' Error= %v",
      signedNumStr, err)
  }

  n2 := new(NumStrDto).New()

  n2.thousandsSeparator = nDto.thousandsSeparator
  n2.decimalSeparator = nDto.decimalSeparator
  n2.currencySymbol = nDto.currencySymbol
  n2.signVal = n1.signVal
  n2.precision = shiftLeftPrecision + n1.precision
  iTotalSpecPrecision := int(n2.precision)
  lenAbsAllNumRunes := len(n1.absAllNumRunes)
  lenAbsIntRunes := n1.GetAbsIntRunesLength()
  lenAbsFracRunes := n1.GetAbsFracRunesLength()

  if nDto.IsNumStrZeroValue(&n1) {

    return nDto.GetZeroNumStrDto(n2.precision), nil
  }

  if iTotalSpecPrecision == lenAbsAllNumRunes {

    n2.absAllNumRunes = append(n2.absAllNumRunes, '0')

  } else if iTotalSpecPrecision > lenAbsAllNumRunes {

    deltaPrecision := iTotalSpecPrecision - lenAbsAllNumRunes + 1

    for i := 0; i < deltaPrecision; i++ {
      n2.absAllNumRunes = append(n2.absAllNumRunes, '0')
    }

  }

  for j := 0; j < lenAbsAllNumRunes; j++ {
    n2.absAllNumRunes = append(n2.absAllNumRunes, n1.absAllNumRunes[j])
  }

  lenAbsAllNumRunes = len(n2.absAllNumRunes)
  lenAbsFracRunes = iTotalSpecPrecision
  lenAbsIntRunes = lenAbsAllNumRunes - lenAbsFracRunes

  if lenAbsIntRunes <= 0 {
    return NumStrDto{}, fmt.Errorf(ePrefix+
      "Calculated number of integer digits is less than or equal to ZERO. "+
      "lenAbsIntRunes= '%v' ",
      lenAbsIntRunes)
  }

  lenAbsFracRunes = n2.GetAbsFracRunesLength()

  err = n2.IsValid(ePrefix)

  if err != nil {
    return NumStrDto{}, err
  }

  return n2, nil
}

// ShiftPrecisionRight - Shifts the existing precision of a number string. The position of
// the decimal point is shifted 'precision' positions to the right.
//
// This is equivalent to: result = signedNumStr X 10^precision or signedNumStr Multiplied
// by 10 raised to the power of precision.
//
// Examples:
// signedNumStr			precision			Result
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
func (nDto *NumStrDto) ShiftPrecisionRight(signedNumStr string, precision uint) (NumStrDto, error) {

  ePrefix := "NumStrDto.ShiftPrecisionRight() "

  if len(signedNumStr) == 0 {
    return NumStrDto{}, errors.New(ePrefix + "Received zero length number string as input!")
  }

  // Set defaults for thousands separators,
  // decimal separators and currency Symbols
  if nDto.thousandsSeparator == 0 {
    nDto.thousandsSeparator = ','
  }

  if nDto.decimalSeparator == 0 {
    nDto.decimalSeparator = '.'
  }

  if nDto.currencySymbol == 0 {
    nDto.currencySymbol = '$'
  }

  n1, err := new(NumStrDto).NewPtr().ParseNumStr(signedNumStr)

  if err != nil {
    return NumStrDto{}, fmt.Errorf(ePrefix+"- Received Error from NumStrDto.ParseNumStr(signedNumStr). str= '%v' Error= %v", signedNumStr, err)
  }

  n2 := new(NumStrDto).New()

  iTotalSpecPrecision := 0
  iPrecision := int(precision)
  iN1Precision := int(n1.precision)

  if iN1Precision > 0 && iPrecision < iN1Precision {
    iTotalSpecPrecision = iN1Precision - iPrecision
  } else {
    iTotalSpecPrecision = 0
  }

  n2.thousandsSeparator = nDto.thousandsSeparator
  n2.decimalSeparator = nDto.decimalSeparator
  n2.currencySymbol = nDto.currencySymbol
  n2.signVal = n1.signVal
  n2.precision = uint(iTotalSpecPrecision)

  lenAbsAllNumRunes := len(n1.absAllNumRunes)

  if nDto.IsNumStrZeroValue(&n1) {

    return nDto.GetZeroNumStrDto(0), nil
  }

  if int(precision) > int(n1.precision) {

    for i := 0; i < lenAbsAllNumRunes; i++ {
      n2.absAllNumRunes = append(n2.absAllNumRunes, n1.absAllNumRunes[i])
    }

    deltaPrecision := int(precision) - int(n1.precision)

    for i := 0; i < deltaPrecision; i++ {
      n2.absAllNumRunes = append(n2.absAllNumRunes, '0')
    }

  } else {

    for i := 0; i < lenAbsAllNumRunes; i++ {
      n2.absAllNumRunes = append(n2.absAllNumRunes, n1.absAllNumRunes[i])
    }

  }

  lenAbsAllNumRunes = len(n2.absAllNumRunes)
  lenAbsFracRunes := iTotalSpecPrecision
  lenAbsIntRunes := lenAbsAllNumRunes - lenAbsFracRunes

  if lenAbsIntRunes <= 0 {
    return NumStrDto{},
      fmt.Errorf(ePrefix+"- Calculated number of integer digits is less than or equal to ZERO. "+
        "lenAbsIntRunes= '%v' ", lenAbsIntRunes)
  }

  lenAbsFracRunes = n2.GetAbsFracRunesLength()

  err = n2.IsValid(ePrefix)

  if err != nil {
    return NumStrDto{}, err
  }

  return n2, nil
}

// SetNumericSeparators
//
//	Used to assign values for the Decimal and Thousands separators
//	as well as the Currency Symbol to be used in displaying number
//	string representations of the numeric value encapsulated by the
//	current NumStrDto instance.
//
//	If zero values are submitted as input for Decimal Separator,
//	Thousands Separator or Currency Symbol, an error will be
//	returned.
//
//	USA Examples
//	============
//
//	Decimal Separator period ('.')    = 123.456
//	Thousands Separator comma (',')   = 1,000,000,000
//	Currency Symbol dollar sign ('$') = $123
func (nDto *NumStrDto) SetNumericSeparators(
  decimalSeparator rune,
  thousandsSeparator rune,
  currencySymbol rune) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetNumericSeparators",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoAtom).setNumericSeparators(
    nDto, decimalSeparator, thousandsSeparator, currencySymbol, ePrefix)
}

// SetNumericSeparatorsDto
//
//	Sets the values of numeric separators:
//
//	  decimal point separator
//	  thousands separator
//	  currency symbol
//
//	Numeric Separator values are transmitted through input
//	parameter 'customSeparators'.
//
//	If any of the values contained in input parameter
//	'customSeparators' is set to zero, an error will be returned.
func (nDto *NumStrDto) SetNumericSeparatorsDto(
  customSeparators NumericSeparatorDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetNumericSeparatorsDto",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoAtom).setNumericSeparatorsDto(
    nDto, customSeparators, ePrefix)
}

// SetNumericSeparatorsToDefaultIfEmpty
//
//	If numeric separators are set to zero or nil, this method will
//	set those numeric separators to the USA defaults. This means
//	that the Decimal separator is set to ('.'), the Thousands
//	separator is set to (',') and the currency symbol is set to
//	'$'.
//
//	If the numeric separators were previously set to a value other
//	than zero or nil, that value is not altered by this method.
//
//	Effectively, this method ensures that numeric separators are
//	set to valid values.
func (nDto *NumStrDto) SetNumericSeparatorsToDefaultIfEmpty() error {

  if nDto.GetDecimalSeparator() == 0 {
    nDto.SetDecimalSeparator('.')
  }

  if nDto.GetThousandsSeparator() == 0 {
    nDto.SetThousandsSeparator(',')
  }

  if nDto.GetCurrencySymbol() == 0 {
    nDto.SetCurrencySymbol('$')
  }

  return nil
}

// SetNumericSeparatorsToUSADefault
//
//	Sets Numeric separatorsto United States of America (USA)
//	defaults.
//
//	  Decimal Point Separator = '.'
//	  Thousands Separator = ','
//	  Currency Symbol = '$'
//
//	Call specific methods to set numeric separators for other
//	countries or cultures:
//
//	  nDto.SetDecimalSeparator()
//	  nDto.SetThousandsSeparator()
//	  nDto.SetCurrencySymbol()
func (nDto *NumStrDto) SetNumericSeparatorsToUSADefault() {
  nDto.SetDecimalSeparator('.')
  nDto.SetThousandsSeparator(',')
  nDto.SetCurrencySymbol('$')
}

// SetNumStr - Sets the value of the current NumStrDto instance
// to the number string received as input.
func (nDto *NumStrDto) SetNumStr(numStr string) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetNumStr",
    "")

  if err != nil {
    return err
  }

  numSeps := nDto.GetNumericSeparatorsDto()

  n2, err := new(NumStrDto).NewNumStr(numStr)

  if err != nil {
    return fmt.Errorf(ePrefix+"Error returned by new(NumStrDto).NewNumStr(numStr). "+
      "numStr='%v' Error='%v' ", numStr, err.Error())
  }

  err = n2.SetNumericSeparatorsDto(numSeps)

  if err != nil {
    return fmt.Errorf(ePrefix+
      "Error returned by n2.SetNumericSeparatorsDto(numSeps) "+
      "Error='%v' \n", err.Error())
  }

  nDto.CopyIn(n2)

  return nil

}

// SetPrecision
//
//		Parses the incoming number string and applies the designated
//		'precision'.
//
//		'precision' determines the number of digits to the right of the
//		decimal place. The boolean parameter 'roundResult' is used to
//		apply rounding in those cases where 'precision' dictates a
//		reduction in the number of digits to the right of the decimal
//		place. See 'Examples' below.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and
//		convert them into numeric values.
//
//		The final NumStrDto result returned by this method will be
//		configured with the Numeric Separators provided by the
//		current instance of 'nDto'. If these Numeric Separators are
//		determined to be invalid, an error will be returned.
//
//		Examples
//		========
//
//		         ------------ Input Parameters ------------
//
//		Example
//		 No       signedNumStr    precision   roundResult     Final Result
//
//		  1       "123456789"         7          false          "123456789.0000000"
//		  2       "123456789"         7          true           "123456789.0000000"
//		  3      "-123456789"         7          false         "-123456789.0000000"
//		  4      "-123456789"         7          true          "-123456789.0000000"
//		  5       "123456.789"        2          true           "123456.79"
//		  6       "123456.789"        2          false          "123456.78"
//		  7       "123456.789"        5          false          "123456.78900"
//		  8       "123.456789"        1          false          "123.4"
//		  9       "123.456789"        1          true           "123.5"
//		 10      "-123.456789"        1          false         "-123.4"
//		 11      "-123.456789"        1          true          "-123.5"
//		 12       "123456.789"        0          true           "123457"
//		 13      "-123456.789"        0          true          "-123457"
//		 14       "123456.789"        0          false          "123456"
//		 15      "-123456.789"        0          false         "-123456"
//		 16       "123457"            1          false          "123457.0"
//		 17       "123457"            1          true           "123457.0"
//		 18      "-123457"            1          false         "-123457.0"
//		 19      "-123457"            1          true          "-123457.0"
//
//		Input Parameters
//		================
//
//		signedNumStr             string
//		  A valid number string
//
//		precision                uint
//		  The 'precision' values designates the number of places to the
//		  right of the decimal point which will be realized upon
//		  completion of this operation.
//
//	   If 'precision' exceeds the maximum limit of 2,147,483,647
//	   an error will be returned
//
//		roundResult              bool
//		  If the 'precision' value is less than the current number of
//		  places to the right of the decimal point, this method will
//		  truncate the existing fractional digits. If 'roundResult' is
//		  set to true, this truncation operation will include rounding
//		  the last digit.
//
//		Return Values
//		=============
//
//		NumStrDto
//		  This returned instance of NumStrDto will contain the result
//		  of the 'set precision' operation described above.
//
//		error
//		  If an error is encountered during processing, this returned
//		  error object will be configured with an appropriate error
//		  message
func (nDto *NumStrDto) SetPrecision(
  signedNumStr string,
  precision uint,
  roundResult bool) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetPrecision",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return NumStrDto{},
      &FuncReturnError{
        ErrPrefix: ePrefix.String(),
        ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
          "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
        ErrContext: "Error extracting Numeric Separators from current NumStrDto instance.",
        ErrMessage: err.Error(),
      }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return NumStrDto{}, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
      ErrContext: "Error: The current instance of NumStrDto ('nDto') is INVALID!\n" +
        "Numeric Separators from 'nDto' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  return new(numStrDtoPhoton).setPrecision(
    numSeps, signedNumStr, precision, roundResult, ePrefix)
}

// SetSignValue
//
//	Sets the sign of the numeric value for the current NumStrDto.
//	Only two values are allowed: +1 and -1. If any other value is
//	passed an error is thrown.
func (nDto *NumStrDto) SetSignValue(newSignVal int) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetSignValue",
    "")

  if err != nil {
    return err
  }

  return new(numStrDtoAtom).setSignValue(
    nDto, true, newSignVal, ePrefix)
}

// SetThisPrecision
//
//	Sets precision for the current NumStrDto instance.
//
//	'precision' identifies the number of decimal places to the
//	 right of the decimal point.
//
//	Numeric Separators
//	==================
//
//	Numeric Separators define the Decimal Separator character,
//	Thousands Separator character, and Currency Symbol character.
//	These separator characters serve two purposes. First they are
//	used to format and display numeric values as number strings.
//	Second, they are also used to parse number strings and
//	convert them into numeric values.
//
//	The Numeric Separators originally configured for the current
//	instance of NumStrDto will remain unchanged. No modifications
//	to Numeric Separators will be made by this method.
//
//	Input Parameters
//	================
//
//	precision                uint
//	  The 'precision' values designates the number of places to the
//	  right of the decimal point which will be realized upon
//	  completion of this operation. The precision operation will be
//	  performed on the number string contained in the NumStrDto
//	  instance.
//
//	roundResult              bool
//	  If the 'precision' value is less than the current number of
//	  places to the right of the decimal point, this method will
//	  truncate the existing fractional digits. If 'roundResult' is
//	  set to 'true', this truncation operation will include
//	  rounding the last digit.
//
//	Return Parameters
//	=================
//
//	error
//	  If an error is encountered during processing, this returned
//	  error object will be configured with an appropriate error
//	  message.
func (nDto *NumStrDto) SetThisPrecision(
  precision uint,
  roundResult bool) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SetThisPrecision",
    "")

  if err != nil {
    return err
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
        "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
      ErrContext: "Error extracting Numeric Separators from current NumStrDto instance.",
      ErrMessage: err.Error(),
    }
  }

  err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

  if err != nil {

    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
      ErrContext: "Error: The current instance of NumStrDto ('nDto') is INVALID!\n" +
        "Numeric Separators from 'nDto' FAILED Validation Tests.",
      ErrMessage: err.Error(),
    }
  }

  err = new(numStrDtoMuon).setPrecisionNumStrDto(
    numSeps, nDto, precision, roundResult, ePrefix)

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(numStrDtoMuon).setPrecisionNumStrDto(\n" +
        "  numSeps, nDto, precision, roundResult, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// Subtract
//
//	Subtracts the value of an input NumStrDto from the current
//	NumStrDto instance.
func (nDto *NumStrDto) Subtract(n2Dto NumStrDto) error {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.Subtract",
    "")

  if err != nil {
    return err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "Error: The current NumStrDto instance is INVALID!",
      ErrMessage: err.Error(),
    }
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
        "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  var differenceNStrDto NumStrDto

  differenceNStrDto, err = new(numStrDtoBoson).subtractNumStrs(
    numSeps, nDto, false, &n2Dto, true, ePrefix.XCpy("nDto - n2Dto"))

  if err != nil {
    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "differenceNStrDto, err = new(numStrDtoBoson).subtractNumStrs(\n" +
        "  numSeps, nDto, false, &n2Dto, true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  // nDto.CopyIn(nResult)
  err = new(numStrDtoMolecule).copy(nDto, &differenceNStrDto, false, ePrefix.XCpy("nResult->n1Dto"))

  if err != nil {

    return &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
        "nDto, &differenceNStrDto, false, ePrefix)",
      ErrContext: "Copying differenceNStrDto->current instance nDto",
      ErrMessage: err.Error(),
    }
  }

  return nil
}

// SubtractNumStrs
//
//		Subtracts the numeric values supplied by two NumStrDto input
//		parameters and returns the result as a new instance of
//		NumStrDto.
//
//	 n1Dto - n2Dto = new returned NumStrDto instance
func (nDto *NumStrDto) SubtractNumStrs(n1Dto NumStrDto, n2Dto NumStrDto) (NumStrDto, error) {

  var ePrefix *ePref.ErrPrefixDto
  var err error

  ePrefix,
    err = ePref.ErrPrefixDto{}.NewIEmpty(
    nil,
    "NumStrDto.SubtractNumStrs",
    "")

  if err != nil {
    return NumStrDto{}, err
  }

  err = new(numStrDtoElectron).isValidNumStrDto(
    nDto, ePrefix.XCpy("Validating 'nDto'"))

  if err != nil {
    return NumStrDto{}, &FuncReturnError{
      ErrPrefix:  ePrefix.String(),
      ReturnFunc: "",
      ErrContext: "Error: The current NumStrDto instance is INVALID!",
      ErrMessage: err.Error(),
    }
  }

  numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
    nDto, ePrefix.XCpy("nDto -> numSeps"))

  if err != nil {
    return NumStrDto{}, &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
        "  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  var differenceNStrDto NumStrDto

  differenceNStrDto, err = new(numStrDtoBoson).subtractNumStrs(
    numSeps, &n1Dto, true, &n2Dto, true, ePrefix.XCpy("n1Dto - n2Dto"))

  if err != nil {
    return NumStrDto{}, &FuncReturnError{
      ErrPrefix: ePrefix.String(),
      ReturnFunc: "differenceNStrDto, err = new(numStrDtoBoson).subtractNumStrs(\n" +
        "  numSeps, &n1Dto, true, &n2Dto, true, ePrefix)",
      ErrContext: "",
      ErrMessage: err.Error(),
    }
  }

  return differenceNStrDto, nil
}
