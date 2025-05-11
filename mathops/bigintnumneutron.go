package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

type bigIntNumNeutron struct {
	lock *sync.Mutex
}

// inverseBigIntNum
//
// Returns the inverseBigIntNum of the current BigIntNum value.
// The inverseBigIntNum of the value is equal to one ('1') divided by the
// numeric value of the current BigIntNum.
//
// The BigIntNum return value for this operation will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) inverseBigIntNum(
	bNum *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.inverseBigIntNum",
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

	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		return BigIntNum{}, nil
	}

	bIOne, err := new(bigIntNumMolecule).newOne(0, ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	err = new(bigIntNumAtom).
		setNumericSeparators(&bIOne, bNum.decimalSeparator,
			bNum.thousandsSeparator, bNum.currencySymbol, ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	bITwo, err := new(bigIntNumUtility).bigIntNumCopyOut(bNum, ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	inverse, err := BigIntMathDivide{}.
		BigIntNumFracQuotient(bIOne, bITwo, maxPrecision)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "inverseBigIntNum, err := BigIntMathDivide{}.\n" +
					"    BigIntNumFracQuotient(bIOne, bITwo, maxPrecision)",
				ErrMessage: err.Error(),
			}
	}

	return inverse, nil
}

// isEvenBigIntNumber
//
// Returns true if the current BigIntNum value is evenly
// divisible by 2.
//
// Even Number Definitions:
//
//	https://www.mathsisfun.com/definitions/even-number.html
//
// "In mathematics, parity is the property of an
// integer's inclusion in one of two categories:
// even or odd. An integer is even if it is evenly
// divisible by two and odd if it is not even."
//
// "Examples of even numbers include −4, 0, 82 and 178."
// In particular, zero is an even number."
//
// https://en.wikipedia.org/wiki/Parity_(mathematics)
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) isEvenBigIntNumber(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.isEvenBigIntNumber",
		"")

	if err != nil {
		return false, err
	}

	if bNum == nil {

		return false,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}
	}

	if bNum.precision > 0 {
		return false, nil
	}

	// Is bNum Zero?
	if bNum.bigInt.Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return false, err
	}

	_, mod, err := BigIntMathDivide{}.
		BigIntNumDivideByTwoQuoMod(bNum2, 50)

	if err != nil {
		return false,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "  _, mod, err := BigIntMathDivide{}.\n" +
					"BigIntNumDivideByTwoQuoMod(bNum2, 50)",
				ErrMessage: err.Error(),
			}
	}

	if mod.bigInt.Cmp(big.NewInt(0)) == 0 {
		return true, nil
	}

	return false, nil
}

// incrementBigIntNum
//
// Adds a value of +1 (plus one) to the numeric value of the
// BigIntNum instance passed as input parameter 'bNum'
//
// The numeric separators (decimal separator, thousands separator
// and currency symbol) from the original BigIntNum will remain
// unchanged.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) incrementBigIntNum(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.incrementBigIntNum",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}

	}

	biNumOne, err := new(bigIntNumMolecule).newOne(
		bNum.precision,
		ePrefix)

	if err != nil {

		return err
	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return err
	}

	bPair, err := new(BigIntPair).NewBigIntNum(bNum2, biNumOne)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "bPair, err := new(BigIntPair).NewBigIntNum(bNum2, biNumOne)",
			ErrMessage: err.Error(),
		}

	}

	result, err := new(BigIntMathAdd).AddPair(bPair)

	if err != nil {

		return fmt.Errorf("%v\n"+
			"Error returned by: \n"+
			" result, err := new(BigIntMathAdd).AddPair(bPair)\n"+
			"Error= %v\n",
			ePrefix,
			err.Error())
	}

	err = new(bigIntNumUtility).bigIntNumCopyIn(
		bNum,
		&result,
		ePrefix)

	if err != nil {

		return err
	}

	return nil
}

