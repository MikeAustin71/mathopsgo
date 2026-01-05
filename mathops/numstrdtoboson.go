package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type numStrDtoBoson struct {
	lock sync.Mutex
}

// addNumStrs
//
//	Adds the values represented by two NumStrDto objects and
//	returns the result as a new instance of NumStrDto.
func (nStrDtoBoson *numStrDtoBoson) addNumStrs(
	numSeps NumericSeparatorDto,
	n1Dto *NumStrDto,
	validateN1Dto bool,
	n2Dto *NumStrDto,
	validateN2Dto bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoBoson.lock.Lock()

	defer nStrDtoBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoBoson.addNumStrs()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if n1Dto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n1Dto'",
			}
	}

	if n2Dto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n2Dto'",
			}
	}

	if validateN1Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n1Dto, ePrefix)",
					ErrContext: "Error: Input parameter 'n1Dto' is invalid!\n" +
						"'n1Dto' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateN2Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoElectron).isValidNumStrDto(\n" +
						"  n2Dto, ePrefix)",
					ErrContext: "Error: Input parameter 'n2Dto' is invalid!\n" +
						"'n2Dto' FAILED validation tests.",
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
				ErrContext: "Error: Numeric Separators input paramter ('numSeps') is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	n1DtoNumStr, err := new(numStrDtoAtom).formatNumStr(
		n1Dto, true, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n1DtoNumStr, err :=new(numStrDtoAtom).\n" +
					"  formatNumStr(LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "Error: Could not convert 'n1Dto' to a valid number string.",
				ErrMessage: err.Error(),
			}
	}

	n2DtoNumStr, err := new(numStrDtoAtom).formatNumStr(
		n2Dto, true, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2DtoNumStr, err :=new(numStrDtoAtom).\n" +
					"  formatNumStr(LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "Error: Could not convert 'n2Dto' to a valid number string.",
				ErrMessage: err.Error(),
			}
	}

	n1DtoSetup, n2DtoSetup, _, _, err :=
		new(numStrDtoPhoton).formatForMathOps(
			numSeps,
			n1Dto,
			false,
			n2Dto,
			false,
			ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n1DtoSetup, n2DtoSetup, _, _, err := new(numStrDtoPhoton).formatForMathOps(\n" +
					"  numSeps, n1Dto, false, n2Dto, false, ePrefix)",
				ErrContext: fmt.Sprintf("n1Dto= '%v'\nn2Dto= '%v'",
					n1DtoNumStr, n2DtoNumStr),
				ErrMessage: err.Error(),
			}
	}

	newSignVal := n1DtoSetup.signVal

	var nDtoOut NumStrDto

	var precision uint
	precision = n1DtoSetup.precision

	originalN1DtoSetupSignVal := n1DtoSetup.signVal

	originalN2DtoSetupSignVal := n2DtoSetup.signVal

	// Set up n1DtoSetup sign values for
	// Addition and Subtraction operations
	// Set Sign Value to '1'
	err = new(numStrDtoAtom).setSignValue(
		&n1DtoSetup, true, 1, ePrefix.XCpy("n1DtoSetup"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " err = new(numStrDtoAtom).setSignValue(&n1DtoSetup, true, 1,\n" +
					"  ePrefix.XCpy(\"n1DtoSetup\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// Set up n2DtoSetup sign values for
	// Addition and Subtraction operations
	// Set Sign Value to '1'
	err = new(numStrDtoAtom).setSignValue(
		&n2DtoSetup, true, 1, ePrefix.XCpy("n2DtoSetup"))

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " err = new(numStrDtoAtom).setSignValue(&n2DtoSetup, true, 1,\n" +
					"  ePrefix.XCpy(\"n2DtoSetup\"))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if originalN1DtoSetupSignVal != originalN2DtoSetupSignVal {
		// Do Subtraction

		// Do subtraction operation
		nDtoOut, err = new(numStrDtoTau).lowLevelSubtraction(
			numSeps,
			&n1DtoSetup,
			true,
			&n2DtoSetup,
			true,
			precision,
			newSignVal,
			true,
			ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nDtoOut, err := new(numStrDtoTau).lowLevelSubtraction(\n" +
						"  numSeps, true, &n2DtoSetup, true, precision,\n" +
						"  newSignVal, true, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// End of n1DtoSetup.signVal != n2DtoSetup.signVal

	} else {

		// MUST BE:
		// n1DtoSetup.signVal == n2DtoSetup.signVal
		// Sign Values are Equal
		// Do addition

		// Do addition operation
		nDtoOut, err = new(numStrDtoTau).lowLevelAddition(
			numSeps,
			&n1DtoSetup,
			true,
			&n2DtoSetup,
			true,
			precision,
			newSignVal,
			true,
			ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nDtoOut, err := new(numStrDtoTau).lowLevelSubtraction(\n" +
						"  numSeps, true, &n2DtoSetup, true, precision,\n" +
						"  newSignVal, true, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return nDtoOut, nil
}

// getSignedBigIntPrecision
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
func (nStrDtoBoson *numStrDtoBoson) getSignedBigIntPrecision(
	nDto *NumStrDto,
	validateNumStrDto bool,
	errPrefDto *ePref.ErrPrefixDto) (
	signedBIntNum *big.Int, precision uint, err error) {

	nStrDtoBoson.lock.Lock()

	defer nStrDtoBoson.lock.Unlock()

	signedBIntNum = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoBoson.getSignedBigIntPrecision",
		"")

	if err != nil {
		return signedBIntNum, precision, err
	}

	if nDto == nil {

		return signedBIntNum, precision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'nDto'",
			}
	}

	if validateNumStrDto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			nDto, ePrefix.XCpy("Validating 'nDto'"))

		if err != nil {
			return signedBIntNum, precision,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: NumStrDto ('nDto') is INVALID!\n" +
						"'nDto' FAILED Validation Tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	absBigInt, biPrecision, err := new(numStrDtoGluon).getAbsoluteBigInt(
		nDto, false, ePrefix)

	if err != nil {

		return signedBIntNum, precision,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "absBigInt, biPrecision, err := new(numStrDtoGluon).\n" +
					"getAbsoluteBigInt(nDto, false, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if nDto.signVal < 0 {
		absBigInt = big.NewInt(0).Neg(absBigInt)
	}

	return big.NewInt(0).Set(absBigInt), biPrecision, nil
}

// subtractNumStrs
//
//	Subtracts the numeric values represented by two NumStrDto
//	objects.
func (nStrDtoBoson *numStrDtoBoson) subtractNumStrs(
	numSeps NumericSeparatorDto,
	n1Dto *NumStrDto,
	validateN1Dto bool,
	n2Dto *NumStrDto,
	validateN2Dto bool,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	nStrDtoBoson.lock.Lock()

	defer nStrDtoBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrDtoBoson.subtractNumStrs()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if n1Dto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n1Dto'",
			}
	}

	if n2Dto == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'n2Dto'",
			}
	}

	if validateN1Dto {

		err = new(numStrDtoElectron).isValidNumStrDto(
			n1Dto, ePrefix.XCpy("Validating 'n1Dto'"))

		if err != nil {

			return NumStrDto{},
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

		err = new(numStrDtoElectron).isValidNumStrDto(
			n2Dto, ePrefix.XCpy("Validating 'n2Dto'"))

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
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

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\"Validating 'numSeps'\").String())",
				ErrContext: "Error: Numeric Separators input paramter 'numSeps' is INVALID!\n" +
					"'numSeps' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	n1DtoNumStr, err := new(numStrDtoAtom).formatNumStr(
		n1Dto, true, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n1DtoNumStr, err :=new(numStrDtoAtom).\n" +
					"  formatNumStr(LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "Error: Could not convert 'n1Dto' to a valid number string.",
				ErrMessage: err.Error(),
			}
	}

	n2DtoNumStr, err := new(numStrDtoAtom).formatNumStr(
		n2Dto, true, LEADMINUSNEGVALFMTMODE, ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "n2DtoNumStr, err :=new(numStrDtoAtom).\n" +
					"  formatNumStr(LEADMINUSNEGVALFMTMODE, ePrefix)",
				ErrContext: "Error: Could not convert 'n2Dto' to a valid number string.",
				ErrMessage: err.Error(),
			}
	}

	// 1. Equalize lengths of internal rune arrays
	// 2. Position n1NumDto >= n2NumDto
	n1DtoSetup, n2DtoSetup, compare, isReversed, err :=
		new(numStrDtoPhoton).formatForMathOps(
			numSeps,
			n1Dto,
			false,
			n2Dto,
			false,
			ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " n1DtoSetup, n2DtoSetup, compare, isReversed, err :=\n" +
					"  new(numStrDtoPhoton).formatForMathOps(\n" +
					"  numSeps, n1Dto, false, n2Dto, false, compare, ePrefix)",
				ErrContext: fmt.Sprintf("n1Dto= '%v'\nn2Dto= '%v'",
					n1DtoNumStr, n2DtoNumStr),
				ErrMessage: err.Error(),
			}
	}

	if compare == 0 {

		return new(numStrDtoMolecule).newZeroNumStrDto(numSeps, n1DtoSetup.precision), nil
	}

	newSignVal := n1DtoSetup.signVal

	var nDtoOut NumStrDto

	var precision uint
	precision = n1DtoSetup.precision

	originalN1DtoSetupSignVal := n1DtoSetup.signVal

	originalN2DtoSetupSignVal := n2DtoSetup.signVal

	if originalN1DtoSetupSignVal != originalN2DtoSetupSignVal {

		// err = n1NumDto.SetSignValue(1)
		err = new(numStrDtoAtom).setSignValue(&n1DtoSetup, false, 1, ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoAtom).setSignValue(\n" +
						"  &n1NumDto, false, 1, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		// err = n2NumDto.SetSignValue(1)
		err = new(numStrDtoAtom).setSignValue(&n2DtoSetup, false, 1, ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(numStrDtoAtom).setSignValue(\n" +
						"  &n2NumDto, false, 1, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		//nOutDto, err = nDto.AddNumStrs(n1NumDto, n2NumDto)
		//
		//if err != nil {
		//	return NumStrDto{},
		//		fmt.Errorf(ePrefix+"- Error from nDto.AddNumStrs(n1NumDto, n2NumDto). "+
		//			"Error= %v", err)
		//}

		// Do subtraction operation
		nDtoOut, err = new(numStrDtoTau).lowLevelAddition(
			numSeps,
			&n1DtoSetup,
			true,
			&n2DtoSetup,
			true,
			precision,
			newSignVal,
			true,
			ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nDtoOut, err := new(numStrDtoTau).lowLevelAddition(\n" +
						"  numSeps, true, &n2DtoSetup, true, precision,\n" +
						"  newSignVal, true, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		//nOutDto.SetSignValue(newSignVal)

		//return nOutDto, nil
		// End Of if n1NumDto.signVal != n2NumDto.signVal
	} else {

		// Must be:
		// n1NumDto.signVal == n2NumDto.signVal
		// Do Subtraction
		// Change sign for subtraction
		newSignVal = n1DtoSetup.signVal

		if isReversed {
			newSignVal = newSignVal * -1
		}

		// Do subtraction operation
		nDtoOut, err = new(numStrDtoTau).lowLevelSubtraction(
			numSeps,
			&n1DtoSetup,
			true,
			&n2DtoSetup,
			true,
			precision,
			newSignVal,
			true,
			ePrefix)

		if err != nil {

			return NumStrDto{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nDtoOut, err := new(numStrDtoTau).lowLevelSubtraction(\n" +
						"  numSeps, true, &n2DtoSetup, true, precision,\n" +
						"  newSignVal, true, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	return nDtoOut, nil
}
