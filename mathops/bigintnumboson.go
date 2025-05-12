package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumBoson struct {
	lock *sync.Mutex
}

// roundToDecimalPlace
//
// Rounds the current BigIntNum instance to a specified number of
// decimal places.
//
// 'precision' equals the number of digits to the right of the
// decimal place.
//
// Example:
//
//	integer= 123456; precision = 3; Numeric Value= 123.456
//
// If the value of BigIntNum.bigInt is zero ('0'), that zero value
// will remain unaltered. However, the BigIntNum.precision value
// will be set equal to input parameter, 'precision'.
//
// If the number of decimal places specified for rounding
// ('precision') is equal to the current BigIntNum.precision, no
// action is taken.
//
// If the number of decimal places specified for rounding
// ('precision') is greater than the current BigIntNum.precision
// value, trailing zeros are added to the current BigIntNum.bigInt
// value and BigIntNum.precision is set equal to input parameter,
// 'precision'.
//
// Finally, if the number of decimal places specified for rounding
// ('precision') is less than the current BigIntNum.precision
// value, the fractional digits will be rounded in accordance with
// the input parameter, 'precision'.
//
// Examples:
//
//		 Original      				'precision'				Resulting
//	   Value							input parameter			  Value
//	 --------------				---------------     -------------
//		654.123456								9							 654.123456000
//		654.123456								4							 654.1235
//
// -654.123456									9							-654.123456000
// -654.123456									4							-654.1235
//
//		 0												3								 0.000
//	   0.000000									0								 0
//
// Existing numeric separators (decimal separator, thousands
// separator and currency symbol) remain unchanged and are not
// altered by this method.
//
//	NOTE:
//	=====
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumBoson *bigIntNumBoson) roundToDecimalPlace(
	bNum *BigIntNum,
	precision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumBoson.lock == nil {
		bIntNumBoson.lock = new(sync.Mutex)
	}

	bIntNumBoson.lock.Lock()

	defer bIntNumBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumBoson.roundToDecimalPlace",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	if bNum.precision == precision {
		// Nothing to do. Specified 'precision' is
		// already implemented.

		return nil
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		bNum,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	numSeps, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		bNum, ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "numSeps, err := new(bigIntNumAtom).getNumericSeparatorsDto(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	// if bigInt == zero, set precision and return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		bNum2, err := new(bigIntNumMechanics).newBigInt(
			big.NewInt(0),
			precision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bNum2, err := new(bigIntNumMechanics).newBigInt(\n" +
					"    big.NewInt(0), precision, ePrefix)",
				ErrContext: "if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {",
				ErrMessage: err.Error(),
			}
		}

		err = new(bigIntNumUtility).bigIntNumCopyIn(
			bNum,
			&bNum2,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumUtility).bigIntNumCopyIn(\n" +
					"    bNum,  &bNum2, ePrefix)",
				ErrContext: "if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {",
				ErrMessage: err.Error(),
			}
		}

		err = new(bigIntNumAtom).setNumericSeparatorsDto(
			bNum,
			numSeps,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
					"    bNum, numSeps, ePrefix)",
				ErrContext: "if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	// If existing precision is less than new specified precision,
	// add trailing zeros, set new precision parameter and return.
	if bNum.precision < precision {

		deltaPrecision := precision - bNum.precision

		err = new(bigIntNumProton).bigIntNumExtendPrecision(
			bNum,
			deltaPrecision,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumProton).bigIntNumExtendPrecision(\n" +
					"    bNum, deltaPrecision, ePrefix)",
				ErrContext: "if bNum.precision < precision {",
				ErrMessage: err.Error(),
			}
		}

		err = new(bigIntNumAtom).setNumericSeparatorsDto(
			bNum,
			numSeps,
			ePrefix)

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
					"    bNum, numSeps, ePrefix)",
				ErrContext: "if bNum.precision < precision {",
				ErrMessage: err.Error(),
			}
		}

		return nil
	}

	// Must be: bNum.precision >  precision

	bigNumRound5, err := new(bigIntNumMechanics).newBigInt(
		big.NewInt(5),
		precision+1,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bigNumRound5, err := new(bigIntNumMechanics).newBigInt(\n" +
				"    big.NewInt(5), precision+1, ePrefix)",
			ErrContext: "Must be: bNum.precision > precision",
			ErrMessage: err.Error(),
		}
	}

	bigNumBase, err := new(bigIntNumMechanics).newBigInt(
		bNum.absBigInt,
		bNum.precision,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bigNumBase, err := new(bigIntNumMechanics).newBigInt(\n" +
				"    bNum.absBigInt, bNum.precision, ePrefix)",
			ErrContext: "Must be: bNum.precision > precision",
			ErrMessage: err.Error(),
		}
	}

	result, err := new(BigIntMathAdd).AddBigIntNums(
		bigNumBase, bigNumRound5)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "result, err := new(BigIntMathAdd).AddBigIntNums(\n" +
				"    bigNumBase, bigNumRound5)",
			ErrContext: "Must be: bNum.precision > precision",
			ErrMessage: err.Error(),
		}
	}

	// 10^deltaPrecision
	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision-precision)), nil)

	result.bigInt = big.NewInt(0).Quo(result.bigInt, scaleVal)

	if bNum.sign < 0 {
		result.bigInt = big.NewInt(0).Neg(result.bigInt)
	}

	err = new(bigIntNumAtom).setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsDto(\n" +
				"    bNum, numSeps, numSeps)",
			ErrContext: "10^deltaPrecision",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumNanobot).setBigInt(
		bNum,
		result.bigInt,
		precision,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: " err = new(bigIntNumNanobot).setBigInt(\n" +
				"    bNum, result.bigInt, precision, ePrefix)",
			ErrContext: "10^deltaPrecision",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
