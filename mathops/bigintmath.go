package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
)

type BigIntMath struct {
	Input  *big.Int
	Output *big.Int
}

// ArithmeticGeometricMean
//
// Computes the Arithmetic-Geometric Mean of two numbers. See:
//
// https://en.wikipedia.org/wiki/Arithmetic–geometric_mean
//
// Dev Note: maxInternalPrecision may need to be 60 to 75 times
// targetPrecision.
func (bIntMath *BigIntMath) ArithmeticGeometricMean(
	aNum *big.Int,
	aNumPrecision *big.Int,
	gNum *big.Int,
	gNumPrecision *big.Int,
	maxInternalPrecision *big.Int,
	targetPrecision *big.Int) (agMean *big.Int,
	agMeanPrecision *big.Int,
	gValue *big.Int,
	gValuePrecision *big.Int,
	cycles uint64,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMath.ArithmeticGeometricMean",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), 0, err
	}

	return new(bigIntMathNanobot).arithmeticGeometricMeanBigInt(
		aNum, aNumPrecision, gNum, gNumPrecision, maxInternalPrecision, targetPrecision, ePrefix)
}

// BigIntPrecisionCmp
// Compares two *big.Int number pairs and adjusts the comparison for
// precision.
//
// Return Values
// num 1 <  num2 == -1
// num 1 == num2 == 0
// num 1 >  num2 == 1
func (bIntMath *BigIntMath) BigIntPrecisionCmp(
	num1 *big.Int,
	num1Precision *big.Int,
	num2 *big.Int,
	num2Precision *big.Int) (int, error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMath.BigIntPrecisionCmp",
		"")

	if err != nil {
		return 0, err
	}

	return new(bigIntMathElectron).precisionCmpBigInt(
		num1, num1Precision, num2, num2Precision, ePrefix)
}

// GetMagnitude
//
// Returns the magnitude of a *big.Int number passed as an input
// parameter ('initialValue'. Magnitude is defined here as the
// power of 10 which generates a value less than or equal to the
// 'target' *big.Int number.
//
//	10^magnitude  <= initialValue
//
// The value of magnitude is returned as a *big.Int number to the
// calling function.
//
//	Examples
//	========
//
//	** target  **      magnitude
//	-------------      ---------
//
//	      963,256           5
//	            2           0
//	           32           1
//	8,456,123,921           9
//
//	Input Parameters
//	================
//
//	initialValue        *big.Int
//
//	An integer of type *big.Int. This method will analyze this
//	integer and return its magnitude.
//
//
//	Return Values
//	=============
//
//	magnitude           *big.Int
//
//	10 raised to the power of magnitude will yield a value which
//	is less than or equal to the input parameter 'initialValue'.
//
//
//	err					        error
//
//	If no errors are encountered during processing, this returned
//	error value will be set to 'nil'.
func (bIntMath *BigIntMath) GetMagnitude(
	initialValue *big.Int) (magnitude *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMath.GetMagnitude",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	return new(bigIntMathElectron).getMagnitudeBigInt(
		initialValue, ePrefix)
}

