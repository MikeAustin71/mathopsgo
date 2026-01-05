package mathops

import (
	"math/big"

	ePref "github.com/MikeAustin71/errpref"
)

type BigIntMathPower struct {
	Base     BigIntNum
	Exponent BigIntNum
	Result   BigIntNum
}

// BigIntPwr
//
//	Raises input parameter 'base' to the power of input parameter
//	'exponent'. Both 'base' and 'exponent' may be positive or negative
//	integer or fractional values.
//
//	This method uses the exponent method ('Exp') provided by the go
//	"math/big" package.
func (bIPwr *BigIntMathPower) BigIntPwr(
	base *big.Int,
	basePrecision *big.Int,
	exponent *big.Int,
	exponentPrecision *big.Int,
	maxPrecision *big.Int) (result *big.Int, resultPrecision *big.Int, err error) {

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntPwr",
		"")

	if err != nil {
		return result, resultPrecision, err
	}

	return new(bigIntMathPowerMechanics).bigIntPwr(
		base, basePrecision, exponent, exponentPrecision, maxPrecision, ePrefix)
}

// BigIntToNegativeFractionalPower
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
func (bIPwr *BigIntMathPower) BigIntToNegativeFractionalPower(
	base *big.Int,
	basePrecision *big.Int,
	exponent *big.Int,
	exponentPrecision *big.Int,
	maxPrecision *big.Int) (result *big.Int, resultPrecision *big.Int, err error) {

	result = big.NewInt(0)

	resultPrecision = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntToNegativeFractionalPower",
		"")

	if err != nil {
		return result, resultPrecision, err
	}

	return new(bigIntMathPowerMacrobot).bigIntToNegativeFractionalPower(
		base, basePrecision, exponent, exponentPrecision, maxPrecision, ePrefix)
}

// BigIntToPositiveFractionalPower
//
//	Raises 'base' to the power of a positive fractional exponent,
//	'exponent'.
//
//	As stated in the function name, this method expects to process
//	only positive exponent values which have fractional digits to the
//	right of the decimal place.
//
//	Examples
//	========
//
//	base  basePrecision  exponent  exponentPrecision  result  resultPrecision
//
//	 5         0            22             1          34493   3 (to 3-decimal places)
//	         (5^2.2 = 34.4932415365303708097515866054
//	         The actual number of decimal places returned in the result
//	         is controlled by input parameter, 'maxPrecision'.)
//
//	18         1            34             1          7378    3 (to 3-decimal places)
//	         (1.8^3.4 = 7.3777938725727533349996174917827
//	         The actual number of decimal places returned in the result
//	         is controlled by input parameter, 'maxPrecision'.)
//
//	Input Parameters
//	================
//
//	base                     *big.Int
//	  The base which will be raised to the power of a positive integer
//	  exponent.
//
//	basePrecision            uint
//	  The precision specification for 'base'. Precision defines the
//	  number of numeric digits to the right of the decimal place.
//	  Taken together, 'base' and 'basePrecision' define a numeric
//	  value with a fixed number of decimal digits to the right of the
//	  decimal place.
//
//	exponent                 *big.Int
//	  The exponent to which 'base' will be raised by this calculation.
//	  By method definition, 'exponent' must be a positive value. If
//	  exponent is less than 0, an error will be triggered.
//
//	exponentPrecision        uint
//	  The precision specification for 'exponent'. Precision defines
//	  the number of numeric digits to the right of the decimal place.
//	  Taken together, 'base' and 'basePrecision' define a numeric
//	  value with a fixed number of decimal digits to the right of the
//	  decimal place. For this method, 'exponentPrecision' must be
//	  greater than zero, thereby designating 'exponent' as a
//	  fractional value. A value of zero for 'exponentPrecision' will
//	  trigger an error.
//
//	maxPrecision             uint
//	  When this method calculates 'base' raised to the power of
//	  'exponent', the maximum number of decimal digits to the right of
//	  the decimal place in the result will be limited by 'maxPrecision'.
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
func (bIPwr *BigIntMathPower) BigIntToPositiveFractionalPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int) (result *big.Int, resultPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntToPositiveFractionalPower",
		"")

	if err != nil {
		return result, resultPrecision, err
	}

	return new(bigIntMathPowerMinibot).bigIntToPositiveFractionalPower(
		base, basePrecision, exponent, exponentPrecision, maxPrecision, ePrefix)
}

