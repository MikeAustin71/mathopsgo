package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathPowerMolecule struct {
	lock *sync.Mutex
}

// bigIntNumRaiseToNegativeFractionalPower - Assumes that input parameter 'exponent' is negative and a
// fractional number (precision > 0 - has fractional digits). If 'exponent' is positive or if
// 'exponent' is an integer number, an error is returned.
//
// If 'exponent' is both a negative number and a fractional number, this method proceeds to
// raise input parameter 'base' to the power of 'exponent' and return the result as a BigIntNum
// type.
func (bIMathPwrMolecule *bigIntMathPowerMolecule) bigIntNumRaiseToNegativeFractionalPower(
	base *BigIntNum,
	validateBase bool,
	exponent *BigIntNum,
	validateExponent bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIMathPwrMolecule.lock == nil {
		bIMathPwrMolecule.lock = new(sync.Mutex)
	}

	bIMathPwrMolecule.lock.Lock()

	defer bIMathPwrMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerMolecule.bigIntNumRaiseToNegativeFractionalPower",
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
				ErrContext: fmt.Sprintf("exponent= '%v",
					exponentNumStr),
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
					"'exponent' is an integer number.\n" +
					"Only fractional exponents can be processed by this method.",
			}
	}

	exponentSignValue, err := exponent.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentSignValue, err = exponent.GetSign()",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
				ErrMessage: err.Error(),
			}
	}

	if exponentSignValue > 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
				ErrMessage: "Error: Input parameter 'exponent' is a positive number.\n" +
					"Only Negative Exponents can be processed by this method!",
			}
	}

	exponentAbsVal, err := exponent.GetAbsoluteBigIntNumValue()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentAbsVal, err := exponent.GetAbsoluteBigIntNumValue()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	exponentAbsValNumStr, err := exponentAbsVal.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentAbsValNumStr, err := exponentAbsVal.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bINumResult, err :=
		new(bigIntMathPowerAtom).bigIntNumRaiseToPositiveFractionalPower(
			base,
			false,
			&exponentAbsVal,
			true,
			maxPrecision,
			ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bINumResult, err := new(bigIntMathPowerAtom).\n" +
					"  bigIntNumRaiseToPositiveFractionalPower(\n" +
					"  base, false, &exponentAbsVal, true, maxPrecision, ePrefix)",
				ErrContext: fmt.Sprintf("base= '%v'\nexponentAbsVal= '%v'\nmaxPrecision= '%v' ",
					baseNumStr, exponentAbsValNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	bINumResultNumStr, err := bINumResult.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumResultNumStr, err := bINumResult.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}
	// exponentNumStr, baseNumStr, exponentAbsValNumStr, bINumResultNumStr

	inverseResult, err := bINumResult.GetInverse(maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "inverseResult, err := bINumResult.GetInverse(maxPrecision)",
				ErrContext: fmt.Sprintf("bINumResult= '%v'\nmaxPrecision='%v'",
					bINumResultNumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	err = inverseResult.IsValid(ePrefix.XCpy("Validating 'inverseResult'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = inverseResult.IsValid(\n" +
					"  ePrefix.XCpy(\"Validating 'inverseResult'\").String()",
				ErrContext: "Error: Final Calculation Result 'inverseResult' is INVALID!\n" +
					"'inverseResult' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return inverseResult, nil
}

// bigIntNumRaiseToNegativeIntegerPower - Assumes that input parameter 'exponent' is negative and an
// integer number (precision==0 - no fractional digits). If 'exponent' is positive or if
// it is NOT an integer number, an error is returned.
//
// If 'exponent' is both a negative number and an integer number, this method proceeds to
// raise input parameter 'base' to the power of 'exponent' and return the result as a BigIntNum
// type.
func (bIMathPwrMolecule *bigIntMathPowerMolecule) bigIntNumRaiseToNegativeIntegerPower(
	base *BigIntNum,
	validateBase bool,
	exponent *BigIntNum,
	validateExponent bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIMathPwrMolecule.lock == nil {
		bIMathPwrMolecule.lock = new(sync.Mutex)
	}

	bIMathPwrMolecule.lock.Lock()

	defer bIMathPwrMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerMolecule.bigIntNumRaiseToNegativeIntegerPower",
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
				ErrContext: fmt.Sprintf("exponent= '%v",
					exponentNumStr),
				ErrMessage: err.Error(),
			}
	}

	if exponentPredisionUint > 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
				ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
					"'exponent' is an fractional number.\n" +
					"Only integer exponents can be processed by this method.",
			}
	}

	exponentSignValue, err := exponent.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentSignValue, err = exponent.GetSign()",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
				ErrMessage: err.Error(),
			}
	}

	if exponentSignValue > 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
				ErrMessage: "Error: Input parameter 'exponent' is a positive number.\n" +
					"Only Negative Exponents can be processed by this method!",
			}
	}

	//bigIBasePrecision := big.NewInt(int64(exponentPredisionUint))

	//bigINewPrecision := big.NewInt(0).Mul(bigIBasePrecision, exponent.absBigInt)

	baseBigIntNumStr := base.bigInt.Text(10)

	exponentAbsBigIntNumStr := exponent.absBigInt.Text(10)

	bigINewPrecision := big.NewInt(0).Mul(base.bigInt, exponent.absBigInt)

	bigINewPrecisionNumStr := bigINewPrecision.Text(10)

	fmt.Printf("baseBigIntNumStr='%v'\n"+
		"exponentAbsBigIntNumStr='%v'\n"+
		"bigINewPrecisionNumStr='%v'\n", baseBigIntNumStr, exponentAbsBigIntNumStr, bigINewPrecisionNumStr)

	//newPrecision := uint(bigINewPrecision.Int64())

	//result := big.NewInt(0).Exp(base.bigInt, exponent.absBigInt, nil)
	_, result, err := new(BigIntMathUtility).BigIntInverse(base.bigInt)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " _, result, err := new(BigIntMathUtility).\n" +
					"BigIntInverse(base.bigInt)\n",
				ErrContext: fmt.Sprintf("base.bigInt= '%v'",
					base.bigInt.Text(10)),
				ErrMessage: err.Error(),
			}

	}

	bINumResult1, err := new(BigIntNum).NewBigInt(result, uint(bigINewPrecision.Uint64()))

	if err != nil {

		//goland:noinspection ALL
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumResult1, err := new(BigIntNum).NewBigInt(result, newPrecision)",
				ErrContext: fmt.Sprintf("result= '%v'\nnewPrecision= '%v'",
					result.Text(10), bigINewPrecision.Uint64()),
				ErrMessage: err.Error(),
			}
	}

	bINumResult1NumStr, err := bINumResult1.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bINumResult1NumStr, err := bINumResult1.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	inverseResult, err := bINumResult1.GetInverse(maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "inverseResult, err := bINumResult1.GetInverse(maxPrecision)",
				ErrContext: fmt.Sprintf("bINumResult1= '%v'\n maxPrecision= '%v'",
					bINumResult1NumStr, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	err = inverseResult.IsValid(ePrefix.XCpy("Validating 'inverseResult'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = inverseResult.IsValid(\n" +
					"  ePrefix.XCpy(\"Validating 'inverseResult'\").String()",
				ErrContext: "Error: Final Calculation Result 'inverseResult' is INVALID!\n" +
					"'inverseResult' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return inverseResult, nil
}
