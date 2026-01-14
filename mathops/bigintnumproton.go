package mathops

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntNumProton struct {
	lock *sync.Mutex
}

// cmpBigInt - Compares the value of the *big.Int integer to that
// contained in an incoming BigIntNum.
//
// For a true comparison of BigIntNum values see Method 'Cmp', above.
//
// Return Values:
// bNum == bigIntNum 				Return  0
// bNum > bigIntNum					Return +1
// bNum < bigIntNum					Return -1
//
//	NOTE:
//	=====
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) cmpBigInt(
	bNum *BigIntNum,
	bigIntNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.cmpBigInt",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0, &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}

	}

	if bigIntNum == nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "FATAL ERROR: Input parameter 'bNum' is a nil pointer.",
			}

	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bigIntNum,
		ePrefix.XCpy(" Validating Input Parameter 'bigIntNum2'"))

	if err != nil {
		return 0, err
	}

	cmpResult := bNum.bigInt.Cmp(bigIntNum.bigInt)

	return cmpResult, nil
}

// bigIntNumCmp - Performs a comparison of two BigIntNum numeric values
// and returns an integer value indicating the relationship
// between the two numeric values (i.e. greater than, less than,
// or equal).
//
// Note: Unlike method CmpBigInt() below, this method does more than
// just compare the root *big.Int. In making the comparision, this
// method takes into account, numeric sign values and precision. Therefore,
// this method effectively compares numeric values. As such, this method
// provides a true and comprehensive picture of the relationship between
// two BigIntNum values.
//
// Return Values:
// bNum == bigIntNum 				Return  0
// bNum > bigIntNum					Return +1
// bNum < bigIntNum					Return -1
//
//	NOTE:
//
// ================
// This method does NOT test the validity of the 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumCmp(
	bNum *BigIntNum,
	bigIntNum2 *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumCmp",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}
	}

	if bigIntNum2 == nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "FATAL ERROR: Input parameter 'bigIntNum2' is a nil pointer.",
			}

	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		bigIntNum2,
		ePrefix.XCpy(" Validating Input Parameter 'bigIntNum2'"))

	if err != nil {
		return 0, err
	}

	b1Sign := bNum.sign
	b2Sign := bigIntNum2.sign

	if b1Sign != b2Sign {

		if b1Sign > b2Sign {

			return b1Sign, nil
		}

		return b2Sign, nil
	}

	// The signs must be equal

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return 0,
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				"  bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(bNum, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	difference, err := new(BigIntMathSubtract).SubtractBigIntNums(
		bNum2,
		*bigIntNum2)

	if err != nil {

		return 0,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " difference, err := BigIntMathSubtract{}.\n" +
					"    SubtractBigIntNums(bNum2, *bigIntNum2)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		&difference,
		ePrefix.XCpy(" Validating calculation result 'difference'"))

	if err != nil {

		return 0, err
	}

	var differenceIsZero bool

	differenceIsZero, err = new(bigIntNumMolecule).isBIntNumZero(
		&difference,
		ePrefix)

	if err != nil {

		return 0, err
	}

	if differenceIsZero {
		return 0, nil
	}

	return difference.sign, nil
}

// bigIntNumDecrement - Subtracts a value of +1 (plus one) from the numeric
// value of the current BigIntNum instance.
//
// The numeric separators (decimal separator, thousands separator
// and currency symbol) from the original BigIntNum will remain
// unchanged.
//
//	NOTE:
//
// ================
// This method does NOT test the validity of the 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDecrement(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDecrement",
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

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {
		return err
	}

	biNumOne, err := new(bigIntNumMolecule).newOne(
		bNum.precision,
		ePrefix)

	if err != nil {
		return err
	}

	bPair, err := new(BigIntPair).NewBigIntNum(bNum2, biNumOne)

	if err != nil {
		return fmt.Errorf("%v\n%w", ePrefix, err)
	}

	result, err := new(BigIntMathSubtract).SubtractPair(bPair)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "result, err := new(BigIntMathSubtract).SubtractPair(bPair)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumUtility).bigIntNumCopyIn(
		bNum,
		&result,
		ePrefix)

	return err
}

// bigIntNumDivide - Performs a division operation. The current BigIntNum instance is the
// 'dividend' divided by the input parameter, 'divisor'. The result of this division
// operation is the 'fracQuotient' which is returned as a BigIntNum type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//		bNum = dividend
//	 -----------------------------
//		dividend / divisor = quotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// quotient. Precision is defined as the number of fractional digits to the right of
// the decimal place. Be advised that these calculations can support very large precision
// values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDivide(
	bNum *BigIntNum,
	divisor *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (fracQuotient BigIntNum, err error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}

	}

	if divisor == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'divisor'",
			}

	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		divisor,
		ePrefix.XCpy("Validating 'divisor'"))

	if err != nil {

		return BigIntNum{}, err
	}

	if divisor.bigInt.Cmp(big.NewInt(0)) == 0 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ErrContext: "if divisor.bigInt.Cmp(big.NewInt(0)) == 0",
				ErrMessage: "Error: Attempted Divide by ZERO!\n" +
					"The value of input parameter 'divisor' is ZERO.",
			}
	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bNum2,
		ePrefix.XCpy("bNum2.NumSeps->numSepsDto"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	return new(BigIntMathDivide).BigIntNumFracQuotient(
		bNum2,
		*divisor,
		numSepsDto,
		maxPrecision)
}

