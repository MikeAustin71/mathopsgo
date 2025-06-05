package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryMathPwrMechanics struct {
	lock sync.Mutex
}

// pwrMultiplyNegativeFractionalExponent
//
//	 Raises 'base' to the power of 'exponent'.
//
//	 exponent
//	 ========
//
//	 Input parameter 'exponent' is expected to represent a negative
//	 fractional value, i.e., a numeric value with digits to the
//	 right of the decimal place.
//
//	 Input parameter 'exponent' must be negative numeric value which
//	 is less than zero.
//
//	 If 'exponent' is a positive value, an error will be returned.
//
//	 If 'exponent' is a numeric value with no digits to the right of
//	 the decimal place, it is by definition an integer value and
//	 NOT a fractional value. In that case, an error will be
//	 returned.
//
//	 maxResultPrecision
//	 ==================
//
//	 Input parameter 'maxResultPrecision' will round the result to
//	 that number of decimal places after the decimal point if the
//	 result is greater than 'maxResultPrecision'.
//
//	 If the value of 'maxResultPrecision' is less than zero, it will
//	 be automatically set to a value of '4096'.
//
//	 minResultPrecision
//	 ==================
//
//	 Input parameter 'minResultPrecision' signals that if the result
//	 precision is less than 'minResultPrecision', zeros will be
//	 added to the right of the decimal place in order to implement
//	 the 'minResultPrecision' specification.
//
//	 If the value of 'minResultPrecision' is less than zero,
//	 'minResultPrecision' will be automatically set to a value of
//	 zero.
//
//	 'power' Operation Result
//	 ========================
//
//	 The result of this power operation is returned as a pointer
//	 to a new 'result' IntAry. None of the input parameters are
//	 altered by this operation.
//
//	 If the value of input parameter 'base' is zero, an error will
//	 be returned.
//
//		This method uses simple multiplication to generate the result.
//
//	 Numeric Separators
//	 ==================
//
//	 The returned 'result' IntAry will contain numeric separators
//	 (decimal separator, thousands separator and currency symbol)
//	 copied from input parameter 'base'.
func (iaMathMech *intAryMathPwrMechanics) pwrMultiplyNegativeFractionalExponent(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision int,
	maxResultPrecision int,
	errPrefDto *ePref.ErrPrefixDto) (*IntAry, error) {

	iaMathMech.lock.Lock()

	defer iaMathMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathPwrMechanics.pwrMultiplyNegativeFractionalExponent",
		"")

	if err != nil {
		return &(IntAry{}), err
	}

	if base == nil {

		return &(IntAry{}),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'base'",
			}
	}

	if exponent == nil {

		return &(IntAry{}),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'exponent'",
			}
	}

	baseIsZero, err := base.IsZero()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseIsZero, err := base.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if baseIsZero {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Input parameter 'base' is INVALID!\n" +
					"The value of 'base' is zero.",
			}
	}

	exponentGetSignVal, err := exponent.GetSign()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentGetSignVal, err := exponent.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
				ErrMessage: err.Error(),
			}
	}

	if exponentGetSignVal != -1 {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if exponentGetSignVal != -1",
				ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a negative fractional value.\n"+
					"Instead, 'exponent' is a positive value!\n"+
					"exponent='%v'", exponentNumStr),
			}
	}

	exponentPrecisionVal := exponent.GetPrecision()

	if exponentPrecisionVal < 1 {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if exponentGetSignVal != -1",
				ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a fractional value.\n"+
					"Instead, 'exponent' is an integer value!\n"+
					"exponent='%v'", exponentNumStr),
			}
	}

	numSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := base.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if maxResultPrecision < 0 {
		maxResultPrecision = 4096
	}

	if minResultPrecision < 0 {
		minResultPrecision = 0
	}

	if minResultPrecision > maxResultPrecision {
		minResultPrecision = maxResultPrecision
	}

	internalMaxPrecision := maxResultPrecision + 100

	newExponent, err := exponent.CopyOut()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newExponent, err := exponent.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// Change sign from minus to plus
	err = newExponent.ChangeSign()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = newExponent.ChangeSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	newInverseBase, err := base.Inverse(internalMaxPrecision)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "newInverseBase, err := base.Inverse(internalMaxPrecision)",
				ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'",
					internalMaxPrecision),
				ErrMessage: err.Error(),
			}
	}

	fracIntAry, err := new(FracIntAry).NewFracIntAry(&newExponent)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracIntAry, err := new(FracIntAry).\n" +
					"  NewFracIntAry(&newExponent)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	internalMaxPrecision += 5

	err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = fracIntAry.ReduceToLowestCommonDenom(\n" +
					"  internalMaxPrecision)",
				ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'",
					internalMaxPrecision),
				ErrMessage: err.Error(),
			}
	}

	internalMaxPrecision += 5

	newBase, err := new(NthRootOp).NewNthRoot(&newInverseBase, &fracIntAry.Denominator, internalMaxPrecision)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "newBase, err := NthRootOp{}.NewNthRoot(\n" +
					"  &newInverseBase, &fracIntAry.Denominator, internalMaxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	result, err := new(IntAry).NewOne(0)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "result, err := new(IntAry).NewOne(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	internalMaxPrecision += 5

	fIAryNumeratorSignVal, err := fracIntAry.Numerator.GetSign()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fIAryNumeratorSignVal, err := fracIntAry.Numerator.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fIAryNumeratorPrecisionVal := fracIntAry.Numerator.GetPrecision()

	if fIAryNumeratorPrecisionVal != 0 {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if fracIntAry Numerator Precision != 0",
				ErrMessage: fmt.Sprintf("Error: FracIntAry Numerator Precision is NOT Equal to Zero!\n"+
					"FracIntAry Numerator Precision = '%v'", fIAryNumeratorPrecisionVal),
			}

	}

	if fIAryNumeratorSignVal == -1 {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if FracIntAry Numerator Sign Value == -1",
				ErrMessage: fmt.Sprintf("Error: FracIntAry Numerator Sign Value is -1.\n"+
					"FracIntAry Numerator is a negative number.\n"+
					"FracIntAry Numerator Sign Value = '%v'", fIAryNumeratorSignVal),
			}

	}

	// fracIntAry.Numerator == new exponent

	fracIntAryNumeratorIsZero, err := fracIntAry.Numerator.IsZero()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracIntAryNumeratorIsZero, err := fracIntAry.Numerator.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	var idx = 0

	for !fracIntAryNumeratorIsZero {

		err = new(IntAryMathMultiply).MultiplyInPlace(&result, &newBase, minResultPrecision, internalMaxPrecision)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = new(IntAryMathMultiply).MultiplyInPlace(&result, &newBase, minResultPrecision, internalMaxPrecision)",
					ErrContext: fmt.Sprintf("Cycle Index= '%v'", idx),
					ErrMessage: err.Error(),
				}
		}

		err = fracIntAry.Numerator.DecrementIntegerOne()

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = fracIntAry.Numerator.DecrementIntegerOne()",
					ErrContext: fmt.Sprintf("Cycle Index= '%v'", idx),
					ErrMessage: err.Error(),
				}
		}

		fracIntAryNumeratorIsZero, err = fracIntAry.Numerator.IsZero()

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "fracIntAryNumeratorIsZero, err := fracIntAry.Numerator.IsZero()",
					ErrContext: fmt.Sprintf("Cycle Index= '%v'", idx),
					ErrMessage: err.Error(),
				}
		}

		idx++
	}

	if result.GetPrecision() > maxResultPrecision {

		err = result.RoundToPrecision(maxResultPrecision)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = result.RoundToPrecision(maxResultPrecision)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	if result.GetPrecision() < minResultPrecision {

		err = result.SetPrecision(minResultPrecision, false)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = result.SetPrecision(minResultPrecision, false)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	err = result.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "err = result.SetNumericSeparatorsDto(numSeps)",
				ErrMessage: err.Error(),
			}
	}

	return &result, nil
}