// modBigIntNum
//
// Performs a modulo operation where the current BigIntNum numeric value is the
// dividend and the divisor is the input parameter, 'divisor'.  The modulo operation finds
// the remainder after division of one number by another (sometimes called modulus).
// (Wikipedia: https://en.wikipedia.org/wiki/Modulo_operation)
//
//		 									dividend = bNum
//	  									dividend % divisor = modulo
//
// The result of this modulo operation is returned as a BigIntNum, 'modulo'. 'modulo' may
// consist of an integer or a floating point value consisting of integer and fractional
// digits.
//
// Input parameter 'maxPrecision' is used to control the maximum precision of the resulting
// floating point 'modulo'. Precision is defined as the number of fractional digits to
// the right of the decimal place. Be advised that these calculations can support very large
// precision values.
//
// The returned BigIntNum instance, 'modulo', will contain numeric separators (decimal
// separator, thousands separator and currency symbol) copied from the current BigIntNum
// instance (bNum).
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
//
// However, this method will vallidate input parameter
// 'divisor'.
func (bNumNeutron *bigIntNumNeutron) modBigIntNum(
	bNum *BigIntNum,
	divisor *BigIntNum,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (modulo BigIntNum, err error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.modBigIntNum",
		"")

	if err != nil {
		return modulo, err
	}

	if bNum == nil {

		return modulo,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}

	}

	if divisor == nil {

		return modulo,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'divisor'",
			}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		divisor,
		ePrefix.XCpy(" Validating 'divisor'"))

	if err != nil {
		return modulo, err
	}

	biNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum, ePrefix)

	if err != nil {

		return modulo, err
	}

	modulo, err = BigIntMathDivide{}.BigIntNumModulo(biNum2, *divisor, maxPrecision)

	if err != nil {

		return modulo,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "modulo, err = BigIntMathDivide{}.BigIntNumModulo(\n" +
					"    biNum2, *divisor, maxPrecision)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return modulo, err
}

// multiplyBigIntNum
//
// Multiplies the numerical value of the current BigIntNum instance
// ('multiplier') times input parameter 'multiplicand'. The 'product' of this
// multiplication operation is returned as a BigIntNum.
//
//	multiplier = bNum
//	multiplier X multiplicand = product
//
// The BigIntNum instance returned by this method, 'product', will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from the current BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
//
// However, this method will vallidate input parameter
// 'multiplicand'.
func (bNumNeutron *bigIntNumNeutron) multiplyBigIntNum(
	bNum *BigIntNum,
	multiplicand *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (product BigIntNum, err error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.multiplyBigIntNum",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	product = new(bigIntNumMechanics).new()

	if bNum == nil {

		return product,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bNum'",
			}

	}

	if multiplicand == nil {

		return product,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'multiplicand'",
			}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		multiplicand,
		ePrefix.XCpy(" Validating 'multiplicand'"))

	if err != nil {
		return product, err
	}

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return product, err
	}

	product, err = new(BigIntMathMultiply).MultiplyBigIntNums(bNum2, *multiplicand)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "product, err = new(BigIntMathMultiply).MultiplyBigIntNums(\n" +
					"    bNum2, *multiplicand)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return product, err
}

// multiplyByFiveBigIntNum
//
// Multiplies the numerical value of the current BigIntNum instance
// times five (5). The product is returned as a BigIntNum.
//
//	product = bNum X 5
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) multiplyByFiveBigIntNum(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.multiplyByFiveBigIntNum",
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

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	mulResult, err := new(BigIntMathMultiply).MultiplyBigIntNumByFive(bNum2)

	if err != nil {

		return BigIntNum{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "mulResult, err := new(BigIntMathMultiply).MultiplyBigIntNumByFive(bNum2)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return mulResult, nil
}

// multiplyByTenBigIntNum
//
// Multiplies the numerical value of the BigIntNum instance
// passed as input parameter 'bNum' by ten (10). The product
// is returned as a BigIntNum.
//
//	product = bNum X 10
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) multiplyByTenBigIntNum(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.multiplyByTenBigIntNum",
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

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	mulResult, err := new(BigIntMathMultiply).MultiplyBigIntNumByTen(bNum2)

	if err != nil {

		return BigIntNum{}, &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "mulResult, err := new(BigIntMathMultiply).MultiplyBigIntNumByFive(bNum2)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return mulResult, nil
}

// multiplyByTenToPowerBigIntNum
//
// Multiplies the numerical value of the current BigIntNum/ instance times
// ten to the power of 'exponent' (10^exponent). The product is returned
// as the new value for the current BigIntNum. The original value of the
// BigIntNum instance will be overwritten and destroyed.
//
//	bNum = bNum X 10^exponent
//
// The BigIntNum instance generated by this method will contain numeric
// separators (decimal separator, thousands separator and currency symbol)
// copied from the original BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) multiplyByTenToPowerBigIntNum(
	bNum *BigIntNum,
	exponent uint,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.multiplyByTenToPowerBigIntNum",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	if bNum.precision >= exponent {

		netPrecision := bNum.precision - exponent

		bInt3, err := new(bigIntNumMechanics).newBigInt(
			bNum.bigInt, netPrecision, ePrefix)

		if err != nil {

			return err
		}

		err = new(bigIntNumUtility).bigIntNumCopyIn(
			bNum,
			&bInt3,
			ePrefix)

		if err != nil {

			return err
		}

	} else {
		// exponent > bNum.precision

		netPrecision := int64(exponent - bNum.precision)

		scaleVal :=
			big.NewInt(0).Exp(big.NewInt(10), big.NewInt(netPrecision), nil)

		newVal := big.NewInt(0).Mul(bNum.bigInt, scaleVal)

		bInt4, err := new(bigIntNumMechanics).newBigInt(
			newVal, 0, ePrefix)

		if err != nil {

			return err

		}

		err = new(bigIntNumUtility).bigIntNumCopyIn(
			bNum,
			&bInt4,
			ePrefix)

		if err != nil {

			return err
		}

	}

	return nil
}

