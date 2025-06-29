package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathPowerMechanics struct {
	lock sync.Mutex
}

// bigIntNumPwr
//
//	Raises 'base' to the power of 'exponent'.  Both 'base' and
//	'exponent' are Type BigIntNum.
//
//	Upon computing the result of 'base' raised to the power of
//	'exponent' (base^exponent), the result is returned as a Type
//	BigIntNum.
//
//	Examples
//	========
//
//	base    exponent    maxPrecision  result
//
//	 2         4            17        16
//	 2        -4            17         0.0625
//	 3.7       2.8          30        38.991040735983142451443031376258
//	 3.7      -2.8          32         0.02564691737189623971146450249457
//
//	-2		    -3.8          32         0.07179364718731468792491418417362
//	-2         3.8          30        13.928809012737986226180320279676
//
//	The return value, a type BigIntNum, represents the result of the
//	base^exponent operation described above.
//
//	Numeric Separators
//	==================
//
//	This returned BigIntNum 'result' will contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from input parameter,'base'.
func (bIMathPwrMech *bigIntMathPowerMechanics) bigIntNumPwr(
	base *BigIntNum,
	validateBase bool,
	exponent *BigIntNum,
	validateExponent bool,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	bIMathPwrMech.lock.Lock()

	defer bIMathPwrMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerMechanics.bigIntNumPwr",
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

	baseNumSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseNumSeps, err := base.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if !validateBase {

		err = baseNumSeps.IsValid(ePrefix.XCpy("baseNumSeps").String())

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "",
					ErrContext: "Error: Input parameter 'base' is INVALID!\n" +
						"'base' Numeric Separators FAILED validation tests.",
					ErrMessage: err.Error(),
				}
		}
	}

	result, err := new(BigIntNum).NewWithNumSeps(baseNumSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: err.Error(),
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

	baseIsZero, err := base.IsZero()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseIsZero, err := base.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if baseIsZero {
		// If 'base' is zero, return zero value.
		return result, nil
	}

	exponentIsZero, err := exponent.IsZero()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentIsZero, err := exponent.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentIsZero {

		result, err = new(BigIntNum).NewBigIntNumSeps(big.NewInt(1), 0, baseNumSeps)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "result, err = new(BigIntNum).\n" +
						"  NewBigIntNumSeps(big.NewInt(1), 0, baseNumSeps)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return result, nil
	}

	exponentPrecisionUint, err := exponent.GetPrecisionUint()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentPrecisionUint, err := exponent.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bigOne, err := new(BigIntNum).NewBigIntNumSeps(big.NewInt(1), 0, baseNumSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bigOne, err := new(BigIntNum).\n" +
					"  NewBigIntNumSeps(big.NewInt(1), 0, baseNumSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	exponentIsEqualBigOne, err := exponent.Equal(bigOne)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentIsEqualBigOne {

		result, err = base.CopyOut()

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "result, err = base.CopyOut()",
					ErrContext: "exponent == bigOne",
					ErrMessage: err.Error(),
				}
		}

		return result, nil
	}

	exponentSignValue, err := exponent.GetSign()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentSignValue, err := exponent.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentPrecisionUint == uint(0) {

		if exponentSignValue > 0 {

			// exponent is a positive number
			result, err = new(bigIntMathPowerNeutron).
				bigIntNumRaiseToPositiveIntegerPower(base, false, exponent, false, ePrefix)

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix: ePrefix.String(),
						ReturnFunc: "result, err =new(bigIntMathPowerNeutron).\n" +
							"  bigIntNumRaiseToPositiveIntegerPower(\n" +
							"  base, false, exponent, false, ePrefix)",
						ErrContext: fmt.Sprintf("base='%v'\nexponent='%v'",
							baseNumStr, exponentNumStr),
						ErrMessage: err.Error(),
					}
			}

		} else {
			// exponent must be a negative number
			result, err = new(bigIntMathPowerMolecule).
				bigIntNumRaiseToNegativeIntegerPower(
					base, false, exponent, false, maxPrecision, ePrefix)

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix: ePrefix.String(),
						ReturnFunc: "result, err = new(bigIntMathPowerMolecule).\n" +
							"  bigIntNumRaiseToNegativeIntegerPower(\n" +
							"  base, false, exponent, false, maxPrecision, ePrefix)",
						ErrContext: fmt.Sprintf("base='%v'\nexponent='%v'\n"+
							"maxPrecision='%v'",
							baseNumStr, exponentNumStr, maxPrecision),
						ErrMessage: err.Error(),
					}
			}
		}

	} else {
		// precision must be greater than zero. exponent is a fractional number

		if exponentSignValue > 0 {
			// fractional exponent is a positive number

			result, err = new(bigIntMathPowerAtom).
				bigIntNumRaiseToPositiveFractionalPower(
					base, false, exponent, false, maxPrecision, ePrefix)

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix: ePrefix.String(),
						ReturnFunc: "result, err = new(bigIntMathPowerAtom).\n" +
							"  bigIntNumRaiseToPositiveFractionalPower(\n" +
							"  base, false, exponent, false, maxPrecision, ePrefix)",
						ErrContext: fmt.Sprintf("base='%v'\nexponent='%v'\n"+
							"maxPrecision='%v'",
							baseNumStr, exponentNumStr, maxPrecision),
						ErrMessage: err.Error(),
					}
			}

		} else {
			// fractional exponent must be a negative number

			result, err = new(bigIntMathPowerMolecule).bigIntNumRaiseToNegativeFractionalPower(base, false, exponent, false, maxPrecision, ePrefix)

			if err != nil {

				return BigIntNum{},
					&FuncReturnError{
						ErrPrefix: ePrefix.String(),
						ReturnFunc: "result, err = new(bigIntMathPowerMolecule).\n" +
							"  bigIntNumRaiseToNegativeFractionalPower(\n" +
							"  base, false, exponent, false, maxPrecision, ePrefix)",
						ErrContext: fmt.Sprintf("base='%v'\nexponent='%v'\n"+
							"maxPrecision='%v'",
							baseNumStr, exponentNumStr, maxPrecision),
						ErrMessage: err.Error(),
					}
			}
		}
	}

	if result.precision > maxPrecision {

		err = result.SetPrecision(maxPrecision)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = result.SetPrecision(maxPrecision)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = result.SetNumericSeparatorsDto(baseNumSeps)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = result.SetNumericSeparatorsDto(baseNumSeps)",
				ErrContext: fmt.Sprintf("baseNumSeps= '%v'", baseNumSeps.String()),
				ErrMessage: err.Error(),
			}
	}

	err = result.IsValid(ePrefix.XCpy("Validating 'result'").String())

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = result.IsValid(\n" +
					"  ePrefix.XCpy(\"Validating 'result'\").String()",
				ErrContext: "Error: Final calculated 'result' is INVALID!\n" +
					"'result' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}
