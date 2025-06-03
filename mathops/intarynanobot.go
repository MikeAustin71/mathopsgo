package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type intAryNanobot struct {
	lock sync.Mutex
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

// hasFractionalDigits
//
//	This method examines the current intAry object to determine if
//	there are non-zero digits to the right of the decimal place. If
//	all digits to the right of the decimal place are zero, this
//	method returns 'false'
//
//	If non-zero digits are present to the right of the decimal
//	place, the method returns 'true'.
func (iaNanobot *intAryNanobot) hasFractionalDigits(
	intAry *IntAry,
	validateIntAry bool,
	errPrefDto *ePref.ErrPrefixDto) (bool, error) {

	iaNanobot.lock.Lock()

	defer iaNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"intAryNanobot.hasFractionalDigits",
		"")

	if err != nil {
		return false, err
	}

	if validateIntAry {
		err = new(intAryElectron).isValidIntAry(
			intAry,
			ePrefix.XCpy("Validating 'intAry'").String())

		if err != nil {

			return false,
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "err = new(intAryElectron).isValidIntAry(\n" +
						"  intAry, ePrefix.XCpy(Validating 'intAry').String())",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}
	}

	if intAry.precision == 0 {
		return false, nil
	}

	err = new(intAryElectron).setIntAryLength(
		intAry, ePrefix.XCpy("Setting 'intAry' IntAry Length"))

	if err != nil {

		return false,
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = new(intAryElectron).setIntAryLength(\n" +
					"intAry, ePrefix.XCpy(Setting 'intAry' IntAry Length))",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	intLen := intAry.intAryLen - intAry.precision

	if intLen < 1 {
		return false,
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "intLen < 1",
				ErrMessage: fmt.Sprintf("Error - Int Array integer length is less than 1.\n"+
					"intLen= '%v'", intLen),
			}
	}

	for i := intLen; i < intAry.intAryLen; i++ {

		if intAry.intAry[i] > 0 {

			return true, nil
		}
	}

	return false, nil
}

// setInternalFlags
//
//	 Sets Array Lengths and test for zero values
//
//		IMPORTANT
//		=========
//
//		The calling function is responsible for verifying the validity
//		of 'ia', the IntAry object.
func (iaNanobot *intAryNanobot) setInternalFlags(
	ia *IntAry,
	errPrefDto *ePref.ErrPrefixDto) error {

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

// setIsZeroValue
//
//	Analyzes the value of the intAry and sets a flag if the value
//	of intAry evaluates to zero.
func (iaNanobot *intAryNanobot) setIsZeroValue(
	ia *IntAry) {

	iaNanobot.lock.Lock()

	defer iaNanobot.lock.Unlock()

	if ia == nil {
		return
	}

	ia.intAryLen = len(ia.intAry)

	ia.isZeroValue = true

	ia.isIntegerZeroValue = true

	intLen := ia.intAryLen - ia.precision

	for i := 0; i < ia.intAryLen; i++ {

		if i < intLen && ia.intAry[i] > 0 {

			ia.isIntegerZeroValue = false
		}

		if ia.intAry[i] > 0 {

			ia.isZeroValue = false

			return
		}
	}

	// ia.isZeroValue == true
	// signVal must be 1
	ia.signVal = 1

	return
}