// BigIntToNegativeIntegerPower
//
//	Raises 'base' to the power of a negative integer exponent,
//	'exponent'.
//
//	As stated in the function name, this method expects to process
//	only negative, integer exponents.
//
//	Examples
//	========
//
//	base  basePrecision  exponent  exponentPrecision  result     resultPrecision
//
//	 5         0            -2            0             4              2
//	                          5^-2 = 0.04
//
//	1131       2            -3            0           00069121    8 (to 8-decimal places)
//	            11.31^-3= 0.00069121345785745610965099525880031
//
//		            The actual number of decimal places returned in
//	            the result is controlled by	input parameter,
//	            'maxPrecision'.
//
//	Input Parameters
//	================
//
//	base                     *big.Int
//	  The base which will be raised to the power of a negative integer
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
//	  exponent is greater than -1, an error will be triggered.
//
//	exponentPrecision        uint
//	  The precision specification for 'exponent'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and 'basePrecision'
//	  define a numeric value with a fixed number of decimal digits to
//	  the right of the decimal place. For this method,
//	  'exponentPrecision' MUST BE SET TO ZERO, thereby designating
//	  'exponent' as an integer value. Values greater than zero will
//	  trigger an error.
//
//	maxPrecision             uint
//	  When this method calculates 'base' raised to the power of
//	  'exponent', the maximum number of decimal digits to the right of
//	  the decimal place in the resulting value will be limited by
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
func (bIPwr *BigIntMathPower) BigIntToNegativeIntegerPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int) (result *big.Int, resultPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntToNegativeIntegerPower",
		"")

	if err != nil {
		return result, resultPrecision, err
	}

	return new(bigIntMathPowerNeutron).bigIntToNegativeIntegerPower(
		base, basePrecision, exponent, exponentPrecision, maxPrecision, ePrefix)
}

// BigIntToPositiveIntegerPower
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
func (bIPwr *BigIntMathPower) BigIntToPositiveIntegerPower(
	base,
	basePrecision,
	exponent,
	exponentPrecision,
	maxPrecision *big.Int) (result *big.Int, resultPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntToPositiveIntegerPower",
		"")

	if err != nil {
		return result, resultPrecision, err
	}

	return new(bigIntMathPowerNanobot).
		bigIntToPositiveIntegerPower(
			base, basePrecision, exponent, exponentPrecision, maxPrecision, ePrefix)
}

// BigIntPwrIteration
//
//	Raises input parameter 'base' to the power of input parameter
//	'exponent'.
//
//	This method of raising a base to an exponent uses iterative
//	multiplication and manages the internal precision of each
//	iterative multiplication. If, during the process of multiplying
//	the base time itself, the internal precision exceeds the
//	'internalMaxPrecision' limit, that intermediate number is rounded
//	down to 'internalMaxPrecision'.
//
//	If the precision of the final result exceeds the limit imposed by
//	input parameter, 'outputMaxPrecision', that final result will be
//	rounded to 'outputMaxPrecision' digits to the right of the decimal
//	place.
//
//	Input Parameter
//	===============
//
//	base                     *big.Int
//	  The base which will be raised to the power of 'exponent'.
//	             baseToPwr = base^exponent
//
//	basePrecision            uint
//	  The number of digits to the right of the decimal place in the
//	  numeric sequence represented by 'base'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and
//	  'basePrecision' define a numeric value with a fixed number of
//	  decimal digits to the right of the decimal place.
//
//	exponent                 uint
//	  This function will raise 'base' to the power of 'exponent'.
//	              baseToPwr = base^exponent
//
//	internalMaxPrecision     uint
//	  This value is imposed as a limit on the precision of internal
//	  calculations necessary to compute the result of this power
//	  operation. If during the calculation an interim or intermediate
//	  result is generated which exceeds this limit, that intermediate
//	  result will be rounded to 'internalMaxPrecision'. The term
//	  precision defines the number of digits to the right of the
//	  decimal place.
//
//	  If 'internalMaxPrecision' is less than 'outputMaxPrecision',
//	  'internalMaxPrecision' will be automatically set to a value of
//	  'outputMaxPrecision' + 100.
//
//	outputMaxPrecision       uint
//	  This value is imposed as a limit on the precision of the final
//	  calculated result of the power operation. If the number of
//	  digits to the right of the decimal point in the final calculated
//	  result exceeds this limit, that final result will be rounded to
//	  'outputMaxPrecision' digits to the right of the decimal place.
//	  The term precision defines the number of digits to the right of
//	  the decimal place.
//
//	Return Values
//	=============
//
//	baseToPwr                *big.Int
//	  This function returns the result of 'base' raised to the power
//	  of 'exponent'. This result, 'baseToPwr' is returned as a type
//	  *big.Int.
//	                   baseToPwr = base^exponent
//
//	baseToPwrPrecision       uint
//	  Specifies the number of digits to the right of the decimal place
//	  in the numeric sequence represented by the calculation result,
//	  'baseToPwr'.
//
//	err                      error
//	  If the calculation encounters an error, an appropriate error
//	  message will be formatted and returned. If the calculation
//	  completes successfully, this return value will be set to 'nil'.
func (bIPwr *BigIntMathPower) BigIntPwrIteration(
	base *big.Int,
	basePrecision uint,
	exponent uint,
	internalMaxPrecision uint,
	outputMaxPrecision uint) (baseToPwr *big.Int, baseToPwrPrecision uint, err error) {

	baseToPwr = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntPwrIteration",
		"")

	if err != nil {
		return baseToPwr, baseToPwrPrecision, err
	}

	return new(bigIntMathPowerNeutron).bigIntPwrIteration(
		base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision, ePrefix)
}

