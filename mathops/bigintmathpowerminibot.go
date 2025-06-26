package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathPowerMinibot struct {
	lock sync.Mutex
}

// bigIntToPositiveFractionalPower - Raises 'base' to the power
// of a positive fractional exponent, 'exponent'.
//
// As stated in the function name, this method expects to process only
// positive exponent values which have fractional digits to the right
// of the decimal place.
//
// Examples:
// base	basePrecision	exponent exponentPrecision result  resultPrecision
//
//		  5			0					   22						1						 34493			 3 (to 3-decimal places)
//	   (5^2.2 = 34.4932415365303708097515866054
//	    The actual number of decimal places returned in the result is controlled by
//	    input parameter, 'maxPrecision'.)
//
//	  18     1            34            1					 7378				 3 (to 3-decimal places)
//	 (1.8^3.4 = 7.3777938725727533349996174917827
//	    The actual number of decimal places returned in the result is controlled by
//	    input parameter, 'maxPrecision'.)
//
// Input Parameters
// ================
//
// base 					*big.Int	- The base which will be raised to the
//
//	power of a positive integer exponent.
//
// basePrecision 			uint	- The precision specification for 'base'.
//
//		Precision defines the number of numeric
//		digits to the right of the decimal place.
//		Taken together, 'base' and 'basePrecision'
//		define a numeric value with a fixed number
//	  of decimal digits to the right of the decimal
//	  place.
//
// exponent				*big.Int	- The exponent to which 'base' will be raised
//
//		by this calculation. By method definition,
//	  'exponent' must be a positive value. If
//	  exponent is less than 0, an error will
//	  be triggered.
//
// exponentPrecision 	uint	- The precision specification for 'exponent'.
//
//	                         	Precision defines the number of numeric
//	                         	digits to the right of the decimal place.
//	                         	Taken together, 'base' and 'basePrecision'
//	                         	define a numeric value with a fixed number
//	                           of decimal digits to the right of the decimal
//	                           place. For this method, 'exponentPrecision'
//															must be greater than zero, thereby designating
//															'exponent' as a fractional value. A value of
//															zero for 'exponentPrecision' will trigger an error.
//
// maxPrecision				uint	- When this method calculates 'base' raised to the
//
//	power of 'exponent', the maximum number of decimal
//	digits to the right of the decimal place in the
//	result will be limited by 'maxPrecision'.
//
// Return Values
// =============
//
// result					*big.Int	- If the calculation completes successfully, this
//
//	return value will be populated with the value
//	of 'base' raised to the power of 'exponent'.
//
// resultPrecision		uint	- The precision specification for 'result'.
//
//		Precision defines the number of numeric
//		digits to the right of the decimal place.
//		Taken together, 'result' and 'resultPrecision'
//		define a numeric value with a fixed number
//	  of decimal digits to the right of the decimal
//	  place.
//
// err							 error	- If the calculation encounters an error, an appropriate
//
//	error message will be formatted and returned. If the
//	calculation completes successfully, this return value
//	will be set to 'nil'.
func (bIMathPwrMinibot *bigIntMathPowerMinibot) bigIntToPositiveFractionalPower(
	base *big.Int,
	basePrecision *big.Int,
	exponent *big.Int,
	exponentPrecision *big.Int,
	maxPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (result *big.Int, resultPrecision *big.Int, err error) {

	bIMathPwrMinibot.lock.Lock()

	defer bIMathPwrMinibot.lock.Unlock()

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerMinibot.bigIntToPositiveFractionalPower",
		"")

	if err != nil {
		return result, resultPrecision, err
	}

	if base == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'base'",
			}
	}

	if basePrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'basePrecision'",
			}
	}

	if exponent == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'exponent'",
			}
	}

	if exponentPrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'exponentPrecision'",
			}
	}

	if maxPrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'maxPrecision'",
			}
	}

	bigZero := big.NewInt(0)

	bigOne := big.NewInt(1)

	internalMaxPrecision := big.NewInt(0).Mul(maxPrecision, big.NewInt(2))

	if base.Cmp(bigZero) == 0 {
		// base is zero result is zero
		return result, resultPrecision, err
	}

	basePrecisionCmpZero := basePrecision.Cmp(bigZero)

	if base.Cmp(bigOne) == 0 &&
		basePrecisionCmpZero == 0 {

		result.Set(bigOne)

		return result, resultPrecision, err
	}

	if basePrecisionCmpZero == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("basePrecision= '%v'",
					basePrecision.Text(10)),
				ErrMessage: "Error: 'basePrecision' is INVALID!\n" +
					"'basePrecision' is a negative value.\n" +
					"Only positive 'basePrecision' values can be used with this method.",
			}
	}

	exponentPrecisionCmpZero := exponentPrecision.Cmp(bigZero)

	if exponentPrecisionCmpZero == 0 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'\nexponentPrecision= '%v'",
					exponent.Text(10), exponentPrecision.Text(10)),
				ErrMessage: "Error: 'exponentPrecision' is INVALID!\n" +
					"'exponentPrecision' is zero. This is an integer exponent.\n" +
					"Only positive fractional exponent values can be used with this method.",
			}
	}

	if exponentPrecisionCmpZero == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'\nexponentPrecision= '%v'",
					exponent.Text(10), exponentPrecision.Text(10)),
				ErrMessage: "Error: 'exponentPrecision' is INVALID!\n" +
					"'exponentPrecision' is a negative value.\n" +
					"Only positive fractional exponent values can be used with this method.",
			}
	}

	// exponentPrecision > 0
	cmpExponentZero := exponent.Cmp(bigZero)

	if cmpExponentZero == -1 {
		err = fmt.Errorf(ePrefix+
			"Error: 'exponent' is a negative value! "+
			"exponent='%v' ", exponent.Text(10))

		return result, resultPrecision, err
	}

	if cmpExponentZero == 0 {
		// Any number raised to a zero power is one
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	bigTen := big.NewInt(10)
	// Get exponent integer value
	scale := big.NewInt(0).Exp(
		bigTen,
		exponentPrecision,
		nil)

	scratch := big.NewInt(0)

	integerExponent, fractionalExponent := big.NewInt(0).QuoRem(exponent, scale, scratch)

	integerResult := big.NewInt(1)
	integerPrecision := big.NewInt(0)

	if integerExponent.Cmp(bigZero) != 0 {

		integerResult, integerPrecision, err =
			new(bigIntMathPowerNanobot).bigIntToPositiveIntegerPower(
				base,
				basePrecision,
				integerExponent,
				big.NewInt(0),
				internalMaxPrecision,
				ePrefix)

		if err != nil {

			return result, resultPrecision,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "integerResult, integerPrecision, err = new(bigIntMathPowerNanobot).\n" +
						"  bigIntToPositiveIntegerPower(base, basePrecision, integerExponent,\n" +
						"    big.NewInt(0), internalMaxPrecision, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	ratFrac := big.NewRat(1, 1).SetFrac(fractionalExponent, scale)

	ratFracExponentNumerator := ratFrac.Num()

	baseToFracExponentNumerator := big.NewInt(0).Exp(base, ratFracExponentNumerator, nil)

	baseToFracExponentNumeratorPrecision := big.NewInt(0).Mul(basePrecision, ratFracExponentNumerator)

	delta := big.NewInt(0)

	bigFive := big.NewInt(5)

	ratFracExponentDenominator := ratFrac.Denom()

	fdr := FixedDecimalNthRoot{}

	fracExponentRoot, fracExponentRootPrecision, err :=
		fdr.CalculatePositiveIntegerNthRoot(
			baseToFracExponentNumerator,
			baseToFracExponentNumeratorPrecision,
			ratFracExponentDenominator,
			big.NewInt(0),
			internalMaxPrecision)

	if err != nil {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracExponentRoot, fracExponentRootPrecision, err :=\n" +
					"  fdr.CalculatePositiveIntegerNthRoot(\n" +
					"  baseToFracExponentNumerator, baseToFracExponentNumeratorPrecision,\n" +
					"  ratFracExponentDenominator, big.NewInt(0), internalMaxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	result, resultPrecision, errx =
		new(BigIntMathMultiply).BigIntMultiply(
			integerResult,
			integerPrecision,
			fracExponentRoot,
			fracExponentRootPrecision)

	if errx != nil {
		result = big.NewInt(0)
		resultPrecision = big.NewInt(0)
		err = fmt.Errorf(ePrefix+
			"%v", errx.Error())

		return result, resultPrecision, err
	}

	bigFive = big.NewInt(5)

	if resultPrecision.Cmp(maxPrecision) == 1 {
		delta = big.NewInt(0).Sub(resultPrecision, maxPrecision)
		delta.Sub(delta, bigOne)
		scale.Exp(big.NewInt(10), delta, nil)
		result.Quo(result, scale)
		if result.Cmp(bigZero) == -1 {
			bigFive.Neg(bigFive)
		}
		result.Add(result, bigFive)
		result.Quo(result, bigTen)
		resultPrecision = big.NewInt(0).Set(maxPrecision)
	}

	return result, resultPrecision, err
}