// bigIntNumDivideByFive - Divides the numerical value of the current BigIntNum by five ('5').
// The result of this division operation is the 'fracQuotient' which is returned as a
// BigIntNum type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 5 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
//
//	NOTE:
//
// ================
// This method does NOT test the validity of the 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDivideByFive(
	bNum *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (fracQuotient BigIntNum, err error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bNum2,
		ePrefix.XCpy("bNum2.NumSeps->numSepsDto"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2Str, err := new(bigIntNumAtom).getBigIntNumStr(
		&bNum2,
		ePrefix.XCpy("bNum2->bNum2Str"))

	if err != nil {

		return BigIntNum{}, err
	}

	fracQuotient, err =
		new(BigIntMathDivide).
			BigIntNumDivideByFiveFracQuo(bNum2, numSepsDto, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumDivideByFiveFracQuo(bNum2, numSepsDto, maxPrecision)",
				ErrContext: fmt.Sprintf("bNum2= '%v; maxPrecision= '%v'",
					bNum2Str, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// DivideByTen - Divides the numerical value of the current BigIntNum by ten ('10'). The
// result of this division operation is the 'fracQuotient' which is returned as a BigIntNum
// type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 10 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
//
//	NOTE:
//
// ================
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDivideByTen(
	bNum *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (fracQuotient BigIntNum, err error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivideByTen",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix.XCpy("bNum Copy Out > bNum2"))

	if err != nil {

		return BigIntNum{}, err
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bNum2,
		ePrefix.XCpy("bNum2.NumSeps->numSepsDto"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	fracQuotient, err =
		new(BigIntMathDivide).BigIntNumDivideByTenFracQuo(
			bNum2, numSepsDto, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumDivideByTenFracQuo(\n" +
					"    bNum2, numSepsDto, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, nil
}

// DivideByTenToPower - Divides the numerical value of the current BigIntNum
// instance by 10 raised to the power of the input parameter, 'exponent'.
//
//	bNum = bNum / (10^exponent)
//
// The original value of the current BigIntNum will be destroyed and overwritten
// by this method.
//
// The final BigIntNum will retain the original numeric separators (decimal separator,
// thousands separator and currency symbol).
//
//	NOTE:
//
// ================
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDivideByTenToPower(
	bNum *BigIntNum,
	exponent uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
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

	newPrecision := bNum.precision + exponent

	result, err := new(bigIntNumMechanics).newBigInt(
		bNum.bigInt,
		newPrecision,
		ePrefix)

	if err != nil {
		return err
	}

	err = new(bigIntNumUtility).bigIntNumCopyIn(
		bNum,
		&result,
		ePrefix)

	return err
}

// bigIntNumDivideByThree - Divides the numerical value of the current BigIntNum by three ('3').
// The result of this division operation is the 'fracQuotient' which is returned as a
// BigIntNum type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 3 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// This returned BigIntNum 'fracQuotient' will contain numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum instance (bNum).
//
//	NOTE:
//
// ================
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDivideByThree(
	bNum *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (fracQuotient BigIntNum, err error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivideByThree",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bNum2,
		ePrefix.XCpy("bNum2.NumSeps->numSepsDto"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	fracQuotient, err =
		new(BigIntMathDivide).
			BigIntNumDivideByThreeFracQuo(bNum2, numSepsDto, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumDivideByThreeFracQuo(\n" +
					"    bNum2, numSepsDto, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, err
}

// bigIntNumDivideByTwo - Divides the numerical value of the current BigIntNum by two ('2').
// The result of this division operation is the 'fracQuotient' which is returned as a
// BigIntNum type.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
// The return value may be an integer or a floating point value as determined by the results
// of the division operation. For floating point values, the number of digits to the right
// of the decimal place is limited by input parameter, 'maxPrecision'.
//
//	bNum / 2 = fracQuotient
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// fractional quotient. Precision is defined as the number of fractional digits to the
// right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDivideByTwo(
	bNum *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (fracQuotient BigIntNum, err error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	fracQuotient = new(bigIntNumMechanics).new()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
		"")

	if err != nil {
		return fracQuotient, err
	}

	if bNum == nil {

		return fracQuotient,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}
	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {
		return fracQuotient, err
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bNum2,
		ePrefix.XCpy("bNum2.NumSeps->numSepsDto"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2Str, err := new(bigIntNumAtom).getBigIntNumStr(
		&bNum2,
		ePrefix.XCpy("bNum2->bNum2Str"))

	if err != nil {

		return BigIntNum{}, err
	}

	fracQuotient, err =
		new(BigIntMathDivide).BigIntNumDivideByTwoFracQuo(bNum2, numSepsDto, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumDivideByTwoFracQuo(bNum2, numSepsDto, maxPrecision)",
				ErrContext: fmt.Sprintf("bNum2= '%v; maxPrecision= '%v'",
					bNum2Str, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return fracQuotient, err
}

// bigIntNumDivideByTwoQuoMod - Divides the numerical value of the current BigIntNum by
// two ('2'). The result of the division operation is returned as an integer quotient,
// 'intQuotient', and a floating point modulo or remainder, 'modulo'.
//
//	bNum / 2 = integer quotient and floating point modulo
//
// If 'modulo' equals zero ('0'), it signals the current BigIntNum numerical value is
// 'even'; that is, it is evenly divisible by two.
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// floating point 'modulo'. Precision is defined as the number of fractional digits to
// the right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// Both returned BigIntNum 'intQuotient' and 'modulo' BigIntNum types will contain numeric
// separators (decimal separator, thousands separator and currency symbol) copied from the
// current BigIntNum instance (bNum).
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumDivideByTwoQuoMod(
	bNum *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (intQuotient BigIntNum, modulo BigIntNum, err error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
		"")

	if err != nil {
		return intQuotient, modulo, err
	}

	if bNum == nil {

		return intQuotient, modulo,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {
		return intQuotient, modulo, err
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bNum2,
		ePrefix.XCpy("bNum2.NumSeps->numSepsDto"))

	if err != nil {
		return BigIntNum{}, BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	bNum2Str, err := new(bigIntNumAtom).getBigIntNumStr(
		&bNum2,
		ePrefix.XCpy("bNum2->bNum2Str"))

	if err != nil {

		return BigIntNum{}, BigIntNum{}, err
	}

	intQuotient, modulo, err =
		new(BigIntMathDivide).BigIntNumDivideByTwoQuoMod(
			bNum2, numSepsDto, maxPrecision)

	if err != nil {

		return BigIntNum{}, BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "fracQuotient, err = new(BigIntMathDivide).\n" +
					"    BigIntNumDivideByTwoQuoMod(bNum2, numSepsDto, maxPrecision)",
				ErrContext: fmt.Sprintf("bNum2= '%v; maxPrecision= '%v'",
					bNum2Str, maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return intQuotient, modulo, nil
}

// ExtendPrecision - Extends the current precision.
//
// Precision is the number of fractional digits to the right
// of the decimal place. This method will extend the number of
// digits to the right of the decimal place by adding trailing
// zeros to the current numeric value of this 'BigIntNum' instance.
// The number of trailing zeros to be added is determined by the
// input parameter, 'deltaPrecision'.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bIntNumProton *bigIntNumProton) bigIntNumExtendPrecision(
	bNum *BigIntNum,
	deltaPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumExtendPrecision",
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

	bigAtom := new(bigIntNumAtom)

	if deltaPrecision == 0 {
		return nil
	}

	err = bigAtom.setNumericSeparatorsToDefaultIfEmpty(
		bNum,
		ePrefix)

	if err != nil {
		return err
	}

	numSeps, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		bNum,
		ePrefix)

	if err != nil {
		return err
	}

	newPrecision := bNum.precision + deltaPrecision

	var bigIUtil = new(bigIntNumUtility)

	var bigIMech = new(bigIntNumMechanics)

	// bigInt == zero, set precision and return
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		bINum3, err := bigIMech.newBigInt(
			big.NewInt(0),
			newPrecision,
			ePrefix)

		if err != nil {
			return err
		}

		err = bigIUtil.bigIntNumCopyIn(
			bNum,
			&bINum3,
			ePrefix)

		if err != nil {
			return err
		}

		err = bigAtom.setNumericSeparatorsDto(
			bNum,
			numSeps,
			ePrefix)

		return err
	}

	base10 := big.NewInt(10)

	scaleVal := big.NewInt(0).Exp(base10, big.NewInt(int64(deltaPrecision)), nil)

	bigINum := big.NewInt(0).Set(bNum.bigInt)

	bigINum = big.NewInt(0).Mul(bigINum, scaleVal)

	bNumResult, err := bigIMech.newBigInt(
		bigINum,
		newPrecision,
		ePrefix)

	if err != nil {
		return err
	}

	err = bigIUtil.bigIntNumCopyIn(
		bNum,
		&bNumResult,
		ePrefix)

	if err != nil {
		return err
	}

	err = bigAtom.setNumericSeparatorsDto(
		bNum,
		numSeps,
		ePrefix)

	return err
}

// bigIntNumFloor - returns the greatest integer less than or
// equal to the numeric value of the current BigIntNum. Reference
// Wikipedia:
// https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//							Initial 			Floor
//	 					 Value				Value
//							-------      -------
//	 						5.95					5
//	 						5.05					5
//	 						5							5
//						 -5.05			 	 -6
//	 						2.4				  	2
//	 						2.9					 	2
//						 -2.7				 	 -3
//						 -2					 	 -2
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumFloor(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumFloor",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	var bigIntMech = new(bigIntNumMechanics)

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {

		bNum2, err := bigIntMech.newBigInt(
			big.NewInt(0),
			0,
			ePrefix)

		if err != nil {
			return BigIntNum{}, err
		}

		return bNum2, nil
	}

	if bNum.precision == 0 {

		bNum3, err := new(bigIntNumUtility).bigIntNumCopyOut(
			bNum,
			ePrefix)

		if err != nil {
			return BigIntNum{}, err
		}

		return bNum3, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision)), nil)

	absQuotient := big.NewInt(0).Quo(bNum.absBigInt, scaleVal)

	// absQuotient IS NOT EQUAL TO bNum.absBigInt

	if bNum.sign > 0 {
		// bNum is positive

		bNum4, err := bigIntMech.newBigInt(
			absQuotient,
			0,
			ePrefix)

		if err != nil {
			return BigIntNum{}, err
		}

		return bNum4, nil
	}

	// bNum is negative
	absQuotient =
		big.NewInt(0).Add(absQuotient, big.NewInt(1))

	bNum5, err := bigIntMech.newBigInt(
		big.NewInt(0).Neg(absQuotient),
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	return bNum5, nil
}

// bigIntNumGetAbsoluteNumStr
//
//	Returns the absolute integer value (positive value) of the
//	*big.Int value encapsulated by this BigIntNum.
//
//	No decimal place is included.
//
//	NOTE:
//	=====
//
//	This method does NOT test the validity of 'bNum' BigIntNum
//	instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetAbsoluteNumStr(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
		"")

	if err != nil {
		return "", err
	}

	if bNum == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	var bigINumMolecule = new(bigIntNumMolecule)

	var bigINumUtility = new(bigIntNumUtility)

	if bNum.sign == 1 {
		// bNum is positive
		numStr, err := bigINumMolecule.formatBigIntNumStr(
			bNum,
			LEADMINUSNEGVALFMTMODE,
			ePrefix)

		if err != nil {

			return "", err
		}

		return numStr, nil
	}

	// bNum is negative
	biNum, err := bigINumUtility.bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return "",
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" biNum, err := new(bigIntNumUtility).bigIntNumCopyOut(bNum, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())

	}

	err = bigINumUtility.bigIntNumChangeSign(
		&biNum,
		ePrefix.XCpy("Changing sign on 'biNum'"))

	if err != nil {

		return "", err
	}

	numStr, err := bigINumMolecule.formatBigIntNumStr(
		&biNum,
		LEADMINUSNEGVALFMTMODE,
		ePrefix)

	return numStr, err
}

// bigIntNumGetAbsoluteBigIntNumValue
//
//	Returns the absolute numeric value of this BigIntNum instance
//	as a new BigIntNum Type.
//
//	If the current BigIntNum value is'-123.456', this method will
//	return '123.456'.
//
//	If the current BigIntNum value is'123.456', this method will
//	return '123.456'.
func (bIntNumProton *bigIntNumProton) bigIntNumGetAbsoluteBigIntNumValue(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetAbsoluteBigIntNumValue",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	bNum2, err := new(bigIntNumMechanics).newBigInt(
		bNum.absBigInt,
		bNum.precision,
		ePrefix.XCpy(fmt.Sprintf("bNum.absBigInt= '%v'\nbNum.precision= '%v'",
			bNum.absBigInt.Text(10), bNum.precision)))

	if err != nil {

		return BigIntNum{}, err
	}

	return bNum2, nil
}

// bigIntNumGetAbsoluteBigIntValue - returns the absolute
// value of the *big.Int value encapsulated by the input
// paramter 'bNum' (type BigIntNum).
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetAbsoluteBigIntValue(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	result := big.NewInt(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
		"")

	if err != nil {
		return result, err
	}

	if bNum == nil {

		return result,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	result = big.NewInt(0).Set(bNum.absBigInt)

	return result, nil
}

// bigIntNumGetBigFloat - Returns the numeric value of
// the input parameter BigIntNum as *big.Float type.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetBigFloat(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Float, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetBigFloat",
		"")

	if err != nil {
		return big.NewFloat(0), err
	}

	if bNum == nil {

		return big.NewFloat(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	numerator := big.NewInt(0).Set(bNum.bigInt)

	denominator := big.NewInt(0).Set(bNum.scaleFactor)

	bRat := big.NewRat(1, 1).SetFrac(numerator, denominator)

	return big.NewFloat(0).SetRat(bRat), nil
}

// GetBigInt - return the numeric value as an integer
// of type *big.int.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetBigInt(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetBigInt",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if bNum == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	return big.NewInt(0).Set(bNum.bigInt), nil
}

// bigIntNumGetBigIntFixedDecimal -
// Returns a BigIntFixedDecimal instance which contains a
// copy of the current BigIntNum values.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetBigIntFixedDecimal(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntFixedDecimal, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetBigIntFixedDecimal",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	if bNum == nil {

		return BigIntFixedDecimal{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	fixedDec, err := new(BigIntFixedDecimal).New(bNum.bigInt, bNum.precision)

	return fixedDec, nil
}

// bigIntNumGetBigRat - Returns the numeric value of
// the passed BigIntNum instance as a type *big.Rat type.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum'
// BigIntNum instance. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetBigRat(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Rat, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetBigRat",
		"")

	if err != nil {
		return big.NewRat(1, 1), err
	}

	if bNum == nil {

		return big.NewRat(1, 1),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	numerator := big.NewInt(0).Set(bNum.bigInt)

	denominator := big.NewInt(0).Set(bNum.scaleFactor)

	return big.NewRat(1, 1).SetFrac(numerator, denominator), nil
}

// bigIntNumGetNumSepSymbol - Returns the character
// currently designated as the separator symbol for
// 'decimal', 'thousands' or 'currency'
//
// The specific symbol returned will depend on the
// input parameter 'numSepSymbol' which is of type
// 'NumSepSymbolCode'
//
// The 'NumSepSymbolCode' must be set to one of
// three constant values:
//
// DECIMALSYMBOL - Returns the decimal symbol
// encapsulated in the BigIntNum passed as
// input parameter 'bNum'.
//
// THOUSANDSYMBOL - Returns the 'thousands' symbol
// encapsulated in input parameter 'bNum.
//
// CURRENCYSYMBOL - Returns the 'Currency' symbol
// encapsulated in input parameter 'bNum.
//
// For a list of Major Currency Unicode Symbols, see constants
// located in: MikeAustin71/mathopsgo/mathops/mathopsconstants.go
//
// USA Example: $123.45
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetNumSepSymbol(
	bNum *BigIntNum,
	nSepSymbol NumSepSymbolCode,
	errPrefDto *ePref.ErrPrefixDto) (rune, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumDivide",
		"")

	if err != nil {
		return '0', err
	}

	if bNum == nil {

		return '0',
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}

	}

	switch nSepSymbol {

	case DECIMALSYMBOL:
		return bNum.decimalSeparator, nil

	case THOUSANDSYMBOL:
		return bNum.thousandsSeparator, nil

	case CURRENCYSYMBOL:
		return bNum.currencySymbol, nil

	}

	err = fmt.Errorf("%v\n"+
		"FATAL ERROR: Input parameter 'nSepSymbol' did not\n"+
		"match a valid Number Separator Symbol Code.\n"+
		"Valid codes are:\n"+
		"DECIMALSYMBOL, THOUSANDSYMBOL or CURRENCYSYMBOL\n",
		ePrefix.String())

	return '0', err
}

// bigIntNumGetFractionalPart
//
// Returns the fractional digits of the BigIntNum passed
// as input parameter 'bNum'. The fractional digit are
// returned in the form of a new BigIntNum instance
// containing those correctly formatted fractional digits.
//
// Examples
// ========
//
//				 Current
//				BigIntNum				 		Return
//	 			  Value						  Value
//				----------				---------
//
//	 			123.456						 0.456
//				 -123.456					-0.456
//				  123								 0
//				 -123								 0
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetFractionalPart(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetFractionalPart",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}
	}

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		// bNum value is ZERO!

		bNum2, err := new(bigIntNumMechanics).newZero(
			0,
			ePrefix)

		if err != nil {

			return BigIntNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: " bNum2, err := new(bigIntNumMechanics).\n" +
						"   newZero(0, ePrefix)",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		return bNum2, nil
	}

	scaleVal := big.NewInt(0).Exp(big.NewInt(10),
		big.NewInt(int64(bNum.precision)), nil)

	modulo := big.NewInt(0).Rem(bNum.bigInt, scaleVal)

	bNum3, err := new(bigIntNumMechanics).newBigInt(
		modulo, bNum.precision, ePrefix)

	if err != nil {

		return BigIntNum{},
			fmt.Errorf("%v\n"+
				"Error returned by: \n"+
				" bNum3, err := new(bigIntNumMechanics).\n"+
				"  newBigInt(modulo, bNum.precision, ePrefix)\n"+
				"Error= %v\n",
				ePrefix,
				err.Error())
	}

	return bNum3, nil
}

// bigIntNumGetInt - Returns a type 'int' containing the 32-bit integer
// value of the current BigIntNum instance.
//
// If the current BigIntNum value is greater than the maximum
// 'int' value, the maximum 32-bit integer value is returned
// in addition to an 'error'.
//
// If the current BigIntNum value is less than the minimum 'int'
// value, the minimum 32-bit integer value is returned along with
// an 'error'.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetInt(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetInt",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}
	}

	var numStr string

	numStr, err = new(bigIntNumAtom).getBigIntNumStr(
		bNum,
		ePrefix)

	if err != nil {
		return 0, err
	}

	bIMaxInt := big.NewInt(int64(math.MaxInt32))

	bIMinInt := big.NewInt(int64(math.MinInt32))

	if bNum.bigInt.Cmp(bIMaxInt) == 1 {

		return math.MaxInt32,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "Error: BigIntNum Value is GREATER than Int32 Maximum!\n" +
					fmt.Sprintf("Int32 Maximum Value= '%v'\nBigIntNum Value='%v'",
						math.MaxInt32, numStr),
			}
	}

	if bNum.bigInt.Cmp(bIMinInt) == -1 {

		return math.MinInt32,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "Error: BigIntNum Value is LESS than Int32 Minmum!\n" +
					fmt.Sprintf("Int32 Minimum Value+ '%v'\nBigIntNum Value='%v'",
						math.MinInt32, numStr),
			}
	}

	return int(bNum.bigInt.Int64()), nil
}

// bigIntNumGetIntAry - Converts the current BigIntNum value to an
// IntAry instance. The resulting number value includes the decimal
// place and fractional digits if they exist.
//
// Note that the BigIntNum settings for 'decimalSeparator', 'thousandsSeparator'
// and 'currencySymbol' are transferred to the new IntAry instance returned to the
// calling function.
//
// The returned IntAry type contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum
// instance.
//
// This method performs a validity test on the current BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetIntAry(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (IntAry, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetIntAry",
		"")

	if err != nil {
		return IntAry{}, err
	}

	if bNum == nil {

		return IntAry{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'bNum'",
			}
	}

	ia, err := new(IntAry).NewBigInt(big.NewInt(0).Set(bNum.bigInt), int(bNum.precision))

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "ia, err := new(IntAry).NewBigInt(big.NewInt(0).\n" +
					"Set(bNum.bigInt), int(bNum.precision))",
				ErrContext: fmt.Sprintf("bNum.bigInt='%v'\nbNum.precision='%v'",
					bNum.bigInt.Text(10), bNum.precision),
				ErrMessage: err.Error(),
			}

	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		bNum,
		ePrefix)

	if err != nil {
		return IntAry{}, err
	}

	err = ia.SetNumericSeparatorsDto(numSepsDto)

	if err != nil {

		return IntAry{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(\n" +
					"  getNumericSeparatorsDto(bNum, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = ia.IsValid(ePrefix.String() + " Testing Validity of ia (IntAry).\n")

	if err != nil {
		return IntAry{}, err
	}

	return ia, nil
}

// GetIntegerValue
// Returns the internal *big.Int number for the current
// BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetIntegerValue(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetIntegerValue",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if bNum == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	return big.NewInt(0).Set(bNum.bigInt), nil
}

// bigIntNumGetInverse
//
// Returns the value of one (1) divided by the input parameter
// 'bNum' (type BigIntNum). The result is returned as new
// BigIntNum Type.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetInverse(
	bNum *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetFractionalPart",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	bINumOne, err := new(bigIntNumMechanics).
		newBigInt(big.NewInt(1), 0, ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	numSepsDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		&bNum2,
		ePrefix.XCpy("bNum2.NumSeps->numSepsDto"))

	if err != nil {
		return BigIntNum{}, err
	}

	numSepsDto.SetDefaultsIfEmpty()

	result, err := new(BigIntMathDivide).
		BigIntNumFracQuotient(bINumOne, bNum2, numSepsDto, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "result, err := new(BigIntMathDivide).\n" +
					"    BigIntNumFracQuotient(\n" +
					"    bINumOne, bNum2, numSepsDto, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}

// bigIntNumGetNumStrDto
//
// Converts the current BigIntNum value to a NumStrDto instance.
// The resulting number string includes the decimal place and
// decimal digits if they exist.
//
// The returned NumStrDto type contains numeric separators (decimal
// separator, thousands separator, and currency symbol) copied from
// the current BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetNumStrDto(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (NumStrDto, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetNumStrDto()",
		"")

	if err != nil {
		return NumStrDto{}, err
	}

	if bNum == nil {

		return NumStrDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}

	}

	nDto, err := new(NumStrDto).NewBigInt(big.NewInt(0).Set(bNum.bigInt), bNum.precision)

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "nDto, err := new(NumStrDto).NewBigInt(big.NewInt(0)\n" +
					".Set(bNum.bigInt), uint(bNum.precision))",
				ErrContext: fmt.Sprintf("bNum.bigInt='%v'\nbNum.precision='%v'",
					bNum.bigInt.Text(10), bNum.precision),
				ErrMessage: err.Error(),
			}
	}

	numSepDto, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		bNum,
		ePrefix)

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "numSepDto, err := new(bigIntNumAtom).\n" +
					"  getNumericSeparatorsDto(bNum, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	numSepDto.SetDefaultsIfEmpty()

	err = nDto.SetNumericSeparatorsDto(numSepDto)

	if err != nil {
		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = nDto.SetNumericSeparatorsDto(numSepDto)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = nDto.IsValid(ePrefix.String() + "'nDto' FAILED Validation Test! ")

	if err != nil {

		return NumStrDto{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = nDto.IsValid(ePrefix.String() + \"'nDto' FAILED Validation Test! \")",
				ErrContext: "Final nDto instance is INVALID!",
				ErrMessage: err.Error(),
			}
	}

	return nDto, nil
}

// bigIntNumGetPrecision
//
// Returns the precision associated with the instance of
// BigIntNum passed as input parameter 'bNum'. The returned
// precision value is formatted as an integer of type 'int'.
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal place in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//					1.234    	GetPrecisionInt() = 3
//							5			GetPrecisionInt() = 0
//				0.12345  		GetPrecisionInt() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetPrecision(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetPrecision",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	bIntMaxInt :=
		big.NewInt(0).SetUint64(uint64(math.MaxInt))

	bUintPrecision :=
		big.NewInt(0).SetUint64(uint64(bNum.precision))

	if bUintPrecision.Cmp(bIntMaxInt) == 1 {

		return 0, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("Max 'int' value= '%v'\n'Precision= '%v'",
				bIntMaxInt.Text(10), bUintPrecision.Text(10)),
			ErrMessage: "Error: The value of uint 'Precision' exceeds the limit for int values!",
		}
	}

	return int(bNum.precision), nil
}