// multiplyByTenToPowerAddBigIntNum
//
// Performs three operations on the BigIntNum instance passed as
// input parameter 'bNum'.
//
// (1) 	First, the method multiplies the numerical value of the current BigIntNum
//
//	instance times ten to the power of 'exponent' (10^exponent).
//
//						bNum1 = bNum X 10^exponent
//
// (2)  Second, the method adds input parameter 'addend' to the product generated
//
//	by operation (1), above.
//
//						bNum2 = bNum1 + 'addend'
//
// (3)  Third and finally, the original value of the current BigIntNum instance
//
//	will be overwritten and replaced by the 'bNum2' value generated in
//	operation (2), above.
//
// The BigIntNum instance generated by this method will contain numeric separators
// (decimal separator, thousands separator and currency symbol) copied from the
// original BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
//
// However, this method WILL test and validate the input
// parameter 'addend'.
func (bNumNeutron *bigIntNumNeutron) multiplyByTenToPowerAddBigIntNum(
	bNum *BigIntNum,
	exponent uint,
	addend *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.multiplyByTenToPowerAddBigIntNum",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	err = new(bigIntNumAtom).isBigIntNumValid(
		addend,
		ePrefix.XCpy("Validating input parameter 'addend'"))

	if err != nil {
		return err
	}

	var bx BigIntNum
	var bigINumMech = new(bigIntNumMechanics)

	if bNum.precision >= exponent {

		netPrecision := bNum.precision - exponent

		bx, err = bigINumMech.newBigInt(
			bNum.bigInt, netPrecision, ePrefix)

		if err != nil {

			return err
		}

	} else {
		// exponent > bNum.precision

		netPrecision := int64(exponent - bNum.precision)

		scaleVal :=
			big.NewInt(0).Exp(
				big.NewInt(10),
				big.NewInt(netPrecision),
				nil)

		newVal := big.NewInt(0).Mul(bNum.bigInt, scaleVal)

		bx, err = bigINumMech.newBigInt(
			newVal, 0, ePrefix)

		if err != nil {

			return err
		}

	}

	result, err := new(BigIntMathAdd).AddBigIntNums(bx, *addend)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "result, err := new(BigIntMathAdd).AddBigIntNums(bx, addend)",
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

