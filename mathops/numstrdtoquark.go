package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type numStrDtoQuark struct {
	lock sync.Mutex
}

// parseBigIntNum
//
//		Receives a BigIntNum instance ('biNum') and coverts it to a
//		NumStrDto instance which is returned to the calling function.
//
//	  If 'biNum' proves to be invalid, an error will be returned.
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
//		configured with the Numeric Separators copied from the input
//	 parameter 'biNum'.
func (nStrDtoQuark *numStrDtoQuark) parseBigIntNum(
	numSeps NumericSeparatorDto,
	biNum *BigIntNum,
	validateBiNum bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoQuark.parseBigIntNum()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if biNum == nil {

		return NumStrDto{}, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'biNum'",
		}
	}

	if validateBiNum {

		err = new(bigIntNumAtom).isBigIntNumValid(
			biNum,
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

	n2Dto := NumStrDto{}

	n2Dto.signVal = biNum.sign

	n2Dto.precision = biNum.precision

	scratchNum := big.NewInt(0).Set(biNum.bigInt)

	if n2Dto.signVal < 0 {
		scratchNum.Neg(scratchNum)
	}

	bigZero := big.NewInt(0)

	bigTen := big.NewInt(int64(10))

	modulo := big.NewInt(0)

	modX := big.NewInt(0)

	n2Dto.absAllNumRunes = make([]rune, 0)

	if scratchNum.Cmp(bigZero) == 0 {

		n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')

	} else {

		for scratchNum.Cmp(bigZero) == 1 {

			scratchNum, modulo = big.NewInt(0).QuoRem(scratchNum, bigTen, modX)

			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, rune(modulo.Int64()+int64(48)))
		}
	}

	lenAllNumRunes := len(n2Dto.absAllNumRunes)

	if int(n2Dto.precision) >= lenAllNumRunes {

		deltaNumRunes := int(n2Dto.precision) - lenAllNumRunes + 1

		for k := 0; k < deltaNumRunes; k++ {

			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')

			lenAllNumRunes++
		}

	}

	tRune := rune(0)

	if lenAllNumRunes > 1 {

		xLen := lenAllNumRunes - 1

		sortLimit := xLen / 2

		yCnt := 0

		for i := xLen; i > sortLimit; i-- {

			tRune = n2Dto.absAllNumRunes[yCnt]

			n2Dto.absAllNumRunes[yCnt] = n2Dto.absAllNumRunes[i]

			n2Dto.absAllNumRunes[i] = tRune

			yCnt++
		}
	}

	err = new(numStrDtoAtom).setNumericSeparatorsDto(
		&n2Dto, numSeps, ePrefix.XCpy("numSeps->n2Dto"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoAtom).setNumericSeparatorsDto(\n" +
					"&n2Dto, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2Dto, ePrefix.XCpy("Validating Final Result 'n2Dto'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err =new(numStrDtoElectron).isValidNumStrDto(\n" +
					"&n2Dto, ePrefix)",
				ErrContext: "Error: The final calculated result ('n2Dto') is INVALID!\n" +
					"'n2Dto' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2Dto, nil
}

// parseNumStr
//
//	Receives a raw string and converts it to a properly formatted
//	number string. The string is returned via a NumStrDto type.
//	Returned number strings may consist of a leading negative sign
//	('-') numeric digits and may include a decimal separator ('.').
//	The NumStrDto breaks the string down into sign, Integer and
//	Fractional components.
//
//	The numeric separators (decimal separator, thousands separator
//	and currency symbol) are taken from the input parameter,
//	'numSeps'. If the NumericSeparatorDto object ('numSeps') is
//	invalid, an error will be returned.
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

// parseSignedBigInt
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
//	  signedBigInt    precision           result
//
//	    946254            3               946.254
//	    946254            0               946254
//	   -946254            3              -946.254
//	   -946254            0              -946254
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
func (nStrDtoQuark *numStrDtoQuark) parseSignedBigInt(
	numSeps NumericSeparatorDto,
	signedBigInt *big.Int,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoQuark.lock.Lock()

	defer nStrDtoQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoQuark.parseSignedBigInt()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if signedBigInt == nil {

		return NumStrDto{}, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'signedBigInt'",
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

	if new(MathProcessUtility).DoesUintExceedMax32BitInt(precision) {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Error: Input parameter 'precision' is INVALID!\n" +
					"'precision' Exceeds the maximum allowable limt of 2,147,483,647.\n" +
					fmt.Sprintf("precision= '%v'", precision),
			}
	}

	n2Dto := NumStrDto{}

	n2Dto.precision = precision

	scratchNum := big.NewInt(0).Set(signedBigInt)

	bigZero := big.NewInt(0)

	n2Dto.signVal = 1

	if scratchNum.Cmp(bigZero) == -1 {

		scratchNum.Neg(scratchNum)

		n2Dto.signVal = -1
	}

	bigTen := big.NewInt(int64(10))

	modulo := big.NewInt(0)

	n2Dto.absAllNumRunes = make([]rune, 0, 100)

	if scratchNum.Cmp(bigZero) == 0 {

		n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')

	} else {

		for scratchNum.Cmp(bigZero) == 1 {

			modulo = big.NewInt(0).Rem(scratchNum, bigTen)

			scratchNum = big.NewInt(0).Quo(scratchNum, bigTen)

			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, rune(modulo.Int64()+int64(48)))
		}
	}

	lenAllNumRunes := len(n2Dto.absAllNumRunes)

	if int(n2Dto.precision) >= lenAllNumRunes {

		deltaNumRunes := int(n2Dto.precision) - lenAllNumRunes + 1

		for k := 0; k < deltaNumRunes; k++ {

			n2Dto.absAllNumRunes = append(n2Dto.absAllNumRunes, '0')

			lenAllNumRunes++
		}
	}

	tRune := rune(0)

	if lenAllNumRunes > 1 {

		xLen := lenAllNumRunes - 1

		sortLimit := xLen / 2

		yCnt := 0

		for i := xLen; i > sortLimit; i-- {

			tRune = n2Dto.absAllNumRunes[yCnt]

			n2Dto.absAllNumRunes[yCnt] = n2Dto.absAllNumRunes[i]

			n2Dto.absAllNumRunes[i] = tRune

			yCnt++
		}
	}

	err = new(numStrDtoAtom).setNumericSeparatorsDto(
		&n2Dto, numSeps, ePrefix.XCpy("numSeps->n2Dto"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoAtom).setNumericSeparatorsDto(\n" +
					"&n2Dto, numSeps, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&n2Dto, ePrefix.XCpy("Validating Final Result 'n2Dto'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err =new(numStrDtoElectron).isValidNumStrDto(\n" +
					"&n2Dto, ePrefix)",
				ErrContext: "Error: The final calculated result ('n2Dto') is INVALID!\n" +
					"'n2Dto' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	return n2Dto, nil
}
