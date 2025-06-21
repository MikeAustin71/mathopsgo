package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoMechanics struct {
	lock sync.Mutex
}

// addNumStrDto
//
//	Adds the numeric values of input parameters 'n1Dto' and 'n2Dto'. The
//	final result is stored in parameter 'n1Dto'.
func (numStrDtoMech *numStrDtoMechanics) addNumStrDto(
	numSeps NumericSeparatorDto,
	n1Dto *NumStrDto,
	validateN1Dto bool,
	n2Dto *NumStrDto,
	validateN2Dto bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.addNumStrDto()",
		"")

	if err != nil {
		return err
	}

	if n1Dto == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n1Dto'",
		}
	}

	if n2Dto == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n2Dto'",
		}
	}

	if validateN1Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  n1Dto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if validateN2Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  n2Dto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
			ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
				"'numSeps' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	nResult, err := new(numStrDtoBoson).addNumStrs(
		numSeps, n1Dto, false, n2Dto, false, ePrefix.XCpy("n1Dto+n2Dto"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "nResult, err := new(numStrDtoBoson).addNumStrs(\n" +
				"  numSeps, n1Dto, false, n2Dto, false, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(numStrDtoMolecule).copy(n1Dto, &nResult, false, ePrefix.XCpy("nResult->n1Dto"))

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
				"  n1Dto, &nResult, false, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// addNumStrs
//
//		Adds the values represented by two NumStrDto objects and
//		returns the result as a new instance of NumStrDto.
//
//	   n1Dto + n2Dto = Returned NumStrDto
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
//		configured with the Numeric Separators passed by input
//	 	parameter 'numSeps'. If these Numeric Separators are
//		determined to be invalid, an error will be returned.
func (numStrDtoMech *numStrDtoMechanics) addNumStrs(
	numSeps NumericSeparatorDto,
	n1Dto *NumStrDto,
	validateN1Dto bool,
	n2Dto *NumStrDto,
	validateN2Dto bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.addNumStrs()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if n1Dto == nil {

		return NumStrDto{}, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n1Dto'",
		}
	}

	if n2Dto == nil {

		return NumStrDto{}, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'n2Dto'",
		}
	}

	if validateN1Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

		if err != nil {

			return NumStrDto{}, &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  n1Dto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if validateN2Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

		if err != nil {

			return NumStrDto{}, &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  n2Dto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
			ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
				"'numSeps' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	var totalNStrDto NumStrDto

	totalNStrDto, err = new(numStrDtoBoson).addNumStrs(
		numSeps, n1Dto, false, n2Dto, false, ePrefix.XCpy("n1Dto + n2Dto"))

	if err != nil {
		return NumStrDto{}, &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "totalNStrDto, err = new(numStrDtoBoson).addNumStrs(\n" +
				"  numSeps, n1Dto, false, n2Dto, false, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return totalNStrDto, nil
}

// divideNumStrs
//
//	Divides NumStrDto input parameters 'n2Dto'.
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
func (numStrDtoMech *numStrDtoMechanics) divideNumStrs(
	numSeps NumericSeparatorDto,
	dividendDto *NumStrDto,
	validateDividend bool,
	divisorDto *NumStrDto,
	validateDivisor bool,
	minimumPrecision int,
	maximumPrecision int,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.divideNumStrs()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if dividendDto == nil {

		return NumStrDto{}, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'dividendDto'",
		}
	}

	if divisorDto == nil {

		return NumStrDto{}, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'divisorDto'",
		}
	}

	if validateDividend {

		err = new(numStrDtoElectron).isValidNumStrDto(
			dividendDto, ePrefix.XCpy("Validating 'dividendDto'"))

		if err != nil {

			return NumStrDto{}, &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  dividendDto, ePrefix)",
				ErrContext: "Error: Input parameter 'dividendDto' is INVALID!\n" +
					"'dividendDto' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	if validateDivisor {

		err = new(numStrDtoElectron).isValidNumStrDto(
			divisorDto, ePrefix.XCpy("Validating 'divisorDto'"))

		if err != nil {

			return NumStrDto{}, &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  divisorDto, ePrefix)",
				ErrContext: "Error: Input parameter 'divisorDto' is INVALID!\n" +
					"'divisorDto' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
			ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
				"'numSeps' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	dividendNumStr, err := new(numStrDtoAtom).formatNumStr(
		dividendDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "dividendNumStr, err := new(numStrDtoAtom).formatNumStr(\n" +
					"  dividendDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	divisorNumStr, err := new(numStrDtoAtom).formatNumStr(
		divisorDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "divisorNumStr, err := new(numStrDtoAtom).formatNumStr(\n" +
					"  divisorDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iaDividend, err := new(IntAry).NewNumStrWithNumSeps(
		dividendNumStr, numSeps)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "iaDividend, err := new(IntAry).NewNumStrWithNumSeps(\n" +
					"dividendNumStr, numSeps)",
				ErrContext: fmt.Sprintf("dividendNumStr= '%v'", dividendNumStr),
				ErrMessage: err.Error(),
			}
	}

	iaDivisor, err := new(IntAry).NewNumStrWithNumSeps(
		divisorNumStr, numSeps)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaDivisor, err := new(IntAry).NewNumStrWithNumSeps(",
				ErrContext: "divisorNumStr, numSeps)",
				ErrMessage: err.Error(),
			}
	}

	iaQuotient, err := new(IntAryMathDivide).Divide(
		&iaDividend, &iaDivisor, minimumPrecision, maximumPrecision)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "iaQuotient, err = new(IntAryMathDivide).Divide(\n" +
					"  &iaDividend, &iaDivisor, minimumPrecision, maximumPrecision)",
				ErrContext: fmt.Sprintf("iaDividend= '%v'\n iaDivisor= '%v'\n"+
					"minPrecision= '%v'\nmaxPrecision= '%v'",
					dividendNumStr, divisorNumStr, minimumPrecision, maximumPrecision),
				ErrMessage: err.Error(),
			}
	}

	resultNumStrDto, err := iaQuotient.GetNumStrDto()

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "resultNumStrDto, err := iaQuotient.GetNumStrDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if resultNumStrDto.GetPrecision() < minimumPrecision {

		err = resultNumStrDto.SetThisPrecision(uint(minimumPrecision), false)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = resultNumStrDto.SetThisPrecision(uint(minimumPrecision), false)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&resultNumStrDto, ePrefix.XCpy("Validating 'resultNumStrDto'"))

	if err != nil {

		return NumStrDto{}, &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
				"  &resultNumStrDto, ePrefix)",
			ErrContext: "Error: Final Result NumStrDto 'resultNumStrDto' is INVALID!\n" +
				"'resultNumStrDto' FAILED Validation Tests.",
			ErrMessage: err.Error(),
		}
	}

	return resultNumStrDto, nil
}