// BigIntegerPwrIteration
//
//	Raises input parameter 'base' to the power of input parameter
//	'exponent'. This version of the function uses *big.Int type for
//	all input parameters.
//
//	This method of raising a base to an exponent uses iterative
//	multiplication and manages the internal precision of each
//	iterative multiplication. If, during the process of multiplying
//	the base times itself, the internal precision exceeds the
//	'internalMaxPrecision' limit, that intermediate number is rounded
//	to 'internalMaxPrecision'.
//
//	If the precision of the final result exceeds the limit imposed by
//	input parameter, 'outputMaxPrecision', that final result will be
//	rounded to 'outputMaxPrecision' digits to the right of the
//	decimal place.
//
//	!!! WARNING !!!
//	===============
//
//	Currently, this method will only accept positive numeric values
//	for input parameter 'exponent'.
//
//	Input Parameter
//	===============
//
//	base                     *big.Int
//	  The base which will be raised to the power of 'exponent'.
//	                  baseToPwr = base^exponent
//
//	basePrecision            *big.Int
//	  The number of digits to the right of the decimal place in the
//	  numeric sequence represented by 'base'.
//
//	  Precision defines the number of numeric digits to the right of
//	  the decimal place. Taken together, 'base' and
//	  'basePrecision' define a numeric value with a fixed number of
//	  decimal digits to the right of the decimal place.
//
//	exponent                 *big.Int
//	  'exponent' is an integer value with zero precision.
//
//	   This function will raise 'base' to the power of 'exponent'.
//									baseToPwr = base^exponent
//
//	   !!! WARNING !!!
//	     Currently, 'exponent' MUST BE a positive numeric value.
//
//	internalMaxPrecision     *big.Int
//	  This value is imposed as a limit on the precision of internal
//	  calculations necessary to compute the result of this power
//	  operation. If during the calculation an interim or intermediate
//	  result is generated which exceeds this limit, that intermediate
//	  result will be rounded to 'internalMaxPrecision'.
//
//	  The term 'precision' defines the number of digits to the right
//	  of the decimal place.
//
//	  If 'internalMaxPrecision' is less than 'outputMaxPrecision',
//	  'internalMaxPrecision' will be automatically set to a value of
//	  'outputMaxPrecision' + 100.
//
//	outputMaxPrecision       *big.Int
//	  This value is imposed as a limit on the precision of the final
//	  calculated result of the power operation. If the number of
//	  digits to the right of the decimal point in the final calculated
//	  result exceeds this limit, that final result will be rounded to
//	  'outputMaxPrecision' digits to the right of the decimal place.
//
//	  The term 'precision' defines the number of digits to the right
//	  of the decimal place. If 'outputMaxPrecision' is less than zero,
//	  an error will be returned.
//
//	Return Values
//	=============
//
//	baseToPwr                *big.Int
//	  This function returns the result of 'base' raised to the power
//	  of 'exponent'. This result, 'baseToPwr' is returned as a type
//	  *big.Int.
//	                    baseToPwr = base^exponent
//
//	baseToPwrPrecision       *big.Int
//	  Specifies the number of digits to the right of the decimal place
//	  in the numeric sequence represented by the calculation result,
//	  'baseToPwr'.
//
//	err                      error
//	  If the function fails to complete successfully, this value is
//	  configured with an appropriate error message and returned to the
//	  caller. If the function completes successfully, this value is
//	  set to 'nil'.
func (bIPwr *BigIntMathPower) BigIntegerPwrIteration(
	base,
	basePrecision,
	exponent,
	internalMaxPrecision,
	outputMaxPrecision *big.Int) (baseToPwr *big.Int, baseToPwrPrecision *big.Int, err error) {

	baseToPwr = big.NewInt(0)
	baseToPwrPrecision = big.NewInt(0)

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntegerPwrIteration",
		"")

	if err != nil {
		return baseToPwr, baseToPwrPrecision, err
	}

	return new(bigIntMathPowerNeutron).bigIntegerPwrIteration(
		base, basePrecision, exponent, internalMaxPrecision, outputMaxPrecision, ePrefix)
}

