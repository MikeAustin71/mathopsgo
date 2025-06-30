package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathNthRootAtom struct {
	lock sync.Mutex
}

// calcPositiveFractionalNthRoot
//
//		Calculates the nth root result for positive fractional nthRoot
//		values.
//
//	 Input Parameters
//	 ================
//
//	 radicand                 BigIntNum
//	   A 'radicand' is the number under a radical symbol (√)
//
//	 nthRoot                  BigIntNum
//	   n√ The index or root. This must be a positive integer value.
//
//	 maxPrecision             uint
//	   The maximum precision specified for the result.
func (bIMathNthrtAtom *bigIntMathNthRootAtom) calcPositiveFractionalNthRoot(
	nthrt *BigIntMathNthRoot,
	radicand *BigIntNum,
	validateRadicand bool,
	nthRoot *BigIntNum,
	validateNthRoot bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	bIMathNthrtAtom.lock.Lock()

	defer bIMathNthrtAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathNthRootAtom.calcPositiveFractionalNthRoot",
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

	numSeps, err := radicand.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := radicand.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if !validateNthRoot {

		err = numSeps.IsValid(ePrefix.XCpy("Validating 'radicand numSeps'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = numSeps.IsValid(ePrefix.XCpy(\n" +
						"  \"Validating 'radicand numSeps'\").String())",
					ErrContext: "Error: Input parameter 'radicand' is INVALID!\n" +
						"Numeric Separators extracted from 'radicand' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
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

	if nthRootSignValue < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
					"This method only processes Positive NthRoot Values.\n" +
					"'nthRoot' is a negative number!",
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

	if nthRootPrecisionUint == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: "Error: Input parameter 'nthRoot' is INVALID!\n" +
					"'nthRoot' is an integer value.\n" +
					"This method only processes Fractional NthRoot Values.",
			}
	}

	modXZero := big.NewInt(0)

	scaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(nthRootPrecisionUint)),
		modXZero)

	nthRootAbsoluteBigInt, err := nthRoot.GetAbsoluteBigIntValue()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootAbsoluteBigInt, err := nthRoot.GetAbsoluteBigIntValue()",
				ErrContext: fmt.Sprintf("nthRoot= '%v'", nthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	ratFraction := big.NewRat(1, 1).SetFrac(nthRootAbsoluteBigInt, scaleFactor)

	// Numerator of ratFraction is new nthRoot
	newNthRoot, err := new(BigIntNum).NewBigInt(ratFraction.Num(), 0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "newNthRoot, err := new(BigIntNum).\n" +
					" NewBigInt(ratFraction.Num(), 0)",
				ErrContext: fmt.Sprintf("ratFraction.Num()= '%v'", ratFraction.Num().Text(10)),
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

	newNthRootIsZero, err := newNthRoot.IsZero()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newNthRootIsZero, err := newNthRoot.IsZero()",
				ErrContext: fmt.Sprintf("newNthRoot= '%v'", newNthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	// If nthRoot is zero, the result will always be '1'
	if newNthRootIsZero {

		bINumOne, err := new(BigIntNum).NewBigIntNumSeps(big.NewInt(1), maxPrecision, numSeps)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bINumOne, err := new(BigIntNum).\n" +
						"  NewBigIntNumSeps(big.NewInt(1), maxPrecision, numSeps)",
					ErrContext: fmt.Sprintf("maxPrecision = '%v'", maxPrecision),
					ErrMessage: err.Error(),
				}
		}

		return bINumOne, nil
	}

	bigINumOne, err := new(BigIntNum).NewOne(0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bigINumOne, err := new(BigIntNum).NewOne(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newNthRootCmpbigINumOne, err := newNthRoot.Cmp(bigINumOne)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "newNthRootCmpbigINumOne, err := \n" +
					"  newNthRoot.Cmp(bigINumOne)",
				ErrContext: fmt.Sprintf("newNthRoot= '%v'", newNthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	// if nthRoot == 1
	if newNthRootCmpbigINumOne == 0 {

		result1, err := radicand.CopyOut()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "result1, err := radicand.CopyOut()",
					ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
					ErrMessage: err.Error(),
				}
		}

		return result1, nil
	}

	// Denominator of ratFraction is exponent for current radicand.
	exponentBigINum, err := new(BigIntNum).NewBigInt(ratFraction.Denom(), 0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "exponentBigINum, err := new(BigIntNum).\n" +
					"NewBigInt(ratFraction.Denom(), 0)",
				ErrContext: fmt.Sprintf("ratFraction.Denom()= '%v'",
					ratFraction.Denom().Text(10)),
				ErrMessage: err.Error(),
			}
	}

	exponentBigINumStr, err := exponentBigINum.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "exponentBigINumStr, err := \n" +
					"exponentBigINum.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	radicandPrecisionUint, err := radicand.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "radicandPrecisionUint, err := radicand.GetPrecisionUint()",
				ErrContext: fmt.Sprintf("radicand= '%v'", radicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	radicandPrecisionBigINum, err := new(BigIntNum).NewBigInt(big.NewInt(int64(radicandPrecisionUint)), 0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "radicandPrecision, err := new(BigIntNum).NewBigInt(\n" +
					"  big.NewInt(int64(radicandPrecisionUint)), 0)",
				ErrContext: fmt.Sprintf("radicandPrecisionUint= '%v'", radicandPrecisionUint),
				ErrMessage: err.Error(),
			}
	}

	newMaxPrecision, err := new(BigIntMathMultiply).MultiplyBigIntNums(exponentBigINum, radicandPrecisionBigINum)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "newMaxPrecision, err := new(BigIntMathMultiply).\n" +
					"  MultiplyBigIntNums(exponentBigINum, radicandPrecisionBigINum)",
				ErrContext: fmt.Sprintf("exponentBigINum= '%v'", exponentBigINumStr),
				ErrMessage: err.Error(),
			}
	}

	newMaxPrecisionNumStr, err := newMaxPrecision.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newMaxPrecisionNumStr, err := newMaxPrecision.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	exponentMaxPrecision, err := newMaxPrecision.GetUInt()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentMaxPrecision, err := newMaxPrecision.GetUInt()",
				ErrContext: fmt.Sprintf("newMaxPrecision= '%v'", newMaxPrecisionNumStr),
				ErrMessage: err.Error(),
			}
	}

	newRadicand, err := new(BigIntMathPower).BigIntNumPwr(*radicand, exponentBigINum, exponentMaxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "newRadicand, err := new(BigIntMathPower).BigIntNumPwr(\n" +
					"  *radicand, exponentBigINum, exponentMaxPrecision)",
				ErrContext: fmt.Sprintf("radicand= '%v'\n"+
					"exponentBigINum= '%v'\n"+
					"exponentMaxPrecision= '%v'",
					radicandNumStr, exponentBigINumStr, exponentMaxPrecision),
				ErrMessage: err.Error(),
			}
	}

	newRadicandNumStr, err := newRadicand.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newRadicandNumStr, err := newRadicand.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newRadicandSignValue, err := newRadicand.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newRadicandSignValue, err :=  newRadicand.GetSign()",
				ErrContext: fmt.Sprintf("newRadicand= '%v'", newRadicandNumStr),
				ErrMessage: err.Error(),
			}
	}

	if newRadicandSignValue == -1 {

		isEvenNum, err := newNthRoot.IsEvenNumber()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "isEvenNum, err := newNthRoot.IsEvenNumber()",
					ErrContext: fmt.Sprintf("newNthRoot= '%v'", newNthRootNumStr),
					ErrMessage: err.Error(),
				}
		}

		if isEvenNum {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: fmt.Sprintf("radicand= '%v'\n"+
						"newNthRoot= '%v'", radicandNumStr, newNthRootNumStr),
					ErrMessage: "Error: Cannot calculate nthRoot of a negative number when nthRoot is even.",
				}
		}
	}

	nthRootResult, err := new(bigIntMathNthRootNeutron).calcNthRootGateway(
		nthrt, &newRadicand, false, &newNthRoot, false, maxPrecision, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthRootResult, err := \n" +
					"  new(bigIntMathNthRootNeutron).calcNthRootGateway(\n" +
					"  nthrt, &newRadicand, false, &newNthRoot, false,\n" +
					"  maxPrecision, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return nthRootResult, nil
}