// bigIntNumGetPrecisionBigInt
//
// Returns the 'precision' of the BigIntNum instance passed
// as input parameter 'bNum'. The precision value is returned
// formatted as a *big.Int Type.
//
// 'precision' is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal place in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//		1.234    	GetPrecisionBigInt() = 3
//				5			GetPrecisionBigInt() = 0
//	0.12345  		GetPrecisionBigInt() = 5
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetPrecisionBigInt(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetPrecisionBigInt",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if bNum == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	return big.NewInt(0).SetUint64(uint64(bNum.precision)), nil

}

// bigIntNumGetPrecisionUint
//
// Returns precision of the BigIntNum instance passed as
// input parameter 'bNum'. The precision value is returned
// formatted as an unsigned integer (uint).
//
// precision is defined as the number of numeric digits to
// the right of the decimal place. To compute the location
// of the decimal place in a string of numeric digits, go
// to the right most digit in the number string and count
// left 'precision' digits.
//
// Example:
//
//					1.234    	GetPrecisionUint() = 3
//							5			GetPrecisionUint() = 0
//				0.12345  		GetPrecisionUint() = 5
//
//	Number String				precision				Fractional Number
//		123456								3								123.456
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetPrecisionUint(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (uint, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetPrecisionUint",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	return bNum.precision, nil
}

// bigIntNumGetScaleFactor
//
// Returns the scale value of the BigIntNum instance passed
// as input parameter 'bNum'.  Scale value is a function of
// 'precision', or the number of digits to the right of the
// decimal place. Therefore, scale factor is defined by 10
// raised to the power of BigIntNum precision.
//
// Example:
// precision = 0 		Scale Factor = 10^0   	Scale Factor =    1
// precision = 1		Scale Factor = 10^1			Scale Factor =   10
// precision = 2		Scale Factor = 10^2			Scale Factor =  100
// precision = 3    Scale Factor = 10^3			Scale Factor = 1000
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetScaleFactor(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetScaleFactor",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if bNum == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	return big.NewInt(0).Set(bNum.scaleFactor), nil
}

// GetSign - Returns the numeric sign associated
// with the current numeric value encapsulated by
// this BigIntNum.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetSign(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetSign",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	return bNum.sign, nil
}

