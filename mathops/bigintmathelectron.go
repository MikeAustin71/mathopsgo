package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntMathElectron struct {
	lock *sync.Mutex
}

// getMagnitudeBigInt
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
func (bIMathElectron *bigIntMathElectron) getMagnitudeBigInt(
	initialValue *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (magnitude *big.Int, err error) {

	if bIMathElectron.lock == nil {
		bIMathElectron.lock = new(sync.Mutex)
	}

	bIMathElectron.lock.Lock()

	defer bIMathElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathElectron.truncateToMaxPrecisionBigInt",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	magnitude = big.NewInt(0)
	err = nil

	if initialValue == nil {

		return magnitude,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'initialValue'",
			}
	}

	target := big.NewInt(0).Set(initialValue)
	compareResult := target.Cmp(big.NewInt(0))

	// Convert to absolute value
	if compareResult == -1 {
		target.Neg(target)
	}

	if compareResult == 0 {
		// magnitude = 0; err=nil
		return magnitude, err
	}

	bigTen := big.NewInt(10)

	if target.Cmp(bigTen) == -1 {
		// magnitude = 0; err=nil
		return magnitude, err
	}

	bitLen := target.BitLen()

	if bitLen <= 0 {

		return magnitude,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("bitLen='%v'", bitLen),
				ErrMessage: "Error: target.BitLen() = %v - Negative value!",
			}
	}

	tenToPowerPrecision := big.NewInt(0)
	bigZero := big.NewInt(0)

	// Note: log10of2To20k is a global constant

	// ************************************
	// target MUST BE <= 2^(bit length)
	// ************************************
	magnitude, tenToPowerPrecision, err =
		new(BigIntMathMultiply).BigIntMultiply(
			big.NewInt(int64(bitLen)),
			big.NewInt(0),
			log10of2To20k.GetInteger(),
			log10of2To20k.GetPrecisionBigInt())

	if err != nil {

		return magnitude,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "magnitude, tenToPowerPrecision, err = \n" +
					"new(BigIntMathMultiply).BigIntMultiply(\n" +
					"    big.NewInt(int64(bitLen)),  big.NewInt(0),\n" +
					"    log10of2To20k.GetInteger(), log10of2To20k.GetPrecisionBigInt())",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if tenToPowerPrecision.Cmp(bigZero) == 1 {
		scale :=
			big.NewInt(0).Exp(
				bigTen,
				tenToPowerPrecision,
				nil)

		magnitude.Quo(magnitude, scale)
	}

	testNum := big.NewInt(0).Exp(bigTen, magnitude, nil)

	if testNum.Cmp(target) == 1 {
		magnitude.Sub(magnitude, big.NewInt(1))
	}

	return magnitude, nil
}

// precisionCmpBigInt
//
// Compares two *big.Int number pairs and adjusts the comparison
// for precision.
//
// Return Values
// num 1 <  num2 == -1
// num 1 == num2 == 0
// num 1 >  num2 == 1
func (bIMathElectron *bigIntMathElectron) precisionCmpBigInt(
	num1 *big.Int,
	num1Precision *big.Int,
	num2 *big.Int,
	num2Precision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIMathElectron.lock == nil {
		bIMathElectron.lock = new(sync.Mutex)
	}

	bIMathElectron.lock.Lock()

	defer bIMathElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathElectron.precisionCmpBigInt",
		"")

	if err != nil {
		return 0, err
	}

	if num1 == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'num1'",
			}
	}

	if num1Precision == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'num1Precision'",
			}
	}

	if num2 == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'num2'",
			}
	}

	if num2Precision == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'num2Precision'",
			}
	}

	bigZero := big.NewInt(0)

	if num1Precision.Cmp(bigZero) == -1 {

		num1Precision = big.NewInt(0)
	}

	if num2Precision.Cmp(bigZero) == -1 {

		num2Precision = big.NewInt(0)
	}

	if num1Precision.Cmp(num2Precision) == 0 {

		return num1.Cmp(num2), nil
	}

	bigTen := big.NewInt(10)
	tNum1 := big.NewInt(0).Set(num1)
	tNum2 := big.NewInt(0).Set(num2)

	if num1Precision.Cmp(num2Precision) == 1 {

		delta := big.NewInt(0).Sub(num1Precision, num2Precision)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		tNum2.Mul(tNum2, scale)
		return tNum1.Cmp(tNum2), nil

	}

	// MUST BE num2Precision > num1Precision
	delta := big.NewInt(0).Sub(num2Precision, num1Precision)
	scale := big.NewInt(0).Exp(bigTen, delta, nil)
	tNum1.Mul(tNum1, scale)

	return tNum1.Cmp(tNum2), nil
}