// RoundToMaxPrecision
//
// Applies maximum precision to a *big.Int number and associated
// numeric precision, 'bigIntNum' and 'bigIntNumPrecision'.
//
// Precision as used here defines the number of digits to the
// right of the decimal place.
//
// If 'bigIntNumPrecision' exceeds 'maxPrecision', the
// 'bigIntNum' and 'bigIntNumPrecision' pair are rounded to
// 'maxPrecision' and returned as 'result' and 'resultPrecision'.
//
//	Input Parameters
//	=================
//
//	bigIntNum                *big.Int
//
//	The integer number of the numeric value to be rounded.
//
//
//	bigIntNumPrecision       *big.Int
//
//	The precision specification associated with 'bigIntNum'. The
//	precision specification defines the number of digits to the
//	right of the decimal place in 'bigIntNum'.
//
//
//	maxPrecision             *big.Int
//
//	If 'bigIntNumPrecision' exceeds 'maxPrecision', the 'bigIntNum'
//	and 'bigIntNumPrecision' pair are rounded to 'maxPrecision'
//	and returned as 'result' and 'resultPrecision'.
//
//	trimTrailingFracZeros    bool
//
//	If trailing fractional zeros are present in the rounded result,
//	this boolean value will determine whether the trailing zeros
//	will be returned in final result. 'true' specifies that all
//	trailing fractional zeros will be deleted.
//
//	    Example: '1.23000'  converted to '1.23'
//
//	Return Values
//	=============
//
//	result                    *big.Int
//
//	The result of the rounding operation expressed a *big.Int
//	value.
//
//	resultPrecision            *big.Int
//
//	The number of digits within 'result' which are fractional
//	digits to be placed at the right of the decimal point.
//
//	err                        error
//
//	If this method completes successfully, this returne error
//	value will be set to 'nil'
//
//	Examples
//	========
//
//	bigIntNum bigIntNumPrecision maxPrecision result  resultPrecision
//
//	  5255            3                2        526         2
//	 52671            4                6       52671        4
func (bIntMath *BigIntMath) RoundToMaxPrecision(
	bigIntNum *big.Int,
	bigIntNumPrecision *big.Int,
	maxPrecision *big.Int,
	trimTrailingFracZeros bool) (result *big.Int, resultPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMath.RoundToMaxPrecision",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	return new(bigIntMathElectron).roundToMaxPrecisionBigInt(
		bigIntNum,
		bigIntNumPrecision,
		maxPrecision,
		trimTrailingFracZeros,
		ePrefix)
}

// TruncateToMaxPrecision
//
// Applies maximum precision to a *big.Int number and associated
// numeric precision, 'bigIntNum' and 'bigIntNumPrecision'.
//
// Precision as used here defines the number of digits to the right
// of the decimal place. If 'bigIntNumPrecision' exceeds
// 'maxPrecision', the 'bigIntNum' and 'bigIntNumPrecision' pair
// are rounded to 'maxPrecision' and returned as 'result' and
// 'resultPrecision'.
//
// Note that if 'bigIntNumPrecision' exceeds 'maxPrecision', the
// returned value will be truncated to (not rounded to)
// 'maxPrecision' decimal digits to the right of the decimal
// place.
//
//	Examples
//	=========
//
//	                                             Result      Result
//	bigIntNum  bigIntNumPrecision  maxPrecision  Integer    Precision
//
//	 5255             3                 2            525        2
//	52671             4                 6          52671        4
func (bIntMath *BigIntMath) TruncateToMaxPrecision(
	bigIntNum *big.Int,
	bigIntNumPrecision *big.Int,
	maxPrecision *big.Int) (result *big.Int,
	resultPrecision *big.Int, err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMath.TruncateToMaxPrecision",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	return new(bigIntMathElectron).truncateToMaxPrecisionBigInt(
		bigIntNum, bigIntNumPrecision, maxPrecision, ePrefix)
}

// TruncateTrailingFractionalZeros
//
// Truncates trailing fractional zeros. If a numeric value has
// trailing fractional zeros, this method will delete those zeros
// and adjust the precision indicator accordingly.
//
//	Example
//	=======
//
//	            Numeric                 Result     Result    Result
//	Number     Precision    Value       Integer   Precision  Value
//
//	12345600       5       123.45600     123456       3      123.456
func (bIntMath *BigIntMath) TruncateTrailingFractionalZeros(
	num,
	numPrecision *big.Int) (result *big.Int,
	resultPrecision *big.Int,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"BigIntMath.TruncateTrailingFractionalZeros",
		"")

	return new(bigIntMathElectron).truncateTrailingFracZerosBigInt(
		num, numPrecision, ePrefix)
}