// bigIntNumGetSignedBigInt
// Returns the integer value of the BigIntNum instance
// passed as input parameter 'bNum'. This value is
// formatted and returned as a signed *big.Int Type.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetSignedBigInt(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (*big.Int, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetSignedBigInt",
		"")

	if err != nil {
		return big.NewInt(0), err
	}

	if bNum == nil {

		return big.NewInt(0),
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	return bNum.bigInt, nil
}

// bigIntNumGetSciNotationNumber
//
// Converts the numeric value of the BigIntNum instance passed
// as input parameter 'bNum'. This converted value is formatted
// as scientific notation and returned as an instance of type
// SciNotationNum (Scientific Notation).
//
// Input Parameter
// ===============
//
// mantissaLen uint	- Specifies the length of the mantissa in the returned
//
//											scientific notation string. If the value of 'mantissaLen'
//											is less than two ('2'), this method will automatically set
//											the 'mantissaLen' to a default value of two ('2').
//
//											Example Scientific Notation:
//											----------------------------
//
//	 										scientific notation string: '2.652e+8'
//
//	 										significand = '2.652'
//	 										significand integer digit = '2'
//												mantissa		= significand factional digits = '.652'
//	 										exponent    = '8'  (10^8)
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetSciNotationNumber(
	bNum *BigIntNum,
	mantissaLen uint,
	errPrefDto *ePref.ErrPrefixDto) (SciNotationNum, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetSciNotationNumber",
		"")

	if err != nil {
		return SciNotationNum{}, err
	}

	if bNum == nil {

		return SciNotationNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	sciNotationNum := new(SciNotationNum).New()

	if mantissaLen < 2 {
		mantissaLen = 2
	}

	var bNumIsZero bool

	bNumIsZero, err = new(bigIntNumMolecule).
		isBIntNumZero(bNum, ePrefix)

	if err != nil {

		return SciNotationNum{}, err
	}

	if bNumIsZero {

		bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
			bNum,
			ePrefix)

		if err != nil {

			return SciNotationNum{}, err

		}

		bNumZero, err := new(bigIntNumMechanics).newZero(0, ePrefix)

		if err != nil {

			return SciNotationNum{}, err
		}

		err = sciNotationNum.SetBigIntNumElements(bNum2, bNumZero)

		if err != nil {

			return SciNotationNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "sciNotationNum.SetBigIntNumElements(bNum2, bNumZero)",
					ErrMessage: err.Error(),
				}
		}

		return sciNotationNum, nil
	}

	bigIntMaxUint32 := big.NewInt(0).SetUint64(math.MaxUint32)

	bINumIntPart, err := new(bigIntNumAtom).getIntegerPart(
		bNum,
		ePrefix)

	if err != nil {

		return SciNotationNum{}, err
	}

	bNumIsZero, err = new(bigIntNumMolecule).
		isBIntNumZero(&bINumIntPart, ePrefix)

	if err != nil {

		return SciNotationNum{}, err

	}

	if !bNumIsZero {

		magnitudeBigInt, err := new(BigIntMath).GetMagnitude(bINumIntPart.bigInt)

		if err != nil {

			return SciNotationNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "  magnitudeBigInt, err := BigIntMath{}.GetMagnitude(\n" +
						"    bINumIntPart.bigInt)",
					ErrContext: fmt.Sprintf("bINumIntPart.bigInt= '%v'",
						bINumIntPart.bigInt.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		if magnitudeBigInt.Cmp(bigIntMaxUint32) == 1 {

			return SciNotationNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ErrContext: "if magnitudeBigInt.Cmp(bigIntMaxUint32) == 1",
					ErrMessage: "Integer Magnitude greater than Max Uint32!",
				}
		}

		uintMagnitude := uint(magnitudeBigInt.Uint64())

		newBINum, err := new(bigIntNumMechanics).newBigInt(
			bNum.bigInt,
			bNum.precision+uintMagnitude,
			ePrefix)

		if err != nil {

			return SciNotationNum{}, err
		}

		biNumExponent, err := new(bigIntNumMechanics).
			newBigInt(magnitudeBigInt, 0, ePrefix)

		if err != nil {

			return SciNotationNum{}, err
		}

		err = sciNotationNum.SetBigIntNumElements(
			newBINum, biNumExponent)

	} else {

		// Must be bINumFracPart > 0
		magnitudeBigInt, err := new(BigIntMath).GetMagnitude(bNum.bigInt)

		if err != nil {
			return SciNotationNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "  magnitudeBigInt, err := BigIntMath{}.\n" +
						"    GetMagnitude(bNum.bigInt)",
					ErrContext: fmt.Sprintf("bNum.bigInt= '%v'",
						bNum.bigInt.Text(10)),
					ErrMessage: err.Error(),
				}
		}

		if magnitudeBigInt.Cmp(bigIntMaxUint32) == 1 {

			return SciNotationNum{},
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ErrContext: "if magnitudeBigInt.Cmp(bigIntMaxUint32) == 1",
					ErrMessage: "Fractional Magnitude greater than Max Uint32!",
				}
		}

		uintMagnitude := uint(magnitudeBigInt.Uint64())

		bINumFracPart, err := new(bigIntNumMechanics).
			newBigInt(bNum.bigInt, uintMagnitude, ePrefix)

		if err != nil {

			return SciNotationNum{}, err
		}

		precisionFrac := int64(uintMagnitude) - int64(bNum.precision)

		bINumScale, err := new(bigIntNumMechanics).
			newInt64Exponent(precisionFrac, 0, ePrefix)

		if err != nil {

			return SciNotationNum{}, err
		}

		err = sciNotationNum.SetBigIntNumElements(bINumFracPart, bINumScale)

		if err != nil {

			return SciNotationNum{},
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "  err = sciNotationNum.SetBigIntNumElements(\n" +
						"    bINumFracPart, bINumScale)",
					ErrMessage: err.Error(),
				}
		}

	}

	sciNotationNum.SetMantissaLength(mantissaLen)

	return sciNotationNum, nil
}

