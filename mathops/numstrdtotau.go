package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"sync"
)

type numStrDtoTau struct {
	lock sync.Mutex
}

// lowLevelAddition
//
//	 Performs addition on two NumStrDto values and returns the
//	 total value as a new instance of NumStrDto.
//
//	 Assumptions
//	 ===========
//
//	 Since this is a low-level routine, the following assumptions
//	 apply to input parameters:
//
//	 1. 'n1NumDto' and 'n2NumDto' have the same sign value
//
//	 2. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    have equal length. Equal length can be achieved by calling
//	    new(numStrDtoPhoton).formatForMathOps() before calling
//	    this method.
//
//	 3. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    contain valid 'runes' with value >= '0' && value <= '9'.
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
//		parameter 'numSeps'. If these Numeric Separators are determined
//	 to be invalid, an error will be returned.
func (nStrDtoTau *numStrDtoTau) lowLevelAddition(
	numSeps NumericSeparatorDto,
	n1NumDto *NumStrDto,
	validateN1NumDto bool,
	n2NumDto *NumStrDto,
	validateN2NumDto bool,
	precision uint,
	newSignVal int,
	validateFinalResult bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoTau.lock.Lock()

	defer nStrDtoTau.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoTau.lowLevelAddition()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if n1NumDto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n1NumDto'",
			}
	}

	if n2NumDto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n2NumDto'",
			}
	}

	if validateN1NumDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1NumDto, ePrefix.XCpy("Validating 'n1NumDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n1NumDto, ePrefix)",
					ErrContext: "Error: Input parameter 'n1NumDto' is invalid!\n" +
						"'n1NumDto' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateN2NumDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2NumDto, ePrefix.XCpy("Validating 'n2NumDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n2NumDto, ePrefix)",
					ErrContext: "Error: Input parameter 'n2NumDto' is invalid!\n" +
						"'n2NumDto' FAILED validation tests.",
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

	lenN1AllRunes := len(n1NumDto.absAllNumRunes)

	if lenN1AllRunes == 0 {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "n1NumDto.absAllNumRunes == 0",
				ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
					"'n1NumDto' internal rune array is empty with zero length. ",
			}

	}

	lenN2AllRunes := len(n2NumDto.absAllNumRunes)

	if lenN1AllRunes != lenN2AllRunes {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "n1NumDto.absAllNumRunes != n2NumDto.absAllNumRunes",
				ErrMessage: "Error: Input parameters 'n1NumDto' and 'n2NumDto'\n" +
					"have runes arrays of unequal length.\n" +
					"'n1NumDto' and 'n2NumDto' internal rune arrays must have equal length. ",
			}
	}

	n3IntAry := make([]int, lenN1AllRunes+1)

	carry := 0

	n1 := 0

	n2 := 0

	n3 := 0

	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = int(n1NumDto.absAllNumRunes[j]) - 48

		if n1 < 0 || n1 > 9 {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "",
					ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
						"The 'n1NumDto' internal rune array contains invalid characters!\n" +
						fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
							j, n1NumDto.absAllNumRunes[j]),
				}

		}

		n2 = int(n2NumDto.absAllNumRunes[j]) - 48

		if n2 < 0 || n2 > 9 {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "",
					ErrMessage: "Error: Input parameter 'n2NumDto' is invalid!\n" +
						"The 'n2NumDto' internal rune array contains invalid characters!\n" +
						fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
							j, n2NumDto.absAllNumRunes[j]),
				}
		}

		n3 = n1 + n2 + carry

		carry = 0

		if n3 > 9 {

			n3 = n3 - 10

			carry = 1
		}

		n3IntAry[j+1] = n3
	}

	if carry > 0 {

		n3IntAry[0] = carry
	}

	// return nDto.FindIntArraySignificantDigitLimits(n3IntAry, precision, newSignVal)

	nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(
		numSeps, n3IntAry, precision, newSignVal, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(",
				ErrContext: "  numSeps, n3IntAry, precision, newSignVal, ePrefix)",
				ErrMessage: err.Error(),
			}
	}

	isZeroValue, err := new(numStrDtoElectron).isNumStrZeroValue(
		&nOutDto, false, ePrefix.XCpy("nDtoOut=0 ?"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " isZeroValue, err :=  new(numStrDtoElectron).isNumStrZeroValue(\n" +
					"  &nDtoOut, true, ePrefix.XCpy(\"nDtoOut=0 ?\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if isZeroValue {
		newSignVal = 1
	}

	err = new(numStrDtoAtom).setSignValue(
		&nOutDto, false, newSignVal, ePrefix.XCpy("n1DtoSetup"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " err = new(numStrDtoAtom).setSignValue(&nDtoOut, true, 1,\n" +
					"  ePrefix.XCpy(\"nDtoOut\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if validateFinalResult {

		err = new(numStrDtoElectron).isValidNumStrDto(
			&nOutDto, ePrefix.XCpy("Validating Final Result 'nOutDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  &nOutDto, ePrefix)",
					ErrContext: "Error: Final Calculated Result 'nOutDto' is invalid!\n" +
						"'nOutDto' FAILED final validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	return nOutDto, nil

}

// lowLevelAddition
//
//	 Performs addition on two NumStrDto values and returns the
//	 total value as a new instance of NumStrDto.
//
//	 Assumptions
//	 ===========
//
//	 Since this is a low-level routine, the following assumptions
//	 apply to input parameters:
//
//	 1. 'n1NumDto' and 'n2NumDto' have the same sign value
//
//	 2. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    have equal length. Equal length can be achieved by calling
//	    new(numStrDtoPhoton).formatForMathOps() before calling
//	    this method.
//
//	 3. With respect to numeric values 'n1NumDto' must be greater
//	    than or equal to 'n2NumDto'. Again, this relationship can
//	    be achieved by calling
//	    new(numStrDtoPhoton).formatForMathOps() before calling this
//	    method.
//
//	 4. The internal rune arrays for 'n1NumDto' and 'n2NumDto'
//	    contain valid 'runes' with value >= '0' && value <= '9'.
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
//		parameter 'numSeps'. If these Numeric Separators are determined
//	 to be invalid, an error will be returned.
func (nStrDtoTau *numStrDtoTau) lowLevelSubtraction(
	numSeps NumericSeparatorDto,
	n1NumDto *NumStrDto,
	validateN1NumDto bool,
	n2NumDto *NumStrDto,
	validateN2NumDto bool,
	precision uint,
	newSignVal int,
	validateFinalResult bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoTau.lock.Lock()

	defer nStrDtoTau.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoTau.lowLevelSubtraction()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if n1NumDto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n1NumDto'",
			}
	}

	if n2NumDto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n2NumDto'",
			}
	}

	if validateN1NumDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1NumDto, ePrefix.XCpy("Validating 'n1NumDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n1NumDto, ePrefix)",
					ErrContext: "Error: Input parameter 'n1NumDto' is invalid!\n" +
						"'n1NumDto' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateN2NumDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2NumDto, ePrefix.XCpy("Validating 'n2NumDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n2NumDto, ePrefix)",
					ErrContext: "Error: Input parameter 'n2NumDto' is invalid!\n" +
						"'n2NumDto' FAILED validation tests.",
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

	lenN1AllRunes := len(n1NumDto.absAllNumRunes)

	if lenN1AllRunes == 0 {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "n1NumDto.absAllNumRunes == 0",
				ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
					"'n1NumDto' internal rune array is empty with zero length. ",
			}

	}

	lenN2AllRunes := len(n2NumDto.absAllNumRunes)

	if lenN1AllRunes != lenN2AllRunes {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "n1NumDto.absAllNumRunes != n2NumDto.absAllNumRunes",
				ErrMessage: "Error: Input parameters 'n1NumDto' and 'n2NumDto'\n" +
					"have runes arrays of unequal length.\n" +
					"'n1NumDto' and 'n2NumDto' internal rune arrays must have equal length. ",
			}
	}

	n1IntAry := make([]int, lenN1AllRunes)

	n2IntAry := make([]int, lenN1AllRunes)

	n3IntAry := make([]int, lenN1AllRunes)

	for i := 0; i < lenN1AllRunes; i++ {

		n1IntAry[i] = int(n1NumDto.absAllNumRunes[i]) - 48

		n2IntAry[i] = int(n2NumDto.absAllNumRunes[i]) - 48

	}

	carry := 0
	n1 := 0
	n2 := 0
	n3 := 0

	// Main Subtraction Routine
	for j := lenN1AllRunes - 1; j >= 0; j-- {

		n1 = n1IntAry[j]

		if n1 < 0 || n1 > 9 {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "",
					ErrMessage: "Error: Input parameter 'n1NumDto' is invalid!\n" +
						"The 'n1NumDto' internal rune array contains invalid characters!\n" +
						fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
							j, n1NumDto.absAllNumRunes[j]),
				}

		}

		n2 = n2IntAry[j]

		if n2 < 0 || n2 > 9 {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "",
					ErrMessage: "Error: Input parameter 'n2NumDto' is invalid!\n" +
						"The 'n2NumDto' internal rune array contains invalid characters!\n" +
						fmt.Sprintf("n1NumDto.absAllNumRunes[%d] = '%v'",
							j, n2NumDto.absAllNumRunes[j]),
				}
		}

		n3 = 0

		if n1-carry-n2 < 0 {

			n1 += 10

			n3 = n1 - n2 - carry

			carry = 1

		} else {

			n3 = n1 - n2 - carry

			carry = 0
		}

		n3IntAry[j] = n3

	}

	//nOutDto, err := nDto.FindIntArraySignificantDigitLimits(n3IntAry, precision, newSignVal)

	nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(
		numSeps, n3IntAry, precision, newSignVal, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nOutDto, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(",
				ErrContext: "  numSeps, n3IntAry, precision, newSignVal, ePrefix)",
				ErrMessage: err.Error(),
			}
	}

	isZeroValue, err := new(numStrDtoElectron).isNumStrZeroValue(
		&nOutDto, false, ePrefix.XCpy("nDtoOut=0 ?"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " isZeroValue, err :=  new(numStrDtoElectron).isNumStrZeroValue(\n" +
					"  &nDtoOut, true, ePrefix.XCpy(\"nDtoOut=0 ?\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if isZeroValue {
		newSignVal = 1
	}

	//nDtoOut.SetSignValue(newSignVal)

	err = new(numStrDtoAtom).setSignValue(
		&nOutDto, false, newSignVal, ePrefix.XCpy("n1DtoSetup"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " err = new(numStrDtoAtom).setSignValue(&nDtoOut, true, 1,\n" +
					"  ePrefix.XCpy(\"nDtoOut\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if validateFinalResult {

		err = new(numStrDtoElectron).isValidNumStrDto(
			&nOutDto, ePrefix.XCpy("Validating Final Result 'nOutDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  &nOutDto, ePrefix)",
					ErrContext: "Error: Final Calculated Result 'nOutDto' is invalid!\n" +
						"'nOutDto' FAILED final validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	return nOutDto, nil
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
//	current instance of 'nDto'. If these Numeric Separators are
//	determined to be invalid, an error will be returned.
func (nStrDtoTau *numStrDtoTau) multiplyNumStrs(
	numSeps NumericSeparatorDto,
	n1NumDto *NumStrDto,
	validateN1NumDto bool,
	n2NumDto *NumStrDto,
	validateN2NumDto bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoTau.lock.Lock()

	defer nStrDtoTau.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoTau.multiplyNumStrs()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if n1NumDto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n1NumDto'",
			}
	}

	if n2NumDto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n2NumDto'",
			}
	}

	if validateN1NumDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1NumDto, ePrefix.XCpy("Validating 'n1NumDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n1NumDto, ePrefix)",
					ErrContext: "Error: Input parameter 'n1NumDto' is invalid!\n" +
						"'n1NumDto' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateN2NumDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2NumDto, ePrefix.XCpy("Validating 'n2NumDto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n2NumDto, ePrefix)",
					ErrContext: "Error: Input parameter 'n2NumDto' is invalid!\n" +
						"'n2NumDto' FAILED validation tests.",
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

	lenN1AbsAllRunes := len(n1NumDto.absAllNumRunes)

	lenN2AbsAllRunes := len(n2NumDto.absAllNumRunes)

	var n1Setup NumStrDto
	var n2Setup NumStrDto

	if lenN2AbsAllRunes > lenN1AbsAllRunes {

		//n1Setup = n2Dto.CopyOut()

		err = new(numStrDtoMolecule).copy(&n1Setup, n2NumDto, false, ePrefix.XCpy("n2NumDto->n1Setup"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
						"  &n1Setup, n2NumDto, false, ePrefix.XCpy(\"n2NumDto->n1Setup\"))",
					ErrContext: "Error Copying n2NumDto->n1Setup",
					ErrMessage: err.Error(),
				}
		}

		//n2Setup = n1Dto.CopyOut()

		err = new(numStrDtoMolecule).copy(&n2Setup, n1NumDto, false, ePrefix.XCpy("n1NumDto->n2Setup"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
						"  &n2Setup, n1NumDto, false, ePrefix.XCpy(\"n1NumDto->n2Setup\"))",
					ErrContext: "Error Copying n1NumDto->n2Setup",
					ErrMessage: err.Error(),
				}
		}

	} else {
		// Must be lenN1AbsAllRunes >= lenN2AbsAllRunes

		// n1Setup = n1Dto.CopyOut()

		err = new(numStrDtoMolecule).copy(&n1Setup, n1NumDto, false, ePrefix.XCpy("n1NumDto->n1Setup"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
						"  &n1Setup, n1NumDto, false, ePrefix.XCpy(\"n1NumDto->n1Setup\"))",
					ErrContext: "Error Copying n1NumDto->n1Setup",
					ErrMessage: err.Error(),
				}
		}

		//n2Setup = n2Dto.CopyOut()

		err = new(numStrDtoMolecule).copy(&n2Setup, n2NumDto, false, ePrefix.XCpy("n2NumDto->n2Setup"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoMolecule).copy(\n" +
						"  &n2Setup, n2NumDto, false, ePrefix.XCpy(\"n2NumDto->n2Setup\"))",
					ErrContext: "Error Copying n2NumDto->n2Setup",
					ErrMessage: err.Error(),
				}
		}

	}

	//else {
	//  // Must be lenN1AbsAllRunes == lenN2AbsAllRunes
	//
	//  n1Setup = n1Dto.CopyOut()
	//
	//  n2Setup = n2Dto.CopyOut()
	//
	//}

	newPrecision := n1Setup.precision + n2Setup.precision

	newSignVal := 1

	if n1Setup.signVal == n2Setup.signVal {

		newSignVal = 1

	} else {

		// Must be n1Setup.signVal != n2Setup.signVal
		newSignVal = -1

	}

	lenN1AbsAllRunes = len(n1Setup.absAllNumRunes)

	lenN2AbsAllRunes = len(n2Setup.absAllNumRunes)

	lenLevels := lenN2AbsAllRunes

	lenNumPlaces := (lenN1AbsAllRunes + lenN2AbsAllRunes) + 1

	intMAry := make([][]int, lenLevels)

	for i := 0; i < lenLevels; i++ {
		intMAry[i] = make([]int, lenNumPlaces)
	}

	intFinalAry := make([]int, lenNumPlaces+1)

	carry := 0

	levels := 0

	place := 0

	n1 := 0

	n2 := 0

	n3 := 0

	n4 := 0

	for i := lenN2AbsAllRunes - 1; i >= 0; i-- {

		place = (lenNumPlaces - 1) - levels

		for j := lenN1AbsAllRunes - 1; j >= 0; j-- {

			n1 = int(n1Setup.absAllNumRunes[j]) - 48

			n2 = int(n2Setup.absAllNumRunes[i]) - 48

			n3 = (n1 * n2) + carry

			//n4 = int(math.Mod(float64(n3), float64(10.00)))
			n4 = int(math.Mod(float64(n3), 10.00))

			intMAry[levels][place] = n4

			//carry = int(n3 / 10)
			carry = n3 / 10

			place--
		}

		intMAry[levels][place] = carry
		carry = 0
		levels++
	}

	carry = 0
	n1 = 0
	n2 = 0
	n3 = 0
	n4 = 0

	for i := 0; i < lenLevels; i++ {

		for j := lenNumPlaces - 1; j >= 0; j-- {

			n1 = intFinalAry[j+1]

			n2 = intMAry[i][j]

			n3 = n1 + n2 + carry

			n4 = 0

			if n3 > 9 {

				//n4 = int(math.Mod(float64(n3), float64(10.0)))
				n4 = int(math.Mod(float64(n3), 10.0))

				carry = n3 / 10

			} else {

				n4 = n3

				carry = 0
			}

			intFinalAry[j+1] = n4
		}

		if carry > 0 {

			intFinalAry[0] = carry
		}

	}

	//numStrOut, err := nDto.FindIntArraySignificantDigitLimits(intFinalAry, newPrecision, newSignVal)
	//
	//if err != nil {
	//  return NumStrDto{},
	//    fmt.Errorf(ePrefix+
	//      "- Error returned from nDto.FindIntArraySignificantDigitLimits(intFinalAry,newPrecision, "+
	//      "newSignVal). Error= %v", err)
	//}

	numStrOut, err := new(numStrDtoAtom).findIntArraySignificantDigitLimits(
		numSeps, intFinalAry, newPrecision, newSignVal, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numStrOut, err := new(numStrDtoAtom).\n" +
					"  findIntArraySignificantDigitLimits(\n" +
					"  numSeps, intFinalAry, newPrecision, newSignVal, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(numStrDtoElectron).isValidNumStrDto(
		&numStrOut, ePrefix.XCpy("Validating Final Result 'numStrOut'"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
					"  &numStrOut, ePrefix)",
				ErrContext: "Error: Final Result 'numStrOut' is invalid!\n" +
					"'numStrOut' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return numStrOut, nil
}
