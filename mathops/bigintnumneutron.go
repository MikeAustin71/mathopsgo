package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntNumNeutron struct {
	lock *sync.Mutex
}

// inverse
//
// Returns the inverse of the current BigIntNum value.
// The inverse of the value is equal to one ('1') divided by the
// numeric value of the current BigIntNum.
//
// The BigIntNum return value for this operation will contain numeric
// separators (decimal separator, thousands separator and currency
// symbol) copied from the original BigIntNum instance.
func (bNumNeutron *bigIntNumNeutron) inverse(
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
		"bigIntNumNeutron.Inverse",
		"")

	if err != nil {
		return BigIntNum{}, err
	}

	if bNum == nil {

		return BigIntNum{},
			&ErrorReturnBasic{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "FATAL ERROR: Input parameter 'bNum' is a nil pointer.",
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
			&ErrorReturnBasic{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "  inverse, err := BigIntMathDivide{}.\n" +
					"    BigIntNumFracQuotient(bIOne, bITwo, maxPrecision)",
				ErrMessage: err.Error(),
			}
	}

	return inverse, nil
}

// isEvenNumber
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
func (bNumNeutron *bigIntNumNeutron) isEvenNumber(
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
		"bigIntNumNeutron.Inverse",
		"")

	if err != nil {
		return false, err
	}

	if bNum == nil {

		return false,
			&ErrorReturnBasic{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "",
				ErrMessage: "FATAL ERROR: Input parameter 'bNum' is a nil pointer.",
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
			&ErrorReturnBasic{
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