// bigIntNumGetUInt
//
// Returns a type 'uint' containing the 32-bit unsigned
// integer value of the BigIntNum instance passed as input
// parameter 'bNum'.
//
// If the current BigIntNum value is greater than the maximum
// 'uint' value, the maximum 32-bit unsigned integer value is returned
// in addition to an 'error'.
//
// If the current BigIntNum value is less than the minimum 'uint'
// value, the minimum 32-bit integer value of zero is returned along
// with an 'error'.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetUInt(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (uint, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetUInt",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	if bNum.bigInt == nil {

		return 0,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "Input parameter 'bNum' is INVALID!\n" +
					"bNum.bigInt == nil",
			}

	}

	bNumBigIntNumStr := bNum.bigInt.Text(10)

	bIMaxUint := big.NewInt(int64(math.MaxUint32))

	bIMaxUintNumStr := bIMaxUint.Text(10)

	if bNum.bigInt.Cmp(big.NewInt(0)) == -1 {

		return uint(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if bNum.bigInt.Cmp(big.NewInt(0)) == -1",
				ErrMessage: "BigIntNum is LESS THAN minimum 'uint' value of zero.\n" +
					fmt.Sprintf("bNum.bigInt= '%v'\n", bNumBigIntNumStr),
			}
	}

	// Maximum Uint Value is 4,294,967,295
	if bNum.bigInt.Cmp(bIMaxUint) == 1 {

		return math.MaxUint32,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if bNum.bigInt.Cmp(bIMaxUint) == 1",
				ErrMessage: "bNum is GREATER THAN maximum 'uint' value.\n" +
					fmt.Sprintf("bNum.bigInt= '%v'\n"+
						"Max 'Uint32' Value= '%v'", bNumBigIntNumStr, bIMaxUintNumStr),
			}
	}

	return uint(bNum.bigInt.Uint64()), nil
}