// FixedDecimalPwrIteration
//
//	Raises input parameter 'base' to the power of input parameter
//	'exponent'.
//
//	This process of raising a base to an exponent uses iterative
//	multiplication and manages the internal precision of each
//	iterative multiplication. If, during the process of multiplying
//	the base times itself, the internal precision exceeds the
//	'internalMaxPrecision' limit, that intermediate number is rounded
//	to 'internalMaxPrecision' digits to the right of the decimal place.
//
//	If the precision of the final result exceeds the limit imposed by
//	input parameter, 'outputMaxPrecision', that final result will be
//	rounded to 'outputMaxPrecision' digits to the right of the decimal
//	place.
//
//	Input Parameter
//	===============
//
//	base                     BigIntFixedDecimal
//	  The base which will be raised to the power of 'exponent'. The
//	  BigIntFixedDecimal type describes a numeric value with a fixed
//	  number of digits to the right of the decimal place. The type
//	  includes a *big.Int integer value and a precision specification.
//								- baseToPwr = base^exponent -
//
//	exponent                 uint
//	  This function will raise 'base' to the power of 'exponent'.
//	              - baseToPwr = base^exponent -
//
//	internalMaxPrecision     uint
//	  This value is imposed as a limit on the precision of internal
//	  calculations necessary to compute the result of this power
//	  operation. If during the calculation an interim or intermediate
//	  result is generated which exceeds this limit, that intermediate
//	  result will be rounded to 'internalMaxPrecision' digits to the
//	  right of the decimal place.
//
//	  The term 'precision' defines the number of digits to the right
//	  of the decimal place.
//
//	  If 'internalMaxPrecision' is less than 'outputMaxPrecision',
//	  'internalMaxPrecision' will be automatically set to a value of
//	  'outputMaxPrecision' + 100.
//
//	outputMaxPrecision       uint
//	  This value is imposed as a limit on the precision of the final
//	  calculated result of the power operation.
//
//	  If the number of digits to the right of the decimal point in the
//	  final calculated result exceeds this limit, that final result
//	  will be rounded to 'outputMaxPrecision' digits to the right of
//	  the decimal place. The term precision defines the number of
//	  digits to the right of the decimal place.
//
//	Return Values
//	=============
//
//	baseToPwr                BigIntFixedDecimal
//	  This function returns the result of 'base' raised to the power
//	  of 'exponent'. This result, 'baseToPwr' is returned as a type
//	  BigIntFixedDecimal.
//
//	  The BigIntFixedDecimal type describes a numeric value with a
//	  fixed number of digits to the right of the decimal place.
//	             				   baseToPwr = base^exponent
//
//	err                      error
//	  If the function fails to complete successfully, this value is
//	  configured with an appropriate error message and returned to the
//	  caller. If the function completes successfully, this value is
//	  set to 'nil'.
func (bIPwr *BigIntMathPower) FixedDecimalPwrIteration(
	base BigIntFixedDecimal,
	exponent uint,
	internalMaxPrecision uint,
	outputMaxPrecision uint) (baseToPwr BigIntFixedDecimal, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.FixedDecimalPwrIteration",
		"")

	if err != nil {
		return baseToPwr, err
	}

	return new(bigIntMathPowerAtom).fixedDecimalPwrIteration(
		base, true, exponent, internalMaxPrecision, outputMaxPrecision, ePrefix)
}