// multiplyByThreeBigIntNum
//
// Multiplies the numerical value of the current BigIntNum
// instance times three (3). The product is returned as a BigIntNum.
//
//	product = bNum X 3
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) multiplyByThreeBigIntNum(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.multiplyByThreeBigIntNum",
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

	bNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	result, err := new(BigIntMathMultiply).
		MultiplyBigIntNumByThree(bNum2)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "result, err := new(BigIntMathMultiply).MultiplyBigIntNumByThree(bNum2)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}

// multiplyByTwoBigIntNum
//
// Multiplies the numerical value of the BigIntNum instance passed
// as input parameter 'bNum', times two (2). The product is
// returned as a BigIntNum.
//
//	product = bNum X 2
//
// The BigIntNum instance returned by this method will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
//
//	NOTE:
//
// This method does NOT test the validity of 'bNum', an
// instance of type BigIntNum. The calling method must
// do this!
func (bNumNeutron *bigIntNumNeutron) multiplyByTwoBigIntNum(
	bNum *BigIntNum,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.multiplyByTenToPowerAddBigIntNum",
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

	biNum2, err := new(bigIntNumUtility).bigIntNumCopyOut(
		bNum,
		ePrefix)

	if err != nil {

		return BigIntNum{}, err

	}

	result, err := new(BigIntMathMultiply).MultiplyBigIntNumByTwo(
		biNum2)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "result, err := new(BigIntMathMultiply).MultiplyBigIntNumByThree(biNum2)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return result, nil
}

// newBigIntNumWithPrecision
//
// Creates a new BigIntNum instance using a *big.Int type and its
// associated precision (also of type *big.Int).
//
// The 'precision' parameter specifies the number of digits to the right
// of the decimal place. The Numeric value is equal to bigI x 10^(precision x -1).
// This effectively locates the decimal place by counting from the extreme right
// of the integer number, 'precision' places to the left. See the example below.
//
//	Precision Example:
//	==================
//
//			Integer Value		precision			Numeric Value
//			  123456					 3					  123.456
//
// Input Parameters
// ================
//
// bigI 			*big.Int
//
//	'bigI' is a type *big.Int and represents the integer
//	value of the number; that is, the numeric value without decimal digits.
//
// precision  *big.Int
//
//	This integer value (always a positive value) identifies
//	the location of the decimal place in the integer value 'bigI'.
//	The decimal place location is calculated by starting with the
//	right most digit in the integer number and counting	left,
//	'precision' places. If precision is greater than the maximum
//	value of an unsigned integer (+4,294,967,295,	which equals
//	2^32 − 1), an error will be triggered. Also, if the 'precision'
//	value is less than zero, an error will be triggered.
//
// Return Parameters
// =================
//
//	BigIntNum - a type BigIntNum numeric value
//
//	error			- If not 'nil', this prameter will
//							transmit any processing errors
//							encountered.
//
//
//		The new BigIntNum instance returned by this method will contain USA default
//		numeric separators (decimal separator, thousands separator and currency
//		symbol). To reconfigure the numeric separators reference method:
//							BigIntNum.SetNumericSeparators()
func (bNumNeutron *bigIntNumNeutron) newBigIntNumWithPrecision(
	bigInt *big.Int,
	precision *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.newBigIntNumWithPrecision",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bigInt == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigInt'",
			}
	}

	if precision == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'precision'",
			}
	}

	if precision.Cmp(big.NewInt(0)) == -1 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("precistion= '%v'", precision.Text(10)),
				ErrMessage: "Error: Input parameter 'precision' IS LESS THAN ZERO!",
			}
	}

	maxUint32 := big.NewInt(0).SetUint64(uint64(math.MaxUint32))

	if precision.Cmp(maxUint32) == 1 {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: fmt.Sprintf("precistion= '%v' math.MaxUint32= '%v'", precision.Text(10), math.MaxUint32),
				ErrMessage: fmt.Sprintf("Error: Input parameter 'precision' exceeds maximum limit of '%v' !", math.MaxUint32),
			}
	}

	bIntNum, err := new(bigIntNumMechanics).newZero(0, ePrefix)

	if err != nil {
		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bIntNum, err := new(bigIntNumMechanics).newZero(0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		&bIntNum, ePrefix)

	if err != nil {

		return BigIntNum{},
		&FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(\n" +
				"    bNum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumNanobot).setBigInt(
		&bIntNum,
		bigInt,
		uint(precision.Uint64()),
		ePrefix)

	if err != nil {

		return BigIntNum{}, err
	}

	return bIntNum, nil
}

// newBigIntExponent
//
// New bigInt Exponent returns a new BigIntNum instance
// in which the numeric value is set using an integer
// multiplied by 10 raised to the power of the 'exponent'
// parameter.
//
//	numeric value = integer X 10^exponent
//
//						OR
//
//	BigIntNum (return value) = bigI X 10^exponent
//
// If exponent is less than +1, precision is set equal to
// exponent and bigI is unchanged.
//
// If exponent is greater than 0, bigI is multiplied by 10
// raised to the power of 'exponent', and precision is set
// equal to zero.
//
// Examples:
//
//	biNum :=
//			new(BigIntNum).
//				NewBigIntExponent(big.NewInt(int64(123456)), -3) =
//								"123.456"  precision = 3
//
//	biNum :=
//			BigIntNum{}.NewBigIntExponent(big.NewInt(int64(123456)), 3) = "123456.000" precision = 3
func (bNumNeutron *bigIntNumNeutron) newBigIntExponent(
	bigI *big.Int,
	exponent int,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.newBigIntExponent",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bigI == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigI' (type *big.Int)",
			}
	}

	bINum, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "b, err := new(bigIntNumMechanics).newZero(0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(
		&bINum, ePrefix)

	if err != nil {

		return BigIntNum{},
		 &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumAtom).setNumericSeparatorsToDefaultIfEmpty(\n" +
				"    bINum, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntNumMolecule).
		setBigIntExponent(&bINum, bigI, exponent, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumMolecule).\n" +
					"    setBigIntExponent(&bINum, bigI, exponent, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bINum, nil
}

