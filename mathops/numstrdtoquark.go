package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrDtoQuark struct {
	lock sync.Mutex
}

// parseNumStr
//
//		Receives a raw string and converts it to a properly formatted
//		number string. The string is returned via a NumStrDto type.
//		Returned number strings may consist of a leading negative sign
//		('-') numeric digits and may include a decimal separator ('.').
//		The NumStrDto breaks the string down into sign, Integer and
//		Fractional components.
//
//		The numeric separators (decimal separator, thousands separator
//		and currency symbol) are taken from the input parameter,
//		'numSeps'. If the NumericSeparatorDto object ('numSeps') is
//	 invalid, an error will be returned.
func (nStrDtoQuark *numStrDtoQuark) parseNumStr(
	numSeps NumericSeparatorDto,
	str string,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoQuark.parseNumStr()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if len(str) == 0 {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'str' is INVALID!\n" +
					"'str' (string) is empty with a zero length.",
			}
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

	nStrMolecule := new(numStrDtoMolecule)

	n2Dto := nStrMolecule.newZeroNumStrDto(numSeps, 0)

	n2Dto.signVal = 1

	baseRunes := []rune(str)

	lBaseRunes := len(baseRunes)

	isStartRunes := false

	isEndRunes := false

	isMinusSignFound := false

	//lCurRunes := len(NumStrCurrencySymbols)
	//isSkip := false
	isFractionalValue := false

	var absFracRunes []rune
	var absIntRunes []rune

	for i := 0; i < lBaseRunes && isEndRunes == false; i++ {

		if baseRunes[i] != '-' &&
			baseRunes[i] != n2Dto.decimalSeparator &&
			(baseRunes[i] < '0' || baseRunes[i] > '9') {

			continue

		} else if baseRunes[i] == '-' &&
			isMinusSignFound == false &&
			isStartRunes == false &&
			i+1 < lBaseRunes &&
			((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
				baseRunes[i+1] == n2Dto.decimalSeparator) {

			isMinusSignFound = true
			n2Dto.signVal = -1
			isStartRunes = true
			continue

		} else if baseRunes[i] >= '0' && baseRunes[i] <= '9' {

			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, baseRunes[i])
			isStartRunes = true

			if isFractionalValue {
				absFracRunes = append(absFracRunes, baseRunes[i])
			} else {
				absIntRunes = append(absIntRunes, baseRunes[i])
			}

		} else if i+1 < lBaseRunes &&
			baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
			baseRunes[i] == n2Dto.decimalSeparator {

			isFractionalValue = true
			continue

		}

		if i == lBaseRunes-1 {

			isEndRunes = true

		}

	}

	// Original Code
	//for i := 0; i < lBaseRunes && isEndRunes == false; i++ {
	//
	//  if baseRunes[i] != '-' &&
	//    baseRunes[i] != n2Dto.decimalSeparator &&
	//    (baseRunes[i] < '0' || baseRunes[i] > '9') {
	//
	//    continue
	//
	//  } else if baseRunes[i] == '-' &&
	//    isMinusSignFound == false &&
	//    isStartRunes == false && isEndRunes == false &&
	//    i+1 < lBaseRunes &&
	//    ((baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9') ||
	//      baseRunes[i+1] == n2Dto.decimalSeparator) {
	//
	//    isMinusSignFound = true
	//    n2Dto.signVal = -1
	//    isStartRunes = true
	//    continue
	//
	//  } else if isEndRunes == false &&
	//    baseRunes[i] >= '0' && baseRunes[i] <= '9' {
	//
	//    n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, baseRunes[i])
	//    isStartRunes = true
	//
	//    if isFractionalValue {
	//      absFracRunes = append(absFracRunes, baseRunes[i])
	//    } else {
	//      absIntRunes = append(absIntRunes, baseRunes[i])
	//    }
	//
	//  } else if isEndRunes == false &&
	//    i+1 < lBaseRunes &&
	//    baseRunes[i+1] >= '0' && baseRunes[i+1] <= '9' &&
	//    baseRunes[i] == n2Dto.decimalSeparator {
	//
	//    isFractionalValue = true
	//    continue
	//
	//  }
	//
	//  if i == lBaseRunes-1 {
	//
	//    isEndRunes = true
	//
	//  }
	//
	//}

	lenAbsAllNumRunes := len(n2Dto.absAllNumRunes)

	if lenAbsAllNumRunes == 0 {

		n2Dto = nStrMolecule.newZeroNumStrDto(numSeps, 0)

		return n2Dto, nil
	}

	lenAbsIntNumRunes := len(absIntRunes)

	if lenAbsIntNumRunes == 0 {

		absIntRunes = append(absIntRunes, '0')
	}

	lenAbsAllNumRunes = len(n2Dto.absAllNumRunes)

	lenAbsIntNumRunes = len(absIntRunes)

	lenAbsFracNumRunes := len(absFracRunes)

	isZeroVal := true

	for i := 0; i < lenAbsAllNumRunes; i++ {

		if n2Dto.absAllNumRunes[i] != '0' {

			isZeroVal = false
		}
	}

	if isZeroVal {

		n2Dto = nStrMolecule.newZeroNumStrDto(numSeps, uint(lenAbsFracNumRunes))
		//nZeroDto := nDto.GetZeroNumStrDto(uint(lenAbsFracNumRunes))
		return n2Dto, nil
	}

	if isFractionalValue {
		n2Dto.precision = uint(len(absFracRunes))
	}

	if lenAbsAllNumRunes != lenAbsIntNumRunes+lenAbsFracNumRunes {

		n2Dto.absAllNumRunes = []rune{}

		newLenAbsAllNumRunes := lenAbsIntNumRunes + lenAbsFracNumRunes

		for i := 0; i < newLenAbsAllNumRunes; i++ {

			if i < lenAbsIntNumRunes {

				n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, absIntRunes[i])

			} else {

				n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, absFracRunes[i-lenAbsIntNumRunes])
			}
		}

		lenAbsAllNumRunes = len(n2Dto.absAllNumRunes)
	}

	// Validate n2Dto object

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2Dto, ePrefix.XCpy("Validating 'n2Dto' Result"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &n2Dto, ePrefix.XCpy(\"Validating 'n2Dto' Result\"))",
				ErrContext: "Error: Calculated result 'n2Dto' is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	return n2Dto, nil

}
