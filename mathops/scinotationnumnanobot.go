package mathops

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type scinotationnumNanobot struct {
	lock sync.Mutex
}

// getSciNotationStr
//
//	Returns a string containing the scientific notation display.  This
//	method differs from method SciNotationNum.GetNumStr() in that this
//	method requires the user to provide input parameter 'mantissaLen'
//	which sets the number of decimal places in the significand of the
//	returned science notation string.
//
//	Default Example
//	---------------
//	"2.652e+8"
func (sciNotNanobot *scinotationnumNanobot) getSciNotationStr(
	sciNotan *SciNotationNum,
	mantissaLen uint,
	errPrefDto *ePref.ErrPrefixDto) (string, error) {

	sciNotNanobot.lock.Lock()

	defer sciNotNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"scinotationnumNanobot.getSciNotationStr",
		"")

	if err != nil {
		return "", err
	}

	if sciNotan == nil {

		return "",
			&InputPtrNilError{
				ErrPrefix:     ePrefix.String(),
				ParameterName: "'sciNotan'",
			}
	}

	var outStr, tempStr string

	if mantissaLen == 0 {
		mantissaLen = 2
	}

	sciNotElectron := new(scinotationnumElectron)

	// sciNotan.SetMantissaLength(mantissaLen)
	sciNotElectron.setMantissaLength(sciNotan, mantissaLen)

	//sciNotan.SetDecimalSeparatorIfEmpty()
	sciNotElectron.setDecimalSeparatorIfEmpty(sciNotan)

	err = sciNotan.significand.SetDecimalSeparator(sciNotan.decimalSeparator)

	if err != nil {

		return "",
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = sciNotan.significand.\n" +
					"  SetDecimalSeparator(sciNotan.decimalSeparator)",
				ErrContext: fmt.Sprintf("sciNotan.decimalSeparator= '%c'",
					sciNotan.decimalSeparator),
				ErrMessage: err.Error(),
			}
	}

	significandPrecisionUint, err := sciNotan.significand.GetPrecisionUint()

	if err != nil {

		return "",
			&FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "significandPrecisionUint, err := \n" +
					"  sciNotan.significand.GetPrecisionUint()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if significandPrecisionUint != sciNotan.mantissaLength {

		bINumSignificand, err := sciNotan.significand.CopyOut()

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "bINum, err := sciNotan.significand.CopyOut()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		bINumSignificandNumStr, err := bINumSignificand.GetNumStr()

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix: ePrefix.String(),
					ReturnFunc: "bINumSignificandNumStr, err :=\n" +
						"  bINumSignificand.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

		err = bINumSignificand.SetPrecision(sciNotan.mantissaLength)

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "err = bINum.SetPrecision(sciNotan.mantissaLength)",
					ErrContext: fmt.Sprintf("sciNotan.mantissaLength= '%v'",
						sciNotan.mantissaLength),
					ErrMessage: err.Error(),
				}
		}

		tempStr, err = bINumSignificand.GetNumStr()

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "tempStr, err = bINumSignificand.GetNumStr()",
					ErrContext: fmt.Sprintf("bINumSignificand= '%v'", bINumSignificandNumStr),
					ErrMessage: err.Error(),
				}
		}

	} else {

		tempStr, err = sciNotan.significand.GetNumStr()

		if err != nil {

			return "",
				&FuncReturnError{
					ErrPrefix:  ePrefix.String(),
					ReturnFunc: "tempStr, err = sciNotan.significand.GetNumStr()",
					ErrContext: "",
					ErrMessage: err.Error(),
				}
		}

	}

	outStr += tempStr

	// outStr += sciNotan.GetExponentChar()
	outStr += string(sciNotan.exponentChar)

	exponentSignVal, err := sciNotan.exponent.GetSign()

	if err != nil {

		return "",
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "exponentSignVal, err := sciNotan.exponent.GetSign()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	if exponentSignVal == 1 &&
		sciNotan.exponentUsesLeadingPlus {

		outStr += "+"

	}

	tempStr, err = sciNotan.exponent.GetNumStr()

	if err != nil {

		return "",
			&FuncReturnError{
				ErrPrefix:  ePrefix.String(),
				ReturnFunc: "tempStr, err = sciNotan.exponent.GetNumStr()",
				ErrContext: "",
				ErrMessage: err.Error(),
			}
	}

	outStr += tempStr

	return outStr, nil
}