// newBigFloat
//
// Returns a new BigIntNum instance using a *big.Float floating
// point input parameter.  The precision of the resulting BigIint
// numeric value is specified by the input parameter,
// 'maxPrecision'.
//
// Input Parameters
// ================
//
// bigFloat *big.Float
//
//	This *big.Float value will be converted into an instance
//	of BigIntNum.
//
// maxPrecision uint
//
//	The maximum precision for the resulting BigIntNum after
//	conversion of input parameter 'bigFloat'. Resulting precision
//	will never be greater than 'maxPrecision'; however, actual
//	precision may be less than 'maxPrecision'.
func (bNumNeutron *bigIntNumNeutron) newBigFloat(
	bNum *BigIntNum,
	bigFloat *big.Float,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.newBigFloat",
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

	if bigFloat == nil {

		return BigIntNum{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'bigFloat'",
			}
	}

	b, err := new(bigIntNumMechanics).newZero(
		0,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "b, err := new(bigIntNumMechanics).newZero(0, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	err = new(bigIntNumNanobot).setBigFloat(
		&b,
		bigFloat,
		maxPrecision,
		ePrefix.XCpy("bigFloat -> 'b' BigIntNum"))

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumNanobot).setBigFloat(b, bigFloat, maxPrecision,\n" +
					"    ePrefix.XCpy(\"bigFloat -> 'b' BigIntNum\"))",
				ErrContext: fmt.Sprintf("bigFloat= '%v' maxPrecision= '%v' ", bigFloat.String(), maxPrecision),
				ErrMessage: err.Error(),
			}
	}

	return b, nil
}

// newDecimal
//
// Receives an input parameter 'decNum' of type Decimal and
// returns a BigIntNum instance configured with the numeric
// value passed by parameter 'decNum'.
//
// Input parameter 'decNum' will be subjected to validation
// testing. If 'decNum' fails these validation tests, an
// error will be returned.
func (bNumNeutron *bigIntNumNeutron) newDecimal(
	bNum *BigIntNum,
	decNum Decimal,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.newDecimal",
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

	validStr := ePrefix.XCpy("Validating 'decNum'").String()

	err = decNum.IsValid(validStr)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = decNum.IsValid(validStr)",
				ErrContext: "Validation on input parameter 'decNum'",
				ErrMessage: err.Error(),
			}
	}

	decNumStr, err := decNum.GetNumStr()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "decNumStr, err := decNum.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	bInt, err := decNum.GetSignedBigInt()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "bInt, err := decNum.GetSignedBigInt()",
				ErrContext: fmt.Sprintf("decNumStr = '%v' ", decNumStr),
				ErrMessage: err.Error(),
			}
	}

	precision := uint(decNum.GetPrecision())

	b := new(bigIntNumMechanics).new()

	new(bigIntNumElectron).empty(&b)

	err = new(bigIntNumNanobot).setBigInt(
		&b,
		bInt,
		precision,
		ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = new(bigIntNumNanobot).setBigInt(&b,bInt,precision,ePrefix)",
				ErrContext: fmt.Sprintf("bInt = '%v' precision= '%v'",
					bInt.Text(10), precision),
				ErrMessage: err.Error(),
			}
	}

	return b, nil
}

