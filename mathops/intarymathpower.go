package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
)

type IntAryMathPower struct {
	Input  IntAryPair
	Result IntAry
}

// MinimumRequiredPrecision
//
//	This method is designed to be used with a power function
//	provided below. It will compute the minimum number of decimal
//	places required to support the result of raising a 'base' value
//	to a specified exponent'. Both the 'base' and the 'exponent'
//	are passed to this function as type *IntAry.
//
//	For example, raising the value 3.12 to the power of 4 means
//	that the result will require at least 8-decimal places to the
//	right of the decimal in order to display a correct result. In
//	the following example with base ='3.12' and exponent = '4',
//	this method will return '8'.
//
//	Example
//	=======
//
//	  3.12^4 = 94.75854336 (8-digits to the right of the decimal)
//
//	The calculated minimum required precision is returned as a
//	positive value of type 'int'.
//
//	If the minimum required precision exceeds the maximum positive
//	value for IntAry precision ( +2,147,483,646, which equals 2^31 − 2),
//	an error message is returned in addition to the maximum positive
//	value for IntAry precision (+2,147,483,646).
func (iaPwr *IntAryMathPower) MinimumRequiredPrecision(
	base, exponent *IntAry) (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathPower.MinimumRequiredPrecision()",
		"")

	if err != nil {
		return 0, err
	}

	//maxValue := 2147483646

	signVal := 1

	basePrecisionUint, err := base.GetPrecisionUint()

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "basePrecisionUint, err := base.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	basePrecision, err := new(IntAry).NewUint(basePrecisionUint, signVal, 0)

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "basePrecision, err := new(IntAry).NewUint(basePrecisionUint, signVal, 0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	tExponent, err := exponent.CopyOut()

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "tExponent, err := exponent.CopyOut()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	exponentSignVal, err := tExponent.GetSign()

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentSignVal, err := tExponent.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentSignVal == -1 {

		err = tExponent.ChangeSign()

		if err != nil {

			return 0,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = tExponent.ChangeSign()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	iaResult := IntAry{}

	maxPrecision := tExponent.GetPrecision() + 5

	err = new(IntAryMathMultiply).Multiply(&basePrecision, &tExponent, &iaResult, 0, maxPrecision)

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(IntAryMathMultiply).Multiply(&basePrecision, &tExponent, &iaResult, 0, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	iaResultPrecisionUint, err := iaResult.GetPrecisionUint()

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaResultPrecisionUint, err := iaResult.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if iaResultPrecisionUint > 0 {

		err = iaResult.RoundToPrecision(0)

		if err != nil {

			return 0,
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = iaResult.RoundToPrecision(0)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	intVal, err := iaResult.GetInt()

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "intVal, err := iaResult.GetInt()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return intVal, nil
}

// Pwr
//
//	Raises input parameter 'base' to the power of 'exponent'. This method
//	uses the 'Power By Twos' technique.
//	See:
//	https://en.wikipedia.org/wiki/Exponentiation_by_squaring
//	https://en.wikipedia.org/wiki/Exponentiation_by_squaring#Computation_by_powers_of_2
//
//	This method is based on revised code taken in part from Ye Lin
//	Aung.
//	https://stackoverflow.com/questions/30182129/calculating-large-exponentiation-in-golang
//
//	This algorithm was modified by Mike Rapp to achieve improved
//	performance.
//
//	The result of raising 'base' to the power of 'exponent' will
//	return the result in 'base'. As such the original value of
//	'base' will be overwritten.
//
//	The returned value 'base' will contain the same numeric
//	separators (decimal separator, thousands separator and currency
//	symbol) as that of the original 'base' instance.'base' numeric
//	separators will therefore remain unchanged.
//
//	Example
//	=======
//
//	The 'power' calculation is computed as follows:
//
//	        base = base^exponent
//
//	maxResultPrecision
//	==================
//
//	Input parameter 'maxResultPrecision' will round the result to
//	this number of decimal places after the decimal point if the
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
//
// 'minResultPrecision' will be automatically set to a value of zero.
func (iaPwr *IntAryMathPower) Pwr(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision,
	maxResultPrecision int) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathPower.Pwr",
		"")

	if err != nil {
		return err
	}

	err = base.IsValid(ePrefix.XCpy(" base IntAry Error").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = base.IsValid(ePrefix.XCpy( base IntAry Error).String())",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = exponent.IsValid(ePrefix.XCpy(" exponent IntAry Error").String())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = exponent.IsValid(ePrefix.XCpy(exponent IntAry Error).String())",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	baseIsZero, err := base.IsZero()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "baseIsZero, err :=  base.IsZero()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if baseIsZero {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "base == Zero",
			ErrContext: "",
			ErrMessage: "Error: Input parameter 'base' has a zero value!",
		}
	}

	exponentIsZero, err := exponent.IsZero()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentIsZero, err := exponent.IsZero()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if exponentIsZero {

		err = base.SetIntAryToOne(minResultPrecision)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.SetIntAryToOne(minResultPrecision)",
				ErrContext: fmt.Sprintf("minResultPrecision = '%v'", minResultPrecision),
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	iaOne, err := new(IntAry).NewOne(exponent.GetPrecision())

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "iaOne, err := new(IntAry).NewOne(exponent.GetPrecision())",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentIsEqualIaOne, err := exponent.Equal(&iaOne)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentIsEqualIaOne, err := exponent.Equal(&iaOne)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if exponentIsEqualIaOne {
		return nil
	}

	exponentPrecision := exponent.GetPrecision()

	exponentSign, err := exponent.GetSign()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentSign, err := exponent.GetSign()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	if exponentPrecision == 0 && exponentSign == 1 {
		return new(intAryMathPwrMechanics).pwrTwoPositiveIntegerExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	if exponentPrecision == 0 && exponentSign == -1 {
		return new(intAryMathPwrMechanics).pwrTwoNegativeIntegerExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	if exponentPrecision > 0 && exponentSign == 1 {
		return new(intAryMathPwrMechanics).pwrTwoPositiveFractionalExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	if exponentPrecision > 0 && exponentSign == -1 {
		return new(intAryMathPwrMechanics).pwrTwoNegativeFractionalExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	return &FuncReturnError{
		ErrPrefix:  ePrefix.String(),
		ReturnFunc: ")",
		ErrContext: "",
		ErrMessage: "Error: Input parameters failed to match valid calculation types!",
	}
}

// PwrByMultiplication
//
//		Raises base to the power of exponent using repetitive
//		multiplication. This method may be slower than the method
//		IntAryMathPower.Pwr(); however, this method is capable of
//		handling very large exponents.
//
//		Effectively, input parameter 'base' is raised to the power
//		of exponent.
//
//		    result = base^exponent
//
//	 'power' Operation Result
//	 ========================
//
//		The result of this operation is returned as pointer to an
//		IntAry instance.
//
//	 Numeric Separators
//	 ==================
//
//		The IntAry instance returned by this method will contain
//		numeric separators (decimal separator, thousands separator and
//		currency symbol) copied from input parameter 'base'.
//
//		maxResultPrecision
//		==================
//
//		Input parameter 'maxResultPrecision' will round the result to
//		this number of decimal places after the decimal point if the
//		result is greater than 'maxResultPrecision'.
//
//		If the value of 'maxResultPrecision' is less than zero, it will
//		be automatically reset to a value of '4096'.
//
//		minResultPrecision
//		==================
//
//		Input parameter 'minResultPrecision' signals that if the result
//		precision is less than 'minResultPrecision', zeros will be
//		added to the right of the decimal place in order to implement
//		the 'minResultPrecision' specification.
//
//		If the value of 'minResultPrecision' is less than zero,
//		'minResultPrecision' will be automatically reset to a value of
//		zero.
func (iaPwr *IntAryMathPower) PwrByMultiplication(
	base *IntAry,
	exponent *IntAry,
	minResultPrecision int,
	maxResultPrecision int) (*IntAry, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"IntAryMathPower.PwrByMultiplication",
		"")

	if err != nil {
		return &IntAry{}, err
	}

	iaReturn, err := new(IntAry).NewZero(0)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaReturn, err := new(IntAry).NewZero(0)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	baseNumSeps, err := base.GetNumericSeparatorsDto()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "baseNumSeps, err := base.GetNumericSeparatorsDto()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = iaReturn.SetNumericSeparatorsDto(baseNumSeps)

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = iaReturn.SetNumericSeparatorsDto(baseNumSeps)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = base.IsValid(ePrefix.XCpy("Invalid 'base'").String())

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = base.IsValid(ePrefix.XCpy(Invalid 'base').String())",
				ErrContext: "Error: Input parameter 'base' is INVALID!\n" +
					"'base' FAILED Validation Tests.",
				ErrMessage: err.Error(),
			}
	}

	err = exponent.IsValid(ePrefix.XCpy("Invalid 'exponent'").String())

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = exponent.IsValid(ePrefix.XCpy(Invalid 'exponent').String())",
				ErrContext: "Error: Input parameter 'exponent' is INVALID!\n" +
					"'exponent' FAILED Validation Tests.",
				ErrMessage: err.Error(),
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
				ErrMessage: "Error: Input parameter 'base' is INVALID!\n" +
					"'base' is zero value.",
			}
	}

	exponentIsZero, err := exponent.IsZero()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentIsZero, err := exponent.IsZero()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentIsZero {

		err = iaReturn.SetIntAryToOne(minResultPrecision)

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = iaReturn.SetIntAryToOne(minResultPrecision)",
					ErrContext: fmt.Sprintf("minResultPrecision = '%v'", minResultPrecision),
					ErrMessage: err.Error(),
				}
		}

		return &iaReturn, nil
	}

	iaOne, err := new(IntAry).NewOne(exponent.GetPrecision())

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "iaOne, err := new(IntAry).NewOne(exponent.GetPrecision())",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponent.Equals(&iaOne) {

		iaReturn, err = base.CopyOut()

		if err != nil {

			return &(IntAry{}),
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "iaReturn, err = base.CopyOut()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return &iaReturn, nil
	}

	exponentPrecisionVal := exponent.GetPrecision()

	exponentSignVal, err := exponent.GetSign()

	if err != nil {

		return &(IntAry{}),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentSign, err := exponent.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentPrecisionVal == 0 && exponentSignVal == 1 {

		return new(intAryMathPwrMechanics).pwrMultiplyPositiveIntegerExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	if exponentPrecisionVal == 0 && exponentSignVal == -1 {

		return new(intAryMathPwrMechanics).pwrMultiplyNegativeIntegerExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	if exponentPrecisionVal > 0 && exponentSignVal == 1 {
		return new(intAryMathPwrMechanics).pwrMultiplyPositiveFractionalExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	if exponentPrecisionVal > 0 && exponentSignVal == -1 {
		return new(intAryMathPwrMechanics).pwrMultiplyNegativeFractionalExponent(
			base,
			exponent,
			minResultPrecision,
			maxResultPrecision,
			ePrefix)
	}

	//
	//return &iaReturn,
	//	errors.New(ePrefix + "Error: input parameters failed to match valid calculation types!")

	return &(IntAry{}),
		&FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: "",
			ErrMessage: "Error: Input parameters failed to match valid calculation types!",
		}
}
