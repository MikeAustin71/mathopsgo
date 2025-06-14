package mathops

import (
	"errors"
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoMuon struct {
	lock sync.Mutex
}

// findNumStrSignificantDigitLimits
//
//	Analyzes an array of characters which constitute a number
//	string are returns the significant digits in a new instance of
//	NumStrDto.
//
//	Example
//	=======
//
//	absAllRunes  precision  signVal  Result
//
//	001236700        4         1     123.67
//	000006700        4         1       0.67
//	001230000        4         1     123.0
func (nStrDtoMuon *numStrDtoMuon) findNumStrSignificantDigitLimits(
	numStrDto *NumStrDto,
	validateNumStrDto bool,
	absAllRunes []rune,
	precision uint,
	signVal int,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoMuon.lock.Lock()

	defer nStrDtoMuon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoMuon.findNumStrSignificantDigitLimits()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if numStrDto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numStrDto'",
			}
	}

	lenAbsAllRunes := len(absAllRunes)

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			numStrDto, ePrefix.XCpy("Validating 'numStrDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  numStrDto, ePrefix)",
					ErrContext: "Error: Input parameter 'numStrDto' is INVALID!",
					ErrMessage: err.Error(),
				}
		}
	} else {

		if signVal != -1 && signVal != 1 {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "",
					ErrMessage: "Error: Input parameter 'signVal' is INVALID!\n" +
						"signVal must be either -1 or 1\n" +
						fmt.Sprintf("signVal=%d", signVal),
				}
		}

		if lenAbsAllRunes < 1 {
			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "len(absAllRunes) == 0",
					ErrMessage: "Error: Input parameter 'absAllRunes' is INVALID!\n" +
						"'absAllRunes' has ZERO length.",
				}
		}

		if precision > uint(lenAbsAllRunes) {
			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "precision > uint(lenAbsAllRunes)",
					ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
						"'precision' value is greater than length of rune array.",
				}
		}

	}

	iPrecision := int(precision)
	firstIntIdx := -1
	lastIntIdx := -1
	lastFracIdx := -1

	isFractional := false

	if iPrecision > 0 {
		isFractional = true
	}

	lenAbsFracRunes := iPrecision

	lenAbsIntRunes := lenAbsAllRunes - lenAbsFracRunes

	for i := 0; i < lenAbsAllRunes; i++ {

		if i < lenAbsIntRunes {

			if firstIntIdx == -1 && absAllRunes[i] > '0' && absAllRunes[i] <= '9' {
				firstIntIdx = i
			}

			lastIntIdx = i
		}

		if isFractional && i >= lenAbsIntRunes && absAllRunes[i] > '0' && absAllRunes[i] <= '9' {

			lastFracIdx = i
		}

	}

	if firstIntIdx == -1 {

		firstIntIdx = lastIntIdx
	}

	if isFractional && lastFracIdx == -1 {

		lastFracIdx = lenAbsIntRunes
	}

	numStrOut := ""

	if signVal < 0 {

		numStrOut = "-"
	}

	numStrOut += string(absAllRunes[firstIntIdx : lastIntIdx+1])

	if isFractional {

		numStrOut += string(numStrDto.decimalSeparator)

		numStrOut += string(absAllRunes[lastIntIdx+1 : lastFracIdx+1])
	}

	nOutDto, err := nDto.ParseNumStr(numStrOut)

	if err != nil {
		return NumStrDto{}, fmt.Errorf("FindSignificantDigitLimits() - Error retuned from nDto.ParseNumStr(numStrOut). numStrOut= '%v' Error= %v", numStrOut, err)
	}

	return nOutDto, nil
}