// pwrMultiplyPositiveFractionalExponent
//
//	Raises 'base' to the power of 'exponent'. Input parameter
//	'exponent' is expected to represent a positive fractional
//	value, i.e., a positive numeric value with digits to the right
//	of the decimal place.
//
//	        result = base^exponent
//
//	exponent
//	========
//
//	If 'exponent' is NOT a positive value, i.e., a value less than
//	zero, an error will be returned.
//
//	If 'exponent' is a numeric value with no digits to the right of
//	the decimal place, it is by definition an integer value and NOT
//	a fractional value. In that case, an error will be returned.
//
//	maxResultPrecision
//	==================
//
//	Input parameter 'maxResultPrecision' will round the result to
//	this number of decimal places after the decimal point if the
//	result is greater than 'maxResultPrecision'.  If the value of
//	'maxResultPrecision' is less than zero, it will be
//	automatically reset to a value of '4096'.
//
//	minResultPrecision
//	==================
//
//	Input parameter 'minResultPrecision' signals that if the result
//	precision is less than 'minResultPrecision', zeros will be
//	added to the right of the decimal place in order to implement
//	the 'minResultPrecision' specification. If the value of
//	'minResultPrecision' is less than zero, 'minResultPrecision'
//	will be automatically reset to a value of zero.
//
//	'power' Operation Result
//	========================
//
//	The result of this power operation is returned as a pointer to
//	a new 'result' IntAry. None of the input parameters are altered
//	by this operation.
//
//	This method uses simple multiplication to generate the result.
//
//	Numeric Separators
//	==================
//
//	The returned 'result' IntAry will contain numeric separators
//	(decimal separator, thousands separator and currency symbol)
//	copied from input parameter 'base'.
func (iaMathMech *intAryMathPwrMechanics) pwrMultiplyPositiveFractionalExponent(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision int,
	maxResultPrecision int,
	errPrefDto *ePref.ErrPrefixDto) (*IntAry, error) {

	iaMathMech.lock.Lock()

	defer iaMathMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathPwrMechanics.pwrMultiplyPositiveFractionalExponent",
		"")

	if err != nil {
		return &(IntAry{}), err
	}

	if base == nil {

		return &(IntAry{}),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'base'",
			}
	}

	if exponent == nil {

		return &(IntAry{}),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'exponent'",
			}
	}

	exponentGetSignVal, err := exponent.GetSign()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentGetSignVal, err := exponent.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
				ErrMessage: err.Error(),
			}
	}

	if exponentGetSignVal != 1 {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if exponentGetSignVal != -1",
				ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a positive fractional value.\n"+
					"Instead, 'exponent' is a negative value!\n"+
					"exponent='%v'", exponentNumStr),
			}
	}

	exponentPrecisionVal := exponent.GetPrecision()

	if exponentPrecisionVal < 1 {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if exponentGetSignVal != -1",
				ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a fractional value.\n"+
					"Instead, 'exponent' is an integer value!\n"+
					"exponent='%v'", exponentNumStr),
			}
	}

	baseIsZero, err := base.IsZero()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseIsZero, err := base.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	basePrecisionVal := base.GetPrecision()

	result, err := new(IntAry).NewZero(uint(basePrecisionVal))

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "result, err := new(IntAry).NewZero(uint(basePrecisionVal))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if baseIsZero {

		return &result, nil

	}

	baseIsOne, err := base.IsOne()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseIsOne, err := base.IsOne()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if baseIsOne {

		result, err = new(IntAry).NewOne(basePrecisionVal)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "result, err = new(IntAry).NewOne(basePrecisionVal)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return &result, nil
	}

	numSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "numSeps, err := base.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if maxResultPrecision < 0 {
		maxResultPrecision = 4096
	}

	if minResultPrecision < 0 {
		minResultPrecision = 0
	}

	if minResultPrecision > maxResultPrecision {
		minResultPrecision = maxResultPrecision
	}

	internalMaxPrecision := maxResultPrecision + 100

	fracIntAry, err := new(FracIntAry).NewFracIntAry(exponent)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracIntAry, err := new(FracIntAry).NewFracIntAry(exponent)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	internalMaxPrecision += 5

	newBase, err := new(NthRootOp).NewNthRoot(base, &fracIntAry.Denominator, internalMaxPrecision)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "newBase, err := NthRootOp{}.NewNthRoot(\n" +
					"  base, &fracIntAry.Denominator, internalMaxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	result, err = new(IntAry).NewOne(0)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "result, err = new(IntAry).NewOne(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	internalMaxPrecision += 5

	fracIntAryNumeratorSignVal, err := fracIntAry.Numerator.GetSign()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracIntAryNumeratorSignVal, err :=  fracIntAry.Numerator.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracIntAryNumeratorNumStr, err := fracIntAry.Numerator.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if fracIntAry.Numerator.GetPrecision() != 0 ||
		fracIntAryNumeratorSignVal == -1 {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: fmt.Sprintf("Error: fracIntAry.Numerator (new exponent) INVALID!\n"+
					"fracIntAry.Numerator='%v' ", fracIntAryNumeratorNumStr),
			}
	}

	fracIntAryNumeratorIsZero, err := fracIntAry.Numerator.IsZero()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fracIntAryNumeratorIsZero, err = fracIntAry.Numerator.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	// fracIntAry.Numerator == new exponent
	for !fracIntAryNumeratorIsZero {

		err = new(IntAryMathMultiply).MultiplyInPlace(&result, &newBase, minResultPrecision, internalMaxPrecision)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "new(IntAryMathMultiply).MultiplyInPlace(\n" +
						"  &result, &newBase, minResultPrecision, internalMaxPrecision)",
					ErrContext: fmt.Sprintf("minResultPrecision= '%v' internalMaxPrecision= '%v'",
						minResultPrecision, internalMaxPrecision),
					ErrMessage: err.Error(),
				}
		}

		err = fracIntAry.Numerator.DecrementIntegerOne()

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = fracIntAry.Numerator.DecrementIntegerOne()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	} // End of for loop

	if result.GetPrecision() > maxResultPrecision {

		err = result.RoundToPrecision(maxResultPrecision)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = result.RoundToPrecision(maxResultPrecision)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	if result.GetPrecision() < minResultPrecision {

		err = result.SetPrecision(minResultPrecision, false)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = result.SetPrecision(minResultPrecision, false)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	err = result.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = result.SetNumericSeparatorsDto(numSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return &result, nil
}

// pwrTwoNegativeFractionalExponent
//
//	Raises 'base' to the power of 'exponent'.
//
//	Input parameter 'exponent' is expected to represent a negative
//	fractional value, i.e., a value less than zero with digits to
//	the right of the decimal point.
//
//	If 'exponent' is a positive value, i.e., a greater than -1, an
//	error will be returned.
//
//	Also, if 'exponent' has a precision value less than '1', an
//	error will be returned.
//
//	Input parameter 'maxResultPrecision' will round the result to
//	this number of decimal places after the decimal point if the
//	result is greater than 'maxResultPrecision'.  If the value of
//	'maxResultPrecision' is less than zero, it will be
//	automatically set to a value of '4096'.
//
//	Input parameter 'minResultPrecision' signals that if the result
//	precision is less than 'minResultPrecision', zeros will be
//	added to the right of the decimal place in order to implement
//	the 'minResultPrecision' specification. If the value of
//	'minResultPrecision' is less than zero, 'minResultPrecision'
//	will be automatically set to a value of zero.
//
//	The result of the power operation is returned in the input
//	parameter 'base'. During this procedure the original value of
//	'base' is destroyed.
//
//	The returned IntAry object 'base' will contain same numeric
//	separators (i.e. decimal separator, thousands separator and
//	currency symbol) as those in the original 'base' instance. As
//	such, the 'base' numeric separators will remain unchanged.
func (iaMathMech *intAryMathPwrMechanics) pwrTwoNegativeFractionalExponent(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision,
	maxResultPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMathMech.lock.Lock()

	defer iaMathMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathPwrMechanics.pwrTwoNegativeFractionalExponent",
		"")

	if err != nil {
		return err
	}

	if base == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'base'",
		}
	}

	if exponent == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'exponent'",
		}
	}

	exponentGetSignVal, err := exponent.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentGetSignVal, err := exponent.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
			ErrMessage: err.Error(),
		}
	}

	if exponentGetSignVal != -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if exponent Sign Value != 1",
			ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a positive fractional value. "+
				"Instead, 'exponent' is negative! exponent='%v'", exponentNumStr),
		}
	}

	exponentPrecisionVal := exponent.GetPrecision()

	if exponentPrecisionVal < 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if exponent precision Value < 1",
			ErrMessage: fmt.Sprintf("Error: 'exponent' precision value is expected to be greater than zero.\n"+
				"precision= '%v'", exponentPrecisionVal),
		}
	}

	numSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "numSeps, err := base.GetNumericSeparatorsDto()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if maxResultPrecision < 0 {
		maxResultPrecision = 4096
	}

	if minResultPrecision < 0 {
		minResultPrecision = 0
	}

	if minResultPrecision > maxResultPrecision {
		minResultPrecision = maxResultPrecision
	}

	internalMaxPrecision := maxResultPrecision + 100

	// Set exponent to a positive value
	newExponent, err := exponent.CopyOut()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "newExponent, err := exponent.CopyOut()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = newExponent.ChangeSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = newExponent.ChangeSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	newInverseBase, err := base.Inverse(internalMaxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "newInverseBase, err := base.Inverse(internalMaxPrecision)",
			ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'",
				internalMaxPrecision),
			ErrMessage: err.Error(),
		}
	}

	fracIntAry, err := new(FracIntAry).NewFracIntAry(&newExponent)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "fracIntAry, err := new(FracIntAry).NewFracIntAry(&newExponent)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	internalMaxPrecision += 5

	err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)",
			ErrContext: fmt.Sprintf("internalMaxPrecision= '%v'",
				internalMaxPrecision),
			ErrMessage: err.Error(),
		}
	}

	internalMaxPrecision += 5

	newBase, err := new(NthRootOp).NewNthRoot(&newInverseBase, &fracIntAry.Denominator, internalMaxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "newBase, err := NthRootOp{}.NewNthRoot(\n" +
				"&newInverseBase, &fracIntAry.Denominator, internalMaxPrecision)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	expBigInt, err := fracIntAry.Numerator.GetBigInt()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "expBigInt, err := fracIntAry.Numerator.GetBigInt()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	internalMaxPrecision += 5

	err = new(intAryMathPwrNanobot).pwrByTwos(&newBase, expBigInt, maxResultPrecision, internalMaxPrecision, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryMathPwrNanobot).pwrByTwos(\n" +
				"  &newBase, expBigInt, maxResultPrecision, internalMaxPrecision,\n" +
				"  ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = base.CopyIn(&newBase, false)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.CopyIn(&newBase, false)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if base.GetPrecision() > maxResultPrecision {

		err = base.RoundToPrecision(maxResultPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.RoundToPrecision(maxResultPrecision)",
				ErrContext: fmt.Sprintf("maxResultPrecision= '%v'",
					maxResultPrecision),
				ErrMessage: err.Error(),
			}
		}

	}

	if base.GetPrecision() < minResultPrecision {

		err = base.SetPrecision(minResultPrecision, false)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.SetPrecision(minResultPrecision, false)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	err = base.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.SetNumericSeparatorsDto(numSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// pwrTwoNegativeIntegerExponent
//
//	Raises 'base' to the power of 'exponent'.
//
//	exponent
//	========
//	Input parameter 'exponent' is expected to represent a negative
//	integer value, i.e., a value less than zero.
//
//	If 'exponent' is a positive integer (value greater than -1), an
//	error will be returned.
//
//	If 'exponent' is a fractional value, i.e., it has digits to the
//	right of the decimal place, it is by definition NOT an integer
//	value and therefore, an error will be returned.
//
//	maxResultPrecision
//	==================
//
//	Input parameter 'maxResultPrecision' will round the result to
//	that number of decimal places after the decimal point if the
//	result is greater than 'maxResultPrecision'.  If the value of
//	'maxResultPrecision' is less than zero, it will be
//	automatically set to a value of '4096'.
//
//	minResultPrecision
//	==================
//
//	Input parameter 'minResultPrecision' signals that if the result
//	precision is less than 'minResultPrecision', zeros will be
//	added to the right of the decimal place in order to implement
//	the 'minResultPrecision' specification. If the value of
//	'minResultPrecision' is less than zero, 'minResultPrecision'
//	will be automatically set to a value of zero.
//
//	Result of the 'power' Operation
//	===============================
//
//	The result of the power operation is returned in the input
//	parameter, 'base'. During this procedure the original value
//	of 'base' is destroyed.
//
//	Numeric Separators
//	==================
//
//	The returned IntAry object 'base' will contain same numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) as those in the original 'base' instance. As such, the
//	'base' numeric separators will remain unchanged.
func (iaMathMech *intAryMathPwrMechanics) pwrTwoNegativeIntegerExponent(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision int,
	maxResultPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMathMech.lock.Lock()

	defer iaMathMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathPwrMechanics.pwrTwoNegativeIntegerExponent",
		"")

	if err != nil {
		return err
	}

	if base == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'base'",
		}
	}

	if exponent == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'exponent'",
		}
	}

	exponentSignVal, err := exponent.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentSignVal, err := exponent.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
			ErrMessage: err.Error(),
		}
	}

	if exponentSignVal != -1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if exponentSignVal != -1",
			ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a negative integer.\n"+
				"Instead, 'exponent' is positive!\nexponent='%v'",
				exponentNumStr),
		}
	}

	if exponent.GetPrecision() != 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if exponent.GetPrecision() != 0",
			ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be an integer value.\n"+
				"Instead, 'exponent' is a fractional value!\n"+
				"exponent='%v'", exponentNumStr),
		}
	}

	numSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "numSeps, err := base.GetNumericSeparatorsDto()",
			ErrMessage: err.Error(),
		}
	}

	if maxResultPrecision < 0 {
		maxResultPrecision = 4096
	}

	if minResultPrecision < 0 {
		minResultPrecision = 0
	}

	if minResultPrecision > maxResultPrecision {
		minResultPrecision = maxResultPrecision
	}

	// This is an integer Exponent!
	bInt, err := exponent.GetBigInt()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bInt, err := exponent.GetBigInt()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	internalPrecision := maxResultPrecision + 100

	err = new(intAryMathPwrNanobot).pwrByTwos(base, bInt, maxResultPrecision, internalPrecision, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryMathPwrNanobot).pwrByTwos(\n" +
				"  base, bInt, maxResultPrecision, internalPrecision, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if base.GetPrecision() > maxResultPrecision {

		err = base.RoundToPrecision(maxResultPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.RoundToPrecision(maxResultPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	if base.GetPrecision() < minResultPrecision {

		err = base.SetPrecision(minResultPrecision, false)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.SetPrecision(minResultPrecision, false)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}
	}

	err = base.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.SetNumericSeparatorsDto(numSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// pwrTwoPositiveFractionalExponent
//
//		Raises 'base' to the power of 'exponent'.
//
//		   Result = base^exponent
//
//		exponent
//		========
//
//		Input parameter 'exponent' is expected to represent a positive
//		fractional value, i.e., a value greater than -1 with digits to
//		the right of the decimal point.
//
//		If 'exponent' is a negative value, i.e., a value less than
//		zero, an error will be returned.
//
//		Also, if 'exponent' has a precision value less than '1', an
//		error will be returned.
//
//		maxResultPrecision
//		==================
//
//		Input parameter 'maxResultPrecision' will round the result to
//		this number of decimal places after the decimal point if the
//		result is greater than 'maxResultPrecision'.  If the value of
//		'maxResultPrecision' is less than zero, it will be
//		automatically set to a value of '4096'.
//
//		minResultPrecision
//		==================
//
//		Input parameter 'minResultPrecision' signals that if the result
//		precision is less than 'minResultPrecision', zeros will be
//		added to the right of the decimal place in order to implement
//		the 'minResultPrecision' specification. If the value of
//		'minResultPrecision' is less than zero, 'minResultPrecision'
//		will be automatically set to a value of zero.
//
//	 'power' Operation Result
//	 ========================
//		The result of the power operation is returned in the input
//		parameter 'base'. During this procedure the original value
//		of 'base' is destroyed.
//
//		Numeric Separators
//		==================
//
//		The returned IntAry object 'base' will contain same numeric
//		separators ( i.e. decimal separator, thousands separator and
//		currency symbol) as those in the original 'base' instance. As
//		such, the 'base' numeric separators will remain unchanged.
func (iaMathMech *intAryMathPwrMechanics) pwrTwoPositiveFractionalExponent(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision int,
	maxResultPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMathMech.lock.Lock()

	defer iaMathMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathPwrMechanics.pwrTwoPositiveFractionalExponent",
		"")

	if err != nil {
		return err
	}

	if base == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'base'",
		}
	}

	if exponent == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'exponent'",
		}
	}

	exponentSignVal, err := exponent.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentSignVal, err := exponent.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
			ErrMessage: err.Error(),
		}
	}

	if exponentSignVal != 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if exponent Sign Value != 1",
			ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a positive fractional value. "+
				"Instead, 'exponent' is negative! exponent='%v'", exponentNumStr),
		}
	}

	exponentPrecisionVal := exponent.GetPrecision()

	if exponentPrecisionVal < 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if exponent precision Value < 1",
			ErrMessage: fmt.Sprintf("Error: 'exponent' precision value is expected to be greater than zero.\n"+
				"precision= '%v'", exponentPrecisionVal),
		}
	}

	numSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "numSeps, err := ia.GetNumericSeparatorsDto()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if maxResultPrecision < 0 {
		maxResultPrecision = 4096
	}

	if minResultPrecision < 0 {
		minResultPrecision = 0
	}

	if minResultPrecision > maxResultPrecision {
		minResultPrecision = maxResultPrecision
	}

	internalMaxPrecision := maxResultPrecision + 100

	// returns exponent / 1
	fracIntAry, err := new(FracIntAry).NewFracIntAry(exponent)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "fracIntAry, err := new(FracIntAry).NewFracIntAry(exponent)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = fracIntAry.ReduceToLowestCommonDenom(internalMaxPrecision)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	internalMaxPrecision += 5

	newBase, err := new(NthRootOp).NewNthRoot(base, &fracIntAry.Denominator, internalMaxPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "newBase, err := NthRootOp{}.NewNthRoot(...)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = base.CopyIn(&newBase, false)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.CopyIn(&newBase, false)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	expBigInt, err := fracIntAry.Numerator.GetBigInt()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "expBigInt, err := fracIntAry.Numerator.GetBigInt()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	internalMaxPrecision += 5

	err = new(intAryMathPwrNanobot).
		pwrByTwos(base, expBigInt, maxResultPrecision, internalMaxPrecision, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryMathPwrNanobot).pwrByTwos(\n" +
				"  base, expBigInt, maxResultPrecision, internalMaxPrecision, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if base.GetPrecision() > maxResultPrecision {

		err = base.RoundToPrecision(maxResultPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.RoundToPrecision(maxResultPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	if base.GetPrecision() < minResultPrecision {

		err = base.SetPrecision(minResultPrecision, false)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.SetPrecision(minResultPrecision, false)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	err = base.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.SetNumericSeparatorsDto(numSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// pwrTwoPositiveIntegerExponent
//
//	raises 'base' to the power of 'exponent'.
//
//	exponent
//	========
//
//	Input parameter 'exponent' is expected to represent a positive
//	integer. If 'exponent' is NOT a positive integer, an error will
//	be returned.
//
//	If 'exponent' is a fractional value, i.e., it has digits to the
//	right of the decimal place, it is by definition NOT an integer
//	value and therefore, an error will be returned.
//
//	maxResultPrecision
//	==================
//
//	Input parameter 'maxResultPrecision' will round the result to
//	that number of decimal places after the decimal point if the
//	result is greater than 'maxResultPrecision'.
//
//	If the value of 'maxResultPrecision' is less than zero, it will
//	be automatically set to a value of '4096'.
//
//	minResultPrecision
//	==================
//
//	Input parameter 'minResultPrecision' signals that if the result
//	precision is less than 'minResultPrecision', zeros will be added
//	to the right of the decimal place in order to implement the
//	'minResultPrecision' specification.
//
//	If the value of 'minResultPrecision' is less than zero,
//	'minResultPrecision' will be automatically set to a value of zero.
//
//	'power' Operation Result
//	========================
//
//	The result of the power operation is returned in the input
//	parameter 'base'. During this procedure the original value of
//	'base' is destroyed.
//
//	Numeric Separators
//	==================
//
//	The returned IntAry object 'base' will contain same numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) as those in the original 'base' instance. As such, the
//	'base' numeric separators will remain unchanged.
func (iaMathMech *intAryMathPwrMechanics) pwrTwoPositiveIntegerExponent(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision int,
	maxResultPrecision int,
	errPrefDto *ePref.ErrPrefixDto) error {

	iaMathMech.lock.Lock()

	defer iaMathMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryMathPwrMechanics.pwrTwoPositiveIntegerExponent",
		"")

	if err != nil {
		return err
	}

	if base == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'base'",
		}
	}

	if exponent == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'exponent'",
		}
	}

	exponentSignVal, err := exponent.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentSignVal, err := exponent.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "exponentNumStr, err := exponent.GetNumStr()",
			ErrMessage: err.Error(),
		}
	}

	if exponentSignVal != 1 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be a positive integer.\n"+
				"Instead, 'exponent' is negative!\n"+
				"exponent='%v'",
				exponentNumStr),
		}
	}

	exponentPrecisionVal := exponent.GetPrecision()

	if exponentPrecisionVal != 0 {
		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "if exponent precision Value < 1",
			ErrMessage: fmt.Sprintf("Error: 'exponent' is expected to be an integer value.\n"+
				"Instead, 'exponent' is a fractional value!\n"+
				"exponent='%v'", exponentNumStr),
		}
	}

	numSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "numSeps, err := base.GetNumericSeparatorsDto()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if maxResultPrecision < 0 {
		maxResultPrecision = 4096
	}

	if minResultPrecision < 0 {
		minResultPrecision = 0
	}

	if minResultPrecision > maxResultPrecision {
		minResultPrecision = maxResultPrecision
	}

	// This is an integer Exponent!
	bInt, err := exponent.GetBigInt()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bInt, err := exponent.GetBigInt()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	internalPrecision := maxResultPrecision + 100

	err = new(intAryMathPwrNanobot).pwrByTwos(base, bInt, maxResultPrecision, internalPrecision, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryMathPwrNanobot).\n" +
				"  pwrByTwos(base, bInt, maxResultPrecision, internalPrecision,\n" +
				"    ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if base.GetPrecision() > maxResultPrecision {

		err = base.RoundToPrecision(maxResultPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.RoundToPrecision(maxResultPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	if base.GetPrecision() < minResultPrecision {

		err = base.SetPrecision(minResultPrecision, false)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
		}

	}

	err = base.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.SetNumericSeparatorsDto(numSeps)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