// new
//
//	Creates and returns an empty SciNotationNum structure. It is a
//	good idea to call this method in order to initialize default
//	settings.
func (sciNotNanobot *scinotationnumNanobot) new() SciNotationNum {

	sciNotNanobot.lock.Lock()

	defer sciNotNanobot.lock.Unlock()

	s2 := SciNotationNum{}

	sciElectron := new(scinotationnumElectron)

	//s2.SetExponentCharIfEmpty()
	sciElectron.setExponentCharIfEmpty(&s2)

	//s2.SetDecimalSeparatorIfEmpty()

	sciElectron.setDecimalSeparatorIfEmpty(&s2)

	s2.significand, _ = new(BigIntNum).NewZero(1)

	s2.exponent, _ = new(BigIntNum).NewZero(0)

	s2.exponentUsesLeadingPlus = true

	return s2
}

// setBigIntNumElements
//
//	Sets the components of the current SciNotationNum instance based
//	on two BigIntNum input parameters.
//
//	Input Parameters
//	================
//
//	significand              BigIntNum
//	  In the example scientific notation '2.652e+8',	the significand
//	  is represented by '2.652'.
//
//	exponent                 BigIntNum
//	  In the example '2.652e+8' the exponent component is represented
//	  by the integer value, '8'. Note: If exponent is NOT an integer
//	  value and contains fractional digits, an error will be
//	  triggered.
func (sciNotNanobot *scinotationnumNanobot) setBigIntNumElements(
	sciNotan *SciNotationNum,
	significand *BigIntNum,
	validateSignificand bool,
	exponent *BigIntNum,
	validateExponent bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	sciNotNanobot.lock.Lock()

	defer sciNotNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"scinotationnumNanobot.setBigIntNumElements",
		"")

	if err != nil {
		return err
	}

	if sciNotan == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'sciNotan'",
		}
	}

	if significand == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'significand'",
		}
	}

	if exponent == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'exponent'",
		}
	}

	if validateSignificand {

		err = significand.IsValid(ePrefix.XCpy("Validating 'significand'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = significand.IsValid(ePrefix.XCpy(\n" +
					"  \"Validating 'significand'\").String())",
				ErrContext: "Error: Input parameter 'significand' is invalid.\n" +
					"'significand' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	if validateExponent {

		err = exponent.IsValid(ePrefix.XCpy("Validating 'exponent'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = exponent.IsValid(ePrefix.XCpy(\n" +
					"  \"Validating 'exponent'\").String())",
				ErrContext: "Error: Input parameter 'exponent' is invalid.\n" +
					"'exponent' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	significandNumStr, err := significand.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: " significandNumStr, err := significand.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentNumStr, err := exponent.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentPrecisionUint, err := exponent.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentPrecisionUint, err := exponent.GetPrecisionUint()",
			ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
			ErrMessage: err.Error(),
		}
	}

	if exponentPrecisionUint > 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
			ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
				"'exponent' contains fractional digits.",
		}
	}

	//sciNotan.SetExponentCharIfEmpty()
	new(scinotationnumElectron).setExponentCharIfEmpty(sciNotan)

	err = sciNotan.significand.CopyIn(significand)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sciNotan.significand.CopyIn(significand)",
			ErrContext: fmt.Sprintf("significand= '%v'", significandNumStr),
			ErrMessage: err.Error(),
		}
	}

	significandDecimalSeparator, err := significand.GetDecimalSeparator()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "significandDecimalSeparator, err := \n" +
				"  significand.GetDecimalSeparator()",
			ErrContext: fmt.Sprintf("significand= '%v'", significandNumStr),
			ErrMessage: err.Error(),
		}
	}

	sciNotan.decimalSeparator = significandDecimalSeparator

	significandPrecisionUint, err := significand.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "significandPrecisionUint, err := \n" +
				"  significand.GetPrecisionUint()",
			ErrContext: fmt.Sprintf("significand= '%v'", significandNumStr),
			ErrMessage: err.Error(),
		}
	}

	sciNotan.mantissaLength = significandPrecisionUint

	err = sciNotan.exponent.CopyIn(exponent)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "sciNotan.exponent.CopyIn(exponent)",
			ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
			ErrMessage: err.Error(),
		}
	}

	return nil
}

