package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type bigIntMathNthRootNanobot struct {
	lock sync.Mutex
}

//                             Stage 2                                             //
// ******************************************************************************* //

func (bIMathNthrtNanobot *bigIntMathNthRootNanobot) calcPositiveIntegerNthRoot(
	nthrt *BigIntMathNthRoot,
	radicand *BigIntNum,
	validateRadicand bool,
	nthRoot *BigIntNum,
	validateNthRoot bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	bIMathNthrtNanobot.lock.Lock()

	defer bIMathNthrtNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootNanobot.calcPositiveIntegerNthRoot",
		"")

	if err != nil {
		return BigIntNum{}, err
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

	if nthRootSignValue != 1 {

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

	if nthRootPrecisionUint != 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
					"'nthRoot' is a fractional value. This method only calculates\n" +
					"results for integer nthRoot values.",
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

	radicandSignValue, err := radicand.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "radicandSignValue, err := radicand.GetSign()",
				ErrContext: fmt.Sprintf("radicand = '%v'", radicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	if radicandSignValue == -1 {

		isEvenNum, err := nthRoot.IsEvenNumber()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "isEvenNum, err := nthRoot.IsEvenNumber()",
					ErrContext: fmt.Sprintf("nthRoot = '%v'", nthRootNumStr),
					ErrMessage: err.Error(),
				}
		}

		if isEvenNum {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: fmt.Sprintf("radicand= '%v'\n"+
						"nthRoot= '%v'",
						radicandNumStr, nthRootNumStr),
					ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
						"'nthRoot' is a negative value and 'nthRoot' is an even number.\n" +
						"Cannot calculate nthRoot of a negative number when nthRoot is even.",
				}
		}
	}

	bINumResult, err := new(bigIntMathNthRootNeutron).calcNthRootGateway(nthrt, radicand, false, nthRoot, false, maxPrecision, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " bINumResult, err := new(bigIntMathNthRootNeutron).\n" +
					"  calcNthRootGateway(nthrt, radicand, false, nthRoot, false,\n" +
					"    maxPrecision, ePrefix)",
				ErrContext: fmt.Sprintf("radicand= '%v'\n"+
					"nthRoot= '%v'\n"+
					"maxPrecision= '%v'",
					radicandNumStr, nthRootNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return bINumResult, nil
}