// bigIntNumGetUInt64
//
// Returns the integer value of BigIntNum.bigInt as a  64-bit
// unsigned integer.
//
// BigIntNum is passed as input parameter 'bNum'.
//
// If the value of BigIntNum.bigInt exceeds that of the maximum
// unsigned 64-bit integer value, an error is returned.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bIntNumProton *bigIntNumProton) bigIntNumGetUInt64(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (uint64, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.bigIntNumGetSciNotationNumber",
		"")

	if err != nil {
		return 0, err
	}

	if bNum == nil {

		return 0,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	bIntMaxUint64 :=
		big.NewInt(0).SetUint64(uint64(math.MaxUint64))

	mathMaxUint64Str := strconv.FormatUint(math.MaxUint64, 10)

	if bNum.bigInt.Cmp(bIntMaxUint64) == 1 {
		return uint64(0),
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "if bNum.bigInt.Cmp(bIntMaxUint64) == 1",
				ErrMessage: fmt.Sprintf("The value of this BigIntNum instance exceeds\n"+
					"the maximum value of the unsigned 64-bit integer.\n"+
					"BigIntNum='%v'\nMaxUint64='%v'",
					bNum.bigInt.Text(10), mathMaxUint64Str),
			}
	}

	return bNum.bigInt.Uint64(), nil
}

