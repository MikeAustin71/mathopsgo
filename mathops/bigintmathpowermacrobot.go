package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathPowerMacrobot struct {
	lock sync.Mutex
}

// bigIntToNegativeFractionalPower
//
//	Raises 'base' to the power of a negative fractional exponent,
//	'exponent'.
//
//	As stated in the function name, this method expects to process
//	only negative exponent values which have fractional digits to the
//	right of the decimal place.
//
//	Examples
//	========
//
//	base  basePrecision  exponent  exponentPrecision  result  resultPrecision
//
//	 5         0           -22             1            29    3 (to 3-decimal places)
//	                 5^-2.2 = 0.02899118654710782125882456003526
//	           The actual number of decimal places returned in the result
//	           is controlled by input parameter, 'maxPrecision'.
//
//	18         1           -34             1           136    3 (to 3-decimal places)
//	                1.8^-3.4 = 0.13554187298692911221722484380209
//	           The actual number of decimal places returned in the result
//	           is controlled by input parameter, 'maxPrecision'.
//
//	Input Parameters
//	================
//
//	base                     *big.Int
//	  The base which will be raised to the power of a positive integer
//	  exponent.
//
//	basePrecision            uint
//	  The precision specification for 'base'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and 'basePrecision'
//	  define a numeric value with a fixed number of decimal digits to
//	  the right of the decimal place.
//
//	exponent                 *big.Int
//	  The exponent to which 'base' will be raised by this calculation.
//	  By method definition, 'exponent' must be a negative value. If
//	  exponent is greater than 0, an error will be triggered.
//
//	exponentPrecision        uint
//	  The precision specification for 'exponent'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'exponent' and
//	  'exponentPrecision' define a numeric value with a fixed number
//	  of decimal digits to the right of the decimal place. For this
//	  method, 'exponentPrecision' must be greater than zero, thereby
//	  designating 'exponent' as a fractional value. A value of zero
//	  for 'exponentPrecision' will trigger an error.
//
//	maxPrecision             uint
//	  When this method calculates 'base' raised to the power of
//	  'exponent', the maximum number of decimal digits to the right of
//	  the decimal place in the result will be limited by
//	  'maxPrecision'.
//
//	Return Values
//	=============
//
//	result                   *big.Int
//	  If the calculation completes successfully, this return value
//	  will be populated with the value of 'base' raised to the power
//	  of 'exponent'.
//
//	resultPrecision          uint
//	  The precision specification for 'result'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'result' and
//	  'resultPrecision' define a numeric value with a fixed number of
//	  decimal digits to the right of the decimal place.
//
//	err                      error
//	  If the calculation encounters an error, an appropriate error
//	  message will be formatted and returned. If the calculation
//	  completes successfully, this return value will be set to 'nil'.
func (bIMathPwrMacrobot *bigIntMathPowerMacrobot) bigIntToNegativeFractionalPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (
	result *big.Int, resultPrecision *big.Int, err error) {

	bIMathPwrMacrobot.lock.Lock()

	defer bIMathPwrMacrobot.lock.Unlock()

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerMacrobot.bigIntToNegativeFractionalPower",
		"")

	if err != nil {
		return result, resultPrecision, err
	}
	bigZero := big.NewInt(0)
	bigOne := big.NewInt(1)

	if base.Cmp(bigZero) == 0 {
		// base is zero result is zero
		return result, resultPrecision, err
	}

	if exponentPrecision.Cmp(bigZero) == 0 {
		err = fmt.Errorf(ePrefix+
			"Error: 'exponentPrecision' is zero. This is an integer exponent. "+
			"exponentPrecision='%v' ", exponentPrecision.Text(10))

		return result, resultPrecision, err
	}

	// exponentPrecision > 0
	cmpExponentZero := exponent.Cmp(bigZero)

	if cmpExponentZero == 1 {
		err = fmt.Errorf(ePrefix+
			"Error: 'exponent' is a positive value! "+
			"exponent='%v' ", exponent.Text(10))

		return result, resultPrecision, err

	}

	if cmpExponentZero == 0 {
		// Any number raised to a zero power is one
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	if maxPrecision.Cmp(bigZero) == -1 {
		err = fmt.Errorf(ePrefix+
			"Error: 'maxPrecision' is a negative value! "+
			"maxPrecision='%v' ", maxPrecision.Text(10))

		return result, resultPrecision, err
	}

	positiveExponent := big.NewInt(0).Set(exponent)
	positiveExponent.Neg(positiveExponent)

	internalMaxPrecision := big.NewInt(0).Add(maxPrecision, big.NewInt(2))

	positiveResult, positivePrecision, errx :=
		BigIntMathPower{}.BigIntToPositiveFractionalPower(
			base,
			basePrecision,
			positiveExponent,
			exponentPrecision,
			internalMaxPrecision)

	if errx != nil {
		err = fmt.Errorf(ePrefix+"%v", errx.Error())
		return result, resultPrecision, err
	}

	result, resultPrecision, errx = BigIntMathDivide{}.BigIntFracQuotient(
		bigOne,
		big.NewInt(0),
		positiveResult,
		positivePrecision,
		maxPrecision)

	if errx != nil {
		result = big.NewInt(0)
		resultPrecision = big.NewInt(0)
		err = fmt.Errorf(ePrefix+"%v", errx.Error())
		return result, resultPrecision, err
	}

	return result, resultPrecision, err
}
