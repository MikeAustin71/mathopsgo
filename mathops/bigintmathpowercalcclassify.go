package mathops

import (
	"fmt"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathPowerCalcClassify struct {
	lock *sync.Mutex
}

// classifyExponentCalcType
//
//	This method classifies the type of exponentiation calculation
//	that is being performed. 16 different types of exponentiation
//	calculations are supported.
//
//	A pointer to an instance of type BigIntMathPowerProfile is
//	passed as an input parameter.
//
//	This BigIntMathPowerProfile instance is populated with the
//	necessary information to perform the exponentiation calculation
//	type classifications.
//
//	The method returns an error value if the input parameter
//	'powerCalcProfile' is 'nil'.
func (bIntMathPwrCalcClass *bigIntMathPowerCalcClassify) classifyExponentCalcType(
	powerCalcProfile *BigIntMathPowerProfile,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bIntMathPwrCalcClass.lock == nil {
		bIntMathPwrCalcClass.lock = new(sync.Mutex)
	}

	bIntMathPwrCalcClass.lock.Lock()

	defer bIntMathPwrCalcClass.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerCalcClassify.classifyExponentCalcType()",
		"")

	if err != nil {
		return err
	}

	if powerCalcProfile == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'bNum'",
		}
	}

	// Calc Type # 1
	if !powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoPlusInt

		return err
	}

	// Calc Type # 2
	if !powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoPlusFrac

		return err
	}

	// Calc Type # 3
	if !powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoMinusInt

		return err
	}

	// Calc Type # 4
	if !powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusIntExpoMinusFrac

		return err
	}

	// Calc Type # 5
	if !powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusFracExpoPlusInt

		return err
	}

	// Calc Type # 6
	if !powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusFracExpoPlusFrac

		return err
	}

	// Calc Type # 7
	if !powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusFracExpoMinusInt

		return err
	}

	// Calc Type # 8
	if !powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBasePlusFracExpoMinusFrac

		return err
	}

	// Calc Type # 9
	if powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusIntExpoPlusInt

		return err
	}

	// Calc Type # 10
	if powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusIntExpoPlusFrac

		return err
	}

	// Calc Type # 11
	if powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusIntExpoMinusInt

		return err
	}

	// Calc Type # 12
	if powerCalcProfile.BaseIsNegative &&
		powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusIntExpoMinusFrac

		return err
	}

	// Calc Type # 13
	if powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusFracExpoPlusInt

		return err
	}

	// Calc Type # 14
	if powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		!powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusFracExpoPlusFrac

		return err
	}

	// Calc Type # 15
	if powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusFracExpoMinusInt

		return err
	}

	// Calc Type # 16
	if powerCalcProfile.BaseIsNegative &&
		!powerCalcProfile.BaseIsInteger &&
		powerCalcProfile.ExponentIsNegative &&
		!powerCalcProfile.ExponentIsInteger {

		powerCalcProfile.ExpoCalcTypeCode = ExpoCalcBaseMinusFracExpoMinusFrac

		return err
	}

	err = fmt.Errorf("%v\n" +
		"Exponent Calculation Type is Invalid!\n")

	return err
}
