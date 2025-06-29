package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type bigIntMathNthRootNeutron struct {
	lock sync.Mutex
}

// calcNthRootGateway
//
//	Calculate the nth root of a radicand. The result is returned as a
//	BigIntNum.
//
//	Input Parameters
//	================
//
//	radicand                 BigIntNum
//	  A 'radicand' is the number under a radical symbol (√)
//
//	nthRoot                  BigIntNum
//	  n√ The index or root. This must be a positive integer value.
//
//	maxPrecision             uint
//	  The maximum precision specified for the result.
func (bIMathNthrtNeutron *bigIntMathNthRootNeutron) calcNthRootGateway(
	nthrt *BigIntMathNthRoot,
	radicand *BigIntNum,
	validateRadicand bool,
	nthRoot *BigIntNum,
	validateNthRoot bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	bIMathNthrtNeutron.lock.Lock()

	defer bIMathNthrtNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootNeutron.calcNthRootGateway",
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

	if nthrt == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'nthrt'",
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
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	if nthRootSignValue < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
					"This method cannot process a negative Nth Root Value.",
			}
	}

	_, err := new(bigIntMathNthRootProton).initializeBigIntMathNthRoot(
		nthrt, radicand, nthRoot, maxPrecision, ePrefix)

	if err != nil {
		return new(BigIntNum).NewZero(0),
			fmt.Errorf(ePrefix+"-Error returned from initialization. Error= %v", err.Error())
	}

	err = nthrt.doRootExtraction()

	if err != nil {
		return new(BigIntNum).NewZero(0),
			fmt.Errorf(ePrefix+"-Error returned from nthrt.doRootExtraction() - %v", err.Error())
	}

	// fmt.Println("FracPrecision Count: ", nthrt.FracPrecision.Text(10))

	return nthrt.ResultBINum, nil
}