// decimalGetDecimal -
// Converts the input parameter 'bNum' (type BigIntNum) value returns that value
// as a Type Decimal instance. The resulting number value includes the decimal
// place and decimal digits if they exist.
//
// The returned Decimal instance contains numeric separators (decimal separator,
// thousands separator and currency symbol) copied from the current BigIntNum
// instance, 'bNum'.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must do this!
func (bIntNumProton *bigIntNumProton) decimalGetDecimal(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (Decimal, error) {

	if bIntNumProton.lock == nil {
		bIntNumProton.lock = new(sync.Mutex)
	}

	bIntNumProton.lock.Lock()

	defer bIntNumProton.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumProton.decimalGetDecimal",
		"")

	if err != nil {
		return Decimal{}, err
	}

	if bNum == nil {

		return Decimal{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	dec, err := new(Decimal).NewBigInt(big.NewInt(0).
		Set(bNum.bigInt), bNum.precision)

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " dec, err := new(Decimal).NewBigInt(\n" +
					"  big.NewInt(0).Set(bNum.bigInt), bNum.precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	numSeps, err := new(bigIntNumAtom).getNumericSeparatorsDto(
		bNum,
		ePrefix)

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: " numSeps, err := new(bigIntNumAtom).\n" +
					"  getNumericSeparatorsDto(bNum, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	err = dec.SetNumericSeparatorsDto(numSeps)

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "dec, err := new(Decimal).NewBigInt(\n" +
					"  big.NewInt(0).Set(bNum.bigInt), bNum.precision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	err = dec.IsValid(ePrefix.String() + "dec INVALID! ")

	if err != nil {

		return Decimal{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = dec.IsValid(ePrefix + \"dec INVALID!\")",
				ErrContext: "",
				ErrMessage: err.Error(),
			}

	}

	return dec, nil
}
