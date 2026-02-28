package mathops

import (
	"fmt"
	"sync"

	ePref "github.com/MikeAustin71/errpref"
)

type bigIntMathPowerCalcRouting struct {
	lock *sync.Mutex
}

func (bigIntPwrCalcRouting *bigIntMathPowerCalcRouting) routeExecutePowerCalc(
	base *BigIntNum,
	exponent *BigIntNum,
	calcProfile *BigIntMathPowerProfile,
	maxPrecision uint,
	errPrefDto *ePref.ErrPrefixDto) (result BigIntNum, err error) {

	if bigIntPwrCalcRouting.lock == nil {
		bigIntPwrCalcRouting.lock = new(sync.Mutex)
	}

	bigIntPwrCalcRouting.lock.Lock()

	defer bigIntPwrCalcRouting.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"bigIntMathPowerCalcRouting.routeExecutePowerCalc",
		"")

	if err != nil {
		return result, err
	}

	if base == nil {

		return result,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'base'",
			}
	}

	if exponent == nil {

		return result,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'exponent'",
			}
	}

	if calcProfile == nil {

		return result,
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ErrContext:    "",
				ParameterName: "'calcProfile'",
			}
	}

	switch calcProfile.ExpoCalcTypeCode {

	case ExpoCalcBasePlusIntExpoMinusInt:

	default:

		return BigIntNum{},
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "",
				ErrContext: "Input parameter 'calcProfile.ExpoCalcTypeCode' is invalid!",
				ErrMessage: fmt.Sprintf("'calcProfile.ExpoCalcTypeCode'\n"+
					"  contains an invalid value: %v", calcProfile.ExpoCalcTypeCode),
			}

	}

	return result, nil
}
