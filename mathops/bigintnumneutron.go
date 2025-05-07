package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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
