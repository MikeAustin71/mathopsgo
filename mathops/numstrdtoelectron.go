package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoElectron struct {
	lock sync.Mutex
}

func (nStrDtoElectron *numStrDtoElectron) emptyNumStrDto(
	numStrDto *NumStrDto) {

	if numStrDto == nil {
		return
	}

	numStrDto.signVal = 0
	numStrDto.absAllNumRunes = []rune{}
	numStrDto.precision = 0
	numStrDto.decimalSeparator = 0
	numStrDto.thousandsSeparator = 0
	numStrDto.currencySymbol = 0

	return
}

// isValidNumStrDto
//
//	Performs a diagnostic review of the current NumStrDto instance
//	and returns 'nil' if the NumStrDto object is valid in all
//	respects.
//
//	If the NumStrDto instance is judged invalid, an error message
//	is returned.
func (nStrDtoElectron *numStrDtoElectron) isValidNumStrDto(
	numStrDto *NumStrDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoElectron.isValidNumStrDto()",
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

	if numStrDto.thousandsSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numStrDto.thousandsSeparator == 0",
			ErrMessage: "Error: Thousands Separator cannot be '0'\n" +
				"NumStrDto Numeric Separators are Invalid!",
		}
	}

	if numStrDto.decimalSeparator == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numStrDto.decimalSeparator == 0",
			ErrMessage: "Error: Decimal Separator cannot be '0'\n" +
				"NumStrDto Numeric Separators are Invalid!",
		}
	}

	if numStrDto.currencySymbol == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numStrDto.currencySymbol == 0",
			ErrMessage: "Error: Currency Separator cannot be '0'\n" +
				"NumStrDto Numeric Separators are Invalid!",
		}
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: " len(numStrDto.absAllNumRunes) == 0",
			ErrMessage: "Error: Number string is a ZERO length array!\n" +
				"NumStrDto object is invalid.",
		}
	}

	if int(numStrDto.precision) >= lenAbsAllNumRunes {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "precision > number string length",
			ErrMessage: "Error: 'precision' does not correlate with number string.\n" +
				"NumStrDto object is invalid!",
		}
	}

	if numStrDto.signVal != 1 && numStrDto.signVal != -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numStrDto.signVal != 1 && numStrDto.signVal != -1",
			ErrMessage: "Error: sign Value is INVALID. Should be +1 or -1.\n" +
				"NumStrDto object is invalid!",
		}
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if numStrDto.absAllNumRunes[i] < '0' || numStrDto.absAllNumRunes[i] > '9' {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "Number string character < 0 Or > 9",
				ErrMessage: "Error: Non-Numeric character found in number string!\n" +
					"NumStrDto object is invalid!",
			}
		}
	}

	return nil
}

// isNumStrZeroValue
//
//	Returns 'true' if all the digits in the number string for the
//	current NumStrDto instance are zero.
func (nStrDtoElectron *numStrDtoElectron) isNumStrZeroValue(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

	nStrDtoElectron.lock.Lock()

	defer nStrDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoElectron.isNumStrZeroValue()",
		"")

	if err != nil {
		return true, err
	}

	if numStrDto == nil {

		return true,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	lenAbsAllNumRunes := len(numStrDto.absAllNumRunes)

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return true,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  numStrDto, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	} else {

		if lenAbsAllNumRunes == 0 {

			return true,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "len(numStrDto.absAllNumRunes) == 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The internal rune array of numeric characters is empty.\n" +
						"This instance of NumStrDto is either corrupted or uninitialized!",
				}
		}

		precision := int(numStrDto.precision)

		if precision < 0 {

			return true,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "precision < 0",
					ErrMessage: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is less than zero.",
				}
		}

		if precision > lenAbsAllNumRunes {

			return true,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: The NumStrDto object is INVALID!\n" +
						"The 'precision' value is greater than the internal numeric digits array.",
					ErrMessage: "",
				}
		}
	}

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if numStrDto.absAllNumRunes[i] != '0' {

			return false, nil

		}
	}

	return true, nil
}
