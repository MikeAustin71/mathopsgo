package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoAtom struct {
	lock sync.Mutex
}

// compareAbsoluteValues
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
func (nStrDtoAtom *numStrDtoAtom) compareAbsoluteValues(
	n1Dto *NumStrDto,
	validateN1Dto bool,
	n2Dto *NumStrDto,
	validateN2Dto bool,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoAtom.compareAbsoluteValues()",
		"")

	if err != nil {
		return 0, err
	}

	if n1Dto == nil {

		return 0, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n1Dto'",
		}
	}

	if n2Dto == nil {

		return 0, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n2Dto'",
		}
	}

	nStrElectron := new(numStrDtoElectron)

	if validateN1Dto {

		err = nStrElectron.isValidNumStrDto(
			n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

		if err != nil {

			return 0,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n1Dto, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateN2Dto {

		err = nStrElectron.isValidNumStrDto(
			n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

		if err != nil {

			return 0,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n2Dto, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	nStrGluon := new(numStrDtoGluon)

	// n1DtoAbsFracRunes := n1Dto.GetAbsFracRunes()
	n1DtoAbsFracRunes, err := nStrGluon.getAbsFracRunes(
		n1Dto, false, ePrefix.XCpy("n1Dto"))

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n1DtoAbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunes(\n" +
					"  n1Dto, validateNumStrDto=false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	//n2DtoAbsFracRunes := n2Dto.GetAbsFracRunes()
	n2DtoAbsFracRunes, err := nStrGluon.getAbsFracRunes(
		n2Dto, false, ePrefix.XCpy("n2Dto"))

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2DtoAbsFracRunes, err := new(numStrDtoGluon).getAbsFracRunes(\n" +
					"  n2Dto, validateNumStrDto=false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	//n1DtoAbsIntRunes := n1Dto.GetAbsIntRunes()
	n1DtoAbsIntRunes, err := nStrGluon.getAbsIntRunes(
		n1Dto, false, ePrefix.XCpy("n1Dto"))

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n1DtoAbsIntRunes, err := new(numStrDtoGluon).getAbsIntRunes(\n" +
					"  n1Dto, validateNumStrDto=false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	//n2DtoAbsIntRunes := n2Dto.GetAbsIntRunes()
	n2DtoAbsIntRunes, err := nStrGluon.getAbsIntRunes(
		n2Dto, false, ePrefix.XCpy("n2Dto"))

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2DtoAbsIntRunes, err := new(numStrDtoGluon).getAbsIntRunes(\n" +
					"  n2Dto, validateNumStrDto=false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	lenN1IntRunes := len(n1DtoAbsIntRunes)

	lenN2IntRunes := len(n2DtoAbsIntRunes)

	//isN1Zero := nDto.IsNumStrZeroValue(n1Dto)

	isN1Zero, err := nStrElectron.isNumStrZeroValue(
		n1Dto, false, ePrefix.XCpy("'n1Dto'"))

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "isN1Zero, err := nStrElectron.isNumStrZeroValue(\n" +
					"  n1Dto, validateNumStrDto=false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	//isN2Zero := nDto.IsNumStrZeroValue(n2Dto)

	isN2Zero, err := nStrElectron.isNumStrZeroValue(
		n2Dto, false, ePrefix.XCpy("'n2Dto'"))

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "isN2Zero, err := nStrElectron.isNumStrZeroValue(\n" +
					"  n2Dto, validateNumStrDto=false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if !isN1Zero && isN2Zero {
		return 1, nil
	}

	if isN1Zero && !isN2Zero {
		return -1, nil
	}

	//if isN1Zero && isN2Zero {
	//	return 0, nil
	//}

	// isN1Zero = isN2Zero = 0
	if isN1Zero {
		return 0, nil
	}

	if lenN1IntRunes > lenN2IntRunes {
		return 1, nil
	}

	if lenN1IntRunes < lenN2IntRunes {
		return -1, nil
	}

	// lenN1IntRunes Must Be Equal to lenN2IntRunes

	for i := 0; i < lenN1IntRunes; i++ {
		n1 := n1DtoAbsIntRunes[i] - 48
		n2 := n2DtoAbsIntRunes[i] - 48

		if n1 > n2 {
			return 1, nil
		}

		if n1 < n2 {
			return -1, nil
		}
	}

	// All the integers are equal
	lenN1FracRunes := len(n1DtoAbsFracRunes)

	lenN2FracRunes := len(n2DtoAbsFracRunes)

	lenFracRunesToTest := lenN1FracRunes

	if lenN2FracRunes < lenN1FracRunes {
		lenFracRunesToTest = lenN2FracRunes
	}

	for j := 0; j < lenFracRunesToTest; j++ {

		n1 := n1DtoAbsFracRunes[j] - 48

		n2 := n2DtoAbsFracRunes[j] - 48

		if n1 > n2 {
			return 1, nil
		}

		if n1 < n2 {
			return -1, nil
		}

	}

	if lenN1FracRunes > lenN2FracRunes {
		return 1, nil
	}

	if lenN1FracRunes < lenN2FracRunes {
		return -1, nil
	}

	return 0, nil
}

// formatNumStr
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
//	Make certain that the decimal separator has been properly
//	configured for the current NumStrDto instance BEFORE you
//	call this method. Note: The USA decimal separator is the
//	period '.'. Decimal separators are used to separate integer
//	and fractional components of a numeric value. If the decimal
//	separator rune is set to zero, an error will be returned.
//
//	Input Parameters
//	================
//
//	negValMode               NegativeValueFmtMode
//	  Specifies the display mode for negative values:
//
//	  LEADMINUSNEGVALFMTMODE    - Negative values formatted with
//	                              a leading minus sign.
//	                              Example: -123456.78
//
//	  PARENTHESESNEGVALFMTMODE  - Negative values formatted with
//	                              surrounding parentheses.
//	                              Example: (123456.78)
func (nStrDtoAtom *numStrDtoAtom) formatNumStr(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	negValMode NegativeValueFmtMode,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoAtom.getNumericSeparatorsDto()",
		"")

	if err != nil {
		return "", err
	}

	if numStrDto == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  numStrDto, ePrefix)",
					ErrContext: "Error: Input parameter 'numStrDto' is INVALID!",
					ErrMessage: err.Error(),
				}
		}
	} else {

		if numStrDto.decimalSeparator == 0 {

			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "numStrDto.decimalSeparator == 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is greater than the internal numeric digits array.",
				}
		}

	}

	lenAllNumRunes := len(numStrDto.absAllNumRunes)

	lenOut := lenAllNumRunes

	lenIntRunes := lenAllNumRunes - int(numStrDto.precision)

	// adjust for negative sign value
	if numStrDto.signVal == -1 {
		if negValMode == LEADMINUSNEGVALFMTMODE {
			lenOut++
		} else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			lenOut += 2
		}

	}

	// adjust for decimal point
	if numStrDto.precision > 0 {
		lenOut++
	}

	outRunes := make([]rune, lenOut)
	outIdx := lenOut - 1

	if numStrDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {
		outRunes[outIdx] = ')'
		outIdx--
	}

	allNumsIdx := lenAllNumRunes - 1

	if numStrDto.precision > 0 {

		for i := 0; i < int(numStrDto.precision); i++ {
			outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
			outIdx--
			allNumsIdx--
		}

		outRunes[outIdx] = numStrDto.decimalSeparator
		outIdx--
	}

	for i := 0; i < lenIntRunes; i++ {

		outRunes[outIdx] = numStrDto.absAllNumRunes[allNumsIdx]
		outIdx--
		allNumsIdx--

	}

	if numStrDto.signVal == -1 {
		if negValMode == LEADMINUSNEGVALFMTMODE {
			outRunes[0] = '-'
		} else {
			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			outRunes[0] = '('
		}

	}

	return string(outRunes), nil
}

// getNumericSeparatorsDto - Returns a structure containing the
// character or rune values for decimal point separator, thousands
// separator and currency symbol.
func (nStrDtoAtom *numStrDtoAtom) getNumericSeparatorsDto(
	numStrDto *NumStrDto,
	errPrefDto *ePref.ErrPrefixDto) (NumericSeparatorDto, error) {

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoAtom.getNumericSeparatorsDto()",
		"")

	if err != nil {
		return NumericSeparatorDto{}, err
	}

	if numStrDto == nil {

		return NumericSeparatorDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	numSeps := NumericSeparatorDto{}
	numSeps.DecimalSeparator = numStrDto.decimalSeparator
	numSeps.ThousandsSeparator = numStrDto.thousandsSeparator
	numSeps.CurrencySymbol = numStrDto.currencySymbol

	return numSeps, nil
}

// setNumericSeparatorsToDefaultIfEmpty
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
func (nStrDtoAtom *numStrDtoAtom) setNumericSeparatorsToDefaultIfEmpty(
	numStrDto *NumStrDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoAtom.setNumericSeparatorsToDefaultIfEmpty()",
		"")

	if err != nil {
		return err
	}

	if numStrDto == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'numStrDto'",
		}
	}

	if numStrDto.GetDecimalSeparator() == 0 {
		numStrDto.SetDecimalSeparator('.')
	}

	if numStrDto.GetThousandsSeparator() == 0 {
		numStrDto.SetThousandsSeparator(',')
	}

	if numStrDto.GetCurrencySymbol() == 0 {
		numStrDto.SetCurrencySymbol('$')
	}

	return nil
}