// newBigIntFixedDecimal
//
// Creates and returns a new BigIntNum instance based
// on input parameter 'fd' of type BigIntFixedDecimal.
//
// The 'fd' numeric value will be converted to type
// a BigIntNum which is then returned by this method.
//
// If input parameter 'fd' proves invalid, an error
// will be returned.
func (bNumNeutron *bigIntNumNeutron) newBigIntFixedDecimal(
	fd BigIntFixedDecimal,
	errPrefDto *ePref.ErrPrefixDto) (BigIntNum, error) {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumNeutron.newBigIntFixedDecimal",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	err = fd.IsValid(" Testing input parameter 'fd'")

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "err = fd.IsValid(\" Testing input parameter 'fd'\")",
				ErrContext: "'fd' is input parameter of type BigIntFixedDecimal",
				ErrMessage: "Error: Input Parameter 'fd' is INVALID!.\n" +
					"'fd' Vallidity Test FAILED!",
			}
	}

	fdIntVal, err := fd.GetInteger()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fdIntVal, err := fd.GetInteger()",
				ErrContext: "'fd' is input parameter of type BigIntFixedDecimal",
				ErrMessage: err.Error(),
			}
	}

	fdPrecision, err := fd.GetPrecision()

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "fdPrecision, err := fd.GetPrecision()",
				ErrContext: "'fd' is input parameter of type BigIntFixedDecimal",
				ErrMessage: err.Error(),
			}
	}

	bid, err := new(bigIntNumMechanics).newBigInt(
		fdIntVal, fdPrecision, ePrefix)

	if err != nil {

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "bid, err := new(bigIntNumMechanics).newBigInt(\n" +
					"    fdIntVal, fdPrecision, ePrefix)",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	return bid, nil
}

// setIntFracStrings
//
// Sets the value of the current BigIntNum instance based on a numeric value
// represented by separate integer and fractional components.
//
// Input parameters 'intStr' and 'fracStr' are strings representing the integer and
// fractional components. They are combined by this method to create a numeric value
// which is assigned to the current BigIntNum instance.
//
// Input parameter 'signVal' must be set to one of two values: +1 or -1. This value is
// used to signal the sign of the resulting numeric value. +1 generates a positive number
// and -1 generates a negative number. If input parameters 'inStr' or 'fracStr' contain
// a leading minus or plus sign character, it will be ignored. The sign of the resulting
// numeric value is controlled strictly by input parameter, 'signVal'.
//
// Existing numeric separators (decimal separator, thousands separator
// and currency symbol) remain unchanged and are not altered by this method.
func (bNumNeutron *bigIntNumNeutron) setIntFracStrings(
	bNum *BigIntNum,
	intStr string,
	fracStr string,
	signVal int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bNumNeutron.lock == nil {
		bNumNeutron.lock = new(sync.Mutex)
	}

	bNumNeutron.lock.Lock()

	defer bNumNeutron.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntNumAtom.setIntFracStrings",
		"")

	if err != nil {
		return err
	}

	if bNum == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bNum'",
		}
	}

	cleanIntRuneAry := make([]rune, 0, 100)

	zeroChar := uint8('0')
	nineChar := uint8('9')

	lStr := len(intStr)

	if lStr == 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "lStr := len(intStr)",
			ErrContext: "if lStr == 0 {",
			ErrMessage: "Error: Input Parameter 'intStr' is zero Length!",
		}
	}

	isFirstRune := true

	// Create pure number string from 'intStr'
	for i := 0; i < lStr; i++ {

		if intStr[i] >= zeroChar &&
			intStr[i] <= nineChar {

			if isFirstRune && signVal == -1 {
				cleanIntRuneAry = append(cleanIntRuneAry, '-')
			}

			isFirstRune = false

			cleanIntRuneAry = append(cleanIntRuneAry, rune(intStr[i]))
		}
	}

	if len(cleanIntRuneAry) == 0 {
		cleanIntRuneAry = append(cleanIntRuneAry, '0')
	}

	lStr = len(fracStr)

	if lStr > 0 {

		isFirstRune = true

		for j := 0; j < lStr; j++ {

			if fracStr[j] >= zeroChar &&
				fracStr[j] <= nineChar {

				if isFirstRune {

					var nSepSymbol NumSepSymbolCode

					nSepSymbol = DECIMALSYMBOL

					decSeparator, err := new(bigIntNumProton).bigIntNumGetNumSepSymbol(
						bNum, nSepSymbol, ePrefix)

					if err != nil {

						return &FuncReturnError{
							ErrPrefix: ePrefix.String(),
							ReturnFunc: "decSeparator, err := new(bigIntNumProton).\n" +
								"    bigIntNumGetNumSepSymbol(bNum, nSepSymbol, ePrefix)",
							ErrContext: "",
							ErrMessage: err.Error(),
						}
					}

					cleanIntRuneAry = append(cleanIntRuneAry, decSeparator)

					isFirstRune = false
				}

				cleanIntRuneAry = append(cleanIntRuneAry, rune(fracStr[j]))
			}

		}
	}

	err = new(bigIntNumMolecule).setNumStr(
		bNum,
		string(cleanIntRuneAry),
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(bigIntNumMolecule).setNumStr(bNum, \n" +
				"    string(cleanIntRuneAry), ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}