// roundToMaxPrecisionBigInt
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
func (bIMathElectron *bigIntMathElectron) roundToMaxPrecisionBigInt(
	bigIntNum *big.Int,
	bigIntNumPrecision *big.Int,
	maxPrecision *big.Int,
	trimTrailingFracZeros bool,
	errPrefDto *ePref.ErrPrefixDto) (result *big.Int, resultPrecision *big.Int, err error) {

	if bIMathElectron.lock == nil {
		bIMathElectron.lock = new(sync.Mutex)
	}

	bIMathElectron.lock.Lock()

	defer bIMathElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathElectron.roundToMaxPrecisionBigInt",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)

	if bigIntNum == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigIntNum'",
			}
	}

	if bigIntNumPrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigIntNumPrecision'",
			}
	}

	if maxPrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'maxPrecision'",
			}
	}

	bigZero := big.NewInt(0)

	if bigIntNum.Cmp(bigZero) == 0 {
		return result, resultPrecision, err
	}

	if bigIntNumPrecision.Cmp(bigZero) == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("bigIntNumPrecision='%v'", bigIntNumPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'bigIntNumPrecision' is LESS THAN ZERO!",
			}
	}

	if maxPrecision.Cmp(bigZero) == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("maxPrecision='%v'", maxPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'maxPrecision' is LESS THAN ZERO!",
			}
	}

	result = big.NewInt(0).Set(bigIntNum)
	resultPrecision = big.NewInt(0).Set(bigIntNumPrecision)

	if resultPrecision.Cmp(maxPrecision) == 1 {
		delta := big.NewInt(0).Sub(resultPrecision, maxPrecision)
		delta.Sub(delta, big.NewInt(1))
		bigTen := big.NewInt(10)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		result.Quo(result, scale)
		bigFive := big.NewInt(5)
		if result.Cmp(bigZero) == -1 {
			bigFive.Neg(bigFive)
		}
		result.Add(result, bigFive)
		result.Quo(result, bigTen)
		resultPrecision = big.NewInt(0).Set(maxPrecision)
	}

	err = nil

	if trimTrailingFracZeros == false {
		return result, resultPrecision, err
	}

	if result.Cmp(bigZero) == 0 {
		resultPrecision = big.NewInt(0)
		return result, resultPrecision, err
	}

	if resultPrecision.Cmp(bigZero) == 1 {
		bigOne := big.NewInt(1)
		scrap := big.NewInt(0)
		biBase10 := big.NewInt(10)
		newProduct, mod10 := big.NewInt(0).QuoRem(result, biBase10, scrap)

		for mod10.Cmp(bigZero) == 0 && resultPrecision.Cmp(bigZero) == 1 {
			result.Set(newProduct)
			resultPrecision.Sub(resultPrecision, bigOne)
			newProduct, mod10 = big.NewInt(0).QuoRem(result, biBase10, scrap)
		}
	}

	return result, resultPrecision, err
}

