package mathops

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryNanobot struct {
	lock *sync.Mutex
}

// setInternalFlags - Sets Array Lengths and
// test for zero values
//
//	IMPORTANT
//	=========
//
//	The calling function is responsible for verifying the validit
//	of 'ia', the Intary object.
func (iaNanobot *intAryNanobot) setInternalFlags(
	ia *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

	if iaNanobot.lock == nil {
		iaNanobot.lock = new(sync.Mutex)
	}

	iaNanobot.lock.Lock()

	defer iaNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNanobot.setInternalFlags()",
		"")

	if err != nil {
		return err
	}

	if ia == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ParameterName: "'ia'",
		}
	}

	err = new(intAryElectron).setSignificantDigitIdxs(
		ia,
		ePrefix)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "err = new(intAryElectron).\n" +
				"  setSignificantDigitIdxs( ia, ePrefix)",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// getNumericSeparatorsDto
//
//	 Receives a pointer to IntAry and extracts the Numeric
//	 Separators. These separators are consolidated and returned as
//	 a NumericSeparatorDto structure containing the character or
//	 'rune' values for decimal point separator, thousands
//	 separator and currency symbol.
//
//		Numeric Separators
//		==================
//
//		Numeric Separators define the Decimal Separator character,
//		Thousands Separator character, and Currency Symbol character.
//		These separator characters serve two purposes. First they are
//		used to format and display numeric values as number strings.
//		Second, they are also used to parse number strings and convert
//		them into numeric values.
func (iaNanobot *intAryNanobot) getNumericSeparatorsDto(
	intAry *IntAry,
	errPrefDto *ePref.ErrPrefixDto) (NumericSeparatorDto, error) {

	if iaNanobot.lock == nil {
		iaNanobot.lock = new(sync.Mutex)
	}

	iaNanobot.lock.Lock()

	defer iaNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNanobot.getNumericSeparatorsDto()",
		"")

	if err != nil {
		return NumericSeparatorDto{}, err
	}

	if intAry == nil {

		return NumericSeparatorDto{},
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'intAry'",
			}
	}

	numSeps := NumericSeparatorDto{}
	numSeps.DecimalSeparator = intAry.GetDecimalSeparator()
	numSeps.ThousandsSeparator = intAry.GetThousandsSeparator()
	numSeps.CurrencySymbol = intAry.GetCurrencySymbol()

	return numSeps, nil
}