// formatCurrencyStr
//
//	Formats the current NumStrDto numeric value as a currency
//	string.
//
//	If the Currency Symbol was not previously set for the
//	NumStrDto instance 'nDto', an error will be returned.
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
func (numStrDtoMech *numStrDtoMechanics) formatCurrencyStr(
	nDto *NumStrDto,
	validateNDto bool,
	negValMode NegativeValueFmtMode,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.formatCurrencyStr()",
		"")

	if err != nil {
		return "", err
	}

	if nDto == nil {

		return "", &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'nDto'",
		}
	}

	if validateNDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {

			return "", &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  nDto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
		nDto, ePrefix.XCpy("nDto -> numSeps"))

	if err != nil {
		return "",
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
					"  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if !validateNDto {

		err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
					ErrContext: "Error: The NumStrDto instance ('nDto') is INVALID!\n" +
						"Numeric Separators from 'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	lenAllNumRunes := len(nDto.absAllNumRunes)

	lenOut := lenAllNumRunes

	lenIntRunes := lenAllNumRunes - int(nDto.precision)

	seps := lenIntRunes / 3

	mod := lenIntRunes - (seps * 3)

	if mod == 0 {

		seps--
	}

	// adjust for thousands delimiters
	lenOut += seps

	// adjust for negative sign value
	if nDto.signVal == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {

			lenOut++

		} else {

			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			lenOut += 2
		}
	}

	// adjust for decimal point
	if nDto.precision > 0 {

		lenOut++
	}

	// adjust for currency symbol
	lenOut++

	outRunes := make([]rune, lenOut)

	outIdx := lenOut - 1

	allNumsIdx := lenAllNumRunes - 1

	// If negative value and parenthesis formatting
	// specified, format trailing parenthesis.
	if nDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {

		outRunes[outIdx] = ')'

		outIdx--
	}

	if nDto.precision > 0 {

		for i := 0; i < int(nDto.precision); i++ {

			outRunes[outIdx] = nDto.absAllNumRunes[allNumsIdx]

			outIdx--

			allNumsIdx--
		}

		outRunes[outIdx] = numSeps.DecimalSeparator

		outIdx--
	}

	sepCnt := 0

	for i := 0; i < lenIntRunes; i++ {

		sepCnt++

		if sepCnt == 4 && seps > 0 {

			sepCnt = 1

			seps--

			outRunes[outIdx] = numSeps.ThousandsSeparator

			outIdx--
		}

		outRunes[outIdx] = nDto.absAllNumRunes[allNumsIdx]

		outIdx--

		allNumsIdx--

	}

	outRunes[outIdx] = numSeps.CurrencySymbol

	// If required, add leading negative
	// value sign
	if nDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {

		outRunes[0] = '('

	} else if nDto.signVal == -1 {

		outRunes[0] = '-'
	}

	return string(outRunes), nil
}

