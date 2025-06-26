package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathPowerAtom struct {
	lock sync.Mutex
}

// bigIntNumRaiseToPositiveFractionalPower - Assumes that input parameter 'exponent' is positive and a
// fractional number (precision > 0 - has fractional digits). If 'exponent' is negative or if
// 'exponent' is an integer number, an error is returned.
//
// If 'exponent' is both a positive number and a fractional number, this method proceeds to
// raise input parameter 'base' to the power of 'exponent' and return the result as a BigIntNum
// type.
func (bIMathPwrAtom *bigIntMathPowerAtom) bigIntNumRaiseToPositiveFractionalPower(
	base *BigIntNum,
	validateBase bool,
	exponent *BigIntNum,
	validateExponent bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	bIMathPwrAtom.lock.Lock()

	defer bIMathPwrAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerAtom.bigIntNumRaiseToPositiveFractionalPower",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if base == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'base'",
			}
	}

	if exponent == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'exponent'",
			}
	}

	if validateBase {

		err = base.IsValid(ePrefix.XCpy("Validating 'base'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = base.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'base'\").String()",
					ErrContext: "Error: Input parameter 'base' is INVALID!\n" +
						"'base' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	baseNumStr, err := base.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseNumStr, err := base.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if validateExponent {

		err = exponent.IsValid(ePrefix.XCpy("Validating 'exponent'").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = exponent.IsValid(\n" +
						"  ePrefix.XCpy(\"Validating 'exponent'\").String()",
					ErrContext: "Error: Input parameter 'exponent' is INVALID!\n" +
						"'exponent' FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentNumStr, err := exponent.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	exponentPredisionUint, err := exponent.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentPredisionUint, err := exponent.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentPredisionUint == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
				ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
					"'exponent' is an integer number. Only fractional exponents\n" +
					"are eligible for processing.",
			}
	}

	exponentSignValue, err := exponent.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentSignValue, err = exponent.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentSignValue < 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
				ErrMessage: "Error: Input parameter 'exponent' is a negative number.\n" +
					"Only Positive Exponents can be processed by this method!",
			}
	}

	modXZero := big.NewInt(0)

	scaleFactor := big.NewInt(0).Exp(
		big.NewInt(10),
		big.NewInt(int64(exponent.precision)),
		modXZero)

	ratExponent := big.NewRat(1, 1).SetFrac(exponent.bigInt, scaleFactor)

	ratExponentNumBInt := ratExponent.Num()

	bINumNumerator, err := new(BigIntNum).NewBigInt(ratExponentNumBInt, 0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINumNumerator, err := new(BigIntNum).\n" +
					"NewBigInt(ratExponentBInt, 0)",
				ErrContext: fmt.Sprintf("ratExponentNumBInt= '%v'",
					ratExponentNumBInt.Text(10)),
				ErrMessage: err.Error(),
			}
	}

	bINumNumeratorNumStr, err := bINumNumerator.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumNumeratorNumStr, err := bINumNumerator.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	pwr1Result, err := new(bigIntMathPowerNeutron).
		bigIntNumRaiseToPositiveIntegerPower(
			base, false, &bINumNumerator, true, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "pwr1Result, err := new(bigIntMathPowerNeutron).\n" +
					"  bigIntNumRaiseToPositiveIntegerPower(\n" +
					"  base, false, &bINumNumerator, true, ePrefix)",
				ErrContext: fmt.Sprintf("base= '%v'\nbINumNumerator = '%v'",
					baseNumStr, bINumNumeratorNumStr),
				ErrMessage: err.Error(),
			}
	}

	pwr1ResultNumStr, err := pwr1Result.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: " pwr1ResultNumStr, err := pwr1Result.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	ratExponentDenomBInt := ratExponent.Denom()

	nthRoot, err := new(BigIntNum).NewBigInt(ratExponentDenomBInt, 0)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nthRoot, err := new(BigIntNum).\n" +
					"  NewBigInt(ratExponentDenomBInt, 0)",
				ErrContext: fmt.Sprintf("ratExponentDenomBInt= '%v'",
					ratExponentDenomBInt.Text(10)),
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

	nthRootIsZero, err := nthRoot.IsZero()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootIsZero, err := nthRoot.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// If nthRoot is zero, the result will always be '1'
	if nthRootIsZero {

		bINumOne, err := new(BigIntNum).NewOne(maxPrecision)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bINumOne, err := new(BigIntNum).NewOne(maxPrecision)",
					ErrContext: "",
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

	nthRootCmpBigINumOne, err := nthRoot.Cmp(bigINumOne)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "nthRootCmpBigINumOne, err := nthRoot.Cmp(bigINumOne)",
				ErrContext: fmt.Sprintf("nthRoot= '%v'",
					nthRootNumStr),
				ErrMessage: err.Error(),
			}
	}

	// Error if nthRoot == 1
	if nthRootCmpBigINumOne == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "nthRoot == 1",
				ErrMessage: "Error - Intermediate result 'nthRoot' INVALID!\n" +
					"'nthRoot' cannot equal 1.",
			}
	}

	pwr1ResultSignValue, err := pwr1Result.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "pwr1ResultSignValue, err :=  pwr1Result.GetSign()",
				ErrContext: fmt.Sprintf("pwr1Result = '%v'",
					pwr1ResultNumStr),
				ErrMessage: err.Error(),
			}
	}
	// pwr1ResultNumStr

	if pwr1ResultSignValue == -1 {

		isEvenNum, err := nthRoot.IsEvenNumber()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "isEvenNum, err := nthRoot.IsEvenNumber()",
					ErrContext: fmt.Sprintf("nthRoot= '%v'",
						nthRootNumStr),
					ErrMessage: err.Error(),
				}
		}

		if isEvenNum {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "",
					ErrMessage: "Error - Intermediate result 'nthRoot' INVALID!\n" +
						"Cannot calculate nthRoot of a negative number when nthRoot is even.",
				}
		}

	}

	biNumResult, err := new(BigIntMathNthRoot).GetNthRoot(pwr1Result, nthRoot, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " biNumResult, err := new(BigIntMathNthRoot).GetNthRoot(\n" +
					" pwr1Result, nthRoot, maxPrecision)",
				ErrContext: fmt.Sprintf("pwr1Result= '%v'\nnthRoot= '%v'\nmaxPrecision= '%v'",
					pwr1ResultNumStr, nthRootNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	err = biNumResult.IsValid(ePrefix.XCpy("Validating 'biNumResult'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = biNumResult.IsValid(\n" +
					"  ePrefix.XCpy(\"Validating 'biNumResult'\").String()",
				ErrContext: "Error: Final Calculation Result 'biNumResult' is INVALID!\n" +
					"'biNumResult' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return biNumResult, nil
}
