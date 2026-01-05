package mathops

import (
	"math/big"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
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
//	This method does NOT test the validity of 'bigIFxDec', an
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

// copyIn
//
// Receives two BigIntFixedDecimal types via input parameters
// 'bigIFxDecDest' and 'bigIFxDecSrc'. This method will copy all
// values from the source BigIntFixedDecimal ('bigIFxDecSrc') to
// the destination BigIntFixedDecimal ('bigIFxDecDest')the values
// to the current BigIntFixedDecimal instance.
//
// If either 'bigIFxDecSrc' or 'bigIFxDecDest' fail the standard
// validation test, an error will be returned.
func (bigIFdUtil *bigIntFixedDecUtility) copyIn(
	bigIFxDecDest *BigIntFixedDecimal,
	bigIFxDecSrc *BigIntFixedDecimal,
	errPrefDto *ePref.ErrPrefixDto) error {

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
		"bigIntFixedDecUtility.copyIn",
		"")

	if err != nil {
		return err
	}

	if bigIFxDecDest == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDecDest'",
		}

		return err
	}

	if bigIFxDecSrc == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDecDest'",
		}

		return err
	}

	err = new(bigIntFixedDecAtom).isBigIntFxDecValid(
		bigIFxDecSrc,
		ePrefix.XCpy("Testing Validity of 'bigIFxDecSrc'"))

	if err != nil {
		return err
	}

	numSepsDto := NumericSeparatorDto{
		DecimalSeparator:   bigIFxDecSrc.decimalSeparator,
		ThousandsSeparator: bigIFxDecSrc.thousandsSeparator,
		CurrencySymbol:     bigIFxDecSrc.currencySymbol,
	}

	intVal, err := bigIFxDecSrc.GetIntegerValue()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "intVal, err := fd.GetIntegerValue()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = new(bigIntFixedDecAtom).setNumericValue(
		bigIFxDecDest,
		big.NewInt(0),
		0,
		numSepsDto,
		ePrefix.XCpy("Setting 'bigIFxDecDest' to zero"))

	if err != nil {
		return err
	}

	bigIFxDecDest.integerNum = big.NewInt(0).Set(intVal)

	bigIFxDecDest.precision, err = bigIFxDecSrc.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "bigIFxDecDest.precision, err = \n" +
				"  bigIFxDecSrc.GetPrecisionUint()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// copyOut
//
// Returns a new BigIntFixedDecimal instance which is
// a deep copy of the current BigIntFixedDecimal instance.
//
// If the current instance of BigIntFixedDecimal is determined
// to be invalid, an error will be returned.
func (bigIFdUtil *bigIntFixedDecUtility) copyOut(
	bigIFxDecSrc *BigIntFixedDecimal,
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
		"bigIntFixedDecUtility.copyIn",
		"")

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	if bigIFxDecSrc == nil {

		err = &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'bigIFxDecSrc'",
		}

		return BigIntFixedDecimal{}, err
	}

	err = new(bigIntFixedDecAtom).isBigIntFxDecValid(
		bigIFxDecSrc,
		ePrefix.XCpy("Testing Validity of 'bigIFxDecSrc'"))

	if err != nil {
		return BigIntFixedDecimal{}, err
	}

	numSepsDto := NumericSeparatorDto{
		DecimalSeparator:   bigIFxDecSrc.decimalSeparator,
		ThousandsSeparator: bigIFxDecSrc.thousandsSeparator,
		CurrencySymbol:     bigIFxDecSrc.currencySymbol,
	}

	bigIFxDecDest := BigIntFixedDecimal{}

	err = new(bigIntFixedDecAtom).setNumericValue(
		&bigIFxDecDest,
		bigIFxDecSrc.integerNum,
		bigIFxDecSrc.precision,
		numSepsDto,
		ePrefix.XCpy("Setting 'bigIFxDecDest' = 'bigIFxDecSrc'"))

	return bigIFxDecDest, err
}