// formatThousandsStr
//
//	Returns the number string delimited with the
//	nDto.thousandsSeparator character plus the Decimal Separator if
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
//	If the NumStrDto instance 'nDto' is invalid, an error will be
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
//	  A number string containing the numeric value of the NumStrDto
//	  instance, 'nDto', formatted with thousands separators.
//	    Example:  123,456.78
//
//	error
//	  If a processing error is encountered, this error object will
//	  be returned formatted with an appropriate error message.
func (numStrDtoMech *numStrDtoMechanics) formatThousandsStr(
	nDto *NumStrDto,
	validateNDto bool,
	negValMode NegativeValueFmtMode,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.formatCurrencyStr()",
		"")

	if err != nil {
		return "", err
	}

	if nDto == nil {

		return "", &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'nDto'",
		}
	}

	if validateNDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {

			return "", &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  nDto, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(
		nDto, ePrefix.XCpy("nDto -> numSeps"))

	if err != nil {
		return "",
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numSeps, err := new(numStrDtoAtom).getNumericSeparatorsDto(\n" +
					"  nDto, ePrefix.XCpy(\"nDto -> numSeps\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if !validateNDto {

		err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
					ErrContext: "Error: The NumStrDto instance ('nDto') is INVALID!\n" +
						"Numeric Separators from 'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	lenAllNumRunes := len(nDto.absAllNumRunes)

	lenOut := lenAllNumRunes

	lenIntRunes := lenAllNumRunes - int(nDto.precision)

	seps := lenIntRunes / 3

	mod := lenIntRunes - (seps * 3)

	if mod == 0 {

		seps--
	}

	// adjust for thousands delimiters
	lenOut += seps

	// adjust for negative sign value
	if nDto.signVal == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {

			lenOut++

		} else {

			// MUST BE negValMode == PARENTHESESNEGVALFMTMODE
			lenOut += 2

		}

	}

	// adjust for decimal point
	if nDto.precision > 0 {

		lenOut++
	}

	outRunes := make([]rune, lenOut)

	outIdx := lenOut - 1

	if nDto.signVal == -1 &&
		negValMode == PARENTHESESNEGVALFMTMODE {

		outRunes[outIdx] = ')'

		outIdx--
	}

	allNumsIdx := lenAllNumRunes - 1

	if nDto.precision > 0 {

		for i := 0; i < int(nDto.precision); i++ {

			outRunes[outIdx] = nDto.absAllNumRunes[allNumsIdx]

			outIdx--

			allNumsIdx--
		}

		outRunes[outIdx] = numSeps.DecimalSeparator

		outIdx--
	}

	sepCnt := 0

	for i := 0; i < lenIntRunes; i++ {

		sepCnt++

		if sepCnt == 4 && seps > 0 {

			sepCnt = 1

			seps--

			outRunes[outIdx] = numSeps.ThousandsSeparator

			outIdx--
		}

		outRunes[outIdx] = nDto.absAllNumRunes[allNumsIdx]

		outIdx--

		allNumsIdx--

	}

	if nDto.signVal == -1 {

		if negValMode == LEADMINUSNEGVALFMTMODE {

			outRunes[0] = '-'

		} else {

			outRunes[0] = '('

		}
	}

	return string(outRunes), nil
}

// getBigIntNum
//
//	Converts the numeric value of the current NumStrDto to a
//	type 'BigIntNum' object and returns it to the calling function.
//
//	IMPORTANT
//	=========
//
//	If the current NumStrDto instance is invalid, an error will be
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
//	The Numeric Separators configured for the current instance of
//	NumStrDto will be copied to the BigIntNum object returned by
//	this method.
func (numStrDtoMech *numStrDtoMechanics) getBigIntNum(
	numSeps NumericSeparatorDto,
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.getBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if nDto == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	signedBIntNum, signedBiNumPrecision, err := new(numStrDtoBoson).
		getSignedBigIntPrecision(nDto, false, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "signedBIntNum, signedBiNumPrecision, err := new(numStrDtoBoson).\n" +
					"  getSignedBigIntPrecision(nDto, false, ePrefix)",
				ErrContext: "Failed to convert 'nDto' to Signed *big.Int!",
				ErrMessage: err.Error(),
			}
	}

	bIntNum, err := new(BigIntNum).NewBigIntNumSeps(signedBIntNum, signedBiNumPrecision, numSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bIntNum, err := new(BigIntNum).NewBigIntNumSeps(\n" +
					"  signedBIntNum, signedBiNumPrecision, numSeps)",
				ErrContext: fmt.Sprintf("signedBIntNum= '%v'\n"+
					"signedBiNumPrecision= '%v'\n"+
					"numSeps= '%v'",
					signedBIntNum.Text(10), signedBiNumPrecision, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	return bIntNum, nil
}

// getDecimal
//
//	Converts the current NumStrDto instance to a Type 'Decimal' and
//	returns it to the calling function.
//
//	The returned Decimal instance will contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from the NumStrDto instance 'nDto'.
func (numStrDtoMech *numStrDtoMechanics) getDecimal(
	numSeps NumericSeparatorDto,
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (Decimal, error) {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.getDecimal",
		"")

	if err != nil {
		return Decimal{}, err
	}

	if nDto == nil {

		return Decimal{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return Decimal{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	numStr, err := new(numStrDtoAtom).formatNumStr(
		nDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numStr, err :=new(numStrDtoAtom).formatNumStr(\n" +
					"  nDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	dec, err := new(Decimal).NewNumStrWithNumSeps(numStr, numSeps)

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "dec, err := new(Decimal).NewNumStrWithNumSeps(\n" +
					"  numStr, numSeps)",
				ErrContext: fmt.Sprintf("numStr= '%v'\n"+
					"numSeps= '%v'", numStr, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	return dec, nil
}

// getIntAry
//
//	Converts the current NumStrDto instance to a Type 'IntAry' and
//	returns it to the calling function.
//
//	The returned IntAry instance will contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from the NumStrDto instance, 'nDto'.
func (numStrDtoMech *numStrDtoMechanics) getIntAry(
	numSeps NumericSeparatorDto,
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	numStrDtoMech.lock.Lock()

	defer numStrDtoMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMechanics.getIntAry",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if nDto == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return IntAry{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	err = numSeps.IsValid(ePrefix.XCpy("Validating 'numSeps'").String())

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	numStr, err := new(numStrDtoAtom).formatNumStr(
		nDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numStr, err :=new(numStrDtoAtom).formatNumStr(\n" +
					"  nDto, false, LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	intAry, err := new(IntAry).NewNumStrWithNumSeps(numStr, numSeps)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "intAry, err := new(IntAry).NewNumStrWithNumSeps(\n" +
					"  numStr, numSeps)",
				ErrContext: fmt.Sprintf("numStr= '%v'\n"+
					"numSeps= '%v'", numStr, numSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	return intAry, nil
}
