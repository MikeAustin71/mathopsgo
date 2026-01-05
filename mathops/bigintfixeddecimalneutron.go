package mathops

import (
	"fmt"
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntFixedDecNeutron struct {
	lock *sync.Mutex
}

// DivideByTenToPower
//
// Divides the numeric value of from input parameter 'bigIFxDec'
// (type BigIntFixedDecimal) by 10 to the power of 'exponent'.
//
//		result = BigIntFixedDecimal / 10^exponent
//
//	 Input Parameters
//	 ================
//
//	 bigIFxDec         *BigIntFixedDecimal
//
//	 An instance of 'BigIntFixedDecimal'. The numeric value of
//	 'bigIFxDec' will be divided by 10 to the power of 'exponent'.
//	 The resulting numeric value will be then be stored in
//	 'bigIFxDec'. The original value of 'bigIFxDec' will be
//	 overwritten.
//
//
//	 exponent          uint
//
//	 The value of BigIntFixedDecimal instance 'bigIFxDec' will be
//	 divided by ten raised to the power of 'exponent'.
//
// This method will destroy and overwrite the previous value of
// the current BigIntFixedDecimal instance with the results of
// this calculation.
//
//	NOTE
//	====
//
// This method does NOT test the validity of 'bigIFxDec', an
// instance of type BigIntNum. The calling method must
// do this!
func (bigIFdNeutron *bigIntFixedDecNeutron) divideByTenToPower(
	bigIFxDec *BigIntFixedDecimal,
	exponent uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdNeutron.lock == nil {
		bigIFdNeutron.lock = new(sync.Mutex)
	}

	bigIFdNeutron.lock.Lock()

	defer bigIFdNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecNeutron.divideByTenToPower",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}
	}

	if bigIFxDec.integerNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec.integerNum'",
		}
	}

	numSeps := NumericSeparatorDto{
		DecimalSeparator:   bigIFxDec.decimalSeparator,
		ThousandsSeparator: bigIFxDec.thousandsSeparator,
		CurrencySymbol:     bigIFxDec.currencySymbol,
	}

	numSeps.SetDefaultsIfEmpty()

	if bigIFxDec.integerNum.Cmp(big.NewInt(0)) == 0 {
		return nil
	}

	scale :=
		big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(exponent)), nil)

	factor, err := new(BigIntFixedDecimal).New(scale, 0)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "factor, err := new(BigIntFixedDecimal).New(scale, 0)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	bigIFxDec2, err := new(bigIntFixedDecUtility).copyOut(
		bigIFxDec,
		ePrefix.XCpy("Copying 'bigIFxDec' -> bigIFxDec2"))

	if err != nil {
		return err
	}

	newPrecision := bigIFxDec2.precision + exponent

	result, err :=
		new(BigIntMathDivide).FixedDecimalFracQuotient(
			bigIFxDec2, factor, numSeps, newPrecision)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "result, err := BigIntMathDivide{}.FixedDecimalFracQuotient(\n" +
				"bigIFxDec2, factor, newPrecision)",
			ErrContext: fmt.Sprintf("factor= '%v' newPrecision= '%v'", factor, newPrecision),
			ErrMessage: err.Error(),
		}
	}

	return new(bigIntFixedDecUtility).copyIn(
		bigIFxDec,
		&result,
		ePrefix.XCpy("Copying 'result' -> 'bigIFxDec'"))
}

// divideByTwoToPower
//
// Receives a BitIntFixedDecimal object via input parameter
// 'bigIFxDec'. This method then proceeds to perform integer
// division on 'bigIFxDec' by two using a 'right-shift' technique.
// Remainders from this division operation are discarded, only the
// integer quotient is returned. When the calculation is completed,
// the value of the integer quotient will replace the old value of
// the BigIntFixedDecimal instance, 'bigIFxDec'.
//
//	 Example
//	 =======
//
//		   exponent =  8
//		   quotient =  BigIntFixedDecimal / 2^(exponent)
//
//	 In this example BigIntFixedDecimal= 33,123.456, so 33,123.456/ 2^8
//
//	 (1) The fractional quotient of 33,123.456/256 (or 2^8) is 129.3885.
//
//	 (2) This method will use a right shift technique on the integer value
//	     33123456 / 2^(8) to generate an integer quotient of 129388.
//
//	     Consider the example BigIntFixedDecimal = 33123456 (no decimal fraction):
//
//	     Dividing 33123456 / 2^8 = fractional quotient = 129388.5
//
//	 Be careful when using this method.
//
// **************************************************************************
//
//		(1) Be sure to consider the outcomes when sending a decimal
//	     fraction to this method.
//
//		(2) Results returned by this method will always have
//	     precision = 0, meaning no decimal digits, only an integer
//	     value result.
//
// **************************************************************************
//
//		NOTE
//		====
//
//	 This method does NOT test the validity of 'bigIFxDec', an
//	 instance of type BigIntNum. The calling method must
//	 do this!
func (bigIFdNeutron *bigIntFixedDecNeutron) divideByTwoToPower(
	bigIFxDec *BigIntFixedDecimal,
	exponent uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bigIFdNeutron.lock == nil {
		bigIFdNeutron.lock = new(sync.Mutex)
	}

	bigIFdNeutron.lock.Lock()

	defer bigIFdNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecNeutron.divideByTenToPower",
		"")

	if err != nil {
		return err
	}

	if bigIFxDec == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}
	}

	if bigIFxDec.integerNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec.integerNum'",
		}
	}

	if bigIFxDec.integerNum.Cmp(big.NewInt(0)) == 0 {
		return nil
	}

	bigIFxDec.integerNum, err =
		new(BigIntMathDivide).BigIntDividedByTwoToPower(
			bigIFxDec.integerNum,
			exponent)

	bigIFxDec.precision = 0

	return nil
}
