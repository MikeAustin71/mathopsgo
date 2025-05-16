package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type bigIntFixedDecUtility struct {
	lock *sync.Mutex
}

// ceiling
//
// Returns the ceiling integer value for the BigIntFixedDecimal
// instance passed as input parameter 'bigIFxDec'.
//
// Ceiling is defined as: The least, or lowest value integer, which
// is greater than or equal to the numeric value of the current
// BigIntFixedDecimal.
//
//	Reference Wikipedia
//	===================
//
//	  https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//	Examples
//	========
//
//	  Initial    Ceiling
//	  Value       Value
//	  -------    -------
//	   5.95         6
//	   5.05         6
//	   5            5
//	  -5.05        -5
//	   2.4          3
//	   2.9          3
//	  -2.7         -2
//	  -2           -2
//
//	NOTE
//	====
//
//	This method does NOT test the validity of 'bNum', an
//	instance of type BigIntNum. The calling method must
//	do this!
func (bigIFdUtil *bigIntFixedDecUtility) ceiling(
	bigIFxDec *BigIntFixedDecimal,
	errPrefDto *ePref.ErrPrefixDto) (BigIntFixedDecimal, error) {

	if bigIFdUtil.lock == nil {
		bigIFdUtil.lock = new(sync.Mutex)
	}

	bigIFdUtil.lock.Lock()

	defer bigIFdUtil.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntFixedDecUtility.ceiling",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	if bigIFxDec == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec'",
		}

		return BigIntFixedDecimal{}, err
	}

	if bigIFxDec.integerNum == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDec.integerNum'",
		}

		return BigIntFixedDecimal{}, err
	}

	numSepsDto := NumericSeparatorDto{
		DecimalSeparator:   bigIFxDec.decimalSeparator,
		ThousandsSeparator: bigIFxDec.thousandsSeparator,
		CurrencySymbol:     bigIFxDec.currencySymbol,
	}

	numSepsDto.SetDefaultsIfEmpty()

	bigIFd2 := BigIntFixedDecimal{}

	bIFdAtom := new(bigIntFixedDecAtom)

	cmpZeroResult := bigIFxDec.integerNum.Cmp(big.NewInt(0))

	if cmpZeroResult == 0 {

		err = bIFdAtom.setNumericValue(
			&bigIFd2,
			big.NewInt(0),
			0,
			numSepsDto,
			ePrefix.XCpy("Setting 'bigIFd2' to Zero"))

		return bigIFd2, err
	}

	ceiling := big.NewInt(0).Set(bigIFxDec.integerNum)

	if bigIFxDec.precision > 0 {

		scale := big.NewInt(0).Exp(
			big.NewInt(10),
			big.NewInt(int64(bigIFxDec.precision)),
			nil)

		ceiling.Quo(ceiling, scale)

		if cmpZeroResult == 1 {
			// signVal must be plus
			ceiling.Add(ceiling, big.NewInt(1))
		}

	}

	err = bIFdAtom.setNumericValue(
		&bigIFd2,
		ceiling,
		0,
		numSepsDto,
		ePrefix.XCpy("Setting 'bigIFd2' to 'ceiling'"))

	return bigIFd2, err
}