// truncateToMaxPrecisionBigInt
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
func (bIMathElectron *bigIntMathElectron) truncateToMaxPrecisionBigInt(
	bigIntNum *big.Int,
	bigIntNumPrecision *big.Int,
	maxPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (result *big.Int,
	resultPrecision *big.Int, err error) {

	if bIMathElectron.lock == nil {
		bIMathElectron.lock = new(sync.Mutex)
	}

	bIMathElectron.lock.Lock()

	defer bIMathElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathElectron.truncateToMaxPrecisionBigInt",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)

	if bigIntNum == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigIntNum'",
			}
	}

	if bigIntNumPrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigIntNumPrecision'",
			}
	}

	if maxPrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'maxPrecision'",
			}
	}

	bigZero := big.NewInt(0)

	if bigIntNum.Cmp(bigZero) == 0 {

		return result, resultPrecision, nil
	}

	if bigIntNumPrecision.Cmp(bigZero) == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("bigIntNumPrecision='%v'", bigIntNumPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'bigIntNumPrecision' is LESS THAN ZERO!",
			}
	}

	if maxPrecision.Cmp(bigZero) == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("maxPrecision='%v'", maxPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'maxPrecision' is LESS THAN ZERO!",
			}
	}

	result = big.NewInt(0).Set(bigIntNum)
	resultPrecision = big.NewInt(0).Set(bigIntNumPrecision)

	if resultPrecision.Cmp(maxPrecision) == 1 {
		delta := big.NewInt(0).Sub(resultPrecision, maxPrecision)
		bigTen := big.NewInt(10)
		scale := big.NewInt(0).Exp(bigTen, delta, nil)
		result.Quo(result, scale)
		resultPrecision = big.NewInt(0).Set(maxPrecision)
	}

	return result, resultPrecision, nil
}

// truncateTrailingFracZerosBigInt
//
// Truncates trailing fractional zeros. If a numeric value has
// trailing fractional zeros, this method will delete those zeros
// and adjust the precision indicator accordingly.
//
//	Example
//	=======
//
//	            Numeric                 Result     Result    Result
//	'num'      Precision    Value       Integer   Precision  Value
//
//	12345600       5       123.45600     123456       3      123.456
func (bIMathElectron *bigIntMathElectron) truncateTrailingFracZerosBigInt(
	num *big.Int,
	numPrecision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (result *big.Int,
	resultPrecision *big.Int,
	err error) {

	if bIMathElectron.lock == nil {
		bIMathElectron.lock = new(sync.Mutex)
	}

	bIMathElectron.lock.Lock()

	defer bIMathElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathElectron.truncateTrailingFracZerosBigInt",
		"")

	if err != nil {
		return big.NewInt(0), big.NewInt(0), err
	}

	result = big.NewInt(0)
	resultPrecision = big.NewInt(0)
	err = nil

	if num == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'num'",
			}
	}

	if numPrecision == nil {

		return result, resultPrecision,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'numPrecision'",
			}
	}

	bigZero := big.NewInt(0)

	if numPrecision.Cmp(bigZero) == -1 {

		return result, resultPrecision,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("numPrecision='%v'", numPrecision.Text(10)),
				ErrMessage: "Error: Input parameter 'numPrecision' is LESS THAN ZERO!",
			}
	}

	result = big.NewInt(0).Set(num)
	resultPrecision = big.NewInt(0).Set(numPrecision)

	if result.Cmp(bigZero) == 0 {
		result = big.NewInt(0)
		resultPrecision = big.NewInt(0)
		err = nil
		return result, resultPrecision, err
	}

	if resultPrecision.Cmp(bigZero) == 1 {
		bigOne := big.NewInt(1)
		scrap := big.NewInt(0)
		biBase10 := big.NewInt(10)
		newProduct, mod10 := big.NewInt(0).QuoRem(result, biBase10, scrap)

		for mod10.Cmp(bigZero) == 0 && resultPrecision.Cmp(bigZero) == 1 {
			result.Set(newProduct)
			resultPrecision.Sub(resultPrecision, bigOne)
			newProduct, mod10 = big.NewInt(0).QuoRem(result, biBase10, scrap)
		}
	}

	err = nil

	return result, resultPrecision, err
}
