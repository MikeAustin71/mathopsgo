package mathops

import (
	"fmt"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathNthRootMacrobot struct {
	lock sync.Mutex
}

//                             Stage 1                                             //
// ******************************************************************************* //

// calcNegativeNthRoot
//
//	Calculates the nth root result of a radicand where the nth root is
//	a negative value.
func (bIMathNthrtMacrobot *bigIntMathNthRootMacrobot) calcNegativeNthRoot(
	nthrt *BigIntMathNthRoot,
	radicand *BigIntNum,
	validateRadicand bool,
	nthRoot *BigIntNum,
	validateNthRoot bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	bIMathNthrtMacrobot.lock.Lock()

	defer bIMathNthrtMacrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootMacrobot.calcNegativeNthRoot",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if nthrt == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthrt'",
			}
	}

	if radicand == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'radicand'",
			}
	}

	if nthRoot == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthRoot'",
			}
	}

	if validateRadicand {

		err = radicand.IsValid(ePrefix.XCpy("Validating 'radicand'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = radicand.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'radicand'\").String()",
					ErrContext: "Error: Input parameter 'radicand' is INVALID!\n" +
						"'radicand' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateNthRoot {

		err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = nthRoot.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'nthRoot'\").String()",
					ErrContext: "Error: Input parameter 'nthRoot' is INVALID!\n" +
						"'nthRoot' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	radicandNumStr, err := radicand.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "radicandNumStr, err := radicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthRootNumStr, err := nthRoot.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthRootSignValue, err := nthRoot.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootSignValue, err := nthRoot.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if nthRootSignValue != -1 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
					"This method only calculates results for negative 'nthRoot' values.",
			}
	}

	newNthRoot, err := nthRoot.CopyOut()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newNthRoot, err := nthRoot.CopyOut()",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	err = newNthRoot.ChangeSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newNthRoot.ChangeSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newNthRootNumStr, err := newNthRoot.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newNthRootNumStr, err := newNthRoot.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var nthRootResult BigIntNum

	newNthRootPrecisionUint, err := newNthRoot.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newNthRootPrecisionUint, err := newNthRoot.GetPrecisionUint()",
				ErrContext: fmt.Sprintf("newNthRoot= '%v'", newNthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	if newNthRootPrecisionUint == 0 {

		nthRootResult, err = new(bigIntMathNthRootNanobot).
			calcPositiveIntegerNthRoot(
				nthrt, radicand, false, &newNthRoot, true, maxPrecision, ePrefix)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nthRootResult, err = new(bigIntMathNthRootNanobot).\n" +
						"  calcPositiveIntegerNthRoot(\n" +
						"  nthrt, radicand, false, &newNthRoot, true, maxPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("radicand= '%v'\n"+
						"newNthRoot= '%v'\n"+
						"maxPrecision= '%v'", radicandNumStr, newNthRootNumStr, maxPrecision),
					ErrMessage: err.Error(),
				}
		}

	} else {

		nthRootResult, err = new(bigIntMathNthRootAtom).calcPositiveFractionalNthRoot(
			nthrt, radicand, true, &newNthRoot, true, maxPrecision, ePrefix)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "nthRootResult, err = new(bigIntMathNthRootAtom).\n" +
						"  calcPositiveFractionalNthRoot(nthrt, radicand, true,\n" +
						"  &newNthRoot, true, maxPrecision, ePrefix)",
					ErrContext: fmt.Sprintf("radicand= '%v'\n"+
						"newNthRoot= '%v'\n"+
						"maxPrecision= '%v'", radicandNumStr, newNthRootNumStr, maxPrecision),
					ErrMessage: err.Error(),
				}
		}
	}

	nthRootResultNumStr, err := nthRootResult.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootResultNumStr, err := nthRootResult.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	inverse, err := nthRootResult.Inverse(maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "inverse, err := nthRootResult.Inverse(maxPrecision)",
				ErrContext: fmt.Sprintf("nthRootResult= '%v'\n"+
					"maxPrecision= '%v", nthRootResultNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	err = inverse.IsValid(ePrefix.XCpy("Validating 'inverse'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = inverse.IsValid(\n" +
					"  ePrefix.XCpy(\"Validating 'inverse'\").String())",
				ErrContext: "Error: Final calculated result 'inverse' is INVALID!\n" +
					"'inverse' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return inverse, nil
}

//                             Stage 1                                             //
// ******************************************************************************* //

// calcPositiveNthRoot
//
//	Calculates the nth root of a radicand where the nth root is a
//	positive value.
func (bIMathNthrtMacrobot *bigIntMathNthRootMacrobot) calcPositiveNthRoot(
	nthrt *BigIntMathNthRoot,
	radicand *BigIntNum,
	validateRadicand bool,
	nthRoot *BigIntNum,
	validateNthRoot bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	bIMathNthrtMacrobot.lock.Lock()

	defer bIMathNthrtMacrobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootMacrobot.calcPositiveNthRoot",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if nthrt == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthrt'",
			}
	}

	if radicand == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'radicand'",
			}
	}

	if nthRoot == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthRoot'",
			}
	}

	if validateRadicand {

		err = radicand.IsValid(ePrefix.XCpy("Validating 'radicand'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = radicand.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'radicand'\").String()",
					ErrContext: "Error: Input parameter 'radicand' is INVALID!\n" +
						"'radicand' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	if validateNthRoot {

		err = nthRoot.IsValid(ePrefix.XCpy("Validating 'nthRoot'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = nthRoot.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'nthRoot'\").String()",
					ErrContext: "Error: Input parameter 'nthRoot' is INVALID!\n" +
						"'nthRoot' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	nthRootNumStr, err := nthRoot.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootNumStr, err := nthRoot.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	nthRootSignValue, err := nthRoot.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootSignValue, err := nthRoot.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if nthRootSignValue == -1 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
					"'nthRoot' is a negative value. This method only calculates\n" +
					"results for positive nthRoot values.",
			}
	}

	nthRootPrecisionUint, err := nthRoot.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootPrecisionUint, err := nthRoot.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if nthRootPrecisionUint > 0 {

		return new(bigIntMathNthRootAtom).calcPositiveFractionalNthRoot(
			nthrt, radicand, false, nthRoot, false, maxPrecision, ePrefix)
	}

	return new(bigIntMathNthRootNanobot).calcPositiveIntegerNthRoot(
		nthrt, radicand, false, nthRoot, false, maxPrecision, ePrefix)
}
