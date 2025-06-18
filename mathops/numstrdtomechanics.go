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