// setNumericSeparatorsToUSADefault
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
func (nStrDtoAtom *numStrDtoAtom) setNumericSeparatorsToUSADefault(
	numStrDto *NumStrDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoAtom.setNumericSeparatorsToUSADefault()",
		"")

	if err != nil {
		return err
	}

	if numStrDto == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'numStrDto'",
		}
	}

	numStrDto.decimalSeparator = '.'

	numStrDto.thousandsSeparator = ','

	numStrDto.currencySymbol = '$'

	return nil
}

// setSignValue
//
//	Sets the sign of the numeric value for the NumStrDto parameter
//	'numStrDto'.
//
//	Only two sign values are allowed: +1 and -1. If any other value
//	is passed an error is returned.
func (nStrDtoAtom *numStrDtoAtom) setSignValue(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	newSignVal int,
	errPrefDto *ePref.ErrPrefixDto) error {

	nStrDtoAtom.lock.Lock()

	defer nStrDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoAtom.setSignValue()",
		"")

	if err != nil {
		return err
	}

	if numStrDto == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'numStrDto'",
		}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  numStrDto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if newSignVal != -1 && newSignVal != 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "newSignVal != -1 && newSignVal != 1",
			ErrMessage: "Error: Input parameter 'newSignVal' is INVALID!\n" +
				"'newSignVal' must be either - or 1 ",
		}
	}

	numStrDto.signVal = newSignVal

	return nil
}
