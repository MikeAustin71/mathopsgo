package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathPowerNanobot struct {
	lock sync.Mutex
}

// bigIntToPositiveIntegerPower
//
//	Raises 'base' to the power of a positive integer exponent,
//	'exponent'.
//
//	As stated in the function name, this method expects to process
//	only positive, integer exponents.
//
//	Examples
//	========
//
//	base  basePrecision  exponent  exponentPrecision  result  resultPrecision
//
//	  5       0             2              0            25          0
//	                            (5^2 = 25)
//
//	 18       1             3              0           5832         3
//	                           (1.8^3= 5.832)
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
//	  By method definition, 'exponent' must be a positive value. If
//	  exponent is less than 0, an error will be triggered.
//
//	exponentPrecision        uint
//	  The precision specification for 'exponent'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and 'basePrecision'
//	  define a numeric value with a fixed number of decimal digits to
//	  the right of the decimal place. For this method,
//	  'exponentPrecision' must be set to zero, designating exponent as
//	  an integer value. Values greater than zero will trigger an
//	  error.
//
//	maxPrecision             uint
//	  When this method calculates 'base' raised to the power of
//	  'exponent', the maximum number of decimal digits to the right of
//	  the decimal place in the resulting value will be limited by
//	  'maxPrecision'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//	  This object encapsulates an error prefix string
//	  which is included in all returned error
//	  messages. Usually, it contains the name of the
//	  calling method or methods listed as a function
//	  chain.
//
//	  If no error prefix information is needed, set
//	  this parameter to 'nil'.
//
//	  Type ErrPrefixDto is included in the 'errpref'
//	  software package:
//	    "github.com/MikeAustin71/errpref".
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
//	  message will be formatted and returned.
func (bIMathPwrNanobot *bigIntMathPowerNanobot) bigIntToPositiveIntegerPower(
	base *big.Int,
	basePrecision *big.Int,
	exponent *big.Int,
	exponentPrecision *big.Int,
	maxPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (result *big.Int, resultPrecision *big.Int, err error) {

	bIMathPwrNanobot.lock.Lock()

	defer bIMathPwrNanobot.lock.Unlock()

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerNanobot.bigIntNumRaiseToPositiveIntegerPower",
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

	if base.Cmp(bigZero) == 0 {
		// base is zero result is zero
		return result, resultPrecision, err
	}

	cmpExponentZero := exponent.Cmp(bigZero)

	if cmpExponentZero == 0 {
		// Any number raised to a zero power is one
		result = big.NewInt(1)
		return result, resultPrecision, err
	}

	if cmpExponentZero == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'", exponent.Text(10)),
				ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
					"Input parameter 'exponent' is a negative value!\n" +
					"Only positive integer exponents can be used with this method.",
			}
	}

	// exponentPrecision > 0
	if exponentPrecision.Cmp(bigZero) == 1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("exponent= '%v'\nexponentPrecision= '%v'",
					exponent.Text(10), exponentPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'exponent' is NOT an integer!\n" +
					"Only integer exponents can be used with this method.",
			}
	}

	bigOne := big.NewInt(1)

	if exponent.Cmp(bigOne) == 0 {

		result.Set(base)

		resultPrecision = basePrecision

		return result, resultPrecision, nil
	}

	result.Exp(base, exponent, nil)

	resultPrecision = big.NewInt(0).Mul(
		basePrecision,
		exponent)

	if resultPrecision.Cmp(maxPrecision) == 1 {

		bigTen := big.NewInt(10)

		delta := big.NewInt(0).Sub(resultPrecision, maxPrecision)

		delta.Sub(delta, bigOne)

		scale := big.NewInt(0).Exp(bigTen, delta, nil)

		result.Quo(result, scale)

		roundFive := big.NewInt(5)

		if result.Cmp(bigZero) == -1 {

			roundFive.Neg(roundFive)

		}

		result.Add(result, roundFive)

		result.Quo(result, bigTen)

		resultPrecision.Set(maxPrecision)

	}

	return result, resultPrecision, nil
}