// BigIntNumMinRequiredPrecision
//
//	Designed to be used with the power function (BigIntNumPwr). This method
//	will compute the minimum number of decimal places required to
//	support the result of raising a 'base' value to a specified
//	exponent. Both the 'base' and the 'exponent' are passed to this
//	function as type BigIntNum.
//
//	For example, raising the value 3.12 to the power of 4 means that
//	the result will require at least 8-decimal places to the right of
//	the decimal in order to display a correct result. In the following
//	example with base ='3.12' and exponent = '4', this method will
//	return '8'.
//
//	Example
//	=======
//
//	3.12^4 = 94.75854336 (2x4 = 8-digits to the right of the decimal)
//	       Minimum Required Precision = precision x exponent
//
//	The calculated minimum required precision is returned as type
//	'uint'.
//
//	If the minimum required precision exceeds the maximum value for
//	type 'uint' (+4,294,967,295, which equals 2^32 − 1), an error
//	message is returned in addition to the maximum uint value
//	(+4,294,967,295).
func (bIPwr *BigIntMathPower) BigIntNumMinRequiredPrecision(
	base, exponent BigIntNum) (uint, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntNumMinRequiredPrecision",
		"")

	if err != nil {
		return 0, err
	}

	return new(bigIntMathPowerNeutron).bigIntNumMinRequiredPrecision(
		&base, true, &exponent, true, ePrefix)
}

// BigIntNumPwr
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
func (bIPwr *BigIntMathPower) BigIntNumPwr(
	base BigIntNum, exponent BigIntNum, maxPrecision uint) (BigIntNum, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMathPower.BigIntNumPwr",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	return new(bigIntMathPowerMechanics).bigIntNumPwr(
		&base, true, &exponent, true, maxPrecision, ePrefix)
}

// computeInternalPrecision - Returns computed internal precision for variables used
// in intermediate calculations. Returned 'internalPrecision' is based on requested
// maximum precision for a specific BigIntMathPower calculation.
/*
func (bIPwr *BigIntMathPower) computeMaxInternalPrecision(maxPrecision *big.Int) *big.Int {

	internalPrecision := big.NewInt(10)

	maxPrecisionCmp25 := maxPrecision.Cmp(big.NewInt(50))

	maxPrecisionCmp50 := maxPrecision.Cmp(big.NewInt(50))

	maxPrecisionCmp200 := maxPrecision.Cmp(big.NewInt(200))

	maxPrecisionCmpTwoThou := maxPrecision.Cmp(big.NewInt(2000))

	maxPrecisionCmpFiveThou := maxPrecision.Cmp(big.NewInt(5000))

	maxPrecisionCmpTenThou := maxPrecision.Cmp(big.NewInt(10000))

	if maxPrecisionCmp25 < 1 {

		internalPrecision.Add(maxPrecision, big.NewInt(50))

	} else if  maxPrecisionCmp25 >= 0 &&   maxPrecisionCmp50 == -1 {

		internalPrecision.Add(maxPrecision, big.NewInt(65))

	} else if maxPrecisionCmp50 >= 0  && maxPrecisionCmp200 == -1 {
		internalPrecision.Add(maxPrecision, big.NewInt(75))

	} else if maxPrecisionCmp200 >= 0 && maxPrecisionCmpTwoThou ==-1 {
		internalPrecision.Add(maxPrecision, big.NewInt(0).Quo(maxPrecision, big.NewInt(4)))

	}else if maxPrecisionCmpTwoThou >= 0 && maxPrecisionCmpFiveThou == -1 {
		internalPrecision.Add(maxPrecision, big.NewInt(0).Quo(maxPrecision, big.NewInt(8)))

	}else if maxPrecisionCmpFiveThou >= 0 && maxPrecisionCmpTenThou == -1 {
		internalPrecision.Add(maxPrecision, big.NewInt(0).Quo(maxPrecision, big.NewInt(100)))

	} else if maxPrecisionCmpTenThou >= 0  {
		internalPrecision.Add(maxPrecision, big.NewInt(100))
	}

	return internalPrecision

}
*/