// setIntAryElements
//
//	Sets the components of the current SciNotationNum instance based
//	on two IntAry input parameters.
//
//	Input Parameters
//	================
//
//	significand              IntAry
//	  In the example scientific notation '2.652e+8', the significand
//	  is represented by '2.652'.
//
//	exponent                 IntAry
//	  In the example '2.652e+8' the exponent component is represented
//	  by the integer value, '8'. Note: If exponent is NOT an integer
//	  value and contains fractional digits, an error will be
//	  returned.
func (sciNotNanobot *scinotationnumNanobot) setIntAryElements(
	sciNotan *SciNotationNum,
	significand *IntAry,
	validateSignificand bool,
	exponent *IntAry,
	validateExponent bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	sciNotNanobot.lock.Lock()

	defer sciNotNanobot.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"scinotationnumNanobot.setIntAryElements",
		"")

	if err != nil {
		return err
	}

	if sciNotan == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'sciNotan'",
		}
	}

	if significand == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'significand'",
		}
	}

	if exponent == nil {

		return &InputPtrNilError{
			ErrPrefix:     ePrefix.String(),
			ErrContext:    "",
			ParameterName: "'exponent'",
		}
	}

	if validateSignificand {

		err = significand.IsValid(ePrefix.XCpy("Validating 'significand'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = significand.IsValid(ePrefix.XCpy(\n" +
					"  \"Validating 'significand'\").String())",
				ErrContext: "Error: Input parameter 'significand' is invalid.\n" +
					"'significand' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	if validateExponent {

		err = exponent.IsValid(ePrefix.XCpy("Validating 'exponent'").String())

		if err != nil {

			return &FuncReturnError{
				ErrPrefix: ePrefix.String(),
				ReturnFunc: "err = exponent.IsValid(ePrefix.XCpy(\n" +
					"  \"Validating 'exponent'\").String())",
				ErrContext: "Error: Input parameter 'exponent' is invalid.\n" +
					"'exponent' FAILED validation tests.",
				ErrMessage: err.Error(),
			}
		}
	}

	significandNumStr, err := significand.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: " significandNumStr, err := significand.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentNumStr, err := exponent.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentNumStr, err := exponent.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	exponentPrecisionUint, err := exponent.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentPrecisionUint, err := exponent.GetPrecisionUint()",
			ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
			ErrMessage: err.Error(),
		}
	}

	if exponentPrecisionUint > 0 {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "",
			ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
			ErrMessage: "Error: Input parameter 'exponent' is INVALID!\n" +
				"'exponent' contains fractional digits.",
		}
	}

	// sciNotan.SetExponentCharIfEmpty()
	new(scinotationnumElectron).setExponentCharIfEmpty(sciNotan)

	biNumSignificand, err := significand.GetBigIntNum()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "biNumSignificand, err := significand.GetBigIntNum()",
			ErrContext: fmt.Sprintf("significand= '%v'", significandNumStr),
			ErrMessage: err.Error(),
		}
	}

	significandNumStr, err = biNumSignificand.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "significandNumStr, err = biNumSignificand.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	//sciNotan.significand = biNumSignificand.CopyOut()

	err = sciNotan.significand.CopyIn(&biNumSignificand)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sciNotan.significand.CopyIn(&biNumSignificand)",
			ErrContext: fmt.Sprintf("biNumSignificand= '%v'", significandNumStr),
			ErrMessage: err.Error(),
		}
	}

	sciNotan.decimalSeparator = significand.GetDecimalSeparator()

	significandPrecisionUint, err := sciNotan.significand.GetPrecisionUint()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix: ePrefix.String(),
			ReturnFunc: "significandPrecisionUint, err := \n" +
				"  sciNotan.significand.GetPrecisionUint()",
			ErrContext: fmt.Sprintf("sciNotan.significand= '%v'", significandNumStr),
			ErrMessage: err.Error(),
		}
	}

	sciNotan.mantissaLength = significandPrecisionUint

	biNumExponent, err := exponent.GetBigIntNum()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "biNumExponent, err := exponent.GetBigIntNum()",
			ErrContext: fmt.Sprintf("exponent= '%v'", exponentNumStr),
			ErrMessage: err.Error(),
		}
	}

	exponentNumStr, err = biNumExponent.GetNumStr()

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "exponentNumStr, err = biNumExponent.GetNumStr()",
			ErrContext: "",
			ErrMessage: err.Error(),
		}
	}

	err = sciNotan.exponent.CopyIn(&biNumExponent)

	if err != nil {

		return &FuncReturnError{
			ErrPrefix:  ePrefix.String(),
			ReturnFunc: "err = sciNotan.exponent.CopyIn(&biNumExponent)",
			ErrContext: fmt.Sprintf("biNumExponent= '%v'", exponentNumStr),
			ErrMessage: err.Error(),
		}
	}

	return nil
}
